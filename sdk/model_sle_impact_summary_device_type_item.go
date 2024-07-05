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

// checks if the SleImpactSummaryDeviceTypeItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleImpactSummaryDeviceTypeItem{}

// SleImpactSummaryDeviceTypeItem struct for SleImpactSummaryDeviceTypeItem
type SleImpactSummaryDeviceTypeItem struct {
	Degraded float32 `json:"degraded"`
	DeviceType string `json:"device_type"`
	Duration float32 `json:"duration"`
	Name string `json:"name"`
	Total float32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _SleImpactSummaryDeviceTypeItem SleImpactSummaryDeviceTypeItem

// NewSleImpactSummaryDeviceTypeItem instantiates a new SleImpactSummaryDeviceTypeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleImpactSummaryDeviceTypeItem(degraded float32, deviceType string, duration float32, name string, total float32) *SleImpactSummaryDeviceTypeItem {
	this := SleImpactSummaryDeviceTypeItem{}
	this.Degraded = degraded
	this.DeviceType = deviceType
	this.Duration = duration
	this.Name = name
	this.Total = total
	return &this
}

// NewSleImpactSummaryDeviceTypeItemWithDefaults instantiates a new SleImpactSummaryDeviceTypeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleImpactSummaryDeviceTypeItemWithDefaults() *SleImpactSummaryDeviceTypeItem {
	this := SleImpactSummaryDeviceTypeItem{}
	return &this
}

// GetDegraded returns the Degraded field value
func (o *SleImpactSummaryDeviceTypeItem) GetDegraded() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Degraded
}

// GetDegradedOk returns a tuple with the Degraded field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryDeviceTypeItem) GetDegradedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Degraded, true
}

// SetDegraded sets field value
func (o *SleImpactSummaryDeviceTypeItem) SetDegraded(v float32) {
	o.Degraded = v
}

// GetDeviceType returns the DeviceType field value
func (o *SleImpactSummaryDeviceTypeItem) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryDeviceTypeItem) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *SleImpactSummaryDeviceTypeItem) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetDuration returns the Duration field value
func (o *SleImpactSummaryDeviceTypeItem) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryDeviceTypeItem) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SleImpactSummaryDeviceTypeItem) SetDuration(v float32) {
	o.Duration = v
}

// GetName returns the Name field value
func (o *SleImpactSummaryDeviceTypeItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryDeviceTypeItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SleImpactSummaryDeviceTypeItem) SetName(v string) {
	o.Name = v
}

// GetTotal returns the Total field value
func (o *SleImpactSummaryDeviceTypeItem) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SleImpactSummaryDeviceTypeItem) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SleImpactSummaryDeviceTypeItem) SetTotal(v float32) {
	o.Total = v
}

func (o SleImpactSummaryDeviceTypeItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleImpactSummaryDeviceTypeItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["degraded"] = o.Degraded
	toSerialize["device_type"] = o.DeviceType
	toSerialize["duration"] = o.Duration
	toSerialize["name"] = o.Name
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleImpactSummaryDeviceTypeItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"degraded",
		"device_type",
		"duration",
		"name",
		"total",
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

	varSleImpactSummaryDeviceTypeItem := _SleImpactSummaryDeviceTypeItem{}

	err = json.Unmarshal(data, &varSleImpactSummaryDeviceTypeItem)

	if err != nil {
		return err
	}

	*o = SleImpactSummaryDeviceTypeItem(varSleImpactSummaryDeviceTypeItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "degraded")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "name")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleImpactSummaryDeviceTypeItem struct {
	value *SleImpactSummaryDeviceTypeItem
	isSet bool
}

func (v NullableSleImpactSummaryDeviceTypeItem) Get() *SleImpactSummaryDeviceTypeItem {
	return v.value
}

func (v *NullableSleImpactSummaryDeviceTypeItem) Set(val *SleImpactSummaryDeviceTypeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSleImpactSummaryDeviceTypeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSleImpactSummaryDeviceTypeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleImpactSummaryDeviceTypeItem(val *SleImpactSummaryDeviceTypeItem) *NullableSleImpactSummaryDeviceTypeItem {
	return &NullableSleImpactSummaryDeviceTypeItem{value: val, isSet: true}
}

func (v NullableSleImpactSummaryDeviceTypeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleImpactSummaryDeviceTypeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


