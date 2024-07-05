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

// checks if the MxedgeUpgradeVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeUpgradeVersion{}

// MxedgeUpgradeVersion version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\\nignored if distro upgrade
type MxedgeUpgradeVersion struct {
	Mxagent string `json:"mxagent"`
	Mxdas *string `json:"mxdas,omitempty"`
	Mxocproxy *string `json:"mxocproxy,omitempty"`
	Radsecproxy *string `json:"radsecproxy,omitempty"`
	Tunterm string `json:"tunterm"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeUpgradeVersion MxedgeUpgradeVersion

// NewMxedgeUpgradeVersion instantiates a new MxedgeUpgradeVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeUpgradeVersion(mxagent string, tunterm string) *MxedgeUpgradeVersion {
	this := MxedgeUpgradeVersion{}
	this.Mxagent = mxagent
	var mxdas string = "current"
	this.Mxdas = &mxdas
	var mxocproxy string = "current"
	this.Mxocproxy = &mxocproxy
	var radsecproxy string = "current"
	this.Radsecproxy = &radsecproxy
	this.Tunterm = tunterm
	return &this
}

// NewMxedgeUpgradeVersionWithDefaults instantiates a new MxedgeUpgradeVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeUpgradeVersionWithDefaults() *MxedgeUpgradeVersion {
	this := MxedgeUpgradeVersion{}
	var mxagent string = "current"
	this.Mxagent = mxagent
	var mxdas string = "current"
	this.Mxdas = &mxdas
	var mxocproxy string = "current"
	this.Mxocproxy = &mxocproxy
	var radsecproxy string = "current"
	this.Radsecproxy = &radsecproxy
	var tunterm string = "current"
	this.Tunterm = tunterm
	return &this
}

// GetMxagent returns the Mxagent field value
func (o *MxedgeUpgradeVersion) GetMxagent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mxagent
}

// GetMxagentOk returns a tuple with the Mxagent field value
// and a boolean to check if the value has been set.
func (o *MxedgeUpgradeVersion) GetMxagentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mxagent, true
}

// SetMxagent sets field value
func (o *MxedgeUpgradeVersion) SetMxagent(v string) {
	o.Mxagent = v
}

// GetMxdas returns the Mxdas field value if set, zero value otherwise.
func (o *MxedgeUpgradeVersion) GetMxdas() string {
	if o == nil || IsNil(o.Mxdas) {
		var ret string
		return ret
	}
	return *o.Mxdas
}

// GetMxdasOk returns a tuple with the Mxdas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeUpgradeVersion) GetMxdasOk() (*string, bool) {
	if o == nil || IsNil(o.Mxdas) {
		return nil, false
	}
	return o.Mxdas, true
}

// HasMxdas returns a boolean if a field has been set.
func (o *MxedgeUpgradeVersion) HasMxdas() bool {
	if o != nil && !IsNil(o.Mxdas) {
		return true
	}

	return false
}

// SetMxdas gets a reference to the given string and assigns it to the Mxdas field.
func (o *MxedgeUpgradeVersion) SetMxdas(v string) {
	o.Mxdas = &v
}

// GetMxocproxy returns the Mxocproxy field value if set, zero value otherwise.
func (o *MxedgeUpgradeVersion) GetMxocproxy() string {
	if o == nil || IsNil(o.Mxocproxy) {
		var ret string
		return ret
	}
	return *o.Mxocproxy
}

// GetMxocproxyOk returns a tuple with the Mxocproxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeUpgradeVersion) GetMxocproxyOk() (*string, bool) {
	if o == nil || IsNil(o.Mxocproxy) {
		return nil, false
	}
	return o.Mxocproxy, true
}

// HasMxocproxy returns a boolean if a field has been set.
func (o *MxedgeUpgradeVersion) HasMxocproxy() bool {
	if o != nil && !IsNil(o.Mxocproxy) {
		return true
	}

	return false
}

// SetMxocproxy gets a reference to the given string and assigns it to the Mxocproxy field.
func (o *MxedgeUpgradeVersion) SetMxocproxy(v string) {
	o.Mxocproxy = &v
}

// GetRadsecproxy returns the Radsecproxy field value if set, zero value otherwise.
func (o *MxedgeUpgradeVersion) GetRadsecproxy() string {
	if o == nil || IsNil(o.Radsecproxy) {
		var ret string
		return ret
	}
	return *o.Radsecproxy
}

// GetRadsecproxyOk returns a tuple with the Radsecproxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeUpgradeVersion) GetRadsecproxyOk() (*string, bool) {
	if o == nil || IsNil(o.Radsecproxy) {
		return nil, false
	}
	return o.Radsecproxy, true
}

// HasRadsecproxy returns a boolean if a field has been set.
func (o *MxedgeUpgradeVersion) HasRadsecproxy() bool {
	if o != nil && !IsNil(o.Radsecproxy) {
		return true
	}

	return false
}

// SetRadsecproxy gets a reference to the given string and assigns it to the Radsecproxy field.
func (o *MxedgeUpgradeVersion) SetRadsecproxy(v string) {
	o.Radsecproxy = &v
}

// GetTunterm returns the Tunterm field value
func (o *MxedgeUpgradeVersion) GetTunterm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tunterm
}

// GetTuntermOk returns a tuple with the Tunterm field value
// and a boolean to check if the value has been set.
func (o *MxedgeUpgradeVersion) GetTuntermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tunterm, true
}

// SetTunterm sets field value
func (o *MxedgeUpgradeVersion) SetTunterm(v string) {
	o.Tunterm = v
}

func (o MxedgeUpgradeVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeUpgradeVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mxagent"] = o.Mxagent
	if !IsNil(o.Mxdas) {
		toSerialize["mxdas"] = o.Mxdas
	}
	if !IsNil(o.Mxocproxy) {
		toSerialize["mxocproxy"] = o.Mxocproxy
	}
	if !IsNil(o.Radsecproxy) {
		toSerialize["radsecproxy"] = o.Radsecproxy
	}
	toSerialize["tunterm"] = o.Tunterm

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeUpgradeVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mxagent",
		"tunterm",
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

	varMxedgeUpgradeVersion := _MxedgeUpgradeVersion{}

	err = json.Unmarshal(data, &varMxedgeUpgradeVersion)

	if err != nil {
		return err
	}

	*o = MxedgeUpgradeVersion(varMxedgeUpgradeVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mxagent")
		delete(additionalProperties, "mxdas")
		delete(additionalProperties, "mxocproxy")
		delete(additionalProperties, "radsecproxy")
		delete(additionalProperties, "tunterm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeUpgradeVersion struct {
	value *MxedgeUpgradeVersion
	isSet bool
}

func (v NullableMxedgeUpgradeVersion) Get() *MxedgeUpgradeVersion {
	return v.value
}

func (v *NullableMxedgeUpgradeVersion) Set(val *MxedgeUpgradeVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeUpgradeVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeUpgradeVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeUpgradeVersion(val *MxedgeUpgradeVersion) *NullableMxedgeUpgradeVersion {
	return &NullableMxedgeUpgradeVersion{value: val, isSet: true}
}

func (v NullableMxedgeUpgradeVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeUpgradeVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


