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

// checks if the Proxy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Proxy{}

// Proxy Proxy Configuration to talk to Mist
type Proxy struct {
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Proxy Proxy

// NewProxy instantiates a new Proxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxy() *Proxy {
	this := Proxy{}
	return &this
}

// NewProxyWithDefaults instantiates a new Proxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyWithDefaults() *Proxy {
	this := Proxy{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Proxy) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Proxy) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Proxy) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Proxy) SetUrl(v string) {
	o.Url = &v
}

func (o Proxy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Proxy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Proxy) UnmarshalJSON(data []byte) (err error) {
	varProxy := _Proxy{}

	err = json.Unmarshal(data, &varProxy)

	if err != nil {
		return err
	}

	*o = Proxy(varProxy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProxy struct {
	value *Proxy
	isSet bool
}

func (v NullableProxy) Get() *Proxy {
	return v.value
}

func (v *NullableProxy) Set(val *Proxy) {
	v.value = val
	v.isSet = true
}

func (v NullableProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxy(val *Proxy) *NullableProxy {
	return &NullableProxy{value: val, isSet: true}
}

func (v NullableProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


