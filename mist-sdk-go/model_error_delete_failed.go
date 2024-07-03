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

// checks if the ErrorDeleteFailed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDeleteFailed{}

// ErrorDeleteFailed struct for ErrorDeleteFailed
type ErrorDeleteFailed struct {
	Detail string `json:"detail"`
	OrgId string `json:"org_id"`
	AdditionalProperties map[string]interface{}
}

type _ErrorDeleteFailed ErrorDeleteFailed

// NewErrorDeleteFailed instantiates a new ErrorDeleteFailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDeleteFailed(detail string, orgId string) *ErrorDeleteFailed {
	this := ErrorDeleteFailed{}
	this.Detail = detail
	this.OrgId = orgId
	return &this
}

// NewErrorDeleteFailedWithDefaults instantiates a new ErrorDeleteFailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDeleteFailedWithDefaults() *ErrorDeleteFailed {
	this := ErrorDeleteFailed{}
	return &this
}

// GetDetail returns the Detail field value
func (o *ErrorDeleteFailed) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorDeleteFailed) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ErrorDeleteFailed) SetDetail(v string) {
	o.Detail = v
}

// GetOrgId returns the OrgId field value
func (o *ErrorDeleteFailed) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ErrorDeleteFailed) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ErrorDeleteFailed) SetOrgId(v string) {
	o.OrgId = v
}

func (o ErrorDeleteFailed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDeleteFailed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	toSerialize["org_id"] = o.OrgId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorDeleteFailed) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"detail",
		"org_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorDeleteFailed := _ErrorDeleteFailed{}

	err = json.Unmarshal(data, &varErrorDeleteFailed)

	if err != nil {
		return err
	}

	*o = ErrorDeleteFailed(varErrorDeleteFailed)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "org_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorDeleteFailed struct {
	value *ErrorDeleteFailed
	isSet bool
}

func (v NullableErrorDeleteFailed) Get() *ErrorDeleteFailed {
	return v.value
}

func (v *NullableErrorDeleteFailed) Set(val *ErrorDeleteFailed) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDeleteFailed) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDeleteFailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDeleteFailed(val *ErrorDeleteFailed) *NullableErrorDeleteFailed {
	return &NullableErrorDeleteFailed{value: val, isSet: true}
}

func (v NullableErrorDeleteFailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDeleteFailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


