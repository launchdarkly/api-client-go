# \UsersApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /users/{projectKey}/{environmentKey}/{userKey} | Delete a user by ID
[**GetSearchUsers**](UsersApi.md#GetSearchUsers) | **Get** /user-search/{projectKey}/{environmentKey} | Search users in LaunchDarkly based on their last active date, or a search query.
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{projectKey}/{environmentKey}/{userKey} | Get a user by key.
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /users/{projectKey}/{environmentKey} | List all users in the environment.


# **DeleteUser**
> DeleteUser($projectKey, $environmentKey, $userKey)

Delete a user by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
 **environmentKey** | **string**| The environment key | 
 **userKey** | **string**| The user&#39;s key | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearchUsers**
> Users GetSearchUsers($projectKey, $environmentKey, $q, $limit, $offset, $after)

Search users in LaunchDarkly based on their last active date, or a search query.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
 **environmentKey** | **string**| The environment key | 
 **q** | **string**| Search query | [optional] 
 **limit** | **float32**| Pagination limit | [optional] 
 **offset** | **float32**| Specifies the first item to return in the collection | [optional] 
 **after** | **float32**| A unix epoch time in milliseconds specifying the maximum last time a user requested a feature flag | [optional] 

### Return type

[**Users**](Users.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser($projectKey, $environmentKey, $userKey)

Get a user by key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
 **environmentKey** | **string**| The environment key | 
 **userKey** | **string**| The user&#39;s key | 

### Return type

[**User**](User.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> Users GetUsers($projectKey, $environmentKey, $limit)

List all users in the environment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
 **environmentKey** | **string**| The environment key | 
 **limit** | **float32**| Pagination limit | [optional] 

### Return type

[**Users**](Users.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

