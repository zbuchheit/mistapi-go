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

// checks if the DswitchesMetricsPoeComplianceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DswitchesMetricsPoeComplianceDetails{}

// DswitchesMetricsPoeComplianceDetails struct for DswitchesMetricsPoeComplianceDetails
type DswitchesMetricsPoeComplianceDetails struct {
	TotalAps int32 `json:"total_aps"`
	TotalPower float32 `json:"total_power"`
	AdditionalProperties map[string]interface{}
}

type _DswitchesMetricsPoeComplianceDetails DswitchesMetricsPoeComplianceDetails

// NewDswitchesMetricsPoeComplianceDetails instantiates a new DswitchesMetricsPoeComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDswitchesMetricsPoeComplianceDetails(totalAps int32, totalPower float32) *DswitchesMetricsPoeComplianceDetails {
	this := DswitchesMetricsPoeComplianceDetails{}
	this.TotalAps = totalAps
	this.TotalPower = totalPower
	return &this
}

// NewDswitchesMetricsPoeComplianceDetailsWithDefaults instantiates a new DswitchesMetricsPoeComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDswitchesMetricsPoeComplianceDetailsWithDefaults() *DswitchesMetricsPoeComplianceDetails {
	this := DswitchesMetricsPoeComplianceDetails{}
	return &this
}

// GetTotalAps returns the TotalAps field value
func (o *DswitchesMetricsPoeComplianceDetails) GetTotalAps() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAps
}

// GetTotalApsOk returns a tuple with the TotalAps field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsPoeComplianceDetails) GetTotalApsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAps, true
}

// SetTotalAps sets field value
func (o *DswitchesMetricsPoeComplianceDetails) SetTotalAps(v int32) {
	o.TotalAps = v
}

// GetTotalPower returns the TotalPower field value
func (o *DswitchesMetricsPoeComplianceDetails) GetTotalPower() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPower
}

// GetTotalPowerOk returns a tuple with the TotalPower field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsPoeComplianceDetails) GetTotalPowerOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPower, true
}

// SetTotalPower sets field value
func (o *DswitchesMetricsPoeComplianceDetails) SetTotalPower(v float32) {
	o.TotalPower = v
}

func (o DswitchesMetricsPoeComplianceDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DswitchesMetricsPoeComplianceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_aps"] = o.TotalAps
	toSerialize["total_power"] = o.TotalPower

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DswitchesMetricsPoeComplianceDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_aps",
		"total_power",
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

	varDswitchesMetricsPoeComplianceDetails := _DswitchesMetricsPoeComplianceDetails{}

	err = json.Unmarshal(data, &varDswitchesMetricsPoeComplianceDetails)

	if err != nil {
		return err
	}

	*o = DswitchesMetricsPoeComplianceDetails(varDswitchesMetricsPoeComplianceDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total_aps")
		delete(additionalProperties, "total_power")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDswitchesMetricsPoeComplianceDetails struct {
	value *DswitchesMetricsPoeComplianceDetails
	isSet bool
}

func (v NullableDswitchesMetricsPoeComplianceDetails) Get() *DswitchesMetricsPoeComplianceDetails {
	return v.value
}

func (v *NullableDswitchesMetricsPoeComplianceDetails) Set(val *DswitchesMetricsPoeComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDswitchesMetricsPoeComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDswitchesMetricsPoeComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDswitchesMetricsPoeComplianceDetails(val *DswitchesMetricsPoeComplianceDetails) *NullableDswitchesMetricsPoeComplianceDetails {
	return &NullableDswitchesMetricsPoeComplianceDetails{value: val, isSet: true}
}

func (v NullableDswitchesMetricsPoeComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDswitchesMetricsPoeComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


