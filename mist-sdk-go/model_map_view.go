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

// MapView when type=google
type MapView string

// List of map_view
const (
	MAPVIEW_EMPTY MapView = ""
	MAPVIEW_ROADMAP MapView = "roadmap"
	MAPVIEW_SATELLITE MapView = "satellite"
	MAPVIEW_HYBRID MapView = "hybrid"
	MAPVIEW_TERRAIN MapView = "terrain"
)

// All allowed values of MapView enum
var AllowedMapViewEnumValues = []MapView{
	"",
	"roadmap",
	"satellite",
	"hybrid",
	"terrain",
}

func (v *MapView) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MapView(value)
	for _, existing := range AllowedMapViewEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MapView", value)
}

// NewMapViewFromValue returns a pointer to a valid MapView
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMapViewFromValue(v string) (*MapView, error) {
	ev := MapView(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MapView: valid values are %v", v, AllowedMapViewEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MapView) IsValid() bool {
	for _, existing := range AllowedMapViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to map_view value
func (v MapView) Ptr() *MapView {
	return &v
}

type NullableMapView struct {
	value *MapView
	isSet bool
}

func (v NullableMapView) Get() *MapView {
	return v.value
}

func (v *NullableMapView) Set(val *MapView) {
	v.value = val
	v.isSet = true
}

func (v NullableMapView) IsSet() bool {
	return v.isSet
}

func (v *NullableMapView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapView(val *MapView) *NullableMapView {
	return &NullableMapView{value: val, isSet: true}
}

func (v NullableMapView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

