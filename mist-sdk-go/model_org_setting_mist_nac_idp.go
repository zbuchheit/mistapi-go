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

// checks if the OrgSettingMistNacIdp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingMistNacIdp{}

// OrgSettingMistNacIdp struct for OrgSettingMistNacIdp
type OrgSettingMistNacIdp struct {
	// when the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org
	ExcludeRealms []string `json:"exclude_realms,omitempty"`
	Id *string `json:"id,omitempty"`
	// which realm should trigger this IDP. we extract user realm from 1. Username-AVP (`mist.com` from john@mist.com) 2. Cert CN
	UserRealms []string `json:"user_realms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingMistNacIdp OrgSettingMistNacIdp

// NewOrgSettingMistNacIdp instantiates a new OrgSettingMistNacIdp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingMistNacIdp() *OrgSettingMistNacIdp {
	this := OrgSettingMistNacIdp{}
	return &this
}

// NewOrgSettingMistNacIdpWithDefaults instantiates a new OrgSettingMistNacIdp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingMistNacIdpWithDefaults() *OrgSettingMistNacIdp {
	this := OrgSettingMistNacIdp{}
	return &this
}

// GetExcludeRealms returns the ExcludeRealms field value if set, zero value otherwise.
func (o *OrgSettingMistNacIdp) GetExcludeRealms() []string {
	if o == nil || IsNil(o.ExcludeRealms) {
		var ret []string
		return ret
	}
	return o.ExcludeRealms
}

// GetExcludeRealmsOk returns a tuple with the ExcludeRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingMistNacIdp) GetExcludeRealmsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeRealms) {
		return nil, false
	}
	return o.ExcludeRealms, true
}

// HasExcludeRealms returns a boolean if a field has been set.
func (o *OrgSettingMistNacIdp) HasExcludeRealms() bool {
	if o != nil && !IsNil(o.ExcludeRealms) {
		return true
	}

	return false
}

// SetExcludeRealms gets a reference to the given []string and assigns it to the ExcludeRealms field.
func (o *OrgSettingMistNacIdp) SetExcludeRealms(v []string) {
	o.ExcludeRealms = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrgSettingMistNacIdp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingMistNacIdp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgSettingMistNacIdp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgSettingMistNacIdp) SetId(v string) {
	o.Id = &v
}

// GetUserRealms returns the UserRealms field value if set, zero value otherwise.
func (o *OrgSettingMistNacIdp) GetUserRealms() []string {
	if o == nil || IsNil(o.UserRealms) {
		var ret []string
		return ret
	}
	return o.UserRealms
}

// GetUserRealmsOk returns a tuple with the UserRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingMistNacIdp) GetUserRealmsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserRealms) {
		return nil, false
	}
	return o.UserRealms, true
}

// HasUserRealms returns a boolean if a field has been set.
func (o *OrgSettingMistNacIdp) HasUserRealms() bool {
	if o != nil && !IsNil(o.UserRealms) {
		return true
	}

	return false
}

// SetUserRealms gets a reference to the given []string and assigns it to the UserRealms field.
func (o *OrgSettingMistNacIdp) SetUserRealms(v []string) {
	o.UserRealms = v
}

func (o OrgSettingMistNacIdp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingMistNacIdp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExcludeRealms) {
		toSerialize["exclude_realms"] = o.ExcludeRealms
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserRealms) {
		toSerialize["user_realms"] = o.UserRealms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingMistNacIdp) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingMistNacIdp := _OrgSettingMistNacIdp{}

	err = json.Unmarshal(data, &varOrgSettingMistNacIdp)

	if err != nil {
		return err
	}

	*o = OrgSettingMistNacIdp(varOrgSettingMistNacIdp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude_realms")
		delete(additionalProperties, "id")
		delete(additionalProperties, "user_realms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingMistNacIdp struct {
	value *OrgSettingMistNacIdp
	isSet bool
}

func (v NullableOrgSettingMistNacIdp) Get() *OrgSettingMistNacIdp {
	return v.value
}

func (v *NullableOrgSettingMistNacIdp) Set(val *OrgSettingMistNacIdp) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingMistNacIdp) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingMistNacIdp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingMistNacIdp(val *OrgSettingMistNacIdp) *NullableOrgSettingMistNacIdp {
	return &NullableOrgSettingMistNacIdp{value: val, isSet: true}
}

func (v NullableOrgSettingMistNacIdp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingMistNacIdp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


