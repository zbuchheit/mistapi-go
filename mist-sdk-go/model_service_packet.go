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

// checks if the ServicePacket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePacket{}

// ServicePacket struct for ServicePacket
type ServicePacket struct {
	// ata from service data
	ServiceData *string `json:"service_data,omitempty"`
	// UUID from service data
	ServiceUuid *string `json:"service_uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicePacket ServicePacket

// NewServicePacket instantiates a new ServicePacket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePacket() *ServicePacket {
	this := ServicePacket{}
	return &this
}

// NewServicePacketWithDefaults instantiates a new ServicePacket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePacketWithDefaults() *ServicePacket {
	this := ServicePacket{}
	return &this
}

// GetServiceData returns the ServiceData field value if set, zero value otherwise.
func (o *ServicePacket) GetServiceData() string {
	if o == nil || IsNil(o.ServiceData) {
		var ret string
		return ret
	}
	return *o.ServiceData
}

// GetServiceDataOk returns a tuple with the ServiceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePacket) GetServiceDataOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceData) {
		return nil, false
	}
	return o.ServiceData, true
}

// HasServiceData returns a boolean if a field has been set.
func (o *ServicePacket) HasServiceData() bool {
	if o != nil && !IsNil(o.ServiceData) {
		return true
	}

	return false
}

// SetServiceData gets a reference to the given string and assigns it to the ServiceData field.
func (o *ServicePacket) SetServiceData(v string) {
	o.ServiceData = &v
}

// GetServiceUuid returns the ServiceUuid field value if set, zero value otherwise.
func (o *ServicePacket) GetServiceUuid() string {
	if o == nil || IsNil(o.ServiceUuid) {
		var ret string
		return ret
	}
	return *o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePacket) GetServiceUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUuid) {
		return nil, false
	}
	return o.ServiceUuid, true
}

// HasServiceUuid returns a boolean if a field has been set.
func (o *ServicePacket) HasServiceUuid() bool {
	if o != nil && !IsNil(o.ServiceUuid) {
		return true
	}

	return false
}

// SetServiceUuid gets a reference to the given string and assigns it to the ServiceUuid field.
func (o *ServicePacket) SetServiceUuid(v string) {
	o.ServiceUuid = &v
}

func (o ServicePacket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePacket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceData) {
		toSerialize["service_data"] = o.ServiceData
	}
	if !IsNil(o.ServiceUuid) {
		toSerialize["service_uuid"] = o.ServiceUuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicePacket) UnmarshalJSON(data []byte) (err error) {
	varServicePacket := _ServicePacket{}

	err = json.Unmarshal(data, &varServicePacket)

	if err != nil {
		return err
	}

	*o = ServicePacket(varServicePacket)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "service_data")
		delete(additionalProperties, "service_uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicePacket struct {
	value *ServicePacket
	isSet bool
}

func (v NullableServicePacket) Get() *ServicePacket {
	return v.value
}

func (v *NullableServicePacket) Set(val *ServicePacket) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePacket) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePacket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePacket(val *ServicePacket) *NullableServicePacket {
	return &NullableServicePacket{value: val, isSet: true}
}

func (v NullableServicePacket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePacket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


