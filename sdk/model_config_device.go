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

// ConfigDevice struct for ConfigDevice
type ConfigDevice struct {
	Ap *Ap
	Gateway *Gateway
	ModelSwitch *ModelSwitch
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ConfigDevice) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Ap
	err = json.Unmarshal(data, &dst.Ap);
	if err == nil {
		jsonAp, _ := json.Marshal(dst.Ap)
		if string(jsonAp) == "{}" { // empty struct
			dst.Ap = nil
		} else {
			return nil // data stored in dst.Ap, return on the first match
		}
	} else {
		dst.Ap = nil
	}

	// try to unmarshal JSON data into Gateway
	err = json.Unmarshal(data, &dst.Gateway);
	if err == nil {
		jsonGateway, _ := json.Marshal(dst.Gateway)
		if string(jsonGateway) == "{}" { // empty struct
			dst.Gateway = nil
		} else {
			return nil // data stored in dst.Gateway, return on the first match
		}
	} else {
		dst.Gateway = nil
	}

	// try to unmarshal JSON data into ModelSwitch
	err = json.Unmarshal(data, &dst.ModelSwitch);
	if err == nil {
		jsonModelSwitch, _ := json.Marshal(dst.ModelSwitch)
		if string(jsonModelSwitch) == "{}" { // empty struct
			dst.ModelSwitch = nil
		} else {
			return nil // data stored in dst.ModelSwitch, return on the first match
		}
	} else {
		dst.ModelSwitch = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ConfigDevice)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ConfigDevice) MarshalJSON() ([]byte, error) {
	if src.Ap != nil {
		return json.Marshal(&src.Ap)
	}

	if src.Gateway != nil {
		return json.Marshal(&src.Gateway)
	}

	if src.ModelSwitch != nil {
		return json.Marshal(&src.ModelSwitch)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableConfigDevice struct {
	value *ConfigDevice
	isSet bool
}

func (v NullableConfigDevice) Get() *ConfigDevice {
	return v.value
}

func (v *NullableConfigDevice) Set(val *ConfigDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigDevice(val *ConfigDevice) *NullableConfigDevice {
	return &NullableConfigDevice{value: val, isSet: true}
}

func (v NullableConfigDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


