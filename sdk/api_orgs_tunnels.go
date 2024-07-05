/*
Mist API

> Version: **2406.1.17** > > Date: **July 5, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.17
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


type OrgsTunnelsAPI interface {

	/*
	CountOrgTunnelsStats countOrgTunnelsStats

	Count Mist Tunnels Stats

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId
	@return ApiCountOrgTunnelsStatsRequest
	*/
	CountOrgTunnelsStats(ctx context.Context, orgId string) ApiCountOrgTunnelsStatsRequest

	// CountOrgTunnelsStatsExecute executes the request
	//  @return RepsonseCount
	CountOrgTunnelsStatsExecute(r ApiCountOrgTunnelsStatsRequest) (*RepsonseCount, *http.Response, error)

	/*
	SearchOrgTunnelsStats searchOrgTunnelsStats

	Search Org Tunnels Stats

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId
	@return ApiSearchOrgTunnelsStatsRequest
	*/
	SearchOrgTunnelsStats(ctx context.Context, orgId string) ApiSearchOrgTunnelsStatsRequest

	// SearchOrgTunnelsStatsExecute executes the request
	//  @return ResponseTunnelSearch
	SearchOrgTunnelsStatsExecute(r ApiSearchOrgTunnelsStatsRequest) (*ResponseTunnelSearch, *http.Response, error)
}

// OrgsTunnelsAPIService OrgsTunnelsAPI service
type OrgsTunnelsAPIService service

type ApiCountOrgTunnelsStatsRequest struct {
	ctx context.Context
	ApiService OrgsTunnelsAPI
	orgId string
	distinct *OrgTunnelCountDistinct
	type_ *OrgTunnelTypeCount
}

// - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;: wxtunnel_id / ap / remote_ip / remote_port / state / mxedge_id / mxcluster_id / site_id / peer_mxedge_id; default is wxtunnel_id  - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;: mac / site_id / node / peer_ip / peer_host/ ip / tunnel_name / protocol / auth_algo / encrypt_algo / ike_version / last_event / up
func (r ApiCountOrgTunnelsStatsRequest) Distinct(distinct OrgTunnelCountDistinct) ApiCountOrgTunnelsStatsRequest {
	r.distinct = &distinct
	return r
}

func (r ApiCountOrgTunnelsStatsRequest) Type_(type_ OrgTunnelTypeCount) ApiCountOrgTunnelsStatsRequest {
	r.type_ = &type_
	return r
}

func (r ApiCountOrgTunnelsStatsRequest) Execute() (*RepsonseCount, *http.Response, error) {
	return r.ApiService.CountOrgTunnelsStatsExecute(r)
}

/*
CountOrgTunnelsStats countOrgTunnelsStats

Count Mist Tunnels Stats

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @return ApiCountOrgTunnelsStatsRequest
*/
func (a *OrgsTunnelsAPIService) CountOrgTunnelsStats(ctx context.Context, orgId string) ApiCountOrgTunnelsStatsRequest {
	return ApiCountOrgTunnelsStatsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return RepsonseCount
func (a *OrgsTunnelsAPIService) CountOrgTunnelsStatsExecute(r ApiCountOrgTunnelsStatsRequest) (*RepsonseCount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepsonseCount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsTunnelsAPIService.CountOrgTunnelsStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/orgs/{org_id}/stats/tunnels/count"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.distinct != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "distinct", r.distinct, "")
	} else {
		var defaultValue OrgTunnelCountDistinct = "wxtunnel_id"
		r.distinct = &defaultValue
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	} else {
		var defaultValue OrgTunnelTypeCount = "wxtunnel"
		r.type_ = &defaultValue
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

type ApiSearchOrgTunnelsStatsRequest struct {
	ctx context.Context
	ApiService OrgsTunnelsAPI
	orgId string
	mxclusterId *string
	siteId *string
	wxtunnelId *string
	ap *string
	mac *string
	node *string
	peerIp *string
	peerHost *string
	ip *string
	tunnelName *string
	protocol *string
	authAlgo *string
	encryptAlgo *string
	ikeVersion *string
	up *string
	type_ *TunnelType
	limit *int32
	start *int32
	end *int32
	duration *string
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) MxclusterId(mxclusterId string) ApiSearchOrgTunnelsStatsRequest {
	r.mxclusterId = &mxclusterId
	return r
}

func (r ApiSearchOrgTunnelsStatsRequest) SiteId(siteId string) ApiSearchOrgTunnelsStatsRequest {
	r.siteId = &siteId
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) WxtunnelId(wxtunnelId string) ApiSearchOrgTunnelsStatsRequest {
	r.wxtunnelId = &wxtunnelId
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) Ap(ap string) ApiSearchOrgTunnelsStatsRequest {
	r.ap = &ap
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) Mac(mac string) ApiSearchOrgTunnelsStatsRequest {
	r.mac = &mac
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) Node(node string) ApiSearchOrgTunnelsStatsRequest {
	r.node = &node
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) PeerIp(peerIp string) ApiSearchOrgTunnelsStatsRequest {
	r.peerIp = &peerIp
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) PeerHost(peerHost string) ApiSearchOrgTunnelsStatsRequest {
	r.peerHost = &peerHost
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) Ip(ip string) ApiSearchOrgTunnelsStatsRequest {
	r.ip = &ip
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) TunnelName(tunnelName string) ApiSearchOrgTunnelsStatsRequest {
	r.tunnelName = &tunnelName
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) Protocol(protocol string) ApiSearchOrgTunnelsStatsRequest {
	r.protocol = &protocol
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) AuthAlgo(authAlgo string) ApiSearchOrgTunnelsStatsRequest {
	r.authAlgo = &authAlgo
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) EncryptAlgo(encryptAlgo string) ApiSearchOrgTunnelsStatsRequest {
	r.encryptAlgo = &encryptAlgo
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) IkeVersion(ikeVersion string) ApiSearchOrgTunnelsStatsRequest {
	r.ikeVersion = &ikeVersion
	return r
}

// if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;
func (r ApiSearchOrgTunnelsStatsRequest) Up(up string) ApiSearchOrgTunnelsStatsRequest {
	r.up = &up
	return r
}

func (r ApiSearchOrgTunnelsStatsRequest) Type_(type_ TunnelType) ApiSearchOrgTunnelsStatsRequest {
	r.type_ = &type_
	return r
}

func (r ApiSearchOrgTunnelsStatsRequest) Limit(limit int32) ApiSearchOrgTunnelsStatsRequest {
	r.limit = &limit
	return r
}

// start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified
func (r ApiSearchOrgTunnelsStatsRequest) Start(start int32) ApiSearchOrgTunnelsStatsRequest {
	r.start = &start
	return r
}

// end datetime, can be epoch or relative time like -1d, -2h; now if not specified
func (r ApiSearchOrgTunnelsStatsRequest) End(end int32) ApiSearchOrgTunnelsStatsRequest {
	r.end = &end
	return r
}

// duration like 7d, 2w
func (r ApiSearchOrgTunnelsStatsRequest) Duration(duration string) ApiSearchOrgTunnelsStatsRequest {
	r.duration = &duration
	return r
}

func (r ApiSearchOrgTunnelsStatsRequest) Execute() (*ResponseTunnelSearch, *http.Response, error) {
	return r.ApiService.SearchOrgTunnelsStatsExecute(r)
}

/*
SearchOrgTunnelsStats searchOrgTunnelsStats

Search Org Tunnels Stats

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @return ApiSearchOrgTunnelsStatsRequest
*/
func (a *OrgsTunnelsAPIService) SearchOrgTunnelsStats(ctx context.Context, orgId string) ApiSearchOrgTunnelsStatsRequest {
	return ApiSearchOrgTunnelsStatsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return ResponseTunnelSearch
func (a *OrgsTunnelsAPIService) SearchOrgTunnelsStatsExecute(r ApiSearchOrgTunnelsStatsRequest) (*ResponseTunnelSearch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseTunnelSearch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsTunnelsAPIService.SearchOrgTunnelsStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/orgs/{org_id}/stats/tunnels/search"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mxclusterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mxcluster_id", r.mxclusterId, "")
	}
	if r.siteId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId, "")
	}
	if r.wxtunnelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "wxtunnel_id", r.wxtunnelId, "")
	}
	if r.ap != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ap", r.ap, "")
	}
	if r.mac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mac", r.mac, "")
	}
	if r.node != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "node", r.node, "")
	}
	if r.peerIp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "peer_ip", r.peerIp, "")
	}
	if r.peerHost != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "peer_host", r.peerHost, "")
	}
	if r.ip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ip", r.ip, "")
	}
	if r.tunnelName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tunnel_name", r.tunnelName, "")
	}
	if r.protocol != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "protocol", r.protocol, "")
	}
	if r.authAlgo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auth_algo", r.authAlgo, "")
	}
	if r.encryptAlgo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "encrypt_algo", r.encryptAlgo, "")
	}
	if r.ikeVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ike_version", r.ikeVersion, "")
	}
	if r.up != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "up", r.up, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	} else {
		var defaultValue TunnelType = "wxtunnel"
		r.type_ = &defaultValue
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
