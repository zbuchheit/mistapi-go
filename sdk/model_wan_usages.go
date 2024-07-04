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

// checks if the WanUsages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WanUsages{}

// WanUsages struct for WanUsages
type WanUsages struct {
	Mac *string `json:"mac,omitempty"`
	PathType *string `json:"path_type,omitempty"`
	PathWeight *int32 `json:"path_weight,omitempty"`
	PeerMac *string `json:"peer_mac,omitempty"`
	PeerPortId *string `json:"peer_port_id,omitempty"`
	Policy *string `json:"policy,omitempty"`
	PortId *string `json:"port_id,omitempty"`
	Tenant *string `json:"tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WanUsages WanUsages

// NewWanUsages instantiates a new WanUsages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWanUsages() *WanUsages {
	this := WanUsages{}
	return &this
}

// NewWanUsagesWithDefaults instantiates a new WanUsages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWanUsagesWithDefaults() *WanUsages {
	this := WanUsages{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *WanUsages) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *WanUsages) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *WanUsages) SetMac(v string) {
	o.Mac = &v
}

// GetPathType returns the PathType field value if set, zero value otherwise.
func (o *WanUsages) GetPathType() string {
	if o == nil || IsNil(o.PathType) {
		var ret string
		return ret
	}
	return *o.PathType
}

// GetPathTypeOk returns a tuple with the PathType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetPathTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PathType) {
		return nil, false
	}
	return o.PathType, true
}

// HasPathType returns a boolean if a field has been set.
func (o *WanUsages) HasPathType() bool {
	if o != nil && !IsNil(o.PathType) {
		return true
	}

	return false
}

// SetPathType gets a reference to the given string and assigns it to the PathType field.
func (o *WanUsages) SetPathType(v string) {
	o.PathType = &v
}

// GetPathWeight returns the PathWeight field value if set, zero value otherwise.
func (o *WanUsages) GetPathWeight() int32 {
	if o == nil || IsNil(o.PathWeight) {
		var ret int32
		return ret
	}
	return *o.PathWeight
}

// GetPathWeightOk returns a tuple with the PathWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetPathWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.PathWeight) {
		return nil, false
	}
	return o.PathWeight, true
}

// HasPathWeight returns a boolean if a field has been set.
func (o *WanUsages) HasPathWeight() bool {
	if o != nil && !IsNil(o.PathWeight) {
		return true
	}

	return false
}

// SetPathWeight gets a reference to the given int32 and assigns it to the PathWeight field.
func (o *WanUsages) SetPathWeight(v int32) {
	o.PathWeight = &v
}

// GetPeerMac returns the PeerMac field value if set, zero value otherwise.
func (o *WanUsages) GetPeerMac() string {
	if o == nil || IsNil(o.PeerMac) {
		var ret string
		return ret
	}
	return *o.PeerMac
}

// GetPeerMacOk returns a tuple with the PeerMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetPeerMacOk() (*string, bool) {
	if o == nil || IsNil(o.PeerMac) {
		return nil, false
	}
	return o.PeerMac, true
}

// HasPeerMac returns a boolean if a field has been set.
func (o *WanUsages) HasPeerMac() bool {
	if o != nil && !IsNil(o.PeerMac) {
		return true
	}

	return false
}

// SetPeerMac gets a reference to the given string and assigns it to the PeerMac field.
func (o *WanUsages) SetPeerMac(v string) {
	o.PeerMac = &v
}

// GetPeerPortId returns the PeerPortId field value if set, zero value otherwise.
func (o *WanUsages) GetPeerPortId() string {
	if o == nil || IsNil(o.PeerPortId) {
		var ret string
		return ret
	}
	return *o.PeerPortId
}

// GetPeerPortIdOk returns a tuple with the PeerPortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetPeerPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PeerPortId) {
		return nil, false
	}
	return o.PeerPortId, true
}

// HasPeerPortId returns a boolean if a field has been set.
func (o *WanUsages) HasPeerPortId() bool {
	if o != nil && !IsNil(o.PeerPortId) {
		return true
	}

	return false
}

// SetPeerPortId gets a reference to the given string and assigns it to the PeerPortId field.
func (o *WanUsages) SetPeerPortId(v string) {
	o.PeerPortId = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *WanUsages) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *WanUsages) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *WanUsages) SetPolicy(v string) {
	o.Policy = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *WanUsages) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *WanUsages) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *WanUsages) SetPortId(v string) {
	o.PortId = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *WanUsages) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WanUsages) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *WanUsages) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *WanUsages) SetTenant(v string) {
	o.Tenant = &v
}

func (o WanUsages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WanUsages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.PathType) {
		toSerialize["path_type"] = o.PathType
	}
	if !IsNil(o.PathWeight) {
		toSerialize["path_weight"] = o.PathWeight
	}
	if !IsNil(o.PeerMac) {
		toSerialize["peer_mac"] = o.PeerMac
	}
	if !IsNil(o.PeerPortId) {
		toSerialize["peer_port_id"] = o.PeerPortId
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WanUsages) UnmarshalJSON(data []byte) (err error) {
	varWanUsages := _WanUsages{}

	err = json.Unmarshal(data, &varWanUsages)

	if err != nil {
		return err
	}

	*o = WanUsages(varWanUsages)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "path_type")
		delete(additionalProperties, "path_weight")
		delete(additionalProperties, "peer_mac")
		delete(additionalProperties, "peer_port_id")
		delete(additionalProperties, "policy")
		delete(additionalProperties, "port_id")
		delete(additionalProperties, "tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWanUsages struct {
	value *WanUsages
	isSet bool
}

func (v NullableWanUsages) Get() *WanUsages {
	return v.value
}

func (v *NullableWanUsages) Set(val *WanUsages) {
	v.value = val
	v.isSet = true
}

func (v NullableWanUsages) IsSet() bool {
	return v.isSet
}

func (v *NullableWanUsages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWanUsages(val *WanUsages) *NullableWanUsages {
	return &NullableWanUsages{value: val, isSet: true}
}

func (v NullableWanUsages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWanUsages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


