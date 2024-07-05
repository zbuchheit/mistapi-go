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

// SiteRogueEventsCountDistinct the model 'SiteRogueEventsCountDistinct'
type SiteRogueEventsCountDistinct string

// List of site_rogue_events_count_distinct
const (
	SITEROGUEEVENTSCOUNTDISTINCT_EMPTY SiteRogueEventsCountDistinct = ""
	SITEROGUEEVENTSCOUNTDISTINCT_BSSID SiteRogueEventsCountDistinct = "bssid"
	SITEROGUEEVENTSCOUNTDISTINCT_SSID SiteRogueEventsCountDistinct = "ssid"
	SITEROGUEEVENTSCOUNTDISTINCT_AP SiteRogueEventsCountDistinct = "ap"
	SITEROGUEEVENTSCOUNTDISTINCT_TYPE SiteRogueEventsCountDistinct = "type"
)

// All allowed values of SiteRogueEventsCountDistinct enum
var AllowedSiteRogueEventsCountDistinctEnumValues = []SiteRogueEventsCountDistinct{
	"",
	"bssid",
	"ssid",
	"ap",
	"type",
}

func (v *SiteRogueEventsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteRogueEventsCountDistinct(value)
	for _, existing := range AllowedSiteRogueEventsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteRogueEventsCountDistinct", value)
}

// NewSiteRogueEventsCountDistinctFromValue returns a pointer to a valid SiteRogueEventsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteRogueEventsCountDistinctFromValue(v string) (*SiteRogueEventsCountDistinct, error) {
	ev := SiteRogueEventsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteRogueEventsCountDistinct: valid values are %v", v, AllowedSiteRogueEventsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteRogueEventsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteRogueEventsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_rogue_events_count_distinct value
func (v SiteRogueEventsCountDistinct) Ptr() *SiteRogueEventsCountDistinct {
	return &v
}

type NullableSiteRogueEventsCountDistinct struct {
	value *SiteRogueEventsCountDistinct
	isSet bool
}

func (v NullableSiteRogueEventsCountDistinct) Get() *SiteRogueEventsCountDistinct {
	return v.value
}

func (v *NullableSiteRogueEventsCountDistinct) Set(val *SiteRogueEventsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteRogueEventsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteRogueEventsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteRogueEventsCountDistinct(val *SiteRogueEventsCountDistinct) *NullableSiteRogueEventsCountDistinct {
	return &NullableSiteRogueEventsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteRogueEventsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteRogueEventsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

