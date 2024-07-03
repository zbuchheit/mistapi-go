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

// JunosPortConfigSpeed the model 'JunosPortConfigSpeed'
type JunosPortConfigSpeed string

// List of junos_port_config_speed
const (
	JUNOSPORTCONFIGSPEED_EMPTY JunosPortConfigSpeed = ""
	JUNOSPORTCONFIGSPEED_AUTO JunosPortConfigSpeed = "auto"
	JUNOSPORTCONFIGSPEED__10M JunosPortConfigSpeed = "10m"
	JUNOSPORTCONFIGSPEED__100M JunosPortConfigSpeed = "100m"
	JUNOSPORTCONFIGSPEED__1G JunosPortConfigSpeed = "1g"
	JUNOSPORTCONFIGSPEED__2_5G JunosPortConfigSpeed = "2.5g"
	JUNOSPORTCONFIGSPEED__5G JunosPortConfigSpeed = "5g"
)

// All allowed values of JunosPortConfigSpeed enum
var AllowedJunosPortConfigSpeedEnumValues = []JunosPortConfigSpeed{
	"",
	"auto",
	"10m",
	"100m",
	"1g",
	"2.5g",
	"5g",
}

func (v *JunosPortConfigSpeed) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JunosPortConfigSpeed(value)
	for _, existing := range AllowedJunosPortConfigSpeedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JunosPortConfigSpeed", value)
}

// NewJunosPortConfigSpeedFromValue returns a pointer to a valid JunosPortConfigSpeed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJunosPortConfigSpeedFromValue(v string) (*JunosPortConfigSpeed, error) {
	ev := JunosPortConfigSpeed(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JunosPortConfigSpeed: valid values are %v", v, AllowedJunosPortConfigSpeedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JunosPortConfigSpeed) IsValid() bool {
	for _, existing := range AllowedJunosPortConfigSpeedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to junos_port_config_speed value
func (v JunosPortConfigSpeed) Ptr() *JunosPortConfigSpeed {
	return &v
}

type NullableJunosPortConfigSpeed struct {
	value *JunosPortConfigSpeed
	isSet bool
}

func (v NullableJunosPortConfigSpeed) Get() *JunosPortConfigSpeed {
	return v.value
}

func (v *NullableJunosPortConfigSpeed) Set(val *JunosPortConfigSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableJunosPortConfigSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableJunosPortConfigSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJunosPortConfigSpeed(val *JunosPortConfigSpeed) *NullableJunosPortConfigSpeed {
	return &NullableJunosPortConfigSpeed{value: val, isSet: true}
}

func (v NullableJunosPortConfigSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJunosPortConfigSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

