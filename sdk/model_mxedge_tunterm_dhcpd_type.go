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

// MxedgeTuntermDhcpdType the model 'MxedgeTuntermDhcpdType'
type MxedgeTuntermDhcpdType string

// List of mxedge_tunterm_dhcpd_type
const (
	MXEDGETUNTERMDHCPDTYPE_EMPTY MxedgeTuntermDhcpdType = ""
	MXEDGETUNTERMDHCPDTYPE_RELAY MxedgeTuntermDhcpdType = "relay"
)

// All allowed values of MxedgeTuntermDhcpdType enum
var AllowedMxedgeTuntermDhcpdTypeEnumValues = []MxedgeTuntermDhcpdType{
	"",
	"relay",
}

func (v *MxedgeTuntermDhcpdType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxedgeTuntermDhcpdType(value)
	for _, existing := range AllowedMxedgeTuntermDhcpdTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxedgeTuntermDhcpdType", value)
}

// NewMxedgeTuntermDhcpdTypeFromValue returns a pointer to a valid MxedgeTuntermDhcpdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxedgeTuntermDhcpdTypeFromValue(v string) (*MxedgeTuntermDhcpdType, error) {
	ev := MxedgeTuntermDhcpdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxedgeTuntermDhcpdType: valid values are %v", v, AllowedMxedgeTuntermDhcpdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxedgeTuntermDhcpdType) IsValid() bool {
	for _, existing := range AllowedMxedgeTuntermDhcpdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxedge_tunterm_dhcpd_type value
func (v MxedgeTuntermDhcpdType) Ptr() *MxedgeTuntermDhcpdType {
	return &v
}

type NullableMxedgeTuntermDhcpdType struct {
	value *MxedgeTuntermDhcpdType
	isSet bool
}

func (v NullableMxedgeTuntermDhcpdType) Get() *MxedgeTuntermDhcpdType {
	return v.value
}

func (v *NullableMxedgeTuntermDhcpdType) Set(val *MxedgeTuntermDhcpdType) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeTuntermDhcpdType) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeTuntermDhcpdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeTuntermDhcpdType(val *MxedgeTuntermDhcpdType) *NullableMxedgeTuntermDhcpdType {
	return &NullableMxedgeTuntermDhcpdType{value: val, isSet: true}
}

func (v NullableMxedgeTuntermDhcpdType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeTuntermDhcpdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

