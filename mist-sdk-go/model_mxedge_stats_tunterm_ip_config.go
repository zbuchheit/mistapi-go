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

// checks if the MxedgeStatsTuntermIpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeStatsTuntermIpConfig{}

// MxedgeStatsTuntermIpConfig struct for MxedgeStatsTuntermIpConfig
type MxedgeStatsTuntermIpConfig struct {
	Gateway *string `json:"gateway,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeStatsTuntermIpConfig MxedgeStatsTuntermIpConfig

// NewMxedgeStatsTuntermIpConfig instantiates a new MxedgeStatsTuntermIpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeStatsTuntermIpConfig() *MxedgeStatsTuntermIpConfig {
	this := MxedgeStatsTuntermIpConfig{}
	return &this
}

// NewMxedgeStatsTuntermIpConfigWithDefaults instantiates a new MxedgeStatsTuntermIpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeStatsTuntermIpConfigWithDefaults() *MxedgeStatsTuntermIpConfig {
	this := MxedgeStatsTuntermIpConfig{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MxedgeStatsTuntermIpConfig) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsTuntermIpConfig) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MxedgeStatsTuntermIpConfig) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *MxedgeStatsTuntermIpConfig) SetGateway(v string) {
	o.Gateway = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *MxedgeStatsTuntermIpConfig) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsTuntermIpConfig) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *MxedgeStatsTuntermIpConfig) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *MxedgeStatsTuntermIpConfig) SetIp(v string) {
	o.Ip = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *MxedgeStatsTuntermIpConfig) GetNetmask() string {
	if o == nil || IsNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsTuntermIpConfig) GetNetmaskOk() (*string, bool) {
	if o == nil || IsNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *MxedgeStatsTuntermIpConfig) HasNetmask() bool {
	if o != nil && !IsNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *MxedgeStatsTuntermIpConfig) SetNetmask(v string) {
	o.Netmask = &v
}

func (o MxedgeStatsTuntermIpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeStatsTuntermIpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Netmask) {
		toSerialize["netmask"] = o.Netmask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeStatsTuntermIpConfig) UnmarshalJSON(data []byte) (err error) {
	varMxedgeStatsTuntermIpConfig := _MxedgeStatsTuntermIpConfig{}

	err = json.Unmarshal(data, &varMxedgeStatsTuntermIpConfig)

	if err != nil {
		return err
	}

	*o = MxedgeStatsTuntermIpConfig(varMxedgeStatsTuntermIpConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "netmask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeStatsTuntermIpConfig struct {
	value *MxedgeStatsTuntermIpConfig
	isSet bool
}

func (v NullableMxedgeStatsTuntermIpConfig) Get() *MxedgeStatsTuntermIpConfig {
	return v.value
}

func (v *NullableMxedgeStatsTuntermIpConfig) Set(val *MxedgeStatsTuntermIpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeStatsTuntermIpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeStatsTuntermIpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeStatsTuntermIpConfig(val *MxedgeStatsTuntermIpConfig) *NullableMxedgeStatsTuntermIpConfig {
	return &NullableMxedgeStatsTuntermIpConfig{value: val, isSet: true}
}

func (v NullableMxedgeStatsTuntermIpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeStatsTuntermIpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


