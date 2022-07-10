package utils

import (
	"fmt"

	"github.com/yahoo/vssh"
)

var user string
var key string

func CreateSession(config *Config, host *HostConfig) (*vssh.VSSH, error) {
	if host.Auth.User == "" {
		user = config.Auth.User
	} else {
		user = host.Auth.User
	}
	if host.Auth.Key == "" {
		key = config.Auth.Key
	} else {
		key = host.Auth.Key
	}

	vs := vssh.New().Start().OnDemand()

	clientConfig, err := vssh.GetConfigPEM(user, key)
	if err != nil {
		return &vssh.VSSH{}, err
	}

	client := fmt.Sprintf("%s:%s", host.Address, fmt.Sprint(host.Port))

	vs.AddClient(client, clientConfig, vssh.SetMaxSessions(2))
	vs.Wait()

	return vs, nil
}
