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

// checks if the GatewayMatchingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayMatchingRule{}

// GatewayMatchingRule struct for GatewayMatchingRule
type GatewayMatchingRule struct {
	// additional CLI commands to append to the generated Junos config  **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	Name *string `json:"name,omitempty"`
	PortConfig *map[string]JunosPortConfig `json:"port_config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GatewayMatchingRule GatewayMatchingRule

// NewGatewayMatchingRule instantiates a new GatewayMatchingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMatchingRule() *GatewayMatchingRule {
	this := GatewayMatchingRule{}
	return &this
}

// NewGatewayMatchingRuleWithDefaults instantiates a new GatewayMatchingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMatchingRuleWithDefaults() *GatewayMatchingRule {
	this := GatewayMatchingRule{}
	return &this
}

// GetAdditionalConfigCmds returns the AdditionalConfigCmds field value if set, zero value otherwise.
func (o *GatewayMatchingRule) GetAdditionalConfigCmds() []string {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		var ret []string
		return ret
	}
	return o.AdditionalConfigCmds
}

// GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMatchingRule) GetAdditionalConfigCmdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		return nil, false
	}
	return o.AdditionalConfigCmds, true
}

// HasAdditionalConfigCmds returns a boolean if a field has been set.
func (o *GatewayMatchingRule) HasAdditionalConfigCmds() bool {
	if o != nil && !IsNil(o.AdditionalConfigCmds) {
		return true
	}

	return false
}

// SetAdditionalConfigCmds gets a reference to the given []string and assigns it to the AdditionalConfigCmds field.
func (o *GatewayMatchingRule) SetAdditionalConfigCmds(v []string) {
	o.AdditionalConfigCmds = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GatewayMatchingRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMatchingRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GatewayMatchingRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GatewayMatchingRule) SetName(v string) {
	o.Name = &v
}

// GetPortConfig returns the PortConfig field value if set, zero value otherwise.
func (o *GatewayMatchingRule) GetPortConfig() map[string]JunosPortConfig {
	if o == nil || IsNil(o.PortConfig) {
		var ret map[string]JunosPortConfig
		return ret
	}
	return *o.PortConfig
}

// GetPortConfigOk returns a tuple with the PortConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMatchingRule) GetPortConfigOk() (*map[string]JunosPortConfig, bool) {
	if o == nil || IsNil(o.PortConfig) {
		return nil, false
	}
	return o.PortConfig, true
}

// HasPortConfig returns a boolean if a field has been set.
func (o *GatewayMatchingRule) HasPortConfig() bool {
	if o != nil && !IsNil(o.PortConfig) {
		return true
	}

	return false
}

// SetPortConfig gets a reference to the given map[string]JunosPortConfig and assigns it to the PortConfig field.
func (o *GatewayMatchingRule) SetPortConfig(v map[string]JunosPortConfig) {
	o.PortConfig = &v
}

func (o GatewayMatchingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayMatchingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalConfigCmds) {
		toSerialize["additional_config_cmds"] = o.AdditionalConfigCmds
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PortConfig) {
		toSerialize["port_config"] = o.PortConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GatewayMatchingRule) UnmarshalJSON(data []byte) (err error) {
	varGatewayMatchingRule := _GatewayMatchingRule{}

	err = json.Unmarshal(data, &varGatewayMatchingRule)

	if err != nil {
		return err
	}

	*o = GatewayMatchingRule(varGatewayMatchingRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_config_cmds")
		delete(additionalProperties, "name")
		delete(additionalProperties, "port_config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGatewayMatchingRule struct {
	value *GatewayMatchingRule
	isSet bool
}

func (v NullableGatewayMatchingRule) Get() *GatewayMatchingRule {
	return v.value
}

func (v *NullableGatewayMatchingRule) Set(val *GatewayMatchingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMatchingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMatchingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMatchingRule(val *GatewayMatchingRule) *NullableGatewayMatchingRule {
	return &NullableGatewayMatchingRule{value: val, isSet: true}
}

func (v NullableGatewayMatchingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMatchingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


