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

// checks if the OrgSettingWiredPma type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingWiredPma{}

// OrgSettingWiredPma struct for OrgSettingWiredPma
type OrgSettingWiredPma struct {
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingWiredPma OrgSettingWiredPma

// NewOrgSettingWiredPma instantiates a new OrgSettingWiredPma object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingWiredPma() *OrgSettingWiredPma {
	this := OrgSettingWiredPma{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewOrgSettingWiredPmaWithDefaults instantiates a new OrgSettingWiredPma object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingWiredPmaWithDefaults() *OrgSettingWiredPma {
	this := OrgSettingWiredPma{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OrgSettingWiredPma) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingWiredPma) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OrgSettingWiredPma) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OrgSettingWiredPma) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o OrgSettingWiredPma) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingWiredPma) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingWiredPma) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingWiredPma := _OrgSettingWiredPma{}

	err = json.Unmarshal(data, &varOrgSettingWiredPma)

	if err != nil {
		return err
	}

	*o = OrgSettingWiredPma(varOrgSettingWiredPma)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingWiredPma struct {
	value *OrgSettingWiredPma
	isSet bool
}

func (v NullableOrgSettingWiredPma) Get() *OrgSettingWiredPma {
	return v.value
}

func (v *NullableOrgSettingWiredPma) Set(val *OrgSettingWiredPma) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingWiredPma) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingWiredPma) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingWiredPma(val *OrgSettingWiredPma) *NullableOrgSettingWiredPma {
	return &NullableOrgSettingWiredPma{value: val, isSet: true}
}

func (v NullableOrgSettingWiredPma) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingWiredPma) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


