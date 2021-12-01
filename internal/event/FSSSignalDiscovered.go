package event

import (
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

func FSSSignalDiscovered(e Entry) (*Event, error) {

	sigName := e["SignalName"].(string)
	if snl, ok := e["SignalName_Localised"].(string); ok && snl != "" {
		sigName = snl
	}

	isStation := ""
	if isSt, ok := e["IsStation"].(bool); ok && isSt {
		isStation = ` <span fgcolor="white">(station)</span>`
	}

	sig := fmt.Sprintf(`Signal: <span size="small" fgcolor="#FFA050">%s%s</span>`, sigName, isStation)

	return &Event{Type: FSS_SIGNALS, Text: sig}, nil
}
