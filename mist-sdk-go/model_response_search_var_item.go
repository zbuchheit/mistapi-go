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
)

// checks if the ResponseSearchVarItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSearchVarItem{}

// ResponseSearchVarItem struct for ResponseSearchVarItem
type ResponseSearchVarItem struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	Src *string `json:"src,omitempty"`
	Var *string `json:"var,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSearchVarItem ResponseSearchVarItem

// NewResponseSearchVarItem instantiates a new ResponseSearchVarItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSearchVarItem() *ResponseSearchVarItem {
	this := ResponseSearchVarItem{}
	return &this
}

// NewResponseSearchVarItemWithDefaults instantiates a new ResponseSearchVarItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSearchVarItemWithDefaults() *ResponseSearchVarItem {
	this := ResponseSearchVarItem{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ResponseSearchVarItem) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchVarItem) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ResponseSearchVarItem) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *ResponseSearchVarItem) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *ResponseSearchVarItem) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchVarItem) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *ResponseSearchVarItem) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *ResponseSearchVarItem) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ResponseSearchVarItem) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchVarItem) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ResponseSearchVarItem) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *ResponseSearchVarItem) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *ResponseSearchVarItem) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchVarItem) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *ResponseSearchVarItem) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *ResponseSearchVarItem) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *ResponseSearchVarItem) GetSrc() string {
	if o == nil || IsNil(o.Src) {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchVarItem) GetSrcOk() (*string, bool) {
	if o == nil || IsNil(o.Src) {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *ResponseSearchVarItem) HasSrc() bool {
	if o != nil && !IsNil(o.Src) {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *ResponseSearchVarItem) SetSrc(v string) {
	o.Src = &v
}

// GetVar returns the Var field value if set, zero value otherwise.
func (o *ResponseSearchVarItem) GetVar() string {
	if o == nil || IsNil(o.Var) {
		var ret string
		return ret
	}
	return *o.Var
}

// GetVarOk returns a tuple with the Var field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSearchVarItem) GetVarOk() (*string, bool) {
	if o == nil || IsNil(o.Var) {
		return nil, false
	}
	return o.Var, true
}

// HasVar returns a boolean if a field has been set.
func (o *ResponseSearchVarItem) HasVar() bool {
	if o != nil && !IsNil(o.Var) {
		return true
	}

	return false
}

// SetVar gets a reference to the given string and assigns it to the Var field.
func (o *ResponseSearchVarItem) SetVar(v string) {
	o.Var = &v
}

func (o ResponseSearchVarItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSearchVarItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Src) {
		toSerialize["src"] = o.Src
	}
	if !IsNil(o.Var) {
		toSerialize["var"] = o.Var
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSearchVarItem) UnmarshalJSON(data []byte) (err error) {
	varResponseSearchVarItem := _ResponseSearchVarItem{}

	err = json.Unmarshal(data, &varResponseSearchVarItem)

	if err != nil {
		return err
	}

	*o = ResponseSearchVarItem(varResponseSearchVarItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "src")
		delete(additionalProperties, "var")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSearchVarItem struct {
	value *ResponseSearchVarItem
	isSet bool
}

func (v NullableResponseSearchVarItem) Get() *ResponseSearchVarItem {
	return v.value
}

func (v *NullableResponseSearchVarItem) Set(val *ResponseSearchVarItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSearchVarItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSearchVarItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSearchVarItem(val *ResponseSearchVarItem) *NullableResponseSearchVarItem {
	return &NullableResponseSearchVarItem{value: val, isSet: true}
}

func (v NullableResponseSearchVarItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSearchVarItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


