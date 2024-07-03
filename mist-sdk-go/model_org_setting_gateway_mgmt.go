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

// checks if the OrgSettingGatewayMgmt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingGatewayMgmt{}

// OrgSettingGatewayMgmt struct for OrgSettingGatewayMgmt
type OrgSettingGatewayMgmt struct {
	AppProbing *OrgSettingGatewayMgmtAppProbing `json:"app_probing,omitempty"`
	AppUsage *bool `json:"app_usage,omitempty"`
	HostOutPolicies *OrgSettingGatewayMgmtHostOutPolicies `json:"host_out_policies,omitempty"`
	OverlayIp *OrgSettingGatewayMgmtOverlayIp `json:"overlay_ip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingGatewayMgmt OrgSettingGatewayMgmt

// NewOrgSettingGatewayMgmt instantiates a new OrgSettingGatewayMgmt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingGatewayMgmt() *OrgSettingGatewayMgmt {
	this := OrgSettingGatewayMgmt{}
	return &this
}

// NewOrgSettingGatewayMgmtWithDefaults instantiates a new OrgSettingGatewayMgmt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingGatewayMgmtWithDefaults() *OrgSettingGatewayMgmt {
	this := OrgSettingGatewayMgmt{}
	return &this
}

// GetAppProbing returns the AppProbing field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmt) GetAppProbing() OrgSettingGatewayMgmtAppProbing {
	if o == nil || IsNil(o.AppProbing) {
		var ret OrgSettingGatewayMgmtAppProbing
		return ret
	}
	return *o.AppProbing
}

// GetAppProbingOk returns a tuple with the AppProbing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmt) GetAppProbingOk() (*OrgSettingGatewayMgmtAppProbing, bool) {
	if o == nil || IsNil(o.AppProbing) {
		return nil, false
	}
	return o.AppProbing, true
}

// HasAppProbing returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmt) HasAppProbing() bool {
	if o != nil && !IsNil(o.AppProbing) {
		return true
	}

	return false
}

// SetAppProbing gets a reference to the given OrgSettingGatewayMgmtAppProbing and assigns it to the AppProbing field.
func (o *OrgSettingGatewayMgmt) SetAppProbing(v OrgSettingGatewayMgmtAppProbing) {
	o.AppProbing = &v
}

// GetAppUsage returns the AppUsage field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmt) GetAppUsage() bool {
	if o == nil || IsNil(o.AppUsage) {
		var ret bool
		return ret
	}
	return *o.AppUsage
}

// GetAppUsageOk returns a tuple with the AppUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmt) GetAppUsageOk() (*bool, bool) {
	if o == nil || IsNil(o.AppUsage) {
		return nil, false
	}
	return o.AppUsage, true
}

// HasAppUsage returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmt) HasAppUsage() bool {
	if o != nil && !IsNil(o.AppUsage) {
		return true
	}

	return false
}

// SetAppUsage gets a reference to the given bool and assigns it to the AppUsage field.
func (o *OrgSettingGatewayMgmt) SetAppUsage(v bool) {
	o.AppUsage = &v
}

// GetHostOutPolicies returns the HostOutPolicies field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmt) GetHostOutPolicies() OrgSettingGatewayMgmtHostOutPolicies {
	if o == nil || IsNil(o.HostOutPolicies) {
		var ret OrgSettingGatewayMgmtHostOutPolicies
		return ret
	}
	return *o.HostOutPolicies
}

// GetHostOutPoliciesOk returns a tuple with the HostOutPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmt) GetHostOutPoliciesOk() (*OrgSettingGatewayMgmtHostOutPolicies, bool) {
	if o == nil || IsNil(o.HostOutPolicies) {
		return nil, false
	}
	return o.HostOutPolicies, true
}

// HasHostOutPolicies returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmt) HasHostOutPolicies() bool {
	if o != nil && !IsNil(o.HostOutPolicies) {
		return true
	}

	return false
}

// SetHostOutPolicies gets a reference to the given OrgSettingGatewayMgmtHostOutPolicies and assigns it to the HostOutPolicies field.
func (o *OrgSettingGatewayMgmt) SetHostOutPolicies(v OrgSettingGatewayMgmtHostOutPolicies) {
	o.HostOutPolicies = &v
}

// GetOverlayIp returns the OverlayIp field value if set, zero value otherwise.
func (o *OrgSettingGatewayMgmt) GetOverlayIp() OrgSettingGatewayMgmtOverlayIp {
	if o == nil || IsNil(o.OverlayIp) {
		var ret OrgSettingGatewayMgmtOverlayIp
		return ret
	}
	return *o.OverlayIp
}

// GetOverlayIpOk returns a tuple with the OverlayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingGatewayMgmt) GetOverlayIpOk() (*OrgSettingGatewayMgmtOverlayIp, bool) {
	if o == nil || IsNil(o.OverlayIp) {
		return nil, false
	}
	return o.OverlayIp, true
}

// HasOverlayIp returns a boolean if a field has been set.
func (o *OrgSettingGatewayMgmt) HasOverlayIp() bool {
	if o != nil && !IsNil(o.OverlayIp) {
		return true
	}

	return false
}

// SetOverlayIp gets a reference to the given OrgSettingGatewayMgmtOverlayIp and assigns it to the OverlayIp field.
func (o *OrgSettingGatewayMgmt) SetOverlayIp(v OrgSettingGatewayMgmtOverlayIp) {
	o.OverlayIp = &v
}

func (o OrgSettingGatewayMgmt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingGatewayMgmt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppProbing) {
		toSerialize["app_probing"] = o.AppProbing
	}
	if !IsNil(o.AppUsage) {
		toSerialize["app_usage"] = o.AppUsage
	}
	if !IsNil(o.HostOutPolicies) {
		toSerialize["host_out_policies"] = o.HostOutPolicies
	}
	if !IsNil(o.OverlayIp) {
		toSerialize["overlay_ip"] = o.OverlayIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingGatewayMgmt) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingGatewayMgmt := _OrgSettingGatewayMgmt{}

	err = json.Unmarshal(data, &varOrgSettingGatewayMgmt)

	if err != nil {
		return err
	}

	*o = OrgSettingGatewayMgmt(varOrgSettingGatewayMgmt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app_probing")
		delete(additionalProperties, "app_usage")
		delete(additionalProperties, "host_out_policies")
		delete(additionalProperties, "overlay_ip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingGatewayMgmt struct {
	value *OrgSettingGatewayMgmt
	isSet bool
}

func (v NullableOrgSettingGatewayMgmt) Get() *OrgSettingGatewayMgmt {
	return v.value
}

func (v *NullableOrgSettingGatewayMgmt) Set(val *OrgSettingGatewayMgmt) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingGatewayMgmt) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingGatewayMgmt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingGatewayMgmt(val *OrgSettingGatewayMgmt) *NullableOrgSettingGatewayMgmt {
	return &NullableOrgSettingGatewayMgmt{value: val, isSet: true}
}

func (v NullableOrgSettingGatewayMgmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingGatewayMgmt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


