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

// GatewayPortWanType if `usage`==`wan`
type GatewayPortWanType string

// List of gateway_port_wan_type
const (
	GATEWAYPORTWANTYPE_EMPTY GatewayPortWanType = ""
	GATEWAYPORTWANTYPE_BROADBAND GatewayPortWanType = "broadband"
	GATEWAYPORTWANTYPE_DSL GatewayPortWanType = "dsl"
	GATEWAYPORTWANTYPE_LTE GatewayPortWanType = "lte"
)

// All allowed values of GatewayPortWanType enum
var AllowedGatewayPortWanTypeEnumValues = []GatewayPortWanType{
	"",
	"broadband",
	"dsl",
	"lte",
}

func (v *GatewayPortWanType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayPortWanType(value)
	for _, existing := range AllowedGatewayPortWanTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayPortWanType", value)
}

// NewGatewayPortWanTypeFromValue returns a pointer to a valid GatewayPortWanType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayPortWanTypeFromValue(v string) (*GatewayPortWanType, error) {
	ev := GatewayPortWanType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayPortWanType: valid values are %v", v, AllowedGatewayPortWanTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayPortWanType) IsValid() bool {
	for _, existing := range AllowedGatewayPortWanTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_port_wan_type value
func (v GatewayPortWanType) Ptr() *GatewayPortWanType {
	return &v
}

type NullableGatewayPortWanType struct {
	value *GatewayPortWanType
	isSet bool
}

func (v NullableGatewayPortWanType) Get() *GatewayPortWanType {
	return v.value
}

func (v *NullableGatewayPortWanType) Set(val *GatewayPortWanType) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortWanType) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortWanType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortWanType(val *GatewayPortWanType) *NullableGatewayPortWanType {
	return &NullableGatewayPortWanType{value: val, isSet: true}
}

func (v NullableGatewayPortWanType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortWanType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

