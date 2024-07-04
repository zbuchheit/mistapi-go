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
)

// checks if the SimpleAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleAlert{}

// SimpleAlert Set of heuristic rules will be enabled when marvis subscription is not available. It triggers when, in a Z minute window, there are more than Y distinct client encountring over X failures
type SimpleAlert struct {
	ArpFailure *SimpleAlertArpFailure `json:"arp_failure,omitempty"`
	DhcpFailure *SimpleAlertDhcpFailure `json:"dhcp_failure,omitempty"`
	DnsFailure *SimpleAlertDnsFailure `json:"dns_failure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimpleAlert SimpleAlert

// NewSimpleAlert instantiates a new SimpleAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleAlert() *SimpleAlert {
	this := SimpleAlert{}
	return &this
}

// NewSimpleAlertWithDefaults instantiates a new SimpleAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleAlertWithDefaults() *SimpleAlert {
	this := SimpleAlert{}
	return &this
}

// GetArpFailure returns the ArpFailure field value if set, zero value otherwise.
func (o *SimpleAlert) GetArpFailure() SimpleAlertArpFailure {
	if o == nil || IsNil(o.ArpFailure) {
		var ret SimpleAlertArpFailure
		return ret
	}
	return *o.ArpFailure
}

// GetArpFailureOk returns a tuple with the ArpFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAlert) GetArpFailureOk() (*SimpleAlertArpFailure, bool) {
	if o == nil || IsNil(o.ArpFailure) {
		return nil, false
	}
	return o.ArpFailure, true
}

// HasArpFailure returns a boolean if a field has been set.
func (o *SimpleAlert) HasArpFailure() bool {
	if o != nil && !IsNil(o.ArpFailure) {
		return true
	}

	return false
}

// SetArpFailure gets a reference to the given SimpleAlertArpFailure and assigns it to the ArpFailure field.
func (o *SimpleAlert) SetArpFailure(v SimpleAlertArpFailure) {
	o.ArpFailure = &v
}

// GetDhcpFailure returns the DhcpFailure field value if set, zero value otherwise.
func (o *SimpleAlert) GetDhcpFailure() SimpleAlertDhcpFailure {
	if o == nil || IsNil(o.DhcpFailure) {
		var ret SimpleAlertDhcpFailure
		return ret
	}
	return *o.DhcpFailure
}

// GetDhcpFailureOk returns a tuple with the DhcpFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAlert) GetDhcpFailureOk() (*SimpleAlertDhcpFailure, bool) {
	if o == nil || IsNil(o.DhcpFailure) {
		return nil, false
	}
	return o.DhcpFailure, true
}

// HasDhcpFailure returns a boolean if a field has been set.
func (o *SimpleAlert) HasDhcpFailure() bool {
	if o != nil && !IsNil(o.DhcpFailure) {
		return true
	}

	return false
}

// SetDhcpFailure gets a reference to the given SimpleAlertDhcpFailure and assigns it to the DhcpFailure field.
func (o *SimpleAlert) SetDhcpFailure(v SimpleAlertDhcpFailure) {
	o.DhcpFailure = &v
}

// GetDnsFailure returns the DnsFailure field value if set, zero value otherwise.
func (o *SimpleAlert) GetDnsFailure() SimpleAlertDnsFailure {
	if o == nil || IsNil(o.DnsFailure) {
		var ret SimpleAlertDnsFailure
		return ret
	}
	return *o.DnsFailure
}

// GetDnsFailureOk returns a tuple with the DnsFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleAlert) GetDnsFailureOk() (*SimpleAlertDnsFailure, bool) {
	if o == nil || IsNil(o.DnsFailure) {
		return nil, false
	}
	return o.DnsFailure, true
}

// HasDnsFailure returns a boolean if a field has been set.
func (o *SimpleAlert) HasDnsFailure() bool {
	if o != nil && !IsNil(o.DnsFailure) {
		return true
	}

	return false
}

// SetDnsFailure gets a reference to the given SimpleAlertDnsFailure and assigns it to the DnsFailure field.
func (o *SimpleAlert) SetDnsFailure(v SimpleAlertDnsFailure) {
	o.DnsFailure = &v
}

func (o SimpleAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArpFailure) {
		toSerialize["arp_failure"] = o.ArpFailure
	}
	if !IsNil(o.DhcpFailure) {
		toSerialize["dhcp_failure"] = o.DhcpFailure
	}
	if !IsNil(o.DnsFailure) {
		toSerialize["dns_failure"] = o.DnsFailure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimpleAlert) UnmarshalJSON(data []byte) (err error) {
	varSimpleAlert := _SimpleAlert{}

	err = json.Unmarshal(data, &varSimpleAlert)

	if err != nil {
		return err
	}

	*o = SimpleAlert(varSimpleAlert)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "arp_failure")
		delete(additionalProperties, "dhcp_failure")
		delete(additionalProperties, "dns_failure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimpleAlert struct {
	value *SimpleAlert
	isSet bool
}

func (v NullableSimpleAlert) Get() *SimpleAlert {
	return v.value
}

func (v *NullableSimpleAlert) Set(val *SimpleAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleAlert(val *SimpleAlert) *NullableSimpleAlert {
	return &NullableSimpleAlert{value: val, isSet: true}
}

func (v NullableSimpleAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


