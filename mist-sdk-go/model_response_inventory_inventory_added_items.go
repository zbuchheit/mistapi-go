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

// checks if the ResponseInventoryInventoryAddedItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseInventoryInventoryAddedItems{}

// ResponseInventoryInventoryAddedItems struct for ResponseInventoryInventoryAddedItems
type ResponseInventoryInventoryAddedItems struct {
	Mac string `json:"mac"`
	Magic string `json:"magic"`
	Model string `json:"model"`
	Serial string `json:"serial"`
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ResponseInventoryInventoryAddedItems ResponseInventoryInventoryAddedItems

// NewResponseInventoryInventoryAddedItems instantiates a new ResponseInventoryInventoryAddedItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseInventoryInventoryAddedItems(mac string, magic string, model string, serial string, type_ string) *ResponseInventoryInventoryAddedItems {
	this := ResponseInventoryInventoryAddedItems{}
	this.Mac = mac
	this.Magic = magic
	this.Model = model
	this.Serial = serial
	this.Type = type_
	return &this
}

// NewResponseInventoryInventoryAddedItemsWithDefaults instantiates a new ResponseInventoryInventoryAddedItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseInventoryInventoryAddedItemsWithDefaults() *ResponseInventoryInventoryAddedItems {
	this := ResponseInventoryInventoryAddedItems{}
	return &this
}

// GetMac returns the Mac field value
func (o *ResponseInventoryInventoryAddedItems) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *ResponseInventoryInventoryAddedItems) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *ResponseInventoryInventoryAddedItems) SetMac(v string) {
	o.Mac = v
}

// GetMagic returns the Magic field value
func (o *ResponseInventoryInventoryAddedItems) GetMagic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Magic
}

// GetMagicOk returns a tuple with the Magic field value
// and a boolean to check if the value has been set.
func (o *ResponseInventoryInventoryAddedItems) GetMagicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Magic, true
}

// SetMagic sets field value
func (o *ResponseInventoryInventoryAddedItems) SetMagic(v string) {
	o.Magic = v
}

// GetModel returns the Model field value
func (o *ResponseInventoryInventoryAddedItems) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *ResponseInventoryInventoryAddedItems) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *ResponseInventoryInventoryAddedItems) SetModel(v string) {
	o.Model = v
}

// GetSerial returns the Serial field value
func (o *ResponseInventoryInventoryAddedItems) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *ResponseInventoryInventoryAddedItems) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *ResponseInventoryInventoryAddedItems) SetSerial(v string) {
	o.Serial = v
}

// GetType returns the Type field value
func (o *ResponseInventoryInventoryAddedItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResponseInventoryInventoryAddedItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResponseInventoryInventoryAddedItems) SetType(v string) {
	o.Type = v
}

func (o ResponseInventoryInventoryAddedItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseInventoryInventoryAddedItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mac"] = o.Mac
	toSerialize["magic"] = o.Magic
	toSerialize["model"] = o.Model
	toSerialize["serial"] = o.Serial
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseInventoryInventoryAddedItems) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mac",
		"magic",
		"model",
		"serial",
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

	varResponseInventoryInventoryAddedItems := _ResponseInventoryInventoryAddedItems{}

	err = json.Unmarshal(data, &varResponseInventoryInventoryAddedItems)

	if err != nil {
		return err
	}

	*o = ResponseInventoryInventoryAddedItems(varResponseInventoryInventoryAddedItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "magic")
		delete(additionalProperties, "model")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseInventoryInventoryAddedItems struct {
	value *ResponseInventoryInventoryAddedItems
	isSet bool
}

func (v NullableResponseInventoryInventoryAddedItems) Get() *ResponseInventoryInventoryAddedItems {
	return v.value
}

func (v *NullableResponseInventoryInventoryAddedItems) Set(val *ResponseInventoryInventoryAddedItems) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseInventoryInventoryAddedItems) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseInventoryInventoryAddedItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseInventoryInventoryAddedItems(val *ResponseInventoryInventoryAddedItems) *NullableResponseInventoryInventoryAddedItems {
	return &NullableResponseInventoryInventoryAddedItems{value: val, isSet: true}
}

func (v NullableResponseInventoryInventoryAddedItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseInventoryInventoryAddedItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


