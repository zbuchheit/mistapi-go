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

// BgpConfigType the model 'BgpConfigType'
type BgpConfigType string

// List of bgp_config_type
const (
	BGPCONFIGTYPE_EMPTY BgpConfigType = ""
	BGPCONFIGTYPE_INTERNAL BgpConfigType = "internal"
	BGPCONFIGTYPE_EXTERNAL BgpConfigType = "external"
)

// All allowed values of BgpConfigType enum
var AllowedBgpConfigTypeEnumValues = []BgpConfigType{
	"",
	"internal",
	"external",
}

func (v *BgpConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BgpConfigType(value)
	for _, existing := range AllowedBgpConfigTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BgpConfigType", value)
}

// NewBgpConfigTypeFromValue returns a pointer to a valid BgpConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBgpConfigTypeFromValue(v string) (*BgpConfigType, error) {
	ev := BgpConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BgpConfigType: valid values are %v", v, AllowedBgpConfigTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BgpConfigType) IsValid() bool {
	for _, existing := range AllowedBgpConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bgp_config_type value
func (v BgpConfigType) Ptr() *BgpConfigType {
	return &v
}

type NullableBgpConfigType struct {
	value *BgpConfigType
	isSet bool
}

func (v NullableBgpConfigType) Get() *BgpConfigType {
	return v.value
}

func (v *NullableBgpConfigType) Set(val *BgpConfigType) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfigType) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfigType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfigType(val *BgpConfigType) *NullableBgpConfigType {
	return &NullableBgpConfigType{value: val, isSet: true}
}

func (v NullableBgpConfigType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfigType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

