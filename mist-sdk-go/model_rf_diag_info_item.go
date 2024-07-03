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

// checks if the RfDiagInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RfDiagInfoItem{}

// RfDiagInfoItem struct for RfDiagInfoItem
type RfDiagInfoItem struct {
	// if `type`==`asset`, id of the asset
	AssetId *string `json:"asset_id,omitempty"`
	// if `type`==`asset`, name of the asset
	AssetName *string `json:"asset_name,omitempty"`
	// if `type`==`client`, hostname of the client
	ClientName *string `json:"client_name,omitempty"`
	// recording length in seconds, max is 120
	Duration int32 `json:"duration"`
	// timestamp of end of recording
	EndTime int32 `json:"end_time"`
	// Number of frames in the output
	FrameCount int32 `json:"frame_count"`
	Id *string `json:"id,omitempty"`
	// if `type`==`client` or `asset`, mac of the device
	Mac *string `json:"mac,omitempty"`
	MapId string `json:"map_id"`
	Name string `json:"name"`
	// Optional. id of the next recoding if present. Only valid for site survey.
	Next *string `json:"next,omitempty"`
	// URL to a JSON file that contains array of raw location diag events
	RawEvents string `json:"raw_events"`
	// whether it’s ready for playback
	Ready bool `json:"ready"`
	// if `type`==`sdkclient`, sdkclient_id of this recording
	SdkclientId *string `json:"sdkclient_id,omitempty"`
	// if `type`==`sdkclient`, name of the sdkclient
	SdkclientName *string `json:"sdkclient_name,omitempty"`
	// if `type`==`sdkclient`, device_id of sdkclient
	SdkclientUuid *string `json:"sdkclient_uuid,omitempty"`
	// timestamp of the recording (the start)
	StartTime int32 `json:"start_time"`
	Type RfClientType `json:"type"`
	// URL to a JSON file that contains an array of frames, each frame is the same format
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _RfDiagInfoItem RfDiagInfoItem

// NewRfDiagInfoItem instantiates a new RfDiagInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRfDiagInfoItem(duration int32, endTime int32, frameCount int32, mapId string, name string, rawEvents string, ready bool, startTime int32, type_ RfClientType, url string) *RfDiagInfoItem {
	this := RfDiagInfoItem{}
	this.Duration = duration
	this.EndTime = endTime
	this.FrameCount = frameCount
	this.MapId = mapId
	this.Name = name
	this.RawEvents = rawEvents
	this.Ready = ready
	this.StartTime = startTime
	this.Type = type_
	this.Url = url
	return &this
}

// NewRfDiagInfoItemWithDefaults instantiates a new RfDiagInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRfDiagInfoItemWithDefaults() *RfDiagInfoItem {
	this := RfDiagInfoItem{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetAssetId() string {
	if o == nil || IsNil(o.AssetId) {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *RfDiagInfoItem) SetAssetId(v string) {
	o.AssetId = &v
}

// GetAssetName returns the AssetName field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetAssetName() string {
	if o == nil || IsNil(o.AssetName) {
		var ret string
		return ret
	}
	return *o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetAssetNameOk() (*string, bool) {
	if o == nil || IsNil(o.AssetName) {
		return nil, false
	}
	return o.AssetName, true
}

// HasAssetName returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasAssetName() bool {
	if o != nil && !IsNil(o.AssetName) {
		return true
	}

	return false
}

// SetAssetName gets a reference to the given string and assigns it to the AssetName field.
func (o *RfDiagInfoItem) SetAssetName(v string) {
	o.AssetName = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *RfDiagInfoItem) SetClientName(v string) {
	o.ClientName = &v
}

// GetDuration returns the Duration field value
func (o *RfDiagInfoItem) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *RfDiagInfoItem) SetDuration(v int32) {
	o.Duration = v
}

// GetEndTime returns the EndTime field value
func (o *RfDiagInfoItem) GetEndTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetEndTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *RfDiagInfoItem) SetEndTime(v int32) {
	o.EndTime = v
}

// GetFrameCount returns the FrameCount field value
func (o *RfDiagInfoItem) GetFrameCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FrameCount
}

// GetFrameCountOk returns a tuple with the FrameCount field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetFrameCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrameCount, true
}

// SetFrameCount sets field value
func (o *RfDiagInfoItem) SetFrameCount(v int32) {
	o.FrameCount = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RfDiagInfoItem) SetId(v string) {
	o.Id = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *RfDiagInfoItem) SetMac(v string) {
	o.Mac = &v
}

// GetMapId returns the MapId field value
func (o *RfDiagInfoItem) GetMapId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapId
}

// GetMapIdOk returns a tuple with the MapId field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetMapIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapId, true
}

// SetMapId sets field value
func (o *RfDiagInfoItem) SetMapId(v string) {
	o.MapId = v
}

// GetName returns the Name field value
func (o *RfDiagInfoItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RfDiagInfoItem) SetName(v string) {
	o.Name = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *RfDiagInfoItem) SetNext(v string) {
	o.Next = &v
}

// GetRawEvents returns the RawEvents field value
func (o *RfDiagInfoItem) GetRawEvents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawEvents
}

// GetRawEventsOk returns a tuple with the RawEvents field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetRawEventsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawEvents, true
}

// SetRawEvents sets field value
func (o *RfDiagInfoItem) SetRawEvents(v string) {
	o.RawEvents = v
}

// GetReady returns the Ready field value
func (o *RfDiagInfoItem) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *RfDiagInfoItem) SetReady(v bool) {
	o.Ready = v
}

// GetSdkclientId returns the SdkclientId field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetSdkclientId() string {
	if o == nil || IsNil(o.SdkclientId) {
		var ret string
		return ret
	}
	return *o.SdkclientId
}

// GetSdkclientIdOk returns a tuple with the SdkclientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetSdkclientIdOk() (*string, bool) {
	if o == nil || IsNil(o.SdkclientId) {
		return nil, false
	}
	return o.SdkclientId, true
}

// HasSdkclientId returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasSdkclientId() bool {
	if o != nil && !IsNil(o.SdkclientId) {
		return true
	}

	return false
}

// SetSdkclientId gets a reference to the given string and assigns it to the SdkclientId field.
func (o *RfDiagInfoItem) SetSdkclientId(v string) {
	o.SdkclientId = &v
}

// GetSdkclientName returns the SdkclientName field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetSdkclientName() string {
	if o == nil || IsNil(o.SdkclientName) {
		var ret string
		return ret
	}
	return *o.SdkclientName
}

// GetSdkclientNameOk returns a tuple with the SdkclientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetSdkclientNameOk() (*string, bool) {
	if o == nil || IsNil(o.SdkclientName) {
		return nil, false
	}
	return o.SdkclientName, true
}

// HasSdkclientName returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasSdkclientName() bool {
	if o != nil && !IsNil(o.SdkclientName) {
		return true
	}

	return false
}

// SetSdkclientName gets a reference to the given string and assigns it to the SdkclientName field.
func (o *RfDiagInfoItem) SetSdkclientName(v string) {
	o.SdkclientName = &v
}

// GetSdkclientUuid returns the SdkclientUuid field value if set, zero value otherwise.
func (o *RfDiagInfoItem) GetSdkclientUuid() string {
	if o == nil || IsNil(o.SdkclientUuid) {
		var ret string
		return ret
	}
	return *o.SdkclientUuid
}

// GetSdkclientUuidOk returns a tuple with the SdkclientUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetSdkclientUuidOk() (*string, bool) {
	if o == nil || IsNil(o.SdkclientUuid) {
		return nil, false
	}
	return o.SdkclientUuid, true
}

// HasSdkclientUuid returns a boolean if a field has been set.
func (o *RfDiagInfoItem) HasSdkclientUuid() bool {
	if o != nil && !IsNil(o.SdkclientUuid) {
		return true
	}

	return false
}

// SetSdkclientUuid gets a reference to the given string and assigns it to the SdkclientUuid field.
func (o *RfDiagInfoItem) SetSdkclientUuid(v string) {
	o.SdkclientUuid = &v
}

// GetStartTime returns the StartTime field value
func (o *RfDiagInfoItem) GetStartTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetStartTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *RfDiagInfoItem) SetStartTime(v int32) {
	o.StartTime = v
}

// GetType returns the Type field value
func (o *RfDiagInfoItem) GetType() RfClientType {
	if o == nil {
		var ret RfClientType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetTypeOk() (*RfClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RfDiagInfoItem) SetType(v RfClientType) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *RfDiagInfoItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RfDiagInfoItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RfDiagInfoItem) SetUrl(v string) {
	o.Url = v
}

func (o RfDiagInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RfDiagInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetId) {
		toSerialize["asset_id"] = o.AssetId
	}
	if !IsNil(o.AssetName) {
		toSerialize["asset_name"] = o.AssetName
	}
	if !IsNil(o.ClientName) {
		toSerialize["client_name"] = o.ClientName
	}
	toSerialize["duration"] = o.Duration
	toSerialize["end_time"] = o.EndTime
	toSerialize["frame_count"] = o.FrameCount
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	toSerialize["map_id"] = o.MapId
	toSerialize["name"] = o.Name
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	toSerialize["raw_events"] = o.RawEvents
	toSerialize["ready"] = o.Ready
	if !IsNil(o.SdkclientId) {
		toSerialize["sdkclient_id"] = o.SdkclientId
	}
	if !IsNil(o.SdkclientName) {
		toSerialize["sdkclient_name"] = o.SdkclientName
	}
	if !IsNil(o.SdkclientUuid) {
		toSerialize["sdkclient_uuid"] = o.SdkclientUuid
	}
	toSerialize["start_time"] = o.StartTime
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RfDiagInfoItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration",
		"end_time",
		"frame_count",
		"map_id",
		"name",
		"raw_events",
		"ready",
		"start_time",
		"type",
		"url",
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

	varRfDiagInfoItem := _RfDiagInfoItem{}

	err = json.Unmarshal(data, &varRfDiagInfoItem)

	if err != nil {
		return err
	}

	*o = RfDiagInfoItem(varRfDiagInfoItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_id")
		delete(additionalProperties, "asset_name")
		delete(additionalProperties, "client_name")
		delete(additionalProperties, "duration")
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "frame_count")
		delete(additionalProperties, "id")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "map_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "next")
		delete(additionalProperties, "raw_events")
		delete(additionalProperties, "ready")
		delete(additionalProperties, "sdkclient_id")
		delete(additionalProperties, "sdkclient_name")
		delete(additionalProperties, "sdkclient_uuid")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRfDiagInfoItem struct {
	value *RfDiagInfoItem
	isSet bool
}

func (v NullableRfDiagInfoItem) Get() *RfDiagInfoItem {
	return v.value
}

func (v *NullableRfDiagInfoItem) Set(val *RfDiagInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRfDiagInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRfDiagInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRfDiagInfoItem(val *RfDiagInfoItem) *NullableRfDiagInfoItem {
	return &NullableRfDiagInfoItem{value: val, isSet: true}
}

func (v NullableRfDiagInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRfDiagInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


