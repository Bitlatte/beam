package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yahoo/vssh"
	"golang.org/x/crypto/ssh"
)

func getConfig() (*ssh.ClientConfig, error) {
	homePath := os.Getenv("HOME")
	config, err := vssh.GetConfigPEM("root", (homePath + "/.ssh/id_rsa"))
	return config, err
}

func Execute(client string, cmd string) {
	vs := vssh.New().Start()
	config, err := getConfig()

	if err != nil {
		log.Fatal(err)
	}

	vs.AddClient(client, config)
	vs.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout, _ := time.ParseDuration("6s")
	respChan := vs.Run(ctx, cmd, timeout)

	resp := <-respChan
	if err := resp.Err(); err != nil {
		log.Fatal(err)
	}

	stream := resp.GetStream()
	defer stream.Close()

	for stream.ScanStdout() {
		txt := stream.TextStdout()
		fmt.Println(txt)
	}
}
