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
)

// checks if the ApStatsAutoPlacementInfoProbabilitySurface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApStatsAutoPlacementInfoProbabilitySurface{}

// ApStatsAutoPlacementInfoProbabilitySurface Coordinates representing a circle where the AP is most likely exists in the event of an inaccurate placement result
type ApStatsAutoPlacementInfoProbabilitySurface struct {
	// The radius representing placement uncertainty, measured in pixels
	Radius *float32 `json:"radius,omitempty"`
	// The radius representing placement uncertainty, measured in meters
	RadiusM *float32 `json:"radius_m,omitempty"`
	// X-coordinate of the potential placement’s center, measured in pixels
	X *float32 `json:"x,omitempty"`
	// Y-coordinate of the potential placement’s center, measured in pixels
	Y *float32 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApStatsAutoPlacementInfoProbabilitySurface ApStatsAutoPlacementInfoProbabilitySurface

// NewApStatsAutoPlacementInfoProbabilitySurface instantiates a new ApStatsAutoPlacementInfoProbabilitySurface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApStatsAutoPlacementInfoProbabilitySurface() *ApStatsAutoPlacementInfoProbabilitySurface {
	this := ApStatsAutoPlacementInfoProbabilitySurface{}
	return &this
}

// NewApStatsAutoPlacementInfoProbabilitySurfaceWithDefaults instantiates a new ApStatsAutoPlacementInfoProbabilitySurface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApStatsAutoPlacementInfoProbabilitySurfaceWithDefaults() *ApStatsAutoPlacementInfoProbabilitySurface {
	this := ApStatsAutoPlacementInfoProbabilitySurface{}
	return &this
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetRadius() float32 {
	if o == nil || IsNil(o.Radius) {
		var ret float32
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetRadiusOk() (*float32, bool) {
	if o == nil || IsNil(o.Radius) {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) HasRadius() bool {
	if o != nil && !IsNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given float32 and assigns it to the Radius field.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) SetRadius(v float32) {
	o.Radius = &v
}

// GetRadiusM returns the RadiusM field value if set, zero value otherwise.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetRadiusM() float32 {
	if o == nil || IsNil(o.RadiusM) {
		var ret float32
		return ret
	}
	return *o.RadiusM
}

// GetRadiusMOk returns a tuple with the RadiusM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetRadiusMOk() (*float32, bool) {
	if o == nil || IsNil(o.RadiusM) {
		return nil, false
	}
	return o.RadiusM, true
}

// HasRadiusM returns a boolean if a field has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) HasRadiusM() bool {
	if o != nil && !IsNil(o.RadiusM) {
		return true
	}

	return false
}

// SetRadiusM gets a reference to the given float32 and assigns it to the RadiusM field.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) SetRadiusM(v float32) {
	o.RadiusM = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetX() float32 {
	if o == nil || IsNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetXOk() (*float32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetY() float32 {
	if o == nil || IsNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) GetYOk() (*float32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *ApStatsAutoPlacementInfoProbabilitySurface) SetY(v float32) {
	o.Y = &v
}

func (o ApStatsAutoPlacementInfoProbabilitySurface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApStatsAutoPlacementInfoProbabilitySurface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !IsNil(o.RadiusM) {
		toSerialize["radius_m"] = o.RadiusM
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

func (o *ApStatsAutoPlacementInfoProbabilitySurface) UnmarshalJSON(data []byte) (err error) {
	varApStatsAutoPlacementInfoProbabilitySurface := _ApStatsAutoPlacementInfoProbabilitySurface{}

	err = json.Unmarshal(data, &varApStatsAutoPlacementInfoProbabilitySurface)

	if err != nil {
		return err
	}

	*o = ApStatsAutoPlacementInfoProbabilitySurface(varApStatsAutoPlacementInfoProbabilitySurface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "radius")
		delete(additionalProperties, "radius_m")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApStatsAutoPlacementInfoProbabilitySurface struct {
	value *ApStatsAutoPlacementInfoProbabilitySurface
	isSet bool
}

func (v NullableApStatsAutoPlacementInfoProbabilitySurface) Get() *ApStatsAutoPlacementInfoProbabilitySurface {
	return v.value
}

func (v *NullableApStatsAutoPlacementInfoProbabilitySurface) Set(val *ApStatsAutoPlacementInfoProbabilitySurface) {
	v.value = val
	v.isSet = true
}

func (v NullableApStatsAutoPlacementInfoProbabilitySurface) IsSet() bool {
	return v.isSet
}

func (v *NullableApStatsAutoPlacementInfoProbabilitySurface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApStatsAutoPlacementInfoProbabilitySurface(val *ApStatsAutoPlacementInfoProbabilitySurface) *NullableApStatsAutoPlacementInfoProbabilitySurface {
	return &NullableApStatsAutoPlacementInfoProbabilitySurface{value: val, isSet: true}
}

func (v NullableApStatsAutoPlacementInfoProbabilitySurface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApStatsAutoPlacementInfoProbabilitySurface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


