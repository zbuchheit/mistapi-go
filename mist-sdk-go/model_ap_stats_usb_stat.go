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

// checks if the ApStatsUsbStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatsUsbStat{}

// ApStatsUsbStat struct for ApStatsUsbStat
type ApStatsUsbStat struct {
	Channel *int32 `json:"channel,omitempty"`
	Connected *bool `json:"connected,omitempty"`
	LastActivity *int32 `json:"last_activity,omitempty"`
	Type *string `json:"type,omitempty"`
	Up *bool `json:"up,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatsUsbStat ApStatsUsbStat

// NewApStatsUsbStat instantiates a new ApStatsUsbStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatsUsbStat() *ApStatsUsbStat {
	this := ApStatsUsbStat{}
	return &this
}

// NewApStatsUsbStatWithDefaults instantiates a new ApStatsUsbStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatsUsbStatWithDefaults() *ApStatsUsbStat {
	this := ApStatsUsbStat{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ApStatsUsbStat) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsUsbStat) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ApStatsUsbStat) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *ApStatsUsbStat) SetChannel(v int32) {
	o.Channel = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *ApStatsUsbStat) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsUsbStat) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *ApStatsUsbStat) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *ApStatsUsbStat) SetConnected(v bool) {
	o.Connected = &v
}

// GetLastActivity returns the LastActivity field value if set, zero value otherwise.
func (o *ApStatsUsbStat) GetLastActivity() int32 {
	if o == nil || IsNil(o.LastActivity) {
		var ret int32
		return ret
	}
	return *o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsUsbStat) GetLastActivityOk() (*int32, bool) {
	if o == nil || IsNil(o.LastActivity) {
		return nil, false
	}
	return o.LastActivity, true
}

// HasLastActivity returns a boolean if a field has been set.
func (o *ApStatsUsbStat) HasLastActivity() bool {
	if o != nil && !IsNil(o.LastActivity) {
		return true
	}

	return false
}

// SetLastActivity gets a reference to the given int32 and assigns it to the LastActivity field.
func (o *ApStatsUsbStat) SetLastActivity(v int32) {
	o.LastActivity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApStatsUsbStat) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsUsbStat) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApStatsUsbStat) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApStatsUsbStat) SetType(v string) {
	o.Type = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *ApStatsUsbStat) GetUp() bool {
	if o == nil || IsNil(o.Up) {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsUsbStat) GetUpOk() (*bool, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *ApStatsUsbStat) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *ApStatsUsbStat) SetUp(v bool) {
	o.Up = &v
}

func (o ApStatsUsbStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatsUsbStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.LastActivity) {
		toSerialize["last_activity"] = o.LastActivity
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApStatsUsbStat) UnmarshalJSON(data []byte) (err error) {
	varApStatsUsbStat := _ApStatsUsbStat{}

	err = json.Unmarshal(data, &varApStatsUsbStat)

	if err != nil {
		return err
	}

	*o = ApStatsUsbStat(varApStatsUsbStat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "connected")
		delete(additionalProperties, "last_activity")
		delete(additionalProperties, "type")
		delete(additionalProperties, "up")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatsUsbStat struct {
	value *ApStatsUsbStat
	isSet bool
}

func (v NullableApStatsUsbStat) Get() *ApStatsUsbStat {
	return v.value
}

func (v *NullableApStatsUsbStat) Set(val *ApStatsUsbStat) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsUsbStat) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsUsbStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsUsbStat(val *ApStatsUsbStat) *NullableApStatsUsbStat {
	return &NullableApStatsUsbStat{value: val, isSet: true}
}

func (v NullableApStatsUsbStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsUsbStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


