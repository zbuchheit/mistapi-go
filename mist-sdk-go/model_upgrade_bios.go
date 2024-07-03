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

// checks if the UpgradeBios type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradeBios{}

// UpgradeBios struct for UpgradeBios
type UpgradeBios struct {
	// Reboot device immediately after upgrade is completed
	Reboot *bool `json:"reboot,omitempty"`
	// specific bios version
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpgradeBios UpgradeBios

// NewUpgradeBios instantiates a new UpgradeBios object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeBios() *UpgradeBios {
	this := UpgradeBios{}
	var reboot bool = false
	this.Reboot = &reboot
	return &this
}

// NewUpgradeBiosWithDefaults instantiates a new UpgradeBios object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeBiosWithDefaults() *UpgradeBios {
	this := UpgradeBios{}
	var reboot bool = false
	this.Reboot = &reboot
	return &this
}

// GetReboot returns the Reboot field value if set, zero value otherwise.
func (o *UpgradeBios) GetReboot() bool {
	if o == nil || IsNil(o.Reboot) {
		var ret bool
		return ret
	}
	return *o.Reboot
}

// GetRebootOk returns a tuple with the Reboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeBios) GetRebootOk() (*bool, bool) {
	if o == nil || IsNil(o.Reboot) {
		return nil, false
	}
	return o.Reboot, true
}

// HasReboot returns a boolean if a field has been set.
func (o *UpgradeBios) HasReboot() bool {
	if o != nil && !IsNil(o.Reboot) {
		return true
	}

	return false
}

// SetReboot gets a reference to the given bool and assigns it to the Reboot field.
func (o *UpgradeBios) SetReboot(v bool) {
	o.Reboot = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpgradeBios) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeBios) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpgradeBios) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UpgradeBios) SetVersion(v string) {
	o.Version = &v
}

func (o UpgradeBios) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradeBios) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reboot) {
		toSerialize["reboot"] = o.Reboot
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpgradeBios) UnmarshalJSON(data []byte) (err error) {
	varUpgradeBios := _UpgradeBios{}

	err = json.Unmarshal(data, &varUpgradeBios)

	if err != nil {
		return err
	}

	*o = UpgradeBios(varUpgradeBios)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reboot")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpgradeBios struct {
	value *UpgradeBios
	isSet bool
}

func (v NullableUpgradeBios) Get() *UpgradeBios {
	return v.value
}

func (v *NullableUpgradeBios) Set(val *UpgradeBios) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeBios) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeBios) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeBios(val *UpgradeBios) *NullableUpgradeBios {
	return &NullableUpgradeBios{value: val, isSet: true}
}

func (v NullableUpgradeBios) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeBios) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


