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

// checks if the RoutingPolicyTermMatchingRouteExists type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingPolicyTermMatchingRouteExists{}

// RoutingPolicyTermMatchingRouteExists struct for RoutingPolicyTermMatchingRouteExists
type RoutingPolicyTermMatchingRouteExists struct {
	Route *string `json:"route,omitempty"`
	// name of the vrf instance it can also be the name of the VPN or wan if they
	VrfName *string `json:"vrf_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingPolicyTermMatchingRouteExists RoutingPolicyTermMatchingRouteExists

// NewRoutingPolicyTermMatchingRouteExists instantiates a new RoutingPolicyTermMatchingRouteExists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingPolicyTermMatchingRouteExists() *RoutingPolicyTermMatchingRouteExists {
	this := RoutingPolicyTermMatchingRouteExists{}
	var vrfName string = "default"
	this.VrfName = &vrfName
	return &this
}

// NewRoutingPolicyTermMatchingRouteExistsWithDefaults instantiates a new RoutingPolicyTermMatchingRouteExists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingPolicyTermMatchingRouteExistsWithDefaults() *RoutingPolicyTermMatchingRouteExists {
	this := RoutingPolicyTermMatchingRouteExists{}
	var vrfName string = "default"
	this.VrfName = &vrfName
	return &this
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *RoutingPolicyTermMatchingRouteExists) GetRoute() string {
	if o == nil || IsNil(o.Route) {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingPolicyTermMatchingRouteExists) GetRouteOk() (*string, bool) {
	if o == nil || IsNil(o.Route) {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *RoutingPolicyTermMatchingRouteExists) HasRoute() bool {
	if o != nil && !IsNil(o.Route) {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *RoutingPolicyTermMatchingRouteExists) SetRoute(v string) {
	o.Route = &v
}

// GetVrfName returns the VrfName field value if set, zero value otherwise.
func (o *RoutingPolicyTermMatchingRouteExists) GetVrfName() string {
	if o == nil || IsNil(o.VrfName) {
		var ret string
		return ret
	}
	return *o.VrfName
}

// GetVrfNameOk returns a tuple with the VrfName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingPolicyTermMatchingRouteExists) GetVrfNameOk() (*string, bool) {
	if o == nil || IsNil(o.VrfName) {
		return nil, false
	}
	return o.VrfName, true
}

// HasVrfName returns a boolean if a field has been set.
func (o *RoutingPolicyTermMatchingRouteExists) HasVrfName() bool {
	if o != nil && !IsNil(o.VrfName) {
		return true
	}

	return false
}

// SetVrfName gets a reference to the given string and assigns it to the VrfName field.
func (o *RoutingPolicyTermMatchingRouteExists) SetVrfName(v string) {
	o.VrfName = &v
}

func (o RoutingPolicyTermMatchingRouteExists) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingPolicyTermMatchingRouteExists) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Route) {
		toSerialize["route"] = o.Route
	}
	if !IsNil(o.VrfName) {
		toSerialize["vrf_name"] = o.VrfName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingPolicyTermMatchingRouteExists) UnmarshalJSON(data []byte) (err error) {
	varRoutingPolicyTermMatchingRouteExists := _RoutingPolicyTermMatchingRouteExists{}

	err = json.Unmarshal(data, &varRoutingPolicyTermMatchingRouteExists)

	if err != nil {
		return err
	}

	*o = RoutingPolicyTermMatchingRouteExists(varRoutingPolicyTermMatchingRouteExists)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "route")
		delete(additionalProperties, "vrf_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingPolicyTermMatchingRouteExists struct {
	value *RoutingPolicyTermMatchingRouteExists
	isSet bool
}

func (v NullableRoutingPolicyTermMatchingRouteExists) Get() *RoutingPolicyTermMatchingRouteExists {
	return v.value
}

func (v *NullableRoutingPolicyTermMatchingRouteExists) Set(val *RoutingPolicyTermMatchingRouteExists) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingPolicyTermMatchingRouteExists) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingPolicyTermMatchingRouteExists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingPolicyTermMatchingRouteExists(val *RoutingPolicyTermMatchingRouteExists) *NullableRoutingPolicyTermMatchingRouteExists {
	return &NullableRoutingPolicyTermMatchingRouteExists{value: val, isSet: true}
}

func (v NullableRoutingPolicyTermMatchingRouteExists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingPolicyTermMatchingRouteExists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


