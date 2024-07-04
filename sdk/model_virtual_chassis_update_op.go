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

// VirtualChassisUpdateOp the model 'VirtualChassisUpdateOp'
type VirtualChassisUpdateOp string

// List of virtual_chassis_update_op
const (
	VIRTUALCHASSISUPDATEOP_EMPTY VirtualChassisUpdateOp = ""
	VIRTUALCHASSISUPDATEOP_ADD VirtualChassisUpdateOp = "add"
	VIRTUALCHASSISUPDATEOP_REMOVE VirtualChassisUpdateOp = "remove"
	VIRTUALCHASSISUPDATEOP_RENUMBER VirtualChassisUpdateOp = "renumber"
	VIRTUALCHASSISUPDATEOP_PREPROVISION VirtualChassisUpdateOp = "preprovision"
)

// All allowed values of VirtualChassisUpdateOp enum
var AllowedVirtualChassisUpdateOpEnumValues = []VirtualChassisUpdateOp{
	"",
	"add",
	"remove",
	"renumber",
	"preprovision",
}

func (v *VirtualChassisUpdateOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualChassisUpdateOp(value)
	for _, existing := range AllowedVirtualChassisUpdateOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualChassisUpdateOp", value)
}

// NewVirtualChassisUpdateOpFromValue returns a pointer to a valid VirtualChassisUpdateOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualChassisUpdateOpFromValue(v string) (*VirtualChassisUpdateOp, error) {
	ev := VirtualChassisUpdateOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualChassisUpdateOp: valid values are %v", v, AllowedVirtualChassisUpdateOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualChassisUpdateOp) IsValid() bool {
	for _, existing := range AllowedVirtualChassisUpdateOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to virtual_chassis_update_op value
func (v VirtualChassisUpdateOp) Ptr() *VirtualChassisUpdateOp {
	return &v
}

type NullableVirtualChassisUpdateOp struct {
	value *VirtualChassisUpdateOp
	isSet bool
}

func (v NullableVirtualChassisUpdateOp) Get() *VirtualChassisUpdateOp {
	return v.value
}

func (v *NullableVirtualChassisUpdateOp) Set(val *VirtualChassisUpdateOp) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualChassisUpdateOp) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualChassisUpdateOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualChassisUpdateOp(val *VirtualChassisUpdateOp) *NullableVirtualChassisUpdateOp {
	return &NullableVirtualChassisUpdateOp{value: val, isSet: true}
}

func (v NullableVirtualChassisUpdateOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualChassisUpdateOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

