package parser

import (
	"edpad/display"
	"fmt"
)

/*
type FSSDiscoveryScan struct {
        BodyCount     int64   `json:"BodyCount,omitempty"`
        NonBodyCount  int64   `json:"NonBodyCount,omitempty"`
        Progress      float64 `json:"Progress,omitempty"`
        SystemAddress int64   `json:"SystemAddress,omitempty"`
        SystemName    string  `json:"SystemName,omitempty"`
        Event         string  `json:"event,omitempty"`
        Timestamp     string  `json:"timestamp,omitempty"`
}
*/

func evFSSDiscoveryScan(e journalEntry) (*display.Data, error) {
	s := fmt.Sprintf("Signals detected: %0.f/%.0f\n", e["BodyCount"].(float64), e["NonBodyCount"].(float64))
	return &display.Data{Id: "evFSSDiscoveryScan", Text: s}, nil
}
