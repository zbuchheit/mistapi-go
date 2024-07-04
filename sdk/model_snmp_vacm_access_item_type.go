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

// SnmpVacmAccessItemType the model 'SnmpVacmAccessItemType'
type SnmpVacmAccessItemType string

// List of snmp_vacm_access_item_type
const (
	SNMPVACMACCESSITEMTYPE_EMPTY SnmpVacmAccessItemType = ""
	SNMPVACMACCESSITEMTYPE_DEFAULT_CONTEXT_PREFIX SnmpVacmAccessItemType = "default_context_prefix"
)

// All allowed values of SnmpVacmAccessItemType enum
var AllowedSnmpVacmAccessItemTypeEnumValues = []SnmpVacmAccessItemType{
	"",
	"default_context_prefix",
}

func (v *SnmpVacmAccessItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SnmpVacmAccessItemType(value)
	for _, existing := range AllowedSnmpVacmAccessItemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SnmpVacmAccessItemType", value)
}

// NewSnmpVacmAccessItemTypeFromValue returns a pointer to a valid SnmpVacmAccessItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnmpVacmAccessItemTypeFromValue(v string) (*SnmpVacmAccessItemType, error) {
	ev := SnmpVacmAccessItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SnmpVacmAccessItemType: valid values are %v", v, AllowedSnmpVacmAccessItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SnmpVacmAccessItemType) IsValid() bool {
	for _, existing := range AllowedSnmpVacmAccessItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to snmp_vacm_access_item_type value
func (v SnmpVacmAccessItemType) Ptr() *SnmpVacmAccessItemType {
	return &v
}

type NullableSnmpVacmAccessItemType struct {
	value *SnmpVacmAccessItemType
	isSet bool
}

func (v NullableSnmpVacmAccessItemType) Get() *SnmpVacmAccessItemType {
	return v.value
}

func (v *NullableSnmpVacmAccessItemType) Set(val *SnmpVacmAccessItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpVacmAccessItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpVacmAccessItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpVacmAccessItemType(val *SnmpVacmAccessItemType) *NullableSnmpVacmAccessItemType {
	return &NullableSnmpVacmAccessItemType{value: val, isSet: true}
}

func (v NullableSnmpVacmAccessItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpVacmAccessItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

