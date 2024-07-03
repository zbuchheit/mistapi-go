/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"encoding/json"
	"fmt"
)

// RadioBand24Usage the model 'RadioBand24Usage'
type RadioBand24Usage string

// List of radio_band_24_usage
const (
	RADIOBAND24USAGE_EMPTY RadioBand24Usage = ""
	RADIOBAND24USAGE__24 RadioBand24Usage = "24"
	RADIOBAND24USAGE__5 RadioBand24Usage = "5"
	RADIOBAND24USAGE__6 RadioBand24Usage = "6"
	RADIOBAND24USAGE_AUTO RadioBand24Usage = "auto"
)

// All allowed values of RadioBand24Usage enum
var AllowedRadioBand24UsageEnumValues = []RadioBand24Usage{
	"",
	"24",
	"5",
	"6",
	"auto",
}

func (v *RadioBand24Usage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RadioBand24Usage(value)
	for _, existing := range AllowedRadioBand24UsageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RadioBand24Usage", value)
}

// NewRadioBand24UsageFromValue returns a pointer to a valid RadioBand24Usage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRadioBand24UsageFromValue(v string) (*RadioBand24Usage, error) {
	ev := RadioBand24Usage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RadioBand24Usage: valid values are %v", v, AllowedRadioBand24UsageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RadioBand24Usage) IsValid() bool {
	for _, existing := range AllowedRadioBand24UsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to radio_band_24_usage value
func (v RadioBand24Usage) Ptr() *RadioBand24Usage {
	return &v
}

type NullableRadioBand24Usage struct {
	value *RadioBand24Usage
	isSet bool
}

func (v NullableRadioBand24Usage) Get() *RadioBand24Usage {
	return v.value
}

func (v *NullableRadioBand24Usage) Set(val *RadioBand24Usage) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioBand24Usage) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioBand24Usage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioBand24Usage(val *RadioBand24Usage) *NullableRadioBand24Usage {
	return &NullableRadioBand24Usage{value: val, isSet: true}
}

func (v NullableRadioBand24Usage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioBand24Usage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

