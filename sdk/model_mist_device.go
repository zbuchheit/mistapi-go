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

// MistDevice - struct for MistDevice
type MistDevice struct {
	Ap *Ap
	Gateway *Gateway
	ModelSwitch *ModelSwitch
}

// ApAsMistDevice is a convenience function that returns Ap wrapped in MistDevice
func ApAsMistDevice(v *Ap) MistDevice {
	return MistDevice{
		Ap: v,
	}
}

// GatewayAsMistDevice is a convenience function that returns Gateway wrapped in MistDevice
func GatewayAsMistDevice(v *Gateway) MistDevice {
	return MistDevice{
		Gateway: v,
	}
}

// ModelSwitchAsMistDevice is a convenience function that returns ModelSwitch wrapped in MistDevice
func ModelSwitchAsMistDevice(v *ModelSwitch) MistDevice {
	return MistDevice{
		ModelSwitch: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MistDevice) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ap
	err = newStrictDecoder(data).Decode(&dst.Ap)
	if err == nil {
		jsonAp, _ := json.Marshal(dst.Ap)
		if string(jsonAp) == "{}" { // empty struct
			dst.Ap = nil
		} else {
			if err = validator.Validate(dst.Ap); err != nil {
				dst.Ap = nil
			} else {
				match++
			}
		}
	} else {
		dst.Ap = nil
	}

	// try to unmarshal data into Gateway
	err = newStrictDecoder(data).Decode(&dst.Gateway)
	if err == nil {
		jsonGateway, _ := json.Marshal(dst.Gateway)
		if string(jsonGateway) == "{}" { // empty struct
			dst.Gateway = nil
		} else {
			if err = validator.Validate(dst.Gateway); err != nil {
				dst.Gateway = nil
			} else {
				match++
			}
		}
	} else {
		dst.Gateway = nil
	}

	// try to unmarshal data into ModelSwitch
	err = newStrictDecoder(data).Decode(&dst.ModelSwitch)
	if err == nil {
		jsonModelSwitch, _ := json.Marshal(dst.ModelSwitch)
		if string(jsonModelSwitch) == "{}" { // empty struct
			dst.ModelSwitch = nil
		} else {
			if err = validator.Validate(dst.ModelSwitch); err != nil {
				dst.ModelSwitch = nil
			} else {
				match++
			}
		}
	} else {
		dst.ModelSwitch = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ap = nil
		dst.Gateway = nil
		dst.ModelSwitch = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MistDevice)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MistDevice)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MistDevice) MarshalJSON() ([]byte, error) {
	if src.Ap != nil {
		return json.Marshal(&src.Ap)
	}

	if src.Gateway != nil {
		return json.Marshal(&src.Gateway)
	}

	if src.ModelSwitch != nil {
		return json.Marshal(&src.ModelSwitch)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MistDevice) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Ap != nil {
		return obj.Ap
	}

	if obj.Gateway != nil {
		return obj.Gateway
	}

	if obj.ModelSwitch != nil {
		return obj.ModelSwitch
	}

	// all schemas are nil
	return nil
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


