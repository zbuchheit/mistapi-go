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

// checks if the ResponseDeviceRadioChannels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeviceRadioChannels{}

// ResponseDeviceRadioChannels struct for ResponseDeviceRadioChannels
type ResponseDeviceRadioChannels struct {
	Band2440mhzAllowed bool `json:"band24_40mhz_allowed"`
	// Property key is the channel width
	Band24Channels map[string][]int32 `json:"band24_channels"`
	Band24Enabled bool `json:"band24_enabled"`
	// Property key is the channel width
	Band5Channels map[string][]int32 `json:"band5_channels"`
	Band5Enabled bool `json:"band5_enabled"`
	// Property key is the channel width
	Band6Channels *map[string][]int32 `json:"band6_channels,omitempty"`
	Band6Enabled *bool `json:"band6_enabled,omitempty"`
	Certified bool `json:"certified"`
	Code int32 `json:"code"`
	DfsOk bool `json:"dfs_ok"`
	Key string `json:"key"`
	Name string `json:"name"`
	Uses string `json:"uses"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDeviceRadioChannels ResponseDeviceRadioChannels

// NewResponseDeviceRadioChannels instantiates a new ResponseDeviceRadioChannels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeviceRadioChannels(band2440mhzAllowed bool, band24Channels map[string][]int32, band24Enabled bool, band5Channels map[string][]int32, band5Enabled bool, certified bool, code int32, dfsOk bool, key string, name string, uses string) *ResponseDeviceRadioChannels {
	this := ResponseDeviceRadioChannels{}
	this.Band2440mhzAllowed = band2440mhzAllowed
	this.Band24Channels = band24Channels
	this.Band24Enabled = band24Enabled
	this.Band5Channels = band5Channels
	this.Band5Enabled = band5Enabled
	this.Certified = certified
	this.Code = code
	this.DfsOk = dfsOk
	this.Key = key
	this.Name = name
	this.Uses = uses
	return &this
}

// NewResponseDeviceRadioChannelsWithDefaults instantiates a new ResponseDeviceRadioChannels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeviceRadioChannelsWithDefaults() *ResponseDeviceRadioChannels {
	this := ResponseDeviceRadioChannels{}
	return &this
}

// GetBand2440mhzAllowed returns the Band2440mhzAllowed field value
func (o *ResponseDeviceRadioChannels) GetBand2440mhzAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Band2440mhzAllowed
}

// GetBand2440mhzAllowedOk returns a tuple with the Band2440mhzAllowed field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetBand2440mhzAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Band2440mhzAllowed, true
}

// SetBand2440mhzAllowed sets field value
func (o *ResponseDeviceRadioChannels) SetBand2440mhzAllowed(v bool) {
	o.Band2440mhzAllowed = v
}

// GetBand24Channels returns the Band24Channels field value
func (o *ResponseDeviceRadioChannels) GetBand24Channels() map[string][]int32 {
	if o == nil {
		var ret map[string][]int32
		return ret
	}

	return o.Band24Channels
}

// GetBand24ChannelsOk returns a tuple with the Band24Channels field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetBand24ChannelsOk() (*map[string][]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Band24Channels, true
}

// SetBand24Channels sets field value
func (o *ResponseDeviceRadioChannels) SetBand24Channels(v map[string][]int32) {
	o.Band24Channels = v
}

// GetBand24Enabled returns the Band24Enabled field value
func (o *ResponseDeviceRadioChannels) GetBand24Enabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Band24Enabled
}

// GetBand24EnabledOk returns a tuple with the Band24Enabled field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetBand24EnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Band24Enabled, true
}

// SetBand24Enabled sets field value
func (o *ResponseDeviceRadioChannels) SetBand24Enabled(v bool) {
	o.Band24Enabled = v
}

// GetBand5Channels returns the Band5Channels field value
func (o *ResponseDeviceRadioChannels) GetBand5Channels() map[string][]int32 {
	if o == nil {
		var ret map[string][]int32
		return ret
	}

	return o.Band5Channels
}

// GetBand5ChannelsOk returns a tuple with the Band5Channels field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetBand5ChannelsOk() (*map[string][]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Band5Channels, true
}

// SetBand5Channels sets field value
func (o *ResponseDeviceRadioChannels) SetBand5Channels(v map[string][]int32) {
	o.Band5Channels = v
}

// GetBand5Enabled returns the Band5Enabled field value
func (o *ResponseDeviceRadioChannels) GetBand5Enabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Band5Enabled
}

// GetBand5EnabledOk returns a tuple with the Band5Enabled field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetBand5EnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Band5Enabled, true
}

// SetBand5Enabled sets field value
func (o *ResponseDeviceRadioChannels) SetBand5Enabled(v bool) {
	o.Band5Enabled = v
}

// GetBand6Channels returns the Band6Channels field value if set, zero value otherwise.
func (o *ResponseDeviceRadioChannels) GetBand6Channels() map[string][]int32 {
	if o == nil || IsNil(o.Band6Channels) {
		var ret map[string][]int32
		return ret
	}
	return *o.Band6Channels
}

// GetBand6ChannelsOk returns a tuple with the Band6Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetBand6ChannelsOk() (*map[string][]int32, bool) {
	if o == nil || IsNil(o.Band6Channels) {
		return nil, false
	}
	return o.Band6Channels, true
}

// HasBand6Channels returns a boolean if a field has been set.
func (o *ResponseDeviceRadioChannels) HasBand6Channels() bool {
	if o != nil && !IsNil(o.Band6Channels) {
		return true
	}

	return false
}

// SetBand6Channels gets a reference to the given map[string][]int32 and assigns it to the Band6Channels field.
func (o *ResponseDeviceRadioChannels) SetBand6Channels(v map[string][]int32) {
	o.Band6Channels = &v
}

// GetBand6Enabled returns the Band6Enabled field value if set, zero value otherwise.
func (o *ResponseDeviceRadioChannels) GetBand6Enabled() bool {
	if o == nil || IsNil(o.Band6Enabled) {
		var ret bool
		return ret
	}
	return *o.Band6Enabled
}

// GetBand6EnabledOk returns a tuple with the Band6Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetBand6EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Band6Enabled) {
		return nil, false
	}
	return o.Band6Enabled, true
}

// HasBand6Enabled returns a boolean if a field has been set.
func (o *ResponseDeviceRadioChannels) HasBand6Enabled() bool {
	if o != nil && !IsNil(o.Band6Enabled) {
		return true
	}

	return false
}

// SetBand6Enabled gets a reference to the given bool and assigns it to the Band6Enabled field.
func (o *ResponseDeviceRadioChannels) SetBand6Enabled(v bool) {
	o.Band6Enabled = &v
}

// GetCertified returns the Certified field value
func (o *ResponseDeviceRadioChannels) GetCertified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Certified
}

// GetCertifiedOk returns a tuple with the Certified field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetCertifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certified, true
}

// SetCertified sets field value
func (o *ResponseDeviceRadioChannels) SetCertified(v bool) {
	o.Certified = v
}

// GetCode returns the Code field value
func (o *ResponseDeviceRadioChannels) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ResponseDeviceRadioChannels) SetCode(v int32) {
	o.Code = v
}

// GetDfsOk returns the DfsOk field value
func (o *ResponseDeviceRadioChannels) GetDfsOk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DfsOk
}

// GetDfsOkOk returns a tuple with the DfsOk field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetDfsOkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DfsOk, true
}

// SetDfsOk sets field value
func (o *ResponseDeviceRadioChannels) SetDfsOk(v bool) {
	o.DfsOk = v
}

// GetKey returns the Key field value
func (o *ResponseDeviceRadioChannels) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResponseDeviceRadioChannels) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ResponseDeviceRadioChannels) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseDeviceRadioChannels) SetName(v string) {
	o.Name = v
}

// GetUses returns the Uses field value
func (o *ResponseDeviceRadioChannels) GetUses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uses
}

// GetUsesOk returns a tuple with the Uses field value
// and a boolean to check if the value has been set.
func (o *ResponseDeviceRadioChannels) GetUsesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uses, true
}

// SetUses sets field value
func (o *ResponseDeviceRadioChannels) SetUses(v string) {
	o.Uses = v
}

func (o ResponseDeviceRadioChannels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeviceRadioChannels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["band24_40mhz_allowed"] = o.Band2440mhzAllowed
	toSerialize["band24_channels"] = o.Band24Channels
	toSerialize["band24_enabled"] = o.Band24Enabled
	toSerialize["band5_channels"] = o.Band5Channels
	toSerialize["band5_enabled"] = o.Band5Enabled
	if !IsNil(o.Band6Channels) {
		toSerialize["band6_channels"] = o.Band6Channels
	}
	if !IsNil(o.Band6Enabled) {
		toSerialize["band6_enabled"] = o.Band6Enabled
	}
	toSerialize["certified"] = o.Certified
	toSerialize["code"] = o.Code
	toSerialize["dfs_ok"] = o.DfsOk
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["uses"] = o.Uses

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDeviceRadioChannels) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"band24_40mhz_allowed",
		"band24_channels",
		"band24_enabled",
		"band5_channels",
		"band5_enabled",
		"certified",
		"code",
		"dfs_ok",
		"key",
		"name",
		"uses",
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

	varResponseDeviceRadioChannels := _ResponseDeviceRadioChannels{}

	err = json.Unmarshal(data, &varResponseDeviceRadioChannels)

	if err != nil {
		return err
	}

	*o = ResponseDeviceRadioChannels(varResponseDeviceRadioChannels)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band24_40mhz_allowed")
		delete(additionalProperties, "band24_channels")
		delete(additionalProperties, "band24_enabled")
		delete(additionalProperties, "band5_channels")
		delete(additionalProperties, "band5_enabled")
		delete(additionalProperties, "band6_channels")
		delete(additionalProperties, "band6_enabled")
		delete(additionalProperties, "certified")
		delete(additionalProperties, "code")
		delete(additionalProperties, "dfs_ok")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "uses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDeviceRadioChannels struct {
	value *ResponseDeviceRadioChannels
	isSet bool
}

func (v NullableResponseDeviceRadioChannels) Get() *ResponseDeviceRadioChannels {
	return v.value
}

func (v *NullableResponseDeviceRadioChannels) Set(val *ResponseDeviceRadioChannels) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeviceRadioChannels) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeviceRadioChannels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeviceRadioChannels(val *ResponseDeviceRadioChannels) *NullableResponseDeviceRadioChannels {
	return &NullableResponseDeviceRadioChannels{value: val, isSet: true}
}

func (v NullableResponseDeviceRadioChannels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeviceRadioChannels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


