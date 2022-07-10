package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

func GetClientConfig(auth map[string]string) (*ssh.ClientConfig, error) {
	// Read Key File
	key, err := ioutil.ReadFile(auth["key"])
	if err != nil {
		return nil, err
	}

	// Create Signer for Private Key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	// Create Client Config
	config := &ssh.ClientConfig{
		User: auth["user"],
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return config, nil
}

func RemoteCopy(address string, user string, path string, config *ssh.ClientConfig) error {
	client := scp.NewClient(address, config)
	err := client.Connect()
	if err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer client.Close()
	defer file.Close()

	if user == "root" {
		err = client.CopyFromFile(context.Background(), *file, fmt.Sprintf("/%s/script.sh", user), "0755")
		if err != nil {
			return err
		}
	} else {
		err = client.CopyFromFile(context.Background(), *file, fmt.Sprintf("/home/%s/script.sh", user), "0755")
		if err != nil {
			return err
		}
	}
	return nil
}
