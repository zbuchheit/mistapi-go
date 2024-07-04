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

// checks if the AppProbing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppProbing{}

// AppProbing struct for AppProbing
type AppProbing struct {
	Apps []string `json:"apps,omitempty"`
	CustomApps []AppProbingCustomApp `json:"custom_apps,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppProbing AppProbing

// NewAppProbing instantiates a new AppProbing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppProbing() *AppProbing {
	this := AppProbing{}
	return &this
}

// NewAppProbingWithDefaults instantiates a new AppProbing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppProbingWithDefaults() *AppProbing {
	this := AppProbing{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *AppProbing) GetApps() []string {
	if o == nil || IsNil(o.Apps) {
		var ret []string
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProbing) GetAppsOk() ([]string, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *AppProbing) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given []string and assigns it to the Apps field.
func (o *AppProbing) SetApps(v []string) {
	o.Apps = v
}

// GetCustomApps returns the CustomApps field value if set, zero value otherwise.
func (o *AppProbing) GetCustomApps() []AppProbingCustomApp {
	if o == nil || IsNil(o.CustomApps) {
		var ret []AppProbingCustomApp
		return ret
	}
	return o.CustomApps
}

// GetCustomAppsOk returns a tuple with the CustomApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProbing) GetCustomAppsOk() ([]AppProbingCustomApp, bool) {
	if o == nil || IsNil(o.CustomApps) {
		return nil, false
	}
	return o.CustomApps, true
}

// HasCustomApps returns a boolean if a field has been set.
func (o *AppProbing) HasCustomApps() bool {
	if o != nil && !IsNil(o.CustomApps) {
		return true
	}

	return false
}

// SetCustomApps gets a reference to the given []AppProbingCustomApp and assigns it to the CustomApps field.
func (o *AppProbing) SetCustomApps(v []AppProbingCustomApp) {
	o.CustomApps = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AppProbing) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppProbing) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AppProbing) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AppProbing) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AppProbing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppProbing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.CustomApps) {
		toSerialize["custom_apps"] = o.CustomApps
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppProbing) UnmarshalJSON(data []byte) (err error) {
	varAppProbing := _AppProbing{}

	err = json.Unmarshal(data, &varAppProbing)

	if err != nil {
		return err
	}

	*o = AppProbing(varAppProbing)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apps")
		delete(additionalProperties, "custom_apps")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppProbing struct {
	value *AppProbing
	isSet bool
}

func (v NullableAppProbing) Get() *AppProbing {
	return v.value
}

func (v *NullableAppProbing) Set(val *AppProbing) {
	v.value = val
	v.isSet = true
}

func (v NullableAppProbing) IsSet() bool {
	return v.isSet
}

func (v *NullableAppProbing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppProbing(val *AppProbing) *NullableAppProbing {
	return &NullableAppProbing{value: val, isSet: true}
}

func (v NullableAppProbing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppProbing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


