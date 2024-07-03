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

// checks if the TacacsAcctServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TacacsAcctServer{}

// TacacsAcctServer struct for TacacsAcctServer
type TacacsAcctServer struct {
	Host *string `json:"host,omitempty"`
	Port *string `json:"port,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Timeout *int32 `json:"timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TacacsAcctServer TacacsAcctServer

// NewTacacsAcctServer instantiates a new TacacsAcctServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTacacsAcctServer() *TacacsAcctServer {
	this := TacacsAcctServer{}
	var timeout int32 = 10
	this.Timeout = &timeout
	return &this
}

// NewTacacsAcctServerWithDefaults instantiates a new TacacsAcctServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTacacsAcctServerWithDefaults() *TacacsAcctServer {
	this := TacacsAcctServer{}
	var timeout int32 = 10
	this.Timeout = &timeout
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *TacacsAcctServer) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAcctServer) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *TacacsAcctServer) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *TacacsAcctServer) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TacacsAcctServer) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAcctServer) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TacacsAcctServer) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *TacacsAcctServer) SetPort(v string) {
	o.Port = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TacacsAcctServer) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAcctServer) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TacacsAcctServer) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TacacsAcctServer) SetSecret(v string) {
	o.Secret = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *TacacsAcctServer) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TacacsAcctServer) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *TacacsAcctServer) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *TacacsAcctServer) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o TacacsAcctServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TacacsAcctServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TacacsAcctServer) UnmarshalJSON(data []byte) (err error) {
	varTacacsAcctServer := _TacacsAcctServer{}

	err = json.Unmarshal(data, &varTacacsAcctServer)

	if err != nil {
		return err
	}

	*o = TacacsAcctServer(varTacacsAcctServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTacacsAcctServer struct {
	value *TacacsAcctServer
	isSet bool
}

func (v NullableTacacsAcctServer) Get() *TacacsAcctServer {
	return v.value
}

func (v *NullableTacacsAcctServer) Set(val *TacacsAcctServer) {
	v.value = val
	v.isSet = true
}

func (v NullableTacacsAcctServer) IsSet() bool {
	return v.isSet
}

func (v *NullableTacacsAcctServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTacacsAcctServer(val *TacacsAcctServer) *NullableTacacsAcctServer {
	return &NullableTacacsAcctServer{value: val, isSet: true}
}

func (v NullableTacacsAcctServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTacacsAcctServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


