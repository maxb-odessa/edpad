package event

type Entry map[string]interface{}

type Type int

type Event struct {
	Type Type
	Text string
}

const (
	FSD_TARGET = iota
	START_JUMP
	MAIN_STAR
	SEC_STAR
	PLANET
	BODY_SIGNALS
	FSS_SIGNALS
)

type Handler func(e Entry) (*Event, error)

var Handlers = map[string]Handler{
	"StartJump":           StartJump,
	"FSDJump":             FSDJump,
	"FSDTarget":           FSDTarget,
	"FSSBodySignals":      FSSBodySignals,
	"FSSDiscoveryScan":    FSSDiscoveryScan,
	"FSSSignalDiscovered": FSSSignalDiscovered,
	"SAAScanComplete":     SAAScanComplete,
	"SAASignalsFound":     SAASignalsFound,
	"Scan":                Scan,
}
