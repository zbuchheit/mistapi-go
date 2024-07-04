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

// OrgLicenseActionOperation to move a license, use the `amend` operation
type OrgLicenseActionOperation string

// List of org_license_action_operation
const (
	ORGLICENSEACTIONOPERATION_EMPTY OrgLicenseActionOperation = ""
	ORGLICENSEACTIONOPERATION_AMEND OrgLicenseActionOperation = "amend"
	ORGLICENSEACTIONOPERATION_UNAMEND OrgLicenseActionOperation = "unamend"
	ORGLICENSEACTIONOPERATION_DELETE OrgLicenseActionOperation = "delete"
	ORGLICENSEACTIONOPERATION_ANNOTATE OrgLicenseActionOperation = "annotate"
)

// All allowed values of OrgLicenseActionOperation enum
var AllowedOrgLicenseActionOperationEnumValues = []OrgLicenseActionOperation{
	"",
	"amend",
	"unamend",
	"delete",
	"annotate",
}

func (v *OrgLicenseActionOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgLicenseActionOperation(value)
	for _, existing := range AllowedOrgLicenseActionOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgLicenseActionOperation", value)
}

// NewOrgLicenseActionOperationFromValue returns a pointer to a valid OrgLicenseActionOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgLicenseActionOperationFromValue(v string) (*OrgLicenseActionOperation, error) {
	ev := OrgLicenseActionOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgLicenseActionOperation: valid values are %v", v, AllowedOrgLicenseActionOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgLicenseActionOperation) IsValid() bool {
	for _, existing := range AllowedOrgLicenseActionOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_license_action_operation value
func (v OrgLicenseActionOperation) Ptr() *OrgLicenseActionOperation {
	return &v
}

type NullableOrgLicenseActionOperation struct {
	value *OrgLicenseActionOperation
	isSet bool
}

func (v NullableOrgLicenseActionOperation) Get() *OrgLicenseActionOperation {
	return v.value
}

func (v *NullableOrgLicenseActionOperation) Set(val *OrgLicenseActionOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgLicenseActionOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgLicenseActionOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgLicenseActionOperation(val *OrgLicenseActionOperation) *NullableOrgLicenseActionOperation {
	return &NullableOrgLicenseActionOperation{value: val, isSet: true}
}

func (v NullableOrgLicenseActionOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgLicenseActionOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

