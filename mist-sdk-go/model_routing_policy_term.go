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

// checks if the RoutingPolicyTerm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingPolicyTerm{}

// RoutingPolicyTerm struct for RoutingPolicyTerm
type RoutingPolicyTerm struct {
	Action *RoutingPolicyTermAction `json:"action,omitempty"`
	Matching *RoutingPolicyTermMatching `json:"matching,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingPolicyTerm RoutingPolicyTerm

// NewRoutingPolicyTerm instantiates a new RoutingPolicyTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingPolicyTerm() *RoutingPolicyTerm {
	this := RoutingPolicyTerm{}
	return &this
}

// NewRoutingPolicyTermWithDefaults instantiates a new RoutingPolicyTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingPolicyTermWithDefaults() *RoutingPolicyTerm {
	this := RoutingPolicyTerm{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RoutingPolicyTerm) GetAction() RoutingPolicyTermAction {
	if o == nil || IsNil(o.Action) {
		var ret RoutingPolicyTermAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingPolicyTerm) GetActionOk() (*RoutingPolicyTermAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RoutingPolicyTerm) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given RoutingPolicyTermAction and assigns it to the Action field.
func (o *RoutingPolicyTerm) SetAction(v RoutingPolicyTermAction) {
	o.Action = &v
}

// GetMatching returns the Matching field value if set, zero value otherwise.
func (o *RoutingPolicyTerm) GetMatching() RoutingPolicyTermMatching {
	if o == nil || IsNil(o.Matching) {
		var ret RoutingPolicyTermMatching
		return ret
	}
	return *o.Matching
}

// GetMatchingOk returns a tuple with the Matching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingPolicyTerm) GetMatchingOk() (*RoutingPolicyTermMatching, bool) {
	if o == nil || IsNil(o.Matching) {
		return nil, false
	}
	return o.Matching, true
}

// HasMatching returns a boolean if a field has been set.
func (o *RoutingPolicyTerm) HasMatching() bool {
	if o != nil && !IsNil(o.Matching) {
		return true
	}

	return false
}

// SetMatching gets a reference to the given RoutingPolicyTermMatching and assigns it to the Matching field.
func (o *RoutingPolicyTerm) SetMatching(v RoutingPolicyTermMatching) {
	o.Matching = &v
}

func (o RoutingPolicyTerm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingPolicyTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Matching) {
		toSerialize["matching"] = o.Matching
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingPolicyTerm) UnmarshalJSON(data []byte) (err error) {
	varRoutingPolicyTerm := _RoutingPolicyTerm{}

	err = json.Unmarshal(data, &varRoutingPolicyTerm)

	if err != nil {
		return err
	}

	*o = RoutingPolicyTerm(varRoutingPolicyTerm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "matching")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingPolicyTerm struct {
	value *RoutingPolicyTerm
	isSet bool
}

func (v NullableRoutingPolicyTerm) Get() *RoutingPolicyTerm {
	return v.value
}

func (v *NullableRoutingPolicyTerm) Set(val *RoutingPolicyTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingPolicyTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingPolicyTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingPolicyTerm(val *RoutingPolicyTerm) *NullableRoutingPolicyTerm {
	return &NullableRoutingPolicyTerm{value: val, isSet: true}
}

func (v NullableRoutingPolicyTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingPolicyTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


