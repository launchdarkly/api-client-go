# \DataExportDestinationsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDestination**](DataExportDestinationsApi.md#DeleteDestination) | **Delete** /api/v2/destinations/{projKey}/{envKey}/{id} | Delete Data Export destination
[**GetDestination**](DataExportDestinationsApi.md#GetDestination) | **Get** /api/v2/destinations/{projKey}/{envKey}/{id} | Get destination
[**GetDestinations**](DataExportDestinationsApi.md#GetDestinations) | **Get** /api/v2/destinations | List destinations
[**PatchDestination**](DataExportDestinationsApi.md#PatchDestination) | **Patch** /api/v2/destinations/{projKey}/{envKey}/{id} | Update Data Export destination
[**PostDestination**](DataExportDestinationsApi.md#PostDestination) | **Post** /api/v2/destinations/{projKey}/{envKey} | Create data export destination



## DeleteDestination

> DeleteDestination(ctx, projKey, envKey, id).Execute()

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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    id := "id_example" // string | The data export destination ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataExportDestinationsApi.DeleteDestination(context.Background(), projKey, envKey, id).Execute()
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
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**id** | **string** | The data export destination ID | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestination

> DestinationRep GetDestination(ctx, projKey, envKey, id).Execute()

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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    id := "id_example" // string | The Data Export destination ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataExportDestinationsApi.GetDestination(context.Background(), projKey, envKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.GetDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestination`: DestinationRep
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.GetDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**id** | **string** | The Data Export destination ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DestinationRep**](DestinationRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinations

> DestinationCollectionRep GetDestinations(ctx).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataExportDestinationsApi.GetDestinations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.GetDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestinations`: DestinationCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.GetDestinations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationsRequest struct via the builder pattern


### Return type

[**DestinationCollectionRep**](DestinationCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDestination

> DestinationRep PatchDestination(ctx, projKey, envKey, id).JSONPatchElt(jSONPatchElt).Execute()

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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    id := "id_example" // string | The Data Export destination ID
    jSONPatchElt := []openapiclient.JSONPatchElt{*openapiclient.NewJSONPatchElt()} // []JSONPatchElt | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataExportDestinationsApi.PatchDestination(context.Background(), projKey, envKey, id).JSONPatchElt(jSONPatchElt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.PatchDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDestination`: DestinationRep
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.PatchDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**id** | **string** | The Data Export destination ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **jSONPatchElt** | [**[]JSONPatchElt**](JSONPatchElt.md) |  | 

### Return type

[**DestinationRep**](DestinationRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDestination

> DestinationRep PostDestination(ctx, projKey, envKey).Body(body).Execute()

Create data export destination



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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataExportDestinationsApi.PostDestination(context.Background(), projKey, envKey).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataExportDestinationsApi.PostDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostDestination`: DestinationRep
    fmt.Fprintf(os.Stdout, "Response from `DataExportDestinationsApi.PostDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

[**DestinationRep**](DestinationRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

