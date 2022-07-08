package main

import (
	"os"
	"flag"
	"fmt"
)

func showHelp() {
	fmt.Println("Usage: beam <options> [host_addresses]")
	fmt.Println("Example:\n	beam -f <path_to_shell_file> [host_addresses]")
	fmt.Println("Commands")
	fmt.Println("	key		Show your ssh key")
	fmt.Println("Options")
	flag.PrintDefaults()
	os.Exit(1)
}