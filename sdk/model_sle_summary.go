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
	"fmt"
)

// checks if the SleSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleSummary{}

// SleSummary struct for SleSummary
type SleSummary struct {
	Classifiers []SleClassifier `json:"classifiers"`
	End float32 `json:"end"`
	Events []map[string]interface{} `json:"events"`
	Impact SleSummaryImpact `json:"impact"`
	Sle SleSummarySle `json:"sle"`
	Start float32 `json:"start"`
	AdditionalProperties map[string]interface{}
}

type _SleSummary SleSummary

// NewSleSummary instantiates a new SleSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleSummary(classifiers []SleClassifier, end float32, events []map[string]interface{}, impact SleSummaryImpact, sle SleSummarySle, start float32) *SleSummary {
	this := SleSummary{}
	this.Classifiers = classifiers
	this.End = end
	this.Events = events
	this.Impact = impact
	this.Sle = sle
	this.Start = start
	return &this
}

// NewSleSummaryWithDefaults instantiates a new SleSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleSummaryWithDefaults() *SleSummary {
	this := SleSummary{}
	return &this
}

// GetClassifiers returns the Classifiers field value
func (o *SleSummary) GetClassifiers() []SleClassifier {
	if o == nil {
		var ret []SleClassifier
		return ret
	}

	return o.Classifiers
}

// GetClassifiersOk returns a tuple with the Classifiers field value
// and a boolean to check if the value has been set.
func (o *SleSummary) GetClassifiersOk() ([]SleClassifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.Classifiers, true
}

// SetClassifiers sets field value
func (o *SleSummary) SetClassifiers(v []SleClassifier) {
	o.Classifiers = v
}

// GetEnd returns the End field value
func (o *SleSummary) GetEnd() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *SleSummary) GetEndOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *SleSummary) SetEnd(v float32) {
	o.End = v
}

// GetEvents returns the Events field value
func (o *SleSummary) GetEvents() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *SleSummary) GetEventsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *SleSummary) SetEvents(v []map[string]interface{}) {
	o.Events = v
}

// GetImpact returns the Impact field value
func (o *SleSummary) GetImpact() SleSummaryImpact {
	if o == nil {
		var ret SleSummaryImpact
		return ret
	}

	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value
// and a boolean to check if the value has been set.
func (o *SleSummary) GetImpactOk() (*SleSummaryImpact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impact, true
}

// SetImpact sets field value
func (o *SleSummary) SetImpact(v SleSummaryImpact) {
	o.Impact = v
}

// GetSle returns the Sle field value
func (o *SleSummary) GetSle() SleSummarySle {
	if o == nil {
		var ret SleSummarySle
		return ret
	}

	return o.Sle
}

// GetSleOk returns a tuple with the Sle field value
// and a boolean to check if the value has been set.
func (o *SleSummary) GetSleOk() (*SleSummarySle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sle, true
}

// SetSle sets field value
func (o *SleSummary) SetSle(v SleSummarySle) {
	o.Sle = v
}

// GetStart returns the Start field value
func (o *SleSummary) GetStart() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SleSummary) GetStartOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *SleSummary) SetStart(v float32) {
	o.Start = v
}

func (o SleSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["classifiers"] = o.Classifiers
	toSerialize["end"] = o.End
	toSerialize["events"] = o.Events
	toSerialize["impact"] = o.Impact
	toSerialize["sle"] = o.Sle
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"classifiers",
		"end",
		"events",
		"impact",
		"sle",
		"start",
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

	varSleSummary := _SleSummary{}

	err = json.Unmarshal(data, &varSleSummary)

	if err != nil {
		return err
	}

	*o = SleSummary(varSleSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "classifiers")
		delete(additionalProperties, "end")
		delete(additionalProperties, "events")
		delete(additionalProperties, "impact")
		delete(additionalProperties, "sle")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleSummary struct {
	value *SleSummary
	isSet bool
}

func (v NullableSleSummary) Get() *SleSummary {
	return v.value
}

func (v *NullableSleSummary) Set(val *SleSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableSleSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSleSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleSummary(val *SleSummary) *NullableSleSummary {
	return &NullableSleSummary{value: val, isSet: true}
}

func (v NullableSleSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


