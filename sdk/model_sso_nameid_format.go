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

// SsoNameidFormat if `idp_type`==`saml`
type SsoNameidFormat string

// List of sso_nameid_format
const (
	SSONAMEIDFORMAT_EMPTY SsoNameidFormat = ""
	SSONAMEIDFORMAT_EMAIL SsoNameidFormat = "email"
	SSONAMEIDFORMAT_UNSPECIFIED SsoNameidFormat = "unspecified"
)

// All allowed values of SsoNameidFormat enum
var AllowedSsoNameidFormatEnumValues = []SsoNameidFormat{
	"",
	"email",
	"unspecified",
}

func (v *SsoNameidFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SsoNameidFormat(value)
	for _, existing := range AllowedSsoNameidFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SsoNameidFormat", value)
}

// NewSsoNameidFormatFromValue returns a pointer to a valid SsoNameidFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSsoNameidFormatFromValue(v string) (*SsoNameidFormat, error) {
	ev := SsoNameidFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SsoNameidFormat: valid values are %v", v, AllowedSsoNameidFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SsoNameidFormat) IsValid() bool {
	for _, existing := range AllowedSsoNameidFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sso_nameid_format value
func (v SsoNameidFormat) Ptr() *SsoNameidFormat {
	return &v
}

type NullableSsoNameidFormat struct {
	value *SsoNameidFormat
	isSet bool
}

func (v NullableSsoNameidFormat) Get() *SsoNameidFormat {
	return v.value
}

func (v *NullableSsoNameidFormat) Set(val *SsoNameidFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoNameidFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoNameidFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoNameidFormat(val *SsoNameidFormat) *NullableSsoNameidFormat {
	return &NullableSsoNameidFormat{value: val, isSet: true}
}

func (v NullableSsoNameidFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoNameidFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

