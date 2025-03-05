package config

import (
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Config struct {
	DefaultProjectID    string `mapstructure:"default_project_id"`
	DefaultProjectColor string `mapstructure:"default_project_color"`
}

var (
	configPath = filepath.Join(xdg.ConfigHome, "tickli", "config.yaml")
	tokenPath  = filepath.Join(xdg.DataHome, "tickli", "token")
)

func InitConfig() error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return errors.Wrap(err, "creating config directory")
	}

	viper.SetDefault("default_project_id", "")
	viper.SetDefault("default_project_color", "#FF1111")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := viper.SafeWriteConfigAs(configPath); err != nil {
			return errors.Wrap(err, "writing default config")
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "reading config")
	}

	return nil
}

func Load() (*Config, error) {
	if err := InitConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "unmarshaling config")
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	viper.Set("default_project_id", cfg.DefaultProjectID)
	viper.Set("default_project_color", cfg.DefaultProjectColor)
	return viper.WriteConfigAs(configPath)
}

func LoadToken() (string, error) {
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return string(data), nil
}

func SaveToken(token string) error {
	if err := os.MkdirAll(filepath.Dir(tokenPath), 0700); err != nil {
		return errors.Wrap(err, "creating token directory")
	}

	return os.WriteFile(tokenPath, []byte(token), 0600)
}

func DeleteToken() error {
	if err := os.Remove(tokenPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
