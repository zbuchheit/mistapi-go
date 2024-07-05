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

// checks if the TuntermDhcpdConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TuntermDhcpdConfig{}

// TuntermDhcpdConfig DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID
type TuntermDhcpdConfig struct {
	Enabled *bool `json:"enabled,omitempty"`
	Servers []string `json:"servers,omitempty"`
	Type *TuntermDhcpdType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TuntermDhcpdConfig TuntermDhcpdConfig

// NewTuntermDhcpdConfig instantiates a new TuntermDhcpdConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTuntermDhcpdConfig() *TuntermDhcpdConfig {
	this := TuntermDhcpdConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	var type_ TuntermDhcpdType = TUNTERMDHCPDTYPE_RELAY
	this.Type = &type_
	return &this
}

// NewTuntermDhcpdConfigWithDefaults instantiates a new TuntermDhcpdConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTuntermDhcpdConfigWithDefaults() *TuntermDhcpdConfig {
	this := TuntermDhcpdConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	var type_ TuntermDhcpdType = TUNTERMDHCPDTYPE_RELAY
	this.Type = &type_
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TuntermDhcpdConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuntermDhcpdConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TuntermDhcpdConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TuntermDhcpdConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *TuntermDhcpdConfig) GetServers() []string {
	if o == nil || IsNil(o.Servers) {
		var ret []string
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuntermDhcpdConfig) GetServersOk() ([]string, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *TuntermDhcpdConfig) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []string and assigns it to the Servers field.
func (o *TuntermDhcpdConfig) SetServers(v []string) {
	o.Servers = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TuntermDhcpdConfig) GetType() TuntermDhcpdType {
	if o == nil || IsNil(o.Type) {
		var ret TuntermDhcpdType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuntermDhcpdConfig) GetTypeOk() (*TuntermDhcpdType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TuntermDhcpdConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TuntermDhcpdType and assigns it to the Type field.
func (o *TuntermDhcpdConfig) SetType(v TuntermDhcpdType) {
	o.Type = &v
}

func (o TuntermDhcpdConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TuntermDhcpdConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TuntermDhcpdConfig) UnmarshalJSON(data []byte) (err error) {
	varTuntermDhcpdConfig := _TuntermDhcpdConfig{}

	err = json.Unmarshal(data, &varTuntermDhcpdConfig)

	if err != nil {
		return err
	}

	*o = TuntermDhcpdConfig(varTuntermDhcpdConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "servers")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTuntermDhcpdConfig struct {
	value *TuntermDhcpdConfig
	isSet bool
}

func (v NullableTuntermDhcpdConfig) Get() *TuntermDhcpdConfig {
	return v.value
}

func (v *NullableTuntermDhcpdConfig) Set(val *TuntermDhcpdConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTuntermDhcpdConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTuntermDhcpdConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTuntermDhcpdConfig(val *TuntermDhcpdConfig) *NullableTuntermDhcpdConfig {
	return &NullableTuntermDhcpdConfig{value: val, isSet: true}
}

func (v NullableTuntermDhcpdConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTuntermDhcpdConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


