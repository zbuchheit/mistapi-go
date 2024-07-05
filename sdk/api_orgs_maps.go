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
	"os"
)


type OrgsMapsAPI interface {

	/*
	ImportOrgMapToSite importOrgMapToSite

	Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches

#### Request

```
"json": a JSON string describing your upload
"file": a binary file
```

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId
	@param siteName
	@return ApiImportOrgMapToSiteRequest
	*/
	ImportOrgMapToSite(ctx context.Context, orgId string, siteName string) ApiImportOrgMapToSiteRequest

	// ImportOrgMapToSiteExecute executes the request
	//  @return ResponseMapImport
	ImportOrgMapToSiteExecute(r ApiImportOrgMapToSiteRequest) (*ResponseMapImport, *http.Response, error)

	/*
	ImportOrgMaps importOrgMaps

	Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches

### CSV File Format
```csv
Vendor AP name,Mist AP Mac
US Office AP-2 - 5c:5b:35:00:00:02,5c5b35000002
```

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId
	@return ApiImportOrgMapsRequest
	*/
	ImportOrgMaps(ctx context.Context, orgId string) ApiImportOrgMapsRequest

	// ImportOrgMapsExecute executes the request
	//  @return ResponseMapImport
	ImportOrgMapsExecute(r ApiImportOrgMapsRequest) (*ResponseMapImport, *http.Response, error)
}

// OrgsMapsAPIService OrgsMapsAPI service
type OrgsMapsAPIService service

type ApiImportOrgMapToSiteRequest struct {
	ctx context.Context
	ApiService OrgsMapsAPI
	orgId string
	siteName string
	autoDeviceprofileAssignment *bool
	csv *os.File
	file *os.File
	json *MapImportJson
}

// whether to auto assign device to deviceprofile by name
func (r ApiImportOrgMapToSiteRequest) AutoDeviceprofileAssignment(autoDeviceprofileAssignment bool) ApiImportOrgMapToSiteRequest {
	r.autoDeviceprofileAssignment = &autoDeviceprofileAssignment
	return r
}

// csv file for ap name mapping, optional
func (r ApiImportOrgMapToSiteRequest) Csv(csv *os.File) ApiImportOrgMapToSiteRequest {
	r.csv = csv
	return r
}

// ekahau or ibwave file
func (r ApiImportOrgMapToSiteRequest) File(file *os.File) ApiImportOrgMapToSiteRequest {
	r.file = file
	return r
}

func (r ApiImportOrgMapToSiteRequest) Json(json MapImportJson) ApiImportOrgMapToSiteRequest {
	r.json = &json
	return r
}

func (r ApiImportOrgMapToSiteRequest) Execute() (*ResponseMapImport, *http.Response, error) {
	return r.ApiService.ImportOrgMapToSiteExecute(r)
}

/*
ImportOrgMapToSite importOrgMapToSite

Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches

#### Request

```
"json": a JSON string describing your upload
"file": a binary file
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @param siteName
 @return ApiImportOrgMapToSiteRequest
*/
func (a *OrgsMapsAPIService) ImportOrgMapToSite(ctx context.Context, orgId string, siteName string) ApiImportOrgMapToSiteRequest {
	return ApiImportOrgMapToSiteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		siteName: siteName,
	}
}

// Execute executes the request
//  @return ResponseMapImport
func (a *OrgsMapsAPIService) ImportOrgMapToSiteExecute(r ApiImportOrgMapToSiteRequest) (*ResponseMapImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseMapImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsMapsAPIService.ImportOrgMapToSite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/orgs/{org_id}/sites/{site_name}/maps/import"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"site_name"+"}", url.PathEscape(parameterValueToString(r.siteName, "siteName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form_data"}

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
	if r.autoDeviceprofileAssignment != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "auto_deviceprofile_assignment", r.autoDeviceprofileAssignment, "")
	}
	var csvLocalVarFormFileName string
	var csvLocalVarFileName     string
	var csvLocalVarFileBytes    []byte

	csvLocalVarFormFileName = "csv"
	csvLocalVarFile := r.csv

	if csvLocalVarFile != nil {
		fbs, _ := io.ReadAll(csvLocalVarFile)

		csvLocalVarFileBytes = fbs
		csvLocalVarFileName = csvLocalVarFile.Name()
		csvLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: csvLocalVarFileBytes, fileName: csvLocalVarFileName, formFileName: csvLocalVarFormFileName})
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.json != nil {
		paramJson, err := parameterToJson(*r.json)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("json", paramJson)
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

type ApiImportOrgMapsRequest struct {
	ctx context.Context
	ApiService OrgsMapsAPI
	orgId string
	autoDeviceprofileAssignment *bool
	csv *os.File
	file *os.File
	json *MapOrgImportFileJson
}

// whether to auto assign device to deviceprofile by name
func (r ApiImportOrgMapsRequest) AutoDeviceprofileAssignment(autoDeviceprofileAssignment bool) ApiImportOrgMapsRequest {
	r.autoDeviceprofileAssignment = &autoDeviceprofileAssignment
	return r
}

// csv file for ap name mapping, optional
func (r ApiImportOrgMapsRequest) Csv(csv *os.File) ApiImportOrgMapsRequest {
	r.csv = csv
	return r
}

// ekahau or ibwave file
func (r ApiImportOrgMapsRequest) File(file *os.File) ApiImportOrgMapsRequest {
	r.file = file
	return r
}

func (r ApiImportOrgMapsRequest) Json(json MapOrgImportFileJson) ApiImportOrgMapsRequest {
	r.json = &json
	return r
}

func (r ApiImportOrgMapsRequest) Execute() (*ResponseMapImport, *http.Response, error) {
	return r.ApiService.ImportOrgMapsExecute(r)
}

/*
ImportOrgMaps importOrgMaps

Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches

### CSV File Format
```csv
Vendor AP name,Mist AP Mac
US Office AP-2 - 5c:5b:35:00:00:02,5c5b35000002
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @return ApiImportOrgMapsRequest
*/
func (a *OrgsMapsAPIService) ImportOrgMaps(ctx context.Context, orgId string) ApiImportOrgMapsRequest {
	return ApiImportOrgMapsRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return ResponseMapImport
func (a *OrgsMapsAPIService) ImportOrgMapsExecute(r ApiImportOrgMapsRequest) (*ResponseMapImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseMapImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgsMapsAPIService.ImportOrgMaps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/orgs/{org_id}/maps/import"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form_data"}

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
	if r.autoDeviceprofileAssignment != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "auto_deviceprofile_assignment", r.autoDeviceprofileAssignment, "")
	}
	var csvLocalVarFormFileName string
	var csvLocalVarFileName     string
	var csvLocalVarFileBytes    []byte

	csvLocalVarFormFileName = "csv"
	csvLocalVarFile := r.csv

	if csvLocalVarFile != nil {
		fbs, _ := io.ReadAll(csvLocalVarFile)

		csvLocalVarFileBytes = fbs
		csvLocalVarFileName = csvLocalVarFile.Name()
		csvLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: csvLocalVarFileBytes, fileName: csvLocalVarFileName, formFileName: csvLocalVarFormFileName})
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.json != nil {
		paramJson, err := parameterToJson(*r.json)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("json", paramJson)
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
