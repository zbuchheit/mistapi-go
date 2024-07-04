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

// checks if the SyntheticTestConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticTestConfig{}

// SyntheticTestConfig struct for SyntheticTestConfig
type SyntheticTestConfig struct {
	Disabled *bool `json:"disabled,omitempty"`
	Vlans []SyntheticTestProperties `json:"vlans,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyntheticTestConfig SyntheticTestConfig

// NewSyntheticTestConfig instantiates a new SyntheticTestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticTestConfig() *SyntheticTestConfig {
	this := SyntheticTestConfig{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// NewSyntheticTestConfigWithDefaults instantiates a new SyntheticTestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticTestConfigWithDefaults() *SyntheticTestConfig {
	this := SyntheticTestConfig{}
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *SyntheticTestConfig) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestConfig) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *SyntheticTestConfig) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *SyntheticTestConfig) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *SyntheticTestConfig) GetVlans() []SyntheticTestProperties {
	if o == nil || IsNil(o.Vlans) {
		var ret []SyntheticTestProperties
		return ret
	}
	return o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestConfig) GetVlansOk() ([]SyntheticTestProperties, bool) {
	if o == nil || IsNil(o.Vlans) {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *SyntheticTestConfig) HasVlans() bool {
	if o != nil && !IsNil(o.Vlans) {
		return true
	}

	return false
}

// SetVlans gets a reference to the given []SyntheticTestProperties and assigns it to the Vlans field.
func (o *SyntheticTestConfig) SetVlans(v []SyntheticTestProperties) {
	o.Vlans = v
}

func (o SyntheticTestConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticTestConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Vlans) {
		toSerialize["vlans"] = o.Vlans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SyntheticTestConfig) UnmarshalJSON(data []byte) (err error) {
	varSyntheticTestConfig := _SyntheticTestConfig{}

	err = json.Unmarshal(data, &varSyntheticTestConfig)

	if err != nil {
		return err
	}

	*o = SyntheticTestConfig(varSyntheticTestConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "vlans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSyntheticTestConfig struct {
	value *SyntheticTestConfig
	isSet bool
}

func (v NullableSyntheticTestConfig) Get() *SyntheticTestConfig {
	return v.value
}

func (v *NullableSyntheticTestConfig) Set(val *SyntheticTestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTestConfig(val *SyntheticTestConfig) *NullableSyntheticTestConfig {
	return &NullableSyntheticTestConfig{value: val, isSet: true}
}

func (v NullableSyntheticTestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


