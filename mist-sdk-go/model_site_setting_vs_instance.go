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

// checks if the SiteSettingVsInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingVsInstance{}

// SiteSettingVsInstance struct for SiteSettingVsInstance
type SiteSettingVsInstance struct {
	Networks []string `json:"networks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingVsInstance SiteSettingVsInstance

// NewSiteSettingVsInstance instantiates a new SiteSettingVsInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingVsInstance() *SiteSettingVsInstance {
	this := SiteSettingVsInstance{}
	return &this
}

// NewSiteSettingVsInstanceWithDefaults instantiates a new SiteSettingVsInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingVsInstanceWithDefaults() *SiteSettingVsInstance {
	this := SiteSettingVsInstance{}
	return &this
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *SiteSettingVsInstance) GetNetworks() []string {
	if o == nil || IsNil(o.Networks) {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingVsInstance) GetNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *SiteSettingVsInstance) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *SiteSettingVsInstance) SetNetworks(v []string) {
	o.Networks = v
}

func (o SiteSettingVsInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingVsInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingVsInstance) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingVsInstance := _SiteSettingVsInstance{}

	err = json.Unmarshal(data, &varSiteSettingVsInstance)

	if err != nil {
		return err
	}

	*o = SiteSettingVsInstance(varSiteSettingVsInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "networks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingVsInstance struct {
	value *SiteSettingVsInstance
	isSet bool
}

func (v NullableSiteSettingVsInstance) Get() *SiteSettingVsInstance {
	return v.value
}

func (v *NullableSiteSettingVsInstance) Set(val *SiteSettingVsInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingVsInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingVsInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingVsInstance(val *SiteSettingVsInstance) *NullableSiteSettingVsInstance {
	return &NullableSiteSettingVsInstance{value: val, isSet: true}
}

func (v NullableSiteSettingVsInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingVsInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


