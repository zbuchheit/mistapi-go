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

// checks if the ResponseSelfOauthLinkFailure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSelfOauthLinkFailure{}

// ResponseSelfOauthLinkFailure struct for ResponseSelfOauthLinkFailure
type ResponseSelfOauthLinkFailure struct {
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSelfOauthLinkFailure ResponseSelfOauthLinkFailure

// NewResponseSelfOauthLinkFailure instantiates a new ResponseSelfOauthLinkFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSelfOauthLinkFailure(error_ string, errorDescription string) *ResponseSelfOauthLinkFailure {
	this := ResponseSelfOauthLinkFailure{}
	this.Error = error_
	this.ErrorDescription = errorDescription
	return &this
}

// NewResponseSelfOauthLinkFailureWithDefaults instantiates a new ResponseSelfOauthLinkFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSelfOauthLinkFailureWithDefaults() *ResponseSelfOauthLinkFailure {
	this := ResponseSelfOauthLinkFailure{}
	return &this
}

// GetError returns the Error field value
func (o *ResponseSelfOauthLinkFailure) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ResponseSelfOauthLinkFailure) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ResponseSelfOauthLinkFailure) SetError(v string) {
	o.Error = v
}

// GetErrorDescription returns the ErrorDescription field value
func (o *ResponseSelfOauthLinkFailure) GetErrorDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value
// and a boolean to check if the value has been set.
func (o *ResponseSelfOauthLinkFailure) GetErrorDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorDescription, true
}

// SetErrorDescription sets field value
func (o *ResponseSelfOauthLinkFailure) SetErrorDescription(v string) {
	o.ErrorDescription = v
}

func (o ResponseSelfOauthLinkFailure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSelfOauthLinkFailure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["error_description"] = o.ErrorDescription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSelfOauthLinkFailure) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
		"error_description",
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

	varResponseSelfOauthLinkFailure := _ResponseSelfOauthLinkFailure{}

	err = json.Unmarshal(data, &varResponseSelfOauthLinkFailure)

	if err != nil {
		return err
	}

	*o = ResponseSelfOauthLinkFailure(varResponseSelfOauthLinkFailure)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "error_description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSelfOauthLinkFailure struct {
	value *ResponseSelfOauthLinkFailure
	isSet bool
}

func (v NullableResponseSelfOauthLinkFailure) Get() *ResponseSelfOauthLinkFailure {
	return v.value
}

func (v *NullableResponseSelfOauthLinkFailure) Set(val *ResponseSelfOauthLinkFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSelfOauthLinkFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSelfOauthLinkFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSelfOauthLinkFailure(val *ResponseSelfOauthLinkFailure) *NullableResponseSelfOauthLinkFailure {
	return &NullableResponseSelfOauthLinkFailure{value: val, isSet: true}
}

func (v NullableResponseSelfOauthLinkFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSelfOauthLinkFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


