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

// SiteClientsCountDistinct the model 'SiteClientsCountDistinct'
type SiteClientsCountDistinct string

// List of site_clients_count_distinct
const (
	SITECLIENTSCOUNTDISTINCT_EMPTY SiteClientsCountDistinct = ""
	SITECLIENTSCOUNTDISTINCT_SSID SiteClientsCountDistinct = "ssid"
	SITECLIENTSCOUNTDISTINCT_AP SiteClientsCountDistinct = "ap"
	SITECLIENTSCOUNTDISTINCT_IP SiteClientsCountDistinct = "ip"
	SITECLIENTSCOUNTDISTINCT_VLAN SiteClientsCountDistinct = "vlan"
	SITECLIENTSCOUNTDISTINCT_HOSTNAME SiteClientsCountDistinct = "hostname"
	SITECLIENTSCOUNTDISTINCT_OS SiteClientsCountDistinct = "os"
	SITECLIENTSCOUNTDISTINCT_MODEL SiteClientsCountDistinct = "model"
	SITECLIENTSCOUNTDISTINCT_DEVICE SiteClientsCountDistinct = "device"
)

// All allowed values of SiteClientsCountDistinct enum
var AllowedSiteClientsCountDistinctEnumValues = []SiteClientsCountDistinct{
	"",
	"ssid",
	"ap",
	"ip",
	"vlan",
	"hostname",
	"os",
	"model",
	"device",
}

func (v *SiteClientsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteClientsCountDistinct(value)
	for _, existing := range AllowedSiteClientsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteClientsCountDistinct", value)
}

// NewSiteClientsCountDistinctFromValue returns a pointer to a valid SiteClientsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteClientsCountDistinctFromValue(v string) (*SiteClientsCountDistinct, error) {
	ev := SiteClientsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteClientsCountDistinct: valid values are %v", v, AllowedSiteClientsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteClientsCountDistinct) IsValid() bool {
	for _, existing := range AllowedSiteClientsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_clients_count_distinct value
func (v SiteClientsCountDistinct) Ptr() *SiteClientsCountDistinct {
	return &v
}

type NullableSiteClientsCountDistinct struct {
	value *SiteClientsCountDistinct
	isSet bool
}

func (v NullableSiteClientsCountDistinct) Get() *SiteClientsCountDistinct {
	return v.value
}

func (v *NullableSiteClientsCountDistinct) Set(val *SiteClientsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteClientsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteClientsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteClientsCountDistinct(val *SiteClientsCountDistinct) *NullableSiteClientsCountDistinct {
	return &NullableSiteClientsCountDistinct{value: val, isSet: true}
}

func (v NullableSiteClientsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteClientsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

