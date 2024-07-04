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

// checks if the DiscoveredSwitchMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveredSwitchMetric{}

// DiscoveredSwitchMetric struct for DiscoveredSwitchMetric
type DiscoveredSwitchMetric struct {
	Adopted *bool `json:"adopted,omitempty"`
	Aps []DiscoveredSwitchMetricAp `json:"aps,omitempty"`
	ChassisId []string `json:"chassis_id,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	MgmtAddr *string `json:"mgmt_addr,omitempty"`
	Model *string `json:"model,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Score *int32 `json:"score,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	SystemDesc *string `json:"system_desc,omitempty"`
	SystemName *string `json:"system_name,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveredSwitchMetric DiscoveredSwitchMetric

// NewDiscoveredSwitchMetric instantiates a new DiscoveredSwitchMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredSwitchMetric() *DiscoveredSwitchMetric {
	this := DiscoveredSwitchMetric{}
	return &this
}

// NewDiscoveredSwitchMetricWithDefaults instantiates a new DiscoveredSwitchMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredSwitchMetricWithDefaults() *DiscoveredSwitchMetric {
	this := DiscoveredSwitchMetric{}
	return &this
}

// GetAdopted returns the Adopted field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetAdopted() bool {
	if o == nil || IsNil(o.Adopted) {
		var ret bool
		return ret
	}
	return *o.Adopted
}

// GetAdoptedOk returns a tuple with the Adopted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetAdoptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Adopted) {
		return nil, false
	}
	return o.Adopted, true
}

// HasAdopted returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasAdopted() bool {
	if o != nil && !IsNil(o.Adopted) {
		return true
	}

	return false
}

// SetAdopted gets a reference to the given bool and assigns it to the Adopted field.
func (o *DiscoveredSwitchMetric) SetAdopted(v bool) {
	o.Adopted = &v
}

// GetAps returns the Aps field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetAps() []DiscoveredSwitchMetricAp {
	if o == nil || IsNil(o.Aps) {
		var ret []DiscoveredSwitchMetricAp
		return ret
	}
	return o.Aps
}

// GetApsOk returns a tuple with the Aps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetApsOk() ([]DiscoveredSwitchMetricAp, bool) {
	if o == nil || IsNil(o.Aps) {
		return nil, false
	}
	return o.Aps, true
}

// HasAps returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasAps() bool {
	if o != nil && !IsNil(o.Aps) {
		return true
	}

	return false
}

// SetAps gets a reference to the given []DiscoveredSwitchMetricAp and assigns it to the Aps field.
func (o *DiscoveredSwitchMetric) SetAps(v []DiscoveredSwitchMetricAp) {
	o.Aps = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetChassisId() []string {
	if o == nil || IsNil(o.ChassisId) {
		var ret []string
		return ret
	}
	return o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetChassisIdOk() ([]string, bool) {
	if o == nil || IsNil(o.ChassisId) {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasChassisId() bool {
	if o != nil && !IsNil(o.ChassisId) {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given []string and assigns it to the ChassisId field.
func (o *DiscoveredSwitchMetric) SetChassisId(v []string) {
	o.ChassisId = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DiscoveredSwitchMetric) SetHostname(v string) {
	o.Hostname = &v
}

// GetMgmtAddr returns the MgmtAddr field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetMgmtAddr() string {
	if o == nil || IsNil(o.MgmtAddr) {
		var ret string
		return ret
	}
	return *o.MgmtAddr
}

// GetMgmtAddrOk returns a tuple with the MgmtAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetMgmtAddrOk() (*string, bool) {
	if o == nil || IsNil(o.MgmtAddr) {
		return nil, false
	}
	return o.MgmtAddr, true
}

// HasMgmtAddr returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasMgmtAddr() bool {
	if o != nil && !IsNil(o.MgmtAddr) {
		return true
	}

	return false
}

// SetMgmtAddr gets a reference to the given string and assigns it to the MgmtAddr field.
func (o *DiscoveredSwitchMetric) SetMgmtAddr(v string) {
	o.MgmtAddr = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DiscoveredSwitchMetric) SetModel(v string) {
	o.Model = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *DiscoveredSwitchMetric) SetOrgId(v string) {
	o.OrgId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *DiscoveredSwitchMetric) SetScope(v string) {
	o.Scope = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *DiscoveredSwitchMetric) SetScore(v int32) {
	o.Score = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DiscoveredSwitchMetric) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSystemDesc returns the SystemDesc field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetSystemDesc() string {
	if o == nil || IsNil(o.SystemDesc) {
		var ret string
		return ret
	}
	return *o.SystemDesc
}

// GetSystemDescOk returns a tuple with the SystemDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetSystemDescOk() (*string, bool) {
	if o == nil || IsNil(o.SystemDesc) {
		return nil, false
	}
	return o.SystemDesc, true
}

// HasSystemDesc returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasSystemDesc() bool {
	if o != nil && !IsNil(o.SystemDesc) {
		return true
	}

	return false
}

// SetSystemDesc gets a reference to the given string and assigns it to the SystemDesc field.
func (o *DiscoveredSwitchMetric) SetSystemDesc(v string) {
	o.SystemDesc = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetSystemName() string {
	if o == nil || IsNil(o.SystemName) {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetSystemNameOk() (*string, bool) {
	if o == nil || IsNil(o.SystemName) {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasSystemName() bool {
	if o != nil && !IsNil(o.SystemName) {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *DiscoveredSwitchMetric) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *DiscoveredSwitchMetric) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DiscoveredSwitchMetric) SetType(v string) {
	o.Type = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *DiscoveredSwitchMetric) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DiscoveredSwitchMetric) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredSwitchMetric) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DiscoveredSwitchMetric) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DiscoveredSwitchMetric) SetVersion(v string) {
	o.Version = &v
}

func (o DiscoveredSwitchMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveredSwitchMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adopted) {
		toSerialize["adopted"] = o.Adopted
	}
	if !IsNil(o.Aps) {
		toSerialize["aps"] = o.Aps
	}
	if !IsNil(o.ChassisId) {
		toSerialize["chassis_id"] = o.ChassisId
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.MgmtAddr) {
		toSerialize["mgmt_addr"] = o.MgmtAddr
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.SystemDesc) {
		toSerialize["system_desc"] = o.SystemDesc
	}
	if !IsNil(o.SystemName) {
		toSerialize["system_name"] = o.SystemName
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveredSwitchMetric) UnmarshalJSON(data []byte) (err error) {
	varDiscoveredSwitchMetric := _DiscoveredSwitchMetric{}

	err = json.Unmarshal(data, &varDiscoveredSwitchMetric)

	if err != nil {
		return err
	}

	*o = DiscoveredSwitchMetric(varDiscoveredSwitchMetric)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "adopted")
		delete(additionalProperties, "aps")
		delete(additionalProperties, "chassis_id")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "mgmt_addr")
		delete(additionalProperties, "model")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "score")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "system_desc")
		delete(additionalProperties, "system_name")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveredSwitchMetric struct {
	value *DiscoveredSwitchMetric
	isSet bool
}

func (v NullableDiscoveredSwitchMetric) Get() *DiscoveredSwitchMetric {
	return v.value
}

func (v *NullableDiscoveredSwitchMetric) Set(val *DiscoveredSwitchMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveredSwitchMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveredSwitchMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveredSwitchMetric(val *DiscoveredSwitchMetric) *NullableDiscoveredSwitchMetric {
	return &NullableDiscoveredSwitchMetric{value: val, isSet: true}
}

func (v NullableDiscoveredSwitchMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveredSwitchMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


