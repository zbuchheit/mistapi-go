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

// checks if the BgpConfigNeighbors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpConfigNeighbors{}

// BgpConfigNeighbors struct for BgpConfigNeighbors
type BgpConfigNeighbors struct {
	// If true, the BGP session to this neighbor will be administratively disabled/shutdown
	Disabled *bool `json:"disabled,omitempty"`
	ExportPolicy *string `json:"export_policy,omitempty"`
	HoldTime *int32 `json:"hold_time,omitempty"`
	ImportPolicy *string `json:"import_policy,omitempty"`
	// assuming BGP neighbor is directly connected
	MultihopTtl *int32 `json:"multihop_ttl,omitempty"`
	NeighborAs *int32 `json:"neighbor_as,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpConfigNeighbors BgpConfigNeighbors

// NewBgpConfigNeighbors instantiates a new BgpConfigNeighbors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfigNeighbors() *BgpConfigNeighbors {
	this := BgpConfigNeighbors{}
	var disabled bool = false
	this.Disabled = &disabled
	var holdTime int32 = 90
	this.HoldTime = &holdTime
	return &this
}

// NewBgpConfigNeighborsWithDefaults instantiates a new BgpConfigNeighbors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigNeighborsWithDefaults() *BgpConfigNeighbors {
	this := BgpConfigNeighbors{}
	var disabled bool = false
	this.Disabled = &disabled
	var holdTime int32 = 90
	this.HoldTime = &holdTime
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *BgpConfigNeighbors) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigNeighbors) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *BgpConfigNeighbors) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *BgpConfigNeighbors) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetExportPolicy returns the ExportPolicy field value if set, zero value otherwise.
func (o *BgpConfigNeighbors) GetExportPolicy() string {
	if o == nil || IsNil(o.ExportPolicy) {
		var ret string
		return ret
	}
	return *o.ExportPolicy
}

// GetExportPolicyOk returns a tuple with the ExportPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigNeighbors) GetExportPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.ExportPolicy) {
		return nil, false
	}
	return o.ExportPolicy, true
}

// HasExportPolicy returns a boolean if a field has been set.
func (o *BgpConfigNeighbors) HasExportPolicy() bool {
	if o != nil && !IsNil(o.ExportPolicy) {
		return true
	}

	return false
}

// SetExportPolicy gets a reference to the given string and assigns it to the ExportPolicy field.
func (o *BgpConfigNeighbors) SetExportPolicy(v string) {
	o.ExportPolicy = &v
}

// GetHoldTime returns the HoldTime field value if set, zero value otherwise.
func (o *BgpConfigNeighbors) GetHoldTime() int32 {
	if o == nil || IsNil(o.HoldTime) {
		var ret int32
		return ret
	}
	return *o.HoldTime
}

// GetHoldTimeOk returns a tuple with the HoldTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigNeighbors) GetHoldTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.HoldTime) {
		return nil, false
	}
	return o.HoldTime, true
}

// HasHoldTime returns a boolean if a field has been set.
func (o *BgpConfigNeighbors) HasHoldTime() bool {
	if o != nil && !IsNil(o.HoldTime) {
		return true
	}

	return false
}

// SetHoldTime gets a reference to the given int32 and assigns it to the HoldTime field.
func (o *BgpConfigNeighbors) SetHoldTime(v int32) {
	o.HoldTime = &v
}

// GetImportPolicy returns the ImportPolicy field value if set, zero value otherwise.
func (o *BgpConfigNeighbors) GetImportPolicy() string {
	if o == nil || IsNil(o.ImportPolicy) {
		var ret string
		return ret
	}
	return *o.ImportPolicy
}

// GetImportPolicyOk returns a tuple with the ImportPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigNeighbors) GetImportPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.ImportPolicy) {
		return nil, false
	}
	return o.ImportPolicy, true
}

// HasImportPolicy returns a boolean if a field has been set.
func (o *BgpConfigNeighbors) HasImportPolicy() bool {
	if o != nil && !IsNil(o.ImportPolicy) {
		return true
	}

	return false
}

// SetImportPolicy gets a reference to the given string and assigns it to the ImportPolicy field.
func (o *BgpConfigNeighbors) SetImportPolicy(v string) {
	o.ImportPolicy = &v
}

// GetMultihopTtl returns the MultihopTtl field value if set, zero value otherwise.
func (o *BgpConfigNeighbors) GetMultihopTtl() int32 {
	if o == nil || IsNil(o.MultihopTtl) {
		var ret int32
		return ret
	}
	return *o.MultihopTtl
}

// GetMultihopTtlOk returns a tuple with the MultihopTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigNeighbors) GetMultihopTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.MultihopTtl) {
		return nil, false
	}
	return o.MultihopTtl, true
}

// HasMultihopTtl returns a boolean if a field has been set.
func (o *BgpConfigNeighbors) HasMultihopTtl() bool {
	if o != nil && !IsNil(o.MultihopTtl) {
		return true
	}

	return false
}

// SetMultihopTtl gets a reference to the given int32 and assigns it to the MultihopTtl field.
func (o *BgpConfigNeighbors) SetMultihopTtl(v int32) {
	o.MultihopTtl = &v
}

// GetNeighborAs returns the NeighborAs field value if set, zero value otherwise.
func (o *BgpConfigNeighbors) GetNeighborAs() int32 {
	if o == nil || IsNil(o.NeighborAs) {
		var ret int32
		return ret
	}
	return *o.NeighborAs
}

// GetNeighborAsOk returns a tuple with the NeighborAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfigNeighbors) GetNeighborAsOk() (*int32, bool) {
	if o == nil || IsNil(o.NeighborAs) {
		return nil, false
	}
	return o.NeighborAs, true
}

// HasNeighborAs returns a boolean if a field has been set.
func (o *BgpConfigNeighbors) HasNeighborAs() bool {
	if o != nil && !IsNil(o.NeighborAs) {
		return true
	}

	return false
}

// SetNeighborAs gets a reference to the given int32 and assigns it to the NeighborAs field.
func (o *BgpConfigNeighbors) SetNeighborAs(v int32) {
	o.NeighborAs = &v
}

func (o BgpConfigNeighbors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpConfigNeighbors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.ExportPolicy) {
		toSerialize["export_policy"] = o.ExportPolicy
	}
	if !IsNil(o.HoldTime) {
		toSerialize["hold_time"] = o.HoldTime
	}
	if !IsNil(o.ImportPolicy) {
		toSerialize["import_policy"] = o.ImportPolicy
	}
	if !IsNil(o.MultihopTtl) {
		toSerialize["multihop_ttl"] = o.MultihopTtl
	}
	if !IsNil(o.NeighborAs) {
		toSerialize["neighbor_as"] = o.NeighborAs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpConfigNeighbors) UnmarshalJSON(data []byte) (err error) {
	varBgpConfigNeighbors := _BgpConfigNeighbors{}

	err = json.Unmarshal(data, &varBgpConfigNeighbors)

	if err != nil {
		return err
	}

	*o = BgpConfigNeighbors(varBgpConfigNeighbors)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "export_policy")
		delete(additionalProperties, "hold_time")
		delete(additionalProperties, "import_policy")
		delete(additionalProperties, "multihop_ttl")
		delete(additionalProperties, "neighbor_as")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpConfigNeighbors struct {
	value *BgpConfigNeighbors
	isSet bool
}

func (v NullableBgpConfigNeighbors) Get() *BgpConfigNeighbors {
	return v.value
}

func (v *NullableBgpConfigNeighbors) Set(val *BgpConfigNeighbors) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfigNeighbors) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfigNeighbors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfigNeighbors(val *BgpConfigNeighbors) *NullableBgpConfigNeighbors {
	return &NullableBgpConfigNeighbors{value: val, isSet: true}
}

func (v NullableBgpConfigNeighbors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfigNeighbors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


