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

// checks if the ConstMxedgeModelPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstMxedgeModelPort{}

// ConstMxedgeModelPort struct for ConstMxedgeModelPort
type ConstMxedgeModelPort struct {
	Display *string `json:"display,omitempty"`
	Speed *int32 `json:"speed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstMxedgeModelPort ConstMxedgeModelPort

// NewConstMxedgeModelPort instantiates a new ConstMxedgeModelPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstMxedgeModelPort() *ConstMxedgeModelPort {
	this := ConstMxedgeModelPort{}
	return &this
}

// NewConstMxedgeModelPortWithDefaults instantiates a new ConstMxedgeModelPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstMxedgeModelPortWithDefaults() *ConstMxedgeModelPort {
	this := ConstMxedgeModelPort{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ConstMxedgeModelPort) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstMxedgeModelPort) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ConstMxedgeModelPort) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ConstMxedgeModelPort) SetDisplay(v string) {
	o.Display = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *ConstMxedgeModelPort) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstMxedgeModelPort) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *ConstMxedgeModelPort) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *ConstMxedgeModelPort) SetSpeed(v int32) {
	o.Speed = &v
}

func (o ConstMxedgeModelPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstMxedgeModelPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstMxedgeModelPort) UnmarshalJSON(data []byte) (err error) {
	varConstMxedgeModelPort := _ConstMxedgeModelPort{}

	err = json.Unmarshal(data, &varConstMxedgeModelPort)

	if err != nil {
		return err
	}

	*o = ConstMxedgeModelPort(varConstMxedgeModelPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "speed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstMxedgeModelPort struct {
	value *ConstMxedgeModelPort
	isSet bool
}

func (v NullableConstMxedgeModelPort) Get() *ConstMxedgeModelPort {
	return v.value
}

func (v *NullableConstMxedgeModelPort) Set(val *ConstMxedgeModelPort) {
	v.value = val
	v.isSet = true
}

func (v NullableConstMxedgeModelPort) IsSet() bool {
	return v.isSet
}

func (v *NullableConstMxedgeModelPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstMxedgeModelPort(val *ConstMxedgeModelPort) *NullableConstMxedgeModelPort {
	return &NullableConstMxedgeModelPort{value: val, isSet: true}
}

func (v NullableConstMxedgeModelPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstMxedgeModelPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


