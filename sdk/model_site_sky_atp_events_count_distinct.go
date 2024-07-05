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

// SiteSkyAtpEventsCountDistinct the model 'SiteSkyAtpEventsCountDistinct'
type SiteSkyAtpEventsCountDistinct string

// List of site_sky_atp_events_count_distinct
const (
	SITESKYATPEVENTSCOUNTDISTINCT_EMPTY SiteSkyAtpEventsCountDistinct = ""
	SITESKYATPEVENTSCOUNTDISTINCT_TYPE SiteSkyAtpEventsCountDistinct = "type"
	SITESKYATPEVENTSCOUNTDISTINCT_MAC SiteSkyAtpEventsCountDistinct = "mac"
	SITESKYATPEVENTSCOUNTDISTINCT_DEVICE_MAC SiteSkyAtpEventsCountDistinct = "device_mac"
	SITESKYATPEVENTSCOUNTDISTINCT_THREAT_LEVEL SiteSkyAtpEventsCountDistinct = "threat_level"
)

// All allowed values of SiteSkyAtpEventsCountDistinct enum
var AllowedSiteSkyAtpEventsCountDistinctEnumValues = []SiteSkyAtpEventsCountDistinct{
	"",
	"type",
	"mac",
	"device_mac",
	"threat_level",
}

func (v *SiteSkyAtpEventsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSkyAtpEventsCountDistinct(value)
	for _, existing := range AllowedSiteSkyAtpEventsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSkyAtpEventsCountDistinct", value)
}

// NewSiteSkyAtpEventsCountDistinctFromValue returns a pointer to a valid SiteSkyAtpEventsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSkyAtpEventsCountDistinctFromValue(v string) (*SiteSkyAtpEventsCountDistinct, error) {
	ev := SiteSkyAtpEventsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSkyAtpEventsCountDistinct: valid values are %v", v, AllowedSiteSkyAtpEventsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSkyAtpEventsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteSkyAtpEventsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sky_atp_events_count_distinct value
func (v SiteSkyAtpEventsCountDistinct) Ptr() *SiteSkyAtpEventsCountDistinct {
	return &v
}

type NullableSiteSkyAtpEventsCountDistinct struct {
	value *SiteSkyAtpEventsCountDistinct
	isSet bool
}

func (v NullableSiteSkyAtpEventsCountDistinct) Get() *SiteSkyAtpEventsCountDistinct {
	return v.value
}

func (v *NullableSiteSkyAtpEventsCountDistinct) Set(val *SiteSkyAtpEventsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSkyAtpEventsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSkyAtpEventsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSkyAtpEventsCountDistinct(val *SiteSkyAtpEventsCountDistinct) *NullableSiteSkyAtpEventsCountDistinct {
	return &NullableSiteSkyAtpEventsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteSkyAtpEventsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSkyAtpEventsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

