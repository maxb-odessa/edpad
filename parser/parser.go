package parser

import (
	"edpad/display"
)

func Start(parserCh chan string, displayCh chan *display.Cmd) error {

	for {
		select {
		case text := <-parserCh:
			if cmd := parse(text); cmd != nil {
				displayCh <- cmd
			}
		}

	}

	return nil
}

func parse(text string) *display.Cmd {

	return nil
}
