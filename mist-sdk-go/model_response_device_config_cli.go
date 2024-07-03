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

// checks if the ResponseDeviceConfigCli type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeviceConfigCli{}

// ResponseDeviceConfigCli struct for ResponseDeviceConfigCli
type ResponseDeviceConfigCli struct {
	Cli []string `json:"cli"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDeviceConfigCli ResponseDeviceConfigCli

// NewResponseDeviceConfigCli instantiates a new ResponseDeviceConfigCli object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeviceConfigCli(cli []string) *ResponseDeviceConfigCli {
	this := ResponseDeviceConfigCli{}
	this.Cli = cli
	return &this
}

// NewResponseDeviceConfigCliWithDefaults instantiates a new ResponseDeviceConfigCli object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeviceConfigCliWithDefaults() *ResponseDeviceConfigCli {
	this := ResponseDeviceConfigCli{}
	return &this
}

// GetCli returns the Cli field value
func (o *ResponseDeviceConfigCli) GetCli() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cli
}

// GetCliOk returns a tuple with the Cli field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceConfigCli) GetCliOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cli, true
}

// SetCli sets field value
func (o *ResponseDeviceConfigCli) SetCli(v []string) {
	o.Cli = v
}

func (o ResponseDeviceConfigCli) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeviceConfigCli) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cli"] = o.Cli

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDeviceConfigCli) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cli",
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

	varResponseDeviceConfigCli := _ResponseDeviceConfigCli{}

	err = json.Unmarshal(data, &varResponseDeviceConfigCli)

	if err != nil {
		return err
	}

	*o = ResponseDeviceConfigCli(varResponseDeviceConfigCli)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cli")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDeviceConfigCli struct {
	value *ResponseDeviceConfigCli
	isSet bool
}

func (v NullableResponseDeviceConfigCli) Get() *ResponseDeviceConfigCli {
	return v.value
}

func (v *NullableResponseDeviceConfigCli) Set(val *ResponseDeviceConfigCli) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeviceConfigCli) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeviceConfigCli) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeviceConfigCli(val *ResponseDeviceConfigCli) *NullableResponseDeviceConfigCli {
	return &NullableResponseDeviceConfigCli{value: val, isSet: true}
}

func (v NullableResponseDeviceConfigCli) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeviceConfigCli) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


