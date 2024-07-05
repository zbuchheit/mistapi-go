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
	"fmt"
)

// checks if the ClientWirelessStatsGuest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientWirelessStatsGuest{}

// ClientWirelessStatsGuest information about this portal
type ClientWirelessStatsGuest struct {
	// whether this guest is authorized
	Authorized bool `json:"authorized"`
	// when the guest authorization will expire
	AuthorizedExpiringTime float32 `json:"authorized_expiring_time"`
	// when the guest is authorized
	AuthorizedTime float32 `json:"authorized_time"`
	Company string `json:"company"`
	Email string `json:"email"`
	Field1 string `json:"field1"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _ClientWirelessStatsGuest ClientWirelessStatsGuest

// NewClientWirelessStatsGuest instantiates a new ClientWirelessStatsGuest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientWirelessStatsGuest(authorized bool, authorizedExpiringTime float32, authorizedTime float32, company string, email string, field1 string, name string) *ClientWirelessStatsGuest {
	this := ClientWirelessStatsGuest{}
	this.Authorized = authorized
	this.AuthorizedExpiringTime = authorizedExpiringTime
	this.AuthorizedTime = authorizedTime
	this.Company = company
	this.Email = email
	this.Field1 = field1
	this.Name = name
	return &this
}

// NewClientWirelessStatsGuestWithDefaults instantiates a new ClientWirelessStatsGuest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWirelessStatsGuestWithDefaults() *ClientWirelessStatsGuest {
	this := ClientWirelessStatsGuest{}
	var authorized bool = false
	this.Authorized = authorized
	return &this
}

// GetAuthorized returns the Authorized field value
func (o *ClientWirelessStatsGuest) GetAuthorized() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Authorized
}

// GetAuthorizedOk returns a tuple with the Authorized field value
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsGuest) GetAuthorizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorized, true
}

// SetAuthorized sets field value
func (o *ClientWirelessStatsGuest) SetAuthorized(v bool) {
	o.Authorized = v
}

// GetAuthorizedExpiringTime returns the AuthorizedExpiringTime field value
func (o *ClientWirelessStatsGuest) GetAuthorizedExpiringTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AuthorizedExpiringTime
}

// GetAuthorizedExpiringTimeOk returns a tuple with the AuthorizedExpiringTime field value
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsGuest) GetAuthorizedExpiringTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizedExpiringTime, true
}

// SetAuthorizedExpiringTime sets field value
func (o *ClientWirelessStatsGuest) SetAuthorizedExpiringTime(v float32) {
	o.AuthorizedExpiringTime = v
}

// GetAuthorizedTime returns the AuthorizedTime field value
func (o *ClientWirelessStatsGuest) GetAuthorizedTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AuthorizedTime
}

// GetAuthorizedTimeOk returns a tuple with the AuthorizedTime field value
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsGuest) GetAuthorizedTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizedTime, true
}

// SetAuthorizedTime sets field value
func (o *ClientWirelessStatsGuest) SetAuthorizedTime(v float32) {
	o.AuthorizedTime = v
}

// GetCompany returns the Company field value
func (o *ClientWirelessStatsGuest) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsGuest) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *ClientWirelessStatsGuest) SetCompany(v string) {
	o.Company = v
}

// GetEmail returns the Email field value
func (o *ClientWirelessStatsGuest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsGuest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ClientWirelessStatsGuest) SetEmail(v string) {
	o.Email = v
}

// GetField1 returns the Field1 field value
func (o *ClientWirelessStatsGuest) GetField1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field1
}

// GetField1Ok returns a tuple with the Field1 field value
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsGuest) GetField1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field1, true
}

// SetField1 sets field value
func (o *ClientWirelessStatsGuest) SetField1(v string) {
	o.Field1 = v
}

// GetName returns the Name field value
func (o *ClientWirelessStatsGuest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientWirelessStatsGuest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientWirelessStatsGuest) SetName(v string) {
	o.Name = v
}

func (o ClientWirelessStatsGuest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientWirelessStatsGuest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorized"] = o.Authorized
	toSerialize["authorized_expiring_time"] = o.AuthorizedExpiringTime
	toSerialize["authorized_time"] = o.AuthorizedTime
	toSerialize["company"] = o.Company
	toSerialize["email"] = o.Email
	toSerialize["field1"] = o.Field1
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientWirelessStatsGuest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorized",
		"authorized_expiring_time",
		"authorized_time",
		"company",
		"email",
		"field1",
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

	varClientWirelessStatsGuest := _ClientWirelessStatsGuest{}

	err = json.Unmarshal(data, &varClientWirelessStatsGuest)

	if err != nil {
		return err
	}

	*o = ClientWirelessStatsGuest(varClientWirelessStatsGuest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorized")
		delete(additionalProperties, "authorized_expiring_time")
		delete(additionalProperties, "authorized_time")
		delete(additionalProperties, "company")
		delete(additionalProperties, "email")
		delete(additionalProperties, "field1")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientWirelessStatsGuest struct {
	value *ClientWirelessStatsGuest
	isSet bool
}

func (v NullableClientWirelessStatsGuest) Get() *ClientWirelessStatsGuest {
	return v.value
}

func (v *NullableClientWirelessStatsGuest) Set(val *ClientWirelessStatsGuest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientWirelessStatsGuest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientWirelessStatsGuest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientWirelessStatsGuest(val *ClientWirelessStatsGuest) *NullableClientWirelessStatsGuest {
	return &NullableClientWirelessStatsGuest{value: val, isSet: true}
}

func (v NullableClientWirelessStatsGuest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientWirelessStatsGuest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


