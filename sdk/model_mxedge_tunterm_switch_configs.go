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

// checks if the MxedgeTuntermSwitchConfigs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeTuntermSwitchConfigs{}

// MxedgeTuntermSwitchConfigs if custom vlan settings are desired
type MxedgeTuntermSwitchConfigs struct {
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeTuntermSwitchConfigs MxedgeTuntermSwitchConfigs

// NewMxedgeTuntermSwitchConfigs instantiates a new MxedgeTuntermSwitchConfigs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeTuntermSwitchConfigs() *MxedgeTuntermSwitchConfigs {
	this := MxedgeTuntermSwitchConfigs{}
	return &this
}

// NewMxedgeTuntermSwitchConfigsWithDefaults instantiates a new MxedgeTuntermSwitchConfigs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeTuntermSwitchConfigsWithDefaults() *MxedgeTuntermSwitchConfigs {
	this := MxedgeTuntermSwitchConfigs{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MxedgeTuntermSwitchConfigs) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeTuntermSwitchConfigs) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MxedgeTuntermSwitchConfigs) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MxedgeTuntermSwitchConfigs) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o MxedgeTuntermSwitchConfigs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeTuntermSwitchConfigs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeTuntermSwitchConfigs) UnmarshalJSON(data []byte) (err error) {
	varMxedgeTuntermSwitchConfigs := _MxedgeTuntermSwitchConfigs{}

	err = json.Unmarshal(data, &varMxedgeTuntermSwitchConfigs)

	if err != nil {
		return err
	}

	*o = MxedgeTuntermSwitchConfigs(varMxedgeTuntermSwitchConfigs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeTuntermSwitchConfigs struct {
	value *MxedgeTuntermSwitchConfigs
	isSet bool
}

func (v NullableMxedgeTuntermSwitchConfigs) Get() *MxedgeTuntermSwitchConfigs {
	return v.value
}

func (v *NullableMxedgeTuntermSwitchConfigs) Set(val *MxedgeTuntermSwitchConfigs) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeTuntermSwitchConfigs) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeTuntermSwitchConfigs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeTuntermSwitchConfigs(val *MxedgeTuntermSwitchConfigs) *NullableMxedgeTuntermSwitchConfigs {
	return &NullableMxedgeTuntermSwitchConfigs{value: val, isSet: true}
}

func (v NullableMxedgeTuntermSwitchConfigs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeTuntermSwitchConfigs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


