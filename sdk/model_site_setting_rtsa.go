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

// checks if the SiteSettingRtsa type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingRtsa{}

// SiteSettingRtsa managed mobility
type SiteSettingRtsa struct {
	AppWaking *bool `json:"app_waking,omitempty"`
	DisableDeadReckoning *bool `json:"disable_dead_reckoning,omitempty"`
	DisablePressureSensor *bool `json:"disable_pressure_sensor,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// asset tracking related
	TrackAsset *bool `json:"track_asset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingRtsa SiteSettingRtsa

// NewSiteSettingRtsa instantiates a new SiteSettingRtsa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingRtsa() *SiteSettingRtsa {
	this := SiteSettingRtsa{}
	var appWaking bool = false
	this.AppWaking = &appWaking
	var disablePressureSensor bool = false
	this.DisablePressureSensor = &disablePressureSensor
	var trackAsset bool = false
	this.TrackAsset = &trackAsset
	return &this
}

// NewSiteSettingRtsaWithDefaults instantiates a new SiteSettingRtsa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingRtsaWithDefaults() *SiteSettingRtsa {
	this := SiteSettingRtsa{}
	var appWaking bool = false
	this.AppWaking = &appWaking
	var disablePressureSensor bool = false
	this.DisablePressureSensor = &disablePressureSensor
	var trackAsset bool = false
	this.TrackAsset = &trackAsset
	return &this
}

// GetAppWaking returns the AppWaking field value if set, zero value otherwise.
func (o *SiteSettingRtsa) GetAppWaking() bool {
	if o == nil || IsNil(o.AppWaking) {
		var ret bool
		return ret
	}
	return *o.AppWaking
}

// GetAppWakingOk returns a tuple with the AppWaking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingRtsa) GetAppWakingOk() (*bool, bool) {
	if o == nil || IsNil(o.AppWaking) {
		return nil, false
	}
	return o.AppWaking, true
}

// HasAppWaking returns a boolean if a field has been set.
func (o *SiteSettingRtsa) HasAppWaking() bool {
	if o != nil && !IsNil(o.AppWaking) {
		return true
	}

	return false
}

// SetAppWaking gets a reference to the given bool and assigns it to the AppWaking field.
func (o *SiteSettingRtsa) SetAppWaking(v bool) {
	o.AppWaking = &v
}

// GetDisableDeadReckoning returns the DisableDeadReckoning field value if set, zero value otherwise.
func (o *SiteSettingRtsa) GetDisableDeadReckoning() bool {
	if o == nil || IsNil(o.DisableDeadReckoning) {
		var ret bool
		return ret
	}
	return *o.DisableDeadReckoning
}

// GetDisableDeadReckoningOk returns a tuple with the DisableDeadReckoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingRtsa) GetDisableDeadReckoningOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableDeadReckoning) {
		return nil, false
	}
	return o.DisableDeadReckoning, true
}

// HasDisableDeadReckoning returns a boolean if a field has been set.
func (o *SiteSettingRtsa) HasDisableDeadReckoning() bool {
	if o != nil && !IsNil(o.DisableDeadReckoning) {
		return true
	}

	return false
}

// SetDisableDeadReckoning gets a reference to the given bool and assigns it to the DisableDeadReckoning field.
func (o *SiteSettingRtsa) SetDisableDeadReckoning(v bool) {
	o.DisableDeadReckoning = &v
}

// GetDisablePressureSensor returns the DisablePressureSensor field value if set, zero value otherwise.
func (o *SiteSettingRtsa) GetDisablePressureSensor() bool {
	if o == nil || IsNil(o.DisablePressureSensor) {
		var ret bool
		return ret
	}
	return *o.DisablePressureSensor
}

// GetDisablePressureSensorOk returns a tuple with the DisablePressureSensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingRtsa) GetDisablePressureSensorOk() (*bool, bool) {
	if o == nil || IsNil(o.DisablePressureSensor) {
		return nil, false
	}
	return o.DisablePressureSensor, true
}

// HasDisablePressureSensor returns a boolean if a field has been set.
func (o *SiteSettingRtsa) HasDisablePressureSensor() bool {
	if o != nil && !IsNil(o.DisablePressureSensor) {
		return true
	}

	return false
}

// SetDisablePressureSensor gets a reference to the given bool and assigns it to the DisablePressureSensor field.
func (o *SiteSettingRtsa) SetDisablePressureSensor(v bool) {
	o.DisablePressureSensor = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SiteSettingRtsa) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingRtsa) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SiteSettingRtsa) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SiteSettingRtsa) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTrackAsset returns the TrackAsset field value if set, zero value otherwise.
func (o *SiteSettingRtsa) GetTrackAsset() bool {
	if o == nil || IsNil(o.TrackAsset) {
		var ret bool
		return ret
	}
	return *o.TrackAsset
}

// GetTrackAssetOk returns a tuple with the TrackAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingRtsa) GetTrackAssetOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackAsset) {
		return nil, false
	}
	return o.TrackAsset, true
}

// HasTrackAsset returns a boolean if a field has been set.
func (o *SiteSettingRtsa) HasTrackAsset() bool {
	if o != nil && !IsNil(o.TrackAsset) {
		return true
	}

	return false
}

// SetTrackAsset gets a reference to the given bool and assigns it to the TrackAsset field.
func (o *SiteSettingRtsa) SetTrackAsset(v bool) {
	o.TrackAsset = &v
}

func (o SiteSettingRtsa) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingRtsa) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppWaking) {
		toSerialize["app_waking"] = o.AppWaking
	}
	if !IsNil(o.DisableDeadReckoning) {
		toSerialize["disable_dead_reckoning"] = o.DisableDeadReckoning
	}
	if !IsNil(o.DisablePressureSensor) {
		toSerialize["disable_pressure_sensor"] = o.DisablePressureSensor
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.TrackAsset) {
		toSerialize["track_asset"] = o.TrackAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingRtsa) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingRtsa := _SiteSettingRtsa{}

	err = json.Unmarshal(data, &varSiteSettingRtsa)

	if err != nil {
		return err
	}

	*o = SiteSettingRtsa(varSiteSettingRtsa)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app_waking")
		delete(additionalProperties, "disable_dead_reckoning")
		delete(additionalProperties, "disable_pressure_sensor")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "track_asset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingRtsa struct {
	value *SiteSettingRtsa
	isSet bool
}

func (v NullableSiteSettingRtsa) Get() *SiteSettingRtsa {
	return v.value
}

func (v *NullableSiteSettingRtsa) Set(val *SiteSettingRtsa) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingRtsa) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingRtsa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingRtsa(val *SiteSettingRtsa) *NullableSiteSettingRtsa {
	return &NullableSiteSettingRtsa{value: val, isSet: true}
}

func (v NullableSiteSettingRtsa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingRtsa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


