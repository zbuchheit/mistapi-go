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

// MxclusterNacClientVendor convention to be followed is : \"<vendor>-<variant>\" <variant> could be an os/platform/model/company for ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, airnonet) etc.
type MxclusterNacClientVendor string

// List of mxcluster_nac_client_vendor
const (
	MXCLUSTERNACCLIENTVENDOR_EMPTY MxclusterNacClientVendor = ""
	MXCLUSTERNACCLIENTVENDOR_GENERIC MxclusterNacClientVendor = "generic"
	MXCLUSTERNACCLIENTVENDOR_JUNIPER MxclusterNacClientVendor = "juniper"
	MXCLUSTERNACCLIENTVENDOR_ARUBA MxclusterNacClientVendor = "aruba"
	MXCLUSTERNACCLIENTVENDOR_PALOALTO MxclusterNacClientVendor = "paloalto"
	MXCLUSTERNACCLIENTVENDOR_CISCO_AIRONET MxclusterNacClientVendor = "cisco-aironet"
	MXCLUSTERNACCLIENTVENDOR_CISCO_IOS MxclusterNacClientVendor = "cisco-ios"
	MXCLUSTERNACCLIENTVENDOR_CISCO_MERAKI MxclusterNacClientVendor = "cisco-meraki"
)

// All allowed values of MxclusterNacClientVendor enum
var AllowedMxclusterNacClientVendorEnumValues = []MxclusterNacClientVendor{
	"",
	"generic",
	"juniper",
	"aruba",
	"paloalto",
	"cisco-aironet",
	"cisco-ios",
	"cisco-meraki",
}

func (v *MxclusterNacClientVendor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MxclusterNacClientVendor(value)
	for _, existing := range AllowedMxclusterNacClientVendorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MxclusterNacClientVendor", value)
}

// NewMxclusterNacClientVendorFromValue returns a pointer to a valid MxclusterNacClientVendor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMxclusterNacClientVendorFromValue(v string) (*MxclusterNacClientVendor, error) {
	ev := MxclusterNacClientVendor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MxclusterNacClientVendor: valid values are %v", v, AllowedMxclusterNacClientVendorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MxclusterNacClientVendor) IsValid() bool {
	for _, existing := range AllowedMxclusterNacClientVendorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to mxcluster_nac_client_vendor value
func (v MxclusterNacClientVendor) Ptr() *MxclusterNacClientVendor {
	return &v
}

type NullableMxclusterNacClientVendor struct {
	value *MxclusterNacClientVendor
	isSet bool
}

func (v NullableMxclusterNacClientVendor) Get() *MxclusterNacClientVendor {
	return v.value
}

func (v *NullableMxclusterNacClientVendor) Set(val *MxclusterNacClientVendor) {
	v.value = val
	v.isSet = true
}

func (v NullableMxclusterNacClientVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableMxclusterNacClientVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMxclusterNacClientVendor(val *MxclusterNacClientVendor) *NullableMxclusterNacClientVendor {
	return &NullableMxclusterNacClientVendor{value: val, isSet: true}
}

func (v NullableMxclusterNacClientVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMxclusterNacClientVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

