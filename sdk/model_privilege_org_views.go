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

// PrivilegeOrgViews Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.  Below are the list of supported UI views. Note that this is UI only feature Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.  You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.  Below are the list of supported UI views. Note that this is UI only feature  | UI View | Description | | --- | --- | | `reporting` | full access to all analytics tools | | `marketing` | can view analytics and location maps | | `location` | can view and manage location maps | | `security` | can view and manage WLAN, rogues and authentication | | `switch_admin` | can view and manage Switch ports | | `mxedge_admin` | can view and manage Mist edges and Mist tunnels | | `lobby_admin` | full access to Org and Site Pre-shared keys |
type PrivilegeOrgViews string

// List of privilege_org_views
const (
	PRIVILEGEORGVIEWS_EMPTY PrivilegeOrgViews = ""
	PRIVILEGEORGVIEWS_REPORTING PrivilegeOrgViews = "reporting"
	PRIVILEGEORGVIEWS_MARKETING PrivilegeOrgViews = "marketing"
	PRIVILEGEORGVIEWS_LOCATION PrivilegeOrgViews = "location"
	PRIVILEGEORGVIEWS_SECURITY PrivilegeOrgViews = "security"
	PRIVILEGEORGVIEWS_SWITCH_ADMIN PrivilegeOrgViews = "switch_admin"
	PRIVILEGEORGVIEWS_MXEDGE_ADMIN PrivilegeOrgViews = "mxedge_admin"
	PRIVILEGEORGVIEWS_LOBBY_ADMIN PrivilegeOrgViews = "lobby_admin"
)

// All allowed values of PrivilegeOrgViews enum
var AllowedPrivilegeOrgViewsEnumValues = []PrivilegeOrgViews{
	"",
	"reporting",
	"marketing",
	"location",
	"security",
	"switch_admin",
	"mxedge_admin",
	"lobby_admin",
}

func (v *PrivilegeOrgViews) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrivilegeOrgViews(value)
	for _, existing := range AllowedPrivilegeOrgViewsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrivilegeOrgViews", value)
}

// NewPrivilegeOrgViewsFromValue returns a pointer to a valid PrivilegeOrgViews
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrivilegeOrgViewsFromValue(v string) (*PrivilegeOrgViews, error) {
	ev := PrivilegeOrgViews(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrivilegeOrgViews: valid values are %v", v, AllowedPrivilegeOrgViewsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrivilegeOrgViews) IsValid() bool {
	for _, existing := range AllowedPrivilegeOrgViewsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to privilege_org_views value
func (v PrivilegeOrgViews) Ptr() *PrivilegeOrgViews {
	return &v
}

type NullablePrivilegeOrgViews struct {
	value *PrivilegeOrgViews
	isSet bool
}

func (v NullablePrivilegeOrgViews) Get() *PrivilegeOrgViews {
	return v.value
}

func (v *NullablePrivilegeOrgViews) Set(val *PrivilegeOrgViews) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegeOrgViews) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegeOrgViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegeOrgViews(val *PrivilegeOrgViews) *NullablePrivilegeOrgViews {
	return &NullablePrivilegeOrgViews{value: val, isSet: true}
}

func (v NullablePrivilegeOrgViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegeOrgViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

