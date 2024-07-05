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

// OspfAreasNetworkAuthType auth type
type OspfAreasNetworkAuthType string

// List of ospf_areas_network_auth_type
const (
	OSPFAREASNETWORKAUTHTYPE_EMPTY OspfAreasNetworkAuthType = ""
	OSPFAREASNETWORKAUTHTYPE_NONE OspfAreasNetworkAuthType = "none"
	OSPFAREASNETWORKAUTHTYPE_MD5 OspfAreasNetworkAuthType = "md5"
	OSPFAREASNETWORKAUTHTYPE_PASSWORD OspfAreasNetworkAuthType = "password"
)

// All allowed values of OspfAreasNetworkAuthType enum
var AllowedOspfAreasNetworkAuthTypeEnumValues = []OspfAreasNetworkAuthType{
	"",
	"none",
	"md5",
	"password",
}

func (v *OspfAreasNetworkAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OspfAreasNetworkAuthType(value)
	for _, existing := range AllowedOspfAreasNetworkAuthTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OspfAreasNetworkAuthType", value)
}

// NewOspfAreasNetworkAuthTypeFromValue returns a pointer to a valid OspfAreasNetworkAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOspfAreasNetworkAuthTypeFromValue(v string) (*OspfAreasNetworkAuthType, error) {
	ev := OspfAreasNetworkAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OspfAreasNetworkAuthType: valid values are %v", v, AllowedOspfAreasNetworkAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OspfAreasNetworkAuthType) IsValid() bool {
	for _, existing := range AllowedOspfAreasNetworkAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ospf_areas_network_auth_type value
func (v OspfAreasNetworkAuthType) Ptr() *OspfAreasNetworkAuthType {
	return &v
}

type NullableOspfAreasNetworkAuthType struct {
	value *OspfAreasNetworkAuthType
	isSet bool
}

func (v NullableOspfAreasNetworkAuthType) Get() *OspfAreasNetworkAuthType {
	return v.value
}

func (v *NullableOspfAreasNetworkAuthType) Set(val *OspfAreasNetworkAuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableOspfAreasNetworkAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableOspfAreasNetworkAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOspfAreasNetworkAuthType(val *OspfAreasNetworkAuthType) *NullableOspfAreasNetworkAuthType {
	return &NullableOspfAreasNetworkAuthType{value: val, isSet: true}
}

func (v NullableOspfAreasNetworkAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOspfAreasNetworkAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

