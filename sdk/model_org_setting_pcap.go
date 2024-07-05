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
)

// checks if the OrgSettingPcap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingPcap{}

// OrgSettingPcap struct for OrgSettingPcap
type OrgSettingPcap struct {
	Bucket *string `json:"bucket,omitempty"`
	// max_len of non-management packets to capture
	MaxPktLen *int32 `json:"max_pkt_len,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingPcap OrgSettingPcap

// NewOrgSettingPcap instantiates a new OrgSettingPcap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingPcap() *OrgSettingPcap {
	this := OrgSettingPcap{}
	var maxPktLen int32 = 128
	this.MaxPktLen = &maxPktLen
	return &this
}

// NewOrgSettingPcapWithDefaults instantiates a new OrgSettingPcap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingPcapWithDefaults() *OrgSettingPcap {
	this := OrgSettingPcap{}
	var maxPktLen int32 = 128
	this.MaxPktLen = &maxPktLen
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *OrgSettingPcap) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingPcap) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *OrgSettingPcap) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *OrgSettingPcap) SetBucket(v string) {
	o.Bucket = &v
}

// GetMaxPktLen returns the MaxPktLen field value if set, zero value otherwise.
func (o *OrgSettingPcap) GetMaxPktLen() int32 {
	if o == nil || IsNil(o.MaxPktLen) {
		var ret int32
		return ret
	}
	return *o.MaxPktLen
}

// GetMaxPktLenOk returns a tuple with the MaxPktLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingPcap) GetMaxPktLenOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPktLen) {
		return nil, false
	}
	return o.MaxPktLen, true
}

// HasMaxPktLen returns a boolean if a field has been set.
func (o *OrgSettingPcap) HasMaxPktLen() bool {
	if o != nil && !IsNil(o.MaxPktLen) {
		return true
	}

	return false
}

// SetMaxPktLen gets a reference to the given int32 and assigns it to the MaxPktLen field.
func (o *OrgSettingPcap) SetMaxPktLen(v int32) {
	o.MaxPktLen = &v
}

func (o OrgSettingPcap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingPcap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.MaxPktLen) {
		toSerialize["max_pkt_len"] = o.MaxPktLen
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingPcap) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingPcap := _OrgSettingPcap{}

	err = json.Unmarshal(data, &varOrgSettingPcap)

	if err != nil {
		return err
	}

	*o = OrgSettingPcap(varOrgSettingPcap)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bucket")
		delete(additionalProperties, "max_pkt_len")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingPcap struct {
	value *OrgSettingPcap
	isSet bool
}

func (v NullableOrgSettingPcap) Get() *OrgSettingPcap {
	return v.value
}

func (v *NullableOrgSettingPcap) Set(val *OrgSettingPcap) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingPcap) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingPcap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingPcap(val *OrgSettingPcap) *NullableOrgSettingPcap {
	return &NullableOrgSettingPcap{value: val, isSet: true}
}

func (v NullableOrgSettingPcap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingPcap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


