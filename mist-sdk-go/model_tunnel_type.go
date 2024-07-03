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

// TunnelType the model 'TunnelType'
type TunnelType string

// List of tunnel_type
const (
	TUNNELTYPE_EMPTY TunnelType = ""
	TUNNELTYPE_WXTUNNEL TunnelType = "wxtunnel"
	TUNNELTYPE_WAN TunnelType = "wan"
)

// All allowed values of TunnelType enum
var AllowedTunnelTypeEnumValues = []TunnelType{
	"",
	"wxtunnel",
	"wan",
}

func (v *TunnelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TunnelType(value)
	for _, existing := range AllowedTunnelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TunnelType", value)
}

// NewTunnelTypeFromValue returns a pointer to a valid TunnelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTunnelTypeFromValue(v string) (*TunnelType, error) {
	ev := TunnelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TunnelType: valid values are %v", v, AllowedTunnelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TunnelType) IsValid() bool {
	for _, existing := range AllowedTunnelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tunnel_type value
func (v TunnelType) Ptr() *TunnelType {
	return &v
}

type NullableTunnelType struct {
	value *TunnelType
	isSet bool
}

func (v NullableTunnelType) Get() *TunnelType {
	return v.value
}

func (v *NullableTunnelType) Set(val *TunnelType) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelType) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelType(val *TunnelType) *NullableTunnelType {
	return &NullableTunnelType{value: val, isSet: true}
}

func (v NullableTunnelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

