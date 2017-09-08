# \WebhooksApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /webhooks/{webhookId} | Delete a webhook by ID
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /webhooks/{webhookId} | Get a webhook by ID
[**GetWebhooks**](WebhooksApi.md#GetWebhooks) | **Get** /webhooks | Fetch a list of all webhooks
[**PatchWebhook**](WebhooksApi.md#PatchWebhook) | **Patch** /webhooks/{webhookId} | Modify a webhook by ID
[**PostWebhook**](WebhooksApi.md#PostWebhook) | **Post** /webhooks | Create a webhook


# **DeleteWebhook**
> DeleteWebhook($webhookId)

Delete a webhook by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string**| The webhook ID | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> Webhook GetWebhook($webhookId)

Get a webhook by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string**| The webhook ID | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhooks**
> Webhooks GetWebhooks()

Fetch a list of all webhooks


### Parameters
This endpoint does not need any parameter.

### Return type

[**Webhooks**](Webhooks.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchWebhook**
> PatchWebhook($webhookId, $patchDelta)

Modify a webhook by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string**| The webhook ID | 
 **patchDelta** | [**[]PatchDelta**](patchDelta.md)| http://jsonpatch.com/ | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostWebhook**
> PostWebhook($webhookPost)

Create a webhook


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookPost** | [**WebhookPost**](WebhookPost.md)| New webhook | 

### Return type

void (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

