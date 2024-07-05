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

// checks if the OrgSettingWirelessPma type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingWirelessPma{}

// OrgSettingWirelessPma struct for OrgSettingWirelessPma
type OrgSettingWirelessPma struct {
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingWirelessPma OrgSettingWirelessPma

// NewOrgSettingWirelessPma instantiates a new OrgSettingWirelessPma object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingWirelessPma() *OrgSettingWirelessPma {
	this := OrgSettingWirelessPma{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewOrgSettingWirelessPmaWithDefaults instantiates a new OrgSettingWirelessPma object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingWirelessPmaWithDefaults() *OrgSettingWirelessPma {
	this := OrgSettingWirelessPma{}
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrgSettingWirelessPma) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingWirelessPma) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrgSettingWirelessPma) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrgSettingWirelessPma) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OrgSettingWirelessPma) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingWirelessPma) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingWirelessPma) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingWirelessPma := _OrgSettingWirelessPma{}

	err = json.Unmarshal(data, &varOrgSettingWirelessPma)

	if err != nil {
		return err
	}

	*o = OrgSettingWirelessPma(varOrgSettingWirelessPma)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingWirelessPma struct {
	value *OrgSettingWirelessPma
	isSet bool
}

func (v NullableOrgSettingWirelessPma) Get() *OrgSettingWirelessPma {
	return v.value
}

func (v *NullableOrgSettingWirelessPma) Set(val *OrgSettingWirelessPma) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingWirelessPma) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingWirelessPma) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingWirelessPma(val *OrgSettingWirelessPma) *NullableOrgSettingWirelessPma {
	return &NullableOrgSettingWirelessPma{value: val, isSet: true}
}

func (v NullableOrgSettingWirelessPma) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingWirelessPma) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


