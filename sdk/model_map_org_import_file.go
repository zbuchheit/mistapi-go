/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"os"
)

// checks if the MapOrgImportFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapOrgImportFile{}

// MapOrgImportFile struct for MapOrgImportFile
type MapOrgImportFile struct {
	// whether to auto assign device to deviceprofile by name
	AutoDeviceprofileAssignment *bool `json:"auto_deviceprofile_assignment,omitempty"`
	// csv file for ap name mapping, optional
	Csv **os.File `json:"csv,omitempty"`
	// ekahau or ibwave file
	File **os.File `json:"file,omitempty"`
	Json *MapOrgImportFileJson `json:"json,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MapOrgImportFile MapOrgImportFile

// NewMapOrgImportFile instantiates a new MapOrgImportFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapOrgImportFile() *MapOrgImportFile {
	this := MapOrgImportFile{}
	return &this
}

// NewMapOrgImportFileWithDefaults instantiates a new MapOrgImportFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapOrgImportFileWithDefaults() *MapOrgImportFile {
	this := MapOrgImportFile{}
	return &this
}

// GetAutoDeviceprofileAssignment returns the AutoDeviceprofileAssignment field value if set, zero value otherwise.
func (o *MapOrgImportFile) GetAutoDeviceprofileAssignment() bool {
	if o == nil || IsNil(o.AutoDeviceprofileAssignment) {
		var ret bool
		return ret
	}
	return *o.AutoDeviceprofileAssignment
}

// GetAutoDeviceprofileAssignmentOk returns a tuple with the AutoDeviceprofileAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapOrgImportFile) GetAutoDeviceprofileAssignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDeviceprofileAssignment) {
		return nil, false
	}
	return o.AutoDeviceprofileAssignment, true
}

// HasAutoDeviceprofileAssignment returns a boolean if a field has been set.
func (o *MapOrgImportFile) HasAutoDeviceprofileAssignment() bool {
	if o != nil && !IsNil(o.AutoDeviceprofileAssignment) {
		return true
	}

	return false
}

// SetAutoDeviceprofileAssignment gets a reference to the given bool and assigns it to the AutoDeviceprofileAssignment field.
func (o *MapOrgImportFile) SetAutoDeviceprofileAssignment(v bool) {
	o.AutoDeviceprofileAssignment = &v
}

// GetCsv returns the Csv field value if set, zero value otherwise.
func (o *MapOrgImportFile) GetCsv() *os.File {
	if o == nil || IsNil(o.Csv) {
		var ret *os.File
		return ret
	}
	return *o.Csv
}

// GetCsvOk returns a tuple with the Csv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapOrgImportFile) GetCsvOk() (**os.File, bool) {
	if o == nil || IsNil(o.Csv) {
		return nil, false
	}
	return o.Csv, true
}

// HasCsv returns a boolean if a field has been set.
func (o *MapOrgImportFile) HasCsv() bool {
	if o != nil && !IsNil(o.Csv) {
		return true
	}

	return false
}

// SetCsv gets a reference to the given *os.File and assigns it to the Csv field.
func (o *MapOrgImportFile) SetCsv(v *os.File) {
	o.Csv = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *MapOrgImportFile) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapOrgImportFile) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *MapOrgImportFile) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *MapOrgImportFile) SetFile(v *os.File) {
	o.File = &v
}

// GetJson returns the Json field value if set, zero value otherwise.
func (o *MapOrgImportFile) GetJson() MapOrgImportFileJson {
	if o == nil || IsNil(o.Json) {
		var ret MapOrgImportFileJson
		return ret
	}
	return *o.Json
}

// GetJsonOk returns a tuple with the Json field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapOrgImportFile) GetJsonOk() (*MapOrgImportFileJson, bool) {
	if o == nil || IsNil(o.Json) {
		return nil, false
	}
	return o.Json, true
}

// HasJson returns a boolean if a field has been set.
func (o *MapOrgImportFile) HasJson() bool {
	if o != nil && !IsNil(o.Json) {
		return true
	}

	return false
}

// SetJson gets a reference to the given MapOrgImportFileJson and assigns it to the Json field.
func (o *MapOrgImportFile) SetJson(v MapOrgImportFileJson) {
	o.Json = &v
}

func (o MapOrgImportFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapOrgImportFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoDeviceprofileAssignment) {
		toSerialize["auto_deviceprofile_assignment"] = o.AutoDeviceprofileAssignment
	}
	if !IsNil(o.Csv) {
		toSerialize["csv"] = o.Csv
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.Json) {
		toSerialize["json"] = o.Json
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MapOrgImportFile) UnmarshalJSON(data []byte) (err error) {
	varMapOrgImportFile := _MapOrgImportFile{}

	err = json.Unmarshal(data, &varMapOrgImportFile)

	if err != nil {
		return err
	}

	*o = MapOrgImportFile(varMapOrgImportFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "auto_deviceprofile_assignment")
		delete(additionalProperties, "csv")
		delete(additionalProperties, "file")
		delete(additionalProperties, "json")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMapOrgImportFile struct {
	value *MapOrgImportFile
	isSet bool
}

func (v NullableMapOrgImportFile) Get() *MapOrgImportFile {
	return v.value
}

func (v *NullableMapOrgImportFile) Set(val *MapOrgImportFile) {
	v.value = val
	v.isSet = true
}

func (v NullableMapOrgImportFile) IsSet() bool {
	return v.isSet
}

func (v *NullableMapOrgImportFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapOrgImportFile(val *MapOrgImportFile) *NullableMapOrgImportFile {
	return &NullableMapOrgImportFile{value: val, isSet: true}
}

func (v NullableMapOrgImportFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapOrgImportFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


