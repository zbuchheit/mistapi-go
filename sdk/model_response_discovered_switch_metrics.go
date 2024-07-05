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
	"fmt"
)

// checks if the ResponseDiscoveredSwitchMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDiscoveredSwitchMetrics{}

// ResponseDiscoveredSwitchMetrics struct for ResponseDiscoveredSwitchMetrics
type ResponseDiscoveredSwitchMetrics struct {
	End float32 `json:"end"`
	Limit int32 `json:"limit"`
	Next *string `json:"next,omitempty"`
	Results []DiscoveredSwitchMetric `json:"results"`
	Start float32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDiscoveredSwitchMetrics ResponseDiscoveredSwitchMetrics

// NewResponseDiscoveredSwitchMetrics instantiates a new ResponseDiscoveredSwitchMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDiscoveredSwitchMetrics(end float32, limit int32, results []DiscoveredSwitchMetric, start float32, total int32) *ResponseDiscoveredSwitchMetrics {
	this := ResponseDiscoveredSwitchMetrics{}
	this.End = end
	this.Limit = limit
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewResponseDiscoveredSwitchMetricsWithDefaults instantiates a new ResponseDiscoveredSwitchMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDiscoveredSwitchMetricsWithDefaults() *ResponseDiscoveredSwitchMetrics {
	this := ResponseDiscoveredSwitchMetrics{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseDiscoveredSwitchMetrics) GetEnd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseDiscoveredSwitchMetrics) GetEndOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseDiscoveredSwitchMetrics) SetEnd(v float32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponseDiscoveredSwitchMetrics) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponseDiscoveredSwitchMetrics) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponseDiscoveredSwitchMetrics) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ResponseDiscoveredSwitchMetrics) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDiscoveredSwitchMetrics) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ResponseDiscoveredSwitchMetrics) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ResponseDiscoveredSwitchMetrics) SetNext(v string) {
	o.Next = &v
}

// GetResults returns the Results field value
func (o *ResponseDiscoveredSwitchMetrics) GetResults() []DiscoveredSwitchMetric {
	if o == nil {
		var ret []DiscoveredSwitchMetric
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseDiscoveredSwitchMetrics) GetResultsOk() ([]DiscoveredSwitchMetric, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseDiscoveredSwitchMetrics) SetResults(v []DiscoveredSwitchMetric) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponseDiscoveredSwitchMetrics) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseDiscoveredSwitchMetrics) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseDiscoveredSwitchMetrics) SetStart(v float32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *ResponseDiscoveredSwitchMetrics) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResponseDiscoveredSwitchMetrics) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResponseDiscoveredSwitchMetrics) SetTotal(v int32) {
	o.Total = v
}

func (o ResponseDiscoveredSwitchMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDiscoveredSwitchMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["limit"] = o.Limit
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDiscoveredSwitchMetrics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"limit",
		"results",
		"start",
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

	varResponseDiscoveredSwitchMetrics := _ResponseDiscoveredSwitchMetrics{}

	err = json.Unmarshal(data, &varResponseDiscoveredSwitchMetrics)

	if err != nil {
		return err
	}

	*o = ResponseDiscoveredSwitchMetrics(varResponseDiscoveredSwitchMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDiscoveredSwitchMetrics struct {
	value *ResponseDiscoveredSwitchMetrics
	isSet bool
}

func (v NullableResponseDiscoveredSwitchMetrics) Get() *ResponseDiscoveredSwitchMetrics {
	return v.value
}

func (v *NullableResponseDiscoveredSwitchMetrics) Set(val *ResponseDiscoveredSwitchMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDiscoveredSwitchMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDiscoveredSwitchMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDiscoveredSwitchMetrics(val *ResponseDiscoveredSwitchMetrics) *NullableResponseDiscoveredSwitchMetrics {
	return &NullableResponseDiscoveredSwitchMetrics{value: val, isSet: true}
}

func (v NullableResponseDiscoveredSwitchMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDiscoveredSwitchMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


