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

// checks if the GatewayTemplateTunnelIpsecProposal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayTemplateTunnelIpsecProposal{}

// GatewayTemplateTunnelIpsecProposal struct for GatewayTemplateTunnelIpsecProposal
type GatewayTemplateTunnelIpsecProposal struct {
	AuthAlgo *TunnelConfigsAuthAlgo `json:"auth_algo,omitempty"`
	DhGroup *TunnelConfigsDhGroup `json:"dh_group,omitempty"`
	EncAlgo NullableTunnelConfigsEncAlgo `json:"enc_algo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayTemplateTunnelIpsecProposal GatewayTemplateTunnelIpsecProposal

// NewGatewayTemplateTunnelIpsecProposal instantiates a new GatewayTemplateTunnelIpsecProposal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayTemplateTunnelIpsecProposal() *GatewayTemplateTunnelIpsecProposal {
	this := GatewayTemplateTunnelIpsecProposal{}
	var dhGroup TunnelConfigsDhGroup = TUNNELCONFIGSDHGROUP__14
	this.DhGroup = &dhGroup
	var encAlgo TunnelConfigsEncAlgo = TUNNELCONFIGSENCALGO_AES256
	this.EncAlgo = *NewNullableTunnelConfigsEncAlgo(&encAlgo)
	return &this
}

// NewGatewayTemplateTunnelIpsecProposalWithDefaults instantiates a new GatewayTemplateTunnelIpsecProposal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayTemplateTunnelIpsecProposalWithDefaults() *GatewayTemplateTunnelIpsecProposal {
	this := GatewayTemplateTunnelIpsecProposal{}
	var dhGroup TunnelConfigsDhGroup = TUNNELCONFIGSDHGROUP__14
	this.DhGroup = &dhGroup
	var encAlgo TunnelConfigsEncAlgo = TUNNELCONFIGSENCALGO_AES256
	this.EncAlgo = *NewNullableTunnelConfigsEncAlgo(&encAlgo)
	return &this
}

// GetAuthAlgo returns the AuthAlgo field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelIpsecProposal) GetAuthAlgo() TunnelConfigsAuthAlgo {
	if o == nil || IsNil(o.AuthAlgo) {
		var ret TunnelConfigsAuthAlgo
		return ret
	}
	return *o.AuthAlgo
}

// GetAuthAlgoOk returns a tuple with the AuthAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelIpsecProposal) GetAuthAlgoOk() (*TunnelConfigsAuthAlgo, bool) {
	if o == nil || IsNil(o.AuthAlgo) {
		return nil, false
	}
	return o.AuthAlgo, true
}

// HasAuthAlgo returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelIpsecProposal) HasAuthAlgo() bool {
	if o != nil && !IsNil(o.AuthAlgo) {
		return true
	}

	return false
}

// SetAuthAlgo gets a reference to the given TunnelConfigsAuthAlgo and assigns it to the AuthAlgo field.
func (o *GatewayTemplateTunnelIpsecProposal) SetAuthAlgo(v TunnelConfigsAuthAlgo) {
	o.AuthAlgo = &v
}

// GetDhGroup returns the DhGroup field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelIpsecProposal) GetDhGroup() TunnelConfigsDhGroup {
	if o == nil || IsNil(o.DhGroup) {
		var ret TunnelConfigsDhGroup
		return ret
	}
	return *o.DhGroup
}

// GetDhGroupOk returns a tuple with the DhGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelIpsecProposal) GetDhGroupOk() (*TunnelConfigsDhGroup, bool) {
	if o == nil || IsNil(o.DhGroup) {
		return nil, false
	}
	return o.DhGroup, true
}

// HasDhGroup returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelIpsecProposal) HasDhGroup() bool {
	if o != nil && !IsNil(o.DhGroup) {
		return true
	}

	return false
}

// SetDhGroup gets a reference to the given TunnelConfigsDhGroup and assigns it to the DhGroup field.
func (o *GatewayTemplateTunnelIpsecProposal) SetDhGroup(v TunnelConfigsDhGroup) {
	o.DhGroup = &v
}

// GetEncAlgo returns the EncAlgo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GatewayTemplateTunnelIpsecProposal) GetEncAlgo() TunnelConfigsEncAlgo {
	if o == nil || IsNil(o.EncAlgo.Get()) {
		var ret TunnelConfigsEncAlgo
		return ret
	}
	return *o.EncAlgo.Get()
}

// GetEncAlgoOk returns a tuple with the EncAlgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GatewayTemplateTunnelIpsecProposal) GetEncAlgoOk() (*TunnelConfigsEncAlgo, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncAlgo.Get(), o.EncAlgo.IsSet()
}

// HasEncAlgo returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelIpsecProposal) HasEncAlgo() bool {
	if o != nil && o.EncAlgo.IsSet() {
		return true
	}

	return false
}

// SetEncAlgo gets a reference to the given NullableTunnelConfigsEncAlgo and assigns it to the EncAlgo field.
func (o *GatewayTemplateTunnelIpsecProposal) SetEncAlgo(v TunnelConfigsEncAlgo) {
	o.EncAlgo.Set(&v)
}
// SetEncAlgoNil sets the value for EncAlgo to be an explicit nil
func (o *GatewayTemplateTunnelIpsecProposal) SetEncAlgoNil() {
	o.EncAlgo.Set(nil)
}

// UnsetEncAlgo ensures that no value is present for EncAlgo, not even an explicit nil
func (o *GatewayTemplateTunnelIpsecProposal) UnsetEncAlgo() {
	o.EncAlgo.Unset()
}

func (o GatewayTemplateTunnelIpsecProposal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayTemplateTunnelIpsecProposal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthAlgo) {
		toSerialize["auth_algo"] = o.AuthAlgo
	}
	if !IsNil(o.DhGroup) {
		toSerialize["dh_group"] = o.DhGroup
	}
	if o.EncAlgo.IsSet() {
		toSerialize["enc_algo"] = o.EncAlgo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayTemplateTunnelIpsecProposal) UnmarshalJSON(data []byte) (err error) {
	varGatewayTemplateTunnelIpsecProposal := _GatewayTemplateTunnelIpsecProposal{}

	err = json.Unmarshal(data, &varGatewayTemplateTunnelIpsecProposal)

	if err != nil {
		return err
	}

	*o = GatewayTemplateTunnelIpsecProposal(varGatewayTemplateTunnelIpsecProposal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auth_algo")
		delete(additionalProperties, "dh_group")
		delete(additionalProperties, "enc_algo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayTemplateTunnelIpsecProposal struct {
	value *GatewayTemplateTunnelIpsecProposal
	isSet bool
}

func (v NullableGatewayTemplateTunnelIpsecProposal) Get() *GatewayTemplateTunnelIpsecProposal {
	return v.value
}

func (v *NullableGatewayTemplateTunnelIpsecProposal) Set(val *GatewayTemplateTunnelIpsecProposal) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateTunnelIpsecProposal) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateTunnelIpsecProposal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateTunnelIpsecProposal(val *GatewayTemplateTunnelIpsecProposal) *NullableGatewayTemplateTunnelIpsecProposal {
	return &NullableGatewayTemplateTunnelIpsecProposal{value: val, isSet: true}
}

func (v NullableGatewayTemplateTunnelIpsecProposal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateTunnelIpsecProposal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


