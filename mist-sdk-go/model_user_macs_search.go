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

// checks if the UserMacsSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMacsSearch{}

// UserMacsSearch struct for UserMacsSearch
type UserMacsSearch struct {
	Limit *int32 `json:"limit,omitempty"`
	Page *int32 `json:"page,omitempty"`
	Results []UserMac `json:"results,omitempty"`
	Total *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserMacsSearch UserMacsSearch

// NewUserMacsSearch instantiates a new UserMacsSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMacsSearch() *UserMacsSearch {
	this := UserMacsSearch{}
	return &this
}

// NewUserMacsSearchWithDefaults instantiates a new UserMacsSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMacsSearchWithDefaults() *UserMacsSearch {
	this := UserMacsSearch{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *UserMacsSearch) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMacsSearch) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *UserMacsSearch) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *UserMacsSearch) SetLimit(v int32) {
	o.Limit = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UserMacsSearch) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMacsSearch) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UserMacsSearch) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UserMacsSearch) SetPage(v int32) {
	o.Page = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *UserMacsSearch) GetResults() []UserMac {
	if o == nil || IsNil(o.Results) {
		var ret []UserMac
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMacsSearch) GetResultsOk() ([]UserMac, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *UserMacsSearch) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []UserMac and assigns it to the Results field.
func (o *UserMacsSearch) SetResults(v []UserMac) {
	o.Results = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UserMacsSearch) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMacsSearch) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UserMacsSearch) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UserMacsSearch) SetTotal(v int32) {
	o.Total = &v
}

func (o UserMacsSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMacsSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserMacsSearch) UnmarshalJSON(data []byte) (err error) {
	varUserMacsSearch := _UserMacsSearch{}

	err = json.Unmarshal(data, &varUserMacsSearch)

	if err != nil {
		return err
	}

	*o = UserMacsSearch(varUserMacsSearch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "page")
		delete(additionalProperties, "results")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserMacsSearch struct {
	value *UserMacsSearch
	isSet bool
}

func (v NullableUserMacsSearch) Get() *UserMacsSearch {
	return v.value
}

func (v *NullableUserMacsSearch) Set(val *UserMacsSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMacsSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMacsSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMacsSearch(val *UserMacsSearch) *NullableUserMacsSearch {
	return &NullableUserMacsSearch{value: val, isSet: true}
}

func (v NullableUserMacsSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMacsSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


