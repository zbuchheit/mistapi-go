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

// checks if the InstallerDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallerDevice{}

// InstallerDevice struct for InstallerDevice
type InstallerDevice struct {
	Connected *bool `json:"connected,omitempty"`
	DeviceprofileName *string `json:"deviceprofile_name,omitempty"`
	ExtIp *string `json:"ext_ip,omitempty"`
	Height *float32 `json:"height,omitempty"`
	Ip *string `json:"ip,omitempty"`
	LastSeen *float32 `json:"last_seen,omitempty"`
	Mac *string `json:"mac,omitempty"`
	MapId *string `json:"map_id,omitempty"`
	Model *string `json:"model,omitempty"`
	Name *string `json:"name,omitempty"`
	Orientation *float32 `json:"orientation,omitempty"`
	Serial *string `json:"serial,omitempty"`
	SiteName *string `json:"site_name,omitempty"`
	Uptime *int32 `json:"uptime,omitempty"`
	VcMac NullableString `json:"vc_mac,omitempty"`
	Version *string `json:"version,omitempty"`
	X *float32 `json:"x,omitempty"`
	Y *float32 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InstallerDevice InstallerDevice

// NewInstallerDevice instantiates a new InstallerDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallerDevice() *InstallerDevice {
	this := InstallerDevice{}
	return &this
}

// NewInstallerDeviceWithDefaults instantiates a new InstallerDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallerDeviceWithDefaults() *InstallerDevice {
	this := InstallerDevice{}
	return &this
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *InstallerDevice) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *InstallerDevice) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *InstallerDevice) SetConnected(v bool) {
	o.Connected = &v
}

// GetDeviceprofileName returns the DeviceprofileName field value if set, zero value otherwise.
func (o *InstallerDevice) GetDeviceprofileName() string {
	if o == nil || IsNil(o.DeviceprofileName) {
		var ret string
		return ret
	}
	return *o.DeviceprofileName
}

// GetDeviceprofileNameOk returns a tuple with the DeviceprofileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetDeviceprofileNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceprofileName) {
		return nil, false
	}
	return o.DeviceprofileName, true
}

// HasDeviceprofileName returns a boolean if a field has been set.
func (o *InstallerDevice) HasDeviceprofileName() bool {
	if o != nil && !IsNil(o.DeviceprofileName) {
		return true
	}

	return false
}

// SetDeviceprofileName gets a reference to the given string and assigns it to the DeviceprofileName field.
func (o *InstallerDevice) SetDeviceprofileName(v string) {
	o.DeviceprofileName = &v
}

// GetExtIp returns the ExtIp field value if set, zero value otherwise.
func (o *InstallerDevice) GetExtIp() string {
	if o == nil || IsNil(o.ExtIp) {
		var ret string
		return ret
	}
	return *o.ExtIp
}

// GetExtIpOk returns a tuple with the ExtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetExtIpOk() (*string, bool) {
	if o == nil || IsNil(o.ExtIp) {
		return nil, false
	}
	return o.ExtIp, true
}

// HasExtIp returns a boolean if a field has been set.
func (o *InstallerDevice) HasExtIp() bool {
	if o != nil && !IsNil(o.ExtIp) {
		return true
	}

	return false
}

// SetExtIp gets a reference to the given string and assigns it to the ExtIp field.
func (o *InstallerDevice) SetExtIp(v string) {
	o.ExtIp = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *InstallerDevice) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *InstallerDevice) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *InstallerDevice) SetHeight(v float32) {
	o.Height = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *InstallerDevice) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *InstallerDevice) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *InstallerDevice) SetIp(v string) {
	o.Ip = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *InstallerDevice) GetLastSeen() float32 {
	if o == nil || IsNil(o.LastSeen) {
		var ret float32
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetLastSeenOk() (*float32, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *InstallerDevice) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given float32 and assigns it to the LastSeen field.
func (o *InstallerDevice) SetLastSeen(v float32) {
	o.LastSeen = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InstallerDevice) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InstallerDevice) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InstallerDevice) SetMac(v string) {
	o.Mac = &v
}

// GetMapId returns the MapId field value if set, zero value otherwise.
func (o *InstallerDevice) GetMapId() string {
	if o == nil || IsNil(o.MapId) {
		var ret string
		return ret
	}
	return *o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetMapIdOk() (*string, bool) {
	if o == nil || IsNil(o.MapId) {
		return nil, false
	}
	return o.MapId, true
}

// HasMapId returns a boolean if a field has been set.
func (o *InstallerDevice) HasMapId() bool {
	if o != nil && !IsNil(o.MapId) {
		return true
	}

	return false
}

// SetMapId gets a reference to the given string and assigns it to the MapId field.
func (o *InstallerDevice) SetMapId(v string) {
	o.MapId = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InstallerDevice) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InstallerDevice) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InstallerDevice) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstallerDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstallerDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstallerDevice) SetName(v string) {
	o.Name = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *InstallerDevice) GetOrientation() float32 {
	if o == nil || IsNil(o.Orientation) {
		var ret float32
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetOrientationOk() (*float32, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *InstallerDevice) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given float32 and assigns it to the Orientation field.
func (o *InstallerDevice) SetOrientation(v float32) {
	o.Orientation = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InstallerDevice) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InstallerDevice) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InstallerDevice) SetSerial(v string) {
	o.Serial = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *InstallerDevice) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *InstallerDevice) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *InstallerDevice) SetSiteName(v string) {
	o.SiteName = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *InstallerDevice) GetUptime() int32 {
	if o == nil || IsNil(o.Uptime) {
		var ret int32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetUptimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Uptime) {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *InstallerDevice) HasUptime() bool {
	if o != nil && !IsNil(o.Uptime) {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int32 and assigns it to the Uptime field.
func (o *InstallerDevice) SetUptime(v int32) {
	o.Uptime = &v
}

// GetVcMac returns the VcMac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstallerDevice) GetVcMac() string {
	if o == nil || IsNil(o.VcMac.Get()) {
		var ret string
		return ret
	}
	return *o.VcMac.Get()
}

// GetVcMacOk returns a tuple with the VcMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstallerDevice) GetVcMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcMac.Get(), o.VcMac.IsSet()
}

// HasVcMac returns a boolean if a field has been set.
func (o *InstallerDevice) HasVcMac() bool {
	if o != nil && o.VcMac.IsSet() {
		return true
	}

	return false
}

// SetVcMac gets a reference to the given NullableString and assigns it to the VcMac field.
func (o *InstallerDevice) SetVcMac(v string) {
	o.VcMac.Set(&v)
}
// SetVcMacNil sets the value for VcMac to be an explicit nil
func (o *InstallerDevice) SetVcMacNil() {
	o.VcMac.Set(nil)
}

// UnsetVcMac ensures that no value is present for VcMac, not even an explicit nil
func (o *InstallerDevice) UnsetVcMac() {
	o.VcMac.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InstallerDevice) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InstallerDevice) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InstallerDevice) SetVersion(v string) {
	o.Version = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *InstallerDevice) GetX() float32 {
	if o == nil || IsNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetXOk() (*float32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InstallerDevice) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *InstallerDevice) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *InstallerDevice) GetY() float32 {
	if o == nil || IsNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallerDevice) GetYOk() (*float32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *InstallerDevice) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *InstallerDevice) SetY(v float32) {
	o.Y = &v
}

func (o InstallerDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallerDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.DeviceprofileName) {
		toSerialize["deviceprofile_name"] = o.DeviceprofileName
	}
	if !IsNil(o.ExtIp) {
		toSerialize["ext_ip"] = o.ExtIp
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.LastSeen) {
		toSerialize["last_seen"] = o.LastSeen
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.MapId) {
		toSerialize["map_id"] = o.MapId
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.SiteName) {
		toSerialize["site_name"] = o.SiteName
	}
	if !IsNil(o.Uptime) {
		toSerialize["uptime"] = o.Uptime
	}
	if o.VcMac.IsSet() {
		toSerialize["vc_mac"] = o.VcMac.Get()
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstallerDevice) UnmarshalJSON(data []byte) (err error) {
	varInstallerDevice := _InstallerDevice{}

	err = json.Unmarshal(data, &varInstallerDevice)

	if err != nil {
		return err
	}

	*o = InstallerDevice(varInstallerDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connected")
		delete(additionalProperties, "deviceprofile_name")
		delete(additionalProperties, "ext_ip")
		delete(additionalProperties, "height")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "last_seen")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "model")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orientation")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "site_name")
		delete(additionalProperties, "uptime")
		delete(additionalProperties, "vc_mac")
		delete(additionalProperties, "version")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstallerDevice struct {
	value *InstallerDevice
	isSet bool
}

func (v NullableInstallerDevice) Get() *InstallerDevice {
	return v.value
}

func (v *NullableInstallerDevice) Set(val *InstallerDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallerDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallerDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallerDevice(val *InstallerDevice) *NullableInstallerDevice {
	return &NullableInstallerDevice{value: val, isSet: true}
}

func (v NullableInstallerDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallerDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


