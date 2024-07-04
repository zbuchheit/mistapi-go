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

// GatewayTemplateTunnelVersion Only if: * `provider`== `custom-gre`  * `provider`== `custom-ipsec`
type GatewayTemplateTunnelVersion string

// List of gateway_template_tunnel_version
const (
	GATEWAYTEMPLATETUNNELVERSION_EMPTY GatewayTemplateTunnelVersion = ""
	GATEWAYTEMPLATETUNNELVERSION__1 GatewayTemplateTunnelVersion = "1"
	GATEWAYTEMPLATETUNNELVERSION__2 GatewayTemplateTunnelVersion = "2"
)

// All allowed values of GatewayTemplateTunnelVersion enum
var AllowedGatewayTemplateTunnelVersionEnumValues = []GatewayTemplateTunnelVersion{
	"",
	"1",
	"2",
}

func (v *GatewayTemplateTunnelVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GatewayTemplateTunnelVersion(value)
	for _, existing := range AllowedGatewayTemplateTunnelVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GatewayTemplateTunnelVersion", value)
}

// NewGatewayTemplateTunnelVersionFromValue returns a pointer to a valid GatewayTemplateTunnelVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGatewayTemplateTunnelVersionFromValue(v string) (*GatewayTemplateTunnelVersion, error) {
	ev := GatewayTemplateTunnelVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GatewayTemplateTunnelVersion: valid values are %v", v, AllowedGatewayTemplateTunnelVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GatewayTemplateTunnelVersion) IsValid() bool {
	for _, existing := range AllowedGatewayTemplateTunnelVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to gateway_template_tunnel_version value
func (v GatewayTemplateTunnelVersion) Ptr() *GatewayTemplateTunnelVersion {
	return &v
}

type NullableGatewayTemplateTunnelVersion struct {
	value *GatewayTemplateTunnelVersion
	isSet bool
}

func (v NullableGatewayTemplateTunnelVersion) Get() *GatewayTemplateTunnelVersion {
	return v.value
}

func (v *NullableGatewayTemplateTunnelVersion) Set(val *GatewayTemplateTunnelVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateTunnelVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateTunnelVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateTunnelVersion(val *GatewayTemplateTunnelVersion) *NullableGatewayTemplateTunnelVersion {
	return &NullableGatewayTemplateTunnelVersion{value: val, isSet: true}
}

func (v NullableGatewayTemplateTunnelVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateTunnelVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

