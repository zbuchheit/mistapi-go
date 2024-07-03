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

// checks if the InsightMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsightMetrics{}

// InsightMetrics struct for InsightMetrics
type InsightMetrics struct {
	End int32 `json:"end"`
	Interval int32 `json:"interval"`
	// results depends on the `metric`
	Results []map[string]interface{} `json:"results"`
	Start int32 `json:"start"`
	AdditionalProperties map[string]interface{}
}

type _InsightMetrics InsightMetrics

// NewInsightMetrics instantiates a new InsightMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsightMetrics(end int32, interval int32, results []map[string]interface{}, start int32) *InsightMetrics {
	this := InsightMetrics{}
	this.End = end
	this.Interval = interval
	this.Results = results
	this.Start = start
	return &this
}

// NewInsightMetricsWithDefaults instantiates a new InsightMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsightMetricsWithDefaults() *InsightMetrics {
	this := InsightMetrics{}
	return &this
}

// GetEnd returns the End field value
func (o *InsightMetrics) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *InsightMetrics) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *InsightMetrics) SetEnd(v int32) {
	o.End = v
}

// GetInterval returns the Interval field value
func (o *InsightMetrics) GetInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *InsightMetrics) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *InsightMetrics) SetInterval(v int32) {
	o.Interval = v
}

// GetResults returns the Results field value
func (o *InsightMetrics) GetResults() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *InsightMetrics) GetResultsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *InsightMetrics) SetResults(v []map[string]interface{}) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *InsightMetrics) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *InsightMetrics) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *InsightMetrics) SetStart(v int32) {
	o.Start = v
}

func (o InsightMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsightMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["interval"] = o.Interval
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InsightMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"interval",
		"results",
		"start",
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

	varInsightMetrics := _InsightMetrics{}

	err = json.Unmarshal(data, &varInsightMetrics)

	if err != nil {
		return err
	}

	*o = InsightMetrics(varInsightMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInsightMetrics struct {
	value *InsightMetrics
	isSet bool
}

func (v NullableInsightMetrics) Get() *InsightMetrics {
	return v.value
}

func (v *NullableInsightMetrics) Set(val *InsightMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableInsightMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableInsightMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsightMetrics(val *InsightMetrics) *NullableInsightMetrics {
	return &NullableInsightMetrics{value: val, isSet: true}
}

func (v NullableInsightMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsightMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


