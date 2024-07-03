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

// checks if the UpgradeBiosMulti type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpgradeBiosMulti{}

// UpgradeBiosMulti struct for UpgradeBiosMulti
type UpgradeBiosMulti struct {
	// list of device id to upgrade bios
	DeviceIds []string `json:"device_ids,omitempty"`
	// list of device model to upgrade bios
	Models []string `json:"models,omitempty"`
	// Reboot device immediately after upgrade is completed
	Reboot *bool `json:"reboot,omitempty"`
	// specific bios version
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpgradeBiosMulti UpgradeBiosMulti

// NewUpgradeBiosMulti instantiates a new UpgradeBiosMulti object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpgradeBiosMulti() *UpgradeBiosMulti {
	this := UpgradeBiosMulti{}
	var reboot bool = false
	this.Reboot = &reboot
	return &this
}

// NewUpgradeBiosMultiWithDefaults instantiates a new UpgradeBiosMulti object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpgradeBiosMultiWithDefaults() *UpgradeBiosMulti {
	this := UpgradeBiosMulti{}
	var reboot bool = false
	this.Reboot = &reboot
	return &this
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *UpgradeBiosMulti) GetDeviceIds() []string {
	if o == nil || IsNil(o.DeviceIds) {
		var ret []string
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeBiosMulti) GetDeviceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceIds) {
		return nil, false
	}
	return o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *UpgradeBiosMulti) HasDeviceIds() bool {
	if o != nil && !IsNil(o.DeviceIds) {
		return true
	}

	return false
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *UpgradeBiosMulti) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetModels returns the Models field value if set, zero value otherwise.
func (o *UpgradeBiosMulti) GetModels() []string {
	if o == nil || IsNil(o.Models) {
		var ret []string
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeBiosMulti) GetModelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Models) {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *UpgradeBiosMulti) HasModels() bool {
	if o != nil && !IsNil(o.Models) {
		return true
	}

	return false
}

// SetModels gets a reference to the given []string and assigns it to the Models field.
func (o *UpgradeBiosMulti) SetModels(v []string) {
	o.Models = v
}

// GetReboot returns the Reboot field value if set, zero value otherwise.
func (o *UpgradeBiosMulti) GetReboot() bool {
	if o == nil || IsNil(o.Reboot) {
		var ret bool
		return ret
	}
	return *o.Reboot
}

// GetRebootOk returns a tuple with the Reboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeBiosMulti) GetRebootOk() (*bool, bool) {
	if o == nil || IsNil(o.Reboot) {
		return nil, false
	}
	return o.Reboot, true
}

// HasReboot returns a boolean if a field has been set.
func (o *UpgradeBiosMulti) HasReboot() bool {
	if o != nil && !IsNil(o.Reboot) {
		return true
	}

	return false
}

// SetReboot gets a reference to the given bool and assigns it to the Reboot field.
func (o *UpgradeBiosMulti) SetReboot(v bool) {
	o.Reboot = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpgradeBiosMulti) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpgradeBiosMulti) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpgradeBiosMulti) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UpgradeBiosMulti) SetVersion(v string) {
	o.Version = &v
}

func (o UpgradeBiosMulti) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpgradeBiosMulti) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceIds) {
		toSerialize["device_ids"] = o.DeviceIds
	}
	if !IsNil(o.Models) {
		toSerialize["models"] = o.Models
	}
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

func (o *UpgradeBiosMulti) UnmarshalJSON(data []byte) (err error) {
	varUpgradeBiosMulti := _UpgradeBiosMulti{}

	err = json.Unmarshal(data, &varUpgradeBiosMulti)

	if err != nil {
		return err
	}

	*o = UpgradeBiosMulti(varUpgradeBiosMulti)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_ids")
		delete(additionalProperties, "models")
		delete(additionalProperties, "reboot")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpgradeBiosMulti struct {
	value *UpgradeBiosMulti
	isSet bool
}

func (v NullableUpgradeBiosMulti) Get() *UpgradeBiosMulti {
	return v.value
}

func (v *NullableUpgradeBiosMulti) Set(val *UpgradeBiosMulti) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeBiosMulti) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeBiosMulti) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeBiosMulti(val *UpgradeBiosMulti) *NullableUpgradeBiosMulti {
	return &NullableUpgradeBiosMulti{value: val, isSet: true}
}

func (v NullableUpgradeBiosMulti) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeBiosMulti) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


