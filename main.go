package main

import (
	"log"
	"os"

	"edpad/conf"
	"edpad/display"
	"edpad/file"
)

func main() {
	confPath := os.Getenv("HOME") + "/.local/etc/edpad.conf"

	if len(os.Args) > 1 {
		confPath = os.Args[1]
	}

	cfg, err := conf.Read(confPath)
	if err != nil {
		log.Println(err)
		return
	}

	// start file reader
	file.Read(cfg)

	// start display and wait for it to finish
	display.Start(cfg)

	return
}
