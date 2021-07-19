/*
 * LaunchDarkly REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0
 * Contact: support@launchdarkly.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ldapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// EnvironmentsApiService EnvironmentsApi service
type EnvironmentsApiService service

type ApiDeleteEnvironmentRequest struct {
	ctx _context.Context
	ApiService *EnvironmentsApiService
	projectKey string
	environmentKey string
}


func (r ApiDeleteEnvironmentRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteEnvironmentExecute(r)
}

/*
 * DeleteEnvironment Delete environment by key
 *  Delete a environment by key.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectKey The project key
 * @param environmentKey The environment key
 * @return ApiDeleteEnvironmentRequest
 */
func (a *EnvironmentsApiService) DeleteEnvironment(ctx _context.Context, projectKey string, environmentKey string) ApiDeleteEnvironmentRequest {
	return ApiDeleteEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		environmentKey: environmentKey,
	}
}

/*
 * Execute executes the request
 */
func (a *EnvironmentsApiService) DeleteEnvironmentExecute(r ApiDeleteEnvironmentRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.DeleteEnvironment")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/environments/{environmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", _neturl.PathEscape(parameterToString(r.projectKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", _neturl.PathEscape(parameterToString(r.environmentKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetEnvironmentRequest struct {
	ctx _context.Context
	ApiService *EnvironmentsApiService
	projectKey string
	environmentKey string
}


func (r ApiGetEnvironmentRequest) Execute() (EnvironmentRep, *_nethttp.Response, error) {
	return r.ApiService.GetEnvironmentExecute(r)
}

/*
 * GetEnvironment Get environment by key
 *  Fetch a single environment by key.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectKey The project key
 * @param environmentKey The environment key
 * @return ApiGetEnvironmentRequest
 */
func (a *EnvironmentsApiService) GetEnvironment(ctx _context.Context, projectKey string, environmentKey string) ApiGetEnvironmentRequest {
	return ApiGetEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		environmentKey: environmentKey,
	}
}

/*
 * Execute executes the request
 * @return EnvironmentRep
 */
func (a *EnvironmentsApiService) GetEnvironmentExecute(r ApiGetEnvironmentRequest) (EnvironmentRep, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EnvironmentRep
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.GetEnvironment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/environments/{environmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", _neturl.PathEscape(parameterToString(r.projectKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", _neturl.PathEscape(parameterToString(r.environmentKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
			if apiKey, ok := auth["ApiKey"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchEnvironmentRequest struct {
	ctx _context.Context
	ApiService *EnvironmentsApiService
	projectKey string
	environmentKey string
	jSONPatchElt *[]JSONPatchElt
}

func (r ApiPatchEnvironmentRequest) JSONPatchElt(jSONPatchElt []JSONPatchElt) ApiPatchEnvironmentRequest {
	r.jSONPatchElt = &jSONPatchElt
	return r
}

func (r ApiPatchEnvironmentRequest) Execute() (EnvironmentRep, *_nethttp.Response, error) {
	return r.ApiService.PatchEnvironmentExecute(r)
}

/*
 * PatchEnvironment Patch environment by key
 *  Patch a environment by key.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectKey The project key
 * @param environmentKey The environment key
 * @return ApiPatchEnvironmentRequest
 */
func (a *EnvironmentsApiService) PatchEnvironment(ctx _context.Context, projectKey string, environmentKey string) ApiPatchEnvironmentRequest {
	return ApiPatchEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		environmentKey: environmentKey,
	}
}

/*
 * Execute executes the request
 * @return EnvironmentRep
 */
func (a *EnvironmentsApiService) PatchEnvironmentExecute(r ApiPatchEnvironmentRequest) (EnvironmentRep, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EnvironmentRep
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.PatchEnvironment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/environments/{environmentKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", _neturl.PathEscape(parameterToString(r.projectKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentKey"+"}", _neturl.PathEscape(parameterToString(r.environmentKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.jSONPatchElt == nil {
		return localVarReturnValue, nil, reportError("jSONPatchElt is required and must be specified")
	}

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
	localVarPostBody = r.jSONPatchElt
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostEnvironmentRequest struct {
	ctx _context.Context
	ApiService *EnvironmentsApiService
	projectKey string
	environmentPost *EnvironmentPost
}

func (r ApiPostEnvironmentRequest) EnvironmentPost(environmentPost EnvironmentPost) ApiPostEnvironmentRequest {
	r.environmentPost = &environmentPost
	return r
}

func (r ApiPostEnvironmentRequest) Execute() (EnvironmentRep, *_nethttp.Response, error) {
	return r.ApiService.PostEnvironmentExecute(r)
}

/*
 * PostEnvironment Create environment
 *  Create a new environment with the given key and name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectKey The project key
 * @return ApiPostEnvironmentRequest
 */
func (a *EnvironmentsApiService) PostEnvironment(ctx _context.Context, projectKey string) ApiPostEnvironmentRequest {
	return ApiPostEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

/*
 * Execute executes the request
 * @return EnvironmentRep
 */
func (a *EnvironmentsApiService) PostEnvironmentExecute(r ApiPostEnvironmentRequest) (EnvironmentRep, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EnvironmentRep
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.PostEnvironment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/projects/{projectKey}/environments"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", _neturl.PathEscape(parameterToString(r.projectKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environmentPost == nil {
		return localVarReturnValue, nil, reportError("environmentPost is required and must be specified")
	}

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
	localVarPostBody = r.environmentPost
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
