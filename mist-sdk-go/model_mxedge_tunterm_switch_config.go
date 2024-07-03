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

// checks if the MxedgeTuntermSwitchConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeTuntermSwitchConfig{}

// MxedgeTuntermSwitchConfig struct for MxedgeTuntermSwitchConfig
type MxedgeTuntermSwitchConfig struct {
	PortVlanId *int32 `json:"port_vlan_id,omitempty"`
	VlanIds []MxedgeTuntermSwitchConfigVlanId `json:"vlan_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeTuntermSwitchConfig MxedgeTuntermSwitchConfig

// NewMxedgeTuntermSwitchConfig instantiates a new MxedgeTuntermSwitchConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeTuntermSwitchConfig() *MxedgeTuntermSwitchConfig {
	this := MxedgeTuntermSwitchConfig{}
	return &this
}

// NewMxedgeTuntermSwitchConfigWithDefaults instantiates a new MxedgeTuntermSwitchConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeTuntermSwitchConfigWithDefaults() *MxedgeTuntermSwitchConfig {
	this := MxedgeTuntermSwitchConfig{}
	return &this
}

// GetPortVlanId returns the PortVlanId field value if set, zero value otherwise.
func (o *MxedgeTuntermSwitchConfig) GetPortVlanId() int32 {
	if o == nil || IsNil(o.PortVlanId) {
		var ret int32
		return ret
	}
	return *o.PortVlanId
}

// GetPortVlanIdOk returns a tuple with the PortVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermSwitchConfig) GetPortVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PortVlanId) {
		return nil, false
	}
	return o.PortVlanId, true
}

// HasPortVlanId returns a boolean if a field has been set.
func (o *MxedgeTuntermSwitchConfig) HasPortVlanId() bool {
	if o != nil && !IsNil(o.PortVlanId) {
		return true
	}

	return false
}

// SetPortVlanId gets a reference to the given int32 and assigns it to the PortVlanId field.
func (o *MxedgeTuntermSwitchConfig) SetPortVlanId(v int32) {
	o.PortVlanId = &v
}

// GetVlanIds returns the VlanIds field value if set, zero value otherwise.
func (o *MxedgeTuntermSwitchConfig) GetVlanIds() []MxedgeTuntermSwitchConfigVlanId {
	if o == nil || IsNil(o.VlanIds) {
		var ret []MxedgeTuntermSwitchConfigVlanId
		return ret
	}
	return o.VlanIds
}

// GetVlanIdsOk returns a tuple with the VlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermSwitchConfig) GetVlanIdsOk() ([]MxedgeTuntermSwitchConfigVlanId, bool) {
	if o == nil || IsNil(o.VlanIds) {
		return nil, false
	}
	return o.VlanIds, true
}

// HasVlanIds returns a boolean if a field has been set.
func (o *MxedgeTuntermSwitchConfig) HasVlanIds() bool {
	if o != nil && !IsNil(o.VlanIds) {
		return true
	}

	return false
}

// SetVlanIds gets a reference to the given []MxedgeTuntermSwitchConfigVlanId and assigns it to the VlanIds field.
func (o *MxedgeTuntermSwitchConfig) SetVlanIds(v []MxedgeTuntermSwitchConfigVlanId) {
	o.VlanIds = v
}

func (o MxedgeTuntermSwitchConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeTuntermSwitchConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortVlanId) {
		toSerialize["port_vlan_id"] = o.PortVlanId
	}
	if !IsNil(o.VlanIds) {
		toSerialize["vlan_ids"] = o.VlanIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeTuntermSwitchConfig) UnmarshalJSON(data []byte) (err error) {
	varMxedgeTuntermSwitchConfig := _MxedgeTuntermSwitchConfig{}

	err = json.Unmarshal(data, &varMxedgeTuntermSwitchConfig)

	if err != nil {
		return err
	}

	*o = MxedgeTuntermSwitchConfig(varMxedgeTuntermSwitchConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "port_vlan_id")
		delete(additionalProperties, "vlan_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeTuntermSwitchConfig struct {
	value *MxedgeTuntermSwitchConfig
	isSet bool
}

func (v NullableMxedgeTuntermSwitchConfig) Get() *MxedgeTuntermSwitchConfig {
	return v.value
}

func (v *NullableMxedgeTuntermSwitchConfig) Set(val *MxedgeTuntermSwitchConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeTuntermSwitchConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeTuntermSwitchConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeTuntermSwitchConfig(val *MxedgeTuntermSwitchConfig) *NullableMxedgeTuntermSwitchConfig {
	return &NullableMxedgeTuntermSwitchConfig{value: val, isSet: true}
}

func (v NullableMxedgeTuntermSwitchConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeTuntermSwitchConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


