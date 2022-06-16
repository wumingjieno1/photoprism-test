package config

import (
	"os"

	"github.com/photoprism/photoprism/pkg/fs"
)

// RawEnabled checks if indexing and conversion of RAW files is enabled.
func (c *Config) RawEnabled() bool {
	return !c.DisableRaw()
}

// RawPresets checks if RAW converter presents should be used (may reduce performance).
func (c *Config) RawPresets() bool {
	return c.options.RawPresets
}

// DarktableBin returns the darktable-cli executable file name.
func (c *Config) DarktableBin() string {
	return findExecutable(c.options.DarktableBin, "darktable-cli")
}

// DarktableBlacklist returns the darktable file extension blacklist.
func (c *Config) DarktableBlacklist() string {
	return c.options.DarktableBlacklist
}

// DarktableConfigPath returns the darktable config directory.
func (c *Config) DarktableConfigPath() string {
	return fs.Abs(c.options.DarktableConfigPath)
}

// DarktableCachePath returns the darktable cache directory.
func (c *Config) DarktableCachePath() string {
	return fs.Abs(c.options.DarktableCachePath)
}

// CreateDarktableCachePath creates and returns the darktable cache directory.
func (c *Config) CreateDarktableCachePath() (string, error) {
	cachePath := c.DarktableCachePath()

	if cachePath == "" {
		return "", nil
	} else if err := os.MkdirAll(cachePath, os.ModePerm); err != nil {
		return cachePath, err
	} else {
		c.options.DarktableCachePath = cachePath
	}

	return cachePath, nil
}

// CreateDarktableConfigPath creates and returns the darktable config directory.
func (c *Config) CreateDarktableConfigPath() (string, error) {
	configPath := c.DarktableConfigPath()

	if configPath == "" {
		return "", nil
	} else if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
		return configPath, err
	} else {
		c.options.DarktableConfigPath = configPath
	}

	return configPath, nil
}

// DarktableEnabled checks if Darktable is enabled for RAW conversion.
func (c *Config) DarktableEnabled() bool {
	return !c.DisableDarktable()
}

// RawtherapeeBin returns the rawtherapee-cli executable file name.
func (c *Config) RawtherapeeBin() string {
	return findExecutable(c.options.RawtherapeeBin, "rawtherapee-cli")
}

// RawtherapeeBlacklist returns the RawTherapee file extension blacklist.
func (c *Config) RawtherapeeBlacklist() string {
	return c.options.RawtherapeeBlacklist
}

// RawtherapeeEnabled checks if Rawtherapee is enabled for RAW conversion.
func (c *Config) RawtherapeeEnabled() bool {
	return !c.DisableRawtherapee()
}

// SipsEnabled checks if SIPS is enabled for RAW conversion.
func (c *Config) SipsEnabled() bool {
	return !c.DisableSips()
}

// SipsBin returns the SIPS executable file name.
func (c *Config) SipsBin() string {
	return findExecutable(c.options.SipsBin, "sips")
}

// HeifConvertBin returns the heif-convert executable file name.
func (c *Config) HeifConvertBin() string {
	return findExecutable(c.options.HeifConvertBin, "heif-convert")
}

// HeifConvertEnabled checks if heif-convert is enabled for HEIF conversion.
func (c *Config) HeifConvertEnabled() bool {
	return !c.DisableHeifConvert()
}
