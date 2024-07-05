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

// SwitchPortUsageDynamicResetDefaultWhen Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage Configuring to none will let the DPC port keep at the current port usage.
type SwitchPortUsageDynamicResetDefaultWhen string

// List of switch_port_usage_dynamic_reset_default_when
const (
	SWITCHPORTUSAGEDYNAMICRESETDEFAULTWHEN_EMPTY SwitchPortUsageDynamicResetDefaultWhen = ""
	SWITCHPORTUSAGEDYNAMICRESETDEFAULTWHEN_NONE SwitchPortUsageDynamicResetDefaultWhen = "none"
	SWITCHPORTUSAGEDYNAMICRESETDEFAULTWHEN_LINK_DOWN SwitchPortUsageDynamicResetDefaultWhen = "link_down"
)

// All allowed values of SwitchPortUsageDynamicResetDefaultWhen enum
var AllowedSwitchPortUsageDynamicResetDefaultWhenEnumValues = []SwitchPortUsageDynamicResetDefaultWhen{
	"",
	"none",
	"link_down",
}

func (v *SwitchPortUsageDynamicResetDefaultWhen) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SwitchPortUsageDynamicResetDefaultWhen(value)
	for _, existing := range AllowedSwitchPortUsageDynamicResetDefaultWhenEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SwitchPortUsageDynamicResetDefaultWhen", value)
}

// NewSwitchPortUsageDynamicResetDefaultWhenFromValue returns a pointer to a valid SwitchPortUsageDynamicResetDefaultWhen
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSwitchPortUsageDynamicResetDefaultWhenFromValue(v string) (*SwitchPortUsageDynamicResetDefaultWhen, error) {
	ev := SwitchPortUsageDynamicResetDefaultWhen(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SwitchPortUsageDynamicResetDefaultWhen: valid values are %v", v, AllowedSwitchPortUsageDynamicResetDefaultWhenEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SwitchPortUsageDynamicResetDefaultWhen) IsValid() bool {
	for _, existing := range AllowedSwitchPortUsageDynamicResetDefaultWhenEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to switch_port_usage_dynamic_reset_default_when value
func (v SwitchPortUsageDynamicResetDefaultWhen) Ptr() *SwitchPortUsageDynamicResetDefaultWhen {
	return &v
}

type NullableSwitchPortUsageDynamicResetDefaultWhen struct {
	value *SwitchPortUsageDynamicResetDefaultWhen
	isSet bool
}

func (v NullableSwitchPortUsageDynamicResetDefaultWhen) Get() *SwitchPortUsageDynamicResetDefaultWhen {
	return v.value
}

func (v *NullableSwitchPortUsageDynamicResetDefaultWhen) Set(val *SwitchPortUsageDynamicResetDefaultWhen) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchPortUsageDynamicResetDefaultWhen) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchPortUsageDynamicResetDefaultWhen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchPortUsageDynamicResetDefaultWhen(val *SwitchPortUsageDynamicResetDefaultWhen) *NullableSwitchPortUsageDynamicResetDefaultWhen {
	return &NullableSwitchPortUsageDynamicResetDefaultWhen{value: val, isSet: true}
}

func (v NullableSwitchPortUsageDynamicResetDefaultWhen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchPortUsageDynamicResetDefaultWhen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

