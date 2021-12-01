package main

import (
	"edpad/internal/cfg"
	"edpad/internal/display"
	"edpad/internal/event"
	"edpad/internal/file"
	"edpad/internal/parser"
	"edpad/internal/socket"
	"edpad/pkg/log"
)

func main() {

	parserCh := make(chan string, 32)
	eventCh := make(chan *event.Event, 32)

	// read cmdline config args
	if err := cfg.Conf(); err != nil {
		return
	}

	// start file reader
	if err := file.Start(parserCh); err != nil {
		log.Fatal("%s\n", err)
	}

	// start socket reader
	if err := socket.Start(parserCh); err != nil {
		log.Fatal("%s\n", err)
	}

	// start json parser
	if err := parser.Start(parserCh, eventCh); err != nil {
		log.Fatal("%s\n", err)
	}

	// start display and wait for it to finish
	if err := display.Start(eventCh); err != nil {
		log.Fatal("%s\n", err)
	}

	return
}
