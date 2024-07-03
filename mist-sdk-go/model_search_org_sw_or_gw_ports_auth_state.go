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

// SearchOrgSwOrGwPortsAuthState the model 'SearchOrgSwOrGwPortsAuthState'
type SearchOrgSwOrGwPortsAuthState string

// List of search_org_sw_or_gw_ports_auth_state
const (
	SEARCHORGSWORGWPORTSAUTHSTATE_EMPTY SearchOrgSwOrGwPortsAuthState = ""
	SEARCHORGSWORGWPORTSAUTHSTATE_INIT SearchOrgSwOrGwPortsAuthState = "init"
	SEARCHORGSWORGWPORTSAUTHSTATE_AUTHENTICATED SearchOrgSwOrGwPortsAuthState = "authenticated"
	SEARCHORGSWORGWPORTSAUTHSTATE_AUTHENTICATING SearchOrgSwOrGwPortsAuthState = "authenticating"
	SEARCHORGSWORGWPORTSAUTHSTATE_HELD SearchOrgSwOrGwPortsAuthState = "held"
)

// All allowed values of SearchOrgSwOrGwPortsAuthState enum
var AllowedSearchOrgSwOrGwPortsAuthStateEnumValues = []SearchOrgSwOrGwPortsAuthState{
	"",
	"init",
	"authenticated",
	"authenticating",
	"held",
}

func (v *SearchOrgSwOrGwPortsAuthState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchOrgSwOrGwPortsAuthState(value)
	for _, existing := range AllowedSearchOrgSwOrGwPortsAuthStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchOrgSwOrGwPortsAuthState", value)
}

// NewSearchOrgSwOrGwPortsAuthStateFromValue returns a pointer to a valid SearchOrgSwOrGwPortsAuthState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchOrgSwOrGwPortsAuthStateFromValue(v string) (*SearchOrgSwOrGwPortsAuthState, error) {
	ev := SearchOrgSwOrGwPortsAuthState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchOrgSwOrGwPortsAuthState: valid values are %v", v, AllowedSearchOrgSwOrGwPortsAuthStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchOrgSwOrGwPortsAuthState) IsValid() bool {
	for _, existing := range AllowedSearchOrgSwOrGwPortsAuthStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to search_org_sw_or_gw_ports_auth_state value
func (v SearchOrgSwOrGwPortsAuthState) Ptr() *SearchOrgSwOrGwPortsAuthState {
	return &v
}

type NullableSearchOrgSwOrGwPortsAuthState struct {
	value *SearchOrgSwOrGwPortsAuthState
	isSet bool
}

func (v NullableSearchOrgSwOrGwPortsAuthState) Get() *SearchOrgSwOrGwPortsAuthState {
	return v.value
}

func (v *NullableSearchOrgSwOrGwPortsAuthState) Set(val *SearchOrgSwOrGwPortsAuthState) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchOrgSwOrGwPortsAuthState) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchOrgSwOrGwPortsAuthState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchOrgSwOrGwPortsAuthState(val *SearchOrgSwOrGwPortsAuthState) *NullableSearchOrgSwOrGwPortsAuthState {
	return &NullableSearchOrgSwOrGwPortsAuthState{value: val, isSet: true}
}

func (v NullableSearchOrgSwOrGwPortsAuthState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchOrgSwOrGwPortsAuthState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

