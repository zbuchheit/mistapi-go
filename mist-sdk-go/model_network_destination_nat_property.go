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

// checks if the NetworkDestinationNatProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDestinationNatProperty{}

// NetworkDestinationNatProperty struct for NetworkDestinationNatProperty
type NetworkDestinationNatProperty struct {
	InternalIp *string `json:"internal_ip,omitempty"`
	Name *string `json:"name,omitempty"`
	Port *int32 `json:"port,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDestinationNatProperty NetworkDestinationNatProperty

// NewNetworkDestinationNatProperty instantiates a new NetworkDestinationNatProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDestinationNatProperty() *NetworkDestinationNatProperty {
	this := NetworkDestinationNatProperty{}
	return &this
}

// NewNetworkDestinationNatPropertyWithDefaults instantiates a new NetworkDestinationNatProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDestinationNatPropertyWithDefaults() *NetworkDestinationNatProperty {
	this := NetworkDestinationNatProperty{}
	return &this
}

// GetInternalIp returns the InternalIp field value if set, zero value otherwise.
func (o *NetworkDestinationNatProperty) GetInternalIp() string {
	if o == nil || IsNil(o.InternalIp) {
		var ret string
		return ret
	}
	return *o.InternalIp
}

// GetInternalIpOk returns a tuple with the InternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDestinationNatProperty) GetInternalIpOk() (*string, bool) {
	if o == nil || IsNil(o.InternalIp) {
		return nil, false
	}
	return o.InternalIp, true
}

// HasInternalIp returns a boolean if a field has been set.
func (o *NetworkDestinationNatProperty) HasInternalIp() bool {
	if o != nil && !IsNil(o.InternalIp) {
		return true
	}

	return false
}

// SetInternalIp gets a reference to the given string and assigns it to the InternalIp field.
func (o *NetworkDestinationNatProperty) SetInternalIp(v string) {
	o.InternalIp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkDestinationNatProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDestinationNatProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkDestinationNatProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkDestinationNatProperty) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NetworkDestinationNatProperty) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDestinationNatProperty) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NetworkDestinationNatProperty) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *NetworkDestinationNatProperty) SetPort(v int32) {
	o.Port = &v
}

func (o NetworkDestinationNatProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDestinationNatProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InternalIp) {
		toSerialize["internal_ip"] = o.InternalIp
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDestinationNatProperty) UnmarshalJSON(data []byte) (err error) {
	varNetworkDestinationNatProperty := _NetworkDestinationNatProperty{}

	err = json.Unmarshal(data, &varNetworkDestinationNatProperty)

	if err != nil {
		return err
	}

	*o = NetworkDestinationNatProperty(varNetworkDestinationNatProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "internal_ip")
		delete(additionalProperties, "name")
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDestinationNatProperty struct {
	value *NetworkDestinationNatProperty
	isSet bool
}

func (v NullableNetworkDestinationNatProperty) Get() *NetworkDestinationNatProperty {
	return v.value
}

func (v *NullableNetworkDestinationNatProperty) Set(val *NetworkDestinationNatProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDestinationNatProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDestinationNatProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDestinationNatProperty(val *NetworkDestinationNatProperty) *NullableNetworkDestinationNatProperty {
	return &NullableNetworkDestinationNatProperty{value: val, isSet: true}
}

func (v NullableNetworkDestinationNatProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDestinationNatProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


