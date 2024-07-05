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

// NacPortalAccessType the model 'NacPortalAccessType'
type NacPortalAccessType string

// List of nac_portal_access_type
const (
	NACPORTALACCESSTYPE_EMPTY NacPortalAccessType = ""
	NACPORTALACCESSTYPE_WIRELESS NacPortalAccessType = "wireless"
	NACPORTALACCESSTYPE_WIRELESSWIRED NacPortalAccessType = "wireless+wired"
)

// All allowed values of NacPortalAccessType enum
var AllowedNacPortalAccessTypeEnumValues = []NacPortalAccessType{
	"",
	"wireless",
	"wireless+wired",
}

func (v *NacPortalAccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NacPortalAccessType(value)
	for _, existing := range AllowedNacPortalAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NacPortalAccessType", value)
}

// NewNacPortalAccessTypeFromValue returns a pointer to a valid NacPortalAccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNacPortalAccessTypeFromValue(v string) (*NacPortalAccessType, error) {
	ev := NacPortalAccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NacPortalAccessType: valid values are %v", v, AllowedNacPortalAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NacPortalAccessType) IsValid() bool {
	for _, existing := range AllowedNacPortalAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to nac_portal_access_type value
func (v NacPortalAccessType) Ptr() *NacPortalAccessType {
	return &v
}

type NullableNacPortalAccessType struct {
	value *NacPortalAccessType
	isSet bool
}

func (v NullableNacPortalAccessType) Get() *NacPortalAccessType {
	return v.value
}

func (v *NullableNacPortalAccessType) Set(val *NacPortalAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableNacPortalAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableNacPortalAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNacPortalAccessType(val *NacPortalAccessType) *NullableNacPortalAccessType {
	return &NullableNacPortalAccessType{value: val, isSet: true}
}

func (v NullableNacPortalAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNacPortalAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

