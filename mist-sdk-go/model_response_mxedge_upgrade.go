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

// checks if the ResponseMxedgeUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMxedgeUpgrade{}

// ResponseMxedgeUpgrade struct for ResponseMxedgeUpgrade
type ResponseMxedgeUpgrade struct {
	Channel string `json:"channel"`
	Counts MxedgeUpgradeResponseCounts `json:"counts"`
	Id string `json:"id"`
	Status string `json:"status"`
	Strategy string `json:"strategy"`
	Versions map[string]interface{} `json:"versions"`
	AdditionalProperties map[string]interface{}
}

type _ResponseMxedgeUpgrade ResponseMxedgeUpgrade

// NewResponseMxedgeUpgrade instantiates a new ResponseMxedgeUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMxedgeUpgrade(channel string, counts MxedgeUpgradeResponseCounts, id string, status string, strategy string, versions map[string]interface{}) *ResponseMxedgeUpgrade {
	this := ResponseMxedgeUpgrade{}
	this.Channel = channel
	this.Counts = counts
	this.Id = id
	this.Status = status
	this.Strategy = strategy
	this.Versions = versions
	return &this
}

// NewResponseMxedgeUpgradeWithDefaults instantiates a new ResponseMxedgeUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMxedgeUpgradeWithDefaults() *ResponseMxedgeUpgrade {
	this := ResponseMxedgeUpgrade{}
	return &this
}

// GetChannel returns the Channel field value
func (o *ResponseMxedgeUpgrade) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeUpgrade) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *ResponseMxedgeUpgrade) SetChannel(v string) {
	o.Channel = v
}

// GetCounts returns the Counts field value
func (o *ResponseMxedgeUpgrade) GetCounts() MxedgeUpgradeResponseCounts {
	if o == nil {
		var ret MxedgeUpgradeResponseCounts
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeUpgrade) GetCountsOk() (*MxedgeUpgradeResponseCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value
func (o *ResponseMxedgeUpgrade) SetCounts(v MxedgeUpgradeResponseCounts) {
	o.Counts = v
}

// GetId returns the Id field value
func (o *ResponseMxedgeUpgrade) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeUpgrade) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseMxedgeUpgrade) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *ResponseMxedgeUpgrade) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeUpgrade) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ResponseMxedgeUpgrade) SetStatus(v string) {
	o.Status = v
}

// GetStrategy returns the Strategy field value
func (o *ResponseMxedgeUpgrade) GetStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeUpgrade) GetStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value
func (o *ResponseMxedgeUpgrade) SetStrategy(v string) {
	o.Strategy = v
}

// GetVersions returns the Versions field value
func (o *ResponseMxedgeUpgrade) GetVersions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *ResponseMxedgeUpgrade) GetVersionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *ResponseMxedgeUpgrade) SetVersions(v map[string]interface{}) {
	o.Versions = v
}

func (o ResponseMxedgeUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMxedgeUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	toSerialize["counts"] = o.Counts
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["strategy"] = o.Strategy
	toSerialize["versions"] = o.Versions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseMxedgeUpgrade) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel",
		"counts",
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

	varResponseMxedgeUpgrade := _ResponseMxedgeUpgrade{}

	err = json.Unmarshal(data, &varResponseMxedgeUpgrade)

	if err != nil {
		return err
	}

	*o = ResponseMxedgeUpgrade(varResponseMxedgeUpgrade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "counts")
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "strategy")
		delete(additionalProperties, "versions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseMxedgeUpgrade struct {
	value *ResponseMxedgeUpgrade
	isSet bool
}

func (v NullableResponseMxedgeUpgrade) Get() *ResponseMxedgeUpgrade {
	return v.value
}

func (v *NullableResponseMxedgeUpgrade) Set(val *ResponseMxedgeUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMxedgeUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMxedgeUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMxedgeUpgrade(val *ResponseMxedgeUpgrade) *NullableResponseMxedgeUpgrade {
	return &NullableResponseMxedgeUpgrade{value: val, isSet: true}
}

func (v NullableResponseMxedgeUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMxedgeUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


