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

// checks if the LatlngBr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LatlngBr{}

// LatlngBr when type=google, latitude / longitude of the bottom-right corner
type LatlngBr struct {
	Lat *string `json:"lat,omitempty"`
	Lng *string `json:"lng,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LatlngBr LatlngBr

// NewLatlngBr instantiates a new LatlngBr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLatlngBr() *LatlngBr {
	this := LatlngBr{}
	return &this
}

// NewLatlngBrWithDefaults instantiates a new LatlngBr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLatlngBrWithDefaults() *LatlngBr {
	this := LatlngBr{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *LatlngBr) GetLat() string {
	if o == nil || IsNil(o.Lat) {
		var ret string
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatlngBr) GetLatOk() (*string, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *LatlngBr) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given string and assigns it to the Lat field.
func (o *LatlngBr) SetLat(v string) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *LatlngBr) GetLng() string {
	if o == nil || IsNil(o.Lng) {
		var ret string
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatlngBr) GetLngOk() (*string, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *LatlngBr) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given string and assigns it to the Lng field.
func (o *LatlngBr) SetLng(v string) {
	o.Lng = &v
}

func (o LatlngBr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LatlngBr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LatlngBr) UnmarshalJSON(data []byte) (err error) {
	varLatlngBr := _LatlngBr{}

	err = json.Unmarshal(data, &varLatlngBr)

	if err != nil {
		return err
	}

	*o = LatlngBr(varLatlngBr)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lat")
		delete(additionalProperties, "lng")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLatlngBr struct {
	value *LatlngBr
	isSet bool
}

func (v NullableLatlngBr) Get() *LatlngBr {
	return v.value
}

func (v *NullableLatlngBr) Set(val *LatlngBr) {
	v.value = val
	v.isSet = true
}

func (v NullableLatlngBr) IsSet() bool {
	return v.isSet
}

func (v *NullableLatlngBr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLatlngBr(val *LatlngBr) *NullableLatlngBr {
	return &NullableLatlngBr{value: val, isSet: true}
}

func (v NullableLatlngBr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLatlngBr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


