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

// checks if the GatewaySearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewaySearch{}

// GatewaySearch struct for GatewaySearch
type GatewaySearch struct {
	ExtIp *string `json:"ext_ip,omitempty"`
	Hostname []string `json:"hostname,omitempty"`
	Ip *string `json:"ip,omitempty"`
	LastHostname *string `json:"last_hostname,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Model *string `json:"model,omitempty"`
	NumMembers *int32 `json:"num_members,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	Timestamp *float32 `json:"timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
	Uptime *int32 `json:"uptime,omitempty"`
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewaySearch GatewaySearch

// NewGatewaySearch instantiates a new GatewaySearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewaySearch() *GatewaySearch {
	this := GatewaySearch{}
	return &this
}

// NewGatewaySearchWithDefaults instantiates a new GatewaySearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewaySearchWithDefaults() *GatewaySearch {
	this := GatewaySearch{}
	return &this
}

// GetExtIp returns the ExtIp field value if set, zero value otherwise.
func (o *GatewaySearch) GetExtIp() string {
	if o == nil || IsNil(o.ExtIp) {
		var ret string
		return ret
	}
	return *o.ExtIp
}

// GetExtIpOk returns a tuple with the ExtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetExtIpOk() (*string, bool) {
	if o == nil || IsNil(o.ExtIp) {
		return nil, false
	}
	return o.ExtIp, true
}

// HasExtIp returns a boolean if a field has been set.
func (o *GatewaySearch) HasExtIp() bool {
	if o != nil && !IsNil(o.ExtIp) {
		return true
	}

	return false
}

// SetExtIp gets a reference to the given string and assigns it to the ExtIp field.
func (o *GatewaySearch) SetExtIp(v string) {
	o.ExtIp = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GatewaySearch) GetHostname() []string {
	if o == nil || IsNil(o.Hostname) {
		var ret []string
		return ret
	}
	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetHostnameOk() ([]string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GatewaySearch) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given []string and assigns it to the Hostname field.
func (o *GatewaySearch) SetHostname(v []string) {
	o.Hostname = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GatewaySearch) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GatewaySearch) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GatewaySearch) SetIp(v string) {
	o.Ip = &v
}

// GetLastHostname returns the LastHostname field value if set, zero value otherwise.
func (o *GatewaySearch) GetLastHostname() string {
	if o == nil || IsNil(o.LastHostname) {
		var ret string
		return ret
	}
	return *o.LastHostname
}

// GetLastHostnameOk returns a tuple with the LastHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetLastHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.LastHostname) {
		return nil, false
	}
	return o.LastHostname, true
}

// HasLastHostname returns a boolean if a field has been set.
func (o *GatewaySearch) HasLastHostname() bool {
	if o != nil && !IsNil(o.LastHostname) {
		return true
	}

	return false
}

// SetLastHostname gets a reference to the given string and assigns it to the LastHostname field.
func (o *GatewaySearch) SetLastHostname(v string) {
	o.LastHostname = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GatewaySearch) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GatewaySearch) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GatewaySearch) SetMac(v string) {
	o.Mac = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GatewaySearch) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GatewaySearch) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GatewaySearch) SetModel(v string) {
	o.Model = &v
}

// GetNumMembers returns the NumMembers field value if set, zero value otherwise.
func (o *GatewaySearch) GetNumMembers() int32 {
	if o == nil || IsNil(o.NumMembers) {
		var ret int32
		return ret
	}
	return *o.NumMembers
}

// GetNumMembersOk returns a tuple with the NumMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetNumMembersOk() (*int32, bool) {
	if o == nil || IsNil(o.NumMembers) {
		return nil, false
	}
	return o.NumMembers, true
}

// HasNumMembers returns a boolean if a field has been set.
func (o *GatewaySearch) HasNumMembers() bool {
	if o != nil && !IsNil(o.NumMembers) {
		return true
	}

	return false
}

// SetNumMembers gets a reference to the given int32 and assigns it to the NumMembers field.
func (o *GatewaySearch) SetNumMembers(v int32) {
	o.NumMembers = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *GatewaySearch) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *GatewaySearch) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *GatewaySearch) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *GatewaySearch) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *GatewaySearch) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *GatewaySearch) SetSiteId(v string) {
	o.SiteId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GatewaySearch) GetTimestamp() float32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret float32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GatewaySearch) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given float32 and assigns it to the Timestamp field.
func (o *GatewaySearch) SetTimestamp(v float32) {
	o.Timestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GatewaySearch) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GatewaySearch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GatewaySearch) SetType(v string) {
	o.Type = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *GatewaySearch) GetUptime() int32 {
	if o == nil || IsNil(o.Uptime) {
		var ret int32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetUptimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Uptime) {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *GatewaySearch) HasUptime() bool {
	if o != nil && !IsNil(o.Uptime) {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int32 and assigns it to the Uptime field.
func (o *GatewaySearch) SetUptime(v int32) {
	o.Uptime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GatewaySearch) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewaySearch) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GatewaySearch) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GatewaySearch) SetVersion(v string) {
	o.Version = &v
}

func (o GatewaySearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewaySearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtIp) {
		toSerialize["ext_ip"] = o.ExtIp
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.LastHostname) {
		toSerialize["last_hostname"] = o.LastHostname
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.NumMembers) {
		toSerialize["num_members"] = o.NumMembers
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uptime) {
		toSerialize["uptime"] = o.Uptime
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewaySearch) UnmarshalJSON(data []byte) (err error) {
	varGatewaySearch := _GatewaySearch{}

	err = json.Unmarshal(data, &varGatewaySearch)

	if err != nil {
		return err
	}

	*o = GatewaySearch(varGatewaySearch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ext_ip")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "last_hostname")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "model")
		delete(additionalProperties, "num_members")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uptime")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewaySearch struct {
	value *GatewaySearch
	isSet bool
}

func (v NullableGatewaySearch) Get() *GatewaySearch {
	return v.value
}

func (v *NullableGatewaySearch) Set(val *GatewaySearch) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewaySearch) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewaySearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewaySearch(val *GatewaySearch) *NullableGatewaySearch {
	return &NullableGatewaySearch{value: val, isSet: true}
}

func (v NullableGatewaySearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewaySearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


