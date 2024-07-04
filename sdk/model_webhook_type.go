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

// WebhookType the model 'WebhookType'
type WebhookType string

// List of webhook_type
const (
	WEBHOOKTYPE_EMPTY WebhookType = ""
	WEBHOOKTYPE_HTTP_POST WebhookType = "http-post"
	WEBHOOKTYPE_SPLUNK WebhookType = "splunk"
	WEBHOOKTYPE_GOOGLE_PUBSUB WebhookType = "google-pubsub"
	WEBHOOKTYPE_AWS_SNS WebhookType = "aws-sns"
	WEBHOOKTYPE_OAUTH2 WebhookType = "oauth2"
)

// All allowed values of WebhookType enum
var AllowedWebhookTypeEnumValues = []WebhookType{
	"",
	"http-post",
	"splunk",
	"google-pubsub",
	"aws-sns",
	"oauth2",
}

func (v *WebhookType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookType(value)
	for _, existing := range AllowedWebhookTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookType", value)
}

// NewWebhookTypeFromValue returns a pointer to a valid WebhookType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookTypeFromValue(v string) (*WebhookType, error) {
	ev := WebhookType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookType: valid values are %v", v, AllowedWebhookTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookType) IsValid() bool {
	for _, existing := range AllowedWebhookTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to webhook_type value
func (v WebhookType) Ptr() *WebhookType {
	return &v
}

type NullableWebhookType struct {
	value *WebhookType
	isSet bool
}

func (v NullableWebhookType) Get() *WebhookType {
	return v.value
}

func (v *NullableWebhookType) Set(val *WebhookType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookType(val *WebhookType) *NullableWebhookType {
	return &NullableWebhookType{value: val, isSet: true}
}

func (v NullableWebhookType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

