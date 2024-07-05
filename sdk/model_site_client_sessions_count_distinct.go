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

// SiteClientSessionsCountDistinct the model 'SiteClientSessionsCountDistinct'
type SiteClientSessionsCountDistinct string

// List of site_client_sessions_count_distinct
const (
	SITECLIENTSESSIONSCOUNTDISTINCT_EMPTY SiteClientSessionsCountDistinct = ""
	SITECLIENTSESSIONSCOUNTDISTINCT_SSID SiteClientSessionsCountDistinct = "ssid"
	SITECLIENTSESSIONSCOUNTDISTINCT_WLAN_ID SiteClientSessionsCountDistinct = "wlan_id"
	SITECLIENTSESSIONSCOUNTDISTINCT_AP SiteClientSessionsCountDistinct = "ap"
	SITECLIENTSESSIONSCOUNTDISTINCT_MAC SiteClientSessionsCountDistinct = "mac"
	SITECLIENTSESSIONSCOUNTDISTINCT_CLIENT_FAMILY SiteClientSessionsCountDistinct = "client_family"
	SITECLIENTSESSIONSCOUNTDISTINCT_CLIENT_MANUFACTURE SiteClientSessionsCountDistinct = "client_manufacture"
	SITECLIENTSESSIONSCOUNTDISTINCT_CLIENT_MODEL SiteClientSessionsCountDistinct = "client_model"
	SITECLIENTSESSIONSCOUNTDISTINCT_CLIENT_OS SiteClientSessionsCountDistinct = "client_os"
)

// All allowed values of SiteClientSessionsCountDistinct enum
var AllowedSiteClientSessionsCountDistinctEnumValues = []SiteClientSessionsCountDistinct{
	"",
	"ssid",
	"wlan_id",
	"ap",
	"mac",
	"client_family",
	"client_manufacture",
	"client_model",
	"client_os",
}

func (v *SiteClientSessionsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteClientSessionsCountDistinct(value)
	for _, existing := range AllowedSiteClientSessionsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteClientSessionsCountDistinct", value)
}

// NewSiteClientSessionsCountDistinctFromValue returns a pointer to a valid SiteClientSessionsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteClientSessionsCountDistinctFromValue(v string) (*SiteClientSessionsCountDistinct, error) {
	ev := SiteClientSessionsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteClientSessionsCountDistinct: valid values are %v", v, AllowedSiteClientSessionsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteClientSessionsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteClientSessionsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_client_sessions_count_distinct value
func (v SiteClientSessionsCountDistinct) Ptr() *SiteClientSessionsCountDistinct {
	return &v
}

type NullableSiteClientSessionsCountDistinct struct {
	value *SiteClientSessionsCountDistinct
	isSet bool
}

func (v NullableSiteClientSessionsCountDistinct) Get() *SiteClientSessionsCountDistinct {
	return v.value
}

func (v *NullableSiteClientSessionsCountDistinct) Set(val *SiteClientSessionsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteClientSessionsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteClientSessionsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteClientSessionsCountDistinct(val *SiteClientSessionsCountDistinct) *NullableSiteClientSessionsCountDistinct {
	return &NullableSiteClientSessionsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteClientSessionsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteClientSessionsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

