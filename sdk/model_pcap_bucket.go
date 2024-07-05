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

// checks if the PcapBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcapBucket{}

// PcapBucket struct for PcapBucket
type PcapBucket struct {
	Bucket string `json:"bucket"`
	AdditionalProperties map[string]interface{}
}

type _PcapBucket PcapBucket

// NewPcapBucket instantiates a new PcapBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcapBucket(bucket string) *PcapBucket {
	this := PcapBucket{}
	this.Bucket = bucket
	return &this
}

// NewPcapBucketWithDefaults instantiates a new PcapBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcapBucketWithDefaults() *PcapBucket {
	this := PcapBucket{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *PcapBucket) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *PcapBucket) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *PcapBucket) SetBucket(v string) {
	o.Bucket = v
}

func (o PcapBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcapBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PcapBucket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket",
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

	varPcapBucket := _PcapBucket{}

	err = json.Unmarshal(data, &varPcapBucket)

	if err != nil {
		return err
	}

	*o = PcapBucket(varPcapBucket)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bucket")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePcapBucket struct {
	value *PcapBucket
	isSet bool
}

func (v NullablePcapBucket) Get() *PcapBucket {
	return v.value
}

func (v *NullablePcapBucket) Set(val *PcapBucket) {
	v.value = val
	v.isSet = true
}

func (v NullablePcapBucket) IsSet() bool {
	return v.isSet
}

func (v *NullablePcapBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcapBucket(val *PcapBucket) *NullablePcapBucket {
	return &NullablePcapBucket{value: val, isSet: true}
}

func (v NullablePcapBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcapBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


