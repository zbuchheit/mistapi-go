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

// checks if the AutoOrient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoOrient{}

// AutoOrient struct for AutoOrient
type AutoOrient struct {
	// If `force_collection`==`false`, the API attempts to start auto orientation with existing BLE data.  If `force_collection`==`true`, the API attempts to start BLE orchestration.
	ForceCollection *bool `json:"force_collection,omitempty"`
	// list of device macs
	Macs []string `json:"macs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoOrient AutoOrient

// NewAutoOrient instantiates a new AutoOrient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoOrient() *AutoOrient {
	this := AutoOrient{}
	var forceCollection bool = false
	this.ForceCollection = &forceCollection
	return &this
}

// NewAutoOrientWithDefaults instantiates a new AutoOrient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoOrientWithDefaults() *AutoOrient {
	this := AutoOrient{}
	var forceCollection bool = false
	this.ForceCollection = &forceCollection
	return &this
}

// GetForceCollection returns the ForceCollection field value if set, zero value otherwise.
func (o *AutoOrient) GetForceCollection() bool {
	if o == nil || IsNil(o.ForceCollection) {
		var ret bool
		return ret
	}
	return *o.ForceCollection
}

// GetForceCollectionOk returns a tuple with the ForceCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoOrient) GetForceCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceCollection) {
		return nil, false
	}
	return o.ForceCollection, true
}

// HasForceCollection returns a boolean if a field has been set.
func (o *AutoOrient) HasForceCollection() bool {
	if o != nil && !IsNil(o.ForceCollection) {
		return true
	}

	return false
}

// SetForceCollection gets a reference to the given bool and assigns it to the ForceCollection field.
func (o *AutoOrient) SetForceCollection(v bool) {
	o.ForceCollection = &v
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *AutoOrient) GetMacs() []string {
	if o == nil || IsNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoOrient) GetMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *AutoOrient) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *AutoOrient) SetMacs(v []string) {
	o.Macs = v
}

func (o AutoOrient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoOrient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForceCollection) {
		toSerialize["force_collection"] = o.ForceCollection
	}
	if !IsNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoOrient) UnmarshalJSON(data []byte) (err error) {
	varAutoOrient := _AutoOrient{}

	err = json.Unmarshal(data, &varAutoOrient)

	if err != nil {
		return err
	}

	*o = AutoOrient(varAutoOrient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "force_collection")
		delete(additionalProperties, "macs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoOrient struct {
	value *AutoOrient
	isSet bool
}

func (v NullableAutoOrient) Get() *AutoOrient {
	return v.value
}

func (v *NullableAutoOrient) Set(val *AutoOrient) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoOrient) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoOrient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoOrient(val *AutoOrient) *NullableAutoOrient {
	return &NullableAutoOrient{value: val, isSet: true}
}

func (v NullableAutoOrient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoOrient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


