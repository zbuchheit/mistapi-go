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

// checks if the SsoRoleOrg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoRoleOrg{}

// SsoRoleOrg SSO Role response
type SsoRoleOrg struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	MspId *string `json:"msp_id,omitempty"`
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	Privileges []PrivilegeOrg `json:"privileges"`
	SiteId *string `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SsoRoleOrg SsoRoleOrg

// NewSsoRoleOrg instantiates a new SsoRoleOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoRoleOrg(name string, privileges []PrivilegeOrg) *SsoRoleOrg {
	this := SsoRoleOrg{}
	this.Name = name
	this.Privileges = privileges
	return &this
}

// NewSsoRoleOrgWithDefaults instantiates a new SsoRoleOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoRoleOrgWithDefaults() *SsoRoleOrg {
	this := SsoRoleOrg{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *SsoRoleOrg) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *SsoRoleOrg) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *SsoRoleOrg) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *SsoRoleOrg) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *SsoRoleOrg) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *SsoRoleOrg) SetForSite(v bool) {
	o.ForSite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SsoRoleOrg) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SsoRoleOrg) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SsoRoleOrg) SetId(v string) {
	o.Id = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *SsoRoleOrg) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *SsoRoleOrg) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *SsoRoleOrg) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMspId returns the MspId field value if set, zero value otherwise.
func (o *SsoRoleOrg) GetMspId() string {
	if o == nil || IsNil(o.MspId) {
		var ret string
		return ret
	}
	return *o.MspId
}

// GetMspIdOk returns a tuple with the MspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetMspIdOk() (*string, bool) {
	if o == nil || IsNil(o.MspId) {
		return nil, false
	}
	return o.MspId, true
}

// HasMspId returns a boolean if a field has been set.
func (o *SsoRoleOrg) HasMspId() bool {
	if o != nil && !IsNil(o.MspId) {
		return true
	}

	return false
}

// SetMspId gets a reference to the given string and assigns it to the MspId field.
func (o *SsoRoleOrg) SetMspId(v string) {
	o.MspId = &v
}

// GetName returns the Name field value
func (o *SsoRoleOrg) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SsoRoleOrg) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SsoRoleOrg) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SsoRoleOrg) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *SsoRoleOrg) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPrivileges returns the Privileges field value
func (o *SsoRoleOrg) GetPrivileges() []PrivilegeOrg {
	if o == nil {
		var ret []PrivilegeOrg
		return ret
	}

	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetPrivilegesOk() ([]PrivilegeOrg, bool) {
	if o == nil {
		return nil, false
	}
	return o.Privileges, true
}

// SetPrivileges sets field value
func (o *SsoRoleOrg) SetPrivileges(v []PrivilegeOrg) {
	o.Privileges = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *SsoRoleOrg) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsoRoleOrg) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *SsoRoleOrg) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *SsoRoleOrg) SetSiteId(v string) {
	o.SiteId = &v
}

func (o SsoRoleOrg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoRoleOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
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
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	toSerialize["privileges"] = o.Privileges
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsoRoleOrg) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"privileges",
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

	varSsoRoleOrg := _SsoRoleOrg{}

	err = json.Unmarshal(data, &varSsoRoleOrg)

	if err != nil {
		return err
	}

	*o = SsoRoleOrg(varSsoRoleOrg)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "msp_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "privileges")
		delete(additionalProperties, "site_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSsoRoleOrg struct {
	value *SsoRoleOrg
	isSet bool
}

func (v NullableSsoRoleOrg) Get() *SsoRoleOrg {
	return v.value
}

func (v *NullableSsoRoleOrg) Set(val *SsoRoleOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoRoleOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoRoleOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoRoleOrg(val *SsoRoleOrg) *NullableSsoRoleOrg {
	return &NullableSsoRoleOrg{value: val, isSet: true}
}

func (v NullableSsoRoleOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoRoleOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


