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

// TunnelConfigsAuthAlgo the model 'TunnelConfigsAuthAlgo'
type TunnelConfigsAuthAlgo string

// List of tunnel_configs_auth_algo
const (
	TUNNELCONFIGSAUTHALGO_EMPTY TunnelConfigsAuthAlgo = ""
	TUNNELCONFIGSAUTHALGO_SHA1 TunnelConfigsAuthAlgo = "sha1"
	TUNNELCONFIGSAUTHALGO_SHA2 TunnelConfigsAuthAlgo = "sha2"
	TUNNELCONFIGSAUTHALGO_MD5 TunnelConfigsAuthAlgo = "md5"
)

// All allowed values of TunnelConfigsAuthAlgo enum
var AllowedTunnelConfigsAuthAlgoEnumValues = []TunnelConfigsAuthAlgo{
	"",
	"sha1",
	"sha2",
	"md5",
}

func (v *TunnelConfigsAuthAlgo) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TunnelConfigsAuthAlgo(value)
	for _, existing := range AllowedTunnelConfigsAuthAlgoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TunnelConfigsAuthAlgo", value)
}

// NewTunnelConfigsAuthAlgoFromValue returns a pointer to a valid TunnelConfigsAuthAlgo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTunnelConfigsAuthAlgoFromValue(v string) (*TunnelConfigsAuthAlgo, error) {
	ev := TunnelConfigsAuthAlgo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TunnelConfigsAuthAlgo: valid values are %v", v, AllowedTunnelConfigsAuthAlgoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TunnelConfigsAuthAlgo) IsValid() bool {
	for _, existing := range AllowedTunnelConfigsAuthAlgoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tunnel_configs_auth_algo value
func (v TunnelConfigsAuthAlgo) Ptr() *TunnelConfigsAuthAlgo {
	return &v
}

type NullableTunnelConfigsAuthAlgo struct {
	value *TunnelConfigsAuthAlgo
	isSet bool
}

func (v NullableTunnelConfigsAuthAlgo) Get() *TunnelConfigsAuthAlgo {
	return v.value
}

func (v *NullableTunnelConfigsAuthAlgo) Set(val *TunnelConfigsAuthAlgo) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelConfigsAuthAlgo) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelConfigsAuthAlgo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelConfigsAuthAlgo(val *TunnelConfigsAuthAlgo) *NullableTunnelConfigsAuthAlgo {
	return &NullableTunnelConfigsAuthAlgo{value: val, isSet: true}
}

func (v NullableTunnelConfigsAuthAlgo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelConfigsAuthAlgo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

