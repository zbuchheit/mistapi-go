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

// CaptureRadiotapwiredType radiotap,wired
type CaptureRadiotapwiredType string

// List of capture_radiotapwired_type
const (
	CAPTURERADIOTAPWIREDTYPE_EMPTY CaptureRadiotapwiredType = ""
	CAPTURERADIOTAPWIREDTYPE_RADIOTAPWIRED CaptureRadiotapwiredType = "radiotap,wired"
)

// All allowed values of CaptureRadiotapwiredType enum
var AllowedCaptureRadiotapwiredTypeEnumValues = []CaptureRadiotapwiredType{
	"",
	"radiotap,wired",
}

func (v *CaptureRadiotapwiredType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureRadiotapwiredType(value)
	for _, existing := range AllowedCaptureRadiotapwiredTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureRadiotapwiredType", value)
}

// NewCaptureRadiotapwiredTypeFromValue returns a pointer to a valid CaptureRadiotapwiredType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureRadiotapwiredTypeFromValue(v string) (*CaptureRadiotapwiredType, error) {
	ev := CaptureRadiotapwiredType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureRadiotapwiredType: valid values are %v", v, AllowedCaptureRadiotapwiredTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureRadiotapwiredType) IsValid() bool {
	for _, existing := range AllowedCaptureRadiotapwiredTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_radiotapwired_type value
func (v CaptureRadiotapwiredType) Ptr() *CaptureRadiotapwiredType {
	return &v
}

type NullableCaptureRadiotapwiredType struct {
	value *CaptureRadiotapwiredType
	isSet bool
}

func (v NullableCaptureRadiotapwiredType) Get() *CaptureRadiotapwiredType {
	return v.value
}

func (v *NullableCaptureRadiotapwiredType) Set(val *CaptureRadiotapwiredType) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureRadiotapwiredType) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureRadiotapwiredType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureRadiotapwiredType(val *CaptureRadiotapwiredType) *NullableCaptureRadiotapwiredType {
	return &NullableCaptureRadiotapwiredType{value: val, isSet: true}
}

func (v NullableCaptureRadiotapwiredType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureRadiotapwiredType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

