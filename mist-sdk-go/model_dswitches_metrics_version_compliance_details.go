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

// checks if the DswitchesMetricsVersionComplianceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DswitchesMetricsVersionComplianceDetails{}

// DswitchesMetricsVersionComplianceDetails struct for DswitchesMetricsVersionComplianceDetails
type DswitchesMetricsVersionComplianceDetails struct {
	MajorVersions []DswitchesComplianceMajorVersion `json:"major_versions"`
	TotalSwitchCount int32 `json:"total_switch_count"`
	AdditionalProperties map[string]interface{}
}

type _DswitchesMetricsVersionComplianceDetails DswitchesMetricsVersionComplianceDetails

// NewDswitchesMetricsVersionComplianceDetails instantiates a new DswitchesMetricsVersionComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDswitchesMetricsVersionComplianceDetails(majorVersions []DswitchesComplianceMajorVersion, totalSwitchCount int32) *DswitchesMetricsVersionComplianceDetails {
	this := DswitchesMetricsVersionComplianceDetails{}
	this.MajorVersions = majorVersions
	this.TotalSwitchCount = totalSwitchCount
	return &this
}

// NewDswitchesMetricsVersionComplianceDetailsWithDefaults instantiates a new DswitchesMetricsVersionComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDswitchesMetricsVersionComplianceDetailsWithDefaults() *DswitchesMetricsVersionComplianceDetails {
	this := DswitchesMetricsVersionComplianceDetails{}
	return &this
}

// GetMajorVersions returns the MajorVersions field value
func (o *DswitchesMetricsVersionComplianceDetails) GetMajorVersions() []DswitchesComplianceMajorVersion {
	if o == nil {
		var ret []DswitchesComplianceMajorVersion
		return ret
	}

	return o.MajorVersions
}

// GetMajorVersionsOk returns a tuple with the MajorVersions field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsVersionComplianceDetails) GetMajorVersionsOk() ([]DswitchesComplianceMajorVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.MajorVersions, true
}

// SetMajorVersions sets field value
func (o *DswitchesMetricsVersionComplianceDetails) SetMajorVersions(v []DswitchesComplianceMajorVersion) {
	o.MajorVersions = v
}

// GetTotalSwitchCount returns the TotalSwitchCount field value
func (o *DswitchesMetricsVersionComplianceDetails) GetTotalSwitchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalSwitchCount
}

// GetTotalSwitchCountOk returns a tuple with the TotalSwitchCount field value
// and a boolean to check if the value has been set.
func (o *DswitchesMetricsVersionComplianceDetails) GetTotalSwitchCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSwitchCount, true
}

// SetTotalSwitchCount sets field value
func (o *DswitchesMetricsVersionComplianceDetails) SetTotalSwitchCount(v int32) {
	o.TotalSwitchCount = v
}

func (o DswitchesMetricsVersionComplianceDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DswitchesMetricsVersionComplianceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["major_versions"] = o.MajorVersions
	toSerialize["total_switch_count"] = o.TotalSwitchCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DswitchesMetricsVersionComplianceDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"major_versions",
		"total_switch_count",
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

	varDswitchesMetricsVersionComplianceDetails := _DswitchesMetricsVersionComplianceDetails{}

	err = json.Unmarshal(data, &varDswitchesMetricsVersionComplianceDetails)

	if err != nil {
		return err
	}

	*o = DswitchesMetricsVersionComplianceDetails(varDswitchesMetricsVersionComplianceDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "major_versions")
		delete(additionalProperties, "total_switch_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDswitchesMetricsVersionComplianceDetails struct {
	value *DswitchesMetricsVersionComplianceDetails
	isSet bool
}

func (v NullableDswitchesMetricsVersionComplianceDetails) Get() *DswitchesMetricsVersionComplianceDetails {
	return v.value
}

func (v *NullableDswitchesMetricsVersionComplianceDetails) Set(val *DswitchesMetricsVersionComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDswitchesMetricsVersionComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDswitchesMetricsVersionComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDswitchesMetricsVersionComplianceDetails(val *DswitchesMetricsVersionComplianceDetails) *NullableDswitchesMetricsVersionComplianceDetails {
	return &NullableDswitchesMetricsVersionComplianceDetails{value: val, isSet: true}
}

func (v NullableDswitchesMetricsVersionComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDswitchesMetricsVersionComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


