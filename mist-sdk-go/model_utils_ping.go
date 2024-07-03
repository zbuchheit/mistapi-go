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
	"fmt"
)

// checks if the UtilsPing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsPing{}

// UtilsPing struct for UtilsPing
type UtilsPing struct {
	Count *int32 `json:"count,omitempty"`
	// Interface through which packet needs to egress
	EgressInterface *string `json:"egress_interface,omitempty"`
	Host string `json:"host"`
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	Size *int32 `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsPing UtilsPing

// NewUtilsPing instantiates a new UtilsPing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsPing(host string) *UtilsPing {
	this := UtilsPing{}
	var count int32 = 10
	this.Count = &count
	this.Host = host
	var size int32 = 56
	this.Size = &size
	return &this
}

// NewUtilsPingWithDefaults instantiates a new UtilsPing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsPingWithDefaults() *UtilsPing {
	this := UtilsPing{}
	var count int32 = 10
	this.Count = &count
	var size int32 = 56
	this.Size = &size
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UtilsPing) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsPing) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UtilsPing) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UtilsPing) SetCount(v int32) {
	o.Count = &v
}

// GetEgressInterface returns the EgressInterface field value if set, zero value otherwise.
func (o *UtilsPing) GetEgressInterface() string {
	if o == nil || IsNil(o.EgressInterface) {
		var ret string
		return ret
	}
	return *o.EgressInterface
}

// GetEgressInterfaceOk returns a tuple with the EgressInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsPing) GetEgressInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.EgressInterface) {
		return nil, false
	}
	return o.EgressInterface, true
}

// HasEgressInterface returns a boolean if a field has been set.
func (o *UtilsPing) HasEgressInterface() bool {
	if o != nil && !IsNil(o.EgressInterface) {
		return true
	}

	return false
}

// SetEgressInterface gets a reference to the given string and assigns it to the EgressInterface field.
func (o *UtilsPing) SetEgressInterface(v string) {
	o.EgressInterface = &v
}

// GetHost returns the Host field value
func (o *UtilsPing) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UtilsPing) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UtilsPing) SetHost(v string) {
	o.Host = v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *UtilsPing) GetNode() HaClusterNodeEnum {
	if o == nil || IsNil(o.Node) {
		var ret HaClusterNodeEnum
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsPing) GetNodeOk() (*HaClusterNodeEnum, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *UtilsPing) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given HaClusterNodeEnum and assigns it to the Node field.
func (o *UtilsPing) SetNode(v HaClusterNodeEnum) {
	o.Node = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UtilsPing) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsPing) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UtilsPing) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *UtilsPing) SetSize(v int32) {
	o.Size = &v
}

func (o UtilsPing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsPing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.EgressInterface) {
		toSerialize["egress_interface"] = o.EgressInterface
	}
	toSerialize["host"] = o.Host
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsPing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
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

	varUtilsPing := _UtilsPing{}

	err = json.Unmarshal(data, &varUtilsPing)

	if err != nil {
		return err
	}

	*o = UtilsPing(varUtilsPing)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "egress_interface")
		delete(additionalProperties, "host")
		delete(additionalProperties, "node")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsPing struct {
	value *UtilsPing
	isSet bool
}

func (v NullableUtilsPing) Get() *UtilsPing {
	return v.value
}

func (v *NullableUtilsPing) Set(val *UtilsPing) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsPing) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsPing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsPing(val *UtilsPing) *NullableUtilsPing {
	return &NullableUtilsPing{value: val, isSet: true}
}

func (v NullableUtilsPing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsPing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


