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

// WebhookDeliveryTopic webhook topic
type WebhookDeliveryTopic string

// List of webhook_delivery_topic
const (
	WEBHOOKDELIVERYTOPIC_EMPTY WebhookDeliveryTopic = ""
	WEBHOOKDELIVERYTOPIC_ALARMS WebhookDeliveryTopic = "alarms"
	WEBHOOKDELIVERYTOPIC_AUDITS WebhookDeliveryTopic = "audits"
	WEBHOOKDELIVERYTOPIC_DEVICE_UPDOWNS WebhookDeliveryTopic = "device-updowns"
	WEBHOOKDELIVERYTOPIC_OCCUPANCY_ALERTS WebhookDeliveryTopic = "occupancy-alerts"
	WEBHOOKDELIVERYTOPIC_PING WebhookDeliveryTopic = "ping"
)

// All allowed values of WebhookDeliveryTopic enum
var AllowedWebhookDeliveryTopicEnumValues = []WebhookDeliveryTopic{
	"",
	"alarms",
	"audits",
	"device-updowns",
	"occupancy-alerts",
	"ping",
}

func (v *WebhookDeliveryTopic) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookDeliveryTopic(value)
	for _, existing := range AllowedWebhookDeliveryTopicEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookDeliveryTopic", value)
}

// NewWebhookDeliveryTopicFromValue returns a pointer to a valid WebhookDeliveryTopic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookDeliveryTopicFromValue(v string) (*WebhookDeliveryTopic, error) {
	ev := WebhookDeliveryTopic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookDeliveryTopic: valid values are %v", v, AllowedWebhookDeliveryTopicEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookDeliveryTopic) IsValid() bool {
	for _, existing := range AllowedWebhookDeliveryTopicEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to webhook_delivery_topic value
func (v WebhookDeliveryTopic) Ptr() *WebhookDeliveryTopic {
	return &v
}

type NullableWebhookDeliveryTopic struct {
	value *WebhookDeliveryTopic
	isSet bool
}

func (v NullableWebhookDeliveryTopic) Get() *WebhookDeliveryTopic {
	return v.value
}

func (v *NullableWebhookDeliveryTopic) Set(val *WebhookDeliveryTopic) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDeliveryTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDeliveryTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDeliveryTopic(val *WebhookDeliveryTopic) *NullableWebhookDeliveryTopic {
	return &NullableWebhookDeliveryTopic{value: val, isSet: true}
}

func (v NullableWebhookDeliveryTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDeliveryTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

