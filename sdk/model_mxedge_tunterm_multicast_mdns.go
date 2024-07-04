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

// checks if the MxedgeTuntermMulticastMdns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeTuntermMulticastMdns{}

// MxedgeTuntermMulticastMdns struct for MxedgeTuntermMulticastMdns
type MxedgeTuntermMulticastMdns struct {
	Enabled *bool `json:"enabled,omitempty"`
	VlanIds []string `json:"vlan_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeTuntermMulticastMdns MxedgeTuntermMulticastMdns

// NewMxedgeTuntermMulticastMdns instantiates a new MxedgeTuntermMulticastMdns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeTuntermMulticastMdns() *MxedgeTuntermMulticastMdns {
	this := MxedgeTuntermMulticastMdns{}
	return &this
}

// NewMxedgeTuntermMulticastMdnsWithDefaults instantiates a new MxedgeTuntermMulticastMdns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeTuntermMulticastMdnsWithDefaults() *MxedgeTuntermMulticastMdns {
	this := MxedgeTuntermMulticastMdns{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MxedgeTuntermMulticastMdns) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermMulticastMdns) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MxedgeTuntermMulticastMdns) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MxedgeTuntermMulticastMdns) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVlanIds returns the VlanIds field value if set, zero value otherwise.
func (o *MxedgeTuntermMulticastMdns) GetVlanIds() []string {
	if o == nil || IsNil(o.VlanIds) {
		var ret []string
		return ret
	}
	return o.VlanIds
}

// GetVlanIdsOk returns a tuple with the VlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermMulticastMdns) GetVlanIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VlanIds) {
		return nil, false
	}
	return o.VlanIds, true
}

// HasVlanIds returns a boolean if a field has been set.
func (o *MxedgeTuntermMulticastMdns) HasVlanIds() bool {
	if o != nil && !IsNil(o.VlanIds) {
		return true
	}

	return false
}

// SetVlanIds gets a reference to the given []string and assigns it to the VlanIds field.
func (o *MxedgeTuntermMulticastMdns) SetVlanIds(v []string) {
	o.VlanIds = v
}

func (o MxedgeTuntermMulticastMdns) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeTuntermMulticastMdns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.VlanIds) {
		toSerialize["vlan_ids"] = o.VlanIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeTuntermMulticastMdns) UnmarshalJSON(data []byte) (err error) {
	varMxedgeTuntermMulticastMdns := _MxedgeTuntermMulticastMdns{}

	err = json.Unmarshal(data, &varMxedgeTuntermMulticastMdns)

	if err != nil {
		return err
	}

	*o = MxedgeTuntermMulticastMdns(varMxedgeTuntermMulticastMdns)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "vlan_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeTuntermMulticastMdns struct {
	value *MxedgeTuntermMulticastMdns
	isSet bool
}

func (v NullableMxedgeTuntermMulticastMdns) Get() *MxedgeTuntermMulticastMdns {
	return v.value
}

func (v *NullableMxedgeTuntermMulticastMdns) Set(val *MxedgeTuntermMulticastMdns) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeTuntermMulticastMdns) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeTuntermMulticastMdns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeTuntermMulticastMdns(val *MxedgeTuntermMulticastMdns) *NullableMxedgeTuntermMulticastMdns {
	return &NullableMxedgeTuntermMulticastMdns{value: val, isSet: true}
}

func (v NullableMxedgeTuntermMulticastMdns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeTuntermMulticastMdns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


