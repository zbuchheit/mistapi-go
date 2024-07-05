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

// checks if the MxedgeStatsTuntermStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeStatsTuntermStat{}

// MxedgeStatsTuntermStat struct for MxedgeStatsTuntermStat
type MxedgeStatsTuntermStat struct {
	MonitoringFailed *bool `json:"monitoring_failed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeStatsTuntermStat MxedgeStatsTuntermStat

// NewMxedgeStatsTuntermStat instantiates a new MxedgeStatsTuntermStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeStatsTuntermStat() *MxedgeStatsTuntermStat {
	this := MxedgeStatsTuntermStat{}
	return &this
}

// NewMxedgeStatsTuntermStatWithDefaults instantiates a new MxedgeStatsTuntermStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeStatsTuntermStatWithDefaults() *MxedgeStatsTuntermStat {
	this := MxedgeStatsTuntermStat{}
	return &this
}

// GetMonitoringFailed returns the MonitoringFailed field value if set, zero value otherwise.
func (o *MxedgeStatsTuntermStat) GetMonitoringFailed() bool {
	if o == nil || IsNil(o.MonitoringFailed) {
		var ret bool
		return ret
	}
	return *o.MonitoringFailed
}

// GetMonitoringFailedOk returns a tuple with the MonitoringFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeStatsTuntermStat) GetMonitoringFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.MonitoringFailed) {
		return nil, false
	}
	return o.MonitoringFailed, true
}

// HasMonitoringFailed returns a boolean if a field has been set.
func (o *MxedgeStatsTuntermStat) HasMonitoringFailed() bool {
	if o != nil && !IsNil(o.MonitoringFailed) {
		return true
	}

	return false
}

// SetMonitoringFailed gets a reference to the given bool and assigns it to the MonitoringFailed field.
func (o *MxedgeStatsTuntermStat) SetMonitoringFailed(v bool) {
	o.MonitoringFailed = &v
}

func (o MxedgeStatsTuntermStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeStatsTuntermStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MonitoringFailed) {
		toSerialize["monitoring_failed"] = o.MonitoringFailed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeStatsTuntermStat) UnmarshalJSON(data []byte) (err error) {
	varMxedgeStatsTuntermStat := _MxedgeStatsTuntermStat{}

	err = json.Unmarshal(data, &varMxedgeStatsTuntermStat)

	if err != nil {
		return err
	}

	*o = MxedgeStatsTuntermStat(varMxedgeStatsTuntermStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "monitoring_failed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeStatsTuntermStat struct {
	value *MxedgeStatsTuntermStat
	isSet bool
}

func (v NullableMxedgeStatsTuntermStat) Get() *MxedgeStatsTuntermStat {
	return v.value
}

func (v *NullableMxedgeStatsTuntermStat) Set(val *MxedgeStatsTuntermStat) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeStatsTuntermStat) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeStatsTuntermStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeStatsTuntermStat(val *MxedgeStatsTuntermStat) *NullableMxedgeStatsTuntermStat {
	return &NullableMxedgeStatsTuntermStat{value: val, isSet: true}
}

func (v NullableMxedgeStatsTuntermStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeStatsTuntermStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


