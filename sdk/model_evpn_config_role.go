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

// EvpnConfigRole the model 'EvpnConfigRole'
type EvpnConfigRole string

// List of evpn_config_role
const (
	EVPNCONFIGROLE_EMPTY EvpnConfigRole = ""
	EVPNCONFIGROLE_CORE EvpnConfigRole = "core"
	EVPNCONFIGROLE_DISTRIBUTION EvpnConfigRole = "distribution"
	EVPNCONFIGROLE_ACCESS EvpnConfigRole = "access"
)

// All allowed values of EvpnConfigRole enum
var AllowedEvpnConfigRoleEnumValues = []EvpnConfigRole{
	"",
	"core",
	"distribution",
	"access",
}

func (v *EvpnConfigRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EvpnConfigRole(value)
	for _, existing := range AllowedEvpnConfigRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EvpnConfigRole", value)
}

// NewEvpnConfigRoleFromValue returns a pointer to a valid EvpnConfigRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEvpnConfigRoleFromValue(v string) (*EvpnConfigRole, error) {
	ev := EvpnConfigRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EvpnConfigRole: valid values are %v", v, AllowedEvpnConfigRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EvpnConfigRole) IsValid() bool {
	for _, existing := range AllowedEvpnConfigRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to evpn_config_role value
func (v EvpnConfigRole) Ptr() *EvpnConfigRole {
	return &v
}

type NullableEvpnConfigRole struct {
	value *EvpnConfigRole
	isSet bool
}

func (v NullableEvpnConfigRole) Get() *EvpnConfigRole {
	return v.value
}

func (v *NullableEvpnConfigRole) Set(val *EvpnConfigRole) {
	v.value = val
	v.isSet = true
}

func (v NullableEvpnConfigRole) IsSet() bool {
	return v.isSet
}

func (v *NullableEvpnConfigRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvpnConfigRole(val *EvpnConfigRole) *NullableEvpnConfigRole {
	return &NullableEvpnConfigRole{value: val, isSet: true}
}

func (v NullableEvpnConfigRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvpnConfigRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

