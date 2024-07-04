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
	"fmt"
)

// checks if the AccountJseConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountJseConfig{}

// AccountJseConfig struct for AccountJseConfig
type AccountJseConfig struct {
	CloudName *string `json:"cloud_name,omitempty"`
	Password string `json:"password"`
	Username string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _AccountJseConfig AccountJseConfig

// NewAccountJseConfig instantiates a new AccountJseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountJseConfig(password string, username string) *AccountJseConfig {
	this := AccountJseConfig{}
	this.Password = password
	this.Username = username
	return &this
}

// NewAccountJseConfigWithDefaults instantiates a new AccountJseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountJseConfigWithDefaults() *AccountJseConfig {
	this := AccountJseConfig{}
	return &this
}

// GetCloudName returns the CloudName field value if set, zero value otherwise.
func (o *AccountJseConfig) GetCloudName() string {
	if o == nil || IsNil(o.CloudName) {
		var ret string
		return ret
	}
	return *o.CloudName
}

// GetCloudNameOk returns a tuple with the CloudName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJseConfig) GetCloudNameOk() (*string, bool) {
	if o == nil || IsNil(o.CloudName) {
		return nil, false
	}
	return o.CloudName, true
}

// HasCloudName returns a boolean if a field has been set.
func (o *AccountJseConfig) HasCloudName() bool {
	if o != nil && !IsNil(o.CloudName) {
		return true
	}

	return false
}

// SetCloudName gets a reference to the given string and assigns it to the CloudName field.
func (o *AccountJseConfig) SetCloudName(v string) {
	o.CloudName = &v
}

// GetPassword returns the Password field value
func (o *AccountJseConfig) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AccountJseConfig) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AccountJseConfig) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *AccountJseConfig) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AccountJseConfig) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AccountJseConfig) SetUsername(v string) {
	o.Username = v
}

func (o AccountJseConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountJseConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudName) {
		toSerialize["cloud_name"] = o.CloudName
	}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountJseConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"username",
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

	varAccountJseConfig := _AccountJseConfig{}

	err = json.Unmarshal(data, &varAccountJseConfig)

	if err != nil {
		return err
	}

	*o = AccountJseConfig(varAccountJseConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloud_name")
		delete(additionalProperties, "password")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountJseConfig struct {
	value *AccountJseConfig
	isSet bool
}

func (v NullableAccountJseConfig) Get() *AccountJseConfig {
	return v.value
}

func (v *NullableAccountJseConfig) Set(val *AccountJseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountJseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountJseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountJseConfig(val *AccountJseConfig) *NullableAccountJseConfig {
	return &NullableAccountJseConfig{value: val, isSet: true}
}

func (v NullableAccountJseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountJseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


