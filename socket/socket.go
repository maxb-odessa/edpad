package socket

import (
	"bufio"
	"log"
	"net"
	"strings"

	"edpad/cfg"
)

func Start(parserCh chan string) error {

	// listening tcp socket is not configured - not reading it
	if cfg.Listen == "" {
		return nil
	}

	lsn, err := net.Listen("tcp4", cfg.Listen)
	if err != nil {
		return err
	}

	go func() {

		// accept remote connections
		for {
			conn, err := lsn.Accept()
			if err != nil {
				log.Println(err)
				continue
			}

			// got a connection - read data from it
			go func() {
				scanner := bufio.NewScanner(conn)
				for scanner.Scan() {
					text := strings.TrimSpace(scanner.Text())
					if text != "" {
						parserCh <- text
					}
				}
			}()

		} //for... accept another connection

	}()

	return nil
}
