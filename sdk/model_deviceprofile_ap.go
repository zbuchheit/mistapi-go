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

// checks if the DeviceprofileAp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceprofileAp{}

// DeviceprofileAp Device Profile
type DeviceprofileAp struct {
	Aeroscout *ApAeroscout `json:"aeroscout,omitempty"`
	BleConfig *BleConfig `json:"ble_config,omitempty"`
	CreatedTime *float32 `json:"created_time,omitempty"`
	// whether to disable eth1 port
	DisableEth1 *bool `json:"disable_eth1,omitempty"`
	// whether to disable eth2 port
	DisableEth2 *bool `json:"disable_eth2,omitempty"`
	// whether to disable eth3 port
	DisableEth3 *bool `json:"disable_eth3,omitempty"`
	// whether to disable module port
	DisableModule *bool `json:"disable_module,omitempty"`
	EslConfig *ApEslConfig `json:"esl_config,omitempty"`
	ForSite *bool `json:"for_site,omitempty"`
	Id *string `json:"id,omitempty"`
	IotConfig *ApIot `json:"iot_config,omitempty"`
	IpConfig *ApIpConfig `json:"ip_config,omitempty"`
	Led *ApLed `json:"led,omitempty"`
	Mesh *ApMesh `json:"mesh,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	Name NullableString `json:"name,omitempty"`
	NtpServers []string `json:"ntp_servers,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	// whether to enable power out through module port (for APH) or eth1 (for APL/BT11)
	PoePassthrough *bool `json:"poe_passthrough,omitempty"`
	// Property key is the interface(s) name (e.g. \"eth1,eth2\")
	PortConfig *map[string]ApPortConfig `json:"port_config,omitempty"`
	PwrConfig *ApPwrConfig `json:"pwr_config,omitempty"`
	RadioConfig *ApRadio `json:"radio_config,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	// Deprecated
	SwitchConfig *ApSwitch `json:"switch_config,omitempty"`
	Type *DeviceType `json:"type,omitempty"`
	UplinkPortConfig *ApUplinkPortConfig `json:"uplink_port_config,omitempty"`
	UsbConfig *ApUsb `json:"usb_config,omitempty"`
	// a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
	Vars *map[string]string `json:"vars,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceprofileAp DeviceprofileAp

// NewDeviceprofileAp instantiates a new DeviceprofileAp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceprofileAp() *DeviceprofileAp {
	this := DeviceprofileAp{}
	var disableEth1 bool = false
	this.DisableEth1 = &disableEth1
	var disableEth2 bool = false
	this.DisableEth2 = &disableEth2
	var disableEth3 bool = false
	this.DisableEth3 = &disableEth3
	var disableModule bool = false
	this.DisableModule = &disableModule
	var poePassthrough bool = false
	this.PoePassthrough = &poePassthrough
	var type_ DeviceType = DEVICETYPE_AP
	this.Type = &type_
	return &this
}

// NewDeviceprofileApWithDefaults instantiates a new DeviceprofileAp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceprofileApWithDefaults() *DeviceprofileAp {
	this := DeviceprofileAp{}
	var disableEth1 bool = false
	this.DisableEth1 = &disableEth1
	var disableEth2 bool = false
	this.DisableEth2 = &disableEth2
	var disableEth3 bool = false
	this.DisableEth3 = &disableEth3
	var disableModule bool = false
	this.DisableModule = &disableModule
	var poePassthrough bool = false
	this.PoePassthrough = &poePassthrough
	var type_ DeviceType = DEVICETYPE_AP
	this.Type = &type_
	return &this
}

// GetAeroscout returns the Aeroscout field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetAeroscout() ApAeroscout {
	if o == nil || IsNil(o.Aeroscout) {
		var ret ApAeroscout
		return ret
	}
	return *o.Aeroscout
}

// GetAeroscoutOk returns a tuple with the Aeroscout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetAeroscoutOk() (*ApAeroscout, bool) {
	if o == nil || IsNil(o.Aeroscout) {
		return nil, false
	}
	return o.Aeroscout, true
}

// HasAeroscout returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasAeroscout() bool {
	if o != nil && !IsNil(o.Aeroscout) {
		return true
	}

	return false
}

// SetAeroscout gets a reference to the given ApAeroscout and assigns it to the Aeroscout field.
func (o *DeviceprofileAp) SetAeroscout(v ApAeroscout) {
	o.Aeroscout = &v
}

// GetBleConfig returns the BleConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetBleConfig() BleConfig {
	if o == nil || IsNil(o.BleConfig) {
		var ret BleConfig
		return ret
	}
	return *o.BleConfig
}

// GetBleConfigOk returns a tuple with the BleConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetBleConfigOk() (*BleConfig, bool) {
	if o == nil || IsNil(o.BleConfig) {
		return nil, false
	}
	return o.BleConfig, true
}

// HasBleConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasBleConfig() bool {
	if o != nil && !IsNil(o.BleConfig) {
		return true
	}

	return false
}

// SetBleConfig gets a reference to the given BleConfig and assigns it to the BleConfig field.
func (o *DeviceprofileAp) SetBleConfig(v BleConfig) {
	o.BleConfig = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *DeviceprofileAp) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetDisableEth1 returns the DisableEth1 field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetDisableEth1() bool {
	if o == nil || IsNil(o.DisableEth1) {
		var ret bool
		return ret
	}
	return *o.DisableEth1
}

// GetDisableEth1Ok returns a tuple with the DisableEth1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetDisableEth1Ok() (*bool, bool) {
	if o == nil || IsNil(o.DisableEth1) {
		return nil, false
	}
	return o.DisableEth1, true
}

// HasDisableEth1 returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasDisableEth1() bool {
	if o != nil && !IsNil(o.DisableEth1) {
		return true
	}

	return false
}

// SetDisableEth1 gets a reference to the given bool and assigns it to the DisableEth1 field.
func (o *DeviceprofileAp) SetDisableEth1(v bool) {
	o.DisableEth1 = &v
}

// GetDisableEth2 returns the DisableEth2 field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetDisableEth2() bool {
	if o == nil || IsNil(o.DisableEth2) {
		var ret bool
		return ret
	}
	return *o.DisableEth2
}

// GetDisableEth2Ok returns a tuple with the DisableEth2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetDisableEth2Ok() (*bool, bool) {
	if o == nil || IsNil(o.DisableEth2) {
		return nil, false
	}
	return o.DisableEth2, true
}

// HasDisableEth2 returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasDisableEth2() bool {
	if o != nil && !IsNil(o.DisableEth2) {
		return true
	}

	return false
}

// SetDisableEth2 gets a reference to the given bool and assigns it to the DisableEth2 field.
func (o *DeviceprofileAp) SetDisableEth2(v bool) {
	o.DisableEth2 = &v
}

// GetDisableEth3 returns the DisableEth3 field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetDisableEth3() bool {
	if o == nil || IsNil(o.DisableEth3) {
		var ret bool
		return ret
	}
	return *o.DisableEth3
}

// GetDisableEth3Ok returns a tuple with the DisableEth3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetDisableEth3Ok() (*bool, bool) {
	if o == nil || IsNil(o.DisableEth3) {
		return nil, false
	}
	return o.DisableEth3, true
}

// HasDisableEth3 returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasDisableEth3() bool {
	if o != nil && !IsNil(o.DisableEth3) {
		return true
	}

	return false
}

// SetDisableEth3 gets a reference to the given bool and assigns it to the DisableEth3 field.
func (o *DeviceprofileAp) SetDisableEth3(v bool) {
	o.DisableEth3 = &v
}

// GetDisableModule returns the DisableModule field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetDisableModule() bool {
	if o == nil || IsNil(o.DisableModule) {
		var ret bool
		return ret
	}
	return *o.DisableModule
}

// GetDisableModuleOk returns a tuple with the DisableModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetDisableModuleOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableModule) {
		return nil, false
	}
	return o.DisableModule, true
}

// HasDisableModule returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasDisableModule() bool {
	if o != nil && !IsNil(o.DisableModule) {
		return true
	}

	return false
}

// SetDisableModule gets a reference to the given bool and assigns it to the DisableModule field.
func (o *DeviceprofileAp) SetDisableModule(v bool) {
	o.DisableModule = &v
}

// GetEslConfig returns the EslConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetEslConfig() ApEslConfig {
	if o == nil || IsNil(o.EslConfig) {
		var ret ApEslConfig
		return ret
	}
	return *o.EslConfig
}

// GetEslConfigOk returns a tuple with the EslConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetEslConfigOk() (*ApEslConfig, bool) {
	if o == nil || IsNil(o.EslConfig) {
		return nil, false
	}
	return o.EslConfig, true
}

// HasEslConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasEslConfig() bool {
	if o != nil && !IsNil(o.EslConfig) {
		return true
	}

	return false
}

// SetEslConfig gets a reference to the given ApEslConfig and assigns it to the EslConfig field.
func (o *DeviceprofileAp) SetEslConfig(v ApEslConfig) {
	o.EslConfig = &v
}

// GetForSite returns the ForSite field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetForSite() bool {
	if o == nil || IsNil(o.ForSite) {
		var ret bool
		return ret
	}
	return *o.ForSite
}

// GetForSiteOk returns a tuple with the ForSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetForSiteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForSite) {
		return nil, false
	}
	return o.ForSite, true
}

// HasForSite returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasForSite() bool {
	if o != nil && !IsNil(o.ForSite) {
		return true
	}

	return false
}

// SetForSite gets a reference to the given bool and assigns it to the ForSite field.
func (o *DeviceprofileAp) SetForSite(v bool) {
	o.ForSite = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceprofileAp) SetId(v string) {
	o.Id = &v
}

// GetIotConfig returns the IotConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetIotConfig() ApIot {
	if o == nil || IsNil(o.IotConfig) {
		var ret ApIot
		return ret
	}
	return *o.IotConfig
}

// GetIotConfigOk returns a tuple with the IotConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetIotConfigOk() (*ApIot, bool) {
	if o == nil || IsNil(o.IotConfig) {
		return nil, false
	}
	return o.IotConfig, true
}

// HasIotConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasIotConfig() bool {
	if o != nil && !IsNil(o.IotConfig) {
		return true
	}

	return false
}

// SetIotConfig gets a reference to the given ApIot and assigns it to the IotConfig field.
func (o *DeviceprofileAp) SetIotConfig(v ApIot) {
	o.IotConfig = &v
}

// GetIpConfig returns the IpConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetIpConfig() ApIpConfig {
	if o == nil || IsNil(o.IpConfig) {
		var ret ApIpConfig
		return ret
	}
	return *o.IpConfig
}

// GetIpConfigOk returns a tuple with the IpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetIpConfigOk() (*ApIpConfig, bool) {
	if o == nil || IsNil(o.IpConfig) {
		return nil, false
	}
	return o.IpConfig, true
}

// HasIpConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasIpConfig() bool {
	if o != nil && !IsNil(o.IpConfig) {
		return true
	}

	return false
}

// SetIpConfig gets a reference to the given ApIpConfig and assigns it to the IpConfig field.
func (o *DeviceprofileAp) SetIpConfig(v ApIpConfig) {
	o.IpConfig = &v
}

// GetLed returns the Led field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetLed() ApLed {
	if o == nil || IsNil(o.Led) {
		var ret ApLed
		return ret
	}
	return *o.Led
}

// GetLedOk returns a tuple with the Led field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetLedOk() (*ApLed, bool) {
	if o == nil || IsNil(o.Led) {
		return nil, false
	}
	return o.Led, true
}

// HasLed returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasLed() bool {
	if o != nil && !IsNil(o.Led) {
		return true
	}

	return false
}

// SetLed gets a reference to the given ApLed and assigns it to the Led field.
func (o *DeviceprofileAp) SetLed(v ApLed) {
	o.Led = &v
}

// GetMesh returns the Mesh field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetMesh() ApMesh {
	if o == nil || IsNil(o.Mesh) {
		var ret ApMesh
		return ret
	}
	return *o.Mesh
}

// GetMeshOk returns a tuple with the Mesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetMeshOk() (*ApMesh, bool) {
	if o == nil || IsNil(o.Mesh) {
		return nil, false
	}
	return o.Mesh, true
}

// HasMesh returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasMesh() bool {
	if o != nil && !IsNil(o.Mesh) {
		return true
	}

	return false
}

// SetMesh gets a reference to the given ApMesh and assigns it to the Mesh field.
func (o *DeviceprofileAp) SetMesh(v ApMesh) {
	o.Mesh = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *DeviceprofileAp) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceprofileAp) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceprofileAp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeviceprofileAp) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DeviceprofileAp) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeviceprofileAp) UnsetName() {
	o.Name.Unset()
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetNtpServers() []string {
	if o == nil || IsNil(o.NtpServers) {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetNtpServersOk() ([]string, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasNtpServers() bool {
	if o != nil && !IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *DeviceprofileAp) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *DeviceprofileAp) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPoePassthrough returns the PoePassthrough field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetPoePassthrough() bool {
	if o == nil || IsNil(o.PoePassthrough) {
		var ret bool
		return ret
	}
	return *o.PoePassthrough
}

// GetPoePassthroughOk returns a tuple with the PoePassthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetPoePassthroughOk() (*bool, bool) {
	if o == nil || IsNil(o.PoePassthrough) {
		return nil, false
	}
	return o.PoePassthrough, true
}

// HasPoePassthrough returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasPoePassthrough() bool {
	if o != nil && !IsNil(o.PoePassthrough) {
		return true
	}

	return false
}

// SetPoePassthrough gets a reference to the given bool and assigns it to the PoePassthrough field.
func (o *DeviceprofileAp) SetPoePassthrough(v bool) {
	o.PoePassthrough = &v
}

// GetPortConfig returns the PortConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetPortConfig() map[string]ApPortConfig {
	if o == nil || IsNil(o.PortConfig) {
		var ret map[string]ApPortConfig
		return ret
	}
	return *o.PortConfig
}

// GetPortConfigOk returns a tuple with the PortConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetPortConfigOk() (*map[string]ApPortConfig, bool) {
	if o == nil || IsNil(o.PortConfig) {
		return nil, false
	}
	return o.PortConfig, true
}

// HasPortConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasPortConfig() bool {
	if o != nil && !IsNil(o.PortConfig) {
		return true
	}

	return false
}

// SetPortConfig gets a reference to the given map[string]ApPortConfig and assigns it to the PortConfig field.
func (o *DeviceprofileAp) SetPortConfig(v map[string]ApPortConfig) {
	o.PortConfig = &v
}

// GetPwrConfig returns the PwrConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetPwrConfig() ApPwrConfig {
	if o == nil || IsNil(o.PwrConfig) {
		var ret ApPwrConfig
		return ret
	}
	return *o.PwrConfig
}

// GetPwrConfigOk returns a tuple with the PwrConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetPwrConfigOk() (*ApPwrConfig, bool) {
	if o == nil || IsNil(o.PwrConfig) {
		return nil, false
	}
	return o.PwrConfig, true
}

// HasPwrConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasPwrConfig() bool {
	if o != nil && !IsNil(o.PwrConfig) {
		return true
	}

	return false
}

// SetPwrConfig gets a reference to the given ApPwrConfig and assigns it to the PwrConfig field.
func (o *DeviceprofileAp) SetPwrConfig(v ApPwrConfig) {
	o.PwrConfig = &v
}

// GetRadioConfig returns the RadioConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetRadioConfig() ApRadio {
	if o == nil || IsNil(o.RadioConfig) {
		var ret ApRadio
		return ret
	}
	return *o.RadioConfig
}

// GetRadioConfigOk returns a tuple with the RadioConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetRadioConfigOk() (*ApRadio, bool) {
	if o == nil || IsNil(o.RadioConfig) {
		return nil, false
	}
	return o.RadioConfig, true
}

// HasRadioConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasRadioConfig() bool {
	if o != nil && !IsNil(o.RadioConfig) {
		return true
	}

	return false
}

// SetRadioConfig gets a reference to the given ApRadio and assigns it to the RadioConfig field.
func (o *DeviceprofileAp) SetRadioConfig(v ApRadio) {
	o.RadioConfig = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DeviceprofileAp) SetSiteId(v string) {
	o.SiteId = &v
}

// GetSwitchConfig returns the SwitchConfig field value if set, zero value otherwise.
// Deprecated
func (o *DeviceprofileAp) GetSwitchConfig() ApSwitch {
	if o == nil || IsNil(o.SwitchConfig) {
		var ret ApSwitch
		return ret
	}
	return *o.SwitchConfig
}

// GetSwitchConfigOk returns a tuple with the SwitchConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceprofileAp) GetSwitchConfigOk() (*ApSwitch, bool) {
	if o == nil || IsNil(o.SwitchConfig) {
		return nil, false
	}
	return o.SwitchConfig, true
}

// HasSwitchConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasSwitchConfig() bool {
	if o != nil && !IsNil(o.SwitchConfig) {
		return true
	}

	return false
}

// SetSwitchConfig gets a reference to the given ApSwitch and assigns it to the SwitchConfig field.
// Deprecated
func (o *DeviceprofileAp) SetSwitchConfig(v ApSwitch) {
	o.SwitchConfig = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetType() DeviceType {
	if o == nil || IsNil(o.Type) {
		var ret DeviceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetTypeOk() (*DeviceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DeviceType and assigns it to the Type field.
func (o *DeviceprofileAp) SetType(v DeviceType) {
	o.Type = &v
}

// GetUplinkPortConfig returns the UplinkPortConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetUplinkPortConfig() ApUplinkPortConfig {
	if o == nil || IsNil(o.UplinkPortConfig) {
		var ret ApUplinkPortConfig
		return ret
	}
	return *o.UplinkPortConfig
}

// GetUplinkPortConfigOk returns a tuple with the UplinkPortConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetUplinkPortConfigOk() (*ApUplinkPortConfig, bool) {
	if o == nil || IsNil(o.UplinkPortConfig) {
		return nil, false
	}
	return o.UplinkPortConfig, true
}

// HasUplinkPortConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasUplinkPortConfig() bool {
	if o != nil && !IsNil(o.UplinkPortConfig) {
		return true
	}

	return false
}

// SetUplinkPortConfig gets a reference to the given ApUplinkPortConfig and assigns it to the UplinkPortConfig field.
func (o *DeviceprofileAp) SetUplinkPortConfig(v ApUplinkPortConfig) {
	o.UplinkPortConfig = &v
}

// GetUsbConfig returns the UsbConfig field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetUsbConfig() ApUsb {
	if o == nil || IsNil(o.UsbConfig) {
		var ret ApUsb
		return ret
	}
	return *o.UsbConfig
}

// GetUsbConfigOk returns a tuple with the UsbConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetUsbConfigOk() (*ApUsb, bool) {
	if o == nil || IsNil(o.UsbConfig) {
		return nil, false
	}
	return o.UsbConfig, true
}

// HasUsbConfig returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasUsbConfig() bool {
	if o != nil && !IsNil(o.UsbConfig) {
		return true
	}

	return false
}

// SetUsbConfig gets a reference to the given ApUsb and assigns it to the UsbConfig field.
func (o *DeviceprofileAp) SetUsbConfig(v ApUsb) {
	o.UsbConfig = &v
}

// GetVars returns the Vars field value if set, zero value otherwise.
func (o *DeviceprofileAp) GetVars() map[string]string {
	if o == nil || IsNil(o.Vars) {
		var ret map[string]string
		return ret
	}
	return *o.Vars
}

// GetVarsOk returns a tuple with the Vars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceprofileAp) GetVarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Vars) {
		return nil, false
	}
	return o.Vars, true
}

// HasVars returns a boolean if a field has been set.
func (o *DeviceprofileAp) HasVars() bool {
	if o != nil && !IsNil(o.Vars) {
		return true
	}

	return false
}

// SetVars gets a reference to the given map[string]string and assigns it to the Vars field.
func (o *DeviceprofileAp) SetVars(v map[string]string) {
	o.Vars = &v
}

func (o DeviceprofileAp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceprofileAp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aeroscout) {
		toSerialize["aeroscout"] = o.Aeroscout
	}
	if !IsNil(o.BleConfig) {
		toSerialize["ble_config"] = o.BleConfig
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.DisableEth1) {
		toSerialize["disable_eth1"] = o.DisableEth1
	}
	if !IsNil(o.DisableEth2) {
		toSerialize["disable_eth2"] = o.DisableEth2
	}
	if !IsNil(o.DisableEth3) {
		toSerialize["disable_eth3"] = o.DisableEth3
	}
	if !IsNil(o.DisableModule) {
		toSerialize["disable_module"] = o.DisableModule
	}
	if !IsNil(o.EslConfig) {
		toSerialize["esl_config"] = o.EslConfig
	}
	if !IsNil(o.ForSite) {
		toSerialize["for_site"] = o.ForSite
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IotConfig) {
		toSerialize["iot_config"] = o.IotConfig
	}
	if !IsNil(o.IpConfig) {
		toSerialize["ip_config"] = o.IpConfig
	}
	if !IsNil(o.Led) {
		toSerialize["led"] = o.Led
	}
	if !IsNil(o.Mesh) {
		toSerialize["mesh"] = o.Mesh
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.NtpServers) {
		toSerialize["ntp_servers"] = o.NtpServers
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.PoePassthrough) {
		toSerialize["poe_passthrough"] = o.PoePassthrough
	}
	if !IsNil(o.PortConfig) {
		toSerialize["port_config"] = o.PortConfig
	}
	if !IsNil(o.PwrConfig) {
		toSerialize["pwr_config"] = o.PwrConfig
	}
	if !IsNil(o.RadioConfig) {
		toSerialize["radio_config"] = o.RadioConfig
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.SwitchConfig) {
		toSerialize["switch_config"] = o.SwitchConfig
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UplinkPortConfig) {
		toSerialize["uplink_port_config"] = o.UplinkPortConfig
	}
	if !IsNil(o.UsbConfig) {
		toSerialize["usb_config"] = o.UsbConfig
	}
	if !IsNil(o.Vars) {
		toSerialize["vars"] = o.Vars
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceprofileAp) UnmarshalJSON(data []byte) (err error) {
	varDeviceprofileAp := _DeviceprofileAp{}

	err = json.Unmarshal(data, &varDeviceprofileAp)

	if err != nil {
		return err
	}

	*o = DeviceprofileAp(varDeviceprofileAp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aeroscout")
		delete(additionalProperties, "ble_config")
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "disable_eth1")
		delete(additionalProperties, "disable_eth2")
		delete(additionalProperties, "disable_eth3")
		delete(additionalProperties, "disable_module")
		delete(additionalProperties, "esl_config")
		delete(additionalProperties, "for_site")
		delete(additionalProperties, "id")
		delete(additionalProperties, "iot_config")
		delete(additionalProperties, "ip_config")
		delete(additionalProperties, "led")
		delete(additionalProperties, "mesh")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ntp_servers")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "poe_passthrough")
		delete(additionalProperties, "port_config")
		delete(additionalProperties, "pwr_config")
		delete(additionalProperties, "radio_config")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "switch_config")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uplink_port_config")
		delete(additionalProperties, "usb_config")
		delete(additionalProperties, "vars")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceprofileAp struct {
	value *DeviceprofileAp
	isSet bool
}

func (v NullableDeviceprofileAp) Get() *DeviceprofileAp {
	return v.value
}

func (v *NullableDeviceprofileAp) Set(val *DeviceprofileAp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceprofileAp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceprofileAp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceprofileAp(val *DeviceprofileAp) *NullableDeviceprofileAp {
	return &NullableDeviceprofileAp{value: val, isSet: true}
}

func (v NullableDeviceprofileAp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceprofileAp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


