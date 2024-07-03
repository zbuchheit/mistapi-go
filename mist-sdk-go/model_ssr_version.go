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

// checks if the SsrVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsrVersion{}

// SsrVersion struct for SsrVersion
type SsrVersion struct {
	Default *bool `json:"default,omitempty"`
	Package string `json:"package"`
	Version string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _SsrVersion SsrVersion

// NewSsrVersion instantiates a new SsrVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsrVersion(package_ string, version string) *SsrVersion {
	this := SsrVersion{}
	this.Package = package_
	this.Version = version
	return &this
}

// NewSsrVersionWithDefaults instantiates a new SsrVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsrVersionWithDefaults() *SsrVersion {
	this := SsrVersion{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *SsrVersion) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsrVersion) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *SsrVersion) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *SsrVersion) SetDefault(v bool) {
	o.Default = &v
}

// GetPackage returns the Package field value
func (o *SsrVersion) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *SsrVersion) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *SsrVersion) SetPackage(v string) {
	o.Package = v
}

// GetVersion returns the Version field value
func (o *SsrVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SsrVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SsrVersion) SetVersion(v string) {
	o.Version = v
}

func (o SsrVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsrVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	toSerialize["package"] = o.Package
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsrVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"package",
		"version",
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

	varSsrVersion := _SsrVersion{}

	err = json.Unmarshal(data, &varSsrVersion)

	if err != nil {
		return err
	}

	*o = SsrVersion(varSsrVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default")
		delete(additionalProperties, "package")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSsrVersion struct {
	value *SsrVersion
	isSet bool
}

func (v NullableSsrVersion) Get() *SsrVersion {
	return v.value
}

func (v *NullableSsrVersion) Set(val *SsrVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableSsrVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableSsrVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsrVersion(val *SsrVersion) *NullableSsrVersion {
	return &NullableSsrVersion{value: val, isSet: true}
}

func (v NullableSsrVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsrVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


