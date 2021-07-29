# \IntegrationsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](IntegrationsApi.md#CreateSubscription) | **Post** /api/v2/integrations/{integrationKey} | Create integration configuration
[**DeleteSubscription**](IntegrationsApi.md#DeleteSubscription) | **Delete** /api/v2/integrations/{integrationKey}/{id} | Delete integration configuration
[**GetConfigurations**](IntegrationsApi.md#GetConfigurations) | **Get** /api/v2/integrations | List integration configurations
[**GetIntegrationSubscriptionByID**](IntegrationsApi.md#GetIntegrationSubscriptionByID) | **Get** /api/v2/integrations/{integrationKey}/{id} | Get configured integration by ID
[**UpdateSubscription**](IntegrationsApi.md#UpdateSubscription) | **Patch** /api/v2/integrations/{integrationKey} | Update subscription



## CreateSubscription

> IntegrationSubscriptionRep CreateSubscription(ctx, integrationKey).IntegrationsSubscriptionPost(integrationsSubscriptionPost).Execute()

Create integration configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    integrationKey := "integrationKey_example" // string | The integration key
    integrationsSubscriptionPost := *openapiclient.NewIntegrationsSubscriptionPost() // IntegrationsSubscriptionPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.CreateSubscription(context.Background(), integrationKey).IntegrationsSubscriptionPost(integrationsSubscriptionPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: IntegrationSubscriptionRep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationsSubscriptionPost** | [**IntegrationsSubscriptionPost**](IntegrationsSubscriptionPost.md) |  | 

### Return type

[**IntegrationSubscriptionRep**](IntegrationSubscriptionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, integrationKey, id).Execute()

Delete integration configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    integrationKey := "integrationKey_example" // string | The integration key
    id := "id_example" // string | The subscription ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.DeleteSubscription(context.Background(), integrationKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurations

> []InlineResponse2001 GetConfigurations(ctx).Execute()

List integration configurations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.GetConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurations`: []InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsRequest struct via the builder pattern


### Return type

[**[]InlineResponse2001**](InlineResponse2001.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationSubscriptionByID

> IntegrationSubscriptionRepCollection GetIntegrationSubscriptionByID(ctx, integrationKey, id).Execute()

Get configured integration by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    integrationKey := "integrationKey_example" // string | The integration key
    id := "id_example" // string | The subscription ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.GetIntegrationSubscriptionByID(context.Background(), integrationKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.GetIntegrationSubscriptionByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationSubscriptionByID`: IntegrationSubscriptionRepCollection
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.GetIntegrationSubscriptionByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationSubscriptionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IntegrationSubscriptionRepCollection**](IntegrationSubscriptionRepCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> IntegrationSubscriptionRep UpdateSubscription(ctx, integrationKey).Execute()

Update subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    integrationKey := "integrationKey_example" // string | The integration key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsApi.UpdateSubscription(context.Background(), integrationKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsApi.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscription`: IntegrationSubscriptionRep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsApi.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationSubscriptionRep**](IntegrationSubscriptionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

