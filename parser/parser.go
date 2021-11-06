package parser

import (
	"encoding/json"

	"edpad/display"
	"edpad/log"
)

type journalEntry map[string]interface{}

func Start(parserCh chan string, displayCh chan *display.Data) error {

	go func() {

		for {
			select {
			case text, ok := <-parserCh:
				if !ok {
					return
				}
				if cmd := parse(text); cmd != nil {
					displayCh <- cmd
				}
			}

		}

	}()

	return nil
}

// ED journal entries have at least 'timestamp' and 'event' entries
// others (i.e. joystick events) - don't
func parse(text string) *display.Data {

	log.Debug("parser: %s\n", text)

	entry := make(journalEntry)

	if err := json.Unmarshal([]byte(text), &entry); err != nil {
		log.Err("json unmarshal: %s\n", err)
		return nil
	}

	log.Debug("%+v\n", entry)

	// check is it ED or Joystick event
	if _, ok := entry["event"]; ok {
		if _, ok := entry["timestamp"]; ok {
			return parseJournal(entry)
		}
	}

	// non ED entry
	return parseJoystick(entry)
}
