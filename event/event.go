package event

type Entry map[string]interface{}

type Type int

type Event struct {
	Type Type
	Text string
}

const (
	FSD_TARGET = iota
	SYSTEM_NAME
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

func starColor(class string) (fgColor string) {

	c := class[0:1]

	switch c {
	case "O":
		fgColor = `#EEEEEE`
	case "B":
		fgColor = `#EEEE80`
	case "A":
		fgColor = `#EEEEAA`
	case "F":
		fgColor = `#EEEECC`
	case "G":
		fgColor = `#EEEE20`
	case "K":
		fgColor = `#EEAA20`
	case "M":
		fgColor = `#EE8080`
	case "N":
		fgColor = `#2020EE`
	case "D":
		fgColor = `#FFFFFF`
	case "T", "Y", "L":
		fgColor = `#AA3030`
	case "H":
		fgColor = `#505050`
	case "W":
		fgColor = `#FFFFFF`
	default:
		fgColor = `#E0E0E0`
	}

	return
}

// must be set by FSDJump() and read by Scan()
var mainStarName string
