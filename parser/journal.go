package parser

import (
	"edpad/event"
	"edpad/log"
)

func parseJournal(entry event.Entry) *event.Event {

	defer func() {
		if r := recover(); r != nil {
			log.Warn("parse journal failed, recoverd in %v", r)
		}
	}()

	log.Debug("got Journal entry: %+v\n", entry)

	evname, ok := entry["event"]
	if !ok {
		log.Warn("broken journal entry: missing 'event' field\n")
		return nil
	}

	if evfunc, ok := event.Handlers[evname.(string)]; !ok {
		log.Warn("handling of event '%s' is not implemented\n", evname)
		return nil
	} else if ev, err := evfunc(entry); err == nil {
		return ev
	}

	return nil
}
