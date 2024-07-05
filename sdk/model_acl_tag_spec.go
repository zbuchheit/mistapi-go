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

// checks if the AclTagSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AclTagSpec{}

// AclTagSpec struct for AclTagSpec
type AclTagSpec struct {
	// matched dst port, \"0\" means any
	PortRange *string `json:"port_range,omitempty"`
	// `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`. `protocol_number` is between 1-254
	Protocol *string `json:"protocol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AclTagSpec AclTagSpec

// NewAclTagSpec instantiates a new AclTagSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclTagSpec() *AclTagSpec {
	this := AclTagSpec{}
	var portRange string = "80"
	this.PortRange = &portRange
	var protocol string = "any"
	this.Protocol = &protocol
	return &this
}

// NewAclTagSpecWithDefaults instantiates a new AclTagSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclTagSpecWithDefaults() *AclTagSpec {
	this := AclTagSpec{}
	var portRange string = "80"
	this.PortRange = &portRange
	var protocol string = "any"
	this.Protocol = &protocol
	return &this
}

// GetPortRange returns the PortRange field value if set, zero value otherwise.
func (o *AclTagSpec) GetPortRange() string {
	if o == nil || IsNil(o.PortRange) {
		var ret string
		return ret
	}
	return *o.PortRange
}

// GetPortRangeOk returns a tuple with the PortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTagSpec) GetPortRangeOk() (*string, bool) {
	if o == nil || IsNil(o.PortRange) {
		return nil, false
	}
	return o.PortRange, true
}

// HasPortRange returns a boolean if a field has been set.
func (o *AclTagSpec) HasPortRange() bool {
	if o != nil && !IsNil(o.PortRange) {
		return true
	}

	return false
}

// SetPortRange gets a reference to the given string and assigns it to the PortRange field.
func (o *AclTagSpec) SetPortRange(v string) {
	o.PortRange = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AclTagSpec) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclTagSpec) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AclTagSpec) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AclTagSpec) SetProtocol(v string) {
	o.Protocol = &v
}

func (o AclTagSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AclTagSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortRange) {
		toSerialize["port_range"] = o.PortRange
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AclTagSpec) UnmarshalJSON(data []byte) (err error) {
	varAclTagSpec := _AclTagSpec{}

	err = json.Unmarshal(data, &varAclTagSpec)

	if err != nil {
		return err
	}

	*o = AclTagSpec(varAclTagSpec)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "port_range")
		delete(additionalProperties, "protocol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAclTagSpec struct {
	value *AclTagSpec
	isSet bool
}

func (v NullableAclTagSpec) Get() *AclTagSpec {
	return v.value
}

func (v *NullableAclTagSpec) Set(val *AclTagSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAclTagSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAclTagSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclTagSpec(val *AclTagSpec) *NullableAclTagSpec {
	return &NullableAclTagSpec{value: val, isSet: true}
}

func (v NullableAclTagSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclTagSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


