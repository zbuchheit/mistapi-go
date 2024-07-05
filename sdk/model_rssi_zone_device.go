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

// checks if the RssiZoneDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RssiZoneDevice{}

// RssiZoneDevice struct for RssiZoneDevice
type RssiZoneDevice struct {
	DeviceId string `json:"device_id"`
	// RSSI threshold
	Rssi int32 `json:"rssi"`
	AdditionalProperties map[string]interface{}
}

type _RssiZoneDevice RssiZoneDevice

// NewRssiZoneDevice instantiates a new RssiZoneDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRssiZoneDevice(deviceId string, rssi int32) *RssiZoneDevice {
	this := RssiZoneDevice{}
	this.DeviceId = deviceId
	this.Rssi = rssi
	return &this
}

// NewRssiZoneDeviceWithDefaults instantiates a new RssiZoneDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRssiZoneDeviceWithDefaults() *RssiZoneDevice {
	this := RssiZoneDevice{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *RssiZoneDevice) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *RssiZoneDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *RssiZoneDevice) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetRssi returns the Rssi field value
func (o *RssiZoneDevice) GetRssi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value
// and a boolean to check if the value has been set.
func (o *RssiZoneDevice) GetRssiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rssi, true
}

// SetRssi sets field value
func (o *RssiZoneDevice) SetRssi(v int32) {
	o.Rssi = v
}

func (o RssiZoneDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RssiZoneDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device_id"] = o.DeviceId
	toSerialize["rssi"] = o.Rssi

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RssiZoneDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device_id",
		"rssi",
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

	varRssiZoneDevice := _RssiZoneDevice{}

	err = json.Unmarshal(data, &varRssiZoneDevice)

	if err != nil {
		return err
	}

	*o = RssiZoneDevice(varRssiZoneDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_id")
		delete(additionalProperties, "rssi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRssiZoneDevice struct {
	value *RssiZoneDevice
	isSet bool
}

func (v NullableRssiZoneDevice) Get() *RssiZoneDevice {
	return v.value
}

func (v *NullableRssiZoneDevice) Set(val *RssiZoneDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableRssiZoneDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableRssiZoneDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRssiZoneDevice(val *RssiZoneDevice) *NullableRssiZoneDevice {
	return &NullableRssiZoneDevice{value: val, isSet: true}
}

func (v NullableRssiZoneDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRssiZoneDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


