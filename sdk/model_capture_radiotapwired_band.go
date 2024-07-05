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

// CaptureRadiotapwiredBand only used for radiotap
type CaptureRadiotapwiredBand string

// List of capture_radiotapwired_band
const (
	CAPTURERADIOTAPWIREDBAND_EMPTY CaptureRadiotapwiredBand = ""
	CAPTURERADIOTAPWIREDBAND__24 CaptureRadiotapwiredBand = "24"
	CAPTURERADIOTAPWIREDBAND__5 CaptureRadiotapwiredBand = "5"
	CAPTURERADIOTAPWIREDBAND__6 CaptureRadiotapwiredBand = "6"
	CAPTURERADIOTAPWIREDBAND__2456 CaptureRadiotapwiredBand = "24,5,6"
)

// All allowed values of CaptureRadiotapwiredBand enum
var AllowedCaptureRadiotapwiredBandEnumValues = []CaptureRadiotapwiredBand{
	"",
	"24",
	"5",
	"6",
	"24,5,6",
}

func (v *CaptureRadiotapwiredBand) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureRadiotapwiredBand(value)
	for _, existing := range AllowedCaptureRadiotapwiredBandEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureRadiotapwiredBand", value)
}

// NewCaptureRadiotapwiredBandFromValue returns a pointer to a valid CaptureRadiotapwiredBand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureRadiotapwiredBandFromValue(v string) (*CaptureRadiotapwiredBand, error) {
	ev := CaptureRadiotapwiredBand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureRadiotapwiredBand: valid values are %v", v, AllowedCaptureRadiotapwiredBandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureRadiotapwiredBand) IsValid() bool {
	for _, existing := range AllowedCaptureRadiotapwiredBandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_radiotapwired_band value
func (v CaptureRadiotapwiredBand) Ptr() *CaptureRadiotapwiredBand {
	return &v
}

type NullableCaptureRadiotapwiredBand struct {
	value *CaptureRadiotapwiredBand
	isSet bool
}

func (v NullableCaptureRadiotapwiredBand) Get() *CaptureRadiotapwiredBand {
	return v.value
}

func (v *NullableCaptureRadiotapwiredBand) Set(val *CaptureRadiotapwiredBand) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureRadiotapwiredBand) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureRadiotapwiredBand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureRadiotapwiredBand(val *CaptureRadiotapwiredBand) *NullableCaptureRadiotapwiredBand {
	return &NullableCaptureRadiotapwiredBand{value: val, isSet: true}
}

func (v NullableCaptureRadiotapwiredBand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureRadiotapwiredBand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

