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

// OrgOtherdevicesEventsCountDistinct the model 'OrgOtherdevicesEventsCountDistinct'
type OrgOtherdevicesEventsCountDistinct string

// List of org_otherdevices_events_count_distinct
const (
	ORGOTHERDEVICESEVENTSCOUNTDISTINCT_EMPTY OrgOtherdevicesEventsCountDistinct = ""
	ORGOTHERDEVICESEVENTSCOUNTDISTINCT_MAC OrgOtherdevicesEventsCountDistinct = "mac"
	ORGOTHERDEVICESEVENTSCOUNTDISTINCT_TYPE OrgOtherdevicesEventsCountDistinct = "type"
	ORGOTHERDEVICESEVENTSCOUNTDISTINCT_VENDOR OrgOtherdevicesEventsCountDistinct = "vendor"
	ORGOTHERDEVICESEVENTSCOUNTDISTINCT_SITE_ID OrgOtherdevicesEventsCountDistinct = "site_id"
)

// All allowed values of OrgOtherdevicesEventsCountDistinct enum
var AllowedOrgOtherdevicesEventsCountDistinctEnumValues = []OrgOtherdevicesEventsCountDistinct{
	"",
	"mac",
	"type",
	"vendor",
	"site_id",
}

func (v *OrgOtherdevicesEventsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgOtherdevicesEventsCountDistinct(value)
	for _, existing := range AllowedOrgOtherdevicesEventsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgOtherdevicesEventsCountDistinct", value)
}

// NewOrgOtherdevicesEventsCountDistinctFromValue returns a pointer to a valid OrgOtherdevicesEventsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgOtherdevicesEventsCountDistinctFromValue(v string) (*OrgOtherdevicesEventsCountDistinct, error) {
	ev := OrgOtherdevicesEventsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgOtherdevicesEventsCountDistinct: valid values are %v", v, AllowedOrgOtherdevicesEventsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgOtherdevicesEventsCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgOtherdevicesEventsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_otherdevices_events_count_distinct value
func (v OrgOtherdevicesEventsCountDistinct) Ptr() *OrgOtherdevicesEventsCountDistinct {
	return &v
}

type NullableOrgOtherdevicesEventsCountDistinct struct {
	value *OrgOtherdevicesEventsCountDistinct
	isSet bool
}

func (v NullableOrgOtherdevicesEventsCountDistinct) Get() *OrgOtherdevicesEventsCountDistinct {
	return v.value
}

func (v *NullableOrgOtherdevicesEventsCountDistinct) Set(val *OrgOtherdevicesEventsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgOtherdevicesEventsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgOtherdevicesEventsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgOtherdevicesEventsCountDistinct(val *OrgOtherdevicesEventsCountDistinct) *NullableOrgOtherdevicesEventsCountDistinct {
	return &NullableOrgOtherdevicesEventsCountDistinct{value: val, isSet: true}
}

func (v NullableOrgOtherdevicesEventsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgOtherdevicesEventsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

