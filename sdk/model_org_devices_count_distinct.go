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

// OrgDevicesCountDistinct the model 'OrgDevicesCountDistinct'
type OrgDevicesCountDistinct string

// List of org_devices_count_distinct
const (
	ORGDEVICESCOUNTDISTINCT_EMPTY OrgDevicesCountDistinct = ""
	ORGDEVICESCOUNTDISTINCT_HOSTNAME OrgDevicesCountDistinct = "hostname"
	ORGDEVICESCOUNTDISTINCT_SITE_ID OrgDevicesCountDistinct = "site_id"
	ORGDEVICESCOUNTDISTINCT_MODEL OrgDevicesCountDistinct = "model"
	ORGDEVICESCOUNTDISTINCT_MAC OrgDevicesCountDistinct = "mac"
	ORGDEVICESCOUNTDISTINCT_VERSION OrgDevicesCountDistinct = "version"
	ORGDEVICESCOUNTDISTINCT_IP OrgDevicesCountDistinct = "ip"
	ORGDEVICESCOUNTDISTINCT_MXTUNNEL_STATUS OrgDevicesCountDistinct = "mxtunnel_status"
	ORGDEVICESCOUNTDISTINCT_MXEDGE_ID OrgDevicesCountDistinct = "mxedge_id"
	ORGDEVICESCOUNTDISTINCT_LLDP_SYSTEM_NAME OrgDevicesCountDistinct = "lldp_system_name"
	ORGDEVICESCOUNTDISTINCT_LLDP_SYSTEM_DESC OrgDevicesCountDistinct = "lldp_system_desc"
	ORGDEVICESCOUNTDISTINCT_LLDP_PORT_ID OrgDevicesCountDistinct = "lldp_port_id"
	ORGDEVICESCOUNTDISTINCT_LLDP_MGMT_ADDR OrgDevicesCountDistinct = "lldp_mgmt_addr"
)

// All allowed values of OrgDevicesCountDistinct enum
var AllowedOrgDevicesCountDistinctEnumValues = []OrgDevicesCountDistinct{
	"",
	"hostname",
	"site_id",
	"model",
	"mac",
	"version",
	"ip",
	"mxtunnel_status",
	"mxedge_id",
	"lldp_system_name",
	"lldp_system_desc",
	"lldp_port_id",
	"lldp_mgmt_addr",
}

func (v *OrgDevicesCountDistinct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrgDevicesCountDistinct(value)
	for _, existing := range AllowedOrgDevicesCountDistinctEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrgDevicesCountDistinct", value)
}

// NewOrgDevicesCountDistinctFromValue returns a pointer to a valid OrgDevicesCountDistinct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrgDevicesCountDistinctFromValue(v string) (*OrgDevicesCountDistinct, error) {
	ev := OrgDevicesCountDistinct(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrgDevicesCountDistinct: valid values are %v", v, AllowedOrgDevicesCountDistinctEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrgDevicesCountDistinct) IsValid() bool {
	for _, existing := range AllowedOrgDevicesCountDistinctEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to org_devices_count_distinct value
func (v OrgDevicesCountDistinct) Ptr() *OrgDevicesCountDistinct {
	return &v
}

type NullableOrgDevicesCountDistinct struct {
	value *OrgDevicesCountDistinct
	isSet bool
}

func (v NullableOrgDevicesCountDistinct) Get() *OrgDevicesCountDistinct {
	return v.value
}

func (v *NullableOrgDevicesCountDistinct) Set(val *OrgDevicesCountDistinct) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgDevicesCountDistinct) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgDevicesCountDistinct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgDevicesCountDistinct(val *OrgDevicesCountDistinct) *NullableOrgDevicesCountDistinct {
	return &NullableOrgDevicesCountDistinct{value: val, isSet: true}
}

func (v NullableOrgDevicesCountDistinct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgDevicesCountDistinct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

