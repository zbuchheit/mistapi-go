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

// MspTier the model 'MspTier'
type MspTier string

// List of msp_tier
const (
	MSPTIER_EMPTY MspTier = ""
	MSPTIER_BASE MspTier = "base"
	MSPTIER_ADVANCED MspTier = "advanced"
)

// All allowed values of MspTier enum
var AllowedMspTierEnumValues = []MspTier{
	"",
	"base",
	"advanced",
}

func (v *MspTier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MspTier(value)
	for _, existing := range AllowedMspTierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MspTier", value)
}

// NewMspTierFromValue returns a pointer to a valid MspTier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMspTierFromValue(v string) (*MspTier, error) {
	ev := MspTier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MspTier: valid values are %v", v, AllowedMspTierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MspTier) IsValid() bool {
	for _, existing := range AllowedMspTierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to msp_tier value
func (v MspTier) Ptr() *MspTier {
	return &v
}

type NullableMspTier struct {
	value *MspTier
	isSet bool
}

func (v NullableMspTier) Get() *MspTier {
	return v.value
}

func (v *NullableMspTier) Set(val *MspTier) {
	v.value = val
	v.isSet = true
}

func (v NullableMspTier) IsSet() bool {
	return v.isSet
}

func (v *NullableMspTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMspTier(val *MspTier) *NullableMspTier {
	return &NullableMspTier{value: val, isSet: true}
}

func (v NullableMspTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMspTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

