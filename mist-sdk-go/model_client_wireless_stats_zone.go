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
)

// checks if the ClientWirelessStatsZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientWirelessStatsZone{}

// ClientWirelessStatsZone struct for ClientWirelessStatsZone
type ClientWirelessStatsZone struct {
	Id *string `json:"id,omitempty"`
	Since *int32 `json:"since,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientWirelessStatsZone ClientWirelessStatsZone

// NewClientWirelessStatsZone instantiates a new ClientWirelessStatsZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientWirelessStatsZone() *ClientWirelessStatsZone {
	this := ClientWirelessStatsZone{}
	return &this
}

// NewClientWirelessStatsZoneWithDefaults instantiates a new ClientWirelessStatsZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWirelessStatsZoneWithDefaults() *ClientWirelessStatsZone {
	this := ClientWirelessStatsZone{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientWirelessStatsZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientWirelessStatsZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientWirelessStatsZone) SetId(v string) {
	o.Id = &v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *ClientWirelessStatsZone) GetSince() int32 {
	if o == nil || IsNil(o.Since) {
		var ret int32
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsZone) GetSinceOk() (*int32, bool) {
	if o == nil || IsNil(o.Since) {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *ClientWirelessStatsZone) HasSince() bool {
	if o != nil && !IsNil(o.Since) {
		return true
	}

	return false
}

// SetSince gets a reference to the given int32 and assigns it to the Since field.
func (o *ClientWirelessStatsZone) SetSince(v int32) {
	o.Since = &v
}

func (o ClientWirelessStatsZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientWirelessStatsZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Since) {
		toSerialize["since"] = o.Since
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientWirelessStatsZone) UnmarshalJSON(data []byte) (err error) {
	varClientWirelessStatsZone := _ClientWirelessStatsZone{}

	err = json.Unmarshal(data, &varClientWirelessStatsZone)

	if err != nil {
		return err
	}

	*o = ClientWirelessStatsZone(varClientWirelessStatsZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "since")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientWirelessStatsZone struct {
	value *ClientWirelessStatsZone
	isSet bool
}

func (v NullableClientWirelessStatsZone) Get() *ClientWirelessStatsZone {
	return v.value
}

func (v *NullableClientWirelessStatsZone) Set(val *ClientWirelessStatsZone) {
	v.value = val
	v.isSet = true
}

func (v NullableClientWirelessStatsZone) IsSet() bool {
	return v.isSet
}

func (v *NullableClientWirelessStatsZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientWirelessStatsZone(val *ClientWirelessStatsZone) *NullableClientWirelessStatsZone {
	return &NullableClientWirelessStatsZone{value: val, isSet: true}
}

func (v NullableClientWirelessStatsZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientWirelessStatsZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


