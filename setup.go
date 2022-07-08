package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"github.com/yahoo/vssh"
)

func setConfig() (config *ssh.ClientConfig) {
	config, err := vssh.GetConfigPEM("root", "/home/bitlatte/.ssh/id_rsa")
	if err != nil {
		log.Fatal(err)
	}
	return config
}