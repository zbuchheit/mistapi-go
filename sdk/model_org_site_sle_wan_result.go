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

// checks if the OrgSiteSleWanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSiteSleWanResult{}

// OrgSiteSleWanResult struct for OrgSiteSleWanResult
type OrgSiteSleWanResult struct {
	ApplicationHealth *float32 `json:"application_health,omitempty"`
	GatewayHealth *float32 `json:"gateway-health,omitempty"`
	NumClients *float32 `json:"num_clients,omitempty"`
	NumGateways *float32 `json:"num_gateways,omitempty"`
	SiteId string `json:"site_id"`
	WanLinkHealth *float32 `json:"wan-link-health,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSiteSleWanResult OrgSiteSleWanResult

// NewOrgSiteSleWanResult instantiates a new OrgSiteSleWanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSiteSleWanResult(siteId string) *OrgSiteSleWanResult {
	this := OrgSiteSleWanResult{}
	this.SiteId = siteId
	return &this
}

// NewOrgSiteSleWanResultWithDefaults instantiates a new OrgSiteSleWanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSiteSleWanResultWithDefaults() *OrgSiteSleWanResult {
	this := OrgSiteSleWanResult{}
	return &this
}

// GetApplicationHealth returns the ApplicationHealth field value if set, zero value otherwise.
func (o *OrgSiteSleWanResult) GetApplicationHealth() float32 {
	if o == nil || IsNil(o.ApplicationHealth) {
		var ret float32
		return ret
	}
	return *o.ApplicationHealth
}

// GetApplicationHealthOk returns a tuple with the ApplicationHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWanResult) GetApplicationHealthOk() (*float32, bool) {
	if o == nil || IsNil(o.ApplicationHealth) {
		return nil, false
	}
	return o.ApplicationHealth, true
}

// HasApplicationHealth returns a boolean if a field has been set.
func (o *OrgSiteSleWanResult) HasApplicationHealth() bool {
	if o != nil && !IsNil(o.ApplicationHealth) {
		return true
	}

	return false
}

// SetApplicationHealth gets a reference to the given float32 and assigns it to the ApplicationHealth field.
func (o *OrgSiteSleWanResult) SetApplicationHealth(v float32) {
	o.ApplicationHealth = &v
}

// GetGatewayHealth returns the GatewayHealth field value if set, zero value otherwise.
func (o *OrgSiteSleWanResult) GetGatewayHealth() float32 {
	if o == nil || IsNil(o.GatewayHealth) {
		var ret float32
		return ret
	}
	return *o.GatewayHealth
}

// GetGatewayHealthOk returns a tuple with the GatewayHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWanResult) GetGatewayHealthOk() (*float32, bool) {
	if o == nil || IsNil(o.GatewayHealth) {
		return nil, false
	}
	return o.GatewayHealth, true
}

// HasGatewayHealth returns a boolean if a field has been set.
func (o *OrgSiteSleWanResult) HasGatewayHealth() bool {
	if o != nil && !IsNil(o.GatewayHealth) {
		return true
	}

	return false
}

// SetGatewayHealth gets a reference to the given float32 and assigns it to the GatewayHealth field.
func (o *OrgSiteSleWanResult) SetGatewayHealth(v float32) {
	o.GatewayHealth = &v
}

// GetNumClients returns the NumClients field value if set, zero value otherwise.
func (o *OrgSiteSleWanResult) GetNumClients() float32 {
	if o == nil || IsNil(o.NumClients) {
		var ret float32
		return ret
	}
	return *o.NumClients
}

// GetNumClientsOk returns a tuple with the NumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWanResult) GetNumClientsOk() (*float32, bool) {
	if o == nil || IsNil(o.NumClients) {
		return nil, false
	}
	return o.NumClients, true
}

// HasNumClients returns a boolean if a field has been set.
func (o *OrgSiteSleWanResult) HasNumClients() bool {
	if o != nil && !IsNil(o.NumClients) {
		return true
	}

	return false
}

// SetNumClients gets a reference to the given float32 and assigns it to the NumClients field.
func (o *OrgSiteSleWanResult) SetNumClients(v float32) {
	o.NumClients = &v
}

// GetNumGateways returns the NumGateways field value if set, zero value otherwise.
func (o *OrgSiteSleWanResult) GetNumGateways() float32 {
	if o == nil || IsNil(o.NumGateways) {
		var ret float32
		return ret
	}
	return *o.NumGateways
}

// GetNumGatewaysOk returns a tuple with the NumGateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWanResult) GetNumGatewaysOk() (*float32, bool) {
	if o == nil || IsNil(o.NumGateways) {
		return nil, false
	}
	return o.NumGateways, true
}

// HasNumGateways returns a boolean if a field has been set.
func (o *OrgSiteSleWanResult) HasNumGateways() bool {
	if o != nil && !IsNil(o.NumGateways) {
		return true
	}

	return false
}

// SetNumGateways gets a reference to the given float32 and assigns it to the NumGateways field.
func (o *OrgSiteSleWanResult) SetNumGateways(v float32) {
	o.NumGateways = &v
}

// GetSiteId returns the SiteId field value
func (o *OrgSiteSleWanResult) GetSiteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWanResult) GetSiteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteId, true
}

// SetSiteId sets field value
func (o *OrgSiteSleWanResult) SetSiteId(v string) {
	o.SiteId = v
}

// GetWanLinkHealth returns the WanLinkHealth field value if set, zero value otherwise.
func (o *OrgSiteSleWanResult) GetWanLinkHealth() float32 {
	if o == nil || IsNil(o.WanLinkHealth) {
		var ret float32
		return ret
	}
	return *o.WanLinkHealth
}

// GetWanLinkHealthOk returns a tuple with the WanLinkHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSiteSleWanResult) GetWanLinkHealthOk() (*float32, bool) {
	if o == nil || IsNil(o.WanLinkHealth) {
		return nil, false
	}
	return o.WanLinkHealth, true
}

// HasWanLinkHealth returns a boolean if a field has been set.
func (o *OrgSiteSleWanResult) HasWanLinkHealth() bool {
	if o != nil && !IsNil(o.WanLinkHealth) {
		return true
	}

	return false
}

// SetWanLinkHealth gets a reference to the given float32 and assigns it to the WanLinkHealth field.
func (o *OrgSiteSleWanResult) SetWanLinkHealth(v float32) {
	o.WanLinkHealth = &v
}

func (o OrgSiteSleWanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSiteSleWanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationHealth) {
		toSerialize["application_health"] = o.ApplicationHealth
	}
	if !IsNil(o.GatewayHealth) {
		toSerialize["gateway-health"] = o.GatewayHealth
	}
	if !IsNil(o.NumClients) {
		toSerialize["num_clients"] = o.NumClients
	}
	if !IsNil(o.NumGateways) {
		toSerialize["num_gateways"] = o.NumGateways
	}
	toSerialize["site_id"] = o.SiteId
	if !IsNil(o.WanLinkHealth) {
		toSerialize["wan-link-health"] = o.WanLinkHealth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSiteSleWanResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"site_id",
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

	varOrgSiteSleWanResult := _OrgSiteSleWanResult{}

	err = json.Unmarshal(data, &varOrgSiteSleWanResult)

	if err != nil {
		return err
	}

	*o = OrgSiteSleWanResult(varOrgSiteSleWanResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "application_health")
		delete(additionalProperties, "gateway-health")
		delete(additionalProperties, "num_clients")
		delete(additionalProperties, "num_gateways")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "wan-link-health")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSiteSleWanResult struct {
	value *OrgSiteSleWanResult
	isSet bool
}

func (v NullableOrgSiteSleWanResult) Get() *OrgSiteSleWanResult {
	return v.value
}

func (v *NullableOrgSiteSleWanResult) Set(val *OrgSiteSleWanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSiteSleWanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSiteSleWanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSiteSleWanResult(val *OrgSiteSleWanResult) *NullableOrgSiteSleWanResult {
	return &NullableOrgSiteSleWanResult{value: val, isSet: true}
}

func (v NullableOrgSiteSleWanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSiteSleWanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


