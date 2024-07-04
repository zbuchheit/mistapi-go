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

// checks if the ResponseSwitchMetricsActivePortsSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSwitchMetricsActivePortsSummary{}

// ResponseSwitchMetricsActivePortsSummary struct for ResponseSwitchMetricsActivePortsSummary
type ResponseSwitchMetricsActivePortsSummary struct {
	Details *SwitchMetricsActivePortsSummaryDetails `json:"details,omitempty"`
	Score *int32 `json:"score,omitempty"`
	TotalSwitchCount *int32 `json:"total_switch_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSwitchMetricsActivePortsSummary ResponseSwitchMetricsActivePortsSummary

// NewResponseSwitchMetricsActivePortsSummary instantiates a new ResponseSwitchMetricsActivePortsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSwitchMetricsActivePortsSummary() *ResponseSwitchMetricsActivePortsSummary {
	this := ResponseSwitchMetricsActivePortsSummary{}
	return &this
}

// NewResponseSwitchMetricsActivePortsSummaryWithDefaults instantiates a new ResponseSwitchMetricsActivePortsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSwitchMetricsActivePortsSummaryWithDefaults() *ResponseSwitchMetricsActivePortsSummary {
	this := ResponseSwitchMetricsActivePortsSummary{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ResponseSwitchMetricsActivePortsSummary) GetDetails() SwitchMetricsActivePortsSummaryDetails {
	if o == nil || IsNil(o.Details) {
		var ret SwitchMetricsActivePortsSummaryDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSwitchMetricsActivePortsSummary) GetDetailsOk() (*SwitchMetricsActivePortsSummaryDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ResponseSwitchMetricsActivePortsSummary) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given SwitchMetricsActivePortsSummaryDetails and assigns it to the Details field.
func (o *ResponseSwitchMetricsActivePortsSummary) SetDetails(v SwitchMetricsActivePortsSummaryDetails) {
	o.Details = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ResponseSwitchMetricsActivePortsSummary) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSwitchMetricsActivePortsSummary) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ResponseSwitchMetricsActivePortsSummary) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *ResponseSwitchMetricsActivePortsSummary) SetScore(v int32) {
	o.Score = &v
}

// GetTotalSwitchCount returns the TotalSwitchCount field value if set, zero value otherwise.
func (o *ResponseSwitchMetricsActivePortsSummary) GetTotalSwitchCount() int32 {
	if o == nil || IsNil(o.TotalSwitchCount) {
		var ret int32
		return ret
	}
	return *o.TotalSwitchCount
}

// GetTotalSwitchCountOk returns a tuple with the TotalSwitchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSwitchMetricsActivePortsSummary) GetTotalSwitchCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSwitchCount) {
		return nil, false
	}
	return o.TotalSwitchCount, true
}

// HasTotalSwitchCount returns a boolean if a field has been set.
func (o *ResponseSwitchMetricsActivePortsSummary) HasTotalSwitchCount() bool {
	if o != nil && !IsNil(o.TotalSwitchCount) {
		return true
	}

	return false
}

// SetTotalSwitchCount gets a reference to the given int32 and assigns it to the TotalSwitchCount field.
func (o *ResponseSwitchMetricsActivePortsSummary) SetTotalSwitchCount(v int32) {
	o.TotalSwitchCount = &v
}

func (o ResponseSwitchMetricsActivePortsSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSwitchMetricsActivePortsSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.TotalSwitchCount) {
		toSerialize["total_switch_count"] = o.TotalSwitchCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSwitchMetricsActivePortsSummary) UnmarshalJSON(data []byte) (err error) {
	varResponseSwitchMetricsActivePortsSummary := _ResponseSwitchMetricsActivePortsSummary{}

	err = json.Unmarshal(data, &varResponseSwitchMetricsActivePortsSummary)

	if err != nil {
		return err
	}

	*o = ResponseSwitchMetricsActivePortsSummary(varResponseSwitchMetricsActivePortsSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "details")
		delete(additionalProperties, "score")
		delete(additionalProperties, "total_switch_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSwitchMetricsActivePortsSummary struct {
	value *ResponseSwitchMetricsActivePortsSummary
	isSet bool
}

func (v NullableResponseSwitchMetricsActivePortsSummary) Get() *ResponseSwitchMetricsActivePortsSummary {
	return v.value
}

func (v *NullableResponseSwitchMetricsActivePortsSummary) Set(val *ResponseSwitchMetricsActivePortsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSwitchMetricsActivePortsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSwitchMetricsActivePortsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSwitchMetricsActivePortsSummary(val *ResponseSwitchMetricsActivePortsSummary) *NullableResponseSwitchMetricsActivePortsSummary {
	return &NullableResponseSwitchMetricsActivePortsSummary{value: val, isSet: true}
}

func (v NullableResponseSwitchMetricsActivePortsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSwitchMetricsActivePortsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


