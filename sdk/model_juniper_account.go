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

// checks if the JuniperAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JuniperAccount{}

// JuniperAccount struct for JuniperAccount
type JuniperAccount struct {
	LinkedBy *string `json:"linked_by,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JuniperAccount JuniperAccount

// NewJuniperAccount instantiates a new JuniperAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJuniperAccount() *JuniperAccount {
	this := JuniperAccount{}
	return &this
}

// NewJuniperAccountWithDefaults instantiates a new JuniperAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJuniperAccountWithDefaults() *JuniperAccount {
	this := JuniperAccount{}
	return &this
}

// GetLinkedBy returns the LinkedBy field value if set, zero value otherwise.
func (o *JuniperAccount) GetLinkedBy() string {
	if o == nil || IsNil(o.LinkedBy) {
		var ret string
		return ret
	}
	return *o.LinkedBy
}

// GetLinkedByOk returns a tuple with the LinkedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JuniperAccount) GetLinkedByOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedBy) {
		return nil, false
	}
	return o.LinkedBy, true
}

// HasLinkedBy returns a boolean if a field has been set.
func (o *JuniperAccount) HasLinkedBy() bool {
	if o != nil && !IsNil(o.LinkedBy) {
		return true
	}

	return false
}

// SetLinkedBy gets a reference to the given string and assigns it to the LinkedBy field.
func (o *JuniperAccount) SetLinkedBy(v string) {
	o.LinkedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JuniperAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JuniperAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JuniperAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JuniperAccount) SetName(v string) {
	o.Name = &v
}

func (o JuniperAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JuniperAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkedBy) {
		toSerialize["linked_by"] = o.LinkedBy
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JuniperAccount) UnmarshalJSON(data []byte) (err error) {
	varJuniperAccount := _JuniperAccount{}

	err = json.Unmarshal(data, &varJuniperAccount)

	if err != nil {
		return err
	}

	*o = JuniperAccount(varJuniperAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "linked_by")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJuniperAccount struct {
	value *JuniperAccount
	isSet bool
}

func (v NullableJuniperAccount) Get() *JuniperAccount {
	return v.value
}

func (v *NullableJuniperAccount) Set(val *JuniperAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableJuniperAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableJuniperAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJuniperAccount(val *JuniperAccount) *NullableJuniperAccount {
	return &NullableJuniperAccount{value: val, isSet: true}
}

func (v NullableJuniperAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJuniperAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


