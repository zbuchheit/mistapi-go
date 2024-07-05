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

// checks if the NacTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NacTag{}

// NacTag struct for NacTag
type NacTag struct {
	// can be set to true to allow the override by usermac result
	AllowUsermacOverride *bool `json:"allow_usermac_override,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	// if `type`==`egress_vlan_names`, list of egress vlans to return
	EgressVlanNames []string `json:"egress_vlan_names,omitempty"`
	// if `type`==`gbp_tag`
	GbpTag *int32 `json:"gbp_tag,omitempty"`
	Id *string `json:"id,omitempty"`
	Match *NacTagMatch `json:"match,omitempty"`
	// This field is applicable only when `type`==`match` * `false`: means it is sufficient to match any of the values (i.e., match-any behavior) * `true`: means all values should be matched (i.e., match-all behavior)   Currently it makes sense to set this field to `true` only if the `match`==`idp_role` or `match`==`usermac_label`
	MatchAll *bool `json:"match_all,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	// if `type`==`radius_attrs`, user can specify a list of one or more standard attributes in the field \"radius_attrs\".  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_attrs in the result of a given rule.
	RadiusAttrs []string `json:"radius_attrs,omitempty"`
	// if `type`==`radius_group`
	RadiusGroup *string `json:"radius_group,omitempty"`
	// if `type`==`radius_vendor_attrs`, user can specify a list of one or more vendor-specific attributes in the field \"radius_vendor_attrs\".  It is the responsibility of the user to provide a syntactically correct string, otherwise it may not work as expected. Note that it is allowed to have more than one radius_vendor_attrs in the result of a given rule.
	RadiusVendorAttrs []string `json:"radius_vendor_attrs,omitempty"`
	// if `type`==`session_timeout, in seconds
	SessionTimeout *int32 `json:"session_timeout,omitempty"`
	Type NacTagType `json:"type"`
	// if `type`==`match`
	Values []string `json:"values,omitempty"`
	// if `type`==`vlan`
	Vlan *string `json:"vlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NacTag NacTag

// NewNacTag instantiates a new NacTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNacTag(name string, type_ NacTagType) *NacTag {
	this := NacTag{}
	var allowUsermacOverride bool = false
	this.AllowUsermacOverride = &allowUsermacOverride
	var matchAll bool = false
	this.MatchAll = &matchAll
	this.Name = name
	this.Type = type_
	return &this
}

// NewNacTagWithDefaults instantiates a new NacTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNacTagWithDefaults() *NacTag {
	this := NacTag{}
	var allowUsermacOverride bool = false
	this.AllowUsermacOverride = &allowUsermacOverride
	var matchAll bool = false
	this.MatchAll = &matchAll
	return &this
}

// GetAllowUsermacOverride returns the AllowUsermacOverride field value if set, zero value otherwise.
func (o *NacTag) GetAllowUsermacOverride() bool {
	if o == nil || IsNil(o.AllowUsermacOverride) {
		var ret bool
		return ret
	}
	return *o.AllowUsermacOverride
}

// GetAllowUsermacOverrideOk returns a tuple with the AllowUsermacOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetAllowUsermacOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowUsermacOverride) {
		return nil, false
	}
	return o.AllowUsermacOverride, true
}

// HasAllowUsermacOverride returns a boolean if a field has been set.
func (o *NacTag) HasAllowUsermacOverride() bool {
	if o != nil && !IsNil(o.AllowUsermacOverride) {
		return true
	}

	return false
}

// SetAllowUsermacOverride gets a reference to the given bool and assigns it to the AllowUsermacOverride field.
func (o *NacTag) SetAllowUsermacOverride(v bool) {
	o.AllowUsermacOverride = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *NacTag) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *NacTag) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *NacTag) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetEgressVlanNames returns the EgressVlanNames field value if set, zero value otherwise.
func (o *NacTag) GetEgressVlanNames() []string {
	if o == nil || IsNil(o.EgressVlanNames) {
		var ret []string
		return ret
	}
	return o.EgressVlanNames
}

// GetEgressVlanNamesOk returns a tuple with the EgressVlanNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetEgressVlanNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.EgressVlanNames) {
		return nil, false
	}
	return o.EgressVlanNames, true
}

// HasEgressVlanNames returns a boolean if a field has been set.
func (o *NacTag) HasEgressVlanNames() bool {
	if o != nil && !IsNil(o.EgressVlanNames) {
		return true
	}

	return false
}

// SetEgressVlanNames gets a reference to the given []string and assigns it to the EgressVlanNames field.
func (o *NacTag) SetEgressVlanNames(v []string) {
	o.EgressVlanNames = v
}

// GetGbpTag returns the GbpTag field value if set, zero value otherwise.
func (o *NacTag) GetGbpTag() int32 {
	if o == nil || IsNil(o.GbpTag) {
		var ret int32
		return ret
	}
	return *o.GbpTag
}

// GetGbpTagOk returns a tuple with the GbpTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetGbpTagOk() (*int32, bool) {
	if o == nil || IsNil(o.GbpTag) {
		return nil, false
	}
	return o.GbpTag, true
}

// HasGbpTag returns a boolean if a field has been set.
func (o *NacTag) HasGbpTag() bool {
	if o != nil && !IsNil(o.GbpTag) {
		return true
	}

	return false
}

// SetGbpTag gets a reference to the given int32 and assigns it to the GbpTag field.
func (o *NacTag) SetGbpTag(v int32) {
	o.GbpTag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NacTag) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NacTag) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NacTag) SetId(v string) {
	o.Id = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *NacTag) GetMatch() NacTagMatch {
	if o == nil || IsNil(o.Match) {
		var ret NacTagMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetMatchOk() (*NacTagMatch, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *NacTag) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given NacTagMatch and assigns it to the Match field.
func (o *NacTag) SetMatch(v NacTagMatch) {
	o.Match = &v
}

// GetMatchAll returns the MatchAll field value if set, zero value otherwise.
func (o *NacTag) GetMatchAll() bool {
	if o == nil || IsNil(o.MatchAll) {
		var ret bool
		return ret
	}
	return *o.MatchAll
}

// GetMatchAllOk returns a tuple with the MatchAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetMatchAllOk() (*bool, bool) {
	if o == nil || IsNil(o.MatchAll) {
		return nil, false
	}
	return o.MatchAll, true
}

// HasMatchAll returns a boolean if a field has been set.
func (o *NacTag) HasMatchAll() bool {
	if o != nil && !IsNil(o.MatchAll) {
		return true
	}

	return false
}

// SetMatchAll gets a reference to the given bool and assigns it to the MatchAll field.
func (o *NacTag) SetMatchAll(v bool) {
	o.MatchAll = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *NacTag) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *NacTag) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *NacTag) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *NacTag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NacTag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NacTag) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *NacTag) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *NacTag) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *NacTag) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRadiusAttrs returns the RadiusAttrs field value if set, zero value otherwise.
func (o *NacTag) GetRadiusAttrs() []string {
	if o == nil || IsNil(o.RadiusAttrs) {
		var ret []string
		return ret
	}
	return o.RadiusAttrs
}

// GetRadiusAttrsOk returns a tuple with the RadiusAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetRadiusAttrsOk() ([]string, bool) {
	if o == nil || IsNil(o.RadiusAttrs) {
		return nil, false
	}
	return o.RadiusAttrs, true
}

// HasRadiusAttrs returns a boolean if a field has been set.
func (o *NacTag) HasRadiusAttrs() bool {
	if o != nil && !IsNil(o.RadiusAttrs) {
		return true
	}

	return false
}

// SetRadiusAttrs gets a reference to the given []string and assigns it to the RadiusAttrs field.
func (o *NacTag) SetRadiusAttrs(v []string) {
	o.RadiusAttrs = v
}

// GetRadiusGroup returns the RadiusGroup field value if set, zero value otherwise.
func (o *NacTag) GetRadiusGroup() string {
	if o == nil || IsNil(o.RadiusGroup) {
		var ret string
		return ret
	}
	return *o.RadiusGroup
}

// GetRadiusGroupOk returns a tuple with the RadiusGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetRadiusGroupOk() (*string, bool) {
	if o == nil || IsNil(o.RadiusGroup) {
		return nil, false
	}
	return o.RadiusGroup, true
}

// HasRadiusGroup returns a boolean if a field has been set.
func (o *NacTag) HasRadiusGroup() bool {
	if o != nil && !IsNil(o.RadiusGroup) {
		return true
	}

	return false
}

// SetRadiusGroup gets a reference to the given string and assigns it to the RadiusGroup field.
func (o *NacTag) SetRadiusGroup(v string) {
	o.RadiusGroup = &v
}

// GetRadiusVendorAttrs returns the RadiusVendorAttrs field value if set, zero value otherwise.
func (o *NacTag) GetRadiusVendorAttrs() []string {
	if o == nil || IsNil(o.RadiusVendorAttrs) {
		var ret []string
		return ret
	}
	return o.RadiusVendorAttrs
}

// GetRadiusVendorAttrsOk returns a tuple with the RadiusVendorAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetRadiusVendorAttrsOk() ([]string, bool) {
	if o == nil || IsNil(o.RadiusVendorAttrs) {
		return nil, false
	}
	return o.RadiusVendorAttrs, true
}

// HasRadiusVendorAttrs returns a boolean if a field has been set.
func (o *NacTag) HasRadiusVendorAttrs() bool {
	if o != nil && !IsNil(o.RadiusVendorAttrs) {
		return true
	}

	return false
}

// SetRadiusVendorAttrs gets a reference to the given []string and assigns it to the RadiusVendorAttrs field.
func (o *NacTag) SetRadiusVendorAttrs(v []string) {
	o.RadiusVendorAttrs = v
}

// GetSessionTimeout returns the SessionTimeout field value if set, zero value otherwise.
func (o *NacTag) GetSessionTimeout() int32 {
	if o == nil || IsNil(o.SessionTimeout) {
		var ret int32
		return ret
	}
	return *o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetSessionTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.SessionTimeout) {
		return nil, false
	}
	return o.SessionTimeout, true
}

// HasSessionTimeout returns a boolean if a field has been set.
func (o *NacTag) HasSessionTimeout() bool {
	if o != nil && !IsNil(o.SessionTimeout) {
		return true
	}

	return false
}

// SetSessionTimeout gets a reference to the given int32 and assigns it to the SessionTimeout field.
func (o *NacTag) SetSessionTimeout(v int32) {
	o.SessionTimeout = &v
}

// GetType returns the Type field value
func (o *NacTag) GetType() NacTagType {
	if o == nil {
		var ret NacTagType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NacTag) GetTypeOk() (*NacTagType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NacTag) SetType(v NacTagType) {
	o.Type = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *NacTag) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *NacTag) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *NacTag) SetValues(v []string) {
	o.Values = v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *NacTag) GetVlan() string {
	if o == nil || IsNil(o.Vlan) {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NacTag) GetVlanOk() (*string, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *NacTag) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *NacTag) SetVlan(v string) {
	o.Vlan = &v
}

func (o NacTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NacTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowUsermacOverride) {
		toSerialize["allow_usermac_override"] = o.AllowUsermacOverride
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.EgressVlanNames) {
		toSerialize["egress_vlan_names"] = o.EgressVlanNames
	}
	if !IsNil(o.GbpTag) {
		toSerialize["gbp_tag"] = o.GbpTag
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.MatchAll) {
		toSerialize["match_all"] = o.MatchAll
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.RadiusAttrs) {
		toSerialize["radius_attrs"] = o.RadiusAttrs
	}
	if !IsNil(o.RadiusGroup) {
		toSerialize["radius_group"] = o.RadiusGroup
	}
	if !IsNil(o.RadiusVendorAttrs) {
		toSerialize["radius_vendor_attrs"] = o.RadiusVendorAttrs
	}
	if !IsNil(o.SessionTimeout) {
		toSerialize["session_timeout"] = o.SessionTimeout
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NacTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varNacTag := _NacTag{}

	err = json.Unmarshal(data, &varNacTag)

	if err != nil {
		return err
	}

	*o = NacTag(varNacTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allow_usermac_override")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "egress_vlan_names")
		delete(additionalProperties, "gbp_tag")
		delete(additionalProperties, "id")
		delete(additionalProperties, "match")
		delete(additionalProperties, "match_all")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "radius_attrs")
		delete(additionalProperties, "radius_group")
		delete(additionalProperties, "radius_vendor_attrs")
		delete(additionalProperties, "session_timeout")
		delete(additionalProperties, "type")
		delete(additionalProperties, "values")
		delete(additionalProperties, "vlan")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNacTag struct {
	value *NacTag
	isSet bool
}

func (v NullableNacTag) Get() *NacTag {
	return v.value
}

func (v *NullableNacTag) Set(val *NacTag) {
	v.value = val
	v.isSet = true
}

func (v NullableNacTag) IsSet() bool {
	return v.isSet
}

func (v *NullableNacTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNacTag(val *NacTag) *NullableNacTag {
	return &NullableNacTag{value: val, isSet: true}
}

func (v NullableNacTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNacTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


