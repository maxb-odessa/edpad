package main

import (
	"log"

	"edpad/cfg"
	"edpad/display"
	"edpad/file"
)

func main() {

	// read cmdline config args
	if err := cfg.Conf(); err != nil {
		return
	}

	// start file reader
	if err := file.Start(); err != nil {
		log.Fatalln(err)
	}

	// start display and wait for it to finish
	if err := display.Start(); err != nil {
		log.Fatalln(err)
	}

	return
}
