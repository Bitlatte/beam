package utils

import (
	"github.com/spf13/viper"
)

type AuthConfig struct {
	Keys     bool   `mapstructure:"keys"`
	Keyfile  string `mapstructure:"keyfile"`
	Password string `mapstructure:"password"`
	User     string `mapstrucutre:"user"`
}

type Config struct {
	Auth AuthConfig `mapstructure:"auth"`
	Port int        `mapstructure:"port"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./utils")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
