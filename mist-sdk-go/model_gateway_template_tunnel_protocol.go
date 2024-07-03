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

// GatewayTemplateTunnelProtocol Only if: * `provider`== `custom-ipsec`
type GatewayTemplateTunnelProtocol string

// List of gateway_template_tunnel_protocol
const (
	GATEWAYTEMPLATETUNNELPROTOCOL_EMPTY GatewayTemplateTunnelProtocol = ""
	GATEWAYTEMPLATETUNNELPROTOCOL_IPSEC GatewayTemplateTunnelProtocol = "ipsec"
	GATEWAYTEMPLATETUNNELPROTOCOL_GRE GatewayTemplateTunnelProtocol = "gre"
)

// All allowed values of GatewayTemplateTunnelProtocol enum
var AllowedGatewayTemplateTunnelProtocolEnumValues = []GatewayTemplateTunnelProtocol{
	"",
	"ipsec",
	"gre",
}

func (v *GatewayTemplateTunnelProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayTemplateTunnelProtocol(value)
	for _, existing := range AllowedGatewayTemplateTunnelProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayTemplateTunnelProtocol", value)
}

// NewGatewayTemplateTunnelProtocolFromValue returns a pointer to a valid GatewayTemplateTunnelProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayTemplateTunnelProtocolFromValue(v string) (*GatewayTemplateTunnelProtocol, error) {
	ev := GatewayTemplateTunnelProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayTemplateTunnelProtocol: valid values are %v", v, AllowedGatewayTemplateTunnelProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayTemplateTunnelProtocol) IsValid() bool {
	for _, existing := range AllowedGatewayTemplateTunnelProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_template_tunnel_protocol value
func (v GatewayTemplateTunnelProtocol) Ptr() *GatewayTemplateTunnelProtocol {
	return &v
}

type NullableGatewayTemplateTunnelProtocol struct {
	value *GatewayTemplateTunnelProtocol
	isSet bool
}

func (v NullableGatewayTemplateTunnelProtocol) Get() *GatewayTemplateTunnelProtocol {
	return v.value
}

func (v *NullableGatewayTemplateTunnelProtocol) Set(val *GatewayTemplateTunnelProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateTunnelProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateTunnelProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateTunnelProtocol(val *GatewayTemplateTunnelProtocol) *NullableGatewayTemplateTunnelProtocol {
	return &NullableGatewayTemplateTunnelProtocol{value: val, isSet: true}
}

func (v NullableGatewayTemplateTunnelProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateTunnelProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

