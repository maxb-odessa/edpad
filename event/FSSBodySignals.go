package event

/*
type FSSBodySignals struct {
        BodyID   int64  `json:"BodyID,omitempty"`
        BodyName string `json:"BodyName,omitempty"`
        Signals  []struct {
                Count         int64  `json:"Count,omitempty"`
                Type          string `json:"Type,omitempty"`
                TypeLocalised string `json:"Type_Localised,omitempty"`
        } `json:"Signals,omitempty"`
        SystemAddress int64  `json:"SystemAddress,omitempty"`
        Event         string `json:"event,omitempty"`
        Timestamp     string `json:"timestamp,omitempty"`
}
*/

func FSSBodySignals(entry Entry) (*Event, error) {
	return nil, nil
}
