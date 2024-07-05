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

// SearchSiteSwOrGwPortsAuthState the model 'SearchSiteSwOrGwPortsAuthState'
type SearchSiteSwOrGwPortsAuthState string

// List of search_site_sw_or_gw_ports_auth_state
const (
	SEARCHSITESWORGWPORTSAUTHSTATE_EMPTY SearchSiteSwOrGwPortsAuthState = ""
	SEARCHSITESWORGWPORTSAUTHSTATE_INIT SearchSiteSwOrGwPortsAuthState = "init"
	SEARCHSITESWORGWPORTSAUTHSTATE_AUTHENTICATED SearchSiteSwOrGwPortsAuthState = "authenticated"
	SEARCHSITESWORGWPORTSAUTHSTATE_AUTHENTICATING SearchSiteSwOrGwPortsAuthState = "authenticating"
	SEARCHSITESWORGWPORTSAUTHSTATE_HELD SearchSiteSwOrGwPortsAuthState = "held"
)

// All allowed values of SearchSiteSwOrGwPortsAuthState enum
var AllowedSearchSiteSwOrGwPortsAuthStateEnumValues = []SearchSiteSwOrGwPortsAuthState{
	"",
	"init",
	"authenticated",
	"authenticating",
	"held",
}

func (v *SearchSiteSwOrGwPortsAuthState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchSiteSwOrGwPortsAuthState(value)
	for _, existing := range AllowedSearchSiteSwOrGwPortsAuthStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchSiteSwOrGwPortsAuthState", value)
}

// NewSearchSiteSwOrGwPortsAuthStateFromValue returns a pointer to a valid SearchSiteSwOrGwPortsAuthState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchSiteSwOrGwPortsAuthStateFromValue(v string) (*SearchSiteSwOrGwPortsAuthState, error) {
	ev := SearchSiteSwOrGwPortsAuthState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchSiteSwOrGwPortsAuthState: valid values are %v", v, AllowedSearchSiteSwOrGwPortsAuthStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchSiteSwOrGwPortsAuthState) IsValid() bool {
	for _, existing := range AllowedSearchSiteSwOrGwPortsAuthStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to search_site_sw_or_gw_ports_auth_state value
func (v SearchSiteSwOrGwPortsAuthState) Ptr() *SearchSiteSwOrGwPortsAuthState {
	return &v
}

type NullableSearchSiteSwOrGwPortsAuthState struct {
	value *SearchSiteSwOrGwPortsAuthState
	isSet bool
}

func (v NullableSearchSiteSwOrGwPortsAuthState) Get() *SearchSiteSwOrGwPortsAuthState {
	return v.value
}

func (v *NullableSearchSiteSwOrGwPortsAuthState) Set(val *SearchSiteSwOrGwPortsAuthState) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSiteSwOrGwPortsAuthState) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSiteSwOrGwPortsAuthState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSiteSwOrGwPortsAuthState(val *SearchSiteSwOrGwPortsAuthState) *NullableSearchSiteSwOrGwPortsAuthState {
	return &NullableSearchSiteSwOrGwPortsAuthState{value: val, isSet: true}
}

func (v NullableSearchSiteSwOrGwPortsAuthState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSiteSwOrGwPortsAuthState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

