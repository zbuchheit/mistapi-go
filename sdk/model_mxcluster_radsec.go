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

// checks if the MxclusterRadsec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxclusterRadsec{}

// MxclusterRadsec MxEdge Radsec Configuration
type MxclusterRadsec struct {
	// list of RADIUS accounting servers, optional, order matters where the first one is treated as primary
	AcctServers []MxclusterRadsecAcctServer `json:"acct_servers,omitempty"`
	// list of RADIUS authentication servers, order matters where the first one is treated as primary
	AuthServers []MxclusterRadsecAuthServer `json:"auth_servers,omitempty"`
	// whether to enable service on Mist Edge i.e. RADIUS proxy over TLS
	Enabled *bool `json:"enabled,omitempty"`
	// whether to match ssid in request message to select from a subset of RADIUS servers
	MatchSsid *bool `json:"match_ssid,omitempty"`
	// hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts`
	ProxyHosts []string `json:"proxy_hosts,omitempty"`
	ServerSelection *MxclusterRadsecServerSelection `json:"server_selection,omitempty"`
	Source *MxclusterRadsecSource `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxclusterRadsec MxclusterRadsec

// NewMxclusterRadsec instantiates a new MxclusterRadsec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxclusterRadsec() *MxclusterRadsec {
	this := MxclusterRadsec{}
	var serverSelection MxclusterRadsecServerSelection = MXCLUSTERRADSECSERVERSELECTION_ORDERED
	this.ServerSelection = &serverSelection
	var source MxclusterRadsecSource = MXCLUSTERRADSECSOURCE_ANY
	this.Source = &source
	return &this
}

// NewMxclusterRadsecWithDefaults instantiates a new MxclusterRadsec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxclusterRadsecWithDefaults() *MxclusterRadsec {
	this := MxclusterRadsec{}
	var serverSelection MxclusterRadsecServerSelection = MXCLUSTERRADSECSERVERSELECTION_ORDERED
	this.ServerSelection = &serverSelection
	var source MxclusterRadsecSource = MXCLUSTERRADSECSOURCE_ANY
	this.Source = &source
	return &this
}

// GetAcctServers returns the AcctServers field value if set, zero value otherwise.
func (o *MxclusterRadsec) GetAcctServers() []MxclusterRadsecAcctServer {
	if o == nil || IsNil(o.AcctServers) {
		var ret []MxclusterRadsecAcctServer
		return ret
	}
	return o.AcctServers
}

// GetAcctServersOk returns a tuple with the AcctServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxclusterRadsec) GetAcctServersOk() ([]MxclusterRadsecAcctServer, bool) {
	if o == nil || IsNil(o.AcctServers) {
		return nil, false
	}
	return o.AcctServers, true
}

// HasAcctServers returns a boolean if a field has been set.
func (o *MxclusterRadsec) HasAcctServers() bool {
	if o != nil && !IsNil(o.AcctServers) {
		return true
	}

	return false
}

// SetAcctServers gets a reference to the given []MxclusterRadsecAcctServer and assigns it to the AcctServers field.
func (o *MxclusterRadsec) SetAcctServers(v []MxclusterRadsecAcctServer) {
	o.AcctServers = v
}

// GetAuthServers returns the AuthServers field value if set, zero value otherwise.
func (o *MxclusterRadsec) GetAuthServers() []MxclusterRadsecAuthServer {
	if o == nil || IsNil(o.AuthServers) {
		var ret []MxclusterRadsecAuthServer
		return ret
	}
	return o.AuthServers
}

// GetAuthServersOk returns a tuple with the AuthServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxclusterRadsec) GetAuthServersOk() ([]MxclusterRadsecAuthServer, bool) {
	if o == nil || IsNil(o.AuthServers) {
		return nil, false
	}
	return o.AuthServers, true
}

// HasAuthServers returns a boolean if a field has been set.
func (o *MxclusterRadsec) HasAuthServers() bool {
	if o != nil && !IsNil(o.AuthServers) {
		return true
	}

	return false
}

// SetAuthServers gets a reference to the given []MxclusterRadsecAuthServer and assigns it to the AuthServers field.
func (o *MxclusterRadsec) SetAuthServers(v []MxclusterRadsecAuthServer) {
	o.AuthServers = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MxclusterRadsec) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxclusterRadsec) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MxclusterRadsec) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MxclusterRadsec) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMatchSsid returns the MatchSsid field value if set, zero value otherwise.
func (o *MxclusterRadsec) GetMatchSsid() bool {
	if o == nil || IsNil(o.MatchSsid) {
		var ret bool
		return ret
	}
	return *o.MatchSsid
}

// GetMatchSsidOk returns a tuple with the MatchSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxclusterRadsec) GetMatchSsidOk() (*bool, bool) {
	if o == nil || IsNil(o.MatchSsid) {
		return nil, false
	}
	return o.MatchSsid, true
}

// HasMatchSsid returns a boolean if a field has been set.
func (o *MxclusterRadsec) HasMatchSsid() bool {
	if o != nil && !IsNil(o.MatchSsid) {
		return true
	}

	return false
}

// SetMatchSsid gets a reference to the given bool and assigns it to the MatchSsid field.
func (o *MxclusterRadsec) SetMatchSsid(v bool) {
	o.MatchSsid = &v
}

// GetProxyHosts returns the ProxyHosts field value if set, zero value otherwise.
func (o *MxclusterRadsec) GetProxyHosts() []string {
	if o == nil || IsNil(o.ProxyHosts) {
		var ret []string
		return ret
	}
	return o.ProxyHosts
}

// GetProxyHostsOk returns a tuple with the ProxyHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxclusterRadsec) GetProxyHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProxyHosts) {
		return nil, false
	}
	return o.ProxyHosts, true
}

// HasProxyHosts returns a boolean if a field has been set.
func (o *MxclusterRadsec) HasProxyHosts() bool {
	if o != nil && !IsNil(o.ProxyHosts) {
		return true
	}

	return false
}

// SetProxyHosts gets a reference to the given []string and assigns it to the ProxyHosts field.
func (o *MxclusterRadsec) SetProxyHosts(v []string) {
	o.ProxyHosts = v
}

// GetServerSelection returns the ServerSelection field value if set, zero value otherwise.
func (o *MxclusterRadsec) GetServerSelection() MxclusterRadsecServerSelection {
	if o == nil || IsNil(o.ServerSelection) {
		var ret MxclusterRadsecServerSelection
		return ret
	}
	return *o.ServerSelection
}

// GetServerSelectionOk returns a tuple with the ServerSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxclusterRadsec) GetServerSelectionOk() (*MxclusterRadsecServerSelection, bool) {
	if o == nil || IsNil(o.ServerSelection) {
		return nil, false
	}
	return o.ServerSelection, true
}

// HasServerSelection returns a boolean if a field has been set.
func (o *MxclusterRadsec) HasServerSelection() bool {
	if o != nil && !IsNil(o.ServerSelection) {
		return true
	}

	return false
}

// SetServerSelection gets a reference to the given MxclusterRadsecServerSelection and assigns it to the ServerSelection field.
func (o *MxclusterRadsec) SetServerSelection(v MxclusterRadsecServerSelection) {
	o.ServerSelection = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MxclusterRadsec) GetSource() MxclusterRadsecSource {
	if o == nil || IsNil(o.Source) {
		var ret MxclusterRadsecSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxclusterRadsec) GetSourceOk() (*MxclusterRadsecSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MxclusterRadsec) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given MxclusterRadsecSource and assigns it to the Source field.
func (o *MxclusterRadsec) SetSource(v MxclusterRadsecSource) {
	o.Source = &v
}

func (o MxclusterRadsec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxclusterRadsec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcctServers) {
		toSerialize["acct_servers"] = o.AcctServers
	}
	if !IsNil(o.AuthServers) {
		toSerialize["auth_servers"] = o.AuthServers
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MatchSsid) {
		toSerialize["match_ssid"] = o.MatchSsid
	}
	if !IsNil(o.ProxyHosts) {
		toSerialize["proxy_hosts"] = o.ProxyHosts
	}
	if !IsNil(o.ServerSelection) {
		toSerialize["server_selection"] = o.ServerSelection
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxclusterRadsec) UnmarshalJSON(data []byte) (err error) {
	varMxclusterRadsec := _MxclusterRadsec{}

	err = json.Unmarshal(data, &varMxclusterRadsec)

	if err != nil {
		return err
	}

	*o = MxclusterRadsec(varMxclusterRadsec)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acct_servers")
		delete(additionalProperties, "auth_servers")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "match_ssid")
		delete(additionalProperties, "proxy_hosts")
		delete(additionalProperties, "server_selection")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxclusterRadsec struct {
	value *MxclusterRadsec
	isSet bool
}

func (v NullableMxclusterRadsec) Get() *MxclusterRadsec {
	return v.value
}

func (v *NullableMxclusterRadsec) Set(val *MxclusterRadsec) {
	v.value = val
	v.isSet = true
}

func (v NullableMxclusterRadsec) IsSet() bool {
	return v.isSet
}

func (v *NullableMxclusterRadsec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxclusterRadsec(val *MxclusterRadsec) *NullableMxclusterRadsec {
	return &NullableMxclusterRadsec{value: val, isSet: true}
}

func (v NullableMxclusterRadsec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxclusterRadsec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


