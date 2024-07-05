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

// checks if the SwitchPortUsageStormControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchPortUsageStormControl{}

// SwitchPortUsageStormControl Switch storm control Only if `mode`!=`dynamic`
type SwitchPortUsageStormControl struct {
	// whether to disable storm control on broadcast traffic
	NoBroadcast *bool `json:"no_broadcast,omitempty"`
	// whether to disable storm control on multicast traffic
	NoMulticast *bool `json:"no_multicast,omitempty"`
	// whether to disable storm control on registered multicast traffic
	NoRegisteredMulticast *bool `json:"no_registered_multicast,omitempty"`
	// whether to disable storm control on unknown unicast traffic
	NoUnknownUnicast *bool `json:"no_unknown_unicast,omitempty"`
	// bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth
	Percentage *int32 `json:"percentage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwitchPortUsageStormControl SwitchPortUsageStormControl

// NewSwitchPortUsageStormControl instantiates a new SwitchPortUsageStormControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchPortUsageStormControl() *SwitchPortUsageStormControl {
	this := SwitchPortUsageStormControl{}
	var noBroadcast bool = false
	this.NoBroadcast = &noBroadcast
	var noMulticast bool = false
	this.NoMulticast = &noMulticast
	var noRegisteredMulticast bool = false
	this.NoRegisteredMulticast = &noRegisteredMulticast
	var noUnknownUnicast bool = false
	this.NoUnknownUnicast = &noUnknownUnicast
	var percentage int32 = 80
	this.Percentage = &percentage
	return &this
}

// NewSwitchPortUsageStormControlWithDefaults instantiates a new SwitchPortUsageStormControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchPortUsageStormControlWithDefaults() *SwitchPortUsageStormControl {
	this := SwitchPortUsageStormControl{}
	var noBroadcast bool = false
	this.NoBroadcast = &noBroadcast
	var noMulticast bool = false
	this.NoMulticast = &noMulticast
	var noRegisteredMulticast bool = false
	this.NoRegisteredMulticast = &noRegisteredMulticast
	var noUnknownUnicast bool = false
	this.NoUnknownUnicast = &noUnknownUnicast
	var percentage int32 = 80
	this.Percentage = &percentage
	return &this
}

// GetNoBroadcast returns the NoBroadcast field value if set, zero value otherwise.
func (o *SwitchPortUsageStormControl) GetNoBroadcast() bool {
	if o == nil || IsNil(o.NoBroadcast) {
		var ret bool
		return ret
	}
	return *o.NoBroadcast
}

// GetNoBroadcastOk returns a tuple with the NoBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchPortUsageStormControl) GetNoBroadcastOk() (*bool, bool) {
	if o == nil || IsNil(o.NoBroadcast) {
		return nil, false
	}
	return o.NoBroadcast, true
}

// HasNoBroadcast returns a boolean if a field has been set.
func (o *SwitchPortUsageStormControl) HasNoBroadcast() bool {
	if o != nil && !IsNil(o.NoBroadcast) {
		return true
	}

	return false
}

// SetNoBroadcast gets a reference to the given bool and assigns it to the NoBroadcast field.
func (o *SwitchPortUsageStormControl) SetNoBroadcast(v bool) {
	o.NoBroadcast = &v
}

// GetNoMulticast returns the NoMulticast field value if set, zero value otherwise.
func (o *SwitchPortUsageStormControl) GetNoMulticast() bool {
	if o == nil || IsNil(o.NoMulticast) {
		var ret bool
		return ret
	}
	return *o.NoMulticast
}

// GetNoMulticastOk returns a tuple with the NoMulticast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchPortUsageStormControl) GetNoMulticastOk() (*bool, bool) {
	if o == nil || IsNil(o.NoMulticast) {
		return nil, false
	}
	return o.NoMulticast, true
}

// HasNoMulticast returns a boolean if a field has been set.
func (o *SwitchPortUsageStormControl) HasNoMulticast() bool {
	if o != nil && !IsNil(o.NoMulticast) {
		return true
	}

	return false
}

// SetNoMulticast gets a reference to the given bool and assigns it to the NoMulticast field.
func (o *SwitchPortUsageStormControl) SetNoMulticast(v bool) {
	o.NoMulticast = &v
}

// GetNoRegisteredMulticast returns the NoRegisteredMulticast field value if set, zero value otherwise.
func (o *SwitchPortUsageStormControl) GetNoRegisteredMulticast() bool {
	if o == nil || IsNil(o.NoRegisteredMulticast) {
		var ret bool
		return ret
	}
	return *o.NoRegisteredMulticast
}

// GetNoRegisteredMulticastOk returns a tuple with the NoRegisteredMulticast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchPortUsageStormControl) GetNoRegisteredMulticastOk() (*bool, bool) {
	if o == nil || IsNil(o.NoRegisteredMulticast) {
		return nil, false
	}
	return o.NoRegisteredMulticast, true
}

// HasNoRegisteredMulticast returns a boolean if a field has been set.
func (o *SwitchPortUsageStormControl) HasNoRegisteredMulticast() bool {
	if o != nil && !IsNil(o.NoRegisteredMulticast) {
		return true
	}

	return false
}

// SetNoRegisteredMulticast gets a reference to the given bool and assigns it to the NoRegisteredMulticast field.
func (o *SwitchPortUsageStormControl) SetNoRegisteredMulticast(v bool) {
	o.NoRegisteredMulticast = &v
}

// GetNoUnknownUnicast returns the NoUnknownUnicast field value if set, zero value otherwise.
func (o *SwitchPortUsageStormControl) GetNoUnknownUnicast() bool {
	if o == nil || IsNil(o.NoUnknownUnicast) {
		var ret bool
		return ret
	}
	return *o.NoUnknownUnicast
}

// GetNoUnknownUnicastOk returns a tuple with the NoUnknownUnicast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchPortUsageStormControl) GetNoUnknownUnicastOk() (*bool, bool) {
	if o == nil || IsNil(o.NoUnknownUnicast) {
		return nil, false
	}
	return o.NoUnknownUnicast, true
}

// HasNoUnknownUnicast returns a boolean if a field has been set.
func (o *SwitchPortUsageStormControl) HasNoUnknownUnicast() bool {
	if o != nil && !IsNil(o.NoUnknownUnicast) {
		return true
	}

	return false
}

// SetNoUnknownUnicast gets a reference to the given bool and assigns it to the NoUnknownUnicast field.
func (o *SwitchPortUsageStormControl) SetNoUnknownUnicast(v bool) {
	o.NoUnknownUnicast = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *SwitchPortUsageStormControl) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchPortUsageStormControl) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *SwitchPortUsageStormControl) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *SwitchPortUsageStormControl) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o SwitchPortUsageStormControl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchPortUsageStormControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoBroadcast) {
		toSerialize["no_broadcast"] = o.NoBroadcast
	}
	if !IsNil(o.NoMulticast) {
		toSerialize["no_multicast"] = o.NoMulticast
	}
	if !IsNil(o.NoRegisteredMulticast) {
		toSerialize["no_registered_multicast"] = o.NoRegisteredMulticast
	}
	if !IsNil(o.NoUnknownUnicast) {
		toSerialize["no_unknown_unicast"] = o.NoUnknownUnicast
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwitchPortUsageStormControl) UnmarshalJSON(data []byte) (err error) {
	varSwitchPortUsageStormControl := _SwitchPortUsageStormControl{}

	err = json.Unmarshal(data, &varSwitchPortUsageStormControl)

	if err != nil {
		return err
	}

	*o = SwitchPortUsageStormControl(varSwitchPortUsageStormControl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "no_broadcast")
		delete(additionalProperties, "no_multicast")
		delete(additionalProperties, "no_registered_multicast")
		delete(additionalProperties, "no_unknown_unicast")
		delete(additionalProperties, "percentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwitchPortUsageStormControl struct {
	value *SwitchPortUsageStormControl
	isSet bool
}

func (v NullableSwitchPortUsageStormControl) Get() *SwitchPortUsageStormControl {
	return v.value
}

func (v *NullableSwitchPortUsageStormControl) Set(val *SwitchPortUsageStormControl) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchPortUsageStormControl) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchPortUsageStormControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchPortUsageStormControl(val *SwitchPortUsageStormControl) *NullableSwitchPortUsageStormControl {
	return &NullableSwitchPortUsageStormControl{value: val, isSet: true}
}

func (v NullableSwitchPortUsageStormControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchPortUsageStormControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


