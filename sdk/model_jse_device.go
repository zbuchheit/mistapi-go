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

// checks if the JseDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JseDevice{}

// JseDevice struct for JseDevice
type JseDevice struct {
	// when available
	ExtIp *string `json:"ext_ip,omitempty"`
	LastSeen *int32 `json:"last_seen,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Model *string `json:"model,omitempty"`
	Serial *string `json:"serial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JseDevice JseDevice

// NewJseDevice instantiates a new JseDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJseDevice() *JseDevice {
	this := JseDevice{}
	return &this
}

// NewJseDeviceWithDefaults instantiates a new JseDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJseDeviceWithDefaults() *JseDevice {
	this := JseDevice{}
	return &this
}

// GetExtIp returns the ExtIp field value if set, zero value otherwise.
func (o *JseDevice) GetExtIp() string {
	if o == nil || IsNil(o.ExtIp) {
		var ret string
		return ret
	}
	return *o.ExtIp
}

// GetExtIpOk returns a tuple with the ExtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseDevice) GetExtIpOk() (*string, bool) {
	if o == nil || IsNil(o.ExtIp) {
		return nil, false
	}
	return o.ExtIp, true
}

// HasExtIp returns a boolean if a field has been set.
func (o *JseDevice) HasExtIp() bool {
	if o != nil && !IsNil(o.ExtIp) {
		return true
	}

	return false
}

// SetExtIp gets a reference to the given string and assigns it to the ExtIp field.
func (o *JseDevice) SetExtIp(v string) {
	o.ExtIp = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *JseDevice) GetLastSeen() int32 {
	if o == nil || IsNil(o.LastSeen) {
		var ret int32
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseDevice) GetLastSeenOk() (*int32, bool) {
	if o == nil || IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *JseDevice) HasLastSeen() bool {
	if o != nil && !IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given int32 and assigns it to the LastSeen field.
func (o *JseDevice) SetLastSeen(v int32) {
	o.LastSeen = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *JseDevice) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseDevice) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *JseDevice) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *JseDevice) SetMac(v string) {
	o.Mac = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *JseDevice) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseDevice) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *JseDevice) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *JseDevice) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *JseDevice) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseDevice) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *JseDevice) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *JseDevice) SetSerial(v string) {
	o.Serial = &v
}

func (o JseDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JseDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtIp) {
		toSerialize["ext_ip"] = o.ExtIp
	}
	if !IsNil(o.LastSeen) {
		toSerialize["last_seen"] = o.LastSeen
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JseDevice) UnmarshalJSON(data []byte) (err error) {
	varJseDevice := _JseDevice{}

	err = json.Unmarshal(data, &varJseDevice)

	if err != nil {
		return err
	}

	*o = JseDevice(varJseDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ext_ip")
		delete(additionalProperties, "last_seen")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "model")
		delete(additionalProperties, "serial")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJseDevice struct {
	value *JseDevice
	isSet bool
}

func (v NullableJseDevice) Get() *JseDevice {
	return v.value
}

func (v *NullableJseDevice) Set(val *JseDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableJseDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableJseDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJseDevice(val *JseDevice) *NullableJseDevice {
	return &NullableJseDevice{value: val, isSet: true}
}

func (v NullableJseDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJseDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


