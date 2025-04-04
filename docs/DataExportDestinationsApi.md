# \DataExportDestinationsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDestination**](DataExportDestinationsApi.md#DeleteDestination) | **Delete** /api/v2/destinations/{projectKey}/{environmentKey}/{id} | Delete Data Export destination
[**GetDestination**](DataExportDestinationsApi.md#GetDestination) | **Get** /api/v2/destinations/{projectKey}/{environmentKey}/{id} | Get destination
[**GetDestinations**](DataExportDestinationsApi.md#GetDestinations) | **Get** /api/v2/destinations | List destinations
[**PatchDestination**](DataExportDestinationsApi.md#PatchDestination) | **Patch** /api/v2/destinations/{projectKey}/{environmentKey}/{id} | Update Data Export destination
[**PostDestination**](DataExportDestinationsApi.md#PostDestination) | **Post** /api/v2/destinations/{projectKey}/{environmentKey} | Create Data Export destination
[**PostGenerateWarehouseDestinationKeyPair**](DataExportDestinationsApi.md#PostGenerateWarehouseDestinationKeyPair) | **Post** /api/v2/destinations/generate-warehouse-destination-key-pair | Generate Snowflake destination key pair



## DeleteDestination

> DeleteDestination(ctx, projectKey, environmentKey, id).Execute()

Delete Data Export destination



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
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The Data Export destination ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataExportDestinationsApi.DeleteDestination(context.Background(), projectKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.DeleteDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The Data Export destination ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestination

> Destination GetDestination(ctx, projectKey, environmentKey, id).Execute()

Get destination



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
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The Data Export destination ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataExportDestinationsApi.GetDestination(context.Background(), projectKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.GetDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestination`: Destination
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.GetDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The Data Export destination ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Destination**](Destination.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinations

> Destinations GetDestinations(ctx).Execute()

List destinations



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
    resp, r, err := apiClient.DataExportDestinationsApi.GetDestinations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.GetDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestinations`: Destinations
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.GetDestinations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationsRequest struct via the builder pattern


### Return type

[**Destinations**](Destinations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDestination

> Destination PatchDestination(ctx, projectKey, environmentKey, id).PatchOperation(patchOperation).Execute()

Update Data Export destination



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
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The Data Export destination ID
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataExportDestinationsApi.PatchDestination(context.Background(), projectKey, environmentKey, id).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.PatchDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDestination`: Destination
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.PatchDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The Data Export destination ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**Destination**](Destination.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDestination

> Destination PostDestination(ctx, projectKey, environmentKey).DestinationPost(destinationPost).Execute()

Create Data Export destination



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
    environmentKey := "environmentKey_example" // string | The environment key
    destinationPost := *openapiclient.NewDestinationPost() // DestinationPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataExportDestinationsApi.PostDestination(context.Background(), projectKey, environmentKey).DestinationPost(destinationPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.PostDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDestination`: Destination
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.PostDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **destinationPost** | [**DestinationPost**](DestinationPost.md) |  | 

### Return type

[**Destination**](Destination.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGenerateWarehouseDestinationKeyPair

> GenerateWarehouseDestinationKeyPairPostRep PostGenerateWarehouseDestinationKeyPair(ctx).Execute()

Generate Snowflake destination key pair



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
    resp, r, err := apiClient.DataExportDestinationsApi.PostGenerateWarehouseDestinationKeyPair(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.PostGenerateWarehouseDestinationKeyPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGenerateWarehouseDestinationKeyPair`: GenerateWarehouseDestinationKeyPairPostRep
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.PostGenerateWarehouseDestinationKeyPair`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGenerateWarehouseDestinationKeyPairRequest struct via the builder pattern


### Return type

[**GenerateWarehouseDestinationKeyPairPostRep**](GenerateWarehouseDestinationKeyPairPostRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

