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
)

// checks if the GatewayComplianceMajorVersionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayComplianceMajorVersionProperties{}

// GatewayComplianceMajorVersionProperties struct for GatewayComplianceMajorVersionProperties
type GatewayComplianceMajorVersionProperties struct {
	MajorCount *int32 `json:"major_count,omitempty"`
	MajorVersion *string `json:"major_version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayComplianceMajorVersionProperties GatewayComplianceMajorVersionProperties

// NewGatewayComplianceMajorVersionProperties instantiates a new GatewayComplianceMajorVersionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayComplianceMajorVersionProperties() *GatewayComplianceMajorVersionProperties {
	this := GatewayComplianceMajorVersionProperties{}
	return &this
}

// NewGatewayComplianceMajorVersionPropertiesWithDefaults instantiates a new GatewayComplianceMajorVersionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayComplianceMajorVersionPropertiesWithDefaults() *GatewayComplianceMajorVersionProperties {
	this := GatewayComplianceMajorVersionProperties{}
	return &this
}

// GetMajorCount returns the MajorCount field value if set, zero value otherwise.
func (o *GatewayComplianceMajorVersionProperties) GetMajorCount() int32 {
	if o == nil || IsNil(o.MajorCount) {
		var ret int32
		return ret
	}
	return *o.MajorCount
}

// GetMajorCountOk returns a tuple with the MajorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayComplianceMajorVersionProperties) GetMajorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MajorCount) {
		return nil, false
	}
	return o.MajorCount, true
}

// HasMajorCount returns a boolean if a field has been set.
func (o *GatewayComplianceMajorVersionProperties) HasMajorCount() bool {
	if o != nil && !IsNil(o.MajorCount) {
		return true
	}

	return false
}

// SetMajorCount gets a reference to the given int32 and assigns it to the MajorCount field.
func (o *GatewayComplianceMajorVersionProperties) SetMajorCount(v int32) {
	o.MajorCount = &v
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *GatewayComplianceMajorVersionProperties) GetMajorVersion() string {
	if o == nil || IsNil(o.MajorVersion) {
		var ret string
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayComplianceMajorVersionProperties) GetMajorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MajorVersion) {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *GatewayComplianceMajorVersionProperties) HasMajorVersion() bool {
	if o != nil && !IsNil(o.MajorVersion) {
		return true
	}

	return false
}

// SetMajorVersion gets a reference to the given string and assigns it to the MajorVersion field.
func (o *GatewayComplianceMajorVersionProperties) SetMajorVersion(v string) {
	o.MajorVersion = &v
}

func (o GatewayComplianceMajorVersionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayComplianceMajorVersionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MajorCount) {
		toSerialize["major_count"] = o.MajorCount
	}
	if !IsNil(o.MajorVersion) {
		toSerialize["major_version"] = o.MajorVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayComplianceMajorVersionProperties) UnmarshalJSON(data []byte) (err error) {
	varGatewayComplianceMajorVersionProperties := _GatewayComplianceMajorVersionProperties{}

	err = json.Unmarshal(data, &varGatewayComplianceMajorVersionProperties)

	if err != nil {
		return err
	}

	*o = GatewayComplianceMajorVersionProperties(varGatewayComplianceMajorVersionProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "major_count")
		delete(additionalProperties, "major_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayComplianceMajorVersionProperties struct {
	value *GatewayComplianceMajorVersionProperties
	isSet bool
}

func (v NullableGatewayComplianceMajorVersionProperties) Get() *GatewayComplianceMajorVersionProperties {
	return v.value
}

func (v *NullableGatewayComplianceMajorVersionProperties) Set(val *GatewayComplianceMajorVersionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayComplianceMajorVersionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayComplianceMajorVersionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayComplianceMajorVersionProperties(val *GatewayComplianceMajorVersionProperties) *NullableGatewayComplianceMajorVersionProperties {
	return &NullableGatewayComplianceMajorVersionProperties{value: val, isSet: true}
}

func (v NullableGatewayComplianceMajorVersionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayComplianceMajorVersionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


