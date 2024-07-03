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

// checks if the ModuleStatTemperaturesItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleStatTemperaturesItems{}

// ModuleStatTemperaturesItems struct for ModuleStatTemperaturesItems
type ModuleStatTemperaturesItems struct {
	Celsius *float32 `json:"celsius,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleStatTemperaturesItems ModuleStatTemperaturesItems

// NewModuleStatTemperaturesItems instantiates a new ModuleStatTemperaturesItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleStatTemperaturesItems() *ModuleStatTemperaturesItems {
	this := ModuleStatTemperaturesItems{}
	return &this
}

// NewModuleStatTemperaturesItemsWithDefaults instantiates a new ModuleStatTemperaturesItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleStatTemperaturesItemsWithDefaults() *ModuleStatTemperaturesItems {
	this := ModuleStatTemperaturesItems{}
	return &this
}

// GetCelsius returns the Celsius field value if set, zero value otherwise.
func (o *ModuleStatTemperaturesItems) GetCelsius() float32 {
	if o == nil || IsNil(o.Celsius) {
		var ret float32
		return ret
	}
	return *o.Celsius
}

// GetCelsiusOk returns a tuple with the Celsius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatTemperaturesItems) GetCelsiusOk() (*float32, bool) {
	if o == nil || IsNil(o.Celsius) {
		return nil, false
	}
	return o.Celsius, true
}

// HasCelsius returns a boolean if a field has been set.
func (o *ModuleStatTemperaturesItems) HasCelsius() bool {
	if o != nil && !IsNil(o.Celsius) {
		return true
	}

	return false
}

// SetCelsius gets a reference to the given float32 and assigns it to the Celsius field.
func (o *ModuleStatTemperaturesItems) SetCelsius(v float32) {
	o.Celsius = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModuleStatTemperaturesItems) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatTemperaturesItems) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModuleStatTemperaturesItems) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModuleStatTemperaturesItems) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModuleStatTemperaturesItems) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleStatTemperaturesItems) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModuleStatTemperaturesItems) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModuleStatTemperaturesItems) SetStatus(v string) {
	o.Status = &v
}

func (o ModuleStatTemperaturesItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleStatTemperaturesItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Celsius) {
		toSerialize["celsius"] = o.Celsius
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModuleStatTemperaturesItems) UnmarshalJSON(data []byte) (err error) {
	varModuleStatTemperaturesItems := _ModuleStatTemperaturesItems{}

	err = json.Unmarshal(data, &varModuleStatTemperaturesItems)

	if err != nil {
		return err
	}

	*o = ModuleStatTemperaturesItems(varModuleStatTemperaturesItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "celsius")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleStatTemperaturesItems struct {
	value *ModuleStatTemperaturesItems
	isSet bool
}

func (v NullableModuleStatTemperaturesItems) Get() *ModuleStatTemperaturesItems {
	return v.value
}

func (v *NullableModuleStatTemperaturesItems) Set(val *ModuleStatTemperaturesItems) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleStatTemperaturesItems) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleStatTemperaturesItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleStatTemperaturesItems(val *ModuleStatTemperaturesItems) *NullableModuleStatTemperaturesItems {
	return &NullableModuleStatTemperaturesItems{value: val, isSet: true}
}

func (v NullableModuleStatTemperaturesItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleStatTemperaturesItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


