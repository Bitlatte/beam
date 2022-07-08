package main

import (
	"time"
	"context"
	"fmt"
	"log"
	"github.com/yahoo/vssh"
)

func execute(client string, cmd string) {
	vs := vssh.New().Start()
	config := setConfig()

	vs.AddClient(client, config)
	vs.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout, _ := time.ParseDuration("6s")
	respChan := vs.Run(ctx, cmd, timeout)

	resp := <- respChan
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