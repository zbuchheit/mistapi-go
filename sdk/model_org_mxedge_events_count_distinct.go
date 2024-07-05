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

// OrgMxedgeEventsCountDistinct the model 'OrgMxedgeEventsCountDistinct'
type OrgMxedgeEventsCountDistinct string

// List of org_mxedge_events_count_distinct
const (
	ORGMXEDGEEVENTSCOUNTDISTINCT_EMPTY OrgMxedgeEventsCountDistinct = ""
	ORGMXEDGEEVENTSCOUNTDISTINCT_MXEDGE_ID OrgMxedgeEventsCountDistinct = "mxedge_id"
	ORGMXEDGEEVENTSCOUNTDISTINCT_TYPE OrgMxedgeEventsCountDistinct = "type"
	ORGMXEDGEEVENTSCOUNTDISTINCT_MXCLUSTER_ID OrgMxedgeEventsCountDistinct = "mxcluster_id"
	ORGMXEDGEEVENTSCOUNTDISTINCT_PACKAGE OrgMxedgeEventsCountDistinct = "package"
)

// All allowed values of OrgMxedgeEventsCountDistinct enum
var AllowedOrgMxedgeEventsCountDistinctEnumValues = []OrgMxedgeEventsCountDistinct{
	"",
	"mxedge_id",
	"type",
	"mxcluster_id",
	"package",
}

func (v *OrgMxedgeEventsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgMxedgeEventsCountDistinct(value)
	for _, existing := range AllowedOrgMxedgeEventsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgMxedgeEventsCountDistinct", value)
}

// NewOrgMxedgeEventsCountDistinctFromValue returns a pointer to a valid OrgMxedgeEventsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgMxedgeEventsCountDistinctFromValue(v string) (*OrgMxedgeEventsCountDistinct, error) {
	ev := OrgMxedgeEventsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgMxedgeEventsCountDistinct: valid values are %v", v, AllowedOrgMxedgeEventsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgMxedgeEventsCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgMxedgeEventsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_mxedge_events_count_distinct value
func (v OrgMxedgeEventsCountDistinct) Ptr() *OrgMxedgeEventsCountDistinct {
	return &v
}

type NullableOrgMxedgeEventsCountDistinct struct {
	value *OrgMxedgeEventsCountDistinct
	isSet bool
}

func (v NullableOrgMxedgeEventsCountDistinct) Get() *OrgMxedgeEventsCountDistinct {
	return v.value
}

func (v *NullableOrgMxedgeEventsCountDistinct) Set(val *OrgMxedgeEventsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgMxedgeEventsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgMxedgeEventsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgMxedgeEventsCountDistinct(val *OrgMxedgeEventsCountDistinct) *NullableOrgMxedgeEventsCountDistinct {
	return &NullableOrgMxedgeEventsCountDistinct{value: val, isSet: true}
}

func (v NullableOrgMxedgeEventsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgMxedgeEventsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

