/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// SwitchRole the model 'SwitchRole'
type SwitchRole string

// List of switch_role
const (
	SWITCHROLE_EMPTY SwitchRole = ""
	SWITCHROLE_ACCESS SwitchRole = "access"
	SWITCHROLE_AGGREGATION SwitchRole = "aggregation"
)

// All allowed values of SwitchRole enum
var AllowedSwitchRoleEnumValues = []SwitchRole{
	"",
	"access",
	"aggregation",
}

func (v *SwitchRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SwitchRole(value)
	for _, existing := range AllowedSwitchRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SwitchRole", value)
}

// NewSwitchRoleFromValue returns a pointer to a valid SwitchRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSwitchRoleFromValue(v string) (*SwitchRole, error) {
	ev := SwitchRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SwitchRole: valid values are %v", v, AllowedSwitchRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SwitchRole) IsValid() bool {
	for _, existing := range AllowedSwitchRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to switch_role value
func (v SwitchRole) Ptr() *SwitchRole {
	return &v
}

type NullableSwitchRole struct {
	value *SwitchRole
	isSet bool
}

func (v NullableSwitchRole) Get() *SwitchRole {
	return v.value
}

func (v *NullableSwitchRole) Set(val *SwitchRole) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchRole) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchRole(val *SwitchRole) *NullableSwitchRole {
	return &NullableSwitchRole{value: val, isSet: true}
}

func (v NullableSwitchRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

