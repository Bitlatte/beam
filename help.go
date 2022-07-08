package main

import (
	"os"
	"flag"
	"fmt"
)

func showHelp() {
	fmt.Println("Usage: shex -f <path_to_shell_file> [host_addresses]")
	flag.PrintDefaults()
	os.Exit(1)
}