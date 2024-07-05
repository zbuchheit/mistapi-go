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

// checks if the MxedgeDasCoaServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MxedgeDasCoaServer{}

// MxedgeDasCoaServer struct for MxedgeDasCoaServer
type MxedgeDasCoaServer struct {
	// whether to disable Event-Timestamp Check
	DisableEventTimestampCheck *bool `json:"disable_event_timestamp_check,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// this server configured to send CoA|DM to mist edges
	Host *string `json:"host,omitempty"`
	// mist edges will allow this host on this port
	Port *int32 `json:"port,omitempty"`
	Secret *string `json:"secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MxedgeDasCoaServer MxedgeDasCoaServer

// NewMxedgeDasCoaServer instantiates a new MxedgeDasCoaServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMxedgeDasCoaServer() *MxedgeDasCoaServer {
	this := MxedgeDasCoaServer{}
	var disableEventTimestampCheck bool = false
	this.DisableEventTimestampCheck = &disableEventTimestampCheck
	var port int32 = 3799
	this.Port = &port
	return &this
}

// NewMxedgeDasCoaServerWithDefaults instantiates a new MxedgeDasCoaServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMxedgeDasCoaServerWithDefaults() *MxedgeDasCoaServer {
	this := MxedgeDasCoaServer{}
	var disableEventTimestampCheck bool = false
	this.DisableEventTimestampCheck = &disableEventTimestampCheck
	var port int32 = 3799
	this.Port = &port
	return &this
}

// GetDisableEventTimestampCheck returns the DisableEventTimestampCheck field value if set, zero value otherwise.
func (o *MxedgeDasCoaServer) GetDisableEventTimestampCheck() bool {
	if o == nil || IsNil(o.DisableEventTimestampCheck) {
		var ret bool
		return ret
	}
	return *o.DisableEventTimestampCheck
}

// GetDisableEventTimestampCheckOk returns a tuple with the DisableEventTimestampCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeDasCoaServer) GetDisableEventTimestampCheckOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableEventTimestampCheck) {
		return nil, false
	}
	return o.DisableEventTimestampCheck, true
}

// HasDisableEventTimestampCheck returns a boolean if a field has been set.
func (o *MxedgeDasCoaServer) HasDisableEventTimestampCheck() bool {
	if o != nil && !IsNil(o.DisableEventTimestampCheck) {
		return true
	}

	return false
}

// SetDisableEventTimestampCheck gets a reference to the given bool and assigns it to the DisableEventTimestampCheck field.
func (o *MxedgeDasCoaServer) SetDisableEventTimestampCheck(v bool) {
	o.DisableEventTimestampCheck = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MxedgeDasCoaServer) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeDasCoaServer) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MxedgeDasCoaServer) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MxedgeDasCoaServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *MxedgeDasCoaServer) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeDasCoaServer) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *MxedgeDasCoaServer) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *MxedgeDasCoaServer) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *MxedgeDasCoaServer) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeDasCoaServer) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *MxedgeDasCoaServer) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *MxedgeDasCoaServer) SetPort(v int32) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *MxedgeDasCoaServer) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MxedgeDasCoaServer) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *MxedgeDasCoaServer) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *MxedgeDasCoaServer) SetSecret(v string) {
	o.Secret = &v
}

func (o MxedgeDasCoaServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MxedgeDasCoaServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisableEventTimestampCheck) {
		toSerialize["disable_event_timestamp_check"] = o.DisableEventTimestampCheck
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MxedgeDasCoaServer) UnmarshalJSON(data []byte) (err error) {
	varMxedgeDasCoaServer := _MxedgeDasCoaServer{}

	err = json.Unmarshal(data, &varMxedgeDasCoaServer)

	if err != nil {
		return err
	}

	*o = MxedgeDasCoaServer(varMxedgeDasCoaServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disable_event_timestamp_check")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMxedgeDasCoaServer struct {
	value *MxedgeDasCoaServer
	isSet bool
}

func (v NullableMxedgeDasCoaServer) Get() *MxedgeDasCoaServer {
	return v.value
}

func (v *NullableMxedgeDasCoaServer) Set(val *MxedgeDasCoaServer) {
	v.value = val
	v.isSet = true
}

func (v NullableMxedgeDasCoaServer) IsSet() bool {
	return v.isSet
}

func (v *NullableMxedgeDasCoaServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxedgeDasCoaServer(val *MxedgeDasCoaServer) *NullableMxedgeDasCoaServer {
	return &NullableMxedgeDasCoaServer{value: val, isSet: true}
}

func (v NullableMxedgeDasCoaServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxedgeDasCoaServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


