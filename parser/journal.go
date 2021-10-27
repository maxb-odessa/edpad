package parser

import (
	"edpad/display"
	"edpad/log"
)

func parseJournal(entry journalEntry) *display.Cmd {

	var data string
	var err error

	log.Debug("got Journal entry: %+v\n", entry)

	evname := entry["event"]
	switch evname {
	case "StartJump":
		data, err = evStartJump(entry)
	case "FSDTarget":
		data, err = evFSDTarget(entry)
	case "FSDJump":
		data, err = evFSDJump(entry)
	case "FSSDiscoveryScan":
		data, err = evFSSDiscoveryScan(entry)
	case "FSSSignalDiscovered":
		data, err = evFSSSignalDiscovered(entry)
	case "FSSBodySignals":
		data, err = evFSSBodySignals(entry)
	case "SAASignalsFound":
		data, err = evSAASignalsFound(entry)
	case "SAAScanComplete":
		data, err = evSAAScanComplete(entry)
	case "Scan":
		data, err = evScan(entry)
	default:
		log.Debug("ignoring Journal event '%s'\n", evname)
	}

	if err != nil {
		log.Err("bad journal entry: %s\n", err)
		return nil
	}

	return &display.Cmd{Command: display.CMD_TEXT, Data: data}
}
