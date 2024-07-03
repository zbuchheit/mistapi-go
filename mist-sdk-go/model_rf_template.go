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

// checks if the RfTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RfTemplate{}

// RfTemplate RF Template
type RfTemplate struct {
	AntGain24 *int32 `json:"ant_gain_24,omitempty"`
	AntGain5 *int32 `json:"ant_gain_5,omitempty"`
	AntGain6 *int32 `json:"ant_gain_6,omitempty"`
	Band24 *RftemplateRadioBand24 `json:"band_24,omitempty"`
	Band24Usage *RadioBand24Usage `json:"band_24_usage,omitempty"`
	Band5 *RftemplateRadioBand5 `json:"band_5,omitempty"`
	Band5On24Radio *RftemplateRadioBand5 `json:"band_5_on_24_radio,omitempty"`
	Band6 *RftemplateRadioBand6 `json:"band_6,omitempty"`
	// optional, country code to use. If specified, this gets applied to all sites using the RF Template
	CountryCode *string `json:"country_code,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Id *string `json:"id,omitempty"`
	// overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. \"AP63\")
	ModelSpecific *map[string]RfTemplateModelSpecificProperty `json:"model_specific,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	// The name of the RF template
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	// whether scanning radio is enabled
	ScanningEnabled *bool `json:"scanning_enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RfTemplate RfTemplate

// NewRfTemplate instantiates a new RfTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRfTemplate(name string) *RfTemplate {
	this := RfTemplate{}
	this.Name = name
	return &this
}

// NewRfTemplateWithDefaults instantiates a new RfTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRfTemplateWithDefaults() *RfTemplate {
	this := RfTemplate{}
	return &this
}

// GetAntGain24 returns the AntGain24 field value if set, zero value otherwise.
func (o *RfTemplate) GetAntGain24() int32 {
	if o == nil || IsNil(o.AntGain24) {
		var ret int32
		return ret
	}
	return *o.AntGain24
}

// GetAntGain24Ok returns a tuple with the AntGain24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetAntGain24Ok() (*int32, bool) {
	if o == nil || IsNil(o.AntGain24) {
		return nil, false
	}
	return o.AntGain24, true
}

// HasAntGain24 returns a boolean if a field has been set.
func (o *RfTemplate) HasAntGain24() bool {
	if o != nil && !IsNil(o.AntGain24) {
		return true
	}

	return false
}

// SetAntGain24 gets a reference to the given int32 and assigns it to the AntGain24 field.
func (o *RfTemplate) SetAntGain24(v int32) {
	o.AntGain24 = &v
}

// GetAntGain5 returns the AntGain5 field value if set, zero value otherwise.
func (o *RfTemplate) GetAntGain5() int32 {
	if o == nil || IsNil(o.AntGain5) {
		var ret int32
		return ret
	}
	return *o.AntGain5
}

// GetAntGain5Ok returns a tuple with the AntGain5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetAntGain5Ok() (*int32, bool) {
	if o == nil || IsNil(o.AntGain5) {
		return nil, false
	}
	return o.AntGain5, true
}

// HasAntGain5 returns a boolean if a field has been set.
func (o *RfTemplate) HasAntGain5() bool {
	if o != nil && !IsNil(o.AntGain5) {
		return true
	}

	return false
}

// SetAntGain5 gets a reference to the given int32 and assigns it to the AntGain5 field.
func (o *RfTemplate) SetAntGain5(v int32) {
	o.AntGain5 = &v
}

// GetAntGain6 returns the AntGain6 field value if set, zero value otherwise.
func (o *RfTemplate) GetAntGain6() int32 {
	if o == nil || IsNil(o.AntGain6) {
		var ret int32
		return ret
	}
	return *o.AntGain6
}

// GetAntGain6Ok returns a tuple with the AntGain6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetAntGain6Ok() (*int32, bool) {
	if o == nil || IsNil(o.AntGain6) {
		return nil, false
	}
	return o.AntGain6, true
}

// HasAntGain6 returns a boolean if a field has been set.
func (o *RfTemplate) HasAntGain6() bool {
	if o != nil && !IsNil(o.AntGain6) {
		return true
	}

	return false
}

// SetAntGain6 gets a reference to the given int32 and assigns it to the AntGain6 field.
func (o *RfTemplate) SetAntGain6(v int32) {
	o.AntGain6 = &v
}

// GetBand24 returns the Band24 field value if set, zero value otherwise.
func (o *RfTemplate) GetBand24() RftemplateRadioBand24 {
	if o == nil || IsNil(o.Band24) {
		var ret RftemplateRadioBand24
		return ret
	}
	return *o.Band24
}

// GetBand24Ok returns a tuple with the Band24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetBand24Ok() (*RftemplateRadioBand24, bool) {
	if o == nil || IsNil(o.Band24) {
		return nil, false
	}
	return o.Band24, true
}

// HasBand24 returns a boolean if a field has been set.
func (o *RfTemplate) HasBand24() bool {
	if o != nil && !IsNil(o.Band24) {
		return true
	}

	return false
}

// SetBand24 gets a reference to the given RftemplateRadioBand24 and assigns it to the Band24 field.
func (o *RfTemplate) SetBand24(v RftemplateRadioBand24) {
	o.Band24 = &v
}

// GetBand24Usage returns the Band24Usage field value if set, zero value otherwise.
func (o *RfTemplate) GetBand24Usage() RadioBand24Usage {
	if o == nil || IsNil(o.Band24Usage) {
		var ret RadioBand24Usage
		return ret
	}
	return *o.Band24Usage
}

// GetBand24UsageOk returns a tuple with the Band24Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetBand24UsageOk() (*RadioBand24Usage, bool) {
	if o == nil || IsNil(o.Band24Usage) {
		return nil, false
	}
	return o.Band24Usage, true
}

// HasBand24Usage returns a boolean if a field has been set.
func (o *RfTemplate) HasBand24Usage() bool {
	if o != nil && !IsNil(o.Band24Usage) {
		return true
	}

	return false
}

// SetBand24Usage gets a reference to the given RadioBand24Usage and assigns it to the Band24Usage field.
func (o *RfTemplate) SetBand24Usage(v RadioBand24Usage) {
	o.Band24Usage = &v
}

// GetBand5 returns the Band5 field value if set, zero value otherwise.
func (o *RfTemplate) GetBand5() RftemplateRadioBand5 {
	if o == nil || IsNil(o.Band5) {
		var ret RftemplateRadioBand5
		return ret
	}
	return *o.Band5
}

// GetBand5Ok returns a tuple with the Band5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetBand5Ok() (*RftemplateRadioBand5, bool) {
	if o == nil || IsNil(o.Band5) {
		return nil, false
	}
	return o.Band5, true
}

// HasBand5 returns a boolean if a field has been set.
func (o *RfTemplate) HasBand5() bool {
	if o != nil && !IsNil(o.Band5) {
		return true
	}

	return false
}

// SetBand5 gets a reference to the given RftemplateRadioBand5 and assigns it to the Band5 field.
func (o *RfTemplate) SetBand5(v RftemplateRadioBand5) {
	o.Band5 = &v
}

// GetBand5On24Radio returns the Band5On24Radio field value if set, zero value otherwise.
func (o *RfTemplate) GetBand5On24Radio() RftemplateRadioBand5 {
	if o == nil || IsNil(o.Band5On24Radio) {
		var ret RftemplateRadioBand5
		return ret
	}
	return *o.Band5On24Radio
}

// GetBand5On24RadioOk returns a tuple with the Band5On24Radio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetBand5On24RadioOk() (*RftemplateRadioBand5, bool) {
	if o == nil || IsNil(o.Band5On24Radio) {
		return nil, false
	}
	return o.Band5On24Radio, true
}

// HasBand5On24Radio returns a boolean if a field has been set.
func (o *RfTemplate) HasBand5On24Radio() bool {
	if o != nil && !IsNil(o.Band5On24Radio) {
		return true
	}

	return false
}

// SetBand5On24Radio gets a reference to the given RftemplateRadioBand5 and assigns it to the Band5On24Radio field.
func (o *RfTemplate) SetBand5On24Radio(v RftemplateRadioBand5) {
	o.Band5On24Radio = &v
}

// GetBand6 returns the Band6 field value if set, zero value otherwise.
func (o *RfTemplate) GetBand6() RftemplateRadioBand6 {
	if o == nil || IsNil(o.Band6) {
		var ret RftemplateRadioBand6
		return ret
	}
	return *o.Band6
}

// GetBand6Ok returns a tuple with the Band6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetBand6Ok() (*RftemplateRadioBand6, bool) {
	if o == nil || IsNil(o.Band6) {
		return nil, false
	}
	return o.Band6, true
}

// HasBand6 returns a boolean if a field has been set.
func (o *RfTemplate) HasBand6() bool {
	if o != nil && !IsNil(o.Band6) {
		return true
	}

	return false
}

// SetBand6 gets a reference to the given RftemplateRadioBand6 and assigns it to the Band6 field.
func (o *RfTemplate) SetBand6(v RftemplateRadioBand6) {
	o.Band6 = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *RfTemplate) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *RfTemplate) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *RfTemplate) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *RfTemplate) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *RfTemplate) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *RfTemplate) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *RfTemplate) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *RfTemplate) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *RfTemplate) SetForSite(v bool) {
	o.ForSite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RfTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RfTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RfTemplate) SetId(v string) {
	o.Id = &v
}

// GetModelSpecific returns the ModelSpecific field value if set, zero value otherwise.
func (o *RfTemplate) GetModelSpecific() map[string]RfTemplateModelSpecificProperty {
	if o == nil || IsNil(o.ModelSpecific) {
		var ret map[string]RfTemplateModelSpecificProperty
		return ret
	}
	return *o.ModelSpecific
}

// GetModelSpecificOk returns a tuple with the ModelSpecific field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetModelSpecificOk() (*map[string]RfTemplateModelSpecificProperty, bool) {
	if o == nil || IsNil(o.ModelSpecific) {
		return nil, false
	}
	return o.ModelSpecific, true
}

// HasModelSpecific returns a boolean if a field has been set.
func (o *RfTemplate) HasModelSpecific() bool {
	if o != nil && !IsNil(o.ModelSpecific) {
		return true
	}

	return false
}

// SetModelSpecific gets a reference to the given map[string]RfTemplateModelSpecificProperty and assigns it to the ModelSpecific field.
func (o *RfTemplate) SetModelSpecific(v map[string]RfTemplateModelSpecificProperty) {
	o.ModelSpecific = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *RfTemplate) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *RfTemplate) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *RfTemplate) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *RfTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RfTemplate) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *RfTemplate) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *RfTemplate) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *RfTemplate) SetOrgId(v string) {
	o.OrgId = &v
}

// GetScanningEnabled returns the ScanningEnabled field value if set, zero value otherwise.
func (o *RfTemplate) GetScanningEnabled() bool {
	if o == nil || IsNil(o.ScanningEnabled) {
		var ret bool
		return ret
	}
	return *o.ScanningEnabled
}

// GetScanningEnabledOk returns a tuple with the ScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfTemplate) GetScanningEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScanningEnabled) {
		return nil, false
	}
	return o.ScanningEnabled, true
}

// HasScanningEnabled returns a boolean if a field has been set.
func (o *RfTemplate) HasScanningEnabled() bool {
	if o != nil && !IsNil(o.ScanningEnabled) {
		return true
	}

	return false
}

// SetScanningEnabled gets a reference to the given bool and assigns it to the ScanningEnabled field.
func (o *RfTemplate) SetScanningEnabled(v bool) {
	o.ScanningEnabled = &v
}

func (o RfTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RfTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AntGain24) {
		toSerialize["ant_gain_24"] = o.AntGain24
	}
	if !IsNil(o.AntGain5) {
		toSerialize["ant_gain_5"] = o.AntGain5
	}
	if !IsNil(o.AntGain6) {
		toSerialize["ant_gain_6"] = o.AntGain6
	}
	if !IsNil(o.Band24) {
		toSerialize["band_24"] = o.Band24
	}
	if !IsNil(o.Band24Usage) {
		toSerialize["band_24_usage"] = o.Band24Usage
	}
	if !IsNil(o.Band5) {
		toSerialize["band_5"] = o.Band5
	}
	if !IsNil(o.Band5On24Radio) {
		toSerialize["band_5_on_24_radio"] = o.Band5On24Radio
	}
	if !IsNil(o.Band6) {
		toSerialize["band_6"] = o.Band6
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ModelSpecific) {
		toSerialize["model_specific"] = o.ModelSpecific
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.ScanningEnabled) {
		toSerialize["scanning_enabled"] = o.ScanningEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RfTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varRfTemplate := _RfTemplate{}

	err = json.Unmarshal(data, &varRfTemplate)

	if err != nil {
		return err
	}

	*o = RfTemplate(varRfTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ant_gain_24")
		delete(additionalProperties, "ant_gain_5")
		delete(additionalProperties, "ant_gain_6")
		delete(additionalProperties, "band_24")
		delete(additionalProperties, "band_24_usage")
		delete(additionalProperties, "band_5")
		delete(additionalProperties, "band_5_on_24_radio")
		delete(additionalProperties, "band_6")
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "id")
		delete(additionalProperties, "model_specific")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "scanning_enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRfTemplate struct {
	value *RfTemplate
	isSet bool
}

func (v NullableRfTemplate) Get() *RfTemplate {
	return v.value
}

func (v *NullableRfTemplate) Set(val *RfTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableRfTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableRfTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRfTemplate(val *RfTemplate) *NullableRfTemplate {
	return &NullableRfTemplate{value: val, isSet: true}
}

func (v NullableRfTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRfTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


