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

// GatewayTemplateTunnelMode the model 'GatewayTemplateTunnelMode'
type GatewayTemplateTunnelMode string

// List of gateway_template_tunnel_mode
const (
	GATEWAYTEMPLATETUNNELMODE_EMPTY GatewayTemplateTunnelMode = ""
	GATEWAYTEMPLATETUNNELMODE_ACTIVE_STANDBY GatewayTemplateTunnelMode = "active-standby"
	GATEWAYTEMPLATETUNNELMODE_ACTIVE_ACTIVE GatewayTemplateTunnelMode = "active-active"
)

// All allowed values of GatewayTemplateTunnelMode enum
var AllowedGatewayTemplateTunnelModeEnumValues = []GatewayTemplateTunnelMode{
	"",
	"active-standby",
	"active-active",
}

func (v *GatewayTemplateTunnelMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayTemplateTunnelMode(value)
	for _, existing := range AllowedGatewayTemplateTunnelModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayTemplateTunnelMode", value)
}

// NewGatewayTemplateTunnelModeFromValue returns a pointer to a valid GatewayTemplateTunnelMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayTemplateTunnelModeFromValue(v string) (*GatewayTemplateTunnelMode, error) {
	ev := GatewayTemplateTunnelMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayTemplateTunnelMode: valid values are %v", v, AllowedGatewayTemplateTunnelModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayTemplateTunnelMode) IsValid() bool {
	for _, existing := range AllowedGatewayTemplateTunnelModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_template_tunnel_mode value
func (v GatewayTemplateTunnelMode) Ptr() *GatewayTemplateTunnelMode {
	return &v
}

type NullableGatewayTemplateTunnelMode struct {
	value *GatewayTemplateTunnelMode
	isSet bool
}

func (v NullableGatewayTemplateTunnelMode) Get() *GatewayTemplateTunnelMode {
	return v.value
}

func (v *NullableGatewayTemplateTunnelMode) Set(val *GatewayTemplateTunnelMode) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateTunnelMode) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateTunnelMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateTunnelMode(val *GatewayTemplateTunnelMode) *NullableGatewayTemplateTunnelMode {
	return &NullableGatewayTemplateTunnelMode{value: val, isSet: true}
}

func (v NullableGatewayTemplateTunnelMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateTunnelMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

