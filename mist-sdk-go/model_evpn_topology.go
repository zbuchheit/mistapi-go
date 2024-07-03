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
	"fmt"
)

// checks if the EvpnTopology type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvpnTopology{}

// EvpnTopology struct for EvpnTopology
type EvpnTopology struct {
	Config *ConfigSwitch `json:"config,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Overwrite *bool `json:"overwrite,omitempty"`
	// Property key is the pod number
	PodNames *map[string]string `json:"pod_names,omitempty"`
	Switches []EvpnTopologySwitch `json:"switches"`
	AdditionalProperties map[string]interface{}
}

type _EvpnTopology EvpnTopology

// NewEvpnTopology instantiates a new EvpnTopology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvpnTopology(switches []EvpnTopologySwitch) *EvpnTopology {
	this := EvpnTopology{}
	this.Switches = switches
	return &this
}

// NewEvpnTopologyWithDefaults instantiates a new EvpnTopology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvpnTopologyWithDefaults() *EvpnTopology {
	this := EvpnTopology{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *EvpnTopology) GetConfig() ConfigSwitch {
	if o == nil || IsNil(o.Config) {
		var ret ConfigSwitch
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvpnTopology) GetConfigOk() (*ConfigSwitch, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *EvpnTopology) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConfigSwitch and assigns it to the Config field.
func (o *EvpnTopology) SetConfig(v ConfigSwitch) {
	o.Config = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EvpnTopology) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvpnTopology) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EvpnTopology) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EvpnTopology) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EvpnTopology) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvpnTopology) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EvpnTopology) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EvpnTopology) SetName(v string) {
	o.Name = &v
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise.
func (o *EvpnTopology) GetOverwrite() bool {
	if o == nil || IsNil(o.Overwrite) {
		var ret bool
		return ret
	}
	return *o.Overwrite
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvpnTopology) GetOverwriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Overwrite) {
		return nil, false
	}
	return o.Overwrite, true
}

// HasOverwrite returns a boolean if a field has been set.
func (o *EvpnTopology) HasOverwrite() bool {
	if o != nil && !IsNil(o.Overwrite) {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given bool and assigns it to the Overwrite field.
func (o *EvpnTopology) SetOverwrite(v bool) {
	o.Overwrite = &v
}

// GetPodNames returns the PodNames field value if set, zero value otherwise.
func (o *EvpnTopology) GetPodNames() map[string]string {
	if o == nil || IsNil(o.PodNames) {
		var ret map[string]string
		return ret
	}
	return *o.PodNames
}

// GetPodNamesOk returns a tuple with the PodNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvpnTopology) GetPodNamesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PodNames) {
		return nil, false
	}
	return o.PodNames, true
}

// HasPodNames returns a boolean if a field has been set.
func (o *EvpnTopology) HasPodNames() bool {
	if o != nil && !IsNil(o.PodNames) {
		return true
	}

	return false
}

// SetPodNames gets a reference to the given map[string]string and assigns it to the PodNames field.
func (o *EvpnTopology) SetPodNames(v map[string]string) {
	o.PodNames = &v
}

// GetSwitches returns the Switches field value
func (o *EvpnTopology) GetSwitches() []EvpnTopologySwitch {
	if o == nil {
		var ret []EvpnTopologySwitch
		return ret
	}

	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value
// and a boolean to check if the value has been set.
func (o *EvpnTopology) GetSwitchesOk() ([]EvpnTopologySwitch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Switches, true
}

// SetSwitches sets field value
func (o *EvpnTopology) SetSwitches(v []EvpnTopologySwitch) {
	o.Switches = v
}

func (o EvpnTopology) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvpnTopology) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Overwrite) {
		toSerialize["overwrite"] = o.Overwrite
	}
	if !IsNil(o.PodNames) {
		toSerialize["pod_names"] = o.PodNames
	}
	toSerialize["switches"] = o.Switches

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EvpnTopology) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"switches",
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

	varEvpnTopology := _EvpnTopology{}

	err = json.Unmarshal(data, &varEvpnTopology)

	if err != nil {
		return err
	}

	*o = EvpnTopology(varEvpnTopology)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "overwrite")
		delete(additionalProperties, "pod_names")
		delete(additionalProperties, "switches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEvpnTopology struct {
	value *EvpnTopology
	isSet bool
}

func (v NullableEvpnTopology) Get() *EvpnTopology {
	return v.value
}

func (v *NullableEvpnTopology) Set(val *EvpnTopology) {
	v.value = val
	v.isSet = true
}

func (v NullableEvpnTopology) IsSet() bool {
	return v.isSet
}

func (v *NullableEvpnTopology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvpnTopology(val *EvpnTopology) *NullableEvpnTopology {
	return &NullableEvpnTopology{value: val, isSet: true}
}

func (v NullableEvpnTopology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvpnTopology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


