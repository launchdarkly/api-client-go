# \FlagsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeatureFlag**](FlagsApi.md#DeleteFeatureFlag) | **Delete** /flags/{projectKey}/{featureFlagKey} | Delete a feature flag by ID
[**GetFeatureFlag**](FlagsApi.md#GetFeatureFlag) | **Get** /flags/{projectKey}/{featureFlagKey} | Get a single feature flag by key.
[**GetFeatureFlagStatus**](FlagsApi.md#GetFeatureFlagStatus) | **Get** /flag-statuses/{projectKey}/{environmentKey} | Get a list of statuses for all feature flags
[**GetFeatureFlagStatuses**](FlagsApi.md#GetFeatureFlagStatuses) | **Get** /flag-statuses/{projectKey}/{environmentKey}/{featureFlagKey} | Get a list of statuses for all feature flags
[**GetFeatureFlags**](FlagsApi.md#GetFeatureFlags) | **Get** /flags/{projectKey} | Get a list of all features in the given project.
[**PatchFeatureFlag**](FlagsApi.md#PatchFeatureFlag) | **Patch** /flags/{projectKey}/{featureFlagKey} | Modify a feature flag by ID
[**PostFeatureFlag**](FlagsApi.md#PostFeatureFlag) | **Post** /flags/{projectKey}/{featureFlagKey} | Create a feature flag


# **DeleteFeatureFlag**
> DeleteFeatureFlag($projectKey, $featureFlagKey)

Delete a feature flag by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key | 
 **featureFlagKey** | **string**| The feature flags key | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlag**
> FeatureFlag GetFeatureFlag($projectKey, $featureFlagKey, $environmentKeyQuery)

Get a single feature flag by key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key | 
 **featureFlagKey** | **string**| The feature flags key | 
 **environmentKeyQuery** | **string**| The environment key | [optional] 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatus**
> FeatureFlagStatuses GetFeatureFlagStatus($projectKey, $environmentKey)

Get a list of statuses for all feature flags


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key | 
 **environmentKey** | **string**| The environment key | 

### Return type

[**FeatureFlagStatuses**](FeatureFlagStatuses.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatuses**
> FeatureFlagStatus GetFeatureFlagStatuses($projectKey, $environmentKey, $featureFlagKey)

Get a list of statuses for all feature flags


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key | 
 **environmentKey** | **string**| The environment key | 
 **featureFlagKey** | **string**| The feature flags key | 

### Return type

[**FeatureFlagStatus**](FeatureFlagStatus.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlags**
> FeatureFlags GetFeatureFlags($projectKey, $environmentKeyQuery, $tag)

Get a list of all features in the given project.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key | 
 **environmentKeyQuery** | **string**| The environment key | [optional] 
 **tag** | **string**| Filter by tag | [optional] 

### Return type

[**FeatureFlags**](FeatureFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFeatureFlag**
> FeatureFlag PatchFeatureFlag($projectKey, $featureFlagKey, $patchDelta)

Modify a feature flag by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key | 
 **featureFlagKey** | **string**| The feature flags key | 
 **patchDelta** | [**[]PatchDelta**](patchDelta.md)| http://jsonpatch.com/ | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFeatureFlag**
> PostFeatureFlag($projectKey, $featureFlagKey, $featureFlagBody)

Create a feature flag


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key | 
 **featureFlagKey** | **string**| The feature flags key | 
 **featureFlagBody** | [**FeatureFlagBody**](FeatureFlagBody.md)| Create a new feature flag | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

