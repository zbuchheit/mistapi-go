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

// checks if the WlanPortalTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanPortalTemplate{}

// WlanPortalTemplate struct for WlanPortalTemplate
type WlanPortalTemplate struct {
	PortalTemplate *WlanPortalTemplateSetting `json:"portal_template,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanPortalTemplate WlanPortalTemplate

// NewWlanPortalTemplate instantiates a new WlanPortalTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanPortalTemplate() *WlanPortalTemplate {
	this := WlanPortalTemplate{}
	return &this
}

// NewWlanPortalTemplateWithDefaults instantiates a new WlanPortalTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanPortalTemplateWithDefaults() *WlanPortalTemplate {
	this := WlanPortalTemplate{}
	return &this
}

// GetPortalTemplate returns the PortalTemplate field value if set, zero value otherwise.
func (o *WlanPortalTemplate) GetPortalTemplate() WlanPortalTemplateSetting {
	if o == nil || IsNil(o.PortalTemplate) {
		var ret WlanPortalTemplateSetting
		return ret
	}
	return *o.PortalTemplate
}

// GetPortalTemplateOk returns a tuple with the PortalTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanPortalTemplate) GetPortalTemplateOk() (*WlanPortalTemplateSetting, bool) {
	if o == nil || IsNil(o.PortalTemplate) {
		return nil, false
	}
	return o.PortalTemplate, true
}

// HasPortalTemplate returns a boolean if a field has been set.
func (o *WlanPortalTemplate) HasPortalTemplate() bool {
	if o != nil && !IsNil(o.PortalTemplate) {
		return true
	}

	return false
}

// SetPortalTemplate gets a reference to the given WlanPortalTemplateSetting and assigns it to the PortalTemplate field.
func (o *WlanPortalTemplate) SetPortalTemplate(v WlanPortalTemplateSetting) {
	o.PortalTemplate = &v
}

func (o WlanPortalTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanPortalTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortalTemplate) {
		toSerialize["portal_template"] = o.PortalTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanPortalTemplate) UnmarshalJSON(data []byte) (err error) {
	varWlanPortalTemplate := _WlanPortalTemplate{}

	err = json.Unmarshal(data, &varWlanPortalTemplate)

	if err != nil {
		return err
	}

	*o = WlanPortalTemplate(varWlanPortalTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "portal_template")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanPortalTemplate struct {
	value *WlanPortalTemplate
	isSet bool
}

func (v NullableWlanPortalTemplate) Get() *WlanPortalTemplate {
	return v.value
}

func (v *NullableWlanPortalTemplate) Set(val *WlanPortalTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanPortalTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanPortalTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanPortalTemplate(val *WlanPortalTemplate) *NullableWlanPortalTemplate {
	return &NullableWlanPortalTemplate{value: val, isSet: true}
}

func (v NullableWlanPortalTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanPortalTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


