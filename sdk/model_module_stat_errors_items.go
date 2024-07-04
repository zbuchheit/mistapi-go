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
	"fmt"
)

// checks if the ModuleStatErrorsItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStatErrorsItems{}

// ModuleStatErrorsItems struct for ModuleStatErrorsItems
type ModuleStatErrorsItems struct {
	Feature *string `json:"feature,omitempty"`
	MinimumVersion *string `json:"minimum_version,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Since int32 `json:"since"`
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ModuleStatErrorsItems ModuleStatErrorsItems

// NewModuleStatErrorsItems instantiates a new ModuleStatErrorsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStatErrorsItems(since int32, type_ string) *ModuleStatErrorsItems {
	this := ModuleStatErrorsItems{}
	this.Since = since
	this.Type = type_
	return &this
}

// NewModuleStatErrorsItemsWithDefaults instantiates a new ModuleStatErrorsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatErrorsItemsWithDefaults() *ModuleStatErrorsItems {
	this := ModuleStatErrorsItems{}
	return &this
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *ModuleStatErrorsItems) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatErrorsItems) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *ModuleStatErrorsItems) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *ModuleStatErrorsItems) SetFeature(v string) {
	o.Feature = &v
}

// GetMinimumVersion returns the MinimumVersion field value if set, zero value otherwise.
func (o *ModuleStatErrorsItems) GetMinimumVersion() string {
	if o == nil || IsNil(o.MinimumVersion) {
		var ret string
		return ret
	}
	return *o.MinimumVersion
}

// GetMinimumVersionOk returns a tuple with the MinimumVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatErrorsItems) GetMinimumVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumVersion) {
		return nil, false
	}
	return o.MinimumVersion, true
}

// HasMinimumVersion returns a boolean if a field has been set.
func (o *ModuleStatErrorsItems) HasMinimumVersion() bool {
	if o != nil && !IsNil(o.MinimumVersion) {
		return true
	}

	return false
}

// SetMinimumVersion gets a reference to the given string and assigns it to the MinimumVersion field.
func (o *ModuleStatErrorsItems) SetMinimumVersion(v string) {
	o.MinimumVersion = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ModuleStatErrorsItems) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatErrorsItems) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ModuleStatErrorsItems) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ModuleStatErrorsItems) SetReason(v string) {
	o.Reason = &v
}

// GetSince returns the Since field value
func (o *ModuleStatErrorsItems) GetSince() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Since
}

// GetSinceOk returns a tuple with the Since field value
// and a boolean to check if the value has been set.
func (o *ModuleStatErrorsItems) GetSinceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Since, true
}

// SetSince sets field value
func (o *ModuleStatErrorsItems) SetSince(v int32) {
	o.Since = v
}

// GetType returns the Type field value
func (o *ModuleStatErrorsItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModuleStatErrorsItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ModuleStatErrorsItems) SetType(v string) {
	o.Type = v
}

func (o ModuleStatErrorsItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStatErrorsItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	if !IsNil(o.MinimumVersion) {
		toSerialize["minimum_version"] = o.MinimumVersion
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["since"] = o.Since
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleStatErrorsItems) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"since",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModuleStatErrorsItems := _ModuleStatErrorsItems{}

	err = json.Unmarshal(data, &varModuleStatErrorsItems)

	if err != nil {
		return err
	}

	*o = ModuleStatErrorsItems(varModuleStatErrorsItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feature")
		delete(additionalProperties, "minimum_version")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "since")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleStatErrorsItems struct {
	value *ModuleStatErrorsItems
	isSet bool
}

func (v NullableModuleStatErrorsItems) Get() *ModuleStatErrorsItems {
	return v.value
}

func (v *NullableModuleStatErrorsItems) Set(val *ModuleStatErrorsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatErrorsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatErrorsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatErrorsItems(val *ModuleStatErrorsItems) *NullableModuleStatErrorsItems {
	return &NullableModuleStatErrorsItems{value: val, isSet: true}
}

func (v NullableModuleStatErrorsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatErrorsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


