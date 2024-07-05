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

// RfClientType the model 'RfClientType'
type RfClientType string

// List of rf_client_type
const (
	RFCLIENTTYPE_EMPTY RfClientType = ""
	RFCLIENTTYPE_SDKCLIENT RfClientType = "sdkclient"
	RFCLIENTTYPE_CLIENT RfClientType = "client"
	RFCLIENTTYPE_ASSET RfClientType = "asset"
)

// All allowed values of RfClientType enum
var AllowedRfClientTypeEnumValues = []RfClientType{
	"",
	"sdkclient",
	"client",
	"asset",
}

func (v *RfClientType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RfClientType(value)
	for _, existing := range AllowedRfClientTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RfClientType", value)
}

// NewRfClientTypeFromValue returns a pointer to a valid RfClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRfClientTypeFromValue(v string) (*RfClientType, error) {
	ev := RfClientType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RfClientType: valid values are %v", v, AllowedRfClientTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RfClientType) IsValid() bool {
	for _, existing := range AllowedRfClientTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to rf_client_type value
func (v RfClientType) Ptr() *RfClientType {
	return &v
}

type NullableRfClientType struct {
	value *RfClientType
	isSet bool
}

func (v NullableRfClientType) Get() *RfClientType {
	return v.value
}

func (v *NullableRfClientType) Set(val *RfClientType) {
	v.value = val
	v.isSet = true
}

func (v NullableRfClientType) IsSet() bool {
	return v.isSet
}

func (v *NullableRfClientType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRfClientType(val *RfClientType) *NullableRfClientType {
	return &NullableRfClientType{value: val, isSet: true}
}

func (v NullableRfClientType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRfClientType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

