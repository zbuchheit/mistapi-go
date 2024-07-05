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

// ResponseStatsDevice struct for ResponseStatsDevice
type ResponseStatsDevice struct {
	ApStats *ApStats
	GatewayStats *GatewayStats
	SwitchStats *SwitchStats
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResponseStatsDevice) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ApStats
	err = json.Unmarshal(data, &dst.ApStats);
	if err == nil {
		jsonApStats, _ := json.Marshal(dst.ApStats)
		if string(jsonApStats) == "{}" { // empty struct
			dst.ApStats = nil
		} else {
			return nil // data stored in dst.ApStats, return on the first match
		}
	} else {
		dst.ApStats = nil
	}

	// try to unmarshal JSON data into GatewayStats
	err = json.Unmarshal(data, &dst.GatewayStats);
	if err == nil {
		jsonGatewayStats, _ := json.Marshal(dst.GatewayStats)
		if string(jsonGatewayStats) == "{}" { // empty struct
			dst.GatewayStats = nil
		} else {
			return nil // data stored in dst.GatewayStats, return on the first match
		}
	} else {
		dst.GatewayStats = nil
	}

	// try to unmarshal JSON data into SwitchStats
	err = json.Unmarshal(data, &dst.SwitchStats);
	if err == nil {
		jsonSwitchStats, _ := json.Marshal(dst.SwitchStats)
		if string(jsonSwitchStats) == "{}" { // empty struct
			dst.SwitchStats = nil
		} else {
			return nil // data stored in dst.SwitchStats, return on the first match
		}
	} else {
		dst.SwitchStats = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResponseStatsDevice)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResponseStatsDevice) MarshalJSON() ([]byte, error) {
	if src.ApStats != nil {
		return json.Marshal(&src.ApStats)
	}

	if src.GatewayStats != nil {
		return json.Marshal(&src.GatewayStats)
	}

	if src.SwitchStats != nil {
		return json.Marshal(&src.SwitchStats)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResponseStatsDevice struct {
	value *ResponseStatsDevice
	isSet bool
}

func (v NullableResponseStatsDevice) Get() *ResponseStatsDevice {
	return v.value
}

func (v *NullableResponseStatsDevice) Set(val *ResponseStatsDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStatsDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStatsDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStatsDevice(val *ResponseStatsDevice) *NullableResponseStatsDevice {
	return &NullableResponseStatsDevice{value: val, isSet: true}
}

func (v NullableResponseStatsDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStatsDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


