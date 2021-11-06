package parser

import (
	"edpad/display"
	"fmt"
)

/*
type FSSSignalDiscovered struct {
        SignalName               string  `json:"SignalName,omitempty"`
        SignalNameLocalised      string  `json:"SignalName_Localised,omitempty"`
        SpawningFaction          string  `json:"SpawningFaction,omitempty"`
        SpawningFactionLocalised string  `json:"SpawningFaction_Localised,omitempty"`
        SpawningState            string  `json:"SpawningState,omitempty"`
        SpawningStateLocalised   string  `json:"SpawningState_Localised,omitempty"`
        SystemAddress            int64   `json:"SystemAddress,omitempty"`
        ThreatLevel              int64   `json:"ThreatLevel,omitempty"`
        TimeRemaining            float64 `json:"TimeRemaining,omitempty"`
        USSType                  string  `json:"USSType,omitempty"`
        USSTypeLocalised         string  `json:"USSType_Localised,omitempty"`
        Event                    string  `json:"event,omitempty"`
        Timestamp                string  `json:"timestamp,omitempty"`
}
*/

func evFSSSignalDiscovered(e journalEntry) (*display.Data, error) {

	sigName := e["SignalName"]
	if sn, ok := e["SignalName_Localised"]; ok {
		sigName = sn
	}

	s := fmt.Sprintf("<span foreground=\"yellow\">FSS signal: <i><b>%s</b></i></span>\n", sigName)

	return &display.Data{Id: "evFSSSignalDiscovered", Text: s}, nil
}
