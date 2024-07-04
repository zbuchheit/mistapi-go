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

// checks if the TestTelstra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestTelstra{}

// TestTelstra struct for TestTelstra
type TestTelstra struct {
	// Telstra client id
	TelstraClientId string `json:"telstra_client_id"`
	// Telstra client secret
	TelstraClientSecret string `json:"telstra_client_secret"`
	// Phone number of the recipient of SMS with country code
	To string `json:"to"`
	AdditionalProperties map[string]interface{}
}

type _TestTelstra TestTelstra

// NewTestTelstra instantiates a new TestTelstra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestTelstra(telstraClientId string, telstraClientSecret string, to string) *TestTelstra {
	this := TestTelstra{}
	this.TelstraClientId = telstraClientId
	this.TelstraClientSecret = telstraClientSecret
	this.To = to
	return &this
}

// NewTestTelstraWithDefaults instantiates a new TestTelstra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestTelstraWithDefaults() *TestTelstra {
	this := TestTelstra{}
	return &this
}

// GetTelstraClientId returns the TelstraClientId field value
func (o *TestTelstra) GetTelstraClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelstraClientId
}

// GetTelstraClientIdOk returns a tuple with the TelstraClientId field value
// and a boolean to check if the value has been set.
func (o *TestTelstra) GetTelstraClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelstraClientId, true
}

// SetTelstraClientId sets field value
func (o *TestTelstra) SetTelstraClientId(v string) {
	o.TelstraClientId = v
}

// GetTelstraClientSecret returns the TelstraClientSecret field value
func (o *TestTelstra) GetTelstraClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelstraClientSecret
}

// GetTelstraClientSecretOk returns a tuple with the TelstraClientSecret field value
// and a boolean to check if the value has been set.
func (o *TestTelstra) GetTelstraClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelstraClientSecret, true
}

// SetTelstraClientSecret sets field value
func (o *TestTelstra) SetTelstraClientSecret(v string) {
	o.TelstraClientSecret = v
}

// GetTo returns the To field value
func (o *TestTelstra) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *TestTelstra) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *TestTelstra) SetTo(v string) {
	o.To = v
}

func (o TestTelstra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestTelstra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["telstra_client_id"] = o.TelstraClientId
	toSerialize["telstra_client_secret"] = o.TelstraClientSecret
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestTelstra) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"telstra_client_id",
		"telstra_client_secret",
		"to",
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

	varTestTelstra := _TestTelstra{}

	err = json.Unmarshal(data, &varTestTelstra)

	if err != nil {
		return err
	}

	*o = TestTelstra(varTestTelstra)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "telstra_client_id")
		delete(additionalProperties, "telstra_client_secret")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestTelstra struct {
	value *TestTelstra
	isSet bool
}

func (v NullableTestTelstra) Get() *TestTelstra {
	return v.value
}

func (v *NullableTestTelstra) Set(val *TestTelstra) {
	v.value = val
	v.isSet = true
}

func (v NullableTestTelstra) IsSet() bool {
	return v.isSet
}

func (v *NullableTestTelstra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestTelstra(val *TestTelstra) *NullableTestTelstra {
	return &NullableTestTelstra{value: val, isSet: true}
}

func (v NullableTestTelstra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestTelstra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


