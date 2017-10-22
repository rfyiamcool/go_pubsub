package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Addr     string `toml:"addr"`
	LogPath  string `toml:"log_path"`
	LogLevel string `toml:"log_level"`
}

type DBConfig struct {
	Password string `toml:"password"`
	DBName   string `toml:"db_name"`
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
