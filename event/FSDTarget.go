package event

import (
	"fmt"
)

/*
type FSDTarget struct {
        Name                  string `json:"Name,omitempty"`
        RemainingJumpsInRoute int64  `json:"RemainingJumpsInRoute,omitempty"`
        StarClass             string `json:"StarClass,omitempty"`
        SystemAddress         int64  `json:"SystemAddress,omitempty"`
        Event                 string `json:"event,omitempty"`
        Timestamp             string `json:"timestamp,omitempty"`
}
*/

func FSDTarget(entry Entry) (*Event, error) {
	s := fmt.Sprintf(`<span foreground="#9090FF"><i><u>Next jump: %s (%s), left: %.0f</u></i></span>`,
		entry["Name"].(string),
		entry["StarClass"].(string),
		entry["RemainingJumpsInRoute"].(float64))

	return &Event{Type: FSD_TARGET, Text: s}, nil
}
