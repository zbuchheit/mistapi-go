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

// VrrpGroupAuthType the model 'VrrpGroupAuthType'
type VrrpGroupAuthType string

// List of vrrp_group_auth_type
const (
	VRRPGROUPAUTHTYPE_EMPTY VrrpGroupAuthType = ""
	VRRPGROUPAUTHTYPE_MD5 VrrpGroupAuthType = "md5"
	VRRPGROUPAUTHTYPE_SIMPLE VrrpGroupAuthType = "simple"
)

// All allowed values of VrrpGroupAuthType enum
var AllowedVrrpGroupAuthTypeEnumValues = []VrrpGroupAuthType{
	"",
	"md5",
	"simple",
}

func (v *VrrpGroupAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VrrpGroupAuthType(value)
	for _, existing := range AllowedVrrpGroupAuthTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VrrpGroupAuthType", value)
}

// NewVrrpGroupAuthTypeFromValue returns a pointer to a valid VrrpGroupAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVrrpGroupAuthTypeFromValue(v string) (*VrrpGroupAuthType, error) {
	ev := VrrpGroupAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VrrpGroupAuthType: valid values are %v", v, AllowedVrrpGroupAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VrrpGroupAuthType) IsValid() bool {
	for _, existing := range AllowedVrrpGroupAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vrrp_group_auth_type value
func (v VrrpGroupAuthType) Ptr() *VrrpGroupAuthType {
	return &v
}

type NullableVrrpGroupAuthType struct {
	value *VrrpGroupAuthType
	isSet bool
}

func (v NullableVrrpGroupAuthType) Get() *VrrpGroupAuthType {
	return v.value
}

func (v *NullableVrrpGroupAuthType) Set(val *VrrpGroupAuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableVrrpGroupAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableVrrpGroupAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrrpGroupAuthType(val *VrrpGroupAuthType) *NullableVrrpGroupAuthType {
	return &NullableVrrpGroupAuthType{value: val, isSet: true}
}

func (v NullableVrrpGroupAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrrpGroupAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

