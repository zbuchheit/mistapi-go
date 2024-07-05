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

// checks if the ApStatMeshUplink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatMeshUplink{}

// ApStatMeshUplink struct for ApStatMeshUplink
type ApStatMeshUplink struct {
	Band *string `json:"band,omitempty"`
	Channel *int32 `json:"channel,omitempty"`
	IdleTime *int32 `json:"idle_time,omitempty"`
	LastSeen *int32 `json:"last_seen,omitempty"`
	Proto *string `json:"proto,omitempty"`
	Rssi *int32 `json:"rssi,omitempty"`
	RxBps *int32 `json:"rx_bps,omitempty"`
	RxBytes *int32 `json:"rx_bytes,omitempty"`
	RxPackets *int32 `json:"rx_packets,omitempty"`
	RxRate *int32 `json:"rx_rate,omitempty"`
	RxRetries *int32 `json:"rx_retries,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	Snr *int32 `json:"snr,omitempty"`
	TxBps *int32 `json:"tx_bps,omitempty"`
	TxBytes *int32 `json:"tx_bytes,omitempty"`
	TxPackets *int32 `json:"tx_packets,omitempty"`
	TxRate *int32 `json:"tx_rate,omitempty"`
	TxRetries *int32 `json:"tx_retries,omitempty"`
	UplinkApId *string `json:"uplink_ap_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatMeshUplink ApStatMeshUplink

// NewApStatMeshUplink instantiates a new ApStatMeshUplink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatMeshUplink() *ApStatMeshUplink {
	this := ApStatMeshUplink{}
	return &this
}

// NewApStatMeshUplinkWithDefaults instantiates a new ApStatMeshUplink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatMeshUplinkWithDefaults() *ApStatMeshUplink {
	this := ApStatMeshUplink{}
	return &this
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetBand() string {
	if o == nil || IsNil(o.Band) {
		var ret string
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetBandOk() (*string, bool) {
	if o == nil || IsNil(o.Band) {
		return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasBand() bool {
	if o != nil && !IsNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given string and assigns it to the Band field.
func (o *ApStatMeshUplink) SetBand(v string) {
	o.Band = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *ApStatMeshUplink) SetChannel(v int32) {
	o.Channel = &v
}

// GetIdleTime returns the IdleTime field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetIdleTime() int32 {
	if o == nil || IsNil(o.IdleTime) {
		var ret int32
		return ret
	}
	return *o.IdleTime
}

// GetIdleTimeOk returns a tuple with the IdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetIdleTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.IdleTime) {
		return nil, false
	}
	return o.IdleTime, true
}

// HasIdleTime returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasIdleTime() bool {
	if o != nil && !IsNil(o.IdleTime) {
		return true
	}

	return false
}

// SetIdleTime gets a reference to the given int32 and assigns it to the IdleTime field.
func (o *ApStatMeshUplink) SetIdleTime(v int32) {
	o.IdleTime = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetLastSeen() int32 {
	if o == nil || IsNil(o.LastSeen) {
		var ret int32
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetLastSeenOk() (*int32, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given int32 and assigns it to the LastSeen field.
func (o *ApStatMeshUplink) SetLastSeen(v int32) {
	o.LastSeen = &v
}

// GetProto returns the Proto field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetProto() string {
	if o == nil || IsNil(o.Proto) {
		var ret string
		return ret
	}
	return *o.Proto
}

// GetProtoOk returns a tuple with the Proto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetProtoOk() (*string, bool) {
	if o == nil || IsNil(o.Proto) {
		return nil, false
	}
	return o.Proto, true
}

// HasProto returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasProto() bool {
	if o != nil && !IsNil(o.Proto) {
		return true
	}

	return false
}

// SetProto gets a reference to the given string and assigns it to the Proto field.
func (o *ApStatMeshUplink) SetProto(v string) {
	o.Proto = &v
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetRssi() int32 {
	if o == nil || IsNil(o.Rssi) {
		var ret int32
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetRssiOk() (*int32, bool) {
	if o == nil || IsNil(o.Rssi) {
		return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasRssi() bool {
	if o != nil && !IsNil(o.Rssi) {
		return true
	}

	return false
}

// SetRssi gets a reference to the given int32 and assigns it to the Rssi field.
func (o *ApStatMeshUplink) SetRssi(v int32) {
	o.Rssi = &v
}

// GetRxBps returns the RxBps field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetRxBps() int32 {
	if o == nil || IsNil(o.RxBps) {
		var ret int32
		return ret
	}
	return *o.RxBps
}

// GetRxBpsOk returns a tuple with the RxBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetRxBpsOk() (*int32, bool) {
	if o == nil || IsNil(o.RxBps) {
		return nil, false
	}
	return o.RxBps, true
}

// HasRxBps returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasRxBps() bool {
	if o != nil && !IsNil(o.RxBps) {
		return true
	}

	return false
}

// SetRxBps gets a reference to the given int32 and assigns it to the RxBps field.
func (o *ApStatMeshUplink) SetRxBps(v int32) {
	o.RxBps = &v
}

// GetRxBytes returns the RxBytes field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetRxBytes() int32 {
	if o == nil || IsNil(o.RxBytes) {
		var ret int32
		return ret
	}
	return *o.RxBytes
}

// GetRxBytesOk returns a tuple with the RxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetRxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.RxBytes) {
		return nil, false
	}
	return o.RxBytes, true
}

// HasRxBytes returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasRxBytes() bool {
	if o != nil && !IsNil(o.RxBytes) {
		return true
	}

	return false
}

// SetRxBytes gets a reference to the given int32 and assigns it to the RxBytes field.
func (o *ApStatMeshUplink) SetRxBytes(v int32) {
	o.RxBytes = &v
}

// GetRxPackets returns the RxPackets field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetRxPackets() int32 {
	if o == nil || IsNil(o.RxPackets) {
		var ret int32
		return ret
	}
	return *o.RxPackets
}

// GetRxPacketsOk returns a tuple with the RxPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetRxPacketsOk() (*int32, bool) {
	if o == nil || IsNil(o.RxPackets) {
		return nil, false
	}
	return o.RxPackets, true
}

// HasRxPackets returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasRxPackets() bool {
	if o != nil && !IsNil(o.RxPackets) {
		return true
	}

	return false
}

// SetRxPackets gets a reference to the given int32 and assigns it to the RxPackets field.
func (o *ApStatMeshUplink) SetRxPackets(v int32) {
	o.RxPackets = &v
}

// GetRxRate returns the RxRate field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetRxRate() int32 {
	if o == nil || IsNil(o.RxRate) {
		var ret int32
		return ret
	}
	return *o.RxRate
}

// GetRxRateOk returns a tuple with the RxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetRxRateOk() (*int32, bool) {
	if o == nil || IsNil(o.RxRate) {
		return nil, false
	}
	return o.RxRate, true
}

// HasRxRate returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasRxRate() bool {
	if o != nil && !IsNil(o.RxRate) {
		return true
	}

	return false
}

// SetRxRate gets a reference to the given int32 and assigns it to the RxRate field.
func (o *ApStatMeshUplink) SetRxRate(v int32) {
	o.RxRate = &v
}

// GetRxRetries returns the RxRetries field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetRxRetries() int32 {
	if o == nil || IsNil(o.RxRetries) {
		var ret int32
		return ret
	}
	return *o.RxRetries
}

// GetRxRetriesOk returns a tuple with the RxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetRxRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.RxRetries) {
		return nil, false
	}
	return o.RxRetries, true
}

// HasRxRetries returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasRxRetries() bool {
	if o != nil && !IsNil(o.RxRetries) {
		return true
	}

	return false
}

// SetRxRetries gets a reference to the given int32 and assigns it to the RxRetries field.
func (o *ApStatMeshUplink) SetRxRetries(v int32) {
	o.RxRetries = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *ApStatMeshUplink) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSnr returns the Snr field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetSnr() int32 {
	if o == nil || IsNil(o.Snr) {
		var ret int32
		return ret
	}
	return *o.Snr
}

// GetSnrOk returns a tuple with the Snr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetSnrOk() (*int32, bool) {
	if o == nil || IsNil(o.Snr) {
		return nil, false
	}
	return o.Snr, true
}

// HasSnr returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasSnr() bool {
	if o != nil && !IsNil(o.Snr) {
		return true
	}

	return false
}

// SetSnr gets a reference to the given int32 and assigns it to the Snr field.
func (o *ApStatMeshUplink) SetSnr(v int32) {
	o.Snr = &v
}

// GetTxBps returns the TxBps field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetTxBps() int32 {
	if o == nil || IsNil(o.TxBps) {
		var ret int32
		return ret
	}
	return *o.TxBps
}

// GetTxBpsOk returns a tuple with the TxBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetTxBpsOk() (*int32, bool) {
	if o == nil || IsNil(o.TxBps) {
		return nil, false
	}
	return o.TxBps, true
}

// HasTxBps returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasTxBps() bool {
	if o != nil && !IsNil(o.TxBps) {
		return true
	}

	return false
}

// SetTxBps gets a reference to the given int32 and assigns it to the TxBps field.
func (o *ApStatMeshUplink) SetTxBps(v int32) {
	o.TxBps = &v
}

// GetTxBytes returns the TxBytes field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetTxBytes() int32 {
	if o == nil || IsNil(o.TxBytes) {
		var ret int32
		return ret
	}
	return *o.TxBytes
}

// GetTxBytesOk returns a tuple with the TxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetTxBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.TxBytes) {
		return nil, false
	}
	return o.TxBytes, true
}

// HasTxBytes returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasTxBytes() bool {
	if o != nil && !IsNil(o.TxBytes) {
		return true
	}

	return false
}

// SetTxBytes gets a reference to the given int32 and assigns it to the TxBytes field.
func (o *ApStatMeshUplink) SetTxBytes(v int32) {
	o.TxBytes = &v
}

// GetTxPackets returns the TxPackets field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetTxPackets() int32 {
	if o == nil || IsNil(o.TxPackets) {
		var ret int32
		return ret
	}
	return *o.TxPackets
}

// GetTxPacketsOk returns a tuple with the TxPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetTxPacketsOk() (*int32, bool) {
	if o == nil || IsNil(o.TxPackets) {
		return nil, false
	}
	return o.TxPackets, true
}

// HasTxPackets returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasTxPackets() bool {
	if o != nil && !IsNil(o.TxPackets) {
		return true
	}

	return false
}

// SetTxPackets gets a reference to the given int32 and assigns it to the TxPackets field.
func (o *ApStatMeshUplink) SetTxPackets(v int32) {
	o.TxPackets = &v
}

// GetTxRate returns the TxRate field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetTxRate() int32 {
	if o == nil || IsNil(o.TxRate) {
		var ret int32
		return ret
	}
	return *o.TxRate
}

// GetTxRateOk returns a tuple with the TxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetTxRateOk() (*int32, bool) {
	if o == nil || IsNil(o.TxRate) {
		return nil, false
	}
	return o.TxRate, true
}

// HasTxRate returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasTxRate() bool {
	if o != nil && !IsNil(o.TxRate) {
		return true
	}

	return false
}

// SetTxRate gets a reference to the given int32 and assigns it to the TxRate field.
func (o *ApStatMeshUplink) SetTxRate(v int32) {
	o.TxRate = &v
}

// GetTxRetries returns the TxRetries field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetTxRetries() int32 {
	if o == nil || IsNil(o.TxRetries) {
		var ret int32
		return ret
	}
	return *o.TxRetries
}

// GetTxRetriesOk returns a tuple with the TxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetTxRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.TxRetries) {
		return nil, false
	}
	return o.TxRetries, true
}

// HasTxRetries returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasTxRetries() bool {
	if o != nil && !IsNil(o.TxRetries) {
		return true
	}

	return false
}

// SetTxRetries gets a reference to the given int32 and assigns it to the TxRetries field.
func (o *ApStatMeshUplink) SetTxRetries(v int32) {
	o.TxRetries = &v
}

// GetUplinkApId returns the UplinkApId field value if set, zero value otherwise.
func (o *ApStatMeshUplink) GetUplinkApId() string {
	if o == nil || IsNil(o.UplinkApId) {
		var ret string
		return ret
	}
	return *o.UplinkApId
}

// GetUplinkApIdOk returns a tuple with the UplinkApId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatMeshUplink) GetUplinkApIdOk() (*string, bool) {
	if o == nil || IsNil(o.UplinkApId) {
		return nil, false
	}
	return o.UplinkApId, true
}

// HasUplinkApId returns a boolean if a field has been set.
func (o *ApStatMeshUplink) HasUplinkApId() bool {
	if o != nil && !IsNil(o.UplinkApId) {
		return true
	}

	return false
}

// SetUplinkApId gets a reference to the given string and assigns it to the UplinkApId field.
func (o *ApStatMeshUplink) SetUplinkApId(v string) {
	o.UplinkApId = &v
}

func (o ApStatMeshUplink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatMeshUplink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.IdleTime) {
		toSerialize["idle_time"] = o.IdleTime
	}
	if !IsNil(o.LastSeen) {
		toSerialize["last_seen"] = o.LastSeen
	}
	if !IsNil(o.Proto) {
		toSerialize["proto"] = o.Proto
	}
	if !IsNil(o.Rssi) {
		toSerialize["rssi"] = o.Rssi
	}
	if !IsNil(o.RxBps) {
		toSerialize["rx_bps"] = o.RxBps
	}
	if !IsNil(o.RxBytes) {
		toSerialize["rx_bytes"] = o.RxBytes
	}
	if !IsNil(o.RxPackets) {
		toSerialize["rx_packets"] = o.RxPackets
	}
	if !IsNil(o.RxRate) {
		toSerialize["rx_rate"] = o.RxRate
	}
	if !IsNil(o.RxRetries) {
		toSerialize["rx_retries"] = o.RxRetries
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Snr) {
		toSerialize["snr"] = o.Snr
	}
	if !IsNil(o.TxBps) {
		toSerialize["tx_bps"] = o.TxBps
	}
	if !IsNil(o.TxBytes) {
		toSerialize["tx_bytes"] = o.TxBytes
	}
	if !IsNil(o.TxPackets) {
		toSerialize["tx_packets"] = o.TxPackets
	}
	if !IsNil(o.TxRate) {
		toSerialize["tx_rate"] = o.TxRate
	}
	if !IsNil(o.TxRetries) {
		toSerialize["tx_retries"] = o.TxRetries
	}
	if !IsNil(o.UplinkApId) {
		toSerialize["uplink_ap_id"] = o.UplinkApId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApStatMeshUplink) UnmarshalJSON(data []byte) (err error) {
	varApStatMeshUplink := _ApStatMeshUplink{}

	err = json.Unmarshal(data, &varApStatMeshUplink)

	if err != nil {
		return err
	}

	*o = ApStatMeshUplink(varApStatMeshUplink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "idle_time")
		delete(additionalProperties, "last_seen")
		delete(additionalProperties, "proto")
		delete(additionalProperties, "rssi")
		delete(additionalProperties, "rx_bps")
		delete(additionalProperties, "rx_bytes")
		delete(additionalProperties, "rx_packets")
		delete(additionalProperties, "rx_rate")
		delete(additionalProperties, "rx_retries")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "snr")
		delete(additionalProperties, "tx_bps")
		delete(additionalProperties, "tx_bytes")
		delete(additionalProperties, "tx_packets")
		delete(additionalProperties, "tx_rate")
		delete(additionalProperties, "tx_retries")
		delete(additionalProperties, "uplink_ap_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatMeshUplink struct {
	value *ApStatMeshUplink
	isSet bool
}

func (v NullableApStatMeshUplink) Get() *ApStatMeshUplink {
	return v.value
}

func (v *NullableApStatMeshUplink) Set(val *ApStatMeshUplink) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatMeshUplink) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatMeshUplink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatMeshUplink(val *ApStatMeshUplink) *NullableApStatMeshUplink {
	return &NullableApStatMeshUplink{value: val, isSet: true}
}

func (v NullableApStatMeshUplink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatMeshUplink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


