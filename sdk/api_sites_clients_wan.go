/*
Mist API

> Version: **2406.1.16** > > Date: **July 4, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.16
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type SitesClientsWanAPI interface {

	/*
	CountSiteWanClientEvents countSiteWanClientEvents

	Count by Distinct Attributes of Site WAN Client-Events

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@return ApiCountSiteWanClientEventsRequest
	*/
	CountSiteWanClientEvents(ctx context.Context, siteId string) ApiCountSiteWanClientEventsRequest

	// CountSiteWanClientEventsExecute executes the request
	//  @return RepsonseCount
	CountSiteWanClientEventsExecute(r ApiCountSiteWanClientEventsRequest) (*RepsonseCount, *http.Response, error)

	/*
	CountSiteWanClients countSiteWanClients

	Count Site WAN Clients

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@return ApiCountSiteWanClientsRequest
	*/
	CountSiteWanClients(ctx context.Context, siteId string) ApiCountSiteWanClientsRequest

	// CountSiteWanClientsExecute executes the request
	//  @return RepsonseCount
	CountSiteWanClientsExecute(r ApiCountSiteWanClientsRequest) (*RepsonseCount, *http.Response, error)

	/*
	SearchSiteWanClientEvents searchSiteWanClientEvents

	Search Site WAN Client Events

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@return ApiSearchSiteWanClientEventsRequest
	*/
	SearchSiteWanClientEvents(ctx context.Context, siteId string) ApiSearchSiteWanClientEventsRequest

	// SearchSiteWanClientEventsExecute executes the request
	//  @return SearchEventsWanClient
	SearchSiteWanClientEventsExecute(r ApiSearchSiteWanClientEventsRequest) (*SearchEventsWanClient, *http.Response, error)

	/*
	SearchSiteWanClients searchSiteWanClients

	Search Site WAN Clients

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@return ApiSearchSiteWanClientsRequest
	*/
	SearchSiteWanClients(ctx context.Context, siteId string) ApiSearchSiteWanClientsRequest

	// SearchSiteWanClientsExecute executes the request
	//  @return SearchWanClient
	SearchSiteWanClientsExecute(r ApiSearchSiteWanClientsRequest) (*SearchWanClient, *http.Response, error)
}

// SitesClientsWanAPIService SitesClientsWanAPI service
type SitesClientsWanAPIService service

type ApiCountSiteWanClientEventsRequest struct {
	ctx context.Context
	ApiService SitesClientsWanAPI
	siteId string
	distinct *SiteWanClientEventsDistinct
	type_ *string
	start *int32
	end *int32
	duration *string
	limit *int32
}

func (r ApiCountSiteWanClientEventsRequest) Distinct(distinct SiteWanClientEventsDistinct) ApiCountSiteWanClientEventsRequest {
	r.distinct = &distinct
	return r
}

// see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions)
func (r ApiCountSiteWanClientEventsRequest) Type_(type_ string) ApiCountSiteWanClientEventsRequest {
	r.type_ = &type_
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiCountSiteWanClientEventsRequest) Start(start int32) ApiCountSiteWanClientEventsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiCountSiteWanClientEventsRequest) End(end int32) ApiCountSiteWanClientEventsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiCountSiteWanClientEventsRequest) Duration(duration string) ApiCountSiteWanClientEventsRequest {
	r.duration = &duration
	return r
}

func (r ApiCountSiteWanClientEventsRequest) Limit(limit int32) ApiCountSiteWanClientEventsRequest {
	r.limit = &limit
	return r
}

func (r ApiCountSiteWanClientEventsRequest) Execute() (*RepsonseCount, *http.Response, error) {
	return r.ApiService.CountSiteWanClientEventsExecute(r)
}

/*
CountSiteWanClientEvents countSiteWanClientEvents

Count by Distinct Attributes of Site WAN Client-Events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @return ApiCountSiteWanClientEventsRequest
*/
func (a *SitesClientsWanAPIService) CountSiteWanClientEvents(ctx context.Context, siteId string) ApiCountSiteWanClientEventsRequest {
	return ApiCountSiteWanClientEventsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return RepsonseCount
func (a *SitesClientsWanAPIService) CountSiteWanClientEventsExecute(r ApiCountSiteWanClientEventsRequest) (*RepsonseCount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepsonseCount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesClientsWanAPIService.CountSiteWanClientEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/wan_client/events/count"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.distinct != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "distinct", r.distinct, "")
	} else {
		var defaultValue SiteWanClientEventsDistinct = "type"
		r.distinct = &defaultValue
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
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

type ApiCountSiteWanClientsRequest struct {
	ctx context.Context
	ApiService SitesClientsWanAPI
	siteId string
	distinct *SiteWanClientsCountDistinct
	start *int32
	end *int32
	duration *string
	limit *int32
	page *int32
}

func (r ApiCountSiteWanClientsRequest) Distinct(distinct SiteWanClientsCountDistinct) ApiCountSiteWanClientsRequest {
	r.distinct = &distinct
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiCountSiteWanClientsRequest) Start(start int32) ApiCountSiteWanClientsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiCountSiteWanClientsRequest) End(end int32) ApiCountSiteWanClientsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiCountSiteWanClientsRequest) Duration(duration string) ApiCountSiteWanClientsRequest {
	r.duration = &duration
	return r
}

func (r ApiCountSiteWanClientsRequest) Limit(limit int32) ApiCountSiteWanClientsRequest {
	r.limit = &limit
	return r
}

func (r ApiCountSiteWanClientsRequest) Page(page int32) ApiCountSiteWanClientsRequest {
	r.page = &page
	return r
}

func (r ApiCountSiteWanClientsRequest) Execute() (*RepsonseCount, *http.Response, error) {
	return r.ApiService.CountSiteWanClientsExecute(r)
}

/*
CountSiteWanClients countSiteWanClients

Count Site WAN Clients

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @return ApiCountSiteWanClientsRequest
*/
func (a *SitesClientsWanAPIService) CountSiteWanClients(ctx context.Context, siteId string) ApiCountSiteWanClientsRequest {
	return ApiCountSiteWanClientsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return RepsonseCount
func (a *SitesClientsWanAPIService) CountSiteWanClientsExecute(r ApiCountSiteWanClientsRequest) (*RepsonseCount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepsonseCount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesClientsWanAPIService.CountSiteWanClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/wan_clients/count"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.distinct != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "distinct", r.distinct, "")
	} else {
		var defaultValue SiteWanClientsCountDistinct = "mac"
		r.distinct = &defaultValue
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
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
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

type ApiSearchSiteWanClientEventsRequest struct {
	ctx context.Context
	ApiService SitesClientsWanAPI
	siteId string
	type_ *string
	mac *string
	hostname *string
	ip *string
	mfg *string
	nacruleId *string
	start *int32
	end *int32
	duration *string
	limit *int32
}

// see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions)
func (r ApiSearchSiteWanClientEventsRequest) Type_(type_ string) ApiSearchSiteWanClientEventsRequest {
	r.type_ = &type_
	return r
}

// partial / full MAC address
func (r ApiSearchSiteWanClientEventsRequest) Mac(mac string) ApiSearchSiteWanClientEventsRequest {
	r.mac = &mac
	return r
}

// partial / full hostname
func (r ApiSearchSiteWanClientEventsRequest) Hostname(hostname string) ApiSearchSiteWanClientEventsRequest {
	r.hostname = &hostname
	return r
}

// client IP
func (r ApiSearchSiteWanClientEventsRequest) Ip(ip string) ApiSearchSiteWanClientEventsRequest {
	r.ip = &ip
	return r
}

// Manufacture
func (r ApiSearchSiteWanClientEventsRequest) Mfg(mfg string) ApiSearchSiteWanClientEventsRequest {
	r.mfg = &mfg
	return r
}

// nacrule_id
func (r ApiSearchSiteWanClientEventsRequest) NacruleId(nacruleId string) ApiSearchSiteWanClientEventsRequest {
	r.nacruleId = &nacruleId
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiSearchSiteWanClientEventsRequest) Start(start int32) ApiSearchSiteWanClientEventsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiSearchSiteWanClientEventsRequest) End(end int32) ApiSearchSiteWanClientEventsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiSearchSiteWanClientEventsRequest) Duration(duration string) ApiSearchSiteWanClientEventsRequest {
	r.duration = &duration
	return r
}

func (r ApiSearchSiteWanClientEventsRequest) Limit(limit int32) ApiSearchSiteWanClientEventsRequest {
	r.limit = &limit
	return r
}

func (r ApiSearchSiteWanClientEventsRequest) Execute() (*SearchEventsWanClient, *http.Response, error) {
	return r.ApiService.SearchSiteWanClientEventsExecute(r)
}

/*
SearchSiteWanClientEvents searchSiteWanClientEvents

Search Site WAN Client Events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @return ApiSearchSiteWanClientEventsRequest
*/
func (a *SitesClientsWanAPIService) SearchSiteWanClientEvents(ctx context.Context, siteId string) ApiSearchSiteWanClientEventsRequest {
	return ApiSearchSiteWanClientEventsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return SearchEventsWanClient
func (a *SitesClientsWanAPIService) SearchSiteWanClientEventsExecute(r ApiSearchSiteWanClientEventsRequest) (*SearchEventsWanClient, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchEventsWanClient
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesClientsWanAPIService.SearchSiteWanClientEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/wan_clients/events/search"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.mac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mac", r.mac, "")
	}
	if r.hostname != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hostname", r.hostname, "")
	}
	if r.ip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ip", r.ip, "")
	}
	if r.mfg != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mfg", r.mfg, "")
	}
	if r.nacruleId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nacrule_id", r.nacruleId, "")
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
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
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

type ApiSearchSiteWanClientsRequest struct {
	ctx context.Context
	ApiService SitesClientsWanAPI
	siteId string
	mac *string
	hostname *string
	ip *string
	mfg *string
	start *int32
	end *int32
	duration *string
	limit *int32
	page *int32
}

// partial / full MAC address
func (r ApiSearchSiteWanClientsRequest) Mac(mac string) ApiSearchSiteWanClientsRequest {
	r.mac = &mac
	return r
}

// partial / full hostname
func (r ApiSearchSiteWanClientsRequest) Hostname(hostname string) ApiSearchSiteWanClientsRequest {
	r.hostname = &hostname
	return r
}

// client IP
func (r ApiSearchSiteWanClientsRequest) Ip(ip string) ApiSearchSiteWanClientsRequest {
	r.ip = &ip
	return r
}

// Manufacture
func (r ApiSearchSiteWanClientsRequest) Mfg(mfg string) ApiSearchSiteWanClientsRequest {
	r.mfg = &mfg
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiSearchSiteWanClientsRequest) Start(start int32) ApiSearchSiteWanClientsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiSearchSiteWanClientsRequest) End(end int32) ApiSearchSiteWanClientsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiSearchSiteWanClientsRequest) Duration(duration string) ApiSearchSiteWanClientsRequest {
	r.duration = &duration
	return r
}

func (r ApiSearchSiteWanClientsRequest) Limit(limit int32) ApiSearchSiteWanClientsRequest {
	r.limit = &limit
	return r
}

func (r ApiSearchSiteWanClientsRequest) Page(page int32) ApiSearchSiteWanClientsRequest {
	r.page = &page
	return r
}

func (r ApiSearchSiteWanClientsRequest) Execute() (*SearchWanClient, *http.Response, error) {
	return r.ApiService.SearchSiteWanClientsExecute(r)
}

/*
SearchSiteWanClients searchSiteWanClients

Search Site WAN Clients

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @return ApiSearchSiteWanClientsRequest
*/
func (a *SitesClientsWanAPIService) SearchSiteWanClients(ctx context.Context, siteId string) ApiSearchSiteWanClientsRequest {
	return ApiSearchSiteWanClientsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return SearchWanClient
func (a *SitesClientsWanAPIService) SearchSiteWanClientsExecute(r ApiSearchSiteWanClientsRequest) (*SearchWanClient, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchWanClient
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesClientsWanAPIService.SearchSiteWanClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/wan_clients/search"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mac", r.mac, "")
	}
	if r.hostname != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hostname", r.hostname, "")
	}
	if r.ip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ip", r.ip, "")
	}
	if r.mfg != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mfg", r.mfg, "")
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
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
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
