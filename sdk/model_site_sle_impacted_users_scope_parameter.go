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

// SiteSleImpactedUsersScopeParameter the model 'SiteSleImpactedUsersScopeParameter'
type SiteSleImpactedUsersScopeParameter string

// List of site_sle_impacted_users_scope_parameter
const (
	SITESLEIMPACTEDUSERSSCOPEPARAMETER_EMPTY SiteSleImpactedUsersScopeParameter = ""
	SITESLEIMPACTEDUSERSSCOPEPARAMETER_SITE SiteSleImpactedUsersScopeParameter = "site"
	SITESLEIMPACTEDUSERSSCOPEPARAMETER_AP SiteSleImpactedUsersScopeParameter = "ap"
)

// All allowed values of SiteSleImpactedUsersScopeParameter enum
var AllowedSiteSleImpactedUsersScopeParameterEnumValues = []SiteSleImpactedUsersScopeParameter{
	"",
	"site",
	"ap",
}

func (v *SiteSleImpactedUsersScopeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSleImpactedUsersScopeParameter(value)
	for _, existing := range AllowedSiteSleImpactedUsersScopeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSleImpactedUsersScopeParameter", value)
}

// NewSiteSleImpactedUsersScopeParameterFromValue returns a pointer to a valid SiteSleImpactedUsersScopeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSleImpactedUsersScopeParameterFromValue(v string) (*SiteSleImpactedUsersScopeParameter, error) {
	ev := SiteSleImpactedUsersScopeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSleImpactedUsersScopeParameter: valid values are %v", v, AllowedSiteSleImpactedUsersScopeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSleImpactedUsersScopeParameter) IsValid() bool {
	for _, existing := range AllowedSiteSleImpactedUsersScopeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sle_impacted_users_scope_parameter value
func (v SiteSleImpactedUsersScopeParameter) Ptr() *SiteSleImpactedUsersScopeParameter {
	return &v
}

type NullableSiteSleImpactedUsersScopeParameter struct {
	value *SiteSleImpactedUsersScopeParameter
	isSet bool
}

func (v NullableSiteSleImpactedUsersScopeParameter) Get() *SiteSleImpactedUsersScopeParameter {
	return v.value
}

func (v *NullableSiteSleImpactedUsersScopeParameter) Set(val *SiteSleImpactedUsersScopeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleImpactedUsersScopeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleImpactedUsersScopeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleImpactedUsersScopeParameter(val *SiteSleImpactedUsersScopeParameter) *NullableSiteSleImpactedUsersScopeParameter {
	return &NullableSiteSleImpactedUsersScopeParameter{value: val, isSet: true}
}

func (v NullableSiteSleImpactedUsersScopeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleImpactedUsersScopeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

