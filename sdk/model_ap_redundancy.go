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

// checks if the ApRedundancy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApRedundancy{}

// ApRedundancy struct for ApRedundancy
type ApRedundancy struct {
	// Property key is the node id
	Modules *map[string]ApRedundancyModule `json:"modules,omitempty"`
	NumAps *int32 `json:"num_aps,omitempty"`
	NumApsWithSwitchRedundancy *int32 `json:"num_aps_with_switch_redundancy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApRedundancy ApRedundancy

// NewApRedundancy instantiates a new ApRedundancy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApRedundancy() *ApRedundancy {
	this := ApRedundancy{}
	return &this
}

// NewApRedundancyWithDefaults instantiates a new ApRedundancy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApRedundancyWithDefaults() *ApRedundancy {
	this := ApRedundancy{}
	return &this
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *ApRedundancy) GetModules() map[string]ApRedundancyModule {
	if o == nil || IsNil(o.Modules) {
		var ret map[string]ApRedundancyModule
		return ret
	}
	return *o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRedundancy) GetModulesOk() (*map[string]ApRedundancyModule, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *ApRedundancy) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given map[string]ApRedundancyModule and assigns it to the Modules field.
func (o *ApRedundancy) SetModules(v map[string]ApRedundancyModule) {
	o.Modules = &v
}

// GetNumAps returns the NumAps field value if set, zero value otherwise.
func (o *ApRedundancy) GetNumAps() int32 {
	if o == nil || IsNil(o.NumAps) {
		var ret int32
		return ret
	}
	return *o.NumAps
}

// GetNumApsOk returns a tuple with the NumAps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRedundancy) GetNumApsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumAps) {
		return nil, false
	}
	return o.NumAps, true
}

// HasNumAps returns a boolean if a field has been set.
func (o *ApRedundancy) HasNumAps() bool {
	if o != nil && !IsNil(o.NumAps) {
		return true
	}

	return false
}

// SetNumAps gets a reference to the given int32 and assigns it to the NumAps field.
func (o *ApRedundancy) SetNumAps(v int32) {
	o.NumAps = &v
}

// GetNumApsWithSwitchRedundancy returns the NumApsWithSwitchRedundancy field value if set, zero value otherwise.
func (o *ApRedundancy) GetNumApsWithSwitchRedundancy() int32 {
	if o == nil || IsNil(o.NumApsWithSwitchRedundancy) {
		var ret int32
		return ret
	}
	return *o.NumApsWithSwitchRedundancy
}

// GetNumApsWithSwitchRedundancyOk returns a tuple with the NumApsWithSwitchRedundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRedundancy) GetNumApsWithSwitchRedundancyOk() (*int32, bool) {
	if o == nil || IsNil(o.NumApsWithSwitchRedundancy) {
		return nil, false
	}
	return o.NumApsWithSwitchRedundancy, true
}

// HasNumApsWithSwitchRedundancy returns a boolean if a field has been set.
func (o *ApRedundancy) HasNumApsWithSwitchRedundancy() bool {
	if o != nil && !IsNil(o.NumApsWithSwitchRedundancy) {
		return true
	}

	return false
}

// SetNumApsWithSwitchRedundancy gets a reference to the given int32 and assigns it to the NumApsWithSwitchRedundancy field.
func (o *ApRedundancy) SetNumApsWithSwitchRedundancy(v int32) {
	o.NumApsWithSwitchRedundancy = &v
}

func (o ApRedundancy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApRedundancy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.NumAps) {
		toSerialize["num_aps"] = o.NumAps
	}
	if !IsNil(o.NumApsWithSwitchRedundancy) {
		toSerialize["num_aps_with_switch_redundancy"] = o.NumApsWithSwitchRedundancy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApRedundancy) UnmarshalJSON(data []byte) (err error) {
	varApRedundancy := _ApRedundancy{}

	err = json.Unmarshal(data, &varApRedundancy)

	if err != nil {
		return err
	}

	*o = ApRedundancy(varApRedundancy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "modules")
		delete(additionalProperties, "num_aps")
		delete(additionalProperties, "num_aps_with_switch_redundancy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApRedundancy struct {
	value *ApRedundancy
	isSet bool
}

func (v NullableApRedundancy) Get() *ApRedundancy {
	return v.value
}

func (v *NullableApRedundancy) Set(val *ApRedundancy) {
	v.value = val
	v.isSet = true
}

func (v NullableApRedundancy) IsSet() bool {
	return v.isSet
}

func (v *NullableApRedundancy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApRedundancy(val *ApRedundancy) *NullableApRedundancy {
	return &NullableApRedundancy{value: val, isSet: true}
}

func (v NullableApRedundancy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApRedundancy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


