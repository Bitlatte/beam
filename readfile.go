package main

import (
	"log"
	"io/ioutil"
)

func readFile(path string)([]byte) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content
}