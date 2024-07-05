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

// checks if the TunnelProviderOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunnelProviderOptions{}

// TunnelProviderOptions struct for TunnelProviderOptions
type TunnelProviderOptions struct {
	Jse *TunnelProviderOptionsJse `json:"jse,omitempty"`
	Zscaler *TunnelProviderOptionsZscaler `json:"zscaler,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelProviderOptions TunnelProviderOptions

// NewTunnelProviderOptions instantiates a new TunnelProviderOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelProviderOptions() *TunnelProviderOptions {
	this := TunnelProviderOptions{}
	return &this
}

// NewTunnelProviderOptionsWithDefaults instantiates a new TunnelProviderOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelProviderOptionsWithDefaults() *TunnelProviderOptions {
	this := TunnelProviderOptions{}
	return &this
}

// GetJse returns the Jse field value if set, zero value otherwise.
func (o *TunnelProviderOptions) GetJse() TunnelProviderOptionsJse {
	if o == nil || IsNil(o.Jse) {
		var ret TunnelProviderOptionsJse
		return ret
	}
	return *o.Jse
}

// GetJseOk returns a tuple with the Jse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptions) GetJseOk() (*TunnelProviderOptionsJse, bool) {
	if o == nil || IsNil(o.Jse) {
		return nil, false
	}
	return o.Jse, true
}

// HasJse returns a boolean if a field has been set.
func (o *TunnelProviderOptions) HasJse() bool {
	if o != nil && !IsNil(o.Jse) {
		return true
	}

	return false
}

// SetJse gets a reference to the given TunnelProviderOptionsJse and assigns it to the Jse field.
func (o *TunnelProviderOptions) SetJse(v TunnelProviderOptionsJse) {
	o.Jse = &v
}

// GetZscaler returns the Zscaler field value if set, zero value otherwise.
func (o *TunnelProviderOptions) GetZscaler() TunnelProviderOptionsZscaler {
	if o == nil || IsNil(o.Zscaler) {
		var ret TunnelProviderOptionsZscaler
		return ret
	}
	return *o.Zscaler
}

// GetZscalerOk returns a tuple with the Zscaler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelProviderOptions) GetZscalerOk() (*TunnelProviderOptionsZscaler, bool) {
	if o == nil || IsNil(o.Zscaler) {
		return nil, false
	}
	return o.Zscaler, true
}

// HasZscaler returns a boolean if a field has been set.
func (o *TunnelProviderOptions) HasZscaler() bool {
	if o != nil && !IsNil(o.Zscaler) {
		return true
	}

	return false
}

// SetZscaler gets a reference to the given TunnelProviderOptionsZscaler and assigns it to the Zscaler field.
func (o *TunnelProviderOptions) SetZscaler(v TunnelProviderOptionsZscaler) {
	o.Zscaler = &v
}

func (o TunnelProviderOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunnelProviderOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jse) {
		toSerialize["jse"] = o.Jse
	}
	if !IsNil(o.Zscaler) {
		toSerialize["zscaler"] = o.Zscaler
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TunnelProviderOptions) UnmarshalJSON(data []byte) (err error) {
	varTunnelProviderOptions := _TunnelProviderOptions{}

	err = json.Unmarshal(data, &varTunnelProviderOptions)

	if err != nil {
		return err
	}

	*o = TunnelProviderOptions(varTunnelProviderOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jse")
		delete(additionalProperties, "zscaler")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnelProviderOptions struct {
	value *TunnelProviderOptions
	isSet bool
}

func (v NullableTunnelProviderOptions) Get() *TunnelProviderOptions {
	return v.value
}

func (v *NullableTunnelProviderOptions) Set(val *TunnelProviderOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelProviderOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelProviderOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelProviderOptions(val *TunnelProviderOptions) *NullableTunnelProviderOptions {
	return &NullableTunnelProviderOptions{value: val, isSet: true}
}

func (v NullableTunnelProviderOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelProviderOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


