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

// checks if the ModuleStatVcLinksItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStatVcLinksItems{}

// ModuleStatVcLinksItems struct for ModuleStatVcLinksItems
type ModuleStatVcLinksItems struct {
	NeighborModuleIdx *int32 `json:"neighbor_module_idx,omitempty"`
	NeighborPortId *string `json:"neighbor_port_id,omitempty"`
	PortId *string `json:"port_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleStatVcLinksItems ModuleStatVcLinksItems

// NewModuleStatVcLinksItems instantiates a new ModuleStatVcLinksItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStatVcLinksItems() *ModuleStatVcLinksItems {
	this := ModuleStatVcLinksItems{}
	return &this
}

// NewModuleStatVcLinksItemsWithDefaults instantiates a new ModuleStatVcLinksItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatVcLinksItemsWithDefaults() *ModuleStatVcLinksItems {
	this := ModuleStatVcLinksItems{}
	return &this
}

// GetNeighborModuleIdx returns the NeighborModuleIdx field value if set, zero value otherwise.
func (o *ModuleStatVcLinksItems) GetNeighborModuleIdx() int32 {
	if o == nil || IsNil(o.NeighborModuleIdx) {
		var ret int32
		return ret
	}
	return *o.NeighborModuleIdx
}

// GetNeighborModuleIdxOk returns a tuple with the NeighborModuleIdx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatVcLinksItems) GetNeighborModuleIdxOk() (*int32, bool) {
	if o == nil || IsNil(o.NeighborModuleIdx) {
		return nil, false
	}
	return o.NeighborModuleIdx, true
}

// HasNeighborModuleIdx returns a boolean if a field has been set.
func (o *ModuleStatVcLinksItems) HasNeighborModuleIdx() bool {
	if o != nil && !IsNil(o.NeighborModuleIdx) {
		return true
	}

	return false
}

// SetNeighborModuleIdx gets a reference to the given int32 and assigns it to the NeighborModuleIdx field.
func (o *ModuleStatVcLinksItems) SetNeighborModuleIdx(v int32) {
	o.NeighborModuleIdx = &v
}

// GetNeighborPortId returns the NeighborPortId field value if set, zero value otherwise.
func (o *ModuleStatVcLinksItems) GetNeighborPortId() string {
	if o == nil || IsNil(o.NeighborPortId) {
		var ret string
		return ret
	}
	return *o.NeighborPortId
}

// GetNeighborPortIdOk returns a tuple with the NeighborPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatVcLinksItems) GetNeighborPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.NeighborPortId) {
		return nil, false
	}
	return o.NeighborPortId, true
}

// HasNeighborPortId returns a boolean if a field has been set.
func (o *ModuleStatVcLinksItems) HasNeighborPortId() bool {
	if o != nil && !IsNil(o.NeighborPortId) {
		return true
	}

	return false
}

// SetNeighborPortId gets a reference to the given string and assigns it to the NeighborPortId field.
func (o *ModuleStatVcLinksItems) SetNeighborPortId(v string) {
	o.NeighborPortId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *ModuleStatVcLinksItems) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatVcLinksItems) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *ModuleStatVcLinksItems) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *ModuleStatVcLinksItems) SetPortId(v string) {
	o.PortId = &v
}

func (o ModuleStatVcLinksItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStatVcLinksItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NeighborModuleIdx) {
		toSerialize["neighbor_module_idx"] = o.NeighborModuleIdx
	}
	if !IsNil(o.NeighborPortId) {
		toSerialize["neighbor_port_id"] = o.NeighborPortId
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleStatVcLinksItems) UnmarshalJSON(data []byte) (err error) {
	varModuleStatVcLinksItems := _ModuleStatVcLinksItems{}

	err = json.Unmarshal(data, &varModuleStatVcLinksItems)

	if err != nil {
		return err
	}

	*o = ModuleStatVcLinksItems(varModuleStatVcLinksItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "neighbor_module_idx")
		delete(additionalProperties, "neighbor_port_id")
		delete(additionalProperties, "port_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleStatVcLinksItems struct {
	value *ModuleStatVcLinksItems
	isSet bool
}

func (v NullableModuleStatVcLinksItems) Get() *ModuleStatVcLinksItems {
	return v.value
}

func (v *NullableModuleStatVcLinksItems) Set(val *ModuleStatVcLinksItems) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatVcLinksItems) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatVcLinksItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatVcLinksItems(val *ModuleStatVcLinksItems) *NullableModuleStatVcLinksItems {
	return &NullableModuleStatVcLinksItems{value: val, isSet: true}
}

func (v NullableModuleStatVcLinksItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatVcLinksItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


