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

// checks if the SiteWidsRepeatedAuthFailures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteWidsRepeatedAuthFailures{}

// SiteWidsRepeatedAuthFailures struct for SiteWidsRepeatedAuthFailures
type SiteWidsRepeatedAuthFailures struct {
	// window where a trigger will be detected and action to be taken (in seconds)
	Duration *int32 `json:"duration,omitempty"`
	// count of events to trigger
	Threshold *int32 `json:"threshold,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteWidsRepeatedAuthFailures SiteWidsRepeatedAuthFailures

// NewSiteWidsRepeatedAuthFailures instantiates a new SiteWidsRepeatedAuthFailures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteWidsRepeatedAuthFailures() *SiteWidsRepeatedAuthFailures {
	this := SiteWidsRepeatedAuthFailures{}
	return &this
}

// NewSiteWidsRepeatedAuthFailuresWithDefaults instantiates a new SiteWidsRepeatedAuthFailures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteWidsRepeatedAuthFailuresWithDefaults() *SiteWidsRepeatedAuthFailures {
	this := SiteWidsRepeatedAuthFailures{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SiteWidsRepeatedAuthFailures) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWidsRepeatedAuthFailures) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SiteWidsRepeatedAuthFailures) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *SiteWidsRepeatedAuthFailures) SetDuration(v int32) {
	o.Duration = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *SiteWidsRepeatedAuthFailures) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteWidsRepeatedAuthFailures) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *SiteWidsRepeatedAuthFailures) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *SiteWidsRepeatedAuthFailures) SetThreshold(v int32) {
	o.Threshold = &v
}

func (o SiteWidsRepeatedAuthFailures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteWidsRepeatedAuthFailures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteWidsRepeatedAuthFailures) UnmarshalJSON(data []byte) (err error) {
	varSiteWidsRepeatedAuthFailures := _SiteWidsRepeatedAuthFailures{}

	err = json.Unmarshal(data, &varSiteWidsRepeatedAuthFailures)

	if err != nil {
		return err
	}

	*o = SiteWidsRepeatedAuthFailures(varSiteWidsRepeatedAuthFailures)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "threshold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteWidsRepeatedAuthFailures struct {
	value *SiteWidsRepeatedAuthFailures
	isSet bool
}

func (v NullableSiteWidsRepeatedAuthFailures) Get() *SiteWidsRepeatedAuthFailures {
	return v.value
}

func (v *NullableSiteWidsRepeatedAuthFailures) Set(val *SiteWidsRepeatedAuthFailures) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteWidsRepeatedAuthFailures) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteWidsRepeatedAuthFailures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteWidsRepeatedAuthFailures(val *SiteWidsRepeatedAuthFailures) *NullableSiteWidsRepeatedAuthFailures {
	return &NullableSiteWidsRepeatedAuthFailures{value: val, isSet: true}
}

func (v NullableSiteWidsRepeatedAuthFailures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteWidsRepeatedAuthFailures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


