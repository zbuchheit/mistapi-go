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

// SwitchMetricScope the model 'SwitchMetricScope'
type SwitchMetricScope string

// List of switch_metric_scope
const (
	SWITCHMETRICSCOPE_EMPTY SwitchMetricScope = ""
	SWITCHMETRICSCOPE_SITE SwitchMetricScope = "site"
	SWITCHMETRICSCOPE_SWITCH SwitchMetricScope = "switch"
)

// All allowed values of SwitchMetricScope enum
var AllowedSwitchMetricScopeEnumValues = []SwitchMetricScope{
	"",
	"site",
	"switch",
}

func (v *SwitchMetricScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SwitchMetricScope(value)
	for _, existing := range AllowedSwitchMetricScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SwitchMetricScope", value)
}

// NewSwitchMetricScopeFromValue returns a pointer to a valid SwitchMetricScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSwitchMetricScopeFromValue(v string) (*SwitchMetricScope, error) {
	ev := SwitchMetricScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SwitchMetricScope: valid values are %v", v, AllowedSwitchMetricScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SwitchMetricScope) IsValid() bool {
	for _, existing := range AllowedSwitchMetricScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to switch_metric_scope value
func (v SwitchMetricScope) Ptr() *SwitchMetricScope {
	return &v
}

type NullableSwitchMetricScope struct {
	value *SwitchMetricScope
	isSet bool
}

func (v NullableSwitchMetricScope) Get() *SwitchMetricScope {
	return v.value
}

func (v *NullableSwitchMetricScope) Set(val *SwitchMetricScope) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchMetricScope) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchMetricScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchMetricScope(val *SwitchMetricScope) *NullableSwitchMetricScope {
	return &NullableSwitchMetricScope{value: val, isSet: true}
}

func (v NullableSwitchMetricScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchMetricScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

