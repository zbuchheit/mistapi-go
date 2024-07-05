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

// SiteSleImpactedGatewaysScopeParameters the model 'SiteSleImpactedGatewaysScopeParameters'
type SiteSleImpactedGatewaysScopeParameters string

// List of site_sle_impacted_gateways_scope_parameters
const (
	SITESLEIMPACTEDGATEWAYSSCOPEPARAMETERS_EMPTY SiteSleImpactedGatewaysScopeParameters = ""
	SITESLEIMPACTEDGATEWAYSSCOPEPARAMETERS_SITE SiteSleImpactedGatewaysScopeParameters = "site"
)

// All allowed values of SiteSleImpactedGatewaysScopeParameters enum
var AllowedSiteSleImpactedGatewaysScopeParametersEnumValues = []SiteSleImpactedGatewaysScopeParameters{
	"",
	"site",
}

func (v *SiteSleImpactedGatewaysScopeParameters) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiteSleImpactedGatewaysScopeParameters(value)
	for _, existing := range AllowedSiteSleImpactedGatewaysScopeParametersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiteSleImpactedGatewaysScopeParameters", value)
}

// NewSiteSleImpactedGatewaysScopeParametersFromValue returns a pointer to a valid SiteSleImpactedGatewaysScopeParameters
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiteSleImpactedGatewaysScopeParametersFromValue(v string) (*SiteSleImpactedGatewaysScopeParameters, error) {
	ev := SiteSleImpactedGatewaysScopeParameters(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiteSleImpactedGatewaysScopeParameters: valid values are %v", v, AllowedSiteSleImpactedGatewaysScopeParametersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiteSleImpactedGatewaysScopeParameters) IsValid() bool {
	for _, existing := range AllowedSiteSleImpactedGatewaysScopeParametersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to site_sle_impacted_gateways_scope_parameters value
func (v SiteSleImpactedGatewaysScopeParameters) Ptr() *SiteSleImpactedGatewaysScopeParameters {
	return &v
}

type NullableSiteSleImpactedGatewaysScopeParameters struct {
	value *SiteSleImpactedGatewaysScopeParameters
	isSet bool
}

func (v NullableSiteSleImpactedGatewaysScopeParameters) Get() *SiteSleImpactedGatewaysScopeParameters {
	return v.value
}

func (v *NullableSiteSleImpactedGatewaysScopeParameters) Set(val *SiteSleImpactedGatewaysScopeParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteSleImpactedGatewaysScopeParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteSleImpactedGatewaysScopeParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteSleImpactedGatewaysScopeParameters(val *SiteSleImpactedGatewaysScopeParameters) *NullableSiteSleImpactedGatewaysScopeParameters {
	return &NullableSiteSleImpactedGatewaysScopeParameters{value: val, isSet: true}
}

func (v NullableSiteSleImpactedGatewaysScopeParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteSleImpactedGatewaysScopeParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

