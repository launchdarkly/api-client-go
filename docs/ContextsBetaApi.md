# \ContextsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContextKindsByProjectKey**](ContextsBetaApi.md#GetContextKindsByProjectKey) | **Get** /api/v2/projects/{projectKey}/context-kinds | Get context kinds
[**PutContextKind**](ContextsBetaApi.md#PutContextKind) | **Put** /api/v2/projects/{projectKey}/context-kinds/{key} | Create or update context kind



## GetContextKindsByProjectKey

> ContextKindsCollectionRep GetContextKindsByProjectKey(ctx, projectKey).Execute()

Get context kinds



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
    projectKey := "projectKey_example" // string | The project key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContextKindsByProjectKey(context.Background(), projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.GetContextKindsByProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextKindsByProjectKey`: ContextKindsCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.GetContextKindsByProjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextKindsByProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContextKindsCollectionRep**](ContextKindsCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutContextKind

> UpsertResponseRep PutContextKind(ctx, projectKey, key).UpsertContextKindPayload(upsertContextKindPayload).Execute()

Create or update context kind



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
    projectKey := "projectKey_example" // string | The project key
    key := "key_example" // string | The context kind key
    upsertContextKindPayload := *openapiclient.NewUpsertContextKindPayload("organization") // UpsertContextKindPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.PutContextKind(context.Background(), projectKey, key).UpsertContextKindPayload(upsertContextKindPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.PutContextKind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutContextKind`: UpsertResponseRep
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.PutContextKind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**key** | **string** | The context kind key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutContextKindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upsertContextKindPayload** | [**UpsertContextKindPayload**](UpsertContextKindPayload.md) |  | 

### Return type

[**UpsertResponseRep**](UpsertResponseRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

