package parser

import "edpad/display"

/*
type FSDJump struct {
        Body      string `json:"Body,omitempty"`
        BodyID    int64  `json:"BodyID,omitempty"`
        BodyType  string `json:"BodyType,omitempty"`
        Conflicts []struct {
                Faction1 struct {
                        Name    string `json:"Name,omitempty"`
                        Stake   string `json:"Stake,omitempty"`
                        WonDays int64  `json:"WonDays,omitempty"`
                } `json:"Faction1,omitempty"`
                Faction2 struct {
                        Name    string `json:"Name,omitempty"`
                        Stake   string `json:"Stake,omitempty"`
                        WonDays int64  `json:"WonDays,omitempty"`
                } `json:"Faction2,omitempty"`
                Status  string `json:"Status,omitempty"`
                WarType string `json:"WarType,omitempty"`
        } `json:"Conflicts,omitempty"`
        Factions []struct {
                ActiveStates []struct {
                        State string `json:"State,omitempty"`
                } `json:"ActiveStates,omitempty"`
                Allegiance         string  `json:"Allegiance,omitempty"`
                FactionState       string  `json:"FactionState,omitempty"`
                Government         string  `json:"Government,omitempty"`
                Happiness          string  `json:"Happiness,omitempty"`
                HappinessLocalised string  `json:"Happiness_Localised,omitempty"`
                Influence          float64 `json:"Influence,omitempty"`
                MyReputation       int64   `json:"MyReputation,omitempty"`
                Name               string  `json:"Name,omitempty"`
                PendingStates      []struct {
                        State string `json:"State,omitempty"`
                        Trend int64  `json:"Trend,omitempty"`
                } `json:"PendingStates,omitempty"`
                RecoveringStates []struct {
                        State string `json:"State,omitempty"`
                        Trend int64  `json:"Trend,omitempty"`
                } `json:"RecoveringStates,omitempty"`
        } `json:"Factions,omitempty"`
        FuelLevel              float64   `json:"FuelLevel,omitempty"`
        FuelUsed               float64   `json:"FuelUsed,omitempty"`
        JumpDist               float64   `json:"JumpDist,omitempty"`
        Multicrew              bool      `json:"Multicrew,omitempty"`
        Population             int64     `json:"Population,omitempty"`
        StarPos                []float64 `json:"StarPos,omitempty"`
        StarSystem             string    `json:"StarSystem,omitempty"`
        SystemAddress          int64     `json:"SystemAddress,omitempty"`
        SystemAllegiance       string    `json:"SystemAllegiance,omitempty"`
        SystemEconomy          string    `json:"SystemEconomy,omitempty"`
        SystemEconomyLocalised string    `json:"SystemEconomy_Localised,omitempty"`
        SystemFaction          struct {
                FactionState string `json:"FactionState,omitempty"`
                Name         string `json:"Name,omitempty"`
        } `json:"SystemFaction,omitempty"`
        SystemGovernment             string `json:"SystemGovernment,omitempty"`
        SystemGovernmentLocalised    string `json:"SystemGovernment_Localised,omitempty"`
        SystemSecondEconomy          string `json:"SystemSecondEconomy,omitempty"`
        SystemSecondEconomyLocalised string `json:"SystemSecondEconomy_Localised,omitempty"`
        SystemSecurity               string `json:"SystemSecurity,omitempty"`
        SystemSecurityLocalised      string `json:"SystemSecurity_Localised,omitempty"`
        Taxi                         bool   `json:"Taxi,omitempty"`
        Event                        string `json:"event,omitempty"`
        Timestamp                    string `json:"timestamp,omitempty"`
}
*/

func evFSDJump(entry journalEntry) (*display.Data, error) {
	return nil, nil
}
