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

// checks if the ResponseLoginFailure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseLoginFailure{}

// ResponseLoginFailure struct for ResponseLoginFailure
type ResponseLoginFailure struct {
	Detail string `json:"detail"`
	ForwardUrl *string `json:"forward_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseLoginFailure ResponseLoginFailure

// NewResponseLoginFailure instantiates a new ResponseLoginFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseLoginFailure(detail string) *ResponseLoginFailure {
	this := ResponseLoginFailure{}
	this.Detail = detail
	return &this
}

// NewResponseLoginFailureWithDefaults instantiates a new ResponseLoginFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseLoginFailureWithDefaults() *ResponseLoginFailure {
	this := ResponseLoginFailure{}
	return &this
}

// GetDetail returns the Detail field value
func (o *ResponseLoginFailure) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ResponseLoginFailure) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ResponseLoginFailure) SetDetail(v string) {
	o.Detail = v
}

// GetForwardUrl returns the ForwardUrl field value if set, zero value otherwise.
func (o *ResponseLoginFailure) GetForwardUrl() string {
	if o == nil || IsNil(o.ForwardUrl) {
		var ret string
		return ret
	}
	return *o.ForwardUrl
}

// GetForwardUrlOk returns a tuple with the ForwardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLoginFailure) GetForwardUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardUrl) {
		return nil, false
	}
	return o.ForwardUrl, true
}

// HasForwardUrl returns a boolean if a field has been set.
func (o *ResponseLoginFailure) HasForwardUrl() bool {
	if o != nil && !IsNil(o.ForwardUrl) {
		return true
	}

	return false
}

// SetForwardUrl gets a reference to the given string and assigns it to the ForwardUrl field.
func (o *ResponseLoginFailure) SetForwardUrl(v string) {
	o.ForwardUrl = &v
}

func (o ResponseLoginFailure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseLoginFailure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	if !IsNil(o.ForwardUrl) {
		toSerialize["forward_url"] = o.ForwardUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseLoginFailure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"detail",
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

	varResponseLoginFailure := _ResponseLoginFailure{}

	err = json.Unmarshal(data, &varResponseLoginFailure)

	if err != nil {
		return err
	}

	*o = ResponseLoginFailure(varResponseLoginFailure)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "forward_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseLoginFailure struct {
	value *ResponseLoginFailure
	isSet bool
}

func (v NullableResponseLoginFailure) Get() *ResponseLoginFailure {
	return v.value
}

func (v *NullableResponseLoginFailure) Set(val *ResponseLoginFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseLoginFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseLoginFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseLoginFailure(val *ResponseLoginFailure) *NullableResponseLoginFailure {
	return &NullableResponseLoginFailure{value: val, isSet: true}
}

func (v NullableResponseLoginFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseLoginFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


