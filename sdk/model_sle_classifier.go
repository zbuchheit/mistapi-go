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

// checks if the SleClassifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleClassifier{}

// SleClassifier struct for SleClassifier
type SleClassifier struct {
	Impact SleClassifierImpact `json:"impact"`
	Interval float32 `json:"interval"`
	Name string `json:"name"`
	Samples *SleClassifierSamples `json:"samples,omitempty"`
	XLabel string `json:"x_label"`
	YLabel string `json:"y_label"`
	AdditionalProperties map[string]interface{}
}

type _SleClassifier SleClassifier

// NewSleClassifier instantiates a new SleClassifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleClassifier(impact SleClassifierImpact, interval float32, name string, xLabel string, yLabel string) *SleClassifier {
	this := SleClassifier{}
	this.Impact = impact
	this.Interval = interval
	this.Name = name
	this.XLabel = xLabel
	this.YLabel = yLabel
	return &this
}

// NewSleClassifierWithDefaults instantiates a new SleClassifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleClassifierWithDefaults() *SleClassifier {
	this := SleClassifier{}
	return &this
}

// GetImpact returns the Impact field value
func (o *SleClassifier) GetImpact() SleClassifierImpact {
	if o == nil {
		var ret SleClassifierImpact
		return ret
	}

	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value
// and a boolean to check if the value has been set.
func (o *SleClassifier) GetImpactOk() (*SleClassifierImpact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impact, true
}

// SetImpact sets field value
func (o *SleClassifier) SetImpact(v SleClassifierImpact) {
	o.Impact = v
}

// GetInterval returns the Interval field value
func (o *SleClassifier) GetInterval() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *SleClassifier) GetIntervalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *SleClassifier) SetInterval(v float32) {
	o.Interval = v
}

// GetName returns the Name field value
func (o *SleClassifier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SleClassifier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SleClassifier) SetName(v string) {
	o.Name = v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *SleClassifier) GetSamples() SleClassifierSamples {
	if o == nil || IsNil(o.Samples) {
		var ret SleClassifierSamples
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleClassifier) GetSamplesOk() (*SleClassifierSamples, bool) {
	if o == nil || IsNil(o.Samples) {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *SleClassifier) HasSamples() bool {
	if o != nil && !IsNil(o.Samples) {
		return true
	}

	return false
}

// SetSamples gets a reference to the given SleClassifierSamples and assigns it to the Samples field.
func (o *SleClassifier) SetSamples(v SleClassifierSamples) {
	o.Samples = &v
}

// GetXLabel returns the XLabel field value
func (o *SleClassifier) GetXLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XLabel
}

// GetXLabelOk returns a tuple with the XLabel field value
// and a boolean to check if the value has been set.
func (o *SleClassifier) GetXLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XLabel, true
}

// SetXLabel sets field value
func (o *SleClassifier) SetXLabel(v string) {
	o.XLabel = v
}

// GetYLabel returns the YLabel field value
func (o *SleClassifier) GetYLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YLabel
}

// GetYLabelOk returns a tuple with the YLabel field value
// and a boolean to check if the value has been set.
func (o *SleClassifier) GetYLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YLabel, true
}

// SetYLabel sets field value
func (o *SleClassifier) SetYLabel(v string) {
	o.YLabel = v
}

func (o SleClassifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleClassifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["impact"] = o.Impact
	toSerialize["interval"] = o.Interval
	toSerialize["name"] = o.Name
	if !IsNil(o.Samples) {
		toSerialize["samples"] = o.Samples
	}
	toSerialize["x_label"] = o.XLabel
	toSerialize["y_label"] = o.YLabel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleClassifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"impact",
		"interval",
		"name",
		"x_label",
		"y_label",
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

	varSleClassifier := _SleClassifier{}

	err = json.Unmarshal(data, &varSleClassifier)

	if err != nil {
		return err
	}

	*o = SleClassifier(varSleClassifier)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "impact")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "name")
		delete(additionalProperties, "samples")
		delete(additionalProperties, "x_label")
		delete(additionalProperties, "y_label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleClassifier struct {
	value *SleClassifier
	isSet bool
}

func (v NullableSleClassifier) Get() *SleClassifier {
	return v.value
}

func (v *NullableSleClassifier) Set(val *SleClassifier) {
	v.value = val
	v.isSet = true
}

func (v NullableSleClassifier) IsSet() bool {
	return v.isSet
}

func (v *NullableSleClassifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleClassifier(val *SleClassifier) *NullableSleClassifier {
	return &NullableSleClassifier{value: val, isSet: true}
}

func (v NullableSleClassifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleClassifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


