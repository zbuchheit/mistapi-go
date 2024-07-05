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

// checks if the TicketCommentsAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TicketCommentsAttachment{}

// TicketCommentsAttachment struct for TicketCommentsAttachment
type TicketCommentsAttachment struct {
	ContentType *string `json:"content_type,omitempty"`
	CreatedAt *int32 `json:"created_at,omitempty"`
	FileName *string `json:"file_name,omitempty"`
	Id *string `json:"id,omitempty"`
	SizeInBytes *int32 `json:"size_in_bytes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TicketCommentsAttachment TicketCommentsAttachment

// NewTicketCommentsAttachment instantiates a new TicketCommentsAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketCommentsAttachment() *TicketCommentsAttachment {
	this := TicketCommentsAttachment{}
	return &this
}

// NewTicketCommentsAttachmentWithDefaults instantiates a new TicketCommentsAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketCommentsAttachmentWithDefaults() *TicketCommentsAttachment {
	this := TicketCommentsAttachment{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *TicketCommentsAttachment) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCommentsAttachment) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *TicketCommentsAttachment) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *TicketCommentsAttachment) SetContentType(v string) {
	o.ContentType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TicketCommentsAttachment) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCommentsAttachment) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TicketCommentsAttachment) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *TicketCommentsAttachment) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *TicketCommentsAttachment) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCommentsAttachment) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *TicketCommentsAttachment) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *TicketCommentsAttachment) SetFileName(v string) {
	o.FileName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TicketCommentsAttachment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCommentsAttachment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TicketCommentsAttachment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TicketCommentsAttachment) SetId(v string) {
	o.Id = &v
}

// GetSizeInBytes returns the SizeInBytes field value if set, zero value otherwise.
func (o *TicketCommentsAttachment) GetSizeInBytes() int32 {
	if o == nil || IsNil(o.SizeInBytes) {
		var ret int32
		return ret
	}
	return *o.SizeInBytes
}

// GetSizeInBytesOk returns a tuple with the SizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketCommentsAttachment) GetSizeInBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.SizeInBytes) {
		return nil, false
	}
	return o.SizeInBytes, true
}

// HasSizeInBytes returns a boolean if a field has been set.
func (o *TicketCommentsAttachment) HasSizeInBytes() bool {
	if o != nil && !IsNil(o.SizeInBytes) {
		return true
	}

	return false
}

// SetSizeInBytes gets a reference to the given int32 and assigns it to the SizeInBytes field.
func (o *TicketCommentsAttachment) SetSizeInBytes(v int32) {
	o.SizeInBytes = &v
}

func (o TicketCommentsAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TicketCommentsAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SizeInBytes) {
		toSerialize["size_in_bytes"] = o.SizeInBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TicketCommentsAttachment) UnmarshalJSON(data []byte) (err error) {
	varTicketCommentsAttachment := _TicketCommentsAttachment{}

	err = json.Unmarshal(data, &varTicketCommentsAttachment)

	if err != nil {
		return err
	}

	*o = TicketCommentsAttachment(varTicketCommentsAttachment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "file_name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "size_in_bytes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTicketCommentsAttachment struct {
	value *TicketCommentsAttachment
	isSet bool
}

func (v NullableTicketCommentsAttachment) Get() *TicketCommentsAttachment {
	return v.value
}

func (v *NullableTicketCommentsAttachment) Set(val *TicketCommentsAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketCommentsAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketCommentsAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketCommentsAttachment(val *TicketCommentsAttachment) *NullableTicketCommentsAttachment {
	return &NullableTicketCommentsAttachment{value: val, isSet: true}
}

func (v NullableTicketCommentsAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketCommentsAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


