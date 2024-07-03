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

// SearchSiteSwitchPortsAuthState the model 'SearchSiteSwitchPortsAuthState'
type SearchSiteSwitchPortsAuthState string

// List of search_site_switch_ports_auth_state
const (
	SEARCHSITESWITCHPORTSAUTHSTATE_EMPTY SearchSiteSwitchPortsAuthState = ""
	SEARCHSITESWITCHPORTSAUTHSTATE_INIT SearchSiteSwitchPortsAuthState = "init"
	SEARCHSITESWITCHPORTSAUTHSTATE_AUTHENTICATED SearchSiteSwitchPortsAuthState = "authenticated"
	SEARCHSITESWITCHPORTSAUTHSTATE_AUTHENTICATING SearchSiteSwitchPortsAuthState = "authenticating"
	SEARCHSITESWITCHPORTSAUTHSTATE_HELD SearchSiteSwitchPortsAuthState = "held"
)

// All allowed values of SearchSiteSwitchPortsAuthState enum
var AllowedSearchSiteSwitchPortsAuthStateEnumValues = []SearchSiteSwitchPortsAuthState{
	"",
	"init",
	"authenticated",
	"authenticating",
	"held",
}

func (v *SearchSiteSwitchPortsAuthState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchSiteSwitchPortsAuthState(value)
	for _, existing := range AllowedSearchSiteSwitchPortsAuthStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchSiteSwitchPortsAuthState", value)
}

// NewSearchSiteSwitchPortsAuthStateFromValue returns a pointer to a valid SearchSiteSwitchPortsAuthState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchSiteSwitchPortsAuthStateFromValue(v string) (*SearchSiteSwitchPortsAuthState, error) {
	ev := SearchSiteSwitchPortsAuthState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchSiteSwitchPortsAuthState: valid values are %v", v, AllowedSearchSiteSwitchPortsAuthStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchSiteSwitchPortsAuthState) IsValid() bool {
	for _, existing := range AllowedSearchSiteSwitchPortsAuthStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to search_site_switch_ports_auth_state value
func (v SearchSiteSwitchPortsAuthState) Ptr() *SearchSiteSwitchPortsAuthState {
	return &v
}

type NullableSearchSiteSwitchPortsAuthState struct {
	value *SearchSiteSwitchPortsAuthState
	isSet bool
}

func (v NullableSearchSiteSwitchPortsAuthState) Get() *SearchSiteSwitchPortsAuthState {
	return v.value
}

func (v *NullableSearchSiteSwitchPortsAuthState) Set(val *SearchSiteSwitchPortsAuthState) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSiteSwitchPortsAuthState) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSiteSwitchPortsAuthState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSiteSwitchPortsAuthState(val *SearchSiteSwitchPortsAuthState) *NullableSearchSiteSwitchPortsAuthState {
	return &NullableSearchSiteSwitchPortsAuthState{value: val, isSet: true}
}

func (v NullableSearchSiteSwitchPortsAuthState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSiteSwitchPortsAuthState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

