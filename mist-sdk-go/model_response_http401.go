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
)

// checks if the ResponseHttp401 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseHttp401{}

// ResponseHttp401 struct for ResponseHttp401
type ResponseHttp401 struct {
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseHttp401 ResponseHttp401

// NewResponseHttp401 instantiates a new ResponseHttp401 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseHttp401() *ResponseHttp401 {
	this := ResponseHttp401{}
	return &this
}

// NewResponseHttp401WithDefaults instantiates a new ResponseHttp401 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseHttp401WithDefaults() *ResponseHttp401 {
	this := ResponseHttp401{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseHttp401) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseHttp401) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseHttp401) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseHttp401) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseHttp401) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseHttp401) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseHttp401) UnmarshalJSON(data []byte) (err error) {
	varResponseHttp401 := _ResponseHttp401{}

	err = json.Unmarshal(data, &varResponseHttp401)

	if err != nil {
		return err
	}

	*o = ResponseHttp401(varResponseHttp401)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseHttp401 struct {
	value *ResponseHttp401
	isSet bool
}

func (v NullableResponseHttp401) Get() *ResponseHttp401 {
	return v.value
}

func (v *NullableResponseHttp401) Set(val *ResponseHttp401) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseHttp401) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseHttp401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseHttp401(val *ResponseHttp401) *NullableResponseHttp401 {
	return &NullableResponseHttp401{value: val, isSet: true}
}

func (v NullableResponseHttp401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseHttp401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


