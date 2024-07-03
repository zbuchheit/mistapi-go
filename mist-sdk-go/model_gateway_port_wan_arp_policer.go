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

// GatewayPortWanArpPolicer when `wan_type`==`broadband`
type GatewayPortWanArpPolicer string

// List of gateway_port_wan_arp_policer
const (
	GATEWAYPORTWANARPPOLICER_EMPTY GatewayPortWanArpPolicer = ""
	GATEWAYPORTWANARPPOLICER_RECOMMENDED GatewayPortWanArpPolicer = "recommended"
	GATEWAYPORTWANARPPOLICER_DEFAULT GatewayPortWanArpPolicer = "default"
	GATEWAYPORTWANARPPOLICER_MAX GatewayPortWanArpPolicer = "max"
)

// All allowed values of GatewayPortWanArpPolicer enum
var AllowedGatewayPortWanArpPolicerEnumValues = []GatewayPortWanArpPolicer{
	"",
	"recommended",
	"default",
	"max",
}

func (v *GatewayPortWanArpPolicer) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayPortWanArpPolicer(value)
	for _, existing := range AllowedGatewayPortWanArpPolicerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayPortWanArpPolicer", value)
}

// NewGatewayPortWanArpPolicerFromValue returns a pointer to a valid GatewayPortWanArpPolicer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayPortWanArpPolicerFromValue(v string) (*GatewayPortWanArpPolicer, error) {
	ev := GatewayPortWanArpPolicer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayPortWanArpPolicer: valid values are %v", v, AllowedGatewayPortWanArpPolicerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayPortWanArpPolicer) IsValid() bool {
	for _, existing := range AllowedGatewayPortWanArpPolicerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_port_wan_arp_policer value
func (v GatewayPortWanArpPolicer) Ptr() *GatewayPortWanArpPolicer {
	return &v
}

type NullableGatewayPortWanArpPolicer struct {
	value *GatewayPortWanArpPolicer
	isSet bool
}

func (v NullableGatewayPortWanArpPolicer) Get() *GatewayPortWanArpPolicer {
	return v.value
}

func (v *NullableGatewayPortWanArpPolicer) Set(val *GatewayPortWanArpPolicer) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPortWanArpPolicer) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPortWanArpPolicer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPortWanArpPolicer(val *GatewayPortWanArpPolicer) *NullableGatewayPortWanArpPolicer {
	return &NullableGatewayPortWanArpPolicer{value: val, isSet: true}
}

func (v NullableGatewayPortWanArpPolicer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPortWanArpPolicer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

