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

// CountOrgSwitchPortsStpState the model 'CountOrgSwitchPortsStpState'
type CountOrgSwitchPortsStpState string

// List of count_org_switch_ports_stp_state
const (
	COUNTORGSWITCHPORTSSTPSTATE_EMPTY CountOrgSwitchPortsStpState = ""
	COUNTORGSWITCHPORTSSTPSTATE_FORWARDING CountOrgSwitchPortsStpState = "forwarding"
	COUNTORGSWITCHPORTSSTPSTATE_BLOCKING CountOrgSwitchPortsStpState = "blocking"
	COUNTORGSWITCHPORTSSTPSTATE_LEARNING CountOrgSwitchPortsStpState = "learning"
	COUNTORGSWITCHPORTSSTPSTATE_LISTENING CountOrgSwitchPortsStpState = "listening"
	COUNTORGSWITCHPORTSSTPSTATE_DISABLED CountOrgSwitchPortsStpState = "disabled"
)

// All allowed values of CountOrgSwitchPortsStpState enum
var AllowedCountOrgSwitchPortsStpStateEnumValues = []CountOrgSwitchPortsStpState{
	"",
	"forwarding",
	"blocking",
	"learning",
	"listening",
	"disabled",
}

func (v *CountOrgSwitchPortsStpState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CountOrgSwitchPortsStpState(value)
	for _, existing := range AllowedCountOrgSwitchPortsStpStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CountOrgSwitchPortsStpState", value)
}

// NewCountOrgSwitchPortsStpStateFromValue returns a pointer to a valid CountOrgSwitchPortsStpState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCountOrgSwitchPortsStpStateFromValue(v string) (*CountOrgSwitchPortsStpState, error) {
	ev := CountOrgSwitchPortsStpState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CountOrgSwitchPortsStpState: valid values are %v", v, AllowedCountOrgSwitchPortsStpStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CountOrgSwitchPortsStpState) IsValid() bool {
	for _, existing := range AllowedCountOrgSwitchPortsStpStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to count_org_switch_ports_stp_state value
func (v CountOrgSwitchPortsStpState) Ptr() *CountOrgSwitchPortsStpState {
	return &v
}

type NullableCountOrgSwitchPortsStpState struct {
	value *CountOrgSwitchPortsStpState
	isSet bool
}

func (v NullableCountOrgSwitchPortsStpState) Get() *CountOrgSwitchPortsStpState {
	return v.value
}

func (v *NullableCountOrgSwitchPortsStpState) Set(val *CountOrgSwitchPortsStpState) {
	v.value = val
	v.isSet = true
}

func (v NullableCountOrgSwitchPortsStpState) IsSet() bool {
	return v.isSet
}

func (v *NullableCountOrgSwitchPortsStpState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountOrgSwitchPortsStpState(val *CountOrgSwitchPortsStpState) *NullableCountOrgSwitchPortsStpState {
	return &NullableCountOrgSwitchPortsStpState{value: val, isSet: true}
}

func (v NullableCountOrgSwitchPortsStpState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountOrgSwitchPortsStpState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

