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

// checks if the GatewayTemplateTunnelNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayTemplateTunnelNode{}

// GatewayTemplateTunnelNode struct for GatewayTemplateTunnelNode
type GatewayTemplateTunnelNode struct {
	Hosts []string `json:"hosts,omitempty"`
	// Only if: * `provider`== `zscaler-gre`  * `provider`== `custom-gre`
	InternalIps []string `json:"internal_ips,omitempty"`
	ProbeIps []string `json:"probe_ips,omitempty"`
	// Only if: * `provider`== `custom-ipsec`
	RemoteIds []string `json:"remote_ids,omitempty"`
	WanNames []string `json:"wan_names,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayTemplateTunnelNode GatewayTemplateTunnelNode

// NewGatewayTemplateTunnelNode instantiates a new GatewayTemplateTunnelNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayTemplateTunnelNode() *GatewayTemplateTunnelNode {
	this := GatewayTemplateTunnelNode{}
	return &this
}

// NewGatewayTemplateTunnelNodeWithDefaults instantiates a new GatewayTemplateTunnelNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayTemplateTunnelNodeWithDefaults() *GatewayTemplateTunnelNode {
	this := GatewayTemplateTunnelNode{}
	return &this
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelNode) GetHosts() []string {
	if o == nil || IsNil(o.Hosts) {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelNode) GetHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelNode) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *GatewayTemplateTunnelNode) SetHosts(v []string) {
	o.Hosts = v
}

// GetInternalIps returns the InternalIps field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelNode) GetInternalIps() []string {
	if o == nil || IsNil(o.InternalIps) {
		var ret []string
		return ret
	}
	return o.InternalIps
}

// GetInternalIpsOk returns a tuple with the InternalIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelNode) GetInternalIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.InternalIps) {
		return nil, false
	}
	return o.InternalIps, true
}

// HasInternalIps returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelNode) HasInternalIps() bool {
	if o != nil && !IsNil(o.InternalIps) {
		return true
	}

	return false
}

// SetInternalIps gets a reference to the given []string and assigns it to the InternalIps field.
func (o *GatewayTemplateTunnelNode) SetInternalIps(v []string) {
	o.InternalIps = v
}

// GetProbeIps returns the ProbeIps field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelNode) GetProbeIps() []string {
	if o == nil || IsNil(o.ProbeIps) {
		var ret []string
		return ret
	}
	return o.ProbeIps
}

// GetProbeIpsOk returns a tuple with the ProbeIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelNode) GetProbeIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProbeIps) {
		return nil, false
	}
	return o.ProbeIps, true
}

// HasProbeIps returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelNode) HasProbeIps() bool {
	if o != nil && !IsNil(o.ProbeIps) {
		return true
	}

	return false
}

// SetProbeIps gets a reference to the given []string and assigns it to the ProbeIps field.
func (o *GatewayTemplateTunnelNode) SetProbeIps(v []string) {
	o.ProbeIps = v
}

// GetRemoteIds returns the RemoteIds field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelNode) GetRemoteIds() []string {
	if o == nil || IsNil(o.RemoteIds) {
		var ret []string
		return ret
	}
	return o.RemoteIds
}

// GetRemoteIdsOk returns a tuple with the RemoteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelNode) GetRemoteIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoteIds) {
		return nil, false
	}
	return o.RemoteIds, true
}

// HasRemoteIds returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelNode) HasRemoteIds() bool {
	if o != nil && !IsNil(o.RemoteIds) {
		return true
	}

	return false
}

// SetRemoteIds gets a reference to the given []string and assigns it to the RemoteIds field.
func (o *GatewayTemplateTunnelNode) SetRemoteIds(v []string) {
	o.RemoteIds = v
}

// GetWanNames returns the WanNames field value if set, zero value otherwise.
func (o *GatewayTemplateTunnelNode) GetWanNames() []string {
	if o == nil || IsNil(o.WanNames) {
		var ret []string
		return ret
	}
	return o.WanNames
}

// GetWanNamesOk returns a tuple with the WanNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTemplateTunnelNode) GetWanNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.WanNames) {
		return nil, false
	}
	return o.WanNames, true
}

// HasWanNames returns a boolean if a field has been set.
func (o *GatewayTemplateTunnelNode) HasWanNames() bool {
	if o != nil && !IsNil(o.WanNames) {
		return true
	}

	return false
}

// SetWanNames gets a reference to the given []string and assigns it to the WanNames field.
func (o *GatewayTemplateTunnelNode) SetWanNames(v []string) {
	o.WanNames = v
}

func (o GatewayTemplateTunnelNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayTemplateTunnelNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.InternalIps) {
		toSerialize["internal_ips"] = o.InternalIps
	}
	if !IsNil(o.ProbeIps) {
		toSerialize["probe_ips"] = o.ProbeIps
	}
	if !IsNil(o.RemoteIds) {
		toSerialize["remote_ids"] = o.RemoteIds
	}
	if !IsNil(o.WanNames) {
		toSerialize["wan_names"] = o.WanNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayTemplateTunnelNode) UnmarshalJSON(data []byte) (err error) {
	varGatewayTemplateTunnelNode := _GatewayTemplateTunnelNode{}

	err = json.Unmarshal(data, &varGatewayTemplateTunnelNode)

	if err != nil {
		return err
	}

	*o = GatewayTemplateTunnelNode(varGatewayTemplateTunnelNode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "internal_ips")
		delete(additionalProperties, "probe_ips")
		delete(additionalProperties, "remote_ids")
		delete(additionalProperties, "wan_names")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayTemplateTunnelNode struct {
	value *GatewayTemplateTunnelNode
	isSet bool
}

func (v NullableGatewayTemplateTunnelNode) Get() *GatewayTemplateTunnelNode {
	return v.value
}

func (v *NullableGatewayTemplateTunnelNode) Set(val *GatewayTemplateTunnelNode) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTemplateTunnelNode) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTemplateTunnelNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTemplateTunnelNode(val *GatewayTemplateTunnelNode) *NullableGatewayTemplateTunnelNode {
	return &NullableGatewayTemplateTunnelNode{value: val, isSet: true}
}

func (v NullableGatewayTemplateTunnelNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTemplateTunnelNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


