/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// WebhookOauth2GrantType required when `type`==`oauth2`
type WebhookOauth2GrantType string

// List of webhook_oauth2_grant_type
const (
	WEBHOOKOAUTH2GRANTTYPE_EMPTY WebhookOauth2GrantType = ""
	WEBHOOKOAUTH2GRANTTYPE_CLIENT_CREDENTIALS WebhookOauth2GrantType = "client_credentials"
	WEBHOOKOAUTH2GRANTTYPE_PASSWORD WebhookOauth2GrantType = "password"
)

// All allowed values of WebhookOauth2GrantType enum
var AllowedWebhookOauth2GrantTypeEnumValues = []WebhookOauth2GrantType{
	"",
	"client_credentials",
	"password",
}

func (v *WebhookOauth2GrantType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookOauth2GrantType(value)
	for _, existing := range AllowedWebhookOauth2GrantTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookOauth2GrantType", value)
}

// NewWebhookOauth2GrantTypeFromValue returns a pointer to a valid WebhookOauth2GrantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookOauth2GrantTypeFromValue(v string) (*WebhookOauth2GrantType, error) {
	ev := WebhookOauth2GrantType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookOauth2GrantType: valid values are %v", v, AllowedWebhookOauth2GrantTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookOauth2GrantType) IsValid() bool {
	for _, existing := range AllowedWebhookOauth2GrantTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to webhook_oauth2_grant_type value
func (v WebhookOauth2GrantType) Ptr() *WebhookOauth2GrantType {
	return &v
}

type NullableWebhookOauth2GrantType struct {
	value *WebhookOauth2GrantType
	isSet bool
}

func (v NullableWebhookOauth2GrantType) Get() *WebhookOauth2GrantType {
	return v.value
}

func (v *NullableWebhookOauth2GrantType) Set(val *WebhookOauth2GrantType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookOauth2GrantType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookOauth2GrantType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookOauth2GrantType(val *WebhookOauth2GrantType) *NullableWebhookOauth2GrantType {
	return &NullableWebhookOauth2GrantType{value: val, isSet: true}
}

func (v NullableWebhookOauth2GrantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookOauth2GrantType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

