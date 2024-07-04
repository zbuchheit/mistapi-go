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

// checks if the RoutingPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingPolicy{}

// RoutingPolicy struct for RoutingPolicy
type RoutingPolicy struct {
	// zero or more criteria/filter can be specified to match the term, all criteria have to be met
	Terms []RoutingPolicyTerm `json:"terms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingPolicy RoutingPolicy

// NewRoutingPolicy instantiates a new RoutingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingPolicy() *RoutingPolicy {
	this := RoutingPolicy{}
	return &this
}

// NewRoutingPolicyWithDefaults instantiates a new RoutingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingPolicyWithDefaults() *RoutingPolicy {
	this := RoutingPolicy{}
	return &this
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *RoutingPolicy) GetTerms() []RoutingPolicyTerm {
	if o == nil || IsNil(o.Terms) {
		var ret []RoutingPolicyTerm
		return ret
	}
	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingPolicy) GetTermsOk() ([]RoutingPolicyTerm, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *RoutingPolicy) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given []RoutingPolicyTerm and assigns it to the Terms field.
func (o *RoutingPolicy) SetTerms(v []RoutingPolicyTerm) {
	o.Terms = v
}

func (o RoutingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingPolicy) UnmarshalJSON(data []byte) (err error) {
	varRoutingPolicy := _RoutingPolicy{}

	err = json.Unmarshal(data, &varRoutingPolicy)

	if err != nil {
		return err
	}

	*o = RoutingPolicy(varRoutingPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "terms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingPolicy struct {
	value *RoutingPolicy
	isSet bool
}

func (v NullableRoutingPolicy) Get() *RoutingPolicy {
	return v.value
}

func (v *NullableRoutingPolicy) Set(val *RoutingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingPolicy(val *RoutingPolicy) *NullableRoutingPolicy {
	return &NullableRoutingPolicy{value: val, isSet: true}
}

func (v NullableRoutingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


