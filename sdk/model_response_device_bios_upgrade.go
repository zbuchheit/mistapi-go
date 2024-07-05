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
)

// checks if the ResponseDeviceBiosUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeviceBiosUpgrade{}

// ResponseDeviceBiosUpgrade struct for ResponseDeviceBiosUpgrade
type ResponseDeviceBiosUpgrade struct {
	Status *string `json:"status,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDeviceBiosUpgrade ResponseDeviceBiosUpgrade

// NewResponseDeviceBiosUpgrade instantiates a new ResponseDeviceBiosUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeviceBiosUpgrade() *ResponseDeviceBiosUpgrade {
	this := ResponseDeviceBiosUpgrade{}
	return &this
}

// NewResponseDeviceBiosUpgradeWithDefaults instantiates a new ResponseDeviceBiosUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeviceBiosUpgradeWithDefaults() *ResponseDeviceBiosUpgrade {
	this := ResponseDeviceBiosUpgrade{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseDeviceBiosUpgrade) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceBiosUpgrade) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseDeviceBiosUpgrade) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResponseDeviceBiosUpgrade) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ResponseDeviceBiosUpgrade) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceBiosUpgrade) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ResponseDeviceBiosUpgrade) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *ResponseDeviceBiosUpgrade) SetTimestamp(v int32) {
	o.Timestamp = &v
}

func (o ResponseDeviceBiosUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeviceBiosUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDeviceBiosUpgrade) UnmarshalJSON(data []byte) (err error) {
	varResponseDeviceBiosUpgrade := _ResponseDeviceBiosUpgrade{}

	err = json.Unmarshal(data, &varResponseDeviceBiosUpgrade)

	if err != nil {
		return err
	}

	*o = ResponseDeviceBiosUpgrade(varResponseDeviceBiosUpgrade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDeviceBiosUpgrade struct {
	value *ResponseDeviceBiosUpgrade
	isSet bool
}

func (v NullableResponseDeviceBiosUpgrade) Get() *ResponseDeviceBiosUpgrade {
	return v.value
}

func (v *NullableResponseDeviceBiosUpgrade) Set(val *ResponseDeviceBiosUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeviceBiosUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeviceBiosUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeviceBiosUpgrade(val *ResponseDeviceBiosUpgrade) *NullableResponseDeviceBiosUpgrade {
	return &NullableResponseDeviceBiosUpgrade{value: val, isSet: true}
}

func (v NullableResponseDeviceBiosUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeviceBiosUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


