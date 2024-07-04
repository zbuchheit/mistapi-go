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

// CountSiteSwitchPortsStpState the model 'CountSiteSwitchPortsStpState'
type CountSiteSwitchPortsStpState string

// List of count_site_switch_ports_stp_state
const (
	COUNTSITESWITCHPORTSSTPSTATE_EMPTY CountSiteSwitchPortsStpState = ""
	COUNTSITESWITCHPORTSSTPSTATE_FORWARDING CountSiteSwitchPortsStpState = "forwarding"
	COUNTSITESWITCHPORTSSTPSTATE_BLOCKING CountSiteSwitchPortsStpState = "blocking"
	COUNTSITESWITCHPORTSSTPSTATE_LEARNING CountSiteSwitchPortsStpState = "learning"
	COUNTSITESWITCHPORTSSTPSTATE_LISTENING CountSiteSwitchPortsStpState = "listening"
	COUNTSITESWITCHPORTSSTPSTATE_DISABLED CountSiteSwitchPortsStpState = "disabled"
)

// All allowed values of CountSiteSwitchPortsStpState enum
var AllowedCountSiteSwitchPortsStpStateEnumValues = []CountSiteSwitchPortsStpState{
	"",
	"forwarding",
	"blocking",
	"learning",
	"listening",
	"disabled",
}

func (v *CountSiteSwitchPortsStpState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CountSiteSwitchPortsStpState(value)
	for _, existing := range AllowedCountSiteSwitchPortsStpStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CountSiteSwitchPortsStpState", value)
}

// NewCountSiteSwitchPortsStpStateFromValue returns a pointer to a valid CountSiteSwitchPortsStpState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountSiteSwitchPortsStpStateFromValue(v string) (*CountSiteSwitchPortsStpState, error) {
	ev := CountSiteSwitchPortsStpState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CountSiteSwitchPortsStpState: valid values are %v", v, AllowedCountSiteSwitchPortsStpStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountSiteSwitchPortsStpState) IsValid() bool {
	for _, existing := range AllowedCountSiteSwitchPortsStpStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to count_site_switch_ports_stp_state value
func (v CountSiteSwitchPortsStpState) Ptr() *CountSiteSwitchPortsStpState {
	return &v
}

type NullableCountSiteSwitchPortsStpState struct {
	value *CountSiteSwitchPortsStpState
	isSet bool
}

func (v NullableCountSiteSwitchPortsStpState) Get() *CountSiteSwitchPortsStpState {
	return v.value
}

func (v *NullableCountSiteSwitchPortsStpState) Set(val *CountSiteSwitchPortsStpState) {
	v.value = val
	v.isSet = true
}

func (v NullableCountSiteSwitchPortsStpState) IsSet() bool {
	return v.isSet
}

func (v *NullableCountSiteSwitchPortsStpState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountSiteSwitchPortsStpState(val *CountSiteSwitchPortsStpState) *NullableCountSiteSwitchPortsStpState {
	return &NullableCountSiteSwitchPortsStpState{value: val, isSet: true}
}

func (v NullableCountSiteSwitchPortsStpState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountSiteSwitchPortsStpState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

