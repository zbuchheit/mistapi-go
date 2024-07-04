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

// checks if the MapNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapNode{}

// MapNode Nodes on maps
type MapNode struct {
	Edges *map[string]string `json:"edges,omitempty"`
	Name string `json:"name"`
	Position *MapNodePosition `json:"position,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MapNode MapNode

// NewMapNode instantiates a new MapNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapNode(name string) *MapNode {
	this := MapNode{}
	this.Name = name
	return &this
}

// NewMapNodeWithDefaults instantiates a new MapNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapNodeWithDefaults() *MapNode {
	this := MapNode{}
	return &this
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *MapNode) GetEdges() map[string]string {
	if o == nil || IsNil(o.Edges) {
		var ret map[string]string
		return ret
	}
	return *o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapNode) GetEdgesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Edges) {
		return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *MapNode) HasEdges() bool {
	if o != nil && !IsNil(o.Edges) {
		return true
	}

	return false
}

// SetEdges gets a reference to the given map[string]string and assigns it to the Edges field.
func (o *MapNode) SetEdges(v map[string]string) {
	o.Edges = &v
}

// GetName returns the Name field value
func (o *MapNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MapNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MapNode) SetName(v string) {
	o.Name = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *MapNode) GetPosition() MapNodePosition {
	if o == nil || IsNil(o.Position) {
		var ret MapNodePosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapNode) GetPositionOk() (*MapNodePosition, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *MapNode) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given MapNodePosition and assigns it to the Position field.
func (o *MapNode) SetPosition(v MapNodePosition) {
	o.Position = &v
}

func (o MapNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Edges) {
		toSerialize["edges"] = o.Edges
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MapNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varMapNode := _MapNode{}

	err = json.Unmarshal(data, &varMapNode)

	if err != nil {
		return err
	}

	*o = MapNode(varMapNode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "edges")
		delete(additionalProperties, "name")
		delete(additionalProperties, "position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMapNode struct {
	value *MapNode
	isSet bool
}

func (v NullableMapNode) Get() *MapNode {
	return v.value
}

func (v *NullableMapNode) Set(val *MapNode) {
	v.value = val
	v.isSet = true
}

func (v NullableMapNode) IsSet() bool {
	return v.isSet
}

func (v *NullableMapNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapNode(val *MapNode) *NullableMapNode {
	return &NullableMapNode{value: val, isSet: true}
}

func (v NullableMapNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


