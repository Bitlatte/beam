package utils

import "github.com/spf13/viper"

type AuthConfig struct {
	Key  string `mapstructure:"key"`
	User string `mapstructure:"user"`
}

type HostConfig struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

type Config struct {
	Auth  AuthConfig   `mapstructure:"auth"`
	Hosts []HostConfig `mapstructure:"hosts"`
}

var config Config

func GetSSHConfig() (*Config, error) {
	err := viper.Unmarshal(&config)
	if err != nil {
		return &Config{}, err
	}
	return &config, nil
}
