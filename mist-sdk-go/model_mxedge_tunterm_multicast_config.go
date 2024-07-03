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

// checks if the MxedgeTuntermMulticastConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeTuntermMulticastConfig{}

// MxedgeTuntermMulticastConfig struct for MxedgeTuntermMulticastConfig
type MxedgeTuntermMulticastConfig struct {
	Mdns *MxedgeTuntermMulticastMdns `json:"mdns,omitempty"`
	Ssdp *MxedgeTuntermMulticastSsdp `json:"ssdp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeTuntermMulticastConfig MxedgeTuntermMulticastConfig

// NewMxedgeTuntermMulticastConfig instantiates a new MxedgeTuntermMulticastConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeTuntermMulticastConfig() *MxedgeTuntermMulticastConfig {
	this := MxedgeTuntermMulticastConfig{}
	return &this
}

// NewMxedgeTuntermMulticastConfigWithDefaults instantiates a new MxedgeTuntermMulticastConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeTuntermMulticastConfigWithDefaults() *MxedgeTuntermMulticastConfig {
	this := MxedgeTuntermMulticastConfig{}
	return &this
}

// GetMdns returns the Mdns field value if set, zero value otherwise.
func (o *MxedgeTuntermMulticastConfig) GetMdns() MxedgeTuntermMulticastMdns {
	if o == nil || IsNil(o.Mdns) {
		var ret MxedgeTuntermMulticastMdns
		return ret
	}
	return *o.Mdns
}

// GetMdnsOk returns a tuple with the Mdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermMulticastConfig) GetMdnsOk() (*MxedgeTuntermMulticastMdns, bool) {
	if o == nil || IsNil(o.Mdns) {
		return nil, false
	}
	return o.Mdns, true
}

// HasMdns returns a boolean if a field has been set.
func (o *MxedgeTuntermMulticastConfig) HasMdns() bool {
	if o != nil && !IsNil(o.Mdns) {
		return true
	}

	return false
}

// SetMdns gets a reference to the given MxedgeTuntermMulticastMdns and assigns it to the Mdns field.
func (o *MxedgeTuntermMulticastConfig) SetMdns(v MxedgeTuntermMulticastMdns) {
	o.Mdns = &v
}

// GetSsdp returns the Ssdp field value if set, zero value otherwise.
func (o *MxedgeTuntermMulticastConfig) GetSsdp() MxedgeTuntermMulticastSsdp {
	if o == nil || IsNil(o.Ssdp) {
		var ret MxedgeTuntermMulticastSsdp
		return ret
	}
	return *o.Ssdp
}

// GetSsdpOk returns a tuple with the Ssdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermMulticastConfig) GetSsdpOk() (*MxedgeTuntermMulticastSsdp, bool) {
	if o == nil || IsNil(o.Ssdp) {
		return nil, false
	}
	return o.Ssdp, true
}

// HasSsdp returns a boolean if a field has been set.
func (o *MxedgeTuntermMulticastConfig) HasSsdp() bool {
	if o != nil && !IsNil(o.Ssdp) {
		return true
	}

	return false
}

// SetSsdp gets a reference to the given MxedgeTuntermMulticastSsdp and assigns it to the Ssdp field.
func (o *MxedgeTuntermMulticastConfig) SetSsdp(v MxedgeTuntermMulticastSsdp) {
	o.Ssdp = &v
}

func (o MxedgeTuntermMulticastConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeTuntermMulticastConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mdns) {
		toSerialize["mdns"] = o.Mdns
	}
	if !IsNil(o.Ssdp) {
		toSerialize["ssdp"] = o.Ssdp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeTuntermMulticastConfig) UnmarshalJSON(data []byte) (err error) {
	varMxedgeTuntermMulticastConfig := _MxedgeTuntermMulticastConfig{}

	err = json.Unmarshal(data, &varMxedgeTuntermMulticastConfig)

	if err != nil {
		return err
	}

	*o = MxedgeTuntermMulticastConfig(varMxedgeTuntermMulticastConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mdns")
		delete(additionalProperties, "ssdp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeTuntermMulticastConfig struct {
	value *MxedgeTuntermMulticastConfig
	isSet bool
}

func (v NullableMxedgeTuntermMulticastConfig) Get() *MxedgeTuntermMulticastConfig {
	return v.value
}

func (v *NullableMxedgeTuntermMulticastConfig) Set(val *MxedgeTuntermMulticastConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeTuntermMulticastConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeTuntermMulticastConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeTuntermMulticastConfig(val *MxedgeTuntermMulticastConfig) *NullableMxedgeTuntermMulticastConfig {
	return &NullableMxedgeTuntermMulticastConfig{value: val, isSet: true}
}

func (v NullableMxedgeTuntermMulticastConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeTuntermMulticastConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


