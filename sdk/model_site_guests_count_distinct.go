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

// SiteGuestsCountDistinct the model 'SiteGuestsCountDistinct'
type SiteGuestsCountDistinct string

// List of site_guests_count_distinct
const (
	SITEGUESTSCOUNTDISTINCT_EMPTY SiteGuestsCountDistinct = ""
	SITEGUESTSCOUNTDISTINCT_AUTH_METHOD SiteGuestsCountDistinct = "auth_method"
	SITEGUESTSCOUNTDISTINCT_SSID SiteGuestsCountDistinct = "ssid"
	SITEGUESTSCOUNTDISTINCT_COMPANY SiteGuestsCountDistinct = "company"
)

// All allowed values of SiteGuestsCountDistinct enum
var AllowedSiteGuestsCountDistinctEnumValues = []SiteGuestsCountDistinct{
	"",
	"auth_method",
	"ssid",
	"company",
}

func (v *SiteGuestsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteGuestsCountDistinct(value)
	for _, existing := range AllowedSiteGuestsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteGuestsCountDistinct", value)
}

// NewSiteGuestsCountDistinctFromValue returns a pointer to a valid SiteGuestsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteGuestsCountDistinctFromValue(v string) (*SiteGuestsCountDistinct, error) {
	ev := SiteGuestsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteGuestsCountDistinct: valid values are %v", v, AllowedSiteGuestsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteGuestsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteGuestsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_guests_count_distinct value
func (v SiteGuestsCountDistinct) Ptr() *SiteGuestsCountDistinct {
	return &v
}

type NullableSiteGuestsCountDistinct struct {
	value *SiteGuestsCountDistinct
	isSet bool
}

func (v NullableSiteGuestsCountDistinct) Get() *SiteGuestsCountDistinct {
	return v.value
}

func (v *NullableSiteGuestsCountDistinct) Set(val *SiteGuestsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteGuestsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteGuestsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteGuestsCountDistinct(val *SiteGuestsCountDistinct) *NullableSiteGuestsCountDistinct {
	return &NullableSiteGuestsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteGuestsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteGuestsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

