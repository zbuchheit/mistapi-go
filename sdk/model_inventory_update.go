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

// checks if the InventoryUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryUpdate{}

// InventoryUpdate struct for InventoryUpdate
type InventoryUpdate struct {
	// if `op`==`assign`, a **cloud-ready** switch/gateway will be managed/configured by Mist by default, this disabled the behavior
	DisableAutoConfig *bool `json:"disable_auto_config,omitempty"`
	// if `op`==`assign`, `op`==`unassign`, `op`==`upgrade_to_mist`or `op`==`downgrade_to_jsi` , list of MAC, e.g. [\"5c5b350e0001\"]
	Macs []string `json:"macs,omitempty"`
	// if `op`==`assign`, an **adopted** switch/gateway will not be managed/configured by Mist by default, this enables the behavior
	Managed *bool `json:"managed,omitempty"`
	// if `op`==`assign`, if true, treat site assignment against an already assigned AP as error
	NoReassign *bool `json:"no_reassign,omitempty"`
	Op InventoryUpdateOperation `json:"op"`
	// if `op`==`delete`, list of serial numbers, e.g. [\"FXLH2015150025\"]
	Serials []string `json:"serials,omitempty"`
	// if `op`==`assign`, target site id
	SiteId *string `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryUpdate InventoryUpdate

// NewInventoryUpdate instantiates a new InventoryUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryUpdate(op InventoryUpdateOperation) *InventoryUpdate {
	this := InventoryUpdate{}
	var disableAutoConfig bool = false
	this.DisableAutoConfig = &disableAutoConfig
	var managed bool = false
	this.Managed = &managed
	this.Op = op
	return &this
}

// NewInventoryUpdateWithDefaults instantiates a new InventoryUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryUpdateWithDefaults() *InventoryUpdate {
	this := InventoryUpdate{}
	var disableAutoConfig bool = false
	this.DisableAutoConfig = &disableAutoConfig
	var managed bool = false
	this.Managed = &managed
	return &this
}

// GetDisableAutoConfig returns the DisableAutoConfig field value if set, zero value otherwise.
func (o *InventoryUpdate) GetDisableAutoConfig() bool {
	if o == nil || IsNil(o.DisableAutoConfig) {
		var ret bool
		return ret
	}
	return *o.DisableAutoConfig
}

// GetDisableAutoConfigOk returns a tuple with the DisableAutoConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetDisableAutoConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableAutoConfig) {
		return nil, false
	}
	return o.DisableAutoConfig, true
}

// HasDisableAutoConfig returns a boolean if a field has been set.
func (o *InventoryUpdate) HasDisableAutoConfig() bool {
	if o != nil && !IsNil(o.DisableAutoConfig) {
		return true
	}

	return false
}

// SetDisableAutoConfig gets a reference to the given bool and assigns it to the DisableAutoConfig field.
func (o *InventoryUpdate) SetDisableAutoConfig(v bool) {
	o.DisableAutoConfig = &v
}

// GetMacs returns the Macs field value if set, zero value otherwise.
func (o *InventoryUpdate) GetMacs() []string {
	if o == nil || IsNil(o.Macs) {
		var ret []string
		return ret
	}
	return o.Macs
}

// GetMacsOk returns a tuple with the Macs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetMacsOk() ([]string, bool) {
	if o == nil || IsNil(o.Macs) {
		return nil, false
	}
	return o.Macs, true
}

// HasMacs returns a boolean if a field has been set.
func (o *InventoryUpdate) HasMacs() bool {
	if o != nil && !IsNil(o.Macs) {
		return true
	}

	return false
}

// SetMacs gets a reference to the given []string and assigns it to the Macs field.
func (o *InventoryUpdate) SetMacs(v []string) {
	o.Macs = v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *InventoryUpdate) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *InventoryUpdate) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *InventoryUpdate) SetManaged(v bool) {
	o.Managed = &v
}

// GetNoReassign returns the NoReassign field value if set, zero value otherwise.
func (o *InventoryUpdate) GetNoReassign() bool {
	if o == nil || IsNil(o.NoReassign) {
		var ret bool
		return ret
	}
	return *o.NoReassign
}

// GetNoReassignOk returns a tuple with the NoReassign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetNoReassignOk() (*bool, bool) {
	if o == nil || IsNil(o.NoReassign) {
		return nil, false
	}
	return o.NoReassign, true
}

// HasNoReassign returns a boolean if a field has been set.
func (o *InventoryUpdate) HasNoReassign() bool {
	if o != nil && !IsNil(o.NoReassign) {
		return true
	}

	return false
}

// SetNoReassign gets a reference to the given bool and assigns it to the NoReassign field.
func (o *InventoryUpdate) SetNoReassign(v bool) {
	o.NoReassign = &v
}

// GetOp returns the Op field value
func (o *InventoryUpdate) GetOp() InventoryUpdateOperation {
	if o == nil {
		var ret InventoryUpdateOperation
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetOpOk() (*InventoryUpdateOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *InventoryUpdate) SetOp(v InventoryUpdateOperation) {
	o.Op = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InventoryUpdate) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InventoryUpdate) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InventoryUpdate) SetSerials(v []string) {
	o.Serials = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *InventoryUpdate) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryUpdate) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *InventoryUpdate) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *InventoryUpdate) SetSiteId(v string) {
	o.SiteId = &v
}

func (o InventoryUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisableAutoConfig) {
		toSerialize["disable_auto_config"] = o.DisableAutoConfig
	}
	if !IsNil(o.Macs) {
		toSerialize["macs"] = o.Macs
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.NoReassign) {
		toSerialize["no_reassign"] = o.NoReassign
	}
	toSerialize["op"] = o.Op
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
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

	varInventoryUpdate := _InventoryUpdate{}

	err = json.Unmarshal(data, &varInventoryUpdate)

	if err != nil {
		return err
	}

	*o = InventoryUpdate(varInventoryUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disable_auto_config")
		delete(additionalProperties, "macs")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "no_reassign")
		delete(additionalProperties, "op")
		delete(additionalProperties, "serials")
		delete(additionalProperties, "site_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryUpdate struct {
	value *InventoryUpdate
	isSet bool
}

func (v NullableInventoryUpdate) Get() *InventoryUpdate {
	return v.value
}

func (v *NullableInventoryUpdate) Set(val *InventoryUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUpdate(val *InventoryUpdate) *NullableInventoryUpdate {
	return &NullableInventoryUpdate{value: val, isSet: true}
}

func (v NullableInventoryUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


