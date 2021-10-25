package parser

import (
	"edpad/display"
	"edpad/event"
	"log"
)

func Parse(in string, cmf *display.Cmd) (string, error) {
	ev := event.Scan{}
	log.Println("parsing")
	return ev.ScanType, nil
}
