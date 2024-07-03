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

// GatewayPortUsage port usage name
type GatewayPortUsage string

// List of gateway_port_usage
const (
	GATEWAYPORTUSAGE_EMPTY GatewayPortUsage = ""
	GATEWAYPORTUSAGE_LAN GatewayPortUsage = "lan"
	GATEWAYPORTUSAGE_WAN GatewayPortUsage = "wan"
	GATEWAYPORTUSAGE_HA_DATA GatewayPortUsage = "ha_data"
	GATEWAYPORTUSAGE_HA_CONTROL GatewayPortUsage = "ha_control"
)

// All allowed values of GatewayPortUsage enum
var AllowedGatewayPortUsageEnumValues = []GatewayPortUsage{
	"",
	"lan",
	"wan",
	"ha_data",
	"ha_control",
}

func (v *GatewayPortUsage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayPortUsage(value)
	for _, existing := range AllowedGatewayPortUsageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayPortUsage", value)
}

// NewGatewayPortUsageFromValue returns a pointer to a valid GatewayPortUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayPortUsageFromValue(v string) (*GatewayPortUsage, error) {
	ev := GatewayPortUsage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayPortUsage: valid values are %v", v, AllowedGatewayPortUsageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayPortUsage) IsValid() bool {
	for _, existing := range AllowedGatewayPortUsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_port_usage value
func (v GatewayPortUsage) Ptr() *GatewayPortUsage {
	return &v
}

type NullableGatewayPortUsage struct {
	value *GatewayPortUsage
	isSet bool
}

func (v NullableGatewayPortUsage) Get() *GatewayPortUsage {
	return v.value
}

func (v *NullableGatewayPortUsage) Set(val *GatewayPortUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortUsage(val *GatewayPortUsage) *NullableGatewayPortUsage {
	return &NullableGatewayPortUsage{value: val, isSet: true}
}

func (v NullableGatewayPortUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

