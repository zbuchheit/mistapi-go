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

// checks if the Anomaly type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Anomaly{}

// Anomaly Anomaly
type Anomaly struct {
	Events []string `json:"events"`
	Since *float32 `json:"since,omitempty"`
	SleBaseline float32 `json:"sle_baseline"`
	SleDeviation float32 `json:"sle_deviation"`
	Timestamp float32 `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _Anomaly Anomaly

// NewAnomaly instantiates a new Anomaly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnomaly(events []string, sleBaseline float32, sleDeviation float32, timestamp float32) *Anomaly {
	this := Anomaly{}
	this.Events = events
	this.SleBaseline = sleBaseline
	this.SleDeviation = sleDeviation
	this.Timestamp = timestamp
	return &this
}

// NewAnomalyWithDefaults instantiates a new Anomaly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnomalyWithDefaults() *Anomaly {
	this := Anomaly{}
	return &this
}

// GetEvents returns the Events field value
func (o *Anomaly) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *Anomaly) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *Anomaly) SetEvents(v []string) {
	o.Events = v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *Anomaly) GetSince() float32 {
	if o == nil || IsNil(o.Since) {
		var ret float32
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Anomaly) GetSinceOk() (*float32, bool) {
	if o == nil || IsNil(o.Since) {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *Anomaly) HasSince() bool {
	if o != nil && !IsNil(o.Since) {
		return true
	}

	return false
}

// SetSince gets a reference to the given float32 and assigns it to the Since field.
func (o *Anomaly) SetSince(v float32) {
	o.Since = &v
}

// GetSleBaseline returns the SleBaseline field value
func (o *Anomaly) GetSleBaseline() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SleBaseline
}

// GetSleBaselineOk returns a tuple with the SleBaseline field value
// and a boolean to check if the value has been set.
func (o *Anomaly) GetSleBaselineOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SleBaseline, true
}

// SetSleBaseline sets field value
func (o *Anomaly) SetSleBaseline(v float32) {
	o.SleBaseline = v
}

// GetSleDeviation returns the SleDeviation field value
func (o *Anomaly) GetSleDeviation() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SleDeviation
}

// GetSleDeviationOk returns a tuple with the SleDeviation field value
// and a boolean to check if the value has been set.
func (o *Anomaly) GetSleDeviationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SleDeviation, true
}

// SetSleDeviation sets field value
func (o *Anomaly) SetSleDeviation(v float32) {
	o.SleDeviation = v
}

// GetTimestamp returns the Timestamp field value
func (o *Anomaly) GetTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Anomaly) GetTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Anomaly) SetTimestamp(v float32) {
	o.Timestamp = v
}

func (o Anomaly) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Anomaly) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	if !IsNil(o.Since) {
		toSerialize["since"] = o.Since
	}
	toSerialize["sle_baseline"] = o.SleBaseline
	toSerialize["sle_deviation"] = o.SleDeviation
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Anomaly) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
		"sle_baseline",
		"sle_deviation",
		"timestamp",
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

	varAnomaly := _Anomaly{}

	err = json.Unmarshal(data, &varAnomaly)

	if err != nil {
		return err
	}

	*o = Anomaly(varAnomaly)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "events")
		delete(additionalProperties, "since")
		delete(additionalProperties, "sle_baseline")
		delete(additionalProperties, "sle_deviation")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnomaly struct {
	value *Anomaly
	isSet bool
}

func (v NullableAnomaly) Get() *Anomaly {
	return v.value
}

func (v *NullableAnomaly) Set(val *Anomaly) {
	v.value = val
	v.isSet = true
}

func (v NullableAnomaly) IsSet() bool {
	return v.isSet
}

func (v *NullableAnomaly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnomaly(val *Anomaly) *NullableAnomaly {
	return &NullableAnomaly{value: val, isSet: true}
}

func (v NullableAnomaly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnomaly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


