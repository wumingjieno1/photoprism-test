package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/i18n"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

// Settings represents user settings for Web UI, indexing, and import.
type Settings struct {
	UI        UISettings       `json:"ui" yaml:"UI"`
	Search    SearchSettings   `json:"search" yaml:"Search"`
	Maps      MapsSettings     `json:"maps" yaml:"Maps"`
	Features  FeatureSettings  `json:"features" yaml:"Features"`
	Import    ImportSettings   `json:"import" yaml:"Import"`
	Index     IndexSettings    `json:"index" yaml:"Index"`
	Stack     StackSettings    `json:"stack" yaml:"Stack"`
	Share     ShareSettings    `json:"share" yaml:"Share"`
	Download  DownloadSettings `json:"download" yaml:"Download"`
	Templates TemplateSettings `json:"templates" yaml:"Templates"`
}

// NewSettings creates a new Settings instance.
func NewSettings(c *Config) *Settings {
	return &Settings{
		UI: UISettings{
			Scrollbar: true,
			Zoom:      false,
			Theme:     c.DefaultTheme(),
			Language:  c.DefaultLocale(),
		},
		Search: SearchSettings{
			BatchSize: 0,
		},
		Maps: MapsSettings{
			Animate: 0,
			Style:   "streets",
		},
		Features: FeatureSettings{
			Upload:    true,
			Download:  true,
			Archive:   true,
			Review:    true,
			Private:   true,
			Files:     true,
			Videos:    true,
			Folders:   true,
			Albums:    true,
			Moments:   true,
			Estimates: true,
			People:    true,
			Labels:    true,
			Places:    true,
			Edit:      true,
			Share:     true,
			Library:   true,
			Import:    true,
			Logs:      true,
		},
		Import: ImportSettings{
			Path: entity.RootPath,
			Move: false,
		},
		Index: IndexSettings{
			Path:    entity.RootPath,
			Rescan:  false,
			Convert: true,
		},
		Stack: StackSettings{
			UUID: true,
			Meta: true,
			Name: false,
		},
		Share: ShareSettings{
			Title: "",
		},
		Download: NewDownloadSettings(),
		Templates: TemplateSettings{
			Default: "index.tmpl",
		},
	}
}

// Propagate updates settings in other packages as needed.
func (s *Settings) Propagate() {
	i18n.SetLocale(s.UI.Language)
}

// StackSequences checks if files should be stacked based on their file name prefix (sequential names).
func (s Settings) StackSequences() bool {
	return s.Stack.Name
}

// StackUUID checks if files should be stacked based on unique image or instance id.
func (s Settings) StackUUID() bool {
	return s.Stack.UUID
}

// StackMeta checks if files should be stacked based on their place and time metadata.
func (s Settings) StackMeta() bool {
	return s.Stack.Meta
}

// Load user settings from file.
func (s *Settings) Load(fileName string) error {
	if !fs.FileExists(fileName) {
		return fmt.Errorf("settings file not found: %s", clean.Log(fileName))
	}

	yamlConfig, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlConfig, s); err != nil {
		return err
	}

	s.Propagate()

	return nil
}

// Save user settings to a file.
func (s *Settings) Save(fileName string) error {
	data, err := yaml.Marshal(s)

	if err != nil {
		return err
	}

	s.Propagate()

	if err := os.WriteFile(fileName, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// initSettings initializes user settings from a config file.
func (c *Config) initSettings() {
	if c.settings != nil {
		return
	}

	c.settings = NewSettings(c)
	fileName := c.SettingsYaml()

	if err := c.settings.Load(fileName); err == nil {
		log.Debugf("settings: loaded from %s ", fileName)
	} else if err := c.settings.Save(fileName); err != nil {
		log.Errorf("settings: could not create %s (%s)", fileName, err)
	} else {
		log.Debugf("settings: saved to %s ", fileName)
	}

	i18n.SetDir(c.LocalesPath())

	c.settings.Propagate()
}

// Settings returns the current user settings.
func (c *Config) Settings() *Settings {
	c.initSettings()

	if c.DisablePlaces() {
		c.settings.Features.Places = false
	}

	return c.settings
}
