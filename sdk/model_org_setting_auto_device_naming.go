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

// checks if the OrgSettingAutoDeviceNaming type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingAutoDeviceNaming{}

// OrgSettingAutoDeviceNaming struct for OrgSettingAutoDeviceNaming
type OrgSettingAutoDeviceNaming struct {
	Enable *bool `json:"enable,omitempty"`
	Rules []OrgAutoRules `json:"rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingAutoDeviceNaming OrgSettingAutoDeviceNaming

// NewOrgSettingAutoDeviceNaming instantiates a new OrgSettingAutoDeviceNaming object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingAutoDeviceNaming() *OrgSettingAutoDeviceNaming {
	this := OrgSettingAutoDeviceNaming{}
	return &this
}

// NewOrgSettingAutoDeviceNamingWithDefaults instantiates a new OrgSettingAutoDeviceNaming object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingAutoDeviceNamingWithDefaults() *OrgSettingAutoDeviceNaming {
	this := OrgSettingAutoDeviceNaming{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *OrgSettingAutoDeviceNaming) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingAutoDeviceNaming) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *OrgSettingAutoDeviceNaming) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *OrgSettingAutoDeviceNaming) SetEnable(v bool) {
	o.Enable = &v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgSettingAutoDeviceNaming) GetRules() []OrgAutoRules {
	if o == nil {
		var ret []OrgAutoRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgSettingAutoDeviceNaming) GetRulesOk() ([]OrgAutoRules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *OrgSettingAutoDeviceNaming) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []OrgAutoRules and assigns it to the Rules field.
func (o *OrgSettingAutoDeviceNaming) SetRules(v []OrgAutoRules) {
	o.Rules = v
}

func (o OrgSettingAutoDeviceNaming) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingAutoDeviceNaming) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingAutoDeviceNaming) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingAutoDeviceNaming := _OrgSettingAutoDeviceNaming{}

	err = json.Unmarshal(data, &varOrgSettingAutoDeviceNaming)

	if err != nil {
		return err
	}

	*o = OrgSettingAutoDeviceNaming(varOrgSettingAutoDeviceNaming)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingAutoDeviceNaming struct {
	value *OrgSettingAutoDeviceNaming
	isSet bool
}

func (v NullableOrgSettingAutoDeviceNaming) Get() *OrgSettingAutoDeviceNaming {
	return v.value
}

func (v *NullableOrgSettingAutoDeviceNaming) Set(val *OrgSettingAutoDeviceNaming) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingAutoDeviceNaming) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingAutoDeviceNaming) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingAutoDeviceNaming(val *OrgSettingAutoDeviceNaming) *NullableOrgSettingAutoDeviceNaming {
	return &NullableOrgSettingAutoDeviceNaming{value: val, isSet: true}
}

func (v NullableOrgSettingAutoDeviceNaming) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingAutoDeviceNaming) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


