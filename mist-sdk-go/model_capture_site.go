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
	"gopkg.in/validator.v2"
	"fmt"
)

// CaptureSite - struct for CaptureSite
type CaptureSite struct {
	CaptureClient *CaptureClient
	CaptureGateway *CaptureGateway
	CaptureNewAssoc *CaptureNewAssoc
	CaptureRadiotap *CaptureRadiotap
	CaptureRadiotapwired *CaptureRadiotapwired
	CaptureScan *CaptureScan
	CaptureSwitch *CaptureSwitch
	CaptureWired *CaptureWired
	CaptureWireless *CaptureWireless
}

// CaptureClientAsCaptureSite is a convenience function that returns CaptureClient wrapped in CaptureSite
func CaptureClientAsCaptureSite(v *CaptureClient) CaptureSite {
	return CaptureSite{
		CaptureClient: v,
	}
}

// CaptureGatewayAsCaptureSite is a convenience function that returns CaptureGateway wrapped in CaptureSite
func CaptureGatewayAsCaptureSite(v *CaptureGateway) CaptureSite {
	return CaptureSite{
		CaptureGateway: v,
	}
}

// CaptureNewAssocAsCaptureSite is a convenience function that returns CaptureNewAssoc wrapped in CaptureSite
func CaptureNewAssocAsCaptureSite(v *CaptureNewAssoc) CaptureSite {
	return CaptureSite{
		CaptureNewAssoc: v,
	}
}

// CaptureRadiotapAsCaptureSite is a convenience function that returns CaptureRadiotap wrapped in CaptureSite
func CaptureRadiotapAsCaptureSite(v *CaptureRadiotap) CaptureSite {
	return CaptureSite{
		CaptureRadiotap: v,
	}
}

// CaptureRadiotapwiredAsCaptureSite is a convenience function that returns CaptureRadiotapwired wrapped in CaptureSite
func CaptureRadiotapwiredAsCaptureSite(v *CaptureRadiotapwired) CaptureSite {
	return CaptureSite{
		CaptureRadiotapwired: v,
	}
}

// CaptureScanAsCaptureSite is a convenience function that returns CaptureScan wrapped in CaptureSite
func CaptureScanAsCaptureSite(v *CaptureScan) CaptureSite {
	return CaptureSite{
		CaptureScan: v,
	}
}

// CaptureSwitchAsCaptureSite is a convenience function that returns CaptureSwitch wrapped in CaptureSite
func CaptureSwitchAsCaptureSite(v *CaptureSwitch) CaptureSite {
	return CaptureSite{
		CaptureSwitch: v,
	}
}

// CaptureWiredAsCaptureSite is a convenience function that returns CaptureWired wrapped in CaptureSite
func CaptureWiredAsCaptureSite(v *CaptureWired) CaptureSite {
	return CaptureSite{
		CaptureWired: v,
	}
}

// CaptureWirelessAsCaptureSite is a convenience function that returns CaptureWireless wrapped in CaptureSite
func CaptureWirelessAsCaptureSite(v *CaptureWireless) CaptureSite {
	return CaptureSite{
		CaptureWireless: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CaptureSite) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CaptureClient
	err = newStrictDecoder(data).Decode(&dst.CaptureClient)
	if err == nil {
		jsonCaptureClient, _ := json.Marshal(dst.CaptureClient)
		if string(jsonCaptureClient) == "{}" { // empty struct
			dst.CaptureClient = nil
		} else {
			if err = validator.Validate(dst.CaptureClient); err != nil {
				dst.CaptureClient = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureClient = nil
	}

	// try to unmarshal data into CaptureGateway
	err = newStrictDecoder(data).Decode(&dst.CaptureGateway)
	if err == nil {
		jsonCaptureGateway, _ := json.Marshal(dst.CaptureGateway)
		if string(jsonCaptureGateway) == "{}" { // empty struct
			dst.CaptureGateway = nil
		} else {
			if err = validator.Validate(dst.CaptureGateway); err != nil {
				dst.CaptureGateway = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureGateway = nil
	}

	// try to unmarshal data into CaptureNewAssoc
	err = newStrictDecoder(data).Decode(&dst.CaptureNewAssoc)
	if err == nil {
		jsonCaptureNewAssoc, _ := json.Marshal(dst.CaptureNewAssoc)
		if string(jsonCaptureNewAssoc) == "{}" { // empty struct
			dst.CaptureNewAssoc = nil
		} else {
			if err = validator.Validate(dst.CaptureNewAssoc); err != nil {
				dst.CaptureNewAssoc = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureNewAssoc = nil
	}

	// try to unmarshal data into CaptureRadiotap
	err = newStrictDecoder(data).Decode(&dst.CaptureRadiotap)
	if err == nil {
		jsonCaptureRadiotap, _ := json.Marshal(dst.CaptureRadiotap)
		if string(jsonCaptureRadiotap) == "{}" { // empty struct
			dst.CaptureRadiotap = nil
		} else {
			if err = validator.Validate(dst.CaptureRadiotap); err != nil {
				dst.CaptureRadiotap = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureRadiotap = nil
	}

	// try to unmarshal data into CaptureRadiotapwired
	err = newStrictDecoder(data).Decode(&dst.CaptureRadiotapwired)
	if err == nil {
		jsonCaptureRadiotapwired, _ := json.Marshal(dst.CaptureRadiotapwired)
		if string(jsonCaptureRadiotapwired) == "{}" { // empty struct
			dst.CaptureRadiotapwired = nil
		} else {
			if err = validator.Validate(dst.CaptureRadiotapwired); err != nil {
				dst.CaptureRadiotapwired = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureRadiotapwired = nil
	}

	// try to unmarshal data into CaptureScan
	err = newStrictDecoder(data).Decode(&dst.CaptureScan)
	if err == nil {
		jsonCaptureScan, _ := json.Marshal(dst.CaptureScan)
		if string(jsonCaptureScan) == "{}" { // empty struct
			dst.CaptureScan = nil
		} else {
			if err = validator.Validate(dst.CaptureScan); err != nil {
				dst.CaptureScan = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureScan = nil
	}

	// try to unmarshal data into CaptureSwitch
	err = newStrictDecoder(data).Decode(&dst.CaptureSwitch)
	if err == nil {
		jsonCaptureSwitch, _ := json.Marshal(dst.CaptureSwitch)
		if string(jsonCaptureSwitch) == "{}" { // empty struct
			dst.CaptureSwitch = nil
		} else {
			if err = validator.Validate(dst.CaptureSwitch); err != nil {
				dst.CaptureSwitch = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureSwitch = nil
	}

	// try to unmarshal data into CaptureWired
	err = newStrictDecoder(data).Decode(&dst.CaptureWired)
	if err == nil {
		jsonCaptureWired, _ := json.Marshal(dst.CaptureWired)
		if string(jsonCaptureWired) == "{}" { // empty struct
			dst.CaptureWired = nil
		} else {
			if err = validator.Validate(dst.CaptureWired); err != nil {
				dst.CaptureWired = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureWired = nil
	}

	// try to unmarshal data into CaptureWireless
	err = newStrictDecoder(data).Decode(&dst.CaptureWireless)
	if err == nil {
		jsonCaptureWireless, _ := json.Marshal(dst.CaptureWireless)
		if string(jsonCaptureWireless) == "{}" { // empty struct
			dst.CaptureWireless = nil
		} else {
			if err = validator.Validate(dst.CaptureWireless); err != nil {
				dst.CaptureWireless = nil
			} else {
				match++
			}
		}
	} else {
		dst.CaptureWireless = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CaptureClient = nil
		dst.CaptureGateway = nil
		dst.CaptureNewAssoc = nil
		dst.CaptureRadiotap = nil
		dst.CaptureRadiotapwired = nil
		dst.CaptureScan = nil
		dst.CaptureSwitch = nil
		dst.CaptureWired = nil
		dst.CaptureWireless = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CaptureSite)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CaptureSite)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CaptureSite) MarshalJSON() ([]byte, error) {
	if src.CaptureClient != nil {
		return json.Marshal(&src.CaptureClient)
	}

	if src.CaptureGateway != nil {
		return json.Marshal(&src.CaptureGateway)
	}

	if src.CaptureNewAssoc != nil {
		return json.Marshal(&src.CaptureNewAssoc)
	}

	if src.CaptureRadiotap != nil {
		return json.Marshal(&src.CaptureRadiotap)
	}

	if src.CaptureRadiotapwired != nil {
		return json.Marshal(&src.CaptureRadiotapwired)
	}

	if src.CaptureScan != nil {
		return json.Marshal(&src.CaptureScan)
	}

	if src.CaptureSwitch != nil {
		return json.Marshal(&src.CaptureSwitch)
	}

	if src.CaptureWired != nil {
		return json.Marshal(&src.CaptureWired)
	}

	if src.CaptureWireless != nil {
		return json.Marshal(&src.CaptureWireless)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CaptureSite) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CaptureClient != nil {
		return obj.CaptureClient
	}

	if obj.CaptureGateway != nil {
		return obj.CaptureGateway
	}

	if obj.CaptureNewAssoc != nil {
		return obj.CaptureNewAssoc
	}

	if obj.CaptureRadiotap != nil {
		return obj.CaptureRadiotap
	}

	if obj.CaptureRadiotapwired != nil {
		return obj.CaptureRadiotapwired
	}

	if obj.CaptureScan != nil {
		return obj.CaptureScan
	}

	if obj.CaptureSwitch != nil {
		return obj.CaptureSwitch
	}

	if obj.CaptureWired != nil {
		return obj.CaptureWired
	}

	if obj.CaptureWireless != nil {
		return obj.CaptureWireless
	}

	// all schemas are nil
	return nil
}

type NullableCaptureSite struct {
	value *CaptureSite
	isSet bool
}

func (v NullableCaptureSite) Get() *CaptureSite {
	return v.value
}

func (v *NullableCaptureSite) Set(val *CaptureSite) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureSite) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureSite(val *CaptureSite) *NullableCaptureSite {
	return &NullableCaptureSite{value: val, isSet: true}
}

func (v NullableCaptureSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


