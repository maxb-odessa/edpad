package parser

import (
	"edpad/display"
	"edpad/log"
)

func Start(parserCh chan string, displayCh chan *display.Cmd) error {

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

func parse(text string) *display.Cmd {

	log.Debug("parser: %s\n", text)

	return &display.Cmd{
		Data:    text,
		Command: display.CMD_TEXT,
	}
}
