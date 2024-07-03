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

// checks if the OrgDeviceUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgDeviceUpgrade{}

// OrgDeviceUpgrade struct for OrgDeviceUpgrade
type OrgDeviceUpgrade struct {
	Id *string `json:"id,omitempty"`
	SiteUpgrades []OrgDeviceUpgradeSiteUpgrade `json:"site_upgrades,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgDeviceUpgrade OrgDeviceUpgrade

// NewOrgDeviceUpgrade instantiates a new OrgDeviceUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgDeviceUpgrade() *OrgDeviceUpgrade {
	this := OrgDeviceUpgrade{}
	return &this
}

// NewOrgDeviceUpgradeWithDefaults instantiates a new OrgDeviceUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgDeviceUpgradeWithDefaults() *OrgDeviceUpgrade {
	this := OrgDeviceUpgrade{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrgDeviceUpgrade) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDeviceUpgrade) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgDeviceUpgrade) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgDeviceUpgrade) SetId(v string) {
	o.Id = &v
}

// GetSiteUpgrades returns the SiteUpgrades field value if set, zero value otherwise.
func (o *OrgDeviceUpgrade) GetSiteUpgrades() []OrgDeviceUpgradeSiteUpgrade {
	if o == nil || IsNil(o.SiteUpgrades) {
		var ret []OrgDeviceUpgradeSiteUpgrade
		return ret
	}
	return o.SiteUpgrades
}

// GetSiteUpgradesOk returns a tuple with the SiteUpgrades field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDeviceUpgrade) GetSiteUpgradesOk() ([]OrgDeviceUpgradeSiteUpgrade, bool) {
	if o == nil || IsNil(o.SiteUpgrades) {
		return nil, false
	}
	return o.SiteUpgrades, true
}

// HasSiteUpgrades returns a boolean if a field has been set.
func (o *OrgDeviceUpgrade) HasSiteUpgrades() bool {
	if o != nil && !IsNil(o.SiteUpgrades) {
		return true
	}

	return false
}

// SetSiteUpgrades gets a reference to the given []OrgDeviceUpgradeSiteUpgrade and assigns it to the SiteUpgrades field.
func (o *OrgDeviceUpgrade) SetSiteUpgrades(v []OrgDeviceUpgradeSiteUpgrade) {
	o.SiteUpgrades = v
}

func (o OrgDeviceUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgDeviceUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SiteUpgrades) {
		toSerialize["site_upgrades"] = o.SiteUpgrades
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgDeviceUpgrade) UnmarshalJSON(data []byte) (err error) {
	varOrgDeviceUpgrade := _OrgDeviceUpgrade{}

	err = json.Unmarshal(data, &varOrgDeviceUpgrade)

	if err != nil {
		return err
	}

	*o = OrgDeviceUpgrade(varOrgDeviceUpgrade)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "site_upgrades")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgDeviceUpgrade struct {
	value *OrgDeviceUpgrade
	isSet bool
}

func (v NullableOrgDeviceUpgrade) Get() *OrgDeviceUpgrade {
	return v.value
}

func (v *NullableOrgDeviceUpgrade) Set(val *OrgDeviceUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgDeviceUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgDeviceUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgDeviceUpgrade(val *OrgDeviceUpgrade) *NullableOrgDeviceUpgrade {
	return &NullableOrgDeviceUpgrade{value: val, isSet: true}
}

func (v NullableOrgDeviceUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgDeviceUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


