# \IntegrationsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIntegrationSubscription**](IntegrationsApi.md#DeleteIntegrationSubscription) | **Delete** /integrations/{integrationKey}/{integrationId} | Delete an integration subscription by ID.
[**GetIntegrationSubscription**](IntegrationsApi.md#GetIntegrationSubscription) | **Get** /integrations/{integrationKey}/{integrationId} | Get a single integration subscription by ID.
[**GetIntegrationSubscriptions**](IntegrationsApi.md#GetIntegrationSubscriptions) | **Get** /integrations/{integrationKey} | Get a list of all configured integrations of a given kind.
[**GetIntegrations**](IntegrationsApi.md#GetIntegrations) | **Get** /integrations | Get a list of all configured audit log event integrations associated with this account.
[**PatchIntegrationSubscription**](IntegrationsApi.md#PatchIntegrationSubscription) | **Patch** /integrations/{integrationKey}/{integrationId} | Modify an integration subscription by ID.
[**PostIntegrationSubscription**](IntegrationsApi.md#PostIntegrationSubscription) | **Post** /integrations/{integrationKey} | Create a new integration subscription of a given kind.


# **DeleteIntegrationSubscription**
> DeleteIntegrationSubscription(ctx, integrationKey, integrationId)
Delete an integration subscription by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationKey** | **string**| The key used to specify the integration kind. | 
  **integrationId** | **string**| The integration ID. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIntegrationSubscription**
> IntegrationSubscription GetIntegrationSubscription(ctx, integrationKey, integrationId)
Get a single integration subscription by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationKey** | **string**| The key used to specify the integration kind. | 
  **integrationId** | **string**| The integration ID. | 

### Return type

[**IntegrationSubscription**](IntegrationSubscription.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIntegrationSubscriptions**
> Integration GetIntegrationSubscriptions(ctx, integrationKey)
Get a list of all configured integrations of a given kind.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationKey** | **string**| The key used to specify the integration kind. | 

### Return type

[**Integration**](Integration.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIntegrations**
> Integrations GetIntegrations(ctx, )
Get a list of all configured audit log event integrations associated with this account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Integrations**](Integrations.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchIntegrationSubscription**
> IntegrationSubscription PatchIntegrationSubscription(ctx, integrationKey, integrationId, patchDelta)
Modify an integration subscription by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationKey** | **string**| The key used to specify the integration kind. | 
  **integrationId** | **string**| The integration ID. | 
  **patchDelta** | [**[]PatchOperation**](PatchOperation.md)| Requires a JSON Patch representation of the desired changes to the project. &#39;http://jsonpatch.com/&#39; | 

### Return type

[**IntegrationSubscription**](IntegrationSubscription.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIntegrationSubscription**
> IntegrationSubscription PostIntegrationSubscription(ctx, integrationKey, subscriptionBody)
Create a new integration subscription of a given kind.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationKey** | **string**| The key used to specify the integration kind. | 
  **subscriptionBody** | [**SubscriptionBody**](SubscriptionBody.md)| Create a new integration subscription. | 

### Return type

[**IntegrationSubscription**](IntegrationSubscription.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

