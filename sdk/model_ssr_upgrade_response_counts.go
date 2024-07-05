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
	"fmt"
)

// checks if the SsrUpgradeResponseCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsrUpgradeResponseCounts{}

// SsrUpgradeResponseCounts struct for SsrUpgradeResponseCounts
type SsrUpgradeResponseCounts struct {
	Failed int32 `json:"failed"`
	Queued int32 `json:"queued"`
	Success int32 `json:"success"`
	Upgrading int32 `json:"upgrading"`
	AdditionalProperties map[string]interface{}
}

type _SsrUpgradeResponseCounts SsrUpgradeResponseCounts

// NewSsrUpgradeResponseCounts instantiates a new SsrUpgradeResponseCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsrUpgradeResponseCounts(failed int32, queued int32, success int32, upgrading int32) *SsrUpgradeResponseCounts {
	this := SsrUpgradeResponseCounts{}
	this.Failed = failed
	this.Queued = queued
	this.Success = success
	this.Upgrading = upgrading
	return &this
}

// NewSsrUpgradeResponseCountsWithDefaults instantiates a new SsrUpgradeResponseCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsrUpgradeResponseCountsWithDefaults() *SsrUpgradeResponseCounts {
	this := SsrUpgradeResponseCounts{}
	return &this
}

// GetFailed returns the Failed field value
func (o *SsrUpgradeResponseCounts) GetFailed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponseCounts) GetFailedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failed, true
}

// SetFailed sets field value
func (o *SsrUpgradeResponseCounts) SetFailed(v int32) {
	o.Failed = v
}

// GetQueued returns the Queued field value
func (o *SsrUpgradeResponseCounts) GetQueued() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponseCounts) GetQueuedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queued, true
}

// SetQueued sets field value
func (o *SsrUpgradeResponseCounts) SetQueued(v int32) {
	o.Queued = v
}

// GetSuccess returns the Success field value
func (o *SsrUpgradeResponseCounts) GetSuccess() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponseCounts) GetSuccessOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *SsrUpgradeResponseCounts) SetSuccess(v int32) {
	o.Success = v
}

// GetUpgrading returns the Upgrading field value
func (o *SsrUpgradeResponseCounts) GetUpgrading() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Upgrading
}

// GetUpgradingOk returns a tuple with the Upgrading field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponseCounts) GetUpgradingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upgrading, true
}

// SetUpgrading sets field value
func (o *SsrUpgradeResponseCounts) SetUpgrading(v int32) {
	o.Upgrading = v
}

func (o SsrUpgradeResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsrUpgradeResponseCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["failed"] = o.Failed
	toSerialize["queued"] = o.Queued
	toSerialize["success"] = o.Success
	toSerialize["upgrading"] = o.Upgrading

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsrUpgradeResponseCounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"failed",
		"queued",
		"success",
		"upgrading",
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

	varSsrUpgradeResponseCounts := _SsrUpgradeResponseCounts{}

	err = json.Unmarshal(data, &varSsrUpgradeResponseCounts)

	if err != nil {
		return err
	}

	*o = SsrUpgradeResponseCounts(varSsrUpgradeResponseCounts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "failed")
		delete(additionalProperties, "queued")
		delete(additionalProperties, "success")
		delete(additionalProperties, "upgrading")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSsrUpgradeResponseCounts struct {
	value *SsrUpgradeResponseCounts
	isSet bool
}

func (v NullableSsrUpgradeResponseCounts) Get() *SsrUpgradeResponseCounts {
	return v.value
}

func (v *NullableSsrUpgradeResponseCounts) Set(val *SsrUpgradeResponseCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableSsrUpgradeResponseCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableSsrUpgradeResponseCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsrUpgradeResponseCounts(val *SsrUpgradeResponseCounts) *NullableSsrUpgradeResponseCounts {
	return &NullableSsrUpgradeResponseCounts{value: val, isSet: true}
}

func (v NullableSsrUpgradeResponseCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsrUpgradeResponseCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


