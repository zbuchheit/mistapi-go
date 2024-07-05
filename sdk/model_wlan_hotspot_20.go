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
)

// checks if the WlanHotspot20 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanHotspot20{}

// WlanHotspot20 hostspot 2.0 wlan settings
type WlanHotspot20 struct {
	DomainName []string `json:"domain_name,omitempty"`
	// whether to enable hotspot 2.0 config
	Enabled *bool `json:"enabled,omitempty"`
	NaiRealms []string `json:"nai_realms,omitempty"`
	// list of operators to support
	Operators []WlanHotspot20OperatorsItem `json:"operators,omitempty"`
	Rcoi []string `json:"rcoi,omitempty"`
	// venue name, default is site name
	VenueName *string `json:"venue_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WlanHotspot20 WlanHotspot20

// NewWlanHotspot20 instantiates a new WlanHotspot20 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanHotspot20() *WlanHotspot20 {
	this := WlanHotspot20{}
	return &this
}

// NewWlanHotspot20WithDefaults instantiates a new WlanHotspot20 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanHotspot20WithDefaults() *WlanHotspot20 {
	this := WlanHotspot20{}
	return &this
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *WlanHotspot20) GetDomainName() []string {
	if o == nil || IsNil(o.DomainName) {
		var ret []string
		return ret
	}
	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanHotspot20) GetDomainNameOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *WlanHotspot20) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given []string and assigns it to the DomainName field.
func (o *WlanHotspot20) SetDomainName(v []string) {
	o.DomainName = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WlanHotspot20) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanHotspot20) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WlanHotspot20) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WlanHotspot20) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNaiRealms returns the NaiRealms field value if set, zero value otherwise.
func (o *WlanHotspot20) GetNaiRealms() []string {
	if o == nil || IsNil(o.NaiRealms) {
		var ret []string
		return ret
	}
	return o.NaiRealms
}

// GetNaiRealmsOk returns a tuple with the NaiRealms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanHotspot20) GetNaiRealmsOk() ([]string, bool) {
	if o == nil || IsNil(o.NaiRealms) {
		return nil, false
	}
	return o.NaiRealms, true
}

// HasNaiRealms returns a boolean if a field has been set.
func (o *WlanHotspot20) HasNaiRealms() bool {
	if o != nil && !IsNil(o.NaiRealms) {
		return true
	}

	return false
}

// SetNaiRealms gets a reference to the given []string and assigns it to the NaiRealms field.
func (o *WlanHotspot20) SetNaiRealms(v []string) {
	o.NaiRealms = v
}

// GetOperators returns the Operators field value if set, zero value otherwise.
func (o *WlanHotspot20) GetOperators() []WlanHotspot20OperatorsItem {
	if o == nil || IsNil(o.Operators) {
		var ret []WlanHotspot20OperatorsItem
		return ret
	}
	return o.Operators
}

// GetOperatorsOk returns a tuple with the Operators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanHotspot20) GetOperatorsOk() ([]WlanHotspot20OperatorsItem, bool) {
	if o == nil || IsNil(o.Operators) {
		return nil, false
	}
	return o.Operators, true
}

// HasOperators returns a boolean if a field has been set.
func (o *WlanHotspot20) HasOperators() bool {
	if o != nil && !IsNil(o.Operators) {
		return true
	}

	return false
}

// SetOperators gets a reference to the given []WlanHotspot20OperatorsItem and assigns it to the Operators field.
func (o *WlanHotspot20) SetOperators(v []WlanHotspot20OperatorsItem) {
	o.Operators = v
}

// GetRcoi returns the Rcoi field value if set, zero value otherwise.
func (o *WlanHotspot20) GetRcoi() []string {
	if o == nil || IsNil(o.Rcoi) {
		var ret []string
		return ret
	}
	return o.Rcoi
}

// GetRcoiOk returns a tuple with the Rcoi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanHotspot20) GetRcoiOk() ([]string, bool) {
	if o == nil || IsNil(o.Rcoi) {
		return nil, false
	}
	return o.Rcoi, true
}

// HasRcoi returns a boolean if a field has been set.
func (o *WlanHotspot20) HasRcoi() bool {
	if o != nil && !IsNil(o.Rcoi) {
		return true
	}

	return false
}

// SetRcoi gets a reference to the given []string and assigns it to the Rcoi field.
func (o *WlanHotspot20) SetRcoi(v []string) {
	o.Rcoi = v
}

// GetVenueName returns the VenueName field value if set, zero value otherwise.
func (o *WlanHotspot20) GetVenueName() string {
	if o == nil || IsNil(o.VenueName) {
		var ret string
		return ret
	}
	return *o.VenueName
}

// GetVenueNameOk returns a tuple with the VenueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanHotspot20) GetVenueNameOk() (*string, bool) {
	if o == nil || IsNil(o.VenueName) {
		return nil, false
	}
	return o.VenueName, true
}

// HasVenueName returns a boolean if a field has been set.
func (o *WlanHotspot20) HasVenueName() bool {
	if o != nil && !IsNil(o.VenueName) {
		return true
	}

	return false
}

// SetVenueName gets a reference to the given string and assigns it to the VenueName field.
func (o *WlanHotspot20) SetVenueName(v string) {
	o.VenueName = &v
}

func (o WlanHotspot20) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanHotspot20) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DomainName) {
		toSerialize["domain_name"] = o.DomainName
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.NaiRealms) {
		toSerialize["nai_realms"] = o.NaiRealms
	}
	if !IsNil(o.Operators) {
		toSerialize["operators"] = o.Operators
	}
	if !IsNil(o.Rcoi) {
		toSerialize["rcoi"] = o.Rcoi
	}
	if !IsNil(o.VenueName) {
		toSerialize["venue_name"] = o.VenueName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WlanHotspot20) UnmarshalJSON(data []byte) (err error) {
	varWlanHotspot20 := _WlanHotspot20{}

	err = json.Unmarshal(data, &varWlanHotspot20)

	if err != nil {
		return err
	}

	*o = WlanHotspot20(varWlanHotspot20)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domain_name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "nai_realms")
		delete(additionalProperties, "operators")
		delete(additionalProperties, "rcoi")
		delete(additionalProperties, "venue_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWlanHotspot20 struct {
	value *WlanHotspot20
	isSet bool
}

func (v NullableWlanHotspot20) Get() *WlanHotspot20 {
	return v.value
}

func (v *NullableWlanHotspot20) Set(val *WlanHotspot20) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanHotspot20) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanHotspot20) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanHotspot20(val *WlanHotspot20) *NullableWlanHotspot20 {
	return &NullableWlanHotspot20{value: val, isSet: true}
}

func (v NullableWlanHotspot20) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanHotspot20) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


