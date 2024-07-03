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

// checks if the ResponseLoginLookup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseLoginLookup{}

// ResponseLoginLookup struct for ResponseLoginLookup
type ResponseLoginLookup struct {
	SsoUrl *string `json:"sso_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseLoginLookup ResponseLoginLookup

// NewResponseLoginLookup instantiates a new ResponseLoginLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseLoginLookup() *ResponseLoginLookup {
	this := ResponseLoginLookup{}
	return &this
}

// NewResponseLoginLookupWithDefaults instantiates a new ResponseLoginLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseLoginLookupWithDefaults() *ResponseLoginLookup {
	this := ResponseLoginLookup{}
	return &this
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise.
func (o *ResponseLoginLookup) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLoginLookup) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}
	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *ResponseLoginLookup) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *ResponseLoginLookup) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

func (o ResponseLoginLookup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseLoginLookup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsoUrl) {
		toSerialize["sso_url"] = o.SsoUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseLoginLookup) UnmarshalJSON(data []byte) (err error) {
	varResponseLoginLookup := _ResponseLoginLookup{}

	err = json.Unmarshal(data, &varResponseLoginLookup)

	if err != nil {
		return err
	}

	*o = ResponseLoginLookup(varResponseLoginLookup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sso_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseLoginLookup struct {
	value *ResponseLoginLookup
	isSet bool
}

func (v NullableResponseLoginLookup) Get() *ResponseLoginLookup {
	return v.value
}

func (v *NullableResponseLoginLookup) Set(val *ResponseLoginLookup) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseLoginLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseLoginLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseLoginLookup(val *ResponseLoginLookup) *NullableResponseLoginLookup {
	return &NullableResponseLoginLookup{value: val, isSet: true}
}

func (v NullableResponseLoginLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseLoginLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


