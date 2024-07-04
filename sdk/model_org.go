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

// checks if the Org type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Org{}

// Org struct for Org
type Org struct {
	AlarmtemplateId NullableString `json:"alarmtemplate_id,omitempty"`
	AllowMist *bool `json:"allow_mist,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	Id *string `json:"id,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	MspId *string `json:"msp_id,omitempty"`
	// logo uploaded by the MSP with advanced tier, only present if provided
	MspLogoUrl *string `json:"msp_logo_url,omitempty"`
	// name of the msp the org belongs to
	MspName *string `json:"msp_name,omitempty"`
	Name string `json:"name"`
	OrggroupIds []string `json:"orggroup_ids,omitempty"`
	SessionExpiry *float32 `json:"session_expiry,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Org Org

// NewOrg instantiates a new Org object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrg(name string) *Org {
	this := Org{}
	var allowMist bool = true
	this.AllowMist = &allowMist
	this.Name = name
	var sessionExpiry float32 = 1440
	this.SessionExpiry = &sessionExpiry
	return &this
}

// NewOrgWithDefaults instantiates a new Org object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgWithDefaults() *Org {
	this := Org{}
	var allowMist bool = true
	this.AllowMist = &allowMist
	var sessionExpiry float32 = 1440
	this.SessionExpiry = &sessionExpiry
	return &this
}

// GetAlarmtemplateId returns the AlarmtemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Org) GetAlarmtemplateId() string {
	if o == nil || IsNil(o.AlarmtemplateId.Get()) {
		var ret string
		return ret
	}
	return *o.AlarmtemplateId.Get()
}

// GetAlarmtemplateIdOk returns a tuple with the AlarmtemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Org) GetAlarmtemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmtemplateId.Get(), o.AlarmtemplateId.IsSet()
}

// HasAlarmtemplateId returns a boolean if a field has been set.
func (o *Org) HasAlarmtemplateId() bool {
	if o != nil && o.AlarmtemplateId.IsSet() {
		return true
	}

	return false
}

// SetAlarmtemplateId gets a reference to the given NullableString and assigns it to the AlarmtemplateId field.
func (o *Org) SetAlarmtemplateId(v string) {
	o.AlarmtemplateId.Set(&v)
}
// SetAlarmtemplateIdNil sets the value for AlarmtemplateId to be an explicit nil
func (o *Org) SetAlarmtemplateIdNil() {
	o.AlarmtemplateId.Set(nil)
}

// UnsetAlarmtemplateId ensures that no value is present for AlarmtemplateId, not even an explicit nil
func (o *Org) UnsetAlarmtemplateId() {
	o.AlarmtemplateId.Unset()
}

// GetAllowMist returns the AllowMist field value if set, zero value otherwise.
func (o *Org) GetAllowMist() bool {
	if o == nil || IsNil(o.AllowMist) {
		var ret bool
		return ret
	}
	return *o.AllowMist
}

// GetAllowMistOk returns a tuple with the AllowMist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetAllowMistOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowMist) {
		return nil, false
	}
	return o.AllowMist, true
}

// HasAllowMist returns a boolean if a field has been set.
func (o *Org) HasAllowMist() bool {
	if o != nil && !IsNil(o.AllowMist) {
		return true
	}

	return false
}

// SetAllowMist gets a reference to the given bool and assigns it to the AllowMist field.
func (o *Org) SetAllowMist(v bool) {
	o.AllowMist = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Org) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Org) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Org) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Org) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Org) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Org) SetId(v string) {
	o.Id = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Org) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Org) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Org) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMspId returns the MspId field value if set, zero value otherwise.
func (o *Org) GetMspId() string {
	if o == nil || IsNil(o.MspId) {
		var ret string
		return ret
	}
	return *o.MspId
}

// GetMspIdOk returns a tuple with the MspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetMspIdOk() (*string, bool) {
	if o == nil || IsNil(o.MspId) {
		return nil, false
	}
	return o.MspId, true
}

// HasMspId returns a boolean if a field has been set.
func (o *Org) HasMspId() bool {
	if o != nil && !IsNil(o.MspId) {
		return true
	}

	return false
}

// SetMspId gets a reference to the given string and assigns it to the MspId field.
func (o *Org) SetMspId(v string) {
	o.MspId = &v
}

// GetMspLogoUrl returns the MspLogoUrl field value if set, zero value otherwise.
func (o *Org) GetMspLogoUrl() string {
	if o == nil || IsNil(o.MspLogoUrl) {
		var ret string
		return ret
	}
	return *o.MspLogoUrl
}

// GetMspLogoUrlOk returns a tuple with the MspLogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetMspLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MspLogoUrl) {
		return nil, false
	}
	return o.MspLogoUrl, true
}

// HasMspLogoUrl returns a boolean if a field has been set.
func (o *Org) HasMspLogoUrl() bool {
	if o != nil && !IsNil(o.MspLogoUrl) {
		return true
	}

	return false
}

// SetMspLogoUrl gets a reference to the given string and assigns it to the MspLogoUrl field.
func (o *Org) SetMspLogoUrl(v string) {
	o.MspLogoUrl = &v
}

// GetMspName returns the MspName field value if set, zero value otherwise.
func (o *Org) GetMspName() string {
	if o == nil || IsNil(o.MspName) {
		var ret string
		return ret
	}
	return *o.MspName
}

// GetMspNameOk returns a tuple with the MspName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetMspNameOk() (*string, bool) {
	if o == nil || IsNil(o.MspName) {
		return nil, false
	}
	return o.MspName, true
}

// HasMspName returns a boolean if a field has been set.
func (o *Org) HasMspName() bool {
	if o != nil && !IsNil(o.MspName) {
		return true
	}

	return false
}

// SetMspName gets a reference to the given string and assigns it to the MspName field.
func (o *Org) SetMspName(v string) {
	o.MspName = &v
}

// GetName returns the Name field value
func (o *Org) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Org) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Org) SetName(v string) {
	o.Name = v
}

// GetOrggroupIds returns the OrggroupIds field value if set, zero value otherwise.
func (o *Org) GetOrggroupIds() []string {
	if o == nil || IsNil(o.OrggroupIds) {
		var ret []string
		return ret
	}
	return o.OrggroupIds
}

// GetOrggroupIdsOk returns a tuple with the OrggroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetOrggroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrggroupIds) {
		return nil, false
	}
	return o.OrggroupIds, true
}

// HasOrggroupIds returns a boolean if a field has been set.
func (o *Org) HasOrggroupIds() bool {
	if o != nil && !IsNil(o.OrggroupIds) {
		return true
	}

	return false
}

// SetOrggroupIds gets a reference to the given []string and assigns it to the OrggroupIds field.
func (o *Org) SetOrggroupIds(v []string) {
	o.OrggroupIds = v
}

// GetSessionExpiry returns the SessionExpiry field value if set, zero value otherwise.
func (o *Org) GetSessionExpiry() float32 {
	if o == nil || IsNil(o.SessionExpiry) {
		var ret float32
		return ret
	}
	return *o.SessionExpiry
}

// GetSessionExpiryOk returns a tuple with the SessionExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org) GetSessionExpiryOk() (*float32, bool) {
	if o == nil || IsNil(o.SessionExpiry) {
		return nil, false
	}
	return o.SessionExpiry, true
}

// HasSessionExpiry returns a boolean if a field has been set.
func (o *Org) HasSessionExpiry() bool {
	if o != nil && !IsNil(o.SessionExpiry) {
		return true
	}

	return false
}

// SetSessionExpiry gets a reference to the given float32 and assigns it to the SessionExpiry field.
func (o *Org) SetSessionExpiry(v float32) {
	o.SessionExpiry = &v
}

func (o Org) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Org) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AlarmtemplateId.IsSet() {
		toSerialize["alarmtemplate_id"] = o.AlarmtemplateId.Get()
	}
	if !IsNil(o.AllowMist) {
		toSerialize["allow_mist"] = o.AllowMist
	}
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
	if !IsNil(o.MspLogoUrl) {
		toSerialize["msp_logo_url"] = o.MspLogoUrl
	}
	if !IsNil(o.MspName) {
		toSerialize["msp_name"] = o.MspName
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OrggroupIds) {
		toSerialize["orggroup_ids"] = o.OrggroupIds
	}
	if !IsNil(o.SessionExpiry) {
		toSerialize["session_expiry"] = o.SessionExpiry
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Org) UnmarshalJSON(data []byte) (err error) {
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

	varOrg := _Org{}

	err = json.Unmarshal(data, &varOrg)

	if err != nil {
		return err
	}

	*o = Org(varOrg)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alarmtemplate_id")
		delete(additionalProperties, "allow_mist")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "id")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "msp_id")
		delete(additionalProperties, "msp_logo_url")
		delete(additionalProperties, "msp_name")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orggroup_ids")
		delete(additionalProperties, "session_expiry")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrg struct {
	value *Org
	isSet bool
}

func (v NullableOrg) Get() *Org {
	return v.value
}

func (v *NullableOrg) Set(val *Org) {
	v.value = val
	v.isSet = true
}

func (v NullableOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrg(val *Org) *NullableOrg {
	return &NullableOrg{value: val, isSet: true}
}

func (v NullableOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


