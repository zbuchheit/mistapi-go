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

// checks if the SiteSettingApMatchingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingApMatchingRule{}

// SiteSettingApMatchingRule struct for SiteSettingApMatchingRule
type SiteSettingApMatchingRule struct {
	MatchModel *string `json:"match_model,omitempty"`
	Name *string `json:"name,omitempty"`
	// Property key is the interface(s) (e.g. \"eth1,eth2\")
	PortConfig *map[string]ApPortConfig `json:"port_config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingApMatchingRule SiteSettingApMatchingRule

// NewSiteSettingApMatchingRule instantiates a new SiteSettingApMatchingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingApMatchingRule() *SiteSettingApMatchingRule {
	this := SiteSettingApMatchingRule{}
	return &this
}

// NewSiteSettingApMatchingRuleWithDefaults instantiates a new SiteSettingApMatchingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingApMatchingRuleWithDefaults() *SiteSettingApMatchingRule {
	this := SiteSettingApMatchingRule{}
	return &this
}

// GetMatchModel returns the MatchModel field value if set, zero value otherwise.
func (o *SiteSettingApMatchingRule) GetMatchModel() string {
	if o == nil || IsNil(o.MatchModel) {
		var ret string
		return ret
	}
	return *o.MatchModel
}

// GetMatchModelOk returns a tuple with the MatchModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingApMatchingRule) GetMatchModelOk() (*string, bool) {
	if o == nil || IsNil(o.MatchModel) {
		return nil, false
	}
	return o.MatchModel, true
}

// HasMatchModel returns a boolean if a field has been set.
func (o *SiteSettingApMatchingRule) HasMatchModel() bool {
	if o != nil && !IsNil(o.MatchModel) {
		return true
	}

	return false
}

// SetMatchModel gets a reference to the given string and assigns it to the MatchModel field.
func (o *SiteSettingApMatchingRule) SetMatchModel(v string) {
	o.MatchModel = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SiteSettingApMatchingRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingApMatchingRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SiteSettingApMatchingRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SiteSettingApMatchingRule) SetName(v string) {
	o.Name = &v
}

// GetPortConfig returns the PortConfig field value if set, zero value otherwise.
func (o *SiteSettingApMatchingRule) GetPortConfig() map[string]ApPortConfig {
	if o == nil || IsNil(o.PortConfig) {
		var ret map[string]ApPortConfig
		return ret
	}
	return *o.PortConfig
}

// GetPortConfigOk returns a tuple with the PortConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingApMatchingRule) GetPortConfigOk() (*map[string]ApPortConfig, bool) {
	if o == nil || IsNil(o.PortConfig) {
		return nil, false
	}
	return o.PortConfig, true
}

// HasPortConfig returns a boolean if a field has been set.
func (o *SiteSettingApMatchingRule) HasPortConfig() bool {
	if o != nil && !IsNil(o.PortConfig) {
		return true
	}

	return false
}

// SetPortConfig gets a reference to the given map[string]ApPortConfig and assigns it to the PortConfig field.
func (o *SiteSettingApMatchingRule) SetPortConfig(v map[string]ApPortConfig) {
	o.PortConfig = &v
}

func (o SiteSettingApMatchingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingApMatchingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchModel) {
		toSerialize["match_model"] = o.MatchModel
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PortConfig) {
		toSerialize["port_config"] = o.PortConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingApMatchingRule) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingApMatchingRule := _SiteSettingApMatchingRule{}

	err = json.Unmarshal(data, &varSiteSettingApMatchingRule)

	if err != nil {
		return err
	}

	*o = SiteSettingApMatchingRule(varSiteSettingApMatchingRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "match_model")
		delete(additionalProperties, "name")
		delete(additionalProperties, "port_config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingApMatchingRule struct {
	value *SiteSettingApMatchingRule
	isSet bool
}

func (v NullableSiteSettingApMatchingRule) Get() *SiteSettingApMatchingRule {
	return v.value
}

func (v *NullableSiteSettingApMatchingRule) Set(val *SiteSettingApMatchingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingApMatchingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingApMatchingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingApMatchingRule(val *SiteSettingApMatchingRule) *NullableSiteSettingApMatchingRule {
	return &NullableSiteSettingApMatchingRule{value: val, isSet: true}
}

func (v NullableSiteSettingApMatchingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingApMatchingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


