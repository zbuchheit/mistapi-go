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
	"fmt"
)

// checks if the WlanBonjour type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanBonjour{}

// WlanBonjour bonjour gateway wlan settings
type WlanBonjour struct {
	// additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
	AdditionalVlanIds []int32 `json:"additional_vlan_ids"`
	// whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false
	Enabled *bool `json:"enabled,omitempty"`
	// what services are allowed.  Property key is the service name
	Services map[string]WlanBonjourServiceProperties `json:"services"`
	AdditionalProperties map[string]interface{}
}

type _WlanBonjour WlanBonjour

// NewWlanBonjour instantiates a new WlanBonjour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanBonjour(additionalVlanIds []int32, services map[string]WlanBonjourServiceProperties) *WlanBonjour {
	this := WlanBonjour{}
	this.AdditionalVlanIds = additionalVlanIds
	var enabled bool = false
	this.Enabled = &enabled
	this.Services = services
	return &this
}

// NewWlanBonjourWithDefaults instantiates a new WlanBonjour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanBonjourWithDefaults() *WlanBonjour {
	this := WlanBonjour{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetAdditionalVlanIds returns the AdditionalVlanIds field value
func (o *WlanBonjour) GetAdditionalVlanIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AdditionalVlanIds
}

// GetAdditionalVlanIdsOk returns a tuple with the AdditionalVlanIds field value
// and a boolean to check if the value has been set.
func (o *WlanBonjour) GetAdditionalVlanIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalVlanIds, true
}

// SetAdditionalVlanIds sets field value
func (o *WlanBonjour) SetAdditionalVlanIds(v []int32) {
	o.AdditionalVlanIds = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WlanBonjour) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanBonjour) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WlanBonjour) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WlanBonjour) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServices returns the Services field value
func (o *WlanBonjour) GetServices() map[string]WlanBonjourServiceProperties {
	if o == nil {
		var ret map[string]WlanBonjourServiceProperties
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *WlanBonjour) GetServicesOk() (*map[string]WlanBonjourServiceProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Services, true
}

// SetServices sets field value
func (o *WlanBonjour) SetServices(v map[string]WlanBonjourServiceProperties) {
	o.Services = v
}

func (o WlanBonjour) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanBonjour) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["additional_vlan_ids"] = o.AdditionalVlanIds
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["services"] = o.Services

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanBonjour) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"additional_vlan_ids",
		"services",
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

	varWlanBonjour := _WlanBonjour{}

	err = json.Unmarshal(data, &varWlanBonjour)

	if err != nil {
		return err
	}

	*o = WlanBonjour(varWlanBonjour)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_vlan_ids")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanBonjour struct {
	value *WlanBonjour
	isSet bool
}

func (v NullableWlanBonjour) Get() *WlanBonjour {
	return v.value
}

func (v *NullableWlanBonjour) Set(val *WlanBonjour) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanBonjour) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanBonjour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanBonjour(val *WlanBonjour) *NullableWlanBonjour {
	return &NullableWlanBonjour{value: val, isSet: true}
}

func (v NullableWlanBonjour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanBonjour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


