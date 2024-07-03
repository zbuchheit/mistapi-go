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

// checks if the ResponseClaimLicenseLicenseErrorItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseClaimLicenseLicenseErrorItem{}

// ResponseClaimLicenseLicenseErrorItem struct for ResponseClaimLicenseLicenseErrorItem
type ResponseClaimLicenseLicenseErrorItem struct {
	Order string `json:"order"`
	Reason string `json:"reason"`
	AdditionalProperties map[string]interface{}
}

type _ResponseClaimLicenseLicenseErrorItem ResponseClaimLicenseLicenseErrorItem

// NewResponseClaimLicenseLicenseErrorItem instantiates a new ResponseClaimLicenseLicenseErrorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseClaimLicenseLicenseErrorItem(order string, reason string) *ResponseClaimLicenseLicenseErrorItem {
	this := ResponseClaimLicenseLicenseErrorItem{}
	this.Order = order
	this.Reason = reason
	return &this
}

// NewResponseClaimLicenseLicenseErrorItemWithDefaults instantiates a new ResponseClaimLicenseLicenseErrorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseClaimLicenseLicenseErrorItemWithDefaults() *ResponseClaimLicenseLicenseErrorItem {
	this := ResponseClaimLicenseLicenseErrorItem{}
	return &this
}

// GetOrder returns the Order field value
func (o *ResponseClaimLicenseLicenseErrorItem) GetOrder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ResponseClaimLicenseLicenseErrorItem) GetOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ResponseClaimLicenseLicenseErrorItem) SetOrder(v string) {
	o.Order = v
}

// GetReason returns the Reason field value
func (o *ResponseClaimLicenseLicenseErrorItem) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ResponseClaimLicenseLicenseErrorItem) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ResponseClaimLicenseLicenseErrorItem) SetReason(v string) {
	o.Reason = v
}

func (o ResponseClaimLicenseLicenseErrorItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseClaimLicenseLicenseErrorItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order"] = o.Order
	toSerialize["reason"] = o.Reason

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseClaimLicenseLicenseErrorItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order",
		"reason",
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

	varResponseClaimLicenseLicenseErrorItem := _ResponseClaimLicenseLicenseErrorItem{}

	err = json.Unmarshal(data, &varResponseClaimLicenseLicenseErrorItem)

	if err != nil {
		return err
	}

	*o = ResponseClaimLicenseLicenseErrorItem(varResponseClaimLicenseLicenseErrorItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "order")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseClaimLicenseLicenseErrorItem struct {
	value *ResponseClaimLicenseLicenseErrorItem
	isSet bool
}

func (v NullableResponseClaimLicenseLicenseErrorItem) Get() *ResponseClaimLicenseLicenseErrorItem {
	return v.value
}

func (v *NullableResponseClaimLicenseLicenseErrorItem) Set(val *ResponseClaimLicenseLicenseErrorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseClaimLicenseLicenseErrorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseClaimLicenseLicenseErrorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseClaimLicenseLicenseErrorItem(val *ResponseClaimLicenseLicenseErrorItem) *NullableResponseClaimLicenseLicenseErrorItem {
	return &NullableResponseClaimLicenseLicenseErrorItem{value: val, isSet: true}
}

func (v NullableResponseClaimLicenseLicenseErrorItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseClaimLicenseLicenseErrorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


