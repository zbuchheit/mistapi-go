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

// checks if the WlanCiscoCwa type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanCiscoCwa{}

// WlanCiscoCwa Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html
type WlanCiscoCwa struct {
	// list of hostnames without http(s):// (matched by substring)
	AllowedHostnames []string `json:"allowed_hostnames,omitempty"`
	// list of CIDRs
	AllowedSubnets []string `json:"allowed_subnets,omitempty"`
	// list of blocked CIDRs
	BlockedSubnets []string `json:"blocked_subnets,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanCiscoCwa WlanCiscoCwa

// NewWlanCiscoCwa instantiates a new WlanCiscoCwa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanCiscoCwa() *WlanCiscoCwa {
	this := WlanCiscoCwa{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewWlanCiscoCwaWithDefaults instantiates a new WlanCiscoCwa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanCiscoCwaWithDefaults() *WlanCiscoCwa {
	this := WlanCiscoCwa{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetAllowedHostnames returns the AllowedHostnames field value if set, zero value otherwise.
func (o *WlanCiscoCwa) GetAllowedHostnames() []string {
	if o == nil || IsNil(o.AllowedHostnames) {
		var ret []string
		return ret
	}
	return o.AllowedHostnames
}

// GetAllowedHostnamesOk returns a tuple with the AllowedHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanCiscoCwa) GetAllowedHostnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedHostnames) {
		return nil, false
	}
	return o.AllowedHostnames, true
}

// HasAllowedHostnames returns a boolean if a field has been set.
func (o *WlanCiscoCwa) HasAllowedHostnames() bool {
	if o != nil && !IsNil(o.AllowedHostnames) {
		return true
	}

	return false
}

// SetAllowedHostnames gets a reference to the given []string and assigns it to the AllowedHostnames field.
func (o *WlanCiscoCwa) SetAllowedHostnames(v []string) {
	o.AllowedHostnames = v
}

// GetAllowedSubnets returns the AllowedSubnets field value if set, zero value otherwise.
func (o *WlanCiscoCwa) GetAllowedSubnets() []string {
	if o == nil || IsNil(o.AllowedSubnets) {
		var ret []string
		return ret
	}
	return o.AllowedSubnets
}

// GetAllowedSubnetsOk returns a tuple with the AllowedSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanCiscoCwa) GetAllowedSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedSubnets) {
		return nil, false
	}
	return o.AllowedSubnets, true
}

// HasAllowedSubnets returns a boolean if a field has been set.
func (o *WlanCiscoCwa) HasAllowedSubnets() bool {
	if o != nil && !IsNil(o.AllowedSubnets) {
		return true
	}

	return false
}

// SetAllowedSubnets gets a reference to the given []string and assigns it to the AllowedSubnets field.
func (o *WlanCiscoCwa) SetAllowedSubnets(v []string) {
	o.AllowedSubnets = v
}

// GetBlockedSubnets returns the BlockedSubnets field value if set, zero value otherwise.
func (o *WlanCiscoCwa) GetBlockedSubnets() []string {
	if o == nil || IsNil(o.BlockedSubnets) {
		var ret []string
		return ret
	}
	return o.BlockedSubnets
}

// GetBlockedSubnetsOk returns a tuple with the BlockedSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanCiscoCwa) GetBlockedSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockedSubnets) {
		return nil, false
	}
	return o.BlockedSubnets, true
}

// HasBlockedSubnets returns a boolean if a field has been set.
func (o *WlanCiscoCwa) HasBlockedSubnets() bool {
	if o != nil && !IsNil(o.BlockedSubnets) {
		return true
	}

	return false
}

// SetBlockedSubnets gets a reference to the given []string and assigns it to the BlockedSubnets field.
func (o *WlanCiscoCwa) SetBlockedSubnets(v []string) {
	o.BlockedSubnets = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WlanCiscoCwa) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanCiscoCwa) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WlanCiscoCwa) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WlanCiscoCwa) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o WlanCiscoCwa) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanCiscoCwa) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedHostnames) {
		toSerialize["allowed_hostnames"] = o.AllowedHostnames
	}
	if !IsNil(o.AllowedSubnets) {
		toSerialize["allowed_subnets"] = o.AllowedSubnets
	}
	if !IsNil(o.BlockedSubnets) {
		toSerialize["blocked_subnets"] = o.BlockedSubnets
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanCiscoCwa) UnmarshalJSON(data []byte) (err error) {
	varWlanCiscoCwa := _WlanCiscoCwa{}

	err = json.Unmarshal(data, &varWlanCiscoCwa)

	if err != nil {
		return err
	}

	*o = WlanCiscoCwa(varWlanCiscoCwa)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowed_hostnames")
		delete(additionalProperties, "allowed_subnets")
		delete(additionalProperties, "blocked_subnets")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanCiscoCwa struct {
	value *WlanCiscoCwa
	isSet bool
}

func (v NullableWlanCiscoCwa) Get() *WlanCiscoCwa {
	return v.value
}

func (v *NullableWlanCiscoCwa) Set(val *WlanCiscoCwa) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanCiscoCwa) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanCiscoCwa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanCiscoCwa(val *WlanCiscoCwa) *NullableWlanCiscoCwa {
	return &NullableWlanCiscoCwa{value: val, isSet: true}
}

func (v NullableWlanCiscoCwa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanCiscoCwa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


