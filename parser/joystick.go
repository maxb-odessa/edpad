package parser

import (
	"edpad/display"
	"edpad/log"
)

func parseJoystick(entry journalEntry) *display.Cmd {
	log.Debug("got joystick entry: %+v\n", entry)
	// TODO
	// note: CMD_CLEAR implement here (on button press?)
	return nil
}
