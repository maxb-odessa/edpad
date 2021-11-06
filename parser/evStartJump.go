package parser

import (
	"edpad/display"
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

func evStartJump(e journalEntry) (*display.Data, error) {

	if jt, ok := e["JumpType"]; !ok || jt != "Hyperspace" {
		return nil, nil
	}

	s := fmt.Sprintf("________________________________\n"+
		"Target: %s (%s)\n", e["StarSystem"].(string), e["StarClass"].(string))

	return &display.Data{Id: "evStartJump", Text: s}, nil
}
