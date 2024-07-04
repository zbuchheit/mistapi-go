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
	"fmt"
)

// checks if the ResponseUpgradeDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseUpgradeDevice{}

// ResponseUpgradeDevice struct for ResponseUpgradeDevice
type ResponseUpgradeDevice struct {
	Status UpgradeInfoStatus `json:"status"`
	// timestamp
	Timestamp float32 `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _ResponseUpgradeDevice ResponseUpgradeDevice

// NewResponseUpgradeDevice instantiates a new ResponseUpgradeDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseUpgradeDevice(status UpgradeInfoStatus, timestamp float32) *ResponseUpgradeDevice {
	this := ResponseUpgradeDevice{}
	this.Status = status
	this.Timestamp = timestamp
	return &this
}

// NewResponseUpgradeDeviceWithDefaults instantiates a new ResponseUpgradeDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseUpgradeDeviceWithDefaults() *ResponseUpgradeDevice {
	this := ResponseUpgradeDevice{}
	return &this
}

// GetStatus returns the Status field value
func (o *ResponseUpgradeDevice) GetStatus() UpgradeInfoStatus {
	if o == nil {
		var ret UpgradeInfoStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeDevice) GetStatusOk() (*UpgradeInfoStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResponseUpgradeDevice) SetStatus(v UpgradeInfoStatus) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *ResponseUpgradeDevice) GetTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ResponseUpgradeDevice) GetTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ResponseUpgradeDevice) SetTimestamp(v float32) {
	o.Timestamp = v
}

func (o ResponseUpgradeDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseUpgradeDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseUpgradeDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"timestamp",
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

	varResponseUpgradeDevice := _ResponseUpgradeDevice{}

	err = json.Unmarshal(data, &varResponseUpgradeDevice)

	if err != nil {
		return err
	}

	*o = ResponseUpgradeDevice(varResponseUpgradeDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseUpgradeDevice struct {
	value *ResponseUpgradeDevice
	isSet bool
}

func (v NullableResponseUpgradeDevice) Get() *ResponseUpgradeDevice {
	return v.value
}

func (v *NullableResponseUpgradeDevice) Set(val *ResponseUpgradeDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseUpgradeDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseUpgradeDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseUpgradeDevice(val *ResponseUpgradeDevice) *NullableResponseUpgradeDevice {
	return &NullableResponseUpgradeDevice{value: val, isSet: true}
}

func (v NullableResponseUpgradeDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseUpgradeDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


