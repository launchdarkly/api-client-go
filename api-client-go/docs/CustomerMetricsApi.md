# \CustomerMetricsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvaluations**](CustomerMetricsApi.md#GetEvaluations) | **Get** /usage/evaluations/{envId}/{flagKey} | Get events usage by event id and the feature flag key.
[**GetEvent**](CustomerMetricsApi.md#GetEvent) | **Get** /usage/events/{type} | Get events usage by event type.
[**GetEvents**](CustomerMetricsApi.md#GetEvents) | **Get** /usage/events | Get events usage endpoints.
[**GetMAU**](CustomerMetricsApi.md#GetMAU) | **Get** /usage/mau | Get monthly active user data.
[**GetMAUByCategory**](CustomerMetricsApi.md#GetMAUByCategory) | **Get** /usage/mau/bycategory | Get monthly active user data by category.
[**GetStream**](CustomerMetricsApi.md#GetStream) | **Get** /usage/streams/{source} | Get a stream endpoint and return timeseries data.
[**GetStreamBySDK**](CustomerMetricsApi.md#GetStreamBySDK) | **Get** /usage/streams/{source}/bysdkversion | Get a stream timeseries data by source show sdk version metadata.
[**GetStreamSDKVersion**](CustomerMetricsApi.md#GetStreamSDKVersion) | **Get** /usage/streams/{source}/sdkversions | Get a stream timeseries data by source and show all sdk version associated.
[**GetStreams**](CustomerMetricsApi.md#GetStreams) | **Get** /usage/streams | Returns a list of all streams.
[**GetUsage**](CustomerMetricsApi.md#GetUsage) | **Get** /usage | Returns of the usage endpoints available.


# **GetEvaluations**
> StreamSdkVersion GetEvaluations(ctx, envId, flagKey)
Get events usage by event id and the feature flag key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **string**| The environment id for the flag evaluations in question. | 
  **flagKey** | **string**| The key of the flag we want metrics for. | 

### Return type

[**StreamSdkVersion**](StreamSDKVersion.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvent**
> StreamSdkVersion GetEvent(ctx, type_)
Get events usage by event type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The type of event we would like to track. | 

### Return type

[**StreamSdkVersion**](StreamSDKVersion.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvents**
> Events GetEvents(ctx, )
Get events usage endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Events**](Events.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMAU**
> Mau GetMAU(ctx, )
Get monthly active user data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Mau**](MAU.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMAUByCategory**
> MaUbyCategory GetMAUByCategory(ctx, )
Get monthly active user data by category.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MaUbyCategory**](MAUbyCategory.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStream**
> Stream GetStream(ctx, source)
Get a stream endpoint and return timeseries data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**| The source of where the stream comes from. | 

### Return type

[**Stream**](Stream.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamBySDK**
> StreamBySdk GetStreamBySDK(ctx, source)
Get a stream timeseries data by source show sdk version metadata.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**| The source of where the stream comes from. | 

### Return type

[**StreamBySdk**](StreamBySDK.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreamSDKVersion**
> StreamSdkVersion GetStreamSDKVersion(ctx, source)
Get a stream timeseries data by source and show all sdk version associated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**| The source of where the stream comes from. | 

### Return type

[**StreamSdkVersion**](StreamSDKVersion.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStreams**
> Streams GetStreams(ctx, )
Returns a list of all streams.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Streams**](Streams.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsage**
> Usage GetUsage(ctx, )
Returns of the usage endpoints available.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Usage**](Usage.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

