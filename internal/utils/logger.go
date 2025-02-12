package utils

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ClientID         string `mapstructure:"client_id"`
	ClientSecret     string `mapstructure:"client_secret"`
	AccessToken      string `mapstructure:"access_token"`
	DefaultProjectID string `mapstructure:"default_project_id"`
}

func getConfigPath() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic("failed to get user home directory")
		}
		xdgConfigHome = filepath.Join(home, ".config")
	}
	return filepath.Join(xdgConfigHome, "tickli", "config.yaml")
}

func Load() (*Config, error) {
	configPath := getConfigPath()
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return &Config{}, nil // No config file found, return empty config
		}
		return nil, errors.Wrap(err, "reading config file")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "parsing config file")
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	configPath := getConfigPath()
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrap(err, "creating config directory")
	}

	viper.Set("client_id", cfg.ClientID)
	viper.Set("client_secret", cfg.ClientSecret)
	viper.Set("access_token", cfg.AccessToken)
	viper.Set("default_project_id", cfg.DefaultProjectID)

	if err := viper.WriteConfigAs(configPath); err != nil {
		return errors.Wrap(err, "writing config file")
	}

	return nil
}
