package utils

import (
	"github.com/spf13/viper"
)

type AuthConfig struct {
	Key  string `mapstructure:"key"`
	User string `mapstructure:"user"`
}

type HostConfig struct {
	Address string     `mapstructure:"address"`
	Port    int        `mapstructure:"port"`
	Auth    AuthConfig `mapstructure:"auth"`
}

type Config struct {
	Auth  AuthConfig   `mapstructure:"auth"`
	Hosts []HostConfig `mapstructure:"hosts"`
}

var config Config

func GetBeamConfig() (*Config, error) {
	err := viper.Unmarshal(&config)
	if err != nil {
		return &Config{}, err
	}
	return &config, nil
}

func GetAuth(config *Config, host *HostConfig) map[string]string {
	auth := map[string]string{
		"user": host.Auth.User,
		"key":  host.Auth.Key,
	}

	if auth["user"] == "" {
		auth["user"] = config.Auth.User
	}
	if auth["key"] == "" {
		auth["key"] = config.Auth.Key
	}

	return auth
}
