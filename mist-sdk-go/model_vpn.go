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

// checks if the Vpn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Vpn{}

// Vpn struct for Vpn
type Vpn struct {
	CreatedTime *int32 `json:"created_time,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedTime *int32 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	Paths map[string]VpnPath `json:"paths"`
	AdditionalProperties map[string]interface{}
}

type _Vpn Vpn

// NewVpn instantiates a new Vpn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpn(name string, paths map[string]VpnPath) *Vpn {
	this := Vpn{}
	this.Name = name
	this.Paths = paths
	return &this
}

// NewVpnWithDefaults instantiates a new Vpn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpnWithDefaults() *Vpn {
	this := Vpn{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Vpn) GetCreatedTime() int32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret int32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpn) GetCreatedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Vpn) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given int32 and assigns it to the CreatedTime field.
func (o *Vpn) SetCreatedTime(v int32) {
	o.CreatedTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Vpn) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpn) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Vpn) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Vpn) SetId(v string) {
	o.Id = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Vpn) GetModifiedTime() int32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret int32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpn) GetModifiedTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Vpn) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given int32 and assigns it to the ModifiedTime field.
func (o *Vpn) SetModifiedTime(v int32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *Vpn) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Vpn) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Vpn) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Vpn) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vpn) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Vpn) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Vpn) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPaths returns the Paths field value
func (o *Vpn) GetPaths() map[string]VpnPath {
	if o == nil {
		var ret map[string]VpnPath
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *Vpn) GetPathsOk() (*map[string]VpnPath, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Paths, true
}

// SetPaths sets field value
func (o *Vpn) SetPaths(v map[string]VpnPath) {
	o.Paths = v
}

func (o Vpn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Vpn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	toSerialize["paths"] = o.Paths

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Vpn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"paths",
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

	varVpn := _Vpn{}

	err = json.Unmarshal(data, &varVpn)

	if err != nil {
		return err
	}

	*o = Vpn(varVpn)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "paths")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVpn struct {
	value *Vpn
	isSet bool
}

func (v NullableVpn) Get() *Vpn {
	return v.value
}

func (v *NullableVpn) Set(val *Vpn) {
	v.value = val
	v.isSet = true
}

func (v NullableVpn) IsSet() bool {
	return v.isSet
}

func (v *NullableVpn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpn(val *Vpn) *NullableVpn {
	return &NullableVpn{value: val, isSet: true}
}

func (v NullableVpn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


