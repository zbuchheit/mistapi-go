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
	"fmt"
)

// checks if the Orggroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Orggroup{}

// Orggroup Organizations Group
type Orggroup struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	MspId *string `json:"msp_id,omitempty"`
	Name string `json:"name"`
	OrgIds []string `json:"org_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Orggroup Orggroup

// NewOrggroup instantiates a new Orggroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrggroup(name string) *Orggroup {
	this := Orggroup{}
	this.Name = name
	return &this
}

// NewOrggroupWithDefaults instantiates a new Orggroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrggroupWithDefaults() *Orggroup {
	this := Orggroup{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Orggroup) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orggroup) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Orggroup) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Orggroup) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Orggroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orggroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Orggroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Orggroup) SetId(v string) {
	o.Id = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Orggroup) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orggroup) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Orggroup) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Orggroup) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMspId returns the MspId field value if set, zero value otherwise.
func (o *Orggroup) GetMspId() string {
	if o == nil || IsNil(o.MspId) {
		var ret string
		return ret
	}
	return *o.MspId
}

// GetMspIdOk returns a tuple with the MspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orggroup) GetMspIdOk() (*string, bool) {
	if o == nil || IsNil(o.MspId) {
		return nil, false
	}
	return o.MspId, true
}

// HasMspId returns a boolean if a field has been set.
func (o *Orggroup) HasMspId() bool {
	if o != nil && !IsNil(o.MspId) {
		return true
	}

	return false
}

// SetMspId gets a reference to the given string and assigns it to the MspId field.
func (o *Orggroup) SetMspId(v string) {
	o.MspId = &v
}

// GetName returns the Name field value
func (o *Orggroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Orggroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Orggroup) SetName(v string) {
	o.Name = v
}

// GetOrgIds returns the OrgIds field value if set, zero value otherwise.
func (o *Orggroup) GetOrgIds() []string {
	if o == nil || IsNil(o.OrgIds) {
		var ret []string
		return ret
	}
	return o.OrgIds
}

// GetOrgIdsOk returns a tuple with the OrgIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Orggroup) GetOrgIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrgIds) {
		return nil, false
	}
	return o.OrgIds, true
}

// HasOrgIds returns a boolean if a field has been set.
func (o *Orggroup) HasOrgIds() bool {
	if o != nil && !IsNil(o.OrgIds) {
		return true
	}

	return false
}

// SetOrgIds gets a reference to the given []string and assigns it to the OrgIds field.
func (o *Orggroup) SetOrgIds(v []string) {
	o.OrgIds = v
}

func (o Orggroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Orggroup) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MspId) {
		toSerialize["msp_id"] = o.MspId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgIds) {
		toSerialize["org_ids"] = o.OrgIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Orggroup) UnmarshalJSON(data []byte) (err error) {
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

	varOrggroup := _Orggroup{}

	err = json.Unmarshal(data, &varOrggroup)

	if err != nil {
		return err
	}

	*o = Orggroup(varOrggroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "msp_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrggroup struct {
	value *Orggroup
	isSet bool
}

func (v NullableOrggroup) Get() *Orggroup {
	return v.value
}

func (v *NullableOrggroup) Set(val *Orggroup) {
	v.value = val
	v.isSet = true
}

func (v NullableOrggroup) IsSet() bool {
	return v.isSet
}

func (v *NullableOrggroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrggroup(val *Orggroup) *NullableOrggroup {
	return &NullableOrggroup{value: val, isSet: true}
}

func (v NullableOrggroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrggroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


