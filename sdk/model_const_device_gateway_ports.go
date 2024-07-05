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

// checks if the ConstDeviceGatewayPorts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstDeviceGatewayPorts{}

// ConstDeviceGatewayPorts Object Key is the interface name (e.g. \"ge-0/0/1\", ...)
type ConstDeviceGatewayPorts struct {
	Display *string `json:"display,omitempty"`
	PciAddress *string `json:"pci_address,omitempty"`
	Speed *int32 `json:"speed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConstDeviceGatewayPorts ConstDeviceGatewayPorts

// NewConstDeviceGatewayPorts instantiates a new ConstDeviceGatewayPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstDeviceGatewayPorts() *ConstDeviceGatewayPorts {
	this := ConstDeviceGatewayPorts{}
	return &this
}

// NewConstDeviceGatewayPortsWithDefaults instantiates a new ConstDeviceGatewayPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstDeviceGatewayPortsWithDefaults() *ConstDeviceGatewayPorts {
	this := ConstDeviceGatewayPorts{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ConstDeviceGatewayPorts) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGatewayPorts) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ConstDeviceGatewayPorts) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ConstDeviceGatewayPorts) SetDisplay(v string) {
	o.Display = &v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *ConstDeviceGatewayPorts) GetPciAddress() string {
	if o == nil || IsNil(o.PciAddress) {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGatewayPorts) GetPciAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PciAddress) {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *ConstDeviceGatewayPorts) HasPciAddress() bool {
	if o != nil && !IsNil(o.PciAddress) {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *ConstDeviceGatewayPorts) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *ConstDeviceGatewayPorts) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstDeviceGatewayPorts) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *ConstDeviceGatewayPorts) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *ConstDeviceGatewayPorts) SetSpeed(v int32) {
	o.Speed = &v
}

func (o ConstDeviceGatewayPorts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstDeviceGatewayPorts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.PciAddress) {
		toSerialize["pci_address"] = o.PciAddress
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConstDeviceGatewayPorts) UnmarshalJSON(data []byte) (err error) {
	varConstDeviceGatewayPorts := _ConstDeviceGatewayPorts{}

	err = json.Unmarshal(data, &varConstDeviceGatewayPorts)

	if err != nil {
		return err
	}

	*o = ConstDeviceGatewayPorts(varConstDeviceGatewayPorts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display")
		delete(additionalProperties, "pci_address")
		delete(additionalProperties, "speed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConstDeviceGatewayPorts struct {
	value *ConstDeviceGatewayPorts
	isSet bool
}

func (v NullableConstDeviceGatewayPorts) Get() *ConstDeviceGatewayPorts {
	return v.value
}

func (v *NullableConstDeviceGatewayPorts) Set(val *ConstDeviceGatewayPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceGatewayPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceGatewayPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceGatewayPorts(val *ConstDeviceGatewayPorts) *NullableConstDeviceGatewayPorts {
	return &NullableConstDeviceGatewayPorts{value: val, isSet: true}
}

func (v NullableConstDeviceGatewayPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceGatewayPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


