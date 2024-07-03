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

// checks if the GatewayComplianceVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayComplianceVersion{}

// GatewayComplianceVersion version compliance score, major version for gateway, type
type GatewayComplianceVersion struct {
	MajorVersion *map[string]GatewayComplianceMajorVersionProperties `json:"major_version,omitempty"`
	Score *float32 `json:"score,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayComplianceVersion GatewayComplianceVersion

// NewGatewayComplianceVersion instantiates a new GatewayComplianceVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayComplianceVersion() *GatewayComplianceVersion {
	this := GatewayComplianceVersion{}
	return &this
}

// NewGatewayComplianceVersionWithDefaults instantiates a new GatewayComplianceVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayComplianceVersionWithDefaults() *GatewayComplianceVersion {
	this := GatewayComplianceVersion{}
	return &this
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *GatewayComplianceVersion) GetMajorVersion() map[string]GatewayComplianceMajorVersionProperties {
	if o == nil || IsNil(o.MajorVersion) {
		var ret map[string]GatewayComplianceMajorVersionProperties
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayComplianceVersion) GetMajorVersionOk() (*map[string]GatewayComplianceMajorVersionProperties, bool) {
	if o == nil || IsNil(o.MajorVersion) {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *GatewayComplianceVersion) HasMajorVersion() bool {
	if o != nil && !IsNil(o.MajorVersion) {
		return true
	}

	return false
}

// SetMajorVersion gets a reference to the given map[string]GatewayComplianceMajorVersionProperties and assigns it to the MajorVersion field.
func (o *GatewayComplianceVersion) SetMajorVersion(v map[string]GatewayComplianceMajorVersionProperties) {
	o.MajorVersion = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *GatewayComplianceVersion) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayComplianceVersion) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *GatewayComplianceVersion) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *GatewayComplianceVersion) SetScore(v float32) {
	o.Score = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GatewayComplianceVersion) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayComplianceVersion) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GatewayComplianceVersion) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GatewayComplianceVersion) SetType(v string) {
	o.Type = &v
}

func (o GatewayComplianceVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayComplianceVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MajorVersion) {
		toSerialize["major_version"] = o.MajorVersion
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayComplianceVersion) UnmarshalJSON(data []byte) (err error) {
	varGatewayComplianceVersion := _GatewayComplianceVersion{}

	err = json.Unmarshal(data, &varGatewayComplianceVersion)

	if err != nil {
		return err
	}

	*o = GatewayComplianceVersion(varGatewayComplianceVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "major_version")
		delete(additionalProperties, "score")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayComplianceVersion struct {
	value *GatewayComplianceVersion
	isSet bool
}

func (v NullableGatewayComplianceVersion) Get() *GatewayComplianceVersion {
	return v.value
}

func (v *NullableGatewayComplianceVersion) Set(val *GatewayComplianceVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayComplianceVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayComplianceVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayComplianceVersion(val *GatewayComplianceVersion) *NullableGatewayComplianceVersion {
	return &NullableGatewayComplianceVersion{value: val, isSet: true}
}

func (v NullableGatewayComplianceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayComplianceVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


