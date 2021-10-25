package file

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"edpad/cfg"
)

func Start(parserCh chan string) error {

	var fp *os.File
	var err error

	// file is not configured - not reading it
	if cfg.FilePipe == "" {
		return nil
	}

	// wait for the file to appear and start reading it
	for {
		if fp, err = os.OpenFile(cfg.FilePipe, os.O_RDWR, 0666); err != nil {
			log.Printf("waiting for the file: %s\n", err)
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}

	// spawn file reader
	go func() {
		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			if text != "" {
				parserCh <- text
			}
		}
	}()

	return nil
}

/*
func reader(scanner *bufio.Scanner) {


		var cmd display.Cmd

		text := scanner.Text()

		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		switch text {
		case "CLEAR:TOP":
			cmd.ViewPort = display.VIEWPORT_TOP
			cmd.Command = display.CMD_CLEAR
		case "CLEAR:CENTER":
			cmd.ViewPort = display.VIEWPORT_CENTER
			cmd.Command = display.CMD_CLEAR
		case "CLEAR:BOTTOM":
			cmd.ViewPort = display.VIEWPORT_BOTTOM
			cmd.Command = display.CMD_CLEAR

		}

		text = strings.ReplaceAll(text, `\n`, "\n")
		text = strings.ReplaceAll(text, `\r`, "\r")
		text = strings.ReplaceAll(text, `\t`, "\t")
		text = strings.ReplaceAll(text, `\\`, "\\")

		if _, err := parser.Parse(text, &cmd); err != nil {
			cmd := new(display.Cmd)

			display.CmdChan <- cmd
		}
	}
}
*/
