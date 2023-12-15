# \PayloadFiltersApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayloadFilter**](PayloadFiltersApi.md#GetPayloadFilter) | **Get** /api/v2/projects/{projectKey}/payload-filters/{payloadFilterKey} | Get payload filter
[**GetPayloadFilters**](PayloadFiltersApi.md#GetPayloadFilters) | **Get** /api/v2/projects/{projectKey}/payload-filters | List payload filters
[**PatchPayloadFilter**](PayloadFiltersApi.md#PatchPayloadFilter) | **Patch** /api/v2/projects/{projectKey}/payload-filters/{payloadFilterKey} | Update payload filter
[**PostPayloadFilters**](PayloadFiltersApi.md#PostPayloadFilters) | **Post** /api/v2/projects/{projectKey}/payload-filters | Create payload filter



## GetPayloadFilter

> FilterRep GetPayloadFilter(ctx, projectKey, payloadFilterKey).Execute()

Get payload filter



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
    payloadFilterKey := "payloadFilterKey_example" // string | The payload filter key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadFiltersApi.GetPayloadFilter(context.Background(), projectKey, payloadFilterKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadFiltersApi.GetPayloadFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayloadFilter`: FilterRep
    fmt.Fprintf(os.Stdout, "Response from `PayloadFiltersApi.GetPayloadFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**payloadFilterKey** | **string** | The payload filter key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayloadFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FilterRep**](FilterRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayloadFilters

> FilterCollectionRep GetPayloadFilters(ctx, projectKey).Limit(limit).Offset(offset).Filter(filter).Execute()

List payload filters



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
    limit := int64(789) // int64 | The number of payload filters to return in the response. Defaults to 20 with a maximum of 1000. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and returns the next `limit` items. (optional)
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is constructed as `field:value`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadFiltersApi.GetPayloadFilters(context.Background(), projectKey).Limit(limit).Offset(offset).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadFiltersApi.GetPayloadFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayloadFilters`: FilterCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `PayloadFiltersApi.GetPayloadFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayloadFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The number of payload filters to return in the response. Defaults to 20 with a maximum of 1000. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and returns the next &#x60;limit&#x60; items. | 
 **filter** | **string** | A comma-separated list of filters. Each filter is constructed as &#x60;field:value&#x60;. | 

### Return type

[**FilterCollectionRep**](FilterCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPayloadFilter

> PatchPayloadFilter(ctx, projectKey, payloadFilterKey).PatchFilterRep(patchFilterRep).Execute()

Update payload filter



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
    payloadFilterKey := "payloadFilterKey_example" // string | The payload filter key
    patchFilterRep := *openapiclient.NewPatchFilterRep() // PatchFilterRep | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadFiltersApi.PatchPayloadFilter(context.Background(), projectKey, payloadFilterKey).PatchFilterRep(patchFilterRep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadFiltersApi.PatchPayloadFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**payloadFilterKey** | **string** | The payload filter key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPayloadFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchFilterRep** | [**PatchFilterRep**](PatchFilterRep.md) |  | 

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


## PostPayloadFilters

> FilterRep PostPayloadFilters(ctx, projectKey).PostFilterRep(postFilterRep).Execute()

Create payload filter



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
    postFilterRep := *openapiclient.NewPostFilterRep("example-filter", "example filter", "This is an example filter", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // PostFilterRep | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadFiltersApi.PostPayloadFilters(context.Background(), projectKey).PostFilterRep(postFilterRep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadFiltersApi.PostPayloadFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPayloadFilters`: FilterRep
    fmt.Fprintf(os.Stdout, "Response from `PayloadFiltersApi.PostPayloadFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPayloadFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postFilterRep** | [**PostFilterRep**](PostFilterRep.md) |  | 

### Return type

[**FilterRep**](FilterRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

