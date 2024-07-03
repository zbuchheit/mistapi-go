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

// CountSiteSwOrGwPortsStpState the model 'CountSiteSwOrGwPortsStpState'
type CountSiteSwOrGwPortsStpState string

// List of count_site_sw_or_gw_ports_stp_state
const (
	COUNTSITESWORGWPORTSSTPSTATE_EMPTY CountSiteSwOrGwPortsStpState = ""
	COUNTSITESWORGWPORTSSTPSTATE_FORWARDING CountSiteSwOrGwPortsStpState = "forwarding"
	COUNTSITESWORGWPORTSSTPSTATE_BLOCKING CountSiteSwOrGwPortsStpState = "blocking"
	COUNTSITESWORGWPORTSSTPSTATE_LEARNING CountSiteSwOrGwPortsStpState = "learning"
	COUNTSITESWORGWPORTSSTPSTATE_LISTENING CountSiteSwOrGwPortsStpState = "listening"
	COUNTSITESWORGWPORTSSTPSTATE_DISABLED CountSiteSwOrGwPortsStpState = "disabled"
)

// All allowed values of CountSiteSwOrGwPortsStpState enum
var AllowedCountSiteSwOrGwPortsStpStateEnumValues = []CountSiteSwOrGwPortsStpState{
	"",
	"forwarding",
	"blocking",
	"learning",
	"listening",
	"disabled",
}

func (v *CountSiteSwOrGwPortsStpState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CountSiteSwOrGwPortsStpState(value)
	for _, existing := range AllowedCountSiteSwOrGwPortsStpStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CountSiteSwOrGwPortsStpState", value)
}

// NewCountSiteSwOrGwPortsStpStateFromValue returns a pointer to a valid CountSiteSwOrGwPortsStpState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountSiteSwOrGwPortsStpStateFromValue(v string) (*CountSiteSwOrGwPortsStpState, error) {
	ev := CountSiteSwOrGwPortsStpState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CountSiteSwOrGwPortsStpState: valid values are %v", v, AllowedCountSiteSwOrGwPortsStpStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountSiteSwOrGwPortsStpState) IsValid() bool {
	for _, existing := range AllowedCountSiteSwOrGwPortsStpStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to count_site_sw_or_gw_ports_stp_state value
func (v CountSiteSwOrGwPortsStpState) Ptr() *CountSiteSwOrGwPortsStpState {
	return &v
}

type NullableCountSiteSwOrGwPortsStpState struct {
	value *CountSiteSwOrGwPortsStpState
	isSet bool
}

func (v NullableCountSiteSwOrGwPortsStpState) Get() *CountSiteSwOrGwPortsStpState {
	return v.value
}

func (v *NullableCountSiteSwOrGwPortsStpState) Set(val *CountSiteSwOrGwPortsStpState) {
	v.value = val
	v.isSet = true
}

func (v NullableCountSiteSwOrGwPortsStpState) IsSet() bool {
	return v.isSet
}

func (v *NullableCountSiteSwOrGwPortsStpState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountSiteSwOrGwPortsStpState(val *CountSiteSwOrGwPortsStpState) *NullableCountSiteSwOrGwPortsStpState {
	return &NullableCountSiteSwOrGwPortsStpState{value: val, isSet: true}
}

func (v NullableCountSiteSwOrGwPortsStpState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountSiteSwOrGwPortsStpState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

