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

// checks if the Sdkinvite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sdkinvite{}

// Sdkinvite SDK invite
type Sdkinvite struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ExpireTime *int32 `json:"expire_time,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	// name, will show up in mobile
	Name string `json:"name"`
	OrgId *string `json:"org_id,omitempty"`
	// number of time this invite can be used
	Quota *int32 `json:"quota,omitempty"`
	// whether quota limiting is enabled
	QuotaLimited *bool `json:"quota_limited,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Sdkinvite Sdkinvite

// NewSdkinvite instantiates a new Sdkinvite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkinvite(name string) *Sdkinvite {
	this := Sdkinvite{}
	var enabled bool = true
	this.Enabled = &enabled
	this.Name = name
	var quotaLimited bool = false
	this.QuotaLimited = &quotaLimited
	return &this
}

// NewSdkinviteWithDefaults instantiates a new Sdkinvite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkinviteWithDefaults() *Sdkinvite {
	this := Sdkinvite{}
	var enabled bool = true
	this.Enabled = &enabled
	var quotaLimited bool = false
	this.QuotaLimited = &quotaLimited
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Sdkinvite) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Sdkinvite) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Sdkinvite) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Sdkinvite) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Sdkinvite) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Sdkinvite) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *Sdkinvite) GetExpireTime() int32 {
	if o == nil || IsNil(o.ExpireTime) {
		var ret int32
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetExpireTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *Sdkinvite) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given int32 and assigns it to the ExpireTime field.
func (o *Sdkinvite) SetExpireTime(v int32) {
	o.ExpireTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Sdkinvite) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Sdkinvite) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Sdkinvite) SetId(v string) {
	o.Id = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Sdkinvite) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Sdkinvite) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Sdkinvite) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value
func (o *Sdkinvite) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Sdkinvite) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Sdkinvite) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Sdkinvite) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Sdkinvite) SetOrgId(v string) {
	o.OrgId = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *Sdkinvite) GetQuota() int32 {
	if o == nil || IsNil(o.Quota) {
		var ret int32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetQuotaOk() (*int32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *Sdkinvite) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int32 and assigns it to the Quota field.
func (o *Sdkinvite) SetQuota(v int32) {
	o.Quota = &v
}

// GetQuotaLimited returns the QuotaLimited field value if set, zero value otherwise.
func (o *Sdkinvite) GetQuotaLimited() bool {
	if o == nil || IsNil(o.QuotaLimited) {
		var ret bool
		return ret
	}
	return *o.QuotaLimited
}

// GetQuotaLimitedOk returns a tuple with the QuotaLimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetQuotaLimitedOk() (*bool, bool) {
	if o == nil || IsNil(o.QuotaLimited) {
		return nil, false
	}
	return o.QuotaLimited, true
}

// HasQuotaLimited returns a boolean if a field has been set.
func (o *Sdkinvite) HasQuotaLimited() bool {
	if o != nil && !IsNil(o.QuotaLimited) {
		return true
	}

	return false
}

// SetQuotaLimited gets a reference to the given bool and assigns it to the QuotaLimited field.
func (o *Sdkinvite) SetQuotaLimited(v bool) {
	o.QuotaLimited = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Sdkinvite) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sdkinvite) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Sdkinvite) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *Sdkinvite) SetSiteId(v string) {
	o.SiteId = &v
}

func (o Sdkinvite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sdkinvite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expire_time"] = o.ExpireTime
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
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.QuotaLimited) {
		toSerialize["quota_limited"] = o.QuotaLimited
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Sdkinvite) UnmarshalJSON(data []byte) (err error) {
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

	varSdkinvite := _Sdkinvite{}

	err = json.Unmarshal(data, &varSdkinvite)

	if err != nil {
		return err
	}

	*o = Sdkinvite(varSdkinvite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "expire_time")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "quota")
		delete(additionalProperties, "quota_limited")
		delete(additionalProperties, "site_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdkinvite struct {
	value *Sdkinvite
	isSet bool
}

func (v NullableSdkinvite) Get() *Sdkinvite {
	return v.value
}

func (v *NullableSdkinvite) Set(val *Sdkinvite) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkinvite) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkinvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkinvite(val *Sdkinvite) *NullableSdkinvite {
	return &NullableSdkinvite{value: val, isSet: true}
}

func (v NullableSdkinvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkinvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


