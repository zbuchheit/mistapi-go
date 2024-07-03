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

// checks if the DswitchesMetricsVersionCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DswitchesMetricsVersionCompliance{}

// DswitchesMetricsVersionCompliance struct for DswitchesMetricsVersionCompliance
type DswitchesMetricsVersionCompliance struct {
	Details DswitchesMetricsVersionComplianceDetails `json:"details"`
	Score float32 `json:"score"`
	AdditionalProperties map[string]interface{}
}

type _DswitchesMetricsVersionCompliance DswitchesMetricsVersionCompliance

// NewDswitchesMetricsVersionCompliance instantiates a new DswitchesMetricsVersionCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDswitchesMetricsVersionCompliance(details DswitchesMetricsVersionComplianceDetails, score float32) *DswitchesMetricsVersionCompliance {
	this := DswitchesMetricsVersionCompliance{}
	this.Details = details
	this.Score = score
	return &this
}

// NewDswitchesMetricsVersionComplianceWithDefaults instantiates a new DswitchesMetricsVersionCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDswitchesMetricsVersionComplianceWithDefaults() *DswitchesMetricsVersionCompliance {
	this := DswitchesMetricsVersionCompliance{}
	return &this
}

// GetDetails returns the Details field value
func (o *DswitchesMetricsVersionCompliance) GetDetails() DswitchesMetricsVersionComplianceDetails {
	if o == nil {
		var ret DswitchesMetricsVersionComplianceDetails
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsVersionCompliance) GetDetailsOk() (*DswitchesMetricsVersionComplianceDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *DswitchesMetricsVersionCompliance) SetDetails(v DswitchesMetricsVersionComplianceDetails) {
	o.Details = v
}

// GetScore returns the Score field value
func (o *DswitchesMetricsVersionCompliance) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsVersionCompliance) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *DswitchesMetricsVersionCompliance) SetScore(v float32) {
	o.Score = v
}

func (o DswitchesMetricsVersionCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DswitchesMetricsVersionCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["details"] = o.Details
	toSerialize["score"] = o.Score

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DswitchesMetricsVersionCompliance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"details",
		"score",
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

	varDswitchesMetricsVersionCompliance := _DswitchesMetricsVersionCompliance{}

	err = json.Unmarshal(data, &varDswitchesMetricsVersionCompliance)

	if err != nil {
		return err
	}

	*o = DswitchesMetricsVersionCompliance(varDswitchesMetricsVersionCompliance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "details")
		delete(additionalProperties, "score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDswitchesMetricsVersionCompliance struct {
	value *DswitchesMetricsVersionCompliance
	isSet bool
}

func (v NullableDswitchesMetricsVersionCompliance) Get() *DswitchesMetricsVersionCompliance {
	return v.value
}

func (v *NullableDswitchesMetricsVersionCompliance) Set(val *DswitchesMetricsVersionCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableDswitchesMetricsVersionCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableDswitchesMetricsVersionCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDswitchesMetricsVersionCompliance(val *DswitchesMetricsVersionCompliance) *NullableDswitchesMetricsVersionCompliance {
	return &NullableDswitchesMetricsVersionCompliance{value: val, isSet: true}
}

func (v NullableDswitchesMetricsVersionCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDswitchesMetricsVersionCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


