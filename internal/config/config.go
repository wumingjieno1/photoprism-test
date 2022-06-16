package config

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/dustin/go-humanize"
	"github.com/klauspost/cpuid/v2"
	"github.com/pbnjay/memory"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/face"
	"github.com/photoprism/photoprism/internal/hub"
	"github.com/photoprism/photoprism/internal/hub/places"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/thumb"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
	"github.com/photoprism/photoprism/pkg/rnd"
)

var log = event.Log
var once sync.Once
var LowMem = false
var TotalMem uint64

// Config holds database, cache and all parameters of photoprism
type Config struct {
	once     sync.Once
	db       *gorm.DB
	options  *Options
	settings *Settings
	hub      *hub.Config
	token    string
	serial   string
	env      string
}

func init() {
	TotalMem = memory.TotalMemory()

	// Check available memory if not running in unsafe mode.
	if Env(EnvUnsafe) {
		// Disable features with high memory requirements?
		LowMem = TotalMem < MinMem
	}

	// Init public thumb sizes for use in client apps.
	for i := len(thumb.DefaultSizes) - 1; i >= 0; i-- {
		name := thumb.DefaultSizes[i]
		t := thumb.Sizes[name]

		if t.Public {
			Thumbs = append(Thumbs, ThumbSize{Size: string(name), Use: t.Use, Width: t.Width, Height: t.Height})
		}
	}
}

func initLogger() {
	once.Do(func() {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})

		if Env(EnvTrace) {
			log.SetLevel(logrus.TraceLevel)
		} else if Env(EnvDebug) {
			log.SetLevel(logrus.DebugLevel)
		} else {
			log.SetLevel(logrus.InfoLevel)
		}
	})
}

// NewConfig initialises a new configuration file
func NewConfig(ctx *cli.Context) *Config {
	// Initialize logger.
	initLogger()

	// Initialize options from config file and CLI context.
	c := &Config{
		options: NewOptions(ctx),
		token:   rnd.GenerateToken(8),
		env:     os.Getenv("DOCKER_ENV"),
	}

	// Overwrite values with options.yml from config path.
	if optionsYaml := c.OptionsYaml(); fs.FileExists(optionsYaml) {
		if err := c.options.Load(optionsYaml); err != nil {
			log.Warnf("config: failed loading values from %s (%s)", clean.Log(optionsYaml), err)
		} else {
			log.Debugf("config: overriding config with values from %s", clean.Log(optionsYaml))
		}
	}

	return c
}

// Unsafe checks if unsafe settings are allowed.
func (c *Config) Unsafe() bool {
	return c.options.Unsafe
}

// Options returns the raw config options.
func (c *Config) Options() *Options {
	if c.options == nil {
		log.Warnf("config: options should not be nil - possible bug")
		c.options = NewOptions(nil)
	}

	return c.options
}

// Propagate updates config options in other packages as needed.
func (c *Config) Propagate() {
	log.SetLevel(c.LogLevel())

	// Set thumbnail generation parameters.
	thumb.StandardRGB = c.ThumbSRGB()
	thumb.SizePrecached = c.ThumbSizePrecached()
	thumb.SizeUncached = c.ThumbSizeUncached()
	thumb.Filter = c.ThumbFilter()
	thumb.JpegQuality = c.JpegQuality()

	// Set geocoding parameters.
	places.UserAgent = c.UserAgent()
	entity.GeoApi = c.GeoApi()

	// Set facial recognition parameters.
	face.ScoreThreshold = c.FaceScore()
	face.OverlapThreshold = c.FaceOverlap()
	face.ClusterScoreThreshold = c.FaceClusterScore()
	face.ClusterSizeThreshold = c.FaceClusterSize()
	face.ClusterCore = c.FaceClusterCore()
	face.ClusterDist = c.FaceClusterDist()
	face.MatchDist = c.FaceMatchDist()

	c.Settings().Propagate()
	c.Hub().Propagate()
}

// Init creates directories, parses additional config files, opens a database connection and initializes dependencies.
func (c *Config) Init() error {
	start := time.Now()

	if err := c.CreateDirectories(); err != nil {
		return err
	}

	if err := c.initSerial(); err != nil {
		return err
	}

	// Show funding info?
	if !c.Sponsor() {
		log.Info(MsgSponsor)
		log.Info(MsgSignUp)
	}

	if insensitive, err := c.CaseInsensitive(); err != nil {
		return err
	} else if insensitive {
		log.Infof("config: case-insensitive file system detected")
		fs.IgnoreCase()
	}

	if cpuName := cpuid.CPU.BrandName; cpuName != "" {
		log.Debugf("config: running on %s, %s memory detected", clean.Log(cpuid.CPU.BrandName), humanize.Bytes(TotalMem))
	}

	// Exit if less than 128 MB RAM was detected.
	if TotalMem < 128*Megabyte {
		return fmt.Errorf("config: %s of memory detected, %d GB required", humanize.Bytes(TotalMem), MinMem/Gigabyte)
	}

	// Show warning if less than 1 GB RAM was detected.
	if LowMem {
		log.Warnf(`config: less than %d GB of memory detected, please upgrade if server becomes unstable or unresponsive`, MinMem/Gigabyte)
		log.Warnf("config: tensorflow as well as indexing and conversion of RAW files have been disabled automatically")
	}

	// Show swap info.
	if TotalMem < RecommendedMem {
		log.Infof("config: make sure your server has enough swap configured to prevent restarts when there are memory usage spikes")
	}

	// Show wakeup interval warning if face recognition is enabled
	// and the worker runs less than once per hour.
	if !c.DisableFaces() && !c.Unsafe() && c.WakeupInterval() > time.Hour {
		log.Warnf("config: the wakeup interval is %s, but must be 1h or less for face recognition to work", c.WakeupInterval().String())
	}

	// Set HTTP user agent.
	places.UserAgent = c.UserAgent()

	c.initSettings()
	c.initHub()

	c.Propagate()

	err := c.connectDb()

	if err == nil {
		log.Debugf("config: successfully initialized [%s]", time.Since(start))
	}

	return err
}

// readSerial reads and returns the current storage serial.
func (c *Config) readSerial() string {
	storageName := filepath.Join(c.StoragePath(), serialName)
	backupName := filepath.Join(c.BackupPath(), serialName)

	if fs.FileExists(storageName) {
		if data, err := os.ReadFile(storageName); err == nil && len(data) == 16 {
			return string(data)
		} else {
			log.Tracef("config: could not read %s (%s)", clean.Log(storageName), err)
		}
	}

	if fs.FileExists(backupName) {
		if data, err := os.ReadFile(backupName); err == nil && len(data) == 16 {
			return string(data)
		} else {
			log.Tracef("config: could not read %s (%s)", clean.Log(backupName), err)
		}
	}

	return ""
}

// initSerial initializes storage directories with a random serial.
func (c *Config) initSerial() (err error) {
	if c.Serial() != "" {
		return nil
	}

	c.serial = rnd.GenerateUID('z')

	storageName := filepath.Join(c.StoragePath(), serialName)
	backupName := filepath.Join(c.BackupPath(), serialName)

	if err = os.WriteFile(storageName, []byte(c.serial), os.ModePerm); err != nil {
		return fmt.Errorf("could not create %s: %s", storageName, err)
	}

	if err = os.WriteFile(backupName, []byte(c.serial), os.ModePerm); err != nil {
		return fmt.Errorf("could not create %s: %s", backupName, err)
	}

	return nil
}

// Serial returns the random storage serial.
func (c *Config) Serial() string {
	if c.serial == "" {
		c.serial = c.readSerial()
	}

	return c.serial
}

// SerialChecksum returns the CRC32 checksum of the storage serial.
func (c *Config) SerialChecksum() string {
	var result []byte

	hash := crc32.New(crc32.MakeTable(crc32.Castagnoli))

	if _, err := hash.Write([]byte(c.Serial())); err != nil {
		log.Warnf("config: %s", err)
	}

	return hex.EncodeToString(hash.Sum(result))
}

// Name returns the app name.
func (c *Config) Name() string {
	if c.options.Name == "" {
		return "PhotoPrism"
	}

	return c.options.Name
}

// Edition returns the app edition.
func (c *Config) Edition() string {
	if c.options.Edition == "" {
		return "PhotoPrism® Dev"
	} else if strings.HasSuffix(c.options.Edition, "CE") && c.Sponsor() {
		return strings.Replace(c.options.Edition, "CE", "Plus", 1)
	}

	return c.options.Edition
}

// Version returns the application version.
func (c *Config) Version() string {
	return c.options.Version
}

// UserAgent returns an HTTP user agent string based on the app config and version.
func (c *Config) UserAgent() string {
	return fmt.Sprintf("%s/%s (%s)", c.Name(), c.Version(), strings.Join(append(c.Flags(), c.Serial()), "; "))
}

// Copyright returns the application copyright.
func (c *Config) Copyright() string {
	return c.options.Copyright
}

// BaseUri returns the site base URI for a given resource.
func (c *Config) BaseUri(res string) string {
	if c.SiteUrl() == "" {
		return res
	}

	u, err := url.Parse(c.SiteUrl())

	if err != nil {
		return res
	}

	return strings.TrimRight(u.EscapedPath(), "/") + res
}

// ApiUri returns the api URI.
func (c *Config) ApiUri() string {
	return c.BaseUri(ApiUri)
}

// CdnUrl returns the optional content delivery network URI without trailing slash.
func (c *Config) CdnUrl(res string) string {
	if c.NoSponsor() {
		return res
	}

	return strings.TrimRight(c.options.CdnUrl, "/") + res
}

// ContentUri returns the content delivery URI.
func (c *Config) ContentUri() string {
	return c.CdnUrl(c.ApiUri())
}

// StaticUri returns the static content URI.
func (c *Config) StaticUri() string {
	return c.CdnUrl(c.BaseUri(StaticUri))
}

// SiteUrl returns the public server URL (default is "http://localhost:2342/").
func (c *Config) SiteUrl() string {
	if c.options.SiteUrl == "" {
		return "http://localhost:2342/"
	}

	return strings.TrimRight(c.options.SiteUrl, "/") + "/"
}

// SiteDomain returns the public server domain.
func (c *Config) SiteDomain() string {
	if u, err := url.Parse(c.SiteUrl()); err != nil {
		return "localhost"
	} else {
		return u.Hostname()
	}
}

// SiteAuthor returns the site author / copyright.
func (c *Config) SiteAuthor() string {
	return c.options.SiteAuthor
}

// SiteTitle returns the main site title (default is application name).
func (c *Config) SiteTitle() string {
	if c.options.SiteTitle == "" || c.NoSponsor() {
		return c.Name()
	}

	return c.options.SiteTitle
}

// SiteCaption returns a short site caption.
func (c *Config) SiteCaption() string {
	return c.options.SiteCaption
}

// SiteDescription returns a long site description.
func (c *Config) SiteDescription() string {
	return c.options.SiteDescription
}

// SitePreview returns the site preview image URL for sharing.
func (c *Config) SitePreview() string {
	if c.options.SitePreview == "" {
		return c.SiteUrl() + "static/img/preview.jpg"
	}

	if !strings.HasPrefix(c.options.SitePreview, "http") {
		return c.SiteUrl() + c.options.SitePreview
	}

	return c.options.SitePreview
}

// Imprint returns the legal info text for the page footer.
func (c *Config) Imprint() string {
	if c.NoSponsor() {
		return MsgSponsor
	}

	return c.options.Imprint
}

// ImprintUrl returns the legal info url.
func (c *Config) ImprintUrl() string {
	if c.NoSponsor() {
		return SignUpURL
	}

	return c.options.ImprintUrl
}

// Debug checks if debug mode is enabled, shows non-essential log messages.
func (c *Config) Debug() bool {
	if c.Trace() {
		return true
	}

	return c.options.Debug
}

// Trace checks if trace mode is enabled, shows all log messages.
func (c *Config) Trace() bool {
	return c.options.Trace
}

// Test checks if test mode is enabled.
func (c *Config) Test() bool {
	return c.options.Test
}

// Demo checks if demo mode is enabled.
func (c *Config) Demo() bool {
	return c.options.Demo
}

// Sponsor reports if your continuous support helps to pay for development and operating expenses.
func (c *Config) Sponsor() bool {
	return c.options.Sponsor || c.Test()
}

// NoSponsor reports if the instance is not operated by a sponsor.
func (c *Config) NoSponsor() bool {
	return !c.Sponsor() && !c.Demo()
}

// Public checks if app runs in public mode and requires no authentication.
func (c *Config) Public() bool {
	if c.Auth() {
		return false
	} else if c.Demo() {
		return true
	}

	return c.options.Public
}

// SetPublic changes authentication while instance is running, for testing purposes only.
func (c *Config) SetPublic(p bool) {
	if c.Debug() {
		c.options.Public = p
	}
}

// Experimental checks if experimental features should be enabled.
func (c *Config) Experimental() bool {
	return c.options.Experimental
}

// ReadOnly checks if photo directories are write protected.
func (c *Config) ReadOnly() bool {
	return c.options.ReadOnly
}

// DetectNSFW checks if NSFW photos should be detected and flagged.
func (c *Config) DetectNSFW() bool {
	return c.options.DetectNSFW
}

// UploadNSFW checks if NSFW photos can be uploaded.
func (c *Config) UploadNSFW() bool {
	return c.options.UploadNSFW
}

// AdminPassword returns the initial admin password.
func (c *Config) AdminPassword() string {
	return c.options.AdminPassword
}

// Auth checks if authentication is always required.
func (c *Config) Auth() bool {
	return c.options.Auth
}

// LogLevel returns the Logrus log level.
func (c *Config) LogLevel() logrus.Level {
	// Normalize string.
	c.options.LogLevel = strings.ToLower(strings.TrimSpace(c.options.LogLevel))

	if c.Trace() {
		c.options.LogLevel = logrus.TraceLevel.String()
	} else if c.Debug() && c.options.LogLevel != logrus.TraceLevel.String() {
		c.options.LogLevel = logrus.DebugLevel.String()
	}

	if logLevel, err := logrus.ParseLevel(c.options.LogLevel); err == nil {
		return logLevel
	} else {
		return logrus.InfoLevel
	}
}

// SetLogLevel sets the Logrus log level.
func (c *Config) SetLogLevel(level logrus.Level) {
	log.SetLevel(level)
}

// Shutdown services and workers.
func (c *Config) Shutdown() {
	mutex.People.Cancel()
	mutex.MainWorker.Cancel()
	mutex.ShareWorker.Cancel()
	mutex.SyncWorker.Cancel()
	mutex.MetaWorker.Cancel()
	mutex.FacesWorker.Cancel()

	if err := c.CloseDb(); err != nil {
		log.Errorf("could not close database connection: %s", err)
	} else {
		log.Info("closed database connection")
	}
}

// Workers returns the number of workers e.g. for indexing files.
func (c *Config) Workers() int {
	// Use one worker on systems with less than the recommended amount of memory.
	if TotalMem < RecommendedMem {
		return 1
	}

	// NumCPU returns the number of logical CPU cores.
	cores := runtime.NumCPU()

	// Limit to physical cores to avoid high load on HT capable CPUs.
	if cores > cpuid.CPU.PhysicalCores {
		cores = cpuid.CPU.PhysicalCores
	}

	// Limit number of workers when using SQLite3 to avoid database locking issues.
	if c.DatabaseDriver() == SQLite3 && (cores >= 8 && c.options.Workers <= 0 || c.options.Workers > 4) {
		return 4
	}

	// Return explicit value if set and not too large.
	if c.options.Workers > runtime.NumCPU() {
		return runtime.NumCPU()
	} else if c.options.Workers > 0 {
		return c.options.Workers
	}

	// Use half the available cores by default.
	if cores > 1 {
		return cores / 2
	}

	return 1
}

// WakeupInterval returns the duration between background worker runs
// required for face recognition and index maintenance(1-86400s).
func (c *Config) WakeupInterval() time.Duration {
	if c.options.WakeupInterval <= 0 {
		if c.options.Unsafe {
			// Worker can be disabled only in unsafe mode.
			return time.Duration(0)
		} else {
			// Default to 15 minutes if no interval is set.
			return DefaultWakeupInterval
		}
	}

	// Do not run more than once per minute.
	if c.options.WakeupInterval < MinWakeupInterval/time.Second {
		return MinWakeupInterval
	} else if c.options.WakeupInterval < MinWakeupInterval {
		c.options.WakeupInterval = c.options.WakeupInterval * time.Second
	}

	// Do not run less than once per day.
	if c.options.WakeupInterval > MaxWakeupInterval {
		return MaxWakeupInterval
	}

	return c.options.WakeupInterval
}

// AutoIndex returns the auto index delay duration.
func (c *Config) AutoIndex() time.Duration {
	if c.options.AutoIndex < 0 {
		return time.Duration(0)
	} else if c.options.AutoIndex == 0 || c.options.AutoIndex > 604800 {
		return time.Duration(DefaultAutoIndexDelay) * time.Second
	}

	return time.Duration(c.options.AutoIndex) * time.Second
}

// AutoImport returns the auto import delay duration.
func (c *Config) AutoImport() time.Duration {
	if c.options.AutoImport < 0 || c.ReadOnly() {
		return time.Duration(0)
	} else if c.options.AutoImport == 0 || c.options.AutoImport > 604800 {
		return time.Duration(DefaultAutoImportDelay) * time.Second
	}

	return time.Duration(c.options.AutoImport) * time.Second
}

// GeoApi returns the preferred geocoding api (none or places).
func (c *Config) GeoApi() string {
	if c.options.DisablePlaces {
		return ""
	}

	return "places"
}

// OriginalsLimit returns the maximum size of originals in MB.
func (c *Config) OriginalsLimit() int {
	if c.options.OriginalsLimit <= 0 || c.options.OriginalsLimit > 100000 {
		return -1
	}

	return c.options.OriginalsLimit
}

// OriginalsLimitBytes returns the maximum size of originals in bytes.
func (c *Config) OriginalsLimitBytes() int64 {
	if result := c.OriginalsLimit(); result <= 0 {
		return -1
	} else {
		return int64(result) * 1024 * 1024
	}
}

// ResolutionLimit returns the maximum resolution of originals in megapixels (width x height).
func (c *Config) ResolutionLimit() int {
	result := c.options.ResolutionLimit

	if result <= 0 {
		return -1
	} else if result > 900 {
		result = 900
	}

	return result
}

// UpdateHub updates backend api credentials for maps & places.
func (c *Config) UpdateHub() {
	if err := c.hub.Refresh(); err != nil {
		log.Debugf("config: %s", err)
	} else if err := c.hub.Save(); err != nil {
		log.Debugf("config: %s", err)
	} else {
		c.hub.Propagate()
	}
}

// initHub initializes PhotoPrism hub config.
func (c *Config) initHub() {
	if c.hub != nil {
		return
	}

	c.hub = hub.NewConfig(c.Version(), c.HubConfigFile(), c.serial, c.env, c.UserAgent(), c.options.PartnerID)

	if err := c.hub.Load(); err == nil {
		// Do nothing.
	} else if err := c.hub.Refresh(); err != nil {
		log.Debugf("config: %s", err)
	} else if err := c.hub.Save(); err != nil {
		log.Debugf("config: %s", err)
	}

	c.hub.Propagate()

	ticker := time.NewTicker(time.Hour * 24)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.UpdateHub()
			}
		}
	}()
}

// Hub returns the PhotoPrism hub config.
func (c *Config) Hub() *hub.Config {
	c.initHub()

	return c.hub
}
