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

// checks if the SiteSettingAutoPlacement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiteSettingAutoPlacement{}

// SiteSettingAutoPlacement if we're able to determine its x/y/orientation, this will be populated
type SiteSettingAutoPlacement struct {
	Orientation *float32 `json:"orientation,omitempty"`
	X *float32 `json:"x,omitempty"`
	Y *float32 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SiteSettingAutoPlacement SiteSettingAutoPlacement

// NewSiteSettingAutoPlacement instantiates a new SiteSettingAutoPlacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteSettingAutoPlacement() *SiteSettingAutoPlacement {
	this := SiteSettingAutoPlacement{}
	return &this
}

// NewSiteSettingAutoPlacementWithDefaults instantiates a new SiteSettingAutoPlacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteSettingAutoPlacementWithDefaults() *SiteSettingAutoPlacement {
	this := SiteSettingAutoPlacement{}
	return &this
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *SiteSettingAutoPlacement) GetOrientation() float32 {
	if o == nil || IsNil(o.Orientation) {
		var ret float32
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingAutoPlacement) GetOrientationOk() (*float32, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *SiteSettingAutoPlacement) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given float32 and assigns it to the Orientation field.
func (o *SiteSettingAutoPlacement) SetOrientation(v float32) {
	o.Orientation = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SiteSettingAutoPlacement) GetX() float32 {
	if o == nil || IsNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingAutoPlacement) GetXOk() (*float32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SiteSettingAutoPlacement) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *SiteSettingAutoPlacement) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SiteSettingAutoPlacement) GetY() float32 {
	if o == nil || IsNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteSettingAutoPlacement) GetYOk() (*float32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SiteSettingAutoPlacement) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *SiteSettingAutoPlacement) SetY(v float32) {
	o.Y = &v
}

func (o SiteSettingAutoPlacement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiteSettingAutoPlacement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SiteSettingAutoPlacement) UnmarshalJSON(data []byte) (err error) {
	varSiteSettingAutoPlacement := _SiteSettingAutoPlacement{}

	err = json.Unmarshal(data, &varSiteSettingAutoPlacement)

	if err != nil {
		return err
	}

	*o = SiteSettingAutoPlacement(varSiteSettingAutoPlacement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orientation")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteSettingAutoPlacement struct {
	value *SiteSettingAutoPlacement
	isSet bool
}

func (v NullableSiteSettingAutoPlacement) Get() *SiteSettingAutoPlacement {
	return v.value
}

func (v *NullableSiteSettingAutoPlacement) Set(val *SiteSettingAutoPlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSettingAutoPlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSettingAutoPlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSettingAutoPlacement(val *SiteSettingAutoPlacement) *NullableSiteSettingAutoPlacement {
	return &NullableSiteSettingAutoPlacement{value: val, isSet: true}
}

func (v NullableSiteSettingAutoPlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSettingAutoPlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


