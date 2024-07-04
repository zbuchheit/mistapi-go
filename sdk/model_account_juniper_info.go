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

// checks if the AccountJuniperInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountJuniperInfo{}

// AccountJuniperInfo struct for AccountJuniperInfo
type AccountJuniperInfo struct {
	Accounts []JuniperAccount `json:"accounts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountJuniperInfo AccountJuniperInfo

// NewAccountJuniperInfo instantiates a new AccountJuniperInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountJuniperInfo() *AccountJuniperInfo {
	this := AccountJuniperInfo{}
	return &this
}

// NewAccountJuniperInfoWithDefaults instantiates a new AccountJuniperInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountJuniperInfoWithDefaults() *AccountJuniperInfo {
	this := AccountJuniperInfo{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *AccountJuniperInfo) GetAccounts() []JuniperAccount {
	if o == nil || IsNil(o.Accounts) {
		var ret []JuniperAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJuniperInfo) GetAccountsOk() ([]JuniperAccount, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *AccountJuniperInfo) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []JuniperAccount and assigns it to the Accounts field.
func (o *AccountJuniperInfo) SetAccounts(v []JuniperAccount) {
	o.Accounts = v
}

func (o AccountJuniperInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountJuniperInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountJuniperInfo) UnmarshalJSON(data []byte) (err error) {
	varAccountJuniperInfo := _AccountJuniperInfo{}

	err = json.Unmarshal(data, &varAccountJuniperInfo)

	if err != nil {
		return err
	}

	*o = AccountJuniperInfo(varAccountJuniperInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountJuniperInfo struct {
	value *AccountJuniperInfo
	isSet bool
}

func (v NullableAccountJuniperInfo) Get() *AccountJuniperInfo {
	return v.value
}

func (v *NullableAccountJuniperInfo) Set(val *AccountJuniperInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountJuniperInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountJuniperInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountJuniperInfo(val *AccountJuniperInfo) *NullableAccountJuniperInfo {
	return &NullableAccountJuniperInfo{value: val, isSet: true}
}

func (v NullableAccountJuniperInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountJuniperInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


