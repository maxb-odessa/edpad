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

func FSDTarget(e Entry) (*Event, error) {
	sClass := e["StarClass"].(string)
	s := fmt.Sprintf(`<span fgcolor="#A0A0FF"><i><u>Next jump: %s (<span fgcolor="%s"><b>%s</b></span>), left: %.0f</u></i></span>`,
		e["Name"],
		starColor(sClass),
		sClass,
		e["RemainingJumpsInRoute"].(float64))

	return &Event{Type: FSD_TARGET, Text: s}, nil
}
