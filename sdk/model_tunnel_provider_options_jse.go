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

// checks if the TunnelProviderOptionsJse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelProviderOptionsJse{}

// TunnelProviderOptionsJse for jse-ipsec, this allow provisioning of adequate resource on JSE. Make sure adequate licenses are added
type TunnelProviderOptionsJse struct {
	Name *string `json:"name,omitempty"`
	NumUsers *int32 `json:"num_users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelProviderOptionsJse TunnelProviderOptionsJse

// NewTunnelProviderOptionsJse instantiates a new TunnelProviderOptionsJse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelProviderOptionsJse() *TunnelProviderOptionsJse {
	this := TunnelProviderOptionsJse{}
	return &this
}

// NewTunnelProviderOptionsJseWithDefaults instantiates a new TunnelProviderOptionsJse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelProviderOptionsJseWithDefaults() *TunnelProviderOptionsJse {
	this := TunnelProviderOptionsJse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TunnelProviderOptionsJse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsJse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TunnelProviderOptionsJse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TunnelProviderOptionsJse) SetName(v string) {
	o.Name = &v
}

// GetNumUsers returns the NumUsers field value if set, zero value otherwise.
func (o *TunnelProviderOptionsJse) GetNumUsers() int32 {
	if o == nil || IsNil(o.NumUsers) {
		var ret int32
		return ret
	}
	return *o.NumUsers
}

// GetNumUsersOk returns a tuple with the NumUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsJse) GetNumUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.NumUsers) {
		return nil, false
	}
	return o.NumUsers, true
}

// HasNumUsers returns a boolean if a field has been set.
func (o *TunnelProviderOptionsJse) HasNumUsers() bool {
	if o != nil && !IsNil(o.NumUsers) {
		return true
	}

	return false
}

// SetNumUsers gets a reference to the given int32 and assigns it to the NumUsers field.
func (o *TunnelProviderOptionsJse) SetNumUsers(v int32) {
	o.NumUsers = &v
}

func (o TunnelProviderOptionsJse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelProviderOptionsJse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumUsers) {
		toSerialize["num_users"] = o.NumUsers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelProviderOptionsJse) UnmarshalJSON(data []byte) (err error) {
	varTunnelProviderOptionsJse := _TunnelProviderOptionsJse{}

	err = json.Unmarshal(data, &varTunnelProviderOptionsJse)

	if err != nil {
		return err
	}

	*o = TunnelProviderOptionsJse(varTunnelProviderOptionsJse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "num_users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelProviderOptionsJse struct {
	value *TunnelProviderOptionsJse
	isSet bool
}

func (v NullableTunnelProviderOptionsJse) Get() *TunnelProviderOptionsJse {
	return v.value
}

func (v *NullableTunnelProviderOptionsJse) Set(val *TunnelProviderOptionsJse) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelProviderOptionsJse) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelProviderOptionsJse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelProviderOptionsJse(val *TunnelProviderOptionsJse) *NullableTunnelProviderOptionsJse {
	return &NullableTunnelProviderOptionsJse{value: val, isSet: true}
}

func (v NullableTunnelProviderOptionsJse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelProviderOptionsJse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


