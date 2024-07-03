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

// checks if the ResponseSwitchPortSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSwitchPortSearch{}

// ResponseSwitchPortSearch struct for ResponseSwitchPortSearch
type ResponseSwitchPortSearch struct {
	End int32 `json:"end"`
	Limit int32 `json:"limit"`
	Results []SwitchPortStats `json:"results"`
	Start int32 `json:"start"`
	Total int32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSwitchPortSearch ResponseSwitchPortSearch

// NewResponseSwitchPortSearch instantiates a new ResponseSwitchPortSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSwitchPortSearch(end int32, limit int32, results []SwitchPortStats, start int32, total int32) *ResponseSwitchPortSearch {
	this := ResponseSwitchPortSearch{}
	this.End = end
	this.Limit = limit
	this.Results = results
	this.Start = start
	this.Total = total
	return &this
}

// NewResponseSwitchPortSearchWithDefaults instantiates a new ResponseSwitchPortSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSwitchPortSearchWithDefaults() *ResponseSwitchPortSearch {
	this := ResponseSwitchPortSearch{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseSwitchPortSearch) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseSwitchPortSearch) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseSwitchPortSearch) SetEnd(v int32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponseSwitchPortSearch) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponseSwitchPortSearch) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponseSwitchPortSearch) SetLimit(v int32) {
	o.Limit = v
}

// GetResults returns the Results field value
func (o *ResponseSwitchPortSearch) GetResults() []SwitchPortStats {
	if o == nil {
		var ret []SwitchPortStats
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseSwitchPortSearch) GetResultsOk() ([]SwitchPortStats, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseSwitchPortSearch) SetResults(v []SwitchPortStats) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponseSwitchPortSearch) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseSwitchPortSearch) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseSwitchPortSearch) SetStart(v int32) {
	o.Start = v
}

// GetTotal returns the Total field value
func (o *ResponseSwitchPortSearch) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResponseSwitchPortSearch) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResponseSwitchPortSearch) SetTotal(v int32) {
	o.Total = v
}

func (o ResponseSwitchPortSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSwitchPortSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *ResponseSwitchPortSearch) UnmarshalJSON(data []byte) (err error) {
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

	varResponseSwitchPortSearch := _ResponseSwitchPortSearch{}

	err = json.Unmarshal(data, &varResponseSwitchPortSearch)

	if err != nil {
		return err
	}

	*o = ResponseSwitchPortSearch(varResponseSwitchPortSearch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSwitchPortSearch struct {
	value *ResponseSwitchPortSearch
	isSet bool
}

func (v NullableResponseSwitchPortSearch) Get() *ResponseSwitchPortSearch {
	return v.value
}

func (v *NullableResponseSwitchPortSearch) Set(val *ResponseSwitchPortSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSwitchPortSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSwitchPortSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSwitchPortSearch(val *ResponseSwitchPortSearch) *NullableResponseSwitchPortSearch {
	return &NullableResponseSwitchPortSearch{value: val, isSet: true}
}

func (v NullableResponseSwitchPortSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSwitchPortSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


