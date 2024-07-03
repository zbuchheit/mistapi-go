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

// checks if the ApTemplateMatching type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApTemplateMatching{}

// ApTemplateMatching struct for ApTemplateMatching
type ApTemplateMatching struct {
	Enabled *bool `json:"enabled,omitempty"`
	Rules []ApTemplateMatchingRule `json:"rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApTemplateMatching ApTemplateMatching

// NewApTemplateMatching instantiates a new ApTemplateMatching object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApTemplateMatching() *ApTemplateMatching {
	this := ApTemplateMatching{}
	return &this
}

// NewApTemplateMatchingWithDefaults instantiates a new ApTemplateMatching object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApTemplateMatchingWithDefaults() *ApTemplateMatching {
	this := ApTemplateMatching{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApTemplateMatching) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplateMatching) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApTemplateMatching) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApTemplateMatching) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ApTemplateMatching) GetRules() []ApTemplateMatchingRule {
	if o == nil || IsNil(o.Rules) {
		var ret []ApTemplateMatchingRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApTemplateMatching) GetRulesOk() ([]ApTemplateMatchingRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ApTemplateMatching) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ApTemplateMatchingRule and assigns it to the Rules field.
func (o *ApTemplateMatching) SetRules(v []ApTemplateMatchingRule) {
	o.Rules = v
}

func (o ApTemplateMatching) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApTemplateMatching) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApTemplateMatching) UnmarshalJSON(data []byte) (err error) {
	varApTemplateMatching := _ApTemplateMatching{}

	err = json.Unmarshal(data, &varApTemplateMatching)

	if err != nil {
		return err
	}

	*o = ApTemplateMatching(varApTemplateMatching)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApTemplateMatching struct {
	value *ApTemplateMatching
	isSet bool
}

func (v NullableApTemplateMatching) Get() *ApTemplateMatching {
	return v.value
}

func (v *NullableApTemplateMatching) Set(val *ApTemplateMatching) {
	v.value = val
	v.isSet = true
}

func (v NullableApTemplateMatching) IsSet() bool {
	return v.isSet
}

func (v *NullableApTemplateMatching) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApTemplateMatching(val *ApTemplateMatching) *NullableApTemplateMatching {
	return &NullableApTemplateMatching{value: val, isSet: true}
}

func (v NullableApTemplateMatching) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApTemplateMatching) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


