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

// checks if the NetworkTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkTemplate{}

// NetworkTemplate Network Template
type NetworkTemplate struct {
	AclPolicies []AclPolicy `json:"acl_policies,omitempty"`
	// ACL Tags to identify traffic source or destination. Key name is the tag name
	AclTags *map[string]AclTag `json:"acl_tags,omitempty"`
	// additional CLI commands to append to the generated Junos config  **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	DhcpSnooping *DhcpSnooping `json:"dhcp_snooping,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	ExtraRoutes *map[string]ExtraRouteProperties `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. \"2a02:1234:420a:10c9::/64\")
	ExtraRoutes6 *map[string]ExtraRoute6Properties `json:"extra_routes6,omitempty"`
	Id *string `json:"id,omitempty"`
	// Org Networks that we'd like to import
	ImportOrgNetworks []string `json:"import_org_networks,omitempty"`
	MistNac *SwitchMistNac `json:"mist_nac,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	Name *string `json:"name,omitempty"`
	// Property key is network name
	Networks *map[string]SwitchNetwork `json:"networks,omitempty"`
	// list of NTP servers specific to this device. By default, those in Site Settings will be used
	NtpServers []string `json:"ntp_servers,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	// Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.
	PortMirrorings *map[string]SwitchPortMirroring `json:"port_mirrorings,omitempty"`
	PortUsages *map[string]SwitchPortUsage `json:"port_usages,omitempty"`
	RadiusConfig *RadiusConfig `json:"radius_config,omitempty"`
	RemoteSyslog *RemoteSyslog `json:"remote_syslog,omitempty"`
	SnmpConfig *SnmpConfig `json:"snmp_config,omitempty"`
	SwitchMatching *SwitchMatching `json:"switch_matching,omitempty"`
	SwitchMgmt *SwitchMgmt `json:"switch_mgmt,omitempty"`
	VrfConfig *VrfConfig `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances *map[string]VrfInstance `json:"vrf_instances,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkTemplate NetworkTemplate

// NewNetworkTemplate instantiates a new NetworkTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkTemplate() *NetworkTemplate {
	this := NetworkTemplate{}
	return &this
}

// NewNetworkTemplateWithDefaults instantiates a new NetworkTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkTemplateWithDefaults() *NetworkTemplate {
	this := NetworkTemplate{}
	return &this
}

// GetAclPolicies returns the AclPolicies field value if set, zero value otherwise.
func (o *NetworkTemplate) GetAclPolicies() []AclPolicy {
	if o == nil || IsNil(o.AclPolicies) {
		var ret []AclPolicy
		return ret
	}
	return o.AclPolicies
}

// GetAclPoliciesOk returns a tuple with the AclPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetAclPoliciesOk() ([]AclPolicy, bool) {
	if o == nil || IsNil(o.AclPolicies) {
		return nil, false
	}
	return o.AclPolicies, true
}

// HasAclPolicies returns a boolean if a field has been set.
func (o *NetworkTemplate) HasAclPolicies() bool {
	if o != nil && !IsNil(o.AclPolicies) {
		return true
	}

	return false
}

// SetAclPolicies gets a reference to the given []AclPolicy and assigns it to the AclPolicies field.
func (o *NetworkTemplate) SetAclPolicies(v []AclPolicy) {
	o.AclPolicies = v
}

// GetAclTags returns the AclTags field value if set, zero value otherwise.
func (o *NetworkTemplate) GetAclTags() map[string]AclTag {
	if o == nil || IsNil(o.AclTags) {
		var ret map[string]AclTag
		return ret
	}
	return *o.AclTags
}

// GetAclTagsOk returns a tuple with the AclTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetAclTagsOk() (*map[string]AclTag, bool) {
	if o == nil || IsNil(o.AclTags) {
		return nil, false
	}
	return o.AclTags, true
}

// HasAclTags returns a boolean if a field has been set.
func (o *NetworkTemplate) HasAclTags() bool {
	if o != nil && !IsNil(o.AclTags) {
		return true
	}

	return false
}

// SetAclTags gets a reference to the given map[string]AclTag and assigns it to the AclTags field.
func (o *NetworkTemplate) SetAclTags(v map[string]AclTag) {
	o.AclTags = &v
}

// GetAdditionalConfigCmds returns the AdditionalConfigCmds field value if set, zero value otherwise.
func (o *NetworkTemplate) GetAdditionalConfigCmds() []string {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		var ret []string
		return ret
	}
	return o.AdditionalConfigCmds
}

// GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetAdditionalConfigCmdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		return nil, false
	}
	return o.AdditionalConfigCmds, true
}

// HasAdditionalConfigCmds returns a boolean if a field has been set.
func (o *NetworkTemplate) HasAdditionalConfigCmds() bool {
	if o != nil && !IsNil(o.AdditionalConfigCmds) {
		return true
	}

	return false
}

// SetAdditionalConfigCmds gets a reference to the given []string and assigns it to the AdditionalConfigCmds field.
func (o *NetworkTemplate) SetAdditionalConfigCmds(v []string) {
	o.AdditionalConfigCmds = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *NetworkTemplate) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *NetworkTemplate) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *NetworkTemplate) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetDhcpSnooping returns the DhcpSnooping field value if set, zero value otherwise.
func (o *NetworkTemplate) GetDhcpSnooping() DhcpSnooping {
	if o == nil || IsNil(o.DhcpSnooping) {
		var ret DhcpSnooping
		return ret
	}
	return *o.DhcpSnooping
}

// GetDhcpSnoopingOk returns a tuple with the DhcpSnooping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetDhcpSnoopingOk() (*DhcpSnooping, bool) {
	if o == nil || IsNil(o.DhcpSnooping) {
		return nil, false
	}
	return o.DhcpSnooping, true
}

// HasDhcpSnooping returns a boolean if a field has been set.
func (o *NetworkTemplate) HasDhcpSnooping() bool {
	if o != nil && !IsNil(o.DhcpSnooping) {
		return true
	}

	return false
}

// SetDhcpSnooping gets a reference to the given DhcpSnooping and assigns it to the DhcpSnooping field.
func (o *NetworkTemplate) SetDhcpSnooping(v DhcpSnooping) {
	o.DhcpSnooping = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *NetworkTemplate) GetDnsServers() []string {
	if o == nil || IsNil(o.DnsServers) {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetDnsServersOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *NetworkTemplate) HasDnsServers() bool {
	if o != nil && !IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *NetworkTemplate) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetDnsSuffix returns the DnsSuffix field value if set, zero value otherwise.
func (o *NetworkTemplate) GetDnsSuffix() []string {
	if o == nil || IsNil(o.DnsSuffix) {
		var ret []string
		return ret
	}
	return o.DnsSuffix
}

// GetDnsSuffixOk returns a tuple with the DnsSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetDnsSuffixOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsSuffix) {
		return nil, false
	}
	return o.DnsSuffix, true
}

// HasDnsSuffix returns a boolean if a field has been set.
func (o *NetworkTemplate) HasDnsSuffix() bool {
	if o != nil && !IsNil(o.DnsSuffix) {
		return true
	}

	return false
}

// SetDnsSuffix gets a reference to the given []string and assigns it to the DnsSuffix field.
func (o *NetworkTemplate) SetDnsSuffix(v []string) {
	o.DnsSuffix = v
}

// GetExtraRoutes returns the ExtraRoutes field value if set, zero value otherwise.
func (o *NetworkTemplate) GetExtraRoutes() map[string]ExtraRouteProperties {
	if o == nil || IsNil(o.ExtraRoutes) {
		var ret map[string]ExtraRouteProperties
		return ret
	}
	return *o.ExtraRoutes
}

// GetExtraRoutesOk returns a tuple with the ExtraRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetExtraRoutesOk() (*map[string]ExtraRouteProperties, bool) {
	if o == nil || IsNil(o.ExtraRoutes) {
		return nil, false
	}
	return o.ExtraRoutes, true
}

// HasExtraRoutes returns a boolean if a field has been set.
func (o *NetworkTemplate) HasExtraRoutes() bool {
	if o != nil && !IsNil(o.ExtraRoutes) {
		return true
	}

	return false
}

// SetExtraRoutes gets a reference to the given map[string]ExtraRouteProperties and assigns it to the ExtraRoutes field.
func (o *NetworkTemplate) SetExtraRoutes(v map[string]ExtraRouteProperties) {
	o.ExtraRoutes = &v
}

// GetExtraRoutes6 returns the ExtraRoutes6 field value if set, zero value otherwise.
func (o *NetworkTemplate) GetExtraRoutes6() map[string]ExtraRoute6Properties {
	if o == nil || IsNil(o.ExtraRoutes6) {
		var ret map[string]ExtraRoute6Properties
		return ret
	}
	return *o.ExtraRoutes6
}

// GetExtraRoutes6Ok returns a tuple with the ExtraRoutes6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetExtraRoutes6Ok() (*map[string]ExtraRoute6Properties, bool) {
	if o == nil || IsNil(o.ExtraRoutes6) {
		return nil, false
	}
	return o.ExtraRoutes6, true
}

// HasExtraRoutes6 returns a boolean if a field has been set.
func (o *NetworkTemplate) HasExtraRoutes6() bool {
	if o != nil && !IsNil(o.ExtraRoutes6) {
		return true
	}

	return false
}

// SetExtraRoutes6 gets a reference to the given map[string]ExtraRoute6Properties and assigns it to the ExtraRoutes6 field.
func (o *NetworkTemplate) SetExtraRoutes6(v map[string]ExtraRoute6Properties) {
	o.ExtraRoutes6 = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkTemplate) SetId(v string) {
	o.Id = &v
}

// GetImportOrgNetworks returns the ImportOrgNetworks field value if set, zero value otherwise.
func (o *NetworkTemplate) GetImportOrgNetworks() []string {
	if o == nil || IsNil(o.ImportOrgNetworks) {
		var ret []string
		return ret
	}
	return o.ImportOrgNetworks
}

// GetImportOrgNetworksOk returns a tuple with the ImportOrgNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetImportOrgNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.ImportOrgNetworks) {
		return nil, false
	}
	return o.ImportOrgNetworks, true
}

// HasImportOrgNetworks returns a boolean if a field has been set.
func (o *NetworkTemplate) HasImportOrgNetworks() bool {
	if o != nil && !IsNil(o.ImportOrgNetworks) {
		return true
	}

	return false
}

// SetImportOrgNetworks gets a reference to the given []string and assigns it to the ImportOrgNetworks field.
func (o *NetworkTemplate) SetImportOrgNetworks(v []string) {
	o.ImportOrgNetworks = v
}

// GetMistNac returns the MistNac field value if set, zero value otherwise.
func (o *NetworkTemplate) GetMistNac() SwitchMistNac {
	if o == nil || IsNil(o.MistNac) {
		var ret SwitchMistNac
		return ret
	}
	return *o.MistNac
}

// GetMistNacOk returns a tuple with the MistNac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetMistNacOk() (*SwitchMistNac, bool) {
	if o == nil || IsNil(o.MistNac) {
		return nil, false
	}
	return o.MistNac, true
}

// HasMistNac returns a boolean if a field has been set.
func (o *NetworkTemplate) HasMistNac() bool {
	if o != nil && !IsNil(o.MistNac) {
		return true
	}

	return false
}

// SetMistNac gets a reference to the given SwitchMistNac and assigns it to the MistNac field.
func (o *NetworkTemplate) SetMistNac(v SwitchMistNac) {
	o.MistNac = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *NetworkTemplate) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *NetworkTemplate) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *NetworkTemplate) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkTemplate) SetName(v string) {
	o.Name = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *NetworkTemplate) GetNetworks() map[string]SwitchNetwork {
	if o == nil || IsNil(o.Networks) {
		var ret map[string]SwitchNetwork
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetNetworksOk() (*map[string]SwitchNetwork, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *NetworkTemplate) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given map[string]SwitchNetwork and assigns it to the Networks field.
func (o *NetworkTemplate) SetNetworks(v map[string]SwitchNetwork) {
	o.Networks = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *NetworkTemplate) GetNtpServers() []string {
	if o == nil || IsNil(o.NtpServers) {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetNtpServersOk() ([]string, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *NetworkTemplate) HasNtpServers() bool {
	if o != nil && !IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *NetworkTemplate) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *NetworkTemplate) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *NetworkTemplate) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *NetworkTemplate) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPortMirrorings returns the PortMirrorings field value if set, zero value otherwise.
func (o *NetworkTemplate) GetPortMirrorings() map[string]SwitchPortMirroring {
	if o == nil || IsNil(o.PortMirrorings) {
		var ret map[string]SwitchPortMirroring
		return ret
	}
	return *o.PortMirrorings
}

// GetPortMirroringsOk returns a tuple with the PortMirrorings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetPortMirroringsOk() (*map[string]SwitchPortMirroring, bool) {
	if o == nil || IsNil(o.PortMirrorings) {
		return nil, false
	}
	return o.PortMirrorings, true
}

// HasPortMirrorings returns a boolean if a field has been set.
func (o *NetworkTemplate) HasPortMirrorings() bool {
	if o != nil && !IsNil(o.PortMirrorings) {
		return true
	}

	return false
}

// SetPortMirrorings gets a reference to the given map[string]SwitchPortMirroring and assigns it to the PortMirrorings field.
func (o *NetworkTemplate) SetPortMirrorings(v map[string]SwitchPortMirroring) {
	o.PortMirrorings = &v
}

// GetPortUsages returns the PortUsages field value if set, zero value otherwise.
func (o *NetworkTemplate) GetPortUsages() map[string]SwitchPortUsage {
	if o == nil || IsNil(o.PortUsages) {
		var ret map[string]SwitchPortUsage
		return ret
	}
	return *o.PortUsages
}

// GetPortUsagesOk returns a tuple with the PortUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetPortUsagesOk() (*map[string]SwitchPortUsage, bool) {
	if o == nil || IsNil(o.PortUsages) {
		return nil, false
	}
	return o.PortUsages, true
}

// HasPortUsages returns a boolean if a field has been set.
func (o *NetworkTemplate) HasPortUsages() bool {
	if o != nil && !IsNil(o.PortUsages) {
		return true
	}

	return false
}

// SetPortUsages gets a reference to the given map[string]SwitchPortUsage and assigns it to the PortUsages field.
func (o *NetworkTemplate) SetPortUsages(v map[string]SwitchPortUsage) {
	o.PortUsages = &v
}

// GetRadiusConfig returns the RadiusConfig field value if set, zero value otherwise.
func (o *NetworkTemplate) GetRadiusConfig() RadiusConfig {
	if o == nil || IsNil(o.RadiusConfig) {
		var ret RadiusConfig
		return ret
	}
	return *o.RadiusConfig
}

// GetRadiusConfigOk returns a tuple with the RadiusConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetRadiusConfigOk() (*RadiusConfig, bool) {
	if o == nil || IsNil(o.RadiusConfig) {
		return nil, false
	}
	return o.RadiusConfig, true
}

// HasRadiusConfig returns a boolean if a field has been set.
func (o *NetworkTemplate) HasRadiusConfig() bool {
	if o != nil && !IsNil(o.RadiusConfig) {
		return true
	}

	return false
}

// SetRadiusConfig gets a reference to the given RadiusConfig and assigns it to the RadiusConfig field.
func (o *NetworkTemplate) SetRadiusConfig(v RadiusConfig) {
	o.RadiusConfig = &v
}

// GetRemoteSyslog returns the RemoteSyslog field value if set, zero value otherwise.
func (o *NetworkTemplate) GetRemoteSyslog() RemoteSyslog {
	if o == nil || IsNil(o.RemoteSyslog) {
		var ret RemoteSyslog
		return ret
	}
	return *o.RemoteSyslog
}

// GetRemoteSyslogOk returns a tuple with the RemoteSyslog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetRemoteSyslogOk() (*RemoteSyslog, bool) {
	if o == nil || IsNil(o.RemoteSyslog) {
		return nil, false
	}
	return o.RemoteSyslog, true
}

// HasRemoteSyslog returns a boolean if a field has been set.
func (o *NetworkTemplate) HasRemoteSyslog() bool {
	if o != nil && !IsNil(o.RemoteSyslog) {
		return true
	}

	return false
}

// SetRemoteSyslog gets a reference to the given RemoteSyslog and assigns it to the RemoteSyslog field.
func (o *NetworkTemplate) SetRemoteSyslog(v RemoteSyslog) {
	o.RemoteSyslog = &v
}

// GetSnmpConfig returns the SnmpConfig field value if set, zero value otherwise.
func (o *NetworkTemplate) GetSnmpConfig() SnmpConfig {
	if o == nil || IsNil(o.SnmpConfig) {
		var ret SnmpConfig
		return ret
	}
	return *o.SnmpConfig
}

// GetSnmpConfigOk returns a tuple with the SnmpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetSnmpConfigOk() (*SnmpConfig, bool) {
	if o == nil || IsNil(o.SnmpConfig) {
		return nil, false
	}
	return o.SnmpConfig, true
}

// HasSnmpConfig returns a boolean if a field has been set.
func (o *NetworkTemplate) HasSnmpConfig() bool {
	if o != nil && !IsNil(o.SnmpConfig) {
		return true
	}

	return false
}

// SetSnmpConfig gets a reference to the given SnmpConfig and assigns it to the SnmpConfig field.
func (o *NetworkTemplate) SetSnmpConfig(v SnmpConfig) {
	o.SnmpConfig = &v
}

// GetSwitchMatching returns the SwitchMatching field value if set, zero value otherwise.
func (o *NetworkTemplate) GetSwitchMatching() SwitchMatching {
	if o == nil || IsNil(o.SwitchMatching) {
		var ret SwitchMatching
		return ret
	}
	return *o.SwitchMatching
}

// GetSwitchMatchingOk returns a tuple with the SwitchMatching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetSwitchMatchingOk() (*SwitchMatching, bool) {
	if o == nil || IsNil(o.SwitchMatching) {
		return nil, false
	}
	return o.SwitchMatching, true
}

// HasSwitchMatching returns a boolean if a field has been set.
func (o *NetworkTemplate) HasSwitchMatching() bool {
	if o != nil && !IsNil(o.SwitchMatching) {
		return true
	}

	return false
}

// SetSwitchMatching gets a reference to the given SwitchMatching and assigns it to the SwitchMatching field.
func (o *NetworkTemplate) SetSwitchMatching(v SwitchMatching) {
	o.SwitchMatching = &v
}

// GetSwitchMgmt returns the SwitchMgmt field value if set, zero value otherwise.
func (o *NetworkTemplate) GetSwitchMgmt() SwitchMgmt {
	if o == nil || IsNil(o.SwitchMgmt) {
		var ret SwitchMgmt
		return ret
	}
	return *o.SwitchMgmt
}

// GetSwitchMgmtOk returns a tuple with the SwitchMgmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetSwitchMgmtOk() (*SwitchMgmt, bool) {
	if o == nil || IsNil(o.SwitchMgmt) {
		return nil, false
	}
	return o.SwitchMgmt, true
}

// HasSwitchMgmt returns a boolean if a field has been set.
func (o *NetworkTemplate) HasSwitchMgmt() bool {
	if o != nil && !IsNil(o.SwitchMgmt) {
		return true
	}

	return false
}

// SetSwitchMgmt gets a reference to the given SwitchMgmt and assigns it to the SwitchMgmt field.
func (o *NetworkTemplate) SetSwitchMgmt(v SwitchMgmt) {
	o.SwitchMgmt = &v
}

// GetVrfConfig returns the VrfConfig field value if set, zero value otherwise.
func (o *NetworkTemplate) GetVrfConfig() VrfConfig {
	if o == nil || IsNil(o.VrfConfig) {
		var ret VrfConfig
		return ret
	}
	return *o.VrfConfig
}

// GetVrfConfigOk returns a tuple with the VrfConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetVrfConfigOk() (*VrfConfig, bool) {
	if o == nil || IsNil(o.VrfConfig) {
		return nil, false
	}
	return o.VrfConfig, true
}

// HasVrfConfig returns a boolean if a field has been set.
func (o *NetworkTemplate) HasVrfConfig() bool {
	if o != nil && !IsNil(o.VrfConfig) {
		return true
	}

	return false
}

// SetVrfConfig gets a reference to the given VrfConfig and assigns it to the VrfConfig field.
func (o *NetworkTemplate) SetVrfConfig(v VrfConfig) {
	o.VrfConfig = &v
}

// GetVrfInstances returns the VrfInstances field value if set, zero value otherwise.
func (o *NetworkTemplate) GetVrfInstances() map[string]VrfInstance {
	if o == nil || IsNil(o.VrfInstances) {
		var ret map[string]VrfInstance
		return ret
	}
	return *o.VrfInstances
}

// GetVrfInstancesOk returns a tuple with the VrfInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTemplate) GetVrfInstancesOk() (*map[string]VrfInstance, bool) {
	if o == nil || IsNil(o.VrfInstances) {
		return nil, false
	}
	return o.VrfInstances, true
}

// HasVrfInstances returns a boolean if a field has been set.
func (o *NetworkTemplate) HasVrfInstances() bool {
	if o != nil && !IsNil(o.VrfInstances) {
		return true
	}

	return false
}

// SetVrfInstances gets a reference to the given map[string]VrfInstance and assigns it to the VrfInstances field.
func (o *NetworkTemplate) SetVrfInstances(v map[string]VrfInstance) {
	o.VrfInstances = &v
}

func (o NetworkTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AclPolicies) {
		toSerialize["acl_policies"] = o.AclPolicies
	}
	if !IsNil(o.AclTags) {
		toSerialize["acl_tags"] = o.AclTags
	}
	if !IsNil(o.AdditionalConfigCmds) {
		toSerialize["additional_config_cmds"] = o.AdditionalConfigCmds
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.DhcpSnooping) {
		toSerialize["dhcp_snooping"] = o.DhcpSnooping
	}
	if !IsNil(o.DnsServers) {
		toSerialize["dns_servers"] = o.DnsServers
	}
	if !IsNil(o.DnsSuffix) {
		toSerialize["dns_suffix"] = o.DnsSuffix
	}
	if !IsNil(o.ExtraRoutes) {
		toSerialize["extra_routes"] = o.ExtraRoutes
	}
	if !IsNil(o.ExtraRoutes6) {
		toSerialize["extra_routes6"] = o.ExtraRoutes6
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImportOrgNetworks) {
		toSerialize["import_org_networks"] = o.ImportOrgNetworks
	}
	if !IsNil(o.MistNac) {
		toSerialize["mist_nac"] = o.MistNac
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.NtpServers) {
		toSerialize["ntp_servers"] = o.NtpServers
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.PortMirrorings) {
		toSerialize["port_mirrorings"] = o.PortMirrorings
	}
	if !IsNil(o.PortUsages) {
		toSerialize["port_usages"] = o.PortUsages
	}
	if !IsNil(o.RadiusConfig) {
		toSerialize["radius_config"] = o.RadiusConfig
	}
	if !IsNil(o.RemoteSyslog) {
		toSerialize["remote_syslog"] = o.RemoteSyslog
	}
	if !IsNil(o.SnmpConfig) {
		toSerialize["snmp_config"] = o.SnmpConfig
	}
	if !IsNil(o.SwitchMatching) {
		toSerialize["switch_matching"] = o.SwitchMatching
	}
	if !IsNil(o.SwitchMgmt) {
		toSerialize["switch_mgmt"] = o.SwitchMgmt
	}
	if !IsNil(o.VrfConfig) {
		toSerialize["vrf_config"] = o.VrfConfig
	}
	if !IsNil(o.VrfInstances) {
		toSerialize["vrf_instances"] = o.VrfInstances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkTemplate) UnmarshalJSON(data []byte) (err error) {
	varNetworkTemplate := _NetworkTemplate{}

	err = json.Unmarshal(data, &varNetworkTemplate)

	if err != nil {
		return err
	}

	*o = NetworkTemplate(varNetworkTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acl_policies")
		delete(additionalProperties, "acl_tags")
		delete(additionalProperties, "additional_config_cmds")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "dhcp_snooping")
		delete(additionalProperties, "dns_servers")
		delete(additionalProperties, "dns_suffix")
		delete(additionalProperties, "extra_routes")
		delete(additionalProperties, "extra_routes6")
		delete(additionalProperties, "id")
		delete(additionalProperties, "import_org_networks")
		delete(additionalProperties, "mist_nac")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "networks")
		delete(additionalProperties, "ntp_servers")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "port_mirrorings")
		delete(additionalProperties, "port_usages")
		delete(additionalProperties, "radius_config")
		delete(additionalProperties, "remote_syslog")
		delete(additionalProperties, "snmp_config")
		delete(additionalProperties, "switch_matching")
		delete(additionalProperties, "switch_mgmt")
		delete(additionalProperties, "vrf_config")
		delete(additionalProperties, "vrf_instances")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkTemplate struct {
	value *NetworkTemplate
	isSet bool
}

func (v NullableNetworkTemplate) Get() *NetworkTemplate {
	return v.value
}

func (v *NullableNetworkTemplate) Set(val *NetworkTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkTemplate(val *NetworkTemplate) *NullableNetworkTemplate {
	return &NullableNetworkTemplate{value: val, isSet: true}
}

func (v NullableNetworkTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


