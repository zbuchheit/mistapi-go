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

// UseAutoApValuesFor The selector to choose auto placement or auto orientation.
type UseAutoApValuesFor string

// List of use_auto_ap_values_for
const (
	USEAUTOAPVALUESFOR_EMPTY UseAutoApValuesFor = ""
	USEAUTOAPVALUESFOR_PLACEMENT UseAutoApValuesFor = "placement"
	USEAUTOAPVALUESFOR_ORIENTATION UseAutoApValuesFor = "orientation"
)

// All allowed values of UseAutoApValuesFor enum
var AllowedUseAutoApValuesForEnumValues = []UseAutoApValuesFor{
	"",
	"placement",
	"orientation",
}

func (v *UseAutoApValuesFor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UseAutoApValuesFor(value)
	for _, existing := range AllowedUseAutoApValuesForEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UseAutoApValuesFor", value)
}

// NewUseAutoApValuesForFromValue returns a pointer to a valid UseAutoApValuesFor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUseAutoApValuesForFromValue(v string) (*UseAutoApValuesFor, error) {
	ev := UseAutoApValuesFor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UseAutoApValuesFor: valid values are %v", v, AllowedUseAutoApValuesForEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UseAutoApValuesFor) IsValid() bool {
	for _, existing := range AllowedUseAutoApValuesForEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to use_auto_ap_values_for value
func (v UseAutoApValuesFor) Ptr() *UseAutoApValuesFor {
	return &v
}

type NullableUseAutoApValuesFor struct {
	value *UseAutoApValuesFor
	isSet bool
}

func (v NullableUseAutoApValuesFor) Get() *UseAutoApValuesFor {
	return v.value
}

func (v *NullableUseAutoApValuesFor) Set(val *UseAutoApValuesFor) {
	v.value = val
	v.isSet = true
}

func (v NullableUseAutoApValuesFor) IsSet() bool {
	return v.isSet
}

func (v *NullableUseAutoApValuesFor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUseAutoApValuesFor(val *UseAutoApValuesFor) *NullableUseAutoApValuesFor {
	return &NullableUseAutoApValuesFor{value: val, isSet: true}
}

func (v NullableUseAutoApValuesFor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUseAutoApValuesFor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

