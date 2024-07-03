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

// ConstInsightMetricsPropertyScope the model 'ConstInsightMetricsPropertyScope'
type ConstInsightMetricsPropertyScope string

// List of const_insight_metrics_property_scope
const (
	CONSTINSIGHTMETRICSPROPERTYSCOPE_EMPTY ConstInsightMetricsPropertyScope = ""
	CONSTINSIGHTMETRICSPROPERTYSCOPE_SITE ConstInsightMetricsPropertyScope = "site"
	CONSTINSIGHTMETRICSPROPERTYSCOPE_AP ConstInsightMetricsPropertyScope = "ap"
	CONSTINSIGHTMETRICSPROPERTYSCOPE_CLIENT ConstInsightMetricsPropertyScope = "client"
	CONSTINSIGHTMETRICSPROPERTYSCOPE_SWITCH ConstInsightMetricsPropertyScope = "switch"
	CONSTINSIGHTMETRICSPROPERTYSCOPE_DEVICE ConstInsightMetricsPropertyScope = "device"
	CONSTINSIGHTMETRICSPROPERTYSCOPE_MXEDGE ConstInsightMetricsPropertyScope = "mxedge"
)

// All allowed values of ConstInsightMetricsPropertyScope enum
var AllowedConstInsightMetricsPropertyScopeEnumValues = []ConstInsightMetricsPropertyScope{
	"",
	"site",
	"ap",
	"client",
	"switch",
	"device",
	"mxedge",
}

func (v *ConstInsightMetricsPropertyScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConstInsightMetricsPropertyScope(value)
	for _, existing := range AllowedConstInsightMetricsPropertyScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConstInsightMetricsPropertyScope", value)
}

// NewConstInsightMetricsPropertyScopeFromValue returns a pointer to a valid ConstInsightMetricsPropertyScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConstInsightMetricsPropertyScopeFromValue(v string) (*ConstInsightMetricsPropertyScope, error) {
	ev := ConstInsightMetricsPropertyScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConstInsightMetricsPropertyScope: valid values are %v", v, AllowedConstInsightMetricsPropertyScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConstInsightMetricsPropertyScope) IsValid() bool {
	for _, existing := range AllowedConstInsightMetricsPropertyScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to const_insight_metrics_property_scope value
func (v ConstInsightMetricsPropertyScope) Ptr() *ConstInsightMetricsPropertyScope {
	return &v
}

type NullableConstInsightMetricsPropertyScope struct {
	value *ConstInsightMetricsPropertyScope
	isSet bool
}

func (v NullableConstInsightMetricsPropertyScope) Get() *ConstInsightMetricsPropertyScope {
	return v.value
}

func (v *NullableConstInsightMetricsPropertyScope) Set(val *ConstInsightMetricsPropertyScope) {
	v.value = val
	v.isSet = true
}

func (v NullableConstInsightMetricsPropertyScope) IsSet() bool {
	return v.isSet
}

func (v *NullableConstInsightMetricsPropertyScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstInsightMetricsPropertyScope(val *ConstInsightMetricsPropertyScope) *NullableConstInsightMetricsPropertyScope {
	return &NullableConstInsightMetricsPropertyScope{value: val, isSet: true}
}

func (v NullableConstInsightMetricsPropertyScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstInsightMetricsPropertyScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

