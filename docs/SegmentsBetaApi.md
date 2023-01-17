# \SegmentsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBigSegmentExport**](SegmentsBetaApi.md#CreateBigSegmentExport) | **Post** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/exports | Create Big Segment export
[**CreateBigSegmentImport**](SegmentsBetaApi.md#CreateBigSegmentImport) | **Post** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/imports | Create Big Segment import
[**GetBigSegmentExport**](SegmentsBetaApi.md#GetBigSegmentExport) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/exports/{exportID} | Get Big Segment export
[**GetBigSegmentImport**](SegmentsBetaApi.md#GetBigSegmentImport) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/imports/{importID} | Get Big Segment import
[**GetContextInstanceSegmentsMembershipByEnv**](SegmentsBetaApi.md#GetContextInstanceSegmentsMembershipByEnv) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/segments/evaluate | List Segment Memberships for Context Instance



## CreateBigSegmentExport

> CreateBigSegmentExport(ctx, projectKey, environmentKey, segmentKey).Execute()

Create Big Segment export



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
    segmentKey := "segmentKey_example" // string | The segment key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsBetaApi.CreateBigSegmentExport(context.Background(), projectKey, environmentKey, segmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsBetaApi.CreateBigSegmentExport``: %v\n", err)
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
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBigSegmentExportRequest struct via the builder pattern


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


## CreateBigSegmentImport

> CreateBigSegmentImport(ctx, projectKey, environmentKey, segmentKey).File(file).Mode(mode).Execute()

Create Big Segment import



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
    segmentKey := "segmentKey_example" // string | The segment key
    file := os.NewFile(1234, "some_file") // *os.File | CSV file containing keys (optional)
    mode := "mode_example" // string | Import mode. Use either `merge` or `replace` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsBetaApi.CreateBigSegmentImport(context.Background(), projectKey, environmentKey, segmentKey).File(file).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsBetaApi.CreateBigSegmentImport``: %v\n", err)
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
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBigSegmentImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | ***os.File** | CSV file containing keys | 
 **mode** | **string** | Import mode. Use either &#x60;merge&#x60; or &#x60;replace&#x60; | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBigSegmentExport

> Export GetBigSegmentExport(ctx, projectKey, environmentKey, segmentKey, exportID).Execute()

Get Big Segment export



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
    segmentKey := "segmentKey_example" // string | The segment key
    exportID := "exportID_example" // string | The export ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsBetaApi.GetBigSegmentExport(context.Background(), projectKey, environmentKey, segmentKey, exportID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsBetaApi.GetBigSegmentExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBigSegmentExport`: Export
    fmt.Fprintf(os.Stdout, "Response from `SegmentsBetaApi.GetBigSegmentExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 
**exportID** | **string** | The export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBigSegmentExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Export**](Export.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBigSegmentImport

> Import GetBigSegmentImport(ctx, projectKey, environmentKey, segmentKey, importID).Execute()

Get Big Segment import



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
    segmentKey := "segmentKey_example" // string | The segment key
    importID := "importID_example" // string | The import ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsBetaApi.GetBigSegmentImport(context.Background(), projectKey, environmentKey, segmentKey, importID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsBetaApi.GetBigSegmentImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBigSegmentImport`: Import
    fmt.Fprintf(os.Stdout, "Response from `SegmentsBetaApi.GetBigSegmentImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 
**importID** | **string** | The import ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBigSegmentImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Import**](Import.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextInstanceSegmentsMembershipByEnv

> ContextInstanceSegmentMemberships GetContextInstanceSegmentsMembershipByEnv(ctx, projectKey, environmentKey).RequestBody(requestBody).Execute()

List Segment Memberships for Context Instance



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsBetaApi.GetContextInstanceSegmentsMembershipByEnv(context.Background(), projectKey, environmentKey).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsBetaApi.GetContextInstanceSegmentsMembershipByEnv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextInstanceSegmentsMembershipByEnv`: ContextInstanceSegmentMemberships
    fmt.Fprintf(os.Stdout, "Response from `SegmentsBetaApi.GetContextInstanceSegmentsMembershipByEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextInstanceSegmentsMembershipByEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**ContextInstanceSegmentMemberships**](ContextInstanceSegmentMemberships.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

