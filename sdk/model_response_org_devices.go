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

// checks if the ResponseOrgDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseOrgDevices{}

// ResponseOrgDevices struct for ResponseOrgDevices
type ResponseOrgDevices struct {
	Results []OrgDevice `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _ResponseOrgDevices ResponseOrgDevices

// NewResponseOrgDevices instantiates a new ResponseOrgDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseOrgDevices(results []OrgDevice) *ResponseOrgDevices {
	this := ResponseOrgDevices{}
	this.Results = results
	return &this
}

// NewResponseOrgDevicesWithDefaults instantiates a new ResponseOrgDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseOrgDevicesWithDefaults() *ResponseOrgDevices {
	this := ResponseOrgDevices{}
	return &this
}

// GetResults returns the Results field value
func (o *ResponseOrgDevices) GetResults() []OrgDevice {
	if o == nil {
		var ret []OrgDevice
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ResponseOrgDevices) GetResultsOk() ([]OrgDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ResponseOrgDevices) SetResults(v []OrgDevice) {
	o.Results = v
}

func (o ResponseOrgDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseOrgDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseOrgDevices) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
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

	varResponseOrgDevices := _ResponseOrgDevices{}

	err = json.Unmarshal(data, &varResponseOrgDevices)

	if err != nil {
		return err
	}

	*o = ResponseOrgDevices(varResponseOrgDevices)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseOrgDevices struct {
	value *ResponseOrgDevices
	isSet bool
}

func (v NullableResponseOrgDevices) Get() *ResponseOrgDevices {
	return v.value
}

func (v *NullableResponseOrgDevices) Set(val *ResponseOrgDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseOrgDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseOrgDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseOrgDevices(val *ResponseOrgDevices) *NullableResponseOrgDevices {
	return &NullableResponseOrgDevices{value: val, isSet: true}
}

func (v NullableResponseOrgDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseOrgDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


