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

// checks if the SwitchMatchingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchMatchingRule{}

// SwitchMatchingRule property key define the type of matching, value is the string to match. e.g: `match_name[0:3]`, `match_name[2:6]`, `match_model`,  `match_model[0-6]`
type SwitchMatchingRule struct {
	// additional CLI commands to append to the generated Junos config  **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	// role to match
	MatchRole *string `json:"match_role,omitempty"`
	Name *string `json:"name,omitempty"`
	// Propery key is the interface name or interface range
	PortConfig *map[string]JunosPortConfig `json:"port_config,omitempty"`
	// Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.
	PortMirroring *map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
	SwitchMgmt *ConfigSwitch `json:"switch_mgmt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwitchMatchingRule SwitchMatchingRule

// NewSwitchMatchingRule instantiates a new SwitchMatchingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchMatchingRule() *SwitchMatchingRule {
	this := SwitchMatchingRule{}
	return &this
}

// NewSwitchMatchingRuleWithDefaults instantiates a new SwitchMatchingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchMatchingRuleWithDefaults() *SwitchMatchingRule {
	this := SwitchMatchingRule{}
	return &this
}

// GetAdditionalConfigCmds returns the AdditionalConfigCmds field value if set, zero value otherwise.
func (o *SwitchMatchingRule) GetAdditionalConfigCmds() []string {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		var ret []string
		return ret
	}
	return o.AdditionalConfigCmds
}

// GetAdditionalConfigCmdsOk returns a tuple with the AdditionalConfigCmds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMatchingRule) GetAdditionalConfigCmdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalConfigCmds) {
		return nil, false
	}
	return o.AdditionalConfigCmds, true
}

// HasAdditionalConfigCmds returns a boolean if a field has been set.
func (o *SwitchMatchingRule) HasAdditionalConfigCmds() bool {
	if o != nil && !IsNil(o.AdditionalConfigCmds) {
		return true
	}

	return false
}

// SetAdditionalConfigCmds gets a reference to the given []string and assigns it to the AdditionalConfigCmds field.
func (o *SwitchMatchingRule) SetAdditionalConfigCmds(v []string) {
	o.AdditionalConfigCmds = v
}

// GetMatchRole returns the MatchRole field value if set, zero value otherwise.
func (o *SwitchMatchingRule) GetMatchRole() string {
	if o == nil || IsNil(o.MatchRole) {
		var ret string
		return ret
	}
	return *o.MatchRole
}

// GetMatchRoleOk returns a tuple with the MatchRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMatchingRule) GetMatchRoleOk() (*string, bool) {
	if o == nil || IsNil(o.MatchRole) {
		return nil, false
	}
	return o.MatchRole, true
}

// HasMatchRole returns a boolean if a field has been set.
func (o *SwitchMatchingRule) HasMatchRole() bool {
	if o != nil && !IsNil(o.MatchRole) {
		return true
	}

	return false
}

// SetMatchRole gets a reference to the given string and assigns it to the MatchRole field.
func (o *SwitchMatchingRule) SetMatchRole(v string) {
	o.MatchRole = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SwitchMatchingRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMatchingRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SwitchMatchingRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SwitchMatchingRule) SetName(v string) {
	o.Name = &v
}

// GetPortConfig returns the PortConfig field value if set, zero value otherwise.
func (o *SwitchMatchingRule) GetPortConfig() map[string]JunosPortConfig {
	if o == nil || IsNil(o.PortConfig) {
		var ret map[string]JunosPortConfig
		return ret
	}
	return *o.PortConfig
}

// GetPortConfigOk returns a tuple with the PortConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMatchingRule) GetPortConfigOk() (*map[string]JunosPortConfig, bool) {
	if o == nil || IsNil(o.PortConfig) {
		return nil, false
	}
	return o.PortConfig, true
}

// HasPortConfig returns a boolean if a field has been set.
func (o *SwitchMatchingRule) HasPortConfig() bool {
	if o != nil && !IsNil(o.PortConfig) {
		return true
	}

	return false
}

// SetPortConfig gets a reference to the given map[string]JunosPortConfig and assigns it to the PortConfig field.
func (o *SwitchMatchingRule) SetPortConfig(v map[string]JunosPortConfig) {
	o.PortConfig = &v
}

// GetPortMirroring returns the PortMirroring field value if set, zero value otherwise.
func (o *SwitchMatchingRule) GetPortMirroring() map[string]SwitchPortMirroringProperty {
	if o == nil || IsNil(o.PortMirroring) {
		var ret map[string]SwitchPortMirroringProperty
		return ret
	}
	return *o.PortMirroring
}

// GetPortMirroringOk returns a tuple with the PortMirroring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMatchingRule) GetPortMirroringOk() (*map[string]SwitchPortMirroringProperty, bool) {
	if o == nil || IsNil(o.PortMirroring) {
		return nil, false
	}
	return o.PortMirroring, true
}

// HasPortMirroring returns a boolean if a field has been set.
func (o *SwitchMatchingRule) HasPortMirroring() bool {
	if o != nil && !IsNil(o.PortMirroring) {
		return true
	}

	return false
}

// SetPortMirroring gets a reference to the given map[string]SwitchPortMirroringProperty and assigns it to the PortMirroring field.
func (o *SwitchMatchingRule) SetPortMirroring(v map[string]SwitchPortMirroringProperty) {
	o.PortMirroring = &v
}

// GetSwitchMgmt returns the SwitchMgmt field value if set, zero value otherwise.
func (o *SwitchMatchingRule) GetSwitchMgmt() ConfigSwitch {
	if o == nil || IsNil(o.SwitchMgmt) {
		var ret ConfigSwitch
		return ret
	}
	return *o.SwitchMgmt
}

// GetSwitchMgmtOk returns a tuple with the SwitchMgmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchMatchingRule) GetSwitchMgmtOk() (*ConfigSwitch, bool) {
	if o == nil || IsNil(o.SwitchMgmt) {
		return nil, false
	}
	return o.SwitchMgmt, true
}

// HasSwitchMgmt returns a boolean if a field has been set.
func (o *SwitchMatchingRule) HasSwitchMgmt() bool {
	if o != nil && !IsNil(o.SwitchMgmt) {
		return true
	}

	return false
}

// SetSwitchMgmt gets a reference to the given ConfigSwitch and assigns it to the SwitchMgmt field.
func (o *SwitchMatchingRule) SetSwitchMgmt(v ConfigSwitch) {
	o.SwitchMgmt = &v
}

func (o SwitchMatchingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchMatchingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalConfigCmds) {
		toSerialize["additional_config_cmds"] = o.AdditionalConfigCmds
	}
	if !IsNil(o.MatchRole) {
		toSerialize["match_role"] = o.MatchRole
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PortConfig) {
		toSerialize["port_config"] = o.PortConfig
	}
	if !IsNil(o.PortMirroring) {
		toSerialize["port_mirroring"] = o.PortMirroring
	}
	if !IsNil(o.SwitchMgmt) {
		toSerialize["switch_mgmt"] = o.SwitchMgmt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwitchMatchingRule) UnmarshalJSON(data []byte) (err error) {
	varSwitchMatchingRule := _SwitchMatchingRule{}

	err = json.Unmarshal(data, &varSwitchMatchingRule)

	if err != nil {
		return err
	}

	*o = SwitchMatchingRule(varSwitchMatchingRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additional_config_cmds")
		delete(additionalProperties, "match_role")
		delete(additionalProperties, "name")
		delete(additionalProperties, "port_config")
		delete(additionalProperties, "port_mirroring")
		delete(additionalProperties, "switch_mgmt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwitchMatchingRule struct {
	value *SwitchMatchingRule
	isSet bool
}

func (v NullableSwitchMatchingRule) Get() *SwitchMatchingRule {
	return v.value
}

func (v *NullableSwitchMatchingRule) Set(val *SwitchMatchingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchMatchingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchMatchingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchMatchingRule(val *SwitchMatchingRule) *NullableSwitchMatchingRule {
	return &NullableSwitchMatchingRule{value: val, isSet: true}
}

func (v NullableSwitchMatchingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchMatchingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


