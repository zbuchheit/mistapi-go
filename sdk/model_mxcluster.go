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

// checks if the Mxcluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mxcluster{}

// Mxcluster mxCluster
type Mxcluster struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Id *string `json:"id,omitempty"`
	MistDas *MxedgeDas `json:"mist_das,omitempty"`
	MistNac *MxclusterNac `json:"mist_nac,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	MxedgeMgmt *MxedgeMgmt `json:"mxedge_mgmt,omitempty"`
	Name *string `json:"name,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	Proxy *Proxy `json:"proxy,omitempty"`
	Radsec *MxclusterRadsec `json:"radsec,omitempty"`
	RadsecTls *MxclusterRadsecTls `json:"radsec_tls,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	// list of subnets where we allow AP to establish Mist Tunnels from
	TuntermApSubnets []string `json:"tunterm_ap_subnets,omitempty"`
	TuntermDhcpdConfig *TuntermDhcpdConfig `json:"tunterm_dhcpd_config,omitempty"`
	// extra routes for Mist Tunneled VLANs. Property key is a CIDR
	TuntermExtraRoutes *map[string]MxclusterTuntermExtraRoute `json:"tunterm_extra_routes,omitempty"`
	// hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP)
	TuntermHosts []string `json:"tunterm_hosts,omitempty"`
	// list of index of tunterm_hosts
	TuntermHostsOrder []int32 `json:"tunterm_hosts_order,omitempty"`
	TuntermHostsSelection *MxclusterTuntermHostsSelection `json:"tunterm_hosts_selection,omitempty"`
	TuntermMonitoring [][]TuntermMonitoringItem `json:"tunterm_monitoring,omitempty"`
	TuntermMonitoringDisabled *bool `json:"tunterm_monitoring_disabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Mxcluster Mxcluster

// NewMxcluster instantiates a new Mxcluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxcluster() *Mxcluster {
	this := Mxcluster{}
	var tuntermHostsSelection MxclusterTuntermHostsSelection = MXCLUSTERTUNTERMHOSTSSELECTION_SHUFFLE
	this.TuntermHostsSelection = &tuntermHostsSelection
	return &this
}

// NewMxclusterWithDefaults instantiates a new Mxcluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxclusterWithDefaults() *Mxcluster {
	this := Mxcluster{}
	var tuntermHostsSelection MxclusterTuntermHostsSelection = MXCLUSTERTUNTERMHOSTSSELECTION_SHUFFLE
	this.TuntermHostsSelection = &tuntermHostsSelection
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Mxcluster) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Mxcluster) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Mxcluster) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *Mxcluster) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *Mxcluster) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *Mxcluster) SetForSite(v bool) {
	o.ForSite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Mxcluster) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Mxcluster) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Mxcluster) SetId(v string) {
	o.Id = &v
}

// GetMistDas returns the MistDas field value if set, zero value otherwise.
func (o *Mxcluster) GetMistDas() MxedgeDas {
	if o == nil || IsNil(o.MistDas) {
		var ret MxedgeDas
		return ret
	}
	return *o.MistDas
}

// GetMistDasOk returns a tuple with the MistDas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetMistDasOk() (*MxedgeDas, bool) {
	if o == nil || IsNil(o.MistDas) {
		return nil, false
	}
	return o.MistDas, true
}

// HasMistDas returns a boolean if a field has been set.
func (o *Mxcluster) HasMistDas() bool {
	if o != nil && !IsNil(o.MistDas) {
		return true
	}

	return false
}

// SetMistDas gets a reference to the given MxedgeDas and assigns it to the MistDas field.
func (o *Mxcluster) SetMistDas(v MxedgeDas) {
	o.MistDas = &v
}

// GetMistNac returns the MistNac field value if set, zero value otherwise.
func (o *Mxcluster) GetMistNac() MxclusterNac {
	if o == nil || IsNil(o.MistNac) {
		var ret MxclusterNac
		return ret
	}
	return *o.MistNac
}

// GetMistNacOk returns a tuple with the MistNac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetMistNacOk() (*MxclusterNac, bool) {
	if o == nil || IsNil(o.MistNac) {
		return nil, false
	}
	return o.MistNac, true
}

// HasMistNac returns a boolean if a field has been set.
func (o *Mxcluster) HasMistNac() bool {
	if o != nil && !IsNil(o.MistNac) {
		return true
	}

	return false
}

// SetMistNac gets a reference to the given MxclusterNac and assigns it to the MistNac field.
func (o *Mxcluster) SetMistNac(v MxclusterNac) {
	o.MistNac = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Mxcluster) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Mxcluster) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Mxcluster) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetMxedgeMgmt returns the MxedgeMgmt field value if set, zero value otherwise.
func (o *Mxcluster) GetMxedgeMgmt() MxedgeMgmt {
	if o == nil || IsNil(o.MxedgeMgmt) {
		var ret MxedgeMgmt
		return ret
	}
	return *o.MxedgeMgmt
}

// GetMxedgeMgmtOk returns a tuple with the MxedgeMgmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetMxedgeMgmtOk() (*MxedgeMgmt, bool) {
	if o == nil || IsNil(o.MxedgeMgmt) {
		return nil, false
	}
	return o.MxedgeMgmt, true
}

// HasMxedgeMgmt returns a boolean if a field has been set.
func (o *Mxcluster) HasMxedgeMgmt() bool {
	if o != nil && !IsNil(o.MxedgeMgmt) {
		return true
	}

	return false
}

// SetMxedgeMgmt gets a reference to the given MxedgeMgmt and assigns it to the MxedgeMgmt field.
func (o *Mxcluster) SetMxedgeMgmt(v MxedgeMgmt) {
	o.MxedgeMgmt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Mxcluster) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Mxcluster) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Mxcluster) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Mxcluster) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Mxcluster) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Mxcluster) SetOrgId(v string) {
	o.OrgId = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *Mxcluster) GetProxy() Proxy {
	if o == nil || IsNil(o.Proxy) {
		var ret Proxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetProxyOk() (*Proxy, bool) {
	if o == nil || IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *Mxcluster) HasProxy() bool {
	if o != nil && !IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given Proxy and assigns it to the Proxy field.
func (o *Mxcluster) SetProxy(v Proxy) {
	o.Proxy = &v
}

// GetRadsec returns the Radsec field value if set, zero value otherwise.
func (o *Mxcluster) GetRadsec() MxclusterRadsec {
	if o == nil || IsNil(o.Radsec) {
		var ret MxclusterRadsec
		return ret
	}
	return *o.Radsec
}

// GetRadsecOk returns a tuple with the Radsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetRadsecOk() (*MxclusterRadsec, bool) {
	if o == nil || IsNil(o.Radsec) {
		return nil, false
	}
	return o.Radsec, true
}

// HasRadsec returns a boolean if a field has been set.
func (o *Mxcluster) HasRadsec() bool {
	if o != nil && !IsNil(o.Radsec) {
		return true
	}

	return false
}

// SetRadsec gets a reference to the given MxclusterRadsec and assigns it to the Radsec field.
func (o *Mxcluster) SetRadsec(v MxclusterRadsec) {
	o.Radsec = &v
}

// GetRadsecTls returns the RadsecTls field value if set, zero value otherwise.
func (o *Mxcluster) GetRadsecTls() MxclusterRadsecTls {
	if o == nil || IsNil(o.RadsecTls) {
		var ret MxclusterRadsecTls
		return ret
	}
	return *o.RadsecTls
}

// GetRadsecTlsOk returns a tuple with the RadsecTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetRadsecTlsOk() (*MxclusterRadsecTls, bool) {
	if o == nil || IsNil(o.RadsecTls) {
		return nil, false
	}
	return o.RadsecTls, true
}

// HasRadsecTls returns a boolean if a field has been set.
func (o *Mxcluster) HasRadsecTls() bool {
	if o != nil && !IsNil(o.RadsecTls) {
		return true
	}

	return false
}

// SetRadsecTls gets a reference to the given MxclusterRadsecTls and assigns it to the RadsecTls field.
func (o *Mxcluster) SetRadsecTls(v MxclusterRadsecTls) {
	o.RadsecTls = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *Mxcluster) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *Mxcluster) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *Mxcluster) SetSiteId(v string) {
	o.SiteId = &v
}

// GetTuntermApSubnets returns the TuntermApSubnets field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermApSubnets() []string {
	if o == nil || IsNil(o.TuntermApSubnets) {
		var ret []string
		return ret
	}
	return o.TuntermApSubnets
}

// GetTuntermApSubnetsOk returns a tuple with the TuntermApSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermApSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.TuntermApSubnets) {
		return nil, false
	}
	return o.TuntermApSubnets, true
}

// HasTuntermApSubnets returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermApSubnets() bool {
	if o != nil && !IsNil(o.TuntermApSubnets) {
		return true
	}

	return false
}

// SetTuntermApSubnets gets a reference to the given []string and assigns it to the TuntermApSubnets field.
func (o *Mxcluster) SetTuntermApSubnets(v []string) {
	o.TuntermApSubnets = v
}

// GetTuntermDhcpdConfig returns the TuntermDhcpdConfig field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermDhcpdConfig() TuntermDhcpdConfig {
	if o == nil || IsNil(o.TuntermDhcpdConfig) {
		var ret TuntermDhcpdConfig
		return ret
	}
	return *o.TuntermDhcpdConfig
}

// GetTuntermDhcpdConfigOk returns a tuple with the TuntermDhcpdConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermDhcpdConfigOk() (*TuntermDhcpdConfig, bool) {
	if o == nil || IsNil(o.TuntermDhcpdConfig) {
		return nil, false
	}
	return o.TuntermDhcpdConfig, true
}

// HasTuntermDhcpdConfig returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermDhcpdConfig() bool {
	if o != nil && !IsNil(o.TuntermDhcpdConfig) {
		return true
	}

	return false
}

// SetTuntermDhcpdConfig gets a reference to the given TuntermDhcpdConfig and assigns it to the TuntermDhcpdConfig field.
func (o *Mxcluster) SetTuntermDhcpdConfig(v TuntermDhcpdConfig) {
	o.TuntermDhcpdConfig = &v
}

// GetTuntermExtraRoutes returns the TuntermExtraRoutes field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermExtraRoutes() map[string]MxclusterTuntermExtraRoute {
	if o == nil || IsNil(o.TuntermExtraRoutes) {
		var ret map[string]MxclusterTuntermExtraRoute
		return ret
	}
	return *o.TuntermExtraRoutes
}

// GetTuntermExtraRoutesOk returns a tuple with the TuntermExtraRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermExtraRoutesOk() (*map[string]MxclusterTuntermExtraRoute, bool) {
	if o == nil || IsNil(o.TuntermExtraRoutes) {
		return nil, false
	}
	return o.TuntermExtraRoutes, true
}

// HasTuntermExtraRoutes returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermExtraRoutes() bool {
	if o != nil && !IsNil(o.TuntermExtraRoutes) {
		return true
	}

	return false
}

// SetTuntermExtraRoutes gets a reference to the given map[string]MxclusterTuntermExtraRoute and assigns it to the TuntermExtraRoutes field.
func (o *Mxcluster) SetTuntermExtraRoutes(v map[string]MxclusterTuntermExtraRoute) {
	o.TuntermExtraRoutes = &v
}

// GetTuntermHosts returns the TuntermHosts field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermHosts() []string {
	if o == nil || IsNil(o.TuntermHosts) {
		var ret []string
		return ret
	}
	return o.TuntermHosts
}

// GetTuntermHostsOk returns a tuple with the TuntermHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.TuntermHosts) {
		return nil, false
	}
	return o.TuntermHosts, true
}

// HasTuntermHosts returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermHosts() bool {
	if o != nil && !IsNil(o.TuntermHosts) {
		return true
	}

	return false
}

// SetTuntermHosts gets a reference to the given []string and assigns it to the TuntermHosts field.
func (o *Mxcluster) SetTuntermHosts(v []string) {
	o.TuntermHosts = v
}

// GetTuntermHostsOrder returns the TuntermHostsOrder field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermHostsOrder() []int32 {
	if o == nil || IsNil(o.TuntermHostsOrder) {
		var ret []int32
		return ret
	}
	return o.TuntermHostsOrder
}

// GetTuntermHostsOrderOk returns a tuple with the TuntermHostsOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermHostsOrderOk() ([]int32, bool) {
	if o == nil || IsNil(o.TuntermHostsOrder) {
		return nil, false
	}
	return o.TuntermHostsOrder, true
}

// HasTuntermHostsOrder returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermHostsOrder() bool {
	if o != nil && !IsNil(o.TuntermHostsOrder) {
		return true
	}

	return false
}

// SetTuntermHostsOrder gets a reference to the given []int32 and assigns it to the TuntermHostsOrder field.
func (o *Mxcluster) SetTuntermHostsOrder(v []int32) {
	o.TuntermHostsOrder = v
}

// GetTuntermHostsSelection returns the TuntermHostsSelection field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermHostsSelection() MxclusterTuntermHostsSelection {
	if o == nil || IsNil(o.TuntermHostsSelection) {
		var ret MxclusterTuntermHostsSelection
		return ret
	}
	return *o.TuntermHostsSelection
}

// GetTuntermHostsSelectionOk returns a tuple with the TuntermHostsSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermHostsSelectionOk() (*MxclusterTuntermHostsSelection, bool) {
	if o == nil || IsNil(o.TuntermHostsSelection) {
		return nil, false
	}
	return o.TuntermHostsSelection, true
}

// HasTuntermHostsSelection returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermHostsSelection() bool {
	if o != nil && !IsNil(o.TuntermHostsSelection) {
		return true
	}

	return false
}

// SetTuntermHostsSelection gets a reference to the given MxclusterTuntermHostsSelection and assigns it to the TuntermHostsSelection field.
func (o *Mxcluster) SetTuntermHostsSelection(v MxclusterTuntermHostsSelection) {
	o.TuntermHostsSelection = &v
}

// GetTuntermMonitoring returns the TuntermMonitoring field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermMonitoring() [][]TuntermMonitoringItem {
	if o == nil || IsNil(o.TuntermMonitoring) {
		var ret [][]TuntermMonitoringItem
		return ret
	}
	return o.TuntermMonitoring
}

// GetTuntermMonitoringOk returns a tuple with the TuntermMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermMonitoringOk() ([][]TuntermMonitoringItem, bool) {
	if o == nil || IsNil(o.TuntermMonitoring) {
		return nil, false
	}
	return o.TuntermMonitoring, true
}

// HasTuntermMonitoring returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermMonitoring() bool {
	if o != nil && !IsNil(o.TuntermMonitoring) {
		return true
	}

	return false
}

// SetTuntermMonitoring gets a reference to the given [][]TuntermMonitoringItem and assigns it to the TuntermMonitoring field.
func (o *Mxcluster) SetTuntermMonitoring(v [][]TuntermMonitoringItem) {
	o.TuntermMonitoring = v
}

// GetTuntermMonitoringDisabled returns the TuntermMonitoringDisabled field value if set, zero value otherwise.
func (o *Mxcluster) GetTuntermMonitoringDisabled() bool {
	if o == nil || IsNil(o.TuntermMonitoringDisabled) {
		var ret bool
		return ret
	}
	return *o.TuntermMonitoringDisabled
}

// GetTuntermMonitoringDisabledOk returns a tuple with the TuntermMonitoringDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mxcluster) GetTuntermMonitoringDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TuntermMonitoringDisabled) {
		return nil, false
	}
	return o.TuntermMonitoringDisabled, true
}

// HasTuntermMonitoringDisabled returns a boolean if a field has been set.
func (o *Mxcluster) HasTuntermMonitoringDisabled() bool {
	if o != nil && !IsNil(o.TuntermMonitoringDisabled) {
		return true
	}

	return false
}

// SetTuntermMonitoringDisabled gets a reference to the given bool and assigns it to the TuntermMonitoringDisabled field.
func (o *Mxcluster) SetTuntermMonitoringDisabled(v bool) {
	o.TuntermMonitoringDisabled = &v
}

func (o Mxcluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mxcluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MistDas) {
		toSerialize["mist_das"] = o.MistDas
	}
	if !IsNil(o.MistNac) {
		toSerialize["mist_nac"] = o.MistNac
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.MxedgeMgmt) {
		toSerialize["mxedge_mgmt"] = o.MxedgeMgmt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !IsNil(o.Radsec) {
		toSerialize["radsec"] = o.Radsec
	}
	if !IsNil(o.RadsecTls) {
		toSerialize["radsec_tls"] = o.RadsecTls
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.TuntermApSubnets) {
		toSerialize["tunterm_ap_subnets"] = o.TuntermApSubnets
	}
	if !IsNil(o.TuntermDhcpdConfig) {
		toSerialize["tunterm_dhcpd_config"] = o.TuntermDhcpdConfig
	}
	if !IsNil(o.TuntermExtraRoutes) {
		toSerialize["tunterm_extra_routes"] = o.TuntermExtraRoutes
	}
	if !IsNil(o.TuntermHosts) {
		toSerialize["tunterm_hosts"] = o.TuntermHosts
	}
	if !IsNil(o.TuntermHostsOrder) {
		toSerialize["tunterm_hosts_order"] = o.TuntermHostsOrder
	}
	if !IsNil(o.TuntermHostsSelection) {
		toSerialize["tunterm_hosts_selection"] = o.TuntermHostsSelection
	}
	if !IsNil(o.TuntermMonitoring) {
		toSerialize["tunterm_monitoring"] = o.TuntermMonitoring
	}
	if !IsNil(o.TuntermMonitoringDisabled) {
		toSerialize["tunterm_monitoring_disabled"] = o.TuntermMonitoringDisabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Mxcluster) UnmarshalJSON(data []byte) (err error) {
	varMxcluster := _Mxcluster{}

	err = json.Unmarshal(data, &varMxcluster)

	if err != nil {
		return err
	}

	*o = Mxcluster(varMxcluster)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mist_das")
		delete(additionalProperties, "mist_nac")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "mxedge_mgmt")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "radsec")
		delete(additionalProperties, "radsec_tls")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "tunterm_ap_subnets")
		delete(additionalProperties, "tunterm_dhcpd_config")
		delete(additionalProperties, "tunterm_extra_routes")
		delete(additionalProperties, "tunterm_hosts")
		delete(additionalProperties, "tunterm_hosts_order")
		delete(additionalProperties, "tunterm_hosts_selection")
		delete(additionalProperties, "tunterm_monitoring")
		delete(additionalProperties, "tunterm_monitoring_disabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxcluster struct {
	value *Mxcluster
	isSet bool
}

func (v NullableMxcluster) Get() *Mxcluster {
	return v.value
}

func (v *NullableMxcluster) Set(val *Mxcluster) {
	v.value = val
	v.isSet = true
}

func (v NullableMxcluster) IsSet() bool {
	return v.isSet
}

func (v *NullableMxcluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxcluster(val *Mxcluster) *NullableMxcluster {
	return &NullableMxcluster{value: val, isSet: true}
}

func (v NullableMxcluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxcluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


