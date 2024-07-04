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

// checks if the NetworkStaticNatProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkStaticNatProperty{}

// NetworkStaticNatProperty struct for NetworkStaticNatProperty
type NetworkStaticNatProperty struct {
	InternalIp *string `json:"internal_ip,omitempty"`
	Name *string `json:"name,omitempty"`
	// If not set, we configure the nat policies against all WAN ports for simplicity
	WanName *string `json:"wan_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkStaticNatProperty NetworkStaticNatProperty

// NewNetworkStaticNatProperty instantiates a new NetworkStaticNatProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkStaticNatProperty() *NetworkStaticNatProperty {
	this := NetworkStaticNatProperty{}
	return &this
}

// NewNetworkStaticNatPropertyWithDefaults instantiates a new NetworkStaticNatProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkStaticNatPropertyWithDefaults() *NetworkStaticNatProperty {
	this := NetworkStaticNatProperty{}
	return &this
}

// GetInternalIp returns the InternalIp field value if set, zero value otherwise.
func (o *NetworkStaticNatProperty) GetInternalIp() string {
	if o == nil || IsNil(o.InternalIp) {
		var ret string
		return ret
	}
	return *o.InternalIp
}

// GetInternalIpOk returns a tuple with the InternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkStaticNatProperty) GetInternalIpOk() (*string, bool) {
	if o == nil || IsNil(o.InternalIp) {
		return nil, false
	}
	return o.InternalIp, true
}

// HasInternalIp returns a boolean if a field has been set.
func (o *NetworkStaticNatProperty) HasInternalIp() bool {
	if o != nil && !IsNil(o.InternalIp) {
		return true
	}

	return false
}

// SetInternalIp gets a reference to the given string and assigns it to the InternalIp field.
func (o *NetworkStaticNatProperty) SetInternalIp(v string) {
	o.InternalIp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkStaticNatProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkStaticNatProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkStaticNatProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkStaticNatProperty) SetName(v string) {
	o.Name = &v
}

// GetWanName returns the WanName field value if set, zero value otherwise.
func (o *NetworkStaticNatProperty) GetWanName() string {
	if o == nil || IsNil(o.WanName) {
		var ret string
		return ret
	}
	return *o.WanName
}

// GetWanNameOk returns a tuple with the WanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkStaticNatProperty) GetWanNameOk() (*string, bool) {
	if o == nil || IsNil(o.WanName) {
		return nil, false
	}
	return o.WanName, true
}

// HasWanName returns a boolean if a field has been set.
func (o *NetworkStaticNatProperty) HasWanName() bool {
	if o != nil && !IsNil(o.WanName) {
		return true
	}

	return false
}

// SetWanName gets a reference to the given string and assigns it to the WanName field.
func (o *NetworkStaticNatProperty) SetWanName(v string) {
	o.WanName = &v
}

func (o NetworkStaticNatProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkStaticNatProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InternalIp) {
		toSerialize["internal_ip"] = o.InternalIp
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.WanName) {
		toSerialize["wan_name"] = o.WanName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkStaticNatProperty) UnmarshalJSON(data []byte) (err error) {
	varNetworkStaticNatProperty := _NetworkStaticNatProperty{}

	err = json.Unmarshal(data, &varNetworkStaticNatProperty)

	if err != nil {
		return err
	}

	*o = NetworkStaticNatProperty(varNetworkStaticNatProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "internal_ip")
		delete(additionalProperties, "name")
		delete(additionalProperties, "wan_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkStaticNatProperty struct {
	value *NetworkStaticNatProperty
	isSet bool
}

func (v NullableNetworkStaticNatProperty) Get() *NetworkStaticNatProperty {
	return v.value
}

func (v *NullableNetworkStaticNatProperty) Set(val *NetworkStaticNatProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkStaticNatProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkStaticNatProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkStaticNatProperty(val *NetworkStaticNatProperty) *NullableNetworkStaticNatProperty {
	return &NullableNetworkStaticNatProperty{value: val, isSet: true}
}

func (v NullableNetworkStaticNatProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkStaticNatProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


