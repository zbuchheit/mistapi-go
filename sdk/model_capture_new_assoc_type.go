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

// CaptureNewAssocType new-assoc
type CaptureNewAssocType string

// List of capture_new_assoc_type
const (
	CAPTURENEWASSOCTYPE_EMPTY CaptureNewAssocType = ""
	CAPTURENEWASSOCTYPE_NEW_ASSOC CaptureNewAssocType = "new_assoc"
)

// All allowed values of CaptureNewAssocType enum
var AllowedCaptureNewAssocTypeEnumValues = []CaptureNewAssocType{
	"",
	"new_assoc",
}

func (v *CaptureNewAssocType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureNewAssocType(value)
	for _, existing := range AllowedCaptureNewAssocTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureNewAssocType", value)
}

// NewCaptureNewAssocTypeFromValue returns a pointer to a valid CaptureNewAssocType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureNewAssocTypeFromValue(v string) (*CaptureNewAssocType, error) {
	ev := CaptureNewAssocType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureNewAssocType: valid values are %v", v, AllowedCaptureNewAssocTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureNewAssocType) IsValid() bool {
	for _, existing := range AllowedCaptureNewAssocTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_new_assoc_type value
func (v CaptureNewAssocType) Ptr() *CaptureNewAssocType {
	return &v
}

type NullableCaptureNewAssocType struct {
	value *CaptureNewAssocType
	isSet bool
}

func (v NullableCaptureNewAssocType) Get() *CaptureNewAssocType {
	return v.value
}

func (v *NullableCaptureNewAssocType) Set(val *CaptureNewAssocType) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureNewAssocType) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureNewAssocType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureNewAssocType(val *CaptureNewAssocType) *NullableCaptureNewAssocType {
	return &NullableCaptureNewAssocType{value: val, isSet: true}
}

func (v NullableCaptureNewAssocType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureNewAssocType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

