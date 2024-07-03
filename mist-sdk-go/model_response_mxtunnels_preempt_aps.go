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

// checks if the ResponseMxtunnelsPreemptAps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMxtunnelsPreemptAps{}

// ResponseMxtunnelsPreemptAps struct for ResponseMxtunnelsPreemptAps
type ResponseMxtunnelsPreemptAps struct {
	PreemptedAps []string `json:"preempted_aps"`
	AdditionalProperties map[string]interface{}
}

type _ResponseMxtunnelsPreemptAps ResponseMxtunnelsPreemptAps

// NewResponseMxtunnelsPreemptAps instantiates a new ResponseMxtunnelsPreemptAps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMxtunnelsPreemptAps(preemptedAps []string) *ResponseMxtunnelsPreemptAps {
	this := ResponseMxtunnelsPreemptAps{}
	this.PreemptedAps = preemptedAps
	return &this
}

// NewResponseMxtunnelsPreemptApsWithDefaults instantiates a new ResponseMxtunnelsPreemptAps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMxtunnelsPreemptApsWithDefaults() *ResponseMxtunnelsPreemptAps {
	this := ResponseMxtunnelsPreemptAps{}
	return &this
}

// GetPreemptedAps returns the PreemptedAps field value
func (o *ResponseMxtunnelsPreemptAps) GetPreemptedAps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PreemptedAps
}

// GetPreemptedApsOk returns a tuple with the PreemptedAps field value
// and a boolean to check if the value has been set.
func (o *ResponseMxtunnelsPreemptAps) GetPreemptedApsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreemptedAps, true
}

// SetPreemptedAps sets field value
func (o *ResponseMxtunnelsPreemptAps) SetPreemptedAps(v []string) {
	o.PreemptedAps = v
}

func (o ResponseMxtunnelsPreemptAps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMxtunnelsPreemptAps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["preempted_aps"] = o.PreemptedAps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseMxtunnelsPreemptAps) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"preempted_aps",
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

	varResponseMxtunnelsPreemptAps := _ResponseMxtunnelsPreemptAps{}

	err = json.Unmarshal(data, &varResponseMxtunnelsPreemptAps)

	if err != nil {
		return err
	}

	*o = ResponseMxtunnelsPreemptAps(varResponseMxtunnelsPreemptAps)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "preempted_aps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseMxtunnelsPreemptAps struct {
	value *ResponseMxtunnelsPreemptAps
	isSet bool
}

func (v NullableResponseMxtunnelsPreemptAps) Get() *ResponseMxtunnelsPreemptAps {
	return v.value
}

func (v *NullableResponseMxtunnelsPreemptAps) Set(val *ResponseMxtunnelsPreemptAps) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMxtunnelsPreemptAps) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMxtunnelsPreemptAps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMxtunnelsPreemptAps(val *ResponseMxtunnelsPreemptAps) *NullableResponseMxtunnelsPreemptAps {
	return &NullableResponseMxtunnelsPreemptAps{value: val, isSet: true}
}

func (v NullableResponseMxtunnelsPreemptAps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMxtunnelsPreemptAps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


