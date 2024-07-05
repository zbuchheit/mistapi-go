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

// checks if the WebsocketSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebsocketSession{}

// WebsocketSession struct for WebsocketSession
type WebsocketSession struct {
	Session string `json:"session"`
	AdditionalProperties map[string]interface{}
}

type _WebsocketSession WebsocketSession

// NewWebsocketSession instantiates a new WebsocketSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketSession(session string) *WebsocketSession {
	this := WebsocketSession{}
	this.Session = session
	return &this
}

// NewWebsocketSessionWithDefaults instantiates a new WebsocketSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketSessionWithDefaults() *WebsocketSession {
	this := WebsocketSession{}
	return &this
}

// GetSession returns the Session field value
func (o *WebsocketSession) GetSession() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *WebsocketSession) GetSessionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *WebsocketSession) SetSession(v string) {
	o.Session = v
}

func (o WebsocketSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebsocketSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebsocketSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"session",
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

	varWebsocketSession := _WebsocketSession{}

	err = json.Unmarshal(data, &varWebsocketSession)

	if err != nil {
		return err
	}

	*o = WebsocketSession(varWebsocketSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "session")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebsocketSession struct {
	value *WebsocketSession
	isSet bool
}

func (v NullableWebsocketSession) Get() *WebsocketSession {
	return v.value
}

func (v *NullableWebsocketSession) Set(val *WebsocketSession) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketSession) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketSession(val *WebsocketSession) *NullableWebsocketSession {
	return &NullableWebsocketSession{value: val, isSet: true}
}

func (v NullableWebsocketSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


