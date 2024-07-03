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

// checks if the UtilsShowForwardingTable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtilsShowForwardingTable{}

// UtilsShowForwardingTable struct for UtilsShowForwardingTable
type UtilsShowForwardingTable struct {
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// IP Prefix
	Prefix *string `json:"prefix,omitempty"`
	// only supported with SSR
	ServiceIp *string `json:"service_ip,omitempty"`
	// only supported with SSR
	ServiceName *string `json:"service_name,omitempty"`
	// only supported with SSR
	ServicePort *int32 `json:"service_port,omitempty"`
	// only supported with SSR
	ServiceProtocol *string `json:"service_protocol,omitempty"`
	// only supported with SSR
	ServiceTenant *string `json:"service_tenant,omitempty"`
	// VRF Name
	Vrf *string `json:"vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UtilsShowForwardingTable UtilsShowForwardingTable

// NewUtilsShowForwardingTable instantiates a new UtilsShowForwardingTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtilsShowForwardingTable() *UtilsShowForwardingTable {
	this := UtilsShowForwardingTable{}
	return &this
}

// NewUtilsShowForwardingTableWithDefaults instantiates a new UtilsShowForwardingTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtilsShowForwardingTableWithDefaults() *UtilsShowForwardingTable {
	this := UtilsShowForwardingTable{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetNode() HaClusterNodeEnum {
	if o == nil || IsNil(o.Node) {
		var ret HaClusterNodeEnum
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetNodeOk() (*HaClusterNodeEnum, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given HaClusterNodeEnum and assigns it to the Node field.
func (o *UtilsShowForwardingTable) SetNode(v HaClusterNodeEnum) {
	o.Node = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *UtilsShowForwardingTable) SetPrefix(v string) {
	o.Prefix = &v
}

// GetServiceIp returns the ServiceIp field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetServiceIp() string {
	if o == nil || IsNil(o.ServiceIp) {
		var ret string
		return ret
	}
	return *o.ServiceIp
}

// GetServiceIpOk returns a tuple with the ServiceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetServiceIpOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceIp) {
		return nil, false
	}
	return o.ServiceIp, true
}

// HasServiceIp returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasServiceIp() bool {
	if o != nil && !IsNil(o.ServiceIp) {
		return true
	}

	return false
}

// SetServiceIp gets a reference to the given string and assigns it to the ServiceIp field.
func (o *UtilsShowForwardingTable) SetServiceIp(v string) {
	o.ServiceIp = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *UtilsShowForwardingTable) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetServicePort returns the ServicePort field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetServicePort() int32 {
	if o == nil || IsNil(o.ServicePort) {
		var ret int32
		return ret
	}
	return *o.ServicePort
}

// GetServicePortOk returns a tuple with the ServicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetServicePortOk() (*int32, bool) {
	if o == nil || IsNil(o.ServicePort) {
		return nil, false
	}
	return o.ServicePort, true
}

// HasServicePort returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasServicePort() bool {
	if o != nil && !IsNil(o.ServicePort) {
		return true
	}

	return false
}

// SetServicePort gets a reference to the given int32 and assigns it to the ServicePort field.
func (o *UtilsShowForwardingTable) SetServicePort(v int32) {
	o.ServicePort = &v
}

// GetServiceProtocol returns the ServiceProtocol field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetServiceProtocol() string {
	if o == nil || IsNil(o.ServiceProtocol) {
		var ret string
		return ret
	}
	return *o.ServiceProtocol
}

// GetServiceProtocolOk returns a tuple with the ServiceProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetServiceProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceProtocol) {
		return nil, false
	}
	return o.ServiceProtocol, true
}

// HasServiceProtocol returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasServiceProtocol() bool {
	if o != nil && !IsNil(o.ServiceProtocol) {
		return true
	}

	return false
}

// SetServiceProtocol gets a reference to the given string and assigns it to the ServiceProtocol field.
func (o *UtilsShowForwardingTable) SetServiceProtocol(v string) {
	o.ServiceProtocol = &v
}

// GetServiceTenant returns the ServiceTenant field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetServiceTenant() string {
	if o == nil || IsNil(o.ServiceTenant) {
		var ret string
		return ret
	}
	return *o.ServiceTenant
}

// GetServiceTenantOk returns a tuple with the ServiceTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetServiceTenantOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceTenant) {
		return nil, false
	}
	return o.ServiceTenant, true
}

// HasServiceTenant returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasServiceTenant() bool {
	if o != nil && !IsNil(o.ServiceTenant) {
		return true
	}

	return false
}

// SetServiceTenant gets a reference to the given string and assigns it to the ServiceTenant field.
func (o *UtilsShowForwardingTable) SetServiceTenant(v string) {
	o.ServiceTenant = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *UtilsShowForwardingTable) GetVrf() string {
	if o == nil || IsNil(o.Vrf) {
		var ret string
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtilsShowForwardingTable) GetVrfOk() (*string, bool) {
	if o == nil || IsNil(o.Vrf) {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *UtilsShowForwardingTable) HasVrf() bool {
	if o != nil && !IsNil(o.Vrf) {
		return true
	}

	return false
}

// SetVrf gets a reference to the given string and assigns it to the Vrf field.
func (o *UtilsShowForwardingTable) SetVrf(v string) {
	o.Vrf = &v
}

func (o UtilsShowForwardingTable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtilsShowForwardingTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.ServiceIp) {
		toSerialize["service_ip"] = o.ServiceIp
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !IsNil(o.ServicePort) {
		toSerialize["service_port"] = o.ServicePort
	}
	if !IsNil(o.ServiceProtocol) {
		toSerialize["service_protocol"] = o.ServiceProtocol
	}
	if !IsNil(o.ServiceTenant) {
		toSerialize["service_tenant"] = o.ServiceTenant
	}
	if !IsNil(o.Vrf) {
		toSerialize["vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UtilsShowForwardingTable) UnmarshalJSON(data []byte) (err error) {
	varUtilsShowForwardingTable := _UtilsShowForwardingTable{}

	err = json.Unmarshal(data, &varUtilsShowForwardingTable)

	if err != nil {
		return err
	}

	*o = UtilsShowForwardingTable(varUtilsShowForwardingTable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "node")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "service_ip")
		delete(additionalProperties, "service_name")
		delete(additionalProperties, "service_port")
		delete(additionalProperties, "service_protocol")
		delete(additionalProperties, "service_tenant")
		delete(additionalProperties, "vrf")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUtilsShowForwardingTable struct {
	value *UtilsShowForwardingTable
	isSet bool
}

func (v NullableUtilsShowForwardingTable) Get() *UtilsShowForwardingTable {
	return v.value
}

func (v *NullableUtilsShowForwardingTable) Set(val *UtilsShowForwardingTable) {
	v.value = val
	v.isSet = true
}

func (v NullableUtilsShowForwardingTable) IsSet() bool {
	return v.isSet
}

func (v *NullableUtilsShowForwardingTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtilsShowForwardingTable(val *UtilsShowForwardingTable) *NullableUtilsShowForwardingTable {
	return &NullableUtilsShowForwardingTable{value: val, isSet: true}
}

func (v NullableUtilsShowForwardingTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtilsShowForwardingTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


