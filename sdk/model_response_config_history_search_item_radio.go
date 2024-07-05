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

// checks if the ResponseConfigHistorySearchItemRadio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseConfigHistorySearchItemRadio{}

// ResponseConfigHistorySearchItemRadio struct for ResponseConfigHistorySearchItemRadio
type ResponseConfigHistorySearchItemRadio struct {
	Band string `json:"band"`
	Channel int32 `json:"channel"`
	AdditionalProperties map[string]interface{}
}

type _ResponseConfigHistorySearchItemRadio ResponseConfigHistorySearchItemRadio

// NewResponseConfigHistorySearchItemRadio instantiates a new ResponseConfigHistorySearchItemRadio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseConfigHistorySearchItemRadio(band string, channel int32) *ResponseConfigHistorySearchItemRadio {
	this := ResponseConfigHistorySearchItemRadio{}
	this.Band = band
	this.Channel = channel
	return &this
}

// NewResponseConfigHistorySearchItemRadioWithDefaults instantiates a new ResponseConfigHistorySearchItemRadio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseConfigHistorySearchItemRadioWithDefaults() *ResponseConfigHistorySearchItemRadio {
	this := ResponseConfigHistorySearchItemRadio{}
	return &this
}

// GetBand returns the Band field value
func (o *ResponseConfigHistorySearchItemRadio) GetBand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Band
}

// GetBandOk returns a tuple with the Band field value
// and a boolean to check if the value has been set.
func (o *ResponseConfigHistorySearchItemRadio) GetBandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Band, true
}

// SetBand sets field value
func (o *ResponseConfigHistorySearchItemRadio) SetBand(v string) {
	o.Band = v
}

// GetChannel returns the Channel field value
func (o *ResponseConfigHistorySearchItemRadio) GetChannel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *ResponseConfigHistorySearchItemRadio) GetChannelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *ResponseConfigHistorySearchItemRadio) SetChannel(v int32) {
	o.Channel = v
}

func (o ResponseConfigHistorySearchItemRadio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseConfigHistorySearchItemRadio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["band"] = o.Band
	toSerialize["channel"] = o.Channel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseConfigHistorySearchItemRadio) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"band",
		"channel",
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

	varResponseConfigHistorySearchItemRadio := _ResponseConfigHistorySearchItemRadio{}

	err = json.Unmarshal(data, &varResponseConfigHistorySearchItemRadio)

	if err != nil {
		return err
	}

	*o = ResponseConfigHistorySearchItemRadio(varResponseConfigHistorySearchItemRadio)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band")
		delete(additionalProperties, "channel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseConfigHistorySearchItemRadio struct {
	value *ResponseConfigHistorySearchItemRadio
	isSet bool
}

func (v NullableResponseConfigHistorySearchItemRadio) Get() *ResponseConfigHistorySearchItemRadio {
	return v.value
}

func (v *NullableResponseConfigHistorySearchItemRadio) Set(val *ResponseConfigHistorySearchItemRadio) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseConfigHistorySearchItemRadio) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseConfigHistorySearchItemRadio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseConfigHistorySearchItemRadio(val *ResponseConfigHistorySearchItemRadio) *NullableResponseConfigHistorySearchItemRadio {
	return &NullableResponseConfigHistorySearchItemRadio{value: val, isSet: true}
}

func (v NullableResponseConfigHistorySearchItemRadio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseConfigHistorySearchItemRadio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


