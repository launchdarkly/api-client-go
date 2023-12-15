# \IntegrationsPrivateApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateBigSegmentStoreIntegrationResponse**](IntegrationsPrivateApi.md#CreatePrivateBigSegmentStoreIntegrationResponse) | **Post** /private/integrations/big-segment-store/{environmentId}/{integrationId}/responses | Create Big Segment store integration response
[**CreatePrivateFeatureStoreIntegrationResponse**](IntegrationsPrivateApi.md#CreatePrivateFeatureStoreIntegrationResponse) | **Post** /private/integrations/featureStore/{envId}/{integrationId}/responses | Create feature store integration response
[**GetAllBigSegmentStoreIntegrationsPrivate**](IntegrationsPrivateApi.md#GetAllBigSegmentStoreIntegrationsPrivate) | **Get** /private/integrations/big-segment-store/{accountId} | Get all enabled Big Segment store integrations
[**PrivateGetFeatureStoreIntegration**](IntegrationsPrivateApi.md#PrivateGetFeatureStoreIntegration) | **Get** /private/integrations/featureStore/{envId}/{integrationId} | Get feature store integration
[**PrivateGetFeatureStoreIntegrationManifest**](IntegrationsPrivateApi.md#PrivateGetFeatureStoreIntegrationManifest) | **Get** /private/integrations/featureStore/{envId}/{integrationId}/manifest | Get feature store integration manifest
[**PrivatePostFeatureStoresRefreshAll**](IntegrationsPrivateApi.md#PrivatePostFeatureStoresRefreshAll) | **Post** /private/integrations/featureStore/refresh | Send notifications to refresh all enabled feature stores



## CreatePrivateBigSegmentStoreIntegrationResponse

> CreatePrivateBigSegmentStoreIntegrationResponse(ctx, environmentId, integrationId).PrivateFeatureStoreResponseRep(privateFeatureStoreResponseRep).Execute()

Create Big Segment store integration response

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
    environmentId := "environmentId_example" // string | The environment ID
    integrationId := "integrationId_example" // string | The integration ID
    privateFeatureStoreResponseRep := *openapiclient.NewPrivateFeatureStoreResponseRep(int32(123)) // PrivateFeatureStoreResponseRep | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsPrivateApi.CreatePrivateBigSegmentStoreIntegrationResponse(context.Background(), environmentId, integrationId).PrivateFeatureStoreResponseRep(privateFeatureStoreResponseRep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPrivateApi.CreatePrivateBigSegmentStoreIntegrationResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The environment ID | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateBigSegmentStoreIntegrationResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateFeatureStoreResponseRep** | [**PrivateFeatureStoreResponseRep**](PrivateFeatureStoreResponseRep.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateFeatureStoreIntegrationResponse

> CreatePrivateFeatureStoreIntegrationResponse(ctx, envId, integrationId).PrivateFeatureStoreResponseRep(privateFeatureStoreResponseRep).Execute()

Create feature store integration response

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
    envId := "envId_example" // string | The environment ID
    integrationId := "integrationId_example" // string | The integration ID
    privateFeatureStoreResponseRep := *openapiclient.NewPrivateFeatureStoreResponseRep(int32(123)) // PrivateFeatureStoreResponseRep | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsPrivateApi.CreatePrivateFeatureStoreIntegrationResponse(context.Background(), envId, integrationId).PrivateFeatureStoreResponseRep(privateFeatureStoreResponseRep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPrivateApi.CreatePrivateFeatureStoreIntegrationResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** | The environment ID | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateFeatureStoreIntegrationResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **privateFeatureStoreResponseRep** | [**PrivateFeatureStoreResponseRep**](PrivateFeatureStoreResponseRep.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBigSegmentStoreIntegrationsPrivate

> PrivateStoreIntegrationCollectionRep GetAllBigSegmentStoreIntegrationsPrivate(ctx, accountID).Execute()

Get all enabled Big Segment store integrations

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
    accountID := "accountID_example" // string | The account ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsPrivateApi.GetAllBigSegmentStoreIntegrationsPrivate(context.Background(), accountID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPrivateApi.GetAllBigSegmentStoreIntegrationsPrivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBigSegmentStoreIntegrationsPrivate`: PrivateStoreIntegrationCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsPrivateApi.GetAllBigSegmentStoreIntegrationsPrivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountID** | **string** | The account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBigSegmentStoreIntegrationsPrivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateStoreIntegrationCollectionRep**](PrivateStoreIntegrationCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateGetFeatureStoreIntegration

> PrivateFeatureStoreRep PrivateGetFeatureStoreIntegration(ctx, envId, integrationId).Execute()

Get feature store integration

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
    envId := "envId_example" // string | The environment key
    integrationId := "integrationId_example" // string | The integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsPrivateApi.PrivateGetFeatureStoreIntegration(context.Background(), envId, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPrivateApi.PrivateGetFeatureStoreIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateGetFeatureStoreIntegration`: PrivateFeatureStoreRep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsPrivateApi.PrivateGetFeatureStoreIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** | The environment key | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateGetFeatureStoreIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PrivateFeatureStoreRep**](PrivateFeatureStoreRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivateGetFeatureStoreIntegrationManifest

> Manifest PrivateGetFeatureStoreIntegrationManifest(ctx, envId, integrationId).Execute()

Get feature store integration manifest

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
    envId := "envId_example" // string | The environment ID
    integrationId := "integrationId_example" // string | The integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsPrivateApi.PrivateGetFeatureStoreIntegrationManifest(context.Background(), envId, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPrivateApi.PrivateGetFeatureStoreIntegrationManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateGetFeatureStoreIntegrationManifest`: Manifest
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsPrivateApi.PrivateGetFeatureStoreIntegrationManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envId** | **string** | The environment ID | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateGetFeatureStoreIntegrationManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Manifest**](Manifest.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrivatePostFeatureStoresRefreshAll

> PrivatePostFeatureStoresRefreshAll(ctx).Execute()

Send notifications to refresh all enabled feature stores

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationsPrivateApi.PrivatePostFeatureStoresRefreshAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsPrivateApi.PrivatePostFeatureStoresRefreshAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivatePostFeatureStoresRefreshAllRequest struct via the builder pattern


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

