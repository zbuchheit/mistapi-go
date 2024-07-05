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

// checks if the ConstDeviceApVble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstDeviceApVble{}

// ConstDeviceApVble struct for ConstDeviceApVble
type ConstDeviceApVble struct {
	BeaconRate *int32 `json:"beacon_rate,omitempty"`
	Beams *int32 `json:"beams,omitempty"`
	Power *int32 `json:"power,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstDeviceApVble ConstDeviceApVble

// NewConstDeviceApVble instantiates a new ConstDeviceApVble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstDeviceApVble() *ConstDeviceApVble {
	this := ConstDeviceApVble{}
	return &this
}

// NewConstDeviceApVbleWithDefaults instantiates a new ConstDeviceApVble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstDeviceApVbleWithDefaults() *ConstDeviceApVble {
	this := ConstDeviceApVble{}
	return &this
}

// GetBeaconRate returns the BeaconRate field value if set, zero value otherwise.
func (o *ConstDeviceApVble) GetBeaconRate() int32 {
	if o == nil || IsNil(o.BeaconRate) {
		var ret int32
		return ret
	}
	return *o.BeaconRate
}

// GetBeaconRateOk returns a tuple with the BeaconRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApVble) GetBeaconRateOk() (*int32, bool) {
	if o == nil || IsNil(o.BeaconRate) {
		return nil, false
	}
	return o.BeaconRate, true
}

// HasBeaconRate returns a boolean if a field has been set.
func (o *ConstDeviceApVble) HasBeaconRate() bool {
	if o != nil && !IsNil(o.BeaconRate) {
		return true
	}

	return false
}

// SetBeaconRate gets a reference to the given int32 and assigns it to the BeaconRate field.
func (o *ConstDeviceApVble) SetBeaconRate(v int32) {
	o.BeaconRate = &v
}

// GetBeams returns the Beams field value if set, zero value otherwise.
func (o *ConstDeviceApVble) GetBeams() int32 {
	if o == nil || IsNil(o.Beams) {
		var ret int32
		return ret
	}
	return *o.Beams
}

// GetBeamsOk returns a tuple with the Beams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApVble) GetBeamsOk() (*int32, bool) {
	if o == nil || IsNil(o.Beams) {
		return nil, false
	}
	return o.Beams, true
}

// HasBeams returns a boolean if a field has been set.
func (o *ConstDeviceApVble) HasBeams() bool {
	if o != nil && !IsNil(o.Beams) {
		return true
	}

	return false
}

// SetBeams gets a reference to the given int32 and assigns it to the Beams field.
func (o *ConstDeviceApVble) SetBeams(v int32) {
	o.Beams = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *ConstDeviceApVble) GetPower() int32 {
	if o == nil || IsNil(o.Power) {
		var ret int32
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceApVble) GetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *ConstDeviceApVble) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given int32 and assigns it to the Power field.
func (o *ConstDeviceApVble) SetPower(v int32) {
	o.Power = &v
}

func (o ConstDeviceApVble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstDeviceApVble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeaconRate) {
		toSerialize["beacon_rate"] = o.BeaconRate
	}
	if !IsNil(o.Beams) {
		toSerialize["beams"] = o.Beams
	}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstDeviceApVble) UnmarshalJSON(data []byte) (err error) {
	varConstDeviceApVble := _ConstDeviceApVble{}

	err = json.Unmarshal(data, &varConstDeviceApVble)

	if err != nil {
		return err
	}

	*o = ConstDeviceApVble(varConstDeviceApVble)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "beacon_rate")
		delete(additionalProperties, "beams")
		delete(additionalProperties, "power")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstDeviceApVble struct {
	value *ConstDeviceApVble
	isSet bool
}

func (v NullableConstDeviceApVble) Get() *ConstDeviceApVble {
	return v.value
}

func (v *NullableConstDeviceApVble) Set(val *ConstDeviceApVble) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceApVble) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceApVble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceApVble(val *ConstDeviceApVble) *NullableConstDeviceApVble {
	return &NullableConstDeviceApVble{value: val, isSet: true}
}

func (v NullableConstDeviceApVble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceApVble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


