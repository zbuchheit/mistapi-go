/*
Mist API

> Version: **2406.1.17** > > Date: **July 5, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.17
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// RadioBandPreamble the model 'RadioBandPreamble'
type RadioBandPreamble string

// List of radio_band_preamble
const (
	RADIOBANDPREAMBLE_EMPTY RadioBandPreamble = ""
	RADIOBANDPREAMBLE_SHORT RadioBandPreamble = "short"
	RADIOBANDPREAMBLE_LONG RadioBandPreamble = "long"
	RADIOBANDPREAMBLE_AUTO RadioBandPreamble = "auto"
)

// All allowed values of RadioBandPreamble enum
var AllowedRadioBandPreambleEnumValues = []RadioBandPreamble{
	"",
	"short",
	"long",
	"auto",
}

func (v *RadioBandPreamble) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RadioBandPreamble(value)
	for _, existing := range AllowedRadioBandPreambleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RadioBandPreamble", value)
}

// NewRadioBandPreambleFromValue returns a pointer to a valid RadioBandPreamble
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRadioBandPreambleFromValue(v string) (*RadioBandPreamble, error) {
	ev := RadioBandPreamble(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RadioBandPreamble: valid values are %v", v, AllowedRadioBandPreambleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RadioBandPreamble) IsValid() bool {
	for _, existing := range AllowedRadioBandPreambleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to radio_band_preamble value
func (v RadioBandPreamble) Ptr() *RadioBandPreamble {
	return &v
}

type NullableRadioBandPreamble struct {
	value *RadioBandPreamble
	isSet bool
}

func (v NullableRadioBandPreamble) Get() *RadioBandPreamble {
	return v.value
}

func (v *NullableRadioBandPreamble) Set(val *RadioBandPreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioBandPreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioBandPreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioBandPreamble(val *RadioBandPreamble) *NullableRadioBandPreamble {
	return &NullableRadioBandPreamble{value: val, isSet: true}
}

func (v NullableRadioBandPreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioBandPreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

