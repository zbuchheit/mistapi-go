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

// checks if the RfDiag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RfDiag{}

// RfDiag RF Diag
type RfDiag struct {
	// recording length in seconds, max is 180. Default value is also 180.
	Duration *int32 `json:"duration,omitempty"`
	// if `type`==`client` or `asset`, mac of the device
	Mac *string `json:"mac,omitempty"`
	// name of the recording, the name of the sdk client would be a good default choice
	Name string `json:"name"`
	// if `type`==`sdkclient`, sdkclient_id of this recording
	SdkclientId *string `json:"sdkclient_id,omitempty"`
	Type RfClientType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _RfDiag RfDiag

// NewRfDiag instantiates a new RfDiag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRfDiag(name string, type_ RfClientType) *RfDiag {
	this := RfDiag{}
	var duration int32 = 180
	this.Duration = &duration
	this.Name = name
	this.Type = type_
	return &this
}

// NewRfDiagWithDefaults instantiates a new RfDiag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRfDiagWithDefaults() *RfDiag {
	this := RfDiag{}
	var duration int32 = 180
	this.Duration = &duration
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *RfDiag) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiag) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *RfDiag) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *RfDiag) SetDuration(v int32) {
	o.Duration = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *RfDiag) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiag) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *RfDiag) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *RfDiag) SetMac(v string) {
	o.Mac = &v
}

// GetName returns the Name field value
func (o *RfDiag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RfDiag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RfDiag) SetName(v string) {
	o.Name = v
}

// GetSdkclientId returns the SdkclientId field value if set, zero value otherwise.
func (o *RfDiag) GetSdkclientId() string {
	if o == nil || IsNil(o.SdkclientId) {
		var ret string
		return ret
	}
	return *o.SdkclientId
}

// GetSdkclientIdOk returns a tuple with the SdkclientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiag) GetSdkclientIdOk() (*string, bool) {
	if o == nil || IsNil(o.SdkclientId) {
		return nil, false
	}
	return o.SdkclientId, true
}

// HasSdkclientId returns a boolean if a field has been set.
func (o *RfDiag) HasSdkclientId() bool {
	if o != nil && !IsNil(o.SdkclientId) {
		return true
	}

	return false
}

// SetSdkclientId gets a reference to the given string and assigns it to the SdkclientId field.
func (o *RfDiag) SetSdkclientId(v string) {
	o.SdkclientId = &v
}

// GetType returns the Type field value
func (o *RfDiag) GetType() RfClientType {
	if o == nil {
		var ret RfClientType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RfDiag) GetTypeOk() (*RfClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RfDiag) SetType(v RfClientType) {
	o.Type = v
}

func (o RfDiag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RfDiag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SdkclientId) {
		toSerialize["sdkclient_id"] = o.SdkclientId
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RfDiag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varRfDiag := _RfDiag{}

	err = json.Unmarshal(data, &varRfDiag)

	if err != nil {
		return err
	}

	*o = RfDiag(varRfDiag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sdkclient_id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRfDiag struct {
	value *RfDiag
	isSet bool
}

func (v NullableRfDiag) Get() *RfDiag {
	return v.value
}

func (v *NullableRfDiag) Set(val *RfDiag) {
	v.value = val
	v.isSet = true
}

func (v NullableRfDiag) IsSet() bool {
	return v.isSet
}

func (v *NullableRfDiag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRfDiag(val *RfDiag) *NullableRfDiag {
	return &NullableRfDiag{value: val, isSet: true}
}

func (v NullableRfDiag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRfDiag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


