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

// UserSettingsApiService UserSettingsApi service
type UserSettingsApiService service

type ApiGetUserFlagSettingRequest struct {
	ctx _context.Context
	ApiService *UserSettingsApiService
	projKey string
	envKey string
	key string
	featureKey string
}


func (r ApiGetUserFlagSettingRequest) Execute() (UserSettingRep, *_nethttp.Response, error) {
	return r.ApiService.GetUserFlagSettingExecute(r)
}

/*
 * GetUserFlagSetting Get flag setting for user
 *  Fetches a single flag setting for a user by key. The most important attribute in the response is the _value. The _value is the current setting that the user sees. For a boolean feature toggle, this is `true`, `false`, or `null`. `Null` returns if there is no defined fallthrough value. The example response indicates that the user Abbie_Braun has the `sort.order` flag enabled and the `alternate.page` flag disabled. The setting attribute indicates whether you've explicitly targeted a user to receive a particular variation. For example, if you have turned off a feature flag for a user, this setting will be `false`. A setting of `null` means that you haven't assigned that user to a specific variation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projKey The project key
 * @param envKey The environment key
 * @param key The user key
 * @param featureKey The feature flag key
 * @return ApiGetUserFlagSettingRequest
 */
func (a *UserSettingsApiService) GetUserFlagSetting(ctx _context.Context, projKey string, envKey string, key string, featureKey string) ApiGetUserFlagSettingRequest {
	return ApiGetUserFlagSettingRequest{
		ApiService: a,
		ctx: ctx,
		projKey: projKey,
		envKey: envKey,
		key: key,
		featureKey: featureKey,
	}
}

/*
 * Execute executes the request
 * @return UserSettingRep
 */
func (a *UserSettingsApiService) GetUserFlagSettingExecute(r ApiGetUserFlagSettingRequest) (UserSettingRep, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserSettingRep
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserSettingsApiService.GetUserFlagSetting")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projKey"+"}", _neturl.PathEscape(parameterToString(r.projKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envKey"+"}", _neturl.PathEscape(parameterToString(r.envKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", _neturl.PathEscape(parameterToString(r.key, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"featureKey"+"}", _neturl.PathEscape(parameterToString(r.featureKey, "")), -1)

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

type ApiGetUserFlagSettingsRequest struct {
	ctx _context.Context
	ApiService *UserSettingsApiService
	projKey string
	envKey string
	key string
}


func (r ApiGetUserFlagSettingsRequest) Execute() (UserSettingsCollection, *_nethttp.Response, error) {
	return r.ApiService.GetUserFlagSettingsExecute(r)
}

/*
 * GetUserFlagSettings List flag settings for user
 *  Lists the current flag settings for a given user. The most important attribute in the response is the _value. The _value is the setting that the user sees. For a boolean feature toggle, this is `true`, `false`, or `null`. `Null` returns if there is no defined fallthrough value. The example response indicates that the user Abbie_Braun has the `sort.order` flag enabled and the `alternate.page` flag disabled. The setting attribute indicates whether you've explicitly targeted a user to receive a particular variation. For example, if you have turned off a feature flag for a user, this setting will be `false`. A setting of `null` means that you haven't assigned that user to a specific variation.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projKey The project key
 * @param envKey The environment key
 * @param key The user key
 * @return ApiGetUserFlagSettingsRequest
 */
func (a *UserSettingsApiService) GetUserFlagSettings(ctx _context.Context, projKey string, envKey string, key string) ApiGetUserFlagSettingsRequest {
	return ApiGetUserFlagSettingsRequest{
		ApiService: a,
		ctx: ctx,
		projKey: projKey,
		envKey: envKey,
		key: key,
	}
}

/*
 * Execute executes the request
 * @return UserSettingsCollection
 */
func (a *UserSettingsApiService) GetUserFlagSettingsExecute(r ApiGetUserFlagSettingsRequest) (UserSettingsCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UserSettingsCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserSettingsApiService.GetUserFlagSettings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/users/{projKey}/{envKey}/{key}/flags"
	localVarPath = strings.Replace(localVarPath, "{"+"projKey"+"}", _neturl.PathEscape(parameterToString(r.projKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envKey"+"}", _neturl.PathEscape(parameterToString(r.envKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", _neturl.PathEscape(parameterToString(r.key, "")), -1)

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

type ApiPutFlagSettingRequest struct {
	ctx _context.Context
	ApiService *UserSettingsApiService
	projKey string
	envKey string
	key string
	featureKey string
	api2ValuePut *Api2ValuePut
}

func (r ApiPutFlagSettingRequest) Api2ValuePut(api2ValuePut Api2ValuePut) ApiPutFlagSettingRequest {
	r.api2ValuePut = &api2ValuePut
	return r
}

func (r ApiPutFlagSettingRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PutFlagSettingExecute(r)
}

/*
 * PutFlagSetting Update flag settings for user
 *  Specifically enable or disable a feature flag for a user based on their key.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projKey The project key
 * @param envKey The environment key
 * @param key The user key
 * @param featureKey The feature flag key
 * @return ApiPutFlagSettingRequest
 */
func (a *UserSettingsApiService) PutFlagSetting(ctx _context.Context, projKey string, envKey string, key string, featureKey string) ApiPutFlagSettingRequest {
	return ApiPutFlagSettingRequest{
		ApiService: a,
		ctx: ctx,
		projKey: projKey,
		envKey: envKey,
		key: key,
		featureKey: featureKey,
	}
}

/*
 * Execute executes the request
 */
func (a *UserSettingsApiService) PutFlagSettingExecute(r ApiPutFlagSettingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserSettingsApiService.PutFlagSetting")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"projKey"+"}", _neturl.PathEscape(parameterToString(r.projKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envKey"+"}", _neturl.PathEscape(parameterToString(r.envKey, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", _neturl.PathEscape(parameterToString(r.key, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"featureKey"+"}", _neturl.PathEscape(parameterToString(r.featureKey, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.api2ValuePut == nil {
		return nil, reportError("api2ValuePut is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.api2ValuePut
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
