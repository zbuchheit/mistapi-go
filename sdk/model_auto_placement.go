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

// checks if the AutoPlacement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoPlacement{}

// AutoPlacement struct for AutoPlacement
type AutoPlacement struct {
	// * If `force_collection`==`false`: the API Iattempts to start localization with existing data.  * If `force_collection`==`true`: maintenance the API attempts to start orchestration.
	ForceCollection *bool `json:"force_collection,omitempty"`
	// list of device macs
	Macs []string `json:"macs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoPlacement AutoPlacement

// NewAutoPlacement instantiates a new AutoPlacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoPlacement() *AutoPlacement {
	this := AutoPlacement{}
	var forceCollection bool = false
	this.ForceCollection = &forceCollection
	return &this
}

// NewAutoPlacementWithDefaults instantiates a new AutoPlacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoPlacementWithDefaults() *AutoPlacement {
	this := AutoPlacement{}
	var forceCollection bool = false
	this.ForceCollection = &forceCollection
	return &this
}

// GetForceCollection returns the ForceCollection field value if set, zero value otherwise.
func (o *AutoPlacement) GetForceCollection() bool {
	if o == nil || IsNil(o.ForceCollection) {
		var ret bool
		return ret
	}
	return *o.ForceCollection
}

// GetForceCollectionOk returns a tuple with the ForceCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPlacement) GetForceCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceCollection) {
		return nil, false
	}
	return o.ForceCollection, true
}

// HasForceCollection returns a boolean if a field has been set.
func (o *AutoPlacement) HasForceCollection() bool {
	if o != nil && !IsNil(o.ForceCollection) {
		return true
	}

	return false
}

// SetForceCollection gets a reference to the given bool and assigns it to the ForceCollection field.
func (o *AutoPlacement) SetForceCollection(v bool) {
	o.ForceCollection = &v
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *AutoPlacement) GetMacs() []string {
	if o == nil || IsNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoPlacement) GetMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *AutoPlacement) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *AutoPlacement) SetMacs(v []string) {
	o.Macs = v
}

func (o AutoPlacement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoPlacement) ToMap() (map[string]interface{}, error) {
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

func (o *AutoPlacement) UnmarshalJSON(data []byte) (err error) {
	varAutoPlacement := _AutoPlacement{}

	err = json.Unmarshal(data, &varAutoPlacement)

	if err != nil {
		return err
	}

	*o = AutoPlacement(varAutoPlacement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "force_collection")
		delete(additionalProperties, "macs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoPlacement struct {
	value *AutoPlacement
	isSet bool
}

func (v NullableAutoPlacement) Get() *AutoPlacement {
	return v.value
}

func (v *NullableAutoPlacement) Set(val *AutoPlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoPlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoPlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoPlacement(val *AutoPlacement) *NullableAutoPlacement {
	return &NullableAutoPlacement{value: val, isSet: true}
}

func (v NullableAutoPlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoPlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


