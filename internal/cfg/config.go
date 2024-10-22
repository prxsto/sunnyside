package cfg

import (
	"errors"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ConfigFile string
}

func FromFile(path string) (*Config, error) {
	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := toml.Unmarshal(cfgFile, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func ToFile(filePath string, cfg *Config) error {
	b, err := toml.Marshaler(cfg)
}
