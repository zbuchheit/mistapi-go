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

// CaptureSwitchFormat the model 'CaptureSwitchFormat'
type CaptureSwitchFormat string

// List of capture_switch_format
const (
	CAPTURESWITCHFORMAT_EMPTY CaptureSwitchFormat = ""
	CAPTURESWITCHFORMAT_STREAM CaptureSwitchFormat = "stream"
)

// All allowed values of CaptureSwitchFormat enum
var AllowedCaptureSwitchFormatEnumValues = []CaptureSwitchFormat{
	"",
	"stream",
}

func (v *CaptureSwitchFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureSwitchFormat(value)
	for _, existing := range AllowedCaptureSwitchFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureSwitchFormat", value)
}

// NewCaptureSwitchFormatFromValue returns a pointer to a valid CaptureSwitchFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureSwitchFormatFromValue(v string) (*CaptureSwitchFormat, error) {
	ev := CaptureSwitchFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureSwitchFormat: valid values are %v", v, AllowedCaptureSwitchFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureSwitchFormat) IsValid() bool {
	for _, existing := range AllowedCaptureSwitchFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_switch_format value
func (v CaptureSwitchFormat) Ptr() *CaptureSwitchFormat {
	return &v
}

type NullableCaptureSwitchFormat struct {
	value *CaptureSwitchFormat
	isSet bool
}

func (v NullableCaptureSwitchFormat) Get() *CaptureSwitchFormat {
	return v.value
}

func (v *NullableCaptureSwitchFormat) Set(val *CaptureSwitchFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureSwitchFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureSwitchFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureSwitchFormat(val *CaptureSwitchFormat) *NullableCaptureSwitchFormat {
	return &NullableCaptureSwitchFormat{value: val, isSet: true}
}

func (v NullableCaptureSwitchFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureSwitchFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

