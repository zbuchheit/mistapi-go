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

// checks if the ResponseEventsOrgsSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEventsOrgsSearch{}

// ResponseEventsOrgsSearch struct for ResponseEventsOrgsSearch
type ResponseEventsOrgsSearch struct {
	End *int32 `json:"end,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Results []OrgEvent `json:"results,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Total *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseEventsOrgsSearch ResponseEventsOrgsSearch

// NewResponseEventsOrgsSearch instantiates a new ResponseEventsOrgsSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEventsOrgsSearch() *ResponseEventsOrgsSearch {
	this := ResponseEventsOrgsSearch{}
	return &this
}

// NewResponseEventsOrgsSearchWithDefaults instantiates a new ResponseEventsOrgsSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEventsOrgsSearchWithDefaults() *ResponseEventsOrgsSearch {
	this := ResponseEventsOrgsSearch{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ResponseEventsOrgsSearch) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEventsOrgsSearch) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ResponseEventsOrgsSearch) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ResponseEventsOrgsSearch) SetEnd(v int32) {
	o.End = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseEventsOrgsSearch) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEventsOrgsSearch) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseEventsOrgsSearch) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ResponseEventsOrgsSearch) SetLimit(v int32) {
	o.Limit = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ResponseEventsOrgsSearch) GetResults() []OrgEvent {
	if o == nil || IsNil(o.Results) {
		var ret []OrgEvent
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEventsOrgsSearch) GetResultsOk() ([]OrgEvent, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ResponseEventsOrgsSearch) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OrgEvent and assigns it to the Results field.
func (o *ResponseEventsOrgsSearch) SetResults(v []OrgEvent) {
	o.Results = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ResponseEventsOrgsSearch) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEventsOrgsSearch) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ResponseEventsOrgsSearch) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ResponseEventsOrgsSearch) SetStart(v int32) {
	o.Start = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponseEventsOrgsSearch) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseEventsOrgsSearch) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponseEventsOrgsSearch) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ResponseEventsOrgsSearch) SetTotal(v int32) {
	o.Total = &v
}

func (o ResponseEventsOrgsSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEventsOrgsSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseEventsOrgsSearch) UnmarshalJSON(data []byte) (err error) {
	varResponseEventsOrgsSearch := _ResponseEventsOrgsSearch{}

	err = json.Unmarshal(data, &varResponseEventsOrgsSearch)

	if err != nil {
		return err
	}

	*o = ResponseEventsOrgsSearch(varResponseEventsOrgsSearch)

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

type NullableResponseEventsOrgsSearch struct {
	value *ResponseEventsOrgsSearch
	isSet bool
}

func (v NullableResponseEventsOrgsSearch) Get() *ResponseEventsOrgsSearch {
	return v.value
}

func (v *NullableResponseEventsOrgsSearch) Set(val *ResponseEventsOrgsSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEventsOrgsSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEventsOrgsSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEventsOrgsSearch(val *ResponseEventsOrgsSearch) *NullableResponseEventsOrgsSearch {
	return &NullableResponseEventsOrgsSearch{value: val, isSet: true}
}

func (v NullableResponseEventsOrgsSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEventsOrgsSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


