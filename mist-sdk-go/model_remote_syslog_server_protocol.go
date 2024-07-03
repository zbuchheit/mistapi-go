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

// RemoteSyslogServerProtocol the model 'RemoteSyslogServerProtocol'
type RemoteSyslogServerProtocol string

// List of remote_syslog_server_protocol
const (
	REMOTESYSLOGSERVERPROTOCOL_EMPTY RemoteSyslogServerProtocol = ""
	REMOTESYSLOGSERVERPROTOCOL_UDP RemoteSyslogServerProtocol = "udp"
	REMOTESYSLOGSERVERPROTOCOL_TCP RemoteSyslogServerProtocol = "tcp"
)

// All allowed values of RemoteSyslogServerProtocol enum
var AllowedRemoteSyslogServerProtocolEnumValues = []RemoteSyslogServerProtocol{
	"",
	"udp",
	"tcp",
}

func (v *RemoteSyslogServerProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RemoteSyslogServerProtocol(value)
	for _, existing := range AllowedRemoteSyslogServerProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RemoteSyslogServerProtocol", value)
}

// NewRemoteSyslogServerProtocolFromValue returns a pointer to a valid RemoteSyslogServerProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRemoteSyslogServerProtocolFromValue(v string) (*RemoteSyslogServerProtocol, error) {
	ev := RemoteSyslogServerProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RemoteSyslogServerProtocol: valid values are %v", v, AllowedRemoteSyslogServerProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RemoteSyslogServerProtocol) IsValid() bool {
	for _, existing := range AllowedRemoteSyslogServerProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to remote_syslog_server_protocol value
func (v RemoteSyslogServerProtocol) Ptr() *RemoteSyslogServerProtocol {
	return &v
}

type NullableRemoteSyslogServerProtocol struct {
	value *RemoteSyslogServerProtocol
	isSet bool
}

func (v NullableRemoteSyslogServerProtocol) Get() *RemoteSyslogServerProtocol {
	return v.value
}

func (v *NullableRemoteSyslogServerProtocol) Set(val *RemoteSyslogServerProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteSyslogServerProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteSyslogServerProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteSyslogServerProtocol(val *RemoteSyslogServerProtocol) *NullableRemoteSyslogServerProtocol {
	return &NullableRemoteSyslogServerProtocol{value: val, isSet: true}
}

func (v NullableRemoteSyslogServerProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteSyslogServerProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

