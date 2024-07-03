/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistsdkgo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type OrgsClientsWiredAPI interface {

	/*
	CountOrgWiredClients countOrgWiredClients

	Count by Distinct Attributes of Clients

Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId
	@return ApiCountOrgWiredClientsRequest
	*/
	CountOrgWiredClients(ctx context.Context, orgId string) ApiCountOrgWiredClientsRequest

	// CountOrgWiredClientsExecute executes the request
	//  @return RepsonseCount
	CountOrgWiredClientsExecute(r ApiCountOrgWiredClientsRequest) (*RepsonseCount, *http.Response, error)

	/*
	SearchOrgWiredClients searchOrgWiredClients

	Search for Wired Clients in org

Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId
	@return ApiSearchOrgWiredClientsRequest
	*/
	SearchOrgWiredClients(ctx context.Context, orgId string) ApiSearchOrgWiredClientsRequest

	// SearchOrgWiredClientsExecute executes the request
	//  @return SearchWiredClient
	SearchOrgWiredClientsExecute(r ApiSearchOrgWiredClientsRequest) (*SearchWiredClient, *http.Response, error)
}

// OrgsClientsWiredAPIService OrgsClientsWiredAPI service
type OrgsClientsWiredAPIService service

type ApiCountOrgWiredClientsRequest struct {
	ctx context.Context
	ApiService OrgsClientsWiredAPI
	orgId string
	distinct *OrgWiredClientsCountDistinct
	page *int32
	limit *int32
	start *int32
	end *int32
	duration *string
}

func (r ApiCountOrgWiredClientsRequest) Distinct(distinct OrgWiredClientsCountDistinct) ApiCountOrgWiredClientsRequest {
	r.distinct = &distinct
	return r
}

func (r ApiCountOrgWiredClientsRequest) Page(page int32) ApiCountOrgWiredClientsRequest {
	r.page = &page
	return r
}

func (r ApiCountOrgWiredClientsRequest) Limit(limit int32) ApiCountOrgWiredClientsRequest {
	r.limit = &limit
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiCountOrgWiredClientsRequest) Start(start int32) ApiCountOrgWiredClientsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiCountOrgWiredClientsRequest) End(end int32) ApiCountOrgWiredClientsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiCountOrgWiredClientsRequest) Duration(duration string) ApiCountOrgWiredClientsRequest {
	r.duration = &duration
	return r
}

func (r ApiCountOrgWiredClientsRequest) Execute() (*RepsonseCount, *http.Response, error) {
	return r.ApiService.CountOrgWiredClientsExecute(r)
}

/*
CountOrgWiredClients countOrgWiredClients

Count by Distinct Attributes of Clients

Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @return ApiCountOrgWiredClientsRequest
*/
func (a *OrgsClientsWiredAPIService) CountOrgWiredClients(ctx context.Context, orgId string) ApiCountOrgWiredClientsRequest {
	return ApiCountOrgWiredClientsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return RepsonseCount
func (a *OrgsClientsWiredAPIService) CountOrgWiredClientsExecute(r ApiCountOrgWiredClientsRequest) (*RepsonseCount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepsonseCount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsClientsWiredAPIService.CountOrgWiredClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/orgs/{org_id}/wired_clients/count"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.distinct != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "distinct", r.distinct, "")
	} else {
		var defaultValue OrgWiredClientsCountDistinct = "mac"
		r.distinct = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.duration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "")
	} else {
		var defaultValue string = "1d"
		r.duration = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponseHttp400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchOrgWiredClientsRequest struct {
	ctx context.Context
	ApiService OrgsClientsWiredAPI
	orgId string
	siteId *string
	deviceMac *string
	mac *string
	portId *string
	vlan *int32
	ip *string
	manufacture *string
	text *string
	nacruleId *string
	limit *int32
	start *int32
	end *int32
	duration *string
}

// Site ID
func (r ApiSearchOrgWiredClientsRequest) SiteId(siteId string) ApiSearchOrgWiredClientsRequest {
	r.siteId = &siteId
	return r
}

// device mac
func (r ApiSearchOrgWiredClientsRequest) DeviceMac(deviceMac string) ApiSearchOrgWiredClientsRequest {
	r.deviceMac = &deviceMac
	return r
}

// client mac
func (r ApiSearchOrgWiredClientsRequest) Mac(mac string) ApiSearchOrgWiredClientsRequest {
	r.mac = &mac
	return r
}

// port id
func (r ApiSearchOrgWiredClientsRequest) PortId(portId string) ApiSearchOrgWiredClientsRequest {
	r.portId = &portId
	return r
}

// vlan
func (r ApiSearchOrgWiredClientsRequest) Vlan(vlan int32) ApiSearchOrgWiredClientsRequest {
	r.vlan = &vlan
	return r
}

// ip
func (r ApiSearchOrgWiredClientsRequest) Ip(ip string) ApiSearchOrgWiredClientsRequest {
	r.ip = &ip
	return r
}

// client manufacturer
func (r ApiSearchOrgWiredClientsRequest) Manufacture(manufacture string) ApiSearchOrgWiredClientsRequest {
	r.manufacture = &manufacture
	return r
}

// single entry of hostname/mac
func (r ApiSearchOrgWiredClientsRequest) Text(text string) ApiSearchOrgWiredClientsRequest {
	r.text = &text
	return r
}

// nacrule_id
func (r ApiSearchOrgWiredClientsRequest) NacruleId(nacruleId string) ApiSearchOrgWiredClientsRequest {
	r.nacruleId = &nacruleId
	return r
}

func (r ApiSearchOrgWiredClientsRequest) Limit(limit int32) ApiSearchOrgWiredClientsRequest {
	r.limit = &limit
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiSearchOrgWiredClientsRequest) Start(start int32) ApiSearchOrgWiredClientsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiSearchOrgWiredClientsRequest) End(end int32) ApiSearchOrgWiredClientsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiSearchOrgWiredClientsRequest) Duration(duration string) ApiSearchOrgWiredClientsRequest {
	r.duration = &duration
	return r
}

func (r ApiSearchOrgWiredClientsRequest) Execute() (*SearchWiredClient, *http.Response, error) {
	return r.ApiService.SearchOrgWiredClientsExecute(r)
}

/*
SearchOrgWiredClients searchOrgWiredClients

Search for Wired Clients in org

Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @return ApiSearchOrgWiredClientsRequest
*/
func (a *OrgsClientsWiredAPIService) SearchOrgWiredClients(ctx context.Context, orgId string) ApiSearchOrgWiredClientsRequest {
	return ApiSearchOrgWiredClientsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return SearchWiredClient
func (a *OrgsClientsWiredAPIService) SearchOrgWiredClientsExecute(r ApiSearchOrgWiredClientsRequest) (*SearchWiredClient, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchWiredClient
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsClientsWiredAPIService.SearchOrgWiredClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/orgs/{org_id}/wired_clients/search"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	if r.deviceMac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device_mac", r.deviceMac, "")
	}
	if r.mac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mac", r.mac, "")
	}
	if r.portId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "port_id", r.portId, "")
	}
	if r.vlan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vlan", r.vlan, "")
	}
	if r.ip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ip", r.ip, "")
	}
	if r.manufacture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "manufacture", r.manufacture, "")
	}
	if r.text != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "text", r.text, "")
	}
	if r.nacruleId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nacrule_id", r.nacruleId, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.duration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "")
	} else {
		var defaultValue string = "1d"
		r.duration = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponseHttp400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
