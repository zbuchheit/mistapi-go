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


type SitesInsightsAPI interface {

	/*
	GetSiteInsightMetrics getSiteInsightMetrics

	Get Site Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param metric see /api/v1/const/insight_metrics for available metrics
	@return ApiGetSiteInsightMetricsRequest
	*/
	GetSiteInsightMetrics(ctx context.Context, siteId string, metric string) ApiGetSiteInsightMetricsRequest

	// GetSiteInsightMetricsExecute executes the request
	//  @return InsightMetrics
	GetSiteInsightMetricsExecute(r ApiGetSiteInsightMetricsRequest) (*InsightMetrics, *http.Response, error)

	/*
	GetSiteInsightMetricsForClient getSiteInsightMetricsForClient

	Get Client Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param clientMac
	@param metric see /api/v1/const/insight_metrics for available metrics
	@return ApiGetSiteInsightMetricsForClientRequest
	*/
	GetSiteInsightMetricsForClient(ctx context.Context, siteId string, clientMac string, metric string) ApiGetSiteInsightMetricsForClientRequest

	// GetSiteInsightMetricsForClientExecute executes the request
	//  @return InsightMetrics
	GetSiteInsightMetricsForClientExecute(r ApiGetSiteInsightMetricsForClientRequest) (*InsightMetrics, *http.Response, error)

	/*
	GetSiteInsightMetricsForDevice getSiteInsightMetricsForDevice

	Get AP Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param metric see /api/v1/const/insight_metrics for available metrics
	@param deviceMac
	@return ApiGetSiteInsightMetricsForDeviceRequest
	*/
	GetSiteInsightMetricsForDevice(ctx context.Context, siteId string, metric string, deviceMac string) ApiGetSiteInsightMetricsForDeviceRequest

	// GetSiteInsightMetricsForDeviceExecute executes the request
	//  @return ResponseDeviceMetrics
	GetSiteInsightMetricsForDeviceExecute(r ApiGetSiteInsightMetricsForDeviceRequest) (*ResponseDeviceMetrics, *http.Response, error)
}

// SitesInsightsAPIService SitesInsightsAPI service
type SitesInsightsAPIService service

type ApiGetSiteInsightMetricsRequest struct {
	ctx context.Context
	ApiService SitesInsightsAPI
	siteId string
	metric string
	page *int32
	limit *int32
	start *int32
	end *int32
	duration *string
	interval *string
}

func (r ApiGetSiteInsightMetricsRequest) Page(page int32) ApiGetSiteInsightMetricsRequest {
	r.page = &page
	return r
}

func (r ApiGetSiteInsightMetricsRequest) Limit(limit int32) ApiGetSiteInsightMetricsRequest {
	r.limit = &limit
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiGetSiteInsightMetricsRequest) Start(start int32) ApiGetSiteInsightMetricsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiGetSiteInsightMetricsRequest) End(end int32) ApiGetSiteInsightMetricsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiGetSiteInsightMetricsRequest) Duration(duration string) ApiGetSiteInsightMetricsRequest {
	r.duration = &duration
	return r
}

// Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to.
func (r ApiGetSiteInsightMetricsRequest) Interval(interval string) ApiGetSiteInsightMetricsRequest {
	r.interval = &interval
	return r
}

func (r ApiGetSiteInsightMetricsRequest) Execute() (*InsightMetrics, *http.Response, error) {
	return r.ApiService.GetSiteInsightMetricsExecute(r)
}

/*
GetSiteInsightMetrics getSiteInsightMetrics

Get Site Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param metric see /api/v1/const/insight_metrics for available metrics
 @return ApiGetSiteInsightMetricsRequest
*/
func (a *SitesInsightsAPIService) GetSiteInsightMetrics(ctx context.Context, siteId string, metric string) ApiGetSiteInsightMetricsRequest {
	return ApiGetSiteInsightMetricsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		metric: metric,
	}
}

// Execute executes the request
//  @return InsightMetrics
func (a *SitesInsightsAPIService) GetSiteInsightMetricsExecute(r ApiGetSiteInsightMetricsRequest) (*InsightMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InsightMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesInsightsAPIService.GetSiteInsightMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/insights/{metric}"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metric"+"}", url.PathEscape(parameterValueToString(r.metric, "metric")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
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

type ApiGetSiteInsightMetricsForClientRequest struct {
	ctx context.Context
	ApiService SitesInsightsAPI
	siteId string
	clientMac string
	metric string
	page *int32
	limit *int32
	start *int32
	end *int32
	duration *string
	interval *string
}

func (r ApiGetSiteInsightMetricsForClientRequest) Page(page int32) ApiGetSiteInsightMetricsForClientRequest {
	r.page = &page
	return r
}

func (r ApiGetSiteInsightMetricsForClientRequest) Limit(limit int32) ApiGetSiteInsightMetricsForClientRequest {
	r.limit = &limit
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiGetSiteInsightMetricsForClientRequest) Start(start int32) ApiGetSiteInsightMetricsForClientRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiGetSiteInsightMetricsForClientRequest) End(end int32) ApiGetSiteInsightMetricsForClientRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiGetSiteInsightMetricsForClientRequest) Duration(duration string) ApiGetSiteInsightMetricsForClientRequest {
	r.duration = &duration
	return r
}

// Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to.
func (r ApiGetSiteInsightMetricsForClientRequest) Interval(interval string) ApiGetSiteInsightMetricsForClientRequest {
	r.interval = &interval
	return r
}

func (r ApiGetSiteInsightMetricsForClientRequest) Execute() (*InsightMetrics, *http.Response, error) {
	return r.ApiService.GetSiteInsightMetricsForClientExecute(r)
}

/*
GetSiteInsightMetricsForClient getSiteInsightMetricsForClient

Get Client Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param clientMac
 @param metric see /api/v1/const/insight_metrics for available metrics
 @return ApiGetSiteInsightMetricsForClientRequest
*/
func (a *SitesInsightsAPIService) GetSiteInsightMetricsForClient(ctx context.Context, siteId string, clientMac string, metric string) ApiGetSiteInsightMetricsForClientRequest {
	return ApiGetSiteInsightMetricsForClientRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		clientMac: clientMac,
		metric: metric,
	}
}

// Execute executes the request
//  @return InsightMetrics
func (a *SitesInsightsAPIService) GetSiteInsightMetricsForClientExecute(r ApiGetSiteInsightMetricsForClientRequest) (*InsightMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InsightMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesInsightsAPIService.GetSiteInsightMetricsForClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/insights/client/{client_mac}/{metric}"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client_mac"+"}", url.PathEscape(parameterValueToString(r.clientMac, "clientMac")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metric"+"}", url.PathEscape(parameterValueToString(r.metric, "metric")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
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

type ApiGetSiteInsightMetricsForDeviceRequest struct {
	ctx context.Context
	ApiService SitesInsightsAPI
	siteId string
	metric string
	deviceMac string
	page *int32
	limit *int32
	start *int32
	end *int32
	duration *string
	interval *string
}

func (r ApiGetSiteInsightMetricsForDeviceRequest) Page(page int32) ApiGetSiteInsightMetricsForDeviceRequest {
	r.page = &page
	return r
}

func (r ApiGetSiteInsightMetricsForDeviceRequest) Limit(limit int32) ApiGetSiteInsightMetricsForDeviceRequest {
	r.limit = &limit
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiGetSiteInsightMetricsForDeviceRequest) Start(start int32) ApiGetSiteInsightMetricsForDeviceRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiGetSiteInsightMetricsForDeviceRequest) End(end int32) ApiGetSiteInsightMetricsForDeviceRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiGetSiteInsightMetricsForDeviceRequest) Duration(duration string) ApiGetSiteInsightMetricsForDeviceRequest {
	r.duration = &duration
	return r
}

// Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to.
func (r ApiGetSiteInsightMetricsForDeviceRequest) Interval(interval string) ApiGetSiteInsightMetricsForDeviceRequest {
	r.interval = &interval
	return r
}

func (r ApiGetSiteInsightMetricsForDeviceRequest) Execute() (*ResponseDeviceMetrics, *http.Response, error) {
	return r.ApiService.GetSiteInsightMetricsForDeviceExecute(r)
}

/*
GetSiteInsightMetricsForDevice getSiteInsightMetricsForDevice

Get AP Insight Metrics
See metrics possibilities at /api/v1/const/insight_metrics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param metric see /api/v1/const/insight_metrics for available metrics
 @param deviceMac
 @return ApiGetSiteInsightMetricsForDeviceRequest
*/
func (a *SitesInsightsAPIService) GetSiteInsightMetricsForDevice(ctx context.Context, siteId string, metric string, deviceMac string) ApiGetSiteInsightMetricsForDeviceRequest {
	return ApiGetSiteInsightMetricsForDeviceRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		metric: metric,
		deviceMac: deviceMac,
	}
}

// Execute executes the request
//  @return ResponseDeviceMetrics
func (a *SitesInsightsAPIService) GetSiteInsightMetricsForDeviceExecute(r ApiGetSiteInsightMetricsForDeviceRequest) (*ResponseDeviceMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseDeviceMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesInsightsAPIService.GetSiteInsightMetricsForDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/insights/device/{device_mac}/{metric}"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metric"+"}", url.PathEscape(parameterValueToString(r.metric, "metric")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_mac"+"}", url.PathEscape(parameterValueToString(r.deviceMac, "deviceMac")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
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
