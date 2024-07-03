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

// checks if the ConstDeviceApBand6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstDeviceApBand6{}

// ConstDeviceApBand6 struct for ConstDeviceApBand6
type ConstDeviceApBand6 struct {
	MaxClients *int32 `json:"max_clients,omitempty"`
	MaxPower *int32 `json:"max_power,omitempty"`
	MinPower *int32 `json:"min_power,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstDeviceApBand6 ConstDeviceApBand6

// NewConstDeviceApBand6 instantiates a new ConstDeviceApBand6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstDeviceApBand6() *ConstDeviceApBand6 {
	this := ConstDeviceApBand6{}
	return &this
}

// NewConstDeviceApBand6WithDefaults instantiates a new ConstDeviceApBand6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstDeviceApBand6WithDefaults() *ConstDeviceApBand6 {
	this := ConstDeviceApBand6{}
	return &this
}

// GetMaxClients returns the MaxClients field value if set, zero value otherwise.
func (o *ConstDeviceApBand6) GetMaxClients() int32 {
	if o == nil || IsNil(o.MaxClients) {
		var ret int32
		return ret
	}
	return *o.MaxClients
}

// GetMaxClientsOk returns a tuple with the MaxClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApBand6) GetMaxClientsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxClients) {
		return nil, false
	}
	return o.MaxClients, true
}

// HasMaxClients returns a boolean if a field has been set.
func (o *ConstDeviceApBand6) HasMaxClients() bool {
	if o != nil && !IsNil(o.MaxClients) {
		return true
	}

	return false
}

// SetMaxClients gets a reference to the given int32 and assigns it to the MaxClients field.
func (o *ConstDeviceApBand6) SetMaxClients(v int32) {
	o.MaxClients = &v
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *ConstDeviceApBand6) GetMaxPower() int32 {
	if o == nil || IsNil(o.MaxPower) {
		var ret int32
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApBand6) GetMaxPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPower) {
		return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *ConstDeviceApBand6) HasMaxPower() bool {
	if o != nil && !IsNil(o.MaxPower) {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given int32 and assigns it to the MaxPower field.
func (o *ConstDeviceApBand6) SetMaxPower(v int32) {
	o.MaxPower = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *ConstDeviceApBand6) GetMinPower() int32 {
	if o == nil || IsNil(o.MinPower) {
		var ret int32
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApBand6) GetMinPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.MinPower) {
		return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *ConstDeviceApBand6) HasMinPower() bool {
	if o != nil && !IsNil(o.MinPower) {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given int32 and assigns it to the MinPower field.
func (o *ConstDeviceApBand6) SetMinPower(v int32) {
	o.MinPower = &v
}

func (o ConstDeviceApBand6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstDeviceApBand6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxClients) {
		toSerialize["max_clients"] = o.MaxClients
	}
	if !IsNil(o.MaxPower) {
		toSerialize["max_power"] = o.MaxPower
	}
	if !IsNil(o.MinPower) {
		toSerialize["min_power"] = o.MinPower
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstDeviceApBand6) UnmarshalJSON(data []byte) (err error) {
	varConstDeviceApBand6 := _ConstDeviceApBand6{}

	err = json.Unmarshal(data, &varConstDeviceApBand6)

	if err != nil {
		return err
	}

	*o = ConstDeviceApBand6(varConstDeviceApBand6)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max_clients")
		delete(additionalProperties, "max_power")
		delete(additionalProperties, "min_power")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstDeviceApBand6 struct {
	value *ConstDeviceApBand6
	isSet bool
}

func (v NullableConstDeviceApBand6) Get() *ConstDeviceApBand6 {
	return v.value
}

func (v *NullableConstDeviceApBand6) Set(val *ConstDeviceApBand6) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceApBand6) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceApBand6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceApBand6(val *ConstDeviceApBand6) *NullableConstDeviceApBand6 {
	return &NullableConstDeviceApBand6{value: val, isSet: true}
}

func (v NullableConstDeviceApBand6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceApBand6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


