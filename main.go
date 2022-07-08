package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"regexp"

	"github.com/google/uuid"
)

var (
	help bool
	file string
	port int
)

func init() {
	flag.BoolVar(&help, "h", false, "show cli help")
	flag.StringVar(&file, "f", "", "path to shell script that should be ran")
	flag.IntVar(&port, "port", 22, "ssh port number")
}

func main() {
	
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		showHelp()
		os.Exit(1)
	}

	for _, arg := range args {

		IP, err := regexp.MatchString(`(\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`, arg)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if help {
			showHelp()
		}
		if file == "" {
			log.Fatal("You must provide a file to run")
		}
		
		if IP {
			client := fmt.Sprintf("%s:%s", arg, fmt.Sprint(port))
			f := readFile(file)
			uuid := uuid.New()

			commands := [...]string{
				fmt.Sprintf("echo %s >> %s.sh", f, uuid.String()),
				fmt.Sprintf("chmod a+x %s.sh", uuid.String()),
				fmt.Sprintf("./%s.sh", uuid.String()),
				fmt.Sprintf("rm -rf *.sh"),
			}
		
			for _, cmd := range commands {
				execute(client, cmd)
			}
		} else {
			log.Fatal("Invalid IP Address Entered")
		}
	}
}