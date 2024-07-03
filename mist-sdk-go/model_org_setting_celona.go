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

// checks if the OrgSettingCelona type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingCelona{}

// OrgSettingCelona struct for OrgSettingCelona
type OrgSettingCelona struct {
	ApiKey *string `json:"api_key,omitempty"`
	ApiPrefix *string `json:"api_prefix,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingCelona OrgSettingCelona

// NewOrgSettingCelona instantiates a new OrgSettingCelona object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingCelona() *OrgSettingCelona {
	this := OrgSettingCelona{}
	return &this
}

// NewOrgSettingCelonaWithDefaults instantiates a new OrgSettingCelona object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingCelonaWithDefaults() *OrgSettingCelona {
	this := OrgSettingCelona{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *OrgSettingCelona) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingCelona) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *OrgSettingCelona) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *OrgSettingCelona) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiPrefix returns the ApiPrefix field value if set, zero value otherwise.
func (o *OrgSettingCelona) GetApiPrefix() string {
	if o == nil || IsNil(o.ApiPrefix) {
		var ret string
		return ret
	}
	return *o.ApiPrefix
}

// GetApiPrefixOk returns a tuple with the ApiPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingCelona) GetApiPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ApiPrefix) {
		return nil, false
	}
	return o.ApiPrefix, true
}

// HasApiPrefix returns a boolean if a field has been set.
func (o *OrgSettingCelona) HasApiPrefix() bool {
	if o != nil && !IsNil(o.ApiPrefix) {
		return true
	}

	return false
}

// SetApiPrefix gets a reference to the given string and assigns it to the ApiPrefix field.
func (o *OrgSettingCelona) SetApiPrefix(v string) {
	o.ApiPrefix = &v
}

func (o OrgSettingCelona) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingCelona) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.ApiPrefix) {
		toSerialize["api_prefix"] = o.ApiPrefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingCelona) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingCelona := _OrgSettingCelona{}

	err = json.Unmarshal(data, &varOrgSettingCelona)

	if err != nil {
		return err
	}

	*o = OrgSettingCelona(varOrgSettingCelona)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "api_prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingCelona struct {
	value *OrgSettingCelona
	isSet bool
}

func (v NullableOrgSettingCelona) Get() *OrgSettingCelona {
	return v.value
}

func (v *NullableOrgSettingCelona) Set(val *OrgSettingCelona) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingCelona) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingCelona) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingCelona(val *OrgSettingCelona) *NullableOrgSettingCelona {
	return &NullableOrgSettingCelona{value: val, isSet: true}
}

func (v NullableOrgSettingCelona) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingCelona) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


