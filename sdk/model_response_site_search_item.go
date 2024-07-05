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
	"fmt"
)

// checks if the ResponseSiteSearchItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSiteSearchItem{}

// ResponseSiteSearchItem struct for ResponseSiteSearchItem
type ResponseSiteSearchItem struct {
	AutoUpgradeEnabled bool `json:"auto_upgrade_enabled"`
	AutoUpgradeVersion string `json:"auto_upgrade_version"`
	CountryCode NullableString `json:"country_code"`
	HoneypotEnabled bool `json:"honeypot_enabled"`
	Id string `json:"id"`
	Name string `json:"name"`
	OrgId string `json:"org_id"`
	SiteId string `json:"site_id"`
	Timestamp float32 `json:"timestamp"`
	Timezone string `json:"timezone"`
	VnaEnabled bool `json:"vna_enabled"`
	WifiEnabled bool `json:"wifi_enabled"`
	AdditionalProperties map[string]interface{}
}

type _ResponseSiteSearchItem ResponseSiteSearchItem

// NewResponseSiteSearchItem instantiates a new ResponseSiteSearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSiteSearchItem(autoUpgradeEnabled bool, autoUpgradeVersion string, countryCode NullableString, honeypotEnabled bool, id string, name string, orgId string, siteId string, timestamp float32, timezone string, vnaEnabled bool, wifiEnabled bool) *ResponseSiteSearchItem {
	this := ResponseSiteSearchItem{}
	this.AutoUpgradeEnabled = autoUpgradeEnabled
	this.AutoUpgradeVersion = autoUpgradeVersion
	this.CountryCode = countryCode
	this.HoneypotEnabled = honeypotEnabled
	this.Id = id
	this.Name = name
	this.OrgId = orgId
	this.SiteId = siteId
	this.Timestamp = timestamp
	this.Timezone = timezone
	this.VnaEnabled = vnaEnabled
	this.WifiEnabled = wifiEnabled
	return &this
}

// NewResponseSiteSearchItemWithDefaults instantiates a new ResponseSiteSearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSiteSearchItemWithDefaults() *ResponseSiteSearchItem {
	this := ResponseSiteSearchItem{}
	return &this
}

// GetAutoUpgradeEnabled returns the AutoUpgradeEnabled field value
func (o *ResponseSiteSearchItem) GetAutoUpgradeEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoUpgradeEnabled
}

// GetAutoUpgradeEnabledOk returns a tuple with the AutoUpgradeEnabled field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetAutoUpgradeEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoUpgradeEnabled, true
}

// SetAutoUpgradeEnabled sets field value
func (o *ResponseSiteSearchItem) SetAutoUpgradeEnabled(v bool) {
	o.AutoUpgradeEnabled = v
}

// GetAutoUpgradeVersion returns the AutoUpgradeVersion field value
func (o *ResponseSiteSearchItem) GetAutoUpgradeVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AutoUpgradeVersion
}

// GetAutoUpgradeVersionOk returns a tuple with the AutoUpgradeVersion field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetAutoUpgradeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoUpgradeVersion, true
}

// SetAutoUpgradeVersion sets field value
func (o *ResponseSiteSearchItem) SetAutoUpgradeVersion(v string) {
	o.AutoUpgradeVersion = v
}

// GetCountryCode returns the CountryCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResponseSiteSearchItem) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseSiteSearchItem) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// SetCountryCode sets field value
func (o *ResponseSiteSearchItem) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}

// GetHoneypotEnabled returns the HoneypotEnabled field value
func (o *ResponseSiteSearchItem) GetHoneypotEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HoneypotEnabled
}

// GetHoneypotEnabledOk returns a tuple with the HoneypotEnabled field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetHoneypotEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HoneypotEnabled, true
}

// SetHoneypotEnabled sets field value
func (o *ResponseSiteSearchItem) SetHoneypotEnabled(v bool) {
	o.HoneypotEnabled = v
}

// GetId returns the Id field value
func (o *ResponseSiteSearchItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseSiteSearchItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseSiteSearchItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseSiteSearchItem) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value
func (o *ResponseSiteSearchItem) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ResponseSiteSearchItem) SetOrgId(v string) {
	o.OrgId = v
}

// GetSiteId returns the SiteId field value
func (o *ResponseSiteSearchItem) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *ResponseSiteSearchItem) SetSiteId(v string) {
	o.SiteId = v
}

// GetTimestamp returns the Timestamp field value
func (o *ResponseSiteSearchItem) GetTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ResponseSiteSearchItem) SetTimestamp(v float32) {
	o.Timestamp = v
}

// GetTimezone returns the Timezone field value
func (o *ResponseSiteSearchItem) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *ResponseSiteSearchItem) SetTimezone(v string) {
	o.Timezone = v
}

// GetVnaEnabled returns the VnaEnabled field value
func (o *ResponseSiteSearchItem) GetVnaEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VnaEnabled
}

// GetVnaEnabledOk returns a tuple with the VnaEnabled field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetVnaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VnaEnabled, true
}

// SetVnaEnabled sets field value
func (o *ResponseSiteSearchItem) SetVnaEnabled(v bool) {
	o.VnaEnabled = v
}

// GetWifiEnabled returns the WifiEnabled field value
func (o *ResponseSiteSearchItem) GetWifiEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WifiEnabled
}

// GetWifiEnabledOk returns a tuple with the WifiEnabled field value
// and a boolean to check if the value has been set.
func (o *ResponseSiteSearchItem) GetWifiEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WifiEnabled, true
}

// SetWifiEnabled sets field value
func (o *ResponseSiteSearchItem) SetWifiEnabled(v bool) {
	o.WifiEnabled = v
}

func (o ResponseSiteSearchItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSiteSearchItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auto_upgrade_enabled"] = o.AutoUpgradeEnabled
	toSerialize["auto_upgrade_version"] = o.AutoUpgradeVersion
	toSerialize["country_code"] = o.CountryCode.Get()
	toSerialize["honeypot_enabled"] = o.HoneypotEnabled
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	toSerialize["site_id"] = o.SiteId
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["timezone"] = o.Timezone
	toSerialize["vna_enabled"] = o.VnaEnabled
	toSerialize["wifi_enabled"] = o.WifiEnabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseSiteSearchItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"auto_upgrade_enabled",
		"auto_upgrade_version",
		"country_code",
		"honeypot_enabled",
		"id",
		"name",
		"org_id",
		"site_id",
		"timestamp",
		"timezone",
		"vna_enabled",
		"wifi_enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResponseSiteSearchItem := _ResponseSiteSearchItem{}

	err = json.Unmarshal(data, &varResponseSiteSearchItem)

	if err != nil {
		return err
	}

	*o = ResponseSiteSearchItem(varResponseSiteSearchItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auto_upgrade_enabled")
		delete(additionalProperties, "auto_upgrade_version")
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "honeypot_enabled")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "vna_enabled")
		delete(additionalProperties, "wifi_enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseSiteSearchItem struct {
	value *ResponseSiteSearchItem
	isSet bool
}

func (v NullableResponseSiteSearchItem) Get() *ResponseSiteSearchItem {
	return v.value
}

func (v *NullableResponseSiteSearchItem) Set(val *ResponseSiteSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSiteSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSiteSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSiteSearchItem(val *ResponseSiteSearchItem) *NullableResponseSiteSearchItem {
	return &NullableResponseSiteSearchItem{value: val, isSet: true}
}

func (v NullableResponseSiteSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSiteSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


