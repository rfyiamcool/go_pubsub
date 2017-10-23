package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var (
	cfg *Config
)

type Config struct {
	Addr     string `toml:"addr"`
	LogPath  string `toml:"log_path"`
	LogLevel string `toml:"log_level"`
	Password string `toml:"password"`
}

func ParseConfigFile(fileName string) (*Config, error) {
	var cfg Config

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	_, err = toml.Decode(string(data), &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
