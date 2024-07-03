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

// TicketStatus Status open: ticket is open, Mist is working on it  * pending: ticket is open and Requester attention is needed (e.g. Mist is asking for some more information)  * solved: ticket is marked as solved / considered by Mist (requester can update it, causing it to re-open; or rate it)  * closed: ticket is archived and cannot be changed
type TicketStatus string

// List of ticket_status
const (
	TICKETSTATUS_EMPTY TicketStatus = ""
	TICKETSTATUS_OPEN TicketStatus = "open"
	TICKETSTATUS_PENDING TicketStatus = "pending"
	TICKETSTATUS_SOLVED TicketStatus = "solved"
	TICKETSTATUS_CLOSED TicketStatus = "closed"
)

// All allowed values of TicketStatus enum
var AllowedTicketStatusEnumValues = []TicketStatus{
	"",
	"open",
	"pending",
	"solved",
	"closed",
}

func (v *TicketStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TicketStatus(value)
	for _, existing := range AllowedTicketStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TicketStatus", value)
}

// NewTicketStatusFromValue returns a pointer to a valid TicketStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTicketStatusFromValue(v string) (*TicketStatus, error) {
	ev := TicketStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TicketStatus: valid values are %v", v, AllowedTicketStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TicketStatus) IsValid() bool {
	for _, existing := range AllowedTicketStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ticket_status value
func (v TicketStatus) Ptr() *TicketStatus {
	return &v
}

type NullableTicketStatus struct {
	value *TicketStatus
	isSet bool
}

func (v NullableTicketStatus) Get() *TicketStatus {
	return v.value
}

func (v *NullableTicketStatus) Set(val *TicketStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketStatus(val *TicketStatus) *NullableTicketStatus {
	return &NullableTicketStatus{value: val, isSet: true}
}

func (v NullableTicketStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

