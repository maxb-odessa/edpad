package event

import (
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

func FSSDiscoveryScan(e Entry) (*Event, error) {
	s := fmt.Sprintf(`<span foreground="cyan">Body signals: %0.f(%.0f)</span>`,
		e["BodyCount"].(float64),
		e["NonBodyCount"].(float64))
	return &Event{Type: BODY_SIGNALS, Text: s}, nil
}
