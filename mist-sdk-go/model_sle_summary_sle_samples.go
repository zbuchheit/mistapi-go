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

// checks if the SleSummarySleSamples type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleSummarySleSamples{}

// SleSummarySleSamples struct for SleSummarySleSamples
type SleSummarySleSamples struct {
	Degraded []float32 `json:"degraded"`
	Total []float32 `json:"total"`
	Value []float32 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _SleSummarySleSamples SleSummarySleSamples

// NewSleSummarySleSamples instantiates a new SleSummarySleSamples object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleSummarySleSamples(degraded []float32, total []float32, value []float32) *SleSummarySleSamples {
	this := SleSummarySleSamples{}
	this.Degraded = degraded
	this.Total = total
	this.Value = value
	return &this
}

// NewSleSummarySleSamplesWithDefaults instantiates a new SleSummarySleSamples object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleSummarySleSamplesWithDefaults() *SleSummarySleSamples {
	this := SleSummarySleSamples{}
	return &this
}

// GetDegraded returns the Degraded field value
func (o *SleSummarySleSamples) GetDegraded() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Degraded
}

// GetDegradedOk returns a tuple with the Degraded field value
// and a boolean to check if the value has been set.
func (o *SleSummarySleSamples) GetDegradedOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Degraded, true
}

// SetDegraded sets field value
func (o *SleSummarySleSamples) SetDegraded(v []float32) {
	o.Degraded = v
}

// GetTotal returns the Total field value
func (o *SleSummarySleSamples) GetTotal() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SleSummarySleSamples) GetTotalOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total, true
}

// SetTotal sets field value
func (o *SleSummarySleSamples) SetTotal(v []float32) {
	o.Total = v
}

// GetValue returns the Value field value
func (o *SleSummarySleSamples) GetValue() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SleSummarySleSamples) GetValueOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *SleSummarySleSamples) SetValue(v []float32) {
	o.Value = v
}

func (o SleSummarySleSamples) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleSummarySleSamples) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["degraded"] = o.Degraded
	toSerialize["total"] = o.Total
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleSummarySleSamples) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"degraded",
		"total",
		"value",
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

	varSleSummarySleSamples := _SleSummarySleSamples{}

	err = json.Unmarshal(data, &varSleSummarySleSamples)

	if err != nil {
		return err
	}

	*o = SleSummarySleSamples(varSleSummarySleSamples)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "degraded")
		delete(additionalProperties, "total")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleSummarySleSamples struct {
	value *SleSummarySleSamples
	isSet bool
}

func (v NullableSleSummarySleSamples) Get() *SleSummarySleSamples {
	return v.value
}

func (v *NullableSleSummarySleSamples) Set(val *SleSummarySleSamples) {
	v.value = val
	v.isSet = true
}

func (v NullableSleSummarySleSamples) IsSet() bool {
	return v.isSet
}

func (v *NullableSleSummarySleSamples) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleSummarySleSamples(val *SleSummarySleSamples) *NullableSleSummarySleSamples {
	return &NullableSleSummarySleSamples{value: val, isSet: true}
}

func (v NullableSleSummarySleSamples) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleSummarySleSamples) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


