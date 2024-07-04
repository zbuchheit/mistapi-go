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

// checks if the AccountJamfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountJamfInfo{}

// AccountJamfInfo OAuth linked Jamf apps account details
type AccountJamfInfo struct {
	// customer account client id
	ClientId *string `json:"client_id,omitempty"`
	// This error is provided when the Jamf account fails to fetch token/data
	Error *string `json:"error,omitempty"`
	// customer account Jamf instance URL
	InstanceUrl *string `json:"instance_url,omitempty"`
	// Is the last data pull for Jamf account is successful or not
	LastStatus *string `json:"last_status,omitempty"`
	// Last data pull timestamp, background jobs that pull Jamf account data
	LastSync *int32 `json:"last_sync,omitempty"`
	// First name of the user who linked the Jamf account
	LinkedBy *string `json:"linked_by,omitempty"`
	// Name of the company whose Jamf account mist has subscribed to
	Name *string `json:"name,omitempty"`
	// smart group membership for determining compliance status
	SmartgroupName *string `json:"smartgroup_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountJamfInfo AccountJamfInfo

// NewAccountJamfInfo instantiates a new AccountJamfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountJamfInfo() *AccountJamfInfo {
	this := AccountJamfInfo{}
	return &this
}

// NewAccountJamfInfoWithDefaults instantiates a new AccountJamfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountJamfInfoWithDefaults() *AccountJamfInfo {
	this := AccountJamfInfo{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AccountJamfInfo) SetClientId(v string) {
	o.ClientId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *AccountJamfInfo) SetError(v string) {
	o.Error = &v
}

// GetInstanceUrl returns the InstanceUrl field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetInstanceUrl() string {
	if o == nil || IsNil(o.InstanceUrl) {
		var ret string
		return ret
	}
	return *o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetInstanceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceUrl) {
		return nil, false
	}
	return o.InstanceUrl, true
}

// HasInstanceUrl returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasInstanceUrl() bool {
	if o != nil && !IsNil(o.InstanceUrl) {
		return true
	}

	return false
}

// SetInstanceUrl gets a reference to the given string and assigns it to the InstanceUrl field.
func (o *AccountJamfInfo) SetInstanceUrl(v string) {
	o.InstanceUrl = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetLastStatus() string {
	if o == nil || IsNil(o.LastStatus) {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetLastStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastStatus) {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasLastStatus() bool {
	if o != nil && !IsNil(o.LastStatus) {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *AccountJamfInfo) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetLastSync() int32 {
	if o == nil || IsNil(o.LastSync) {
		var ret int32
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetLastSyncOk() (*int32, bool) {
	if o == nil || IsNil(o.LastSync) {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasLastSync() bool {
	if o != nil && !IsNil(o.LastSync) {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given int32 and assigns it to the LastSync field.
func (o *AccountJamfInfo) SetLastSync(v int32) {
	o.LastSync = &v
}

// GetLinkedBy returns the LinkedBy field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetLinkedBy() string {
	if o == nil || IsNil(o.LinkedBy) {
		var ret string
		return ret
	}
	return *o.LinkedBy
}

// GetLinkedByOk returns a tuple with the LinkedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetLinkedByOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedBy) {
		return nil, false
	}
	return o.LinkedBy, true
}

// HasLinkedBy returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasLinkedBy() bool {
	if o != nil && !IsNil(o.LinkedBy) {
		return true
	}

	return false
}

// SetLinkedBy gets a reference to the given string and assigns it to the LinkedBy field.
func (o *AccountJamfInfo) SetLinkedBy(v string) {
	o.LinkedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountJamfInfo) SetName(v string) {
	o.Name = &v
}

// GetSmartgroupName returns the SmartgroupName field value if set, zero value otherwise.
func (o *AccountJamfInfo) GetSmartgroupName() string {
	if o == nil || IsNil(o.SmartgroupName) {
		var ret string
		return ret
	}
	return *o.SmartgroupName
}

// GetSmartgroupNameOk returns a tuple with the SmartgroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJamfInfo) GetSmartgroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.SmartgroupName) {
		return nil, false
	}
	return o.SmartgroupName, true
}

// HasSmartgroupName returns a boolean if a field has been set.
func (o *AccountJamfInfo) HasSmartgroupName() bool {
	if o != nil && !IsNil(o.SmartgroupName) {
		return true
	}

	return false
}

// SetSmartgroupName gets a reference to the given string and assigns it to the SmartgroupName field.
func (o *AccountJamfInfo) SetSmartgroupName(v string) {
	o.SmartgroupName = &v
}

func (o AccountJamfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountJamfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.InstanceUrl) {
		toSerialize["instance_url"] = o.InstanceUrl
	}
	if !IsNil(o.LastStatus) {
		toSerialize["last_status"] = o.LastStatus
	}
	if !IsNil(o.LastSync) {
		toSerialize["last_sync"] = o.LastSync
	}
	if !IsNil(o.LinkedBy) {
		toSerialize["linked_by"] = o.LinkedBy
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SmartgroupName) {
		toSerialize["smartgroup_name"] = o.SmartgroupName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountJamfInfo) UnmarshalJSON(data []byte) (err error) {
	varAccountJamfInfo := _AccountJamfInfo{}

	err = json.Unmarshal(data, &varAccountJamfInfo)

	if err != nil {
		return err
	}

	*o = AccountJamfInfo(varAccountJamfInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "error")
		delete(additionalProperties, "instance_url")
		delete(additionalProperties, "last_status")
		delete(additionalProperties, "last_sync")
		delete(additionalProperties, "linked_by")
		delete(additionalProperties, "name")
		delete(additionalProperties, "smartgroup_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountJamfInfo struct {
	value *AccountJamfInfo
	isSet bool
}

func (v NullableAccountJamfInfo) Get() *AccountJamfInfo {
	return v.value
}

func (v *NullableAccountJamfInfo) Set(val *AccountJamfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountJamfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountJamfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountJamfInfo(val *AccountJamfInfo) *NullableAccountJamfInfo {
	return &NullableAccountJamfInfo{value: val, isSet: true}
}

func (v NullableAccountJamfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountJamfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


