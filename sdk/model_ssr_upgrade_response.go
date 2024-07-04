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

// checks if the SsrUpgradeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsrUpgradeResponse{}

// SsrUpgradeResponse struct for SsrUpgradeResponse
type SsrUpgradeResponse struct {
	Channel string `json:"channel"`
	Counts SsrUpgradeResponseCounts `json:"counts"`
	DeviceType string `json:"device_type"`
	Id string `json:"id"`
	Status string `json:"status"`
	Strategy string `json:"strategy"`
	Versions map[string]string `json:"versions"`
	AdditionalProperties map[string]interface{}
}

type _SsrUpgradeResponse SsrUpgradeResponse

// NewSsrUpgradeResponse instantiates a new SsrUpgradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsrUpgradeResponse(channel string, counts SsrUpgradeResponseCounts, deviceType string, id string, status string, strategy string, versions map[string]string) *SsrUpgradeResponse {
	this := SsrUpgradeResponse{}
	this.Channel = channel
	this.Counts = counts
	this.DeviceType = deviceType
	this.Id = id
	this.Status = status
	this.Strategy = strategy
	this.Versions = versions
	return &this
}

// NewSsrUpgradeResponseWithDefaults instantiates a new SsrUpgradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsrUpgradeResponseWithDefaults() *SsrUpgradeResponse {
	this := SsrUpgradeResponse{}
	return &this
}

// GetChannel returns the Channel field value
func (o *SsrUpgradeResponse) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponse) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *SsrUpgradeResponse) SetChannel(v string) {
	o.Channel = v
}

// GetCounts returns the Counts field value
func (o *SsrUpgradeResponse) GetCounts() SsrUpgradeResponseCounts {
	if o == nil {
		var ret SsrUpgradeResponseCounts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponse) GetCountsOk() (*SsrUpgradeResponseCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *SsrUpgradeResponse) SetCounts(v SsrUpgradeResponseCounts) {
	o.Counts = v
}

// GetDeviceType returns the DeviceType field value
func (o *SsrUpgradeResponse) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponse) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *SsrUpgradeResponse) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetId returns the Id field value
func (o *SsrUpgradeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SsrUpgradeResponse) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *SsrUpgradeResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SsrUpgradeResponse) SetStatus(v string) {
	o.Status = v
}

// GetStrategy returns the Strategy field value
func (o *SsrUpgradeResponse) GetStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponse) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *SsrUpgradeResponse) SetStrategy(v string) {
	o.Strategy = v
}

// GetVersions returns the Versions field value
func (o *SsrUpgradeResponse) GetVersions() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *SsrUpgradeResponse) GetVersionsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versions, true
}

// SetVersions sets field value
func (o *SsrUpgradeResponse) SetVersions(v map[string]string) {
	o.Versions = v
}

func (o SsrUpgradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsrUpgradeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["counts"] = o.Counts
	toSerialize["device_type"] = o.DeviceType
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["strategy"] = o.Strategy
	toSerialize["versions"] = o.Versions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsrUpgradeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel",
		"counts",
		"device_type",
		"id",
		"status",
		"strategy",
		"versions",
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

	varSsrUpgradeResponse := _SsrUpgradeResponse{}

	err = json.Unmarshal(data, &varSsrUpgradeResponse)

	if err != nil {
		return err
	}

	*o = SsrUpgradeResponse(varSsrUpgradeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "counts")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "strategy")
		delete(additionalProperties, "versions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSsrUpgradeResponse struct {
	value *SsrUpgradeResponse
	isSet bool
}

func (v NullableSsrUpgradeResponse) Get() *SsrUpgradeResponse {
	return v.value
}

func (v *NullableSsrUpgradeResponse) Set(val *SsrUpgradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSsrUpgradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSsrUpgradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsrUpgradeResponse(val *SsrUpgradeResponse) *NullableSsrUpgradeResponse {
	return &NullableSsrUpgradeResponse{value: val, isSet: true}
}

func (v NullableSsrUpgradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsrUpgradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


