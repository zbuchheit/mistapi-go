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

// WlanAuthType the model 'WlanAuthType'
type WlanAuthType string

// List of wlan_auth_type
const (
	WLANAUTHTYPE_EMPTY WlanAuthType = ""
	WLANAUTHTYPE_OPEN WlanAuthType = "open"
	WLANAUTHTYPE_PSK WlanAuthType = "psk"
	WLANAUTHTYPE_WEP WlanAuthType = "wep"
	WLANAUTHTYPE_EAP WlanAuthType = "eap"
	WLANAUTHTYPE_EAP192 WlanAuthType = "eap192"
	WLANAUTHTYPE_PSK_TKIP WlanAuthType = "psk-tkip"
	WLANAUTHTYPE_PSK_WPA2_TKIP WlanAuthType = "psk-wpa2-tkip"
)

// All allowed values of WlanAuthType enum
var AllowedWlanAuthTypeEnumValues = []WlanAuthType{
	"",
	"open",
	"psk",
	"wep",
	"eap",
	"eap192",
	"psk-tkip",
	"psk-wpa2-tkip",
}

func (v *WlanAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanAuthType(value)
	for _, existing := range AllowedWlanAuthTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanAuthType", value)
}

// NewWlanAuthTypeFromValue returns a pointer to a valid WlanAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanAuthTypeFromValue(v string) (*WlanAuthType, error) {
	ev := WlanAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanAuthType: valid values are %v", v, AllowedWlanAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanAuthType) IsValid() bool {
	for _, existing := range AllowedWlanAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wlan_auth_type value
func (v WlanAuthType) Ptr() *WlanAuthType {
	return &v
}

type NullableWlanAuthType struct {
	value *WlanAuthType
	isSet bool
}

func (v NullableWlanAuthType) Get() *WlanAuthType {
	return v.value
}

func (v *NullableWlanAuthType) Set(val *WlanAuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanAuthType(val *WlanAuthType) *NullableWlanAuthType {
	return &NullableWlanAuthType{value: val, isSet: true}
}

func (v NullableWlanAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

