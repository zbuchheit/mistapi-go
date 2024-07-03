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

// checks if the SnmpConfigViews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpConfigViews{}

// SnmpConfigViews struct for SnmpConfigViews
type SnmpConfigViews struct {
	// if the root oid configured is included
	Include *bool `json:"include,omitempty"`
	Oid *string `json:"oid,omitempty"`
	ViewName *string `json:"view_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpConfigViews SnmpConfigViews

// NewSnmpConfigViews instantiates a new SnmpConfigViews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpConfigViews() *SnmpConfigViews {
	this := SnmpConfigViews{}
	return &this
}

// NewSnmpConfigViewsWithDefaults instantiates a new SnmpConfigViews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpConfigViewsWithDefaults() *SnmpConfigViews {
	this := SnmpConfigViews{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *SnmpConfigViews) GetInclude() bool {
	if o == nil || IsNil(o.Include) {
		var ret bool
		return ret
	}
	return *o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigViews) GetIncludeOk() (*bool, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *SnmpConfigViews) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given bool and assigns it to the Include field.
func (o *SnmpConfigViews) SetInclude(v bool) {
	o.Include = &v
}

// GetOid returns the Oid field value if set, zero value otherwise.
func (o *SnmpConfigViews) GetOid() string {
	if o == nil || IsNil(o.Oid) {
		var ret string
		return ret
	}
	return *o.Oid
}

// GetOidOk returns a tuple with the Oid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigViews) GetOidOk() (*string, bool) {
	if o == nil || IsNil(o.Oid) {
		return nil, false
	}
	return o.Oid, true
}

// HasOid returns a boolean if a field has been set.
func (o *SnmpConfigViews) HasOid() bool {
	if o != nil && !IsNil(o.Oid) {
		return true
	}

	return false
}

// SetOid gets a reference to the given string and assigns it to the Oid field.
func (o *SnmpConfigViews) SetOid(v string) {
	o.Oid = &v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *SnmpConfigViews) GetViewName() string {
	if o == nil || IsNil(o.ViewName) {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigViews) GetViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.ViewName) {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *SnmpConfigViews) HasViewName() bool {
	if o != nil && !IsNil(o.ViewName) {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *SnmpConfigViews) SetViewName(v string) {
	o.ViewName = &v
}

func (o SnmpConfigViews) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpConfigViews) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	if !IsNil(o.Oid) {
		toSerialize["oid"] = o.Oid
	}
	if !IsNil(o.ViewName) {
		toSerialize["view_name"] = o.ViewName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SnmpConfigViews) UnmarshalJSON(data []byte) (err error) {
	varSnmpConfigViews := _SnmpConfigViews{}

	err = json.Unmarshal(data, &varSnmpConfigViews)

	if err != nil {
		return err
	}

	*o = SnmpConfigViews(varSnmpConfigViews)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		delete(additionalProperties, "oid")
		delete(additionalProperties, "view_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpConfigViews struct {
	value *SnmpConfigViews
	isSet bool
}

func (v NullableSnmpConfigViews) Get() *SnmpConfigViews {
	return v.value
}

func (v *NullableSnmpConfigViews) Set(val *SnmpConfigViews) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpConfigViews) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpConfigViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpConfigViews(val *SnmpConfigViews) *NullableSnmpConfigViews {
	return &NullableSnmpConfigViews{value: val, isSet: true}
}

func (v NullableSnmpConfigViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpConfigViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


