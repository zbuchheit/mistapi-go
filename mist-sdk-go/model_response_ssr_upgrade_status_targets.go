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
	"fmt"
)

// checks if the ResponseSsrUpgradeStatusTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSsrUpgradeStatusTargets{}

// ResponseSsrUpgradeStatusTargets struct for ResponseSsrUpgradeStatusTargets
type ResponseSsrUpgradeStatusTargets struct {
	Failed []string `json:"failed"`
	Queued []string `json:"queued"`
	Success []string `json:"success"`
	Upgrading []string `json:"upgrading"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSsrUpgradeStatusTargets ResponseSsrUpgradeStatusTargets

// NewResponseSsrUpgradeStatusTargets instantiates a new ResponseSsrUpgradeStatusTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSsrUpgradeStatusTargets(failed []string, queued []string, success []string, upgrading []string) *ResponseSsrUpgradeStatusTargets {
	this := ResponseSsrUpgradeStatusTargets{}
	this.Failed = failed
	this.Queued = queued
	this.Success = success
	this.Upgrading = upgrading
	return &this
}

// NewResponseSsrUpgradeStatusTargetsWithDefaults instantiates a new ResponseSsrUpgradeStatusTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSsrUpgradeStatusTargetsWithDefaults() *ResponseSsrUpgradeStatusTargets {
	this := ResponseSsrUpgradeStatusTargets{}
	return &this
}

// GetFailed returns the Failed field value
func (o *ResponseSsrUpgradeStatusTargets) GetFailed() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatusTargets) GetFailedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Failed, true
}

// SetFailed sets field value
func (o *ResponseSsrUpgradeStatusTargets) SetFailed(v []string) {
	o.Failed = v
}

// GetQueued returns the Queued field value
func (o *ResponseSsrUpgradeStatusTargets) GetQueued() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatusTargets) GetQueuedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queued, true
}

// SetQueued sets field value
func (o *ResponseSsrUpgradeStatusTargets) SetQueued(v []string) {
	o.Queued = v
}

// GetSuccess returns the Success field value
func (o *ResponseSsrUpgradeStatusTargets) GetSuccess() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatusTargets) GetSuccessOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Success, true
}

// SetSuccess sets field value
func (o *ResponseSsrUpgradeStatusTargets) SetSuccess(v []string) {
	o.Success = v
}

// GetUpgrading returns the Upgrading field value
func (o *ResponseSsrUpgradeStatusTargets) GetUpgrading() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Upgrading
}

// GetUpgradingOk returns a tuple with the Upgrading field value
// and a boolean to check if the value has been set.
func (o *ResponseSsrUpgradeStatusTargets) GetUpgradingOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Upgrading, true
}

// SetUpgrading sets field value
func (o *ResponseSsrUpgradeStatusTargets) SetUpgrading(v []string) {
	o.Upgrading = v
}

func (o ResponseSsrUpgradeStatusTargets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSsrUpgradeStatusTargets) ToMap() (map[string]interface{}, error) {
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

func (o *ResponseSsrUpgradeStatusTargets) UnmarshalJSON(data []byte) (err error) {
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

	varResponseSsrUpgradeStatusTargets := _ResponseSsrUpgradeStatusTargets{}

	err = json.Unmarshal(data, &varResponseSsrUpgradeStatusTargets)

	if err != nil {
		return err
	}

	*o = ResponseSsrUpgradeStatusTargets(varResponseSsrUpgradeStatusTargets)

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

type NullableResponseSsrUpgradeStatusTargets struct {
	value *ResponseSsrUpgradeStatusTargets
	isSet bool
}

func (v NullableResponseSsrUpgradeStatusTargets) Get() *ResponseSsrUpgradeStatusTargets {
	return v.value
}

func (v *NullableResponseSsrUpgradeStatusTargets) Set(val *ResponseSsrUpgradeStatusTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSsrUpgradeStatusTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSsrUpgradeStatusTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSsrUpgradeStatusTargets(val *ResponseSsrUpgradeStatusTargets) *NullableResponseSsrUpgradeStatusTargets {
	return &NullableResponseSsrUpgradeStatusTargets{value: val, isSet: true}
}

func (v NullableResponseSsrUpgradeStatusTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSsrUpgradeStatusTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


