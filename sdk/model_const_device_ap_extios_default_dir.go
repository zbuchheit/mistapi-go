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
	"fmt"
)

// ConstDeviceApExtiosDefaultDir the model 'ConstDeviceApExtiosDefaultDir'
type ConstDeviceApExtiosDefaultDir string

// List of const_device_ap_extios_default_dir
const (
	CONSTDEVICEAPEXTIOSDEFAULTDIR_EMPTY ConstDeviceApExtiosDefaultDir = ""
	CONSTDEVICEAPEXTIOSDEFAULTDIR_IN ConstDeviceApExtiosDefaultDir = "IN"
	CONSTDEVICEAPEXTIOSDEFAULTDIR_OUT ConstDeviceApExtiosDefaultDir = "OUT"
)

// All allowed values of ConstDeviceApExtiosDefaultDir enum
var AllowedConstDeviceApExtiosDefaultDirEnumValues = []ConstDeviceApExtiosDefaultDir{
	"",
	"IN",
	"OUT",
}

func (v *ConstDeviceApExtiosDefaultDir) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConstDeviceApExtiosDefaultDir(value)
	for _, existing := range AllowedConstDeviceApExtiosDefaultDirEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConstDeviceApExtiosDefaultDir", value)
}

// NewConstDeviceApExtiosDefaultDirFromValue returns a pointer to a valid ConstDeviceApExtiosDefaultDir
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConstDeviceApExtiosDefaultDirFromValue(v string) (*ConstDeviceApExtiosDefaultDir, error) {
	ev := ConstDeviceApExtiosDefaultDir(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConstDeviceApExtiosDefaultDir: valid values are %v", v, AllowedConstDeviceApExtiosDefaultDirEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConstDeviceApExtiosDefaultDir) IsValid() bool {
	for _, existing := range AllowedConstDeviceApExtiosDefaultDirEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to const_device_ap_extios_default_dir value
func (v ConstDeviceApExtiosDefaultDir) Ptr() *ConstDeviceApExtiosDefaultDir {
	return &v
}

type NullableConstDeviceApExtiosDefaultDir struct {
	value *ConstDeviceApExtiosDefaultDir
	isSet bool
}

func (v NullableConstDeviceApExtiosDefaultDir) Get() *ConstDeviceApExtiosDefaultDir {
	return v.value
}

func (v *NullableConstDeviceApExtiosDefaultDir) Set(val *ConstDeviceApExtiosDefaultDir) {
	v.value = val
	v.isSet = true
}

func (v NullableConstDeviceApExtiosDefaultDir) IsSet() bool {
	return v.isSet
}

func (v *NullableConstDeviceApExtiosDefaultDir) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstDeviceApExtiosDefaultDir(val *ConstDeviceApExtiosDefaultDir) *NullableConstDeviceApExtiosDefaultDir {
	return &NullableConstDeviceApExtiosDefaultDir{value: val, isSet: true}
}

func (v NullableConstDeviceApExtiosDefaultDir) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstDeviceApExtiosDefaultDir) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

