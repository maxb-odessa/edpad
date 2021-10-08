package file

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"edpad/conf"
	"edpad/display"
	"edpad/parser"
)

func Read(cfg *conf.Conf) {

	fname, _ := cfg.Get("pipe_path")

	for {
		if fp, err := os.OpenFile(fname, os.O_RDWR, 0666); err != nil {
			log.Printf("waiting for the file: %s\n", err)
			time.Sleep(time.Second * 1)
		} else {
			scanner := bufio.NewScanner(fp)
			go reader(scanner)
			return
		}
	}
}

func reader(scanner *bufio.Scanner) {
	for scanner.Scan() {
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
