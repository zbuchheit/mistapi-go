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
	"fmt"
)

// checks if the SleClassifierSamples type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleClassifierSamples{}

// SleClassifierSamples struct for SleClassifierSamples
type SleClassifierSamples struct {
	Degraded []float32 `json:"degraded"`
	Duration []float32 `json:"duration"`
	Total []float32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _SleClassifierSamples SleClassifierSamples

// NewSleClassifierSamples instantiates a new SleClassifierSamples object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleClassifierSamples(degraded []float32, duration []float32, total []float32) *SleClassifierSamples {
	this := SleClassifierSamples{}
	this.Degraded = degraded
	this.Duration = duration
	this.Total = total
	return &this
}

// NewSleClassifierSamplesWithDefaults instantiates a new SleClassifierSamples object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleClassifierSamplesWithDefaults() *SleClassifierSamples {
	this := SleClassifierSamples{}
	return &this
}

// GetDegraded returns the Degraded field value
func (o *SleClassifierSamples) GetDegraded() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Degraded
}

// GetDegradedOk returns a tuple with the Degraded field value
// and a boolean to check if the value has been set.
func (o *SleClassifierSamples) GetDegradedOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Degraded, true
}

// SetDegraded sets field value
func (o *SleClassifierSamples) SetDegraded(v []float32) {
	o.Degraded = v
}

// GetDuration returns the Duration field value
func (o *SleClassifierSamples) GetDuration() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SleClassifierSamples) GetDurationOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration, true
}

// SetDuration sets field value
func (o *SleClassifierSamples) SetDuration(v []float32) {
	o.Duration = v
}

// GetTotal returns the Total field value
func (o *SleClassifierSamples) GetTotal() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SleClassifierSamples) GetTotalOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total, true
}

// SetTotal sets field value
func (o *SleClassifierSamples) SetTotal(v []float32) {
	o.Total = v
}

func (o SleClassifierSamples) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleClassifierSamples) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["degraded"] = o.Degraded
	toSerialize["duration"] = o.Duration
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleClassifierSamples) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"degraded",
		"duration",
		"total",
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

	varSleClassifierSamples := _SleClassifierSamples{}

	err = json.Unmarshal(data, &varSleClassifierSamples)

	if err != nil {
		return err
	}

	*o = SleClassifierSamples(varSleClassifierSamples)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "degraded")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleClassifierSamples struct {
	value *SleClassifierSamples
	isSet bool
}

func (v NullableSleClassifierSamples) Get() *SleClassifierSamples {
	return v.value
}

func (v *NullableSleClassifierSamples) Set(val *SleClassifierSamples) {
	v.value = val
	v.isSet = true
}

func (v NullableSleClassifierSamples) IsSet() bool {
	return v.isSet
}

func (v *NullableSleClassifierSamples) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleClassifierSamples(val *SleClassifierSamples) *NullableSleClassifierSamples {
	return &NullableSleClassifierSamples{value: val, isSet: true}
}

func (v NullableSleClassifierSamples) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleClassifierSamples) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


