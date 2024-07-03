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

// checks if the ApRadioBand6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApRadioBand6{}

// ApRadioBand6 Radio Band AP settings
type ApRadioBand6 struct {
	AllowRrmDisable *bool `json:"allow_rrm_disable,omitempty"`
	AntGain NullableInt32 `json:"ant_gain,omitempty"`
	AntennaMode *RadioBandAntennaMode `json:"antenna_mode,omitempty"`
	Bandwidth *Dot11Bandwidth6 `json:"bandwidth,omitempty"`
	// For Device. (primary) channel for the band, 0 means using the Site Setting
	Channel NullableInt32 `json:"channel,omitempty"`
	// For RFTemplates. List of channels, null or empty array means auto
	Channels []int32 `json:"channels,omitempty"`
	// whether to disable the radio
	Disabled *bool `json:"disabled,omitempty"`
	// TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / …
	Power NullableInt32 `json:"power,omitempty"`
	// when power=0, max tx power to use, HW-specific values will be used if not set
	PowerMax NullableInt32 `json:"power_max,omitempty"`
	// when power=0, min tx power to use, HW-specific values will be used if not set
	PowerMin NullableInt32 `json:"power_min,omitempty"`
	Preamble *RadioBandPreamble `json:"preamble,omitempty"`
	// for 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed and we'll fallback to Low Power Indoor if AFC failed
	StandardPower *bool `json:"standard_power,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApRadioBand6 ApRadioBand6

// NewApRadioBand6 instantiates a new ApRadioBand6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApRadioBand6() *ApRadioBand6 {
	this := ApRadioBand6{}
	var allowRrmDisable bool = false
	this.AllowRrmDisable = &allowRrmDisable
	var antGain int32 = 0
	this.AntGain = *NewNullableInt32(&antGain)
	var antennaMode RadioBandAntennaMode = RADIOBANDANTENNAMODE_DEFAULT
	this.AntennaMode = &antennaMode
	var bandwidth Dot11Bandwidth6 = DOT11BANDWIDTH6__80
	this.Bandwidth = &bandwidth
	var disabled bool = false
	this.Disabled = &disabled
	var powerMax int32 = 18
	this.PowerMax = *NewNullableInt32(&powerMax)
	var powerMin int32 = 8
	this.PowerMin = *NewNullableInt32(&powerMin)
	var preamble RadioBandPreamble = RADIOBANDPREAMBLE_SHORT
	this.Preamble = &preamble
	var standardPower bool = false
	this.StandardPower = &standardPower
	return &this
}

// NewApRadioBand6WithDefaults instantiates a new ApRadioBand6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApRadioBand6WithDefaults() *ApRadioBand6 {
	this := ApRadioBand6{}
	var allowRrmDisable bool = false
	this.AllowRrmDisable = &allowRrmDisable
	var antGain int32 = 0
	this.AntGain = *NewNullableInt32(&antGain)
	var antennaMode RadioBandAntennaMode = RADIOBANDANTENNAMODE_DEFAULT
	this.AntennaMode = &antennaMode
	var bandwidth Dot11Bandwidth6 = DOT11BANDWIDTH6__80
	this.Bandwidth = &bandwidth
	var disabled bool = false
	this.Disabled = &disabled
	var powerMax int32 = 18
	this.PowerMax = *NewNullableInt32(&powerMax)
	var powerMin int32 = 8
	this.PowerMin = *NewNullableInt32(&powerMin)
	var preamble RadioBandPreamble = RADIOBANDPREAMBLE_SHORT
	this.Preamble = &preamble
	var standardPower bool = false
	this.StandardPower = &standardPower
	return &this
}

// GetAllowRrmDisable returns the AllowRrmDisable field value if set, zero value otherwise.
func (o *ApRadioBand6) GetAllowRrmDisable() bool {
	if o == nil || IsNil(o.AllowRrmDisable) {
		var ret bool
		return ret
	}
	return *o.AllowRrmDisable
}

// GetAllowRrmDisableOk returns a tuple with the AllowRrmDisable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioBand6) GetAllowRrmDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRrmDisable) {
		return nil, false
	}
	return o.AllowRrmDisable, true
}

// HasAllowRrmDisable returns a boolean if a field has been set.
func (o *ApRadioBand6) HasAllowRrmDisable() bool {
	if o != nil && !IsNil(o.AllowRrmDisable) {
		return true
	}

	return false
}

// SetAllowRrmDisable gets a reference to the given bool and assigns it to the AllowRrmDisable field.
func (o *ApRadioBand6) SetAllowRrmDisable(v bool) {
	o.AllowRrmDisable = &v
}

// GetAntGain returns the AntGain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApRadioBand6) GetAntGain() int32 {
	if o == nil || IsNil(o.AntGain.Get()) {
		var ret int32
		return ret
	}
	return *o.AntGain.Get()
}

// GetAntGainOk returns a tuple with the AntGain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApRadioBand6) GetAntGainOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AntGain.Get(), o.AntGain.IsSet()
}

// HasAntGain returns a boolean if a field has been set.
func (o *ApRadioBand6) HasAntGain() bool {
	if o != nil && o.AntGain.IsSet() {
		return true
	}

	return false
}

// SetAntGain gets a reference to the given NullableInt32 and assigns it to the AntGain field.
func (o *ApRadioBand6) SetAntGain(v int32) {
	o.AntGain.Set(&v)
}
// SetAntGainNil sets the value for AntGain to be an explicit nil
func (o *ApRadioBand6) SetAntGainNil() {
	o.AntGain.Set(nil)
}

// UnsetAntGain ensures that no value is present for AntGain, not even an explicit nil
func (o *ApRadioBand6) UnsetAntGain() {
	o.AntGain.Unset()
}

// GetAntennaMode returns the AntennaMode field value if set, zero value otherwise.
func (o *ApRadioBand6) GetAntennaMode() RadioBandAntennaMode {
	if o == nil || IsNil(o.AntennaMode) {
		var ret RadioBandAntennaMode
		return ret
	}
	return *o.AntennaMode
}

// GetAntennaModeOk returns a tuple with the AntennaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioBand6) GetAntennaModeOk() (*RadioBandAntennaMode, bool) {
	if o == nil || IsNil(o.AntennaMode) {
		return nil, false
	}
	return o.AntennaMode, true
}

// HasAntennaMode returns a boolean if a field has been set.
func (o *ApRadioBand6) HasAntennaMode() bool {
	if o != nil && !IsNil(o.AntennaMode) {
		return true
	}

	return false
}

// SetAntennaMode gets a reference to the given RadioBandAntennaMode and assigns it to the AntennaMode field.
func (o *ApRadioBand6) SetAntennaMode(v RadioBandAntennaMode) {
	o.AntennaMode = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *ApRadioBand6) GetBandwidth() Dot11Bandwidth6 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret Dot11Bandwidth6
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioBand6) GetBandwidthOk() (*Dot11Bandwidth6, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *ApRadioBand6) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given Dot11Bandwidth6 and assigns it to the Bandwidth field.
func (o *ApRadioBand6) SetBandwidth(v Dot11Bandwidth6) {
	o.Bandwidth = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApRadioBand6) GetChannel() int32 {
	if o == nil || IsNil(o.Channel.Get()) {
		var ret int32
		return ret
	}
	return *o.Channel.Get()
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApRadioBand6) GetChannelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channel.Get(), o.Channel.IsSet()
}

// HasChannel returns a boolean if a field has been set.
func (o *ApRadioBand6) HasChannel() bool {
	if o != nil && o.Channel.IsSet() {
		return true
	}

	return false
}

// SetChannel gets a reference to the given NullableInt32 and assigns it to the Channel field.
func (o *ApRadioBand6) SetChannel(v int32) {
	o.Channel.Set(&v)
}
// SetChannelNil sets the value for Channel to be an explicit nil
func (o *ApRadioBand6) SetChannelNil() {
	o.Channel.Set(nil)
}

// UnsetChannel ensures that no value is present for Channel, not even an explicit nil
func (o *ApRadioBand6) UnsetChannel() {
	o.Channel.Unset()
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApRadioBand6) GetChannels() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApRadioBand6) GetChannelsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *ApRadioBand6) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []int32 and assigns it to the Channels field.
func (o *ApRadioBand6) SetChannels(v []int32) {
	o.Channels = v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ApRadioBand6) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioBand6) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ApRadioBand6) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ApRadioBand6) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetPower returns the Power field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApRadioBand6) GetPower() int32 {
	if o == nil || IsNil(o.Power.Get()) {
		var ret int32
		return ret
	}
	return *o.Power.Get()
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApRadioBand6) GetPowerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Power.Get(), o.Power.IsSet()
}

// HasPower returns a boolean if a field has been set.
func (o *ApRadioBand6) HasPower() bool {
	if o != nil && o.Power.IsSet() {
		return true
	}

	return false
}

// SetPower gets a reference to the given NullableInt32 and assigns it to the Power field.
func (o *ApRadioBand6) SetPower(v int32) {
	o.Power.Set(&v)
}
// SetPowerNil sets the value for Power to be an explicit nil
func (o *ApRadioBand6) SetPowerNil() {
	o.Power.Set(nil)
}

// UnsetPower ensures that no value is present for Power, not even an explicit nil
func (o *ApRadioBand6) UnsetPower() {
	o.Power.Unset()
}

// GetPowerMax returns the PowerMax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApRadioBand6) GetPowerMax() int32 {
	if o == nil || IsNil(o.PowerMax.Get()) {
		var ret int32
		return ret
	}
	return *o.PowerMax.Get()
}

// GetPowerMaxOk returns a tuple with the PowerMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApRadioBand6) GetPowerMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerMax.Get(), o.PowerMax.IsSet()
}

// HasPowerMax returns a boolean if a field has been set.
func (o *ApRadioBand6) HasPowerMax() bool {
	if o != nil && o.PowerMax.IsSet() {
		return true
	}

	return false
}

// SetPowerMax gets a reference to the given NullableInt32 and assigns it to the PowerMax field.
func (o *ApRadioBand6) SetPowerMax(v int32) {
	o.PowerMax.Set(&v)
}
// SetPowerMaxNil sets the value for PowerMax to be an explicit nil
func (o *ApRadioBand6) SetPowerMaxNil() {
	o.PowerMax.Set(nil)
}

// UnsetPowerMax ensures that no value is present for PowerMax, not even an explicit nil
func (o *ApRadioBand6) UnsetPowerMax() {
	o.PowerMax.Unset()
}

// GetPowerMin returns the PowerMin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApRadioBand6) GetPowerMin() int32 {
	if o == nil || IsNil(o.PowerMin.Get()) {
		var ret int32
		return ret
	}
	return *o.PowerMin.Get()
}

// GetPowerMinOk returns a tuple with the PowerMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApRadioBand6) GetPowerMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerMin.Get(), o.PowerMin.IsSet()
}

// HasPowerMin returns a boolean if a field has been set.
func (o *ApRadioBand6) HasPowerMin() bool {
	if o != nil && o.PowerMin.IsSet() {
		return true
	}

	return false
}

// SetPowerMin gets a reference to the given NullableInt32 and assigns it to the PowerMin field.
func (o *ApRadioBand6) SetPowerMin(v int32) {
	o.PowerMin.Set(&v)
}
// SetPowerMinNil sets the value for PowerMin to be an explicit nil
func (o *ApRadioBand6) SetPowerMinNil() {
	o.PowerMin.Set(nil)
}

// UnsetPowerMin ensures that no value is present for PowerMin, not even an explicit nil
func (o *ApRadioBand6) UnsetPowerMin() {
	o.PowerMin.Unset()
}

// GetPreamble returns the Preamble field value if set, zero value otherwise.
func (o *ApRadioBand6) GetPreamble() RadioBandPreamble {
	if o == nil || IsNil(o.Preamble) {
		var ret RadioBandPreamble
		return ret
	}
	return *o.Preamble
}

// GetPreambleOk returns a tuple with the Preamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioBand6) GetPreambleOk() (*RadioBandPreamble, bool) {
	if o == nil || IsNil(o.Preamble) {
		return nil, false
	}
	return o.Preamble, true
}

// HasPreamble returns a boolean if a field has been set.
func (o *ApRadioBand6) HasPreamble() bool {
	if o != nil && !IsNil(o.Preamble) {
		return true
	}

	return false
}

// SetPreamble gets a reference to the given RadioBandPreamble and assigns it to the Preamble field.
func (o *ApRadioBand6) SetPreamble(v RadioBandPreamble) {
	o.Preamble = &v
}

// GetStandardPower returns the StandardPower field value if set, zero value otherwise.
func (o *ApRadioBand6) GetStandardPower() bool {
	if o == nil || IsNil(o.StandardPower) {
		var ret bool
		return ret
	}
	return *o.StandardPower
}

// GetStandardPowerOk returns a tuple with the StandardPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioBand6) GetStandardPowerOk() (*bool, bool) {
	if o == nil || IsNil(o.StandardPower) {
		return nil, false
	}
	return o.StandardPower, true
}

// HasStandardPower returns a boolean if a field has been set.
func (o *ApRadioBand6) HasStandardPower() bool {
	if o != nil && !IsNil(o.StandardPower) {
		return true
	}

	return false
}

// SetStandardPower gets a reference to the given bool and assigns it to the StandardPower field.
func (o *ApRadioBand6) SetStandardPower(v bool) {
	o.StandardPower = &v
}

func (o ApRadioBand6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApRadioBand6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowRrmDisable) {
		toSerialize["allow_rrm_disable"] = o.AllowRrmDisable
	}
	if o.AntGain.IsSet() {
		toSerialize["ant_gain"] = o.AntGain.Get()
	}
	if !IsNil(o.AntennaMode) {
		toSerialize["antenna_mode"] = o.AntennaMode
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.Channel.IsSet() {
		toSerialize["channel"] = o.Channel.Get()
	}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Power.IsSet() {
		toSerialize["power"] = o.Power.Get()
	}
	if o.PowerMax.IsSet() {
		toSerialize["power_max"] = o.PowerMax.Get()
	}
	if o.PowerMin.IsSet() {
		toSerialize["power_min"] = o.PowerMin.Get()
	}
	if !IsNil(o.Preamble) {
		toSerialize["preamble"] = o.Preamble
	}
	if !IsNil(o.StandardPower) {
		toSerialize["standard_power"] = o.StandardPower
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApRadioBand6) UnmarshalJSON(data []byte) (err error) {
	varApRadioBand6 := _ApRadioBand6{}

	err = json.Unmarshal(data, &varApRadioBand6)

	if err != nil {
		return err
	}

	*o = ApRadioBand6(varApRadioBand6)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allow_rrm_disable")
		delete(additionalProperties, "ant_gain")
		delete(additionalProperties, "antenna_mode")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "channels")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "power")
		delete(additionalProperties, "power_max")
		delete(additionalProperties, "power_min")
		delete(additionalProperties, "preamble")
		delete(additionalProperties, "standard_power")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApRadioBand6 struct {
	value *ApRadioBand6
	isSet bool
}

func (v NullableApRadioBand6) Get() *ApRadioBand6 {
	return v.value
}

func (v *NullableApRadioBand6) Set(val *ApRadioBand6) {
	v.value = val
	v.isSet = true
}

func (v NullableApRadioBand6) IsSet() bool {
	return v.isSet
}

func (v *NullableApRadioBand6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApRadioBand6(val *ApRadioBand6) *NullableApRadioBand6 {
	return &NullableApRadioBand6{value: val, isSet: true}
}

func (v NullableApRadioBand6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApRadioBand6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


