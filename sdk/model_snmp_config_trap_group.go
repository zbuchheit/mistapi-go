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

// checks if the SnmpConfigTrapGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnmpConfigTrapGroup{}

// SnmpConfigTrapGroup struct for SnmpConfigTrapGroup
type SnmpConfigTrapGroup struct {
	Categories []string `json:"categories,omitempty"`
	// Categories list can refer to https://www.juniper.net/documentation/software/topics/task/configuration/snmp_trap-groups-configuring-junos-nm.html
	GroupName *string `json:"group_name,omitempty"`
	Targets []string `json:"targets,omitempty"`
	Version *SnmpConfigTrapVerion `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpConfigTrapGroup SnmpConfigTrapGroup

// NewSnmpConfigTrapGroup instantiates a new SnmpConfigTrapGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpConfigTrapGroup() *SnmpConfigTrapGroup {
	this := SnmpConfigTrapGroup{}
	var version SnmpConfigTrapVerion = SNMPCONFIGTRAPVERION_V2
	this.Version = &version
	return &this
}

// NewSnmpConfigTrapGroupWithDefaults instantiates a new SnmpConfigTrapGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpConfigTrapGroupWithDefaults() *SnmpConfigTrapGroup {
	this := SnmpConfigTrapGroup{}
	var version SnmpConfigTrapVerion = SNMPCONFIGTRAPVERION_V2
	this.Version = &version
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SnmpConfigTrapGroup) GetCategories() []string {
	if o == nil || IsNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigTrapGroup) GetCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SnmpConfigTrapGroup) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *SnmpConfigTrapGroup) SetCategories(v []string) {
	o.Categories = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *SnmpConfigTrapGroup) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigTrapGroup) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *SnmpConfigTrapGroup) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *SnmpConfigTrapGroup) SetGroupName(v string) {
	o.GroupName = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *SnmpConfigTrapGroup) GetTargets() []string {
	if o == nil || IsNil(o.Targets) {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigTrapGroup) GetTargetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *SnmpConfigTrapGroup) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *SnmpConfigTrapGroup) SetTargets(v []string) {
	o.Targets = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SnmpConfigTrapGroup) GetVersion() SnmpConfigTrapVerion {
	if o == nil || IsNil(o.Version) {
		var ret SnmpConfigTrapVerion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpConfigTrapGroup) GetVersionOk() (*SnmpConfigTrapVerion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SnmpConfigTrapGroup) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given SnmpConfigTrapVerion and assigns it to the Version field.
func (o *SnmpConfigTrapGroup) SetVersion(v SnmpConfigTrapVerion) {
	o.Version = &v
}

func (o SnmpConfigTrapGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnmpConfigTrapGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SnmpConfigTrapGroup) UnmarshalJSON(data []byte) (err error) {
	varSnmpConfigTrapGroup := _SnmpConfigTrapGroup{}

	err = json.Unmarshal(data, &varSnmpConfigTrapGroup)

	if err != nil {
		return err
	}

	*o = SnmpConfigTrapGroup(varSnmpConfigTrapGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "categories")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "targets")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpConfigTrapGroup struct {
	value *SnmpConfigTrapGroup
	isSet bool
}

func (v NullableSnmpConfigTrapGroup) Get() *SnmpConfigTrapGroup {
	return v.value
}

func (v *NullableSnmpConfigTrapGroup) Set(val *SnmpConfigTrapGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpConfigTrapGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpConfigTrapGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpConfigTrapGroup(val *SnmpConfigTrapGroup) *NullableSnmpConfigTrapGroup {
	return &NullableSnmpConfigTrapGroup{value: val, isSet: true}
}

func (v NullableSnmpConfigTrapGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpConfigTrapGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


