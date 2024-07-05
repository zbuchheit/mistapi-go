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
)

// checks if the TicketAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TicketAttachment{}

// TicketAttachment struct for TicketAttachment
type TicketAttachment struct {
	ContentUrl *string `json:"content_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TicketAttachment TicketAttachment

// NewTicketAttachment instantiates a new TicketAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketAttachment() *TicketAttachment {
	this := TicketAttachment{}
	return &this
}

// NewTicketAttachmentWithDefaults instantiates a new TicketAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketAttachmentWithDefaults() *TicketAttachment {
	this := TicketAttachment{}
	return &this
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise.
func (o *TicketAttachment) GetContentUrl() string {
	if o == nil || IsNil(o.ContentUrl) {
		var ret string
		return ret
	}
	return *o.ContentUrl
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketAttachment) GetContentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ContentUrl) {
		return nil, false
	}
	return o.ContentUrl, true
}

// HasContentUrl returns a boolean if a field has been set.
func (o *TicketAttachment) HasContentUrl() bool {
	if o != nil && !IsNil(o.ContentUrl) {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.
func (o *TicketAttachment) SetContentUrl(v string) {
	o.ContentUrl = &v
}

func (o TicketAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TicketAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentUrl) {
		toSerialize["content_url"] = o.ContentUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TicketAttachment) UnmarshalJSON(data []byte) (err error) {
	varTicketAttachment := _TicketAttachment{}

	err = json.Unmarshal(data, &varTicketAttachment)

	if err != nil {
		return err
	}

	*o = TicketAttachment(varTicketAttachment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTicketAttachment struct {
	value *TicketAttachment
	isSet bool
}

func (v NullableTicketAttachment) Get() *TicketAttachment {
	return v.value
}

func (v *NullableTicketAttachment) Set(val *TicketAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketAttachment(val *TicketAttachment) *NullableTicketAttachment {
	return &NullableTicketAttachment{value: val, isSet: true}
}

func (v NullableTicketAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


