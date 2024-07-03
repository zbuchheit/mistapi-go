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

// checks if the WlanAppQos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanAppQos{}

// WlanAppQos app qos wlan settings
type WlanAppQos struct {
	Apps *map[string]WlanAppQosAppsProperties `json:"apps,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Others []WlanAppQosOthersItem `json:"others,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanAppQos WlanAppQos

// NewWlanAppQos instantiates a new WlanAppQos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanAppQos() *WlanAppQos {
	this := WlanAppQos{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewWlanAppQosWithDefaults instantiates a new WlanAppQos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanAppQosWithDefaults() *WlanAppQos {
	this := WlanAppQos{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *WlanAppQos) GetApps() map[string]WlanAppQosAppsProperties {
	if o == nil || IsNil(o.Apps) {
		var ret map[string]WlanAppQosAppsProperties
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppQos) GetAppsOk() (*map[string]WlanAppQosAppsProperties, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *WlanAppQos) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given map[string]WlanAppQosAppsProperties and assigns it to the Apps field.
func (o *WlanAppQos) SetApps(v map[string]WlanAppQosAppsProperties) {
	o.Apps = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WlanAppQos) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppQos) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WlanAppQos) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WlanAppQos) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOthers returns the Others field value if set, zero value otherwise.
func (o *WlanAppQos) GetOthers() []WlanAppQosOthersItem {
	if o == nil || IsNil(o.Others) {
		var ret []WlanAppQosOthersItem
		return ret
	}
	return o.Others
}

// GetOthersOk returns a tuple with the Others field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanAppQos) GetOthersOk() ([]WlanAppQosOthersItem, bool) {
	if o == nil || IsNil(o.Others) {
		return nil, false
	}
	return o.Others, true
}

// HasOthers returns a boolean if a field has been set.
func (o *WlanAppQos) HasOthers() bool {
	if o != nil && !IsNil(o.Others) {
		return true
	}

	return false
}

// SetOthers gets a reference to the given []WlanAppQosOthersItem and assigns it to the Others field.
func (o *WlanAppQos) SetOthers(v []WlanAppQosOthersItem) {
	o.Others = v
}

func (o WlanAppQos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanAppQos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Others) {
		toSerialize["others"] = o.Others
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanAppQos) UnmarshalJSON(data []byte) (err error) {
	varWlanAppQos := _WlanAppQos{}

	err = json.Unmarshal(data, &varWlanAppQos)

	if err != nil {
		return err
	}

	*o = WlanAppQos(varWlanAppQos)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apps")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "others")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanAppQos struct {
	value *WlanAppQos
	isSet bool
}

func (v NullableWlanAppQos) Get() *WlanAppQos {
	return v.value
}

func (v *NullableWlanAppQos) Set(val *WlanAppQos) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanAppQos) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanAppQos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanAppQos(val *WlanAppQos) *NullableWlanAppQos {
	return &NullableWlanAppQos{value: val, isSet: true}
}

func (v NullableWlanAppQos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanAppQos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


