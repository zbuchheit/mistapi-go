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

// checks if the AclTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AclTag{}

// AclTag struct for AclTag
type AclTag struct {
	// required if - `type`==`dynamic_gbp` (gbp_tag received from RADIUS) - `type`==`static_gbp` (applying gbp tag against matching conditions)
	GbpTag *int32 `json:"gbp_tag,omitempty"`
	// required if  - `type`==`mac` - `type`==`static_gbp` if from matching mac
	Macs []string `json:"macs,omitempty"`
	// if: - `type`==`mac` (optional. default is `any`) - `type`==`subnet` (optional. default is `any`) - `type`==`network` - `type`==`resource` (optional. default is `any`) - `type`==`static_gbp` if from matching network (vlan)
	Network *string `json:"network,omitempty"`
	// required if  - `type`==`radius_group`  - `type`==`static_gbp` if from matching radius_group
	RadiusGroup *string `json:"radius_group,omitempty"`
	// if `type`==`resource` empty means unrestricted, i.e. any
	Specs []AclTagSpec `json:"specs,omitempty"`
	// if  - `type`==`subnet`  - `type`==`resource` (optional. default is `any`) - `type`==`static_gbp` if from matching subnet
	Subnets []string `json:"subnets,omitempty"`
	Type AclTagType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AclTag AclTag

// NewAclTag instantiates a new AclTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclTag(type_ AclTagType) *AclTag {
	this := AclTag{}
	this.Type = type_
	return &this
}

// NewAclTagWithDefaults instantiates a new AclTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclTagWithDefaults() *AclTag {
	this := AclTag{}
	return &this
}

// GetGbpTag returns the GbpTag field value if set, zero value otherwise.
func (o *AclTag) GetGbpTag() int32 {
	if o == nil || IsNil(o.GbpTag) {
		var ret int32
		return ret
	}
	return *o.GbpTag
}

// GetGbpTagOk returns a tuple with the GbpTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTag) GetGbpTagOk() (*int32, bool) {
	if o == nil || IsNil(o.GbpTag) {
		return nil, false
	}
	return o.GbpTag, true
}

// HasGbpTag returns a boolean if a field has been set.
func (o *AclTag) HasGbpTag() bool {
	if o != nil && !IsNil(o.GbpTag) {
		return true
	}

	return false
}

// SetGbpTag gets a reference to the given int32 and assigns it to the GbpTag field.
func (o *AclTag) SetGbpTag(v int32) {
	o.GbpTag = &v
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *AclTag) GetMacs() []string {
	if o == nil || IsNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTag) GetMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *AclTag) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *AclTag) SetMacs(v []string) {
	o.Macs = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AclTag) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTag) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AclTag) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *AclTag) SetNetwork(v string) {
	o.Network = &v
}

// GetRadiusGroup returns the RadiusGroup field value if set, zero value otherwise.
func (o *AclTag) GetRadiusGroup() string {
	if o == nil || IsNil(o.RadiusGroup) {
		var ret string
		return ret
	}
	return *o.RadiusGroup
}

// GetRadiusGroupOk returns a tuple with the RadiusGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTag) GetRadiusGroupOk() (*string, bool) {
	if o == nil || IsNil(o.RadiusGroup) {
		return nil, false
	}
	return o.RadiusGroup, true
}

// HasRadiusGroup returns a boolean if a field has been set.
func (o *AclTag) HasRadiusGroup() bool {
	if o != nil && !IsNil(o.RadiusGroup) {
		return true
	}

	return false
}

// SetRadiusGroup gets a reference to the given string and assigns it to the RadiusGroup field.
func (o *AclTag) SetRadiusGroup(v string) {
	o.RadiusGroup = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *AclTag) GetSpecs() []AclTagSpec {
	if o == nil || IsNil(o.Specs) {
		var ret []AclTagSpec
		return ret
	}
	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTag) GetSpecsOk() ([]AclTagSpec, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *AclTag) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given []AclTagSpec and assigns it to the Specs field.
func (o *AclTag) SetSpecs(v []AclTagSpec) {
	o.Specs = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *AclTag) GetSubnets() []string {
	if o == nil || IsNil(o.Subnets) {
		var ret []string
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTag) GetSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *AclTag) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *AclTag) SetSubnets(v []string) {
	o.Subnets = v
}

// GetType returns the Type field value
func (o *AclTag) GetType() AclTagType {
	if o == nil {
		var ret AclTagType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AclTag) GetTypeOk() (*AclTagType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AclTag) SetType(v AclTagType) {
	o.Type = v
}

func (o AclTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AclTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GbpTag) {
		toSerialize["gbp_tag"] = o.GbpTag
	}
	if !IsNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.RadiusGroup) {
		toSerialize["radius_group"] = o.RadiusGroup
	}
	if !IsNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}
	if !IsNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AclTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varAclTag := _AclTag{}

	err = json.Unmarshal(data, &varAclTag)

	if err != nil {
		return err
	}

	*o = AclTag(varAclTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gbp_tag")
		delete(additionalProperties, "macs")
		delete(additionalProperties, "network")
		delete(additionalProperties, "radius_group")
		delete(additionalProperties, "specs")
		delete(additionalProperties, "subnets")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAclTag struct {
	value *AclTag
	isSet bool
}

func (v NullableAclTag) Get() *AclTag {
	return v.value
}

func (v *NullableAclTag) Set(val *AclTag) {
	v.value = val
	v.isSet = true
}

func (v NullableAclTag) IsSet() bool {
	return v.isSet
}

func (v *NullableAclTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclTag(val *AclTag) *NullableAclTag {
	return &NullableAclTag{value: val, isSet: true}
}

func (v NullableAclTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


