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

// OrgTicketsCountDistinct the model 'OrgTicketsCountDistinct'
type OrgTicketsCountDistinct string

// List of org_tickets_count_distinct
const (
	ORGTICKETSCOUNTDISTINCT_EMPTY OrgTicketsCountDistinct = ""
	ORGTICKETSCOUNTDISTINCT_STATUS OrgTicketsCountDistinct = "status"
	ORGTICKETSCOUNTDISTINCT_TYPE OrgTicketsCountDistinct = "type"
)

// All allowed values of OrgTicketsCountDistinct enum
var AllowedOrgTicketsCountDistinctEnumValues = []OrgTicketsCountDistinct{
	"",
	"status",
	"type",
}

func (v *OrgTicketsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgTicketsCountDistinct(value)
	for _, existing := range AllowedOrgTicketsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgTicketsCountDistinct", value)
}

// NewOrgTicketsCountDistinctFromValue returns a pointer to a valid OrgTicketsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgTicketsCountDistinctFromValue(v string) (*OrgTicketsCountDistinct, error) {
	ev := OrgTicketsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgTicketsCountDistinct: valid values are %v", v, AllowedOrgTicketsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgTicketsCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgTicketsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_tickets_count_distinct value
func (v OrgTicketsCountDistinct) Ptr() *OrgTicketsCountDistinct {
	return &v
}

type NullableOrgTicketsCountDistinct struct {
	value *OrgTicketsCountDistinct
	isSet bool
}

func (v NullableOrgTicketsCountDistinct) Get() *OrgTicketsCountDistinct {
	return v.value
}

func (v *NullableOrgTicketsCountDistinct) Set(val *OrgTicketsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgTicketsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgTicketsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgTicketsCountDistinct(val *OrgTicketsCountDistinct) *NullableOrgTicketsCountDistinct {
	return &NullableOrgTicketsCountDistinct{value: val, isSet: true}
}

func (v NullableOrgTicketsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgTicketsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

