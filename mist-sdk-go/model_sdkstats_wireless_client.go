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

// checks if the SdkstatsWirelessClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdkstatsWirelessClient{}

// SdkstatsWirelessClient SDK Client Details statistics
type SdkstatsWirelessClient struct {
	Id string `json:"id"`
	// last seen timestamp
	LastSeen float32 `json:"last_seen"`
	// map_id of the sdk client (if known), or null
	MapId NullableString `json:"map_id,omitempty"`
	// name of the sdk client (if provided)
	Name *string `json:"name,omitempty"`
	NetworkConnection *SdkclientStatNetworkConnection `json:"network_connection,omitempty"`
	// uuid of the sdk client
	Uuid string `json:"uuid"`
	// list of beacon_id’s of the sdk client is in and since when (if known)
	Vbeacons []SdkclientWirelessStatsVbeacon `json:"vbeacons,omitempty"`
	// x (in pixels) of user location on the map (if known)
	X *float32 `json:"x,omitempty"`
	// y (in pixels) of user location on the map (if known)
	Y *float32 `json:"y,omitempty"`
	// list of zone_id’s of the sdk client is in and since when (if known)
	Zones []SdkclientWirelessStatsZone `json:"zones,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdkstatsWirelessClient SdkstatsWirelessClient

// NewSdkstatsWirelessClient instantiates a new SdkstatsWirelessClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdkstatsWirelessClient(id string, lastSeen float32, uuid string) *SdkstatsWirelessClient {
	this := SdkstatsWirelessClient{}
	this.Id = id
	this.LastSeen = lastSeen
	this.Uuid = uuid
	return &this
}

// NewSdkstatsWirelessClientWithDefaults instantiates a new SdkstatsWirelessClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdkstatsWirelessClientWithDefaults() *SdkstatsWirelessClient {
	this := SdkstatsWirelessClient{}
	return &this
}

// GetId returns the Id field value
func (o *SdkstatsWirelessClient) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SdkstatsWirelessClient) SetId(v string) {
	o.Id = v
}

// GetLastSeen returns the LastSeen field value
func (o *SdkstatsWirelessClient) GetLastSeen() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetLastSeenOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value
func (o *SdkstatsWirelessClient) SetLastSeen(v float32) {
	o.LastSeen = v
}

// GetMapId returns the MapId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdkstatsWirelessClient) GetMapId() string {
	if o == nil || IsNil(o.MapId.Get()) {
		var ret string
		return ret
	}
	return *o.MapId.Get()
}

// GetMapIdOk returns a tuple with the MapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdkstatsWirelessClient) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MapId.Get(), o.MapId.IsSet()
}

// HasMapId returns a boolean if a field has been set.
func (o *SdkstatsWirelessClient) HasMapId() bool {
	if o != nil && o.MapId.IsSet() {
		return true
	}

	return false
}

// SetMapId gets a reference to the given NullableString and assigns it to the MapId field.
func (o *SdkstatsWirelessClient) SetMapId(v string) {
	o.MapId.Set(&v)
}
// SetMapIdNil sets the value for MapId to be an explicit nil
func (o *SdkstatsWirelessClient) SetMapIdNil() {
	o.MapId.Set(nil)
}

// UnsetMapId ensures that no value is present for MapId, not even an explicit nil
func (o *SdkstatsWirelessClient) UnsetMapId() {
	o.MapId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SdkstatsWirelessClient) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SdkstatsWirelessClient) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SdkstatsWirelessClient) SetName(v string) {
	o.Name = &v
}

// GetNetworkConnection returns the NetworkConnection field value if set, zero value otherwise.
func (o *SdkstatsWirelessClient) GetNetworkConnection() SdkclientStatNetworkConnection {
	if o == nil || IsNil(o.NetworkConnection) {
		var ret SdkclientStatNetworkConnection
		return ret
	}
	return *o.NetworkConnection
}

// GetNetworkConnectionOk returns a tuple with the NetworkConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetNetworkConnectionOk() (*SdkclientStatNetworkConnection, bool) {
	if o == nil || IsNil(o.NetworkConnection) {
		return nil, false
	}
	return o.NetworkConnection, true
}

// HasNetworkConnection returns a boolean if a field has been set.
func (o *SdkstatsWirelessClient) HasNetworkConnection() bool {
	if o != nil && !IsNil(o.NetworkConnection) {
		return true
	}

	return false
}

// SetNetworkConnection gets a reference to the given SdkclientStatNetworkConnection and assigns it to the NetworkConnection field.
func (o *SdkstatsWirelessClient) SetNetworkConnection(v SdkclientStatNetworkConnection) {
	o.NetworkConnection = &v
}

// GetUuid returns the Uuid field value
func (o *SdkstatsWirelessClient) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SdkstatsWirelessClient) SetUuid(v string) {
	o.Uuid = v
}

// GetVbeacons returns the Vbeacons field value if set, zero value otherwise.
func (o *SdkstatsWirelessClient) GetVbeacons() []SdkclientWirelessStatsVbeacon {
	if o == nil || IsNil(o.Vbeacons) {
		var ret []SdkclientWirelessStatsVbeacon
		return ret
	}
	return o.Vbeacons
}

// GetVbeaconsOk returns a tuple with the Vbeacons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetVbeaconsOk() ([]SdkclientWirelessStatsVbeacon, bool) {
	if o == nil || IsNil(o.Vbeacons) {
		return nil, false
	}
	return o.Vbeacons, true
}

// HasVbeacons returns a boolean if a field has been set.
func (o *SdkstatsWirelessClient) HasVbeacons() bool {
	if o != nil && !IsNil(o.Vbeacons) {
		return true
	}

	return false
}

// SetVbeacons gets a reference to the given []SdkclientWirelessStatsVbeacon and assigns it to the Vbeacons field.
func (o *SdkstatsWirelessClient) SetVbeacons(v []SdkclientWirelessStatsVbeacon) {
	o.Vbeacons = v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SdkstatsWirelessClient) GetX() float32 {
	if o == nil || IsNil(o.X) {
		var ret float32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetXOk() (*float32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SdkstatsWirelessClient) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float32 and assigns it to the X field.
func (o *SdkstatsWirelessClient) SetX(v float32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SdkstatsWirelessClient) GetY() float32 {
	if o == nil || IsNil(o.Y) {
		var ret float32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetYOk() (*float32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SdkstatsWirelessClient) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float32 and assigns it to the Y field.
func (o *SdkstatsWirelessClient) SetY(v float32) {
	o.Y = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *SdkstatsWirelessClient) GetZones() []SdkclientWirelessStatsZone {
	if o == nil || IsNil(o.Zones) {
		var ret []SdkclientWirelessStatsZone
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdkstatsWirelessClient) GetZonesOk() ([]SdkclientWirelessStatsZone, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *SdkstatsWirelessClient) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []SdkclientWirelessStatsZone and assigns it to the Zones field.
func (o *SdkstatsWirelessClient) SetZones(v []SdkclientWirelessStatsZone) {
	o.Zones = v
}

func (o SdkstatsWirelessClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdkstatsWirelessClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["last_seen"] = o.LastSeen
	if o.MapId.IsSet() {
		toSerialize["map_id"] = o.MapId.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetworkConnection) {
		toSerialize["network_connection"] = o.NetworkConnection
	}
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Vbeacons) {
		toSerialize["vbeacons"] = o.Vbeacons
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SdkstatsWirelessClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"last_seen",
		"uuid",
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

	varSdkstatsWirelessClient := _SdkstatsWirelessClient{}

	err = json.Unmarshal(data, &varSdkstatsWirelessClient)

	if err != nil {
		return err
	}

	*o = SdkstatsWirelessClient(varSdkstatsWirelessClient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "last_seen")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "network_connection")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "vbeacons")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "zones")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdkstatsWirelessClient struct {
	value *SdkstatsWirelessClient
	isSet bool
}

func (v NullableSdkstatsWirelessClient) Get() *SdkstatsWirelessClient {
	return v.value
}

func (v *NullableSdkstatsWirelessClient) Set(val *SdkstatsWirelessClient) {
	v.value = val
	v.isSet = true
}

func (v NullableSdkstatsWirelessClient) IsSet() bool {
	return v.isSet
}

func (v *NullableSdkstatsWirelessClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdkstatsWirelessClient(val *SdkstatsWirelessClient) *NullableSdkstatsWirelessClient {
	return &NullableSdkstatsWirelessClient{value: val, isSet: true}
}

func (v NullableSdkstatsWirelessClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdkstatsWirelessClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


