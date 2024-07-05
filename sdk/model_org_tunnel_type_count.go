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

// OrgTunnelTypeCount the model 'OrgTunnelTypeCount'
type OrgTunnelTypeCount string

// List of org_tunnel_type_count
const (
	ORGTUNNELTYPECOUNT_EMPTY OrgTunnelTypeCount = ""
	ORGTUNNELTYPECOUNT_WXTUNNEL OrgTunnelTypeCount = "wxtunnel"
	ORGTUNNELTYPECOUNT_WAN OrgTunnelTypeCount = "wan"
)

// All allowed values of OrgTunnelTypeCount enum
var AllowedOrgTunnelTypeCountEnumValues = []OrgTunnelTypeCount{
	"",
	"wxtunnel",
	"wan",
}

func (v *OrgTunnelTypeCount) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgTunnelTypeCount(value)
	for _, existing := range AllowedOrgTunnelTypeCountEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgTunnelTypeCount", value)
}

// NewOrgTunnelTypeCountFromValue returns a pointer to a valid OrgTunnelTypeCount
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgTunnelTypeCountFromValue(v string) (*OrgTunnelTypeCount, error) {
	ev := OrgTunnelTypeCount(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgTunnelTypeCount: valid values are %v", v, AllowedOrgTunnelTypeCountEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgTunnelTypeCount) IsValid() bool {
	for _, existing := range AllowedOrgTunnelTypeCountEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_tunnel_type_count value
func (v OrgTunnelTypeCount) Ptr() *OrgTunnelTypeCount {
	return &v
}

type NullableOrgTunnelTypeCount struct {
	value *OrgTunnelTypeCount
	isSet bool
}

func (v NullableOrgTunnelTypeCount) Get() *OrgTunnelTypeCount {
	return v.value
}

func (v *NullableOrgTunnelTypeCount) Set(val *OrgTunnelTypeCount) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgTunnelTypeCount) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgTunnelTypeCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgTunnelTypeCount(val *OrgTunnelTypeCount) *NullableOrgTunnelTypeCount {
	return &NullableOrgTunnelTypeCount{value: val, isSet: true}
}

func (v NullableOrgTunnelTypeCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgTunnelTypeCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

