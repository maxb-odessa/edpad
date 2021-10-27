package parser

import (
	"edpad/display"
	"edpad/log"
)

func parseJoystick(entry journalEntry) *display.Cmd {
	log.Debug("got joystick entry: %+v\n", entry)
	// TODO
	return nil
}
