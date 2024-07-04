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


type SitesMapsAutoPlacementAPI interface {

	/*
	ClearSiteApAutoplacement clearSiteApAutoplacement

	This API is used to destroy the cached autoplacement locations of a map or subset of APs on a map.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param mapId
	@return ApiClearSiteApAutoplacementRequest
	*/
	ClearSiteApAutoplacement(ctx context.Context, siteId string, mapId string) ApiClearSiteApAutoplacementRequest

	// ClearSiteApAutoplacementExecute executes the request
	ClearSiteApAutoplacementExecute(r ApiClearSiteApAutoplacementRequest) (*http.Response, error)

	/*
	ConfirmSiteApLocalizationData confirmSiteApLocalizationData

	This API is used to accept or reject the cached autoplacement and auto orientation values of a map or subset of APs on a map. A rejected AP will retain its current X,Y and orientation until accpeted.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param mapId
	@return ApiConfirmSiteApLocalizationDataRequest
	*/
	ConfirmSiteApLocalizationData(ctx context.Context, siteId string, mapId string) ApiConfirmSiteApLocalizationDataRequest

	// ConfirmSiteApLocalizationDataExecute executes the request
	ConfirmSiteApLocalizationDataExecute(r ApiConfirmSiteApLocalizationDataRequest) (*http.Response, error)

	/*
	DeleteSiteApAutoplacement deleteSiteApAutoplacement

	This API is called to force stop auto placement for a given map

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param mapId
	@return ApiDeleteSiteApAutoplacementRequest
	*/
	DeleteSiteApAutoplacement(ctx context.Context, siteId string, mapId string) ApiDeleteSiteApAutoplacementRequest

	// DeleteSiteApAutoplacementExecute executes the request
	DeleteSiteApAutoplacementExecute(r ApiDeleteSiteApAutoplacementRequest) (*http.Response, error)

	/*
	GetSiteApAutoPlacement getSiteApAutoplacement

	This API is called to view the current status of auto placement for a given map.


#### Status Descriptions

| Status | Description |
| --- | --- |
| `pending` | Autoplacement has not been requested for this map |
| `inprogress` | Autoplacement is currently processing |
| `done` | The autoplacement process has completed |
| `data_needed` | Additional position data is required for autoplacement. Users should verify the requested anchor APs have a position on the map |
| `invalid_model` | Autoplacement is not supported on the model of the APs on the map |
| `invalid_version` | Autoplacement is not supported with the APs current firmware version |
| `error` | There was an error in the autoplacement process |

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param mapId
	@return ApiGetSiteApAutoPlacementRequest
	*/
	GetSiteApAutoPlacement(ctx context.Context, siteId string, mapId string) ApiGetSiteApAutoPlacementRequest

	// GetSiteApAutoPlacementExecute executes the request
	//  @return ResponseAutoPlacementInfo
	GetSiteApAutoPlacementExecute(r ApiGetSiteApAutoPlacementRequest) (*ResponseAutoPlacementInfo, *http.Response, error)

	/*
	RunSiteApAutoplacement runSiteApAutoplacement

	This API is called to trigger a map for auto placement. For auto placement feature to work, RTT-FTM data need to be collected from the APs on the map. This scan is disruptive and therefore the user must be notified of service disrution during the functioning of auto placement Repeated POST to this endpoint while a map is still running will be rejected.

List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide autoplacement suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param mapId
	@return ApiRunSiteApAutoplacementRequest
	*/
	RunSiteApAutoplacement(ctx context.Context, siteId string, mapId string) ApiRunSiteApAutoplacementRequest

	// RunSiteApAutoplacementExecute executes the request
	RunSiteApAutoplacementExecute(r ApiRunSiteApAutoplacementRequest) (*http.Response, error)
}

// SitesMapsAutoPlacementAPIService SitesMapsAutoPlacementAPI service
type SitesMapsAutoPlacementAPIService service

type ApiClearSiteApAutoplacementRequest struct {
	ctx context.Context
	ApiService SitesMapsAutoPlacementAPI
	siteId string
	mapId string
	macAddresses *MacAddresses
}

func (r ApiClearSiteApAutoplacementRequest) MacAddresses(macAddresses MacAddresses) ApiClearSiteApAutoplacementRequest {
	r.macAddresses = &macAddresses
	return r
}

func (r ApiClearSiteApAutoplacementRequest) Execute() (*http.Response, error) {
	return r.ApiService.ClearSiteApAutoplacementExecute(r)
}

/*
ClearSiteApAutoplacement clearSiteApAutoplacement

This API is used to destroy the cached autoplacement locations of a map or subset of APs on a map.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param mapId
 @return ApiClearSiteApAutoplacementRequest
*/
func (a *SitesMapsAutoPlacementAPIService) ClearSiteApAutoplacement(ctx context.Context, siteId string, mapId string) ApiClearSiteApAutoplacementRequest {
	return ApiClearSiteApAutoplacementRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		mapId: mapId,
	}
}

// Execute executes the request
func (a *SitesMapsAutoPlacementAPIService) ClearSiteApAutoplacementExecute(r ApiClearSiteApAutoplacementRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesMapsAutoPlacementAPIService.ClearSiteApAutoplacement")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/maps/{map_id}/clear_autoplacement"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_id"+"}", url.PathEscape(parameterValueToString(r.mapId, "mapId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.macAddresses
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiConfirmSiteApLocalizationDataRequest struct {
	ctx context.Context
	ApiService SitesMapsAutoPlacementAPI
	siteId string
	mapId string
	useAutoApValues *UseAutoApValues
}

func (r ApiConfirmSiteApLocalizationDataRequest) UseAutoApValues(useAutoApValues UseAutoApValues) ApiConfirmSiteApLocalizationDataRequest {
	r.useAutoApValues = &useAutoApValues
	return r
}

func (r ApiConfirmSiteApLocalizationDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.ConfirmSiteApLocalizationDataExecute(r)
}

/*
ConfirmSiteApLocalizationData confirmSiteApLocalizationData

This API is used to accept or reject the cached autoplacement and auto orientation values of a map or subset of APs on a map. A rejected AP will retain its current X,Y and orientation until accpeted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param mapId
 @return ApiConfirmSiteApLocalizationDataRequest
*/
func (a *SitesMapsAutoPlacementAPIService) ConfirmSiteApLocalizationData(ctx context.Context, siteId string, mapId string) ApiConfirmSiteApLocalizationDataRequest {
	return ApiConfirmSiteApLocalizationDataRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		mapId: mapId,
	}
}

// Execute executes the request
func (a *SitesMapsAutoPlacementAPIService) ConfirmSiteApLocalizationDataExecute(r ApiConfirmSiteApLocalizationDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesMapsAutoPlacementAPIService.ConfirmSiteApLocalizationData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/maps/{map_id}/use_auto_ap_values"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_id"+"}", url.PathEscape(parameterValueToString(r.mapId, "mapId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.useAutoApValues
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteSiteApAutoplacementRequest struct {
	ctx context.Context
	ApiService SitesMapsAutoPlacementAPI
	siteId string
	mapId string
}

func (r ApiDeleteSiteApAutoplacementRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSiteApAutoplacementExecute(r)
}

/*
DeleteSiteApAutoplacement deleteSiteApAutoplacement

This API is called to force stop auto placement for a given map

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param mapId
 @return ApiDeleteSiteApAutoplacementRequest
*/
func (a *SitesMapsAutoPlacementAPIService) DeleteSiteApAutoplacement(ctx context.Context, siteId string, mapId string) ApiDeleteSiteApAutoplacementRequest {
	return ApiDeleteSiteApAutoplacementRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		mapId: mapId,
	}
}

// Execute executes the request
func (a *SitesMapsAutoPlacementAPIService) DeleteSiteApAutoplacementExecute(r ApiDeleteSiteApAutoplacementRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesMapsAutoPlacementAPIService.DeleteSiteApAutoplacement")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/maps/{map_id}/auto_placement"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_id"+"}", url.PathEscape(parameterValueToString(r.mapId, "mapId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSiteApAutoPlacementRequest struct {
	ctx context.Context
	ApiService SitesMapsAutoPlacementAPI
	siteId string
	mapId string
}

func (r ApiGetSiteApAutoPlacementRequest) Execute() (*ResponseAutoPlacementInfo, *http.Response, error) {
	return r.ApiService.GetSiteApAutoPlacementExecute(r)
}

/*
GetSiteApAutoPlacement getSiteApAutoplacement

This API is called to view the current status of auto placement for a given map.


#### Status Descriptions

| Status | Description |
| --- | --- |
| `pending` | Autoplacement has not been requested for this map |
| `inprogress` | Autoplacement is currently processing |
| `done` | The autoplacement process has completed |
| `data_needed` | Additional position data is required for autoplacement. Users should verify the requested anchor APs have a position on the map |
| `invalid_model` | Autoplacement is not supported on the model of the APs on the map |
| `invalid_version` | Autoplacement is not supported with the APs current firmware version |
| `error` | There was an error in the autoplacement process |

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param mapId
 @return ApiGetSiteApAutoPlacementRequest
*/
func (a *SitesMapsAutoPlacementAPIService) GetSiteApAutoPlacement(ctx context.Context, siteId string, mapId string) ApiGetSiteApAutoPlacementRequest {
	return ApiGetSiteApAutoPlacementRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		mapId: mapId,
	}
}

// Execute executes the request
//  @return ResponseAutoPlacementInfo
func (a *SitesMapsAutoPlacementAPIService) GetSiteApAutoPlacementExecute(r ApiGetSiteApAutoPlacementRequest) (*ResponseAutoPlacementInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseAutoPlacementInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesMapsAutoPlacementAPIService.GetSiteApAutoPlacement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/maps/{map_id}/auto_placement"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_id"+"}", url.PathEscape(parameterValueToString(r.mapId, "mapId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiRunSiteApAutoplacementRequest struct {
	ctx context.Context
	ApiService SitesMapsAutoPlacementAPI
	siteId string
	mapId string
	autoPlacement *AutoPlacement
}

func (r ApiRunSiteApAutoplacementRequest) AutoPlacement(autoPlacement AutoPlacement) ApiRunSiteApAutoplacementRequest {
	r.autoPlacement = &autoPlacement
	return r
}

func (r ApiRunSiteApAutoplacementRequest) Execute() (*http.Response, error) {
	return r.ApiService.RunSiteApAutoplacementExecute(r)
}

/*
RunSiteApAutoplacement runSiteApAutoplacement

This API is called to trigger a map for auto placement. For auto placement feature to work, RTT-FTM data need to be collected from the APs on the map. This scan is disruptive and therefore the user must be notified of service disrution during the functioning of auto placement Repeated POST to this endpoint while a map is still running will be rejected.

List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide autoplacement suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param mapId
 @return ApiRunSiteApAutoplacementRequest
*/
func (a *SitesMapsAutoPlacementAPIService) RunSiteApAutoplacement(ctx context.Context, siteId string, mapId string) ApiRunSiteApAutoplacementRequest {
	return ApiRunSiteApAutoplacementRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		mapId: mapId,
	}
}

// Execute executes the request
func (a *SitesMapsAutoPlacementAPIService) RunSiteApAutoplacementExecute(r ApiRunSiteApAutoplacementRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesMapsAutoPlacementAPIService.RunSiteApAutoplacement")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/maps/{map_id}/auto_placement"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_id"+"}", url.PathEscape(parameterValueToString(r.mapId, "mapId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.autoPlacement
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
