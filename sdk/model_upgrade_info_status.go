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

// UpgradeInfoStatus the model 'UpgradeInfoStatus'
type UpgradeInfoStatus string

// List of upgrade_info_status
const (
	UPGRADEINFOSTATUS_EMPTY UpgradeInfoStatus = ""
	UPGRADEINFOSTATUS_STARTING UpgradeInfoStatus = "starting"
	UPGRADEINFOSTATUS_INPROGRESS UpgradeInfoStatus = "inprogress"
	UPGRADEINFOSTATUS_SUCCESS UpgradeInfoStatus = "success"
	UPGRADEINFOSTATUS_ERROR UpgradeInfoStatus = "error"
	UPGRADEINFOSTATUS_SCHEDULED UpgradeInfoStatus = "scheduled"
)

// All allowed values of UpgradeInfoStatus enum
var AllowedUpgradeInfoStatusEnumValues = []UpgradeInfoStatus{
	"",
	"starting",
	"inprogress",
	"success",
	"error",
	"scheduled",
}

func (v *UpgradeInfoStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpgradeInfoStatus(value)
	for _, existing := range AllowedUpgradeInfoStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpgradeInfoStatus", value)
}

// NewUpgradeInfoStatusFromValue returns a pointer to a valid UpgradeInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpgradeInfoStatusFromValue(v string) (*UpgradeInfoStatus, error) {
	ev := UpgradeInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpgradeInfoStatus: valid values are %v", v, AllowedUpgradeInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpgradeInfoStatus) IsValid() bool {
	for _, existing := range AllowedUpgradeInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to upgrade_info_status value
func (v UpgradeInfoStatus) Ptr() *UpgradeInfoStatus {
	return &v
}

type NullableUpgradeInfoStatus struct {
	value *UpgradeInfoStatus
	isSet bool
}

func (v NullableUpgradeInfoStatus) Get() *UpgradeInfoStatus {
	return v.value
}

func (v *NullableUpgradeInfoStatus) Set(val *UpgradeInfoStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUpgradeInfoStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUpgradeInfoStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpgradeInfoStatus(val *UpgradeInfoStatus) *NullableUpgradeInfoStatus {
	return &NullableUpgradeInfoStatus{value: val, isSet: true}
}

func (v NullableUpgradeInfoStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpgradeInfoStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

