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

// MistDevice struct for MistDevice
type MistDevice struct {
	DeviceAp *DeviceAp
	DeviceGateway *DeviceGateway
	DeviceSwitch *DeviceSwitch
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MistDevice) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DeviceAp
	err = json.Unmarshal(data, &dst.DeviceAp);
	if err == nil {
		jsonDeviceAp, _ := json.Marshal(dst.DeviceAp)
		if string(jsonDeviceAp) == "{}" { // empty struct
			dst.DeviceAp = nil
		} else {
			return nil // data stored in dst.DeviceAp, return on the first match
		}
	} else {
		dst.DeviceAp = nil
	}

	// try to unmarshal JSON data into DeviceGateway
	err = json.Unmarshal(data, &dst.DeviceGateway);
	if err == nil {
		jsonDeviceGateway, _ := json.Marshal(dst.DeviceGateway)
		if string(jsonDeviceGateway) == "{}" { // empty struct
			dst.DeviceGateway = nil
		} else {
			return nil // data stored in dst.DeviceGateway, return on the first match
		}
	} else {
		dst.DeviceGateway = nil
	}

	// try to unmarshal JSON data into DeviceSwitch
	err = json.Unmarshal(data, &dst.DeviceSwitch);
	if err == nil {
		jsonDeviceSwitch, _ := json.Marshal(dst.DeviceSwitch)
		if string(jsonDeviceSwitch) == "{}" { // empty struct
			dst.DeviceSwitch = nil
		} else {
			return nil // data stored in dst.DeviceSwitch, return on the first match
		}
	} else {
		dst.DeviceSwitch = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MistDevice)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MistDevice) MarshalJSON() ([]byte, error) {
	if src.DeviceAp != nil {
		return json.Marshal(&src.DeviceAp)
	}

	if src.DeviceGateway != nil {
		return json.Marshal(&src.DeviceGateway)
	}

	if src.DeviceSwitch != nil {
		return json.Marshal(&src.DeviceSwitch)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMistDevice struct {
	value *MistDevice
	isSet bool
}

func (v NullableMistDevice) Get() *MistDevice {
	return v.value
}

func (v *NullableMistDevice) Set(val *MistDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableMistDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableMistDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMistDevice(val *MistDevice) *NullableMistDevice {
	return &NullableMistDevice{value: val, isSet: true}
}

func (v NullableMistDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMistDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


