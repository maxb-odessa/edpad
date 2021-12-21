package parser

import (
	"edpad/internal/event"
	"edpad/pkg/log"
)

func parseJoystick(entry event.Entry) *event.Event {
	log.Debug("got joystick entry: %+v\n", entry)
	// TODO
	// note: CMD_CLEAR implement here (on button press?)
	return nil
}
