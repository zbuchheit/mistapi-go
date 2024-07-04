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

// checks if the MxtunnelIpsec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxtunnelIpsec{}

// MxtunnelIpsec struct for MxtunnelIpsec
type MxtunnelIpsec struct {
	DnsServers []string `json:"dns_servers,omitempty"`
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ExtraRoutes []MxtunnelIpsecExtraRoute `json:"extra_routes,omitempty"`
	SplitTunnel *bool `json:"split_tunnel,omitempty"`
	UseMxedge *bool `json:"use_mxedge,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxtunnelIpsec MxtunnelIpsec

// NewMxtunnelIpsec instantiates a new MxtunnelIpsec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxtunnelIpsec() *MxtunnelIpsec {
	this := MxtunnelIpsec{}
	return &this
}

// NewMxtunnelIpsecWithDefaults instantiates a new MxtunnelIpsec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxtunnelIpsecWithDefaults() *MxtunnelIpsec {
	this := MxtunnelIpsec{}
	return &this
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MxtunnelIpsec) GetDnsServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MxtunnelIpsec) GetDnsServersOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *MxtunnelIpsec) HasDnsServers() bool {
	if o != nil && !IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *MxtunnelIpsec) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetDnsSuffix returns the DnsSuffix field value if set, zero value otherwise.
func (o *MxtunnelIpsec) GetDnsSuffix() []string {
	if o == nil || IsNil(o.DnsSuffix) {
		var ret []string
		return ret
	}
	return o.DnsSuffix
}

// GetDnsSuffixOk returns a tuple with the DnsSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxtunnelIpsec) GetDnsSuffixOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsSuffix) {
		return nil, false
	}
	return o.DnsSuffix, true
}

// HasDnsSuffix returns a boolean if a field has been set.
func (o *MxtunnelIpsec) HasDnsSuffix() bool {
	if o != nil && !IsNil(o.DnsSuffix) {
		return true
	}

	return false
}

// SetDnsSuffix gets a reference to the given []string and assigns it to the DnsSuffix field.
func (o *MxtunnelIpsec) SetDnsSuffix(v []string) {
	o.DnsSuffix = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MxtunnelIpsec) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxtunnelIpsec) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MxtunnelIpsec) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MxtunnelIpsec) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExtraRoutes returns the ExtraRoutes field value if set, zero value otherwise.
func (o *MxtunnelIpsec) GetExtraRoutes() []MxtunnelIpsecExtraRoute {
	if o == nil || IsNil(o.ExtraRoutes) {
		var ret []MxtunnelIpsecExtraRoute
		return ret
	}
	return o.ExtraRoutes
}

// GetExtraRoutesOk returns a tuple with the ExtraRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxtunnelIpsec) GetExtraRoutesOk() ([]MxtunnelIpsecExtraRoute, bool) {
	if o == nil || IsNil(o.ExtraRoutes) {
		return nil, false
	}
	return o.ExtraRoutes, true
}

// HasExtraRoutes returns a boolean if a field has been set.
func (o *MxtunnelIpsec) HasExtraRoutes() bool {
	if o != nil && !IsNil(o.ExtraRoutes) {
		return true
	}

	return false
}

// SetExtraRoutes gets a reference to the given []MxtunnelIpsecExtraRoute and assigns it to the ExtraRoutes field.
func (o *MxtunnelIpsec) SetExtraRoutes(v []MxtunnelIpsecExtraRoute) {
	o.ExtraRoutes = v
}

// GetSplitTunnel returns the SplitTunnel field value if set, zero value otherwise.
func (o *MxtunnelIpsec) GetSplitTunnel() bool {
	if o == nil || IsNil(o.SplitTunnel) {
		var ret bool
		return ret
	}
	return *o.SplitTunnel
}

// GetSplitTunnelOk returns a tuple with the SplitTunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxtunnelIpsec) GetSplitTunnelOk() (*bool, bool) {
	if o == nil || IsNil(o.SplitTunnel) {
		return nil, false
	}
	return o.SplitTunnel, true
}

// HasSplitTunnel returns a boolean if a field has been set.
func (o *MxtunnelIpsec) HasSplitTunnel() bool {
	if o != nil && !IsNil(o.SplitTunnel) {
		return true
	}

	return false
}

// SetSplitTunnel gets a reference to the given bool and assigns it to the SplitTunnel field.
func (o *MxtunnelIpsec) SetSplitTunnel(v bool) {
	o.SplitTunnel = &v
}

// GetUseMxedge returns the UseMxedge field value if set, zero value otherwise.
func (o *MxtunnelIpsec) GetUseMxedge() bool {
	if o == nil || IsNil(o.UseMxedge) {
		var ret bool
		return ret
	}
	return *o.UseMxedge
}

// GetUseMxedgeOk returns a tuple with the UseMxedge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxtunnelIpsec) GetUseMxedgeOk() (*bool, bool) {
	if o == nil || IsNil(o.UseMxedge) {
		return nil, false
	}
	return o.UseMxedge, true
}

// HasUseMxedge returns a boolean if a field has been set.
func (o *MxtunnelIpsec) HasUseMxedge() bool {
	if o != nil && !IsNil(o.UseMxedge) {
		return true
	}

	return false
}

// SetUseMxedge gets a reference to the given bool and assigns it to the UseMxedge field.
func (o *MxtunnelIpsec) SetUseMxedge(v bool) {
	o.UseMxedge = &v
}

func (o MxtunnelIpsec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxtunnelIpsec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsServers != nil {
		toSerialize["dns_servers"] = o.DnsServers
	}
	if !IsNil(o.DnsSuffix) {
		toSerialize["dns_suffix"] = o.DnsSuffix
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ExtraRoutes) {
		toSerialize["extra_routes"] = o.ExtraRoutes
	}
	if !IsNil(o.SplitTunnel) {
		toSerialize["split_tunnel"] = o.SplitTunnel
	}
	if !IsNil(o.UseMxedge) {
		toSerialize["use_mxedge"] = o.UseMxedge
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxtunnelIpsec) UnmarshalJSON(data []byte) (err error) {
	varMxtunnelIpsec := _MxtunnelIpsec{}

	err = json.Unmarshal(data, &varMxtunnelIpsec)

	if err != nil {
		return err
	}

	*o = MxtunnelIpsec(varMxtunnelIpsec)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_servers")
		delete(additionalProperties, "dns_suffix")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "extra_routes")
		delete(additionalProperties, "split_tunnel")
		delete(additionalProperties, "use_mxedge")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxtunnelIpsec struct {
	value *MxtunnelIpsec
	isSet bool
}

func (v NullableMxtunnelIpsec) Get() *MxtunnelIpsec {
	return v.value
}

func (v *NullableMxtunnelIpsec) Set(val *MxtunnelIpsec) {
	v.value = val
	v.isSet = true
}

func (v NullableMxtunnelIpsec) IsSet() bool {
	return v.isSet
}

func (v *NullableMxtunnelIpsec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxtunnelIpsec(val *MxtunnelIpsec) *NullableMxtunnelIpsec {
	return &NullableMxtunnelIpsec{value: val, isSet: true}
}

func (v NullableMxtunnelIpsec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxtunnelIpsec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


