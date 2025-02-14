package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DefaultProjectID string `yaml:"default_project_id"`
}

const configPath = "~/.config/tickli/config.yaml"

func Load() (*Config, error) {
	path := expandPath(configPath)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &Config{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.Wrap(err, "parsing config file")
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	path := expandPath(configPath)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return errors.Wrap(err, "creating config directory")
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "marshaling config")
	}

	return os.WriteFile(path, data, 0600)
}

func expandPath(path string) string {
	if path[0] == '~' {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, path[1:])
	}
	return path
}

//TODO: setup the configuration

func getTokenPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local/share/tickli/token"), nil
}

func LoadToken() (string, error) {
	path, err := getTokenPath()
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	return string(data), nil
}

func SaveToken(token string) error {
	path, err := getTokenPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return errors.Wrap(err, "creating token directory")
	}

	return os.WriteFile(path, []byte(token), 0600)
}

func DeleteToken() error {
	path, err := getTokenPath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
