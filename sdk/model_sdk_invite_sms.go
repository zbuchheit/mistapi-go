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

// checks if the SdkInviteSms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdkInviteSms{}

// SdkInviteSms struct for SdkInviteSms
type SdkInviteSms struct {
	Number string `json:"number"`
	AdditionalProperties map[string]interface{}
}

type _SdkInviteSms SdkInviteSms

// NewSdkInviteSms instantiates a new SdkInviteSms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkInviteSms(number string) *SdkInviteSms {
	this := SdkInviteSms{}
	this.Number = number
	return &this
}

// NewSdkInviteSmsWithDefaults instantiates a new SdkInviteSms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkInviteSmsWithDefaults() *SdkInviteSms {
	this := SdkInviteSms{}
	return &this
}

// GetNumber returns the Number field value
func (o *SdkInviteSms) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *SdkInviteSms) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *SdkInviteSms) SetNumber(v string) {
	o.Number = v
}

func (o SdkInviteSms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdkInviteSms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SdkInviteSms) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"number",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSdkInviteSms := _SdkInviteSms{}

	err = json.Unmarshal(data, &varSdkInviteSms)

	if err != nil {
		return err
	}

	*o = SdkInviteSms(varSdkInviteSms)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdkInviteSms struct {
	value *SdkInviteSms
	isSet bool
}

func (v NullableSdkInviteSms) Get() *SdkInviteSms {
	return v.value
}

func (v *NullableSdkInviteSms) Set(val *SdkInviteSms) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkInviteSms) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkInviteSms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkInviteSms(val *SdkInviteSms) *NullableSdkInviteSms {
	return &NullableSdkInviteSms{value: val, isSet: true}
}

func (v NullableSdkInviteSms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkInviteSms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


