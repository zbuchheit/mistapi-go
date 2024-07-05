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

// checks if the TunnelConfigsAutoProvisionNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelConfigsAutoProvisionNode{}

// TunnelConfigsAutoProvisionNode struct for TunnelConfigsAutoProvisionNode
type TunnelConfigsAutoProvisionNode struct {
	NumHosts *string `json:"num_hosts,omitempty"`
	// optional, only needed if `vars_only`==`false`
	WanNames []string `json:"wan_names,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelConfigsAutoProvisionNode TunnelConfigsAutoProvisionNode

// NewTunnelConfigsAutoProvisionNode instantiates a new TunnelConfigsAutoProvisionNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelConfigsAutoProvisionNode() *TunnelConfigsAutoProvisionNode {
	this := TunnelConfigsAutoProvisionNode{}
	return &this
}

// NewTunnelConfigsAutoProvisionNodeWithDefaults instantiates a new TunnelConfigsAutoProvisionNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelConfigsAutoProvisionNodeWithDefaults() *TunnelConfigsAutoProvisionNode {
	this := TunnelConfigsAutoProvisionNode{}
	return &this
}

// GetNumHosts returns the NumHosts field value if set, zero value otherwise.
func (o *TunnelConfigsAutoProvisionNode) GetNumHosts() string {
	if o == nil || IsNil(o.NumHosts) {
		var ret string
		return ret
	}
	return *o.NumHosts
}

// GetNumHostsOk returns a tuple with the NumHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelConfigsAutoProvisionNode) GetNumHostsOk() (*string, bool) {
	if o == nil || IsNil(o.NumHosts) {
		return nil, false
	}
	return o.NumHosts, true
}

// HasNumHosts returns a boolean if a field has been set.
func (o *TunnelConfigsAutoProvisionNode) HasNumHosts() bool {
	if o != nil && !IsNil(o.NumHosts) {
		return true
	}

	return false
}

// SetNumHosts gets a reference to the given string and assigns it to the NumHosts field.
func (o *TunnelConfigsAutoProvisionNode) SetNumHosts(v string) {
	o.NumHosts = &v
}

// GetWanNames returns the WanNames field value if set, zero value otherwise.
func (o *TunnelConfigsAutoProvisionNode) GetWanNames() []string {
	if o == nil || IsNil(o.WanNames) {
		var ret []string
		return ret
	}
	return o.WanNames
}

// GetWanNamesOk returns a tuple with the WanNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelConfigsAutoProvisionNode) GetWanNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.WanNames) {
		return nil, false
	}
	return o.WanNames, true
}

// HasWanNames returns a boolean if a field has been set.
func (o *TunnelConfigsAutoProvisionNode) HasWanNames() bool {
	if o != nil && !IsNil(o.WanNames) {
		return true
	}

	return false
}

// SetWanNames gets a reference to the given []string and assigns it to the WanNames field.
func (o *TunnelConfigsAutoProvisionNode) SetWanNames(v []string) {
	o.WanNames = v
}

func (o TunnelConfigsAutoProvisionNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelConfigsAutoProvisionNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumHosts) {
		toSerialize["num_hosts"] = o.NumHosts
	}
	if !IsNil(o.WanNames) {
		toSerialize["wan_names"] = o.WanNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelConfigsAutoProvisionNode) UnmarshalJSON(data []byte) (err error) {
	varTunnelConfigsAutoProvisionNode := _TunnelConfigsAutoProvisionNode{}

	err = json.Unmarshal(data, &varTunnelConfigsAutoProvisionNode)

	if err != nil {
		return err
	}

	*o = TunnelConfigsAutoProvisionNode(varTunnelConfigsAutoProvisionNode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "num_hosts")
		delete(additionalProperties, "wan_names")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelConfigsAutoProvisionNode struct {
	value *TunnelConfigsAutoProvisionNode
	isSet bool
}

func (v NullableTunnelConfigsAutoProvisionNode) Get() *TunnelConfigsAutoProvisionNode {
	return v.value
}

func (v *NullableTunnelConfigsAutoProvisionNode) Set(val *TunnelConfigsAutoProvisionNode) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelConfigsAutoProvisionNode) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelConfigsAutoProvisionNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelConfigsAutoProvisionNode(val *TunnelConfigsAutoProvisionNode) *NullableTunnelConfigsAutoProvisionNode {
	return &NullableTunnelConfigsAutoProvisionNode{value: val, isSet: true}
}

func (v NullableTunnelConfigsAutoProvisionNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelConfigsAutoProvisionNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


