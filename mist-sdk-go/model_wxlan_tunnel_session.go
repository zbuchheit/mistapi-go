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
)

// checks if the WxlanTunnelSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WxlanTunnelSession{}

// WxlanTunnelSession struct for WxlanTunnelSession
type WxlanTunnelSession struct {
	// if `use_ap_as_session_ids`==`true`, only apmac is supported right now. This is the name WLAN should use for wxtunnel_remote_id
	ApAsSessionId *string `json:"ap_as_session_id,omitempty"`
	// optional, user-specified string for display purpose
	Comment *string `json:"comment,omitempty"`
	EnableCookie *bool `json:"enable_cookie,omitempty"`
	Ethertype *WxlanTunnelSessionEthertype `json:"ethertype,omitempty"`
	// 1-4294967295
	LocalSessionId *int32 `json:"local_session_id,omitempty"`
	// optional. Enables the pseudo 802.1ad QinQ mode where the AP device drops the outer vlan tag (QinQ). This mode is useful when tunneling Mist AP’s to some aggregation routers.
	Pseudo8021adEnabled *bool `json:"pseudo_802.1ad_enabled,omitempty"`
	// remote-id of the session, has to be unique in the same tunnel
	RemoteId *string `json:"remote_id,omitempty"`
	// 1-4294967295
	RemoteSessionId *int32 `json:"remote_session_id,omitempty"`
	// whether to use AP (last 4 bytes of MAC currently) as session ids
	UseApAsSessionIds *bool `json:"use_ap_as_session_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WxlanTunnelSession WxlanTunnelSession

// NewWxlanTunnelSession instantiates a new WxlanTunnelSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWxlanTunnelSession() *WxlanTunnelSession {
	this := WxlanTunnelSession{}
	var pseudo8021adEnabled bool = false
	this.Pseudo8021adEnabled = &pseudo8021adEnabled
	var useApAsSessionIds bool = false
	this.UseApAsSessionIds = &useApAsSessionIds
	return &this
}

// NewWxlanTunnelSessionWithDefaults instantiates a new WxlanTunnelSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWxlanTunnelSessionWithDefaults() *WxlanTunnelSession {
	this := WxlanTunnelSession{}
	var pseudo8021adEnabled bool = false
	this.Pseudo8021adEnabled = &pseudo8021adEnabled
	var useApAsSessionIds bool = false
	this.UseApAsSessionIds = &useApAsSessionIds
	return &this
}

// GetApAsSessionId returns the ApAsSessionId field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetApAsSessionId() string {
	if o == nil || IsNil(o.ApAsSessionId) {
		var ret string
		return ret
	}
	return *o.ApAsSessionId
}

// GetApAsSessionIdOk returns a tuple with the ApAsSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetApAsSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApAsSessionId) {
		return nil, false
	}
	return o.ApAsSessionId, true
}

// HasApAsSessionId returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasApAsSessionId() bool {
	if o != nil && !IsNil(o.ApAsSessionId) {
		return true
	}

	return false
}

// SetApAsSessionId gets a reference to the given string and assigns it to the ApAsSessionId field.
func (o *WxlanTunnelSession) SetApAsSessionId(v string) {
	o.ApAsSessionId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *WxlanTunnelSession) SetComment(v string) {
	o.Comment = &v
}

// GetEnableCookie returns the EnableCookie field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetEnableCookie() bool {
	if o == nil || IsNil(o.EnableCookie) {
		var ret bool
		return ret
	}
	return *o.EnableCookie
}

// GetEnableCookieOk returns a tuple with the EnableCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetEnableCookieOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCookie) {
		return nil, false
	}
	return o.EnableCookie, true
}

// HasEnableCookie returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasEnableCookie() bool {
	if o != nil && !IsNil(o.EnableCookie) {
		return true
	}

	return false
}

// SetEnableCookie gets a reference to the given bool and assigns it to the EnableCookie field.
func (o *WxlanTunnelSession) SetEnableCookie(v bool) {
	o.EnableCookie = &v
}

// GetEthertype returns the Ethertype field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetEthertype() WxlanTunnelSessionEthertype {
	if o == nil || IsNil(o.Ethertype) {
		var ret WxlanTunnelSessionEthertype
		return ret
	}
	return *o.Ethertype
}

// GetEthertypeOk returns a tuple with the Ethertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetEthertypeOk() (*WxlanTunnelSessionEthertype, bool) {
	if o == nil || IsNil(o.Ethertype) {
		return nil, false
	}
	return o.Ethertype, true
}

// HasEthertype returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasEthertype() bool {
	if o != nil && !IsNil(o.Ethertype) {
		return true
	}

	return false
}

// SetEthertype gets a reference to the given WxlanTunnelSessionEthertype and assigns it to the Ethertype field.
func (o *WxlanTunnelSession) SetEthertype(v WxlanTunnelSessionEthertype) {
	o.Ethertype = &v
}

// GetLocalSessionId returns the LocalSessionId field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetLocalSessionId() int32 {
	if o == nil || IsNil(o.LocalSessionId) {
		var ret int32
		return ret
	}
	return *o.LocalSessionId
}

// GetLocalSessionIdOk returns a tuple with the LocalSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetLocalSessionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalSessionId) {
		return nil, false
	}
	return o.LocalSessionId, true
}

// HasLocalSessionId returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasLocalSessionId() bool {
	if o != nil && !IsNil(o.LocalSessionId) {
		return true
	}

	return false
}

// SetLocalSessionId gets a reference to the given int32 and assigns it to the LocalSessionId field.
func (o *WxlanTunnelSession) SetLocalSessionId(v int32) {
	o.LocalSessionId = &v
}

// GetPseudo8021adEnabled returns the Pseudo8021adEnabled field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetPseudo8021adEnabled() bool {
	if o == nil || IsNil(o.Pseudo8021adEnabled) {
		var ret bool
		return ret
	}
	return *o.Pseudo8021adEnabled
}

// GetPseudo8021adEnabledOk returns a tuple with the Pseudo8021adEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetPseudo8021adEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Pseudo8021adEnabled) {
		return nil, false
	}
	return o.Pseudo8021adEnabled, true
}

// HasPseudo8021adEnabled returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasPseudo8021adEnabled() bool {
	if o != nil && !IsNil(o.Pseudo8021adEnabled) {
		return true
	}

	return false
}

// SetPseudo8021adEnabled gets a reference to the given bool and assigns it to the Pseudo8021adEnabled field.
func (o *WxlanTunnelSession) SetPseudo8021adEnabled(v bool) {
	o.Pseudo8021adEnabled = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *WxlanTunnelSession) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetRemoteSessionId returns the RemoteSessionId field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetRemoteSessionId() int32 {
	if o == nil || IsNil(o.RemoteSessionId) {
		var ret int32
		return ret
	}
	return *o.RemoteSessionId
}

// GetRemoteSessionIdOk returns a tuple with the RemoteSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetRemoteSessionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RemoteSessionId) {
		return nil, false
	}
	return o.RemoteSessionId, true
}

// HasRemoteSessionId returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasRemoteSessionId() bool {
	if o != nil && !IsNil(o.RemoteSessionId) {
		return true
	}

	return false
}

// SetRemoteSessionId gets a reference to the given int32 and assigns it to the RemoteSessionId field.
func (o *WxlanTunnelSession) SetRemoteSessionId(v int32) {
	o.RemoteSessionId = &v
}

// GetUseApAsSessionIds returns the UseApAsSessionIds field value if set, zero value otherwise.
func (o *WxlanTunnelSession) GetUseApAsSessionIds() bool {
	if o == nil || IsNil(o.UseApAsSessionIds) {
		var ret bool
		return ret
	}
	return *o.UseApAsSessionIds
}

// GetUseApAsSessionIdsOk returns a tuple with the UseApAsSessionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WxlanTunnelSession) GetUseApAsSessionIdsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseApAsSessionIds) {
		return nil, false
	}
	return o.UseApAsSessionIds, true
}

// HasUseApAsSessionIds returns a boolean if a field has been set.
func (o *WxlanTunnelSession) HasUseApAsSessionIds() bool {
	if o != nil && !IsNil(o.UseApAsSessionIds) {
		return true
	}

	return false
}

// SetUseApAsSessionIds gets a reference to the given bool and assigns it to the UseApAsSessionIds field.
func (o *WxlanTunnelSession) SetUseApAsSessionIds(v bool) {
	o.UseApAsSessionIds = &v
}

func (o WxlanTunnelSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WxlanTunnelSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApAsSessionId) {
		toSerialize["ap_as_session_id"] = o.ApAsSessionId
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.EnableCookie) {
		toSerialize["enable_cookie"] = o.EnableCookie
	}
	if !IsNil(o.Ethertype) {
		toSerialize["ethertype"] = o.Ethertype
	}
	if !IsNil(o.LocalSessionId) {
		toSerialize["local_session_id"] = o.LocalSessionId
	}
	if !IsNil(o.Pseudo8021adEnabled) {
		toSerialize["pseudo_802.1ad_enabled"] = o.Pseudo8021adEnabled
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remote_id"] = o.RemoteId
	}
	if !IsNil(o.RemoteSessionId) {
		toSerialize["remote_session_id"] = o.RemoteSessionId
	}
	if !IsNil(o.UseApAsSessionIds) {
		toSerialize["use_ap_as_session_ids"] = o.UseApAsSessionIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WxlanTunnelSession) UnmarshalJSON(data []byte) (err error) {
	varWxlanTunnelSession := _WxlanTunnelSession{}

	err = json.Unmarshal(data, &varWxlanTunnelSession)

	if err != nil {
		return err
	}

	*o = WxlanTunnelSession(varWxlanTunnelSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_as_session_id")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "enable_cookie")
		delete(additionalProperties, "ethertype")
		delete(additionalProperties, "local_session_id")
		delete(additionalProperties, "pseudo_802.1ad_enabled")
		delete(additionalProperties, "remote_id")
		delete(additionalProperties, "remote_session_id")
		delete(additionalProperties, "use_ap_as_session_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWxlanTunnelSession struct {
	value *WxlanTunnelSession
	isSet bool
}

func (v NullableWxlanTunnelSession) Get() *WxlanTunnelSession {
	return v.value
}

func (v *NullableWxlanTunnelSession) Set(val *WxlanTunnelSession) {
	v.value = val
	v.isSet = true
}

func (v NullableWxlanTunnelSession) IsSet() bool {
	return v.isSet
}

func (v *NullableWxlanTunnelSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWxlanTunnelSession(val *WxlanTunnelSession) *NullableWxlanTunnelSession {
	return &NullableWxlanTunnelSession{value: val, isSet: true}
}

func (v NullableWxlanTunnelSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWxlanTunnelSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


