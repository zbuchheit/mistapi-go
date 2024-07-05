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

// WlanPortalSmsProvider the model 'WlanPortalSmsProvider'
type WlanPortalSmsProvider string

// List of wlan_portal_sms_provider
const (
	WLANPORTALSMSPROVIDER_EMPTY WlanPortalSmsProvider = ""
	WLANPORTALSMSPROVIDER_MANUAL WlanPortalSmsProvider = "manual"
	WLANPORTALSMSPROVIDER_TWILIO WlanPortalSmsProvider = "twilio"
	WLANPORTALSMSPROVIDER_BROADNET WlanPortalSmsProvider = "broadnet"
	WLANPORTALSMSPROVIDER_CLICKATELL WlanPortalSmsProvider = "clickatell"
	WLANPORTALSMSPROVIDER_PUZZEL WlanPortalSmsProvider = "puzzel"
	WLANPORTALSMSPROVIDER_GUPSHUP WlanPortalSmsProvider = "gupshup"
	WLANPORTALSMSPROVIDER_TELSTRA WlanPortalSmsProvider = "telstra"
)

// All allowed values of WlanPortalSmsProvider enum
var AllowedWlanPortalSmsProviderEnumValues = []WlanPortalSmsProvider{
	"",
	"manual",
	"twilio",
	"broadnet",
	"clickatell",
	"puzzel",
	"gupshup",
	"telstra",
}

func (v *WlanPortalSmsProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanPortalSmsProvider(value)
	for _, existing := range AllowedWlanPortalSmsProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanPortalSmsProvider", value)
}

// NewWlanPortalSmsProviderFromValue returns a pointer to a valid WlanPortalSmsProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanPortalSmsProviderFromValue(v string) (*WlanPortalSmsProvider, error) {
	ev := WlanPortalSmsProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanPortalSmsProvider: valid values are %v", v, AllowedWlanPortalSmsProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanPortalSmsProvider) IsValid() bool {
	for _, existing := range AllowedWlanPortalSmsProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to wlan_portal_sms_provider value
func (v WlanPortalSmsProvider) Ptr() *WlanPortalSmsProvider {
	return &v
}

type NullableWlanPortalSmsProvider struct {
	value *WlanPortalSmsProvider
	isSet bool
}

func (v NullableWlanPortalSmsProvider) Get() *WlanPortalSmsProvider {
	return v.value
}

func (v *NullableWlanPortalSmsProvider) Set(val *WlanPortalSmsProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanPortalSmsProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanPortalSmsProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanPortalSmsProvider(val *WlanPortalSmsProvider) *NullableWlanPortalSmsProvider {
	return &NullableWlanPortalSmsProvider{value: val, isSet: true}
}

func (v NullableWlanPortalSmsProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanPortalSmsProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

