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

// checks if the MapWallPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapWallPath{}

// MapWallPath a JSON blob for wall definition (same format as wayfinding_path)
type MapWallPath struct {
	Coordinate *string `json:"coordinate,omitempty"`
	Nodes []MapNode `json:"nodes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MapWallPath MapWallPath

// NewMapWallPath instantiates a new MapWallPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapWallPath() *MapWallPath {
	this := MapWallPath{}
	return &this
}

// NewMapWallPathWithDefaults instantiates a new MapWallPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapWallPathWithDefaults() *MapWallPath {
	this := MapWallPath{}
	return &this
}

// GetCoordinate returns the Coordinate field value if set, zero value otherwise.
func (o *MapWallPath) GetCoordinate() string {
	if o == nil || IsNil(o.Coordinate) {
		var ret string
		return ret
	}
	return *o.Coordinate
}

// GetCoordinateOk returns a tuple with the Coordinate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapWallPath) GetCoordinateOk() (*string, bool) {
	if o == nil || IsNil(o.Coordinate) {
		return nil, false
	}
	return o.Coordinate, true
}

// HasCoordinate returns a boolean if a field has been set.
func (o *MapWallPath) HasCoordinate() bool {
	if o != nil && !IsNil(o.Coordinate) {
		return true
	}

	return false
}

// SetCoordinate gets a reference to the given string and assigns it to the Coordinate field.
func (o *MapWallPath) SetCoordinate(v string) {
	o.Coordinate = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *MapWallPath) GetNodes() []MapNode {
	if o == nil || IsNil(o.Nodes) {
		var ret []MapNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapWallPath) GetNodesOk() ([]MapNode, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *MapWallPath) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []MapNode and assigns it to the Nodes field.
func (o *MapWallPath) SetNodes(v []MapNode) {
	o.Nodes = v
}

func (o MapWallPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapWallPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coordinate) {
		toSerialize["coordinate"] = o.Coordinate
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MapWallPath) UnmarshalJSON(data []byte) (err error) {
	varMapWallPath := _MapWallPath{}

	err = json.Unmarshal(data, &varMapWallPath)

	if err != nil {
		return err
	}

	*o = MapWallPath(varMapWallPath)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "coordinate")
		delete(additionalProperties, "nodes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMapWallPath struct {
	value *MapWallPath
	isSet bool
}

func (v NullableMapWallPath) Get() *MapWallPath {
	return v.value
}

func (v *NullableMapWallPath) Set(val *MapWallPath) {
	v.value = val
	v.isSet = true
}

func (v NullableMapWallPath) IsSet() bool {
	return v.isSet
}

func (v *NullableMapWallPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapWallPath(val *MapWallPath) *NullableMapWallPath {
	return &NullableMapWallPath{value: val, isSet: true}
}

func (v NullableMapWallPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapWallPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


