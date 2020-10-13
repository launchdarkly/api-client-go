# \EnvironmentsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEnvironment**](EnvironmentsApi.md#DeleteEnvironment) | **Delete** /projects/{projectKey}/environments/{environmentKey} | Delete an environment in a specific project.
[**GetEnvironment**](EnvironmentsApi.md#GetEnvironment) | **Get** /projects/{projectKey}/environments/{environmentKey} | Get an environment given a project and key.
[**PatchEnvironment**](EnvironmentsApi.md#PatchEnvironment) | **Patch** /projects/{projectKey}/environments/{environmentKey} | Modify an environment by ID.
[**PostEnvironment**](EnvironmentsApi.md#PostEnvironment) | **Post** /projects/{projectKey}/environments | Create a new environment in a specified project with a given name, key, and swatch color.
[**ResetEnvironmentMobileKey**](EnvironmentsApi.md#ResetEnvironmentMobileKey) | **Post** /projects/{projectKey}/environments/{environmentKey}/mobileKey | Reset an environment&#39;s mobile key. The optional expiry for the old key is deprecated for this endpoint, so the old key will always expire immediately.
[**ResetEnvironmentSDKKey**](EnvironmentsApi.md#ResetEnvironmentSDKKey) | **Post** /projects/{projectKey}/environments/{environmentKey}/apiKey | Reset an environment&#39;s SDK key with an optional expiry time for the old key.


# **DeleteEnvironment**
> DeleteEnvironment(ctx, projectKey, environmentKey)
Delete an environment in a specific project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironment**
> Environment GetEnvironment(ctx, projectKey, environmentKey)
Get an environment given a project and key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchEnvironment**
> Environment PatchEnvironment(ctx, projectKey, environmentKey, patchDelta)
Modify an environment by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **patchDelta** | [**[]PatchOperation**](PatchOperation.md)| Requires a JSON Patch representation of the desired changes to the project. &#39;http://jsonpatch.com/&#39; | 

### Return type

[**Environment**](Environment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEnvironment**
> Environment PostEnvironment(ctx, projectKey, environmentBody)
Create a new environment in a specified project with a given name, key, and swatch color.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentBody** | [**EnvironmentPost**](EnvironmentPost.md)| New environment. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetEnvironmentMobileKey**
> Environment ResetEnvironmentMobileKey(ctx, projectKey, environmentKey, optional)
Reset an environment's mobile key. The optional expiry for the old key is deprecated for this endpoint, so the old key will always expire immediately.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
 **optional** | ***ResetEnvironmentMobileKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResetEnvironmentMobileKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expiry** | **optional.Int64**| The expiry parameter is deprecated for this endpoint, so the old mobile key will always expire immediately. This parameter will be removed in an upcoming major API client version. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetEnvironmentSDKKey**
> Environment ResetEnvironmentSDKKey(ctx, projectKey, environmentKey, optional)
Reset an environment's SDK key with an optional expiry time for the old key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
 **optional** | ***ResetEnvironmentSDKKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResetEnvironmentSDKKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expiry** | **optional.Int64**| An expiration time for the old environment SDK key, expressed as a Unix epoch time in milliseconds. By default, the key will expire immediately. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

