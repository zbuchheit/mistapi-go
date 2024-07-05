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

// checks if the AssetImport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetImport{}

// AssetImport struct for AssetImport
type AssetImport struct {
	Mac string `json:"mac"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AssetImport AssetImport

// NewAssetImport instantiates a new AssetImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetImport(mac string, name string) *AssetImport {
	this := AssetImport{}
	this.Mac = mac
	this.Name = name
	return &this
}

// NewAssetImportWithDefaults instantiates a new AssetImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetImportWithDefaults() *AssetImport {
	this := AssetImport{}
	return &this
}

// GetMac returns the Mac field value
func (o *AssetImport) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *AssetImport) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *AssetImport) SetMac(v string) {
	o.Mac = v
}

// GetName returns the Name field value
func (o *AssetImport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssetImport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AssetImport) SetName(v string) {
	o.Name = v
}

func (o AssetImport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetImport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mac"] = o.Mac
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetImport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mac",
		"name",
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

	varAssetImport := _AssetImport{}

	err = json.Unmarshal(data, &varAssetImport)

	if err != nil {
		return err
	}

	*o = AssetImport(varAssetImport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mac")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetImport struct {
	value *AssetImport
	isSet bool
}

func (v NullableAssetImport) Get() *AssetImport {
	return v.value
}

func (v *NullableAssetImport) Set(val *AssetImport) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetImport) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetImport(val *AssetImport) *NullableAssetImport {
	return &NullableAssetImport{value: val, isSet: true}
}

func (v NullableAssetImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


