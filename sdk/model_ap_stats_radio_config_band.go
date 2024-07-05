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

// checks if the ApStatsRadioConfigBand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatsRadioConfigBand{}

// ApStatsRadioConfigBand struct for ApStatsRadioConfigBand
type ApStatsRadioConfigBand struct {
	Bandwidth *float32 `json:"bandwidth,omitempty"`
	Channel *int32 `json:"channel,omitempty"`
	DynamicChainingEnabled *bool `json:"dynamic_chaining_enabled,omitempty"`
	Power *float32 `json:"power,omitempty"`
	RxChain *int32 `json:"rx_chain,omitempty"`
	TxChain *int32 `json:"tx_chain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatsRadioConfigBand ApStatsRadioConfigBand

// NewApStatsRadioConfigBand instantiates a new ApStatsRadioConfigBand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatsRadioConfigBand() *ApStatsRadioConfigBand {
	this := ApStatsRadioConfigBand{}
	return &this
}

// NewApStatsRadioConfigBandWithDefaults instantiates a new ApStatsRadioConfigBand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatsRadioConfigBandWithDefaults() *ApStatsRadioConfigBand {
	this := ApStatsRadioConfigBand{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *ApStatsRadioConfigBand) GetBandwidth() float32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret float32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfigBand) GetBandwidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *ApStatsRadioConfigBand) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given float32 and assigns it to the Bandwidth field.
func (o *ApStatsRadioConfigBand) SetBandwidth(v float32) {
	o.Bandwidth = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ApStatsRadioConfigBand) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfigBand) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ApStatsRadioConfigBand) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *ApStatsRadioConfigBand) SetChannel(v int32) {
	o.Channel = &v
}

// GetDynamicChainingEnabled returns the DynamicChainingEnabled field value if set, zero value otherwise.
func (o *ApStatsRadioConfigBand) GetDynamicChainingEnabled() bool {
	if o == nil || IsNil(o.DynamicChainingEnabled) {
		var ret bool
		return ret
	}
	return *o.DynamicChainingEnabled
}

// GetDynamicChainingEnabledOk returns a tuple with the DynamicChainingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfigBand) GetDynamicChainingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicChainingEnabled) {
		return nil, false
	}
	return o.DynamicChainingEnabled, true
}

// HasDynamicChainingEnabled returns a boolean if a field has been set.
func (o *ApStatsRadioConfigBand) HasDynamicChainingEnabled() bool {
	if o != nil && !IsNil(o.DynamicChainingEnabled) {
		return true
	}

	return false
}

// SetDynamicChainingEnabled gets a reference to the given bool and assigns it to the DynamicChainingEnabled field.
func (o *ApStatsRadioConfigBand) SetDynamicChainingEnabled(v bool) {
	o.DynamicChainingEnabled = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *ApStatsRadioConfigBand) GetPower() float32 {
	if o == nil || IsNil(o.Power) {
		var ret float32
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfigBand) GetPowerOk() (*float32, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *ApStatsRadioConfigBand) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given float32 and assigns it to the Power field.
func (o *ApStatsRadioConfigBand) SetPower(v float32) {
	o.Power = &v
}

// GetRxChain returns the RxChain field value if set, zero value otherwise.
func (o *ApStatsRadioConfigBand) GetRxChain() int32 {
	if o == nil || IsNil(o.RxChain) {
		var ret int32
		return ret
	}
	return *o.RxChain
}

// GetRxChainOk returns a tuple with the RxChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfigBand) GetRxChainOk() (*int32, bool) {
	if o == nil || IsNil(o.RxChain) {
		return nil, false
	}
	return o.RxChain, true
}

// HasRxChain returns a boolean if a field has been set.
func (o *ApStatsRadioConfigBand) HasRxChain() bool {
	if o != nil && !IsNil(o.RxChain) {
		return true
	}

	return false
}

// SetRxChain gets a reference to the given int32 and assigns it to the RxChain field.
func (o *ApStatsRadioConfigBand) SetRxChain(v int32) {
	o.RxChain = &v
}

// GetTxChain returns the TxChain field value if set, zero value otherwise.
func (o *ApStatsRadioConfigBand) GetTxChain() int32 {
	if o == nil || IsNil(o.TxChain) {
		var ret int32
		return ret
	}
	return *o.TxChain
}

// GetTxChainOk returns a tuple with the TxChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsRadioConfigBand) GetTxChainOk() (*int32, bool) {
	if o == nil || IsNil(o.TxChain) {
		return nil, false
	}
	return o.TxChain, true
}

// HasTxChain returns a boolean if a field has been set.
func (o *ApStatsRadioConfigBand) HasTxChain() bool {
	if o != nil && !IsNil(o.TxChain) {
		return true
	}

	return false
}

// SetTxChain gets a reference to the given int32 and assigns it to the TxChain field.
func (o *ApStatsRadioConfigBand) SetTxChain(v int32) {
	o.TxChain = &v
}

func (o ApStatsRadioConfigBand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatsRadioConfigBand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.DynamicChainingEnabled) {
		toSerialize["dynamic_chaining_enabled"] = o.DynamicChainingEnabled
	}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !IsNil(o.RxChain) {
		toSerialize["rx_chain"] = o.RxChain
	}
	if !IsNil(o.TxChain) {
		toSerialize["tx_chain"] = o.TxChain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApStatsRadioConfigBand) UnmarshalJSON(data []byte) (err error) {
	varApStatsRadioConfigBand := _ApStatsRadioConfigBand{}

	err = json.Unmarshal(data, &varApStatsRadioConfigBand)

	if err != nil {
		return err
	}

	*o = ApStatsRadioConfigBand(varApStatsRadioConfigBand)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "dynamic_chaining_enabled")
		delete(additionalProperties, "power")
		delete(additionalProperties, "rx_chain")
		delete(additionalProperties, "tx_chain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatsRadioConfigBand struct {
	value *ApStatsRadioConfigBand
	isSet bool
}

func (v NullableApStatsRadioConfigBand) Get() *ApStatsRadioConfigBand {
	return v.value
}

func (v *NullableApStatsRadioConfigBand) Set(val *ApStatsRadioConfigBand) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsRadioConfigBand) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsRadioConfigBand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsRadioConfigBand(val *ApStatsRadioConfigBand) *NullableApStatsRadioConfigBand {
	return &NullableApStatsRadioConfigBand{value: val, isSet: true}
}

func (v NullableApStatsRadioConfigBand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsRadioConfigBand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


