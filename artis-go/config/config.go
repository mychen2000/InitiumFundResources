package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// 运行环境: either `development` or `production`
	Env string `yaml:"env" envconfig:"env"`

	Alpaca struct {
		ApiKeyID     string `yaml:"APIKeyID" envconfig:"APCA_API_KEY_ID"`
		ApiSecretKey string `yaml:"APISecretKey" envconfig:"APCA_API_SECRET_KEY"`
	} `yaml:"Alpaca"`

	Log struct {
		Directory string `yaml:"Directory"`
		FileName  string `yaml:"FileName"`
		LogLevel  string `yaml:"Level"`
	} `yaml:"Log"`
}

func (c Config) IsDevelopment() bool {
	return c.Env == "development"
}

func LoadConfig(path string) *Config {
	var cfg Config

	if path != "" {
		path, err := filepath.Abs(path)
		if err != nil {
			panic("Error while opening config files: " + err.Error())
		}
		LoadConfigFromYaml(&cfg, path)
	}

	// TODO: load config from environment variables

	return &cfg
}

func LoadConfigFromYaml(cfg *Config, path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("Error while opening config files: " + err.Error())
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		panic("Error while initializing configs: " + err.Error())
	}
}
