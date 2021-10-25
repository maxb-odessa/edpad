package main

import (
	"log"

	"edpad/cfg"
	"edpad/display"
	"edpad/file"
	"edpad/parser"
	"edpad/socket"
)

func main() {

	parserCh := make(chan string, 32)
	displayCh := make(chan *display.Cmd, 32)

	// read cmdline config args
	if err := cfg.Conf(); err != nil {
		return
	}

	// start file reader
	if err := file.Start(parserCh); err != nil {
		log.Fatalln(err)
	}

	// start socket reader
	if err := socket.Start(parserCh); err != nil {
		log.Fatalln(err)
	}

	// start json parser
	if err := parser.Start(parserCh, displayCh); err != nil {
		log.Fatalln(err)
	}

	// start display and wait for it to finish
	if err := display.Start(displayCh); err != nil {
		log.Fatalln(err)
	}

	return
}
