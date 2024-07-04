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

// checks if the IdpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpConfig{}

// IdpConfig struct for IdpConfig
type IdpConfig struct {
	AlertOnly *bool `json:"alert_only,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// org_level IDP Profile can be used, this takes precedence over `profile`
	IdpprofileId *string `json:"idpprofile_id,omitempty"`
	// `strict` (default) / `standard` / or keys from from idp_profiles
	Profile *string `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpConfig IdpConfig

// NewIdpConfig instantiates a new IdpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpConfig() *IdpConfig {
	this := IdpConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	var profile string = "strict"
	this.Profile = &profile
	return &this
}

// NewIdpConfigWithDefaults instantiates a new IdpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpConfigWithDefaults() *IdpConfig {
	this := IdpConfig{}
	var enabled bool = false
	this.Enabled = &enabled
	var profile string = "strict"
	this.Profile = &profile
	return &this
}

// GetAlertOnly returns the AlertOnly field value if set, zero value otherwise.
func (o *IdpConfig) GetAlertOnly() bool {
	if o == nil || IsNil(o.AlertOnly) {
		var ret bool
		return ret
	}
	return *o.AlertOnly
}

// GetAlertOnlyOk returns a tuple with the AlertOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConfig) GetAlertOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.AlertOnly) {
		return nil, false
	}
	return o.AlertOnly, true
}

// HasAlertOnly returns a boolean if a field has been set.
func (o *IdpConfig) HasAlertOnly() bool {
	if o != nil && !IsNil(o.AlertOnly) {
		return true
	}

	return false
}

// SetAlertOnly gets a reference to the given bool and assigns it to the AlertOnly field.
func (o *IdpConfig) SetAlertOnly(v bool) {
	o.AlertOnly = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IdpConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IdpConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IdpConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIdpprofileId returns the IdpprofileId field value if set, zero value otherwise.
func (o *IdpConfig) GetIdpprofileId() string {
	if o == nil || IsNil(o.IdpprofileId) {
		var ret string
		return ret
	}
	return *o.IdpprofileId
}

// GetIdpprofileIdOk returns a tuple with the IdpprofileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConfig) GetIdpprofileIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdpprofileId) {
		return nil, false
	}
	return o.IdpprofileId, true
}

// HasIdpprofileId returns a boolean if a field has been set.
func (o *IdpConfig) HasIdpprofileId() bool {
	if o != nil && !IsNil(o.IdpprofileId) {
		return true
	}

	return false
}

// SetIdpprofileId gets a reference to the given string and assigns it to the IdpprofileId field.
func (o *IdpConfig) SetIdpprofileId(v string) {
	o.IdpprofileId = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *IdpConfig) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpConfig) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *IdpConfig) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *IdpConfig) SetProfile(v string) {
	o.Profile = &v
}

func (o IdpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlertOnly) {
		toSerialize["alert_only"] = o.AlertOnly
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.IdpprofileId) {
		toSerialize["idpprofile_id"] = o.IdpprofileId
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdpConfig) UnmarshalJSON(data []byte) (err error) {
	varIdpConfig := _IdpConfig{}

	err = json.Unmarshal(data, &varIdpConfig)

	if err != nil {
		return err
	}

	*o = IdpConfig(varIdpConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alert_only")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "idpprofile_id")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdpConfig struct {
	value *IdpConfig
	isSet bool
}

func (v NullableIdpConfig) Get() *IdpConfig {
	return v.value
}

func (v *NullableIdpConfig) Set(val *IdpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpConfig(val *IdpConfig) *NullableIdpConfig {
	return &NullableIdpConfig{value: val, isSet: true}
}

func (v NullableIdpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


