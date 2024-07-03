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

// checks if the MarvisClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarvisClient{}

// MarvisClient struct for MarvisClient
type MarvisClient struct {
	Diabled *bool `json:"diabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// in MDM, add `--provision_url <provision_url>` to the instlal command
	ProvisionUrl *string `json:"provision_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarvisClient MarvisClient

// NewMarvisClient instantiates a new MarvisClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarvisClient() *MarvisClient {
	this := MarvisClient{}
	var diabled bool = false
	this.Diabled = &diabled
	return &this
}

// NewMarvisClientWithDefaults instantiates a new MarvisClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarvisClientWithDefaults() *MarvisClient {
	this := MarvisClient{}
	var diabled bool = false
	this.Diabled = &diabled
	return &this
}

// GetDiabled returns the Diabled field value if set, zero value otherwise.
func (o *MarvisClient) GetDiabled() bool {
	if o == nil || IsNil(o.Diabled) {
		var ret bool
		return ret
	}
	return *o.Diabled
}

// GetDiabledOk returns a tuple with the Diabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarvisClient) GetDiabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Diabled) {
		return nil, false
	}
	return o.Diabled, true
}

// HasDiabled returns a boolean if a field has been set.
func (o *MarvisClient) HasDiabled() bool {
	if o != nil && !IsNil(o.Diabled) {
		return true
	}

	return false
}

// SetDiabled gets a reference to the given bool and assigns it to the Diabled field.
func (o *MarvisClient) SetDiabled(v bool) {
	o.Diabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MarvisClient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarvisClient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MarvisClient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MarvisClient) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarvisClient) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarvisClient) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarvisClient) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarvisClient) SetName(v string) {
	o.Name = &v
}

// GetProvisionUrl returns the ProvisionUrl field value if set, zero value otherwise.
func (o *MarvisClient) GetProvisionUrl() string {
	if o == nil || IsNil(o.ProvisionUrl) {
		var ret string
		return ret
	}
	return *o.ProvisionUrl
}

// GetProvisionUrlOk returns a tuple with the ProvisionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarvisClient) GetProvisionUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProvisionUrl) {
		return nil, false
	}
	return o.ProvisionUrl, true
}

// HasProvisionUrl returns a boolean if a field has been set.
func (o *MarvisClient) HasProvisionUrl() bool {
	if o != nil && !IsNil(o.ProvisionUrl) {
		return true
	}

	return false
}

// SetProvisionUrl gets a reference to the given string and assigns it to the ProvisionUrl field.
func (o *MarvisClient) SetProvisionUrl(v string) {
	o.ProvisionUrl = &v
}

func (o MarvisClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarvisClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Diabled) {
		toSerialize["diabled"] = o.Diabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProvisionUrl) {
		toSerialize["provision_url"] = o.ProvisionUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarvisClient) UnmarshalJSON(data []byte) (err error) {
	varMarvisClient := _MarvisClient{}

	err = json.Unmarshal(data, &varMarvisClient)

	if err != nil {
		return err
	}

	*o = MarvisClient(varMarvisClient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "diabled")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "provision_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarvisClient struct {
	value *MarvisClient
	isSet bool
}

func (v NullableMarvisClient) Get() *MarvisClient {
	return v.value
}

func (v *NullableMarvisClient) Set(val *MarvisClient) {
	v.value = val
	v.isSet = true
}

func (v NullableMarvisClient) IsSet() bool {
	return v.isSet
}

func (v *NullableMarvisClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarvisClient(val *MarvisClient) *NullableMarvisClient {
	return &NullableMarvisClient{value: val, isSet: true}
}

func (v NullableMarvisClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarvisClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


