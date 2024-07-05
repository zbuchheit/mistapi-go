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

// checks if the ServicePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePolicy{}

// ServicePolicy struct for ServicePolicy
type ServicePolicy struct {
	Action *AllowDeny `json:"action,omitempty"`
	Appqoe *ServicePolicyAppqoe `json:"appqoe,omitempty"`
	Ewf []ServicePolicyEwfRule `json:"ewf,omitempty"`
	Idp *IdpConfig `json:"idp,omitempty"`
	// access within the same VRF
	LocalRouting *bool `json:"local_routing,omitempty"`
	Name *string `json:"name,omitempty"`
	// by default, we derive all paths available and use them optionally, you can customize by using `path_preference`
	PathPreferences *string `json:"path_preferences,omitempty"`
	// used to link servicepolicy defined at org level and overwrite some attributes
	ServicepolicyId *string `json:"servicepolicy_id,omitempty"`
	Services []string `json:"services,omitempty"`
	Tenants []string `json:"tenants,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicePolicy ServicePolicy

// NewServicePolicy instantiates a new ServicePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePolicy() *ServicePolicy {
	this := ServicePolicy{}
	var action AllowDeny = ALLOWDENY_ALLOW
	this.Action = &action
	var localRouting bool = false
	this.LocalRouting = &localRouting
	return &this
}

// NewServicePolicyWithDefaults instantiates a new ServicePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePolicyWithDefaults() *ServicePolicy {
	this := ServicePolicy{}
	var action AllowDeny = ALLOWDENY_ALLOW
	this.Action = &action
	var localRouting bool = false
	this.LocalRouting = &localRouting
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ServicePolicy) GetAction() AllowDeny {
	if o == nil || IsNil(o.Action) {
		var ret AllowDeny
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetActionOk() (*AllowDeny, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ServicePolicy) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given AllowDeny and assigns it to the Action field.
func (o *ServicePolicy) SetAction(v AllowDeny) {
	o.Action = &v
}

// GetAppqoe returns the Appqoe field value if set, zero value otherwise.
func (o *ServicePolicy) GetAppqoe() ServicePolicyAppqoe {
	if o == nil || IsNil(o.Appqoe) {
		var ret ServicePolicyAppqoe
		return ret
	}
	return *o.Appqoe
}

// GetAppqoeOk returns a tuple with the Appqoe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetAppqoeOk() (*ServicePolicyAppqoe, bool) {
	if o == nil || IsNil(o.Appqoe) {
		return nil, false
	}
	return o.Appqoe, true
}

// HasAppqoe returns a boolean if a field has been set.
func (o *ServicePolicy) HasAppqoe() bool {
	if o != nil && !IsNil(o.Appqoe) {
		return true
	}

	return false
}

// SetAppqoe gets a reference to the given ServicePolicyAppqoe and assigns it to the Appqoe field.
func (o *ServicePolicy) SetAppqoe(v ServicePolicyAppqoe) {
	o.Appqoe = &v
}

// GetEwf returns the Ewf field value if set, zero value otherwise.
func (o *ServicePolicy) GetEwf() []ServicePolicyEwfRule {
	if o == nil || IsNil(o.Ewf) {
		var ret []ServicePolicyEwfRule
		return ret
	}
	return o.Ewf
}

// GetEwfOk returns a tuple with the Ewf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetEwfOk() ([]ServicePolicyEwfRule, bool) {
	if o == nil || IsNil(o.Ewf) {
		return nil, false
	}
	return o.Ewf, true
}

// HasEwf returns a boolean if a field has been set.
func (o *ServicePolicy) HasEwf() bool {
	if o != nil && !IsNil(o.Ewf) {
		return true
	}

	return false
}

// SetEwf gets a reference to the given []ServicePolicyEwfRule and assigns it to the Ewf field.
func (o *ServicePolicy) SetEwf(v []ServicePolicyEwfRule) {
	o.Ewf = v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *ServicePolicy) GetIdp() IdpConfig {
	if o == nil || IsNil(o.Idp) {
		var ret IdpConfig
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetIdpOk() (*IdpConfig, bool) {
	if o == nil || IsNil(o.Idp) {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *ServicePolicy) HasIdp() bool {
	if o != nil && !IsNil(o.Idp) {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IdpConfig and assigns it to the Idp field.
func (o *ServicePolicy) SetIdp(v IdpConfig) {
	o.Idp = &v
}

// GetLocalRouting returns the LocalRouting field value if set, zero value otherwise.
func (o *ServicePolicy) GetLocalRouting() bool {
	if o == nil || IsNil(o.LocalRouting) {
		var ret bool
		return ret
	}
	return *o.LocalRouting
}

// GetLocalRoutingOk returns a tuple with the LocalRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetLocalRoutingOk() (*bool, bool) {
	if o == nil || IsNil(o.LocalRouting) {
		return nil, false
	}
	return o.LocalRouting, true
}

// HasLocalRouting returns a boolean if a field has been set.
func (o *ServicePolicy) HasLocalRouting() bool {
	if o != nil && !IsNil(o.LocalRouting) {
		return true
	}

	return false
}

// SetLocalRouting gets a reference to the given bool and assigns it to the LocalRouting field.
func (o *ServicePolicy) SetLocalRouting(v bool) {
	o.LocalRouting = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServicePolicy) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServicePolicy) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServicePolicy) SetName(v string) {
	o.Name = &v
}

// GetPathPreferences returns the PathPreferences field value if set, zero value otherwise.
func (o *ServicePolicy) GetPathPreferences() string {
	if o == nil || IsNil(o.PathPreferences) {
		var ret string
		return ret
	}
	return *o.PathPreferences
}

// GetPathPreferencesOk returns a tuple with the PathPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetPathPreferencesOk() (*string, bool) {
	if o == nil || IsNil(o.PathPreferences) {
		return nil, false
	}
	return o.PathPreferences, true
}

// HasPathPreferences returns a boolean if a field has been set.
func (o *ServicePolicy) HasPathPreferences() bool {
	if o != nil && !IsNil(o.PathPreferences) {
		return true
	}

	return false
}

// SetPathPreferences gets a reference to the given string and assigns it to the PathPreferences field.
func (o *ServicePolicy) SetPathPreferences(v string) {
	o.PathPreferences = &v
}

// GetServicepolicyId returns the ServicepolicyId field value if set, zero value otherwise.
func (o *ServicePolicy) GetServicepolicyId() string {
	if o == nil || IsNil(o.ServicepolicyId) {
		var ret string
		return ret
	}
	return *o.ServicepolicyId
}

// GetServicepolicyIdOk returns a tuple with the ServicepolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetServicepolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServicepolicyId) {
		return nil, false
	}
	return o.ServicepolicyId, true
}

// HasServicepolicyId returns a boolean if a field has been set.
func (o *ServicePolicy) HasServicepolicyId() bool {
	if o != nil && !IsNil(o.ServicepolicyId) {
		return true
	}

	return false
}

// SetServicepolicyId gets a reference to the given string and assigns it to the ServicepolicyId field.
func (o *ServicePolicy) SetServicepolicyId(v string) {
	o.ServicepolicyId = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ServicePolicy) GetServices() []string {
	if o == nil || IsNil(o.Services) {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ServicePolicy) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *ServicePolicy) SetServices(v []string) {
	o.Services = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ServicePolicy) GetTenants() []string {
	if o == nil || IsNil(o.Tenants) {
		var ret []string
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePolicy) GetTenantsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ServicePolicy) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []string and assigns it to the Tenants field.
func (o *ServicePolicy) SetTenants(v []string) {
	o.Tenants = v
}

func (o ServicePolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Appqoe) {
		toSerialize["appqoe"] = o.Appqoe
	}
	if !IsNil(o.Ewf) {
		toSerialize["ewf"] = o.Ewf
	}
	if !IsNil(o.Idp) {
		toSerialize["idp"] = o.Idp
	}
	if !IsNil(o.LocalRouting) {
		toSerialize["local_routing"] = o.LocalRouting
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PathPreferences) {
		toSerialize["path_preferences"] = o.PathPreferences
	}
	if !IsNil(o.ServicepolicyId) {
		toSerialize["servicepolicy_id"] = o.ServicepolicyId
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicePolicy) UnmarshalJSON(data []byte) (err error) {
	varServicePolicy := _ServicePolicy{}

	err = json.Unmarshal(data, &varServicePolicy)

	if err != nil {
		return err
	}

	*o = ServicePolicy(varServicePolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "appqoe")
		delete(additionalProperties, "ewf")
		delete(additionalProperties, "idp")
		delete(additionalProperties, "local_routing")
		delete(additionalProperties, "name")
		delete(additionalProperties, "path_preferences")
		delete(additionalProperties, "servicepolicy_id")
		delete(additionalProperties, "services")
		delete(additionalProperties, "tenants")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicePolicy struct {
	value *ServicePolicy
	isSet bool
}

func (v NullableServicePolicy) Get() *ServicePolicy {
	return v.value
}

func (v *NullableServicePolicy) Set(val *ServicePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePolicy(val *ServicePolicy) *NullableServicePolicy {
	return &NullableServicePolicy{value: val, isSet: true}
}

func (v NullableServicePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


