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

// checks if the DhcpdConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpdConfig{}

// DhcpdConfig struct for DhcpdConfig
type DhcpdConfig struct {
	// if `type`==`local` - optional, if not defined, system one will be used
	DnsServers []string `json:"dns_servers,omitempty"`
	// if `type`==`local` - optional, if not defined, system one will be used
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// Property key is the MAC Address
	FixedBindings *map[string]DhcpdConfigFixedBinding `json:"fixed_bindings,omitempty"`
	// if `type`==`local` - optional, `ip` will be used if not provided
	Gateway *string `json:"gateway,omitempty"`
	// if `type`==`local`
	IpEnd *string `json:"ip_end,omitempty"`
	// if `type6`==`local`
	IpEnd6 *string `json:"ip_end6,omitempty"`
	// if `type`==`local`
	IpStart *string `json:"ip_start,omitempty"`
	// if `type6`==`local`
	IpStart6 *string `json:"ip_start6,omitempty"`
	// in seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day]
	LeaseTime *int32 `json:"lease_time,omitempty"`
	// Property key is the DHCP option number
	Options *map[string]DhcpdConfigOption `json:"options,omitempty"`
	// `server_id_override`==`true` means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,  should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address.
	ServerIdOverride *bool `json:"server_id_override,omitempty"`
	// if `type`==`relay`
	Servers []string `json:"servers,omitempty"`
	// if `type6`==`relay`
	Servers6 []string `json:"servers6,omitempty"`
	Type *DhcpdConfigType `json:"type,omitempty"`
	Type6 *DhcpdConfigType6 `json:"type6,omitempty"`
	// Property key is <enterprise number>:<sub option code>, with * enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers) * sub option code: 1-255, sub-option code
	VendorEncapulated *map[string]DhcpdConfigVendorOption `json:"vendor_encapulated,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DhcpdConfig DhcpdConfig

// NewDhcpdConfig instantiates a new DhcpdConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpdConfig() *DhcpdConfig {
	this := DhcpdConfig{}
	var leaseTime int32 = 86400
	this.LeaseTime = &leaseTime
	var serverIdOverride bool = false
	this.ServerIdOverride = &serverIdOverride
	var type_ DhcpdConfigType = DHCPDCONFIGTYPE_LOCAL
	this.Type = &type_
	var type6 DhcpdConfigType6 = DHCPDCONFIGTYPE6_NONE
	this.Type6 = &type6
	return &this
}

// NewDhcpdConfigWithDefaults instantiates a new DhcpdConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpdConfigWithDefaults() *DhcpdConfig {
	this := DhcpdConfig{}
	var leaseTime int32 = 86400
	this.LeaseTime = &leaseTime
	var serverIdOverride bool = false
	this.ServerIdOverride = &serverIdOverride
	var type_ DhcpdConfigType = DHCPDCONFIGTYPE_LOCAL
	this.Type = &type_
	var type6 DhcpdConfigType6 = DHCPDCONFIGTYPE6_NONE
	this.Type6 = &type6
	return &this
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *DhcpdConfig) GetDnsServers() []string {
	if o == nil || IsNil(o.DnsServers) {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetDnsServersOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *DhcpdConfig) HasDnsServers() bool {
	if o != nil && !IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *DhcpdConfig) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetDnsSuffix returns the DnsSuffix field value if set, zero value otherwise.
func (o *DhcpdConfig) GetDnsSuffix() []string {
	if o == nil || IsNil(o.DnsSuffix) {
		var ret []string
		return ret
	}
	return o.DnsSuffix
}

// GetDnsSuffixOk returns a tuple with the DnsSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetDnsSuffixOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsSuffix) {
		return nil, false
	}
	return o.DnsSuffix, true
}

// HasDnsSuffix returns a boolean if a field has been set.
func (o *DhcpdConfig) HasDnsSuffix() bool {
	if o != nil && !IsNil(o.DnsSuffix) {
		return true
	}

	return false
}

// SetDnsSuffix gets a reference to the given []string and assigns it to the DnsSuffix field.
func (o *DhcpdConfig) SetDnsSuffix(v []string) {
	o.DnsSuffix = v
}

// GetFixedBindings returns the FixedBindings field value if set, zero value otherwise.
func (o *DhcpdConfig) GetFixedBindings() map[string]DhcpdConfigFixedBinding {
	if o == nil || IsNil(o.FixedBindings) {
		var ret map[string]DhcpdConfigFixedBinding
		return ret
	}
	return *o.FixedBindings
}

// GetFixedBindingsOk returns a tuple with the FixedBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetFixedBindingsOk() (*map[string]DhcpdConfigFixedBinding, bool) {
	if o == nil || IsNil(o.FixedBindings) {
		return nil, false
	}
	return o.FixedBindings, true
}

// HasFixedBindings returns a boolean if a field has been set.
func (o *DhcpdConfig) HasFixedBindings() bool {
	if o != nil && !IsNil(o.FixedBindings) {
		return true
	}

	return false
}

// SetFixedBindings gets a reference to the given map[string]DhcpdConfigFixedBinding and assigns it to the FixedBindings field.
func (o *DhcpdConfig) SetFixedBindings(v map[string]DhcpdConfigFixedBinding) {
	o.FixedBindings = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *DhcpdConfig) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *DhcpdConfig) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *DhcpdConfig) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpEnd returns the IpEnd field value if set, zero value otherwise.
func (o *DhcpdConfig) GetIpEnd() string {
	if o == nil || IsNil(o.IpEnd) {
		var ret string
		return ret
	}
	return *o.IpEnd
}

// GetIpEndOk returns a tuple with the IpEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetIpEndOk() (*string, bool) {
	if o == nil || IsNil(o.IpEnd) {
		return nil, false
	}
	return o.IpEnd, true
}

// HasIpEnd returns a boolean if a field has been set.
func (o *DhcpdConfig) HasIpEnd() bool {
	if o != nil && !IsNil(o.IpEnd) {
		return true
	}

	return false
}

// SetIpEnd gets a reference to the given string and assigns it to the IpEnd field.
func (o *DhcpdConfig) SetIpEnd(v string) {
	o.IpEnd = &v
}

// GetIpEnd6 returns the IpEnd6 field value if set, zero value otherwise.
func (o *DhcpdConfig) GetIpEnd6() string {
	if o == nil || IsNil(o.IpEnd6) {
		var ret string
		return ret
	}
	return *o.IpEnd6
}

// GetIpEnd6Ok returns a tuple with the IpEnd6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetIpEnd6Ok() (*string, bool) {
	if o == nil || IsNil(o.IpEnd6) {
		return nil, false
	}
	return o.IpEnd6, true
}

// HasIpEnd6 returns a boolean if a field has been set.
func (o *DhcpdConfig) HasIpEnd6() bool {
	if o != nil && !IsNil(o.IpEnd6) {
		return true
	}

	return false
}

// SetIpEnd6 gets a reference to the given string and assigns it to the IpEnd6 field.
func (o *DhcpdConfig) SetIpEnd6(v string) {
	o.IpEnd6 = &v
}

// GetIpStart returns the IpStart field value if set, zero value otherwise.
func (o *DhcpdConfig) GetIpStart() string {
	if o == nil || IsNil(o.IpStart) {
		var ret string
		return ret
	}
	return *o.IpStart
}

// GetIpStartOk returns a tuple with the IpStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetIpStartOk() (*string, bool) {
	if o == nil || IsNil(o.IpStart) {
		return nil, false
	}
	return o.IpStart, true
}

// HasIpStart returns a boolean if a field has been set.
func (o *DhcpdConfig) HasIpStart() bool {
	if o != nil && !IsNil(o.IpStart) {
		return true
	}

	return false
}

// SetIpStart gets a reference to the given string and assigns it to the IpStart field.
func (o *DhcpdConfig) SetIpStart(v string) {
	o.IpStart = &v
}

// GetIpStart6 returns the IpStart6 field value if set, zero value otherwise.
func (o *DhcpdConfig) GetIpStart6() string {
	if o == nil || IsNil(o.IpStart6) {
		var ret string
		return ret
	}
	return *o.IpStart6
}

// GetIpStart6Ok returns a tuple with the IpStart6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetIpStart6Ok() (*string, bool) {
	if o == nil || IsNil(o.IpStart6) {
		return nil, false
	}
	return o.IpStart6, true
}

// HasIpStart6 returns a boolean if a field has been set.
func (o *DhcpdConfig) HasIpStart6() bool {
	if o != nil && !IsNil(o.IpStart6) {
		return true
	}

	return false
}

// SetIpStart6 gets a reference to the given string and assigns it to the IpStart6 field.
func (o *DhcpdConfig) SetIpStart6(v string) {
	o.IpStart6 = &v
}

// GetLeaseTime returns the LeaseTime field value if set, zero value otherwise.
func (o *DhcpdConfig) GetLeaseTime() int32 {
	if o == nil || IsNil(o.LeaseTime) {
		var ret int32
		return ret
	}
	return *o.LeaseTime
}

// GetLeaseTimeOk returns a tuple with the LeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetLeaseTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.LeaseTime) {
		return nil, false
	}
	return o.LeaseTime, true
}

// HasLeaseTime returns a boolean if a field has been set.
func (o *DhcpdConfig) HasLeaseTime() bool {
	if o != nil && !IsNil(o.LeaseTime) {
		return true
	}

	return false
}

// SetLeaseTime gets a reference to the given int32 and assigns it to the LeaseTime field.
func (o *DhcpdConfig) SetLeaseTime(v int32) {
	o.LeaseTime = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DhcpdConfig) GetOptions() map[string]DhcpdConfigOption {
	if o == nil || IsNil(o.Options) {
		var ret map[string]DhcpdConfigOption
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetOptionsOk() (*map[string]DhcpdConfigOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DhcpdConfig) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]DhcpdConfigOption and assigns it to the Options field.
func (o *DhcpdConfig) SetOptions(v map[string]DhcpdConfigOption) {
	o.Options = &v
}

// GetServerIdOverride returns the ServerIdOverride field value if set, zero value otherwise.
func (o *DhcpdConfig) GetServerIdOverride() bool {
	if o == nil || IsNil(o.ServerIdOverride) {
		var ret bool
		return ret
	}
	return *o.ServerIdOverride
}

// GetServerIdOverrideOk returns a tuple with the ServerIdOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetServerIdOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerIdOverride) {
		return nil, false
	}
	return o.ServerIdOverride, true
}

// HasServerIdOverride returns a boolean if a field has been set.
func (o *DhcpdConfig) HasServerIdOverride() bool {
	if o != nil && !IsNil(o.ServerIdOverride) {
		return true
	}

	return false
}

// SetServerIdOverride gets a reference to the given bool and assigns it to the ServerIdOverride field.
func (o *DhcpdConfig) SetServerIdOverride(v bool) {
	o.ServerIdOverride = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *DhcpdConfig) GetServers() []string {
	if o == nil || IsNil(o.Servers) {
		var ret []string
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetServersOk() ([]string, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *DhcpdConfig) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []string and assigns it to the Servers field.
func (o *DhcpdConfig) SetServers(v []string) {
	o.Servers = v
}

// GetServers6 returns the Servers6 field value if set, zero value otherwise.
func (o *DhcpdConfig) GetServers6() []string {
	if o == nil || IsNil(o.Servers6) {
		var ret []string
		return ret
	}
	return o.Servers6
}

// GetServers6Ok returns a tuple with the Servers6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetServers6Ok() ([]string, bool) {
	if o == nil || IsNil(o.Servers6) {
		return nil, false
	}
	return o.Servers6, true
}

// HasServers6 returns a boolean if a field has been set.
func (o *DhcpdConfig) HasServers6() bool {
	if o != nil && !IsNil(o.Servers6) {
		return true
	}

	return false
}

// SetServers6 gets a reference to the given []string and assigns it to the Servers6 field.
func (o *DhcpdConfig) SetServers6(v []string) {
	o.Servers6 = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DhcpdConfig) GetType() DhcpdConfigType {
	if o == nil || IsNil(o.Type) {
		var ret DhcpdConfigType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetTypeOk() (*DhcpdConfigType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DhcpdConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DhcpdConfigType and assigns it to the Type field.
func (o *DhcpdConfig) SetType(v DhcpdConfigType) {
	o.Type = &v
}

// GetType6 returns the Type6 field value if set, zero value otherwise.
func (o *DhcpdConfig) GetType6() DhcpdConfigType6 {
	if o == nil || IsNil(o.Type6) {
		var ret DhcpdConfigType6
		return ret
	}
	return *o.Type6
}

// GetType6Ok returns a tuple with the Type6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetType6Ok() (*DhcpdConfigType6, bool) {
	if o == nil || IsNil(o.Type6) {
		return nil, false
	}
	return o.Type6, true
}

// HasType6 returns a boolean if a field has been set.
func (o *DhcpdConfig) HasType6() bool {
	if o != nil && !IsNil(o.Type6) {
		return true
	}

	return false
}

// SetType6 gets a reference to the given DhcpdConfigType6 and assigns it to the Type6 field.
func (o *DhcpdConfig) SetType6(v DhcpdConfigType6) {
	o.Type6 = &v
}

// GetVendorEncapulated returns the VendorEncapulated field value if set, zero value otherwise.
func (o *DhcpdConfig) GetVendorEncapulated() map[string]DhcpdConfigVendorOption {
	if o == nil || IsNil(o.VendorEncapulated) {
		var ret map[string]DhcpdConfigVendorOption
		return ret
	}
	return *o.VendorEncapulated
}

// GetVendorEncapulatedOk returns a tuple with the VendorEncapulated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpdConfig) GetVendorEncapulatedOk() (*map[string]DhcpdConfigVendorOption, bool) {
	if o == nil || IsNil(o.VendorEncapulated) {
		return nil, false
	}
	return o.VendorEncapulated, true
}

// HasVendorEncapulated returns a boolean if a field has been set.
func (o *DhcpdConfig) HasVendorEncapulated() bool {
	if o != nil && !IsNil(o.VendorEncapulated) {
		return true
	}

	return false
}

// SetVendorEncapulated gets a reference to the given map[string]DhcpdConfigVendorOption and assigns it to the VendorEncapulated field.
func (o *DhcpdConfig) SetVendorEncapulated(v map[string]DhcpdConfigVendorOption) {
	o.VendorEncapulated = &v
}

func (o DhcpdConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpdConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnsServers) {
		toSerialize["dns_servers"] = o.DnsServers
	}
	if !IsNil(o.DnsSuffix) {
		toSerialize["dns_suffix"] = o.DnsSuffix
	}
	if !IsNil(o.FixedBindings) {
		toSerialize["fixed_bindings"] = o.FixedBindings
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.IpEnd) {
		toSerialize["ip_end"] = o.IpEnd
	}
	if !IsNil(o.IpEnd6) {
		toSerialize["ip_end6"] = o.IpEnd6
	}
	if !IsNil(o.IpStart) {
		toSerialize["ip_start"] = o.IpStart
	}
	if !IsNil(o.IpStart6) {
		toSerialize["ip_start6"] = o.IpStart6
	}
	if !IsNil(o.LeaseTime) {
		toSerialize["lease_time"] = o.LeaseTime
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.ServerIdOverride) {
		toSerialize["server_id_override"] = o.ServerIdOverride
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Servers6) {
		toSerialize["servers6"] = o.Servers6
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Type6) {
		toSerialize["type6"] = o.Type6
	}
	if !IsNil(o.VendorEncapulated) {
		toSerialize["vendor_encapulated"] = o.VendorEncapulated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DhcpdConfig) UnmarshalJSON(data []byte) (err error) {
	varDhcpdConfig := _DhcpdConfig{}

	err = json.Unmarshal(data, &varDhcpdConfig)

	if err != nil {
		return err
	}

	*o = DhcpdConfig(varDhcpdConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_servers")
		delete(additionalProperties, "dns_suffix")
		delete(additionalProperties, "fixed_bindings")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "ip_end")
		delete(additionalProperties, "ip_end6")
		delete(additionalProperties, "ip_start")
		delete(additionalProperties, "ip_start6")
		delete(additionalProperties, "lease_time")
		delete(additionalProperties, "options")
		delete(additionalProperties, "server_id_override")
		delete(additionalProperties, "servers")
		delete(additionalProperties, "servers6")
		delete(additionalProperties, "type")
		delete(additionalProperties, "type6")
		delete(additionalProperties, "vendor_encapulated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDhcpdConfig struct {
	value *DhcpdConfig
	isSet bool
}

func (v NullableDhcpdConfig) Get() *DhcpdConfig {
	return v.value
}

func (v *NullableDhcpdConfig) Set(val *DhcpdConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpdConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpdConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpdConfig(val *DhcpdConfig) *NullableDhcpdConfig {
	return &NullableDhcpdConfig{value: val, isSet: true}
}

func (v NullableDhcpdConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpdConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


