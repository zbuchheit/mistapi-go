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

// checks if the SleImpactedSwitches type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SleImpactedSwitches{}

// SleImpactedSwitches struct for SleImpactedSwitches
type SleImpactedSwitches struct {
	Classifier *string `json:"classifier,omitempty"`
	End *int32 `json:"end,omitempty"`
	Failure *string `json:"failure,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Metric *string `json:"metric,omitempty"`
	Page *int32 `json:"page,omitempty"`
	Start *int32 `json:"start,omitempty"`
	Switches []SleImpactedSwitchesSwitch `json:"switches,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SleImpactedSwitches SleImpactedSwitches

// NewSleImpactedSwitches instantiates a new SleImpactedSwitches object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSleImpactedSwitches() *SleImpactedSwitches {
	this := SleImpactedSwitches{}
	return &this
}

// NewSleImpactedSwitchesWithDefaults instantiates a new SleImpactedSwitches object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSleImpactedSwitchesWithDefaults() *SleImpactedSwitches {
	this := SleImpactedSwitches{}
	return &this
}

// GetClassifier returns the Classifier field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetClassifier() string {
	if o == nil || IsNil(o.Classifier) {
		var ret string
		return ret
	}
	return *o.Classifier
}

// GetClassifierOk returns a tuple with the Classifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetClassifierOk() (*string, bool) {
	if o == nil || IsNil(o.Classifier) {
		return nil, false
	}
	return o.Classifier, true
}

// HasClassifier returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasClassifier() bool {
	if o != nil && !IsNil(o.Classifier) {
		return true
	}

	return false
}

// SetClassifier gets a reference to the given string and assigns it to the Classifier field.
func (o *SleImpactedSwitches) SetClassifier(v string) {
	o.Classifier = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetEnd() int32 {
	if o == nil || IsNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetEndOk() (*int32, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *SleImpactedSwitches) SetEnd(v int32) {
	o.End = &v
}

// GetFailure returns the Failure field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetFailure() string {
	if o == nil || IsNil(o.Failure) {
		var ret string
		return ret
	}
	return *o.Failure
}

// GetFailureOk returns a tuple with the Failure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetFailureOk() (*string, bool) {
	if o == nil || IsNil(o.Failure) {
		return nil, false
	}
	return o.Failure, true
}

// HasFailure returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasFailure() bool {
	if o != nil && !IsNil(o.Failure) {
		return true
	}

	return false
}

// SetFailure gets a reference to the given string and assigns it to the Failure field.
func (o *SleImpactedSwitches) SetFailure(v string) {
	o.Failure = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SleImpactedSwitches) SetLimit(v int32) {
	o.Limit = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetMetric() string {
	if o == nil || IsNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetMetricOk() (*string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *SleImpactedSwitches) SetMetric(v string) {
	o.Metric = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *SleImpactedSwitches) SetPage(v int32) {
	o.Page = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *SleImpactedSwitches) SetStart(v int32) {
	o.Start = &v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetSwitches() []SleImpactedSwitchesSwitch {
	if o == nil || IsNil(o.Switches) {
		var ret []SleImpactedSwitchesSwitch
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetSwitchesOk() ([]SleImpactedSwitchesSwitch, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []SleImpactedSwitchesSwitch and assigns it to the Switches field.
func (o *SleImpactedSwitches) SetSwitches(v []SleImpactedSwitchesSwitch) {
	o.Switches = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *SleImpactedSwitches) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SleImpactedSwitches) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *SleImpactedSwitches) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *SleImpactedSwitches) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o SleImpactedSwitches) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SleImpactedSwitches) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Classifier) {
		toSerialize["classifier"] = o.Classifier
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Failure) {
		toSerialize["failure"] = o.Failure
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SleImpactedSwitches) UnmarshalJSON(data []byte) (err error) {
	varSleImpactedSwitches := _SleImpactedSwitches{}

	err = json.Unmarshal(data, &varSleImpactedSwitches)

	if err != nil {
		return err
	}

	*o = SleImpactedSwitches(varSleImpactedSwitches)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "classifier")
		delete(additionalProperties, "end")
		delete(additionalProperties, "failure")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "metric")
		delete(additionalProperties, "page")
		delete(additionalProperties, "start")
		delete(additionalProperties, "switches")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSleImpactedSwitches struct {
	value *SleImpactedSwitches
	isSet bool
}

func (v NullableSleImpactedSwitches) Get() *SleImpactedSwitches {
	return v.value
}

func (v *NullableSleImpactedSwitches) Set(val *SleImpactedSwitches) {
	v.value = val
	v.isSet = true
}

func (v NullableSleImpactedSwitches) IsSet() bool {
	return v.isSet
}

func (v *NullableSleImpactedSwitches) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSleImpactedSwitches(val *SleImpactedSwitches) *NullableSleImpactedSwitches {
	return &NullableSleImpactedSwitches{value: val, isSet: true}
}

func (v NullableSleImpactedSwitches) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSleImpactedSwitches) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


