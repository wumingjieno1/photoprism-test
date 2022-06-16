package hub

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/photoprism/photoprism/internal/hub/places"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

// Config represents backend api credentials for maps & geodata.
type Config struct {
	Version   string `json:"version" yaml:"Version"`
	FileName  string `json:"-" yaml:"-"`
	Key       string `json:"key" yaml:"Key"`
	Secret    string `json:"secret" yaml:"Secret"`
	Session   string `json:"session" yaml:"Session"`
	Status    string `json:"status" yaml:"Status"`
	Serial    string `json:"serial" yaml:"Serial"`
	Env       string `json:"-" yaml:"-"`
	UserAgent string `json:"-" yaml:"-"`
	PartnerID string `json:"-" yaml:"-"`
}

// NewConfig creates a new backend api credentials instance.
func NewConfig(version, fileName, serial, env, userAgent, partnerId string) *Config {
	return &Config{
		Version:   version,
		FileName:  fileName,
		Key:       "",
		Secret:    "",
		Session:   "",
		Status:    "",
		Serial:    serial,
		Env:       env,
		UserAgent: userAgent,
		PartnerID: partnerId,
	}
}

// MapKey returns the maps api key.
func (c *Config) MapKey() string {
	if sess, err := c.DecodeSession(); err != nil {
		return ""
	} else {
		return sess.MapKey
	}
}

// Propagate updates backend api credentials in other packages.
func (c *Config) Propagate() {
	places.Key = c.Key
	places.Secret = c.Secret
}

// Sanitize verifies and sanitizes backend api credentials.
func (c *Config) Sanitize() {
	c.Key = strings.ToLower(c.Key)

	if c.Secret != "" {
		if c.Key != fmt.Sprintf("%x", sha1.Sum([]byte(c.Secret))) {
			c.Key = ""
			c.Secret = ""
			c.Session = ""
			c.Status = ""
		}
	}
}

// DecodeSession decodes backend api session data.
func (c *Config) DecodeSession() (Session, error) {
	c.Sanitize()

	result := Session{}

	if c.Session == "" {
		return result, fmt.Errorf("empty session")
	}

	s, err := hex.DecodeString(c.Session)

	if err != nil {
		return result, err
	}

	hash := sha256.New()
	if _, err := hash.Write([]byte(c.Secret)); err != nil {
		return result, err
	}

	var b []byte

	block, err := aes.NewCipher(hash.Sum(b))

	if err != nil {
		return result, err
	}

	iv := s[:aes.BlockSize]

	plaintext := make([]byte, len(s))

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, s[aes.BlockSize:])

	plaintext = bytes.Trim(plaintext, "\x00")

	if err := json.Unmarshal(plaintext, &result); err != nil {
		return result, err
	}

	return result, nil
}

// Refresh updates backend api credentials.
func (c *Config) Refresh() (err error) {
	mutex.Lock()
	defer mutex.Unlock()

	if err := os.MkdirAll(filepath.Dir(c.FileName), os.ModePerm); err != nil {
		return err
	}

	c.Sanitize()

	// Create new http.Client instance.
	//
	// NOTE: Timeout specifies a time limit for requests made by
	// this Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	client := &http.Client{Timeout: 60 * time.Second}

	url := ServiceURL
	method := http.MethodPost

	var req *http.Request

	if c.Key != "" {
		url = fmt.Sprintf(ServiceURL+"/%s", c.Key)
		method = http.MethodPut
		log.Debugf("config: requesting updated api key for maps and places")
	} else {
		log.Debugf("config: requesting new api key for maps and places")
	}

	// Create request.
	if j, err := json.Marshal(NewRequest(c.Version, c.Serial, c.Env, c.PartnerID)); err != nil {
		return err
	} else if req, err = http.NewRequest(method, url, bytes.NewReader(j)); err != nil {
		return err
	}

	// Set user agent.
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	} else {
		req.Header.Set("User-Agent", "PhotoPrism/Test")
	}

	// Add Content-Type header.
	req.Header.Add("Content-Type", "application/json")

	var r *http.Response

	// Send request.
	for i := 0; i < 3; i++ {
		r, err = client.Do(req)

		if err == nil {
			break
		}
	}

	// Ok?
	if err != nil {
		return err
	} else if r.StatusCode >= 400 {
		err = fmt.Errorf("fetching api key from %s failed (error %d)", ApiHost(), r.StatusCode)
		return err
	}

	// Decode JSON response.
	err = json.NewDecoder(r.Body).Decode(c)

	if err != nil {
		return err
	}

	return nil
}

// Load backend api credentials from a YAML file.
func (c *Config) Load() error {
	if !fs.FileExists(c.FileName) {
		return fmt.Errorf("settings file not found: %s", clean.Log(c.FileName))
	}

	mutex.Lock()
	defer mutex.Unlock()

	yamlConfig, err := os.ReadFile(c.FileName)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlConfig, c); err != nil {
		return err
	}

	c.Sanitize()
	c.Propagate()

	if sess, err := c.DecodeSession(); err != nil {
		return err
	} else if sess.Expired() {
		return errors.New("session expired")
	}

	return nil
}

// Save backend api credentials to a YAML file.
func (c *Config) Save() error {
	mutex.Lock()
	defer mutex.Unlock()

	c.Sanitize()

	data, err := yaml.Marshal(c)

	if err != nil {
		return err
	}

	c.Propagate()

	if err := os.MkdirAll(filepath.Dir(c.FileName), os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(c.FileName, data, os.ModePerm); err != nil {
		return err
	}

	c.Propagate()

	return nil
}
