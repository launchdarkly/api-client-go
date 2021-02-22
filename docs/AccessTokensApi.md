# \AccessTokensApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToken**](AccessTokensApi.md#DeleteToken) | **Delete** /tokens/{tokenId} | Delete an access token by ID.
[**GetToken**](AccessTokensApi.md#GetToken) | **Get** /tokens/{tokenId} | Get a single access token by ID.
[**GetTokens**](AccessTokensApi.md#GetTokens) | **Get** /tokens | Returns a list of tokens in the account.
[**PatchToken**](AccessTokensApi.md#PatchToken) | **Patch** /tokens/{tokenId} | Modify an access token by ID.
[**PostToken**](AccessTokensApi.md#PostToken) | **Post** /tokens | Create a new token.
[**ResetToken**](AccessTokensApi.md#ResetToken) | **Post** /tokens/{tokenId}/reset | Reset an access token&#39;s secret key with an optional expiry time for the old key.


# **DeleteToken**
> DeleteToken(ctx, tokenId)
Delete an access token by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| The access token ID. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetToken**
> Token GetToken(ctx, tokenId)
Get a single access token by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| The access token ID. | 

### Return type

[**Token**](Token.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokens**
> Tokens GetTokens(ctx, optional)
Returns a list of tokens in the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccessTokensApiGetTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessTokensApiGetTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showAll** | **optional.Bool**| If set to true, and the authentication access token has the \&quot;Admin\&quot; role, personal access tokens for all members will be retrieved. | 

### Return type

[**Tokens**](Tokens.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchToken**
> Token PatchToken(ctx, tokenId, patchDelta)
Modify an access token by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| The access token ID. | 
  **patchDelta** | [**[]PatchOperation**](PatchOperation.md)| Requires a JSON Patch representation of the desired changes to the project. &#39;http://jsonpatch.com/&#39; | 

### Return type

[**Token**](Token.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostToken**
> Token PostToken(ctx, tokenBody)
Create a new token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenBody** | [**TokenBody**](TokenBody.md)| Create a new access token. | 

### Return type

[**Token**](Token.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetToken**
> Token ResetToken(ctx, tokenId, optional)
Reset an access token's secret key with an optional expiry time for the old key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| The access token ID. | 
 **optional** | ***AccessTokensApiResetTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessTokensApiResetTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiry** | **optional.Int64**| An expiration time for the old token key, expressed as a Unix epoch time in milliseconds. By default, the token will expire immediately. | 

### Return type

[**Token**](Token.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

