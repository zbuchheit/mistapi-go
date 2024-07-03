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

// GatewayPathStrategy the model 'GatewayPathStrategy'
type GatewayPathStrategy string

// List of gateway_path_strategy
const (
	GATEWAYPATHSTRATEGY_EMPTY GatewayPathStrategy = ""
	GATEWAYPATHSTRATEGY_ORDERED GatewayPathStrategy = "ordered"
	GATEWAYPATHSTRATEGY_WEIGHTED GatewayPathStrategy = "weighted"
	GATEWAYPATHSTRATEGY_ECMP GatewayPathStrategy = "ecmp"
)

// All allowed values of GatewayPathStrategy enum
var AllowedGatewayPathStrategyEnumValues = []GatewayPathStrategy{
	"",
	"ordered",
	"weighted",
	"ecmp",
}

func (v *GatewayPathStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayPathStrategy(value)
	for _, existing := range AllowedGatewayPathStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayPathStrategy", value)
}

// NewGatewayPathStrategyFromValue returns a pointer to a valid GatewayPathStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayPathStrategyFromValue(v string) (*GatewayPathStrategy, error) {
	ev := GatewayPathStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayPathStrategy: valid values are %v", v, AllowedGatewayPathStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayPathStrategy) IsValid() bool {
	for _, existing := range AllowedGatewayPathStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_path_strategy value
func (v GatewayPathStrategy) Ptr() *GatewayPathStrategy {
	return &v
}

type NullableGatewayPathStrategy struct {
	value *GatewayPathStrategy
	isSet bool
}

func (v NullableGatewayPathStrategy) Get() *GatewayPathStrategy {
	return v.value
}

func (v *NullableGatewayPathStrategy) Set(val *GatewayPathStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayPathStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayPathStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayPathStrategy(val *GatewayPathStrategy) *NullableGatewayPathStrategy {
	return &NullableGatewayPathStrategy{value: val, isSet: true}
}

func (v NullableGatewayPathStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayPathStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

