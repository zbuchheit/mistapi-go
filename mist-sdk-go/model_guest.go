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

// checks if the Guest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Guest{}

// Guest Guest
type Guest struct {
	// whether the guest is current authorized
	Authorized *bool `json:"authorized,omitempty"`
	// when the authorization would expire
	AuthorizedExpiringTime *float32 `json:"authorized_expiring_time,omitempty"`
	// when the guest was authorized
	AuthorizedTime *float32 `json:"authorized_time,omitempty"`
	// optional, the info provided by user
	Company *string `json:"company,omitempty"`
	// optional, the info provided by user
	Email *string `json:"email,omitempty"`
	// optional, the info provided by user
	Field1 *string `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	Field3 *string `json:"field3,omitempty"`
	Field4 *string `json:"field4,omitempty"`
	// mac
	Mac *string `json:"mac,omitempty"`
	// minutes, the maximum is 259200 (180 days)
	Minutes *int32 `json:"minutes,omitempty"`
	// optional, the info provided by user
	Name *string `json:"name,omitempty"`
	Ssid *string `json:"ssid,omitempty"`
	WlanId *string `json:"wlan_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Guest Guest

// NewGuest instantiates a new Guest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuest() *Guest {
	this := Guest{}
	return &this
}

// NewGuestWithDefaults instantiates a new Guest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestWithDefaults() *Guest {
	this := Guest{}
	return &this
}

// GetAuthorized returns the Authorized field value if set, zero value otherwise.
func (o *Guest) GetAuthorized() bool {
	if o == nil || IsNil(o.Authorized) {
		var ret bool
		return ret
	}
	return *o.Authorized
}

// GetAuthorizedOk returns a tuple with the Authorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetAuthorizedOk() (*bool, bool) {
	if o == nil || IsNil(o.Authorized) {
		return nil, false
	}
	return o.Authorized, true
}

// HasAuthorized returns a boolean if a field has been set.
func (o *Guest) HasAuthorized() bool {
	if o != nil && !IsNil(o.Authorized) {
		return true
	}

	return false
}

// SetAuthorized gets a reference to the given bool and assigns it to the Authorized field.
func (o *Guest) SetAuthorized(v bool) {
	o.Authorized = &v
}

// GetAuthorizedExpiringTime returns the AuthorizedExpiringTime field value if set, zero value otherwise.
func (o *Guest) GetAuthorizedExpiringTime() float32 {
	if o == nil || IsNil(o.AuthorizedExpiringTime) {
		var ret float32
		return ret
	}
	return *o.AuthorizedExpiringTime
}

// GetAuthorizedExpiringTimeOk returns a tuple with the AuthorizedExpiringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetAuthorizedExpiringTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.AuthorizedExpiringTime) {
		return nil, false
	}
	return o.AuthorizedExpiringTime, true
}

// HasAuthorizedExpiringTime returns a boolean if a field has been set.
func (o *Guest) HasAuthorizedExpiringTime() bool {
	if o != nil && !IsNil(o.AuthorizedExpiringTime) {
		return true
	}

	return false
}

// SetAuthorizedExpiringTime gets a reference to the given float32 and assigns it to the AuthorizedExpiringTime field.
func (o *Guest) SetAuthorizedExpiringTime(v float32) {
	o.AuthorizedExpiringTime = &v
}

// GetAuthorizedTime returns the AuthorizedTime field value if set, zero value otherwise.
func (o *Guest) GetAuthorizedTime() float32 {
	if o == nil || IsNil(o.AuthorizedTime) {
		var ret float32
		return ret
	}
	return *o.AuthorizedTime
}

// GetAuthorizedTimeOk returns a tuple with the AuthorizedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetAuthorizedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.AuthorizedTime) {
		return nil, false
	}
	return o.AuthorizedTime, true
}

// HasAuthorizedTime returns a boolean if a field has been set.
func (o *Guest) HasAuthorizedTime() bool {
	if o != nil && !IsNil(o.AuthorizedTime) {
		return true
	}

	return false
}

// SetAuthorizedTime gets a reference to the given float32 and assigns it to the AuthorizedTime field.
func (o *Guest) SetAuthorizedTime(v float32) {
	o.AuthorizedTime = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Guest) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Guest) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *Guest) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Guest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Guest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Guest) SetEmail(v string) {
	o.Email = &v
}

// GetField1 returns the Field1 field value if set, zero value otherwise.
func (o *Guest) GetField1() string {
	if o == nil || IsNil(o.Field1) {
		var ret string
		return ret
	}
	return *o.Field1
}

// GetField1Ok returns a tuple with the Field1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetField1Ok() (*string, bool) {
	if o == nil || IsNil(o.Field1) {
		return nil, false
	}
	return o.Field1, true
}

// HasField1 returns a boolean if a field has been set.
func (o *Guest) HasField1() bool {
	if o != nil && !IsNil(o.Field1) {
		return true
	}

	return false
}

// SetField1 gets a reference to the given string and assigns it to the Field1 field.
func (o *Guest) SetField1(v string) {
	o.Field1 = &v
}

// GetField2 returns the Field2 field value if set, zero value otherwise.
func (o *Guest) GetField2() string {
	if o == nil || IsNil(o.Field2) {
		var ret string
		return ret
	}
	return *o.Field2
}

// GetField2Ok returns a tuple with the Field2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetField2Ok() (*string, bool) {
	if o == nil || IsNil(o.Field2) {
		return nil, false
	}
	return o.Field2, true
}

// HasField2 returns a boolean if a field has been set.
func (o *Guest) HasField2() bool {
	if o != nil && !IsNil(o.Field2) {
		return true
	}

	return false
}

// SetField2 gets a reference to the given string and assigns it to the Field2 field.
func (o *Guest) SetField2(v string) {
	o.Field2 = &v
}

// GetField3 returns the Field3 field value if set, zero value otherwise.
func (o *Guest) GetField3() string {
	if o == nil || IsNil(o.Field3) {
		var ret string
		return ret
	}
	return *o.Field3
}

// GetField3Ok returns a tuple with the Field3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetField3Ok() (*string, bool) {
	if o == nil || IsNil(o.Field3) {
		return nil, false
	}
	return o.Field3, true
}

// HasField3 returns a boolean if a field has been set.
func (o *Guest) HasField3() bool {
	if o != nil && !IsNil(o.Field3) {
		return true
	}

	return false
}

// SetField3 gets a reference to the given string and assigns it to the Field3 field.
func (o *Guest) SetField3(v string) {
	o.Field3 = &v
}

// GetField4 returns the Field4 field value if set, zero value otherwise.
func (o *Guest) GetField4() string {
	if o == nil || IsNil(o.Field4) {
		var ret string
		return ret
	}
	return *o.Field4
}

// GetField4Ok returns a tuple with the Field4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetField4Ok() (*string, bool) {
	if o == nil || IsNil(o.Field4) {
		return nil, false
	}
	return o.Field4, true
}

// HasField4 returns a boolean if a field has been set.
func (o *Guest) HasField4() bool {
	if o != nil && !IsNil(o.Field4) {
		return true
	}

	return false
}

// SetField4 gets a reference to the given string and assigns it to the Field4 field.
func (o *Guest) SetField4(v string) {
	o.Field4 = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *Guest) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *Guest) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *Guest) SetMac(v string) {
	o.Mac = &v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *Guest) GetMinutes() int32 {
	if o == nil || IsNil(o.Minutes) {
		var ret int32
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *Guest) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given int32 and assigns it to the Minutes field.
func (o *Guest) SetMinutes(v int32) {
	o.Minutes = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Guest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Guest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Guest) SetName(v string) {
	o.Name = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *Guest) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *Guest) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *Guest) SetSsid(v string) {
	o.Ssid = &v
}

// GetWlanId returns the WlanId field value if set, zero value otherwise.
func (o *Guest) GetWlanId() string {
	if o == nil || IsNil(o.WlanId) {
		var ret string
		return ret
	}
	return *o.WlanId
}

// GetWlanIdOk returns a tuple with the WlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Guest) GetWlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.WlanId) {
		return nil, false
	}
	return o.WlanId, true
}

// HasWlanId returns a boolean if a field has been set.
func (o *Guest) HasWlanId() bool {
	if o != nil && !IsNil(o.WlanId) {
		return true
	}

	return false
}

// SetWlanId gets a reference to the given string and assigns it to the WlanId field.
func (o *Guest) SetWlanId(v string) {
	o.WlanId = &v
}

func (o Guest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Guest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authorized) {
		toSerialize["authorized"] = o.Authorized
	}
	if !IsNil(o.AuthorizedExpiringTime) {
		toSerialize["authorized_expiring_time"] = o.AuthorizedExpiringTime
	}
	if !IsNil(o.AuthorizedTime) {
		toSerialize["authorized_time"] = o.AuthorizedTime
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Field1) {
		toSerialize["field1"] = o.Field1
	}
	if !IsNil(o.Field2) {
		toSerialize["field2"] = o.Field2
	}
	if !IsNil(o.Field3) {
		toSerialize["field3"] = o.Field3
	}
	if !IsNil(o.Field4) {
		toSerialize["field4"] = o.Field4
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.WlanId) {
		toSerialize["wlan_id"] = o.WlanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Guest) UnmarshalJSON(data []byte) (err error) {
	varGuest := _Guest{}

	err = json.Unmarshal(data, &varGuest)

	if err != nil {
		return err
	}

	*o = Guest(varGuest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorized")
		delete(additionalProperties, "authorized_expiring_time")
		delete(additionalProperties, "authorized_time")
		delete(additionalProperties, "company")
		delete(additionalProperties, "email")
		delete(additionalProperties, "field1")
		delete(additionalProperties, "field2")
		delete(additionalProperties, "field3")
		delete(additionalProperties, "field4")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "minutes")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "wlan_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGuest struct {
	value *Guest
	isSet bool
}

func (v NullableGuest) Get() *Guest {
	return v.value
}

func (v *NullableGuest) Set(val *Guest) {
	v.value = val
	v.isSet = true
}

func (v NullableGuest) IsSet() bool {
	return v.isSet
}

func (v *NullableGuest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuest(val *Guest) *NullableGuest {
	return &NullableGuest{value: val, isSet: true}
}

func (v NullableGuest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


