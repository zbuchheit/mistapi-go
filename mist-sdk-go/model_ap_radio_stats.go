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

// checks if the ApRadioStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApRadioStats{}

// ApRadioStats radio stat
type ApRadioStats struct {
	Bandwidth *Dot11Bandwidth `json:"bandwidth,omitempty"`
	// current channel the radio is running on
	Channel *int32 `json:"channel,omitempty"`
	// Use dynamic chaining for downlink
	DynamicChainingEnalbed *bool `json:"dynamic_chaining_enalbed,omitempty"`
	// radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af)
	Mac *string `json:"mac,omitempty"`
	NumClients *int32 `json:"num_clients,omitempty"`
	// transmit power (in dBm)
	Power *int32 `json:"power,omitempty"`
	RxBytes *float32 `json:"rx_bytes,omitempty"`
	RxPkts *float32 `json:"rx_pkts,omitempty"`
	TxBytes *float32 `json:"tx_bytes,omitempty"`
	TxPkts *float32 `json:"tx_pkts,omitempty"`
	// all utilization in percentage
	UtilAll *int32 `json:"util_all,omitempty"`
	// reception of “No Packets” utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise
	UtilNonWifi *int32 `json:"util_non_wifi,omitempty"`
	// reception of “In BSS” utilization in percentage, only frames that are received from AP/STAs within the BSS
	UtilRxInBss *int32 `json:"util_rx_in_bss,omitempty"`
	// reception of “Other BSS” utilization in percentage, all frames received from AP/STAs that are outside the BSS
	UtilRxOtherBss *int32 `json:"util_rx_other_bss,omitempty"`
	// transmission utilization in percentage
	UtilTx *int32 `json:"util_tx,omitempty"`
	// reception of “UnDecodable Wifi“ utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio
	UtilUndecodableWifi *int32 `json:"util_undecodable_wifi,omitempty"`
	// reception of “No Category” utilization in percentage, all 802.11 frames that are corrupted at the receiver
	UtilUnknownWifi *int32 `json:"util_unknown_wifi,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApRadioStats ApRadioStats

// NewApRadioStats instantiates a new ApRadioStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApRadioStats() *ApRadioStats {
	this := ApRadioStats{}
	return &this
}

// NewApRadioStatsWithDefaults instantiates a new ApRadioStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApRadioStatsWithDefaults() *ApRadioStats {
	this := ApRadioStats{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *ApRadioStats) GetBandwidth() Dot11Bandwidth {
	if o == nil || IsNil(o.Bandwidth) {
		var ret Dot11Bandwidth
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetBandwidthOk() (*Dot11Bandwidth, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *ApRadioStats) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given Dot11Bandwidth and assigns it to the Bandwidth field.
func (o *ApRadioStats) SetBandwidth(v Dot11Bandwidth) {
	o.Bandwidth = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ApRadioStats) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ApRadioStats) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *ApRadioStats) SetChannel(v int32) {
	o.Channel = &v
}

// GetDynamicChainingEnalbed returns the DynamicChainingEnalbed field value if set, zero value otherwise.
func (o *ApRadioStats) GetDynamicChainingEnalbed() bool {
	if o == nil || IsNil(o.DynamicChainingEnalbed) {
		var ret bool
		return ret
	}
	return *o.DynamicChainingEnalbed
}

// GetDynamicChainingEnalbedOk returns a tuple with the DynamicChainingEnalbed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetDynamicChainingEnalbedOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicChainingEnalbed) {
		return nil, false
	}
	return o.DynamicChainingEnalbed, true
}

// HasDynamicChainingEnalbed returns a boolean if a field has been set.
func (o *ApRadioStats) HasDynamicChainingEnalbed() bool {
	if o != nil && !IsNil(o.DynamicChainingEnalbed) {
		return true
	}

	return false
}

// SetDynamicChainingEnalbed gets a reference to the given bool and assigns it to the DynamicChainingEnalbed field.
func (o *ApRadioStats) SetDynamicChainingEnalbed(v bool) {
	o.DynamicChainingEnalbed = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *ApRadioStats) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *ApRadioStats) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *ApRadioStats) SetMac(v string) {
	o.Mac = &v
}

// GetNumClients returns the NumClients field value if set, zero value otherwise.
func (o *ApRadioStats) GetNumClients() int32 {
	if o == nil || IsNil(o.NumClients) {
		var ret int32
		return ret
	}
	return *o.NumClients
}

// GetNumClientsOk returns a tuple with the NumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetNumClientsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumClients) {
		return nil, false
	}
	return o.NumClients, true
}

// HasNumClients returns a boolean if a field has been set.
func (o *ApRadioStats) HasNumClients() bool {
	if o != nil && !IsNil(o.NumClients) {
		return true
	}

	return false
}

// SetNumClients gets a reference to the given int32 and assigns it to the NumClients field.
func (o *ApRadioStats) SetNumClients(v int32) {
	o.NumClients = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *ApRadioStats) GetPower() int32 {
	if o == nil || IsNil(o.Power) {
		var ret int32
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *ApRadioStats) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given int32 and assigns it to the Power field.
func (o *ApRadioStats) SetPower(v int32) {
	o.Power = &v
}

// GetRxBytes returns the RxBytes field value if set, zero value otherwise.
func (o *ApRadioStats) GetRxBytes() float32 {
	if o == nil || IsNil(o.RxBytes) {
		var ret float32
		return ret
	}
	return *o.RxBytes
}

// GetRxBytesOk returns a tuple with the RxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetRxBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.RxBytes) {
		return nil, false
	}
	return o.RxBytes, true
}

// HasRxBytes returns a boolean if a field has been set.
func (o *ApRadioStats) HasRxBytes() bool {
	if o != nil && !IsNil(o.RxBytes) {
		return true
	}

	return false
}

// SetRxBytes gets a reference to the given float32 and assigns it to the RxBytes field.
func (o *ApRadioStats) SetRxBytes(v float32) {
	o.RxBytes = &v
}

// GetRxPkts returns the RxPkts field value if set, zero value otherwise.
func (o *ApRadioStats) GetRxPkts() float32 {
	if o == nil || IsNil(o.RxPkts) {
		var ret float32
		return ret
	}
	return *o.RxPkts
}

// GetRxPktsOk returns a tuple with the RxPkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetRxPktsOk() (*float32, bool) {
	if o == nil || IsNil(o.RxPkts) {
		return nil, false
	}
	return o.RxPkts, true
}

// HasRxPkts returns a boolean if a field has been set.
func (o *ApRadioStats) HasRxPkts() bool {
	if o != nil && !IsNil(o.RxPkts) {
		return true
	}

	return false
}

// SetRxPkts gets a reference to the given float32 and assigns it to the RxPkts field.
func (o *ApRadioStats) SetRxPkts(v float32) {
	o.RxPkts = &v
}

// GetTxBytes returns the TxBytes field value if set, zero value otherwise.
func (o *ApRadioStats) GetTxBytes() float32 {
	if o == nil || IsNil(o.TxBytes) {
		var ret float32
		return ret
	}
	return *o.TxBytes
}

// GetTxBytesOk returns a tuple with the TxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetTxBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.TxBytes) {
		return nil, false
	}
	return o.TxBytes, true
}

// HasTxBytes returns a boolean if a field has been set.
func (o *ApRadioStats) HasTxBytes() bool {
	if o != nil && !IsNil(o.TxBytes) {
		return true
	}

	return false
}

// SetTxBytes gets a reference to the given float32 and assigns it to the TxBytes field.
func (o *ApRadioStats) SetTxBytes(v float32) {
	o.TxBytes = &v
}

// GetTxPkts returns the TxPkts field value if set, zero value otherwise.
func (o *ApRadioStats) GetTxPkts() float32 {
	if o == nil || IsNil(o.TxPkts) {
		var ret float32
		return ret
	}
	return *o.TxPkts
}

// GetTxPktsOk returns a tuple with the TxPkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetTxPktsOk() (*float32, bool) {
	if o == nil || IsNil(o.TxPkts) {
		return nil, false
	}
	return o.TxPkts, true
}

// HasTxPkts returns a boolean if a field has been set.
func (o *ApRadioStats) HasTxPkts() bool {
	if o != nil && !IsNil(o.TxPkts) {
		return true
	}

	return false
}

// SetTxPkts gets a reference to the given float32 and assigns it to the TxPkts field.
func (o *ApRadioStats) SetTxPkts(v float32) {
	o.TxPkts = &v
}

// GetUtilAll returns the UtilAll field value if set, zero value otherwise.
func (o *ApRadioStats) GetUtilAll() int32 {
	if o == nil || IsNil(o.UtilAll) {
		var ret int32
		return ret
	}
	return *o.UtilAll
}

// GetUtilAllOk returns a tuple with the UtilAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetUtilAllOk() (*int32, bool) {
	if o == nil || IsNil(o.UtilAll) {
		return nil, false
	}
	return o.UtilAll, true
}

// HasUtilAll returns a boolean if a field has been set.
func (o *ApRadioStats) HasUtilAll() bool {
	if o != nil && !IsNil(o.UtilAll) {
		return true
	}

	return false
}

// SetUtilAll gets a reference to the given int32 and assigns it to the UtilAll field.
func (o *ApRadioStats) SetUtilAll(v int32) {
	o.UtilAll = &v
}

// GetUtilNonWifi returns the UtilNonWifi field value if set, zero value otherwise.
func (o *ApRadioStats) GetUtilNonWifi() int32 {
	if o == nil || IsNil(o.UtilNonWifi) {
		var ret int32
		return ret
	}
	return *o.UtilNonWifi
}

// GetUtilNonWifiOk returns a tuple with the UtilNonWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetUtilNonWifiOk() (*int32, bool) {
	if o == nil || IsNil(o.UtilNonWifi) {
		return nil, false
	}
	return o.UtilNonWifi, true
}

// HasUtilNonWifi returns a boolean if a field has been set.
func (o *ApRadioStats) HasUtilNonWifi() bool {
	if o != nil && !IsNil(o.UtilNonWifi) {
		return true
	}

	return false
}

// SetUtilNonWifi gets a reference to the given int32 and assigns it to the UtilNonWifi field.
func (o *ApRadioStats) SetUtilNonWifi(v int32) {
	o.UtilNonWifi = &v
}

// GetUtilRxInBss returns the UtilRxInBss field value if set, zero value otherwise.
func (o *ApRadioStats) GetUtilRxInBss() int32 {
	if o == nil || IsNil(o.UtilRxInBss) {
		var ret int32
		return ret
	}
	return *o.UtilRxInBss
}

// GetUtilRxInBssOk returns a tuple with the UtilRxInBss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetUtilRxInBssOk() (*int32, bool) {
	if o == nil || IsNil(o.UtilRxInBss) {
		return nil, false
	}
	return o.UtilRxInBss, true
}

// HasUtilRxInBss returns a boolean if a field has been set.
func (o *ApRadioStats) HasUtilRxInBss() bool {
	if o != nil && !IsNil(o.UtilRxInBss) {
		return true
	}

	return false
}

// SetUtilRxInBss gets a reference to the given int32 and assigns it to the UtilRxInBss field.
func (o *ApRadioStats) SetUtilRxInBss(v int32) {
	o.UtilRxInBss = &v
}

// GetUtilRxOtherBss returns the UtilRxOtherBss field value if set, zero value otherwise.
func (o *ApRadioStats) GetUtilRxOtherBss() int32 {
	if o == nil || IsNil(o.UtilRxOtherBss) {
		var ret int32
		return ret
	}
	return *o.UtilRxOtherBss
}

// GetUtilRxOtherBssOk returns a tuple with the UtilRxOtherBss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetUtilRxOtherBssOk() (*int32, bool) {
	if o == nil || IsNil(o.UtilRxOtherBss) {
		return nil, false
	}
	return o.UtilRxOtherBss, true
}

// HasUtilRxOtherBss returns a boolean if a field has been set.
func (o *ApRadioStats) HasUtilRxOtherBss() bool {
	if o != nil && !IsNil(o.UtilRxOtherBss) {
		return true
	}

	return false
}

// SetUtilRxOtherBss gets a reference to the given int32 and assigns it to the UtilRxOtherBss field.
func (o *ApRadioStats) SetUtilRxOtherBss(v int32) {
	o.UtilRxOtherBss = &v
}

// GetUtilTx returns the UtilTx field value if set, zero value otherwise.
func (o *ApRadioStats) GetUtilTx() int32 {
	if o == nil || IsNil(o.UtilTx) {
		var ret int32
		return ret
	}
	return *o.UtilTx
}

// GetUtilTxOk returns a tuple with the UtilTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetUtilTxOk() (*int32, bool) {
	if o == nil || IsNil(o.UtilTx) {
		return nil, false
	}
	return o.UtilTx, true
}

// HasUtilTx returns a boolean if a field has been set.
func (o *ApRadioStats) HasUtilTx() bool {
	if o != nil && !IsNil(o.UtilTx) {
		return true
	}

	return false
}

// SetUtilTx gets a reference to the given int32 and assigns it to the UtilTx field.
func (o *ApRadioStats) SetUtilTx(v int32) {
	o.UtilTx = &v
}

// GetUtilUndecodableWifi returns the UtilUndecodableWifi field value if set, zero value otherwise.
func (o *ApRadioStats) GetUtilUndecodableWifi() int32 {
	if o == nil || IsNil(o.UtilUndecodableWifi) {
		var ret int32
		return ret
	}
	return *o.UtilUndecodableWifi
}

// GetUtilUndecodableWifiOk returns a tuple with the UtilUndecodableWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetUtilUndecodableWifiOk() (*int32, bool) {
	if o == nil || IsNil(o.UtilUndecodableWifi) {
		return nil, false
	}
	return o.UtilUndecodableWifi, true
}

// HasUtilUndecodableWifi returns a boolean if a field has been set.
func (o *ApRadioStats) HasUtilUndecodableWifi() bool {
	if o != nil && !IsNil(o.UtilUndecodableWifi) {
		return true
	}

	return false
}

// SetUtilUndecodableWifi gets a reference to the given int32 and assigns it to the UtilUndecodableWifi field.
func (o *ApRadioStats) SetUtilUndecodableWifi(v int32) {
	o.UtilUndecodableWifi = &v
}

// GetUtilUnknownWifi returns the UtilUnknownWifi field value if set, zero value otherwise.
func (o *ApRadioStats) GetUtilUnknownWifi() int32 {
	if o == nil || IsNil(o.UtilUnknownWifi) {
		var ret int32
		return ret
	}
	return *o.UtilUnknownWifi
}

// GetUtilUnknownWifiOk returns a tuple with the UtilUnknownWifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApRadioStats) GetUtilUnknownWifiOk() (*int32, bool) {
	if o == nil || IsNil(o.UtilUnknownWifi) {
		return nil, false
	}
	return o.UtilUnknownWifi, true
}

// HasUtilUnknownWifi returns a boolean if a field has been set.
func (o *ApRadioStats) HasUtilUnknownWifi() bool {
	if o != nil && !IsNil(o.UtilUnknownWifi) {
		return true
	}

	return false
}

// SetUtilUnknownWifi gets a reference to the given int32 and assigns it to the UtilUnknownWifi field.
func (o *ApRadioStats) SetUtilUnknownWifi(v int32) {
	o.UtilUnknownWifi = &v
}

func (o ApRadioStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApRadioStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.DynamicChainingEnalbed) {
		toSerialize["dynamic_chaining_enalbed"] = o.DynamicChainingEnalbed
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.NumClients) {
		toSerialize["num_clients"] = o.NumClients
	}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !IsNil(o.RxBytes) {
		toSerialize["rx_bytes"] = o.RxBytes
	}
	if !IsNil(o.RxPkts) {
		toSerialize["rx_pkts"] = o.RxPkts
	}
	if !IsNil(o.TxBytes) {
		toSerialize["tx_bytes"] = o.TxBytes
	}
	if !IsNil(o.TxPkts) {
		toSerialize["tx_pkts"] = o.TxPkts
	}
	if !IsNil(o.UtilAll) {
		toSerialize["util_all"] = o.UtilAll
	}
	if !IsNil(o.UtilNonWifi) {
		toSerialize["util_non_wifi"] = o.UtilNonWifi
	}
	if !IsNil(o.UtilRxInBss) {
		toSerialize["util_rx_in_bss"] = o.UtilRxInBss
	}
	if !IsNil(o.UtilRxOtherBss) {
		toSerialize["util_rx_other_bss"] = o.UtilRxOtherBss
	}
	if !IsNil(o.UtilTx) {
		toSerialize["util_tx"] = o.UtilTx
	}
	if !IsNil(o.UtilUndecodableWifi) {
		toSerialize["util_undecodable_wifi"] = o.UtilUndecodableWifi
	}
	if !IsNil(o.UtilUnknownWifi) {
		toSerialize["util_unknown_wifi"] = o.UtilUnknownWifi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApRadioStats) UnmarshalJSON(data []byte) (err error) {
	varApRadioStats := _ApRadioStats{}

	err = json.Unmarshal(data, &varApRadioStats)

	if err != nil {
		return err
	}

	*o = ApRadioStats(varApRadioStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "dynamic_chaining_enalbed")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "num_clients")
		delete(additionalProperties, "power")
		delete(additionalProperties, "rx_bytes")
		delete(additionalProperties, "rx_pkts")
		delete(additionalProperties, "tx_bytes")
		delete(additionalProperties, "tx_pkts")
		delete(additionalProperties, "util_all")
		delete(additionalProperties, "util_non_wifi")
		delete(additionalProperties, "util_rx_in_bss")
		delete(additionalProperties, "util_rx_other_bss")
		delete(additionalProperties, "util_tx")
		delete(additionalProperties, "util_undecodable_wifi")
		delete(additionalProperties, "util_unknown_wifi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApRadioStats struct {
	value *ApRadioStats
	isSet bool
}

func (v NullableApRadioStats) Get() *ApRadioStats {
	return v.value
}

func (v *NullableApRadioStats) Set(val *ApRadioStats) {
	v.value = val
	v.isSet = true
}

func (v NullableApRadioStats) IsSet() bool {
	return v.isSet
}

func (v *NullableApRadioStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApRadioStats(val *ApRadioStats) *NullableApRadioStats {
	return &NullableApRadioStats{value: val, isSet: true}
}

func (v NullableApRadioStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApRadioStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


