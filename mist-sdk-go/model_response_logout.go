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

// checks if the ResponseLogout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseLogout{}

// ResponseLogout struct for ResponseLogout
type ResponseLogout struct {
	// if configured in SSO as custom_logout_url
	ForwardUrl *string `json:"forward_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseLogout ResponseLogout

// NewResponseLogout instantiates a new ResponseLogout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseLogout() *ResponseLogout {
	this := ResponseLogout{}
	return &this
}

// NewResponseLogoutWithDefaults instantiates a new ResponseLogout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseLogoutWithDefaults() *ResponseLogout {
	this := ResponseLogout{}
	return &this
}

// GetForwardUrl returns the ForwardUrl field value if set, zero value otherwise.
func (o *ResponseLogout) GetForwardUrl() string {
	if o == nil || IsNil(o.ForwardUrl) {
		var ret string
		return ret
	}
	return *o.ForwardUrl
}

// GetForwardUrlOk returns a tuple with the ForwardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseLogout) GetForwardUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardUrl) {
		return nil, false
	}
	return o.ForwardUrl, true
}

// HasForwardUrl returns a boolean if a field has been set.
func (o *ResponseLogout) HasForwardUrl() bool {
	if o != nil && !IsNil(o.ForwardUrl) {
		return true
	}

	return false
}

// SetForwardUrl gets a reference to the given string and assigns it to the ForwardUrl field.
func (o *ResponseLogout) SetForwardUrl(v string) {
	o.ForwardUrl = &v
}

func (o ResponseLogout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseLogout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForwardUrl) {
		toSerialize["forward_url"] = o.ForwardUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseLogout) UnmarshalJSON(data []byte) (err error) {
	varResponseLogout := _ResponseLogout{}

	err = json.Unmarshal(data, &varResponseLogout)

	if err != nil {
		return err
	}

	*o = ResponseLogout(varResponseLogout)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "forward_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseLogout struct {
	value *ResponseLogout
	isSet bool
}

func (v NullableResponseLogout) Get() *ResponseLogout {
	return v.value
}

func (v *NullableResponseLogout) Set(val *ResponseLogout) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseLogout) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseLogout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseLogout(val *ResponseLogout) *NullableResponseLogout {
	return &NullableResponseLogout{value: val, isSet: true}
}

func (v NullableResponseLogout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseLogout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


