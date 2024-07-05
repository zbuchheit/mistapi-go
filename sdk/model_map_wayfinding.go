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

// checks if the MapWayfinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapWayfinding{}

// MapWayfinding properties related to wayfinding
type MapWayfinding struct {
	Micello *MapWayfindingMicello `json:"micello,omitempty"`
	SnapToPath *bool `json:"snap_to_path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MapWayfinding MapWayfinding

// NewMapWayfinding instantiates a new MapWayfinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapWayfinding() *MapWayfinding {
	this := MapWayfinding{}
	return &this
}

// NewMapWayfindingWithDefaults instantiates a new MapWayfinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapWayfindingWithDefaults() *MapWayfinding {
	this := MapWayfinding{}
	return &this
}

// GetMicello returns the Micello field value if set, zero value otherwise.
func (o *MapWayfinding) GetMicello() MapWayfindingMicello {
	if o == nil || IsNil(o.Micello) {
		var ret MapWayfindingMicello
		return ret
	}
	return *o.Micello
}

// GetMicelloOk returns a tuple with the Micello field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapWayfinding) GetMicelloOk() (*MapWayfindingMicello, bool) {
	if o == nil || IsNil(o.Micello) {
		return nil, false
	}
	return o.Micello, true
}

// HasMicello returns a boolean if a field has been set.
func (o *MapWayfinding) HasMicello() bool {
	if o != nil && !IsNil(o.Micello) {
		return true
	}

	return false
}

// SetMicello gets a reference to the given MapWayfindingMicello and assigns it to the Micello field.
func (o *MapWayfinding) SetMicello(v MapWayfindingMicello) {
	o.Micello = &v
}

// GetSnapToPath returns the SnapToPath field value if set, zero value otherwise.
func (o *MapWayfinding) GetSnapToPath() bool {
	if o == nil || IsNil(o.SnapToPath) {
		var ret bool
		return ret
	}
	return *o.SnapToPath
}

// GetSnapToPathOk returns a tuple with the SnapToPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapWayfinding) GetSnapToPathOk() (*bool, bool) {
	if o == nil || IsNil(o.SnapToPath) {
		return nil, false
	}
	return o.SnapToPath, true
}

// HasSnapToPath returns a boolean if a field has been set.
func (o *MapWayfinding) HasSnapToPath() bool {
	if o != nil && !IsNil(o.SnapToPath) {
		return true
	}

	return false
}

// SetSnapToPath gets a reference to the given bool and assigns it to the SnapToPath field.
func (o *MapWayfinding) SetSnapToPath(v bool) {
	o.SnapToPath = &v
}

func (o MapWayfinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapWayfinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Micello) {
		toSerialize["micello"] = o.Micello
	}
	if !IsNil(o.SnapToPath) {
		toSerialize["snap_to_path"] = o.SnapToPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MapWayfinding) UnmarshalJSON(data []byte) (err error) {
	varMapWayfinding := _MapWayfinding{}

	err = json.Unmarshal(data, &varMapWayfinding)

	if err != nil {
		return err
	}

	*o = MapWayfinding(varMapWayfinding)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "micello")
		delete(additionalProperties, "snap_to_path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMapWayfinding struct {
	value *MapWayfinding
	isSet bool
}

func (v NullableMapWayfinding) Get() *MapWayfinding {
	return v.value
}

func (v *NullableMapWayfinding) Set(val *MapWayfinding) {
	v.value = val
	v.isSet = true
}

func (v NullableMapWayfinding) IsSet() bool {
	return v.isSet
}

func (v *NullableMapWayfinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapWayfinding(val *MapWayfinding) *NullableMapWayfinding {
	return &NullableMapWayfinding{value: val, isSet: true}
}

func (v NullableMapWayfinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapWayfinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


