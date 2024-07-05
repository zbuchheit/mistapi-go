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

// AutoOrientationState The state of auto orient for a given map derived from an Enum
type AutoOrientationState string

// List of auto_orientation_state
const (
	AUTOORIENTATIONSTATE_EMPTY AutoOrientationState = ""
	AUTOORIENTATIONSTATE_NOT_STARTED AutoOrientationState = "Not Started"
	AUTOORIENTATIONSTATE_ENQUEUED AutoOrientationState = "Enqueued"
	AUTOORIENTATIONSTATE_ORIENTED AutoOrientationState = "Oriented"
)

// All allowed values of AutoOrientationState enum
var AllowedAutoOrientationStateEnumValues = []AutoOrientationState{
	"",
	"Not Started",
	"Enqueued",
	"Oriented",
}

func (v *AutoOrientationState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutoOrientationState(value)
	for _, existing := range AllowedAutoOrientationStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutoOrientationState", value)
}

// NewAutoOrientationStateFromValue returns a pointer to a valid AutoOrientationState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutoOrientationStateFromValue(v string) (*AutoOrientationState, error) {
	ev := AutoOrientationState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutoOrientationState: valid values are %v", v, AllowedAutoOrientationStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutoOrientationState) IsValid() bool {
	for _, existing := range AllowedAutoOrientationStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to auto_orientation_state value
func (v AutoOrientationState) Ptr() *AutoOrientationState {
	return &v
}

type NullableAutoOrientationState struct {
	value *AutoOrientationState
	isSet bool
}

func (v NullableAutoOrientationState) Get() *AutoOrientationState {
	return v.value
}

func (v *NullableAutoOrientationState) Set(val *AutoOrientationState) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoOrientationState) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoOrientationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoOrientationState(val *AutoOrientationState) *NullableAutoOrientationState {
	return &NullableAutoOrientationState{value: val, isSet: true}
}

func (v NullableAutoOrientationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoOrientationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

