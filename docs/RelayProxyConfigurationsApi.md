# \RelayProxyConfigurationsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRelayProxyConfig**](RelayProxyConfigurationsApi.md#DeleteRelayProxyConfig) | **Delete** /account/relay-auto-configs/{id} | Delete a relay proxy configuration by ID.
[**GetRelayProxyConfig**](RelayProxyConfigurationsApi.md#GetRelayProxyConfig) | **Get** /account/relay-auto-configs/{id} | Get a single relay proxy configuration by ID.
[**GetRelayProxyConfigs**](RelayProxyConfigurationsApi.md#GetRelayProxyConfigs) | **Get** /account/relay-auto-configs | Returns a list of relay proxy configurations in the account.
[**PatchRelayProxyConfig**](RelayProxyConfigurationsApi.md#PatchRelayProxyConfig) | **Patch** /account/relay-auto-configs/{id} | Modify a relay proxy configuration by ID.
[**PostRelayAutoConfig**](RelayProxyConfigurationsApi.md#PostRelayAutoConfig) | **Post** /account/relay-auto-configs | Create a new relay proxy config.
[**ResetRelayProxyConfig**](RelayProxyConfigurationsApi.md#ResetRelayProxyConfig) | **Post** /account/relay-auto-configs/{id}/reset | Reset a relay proxy configuration&#39;s secret key with an optional expiry time for the old key.


# **DeleteRelayProxyConfig**
> DeleteRelayProxyConfig(ctx, id)
Delete a relay proxy configuration by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The relay proxy configuration ID | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelayProxyConfig**
> RelayProxyConfig GetRelayProxyConfig(ctx, id)
Get a single relay proxy configuration by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The relay proxy configuration ID | 

### Return type

[**RelayProxyConfig**](RelayProxyConfig.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelayProxyConfigs**
> RelayProxyConfigs GetRelayProxyConfigs(ctx, )
Returns a list of relay proxy configurations in the account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RelayProxyConfigs**](RelayProxyConfigs.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRelayProxyConfig**
> RelayProxyConfig PatchRelayProxyConfig(ctx, id, patchDelta)
Modify a relay proxy configuration by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The relay proxy configuration ID | 
  **patchDelta** | [**[]PatchOperation**](PatchOperation.md)| Requires a JSON Patch representation of the desired changes to the project. &#39;http://jsonpatch.com/&#39; | 

### Return type

[**RelayProxyConfig**](RelayProxyConfig.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostRelayAutoConfig**
> RelayProxyConfig PostRelayAutoConfig(ctx, relayProxyConfigBody)
Create a new relay proxy config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **relayProxyConfigBody** | [**RelayProxyConfigBody**](RelayProxyConfigBody.md)| Create a new relay proxy configuration | 

### Return type

[**RelayProxyConfig**](RelayProxyConfig.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetRelayProxyConfig**
> RelayProxyConfig ResetRelayProxyConfig(ctx, id, optional)
Reset a relay proxy configuration's secret key with an optional expiry time for the old key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The relay proxy configuration ID | 
 **optional** | ***RelayProxyConfigurationsApiResetRelayProxyConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RelayProxyConfigurationsApiResetRelayProxyConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiry** | **optional.Int64**| An expiration time for the old relay proxy configuration key, expressed as a Unix epoch time in milliseconds. By default, the relay proxy configuration will expire immediately | 

### Return type

[**RelayProxyConfig**](RelayProxyConfig.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

