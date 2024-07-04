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

// checks if the ResponseMxedgeEventsSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMxedgeEventsSearch{}

// ResponseMxedgeEventsSearch struct for ResponseMxedgeEventsSearch
type ResponseMxedgeEventsSearch struct {
	End *int32 `json:"end,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Page *int32 `json:"page,omitempty"`
	Results []MxedgeEvent `json:"results,omitempty"`
	Start *int32 `json:"start,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseMxedgeEventsSearch ResponseMxedgeEventsSearch

// NewResponseMxedgeEventsSearch instantiates a new ResponseMxedgeEventsSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMxedgeEventsSearch() *ResponseMxedgeEventsSearch {
	this := ResponseMxedgeEventsSearch{}
	return &this
}

// NewResponseMxedgeEventsSearchWithDefaults instantiates a new ResponseMxedgeEventsSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMxedgeEventsSearchWithDefaults() *ResponseMxedgeEventsSearch {
	this := ResponseMxedgeEventsSearch{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ResponseMxedgeEventsSearch) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeEventsSearch) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ResponseMxedgeEventsSearch) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ResponseMxedgeEventsSearch) SetEnd(v int32) {
	o.End = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseMxedgeEventsSearch) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeEventsSearch) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseMxedgeEventsSearch) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ResponseMxedgeEventsSearch) SetLimit(v int32) {
	o.Limit = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ResponseMxedgeEventsSearch) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeEventsSearch) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ResponseMxedgeEventsSearch) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ResponseMxedgeEventsSearch) SetPage(v int32) {
	o.Page = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ResponseMxedgeEventsSearch) GetResults() []MxedgeEvent {
	if o == nil || IsNil(o.Results) {
		var ret []MxedgeEvent
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeEventsSearch) GetResultsOk() ([]MxedgeEvent, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ResponseMxedgeEventsSearch) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MxedgeEvent and assigns it to the Results field.
func (o *ResponseMxedgeEventsSearch) SetResults(v []MxedgeEvent) {
	o.Results = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ResponseMxedgeEventsSearch) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeEventsSearch) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ResponseMxedgeEventsSearch) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ResponseMxedgeEventsSearch) SetStart(v int32) {
	o.Start = &v
}

func (o ResponseMxedgeEventsSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMxedgeEventsSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseMxedgeEventsSearch) UnmarshalJSON(data []byte) (err error) {
	varResponseMxedgeEventsSearch := _ResponseMxedgeEventsSearch{}

	err = json.Unmarshal(data, &varResponseMxedgeEventsSearch)

	if err != nil {
		return err
	}

	*o = ResponseMxedgeEventsSearch(varResponseMxedgeEventsSearch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "page")
		delete(additionalProperties, "results")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseMxedgeEventsSearch struct {
	value *ResponseMxedgeEventsSearch
	isSet bool
}

func (v NullableResponseMxedgeEventsSearch) Get() *ResponseMxedgeEventsSearch {
	return v.value
}

func (v *NullableResponseMxedgeEventsSearch) Set(val *ResponseMxedgeEventsSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMxedgeEventsSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMxedgeEventsSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMxedgeEventsSearch(val *ResponseMxedgeEventsSearch) *NullableResponseMxedgeEventsSearch {
	return &NullableResponseMxedgeEventsSearch{value: val, isSet: true}
}

func (v NullableResponseMxedgeEventsSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMxedgeEventsSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


