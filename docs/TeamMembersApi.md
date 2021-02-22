# \TeamMembersApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMember**](TeamMembersApi.md#DeleteMember) | **Delete** /members/{memberId} | Delete a team member by ID.
[**GetMe**](TeamMembersApi.md#GetMe) | **Get** /members/me | Get the current team member associated with the token
[**GetMember**](TeamMembersApi.md#GetMember) | **Get** /members/{memberId} | Get a single team member by ID.
[**GetMembers**](TeamMembersApi.md#GetMembers) | **Get** /members | Returns a list of all members in the account.
[**PatchMember**](TeamMembersApi.md#PatchMember) | **Patch** /members/{memberId} | Modify a team member by ID.
[**PostMembers**](TeamMembersApi.md#PostMembers) | **Post** /members | Invite new members.


# **DeleteMember**
> DeleteMember(ctx, memberId)
Delete a team member by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **memberId** | **string**| The member ID. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMe**
> Member GetMe(ctx, )
Get the current team member associated with the token

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Member**](Member.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMember**
> Member GetMember(ctx, memberId)
Get a single team member by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **memberId** | **string**| The member ID. | 

### Return type

[**Member**](Member.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMembers**
> Members GetMembers(ctx, optional)
Returns a list of all members in the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamMembersApiGetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamMembersApiGetMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Float32**| The number of objects to return. Defaults to -1, which returns everything. | 
 **offset** | **optional.Float32**| Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first 10 items and then return the next limit items. | 
 **filter** | **optional.String**| A comma-separated list of filters. Each filter is of the form field:value. | 
 **sort** | **optional.String**| A comma-separated list of fields to sort by. A field prefixed by a - will be sorted in descending order. | 

### Return type

[**Members**](Members.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchMember**
> Member PatchMember(ctx, memberId, patchDelta)
Modify a team member by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **memberId** | **string**| The member ID. | 
  **patchDelta** | [**[]PatchOperation**](PatchOperation.md)| Requires a JSON Patch representation of the desired changes to the project. &#39;http://jsonpatch.com/&#39; | 

### Return type

[**Member**](Member.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMembers**
> Members PostMembers(ctx, membersBody)
Invite new members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membersBody** | [**[]MembersBody**](membersBody.md)| New members to invite. | 

### Return type

[**Members**](Members.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

