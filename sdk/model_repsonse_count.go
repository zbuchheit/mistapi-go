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

// checks if the RepsonseCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepsonseCount{}

// RepsonseCount struct for RepsonseCount
type RepsonseCount struct {
	Distinct string `json:"distinct"`
	End int32 `json:"end"`
	Limit int32 `json:"limit"`
	Results []CountResult `json:"results"`
	Start int32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _RepsonseCount RepsonseCount

// NewRepsonseCount instantiates a new RepsonseCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepsonseCount(distinct string, end int32, limit int32, results []CountResult, start int32, total int32) *RepsonseCount {
	this := RepsonseCount{}
	this.Distinct = distinct
	this.End = end
	this.Limit = limit
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewRepsonseCountWithDefaults instantiates a new RepsonseCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepsonseCountWithDefaults() *RepsonseCount {
	this := RepsonseCount{}
	return &this
}

// GetDistinct returns the Distinct field value
func (o *RepsonseCount) GetDistinct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distinct
}

// GetDistinctOk returns a tuple with the Distinct field value
// and a boolean to check if the value has been set.
func (o *RepsonseCount) GetDistinctOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distinct, true
}

// SetDistinct sets field value
func (o *RepsonseCount) SetDistinct(v string) {
	o.Distinct = v
}

// GetEnd returns the End field value
func (o *RepsonseCount) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *RepsonseCount) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *RepsonseCount) SetEnd(v int32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *RepsonseCount) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *RepsonseCount) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *RepsonseCount) SetLimit(v int32) {
	o.Limit = v
}

// GetResults returns the Results field value
func (o *RepsonseCount) GetResults() []CountResult {
	if o == nil {
		var ret []CountResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *RepsonseCount) GetResultsOk() ([]CountResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *RepsonseCount) SetResults(v []CountResult) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *RepsonseCount) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *RepsonseCount) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *RepsonseCount) SetStart(v int32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *RepsonseCount) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *RepsonseCount) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *RepsonseCount) SetTotal(v int32) {
	o.Total = v
}

func (o RepsonseCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepsonseCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["distinct"] = o.Distinct
	toSerialize["end"] = o.End
	toSerialize["limit"] = o.Limit
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RepsonseCount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"distinct",
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

	varRepsonseCount := _RepsonseCount{}

	err = json.Unmarshal(data, &varRepsonseCount)

	if err != nil {
		return err
	}

	*o = RepsonseCount(varRepsonseCount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "distinct")
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRepsonseCount struct {
	value *RepsonseCount
	isSet bool
}

func (v NullableRepsonseCount) Get() *RepsonseCount {
	return v.value
}

func (v *NullableRepsonseCount) Set(val *RepsonseCount) {
	v.value = val
	v.isSet = true
}

func (v NullableRepsonseCount) IsSet() bool {
	return v.isSet
}

func (v *NullableRepsonseCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepsonseCount(val *RepsonseCount) *NullableRepsonseCount {
	return &NullableRepsonseCount{value: val, isSet: true}
}

func (v NullableRepsonseCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepsonseCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


