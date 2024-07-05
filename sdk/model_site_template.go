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

// checks if the SiteTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteTemplate{}

// SiteTemplate struct for SiteTemplate
type SiteTemplate struct {
	AutoUpgrade *SiteTemplateAutoUpgrade `json:"auto_upgrade,omitempty"`
	Name *string `json:"name,omitempty"`
	// a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
	Vars *map[string]string `json:"vars,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteTemplate SiteTemplate

// NewSiteTemplate instantiates a new SiteTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteTemplate() *SiteTemplate {
	this := SiteTemplate{}
	return &this
}

// NewSiteTemplateWithDefaults instantiates a new SiteTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteTemplateWithDefaults() *SiteTemplate {
	this := SiteTemplate{}
	return &this
}

// GetAutoUpgrade returns the AutoUpgrade field value if set, zero value otherwise.
func (o *SiteTemplate) GetAutoUpgrade() SiteTemplateAutoUpgrade {
	if o == nil || IsNil(o.AutoUpgrade) {
		var ret SiteTemplateAutoUpgrade
		return ret
	}
	return *o.AutoUpgrade
}

// GetAutoUpgradeOk returns a tuple with the AutoUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteTemplate) GetAutoUpgradeOk() (*SiteTemplateAutoUpgrade, bool) {
	if o == nil || IsNil(o.AutoUpgrade) {
		return nil, false
	}
	return o.AutoUpgrade, true
}

// HasAutoUpgrade returns a boolean if a field has been set.
func (o *SiteTemplate) HasAutoUpgrade() bool {
	if o != nil && !IsNil(o.AutoUpgrade) {
		return true
	}

	return false
}

// SetAutoUpgrade gets a reference to the given SiteTemplateAutoUpgrade and assigns it to the AutoUpgrade field.
func (o *SiteTemplate) SetAutoUpgrade(v SiteTemplateAutoUpgrade) {
	o.AutoUpgrade = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SiteTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SiteTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SiteTemplate) SetName(v string) {
	o.Name = &v
}

// GetVars returns the Vars field value if set, zero value otherwise.
func (o *SiteTemplate) GetVars() map[string]string {
	if o == nil || IsNil(o.Vars) {
		var ret map[string]string
		return ret
	}
	return *o.Vars
}

// GetVarsOk returns a tuple with the Vars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteTemplate) GetVarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Vars) {
		return nil, false
	}
	return o.Vars, true
}

// HasVars returns a boolean if a field has been set.
func (o *SiteTemplate) HasVars() bool {
	if o != nil && !IsNil(o.Vars) {
		return true
	}

	return false
}

// SetVars gets a reference to the given map[string]string and assigns it to the Vars field.
func (o *SiteTemplate) SetVars(v map[string]string) {
	o.Vars = &v
}

func (o SiteTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoUpgrade) {
		toSerialize["auto_upgrade"] = o.AutoUpgrade
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Vars) {
		toSerialize["vars"] = o.Vars
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteTemplate) UnmarshalJSON(data []byte) (err error) {
	varSiteTemplate := _SiteTemplate{}

	err = json.Unmarshal(data, &varSiteTemplate)

	if err != nil {
		return err
	}

	*o = SiteTemplate(varSiteTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auto_upgrade")
		delete(additionalProperties, "name")
		delete(additionalProperties, "vars")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteTemplate struct {
	value *SiteTemplate
	isSet bool
}

func (v NullableSiteTemplate) Get() *SiteTemplate {
	return v.value
}

func (v *NullableSiteTemplate) Set(val *SiteTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteTemplate(val *SiteTemplate) *NullableSiteTemplate {
	return &NullableSiteTemplate{value: val, isSet: true}
}

func (v NullableSiteTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


