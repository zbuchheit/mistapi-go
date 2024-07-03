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

// CaptureWirelessFormat pcap format
type CaptureWirelessFormat string

// List of capture_wireless_format
const (
	CAPTUREWIRELESSFORMAT_EMPTY CaptureWirelessFormat = ""
	CAPTUREWIRELESSFORMAT_PCAP CaptureWirelessFormat = "pcap"
	CAPTUREWIRELESSFORMAT_STREAM CaptureWirelessFormat = "stream"
)

// All allowed values of CaptureWirelessFormat enum
var AllowedCaptureWirelessFormatEnumValues = []CaptureWirelessFormat{
	"",
	"pcap",
	"stream",
}

func (v *CaptureWirelessFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaptureWirelessFormat(value)
	for _, existing := range AllowedCaptureWirelessFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaptureWirelessFormat", value)
}

// NewCaptureWirelessFormatFromValue returns a pointer to a valid CaptureWirelessFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaptureWirelessFormatFromValue(v string) (*CaptureWirelessFormat, error) {
	ev := CaptureWirelessFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaptureWirelessFormat: valid values are %v", v, AllowedCaptureWirelessFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaptureWirelessFormat) IsValid() bool {
	for _, existing := range AllowedCaptureWirelessFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to capture_wireless_format value
func (v CaptureWirelessFormat) Ptr() *CaptureWirelessFormat {
	return &v
}

type NullableCaptureWirelessFormat struct {
	value *CaptureWirelessFormat
	isSet bool
}

func (v NullableCaptureWirelessFormat) Get() *CaptureWirelessFormat {
	return v.value
}

func (v *NullableCaptureWirelessFormat) Set(val *CaptureWirelessFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureWirelessFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureWirelessFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureWirelessFormat(val *CaptureWirelessFormat) *NullableCaptureWirelessFormat {
	return &NullableCaptureWirelessFormat{value: val, isSet: true}
}

func (v NullableCaptureWirelessFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureWirelessFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

