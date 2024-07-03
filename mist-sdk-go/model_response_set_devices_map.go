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
)

// checks if the ResponseSetDevicesMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSetDevicesMap{}

// ResponseSetDevicesMap struct for ResponseSetDevicesMap
type ResponseSetDevicesMap struct {
	Locked []string `json:"locked,omitempty"`
	Moved []string `json:"moved,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSetDevicesMap ResponseSetDevicesMap

// NewResponseSetDevicesMap instantiates a new ResponseSetDevicesMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSetDevicesMap() *ResponseSetDevicesMap {
	this := ResponseSetDevicesMap{}
	return &this
}

// NewResponseSetDevicesMapWithDefaults instantiates a new ResponseSetDevicesMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSetDevicesMapWithDefaults() *ResponseSetDevicesMap {
	this := ResponseSetDevicesMap{}
	return &this
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *ResponseSetDevicesMap) GetLocked() []string {
	if o == nil || IsNil(o.Locked) {
		var ret []string
		return ret
	}
	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSetDevicesMap) GetLockedOk() ([]string, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *ResponseSetDevicesMap) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given []string and assigns it to the Locked field.
func (o *ResponseSetDevicesMap) SetLocked(v []string) {
	o.Locked = v
}

// GetMoved returns the Moved field value if set, zero value otherwise.
func (o *ResponseSetDevicesMap) GetMoved() []string {
	if o == nil || IsNil(o.Moved) {
		var ret []string
		return ret
	}
	return o.Moved
}

// GetMovedOk returns a tuple with the Moved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSetDevicesMap) GetMovedOk() ([]string, bool) {
	if o == nil || IsNil(o.Moved) {
		return nil, false
	}
	return o.Moved, true
}

// HasMoved returns a boolean if a field has been set.
func (o *ResponseSetDevicesMap) HasMoved() bool {
	if o != nil && !IsNil(o.Moved) {
		return true
	}

	return false
}

// SetMoved gets a reference to the given []string and assigns it to the Moved field.
func (o *ResponseSetDevicesMap) SetMoved(v []string) {
	o.Moved = v
}

func (o ResponseSetDevicesMap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSetDevicesMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Moved) {
		toSerialize["moved"] = o.Moved
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSetDevicesMap) UnmarshalJSON(data []byte) (err error) {
	varResponseSetDevicesMap := _ResponseSetDevicesMap{}

	err = json.Unmarshal(data, &varResponseSetDevicesMap)

	if err != nil {
		return err
	}

	*o = ResponseSetDevicesMap(varResponseSetDevicesMap)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "locked")
		delete(additionalProperties, "moved")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSetDevicesMap struct {
	value *ResponseSetDevicesMap
	isSet bool
}

func (v NullableResponseSetDevicesMap) Get() *ResponseSetDevicesMap {
	return v.value
}

func (v *NullableResponseSetDevicesMap) Set(val *ResponseSetDevicesMap) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSetDevicesMap) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSetDevicesMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSetDevicesMap(val *ResponseSetDevicesMap) *NullableResponseSetDevicesMap {
	return &NullableResponseSetDevicesMap{value: val, isSet: true}
}

func (v NullableResponseSetDevicesMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSetDevicesMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


