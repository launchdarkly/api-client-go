# \ProjectsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProject**](ProjectsApi.md#DeleteProject) | **Delete** /projects/{projectKey} | Delete a project by ID
[**GetProject**](ProjectsApi.md#GetProject) | **Get** /projects/{projectKey} | Fetch a single project by key.
[**GetProjects**](ProjectsApi.md#GetProjects) | **Get** /projects | Returns a list of all projects in the account.
[**PatchProject**](ProjectsApi.md#PatchProject) | **Patch** /projects/{projectKey} | Modify a project by ID
[**PostProject**](ProjectsApi.md#PostProject) | **Post** /projects | Create a project


# **DeleteProject**
> DeleteProject($projectKey)

Delete a project by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> Project GetProject($projectKey)

Fetch a single project by key.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 

### Return type

[**Project**](Project.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjects**
> Projects GetProjects()

Returns a list of all projects in the account.


### Parameters
This endpoint does not need any parameter.

### Return type

[**Projects**](Projects.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProject**
> PatchProject($projectKey, $patchDelta)

Modify a project by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
 **patchDelta** | [**[]PatchDelta**](patchDelta.md)| http://jsonpatch.com/ | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostProject**
> PostProject($projectBody)

Create a project


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectBody** | [**ProjectBody**](ProjectBody.md)| New project | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

