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

// OrgNacClientsCountDistinct the model 'OrgNacClientsCountDistinct'
type OrgNacClientsCountDistinct string

// List of org_nac_clients_count_distinct
const (
	ORGNACCLIENTSCOUNTDISTINCT_EMPTY OrgNacClientsCountDistinct = ""
	ORGNACCLIENTSCOUNTDISTINCT_TYPE OrgNacClientsCountDistinct = "type"
	ORGNACCLIENTSCOUNTDISTINCT_LAST_NACRULE_ID OrgNacClientsCountDistinct = "last_nacrule_id"
	ORGNACCLIENTSCOUNTDISTINCT_AUTH_TYPE OrgNacClientsCountDistinct = "auth_type"
	ORGNACCLIENTSCOUNTDISTINCT_LAST_VLAN OrgNacClientsCountDistinct = "last_vlan"
	ORGNACCLIENTSCOUNTDISTINCT_LAST_NAS_VENDOR OrgNacClientsCountDistinct = "last_nas_vendor"
	ORGNACCLIENTSCOUNTDISTINCT_LAST_USERNAME OrgNacClientsCountDistinct = "last_username"
	ORGNACCLIENTSCOUNTDISTINCT_LAST_AP OrgNacClientsCountDistinct = "last_ap"
	ORGNACCLIENTSCOUNTDISTINCT_MAC OrgNacClientsCountDistinct = "mac"
	ORGNACCLIENTSCOUNTDISTINCT_LAST_SSID OrgNacClientsCountDistinct = "last_ssid"
	ORGNACCLIENTSCOUNTDISTINCT_LAST_STATUS OrgNacClientsCountDistinct = "last_status"
	ORGNACCLIENTSCOUNTDISTINCT_MDM_COMPLIANCE OrgNacClientsCountDistinct = "mdm_compliance"
	ORGNACCLIENTSCOUNTDISTINCT_MDM_PROVIDER OrgNacClientsCountDistinct = "mdm_provider"
)

// All allowed values of OrgNacClientsCountDistinct enum
var AllowedOrgNacClientsCountDistinctEnumValues = []OrgNacClientsCountDistinct{
	"",
	"type",
	"last_nacrule_id",
	"auth_type",
	"last_vlan",
	"last_nas_vendor",
	"last_username",
	"last_ap",
	"mac",
	"last_ssid",
	"last_status",
	"mdm_compliance",
	"mdm_provider",
}

func (v *OrgNacClientsCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgNacClientsCountDistinct(value)
	for _, existing := range AllowedOrgNacClientsCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgNacClientsCountDistinct", value)
}

// NewOrgNacClientsCountDistinctFromValue returns a pointer to a valid OrgNacClientsCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgNacClientsCountDistinctFromValue(v string) (*OrgNacClientsCountDistinct, error) {
	ev := OrgNacClientsCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgNacClientsCountDistinct: valid values are %v", v, AllowedOrgNacClientsCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgNacClientsCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgNacClientsCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_nac_clients_count_distinct value
func (v OrgNacClientsCountDistinct) Ptr() *OrgNacClientsCountDistinct {
	return &v
}

type NullableOrgNacClientsCountDistinct struct {
	value *OrgNacClientsCountDistinct
	isSet bool
}

func (v NullableOrgNacClientsCountDistinct) Get() *OrgNacClientsCountDistinct {
	return v.value
}

func (v *NullableOrgNacClientsCountDistinct) Set(val *OrgNacClientsCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgNacClientsCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgNacClientsCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgNacClientsCountDistinct(val *OrgNacClientsCountDistinct) *NullableOrgNacClientsCountDistinct {
	return &NullableOrgNacClientsCountDistinct{value: val, isSet: true}
}

func (v NullableOrgNacClientsCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgNacClientsCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

