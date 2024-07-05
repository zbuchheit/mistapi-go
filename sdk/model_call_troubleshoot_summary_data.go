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

// checks if the CallTroubleshootSummaryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallTroubleshootSummaryData{}

// CallTroubleshootSummaryData struct for CallTroubleshootSummaryData
type CallTroubleshootSummaryData struct {
	ApNumClients *float32 `json:"ap_num_clients,omitempty"`
	ApRtt *float32 `json:"ap_rtt,omitempty"`
	ClientCpu *float32 `json:"client_cpu,omitempty"`
	ClientNStreams *float32 `json:"client_n_streams,omitempty"`
	ClientRadioBand *float32 `json:"client_radio_band,omitempty"`
	ClientRssi *float32 `json:"client_rssi,omitempty"`
	ClientRxBytes *float32 `json:"client_rx_bytes,omitempty"`
	ClientRxRates *float32 `json:"client_rx_rates,omitempty"`
	ClientTxBytes *float32 `json:"client_tx_bytes,omitempty"`
	ClientTxRates *float32 `json:"client_tx_rates,omitempty"`
	ClientTxRetries *float32 `json:"client_tx_retries,omitempty"`
	ClientVpnDistaince *float32 `json:"client_vpn_distaince,omitempty"`
	ClientWifiVersion *float32 `json:"client_wifi_version,omitempty"`
	Expected *float32 `json:"expected,omitempty"`
	RadioBandwidth *float32 `json:"radio_bandwidth,omitempty"`
	RadioChannel *float32 `json:"radio_channel,omitempty"`
	RadioTxPower *float32 `json:"radio_tx_power,omitempty"`
	RadioUtil *float32 `json:"radio_util,omitempty"`
	RadioUtilInterference *float32 `json:"radio_util_interference,omitempty"`
	SiteNumClients *float32 `json:"site_num_clients,omitempty"`
	WanAvgDownloadMbps *float32 `json:"wan_avg_download_mbps,omitempty"`
	WanAvgUploadMbps *float32 `json:"wan_avg_upload_mbps,omitempty"`
	WanJitter *float32 `json:"wan_jitter,omitempty"`
	WanMaxDownloadMbps *float32 `json:"wan_max_download_mbps,omitempty"`
	WanMaxUploadMbps *float32 `json:"wan_max_upload_mbps,omitempty"`
	WanRtt *float32 `json:"wan_rtt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallTroubleshootSummaryData CallTroubleshootSummaryData

// NewCallTroubleshootSummaryData instantiates a new CallTroubleshootSummaryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallTroubleshootSummaryData() *CallTroubleshootSummaryData {
	this := CallTroubleshootSummaryData{}
	return &this
}

// NewCallTroubleshootSummaryDataWithDefaults instantiates a new CallTroubleshootSummaryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallTroubleshootSummaryDataWithDefaults() *CallTroubleshootSummaryData {
	this := CallTroubleshootSummaryData{}
	return &this
}

// GetApNumClients returns the ApNumClients field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetApNumClients() float32 {
	if o == nil || IsNil(o.ApNumClients) {
		var ret float32
		return ret
	}
	return *o.ApNumClients
}

// GetApNumClientsOk returns a tuple with the ApNumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetApNumClientsOk() (*float32, bool) {
	if o == nil || IsNil(o.ApNumClients) {
		return nil, false
	}
	return o.ApNumClients, true
}

// HasApNumClients returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasApNumClients() bool {
	if o != nil && !IsNil(o.ApNumClients) {
		return true
	}

	return false
}

// SetApNumClients gets a reference to the given float32 and assigns it to the ApNumClients field.
func (o *CallTroubleshootSummaryData) SetApNumClients(v float32) {
	o.ApNumClients = &v
}

// GetApRtt returns the ApRtt field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetApRtt() float32 {
	if o == nil || IsNil(o.ApRtt) {
		var ret float32
		return ret
	}
	return *o.ApRtt
}

// GetApRttOk returns a tuple with the ApRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetApRttOk() (*float32, bool) {
	if o == nil || IsNil(o.ApRtt) {
		return nil, false
	}
	return o.ApRtt, true
}

// HasApRtt returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasApRtt() bool {
	if o != nil && !IsNil(o.ApRtt) {
		return true
	}

	return false
}

// SetApRtt gets a reference to the given float32 and assigns it to the ApRtt field.
func (o *CallTroubleshootSummaryData) SetApRtt(v float32) {
	o.ApRtt = &v
}

// GetClientCpu returns the ClientCpu field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientCpu() float32 {
	if o == nil || IsNil(o.ClientCpu) {
		var ret float32
		return ret
	}
	return *o.ClientCpu
}

// GetClientCpuOk returns a tuple with the ClientCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientCpuOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientCpu) {
		return nil, false
	}
	return o.ClientCpu, true
}

// HasClientCpu returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientCpu() bool {
	if o != nil && !IsNil(o.ClientCpu) {
		return true
	}

	return false
}

// SetClientCpu gets a reference to the given float32 and assigns it to the ClientCpu field.
func (o *CallTroubleshootSummaryData) SetClientCpu(v float32) {
	o.ClientCpu = &v
}

// GetClientNStreams returns the ClientNStreams field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientNStreams() float32 {
	if o == nil || IsNil(o.ClientNStreams) {
		var ret float32
		return ret
	}
	return *o.ClientNStreams
}

// GetClientNStreamsOk returns a tuple with the ClientNStreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientNStreamsOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientNStreams) {
		return nil, false
	}
	return o.ClientNStreams, true
}

// HasClientNStreams returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientNStreams() bool {
	if o != nil && !IsNil(o.ClientNStreams) {
		return true
	}

	return false
}

// SetClientNStreams gets a reference to the given float32 and assigns it to the ClientNStreams field.
func (o *CallTroubleshootSummaryData) SetClientNStreams(v float32) {
	o.ClientNStreams = &v
}

// GetClientRadioBand returns the ClientRadioBand field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientRadioBand() float32 {
	if o == nil || IsNil(o.ClientRadioBand) {
		var ret float32
		return ret
	}
	return *o.ClientRadioBand
}

// GetClientRadioBandOk returns a tuple with the ClientRadioBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientRadioBandOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientRadioBand) {
		return nil, false
	}
	return o.ClientRadioBand, true
}

// HasClientRadioBand returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientRadioBand() bool {
	if o != nil && !IsNil(o.ClientRadioBand) {
		return true
	}

	return false
}

// SetClientRadioBand gets a reference to the given float32 and assigns it to the ClientRadioBand field.
func (o *CallTroubleshootSummaryData) SetClientRadioBand(v float32) {
	o.ClientRadioBand = &v
}

// GetClientRssi returns the ClientRssi field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientRssi() float32 {
	if o == nil || IsNil(o.ClientRssi) {
		var ret float32
		return ret
	}
	return *o.ClientRssi
}

// GetClientRssiOk returns a tuple with the ClientRssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientRssiOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientRssi) {
		return nil, false
	}
	return o.ClientRssi, true
}

// HasClientRssi returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientRssi() bool {
	if o != nil && !IsNil(o.ClientRssi) {
		return true
	}

	return false
}

// SetClientRssi gets a reference to the given float32 and assigns it to the ClientRssi field.
func (o *CallTroubleshootSummaryData) SetClientRssi(v float32) {
	o.ClientRssi = &v
}

// GetClientRxBytes returns the ClientRxBytes field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientRxBytes() float32 {
	if o == nil || IsNil(o.ClientRxBytes) {
		var ret float32
		return ret
	}
	return *o.ClientRxBytes
}

// GetClientRxBytesOk returns a tuple with the ClientRxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientRxBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientRxBytes) {
		return nil, false
	}
	return o.ClientRxBytes, true
}

// HasClientRxBytes returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientRxBytes() bool {
	if o != nil && !IsNil(o.ClientRxBytes) {
		return true
	}

	return false
}

// SetClientRxBytes gets a reference to the given float32 and assigns it to the ClientRxBytes field.
func (o *CallTroubleshootSummaryData) SetClientRxBytes(v float32) {
	o.ClientRxBytes = &v
}

// GetClientRxRates returns the ClientRxRates field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientRxRates() float32 {
	if o == nil || IsNil(o.ClientRxRates) {
		var ret float32
		return ret
	}
	return *o.ClientRxRates
}

// GetClientRxRatesOk returns a tuple with the ClientRxRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientRxRatesOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientRxRates) {
		return nil, false
	}
	return o.ClientRxRates, true
}

// HasClientRxRates returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientRxRates() bool {
	if o != nil && !IsNil(o.ClientRxRates) {
		return true
	}

	return false
}

// SetClientRxRates gets a reference to the given float32 and assigns it to the ClientRxRates field.
func (o *CallTroubleshootSummaryData) SetClientRxRates(v float32) {
	o.ClientRxRates = &v
}

// GetClientTxBytes returns the ClientTxBytes field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientTxBytes() float32 {
	if o == nil || IsNil(o.ClientTxBytes) {
		var ret float32
		return ret
	}
	return *o.ClientTxBytes
}

// GetClientTxBytesOk returns a tuple with the ClientTxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientTxBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientTxBytes) {
		return nil, false
	}
	return o.ClientTxBytes, true
}

// HasClientTxBytes returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientTxBytes() bool {
	if o != nil && !IsNil(o.ClientTxBytes) {
		return true
	}

	return false
}

// SetClientTxBytes gets a reference to the given float32 and assigns it to the ClientTxBytes field.
func (o *CallTroubleshootSummaryData) SetClientTxBytes(v float32) {
	o.ClientTxBytes = &v
}

// GetClientTxRates returns the ClientTxRates field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientTxRates() float32 {
	if o == nil || IsNil(o.ClientTxRates) {
		var ret float32
		return ret
	}
	return *o.ClientTxRates
}

// GetClientTxRatesOk returns a tuple with the ClientTxRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientTxRatesOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientTxRates) {
		return nil, false
	}
	return o.ClientTxRates, true
}

// HasClientTxRates returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientTxRates() bool {
	if o != nil && !IsNil(o.ClientTxRates) {
		return true
	}

	return false
}

// SetClientTxRates gets a reference to the given float32 and assigns it to the ClientTxRates field.
func (o *CallTroubleshootSummaryData) SetClientTxRates(v float32) {
	o.ClientTxRates = &v
}

// GetClientTxRetries returns the ClientTxRetries field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientTxRetries() float32 {
	if o == nil || IsNil(o.ClientTxRetries) {
		var ret float32
		return ret
	}
	return *o.ClientTxRetries
}

// GetClientTxRetriesOk returns a tuple with the ClientTxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientTxRetriesOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientTxRetries) {
		return nil, false
	}
	return o.ClientTxRetries, true
}

// HasClientTxRetries returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientTxRetries() bool {
	if o != nil && !IsNil(o.ClientTxRetries) {
		return true
	}

	return false
}

// SetClientTxRetries gets a reference to the given float32 and assigns it to the ClientTxRetries field.
func (o *CallTroubleshootSummaryData) SetClientTxRetries(v float32) {
	o.ClientTxRetries = &v
}

// GetClientVpnDistaince returns the ClientVpnDistaince field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientVpnDistaince() float32 {
	if o == nil || IsNil(o.ClientVpnDistaince) {
		var ret float32
		return ret
	}
	return *o.ClientVpnDistaince
}

// GetClientVpnDistainceOk returns a tuple with the ClientVpnDistaince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientVpnDistainceOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientVpnDistaince) {
		return nil, false
	}
	return o.ClientVpnDistaince, true
}

// HasClientVpnDistaince returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientVpnDistaince() bool {
	if o != nil && !IsNil(o.ClientVpnDistaince) {
		return true
	}

	return false
}

// SetClientVpnDistaince gets a reference to the given float32 and assigns it to the ClientVpnDistaince field.
func (o *CallTroubleshootSummaryData) SetClientVpnDistaince(v float32) {
	o.ClientVpnDistaince = &v
}

// GetClientWifiVersion returns the ClientWifiVersion field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetClientWifiVersion() float32 {
	if o == nil || IsNil(o.ClientWifiVersion) {
		var ret float32
		return ret
	}
	return *o.ClientWifiVersion
}

// GetClientWifiVersionOk returns a tuple with the ClientWifiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetClientWifiVersionOk() (*float32, bool) {
	if o == nil || IsNil(o.ClientWifiVersion) {
		return nil, false
	}
	return o.ClientWifiVersion, true
}

// HasClientWifiVersion returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasClientWifiVersion() bool {
	if o != nil && !IsNil(o.ClientWifiVersion) {
		return true
	}

	return false
}

// SetClientWifiVersion gets a reference to the given float32 and assigns it to the ClientWifiVersion field.
func (o *CallTroubleshootSummaryData) SetClientWifiVersion(v float32) {
	o.ClientWifiVersion = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetExpected() float32 {
	if o == nil || IsNil(o.Expected) {
		var ret float32
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetExpectedOk() (*float32, bool) {
	if o == nil || IsNil(o.Expected) {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasExpected() bool {
	if o != nil && !IsNil(o.Expected) {
		return true
	}

	return false
}

// SetExpected gets a reference to the given float32 and assigns it to the Expected field.
func (o *CallTroubleshootSummaryData) SetExpected(v float32) {
	o.Expected = &v
}

// GetRadioBandwidth returns the RadioBandwidth field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetRadioBandwidth() float32 {
	if o == nil || IsNil(o.RadioBandwidth) {
		var ret float32
		return ret
	}
	return *o.RadioBandwidth
}

// GetRadioBandwidthOk returns a tuple with the RadioBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetRadioBandwidthOk() (*float32, bool) {
	if o == nil || IsNil(o.RadioBandwidth) {
		return nil, false
	}
	return o.RadioBandwidth, true
}

// HasRadioBandwidth returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasRadioBandwidth() bool {
	if o != nil && !IsNil(o.RadioBandwidth) {
		return true
	}

	return false
}

// SetRadioBandwidth gets a reference to the given float32 and assigns it to the RadioBandwidth field.
func (o *CallTroubleshootSummaryData) SetRadioBandwidth(v float32) {
	o.RadioBandwidth = &v
}

// GetRadioChannel returns the RadioChannel field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetRadioChannel() float32 {
	if o == nil || IsNil(o.RadioChannel) {
		var ret float32
		return ret
	}
	return *o.RadioChannel
}

// GetRadioChannelOk returns a tuple with the RadioChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetRadioChannelOk() (*float32, bool) {
	if o == nil || IsNil(o.RadioChannel) {
		return nil, false
	}
	return o.RadioChannel, true
}

// HasRadioChannel returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasRadioChannel() bool {
	if o != nil && !IsNil(o.RadioChannel) {
		return true
	}

	return false
}

// SetRadioChannel gets a reference to the given float32 and assigns it to the RadioChannel field.
func (o *CallTroubleshootSummaryData) SetRadioChannel(v float32) {
	o.RadioChannel = &v
}

// GetRadioTxPower returns the RadioTxPower field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetRadioTxPower() float32 {
	if o == nil || IsNil(o.RadioTxPower) {
		var ret float32
		return ret
	}
	return *o.RadioTxPower
}

// GetRadioTxPowerOk returns a tuple with the RadioTxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetRadioTxPowerOk() (*float32, bool) {
	if o == nil || IsNil(o.RadioTxPower) {
		return nil, false
	}
	return o.RadioTxPower, true
}

// HasRadioTxPower returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasRadioTxPower() bool {
	if o != nil && !IsNil(o.RadioTxPower) {
		return true
	}

	return false
}

// SetRadioTxPower gets a reference to the given float32 and assigns it to the RadioTxPower field.
func (o *CallTroubleshootSummaryData) SetRadioTxPower(v float32) {
	o.RadioTxPower = &v
}

// GetRadioUtil returns the RadioUtil field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetRadioUtil() float32 {
	if o == nil || IsNil(o.RadioUtil) {
		var ret float32
		return ret
	}
	return *o.RadioUtil
}

// GetRadioUtilOk returns a tuple with the RadioUtil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetRadioUtilOk() (*float32, bool) {
	if o == nil || IsNil(o.RadioUtil) {
		return nil, false
	}
	return o.RadioUtil, true
}

// HasRadioUtil returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasRadioUtil() bool {
	if o != nil && !IsNil(o.RadioUtil) {
		return true
	}

	return false
}

// SetRadioUtil gets a reference to the given float32 and assigns it to the RadioUtil field.
func (o *CallTroubleshootSummaryData) SetRadioUtil(v float32) {
	o.RadioUtil = &v
}

// GetRadioUtilInterference returns the RadioUtilInterference field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetRadioUtilInterference() float32 {
	if o == nil || IsNil(o.RadioUtilInterference) {
		var ret float32
		return ret
	}
	return *o.RadioUtilInterference
}

// GetRadioUtilInterferenceOk returns a tuple with the RadioUtilInterference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetRadioUtilInterferenceOk() (*float32, bool) {
	if o == nil || IsNil(o.RadioUtilInterference) {
		return nil, false
	}
	return o.RadioUtilInterference, true
}

// HasRadioUtilInterference returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasRadioUtilInterference() bool {
	if o != nil && !IsNil(o.RadioUtilInterference) {
		return true
	}

	return false
}

// SetRadioUtilInterference gets a reference to the given float32 and assigns it to the RadioUtilInterference field.
func (o *CallTroubleshootSummaryData) SetRadioUtilInterference(v float32) {
	o.RadioUtilInterference = &v
}

// GetSiteNumClients returns the SiteNumClients field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetSiteNumClients() float32 {
	if o == nil || IsNil(o.SiteNumClients) {
		var ret float32
		return ret
	}
	return *o.SiteNumClients
}

// GetSiteNumClientsOk returns a tuple with the SiteNumClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetSiteNumClientsOk() (*float32, bool) {
	if o == nil || IsNil(o.SiteNumClients) {
		return nil, false
	}
	return o.SiteNumClients, true
}

// HasSiteNumClients returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasSiteNumClients() bool {
	if o != nil && !IsNil(o.SiteNumClients) {
		return true
	}

	return false
}

// SetSiteNumClients gets a reference to the given float32 and assigns it to the SiteNumClients field.
func (o *CallTroubleshootSummaryData) SetSiteNumClients(v float32) {
	o.SiteNumClients = &v
}

// GetWanAvgDownloadMbps returns the WanAvgDownloadMbps field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetWanAvgDownloadMbps() float32 {
	if o == nil || IsNil(o.WanAvgDownloadMbps) {
		var ret float32
		return ret
	}
	return *o.WanAvgDownloadMbps
}

// GetWanAvgDownloadMbpsOk returns a tuple with the WanAvgDownloadMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetWanAvgDownloadMbpsOk() (*float32, bool) {
	if o == nil || IsNil(o.WanAvgDownloadMbps) {
		return nil, false
	}
	return o.WanAvgDownloadMbps, true
}

// HasWanAvgDownloadMbps returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasWanAvgDownloadMbps() bool {
	if o != nil && !IsNil(o.WanAvgDownloadMbps) {
		return true
	}

	return false
}

// SetWanAvgDownloadMbps gets a reference to the given float32 and assigns it to the WanAvgDownloadMbps field.
func (o *CallTroubleshootSummaryData) SetWanAvgDownloadMbps(v float32) {
	o.WanAvgDownloadMbps = &v
}

// GetWanAvgUploadMbps returns the WanAvgUploadMbps field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetWanAvgUploadMbps() float32 {
	if o == nil || IsNil(o.WanAvgUploadMbps) {
		var ret float32
		return ret
	}
	return *o.WanAvgUploadMbps
}

// GetWanAvgUploadMbpsOk returns a tuple with the WanAvgUploadMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetWanAvgUploadMbpsOk() (*float32, bool) {
	if o == nil || IsNil(o.WanAvgUploadMbps) {
		return nil, false
	}
	return o.WanAvgUploadMbps, true
}

// HasWanAvgUploadMbps returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasWanAvgUploadMbps() bool {
	if o != nil && !IsNil(o.WanAvgUploadMbps) {
		return true
	}

	return false
}

// SetWanAvgUploadMbps gets a reference to the given float32 and assigns it to the WanAvgUploadMbps field.
func (o *CallTroubleshootSummaryData) SetWanAvgUploadMbps(v float32) {
	o.WanAvgUploadMbps = &v
}

// GetWanJitter returns the WanJitter field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetWanJitter() float32 {
	if o == nil || IsNil(o.WanJitter) {
		var ret float32
		return ret
	}
	return *o.WanJitter
}

// GetWanJitterOk returns a tuple with the WanJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetWanJitterOk() (*float32, bool) {
	if o == nil || IsNil(o.WanJitter) {
		return nil, false
	}
	return o.WanJitter, true
}

// HasWanJitter returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasWanJitter() bool {
	if o != nil && !IsNil(o.WanJitter) {
		return true
	}

	return false
}

// SetWanJitter gets a reference to the given float32 and assigns it to the WanJitter field.
func (o *CallTroubleshootSummaryData) SetWanJitter(v float32) {
	o.WanJitter = &v
}

// GetWanMaxDownloadMbps returns the WanMaxDownloadMbps field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetWanMaxDownloadMbps() float32 {
	if o == nil || IsNil(o.WanMaxDownloadMbps) {
		var ret float32
		return ret
	}
	return *o.WanMaxDownloadMbps
}

// GetWanMaxDownloadMbpsOk returns a tuple with the WanMaxDownloadMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetWanMaxDownloadMbpsOk() (*float32, bool) {
	if o == nil || IsNil(o.WanMaxDownloadMbps) {
		return nil, false
	}
	return o.WanMaxDownloadMbps, true
}

// HasWanMaxDownloadMbps returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasWanMaxDownloadMbps() bool {
	if o != nil && !IsNil(o.WanMaxDownloadMbps) {
		return true
	}

	return false
}

// SetWanMaxDownloadMbps gets a reference to the given float32 and assigns it to the WanMaxDownloadMbps field.
func (o *CallTroubleshootSummaryData) SetWanMaxDownloadMbps(v float32) {
	o.WanMaxDownloadMbps = &v
}

// GetWanMaxUploadMbps returns the WanMaxUploadMbps field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetWanMaxUploadMbps() float32 {
	if o == nil || IsNil(o.WanMaxUploadMbps) {
		var ret float32
		return ret
	}
	return *o.WanMaxUploadMbps
}

// GetWanMaxUploadMbpsOk returns a tuple with the WanMaxUploadMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetWanMaxUploadMbpsOk() (*float32, bool) {
	if o == nil || IsNil(o.WanMaxUploadMbps) {
		return nil, false
	}
	return o.WanMaxUploadMbps, true
}

// HasWanMaxUploadMbps returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasWanMaxUploadMbps() bool {
	if o != nil && !IsNil(o.WanMaxUploadMbps) {
		return true
	}

	return false
}

// SetWanMaxUploadMbps gets a reference to the given float32 and assigns it to the WanMaxUploadMbps field.
func (o *CallTroubleshootSummaryData) SetWanMaxUploadMbps(v float32) {
	o.WanMaxUploadMbps = &v
}

// GetWanRtt returns the WanRtt field value if set, zero value otherwise.
func (o *CallTroubleshootSummaryData) GetWanRtt() float32 {
	if o == nil || IsNil(o.WanRtt) {
		var ret float32
		return ret
	}
	return *o.WanRtt
}

// GetWanRttOk returns a tuple with the WanRtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallTroubleshootSummaryData) GetWanRttOk() (*float32, bool) {
	if o == nil || IsNil(o.WanRtt) {
		return nil, false
	}
	return o.WanRtt, true
}

// HasWanRtt returns a boolean if a field has been set.
func (o *CallTroubleshootSummaryData) HasWanRtt() bool {
	if o != nil && !IsNil(o.WanRtt) {
		return true
	}

	return false
}

// SetWanRtt gets a reference to the given float32 and assigns it to the WanRtt field.
func (o *CallTroubleshootSummaryData) SetWanRtt(v float32) {
	o.WanRtt = &v
}

func (o CallTroubleshootSummaryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallTroubleshootSummaryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApNumClients) {
		toSerialize["ap_num_clients"] = o.ApNumClients
	}
	if !IsNil(o.ApRtt) {
		toSerialize["ap_rtt"] = o.ApRtt
	}
	if !IsNil(o.ClientCpu) {
		toSerialize["client_cpu"] = o.ClientCpu
	}
	if !IsNil(o.ClientNStreams) {
		toSerialize["client_n_streams"] = o.ClientNStreams
	}
	if !IsNil(o.ClientRadioBand) {
		toSerialize["client_radio_band"] = o.ClientRadioBand
	}
	if !IsNil(o.ClientRssi) {
		toSerialize["client_rssi"] = o.ClientRssi
	}
	if !IsNil(o.ClientRxBytes) {
		toSerialize["client_rx_bytes"] = o.ClientRxBytes
	}
	if !IsNil(o.ClientRxRates) {
		toSerialize["client_rx_rates"] = o.ClientRxRates
	}
	if !IsNil(o.ClientTxBytes) {
		toSerialize["client_tx_bytes"] = o.ClientTxBytes
	}
	if !IsNil(o.ClientTxRates) {
		toSerialize["client_tx_rates"] = o.ClientTxRates
	}
	if !IsNil(o.ClientTxRetries) {
		toSerialize["client_tx_retries"] = o.ClientTxRetries
	}
	if !IsNil(o.ClientVpnDistaince) {
		toSerialize["client_vpn_distaince"] = o.ClientVpnDistaince
	}
	if !IsNil(o.ClientWifiVersion) {
		toSerialize["client_wifi_version"] = o.ClientWifiVersion
	}
	if !IsNil(o.Expected) {
		toSerialize["expected"] = o.Expected
	}
	if !IsNil(o.RadioBandwidth) {
		toSerialize["radio_bandwidth"] = o.RadioBandwidth
	}
	if !IsNil(o.RadioChannel) {
		toSerialize["radio_channel"] = o.RadioChannel
	}
	if !IsNil(o.RadioTxPower) {
		toSerialize["radio_tx_power"] = o.RadioTxPower
	}
	if !IsNil(o.RadioUtil) {
		toSerialize["radio_util"] = o.RadioUtil
	}
	if !IsNil(o.RadioUtilInterference) {
		toSerialize["radio_util_interference"] = o.RadioUtilInterference
	}
	if !IsNil(o.SiteNumClients) {
		toSerialize["site_num_clients"] = o.SiteNumClients
	}
	if !IsNil(o.WanAvgDownloadMbps) {
		toSerialize["wan_avg_download_mbps"] = o.WanAvgDownloadMbps
	}
	if !IsNil(o.WanAvgUploadMbps) {
		toSerialize["wan_avg_upload_mbps"] = o.WanAvgUploadMbps
	}
	if !IsNil(o.WanJitter) {
		toSerialize["wan_jitter"] = o.WanJitter
	}
	if !IsNil(o.WanMaxDownloadMbps) {
		toSerialize["wan_max_download_mbps"] = o.WanMaxDownloadMbps
	}
	if !IsNil(o.WanMaxUploadMbps) {
		toSerialize["wan_max_upload_mbps"] = o.WanMaxUploadMbps
	}
	if !IsNil(o.WanRtt) {
		toSerialize["wan_rtt"] = o.WanRtt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallTroubleshootSummaryData) UnmarshalJSON(data []byte) (err error) {
	varCallTroubleshootSummaryData := _CallTroubleshootSummaryData{}

	err = json.Unmarshal(data, &varCallTroubleshootSummaryData)

	if err != nil {
		return err
	}

	*o = CallTroubleshootSummaryData(varCallTroubleshootSummaryData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ap_num_clients")
		delete(additionalProperties, "ap_rtt")
		delete(additionalProperties, "client_cpu")
		delete(additionalProperties, "client_n_streams")
		delete(additionalProperties, "client_radio_band")
		delete(additionalProperties, "client_rssi")
		delete(additionalProperties, "client_rx_bytes")
		delete(additionalProperties, "client_rx_rates")
		delete(additionalProperties, "client_tx_bytes")
		delete(additionalProperties, "client_tx_rates")
		delete(additionalProperties, "client_tx_retries")
		delete(additionalProperties, "client_vpn_distaince")
		delete(additionalProperties, "client_wifi_version")
		delete(additionalProperties, "expected")
		delete(additionalProperties, "radio_bandwidth")
		delete(additionalProperties, "radio_channel")
		delete(additionalProperties, "radio_tx_power")
		delete(additionalProperties, "radio_util")
		delete(additionalProperties, "radio_util_interference")
		delete(additionalProperties, "site_num_clients")
		delete(additionalProperties, "wan_avg_download_mbps")
		delete(additionalProperties, "wan_avg_upload_mbps")
		delete(additionalProperties, "wan_jitter")
		delete(additionalProperties, "wan_max_download_mbps")
		delete(additionalProperties, "wan_max_upload_mbps")
		delete(additionalProperties, "wan_rtt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallTroubleshootSummaryData struct {
	value *CallTroubleshootSummaryData
	isSet bool
}

func (v NullableCallTroubleshootSummaryData) Get() *CallTroubleshootSummaryData {
	return v.value
}

func (v *NullableCallTroubleshootSummaryData) Set(val *CallTroubleshootSummaryData) {
	v.value = val
	v.isSet = true
}

func (v NullableCallTroubleshootSummaryData) IsSet() bool {
	return v.isSet
}

func (v *NullableCallTroubleshootSummaryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallTroubleshootSummaryData(val *CallTroubleshootSummaryData) *NullableCallTroubleshootSummaryData {
	return &NullableCallTroubleshootSummaryData{value: val, isSet: true}
}

func (v NullableCallTroubleshootSummaryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallTroubleshootSummaryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


