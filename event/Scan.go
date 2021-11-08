package event

import (
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
		t, s, err := scanStar(entry)
		return &Event{Type: t, Text: s}, err
	}

	// Planet scan
	if _, ok := entry["PlanetClass"]; ok {
		t, p, err := scanPlanet(entry)
		return &Event{Type: t, Text: p}, err
	}

	// huh?
	return nil, fmt.Errorf("unknown body type scan")
}

const SOLAR_RADIUS = 696340000.0

func isMainStar(e Entry) bool {
	defer func() {
		recover()
	}()

	v, ok := e["Parents"].([]interface{})[0].(map[string]interface{})["Null"].(float64)
	if ok && v != 0.0 {
		return false
	}

	return true
}

func scanStar(e Entry) (Type, string, error) {

	var stype Type
	var sname string
	var discovered string

	if isMainStar(e) {
		stype = MAIN_STAR
		sname = `<span weight="bold" foreground="yellow">Star</span>`
	} else {
		stype = SEC_STAR
		sname = `<span foreground="red">Star</span>`
	}

	if stype == MAIN_STAR && e["WasDiscovered"].(bool) {
		discovered = `<span foreground="green"><i> Discovered!</i></span>`
	}

	star := fmt.Sprintf("%s: %s%.0f, m:%.2f, r:%.2f, t:%.0f%s",
		sname,
		e["StarType"].(string),
		e["Subclass"].(float64),
		e["StellarMass"].(float64),
		e["Radius"].(float64)/SOLAR_RADIUS,
		e["SurfaceTemperature"].(float64),
		discovered)

	return stype, star, nil
}

func scanPlanet(entry Entry) (Type, string, error) {

	return PLANET, "", nil
}
