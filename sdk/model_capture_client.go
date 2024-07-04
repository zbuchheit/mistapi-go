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
	"fmt"
)

// checks if the CaptureClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureClient{}

// CaptureClient Initiate a Client Packet Capture
type CaptureClient struct {
	ApMac NullableString `json:"ap_mac,omitempty"`
	// client mac, required if `type`==`client`; optional otherwise
	ClientMac NullableString `json:"client_mac,omitempty"`
	// duration of the capture, in seconds
	Duration NullableInt32 `json:"duration,omitempty"`
	IncludesMcast *bool `json:"includes_mcast,omitempty"`
	MaxPktLen NullableInt32 `json:"max_pkt_len,omitempty"`
	// number of packets to capture, 0 for unlimited, default is 1024 for client-capture
	NumPackets NullableInt32 `json:"num_packets,omitempty"`
	// optional filter by ssid
	Ssid NullableString `json:"ssid,omitempty"`
	Type CaptureClientType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _CaptureClient CaptureClient

// NewCaptureClient instantiates a new CaptureClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureClient(type_ CaptureClientType) *CaptureClient {
	this := CaptureClient{}
	var duration int32 = 600
	this.Duration = *NewNullableInt32(&duration)
	var includesMcast bool = false
	this.IncludesMcast = &includesMcast
	var maxPktLen int32 = 128
	this.MaxPktLen = *NewNullableInt32(&maxPktLen)
	var numPackets int32 = 1024
	this.NumPackets = *NewNullableInt32(&numPackets)
	this.Type = type_
	return &this
}

// NewCaptureClientWithDefaults instantiates a new CaptureClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureClientWithDefaults() *CaptureClient {
	this := CaptureClient{}
	var duration int32 = 600
	this.Duration = *NewNullableInt32(&duration)
	var includesMcast bool = false
	this.IncludesMcast = &includesMcast
	var maxPktLen int32 = 128
	this.MaxPktLen = *NewNullableInt32(&maxPktLen)
	var numPackets int32 = 1024
	this.NumPackets = *NewNullableInt32(&numPackets)
	return &this
}

// GetApMac returns the ApMac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureClient) GetApMac() string {
	if o == nil || IsNil(o.ApMac.Get()) {
		var ret string
		return ret
	}
	return *o.ApMac.Get()
}

// GetApMacOk returns a tuple with the ApMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureClient) GetApMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApMac.Get(), o.ApMac.IsSet()
}

// HasApMac returns a boolean if a field has been set.
func (o *CaptureClient) HasApMac() bool {
	if o != nil && o.ApMac.IsSet() {
		return true
	}

	return false
}

// SetApMac gets a reference to the given NullableString and assigns it to the ApMac field.
func (o *CaptureClient) SetApMac(v string) {
	o.ApMac.Set(&v)
}
// SetApMacNil sets the value for ApMac to be an explicit nil
func (o *CaptureClient) SetApMacNil() {
	o.ApMac.Set(nil)
}

// UnsetApMac ensures that no value is present for ApMac, not even an explicit nil
func (o *CaptureClient) UnsetApMac() {
	o.ApMac.Unset()
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureClient) GetClientMac() string {
	if o == nil || IsNil(o.ClientMac.Get()) {
		var ret string
		return ret
	}
	return *o.ClientMac.Get()
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureClient) GetClientMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientMac.Get(), o.ClientMac.IsSet()
}

// HasClientMac returns a boolean if a field has been set.
func (o *CaptureClient) HasClientMac() bool {
	if o != nil && o.ClientMac.IsSet() {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given NullableString and assigns it to the ClientMac field.
func (o *CaptureClient) SetClientMac(v string) {
	o.ClientMac.Set(&v)
}
// SetClientMacNil sets the value for ClientMac to be an explicit nil
func (o *CaptureClient) SetClientMacNil() {
	o.ClientMac.Set(nil)
}

// UnsetClientMac ensures that no value is present for ClientMac, not even an explicit nil
func (o *CaptureClient) UnsetClientMac() {
	o.ClientMac.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureClient) GetDuration() int32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureClient) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *CaptureClient) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *CaptureClient) SetDuration(v int32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *CaptureClient) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *CaptureClient) UnsetDuration() {
	o.Duration.Unset()
}

// GetIncludesMcast returns the IncludesMcast field value if set, zero value otherwise.
func (o *CaptureClient) GetIncludesMcast() bool {
	if o == nil || IsNil(o.IncludesMcast) {
		var ret bool
		return ret
	}
	return *o.IncludesMcast
}

// GetIncludesMcastOk returns a tuple with the IncludesMcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureClient) GetIncludesMcastOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludesMcast) {
		return nil, false
	}
	return o.IncludesMcast, true
}

// HasIncludesMcast returns a boolean if a field has been set.
func (o *CaptureClient) HasIncludesMcast() bool {
	if o != nil && !IsNil(o.IncludesMcast) {
		return true
	}

	return false
}

// SetIncludesMcast gets a reference to the given bool and assigns it to the IncludesMcast field.
func (o *CaptureClient) SetIncludesMcast(v bool) {
	o.IncludesMcast = &v
}

// GetMaxPktLen returns the MaxPktLen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureClient) GetMaxPktLen() int32 {
	if o == nil || IsNil(o.MaxPktLen.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPktLen.Get()
}

// GetMaxPktLenOk returns a tuple with the MaxPktLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureClient) GetMaxPktLenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPktLen.Get(), o.MaxPktLen.IsSet()
}

// HasMaxPktLen returns a boolean if a field has been set.
func (o *CaptureClient) HasMaxPktLen() bool {
	if o != nil && o.MaxPktLen.IsSet() {
		return true
	}

	return false
}

// SetMaxPktLen gets a reference to the given NullableInt32 and assigns it to the MaxPktLen field.
func (o *CaptureClient) SetMaxPktLen(v int32) {
	o.MaxPktLen.Set(&v)
}
// SetMaxPktLenNil sets the value for MaxPktLen to be an explicit nil
func (o *CaptureClient) SetMaxPktLenNil() {
	o.MaxPktLen.Set(nil)
}

// UnsetMaxPktLen ensures that no value is present for MaxPktLen, not even an explicit nil
func (o *CaptureClient) UnsetMaxPktLen() {
	o.MaxPktLen.Unset()
}

// GetNumPackets returns the NumPackets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureClient) GetNumPackets() int32 {
	if o == nil || IsNil(o.NumPackets.Get()) {
		var ret int32
		return ret
	}
	return *o.NumPackets.Get()
}

// GetNumPacketsOk returns a tuple with the NumPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureClient) GetNumPacketsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumPackets.Get(), o.NumPackets.IsSet()
}

// HasNumPackets returns a boolean if a field has been set.
func (o *CaptureClient) HasNumPackets() bool {
	if o != nil && o.NumPackets.IsSet() {
		return true
	}

	return false
}

// SetNumPackets gets a reference to the given NullableInt32 and assigns it to the NumPackets field.
func (o *CaptureClient) SetNumPackets(v int32) {
	o.NumPackets.Set(&v)
}
// SetNumPacketsNil sets the value for NumPackets to be an explicit nil
func (o *CaptureClient) SetNumPacketsNil() {
	o.NumPackets.Set(nil)
}

// UnsetNumPackets ensures that no value is present for NumPackets, not even an explicit nil
func (o *CaptureClient) UnsetNumPackets() {
	o.NumPackets.Unset()
}

// GetSsid returns the Ssid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureClient) GetSsid() string {
	if o == nil || IsNil(o.Ssid.Get()) {
		var ret string
		return ret
	}
	return *o.Ssid.Get()
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureClient) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ssid.Get(), o.Ssid.IsSet()
}

// HasSsid returns a boolean if a field has been set.
func (o *CaptureClient) HasSsid() bool {
	if o != nil && o.Ssid.IsSet() {
		return true
	}

	return false
}

// SetSsid gets a reference to the given NullableString and assigns it to the Ssid field.
func (o *CaptureClient) SetSsid(v string) {
	o.Ssid.Set(&v)
}
// SetSsidNil sets the value for Ssid to be an explicit nil
func (o *CaptureClient) SetSsidNil() {
	o.Ssid.Set(nil)
}

// UnsetSsid ensures that no value is present for Ssid, not even an explicit nil
func (o *CaptureClient) UnsetSsid() {
	o.Ssid.Unset()
}

// GetType returns the Type field value
func (o *CaptureClient) GetType() CaptureClientType {
	if o == nil {
		var ret CaptureClientType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaptureClient) GetTypeOk() (*CaptureClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureClient) SetType(v CaptureClientType) {
	o.Type = v
}

func (o CaptureClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApMac.IsSet() {
		toSerialize["ap_mac"] = o.ApMac.Get()
	}
	if o.ClientMac.IsSet() {
		toSerialize["client_mac"] = o.ClientMac.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if !IsNil(o.IncludesMcast) {
		toSerialize["includes_mcast"] = o.IncludesMcast
	}
	if o.MaxPktLen.IsSet() {
		toSerialize["max_pkt_len"] = o.MaxPktLen.Get()
	}
	if o.NumPackets.IsSet() {
		toSerialize["num_packets"] = o.NumPackets.Get()
	}
	if o.Ssid.IsSet() {
		toSerialize["ssid"] = o.Ssid.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaptureClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCaptureClient := _CaptureClient{}

	err = json.Unmarshal(data, &varCaptureClient)

	if err != nil {
		return err
	}

	*o = CaptureClient(varCaptureClient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_mac")
		delete(additionalProperties, "client_mac")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "includes_mcast")
		delete(additionalProperties, "max_pkt_len")
		delete(additionalProperties, "num_packets")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaptureClient struct {
	value *CaptureClient
	isSet bool
}

func (v NullableCaptureClient) Get() *CaptureClient {
	return v.value
}

func (v *NullableCaptureClient) Set(val *CaptureClient) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureClient) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureClient(val *CaptureClient) *NullableCaptureClient {
	return &NullableCaptureClient{value: val, isSet: true}
}

func (v NullableCaptureClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


