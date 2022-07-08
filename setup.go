package main

import (
	"os"
	"log"

	"golang.org/x/crypto/ssh"
	"github.com/yahoo/vssh"
)

func setConfig() (config *ssh.ClientConfig) {
	home := os.Getenv("HOME")
	keyPath := home + "/.ssh/id_rsa"
	config, err := vssh.GetConfigPEM("root", keyPath)
	if err != nil {
		log.Fatal(err)
	}
	return config
}