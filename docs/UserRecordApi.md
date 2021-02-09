# \UserRecordApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUser**](UserRecordApi.md#GetUser) | **Get** /users/{projectKey}/{environmentKey}/{userKey} | Get a user by key.


# **GetUser**
> User GetUser(ctx, projectKey, environmentKey, userKey)
Get a user by key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **userKey** | **string**| The user&#39;s key. | 

### Return type

[**User**](User.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

