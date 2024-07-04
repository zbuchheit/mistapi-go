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

// MxclusterTuntermHostsSelection Ordering of tunterm_hosts for mxedge within the same mxcluster.  * When `shuffle`, the ordering of tunterm_hosts is randomized by the device’s MAC.  * When `shuffle-by-site`, we shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels).  * When `ordered`, the order is decided by tunterm_hosts_order
type MxclusterTuntermHostsSelection string

// List of mxcluster_tunterm_hosts_selection
const (
	MXCLUSTERTUNTERMHOSTSSELECTION_EMPTY MxclusterTuntermHostsSelection = ""
	MXCLUSTERTUNTERMHOSTSSELECTION_SHUFFLE MxclusterTuntermHostsSelection = "shuffle"
	MXCLUSTERTUNTERMHOSTSSELECTION_SHUFFLE_BY_SITE MxclusterTuntermHostsSelection = "shuffle-by-site"
	MXCLUSTERTUNTERMHOSTSSELECTION_ORDERED MxclusterTuntermHostsSelection = "ordered"
)

// All allowed values of MxclusterTuntermHostsSelection enum
var AllowedMxclusterTuntermHostsSelectionEnumValues = []MxclusterTuntermHostsSelection{
	"",
	"shuffle",
	"shuffle-by-site",
	"ordered",
}

func (v *MxclusterTuntermHostsSelection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxclusterTuntermHostsSelection(value)
	for _, existing := range AllowedMxclusterTuntermHostsSelectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxclusterTuntermHostsSelection", value)
}

// NewMxclusterTuntermHostsSelectionFromValue returns a pointer to a valid MxclusterTuntermHostsSelection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxclusterTuntermHostsSelectionFromValue(v string) (*MxclusterTuntermHostsSelection, error) {
	ev := MxclusterTuntermHostsSelection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxclusterTuntermHostsSelection: valid values are %v", v, AllowedMxclusterTuntermHostsSelectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxclusterTuntermHostsSelection) IsValid() bool {
	for _, existing := range AllowedMxclusterTuntermHostsSelectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxcluster_tunterm_hosts_selection value
func (v MxclusterTuntermHostsSelection) Ptr() *MxclusterTuntermHostsSelection {
	return &v
}

type NullableMxclusterTuntermHostsSelection struct {
	value *MxclusterTuntermHostsSelection
	isSet bool
}

func (v NullableMxclusterTuntermHostsSelection) Get() *MxclusterTuntermHostsSelection {
	return v.value
}

func (v *NullableMxclusterTuntermHostsSelection) Set(val *MxclusterTuntermHostsSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableMxclusterTuntermHostsSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableMxclusterTuntermHostsSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxclusterTuntermHostsSelection(val *MxclusterTuntermHostsSelection) *NullableMxclusterTuntermHostsSelection {
	return &NullableMxclusterTuntermHostsSelection{value: val, isSet: true}
}

func (v NullableMxclusterTuntermHostsSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxclusterTuntermHostsSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

