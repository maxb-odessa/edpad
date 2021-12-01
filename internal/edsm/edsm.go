package edsm

import "edpad/pkg/log"

type EDSM struct {
	cmdrName string
	apiKey   string
	apiURL   string
	softName string
	softVer  string
	ch       chan string
	inited   bool
}

func Init() *EDSM {
	e := new(EDSM)

	// init, check conn etc

	e.ch = make(chan string, 32)
	go e.sendData()

	return e
}

func (e *EDSM) Send(evType string, text string) {

	if !e.inited {
		return
	}

	if !acceptEvType(evType) {
		return
	}

	e.ch <- text
}

func (e *EDSM) sendData() {
	select {
	case text := <-e.ch:
		log.Debug("TO EDSM: %s\n", text)
	}
}
