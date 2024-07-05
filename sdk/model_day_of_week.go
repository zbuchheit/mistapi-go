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

// DayOfWeek the model 'DayOfWeek'
type DayOfWeek string

// List of day_of_week
const (
	DAYOFWEEK_EMPTY DayOfWeek = ""
	DAYOFWEEK_ANY DayOfWeek = "any"
	DAYOFWEEK_MON DayOfWeek = "mon"
	DAYOFWEEK_TUE DayOfWeek = "tue"
	DAYOFWEEK_WED DayOfWeek = "wed"
	DAYOFWEEK_THU DayOfWeek = "thu"
	DAYOFWEEK_FRI DayOfWeek = "fri"
	DAYOFWEEK_SAT DayOfWeek = "sat"
	DAYOFWEEK_SUN DayOfWeek = "sun"
)

// All allowed values of DayOfWeek enum
var AllowedDayOfWeekEnumValues = []DayOfWeek{
	"",
	"any",
	"mon",
	"tue",
	"wed",
	"thu",
	"fri",
	"sat",
	"sun",
}

func (v *DayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DayOfWeek(value)
	for _, existing := range AllowedDayOfWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DayOfWeek", value)
}

// NewDayOfWeekFromValue returns a pointer to a valid DayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDayOfWeekFromValue(v string) (*DayOfWeek, error) {
	ev := DayOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DayOfWeek: valid values are %v", v, AllowedDayOfWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DayOfWeek) IsValid() bool {
	for _, existing := range AllowedDayOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to day_of_week value
func (v DayOfWeek) Ptr() *DayOfWeek {
	return &v
}

type NullableDayOfWeek struct {
	value *DayOfWeek
	isSet bool
}

func (v NullableDayOfWeek) Get() *DayOfWeek {
	return v.value
}

func (v *NullableDayOfWeek) Set(val *DayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableDayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableDayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayOfWeek(val *DayOfWeek) *NullableDayOfWeek {
	return &NullableDayOfWeek{value: val, isSet: true}
}

func (v NullableDayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

