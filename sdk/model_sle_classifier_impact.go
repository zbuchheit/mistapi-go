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

// checks if the SleClassifierImpact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleClassifierImpact{}

// SleClassifierImpact struct for SleClassifierImpact
type SleClassifierImpact struct {
	NumAps float32 `json:"num_aps"`
	NumUsers float32 `json:"num_users"`
	TotalAps float32 `json:"total_aps"`
	TotalUsers float32 `json:"total_users"`
	AdditionalProperties map[string]interface{}
}

type _SleClassifierImpact SleClassifierImpact

// NewSleClassifierImpact instantiates a new SleClassifierImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleClassifierImpact(numAps float32, numUsers float32, totalAps float32, totalUsers float32) *SleClassifierImpact {
	this := SleClassifierImpact{}
	this.NumAps = numAps
	this.NumUsers = numUsers
	this.TotalAps = totalAps
	this.TotalUsers = totalUsers
	return &this
}

// NewSleClassifierImpactWithDefaults instantiates a new SleClassifierImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleClassifierImpactWithDefaults() *SleClassifierImpact {
	this := SleClassifierImpact{}
	return &this
}

// GetNumAps returns the NumAps field value
func (o *SleClassifierImpact) GetNumAps() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NumAps
}

// GetNumApsOk returns a tuple with the NumAps field value
// and a boolean to check if the value has been set.
func (o *SleClassifierImpact) GetNumApsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumAps, true
}

// SetNumAps sets field value
func (o *SleClassifierImpact) SetNumAps(v float32) {
	o.NumAps = v
}

// GetNumUsers returns the NumUsers field value
func (o *SleClassifierImpact) GetNumUsers() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NumUsers
}

// GetNumUsersOk returns a tuple with the NumUsers field value
// and a boolean to check if the value has been set.
func (o *SleClassifierImpact) GetNumUsersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumUsers, true
}

// SetNumUsers sets field value
func (o *SleClassifierImpact) SetNumUsers(v float32) {
	o.NumUsers = v
}

// GetTotalAps returns the TotalAps field value
func (o *SleClassifierImpact) GetTotalAps() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalAps
}

// GetTotalApsOk returns a tuple with the TotalAps field value
// and a boolean to check if the value has been set.
func (o *SleClassifierImpact) GetTotalApsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAps, true
}

// SetTotalAps sets field value
func (o *SleClassifierImpact) SetTotalAps(v float32) {
	o.TotalAps = v
}

// GetTotalUsers returns the TotalUsers field value
func (o *SleClassifierImpact) GetTotalUsers() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalUsers
}

// GetTotalUsersOk returns a tuple with the TotalUsers field value
// and a boolean to check if the value has been set.
func (o *SleClassifierImpact) GetTotalUsersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalUsers, true
}

// SetTotalUsers sets field value
func (o *SleClassifierImpact) SetTotalUsers(v float32) {
	o.TotalUsers = v
}

func (o SleClassifierImpact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleClassifierImpact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num_aps"] = o.NumAps
	toSerialize["num_users"] = o.NumUsers
	toSerialize["total_aps"] = o.TotalAps
	toSerialize["total_users"] = o.TotalUsers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleClassifierImpact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num_aps",
		"num_users",
		"total_aps",
		"total_users",
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

	varSleClassifierImpact := _SleClassifierImpact{}

	err = json.Unmarshal(data, &varSleClassifierImpact)

	if err != nil {
		return err
	}

	*o = SleClassifierImpact(varSleClassifierImpact)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "num_aps")
		delete(additionalProperties, "num_users")
		delete(additionalProperties, "total_aps")
		delete(additionalProperties, "total_users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleClassifierImpact struct {
	value *SleClassifierImpact
	isSet bool
}

func (v NullableSleClassifierImpact) Get() *SleClassifierImpact {
	return v.value
}

func (v *NullableSleClassifierImpact) Set(val *SleClassifierImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableSleClassifierImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableSleClassifierImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleClassifierImpact(val *SleClassifierImpact) *NullableSleClassifierImpact {
	return &NullableSleClassifierImpact{value: val, isSet: true}
}

func (v NullableSleClassifierImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleClassifierImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


