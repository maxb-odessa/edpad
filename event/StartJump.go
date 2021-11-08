package event

import (
	"fmt"
)

/*
type StartJump struct {
        JumpType      string `json:"JumpType,omitempty"`
        StarClass     string `json:"StarClass,omitempty"`
        StarSystem    string `json:"StarSystem,omitempty"`
        SystemAddress int64  `json:"SystemAddress,omitempty"`
        Event         string `json:"event,omitempty"`
        Timestamp     string `json:"timestamp,omitempty"`
}
*/

func StartJump(e Entry) (*Event, error) {

	if jt, ok := e["JumpType"]; !ok || jt != "Hyperspace" {
		return nil, nil
	}

	s := fmt.Sprintf("________________________________"+
		"Target: %s (%s)\n", e["StarSystem"].(string), e["StarClass"].(string))

	return &Event{Type: START_JUMP, Text: s}, nil
}
