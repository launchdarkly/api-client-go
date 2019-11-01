# \UsageApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvaluations**](UsageApi.md#GetEvaluations) | **Get** /usage/evaluations/{envId}/{flagKey} | [BETA] Get events usage by event id and the feature flag key.
[**GetEvent**](UsageApi.md#GetEvent) | **Get** /usage/events/{type} | [BETA] Get events usage by event type.
[**GetEvents**](UsageApi.md#GetEvents) | **Get** /usage/events | [BETA] Get events usage endpoints.
[**GetMAU**](UsageApi.md#GetMAU) | **Get** /usage/mau | [BETA] Get monthly active user data.
[**GetMAUByCategory**](UsageApi.md#GetMAUByCategory) | **Get** /usage/mau/bycategory | [BETA] Get monthly active user data by category.
[**GetStream**](UsageApi.md#GetStream) | **Get** /usage/streams/{source} | [BETA] Get a stream endpoint and return timeseries data.
[**GetStreamBySDK**](UsageApi.md#GetStreamBySDK) | **Get** /usage/streams/{source}/bysdkversion | [BETA] Get a stream timeseries data by source show sdk version metadata.
[**GetStreamSDKVersion**](UsageApi.md#GetStreamSDKVersion) | **Get** /usage/streams/{source}/sdkversions | [BETA] Get a stream timeseries data by source and show all sdk version associated.
[**GetStreams**](UsageApi.md#GetStreams) | **Get** /usage/streams | [BETA] Returns a list of all streams.
[**GetUsage**](UsageApi.md#GetUsage) | **Get** /usage | [BETA] Returns of the usage endpoints available.


# **GetEvaluations**
> StreamSdkVersion GetEvaluations(ctx, envId, flagKey)
[BETA] Get events usage by event id and the feature flag key.

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
[BETA] Get events usage by event type.

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
[BETA] Get events usage endpoints.

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
[BETA] Get monthly active user data.

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
[BETA] Get monthly active user data by category.

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
[BETA] Get a stream endpoint and return timeseries data.

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
[BETA] Get a stream timeseries data by source show sdk version metadata.

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
[BETA] Get a stream timeseries data by source and show all sdk version associated.

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
[BETA] Returns a list of all streams.

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
[BETA] Returns of the usage endpoints available.

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

