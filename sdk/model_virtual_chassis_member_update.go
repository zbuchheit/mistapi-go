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

// checks if the VirtualChassisMemberUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualChassisMemberUpdate{}

// VirtualChassisMemberUpdate struct for VirtualChassisMemberUpdate
type VirtualChassisMemberUpdate struct {
	// Required if `op`==`add` or `op`==`preprovision`.
	Mac *string `json:"mac,omitempty"`
	// Required if `op`==`remove` or `op`==`preprovision`. Optional if `op`==`add`
	Member *int32 `json:"member,omitempty"`
	// Required if `op`==`add` or `op`==`preprovision`
	VcPorts []string `json:"vc_ports,omitempty"`
	VcRole *VirtualChassisMemberUpdateVcRole `json:"vc_role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualChassisMemberUpdate VirtualChassisMemberUpdate

// NewVirtualChassisMemberUpdate instantiates a new VirtualChassisMemberUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualChassisMemberUpdate() *VirtualChassisMemberUpdate {
	this := VirtualChassisMemberUpdate{}
	return &this
}

// NewVirtualChassisMemberUpdateWithDefaults instantiates a new VirtualChassisMemberUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualChassisMemberUpdateWithDefaults() *VirtualChassisMemberUpdate {
	this := VirtualChassisMemberUpdate{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *VirtualChassisMemberUpdate) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisMemberUpdate) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *VirtualChassisMemberUpdate) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *VirtualChassisMemberUpdate) SetMac(v string) {
	o.Mac = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *VirtualChassisMemberUpdate) GetMember() int32 {
	if o == nil || IsNil(o.Member) {
		var ret int32
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisMemberUpdate) GetMemberOk() (*int32, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *VirtualChassisMemberUpdate) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given int32 and assigns it to the Member field.
func (o *VirtualChassisMemberUpdate) SetMember(v int32) {
	o.Member = &v
}

// GetVcPorts returns the VcPorts field value if set, zero value otherwise.
func (o *VirtualChassisMemberUpdate) GetVcPorts() []string {
	if o == nil || IsNil(o.VcPorts) {
		var ret []string
		return ret
	}
	return o.VcPorts
}

// GetVcPortsOk returns a tuple with the VcPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisMemberUpdate) GetVcPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.VcPorts) {
		return nil, false
	}
	return o.VcPorts, true
}

// HasVcPorts returns a boolean if a field has been set.
func (o *VirtualChassisMemberUpdate) HasVcPorts() bool {
	if o != nil && !IsNil(o.VcPorts) {
		return true
	}

	return false
}

// SetVcPorts gets a reference to the given []string and assigns it to the VcPorts field.
func (o *VirtualChassisMemberUpdate) SetVcPorts(v []string) {
	o.VcPorts = v
}

// GetVcRole returns the VcRole field value if set, zero value otherwise.
func (o *VirtualChassisMemberUpdate) GetVcRole() VirtualChassisMemberUpdateVcRole {
	if o == nil || IsNil(o.VcRole) {
		var ret VirtualChassisMemberUpdateVcRole
		return ret
	}
	return *o.VcRole
}

// GetVcRoleOk returns a tuple with the VcRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualChassisMemberUpdate) GetVcRoleOk() (*VirtualChassisMemberUpdateVcRole, bool) {
	if o == nil || IsNil(o.VcRole) {
		return nil, false
	}
	return o.VcRole, true
}

// HasVcRole returns a boolean if a field has been set.
func (o *VirtualChassisMemberUpdate) HasVcRole() bool {
	if o != nil && !IsNil(o.VcRole) {
		return true
	}

	return false
}

// SetVcRole gets a reference to the given VirtualChassisMemberUpdateVcRole and assigns it to the VcRole field.
func (o *VirtualChassisMemberUpdate) SetVcRole(v VirtualChassisMemberUpdateVcRole) {
	o.VcRole = &v
}

func (o VirtualChassisMemberUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualChassisMemberUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.VcPorts) {
		toSerialize["vc_ports"] = o.VcPorts
	}
	if !IsNil(o.VcRole) {
		toSerialize["vc_role"] = o.VcRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualChassisMemberUpdate) UnmarshalJSON(data []byte) (err error) {
	varVirtualChassisMemberUpdate := _VirtualChassisMemberUpdate{}

	err = json.Unmarshal(data, &varVirtualChassisMemberUpdate)

	if err != nil {
		return err
	}

	*o = VirtualChassisMemberUpdate(varVirtualChassisMemberUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "member")
		delete(additionalProperties, "vc_ports")
		delete(additionalProperties, "vc_role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualChassisMemberUpdate struct {
	value *VirtualChassisMemberUpdate
	isSet bool
}

func (v NullableVirtualChassisMemberUpdate) Get() *VirtualChassisMemberUpdate {
	return v.value
}

func (v *NullableVirtualChassisMemberUpdate) Set(val *VirtualChassisMemberUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualChassisMemberUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualChassisMemberUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualChassisMemberUpdate(val *VirtualChassisMemberUpdate) *NullableVirtualChassisMemberUpdate {
	return &NullableVirtualChassisMemberUpdate{value: val, isSet: true}
}

func (v NullableVirtualChassisMemberUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualChassisMemberUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


