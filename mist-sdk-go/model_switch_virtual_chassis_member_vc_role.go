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

// SwitchVirtualChassisMemberVcRole Both vc_role master and backup will be matched to routing-engine role in Junos preprovisioned VC config
type SwitchVirtualChassisMemberVcRole string

// List of switch_virtual_chassis_member_vc_role
const (
	SWITCHVIRTUALCHASSISMEMBERVCROLE_EMPTY SwitchVirtualChassisMemberVcRole = ""
	SWITCHVIRTUALCHASSISMEMBERVCROLE_MASTER SwitchVirtualChassisMemberVcRole = "master"
	SWITCHVIRTUALCHASSISMEMBERVCROLE_LINECARD SwitchVirtualChassisMemberVcRole = "linecard"
	SWITCHVIRTUALCHASSISMEMBERVCROLE_BACKUP SwitchVirtualChassisMemberVcRole = "backup"
)

// All allowed values of SwitchVirtualChassisMemberVcRole enum
var AllowedSwitchVirtualChassisMemberVcRoleEnumValues = []SwitchVirtualChassisMemberVcRole{
	"",
	"master",
	"linecard",
	"backup",
}

func (v *SwitchVirtualChassisMemberVcRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SwitchVirtualChassisMemberVcRole(value)
	for _, existing := range AllowedSwitchVirtualChassisMemberVcRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SwitchVirtualChassisMemberVcRole", value)
}

// NewSwitchVirtualChassisMemberVcRoleFromValue returns a pointer to a valid SwitchVirtualChassisMemberVcRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSwitchVirtualChassisMemberVcRoleFromValue(v string) (*SwitchVirtualChassisMemberVcRole, error) {
	ev := SwitchVirtualChassisMemberVcRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SwitchVirtualChassisMemberVcRole: valid values are %v", v, AllowedSwitchVirtualChassisMemberVcRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SwitchVirtualChassisMemberVcRole) IsValid() bool {
	for _, existing := range AllowedSwitchVirtualChassisMemberVcRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to switch_virtual_chassis_member_vc_role value
func (v SwitchVirtualChassisMemberVcRole) Ptr() *SwitchVirtualChassisMemberVcRole {
	return &v
}

type NullableSwitchVirtualChassisMemberVcRole struct {
	value *SwitchVirtualChassisMemberVcRole
	isSet bool
}

func (v NullableSwitchVirtualChassisMemberVcRole) Get() *SwitchVirtualChassisMemberVcRole {
	return v.value
}

func (v *NullableSwitchVirtualChassisMemberVcRole) Set(val *SwitchVirtualChassisMemberVcRole) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchVirtualChassisMemberVcRole) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchVirtualChassisMemberVcRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchVirtualChassisMemberVcRole(val *SwitchVirtualChassisMemberVcRole) *NullableSwitchVirtualChassisMemberVcRole {
	return &NullableSwitchVirtualChassisMemberVcRole{value: val, isSet: true}
}

func (v NullableSwitchVirtualChassisMemberVcRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchVirtualChassisMemberVcRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

