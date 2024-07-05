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

// checks if the ResponseDeviceUpgradeCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeviceUpgradeCounts{}

// ResponseDeviceUpgradeCounts struct for ResponseDeviceUpgradeCounts
type ResponseDeviceUpgradeCounts struct {
	// list of devices MAC Addresses which cloud has requested to download firmware
	DownloadRequested []string `json:"download_requested,omitempty"`
	// list of devices MAC Addresses which have the firmware downloaded
	Downloaded []string `json:"downloaded,omitempty"`
	// list of devices MAC Addresses which have failed to upgrade
	Failed []string `json:"failed,omitempty"`
	// list of devices MAC Addresses which are rebooting
	RebootInProgress []string `json:"reboot_in_progress,omitempty"`
	// list of devices MAC Addresses which have rebooted successfully
	Rebooted []string `json:"rebooted,omitempty"`
	// list of devices MAC Addresses which cloud has scheduled an upgrade for
	Scheduled []string `json:"scheduled,omitempty"`
	// list of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade
	Skipped []string `json:"skipped,omitempty"`
	// count of devices part of this upgrade
	Total *int32 `json:"total,omitempty"`
	// count of devices which have upgraded successfully
	Upgraded []string `json:"upgraded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResponseDeviceUpgradeCounts ResponseDeviceUpgradeCounts

// NewResponseDeviceUpgradeCounts instantiates a new ResponseDeviceUpgradeCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeviceUpgradeCounts() *ResponseDeviceUpgradeCounts {
	this := ResponseDeviceUpgradeCounts{}
	return &this
}

// NewResponseDeviceUpgradeCountsWithDefaults instantiates a new ResponseDeviceUpgradeCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeviceUpgradeCountsWithDefaults() *ResponseDeviceUpgradeCounts {
	this := ResponseDeviceUpgradeCounts{}
	return &this
}

// GetDownloadRequested returns the DownloadRequested field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetDownloadRequested() []string {
	if o == nil || IsNil(o.DownloadRequested) {
		var ret []string
		return ret
	}
	return o.DownloadRequested
}

// GetDownloadRequestedOk returns a tuple with the DownloadRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetDownloadRequestedOk() ([]string, bool) {
	if o == nil || IsNil(o.DownloadRequested) {
		return nil, false
	}
	return o.DownloadRequested, true
}

// HasDownloadRequested returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasDownloadRequested() bool {
	if o != nil && !IsNil(o.DownloadRequested) {
		return true
	}

	return false
}

// SetDownloadRequested gets a reference to the given []string and assigns it to the DownloadRequested field.
func (o *ResponseDeviceUpgradeCounts) SetDownloadRequested(v []string) {
	o.DownloadRequested = v
}

// GetDownloaded returns the Downloaded field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetDownloaded() []string {
	if o == nil || IsNil(o.Downloaded) {
		var ret []string
		return ret
	}
	return o.Downloaded
}

// GetDownloadedOk returns a tuple with the Downloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetDownloadedOk() ([]string, bool) {
	if o == nil || IsNil(o.Downloaded) {
		return nil, false
	}
	return o.Downloaded, true
}

// HasDownloaded returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasDownloaded() bool {
	if o != nil && !IsNil(o.Downloaded) {
		return true
	}

	return false
}

// SetDownloaded gets a reference to the given []string and assigns it to the Downloaded field.
func (o *ResponseDeviceUpgradeCounts) SetDownloaded(v []string) {
	o.Downloaded = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetFailed() []string {
	if o == nil || IsNil(o.Failed) {
		var ret []string
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetFailedOk() ([]string, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given []string and assigns it to the Failed field.
func (o *ResponseDeviceUpgradeCounts) SetFailed(v []string) {
	o.Failed = v
}

// GetRebootInProgress returns the RebootInProgress field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetRebootInProgress() []string {
	if o == nil || IsNil(o.RebootInProgress) {
		var ret []string
		return ret
	}
	return o.RebootInProgress
}

// GetRebootInProgressOk returns a tuple with the RebootInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetRebootInProgressOk() ([]string, bool) {
	if o == nil || IsNil(o.RebootInProgress) {
		return nil, false
	}
	return o.RebootInProgress, true
}

// HasRebootInProgress returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasRebootInProgress() bool {
	if o != nil && !IsNil(o.RebootInProgress) {
		return true
	}

	return false
}

// SetRebootInProgress gets a reference to the given []string and assigns it to the RebootInProgress field.
func (o *ResponseDeviceUpgradeCounts) SetRebootInProgress(v []string) {
	o.RebootInProgress = v
}

// GetRebooted returns the Rebooted field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetRebooted() []string {
	if o == nil || IsNil(o.Rebooted) {
		var ret []string
		return ret
	}
	return o.Rebooted
}

// GetRebootedOk returns a tuple with the Rebooted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetRebootedOk() ([]string, bool) {
	if o == nil || IsNil(o.Rebooted) {
		return nil, false
	}
	return o.Rebooted, true
}

// HasRebooted returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasRebooted() bool {
	if o != nil && !IsNil(o.Rebooted) {
		return true
	}

	return false
}

// SetRebooted gets a reference to the given []string and assigns it to the Rebooted field.
func (o *ResponseDeviceUpgradeCounts) SetRebooted(v []string) {
	o.Rebooted = v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetScheduled() []string {
	if o == nil || IsNil(o.Scheduled) {
		var ret []string
		return ret
	}
	return o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetScheduledOk() ([]string, bool) {
	if o == nil || IsNil(o.Scheduled) {
		return nil, false
	}
	return o.Scheduled, true
}

// HasScheduled returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasScheduled() bool {
	if o != nil && !IsNil(o.Scheduled) {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given []string and assigns it to the Scheduled field.
func (o *ResponseDeviceUpgradeCounts) SetScheduled(v []string) {
	o.Scheduled = v
}

// GetSkipped returns the Skipped field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetSkipped() []string {
	if o == nil || IsNil(o.Skipped) {
		var ret []string
		return ret
	}
	return o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetSkippedOk() ([]string, bool) {
	if o == nil || IsNil(o.Skipped) {
		return nil, false
	}
	return o.Skipped, true
}

// HasSkipped returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasSkipped() bool {
	if o != nil && !IsNil(o.Skipped) {
		return true
	}

	return false
}

// SetSkipped gets a reference to the given []string and assigns it to the Skipped field.
func (o *ResponseDeviceUpgradeCounts) SetSkipped(v []string) {
	o.Skipped = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ResponseDeviceUpgradeCounts) SetTotal(v int32) {
	o.Total = &v
}

// GetUpgraded returns the Upgraded field value if set, zero value otherwise.
func (o *ResponseDeviceUpgradeCounts) GetUpgraded() []string {
	if o == nil || IsNil(o.Upgraded) {
		var ret []string
		return ret
	}
	return o.Upgraded
}

// GetUpgradedOk returns a tuple with the Upgraded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDeviceUpgradeCounts) GetUpgradedOk() ([]string, bool) {
	if o == nil || IsNil(o.Upgraded) {
		return nil, false
	}
	return o.Upgraded, true
}

// HasUpgraded returns a boolean if a field has been set.
func (o *ResponseDeviceUpgradeCounts) HasUpgraded() bool {
	if o != nil && !IsNil(o.Upgraded) {
		return true
	}

	return false
}

// SetUpgraded gets a reference to the given []string and assigns it to the Upgraded field.
func (o *ResponseDeviceUpgradeCounts) SetUpgraded(v []string) {
	o.Upgraded = v
}

func (o ResponseDeviceUpgradeCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeviceUpgradeCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadRequested) {
		toSerialize["download_requested"] = o.DownloadRequested
	}
	if !IsNil(o.Downloaded) {
		toSerialize["downloaded"] = o.Downloaded
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.RebootInProgress) {
		toSerialize["reboot_in_progress"] = o.RebootInProgress
	}
	if !IsNil(o.Rebooted) {
		toSerialize["rebooted"] = o.Rebooted
	}
	if !IsNil(o.Scheduled) {
		toSerialize["scheduled"] = o.Scheduled
	}
	if !IsNil(o.Skipped) {
		toSerialize["skipped"] = o.Skipped
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Upgraded) {
		toSerialize["upgraded"] = o.Upgraded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResponseDeviceUpgradeCounts) UnmarshalJSON(data []byte) (err error) {
	varResponseDeviceUpgradeCounts := _ResponseDeviceUpgradeCounts{}

	err = json.Unmarshal(data, &varResponseDeviceUpgradeCounts)

	if err != nil {
		return err
	}

	*o = ResponseDeviceUpgradeCounts(varResponseDeviceUpgradeCounts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "download_requested")
		delete(additionalProperties, "downloaded")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "reboot_in_progress")
		delete(additionalProperties, "rebooted")
		delete(additionalProperties, "scheduled")
		delete(additionalProperties, "skipped")
		delete(additionalProperties, "total")
		delete(additionalProperties, "upgraded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResponseDeviceUpgradeCounts struct {
	value *ResponseDeviceUpgradeCounts
	isSet bool
}

func (v NullableResponseDeviceUpgradeCounts) Get() *ResponseDeviceUpgradeCounts {
	return v.value
}

func (v *NullableResponseDeviceUpgradeCounts) Set(val *ResponseDeviceUpgradeCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeviceUpgradeCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeviceUpgradeCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeviceUpgradeCounts(val *ResponseDeviceUpgradeCounts) *NullableResponseDeviceUpgradeCounts {
	return &NullableResponseDeviceUpgradeCounts{value: val, isSet: true}
}

func (v NullableResponseDeviceUpgradeCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeviceUpgradeCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


