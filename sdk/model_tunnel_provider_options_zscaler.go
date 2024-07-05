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

// checks if the TunnelProviderOptionsZscaler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelProviderOptionsZscaler{}

// TunnelProviderOptionsZscaler for zscaler-ipsec and zscaler-gre
type TunnelProviderOptionsZscaler struct {
	AupAcceptanceRequired *bool `json:"aup_acceptance_required,omitempty"`
	// days before AUP is requested again
	AupExpire *int32 `json:"aup_expire,omitempty"`
	// proxy HTTPs traffic, requiring Zscaler cert to be installed in browser
	AupSslProxy *bool `json:"aup_ssl_proxy,omitempty"`
	// the download bandwidth cap of the link, in Mbps
	DownloadMbps *int32 `json:"download_mbps,omitempty"`
	// if `use_xff`==`true`, display Acceptable Use Policy (AUP)
	EnableAup *bool `json:"enable_aup,omitempty"`
	// when `enforce_authentication`==`false`, display caution notification for non-authenticated users
	EnableCaution *bool `json:"enable_caution,omitempty"`
	EnforceAuthentication *bool `json:"enforce_authentication,omitempty"`
	Name *string `json:"name,omitempty"`
	// if `use_xff`==`true`
	SubLocations []TunnelProviderOptionsZscalerSubLocation `json:"sub_locations,omitempty"`
	// the download bandwidth cap of the link, in Mbps
	UploadMbps *int32 `json:"upload_mbps,omitempty"`
	// location uses proxy chaining to forward traffic
	UseXff *bool `json:"use_xff,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelProviderOptionsZscaler TunnelProviderOptionsZscaler

// NewTunnelProviderOptionsZscaler instantiates a new TunnelProviderOptionsZscaler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelProviderOptionsZscaler() *TunnelProviderOptionsZscaler {
	this := TunnelProviderOptionsZscaler{}
	var aupAcceptanceRequired bool = true
	this.AupAcceptanceRequired = &aupAcceptanceRequired
	var aupExpire int32 = 1
	this.AupExpire = &aupExpire
	var aupSslProxy bool = false
	this.AupSslProxy = &aupSslProxy
	var enableAup bool = false
	this.EnableAup = &enableAup
	var enableCaution bool = false
	this.EnableCaution = &enableCaution
	var enforceAuthentication bool = false
	this.EnforceAuthentication = &enforceAuthentication
	return &this
}

// NewTunnelProviderOptionsZscalerWithDefaults instantiates a new TunnelProviderOptionsZscaler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelProviderOptionsZscalerWithDefaults() *TunnelProviderOptionsZscaler {
	this := TunnelProviderOptionsZscaler{}
	var aupAcceptanceRequired bool = true
	this.AupAcceptanceRequired = &aupAcceptanceRequired
	var aupExpire int32 = 1
	this.AupExpire = &aupExpire
	var aupSslProxy bool = false
	this.AupSslProxy = &aupSslProxy
	var enableAup bool = false
	this.EnableAup = &enableAup
	var enableCaution bool = false
	this.EnableCaution = &enableCaution
	var enforceAuthentication bool = false
	this.EnforceAuthentication = &enforceAuthentication
	return &this
}

// GetAupAcceptanceRequired returns the AupAcceptanceRequired field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetAupAcceptanceRequired() bool {
	if o == nil || IsNil(o.AupAcceptanceRequired) {
		var ret bool
		return ret
	}
	return *o.AupAcceptanceRequired
}

// GetAupAcceptanceRequiredOk returns a tuple with the AupAcceptanceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetAupAcceptanceRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.AupAcceptanceRequired) {
		return nil, false
	}
	return o.AupAcceptanceRequired, true
}

// HasAupAcceptanceRequired returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasAupAcceptanceRequired() bool {
	if o != nil && !IsNil(o.AupAcceptanceRequired) {
		return true
	}

	return false
}

// SetAupAcceptanceRequired gets a reference to the given bool and assigns it to the AupAcceptanceRequired field.
func (o *TunnelProviderOptionsZscaler) SetAupAcceptanceRequired(v bool) {
	o.AupAcceptanceRequired = &v
}

// GetAupExpire returns the AupExpire field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetAupExpire() int32 {
	if o == nil || IsNil(o.AupExpire) {
		var ret int32
		return ret
	}
	return *o.AupExpire
}

// GetAupExpireOk returns a tuple with the AupExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetAupExpireOk() (*int32, bool) {
	if o == nil || IsNil(o.AupExpire) {
		return nil, false
	}
	return o.AupExpire, true
}

// HasAupExpire returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasAupExpire() bool {
	if o != nil && !IsNil(o.AupExpire) {
		return true
	}

	return false
}

// SetAupExpire gets a reference to the given int32 and assigns it to the AupExpire field.
func (o *TunnelProviderOptionsZscaler) SetAupExpire(v int32) {
	o.AupExpire = &v
}

// GetAupSslProxy returns the AupSslProxy field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetAupSslProxy() bool {
	if o == nil || IsNil(o.AupSslProxy) {
		var ret bool
		return ret
	}
	return *o.AupSslProxy
}

// GetAupSslProxyOk returns a tuple with the AupSslProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetAupSslProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.AupSslProxy) {
		return nil, false
	}
	return o.AupSslProxy, true
}

// HasAupSslProxy returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasAupSslProxy() bool {
	if o != nil && !IsNil(o.AupSslProxy) {
		return true
	}

	return false
}

// SetAupSslProxy gets a reference to the given bool and assigns it to the AupSslProxy field.
func (o *TunnelProviderOptionsZscaler) SetAupSslProxy(v bool) {
	o.AupSslProxy = &v
}

// GetDownloadMbps returns the DownloadMbps field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetDownloadMbps() int32 {
	if o == nil || IsNil(o.DownloadMbps) {
		var ret int32
		return ret
	}
	return *o.DownloadMbps
}

// GetDownloadMbpsOk returns a tuple with the DownloadMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetDownloadMbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.DownloadMbps) {
		return nil, false
	}
	return o.DownloadMbps, true
}

// HasDownloadMbps returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasDownloadMbps() bool {
	if o != nil && !IsNil(o.DownloadMbps) {
		return true
	}

	return false
}

// SetDownloadMbps gets a reference to the given int32 and assigns it to the DownloadMbps field.
func (o *TunnelProviderOptionsZscaler) SetDownloadMbps(v int32) {
	o.DownloadMbps = &v
}

// GetEnableAup returns the EnableAup field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetEnableAup() bool {
	if o == nil || IsNil(o.EnableAup) {
		var ret bool
		return ret
	}
	return *o.EnableAup
}

// GetEnableAupOk returns a tuple with the EnableAup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetEnableAupOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAup) {
		return nil, false
	}
	return o.EnableAup, true
}

// HasEnableAup returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasEnableAup() bool {
	if o != nil && !IsNil(o.EnableAup) {
		return true
	}

	return false
}

// SetEnableAup gets a reference to the given bool and assigns it to the EnableAup field.
func (o *TunnelProviderOptionsZscaler) SetEnableAup(v bool) {
	o.EnableAup = &v
}

// GetEnableCaution returns the EnableCaution field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetEnableCaution() bool {
	if o == nil || IsNil(o.EnableCaution) {
		var ret bool
		return ret
	}
	return *o.EnableCaution
}

// GetEnableCautionOk returns a tuple with the EnableCaution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetEnableCautionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCaution) {
		return nil, false
	}
	return o.EnableCaution, true
}

// HasEnableCaution returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasEnableCaution() bool {
	if o != nil && !IsNil(o.EnableCaution) {
		return true
	}

	return false
}

// SetEnableCaution gets a reference to the given bool and assigns it to the EnableCaution field.
func (o *TunnelProviderOptionsZscaler) SetEnableCaution(v bool) {
	o.EnableCaution = &v
}

// GetEnforceAuthentication returns the EnforceAuthentication field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetEnforceAuthentication() bool {
	if o == nil || IsNil(o.EnforceAuthentication) {
		var ret bool
		return ret
	}
	return *o.EnforceAuthentication
}

// GetEnforceAuthenticationOk returns a tuple with the EnforceAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetEnforceAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceAuthentication) {
		return nil, false
	}
	return o.EnforceAuthentication, true
}

// HasEnforceAuthentication returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasEnforceAuthentication() bool {
	if o != nil && !IsNil(o.EnforceAuthentication) {
		return true
	}

	return false
}

// SetEnforceAuthentication gets a reference to the given bool and assigns it to the EnforceAuthentication field.
func (o *TunnelProviderOptionsZscaler) SetEnforceAuthentication(v bool) {
	o.EnforceAuthentication = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TunnelProviderOptionsZscaler) SetName(v string) {
	o.Name = &v
}

// GetSubLocations returns the SubLocations field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetSubLocations() []TunnelProviderOptionsZscalerSubLocation {
	if o == nil || IsNil(o.SubLocations) {
		var ret []TunnelProviderOptionsZscalerSubLocation
		return ret
	}
	return o.SubLocations
}

// GetSubLocationsOk returns a tuple with the SubLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetSubLocationsOk() ([]TunnelProviderOptionsZscalerSubLocation, bool) {
	if o == nil || IsNil(o.SubLocations) {
		return nil, false
	}
	return o.SubLocations, true
}

// HasSubLocations returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasSubLocations() bool {
	if o != nil && !IsNil(o.SubLocations) {
		return true
	}

	return false
}

// SetSubLocations gets a reference to the given []TunnelProviderOptionsZscalerSubLocation and assigns it to the SubLocations field.
func (o *TunnelProviderOptionsZscaler) SetSubLocations(v []TunnelProviderOptionsZscalerSubLocation) {
	o.SubLocations = v
}

// GetUploadMbps returns the UploadMbps field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetUploadMbps() int32 {
	if o == nil || IsNil(o.UploadMbps) {
		var ret int32
		return ret
	}
	return *o.UploadMbps
}

// GetUploadMbpsOk returns a tuple with the UploadMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetUploadMbpsOk() (*int32, bool) {
	if o == nil || IsNil(o.UploadMbps) {
		return nil, false
	}
	return o.UploadMbps, true
}

// HasUploadMbps returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasUploadMbps() bool {
	if o != nil && !IsNil(o.UploadMbps) {
		return true
	}

	return false
}

// SetUploadMbps gets a reference to the given int32 and assigns it to the UploadMbps field.
func (o *TunnelProviderOptionsZscaler) SetUploadMbps(v int32) {
	o.UploadMbps = &v
}

// GetUseXff returns the UseXff field value if set, zero value otherwise.
func (o *TunnelProviderOptionsZscaler) GetUseXff() bool {
	if o == nil || IsNil(o.UseXff) {
		var ret bool
		return ret
	}
	return *o.UseXff
}

// GetUseXffOk returns a tuple with the UseXff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptionsZscaler) GetUseXffOk() (*bool, bool) {
	if o == nil || IsNil(o.UseXff) {
		return nil, false
	}
	return o.UseXff, true
}

// HasUseXff returns a boolean if a field has been set.
func (o *TunnelProviderOptionsZscaler) HasUseXff() bool {
	if o != nil && !IsNil(o.UseXff) {
		return true
	}

	return false
}

// SetUseXff gets a reference to the given bool and assigns it to the UseXff field.
func (o *TunnelProviderOptionsZscaler) SetUseXff(v bool) {
	o.UseXff = &v
}

func (o TunnelProviderOptionsZscaler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelProviderOptionsZscaler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AupAcceptanceRequired) {
		toSerialize["aup_acceptance_required"] = o.AupAcceptanceRequired
	}
	if !IsNil(o.AupExpire) {
		toSerialize["aup_expire"] = o.AupExpire
	}
	if !IsNil(o.AupSslProxy) {
		toSerialize["aup_ssl_proxy"] = o.AupSslProxy
	}
	if !IsNil(o.DownloadMbps) {
		toSerialize["download_mbps"] = o.DownloadMbps
	}
	if !IsNil(o.EnableAup) {
		toSerialize["enable_aup"] = o.EnableAup
	}
	if !IsNil(o.EnableCaution) {
		toSerialize["enable_caution"] = o.EnableCaution
	}
	if !IsNil(o.EnforceAuthentication) {
		toSerialize["enforce_authentication"] = o.EnforceAuthentication
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SubLocations) {
		toSerialize["sub_locations"] = o.SubLocations
	}
	if !IsNil(o.UploadMbps) {
		toSerialize["upload_mbps"] = o.UploadMbps
	}
	if !IsNil(o.UseXff) {
		toSerialize["use_xff"] = o.UseXff
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelProviderOptionsZscaler) UnmarshalJSON(data []byte) (err error) {
	varTunnelProviderOptionsZscaler := _TunnelProviderOptionsZscaler{}

	err = json.Unmarshal(data, &varTunnelProviderOptionsZscaler)

	if err != nil {
		return err
	}

	*o = TunnelProviderOptionsZscaler(varTunnelProviderOptionsZscaler)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aup_acceptance_required")
		delete(additionalProperties, "aup_expire")
		delete(additionalProperties, "aup_ssl_proxy")
		delete(additionalProperties, "download_mbps")
		delete(additionalProperties, "enable_aup")
		delete(additionalProperties, "enable_caution")
		delete(additionalProperties, "enforce_authentication")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sub_locations")
		delete(additionalProperties, "upload_mbps")
		delete(additionalProperties, "use_xff")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelProviderOptionsZscaler struct {
	value *TunnelProviderOptionsZscaler
	isSet bool
}

func (v NullableTunnelProviderOptionsZscaler) Get() *TunnelProviderOptionsZscaler {
	return v.value
}

func (v *NullableTunnelProviderOptionsZscaler) Set(val *TunnelProviderOptionsZscaler) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelProviderOptionsZscaler) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelProviderOptionsZscaler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelProviderOptionsZscaler(val *TunnelProviderOptionsZscaler) *NullableTunnelProviderOptionsZscaler {
	return &NullableTunnelProviderOptionsZscaler{value: val, isSet: true}
}

func (v NullableTunnelProviderOptionsZscaler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelProviderOptionsZscaler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


