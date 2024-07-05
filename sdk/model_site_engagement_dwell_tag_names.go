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

// checks if the SiteEngagementDwellTagNames type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteEngagementDwellTagNames{}

// SiteEngagementDwellTagNames struct for SiteEngagementDwellTagNames
type SiteEngagementDwellTagNames struct {
	Bounce *string `json:"bounce,omitempty"`
	Engaged *string `json:"engaged,omitempty"`
	Passerby *string `json:"passerby,omitempty"`
	Stationed *string `json:"stationed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteEngagementDwellTagNames SiteEngagementDwellTagNames

// NewSiteEngagementDwellTagNames instantiates a new SiteEngagementDwellTagNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteEngagementDwellTagNames() *SiteEngagementDwellTagNames {
	this := SiteEngagementDwellTagNames{}
	return &this
}

// NewSiteEngagementDwellTagNamesWithDefaults instantiates a new SiteEngagementDwellTagNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteEngagementDwellTagNamesWithDefaults() *SiteEngagementDwellTagNames {
	this := SiteEngagementDwellTagNames{}
	return &this
}

// GetBounce returns the Bounce field value if set, zero value otherwise.
func (o *SiteEngagementDwellTagNames) GetBounce() string {
	if o == nil || IsNil(o.Bounce) {
		var ret string
		return ret
	}
	return *o.Bounce
}

// GetBounceOk returns a tuple with the Bounce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetBounceOk() (*string, bool) {
	if o == nil || IsNil(o.Bounce) {
		return nil, false
	}
	return o.Bounce, true
}

// HasBounce returns a boolean if a field has been set.
func (o *SiteEngagementDwellTagNames) HasBounce() bool {
	if o != nil && !IsNil(o.Bounce) {
		return true
	}

	return false
}

// SetBounce gets a reference to the given string and assigns it to the Bounce field.
func (o *SiteEngagementDwellTagNames) SetBounce(v string) {
	o.Bounce = &v
}

// GetEngaged returns the Engaged field value if set, zero value otherwise.
func (o *SiteEngagementDwellTagNames) GetEngaged() string {
	if o == nil || IsNil(o.Engaged) {
		var ret string
		return ret
	}
	return *o.Engaged
}

// GetEngagedOk returns a tuple with the Engaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetEngagedOk() (*string, bool) {
	if o == nil || IsNil(o.Engaged) {
		return nil, false
	}
	return o.Engaged, true
}

// HasEngaged returns a boolean if a field has been set.
func (o *SiteEngagementDwellTagNames) HasEngaged() bool {
	if o != nil && !IsNil(o.Engaged) {
		return true
	}

	return false
}

// SetEngaged gets a reference to the given string and assigns it to the Engaged field.
func (o *SiteEngagementDwellTagNames) SetEngaged(v string) {
	o.Engaged = &v
}

// GetPasserby returns the Passerby field value if set, zero value otherwise.
func (o *SiteEngagementDwellTagNames) GetPasserby() string {
	if o == nil || IsNil(o.Passerby) {
		var ret string
		return ret
	}
	return *o.Passerby
}

// GetPasserbyOk returns a tuple with the Passerby field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetPasserbyOk() (*string, bool) {
	if o == nil || IsNil(o.Passerby) {
		return nil, false
	}
	return o.Passerby, true
}

// HasPasserby returns a boolean if a field has been set.
func (o *SiteEngagementDwellTagNames) HasPasserby() bool {
	if o != nil && !IsNil(o.Passerby) {
		return true
	}

	return false
}

// SetPasserby gets a reference to the given string and assigns it to the Passerby field.
func (o *SiteEngagementDwellTagNames) SetPasserby(v string) {
	o.Passerby = &v
}

// GetStationed returns the Stationed field value if set, zero value otherwise.
func (o *SiteEngagementDwellTagNames) GetStationed() string {
	if o == nil || IsNil(o.Stationed) {
		var ret string
		return ret
	}
	return *o.Stationed
}

// GetStationedOk returns a tuple with the Stationed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteEngagementDwellTagNames) GetStationedOk() (*string, bool) {
	if o == nil || IsNil(o.Stationed) {
		return nil, false
	}
	return o.Stationed, true
}

// HasStationed returns a boolean if a field has been set.
func (o *SiteEngagementDwellTagNames) HasStationed() bool {
	if o != nil && !IsNil(o.Stationed) {
		return true
	}

	return false
}

// SetStationed gets a reference to the given string and assigns it to the Stationed field.
func (o *SiteEngagementDwellTagNames) SetStationed(v string) {
	o.Stationed = &v
}

func (o SiteEngagementDwellTagNames) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteEngagementDwellTagNames) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bounce) {
		toSerialize["bounce"] = o.Bounce
	}
	if !IsNil(o.Engaged) {
		toSerialize["engaged"] = o.Engaged
	}
	if !IsNil(o.Passerby) {
		toSerialize["passerby"] = o.Passerby
	}
	if !IsNil(o.Stationed) {
		toSerialize["stationed"] = o.Stationed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteEngagementDwellTagNames) UnmarshalJSON(data []byte) (err error) {
	varSiteEngagementDwellTagNames := _SiteEngagementDwellTagNames{}

	err = json.Unmarshal(data, &varSiteEngagementDwellTagNames)

	if err != nil {
		return err
	}

	*o = SiteEngagementDwellTagNames(varSiteEngagementDwellTagNames)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bounce")
		delete(additionalProperties, "engaged")
		delete(additionalProperties, "passerby")
		delete(additionalProperties, "stationed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteEngagementDwellTagNames struct {
	value *SiteEngagementDwellTagNames
	isSet bool
}

func (v NullableSiteEngagementDwellTagNames) Get() *SiteEngagementDwellTagNames {
	return v.value
}

func (v *NullableSiteEngagementDwellTagNames) Set(val *SiteEngagementDwellTagNames) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteEngagementDwellTagNames) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteEngagementDwellTagNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteEngagementDwellTagNames(val *SiteEngagementDwellTagNames) *NullableSiteEngagementDwellTagNames {
	return &NullableSiteEngagementDwellTagNames{value: val, isSet: true}
}

func (v NullableSiteEngagementDwellTagNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteEngagementDwellTagNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


