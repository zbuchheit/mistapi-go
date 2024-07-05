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

// checks if the SleImpactedClientsClientSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleImpactedClientsClientSwitch{}

// SleImpactedClientsClientSwitch struct for SleImpactedClientsClientSwitch
type SleImpactedClientsClientSwitch struct {
	Interfaces []string `json:"interfaces,omitempty"`
	SwitchMac *string `json:"switch_mac,omitempty"`
	SwitchName *string `json:"switch_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SleImpactedClientsClientSwitch SleImpactedClientsClientSwitch

// NewSleImpactedClientsClientSwitch instantiates a new SleImpactedClientsClientSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleImpactedClientsClientSwitch() *SleImpactedClientsClientSwitch {
	this := SleImpactedClientsClientSwitch{}
	return &this
}

// NewSleImpactedClientsClientSwitchWithDefaults instantiates a new SleImpactedClientsClientSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleImpactedClientsClientSwitchWithDefaults() *SleImpactedClientsClientSwitch {
	this := SleImpactedClientsClientSwitch{}
	return &this
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *SleImpactedClientsClientSwitch) GetInterfaces() []string {
	if o == nil || IsNil(o.Interfaces) {
		var ret []string
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedClientsClientSwitch) GetInterfacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *SleImpactedClientsClientSwitch) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *SleImpactedClientsClientSwitch) SetInterfaces(v []string) {
	o.Interfaces = v
}

// GetSwitchMac returns the SwitchMac field value if set, zero value otherwise.
func (o *SleImpactedClientsClientSwitch) GetSwitchMac() string {
	if o == nil || IsNil(o.SwitchMac) {
		var ret string
		return ret
	}
	return *o.SwitchMac
}

// GetSwitchMacOk returns a tuple with the SwitchMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedClientsClientSwitch) GetSwitchMacOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchMac) {
		return nil, false
	}
	return o.SwitchMac, true
}

// HasSwitchMac returns a boolean if a field has been set.
func (o *SleImpactedClientsClientSwitch) HasSwitchMac() bool {
	if o != nil && !IsNil(o.SwitchMac) {
		return true
	}

	return false
}

// SetSwitchMac gets a reference to the given string and assigns it to the SwitchMac field.
func (o *SleImpactedClientsClientSwitch) SetSwitchMac(v string) {
	o.SwitchMac = &v
}

// GetSwitchName returns the SwitchName field value if set, zero value otherwise.
func (o *SleImpactedClientsClientSwitch) GetSwitchName() string {
	if o == nil || IsNil(o.SwitchName) {
		var ret string
		return ret
	}
	return *o.SwitchName
}

// GetSwitchNameOk returns a tuple with the SwitchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedClientsClientSwitch) GetSwitchNameOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchName) {
		return nil, false
	}
	return o.SwitchName, true
}

// HasSwitchName returns a boolean if a field has been set.
func (o *SleImpactedClientsClientSwitch) HasSwitchName() bool {
	if o != nil && !IsNil(o.SwitchName) {
		return true
	}

	return false
}

// SetSwitchName gets a reference to the given string and assigns it to the SwitchName field.
func (o *SleImpactedClientsClientSwitch) SetSwitchName(v string) {
	o.SwitchName = &v
}

func (o SleImpactedClientsClientSwitch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleImpactedClientsClientSwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	if !IsNil(o.SwitchMac) {
		toSerialize["switch_mac"] = o.SwitchMac
	}
	if !IsNil(o.SwitchName) {
		toSerialize["switch_name"] = o.SwitchName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleImpactedClientsClientSwitch) UnmarshalJSON(data []byte) (err error) {
	varSleImpactedClientsClientSwitch := _SleImpactedClientsClientSwitch{}

	err = json.Unmarshal(data, &varSleImpactedClientsClientSwitch)

	if err != nil {
		return err
	}

	*o = SleImpactedClientsClientSwitch(varSleImpactedClientsClientSwitch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interfaces")
		delete(additionalProperties, "switch_mac")
		delete(additionalProperties, "switch_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleImpactedClientsClientSwitch struct {
	value *SleImpactedClientsClientSwitch
	isSet bool
}

func (v NullableSleImpactedClientsClientSwitch) Get() *SleImpactedClientsClientSwitch {
	return v.value
}

func (v *NullableSleImpactedClientsClientSwitch) Set(val *SleImpactedClientsClientSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableSleImpactedClientsClientSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableSleImpactedClientsClientSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleImpactedClientsClientSwitch(val *SleImpactedClientsClientSwitch) *NullableSleImpactedClientsClientSwitch {
	return &NullableSleImpactedClientsClientSwitch{value: val, isSet: true}
}

func (v NullableSleImpactedClientsClientSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleImpactedClientsClientSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


