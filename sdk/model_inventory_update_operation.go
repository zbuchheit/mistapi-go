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

// InventoryUpdateOperation * if `op`== `upgrade_to_mist`: Upgrade to mist-managed * if `op`== `downgrade_to_jsi`: Downgrade to basic monitoring. When downgrading a VC member to jsi, we will move the cloud connection of the VC to jsi-terminator and keep all VC device/inventories intact for pain-free upgrading back to mist. * if `op`== `assign`: Assign inventory to a site * if `op`== `unassign`: Unassign inventory from a site * if `op`== `delete`: Delete multiple inventory from org. If the device is already assigned to a site, it will be unassigned.
type InventoryUpdateOperation string

// List of inventory_update_operation
const (
	INVENTORYUPDATEOPERATION_EMPTY InventoryUpdateOperation = ""
	INVENTORYUPDATEOPERATION_ASSIGN InventoryUpdateOperation = "assign"
	INVENTORYUPDATEOPERATION_UNASSIGN InventoryUpdateOperation = "unassign"
	INVENTORYUPDATEOPERATION_DELETE InventoryUpdateOperation = "delete"
	INVENTORYUPDATEOPERATION_UPGRADE_TO_MIST InventoryUpdateOperation = "upgrade_to_mist"
	INVENTORYUPDATEOPERATION_DOWNGRADE_TO_JSI InventoryUpdateOperation = "downgrade_to_jsi"
)

// All allowed values of InventoryUpdateOperation enum
var AllowedInventoryUpdateOperationEnumValues = []InventoryUpdateOperation{
	"",
	"assign",
	"unassign",
	"delete",
	"upgrade_to_mist",
	"downgrade_to_jsi",
}

func (v *InventoryUpdateOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InventoryUpdateOperation(value)
	for _, existing := range AllowedInventoryUpdateOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InventoryUpdateOperation", value)
}

// NewInventoryUpdateOperationFromValue returns a pointer to a valid InventoryUpdateOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInventoryUpdateOperationFromValue(v string) (*InventoryUpdateOperation, error) {
	ev := InventoryUpdateOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InventoryUpdateOperation: valid values are %v", v, AllowedInventoryUpdateOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InventoryUpdateOperation) IsValid() bool {
	for _, existing := range AllowedInventoryUpdateOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to inventory_update_operation value
func (v InventoryUpdateOperation) Ptr() *InventoryUpdateOperation {
	return &v
}

type NullableInventoryUpdateOperation struct {
	value *InventoryUpdateOperation
	isSet bool
}

func (v NullableInventoryUpdateOperation) Get() *InventoryUpdateOperation {
	return v.value
}

func (v *NullableInventoryUpdateOperation) Set(val *InventoryUpdateOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryUpdateOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryUpdateOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryUpdateOperation(val *InventoryUpdateOperation) *NullableInventoryUpdateOperation {
	return &NullableInventoryUpdateOperation{value: val, isSet: true}
}

func (v NullableInventoryUpdateOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryUpdateOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

