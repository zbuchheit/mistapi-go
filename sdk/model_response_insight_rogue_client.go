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
	"fmt"
)

// checks if the ResponseInsightRogueClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseInsightRogueClient{}

// ResponseInsightRogueClient struct for ResponseInsightRogueClient
type ResponseInsightRogueClient struct {
	End int32 `json:"end"`
	Limit int32 `json:"limit"`
	Next string `json:"next"`
	Results []InsightRogueClient `json:"results"`
	Start int32 `json:"start"`
	AdditionalProperties map[string]interface{}
}

type _ResponseInsightRogueClient ResponseInsightRogueClient

// NewResponseInsightRogueClient instantiates a new ResponseInsightRogueClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseInsightRogueClient(end int32, limit int32, next string, results []InsightRogueClient, start int32) *ResponseInsightRogueClient {
	this := ResponseInsightRogueClient{}
	this.End = end
	this.Limit = limit
	this.Next = next
	this.Results = results
	this.Start = start
	return &this
}

// NewResponseInsightRogueClientWithDefaults instantiates a new ResponseInsightRogueClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseInsightRogueClientWithDefaults() *ResponseInsightRogueClient {
	this := ResponseInsightRogueClient{}
	return &this
}

// GetEnd returns the End field value
func (o *ResponseInsightRogueClient) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ResponseInsightRogueClient) GetEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *ResponseInsightRogueClient) SetEnd(v int32) {
	o.End = v
}

// GetLimit returns the Limit field value
func (o *ResponseInsightRogueClient) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ResponseInsightRogueClient) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ResponseInsightRogueClient) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value
func (o *ResponseInsightRogueClient) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *ResponseInsightRogueClient) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *ResponseInsightRogueClient) SetNext(v string) {
	o.Next = v
}

// GetResults returns the Results field value
func (o *ResponseInsightRogueClient) GetResults() []InsightRogueClient {
	if o == nil {
		var ret []InsightRogueClient
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseInsightRogueClient) GetResultsOk() ([]InsightRogueClient, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseInsightRogueClient) SetResults(v []InsightRogueClient) {
	o.Results = v
}

// GetStart returns the Start field value
func (o *ResponseInsightRogueClient) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ResponseInsightRogueClient) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *ResponseInsightRogueClient) SetStart(v int32) {
	o.Start = v
}

func (o ResponseInsightRogueClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseInsightRogueClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["limit"] = o.Limit
	toSerialize["next"] = o.Next
	toSerialize["results"] = o.Results
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseInsightRogueClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end",
		"limit",
		"next",
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

	varResponseInsightRogueClient := _ResponseInsightRogueClient{}

	err = json.Unmarshal(data, &varResponseInsightRogueClient)

	if err != nil {
		return err
	}

	*o = ResponseInsightRogueClient(varResponseInsightRogueClient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseInsightRogueClient struct {
	value *ResponseInsightRogueClient
	isSet bool
}

func (v NullableResponseInsightRogueClient) Get() *ResponseInsightRogueClient {
	return v.value
}

func (v *NullableResponseInsightRogueClient) Set(val *ResponseInsightRogueClient) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseInsightRogueClient) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseInsightRogueClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseInsightRogueClient(val *ResponseInsightRogueClient) *NullableResponseInsightRogueClient {
	return &NullableResponseInsightRogueClient{value: val, isSet: true}
}

func (v NullableResponseInsightRogueClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseInsightRogueClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


