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

// SiteMxedgeEventsCountDistinct the model 'SiteMxedgeEventsCountDistinct'
type SiteMxedgeEventsCountDistinct string

// List of site_mxedge_events_count_distinct
const (
	SITEMXEDGEEVENTSCOUNTDISTINCT_EMPTY SiteMxedgeEventsCountDistinct = ""
	SITEMXEDGEEVENTSCOUNTDISTINCT_MXEDGE_ID SiteMxedgeEventsCountDistinct = "mxedge_id"
	SITEMXEDGEEVENTSCOUNTDISTINCT_TYPE SiteMxedgeEventsCountDistinct = "type"
	SITEMXEDGEEVENTSCOUNTDISTINCT_MXCLUSTER_ID SiteMxedgeEventsCountDistinct = "mxcluster_id"
	SITEMXEDGEEVENTSCOUNTDISTINCT_PACKAGE SiteMxedgeEventsCountDistinct = "package"
)

// All allowed values of SiteMxedgeEventsCountDistinct enum
var AllowedSiteMxedgeEventsCountDistinctEnumValues = []SiteMxedgeEventsCountDistinct{
	"",
	"mxedge_id",
	"type",
	"mxcluster_id",
	"package",
}

func (v *SiteMxedgeEventsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteMxedgeEventsCountDistinct(value)
	for _, existing := range AllowedSiteMxedgeEventsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteMxedgeEventsCountDistinct", value)
}

// NewSiteMxedgeEventsCountDistinctFromValue returns a pointer to a valid SiteMxedgeEventsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteMxedgeEventsCountDistinctFromValue(v string) (*SiteMxedgeEventsCountDistinct, error) {
	ev := SiteMxedgeEventsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteMxedgeEventsCountDistinct: valid values are %v", v, AllowedSiteMxedgeEventsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteMxedgeEventsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteMxedgeEventsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_mxedge_events_count_distinct value
func (v SiteMxedgeEventsCountDistinct) Ptr() *SiteMxedgeEventsCountDistinct {
	return &v
}

type NullableSiteMxedgeEventsCountDistinct struct {
	value *SiteMxedgeEventsCountDistinct
	isSet bool
}

func (v NullableSiteMxedgeEventsCountDistinct) Get() *SiteMxedgeEventsCountDistinct {
	return v.value
}

func (v *NullableSiteMxedgeEventsCountDistinct) Set(val *SiteMxedgeEventsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteMxedgeEventsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteMxedgeEventsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteMxedgeEventsCountDistinct(val *SiteMxedgeEventsCountDistinct) *NullableSiteMxedgeEventsCountDistinct {
	return &NullableSiteMxedgeEventsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteMxedgeEventsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteMxedgeEventsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

