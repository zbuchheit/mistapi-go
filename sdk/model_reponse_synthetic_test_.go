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

// checks if the ReponseSyntheticTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReponseSyntheticTest{}

// ReponseSyntheticTest struct for ReponseSyntheticTest
type ReponseSyntheticTest struct {
	Id *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReponseSyntheticTest ReponseSyntheticTest

// NewReponseSyntheticTest instantiates a new ReponseSyntheticTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReponseSyntheticTest() *ReponseSyntheticTest {
	this := ReponseSyntheticTest{}
	return &this
}

// NewReponseSyntheticTestWithDefaults instantiates a new ReponseSyntheticTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReponseSyntheticTestWithDefaults() *ReponseSyntheticTest {
	this := ReponseSyntheticTest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReponseSyntheticTest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseSyntheticTest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReponseSyntheticTest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReponseSyntheticTest) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ReponseSyntheticTest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseSyntheticTest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ReponseSyntheticTest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ReponseSyntheticTest) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReponseSyntheticTest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReponseSyntheticTest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReponseSyntheticTest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReponseSyntheticTest) SetStatus(v string) {
	o.Status = &v
}

func (o ReponseSyntheticTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReponseSyntheticTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReponseSyntheticTest) UnmarshalJSON(data []byte) (err error) {
	varReponseSyntheticTest := _ReponseSyntheticTest{}

	err = json.Unmarshal(data, &varReponseSyntheticTest)

	if err != nil {
		return err
	}

	*o = ReponseSyntheticTest(varReponseSyntheticTest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "message")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReponseSyntheticTest struct {
	value *ReponseSyntheticTest
	isSet bool
}

func (v NullableReponseSyntheticTest) Get() *ReponseSyntheticTest {
	return v.value
}

func (v *NullableReponseSyntheticTest) Set(val *ReponseSyntheticTest) {
	v.value = val
	v.isSet = true
}

func (v NullableReponseSyntheticTest) IsSet() bool {
	return v.isSet
}

func (v *NullableReponseSyntheticTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReponseSyntheticTest(val *ReponseSyntheticTest) *NullableReponseSyntheticTest {
	return &NullableReponseSyntheticTest{value: val, isSet: true}
}

func (v NullableReponseSyntheticTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReponseSyntheticTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


