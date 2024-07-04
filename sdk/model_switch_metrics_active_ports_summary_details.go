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
)

// checks if the SwitchMetricsActivePortsSummaryDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchMetricsActivePortsSummaryDetails{}

// SwitchMetricsActivePortsSummaryDetails struct for SwitchMetricsActivePortsSummaryDetails
type SwitchMetricsActivePortsSummaryDetails struct {
	ActivePortCount *int32 `json:"active_port_count,omitempty"`
	TotalPortCount *int32 `json:"total_port_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwitchMetricsActivePortsSummaryDetails SwitchMetricsActivePortsSummaryDetails

// NewSwitchMetricsActivePortsSummaryDetails instantiates a new SwitchMetricsActivePortsSummaryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchMetricsActivePortsSummaryDetails() *SwitchMetricsActivePortsSummaryDetails {
	this := SwitchMetricsActivePortsSummaryDetails{}
	return &this
}

// NewSwitchMetricsActivePortsSummaryDetailsWithDefaults instantiates a new SwitchMetricsActivePortsSummaryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchMetricsActivePortsSummaryDetailsWithDefaults() *SwitchMetricsActivePortsSummaryDetails {
	this := SwitchMetricsActivePortsSummaryDetails{}
	return &this
}

// GetActivePortCount returns the ActivePortCount field value if set, zero value otherwise.
func (o *SwitchMetricsActivePortsSummaryDetails) GetActivePortCount() int32 {
	if o == nil || IsNil(o.ActivePortCount) {
		var ret int32
		return ret
	}
	return *o.ActivePortCount
}

// GetActivePortCountOk returns a tuple with the ActivePortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMetricsActivePortsSummaryDetails) GetActivePortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActivePortCount) {
		return nil, false
	}
	return o.ActivePortCount, true
}

// HasActivePortCount returns a boolean if a field has been set.
func (o *SwitchMetricsActivePortsSummaryDetails) HasActivePortCount() bool {
	if o != nil && !IsNil(o.ActivePortCount) {
		return true
	}

	return false
}

// SetActivePortCount gets a reference to the given int32 and assigns it to the ActivePortCount field.
func (o *SwitchMetricsActivePortsSummaryDetails) SetActivePortCount(v int32) {
	o.ActivePortCount = &v
}

// GetTotalPortCount returns the TotalPortCount field value if set, zero value otherwise.
func (o *SwitchMetricsActivePortsSummaryDetails) GetTotalPortCount() int32 {
	if o == nil || IsNil(o.TotalPortCount) {
		var ret int32
		return ret
	}
	return *o.TotalPortCount
}

// GetTotalPortCountOk returns a tuple with the TotalPortCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMetricsActivePortsSummaryDetails) GetTotalPortCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPortCount) {
		return nil, false
	}
	return o.TotalPortCount, true
}

// HasTotalPortCount returns a boolean if a field has been set.
func (o *SwitchMetricsActivePortsSummaryDetails) HasTotalPortCount() bool {
	if o != nil && !IsNil(o.TotalPortCount) {
		return true
	}

	return false
}

// SetTotalPortCount gets a reference to the given int32 and assigns it to the TotalPortCount field.
func (o *SwitchMetricsActivePortsSummaryDetails) SetTotalPortCount(v int32) {
	o.TotalPortCount = &v
}

func (o SwitchMetricsActivePortsSummaryDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchMetricsActivePortsSummaryDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivePortCount) {
		toSerialize["active_port_count"] = o.ActivePortCount
	}
	if !IsNil(o.TotalPortCount) {
		toSerialize["total_port_count"] = o.TotalPortCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwitchMetricsActivePortsSummaryDetails) UnmarshalJSON(data []byte) (err error) {
	varSwitchMetricsActivePortsSummaryDetails := _SwitchMetricsActivePortsSummaryDetails{}

	err = json.Unmarshal(data, &varSwitchMetricsActivePortsSummaryDetails)

	if err != nil {
		return err
	}

	*o = SwitchMetricsActivePortsSummaryDetails(varSwitchMetricsActivePortsSummaryDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_port_count")
		delete(additionalProperties, "total_port_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwitchMetricsActivePortsSummaryDetails struct {
	value *SwitchMetricsActivePortsSummaryDetails
	isSet bool
}

func (v NullableSwitchMetricsActivePortsSummaryDetails) Get() *SwitchMetricsActivePortsSummaryDetails {
	return v.value
}

func (v *NullableSwitchMetricsActivePortsSummaryDetails) Set(val *SwitchMetricsActivePortsSummaryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchMetricsActivePortsSummaryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchMetricsActivePortsSummaryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchMetricsActivePortsSummaryDetails(val *SwitchMetricsActivePortsSummaryDetails) *NullableSwitchMetricsActivePortsSummaryDetails {
	return &NullableSwitchMetricsActivePortsSummaryDetails{value: val, isSet: true}
}

func (v NullableSwitchMetricsActivePortsSummaryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchMetricsActivePortsSummaryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


