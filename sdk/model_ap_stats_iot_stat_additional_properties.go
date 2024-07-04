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
)

// checks if the ApStatsIotStatAdditionalProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatsIotStatAdditionalProperties{}

// ApStatsIotStatAdditionalProperties struct for ApStatsIotStatAdditionalProperties
type ApStatsIotStatAdditionalProperties struct {
	Value *int32 `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatsIotStatAdditionalProperties ApStatsIotStatAdditionalProperties

// NewApStatsIotStatAdditionalProperties instantiates a new ApStatsIotStatAdditionalProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatsIotStatAdditionalProperties() *ApStatsIotStatAdditionalProperties {
	this := ApStatsIotStatAdditionalProperties{}
	return &this
}

// NewApStatsIotStatAdditionalPropertiesWithDefaults instantiates a new ApStatsIotStatAdditionalProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatsIotStatAdditionalPropertiesWithDefaults() *ApStatsIotStatAdditionalProperties {
	this := ApStatsIotStatAdditionalProperties{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApStatsIotStatAdditionalProperties) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsIotStatAdditionalProperties) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApStatsIotStatAdditionalProperties) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *ApStatsIotStatAdditionalProperties) SetValue(v int32) {
	o.Value = &v
}

func (o ApStatsIotStatAdditionalProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatsIotStatAdditionalProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApStatsIotStatAdditionalProperties) UnmarshalJSON(data []byte) (err error) {
	varApStatsIotStatAdditionalProperties := _ApStatsIotStatAdditionalProperties{}

	err = json.Unmarshal(data, &varApStatsIotStatAdditionalProperties)

	if err != nil {
		return err
	}

	*o = ApStatsIotStatAdditionalProperties(varApStatsIotStatAdditionalProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatsIotStatAdditionalProperties struct {
	value *ApStatsIotStatAdditionalProperties
	isSet bool
}

func (v NullableApStatsIotStatAdditionalProperties) Get() *ApStatsIotStatAdditionalProperties {
	return v.value
}

func (v *NullableApStatsIotStatAdditionalProperties) Set(val *ApStatsIotStatAdditionalProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsIotStatAdditionalProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsIotStatAdditionalProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsIotStatAdditionalProperties(val *ApStatsIotStatAdditionalProperties) *NullableApStatsIotStatAdditionalProperties {
	return &NullableApStatsIotStatAdditionalProperties{value: val, isSet: true}
}

func (v NullableApStatsIotStatAdditionalProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsIotStatAdditionalProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


