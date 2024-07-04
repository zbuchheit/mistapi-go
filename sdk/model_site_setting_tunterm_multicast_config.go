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
)

// checks if the SiteSettingTuntermMulticastConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingTuntermMulticastConfig{}

// SiteSettingTuntermMulticastConfig struct for SiteSettingTuntermMulticastConfig
type SiteSettingTuntermMulticastConfig struct {
	Mdns *SiteSettingTuntermMulticastConfigMdns `json:"mdns,omitempty"`
	MulticastAll *bool `json:"multicast_all,omitempty"`
	Ssdp *SiteSettingTuntermMulticastConfigSsdp `json:"ssdp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingTuntermMulticastConfig SiteSettingTuntermMulticastConfig

// NewSiteSettingTuntermMulticastConfig instantiates a new SiteSettingTuntermMulticastConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingTuntermMulticastConfig() *SiteSettingTuntermMulticastConfig {
	this := SiteSettingTuntermMulticastConfig{}
	var multicastAll bool = false
	this.MulticastAll = &multicastAll
	return &this
}

// NewSiteSettingTuntermMulticastConfigWithDefaults instantiates a new SiteSettingTuntermMulticastConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingTuntermMulticastConfigWithDefaults() *SiteSettingTuntermMulticastConfig {
	this := SiteSettingTuntermMulticastConfig{}
	var multicastAll bool = false
	this.MulticastAll = &multicastAll
	return &this
}

// GetMdns returns the Mdns field value if set, zero value otherwise.
func (o *SiteSettingTuntermMulticastConfig) GetMdns() SiteSettingTuntermMulticastConfigMdns {
	if o == nil || IsNil(o.Mdns) {
		var ret SiteSettingTuntermMulticastConfigMdns
		return ret
	}
	return *o.Mdns
}

// GetMdnsOk returns a tuple with the Mdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingTuntermMulticastConfig) GetMdnsOk() (*SiteSettingTuntermMulticastConfigMdns, bool) {
	if o == nil || IsNil(o.Mdns) {
		return nil, false
	}
	return o.Mdns, true
}

// HasMdns returns a boolean if a field has been set.
func (o *SiteSettingTuntermMulticastConfig) HasMdns() bool {
	if o != nil && !IsNil(o.Mdns) {
		return true
	}

	return false
}

// SetMdns gets a reference to the given SiteSettingTuntermMulticastConfigMdns and assigns it to the Mdns field.
func (o *SiteSettingTuntermMulticastConfig) SetMdns(v SiteSettingTuntermMulticastConfigMdns) {
	o.Mdns = &v
}

// GetMulticastAll returns the MulticastAll field value if set, zero value otherwise.
func (o *SiteSettingTuntermMulticastConfig) GetMulticastAll() bool {
	if o == nil || IsNil(o.MulticastAll) {
		var ret bool
		return ret
	}
	return *o.MulticastAll
}

// GetMulticastAllOk returns a tuple with the MulticastAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingTuntermMulticastConfig) GetMulticastAllOk() (*bool, bool) {
	if o == nil || IsNil(o.MulticastAll) {
		return nil, false
	}
	return o.MulticastAll, true
}

// HasMulticastAll returns a boolean if a field has been set.
func (o *SiteSettingTuntermMulticastConfig) HasMulticastAll() bool {
	if o != nil && !IsNil(o.MulticastAll) {
		return true
	}

	return false
}

// SetMulticastAll gets a reference to the given bool and assigns it to the MulticastAll field.
func (o *SiteSettingTuntermMulticastConfig) SetMulticastAll(v bool) {
	o.MulticastAll = &v
}

// GetSsdp returns the Ssdp field value if set, zero value otherwise.
func (o *SiteSettingTuntermMulticastConfig) GetSsdp() SiteSettingTuntermMulticastConfigSsdp {
	if o == nil || IsNil(o.Ssdp) {
		var ret SiteSettingTuntermMulticastConfigSsdp
		return ret
	}
	return *o.Ssdp
}

// GetSsdpOk returns a tuple with the Ssdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingTuntermMulticastConfig) GetSsdpOk() (*SiteSettingTuntermMulticastConfigSsdp, bool) {
	if o == nil || IsNil(o.Ssdp) {
		return nil, false
	}
	return o.Ssdp, true
}

// HasSsdp returns a boolean if a field has been set.
func (o *SiteSettingTuntermMulticastConfig) HasSsdp() bool {
	if o != nil && !IsNil(o.Ssdp) {
		return true
	}

	return false
}

// SetSsdp gets a reference to the given SiteSettingTuntermMulticastConfigSsdp and assigns it to the Ssdp field.
func (o *SiteSettingTuntermMulticastConfig) SetSsdp(v SiteSettingTuntermMulticastConfigSsdp) {
	o.Ssdp = &v
}

func (o SiteSettingTuntermMulticastConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingTuntermMulticastConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mdns) {
		toSerialize["mdns"] = o.Mdns
	}
	if !IsNil(o.MulticastAll) {
		toSerialize["multicast_all"] = o.MulticastAll
	}
	if !IsNil(o.Ssdp) {
		toSerialize["ssdp"] = o.Ssdp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingTuntermMulticastConfig) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingTuntermMulticastConfig := _SiteSettingTuntermMulticastConfig{}

	err = json.Unmarshal(data, &varSiteSettingTuntermMulticastConfig)

	if err != nil {
		return err
	}

	*o = SiteSettingTuntermMulticastConfig(varSiteSettingTuntermMulticastConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mdns")
		delete(additionalProperties, "multicast_all")
		delete(additionalProperties, "ssdp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingTuntermMulticastConfig struct {
	value *SiteSettingTuntermMulticastConfig
	isSet bool
}

func (v NullableSiteSettingTuntermMulticastConfig) Get() *SiteSettingTuntermMulticastConfig {
	return v.value
}

func (v *NullableSiteSettingTuntermMulticastConfig) Set(val *SiteSettingTuntermMulticastConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingTuntermMulticastConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingTuntermMulticastConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingTuntermMulticastConfig(val *SiteSettingTuntermMulticastConfig) *NullableSiteSettingTuntermMulticastConfig {
	return &NullableSiteSettingTuntermMulticastConfig{value: val, isSet: true}
}

func (v NullableSiteSettingTuntermMulticastConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingTuntermMulticastConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


