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
)

// checks if the CaptureScanAps type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureScanAps{}

// CaptureScanAps Property key is the AP MAC address (e.g. \"5c5b35000001\"). All optionals, parent parameters will be used if not defined
type CaptureScanAps struct {
	Band *CaptureScanApsBand `json:"band,omitempty"`
	// specify the channel value where scan PCAP has to be started
	Channel *string `json:"channel,omitempty"`
	// tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
	TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
	// specify the bandwidth value with respect to the channel.
	Width *string `json:"width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CaptureScanAps CaptureScanAps

// NewCaptureScanAps instantiates a new CaptureScanAps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureScanAps() *CaptureScanAps {
	this := CaptureScanAps{}
	var band CaptureScanApsBand = CAPTURESCANAPSBAND__24
	this.Band = &band
	return &this
}

// NewCaptureScanApsWithDefaults instantiates a new CaptureScanAps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureScanApsWithDefaults() *CaptureScanAps {
	this := CaptureScanAps{}
	var band CaptureScanApsBand = CAPTURESCANAPSBAND__24
	this.Band = &band
	return &this
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *CaptureScanAps) GetBand() CaptureScanApsBand {
	if o == nil || IsNil(o.Band) {
		var ret CaptureScanApsBand
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureScanAps) GetBandOk() (*CaptureScanApsBand, bool) {
	if o == nil || IsNil(o.Band) {
		return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *CaptureScanAps) HasBand() bool {
	if o != nil && !IsNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given CaptureScanApsBand and assigns it to the Band field.
func (o *CaptureScanAps) SetBand(v CaptureScanApsBand) {
	o.Band = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *CaptureScanAps) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureScanAps) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *CaptureScanAps) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *CaptureScanAps) SetChannel(v string) {
	o.Channel = &v
}

// GetTcpdumpExpression returns the TcpdumpExpression field value if set, zero value otherwise.
func (o *CaptureScanAps) GetTcpdumpExpression() string {
	if o == nil || IsNil(o.TcpdumpExpression) {
		var ret string
		return ret
	}
	return *o.TcpdumpExpression
}

// GetTcpdumpExpressionOk returns a tuple with the TcpdumpExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureScanAps) GetTcpdumpExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.TcpdumpExpression) {
		return nil, false
	}
	return o.TcpdumpExpression, true
}

// HasTcpdumpExpression returns a boolean if a field has been set.
func (o *CaptureScanAps) HasTcpdumpExpression() bool {
	if o != nil && !IsNil(o.TcpdumpExpression) {
		return true
	}

	return false
}

// SetTcpdumpExpression gets a reference to the given string and assigns it to the TcpdumpExpression field.
func (o *CaptureScanAps) SetTcpdumpExpression(v string) {
	o.TcpdumpExpression = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CaptureScanAps) GetWidth() string {
	if o == nil || IsNil(o.Width) {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureScanAps) GetWidthOk() (*string, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CaptureScanAps) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *CaptureScanAps) SetWidth(v string) {
	o.Width = &v
}

func (o CaptureScanAps) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureScanAps) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.TcpdumpExpression) {
		toSerialize["tcpdump_expression"] = o.TcpdumpExpression
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaptureScanAps) UnmarshalJSON(data []byte) (err error) {
	varCaptureScanAps := _CaptureScanAps{}

	err = json.Unmarshal(data, &varCaptureScanAps)

	if err != nil {
		return err
	}

	*o = CaptureScanAps(varCaptureScanAps)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "band")
		delete(additionalProperties, "channel")
		delete(additionalProperties, "tcpdump_expression")
		delete(additionalProperties, "width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaptureScanAps struct {
	value *CaptureScanAps
	isSet bool
}

func (v NullableCaptureScanAps) Get() *CaptureScanAps {
	return v.value
}

func (v *NullableCaptureScanAps) Set(val *CaptureScanAps) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureScanAps) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureScanAps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureScanAps(val *CaptureScanAps) *NullableCaptureScanAps {
	return &NullableCaptureScanAps{value: val, isSet: true}
}

func (v NullableCaptureScanAps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureScanAps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


