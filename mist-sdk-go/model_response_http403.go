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

// checks if the ResponseHttp403 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseHttp403{}

// ResponseHttp403 struct for ResponseHttp403
type ResponseHttp403 struct {
	Detail *string `json:"detail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseHttp403 ResponseHttp403

// NewResponseHttp403 instantiates a new ResponseHttp403 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseHttp403() *ResponseHttp403 {
	this := ResponseHttp403{}
	return &this
}

// NewResponseHttp403WithDefaults instantiates a new ResponseHttp403 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseHttp403WithDefaults() *ResponseHttp403 {
	this := ResponseHttp403{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseHttp403) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseHttp403) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseHttp403) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseHttp403) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseHttp403) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseHttp403) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseHttp403) UnmarshalJSON(data []byte) (err error) {
	varResponseHttp403 := _ResponseHttp403{}

	err = json.Unmarshal(data, &varResponseHttp403)

	if err != nil {
		return err
	}

	*o = ResponseHttp403(varResponseHttp403)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseHttp403 struct {
	value *ResponseHttp403
	isSet bool
}

func (v NullableResponseHttp403) Get() *ResponseHttp403 {
	return v.value
}

func (v *NullableResponseHttp403) Set(val *ResponseHttp403) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseHttp403) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseHttp403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseHttp403(val *ResponseHttp403) *NullableResponseHttp403 {
	return &NullableResponseHttp403{value: val, isSet: true}
}

func (v NullableResponseHttp403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseHttp403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


