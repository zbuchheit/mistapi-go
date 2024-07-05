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

// checks if the SiteZoneOccupancyAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteZoneOccupancyAlert{}

// SiteZoneOccupancyAlert Zone Occupancy alert site settings
type SiteZoneOccupancyAlert struct {
	// list of email addresses to send email notifications when the alert threshold is reached
	EmailNotifiers []string `json:"email_notifiers,omitempty"`
	// indicate whether zone occupancy alert is enabled for the site
	Enabled *bool `json:"enabled,omitempty"`
	// sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy > occupancy_limit) for a minimum duration specified in the threshold, in minutes
	Threshold *int32 `json:"threshold,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteZoneOccupancyAlert SiteZoneOccupancyAlert

// NewSiteZoneOccupancyAlert instantiates a new SiteZoneOccupancyAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteZoneOccupancyAlert() *SiteZoneOccupancyAlert {
	this := SiteZoneOccupancyAlert{}
	var enabled bool = false
	this.Enabled = &enabled
	var threshold int32 = 5
	this.Threshold = &threshold
	return &this
}

// NewSiteZoneOccupancyAlertWithDefaults instantiates a new SiteZoneOccupancyAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteZoneOccupancyAlertWithDefaults() *SiteZoneOccupancyAlert {
	this := SiteZoneOccupancyAlert{}
	var enabled bool = false
	this.Enabled = &enabled
	var threshold int32 = 5
	this.Threshold = &threshold
	return &this
}

// GetEmailNotifiers returns the EmailNotifiers field value if set, zero value otherwise.
func (o *SiteZoneOccupancyAlert) GetEmailNotifiers() []string {
	if o == nil || IsNil(o.EmailNotifiers) {
		var ret []string
		return ret
	}
	return o.EmailNotifiers
}

// GetEmailNotifiersOk returns a tuple with the EmailNotifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteZoneOccupancyAlert) GetEmailNotifiersOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailNotifiers) {
		return nil, false
	}
	return o.EmailNotifiers, true
}

// HasEmailNotifiers returns a boolean if a field has been set.
func (o *SiteZoneOccupancyAlert) HasEmailNotifiers() bool {
	if o != nil && !IsNil(o.EmailNotifiers) {
		return true
	}

	return false
}

// SetEmailNotifiers gets a reference to the given []string and assigns it to the EmailNotifiers field.
func (o *SiteZoneOccupancyAlert) SetEmailNotifiers(v []string) {
	o.EmailNotifiers = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SiteZoneOccupancyAlert) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteZoneOccupancyAlert) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SiteZoneOccupancyAlert) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SiteZoneOccupancyAlert) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *SiteZoneOccupancyAlert) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteZoneOccupancyAlert) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *SiteZoneOccupancyAlert) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *SiteZoneOccupancyAlert) SetThreshold(v int32) {
	o.Threshold = &v
}

func (o SiteZoneOccupancyAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteZoneOccupancyAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailNotifiers) {
		toSerialize["email_notifiers"] = o.EmailNotifiers
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteZoneOccupancyAlert) UnmarshalJSON(data []byte) (err error) {
	varSiteZoneOccupancyAlert := _SiteZoneOccupancyAlert{}

	err = json.Unmarshal(data, &varSiteZoneOccupancyAlert)

	if err != nil {
		return err
	}

	*o = SiteZoneOccupancyAlert(varSiteZoneOccupancyAlert)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email_notifiers")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "threshold")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteZoneOccupancyAlert struct {
	value *SiteZoneOccupancyAlert
	isSet bool
}

func (v NullableSiteZoneOccupancyAlert) Get() *SiteZoneOccupancyAlert {
	return v.value
}

func (v *NullableSiteZoneOccupancyAlert) Set(val *SiteZoneOccupancyAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteZoneOccupancyAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteZoneOccupancyAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteZoneOccupancyAlert(val *SiteZoneOccupancyAlert) *NullableSiteZoneOccupancyAlert {
	return &NullableSiteZoneOccupancyAlert{value: val, isSet: true}
}

func (v NullableSiteZoneOccupancyAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteZoneOccupancyAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


