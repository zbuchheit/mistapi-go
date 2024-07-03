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


type OrgsMarvisAPI interface {

	/*
	TroubleshootOrg troubleshootOrg

	Troubleshoot sites, devices, clients, and wired clientsfor maximum of last 7 days from current time. See search APIs for device information:
- [search Device]($e/Orgs%20Devices/searchOrgDevices)
- [search Wireless Client]($e/Orgs%20Clients%20-%20Wireless/searchOrgWirelessClients)
- [search Wired Client]($e/Orgs%20Clients%20-%20Wired/searchOrgWiredClients)
- [search Wan Client]($e/Orgs%20Clients%20-%20Wan/searchOrgWanClients)

**NOTE**: requires Marvis subscription license

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId
	@return ApiTroubleshootOrgRequest
	*/
	TroubleshootOrg(ctx context.Context, orgId string) ApiTroubleshootOrgRequest

	// TroubleshootOrgExecute executes the request
	//  @return ResponseTroubleshoot
	TroubleshootOrgExecute(r ApiTroubleshootOrgRequest) (*ResponseTroubleshoot, *http.Response, error)
}

// OrgsMarvisAPIService OrgsMarvisAPI service
type OrgsMarvisAPIService service

type ApiTroubleshootOrgRequest struct {
	ctx context.Context
	ApiService OrgsMarvisAPI
	orgId string
	mac *string
	siteId *string
	start *int32
	end *int32
	type_ *TroubleshootType
}

// **required** when troubleshooting device or a client
func (r ApiTroubleshootOrgRequest) Mac(mac string) ApiTroubleshootOrgRequest {
	r.mac = &mac
	return r
}

// **required** when troubleshooting site
func (r ApiTroubleshootOrgRequest) SiteId(siteId string) ApiTroubleshootOrgRequest {
	r.siteId = &siteId
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiTroubleshootOrgRequest) Start(start int32) ApiTroubleshootOrgRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiTroubleshootOrgRequest) End(end int32) ApiTroubleshootOrgRequest {
	r.end = &end
	return r
}

// when troubleshooting site, type of network to troubleshoot
func (r ApiTroubleshootOrgRequest) Type_(type_ TroubleshootType) ApiTroubleshootOrgRequest {
	r.type_ = &type_
	return r
}

func (r ApiTroubleshootOrgRequest) Execute() (*ResponseTroubleshoot, *http.Response, error) {
	return r.ApiService.TroubleshootOrgExecute(r)
}

/*
TroubleshootOrg troubleshootOrg

Troubleshoot sites, devices, clients, and wired clientsfor maximum of last 7 days from current time. See search APIs for device information:
- [search Device]($e/Orgs%20Devices/searchOrgDevices)
- [search Wireless Client]($e/Orgs%20Clients%20-%20Wireless/searchOrgWirelessClients)
- [search Wired Client]($e/Orgs%20Clients%20-%20Wired/searchOrgWiredClients)
- [search Wan Client]($e/Orgs%20Clients%20-%20Wan/searchOrgWanClients)

**NOTE**: requires Marvis subscription license

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @return ApiTroubleshootOrgRequest
*/
func (a *OrgsMarvisAPIService) TroubleshootOrg(ctx context.Context, orgId string) ApiTroubleshootOrgRequest {
	return ApiTroubleshootOrgRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return ResponseTroubleshoot
func (a *OrgsMarvisAPIService) TroubleshootOrgExecute(r ApiTroubleshootOrgRequest) (*ResponseTroubleshoot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseTroubleshoot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsMarvisAPIService.TroubleshootOrg")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/orgs/{org_id}/troubleshoot"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mac", r.mac, "")
	}
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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
