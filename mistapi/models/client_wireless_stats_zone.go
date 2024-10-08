package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientWirelessStatsZone represents a ClientWirelessStatsZone struct.
type ClientWirelessStatsZone struct {
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Since                *int           `json:"since,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWirelessStatsZone.
// It customizes the JSON marshaling process for ClientWirelessStatsZone objects.
func (c ClientWirelessStatsZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWirelessStatsZone object to a map representation for JSON marshaling.
func (c ClientWirelessStatsZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Since != nil {
        structMap["since"] = c.Since
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWirelessStatsZone.
// It customizes the JSON unmarshaling process for ClientWirelessStatsZone objects.
func (c *ClientWirelessStatsZone) UnmarshalJSON(input []byte) error {
    var temp clientWirelessStatsZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "since")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Id = temp.Id
    c.Since = temp.Since
    return nil
}

// clientWirelessStatsZone is a temporary struct used for validating the fields of ClientWirelessStatsZone.
type clientWirelessStatsZone  struct {
    Id    *uuid.UUID `json:"id,omitempty"`
    Since *int       `json:"since,omitempty"`
}
