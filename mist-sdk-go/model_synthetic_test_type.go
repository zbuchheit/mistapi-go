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

// SyntheticTestType the model 'SyntheticTestType'
type SyntheticTestType string

// List of synthetic_test_type
const (
	SYNTHETICTESTTYPE_EMPTY SyntheticTestType = ""
	SYNTHETICTESTTYPE_DNS SyntheticTestType = "dns"
	SYNTHETICTESTTYPE_ARP SyntheticTestType = "arp"
	SYNTHETICTESTTYPE_DHCP SyntheticTestType = "dhcp"
	SYNTHETICTESTTYPE_CURL SyntheticTestType = "curl"
	SYNTHETICTESTTYPE_RADIUS SyntheticTestType = "radius"
	SYNTHETICTESTTYPE_SPEEDTEST SyntheticTestType = "speedtest"
	SYNTHETICTESTTYPE_DHCP6 SyntheticTestType = "dhcp6"
)

// All allowed values of SyntheticTestType enum
var AllowedSyntheticTestTypeEnumValues = []SyntheticTestType{
	"",
	"dns",
	"arp",
	"dhcp",
	"curl",
	"radius",
	"speedtest",
	"dhcp6",
}

func (v *SyntheticTestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyntheticTestType(value)
	for _, existing := range AllowedSyntheticTestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyntheticTestType", value)
}

// NewSyntheticTestTypeFromValue returns a pointer to a valid SyntheticTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyntheticTestTypeFromValue(v string) (*SyntheticTestType, error) {
	ev := SyntheticTestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyntheticTestType: valid values are %v", v, AllowedSyntheticTestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyntheticTestType) IsValid() bool {
	for _, existing := range AllowedSyntheticTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to synthetic_test_type value
func (v SyntheticTestType) Ptr() *SyntheticTestType {
	return &v
}

type NullableSyntheticTestType struct {
	value *SyntheticTestType
	isSet bool
}

func (v NullableSyntheticTestType) Get() *SyntheticTestType {
	return v.value
}

func (v *NullableSyntheticTestType) Set(val *SyntheticTestType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTestType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTestType(val *SyntheticTestType) *NullableSyntheticTestType {
	return &NullableSyntheticTestType{value: val, isSet: true}
}

func (v NullableSyntheticTestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

