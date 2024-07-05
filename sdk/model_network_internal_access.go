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

// checks if the NetworkInternalAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkInternalAccess{}

// NetworkInternalAccess struct for NetworkInternalAccess
type NetworkInternalAccess struct {
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkInternalAccess NetworkInternalAccess

// NewNetworkInternalAccess instantiates a new NetworkInternalAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInternalAccess() *NetworkInternalAccess {
	this := NetworkInternalAccess{}
	return &this
}

// NewNetworkInternalAccessWithDefaults instantiates a new NetworkInternalAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInternalAccessWithDefaults() *NetworkInternalAccess {
	this := NetworkInternalAccess{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkInternalAccess) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInternalAccess) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworkInternalAccess) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkInternalAccess) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NetworkInternalAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkInternalAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkInternalAccess) UnmarshalJSON(data []byte) (err error) {
	varNetworkInternalAccess := _NetworkInternalAccess{}

	err = json.Unmarshal(data, &varNetworkInternalAccess)

	if err != nil {
		return err
	}

	*o = NetworkInternalAccess(varNetworkInternalAccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkInternalAccess struct {
	value *NetworkInternalAccess
	isSet bool
}

func (v NullableNetworkInternalAccess) Get() *NetworkInternalAccess {
	return v.value
}

func (v *NullableNetworkInternalAccess) Set(val *NetworkInternalAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInternalAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInternalAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInternalAccess(val *NetworkInternalAccess) *NullableNetworkInternalAccess {
	return &NullableNetworkInternalAccess{value: val, isSet: true}
}

func (v NullableNetworkInternalAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInternalAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


