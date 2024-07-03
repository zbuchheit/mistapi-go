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

// checks if the UtilsShowArp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsShowArp{}

// UtilsShowArp struct for UtilsShowArp
type UtilsShowArp struct {
	// duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
	Duration *int32 `json:"duration,omitempty"`
	// rate at which output will refresh
	Interval *int32 `json:"interval,omitempty"`
	// IP Address
	Ip *string `json:"ip,omitempty"`
	// device Port ID
	PortId *string `json:"port_id,omitempty"`
	// VRF Name
	Vrf *string `json:"vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsShowArp UtilsShowArp

// NewUtilsShowArp instantiates a new UtilsShowArp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsShowArp() *UtilsShowArp {
	this := UtilsShowArp{}
	var duration int32 = 0
	this.Duration = &duration
	var interval int32 = 0
	this.Interval = &interval
	return &this
}

// NewUtilsShowArpWithDefaults instantiates a new UtilsShowArp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsShowArpWithDefaults() *UtilsShowArp {
	this := UtilsShowArp{}
	var duration int32 = 0
	this.Duration = &duration
	var interval int32 = 0
	this.Interval = &interval
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UtilsShowArp) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowArp) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UtilsShowArp) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *UtilsShowArp) SetDuration(v int32) {
	o.Duration = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *UtilsShowArp) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowArp) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *UtilsShowArp) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *UtilsShowArp) SetInterval(v int32) {
	o.Interval = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *UtilsShowArp) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowArp) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *UtilsShowArp) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *UtilsShowArp) SetIp(v string) {
	o.Ip = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *UtilsShowArp) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowArp) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *UtilsShowArp) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *UtilsShowArp) SetPortId(v string) {
	o.PortId = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *UtilsShowArp) GetVrf() string {
	if o == nil || IsNil(o.Vrf) {
		var ret string
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowArp) GetVrfOk() (*string, bool) {
	if o == nil || IsNil(o.Vrf) {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *UtilsShowArp) HasVrf() bool {
	if o != nil && !IsNil(o.Vrf) {
		return true
	}

	return false
}

// SetVrf gets a reference to the given string and assigns it to the Vrf field.
func (o *UtilsShowArp) SetVrf(v string) {
	o.Vrf = &v
}

func (o UtilsShowArp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsShowArp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	if !IsNil(o.Vrf) {
		toSerialize["vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsShowArp) UnmarshalJSON(data []byte) (err error) {
	varUtilsShowArp := _UtilsShowArp{}

	err = json.Unmarshal(data, &varUtilsShowArp)

	if err != nil {
		return err
	}

	*o = UtilsShowArp(varUtilsShowArp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "port_id")
		delete(additionalProperties, "vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsShowArp struct {
	value *UtilsShowArp
	isSet bool
}

func (v NullableUtilsShowArp) Get() *UtilsShowArp {
	return v.value
}

func (v *NullableUtilsShowArp) Set(val *UtilsShowArp) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsShowArp) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsShowArp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsShowArp(val *UtilsShowArp) *NullableUtilsShowArp {
	return &NullableUtilsShowArp{value: val, isSet: true}
}

func (v NullableUtilsShowArp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsShowArp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


