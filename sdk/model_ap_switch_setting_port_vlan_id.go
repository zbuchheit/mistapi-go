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
	"gopkg.in/validator.v2"
	"fmt"
)

// ApSwitchSettingPortVlanId - native VLAN id, optional
type ApSwitchSettingPortVlanId struct {
	Int32 *int32
	String *string
}

// int32AsApSwitchSettingPortVlanId is a convenience function that returns int32 wrapped in ApSwitchSettingPortVlanId
func Int32AsApSwitchSettingPortVlanId(v *int32) ApSwitchSettingPortVlanId {
	return ApSwitchSettingPortVlanId{
		Int32: v,
	}
}

// stringAsApSwitchSettingPortVlanId is a convenience function that returns string wrapped in ApSwitchSettingPortVlanId
func StringAsApSwitchSettingPortVlanId(v *string) ApSwitchSettingPortVlanId {
	return ApSwitchSettingPortVlanId{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApSwitchSettingPortVlanId) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApSwitchSettingPortVlanId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApSwitchSettingPortVlanId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApSwitchSettingPortVlanId) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApSwitchSettingPortVlanId) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableApSwitchSettingPortVlanId struct {
	value *ApSwitchSettingPortVlanId
	isSet bool
}

func (v NullableApSwitchSettingPortVlanId) Get() *ApSwitchSettingPortVlanId {
	return v.value
}

func (v *NullableApSwitchSettingPortVlanId) Set(val *ApSwitchSettingPortVlanId) {
	v.value = val
	v.isSet = true
}

func (v NullableApSwitchSettingPortVlanId) IsSet() bool {
	return v.isSet
}

func (v *NullableApSwitchSettingPortVlanId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApSwitchSettingPortVlanId(val *ApSwitchSettingPortVlanId) *NullableApSwitchSettingPortVlanId {
	return &NullableApSwitchSettingPortVlanId{value: val, isSet: true}
}

func (v NullableApSwitchSettingPortVlanId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApSwitchSettingPortVlanId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


