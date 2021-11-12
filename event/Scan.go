package event

import (
	"edpad/log"
	"fmt"
)

/* PLANET SCAN
type Scan struct {
        AscendingNode         float64 `json:"AscendingNode,omitempty"`
        Atmosphere            string  `json:"Atmosphere,omitempty"`
        AtmosphereComposition []struct {
                Name    string  `json:"Name,omitempty"`
                Percent float64 `json:"Percent,omitempty"`
        } `json:"AtmosphereComposition,omitempty"`
        AtmosphereType string  `json:"AtmosphereType,omitempty"`
        AxialTilt      float64 `json:"AxialTilt,omitempty"`
        BodyID         int64   `json:"BodyID,omitempty"`
        BodyName       string  `json:"BodyName,omitempty"`
        Composition    struct {
                Ice   float64 `json:"Ice,omitempty"`
                Metal float64 `json:"Metal,omitempty"`
                Rock  float64 `json:"Rock,omitempty"`
        } `json:"Composition,omitempty"`
        DistanceFromArrivalLS float64 `json:"DistanceFromArrivalLS,omitempty"`
        Eccentricity          int64   `json:"Eccentricity,omitempty"`
        Landable              bool    `json:"Landable,omitempty"`
        MassEM                float64 `json:"MassEM,omitempty"`
        MeanAnomaly           float64 `json:"MeanAnomaly,omitempty"`
        OrbitalInclination    float64 `json:"OrbitalInclination,omitempty"`
        OrbitalPeriod         float64 `json:"OrbitalPeriod,omitempty"`
        Parents               []struct {
                Null   int64 `json:"Null,omitempty"`
                Planet int64 `json:"Planet,omitempty"`
                Star   int64 `json:"Star,omitempty"`
        } `json:"Parents,omitempty"`
        Periapsis    float64 `json:"Periapsis,omitempty"`
        PlanetClass  string  `json:"PlanetClass,omitempty"`
        Radius       float64 `json:"Radius,omitempty"`
        ReserveLevel string  `json:"ReserveLevel,omitempty"`
        Rings        []struct {
                InnerRad  int64  `json:"InnerRad,omitempty"`
                MassMT    int64  `json:"MassMT,omitempty"`
                Name      string `json:"Name,omitempty"`
                OuterRad  int64  `json:"OuterRad,omitempty"`
                RingClass string `json:"RingClass,omitempty"`
        } `json:"Rings,omitempty"`
        RotationPeriod     float64 `json:"RotationPeriod,omitempty"`
        ScanType           string  `json:"ScanType,omitempty"`
        SemiMajorAxis      float64 `json:"SemiMajorAxis,omitempty"`
        StarSystem         string  `json:"StarSystem,omitempty"`
        SurfaceGravity     float64 `json:"SurfaceGravity,omitempty"`
        SurfacePressure    float64 `json:"SurfacePressure,omitempty"`
        SurfaceTemperature float64 `json:"SurfaceTemperature,omitempty"`
        SystemAddress      int64   `json:"SystemAddress,omitempty"`
        TerraformState     string  `json:"TerraformState,omitempty"`
        TidalLock          bool    `json:"TidalLock,omitempty"`
        Volcanism          string  `json:"Volcanism,omitempty"`
        WasDiscovered      bool    `json:"WasDiscovered,omitempty"`
        WasMapped          bool    `json:"WasMapped,omitempty"`
        Event              string  `json:"event,omitempty"`
        Timestamp          string  `json:"timestamp,omitempty"`
}
*/

/* STAR SCAN
type Scan struct {
	AbsoluteMagnitude     float64 `json:"AbsoluteMagnitude"`
	AgeMY                 int64   `json:"Age_MY"`
	AxialTilt             int64   `json:"AxialTilt"`
	BodyID                int64   `json:"BodyID"`
	BodyName              string  `json:"BodyName"`
	DistanceFromArrivalLS int64   `json:"DistanceFromArrivalLS"`
	Luminosity            string  `json:"Luminosity"`
	Radius                int64   `json:"Radius"`
	Rings                 []struct {
		InnerRad  int64  `json:"InnerRad"`
		MassMT    int64  `json:"MassMT"`
		Name      string `json:"Name"`
		OuterRad  int64  `json:"OuterRad"`
		RingClass string `json:"RingClass"`
	} `json:"Rings"`
	RotationPeriod     float64 `json:"RotationPeriod"`
	ScanType           string  `json:"ScanType"`
	StarSystem         string  `json:"StarSystem"`
	StarType           string  `json:"StarType"`
	StellarMass        float64 `json:"StellarMass"`
	Subclass           int64   `json:"Subclass"`
	SurfaceTemperature int64   `json:"SurfaceTemperature"`
	SystemAddress      int64   `json:"SystemAddress"`
	WasDiscovered      bool    `json:"WasDiscovered"`
	WasMapped          bool    `json:"WasMapped"`
	Event              string  `json:"event"`
	Timestamp          string  `json:"timestamp"`
}
*/

func Scan(entry Entry) (*Event, error) {

	// Star scan
	if _, ok := entry["StarType"]; ok {
		t, s := scanStar(entry)
		return &Event{Type: t, Text: s}, nil
	}

	// Planet scan
	if _, ok := entry["PlanetClass"]; ok {
		t, p := scanPlanet(entry)
		// no text is ok - skip this event
		if p == "" {
			return nil, nil
		}
		return &Event{Type: t, Text: p}, nil
	}

	// huh?
	return nil, fmt.Errorf("unknown body type scan")
}

// in meters
const SOLAR_RADIUS = 696340000.0
const EARTH_RADIUS = 6371.0 * 1000.0
const LIGHT_SECOND = 299792.0 * 1000.

// in light seconds
const MIN_RING_RAD = 10.0

// must be set by FSDJump()
var mainStarName string

func isMainStar(e Entry) bool {

	log.Debug("main star:%s saved main star:%s\n", e["BodyName"].(string), mainStarName)
	if mainStarName == e["BodyName"].(string) {
		return true
	}

	return false
}

func scanStar(e Entry) (t Type, star string) {

	var isMain Type
	var prefix string

	if isMainStar(e) {
		isMain = MAIN_STAR
		prefix = "Star:"
	} else {
		isMain = SEC_STAR
		prefix = "   +:"
	}

	defer func() {
		if isMain == MAIN_STAR && e["WasDiscovered"].(bool) && star != "" {
			star += ` <span size="smaller" fgcolor="yellow"><b>(!)</b></span>`
		}
	}()

	fgColor := `#FFFFFF`
	sType := e["StarType"].(string)
	sClass := fmt.Sprintf("%.0f", e["Subclass"].(float64))

	if sType[0:1] == "O" {
		fgColor = `#EEEEEE`
	} else if sType[0:1] == "B" {
		fgColor = `#EEEE80`
	} else if sType[0:1] == "A" {
		fgColor = `#EEEEAA`
	} else if sType[0:1] == "F" {
		fgColor = `#EEEECC`
	} else if sType[0:1] == "G" {
		fgColor = `#EEEE20`
	} else if sType[0:1] == "K" {
		fgColor = `#EEAA20`
	} else if sType[0:1] == "M" {
		fgColor = `#EE8080`
	} else if sType[0:1] == "N" {
		fgColor = `#2020EE`
	} else if sType[0:1] == "D" {
		fgColor = `#FFFFFF`
	} else if sType[0:1] == "T" || sType[0:1] == "Y" || sType[0:1] == "L" {
		fgColor = `#AA3030`
	} else if sType[0:1] == "H" {
		fgColor = `#505050`
	}

	starType := `<span size="larger" fgcolor="` + fgColor + `">` + sType + sClass + `</span>`

	star = fmt.Sprintf("%s %s, solM:%.2f, solR:%.2f, tK:%.0f",
		prefix,
		starType,
		e["StellarMass"].(float64),
		e["Radius"].(float64)/SOLAR_RADIUS,
		e["SurfaceTemperature"].(float64))

	return isMain, star
}

func scanPlanet(e Entry) (Type, string) {

	if p := rarePlanet(e); p != "" {
		return PLANET, p
	}

	if p := wideRing(e); p != "" {
		return RING, p
	}

	return PLANET, ""
}

func rarePlanet(e Entry) string {

	pClass := e["PlanetClass"].(string)

	pColor := "#808080"
	switch pClass {
	case "Earthlike body":
		pColor = "#00FF30"
	case "Water world":
		pColor = "#0070FF"
	case "Ammonia world":
		pColor = "#FF3000"
	default:
		return "" // not interested in other planets
	}

	ringed := ""
	if _, ok := e["Rings"]; ok {
		ringed = " <i>(ringed)</i>"
	}

	pMass := e["MassEM"].(float64)
	pRad := e["Radius"].(float64) / EARTH_RADIUS

	planet := fmt.Sprintf(`Body: <span fgcolor="%s">%-14s</span> eM:%.2f, eR:%.2f%s`,
		pColor,
		pClass,
		pMass,
		pRad,
		ringed)

	if e["WasDiscovered"].(bool) {
		planet += ` <span size="smaller" fgcolor="yellow"><b>(!)</b></span>`
	}

	return planet
}

func wideRing(e Entry) string {

	rings, ok := e["Rings"].([]interface{})
	if !ok {
		return ""
	}

	maxOutRad := 0.0
	rNum := 0

	for _, r := range rings {

		rNum++

		ring := r.(map[string]interface{})
		outRad := ring["OuterRad"].(float64)

		if maxOutRad < outRad {
			maxOutRad = outRad
		}

	}

	if maxOutRad >= MIN_RING_RAD {
		return fmt.Sprintf(`Ring: <span fgcolor="gray">bodyID:%.0f, rNum:%d, lsRad:%.2f</span>`,
			e["BodyID"].(float64),
			rNum,
			maxOutRad/LIGHT_SECOND)
	}

	return ""
}
