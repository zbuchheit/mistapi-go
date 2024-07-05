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
)

// checks if the GatewayMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayMetrics{}

// GatewayMetrics struct for GatewayMetrics
type GatewayMetrics struct {
	// config success score
	ConfigSuccess *float32 `json:"config_success,omitempty"`
	VersionCompliance *GatewayComplianceVersion `json:"version_compliance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayMetrics GatewayMetrics

// NewGatewayMetrics instantiates a new GatewayMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMetrics() *GatewayMetrics {
	this := GatewayMetrics{}
	return &this
}

// NewGatewayMetricsWithDefaults instantiates a new GatewayMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMetricsWithDefaults() *GatewayMetrics {
	this := GatewayMetrics{}
	return &this
}

// GetConfigSuccess returns the ConfigSuccess field value if set, zero value otherwise.
func (o *GatewayMetrics) GetConfigSuccess() float32 {
	if o == nil || IsNil(o.ConfigSuccess) {
		var ret float32
		return ret
	}
	return *o.ConfigSuccess
}

// GetConfigSuccessOk returns a tuple with the ConfigSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMetrics) GetConfigSuccessOk() (*float32, bool) {
	if o == nil || IsNil(o.ConfigSuccess) {
		return nil, false
	}
	return o.ConfigSuccess, true
}

// HasConfigSuccess returns a boolean if a field has been set.
func (o *GatewayMetrics) HasConfigSuccess() bool {
	if o != nil && !IsNil(o.ConfigSuccess) {
		return true
	}

	return false
}

// SetConfigSuccess gets a reference to the given float32 and assigns it to the ConfigSuccess field.
func (o *GatewayMetrics) SetConfigSuccess(v float32) {
	o.ConfigSuccess = &v
}

// GetVersionCompliance returns the VersionCompliance field value if set, zero value otherwise.
func (o *GatewayMetrics) GetVersionCompliance() GatewayComplianceVersion {
	if o == nil || IsNil(o.VersionCompliance) {
		var ret GatewayComplianceVersion
		return ret
	}
	return *o.VersionCompliance
}

// GetVersionComplianceOk returns a tuple with the VersionCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMetrics) GetVersionComplianceOk() (*GatewayComplianceVersion, bool) {
	if o == nil || IsNil(o.VersionCompliance) {
		return nil, false
	}
	return o.VersionCompliance, true
}

// HasVersionCompliance returns a boolean if a field has been set.
func (o *GatewayMetrics) HasVersionCompliance() bool {
	if o != nil && !IsNil(o.VersionCompliance) {
		return true
	}

	return false
}

// SetVersionCompliance gets a reference to the given GatewayComplianceVersion and assigns it to the VersionCompliance field.
func (o *GatewayMetrics) SetVersionCompliance(v GatewayComplianceVersion) {
	o.VersionCompliance = &v
}

func (o GatewayMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigSuccess) {
		toSerialize["config_success"] = o.ConfigSuccess
	}
	if !IsNil(o.VersionCompliance) {
		toSerialize["version_compliance"] = o.VersionCompliance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayMetrics) UnmarshalJSON(data []byte) (err error) {
	varGatewayMetrics := _GatewayMetrics{}

	err = json.Unmarshal(data, &varGatewayMetrics)

	if err != nil {
		return err
	}

	*o = GatewayMetrics(varGatewayMetrics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config_success")
		delete(additionalProperties, "version_compliance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayMetrics struct {
	value *GatewayMetrics
	isSet bool
}

func (v NullableGatewayMetrics) Get() *GatewayMetrics {
	return v.value
}

func (v *NullableGatewayMetrics) Set(val *GatewayMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMetrics(val *GatewayMetrics) *NullableGatewayMetrics {
	return &NullableGatewayMetrics{value: val, isSet: true}
}

func (v NullableGatewayMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


