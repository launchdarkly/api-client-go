# \ApplicationsInternalApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationVersionSupported**](ApplicationsInternalApi.md#GetApplicationVersionSupported) | **Get** /internal/applications/{applicationKey}/versions | Get number of supported versions for an application
[**InternalGetApplicationVersionAdoption**](ApplicationsInternalApi.md#InternalGetApplicationVersionAdoption) | **Get** /internal/applications/{applicationKey}/version-adoption | Get application version adoption data by application key



## GetApplicationVersionSupported

> ApplicationVersionBySupportedState GetApplicationVersionSupported(ctx, applicationKey).Supported(supported).Execute()

Get number of supported versions for an application



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
    applicationKey := "applicationKey_example" // string | The application key
    supported := "supported_example" // string | Filter by supported state

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsInternalApi.GetApplicationVersionSupported(context.Background(), applicationKey).Supported(supported).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsInternalApi.GetApplicationVersionSupported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersionSupported`: ApplicationVersionBySupportedState
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsInternalApi.GetApplicationVersionSupported`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationVersionSupportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supported** | **string** | Filter by supported state | 

### Return type

[**ApplicationVersionBySupportedState**](ApplicationVersionBySupportedState.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalGetApplicationVersionAdoption

> InternalApplicationVersionAdoptionCollection InternalGetApplicationVersionAdoption(ctx, applicationKey).Execute()

Get application version adoption data by application key



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
    applicationKey := "applicationKey_example" // string | The application key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsInternalApi.InternalGetApplicationVersionAdoption(context.Background(), applicationKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsInternalApi.InternalGetApplicationVersionAdoption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InternalGetApplicationVersionAdoption`: InternalApplicationVersionAdoptionCollection
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsInternalApi.InternalGetApplicationVersionAdoption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInternalGetApplicationVersionAdoptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InternalApplicationVersionAdoptionCollection**](InternalApplicationVersionAdoptionCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

