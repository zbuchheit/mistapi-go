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

// checks if the RadiusAuthServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RadiusAuthServer{}

// RadiusAuthServer Authentication Server
type RadiusAuthServer struct {
	// ip / hostname of RADIUS server
	Host string `json:"host"`
	KeywrapEnabled *bool `json:"keywrap_enabled,omitempty"`
	KeywrapFormat *RadiusKeywrapFormat `json:"keywrap_format,omitempty"`
	KeywrapKek *string `json:"keywrap_kek,omitempty"`
	KeywrapMack *string `json:"keywrap_mack,omitempty"`
	// Auth port of RADIUS server
	Port int32 `json:"port"`
	// secret of RADIUS server
	Secret string `json:"secret"`
	AdditionalProperties map[string]interface{}
}

type _RadiusAuthServer RadiusAuthServer

// NewRadiusAuthServer instantiates a new RadiusAuthServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusAuthServer(host string, port int32, secret string) *RadiusAuthServer {
	this := RadiusAuthServer{}
	this.Host = host
	this.Port = port
	this.Secret = secret
	return &this
}

// NewRadiusAuthServerWithDefaults instantiates a new RadiusAuthServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusAuthServerWithDefaults() *RadiusAuthServer {
	this := RadiusAuthServer{}
	var port int32 = 1812
	this.Port = port
	return &this
}

// GetHost returns the Host field value
func (o *RadiusAuthServer) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *RadiusAuthServer) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *RadiusAuthServer) SetHost(v string) {
	o.Host = v
}

// GetKeywrapEnabled returns the KeywrapEnabled field value if set, zero value otherwise.
func (o *RadiusAuthServer) GetKeywrapEnabled() bool {
	if o == nil || IsNil(o.KeywrapEnabled) {
		var ret bool
		return ret
	}
	return *o.KeywrapEnabled
}

// GetKeywrapEnabledOk returns a tuple with the KeywrapEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthServer) GetKeywrapEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.KeywrapEnabled) {
		return nil, false
	}
	return o.KeywrapEnabled, true
}

// HasKeywrapEnabled returns a boolean if a field has been set.
func (o *RadiusAuthServer) HasKeywrapEnabled() bool {
	if o != nil && !IsNil(o.KeywrapEnabled) {
		return true
	}

	return false
}

// SetKeywrapEnabled gets a reference to the given bool and assigns it to the KeywrapEnabled field.
func (o *RadiusAuthServer) SetKeywrapEnabled(v bool) {
	o.KeywrapEnabled = &v
}

// GetKeywrapFormat returns the KeywrapFormat field value if set, zero value otherwise.
func (o *RadiusAuthServer) GetKeywrapFormat() RadiusKeywrapFormat {
	if o == nil || IsNil(o.KeywrapFormat) {
		var ret RadiusKeywrapFormat
		return ret
	}
	return *o.KeywrapFormat
}

// GetKeywrapFormatOk returns a tuple with the KeywrapFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthServer) GetKeywrapFormatOk() (*RadiusKeywrapFormat, bool) {
	if o == nil || IsNil(o.KeywrapFormat) {
		return nil, false
	}
	return o.KeywrapFormat, true
}

// HasKeywrapFormat returns a boolean if a field has been set.
func (o *RadiusAuthServer) HasKeywrapFormat() bool {
	if o != nil && !IsNil(o.KeywrapFormat) {
		return true
	}

	return false
}

// SetKeywrapFormat gets a reference to the given RadiusKeywrapFormat and assigns it to the KeywrapFormat field.
func (o *RadiusAuthServer) SetKeywrapFormat(v RadiusKeywrapFormat) {
	o.KeywrapFormat = &v
}

// GetKeywrapKek returns the KeywrapKek field value if set, zero value otherwise.
func (o *RadiusAuthServer) GetKeywrapKek() string {
	if o == nil || IsNil(o.KeywrapKek) {
		var ret string
		return ret
	}
	return *o.KeywrapKek
}

// GetKeywrapKekOk returns a tuple with the KeywrapKek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthServer) GetKeywrapKekOk() (*string, bool) {
	if o == nil || IsNil(o.KeywrapKek) {
		return nil, false
	}
	return o.KeywrapKek, true
}

// HasKeywrapKek returns a boolean if a field has been set.
func (o *RadiusAuthServer) HasKeywrapKek() bool {
	if o != nil && !IsNil(o.KeywrapKek) {
		return true
	}

	return false
}

// SetKeywrapKek gets a reference to the given string and assigns it to the KeywrapKek field.
func (o *RadiusAuthServer) SetKeywrapKek(v string) {
	o.KeywrapKek = &v
}

// GetKeywrapMack returns the KeywrapMack field value if set, zero value otherwise.
func (o *RadiusAuthServer) GetKeywrapMack() string {
	if o == nil || IsNil(o.KeywrapMack) {
		var ret string
		return ret
	}
	return *o.KeywrapMack
}

// GetKeywrapMackOk returns a tuple with the KeywrapMack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusAuthServer) GetKeywrapMackOk() (*string, bool) {
	if o == nil || IsNil(o.KeywrapMack) {
		return nil, false
	}
	return o.KeywrapMack, true
}

// HasKeywrapMack returns a boolean if a field has been set.
func (o *RadiusAuthServer) HasKeywrapMack() bool {
	if o != nil && !IsNil(o.KeywrapMack) {
		return true
	}

	return false
}

// SetKeywrapMack gets a reference to the given string and assigns it to the KeywrapMack field.
func (o *RadiusAuthServer) SetKeywrapMack(v string) {
	o.KeywrapMack = &v
}

// GetPort returns the Port field value
func (o *RadiusAuthServer) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *RadiusAuthServer) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *RadiusAuthServer) SetPort(v int32) {
	o.Port = v
}

// GetSecret returns the Secret field value
func (o *RadiusAuthServer) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *RadiusAuthServer) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *RadiusAuthServer) SetSecret(v string) {
	o.Secret = v
}

func (o RadiusAuthServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RadiusAuthServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	if !IsNil(o.KeywrapEnabled) {
		toSerialize["keywrap_enabled"] = o.KeywrapEnabled
	}
	if !IsNil(o.KeywrapFormat) {
		toSerialize["keywrap_format"] = o.KeywrapFormat
	}
	if !IsNil(o.KeywrapKek) {
		toSerialize["keywrap_kek"] = o.KeywrapKek
	}
	if !IsNil(o.KeywrapMack) {
		toSerialize["keywrap_mack"] = o.KeywrapMack
	}
	toSerialize["port"] = o.Port
	toSerialize["secret"] = o.Secret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RadiusAuthServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"port",
		"secret",
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

	varRadiusAuthServer := _RadiusAuthServer{}

	err = json.Unmarshal(data, &varRadiusAuthServer)

	if err != nil {
		return err
	}

	*o = RadiusAuthServer(varRadiusAuthServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "host")
		delete(additionalProperties, "keywrap_enabled")
		delete(additionalProperties, "keywrap_format")
		delete(additionalProperties, "keywrap_kek")
		delete(additionalProperties, "keywrap_mack")
		delete(additionalProperties, "port")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRadiusAuthServer struct {
	value *RadiusAuthServer
	isSet bool
}

func (v NullableRadiusAuthServer) Get() *RadiusAuthServer {
	return v.value
}

func (v *NullableRadiusAuthServer) Set(val *RadiusAuthServer) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusAuthServer) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusAuthServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusAuthServer(val *RadiusAuthServer) *NullableRadiusAuthServer {
	return &NullableRadiusAuthServer{value: val, isSet: true}
}

func (v NullableRadiusAuthServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusAuthServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


