package event

import (
	"edpad/pkg/log"
	"fmt"
	"strings"
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
const (
	SOLAR_RADIUS     = 696340000.0
	EARTH_RADIUS     = 6371.0 * 1000.0
	LIGHT_SECOND     = 299792.0 * 1000.
	MIN_RING_OUT_RAD = 20.0 * LIGHT_SECOND
)

func isMainStar(e Entry) bool {

	log.Debug("main star:%s saved main star:%s\n", e["BodyName"].(string), mainStarName)
	if mainStarName == e["BodyName"].(string) {
		return true
	}

	return false
}

func scanStar(e Entry) (Type, string) {

	var isMain Type
	var prefix string

	if isMainStar(e) {
		isMain = MAIN_STAR
		prefix = "Star"
	} else {
		isMain = SEC_STAR
		prefix = "   *"
	}

	discovered := ""
	if isMain == MAIN_STAR && e["WasDiscovered"].(bool) {
		discovered = ` <span fgcolor="yellow">+D</span>`
	}

	sType := e["StarType"].(string)
	sClass := fmt.Sprintf("%.0f", e["Subclass"].(float64))

	fgColor := starColor(sType)

	starType := `<span size="larger" fgcolor="` + fgColor + `"><b>` + sType + sClass + `</b></span>`

	ringed := ""
	if _, ok := e["Rings"].([]interface{}); ok {
		ringed = ` +R`
	}

	star := fmt.Sprintf(`%s: %s, <span size="small"><i>sM:%.2f sR:%.2f tK:%.0fk%s%s</i></span>`,
		prefix,
		starType,
		e["StellarMass"].(float64),
		e["Radius"].(float64)/SOLAR_RADIUS,
		e["SurfaceTemperature"].(float64)/1000,
		discovered,
		ringed)

	log.Debug("STAR: %s\n", star)

	return isMain, star
}

func scanPlanet(e Entry) (Type, string) {

	var planets []string

	if p := rarePlanet(e); p != "" {
		planets = append(planets, p)
	}

	if p := wideRing(e); p != "" {
		planets = append(planets, p)
	}

	return PLANET, strings.Join(planets[:], "\n")
}

func rarePlanet(e Entry) string {

	pColor := "#C0C0C0"
	pClass := e["PlanetClass"].(string)
	terraFormable := ""

	switch pClass {
	case "Earthlike body":
		pColor = "#70FF70"
	case "Water world":
		pColor = "#7070FF"
	case "Ammonia world":
		pColor = "#FF7070"
	default:
		if tf, ok := e["TerraformState"].(string); ok && tf != "" {
			terraFormable = ` +TF`
			if pClass == "High metal content body" {
				pClass = "HMC"
			}
		} else {
			return "" // not interested in other planets
		}
	}

	ringed := ""
	if _, ok := e["Rings"]; ok {
		ringed = ` +R`
	}

	discovered := ""
	if e["WasDiscovered"].(bool) {
		discovered = ` <span fgcolor="yellow">+D</span>`
	}

	pMass := e["MassEM"].(float64)
	pRad := e["Radius"].(float64) / EARTH_RADIUS

	planet := fmt.Sprintf(`Body: id:%.0f, <span fgcolor="%s">%s</span>, <span size="small"><i>eM:%.2f, eR:%.2f<b>%s%s%s</b></i></span>`,
		e["BodyID"].(float64),
		pColor,
		pClass,
		pMass,
		pRad,
		discovered,
		ringed,
		terraFormable)

	return planet
}

func getWideRing(e Entry) (nRings int, outRad float64, yes bool) {

	rings, ok := e["Rings"].([]interface{})
	if !ok {
		return
	}

	maxOutRad := 0.0

	for _, r := range rings {

		nRings++

		ring := r.(map[string]interface{})
		outRad = ring["OuterRad"].(float64)

		if maxOutRad < outRad {
			maxOutRad = outRad
		}

	}

	if maxOutRad < MIN_RING_OUT_RAD {
		return
	}

	yes = true
	return
}

func wideRing(e Entry) string {

	if nrs, or, yes := getWideRing(e); yes {

		planet := fmt.Sprintf(`Body: id:%.0f, <span fgcolor="gray">Wide Ring</span>, <span size="small">nR:%d, lsR:%.2f</span>`,
			e["BodyID"].(float64),
			nrs,
			or/LIGHT_SECOND)

		if e["WasDiscovered"].(bool) {
			planet += ` <span size="small" fgcolor="yellow"><b>+D</b></span>`
		}

		return planet
	}

	return ""
}