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

// checks if the OrgSleStatUserMinutes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSleStatUserMinutes{}

// OrgSleStatUserMinutes struct for OrgSleStatUserMinutes
type OrgSleStatUserMinutes struct {
	Ok float32 `json:"ok"`
	Total float32 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _OrgSleStatUserMinutes OrgSleStatUserMinutes

// NewOrgSleStatUserMinutes instantiates a new OrgSleStatUserMinutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSleStatUserMinutes(ok float32, total float32) *OrgSleStatUserMinutes {
	this := OrgSleStatUserMinutes{}
	this.Ok = ok
	this.Total = total
	return &this
}

// NewOrgSleStatUserMinutesWithDefaults instantiates a new OrgSleStatUserMinutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSleStatUserMinutesWithDefaults() *OrgSleStatUserMinutes {
	this := OrgSleStatUserMinutes{}
	return &this
}

// GetOk returns the Ok field value
func (o *OrgSleStatUserMinutes) GetOk() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value
// and a boolean to check if the value has been set.
func (o *OrgSleStatUserMinutes) GetOkOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ok, true
}

// SetOk sets field value
func (o *OrgSleStatUserMinutes) SetOk(v float32) {
	o.Ok = v
}

// GetTotal returns the Total field value
func (o *OrgSleStatUserMinutes) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *OrgSleStatUserMinutes) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *OrgSleStatUserMinutes) SetTotal(v float32) {
	o.Total = v
}

func (o OrgSleStatUserMinutes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSleStatUserMinutes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ok"] = o.Ok
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSleStatUserMinutes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ok",
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

	varOrgSleStatUserMinutes := _OrgSleStatUserMinutes{}

	err = json.Unmarshal(data, &varOrgSleStatUserMinutes)

	if err != nil {
		return err
	}

	*o = OrgSleStatUserMinutes(varOrgSleStatUserMinutes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ok")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSleStatUserMinutes struct {
	value *OrgSleStatUserMinutes
	isSet bool
}

func (v NullableOrgSleStatUserMinutes) Get() *OrgSleStatUserMinutes {
	return v.value
}

func (v *NullableOrgSleStatUserMinutes) Set(val *OrgSleStatUserMinutes) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSleStatUserMinutes) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSleStatUserMinutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSleStatUserMinutes(val *OrgSleStatUserMinutes) *NullableOrgSleStatUserMinutes {
	return &NullableOrgSleStatUserMinutes{value: val, isSet: true}
}

func (v NullableOrgSleStatUserMinutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSleStatUserMinutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


