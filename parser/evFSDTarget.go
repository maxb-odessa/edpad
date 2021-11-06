package parser

import (
	"edpad/display"
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

func evFSDTarget(entry journalEntry) (*display.Data, error) {
	s := fmt.Sprintf("<span foreground=\"red\"><i>Jumping to: %s (%s), jumps left: %.0f</i></span>\n",
		entry["Name"].(string),
		entry["StarClass"].(string),
		entry["RemainingJumpsInRoute"].(float64))

	return &display.Data{Id: "FSDTarget", Text: s}, nil
}
