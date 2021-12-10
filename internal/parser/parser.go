package parser

import (
	"encoding/json"

	"edpad/internal/event"
	"edpad/pkg/log"
)

func Start(parserCh chan string, eventCh chan *event.Event) error {

	go func() {

		for {
			select {
			case text, ok := <-parserCh:
				if !ok {
					return
				}
				if ev := parse(text); ev != nil {
					eventCh <- ev
				}
			}

		}

	}()

	return nil
}

var prevEDtext string

// ED journal entries have at least 'timestamp' and 'event' entries
// others (i.e. joystick events) - don't
func parse(text string) *event.Event {

	log.Debug("parser: %s\n", text)

	var entry event.Entry

	if err := json.Unmarshal([]byte(text), &entry); err != nil {
		log.Err("json unmarshal: %s\n", err)
		return nil
	}

	log.Debug("%+v\n", entry)

	// check is it ED or Joystick event
	if _, ok := entry["event"]; ok {
		if _, ok := entry["timestamp"]; ok {
			// anti-dup check (ED sometimes writes duplicate journal lines)
			if text == prevEDtext {
				log.Warn("skipping duplicate ED Journal line: %s\n", text)
				return nil
			} else {
				prevEDtext = text
				return parseJournal(entry)
			}
		}
	}

	// non ED entry
	return parseJoystick(entry)
}
