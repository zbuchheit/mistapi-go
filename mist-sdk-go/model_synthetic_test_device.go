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

// checks if the SyntheticTestDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticTestDevice{}

// SyntheticTestDevice struct for SyntheticTestDevice
type SyntheticTestDevice struct {
	// if `type`==`dns`
	Hostname *string `json:"hostname,omitempty"`
	// if `type`==`arp`
	Ip *string `json:"ip,omitempty"`
	// if `type`==`radius`
	Password *string `json:"password,omitempty"`
	// if `type`==`ssr`
	PortId *string `json:"port_id,omitempty"`
	Type SyntheticTestType `json:"type"`
	// if `type`==`curl`
	Url *string `json:"url,omitempty"`
	// if `type`==`radius`
	Username *string `json:"username,omitempty"`
	// required for AP
	VlanId *int32 `json:"vlan_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyntheticTestDevice SyntheticTestDevice

// NewSyntheticTestDevice instantiates a new SyntheticTestDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticTestDevice(type_ SyntheticTestType) *SyntheticTestDevice {
	this := SyntheticTestDevice{}
	this.Type = type_
	return &this
}

// NewSyntheticTestDeviceWithDefaults instantiates a new SyntheticTestDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticTestDeviceWithDefaults() *SyntheticTestDevice {
	this := SyntheticTestDevice{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SyntheticTestDevice) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SyntheticTestDevice) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SyntheticTestDevice) SetHostname(v string) {
	o.Hostname = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *SyntheticTestDevice) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *SyntheticTestDevice) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *SyntheticTestDevice) SetIp(v string) {
	o.Ip = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SyntheticTestDevice) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SyntheticTestDevice) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SyntheticTestDevice) SetPassword(v string) {
	o.Password = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *SyntheticTestDevice) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *SyntheticTestDevice) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *SyntheticTestDevice) SetPortId(v string) {
	o.PortId = &v
}

// GetType returns the Type field value
func (o *SyntheticTestDevice) GetType() SyntheticTestType {
	if o == nil {
		var ret SyntheticTestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetTypeOk() (*SyntheticTestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticTestDevice) SetType(v SyntheticTestType) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticTestDevice) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticTestDevice) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticTestDevice) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SyntheticTestDevice) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SyntheticTestDevice) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SyntheticTestDevice) SetUsername(v string) {
	o.Username = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *SyntheticTestDevice) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticTestDevice) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *SyntheticTestDevice) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *SyntheticTestDevice) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o SyntheticTestDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticTestDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SyntheticTestDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varSyntheticTestDevice := _SyntheticTestDevice{}

	err = json.Unmarshal(data, &varSyntheticTestDevice)

	if err != nil {
		return err
	}

	*o = SyntheticTestDevice(varSyntheticTestDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "password")
		delete(additionalProperties, "port_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		delete(additionalProperties, "username")
		delete(additionalProperties, "vlan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSyntheticTestDevice struct {
	value *SyntheticTestDevice
	isSet bool
}

func (v NullableSyntheticTestDevice) Get() *SyntheticTestDevice {
	return v.value
}

func (v *NullableSyntheticTestDevice) Set(val *SyntheticTestDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticTestDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticTestDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticTestDevice(val *SyntheticTestDevice) *NullableSyntheticTestDevice {
	return &NullableSyntheticTestDevice{value: val, isSet: true}
}

func (v NullableSyntheticTestDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticTestDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


