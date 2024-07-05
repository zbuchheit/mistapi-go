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

// checks if the AccountZscalerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountZscalerInfo{}

// AccountZscalerInfo OAuth linked Zscaler apps account details
type AccountZscalerInfo struct {
	CloudName *string `json:"cloud_name,omitempty"`
	PartnerKey *string `json:"partner_key,omitempty"`
	// customer account user name
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountZscalerInfo AccountZscalerInfo

// NewAccountZscalerInfo instantiates a new AccountZscalerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountZscalerInfo() *AccountZscalerInfo {
	this := AccountZscalerInfo{}
	return &this
}

// NewAccountZscalerInfoWithDefaults instantiates a new AccountZscalerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountZscalerInfoWithDefaults() *AccountZscalerInfo {
	this := AccountZscalerInfo{}
	return &this
}

// GetCloudName returns the CloudName field value if set, zero value otherwise.
func (o *AccountZscalerInfo) GetCloudName() string {
	if o == nil || IsNil(o.CloudName) {
		var ret string
		return ret
	}
	return *o.CloudName
}

// GetCloudNameOk returns a tuple with the CloudName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountZscalerInfo) GetCloudNameOk() (*string, bool) {
	if o == nil || IsNil(o.CloudName) {
		return nil, false
	}
	return o.CloudName, true
}

// HasCloudName returns a boolean if a field has been set.
func (o *AccountZscalerInfo) HasCloudName() bool {
	if o != nil && !IsNil(o.CloudName) {
		return true
	}

	return false
}

// SetCloudName gets a reference to the given string and assigns it to the CloudName field.
func (o *AccountZscalerInfo) SetCloudName(v string) {
	o.CloudName = &v
}

// GetPartnerKey returns the PartnerKey field value if set, zero value otherwise.
func (o *AccountZscalerInfo) GetPartnerKey() string {
	if o == nil || IsNil(o.PartnerKey) {
		var ret string
		return ret
	}
	return *o.PartnerKey
}

// GetPartnerKeyOk returns a tuple with the PartnerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountZscalerInfo) GetPartnerKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerKey) {
		return nil, false
	}
	return o.PartnerKey, true
}

// HasPartnerKey returns a boolean if a field has been set.
func (o *AccountZscalerInfo) HasPartnerKey() bool {
	if o != nil && !IsNil(o.PartnerKey) {
		return true
	}

	return false
}

// SetPartnerKey gets a reference to the given string and assigns it to the PartnerKey field.
func (o *AccountZscalerInfo) SetPartnerKey(v string) {
	o.PartnerKey = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AccountZscalerInfo) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountZscalerInfo) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AccountZscalerInfo) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AccountZscalerInfo) SetUsername(v string) {
	o.Username = &v
}

func (o AccountZscalerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountZscalerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudName) {
		toSerialize["cloud_name"] = o.CloudName
	}
	if !IsNil(o.PartnerKey) {
		toSerialize["partner_key"] = o.PartnerKey
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountZscalerInfo) UnmarshalJSON(data []byte) (err error) {
	varAccountZscalerInfo := _AccountZscalerInfo{}

	err = json.Unmarshal(data, &varAccountZscalerInfo)

	if err != nil {
		return err
	}

	*o = AccountZscalerInfo(varAccountZscalerInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloud_name")
		delete(additionalProperties, "partner_key")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountZscalerInfo struct {
	value *AccountZscalerInfo
	isSet bool
}

func (v NullableAccountZscalerInfo) Get() *AccountZscalerInfo {
	return v.value
}

func (v *NullableAccountZscalerInfo) Set(val *AccountZscalerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountZscalerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountZscalerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountZscalerInfo(val *AccountZscalerInfo) *NullableAccountZscalerInfo {
	return &NullableAccountZscalerInfo{value: val, isSet: true}
}

func (v NullableAccountZscalerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountZscalerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


