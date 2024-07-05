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

// checks if the Recaptcha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Recaptcha{}

// Recaptcha struct for Recaptcha
type Recaptcha struct {
	Flavor *RecaptchaFlavor `json:"flavor,omitempty"`
	Required *bool `json:"required,omitempty"`
	Sitekey *string `json:"sitekey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Recaptcha Recaptcha

// NewRecaptcha instantiates a new Recaptcha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecaptcha() *Recaptcha {
	this := Recaptcha{}
	var flavor RecaptchaFlavor = RECAPTCHAFLAVOR_GOOGLE
	this.Flavor = &flavor
	return &this
}

// NewRecaptchaWithDefaults instantiates a new Recaptcha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecaptchaWithDefaults() *Recaptcha {
	this := Recaptcha{}
	var flavor RecaptchaFlavor = RECAPTCHAFLAVOR_GOOGLE
	this.Flavor = &flavor
	return &this
}

// GetFlavor returns the Flavor field value if set, zero value otherwise.
func (o *Recaptcha) GetFlavor() RecaptchaFlavor {
	if o == nil || IsNil(o.Flavor) {
		var ret RecaptchaFlavor
		return ret
	}
	return *o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recaptcha) GetFlavorOk() (*RecaptchaFlavor, bool) {
	if o == nil || IsNil(o.Flavor) {
		return nil, false
	}
	return o.Flavor, true
}

// HasFlavor returns a boolean if a field has been set.
func (o *Recaptcha) HasFlavor() bool {
	if o != nil && !IsNil(o.Flavor) {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given RecaptchaFlavor and assigns it to the Flavor field.
func (o *Recaptcha) SetFlavor(v RecaptchaFlavor) {
	o.Flavor = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Recaptcha) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recaptcha) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *Recaptcha) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *Recaptcha) SetRequired(v bool) {
	o.Required = &v
}

// GetSitekey returns the Sitekey field value if set, zero value otherwise.
func (o *Recaptcha) GetSitekey() string {
	if o == nil || IsNil(o.Sitekey) {
		var ret string
		return ret
	}
	return *o.Sitekey
}

// GetSitekeyOk returns a tuple with the Sitekey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Recaptcha) GetSitekeyOk() (*string, bool) {
	if o == nil || IsNil(o.Sitekey) {
		return nil, false
	}
	return o.Sitekey, true
}

// HasSitekey returns a boolean if a field has been set.
func (o *Recaptcha) HasSitekey() bool {
	if o != nil && !IsNil(o.Sitekey) {
		return true
	}

	return false
}

// SetSitekey gets a reference to the given string and assigns it to the Sitekey field.
func (o *Recaptcha) SetSitekey(v string) {
	o.Sitekey = &v
}

func (o Recaptcha) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Recaptcha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flavor) {
		toSerialize["flavor"] = o.Flavor
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Sitekey) {
		toSerialize["sitekey"] = o.Sitekey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Recaptcha) UnmarshalJSON(data []byte) (err error) {
	varRecaptcha := _Recaptcha{}

	err = json.Unmarshal(data, &varRecaptcha)

	if err != nil {
		return err
	}

	*o = Recaptcha(varRecaptcha)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "flavor")
		delete(additionalProperties, "required")
		delete(additionalProperties, "sitekey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecaptcha struct {
	value *Recaptcha
	isSet bool
}

func (v NullableRecaptcha) Get() *Recaptcha {
	return v.value
}

func (v *NullableRecaptcha) Set(val *Recaptcha) {
	v.value = val
	v.isSet = true
}

func (v NullableRecaptcha) IsSet() bool {
	return v.isSet
}

func (v *NullableRecaptcha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecaptcha(val *Recaptcha) *NullableRecaptcha {
	return &NullableRecaptcha{value: val, isSet: true}
}

func (v NullableRecaptcha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecaptcha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


