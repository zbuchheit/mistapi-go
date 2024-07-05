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

// Snmpv3ConfigTargetParamSecurityLevel the model 'Snmpv3ConfigTargetParamSecurityLevel'
type Snmpv3ConfigTargetParamSecurityLevel string

// List of snmpv3_config_target_param_security_level
const (
	SNMPV3CONFIGTARGETPARAMSECURITYLEVEL_EMPTY Snmpv3ConfigTargetParamSecurityLevel = ""
	SNMPV3CONFIGTARGETPARAMSECURITYLEVEL_AUTHENTICATION Snmpv3ConfigTargetParamSecurityLevel = "authentication"
	SNMPV3CONFIGTARGETPARAMSECURITYLEVEL_NONE Snmpv3ConfigTargetParamSecurityLevel = "none"
	SNMPV3CONFIGTARGETPARAMSECURITYLEVEL_PRIVACY Snmpv3ConfigTargetParamSecurityLevel = "privacy"
)

// All allowed values of Snmpv3ConfigTargetParamSecurityLevel enum
var AllowedSnmpv3ConfigTargetParamSecurityLevelEnumValues = []Snmpv3ConfigTargetParamSecurityLevel{
	"",
	"authentication",
	"none",
	"privacy",
}

func (v *Snmpv3ConfigTargetParamSecurityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Snmpv3ConfigTargetParamSecurityLevel(value)
	for _, existing := range AllowedSnmpv3ConfigTargetParamSecurityLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Snmpv3ConfigTargetParamSecurityLevel", value)
}

// NewSnmpv3ConfigTargetParamSecurityLevelFromValue returns a pointer to a valid Snmpv3ConfigTargetParamSecurityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnmpv3ConfigTargetParamSecurityLevelFromValue(v string) (*Snmpv3ConfigTargetParamSecurityLevel, error) {
	ev := Snmpv3ConfigTargetParamSecurityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Snmpv3ConfigTargetParamSecurityLevel: valid values are %v", v, AllowedSnmpv3ConfigTargetParamSecurityLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Snmpv3ConfigTargetParamSecurityLevel) IsValid() bool {
	for _, existing := range AllowedSnmpv3ConfigTargetParamSecurityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to snmpv3_config_target_param_security_level value
func (v Snmpv3ConfigTargetParamSecurityLevel) Ptr() *Snmpv3ConfigTargetParamSecurityLevel {
	return &v
}

type NullableSnmpv3ConfigTargetParamSecurityLevel struct {
	value *Snmpv3ConfigTargetParamSecurityLevel
	isSet bool
}

func (v NullableSnmpv3ConfigTargetParamSecurityLevel) Get() *Snmpv3ConfigTargetParamSecurityLevel {
	return v.value
}

func (v *NullableSnmpv3ConfigTargetParamSecurityLevel) Set(val *Snmpv3ConfigTargetParamSecurityLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpv3ConfigTargetParamSecurityLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpv3ConfigTargetParamSecurityLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpv3ConfigTargetParamSecurityLevel(val *Snmpv3ConfigTargetParamSecurityLevel) *NullableSnmpv3ConfigTargetParamSecurityLevel {
	return &NullableSnmpv3ConfigTargetParamSecurityLevel{value: val, isSet: true}
}

func (v NullableSnmpv3ConfigTargetParamSecurityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpv3ConfigTargetParamSecurityLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

