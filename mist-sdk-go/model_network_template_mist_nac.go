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

// checks if the NetworkTemplateMistNac type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkTemplateMistNac{}

// NetworkTemplateMistNac enable mist_nac to use radsec
type NetworkTemplateMistNac struct {
	Enabled *bool `json:"enabled,omitempty"`
	Network *string `json:"network,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkTemplateMistNac NetworkTemplateMistNac

// NewNetworkTemplateMistNac instantiates a new NetworkTemplateMistNac object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkTemplateMistNac() *NetworkTemplateMistNac {
	this := NetworkTemplateMistNac{}
	return &this
}

// NewNetworkTemplateMistNacWithDefaults instantiates a new NetworkTemplateMistNac object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkTemplateMistNacWithDefaults() *NetworkTemplateMistNac {
	this := NetworkTemplateMistNac{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkTemplateMistNac) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplateMistNac) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworkTemplateMistNac) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkTemplateMistNac) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkTemplateMistNac) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplateMistNac) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkTemplateMistNac) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *NetworkTemplateMistNac) SetNetwork(v string) {
	o.Network = &v
}

func (o NetworkTemplateMistNac) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkTemplateMistNac) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkTemplateMistNac) UnmarshalJSON(data []byte) (err error) {
	varNetworkTemplateMistNac := _NetworkTemplateMistNac{}

	err = json.Unmarshal(data, &varNetworkTemplateMistNac)

	if err != nil {
		return err
	}

	*o = NetworkTemplateMistNac(varNetworkTemplateMistNac)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "network")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkTemplateMistNac struct {
	value *NetworkTemplateMistNac
	isSet bool
}

func (v NullableNetworkTemplateMistNac) Get() *NetworkTemplateMistNac {
	return v.value
}

func (v *NullableNetworkTemplateMistNac) Set(val *NetworkTemplateMistNac) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkTemplateMistNac) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkTemplateMistNac) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkTemplateMistNac(val *NetworkTemplateMistNac) *NullableNetworkTemplateMistNac {
	return &NullableNetworkTemplateMistNac{value: val, isSet: true}
}

func (v NullableNetworkTemplateMistNac) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkTemplateMistNac) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


