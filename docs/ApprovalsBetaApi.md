# \ApprovalsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApprovalRequest**](ApprovalsBetaApi.md#DeleteApprovalRequest) | **Delete** /api/v2/approval-requests/{id} | Delete approval request
[**GetApprovalRequest**](ApprovalsBetaApi.md#GetApprovalRequest) | **Get** /api/v2/approval-requests/{id} | Get approval request
[**GetApprovalRequests**](ApprovalsBetaApi.md#GetApprovalRequests) | **Get** /api/v2/approval-requests | List approval requests
[**PostApprovalRequest**](ApprovalsBetaApi.md#PostApprovalRequest) | **Post** /api/v2/approval-requests | Create approval request
[**PostApprovalRequestApply**](ApprovalsBetaApi.md#PostApprovalRequestApply) | **Post** /api/v2/approval-requests/{id}/apply | Apply approval request
[**PostApprovalRequestReview**](ApprovalsBetaApi.md#PostApprovalRequestReview) | **Post** /api/v2/approval-requests/{id}/reviews | Review approval request



## DeleteApprovalRequest

> DeleteApprovalRequest(ctx, id).Execute()

Delete approval request



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
    id := "id_example" // string | The approval request ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsBetaApi.DeleteApprovalRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.DeleteApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApprovalRequestRequest struct via the builder pattern


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


## GetApprovalRequest

> ExpandableApprovalRequestResponse GetApprovalRequest(ctx, id).Expand(expand).Execute()

Get approval request



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
    id := "id_example" // string | The approval request ID
    expand := "expand_example" // string | A comma-separated list of fields to expand in the response. Supported fields are explained above. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsBetaApi.GetApprovalRequest(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.GetApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequest`: ExpandableApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.GetApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | A comma-separated list of fields to expand in the response. Supported fields are explained above. | 

### Return type

[**ExpandableApprovalRequestResponse**](ExpandableApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalRequests

> ExpandableApprovalRequestsResponse GetApprovalRequests(ctx).Filter(filter).Expand(expand).Limit(limit).Offset(offset).Execute()

List approval requests



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
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form `field operator value`. Supported fields are explained above. (optional)
    expand := "expand_example" // string | A comma-separated list of fields to expand in the response. Supported fields are explained above. (optional)
    limit := int64(789) // int64 | The number of approvals to return. Defaults to -1, which returns all approvals. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsBetaApi.GetApprovalRequests(context.Background()).Filter(filter).Expand(expand).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.GetApprovalRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequests`: ExpandableApprovalRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.GetApprovalRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form &#x60;field operator value&#x60;. Supported fields are explained above. | 
 **expand** | **string** | A comma-separated list of fields to expand in the response. Supported fields are explained above. | 
 **limit** | **int64** | The number of approvals to return. Defaults to -1, which returns all approvals. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**ExpandableApprovalRequestsResponse**](ExpandableApprovalRequestsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApprovalRequest

> ApprovalRequestResponse PostApprovalRequest(ctx).CreateApprovalRequestRequest(createApprovalRequestRequest).Execute()

Create approval request



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
    createApprovalRequestRequest := *openapiclient.NewCreateApprovalRequestRequest("ResourceId_example", "Requesting to update targeting", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // CreateApprovalRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsBetaApi.PostApprovalRequest(context.Background()).CreateApprovalRequestRequest(createApprovalRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.PostApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequest`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.PostApprovalRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApprovalRequestRequest** | [**CreateApprovalRequestRequest**](CreateApprovalRequestRequest.md) |  | 

### Return type

[**ApprovalRequestResponse**](ApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApprovalRequestApply

> ApprovalRequestResponse PostApprovalRequestApply(ctx, id).PostApprovalRequestApplyRequest(postApprovalRequestApplyRequest).Execute()

Apply approval request



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
    id := "id_example" // string | The feature flag approval request ID
    postApprovalRequestApplyRequest := *openapiclient.NewPostApprovalRequestApplyRequest() // PostApprovalRequestApplyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsBetaApi.PostApprovalRequestApply(context.Background(), id).PostApprovalRequestApplyRequest(postApprovalRequestApplyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.PostApprovalRequestApply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestApply`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.PostApprovalRequestApply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApprovalRequestApplyRequest** | [**PostApprovalRequestApplyRequest**](PostApprovalRequestApplyRequest.md) |  | 

### Return type

[**ApprovalRequestResponse**](ApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApprovalRequestReview

> ApprovalRequestResponse PostApprovalRequestReview(ctx, id).PostApprovalRequestReviewRequest(postApprovalRequestReviewRequest).Execute()

Review approval request



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
    id := "id_example" // string | The approval request ID
    postApprovalRequestReviewRequest := *openapiclient.NewPostApprovalRequestReviewRequest() // PostApprovalRequestReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsBetaApi.PostApprovalRequestReview(context.Background(), id).PostApprovalRequestReviewRequest(postApprovalRequestReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.PostApprovalRequestReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestReview`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.PostApprovalRequestReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postApprovalRequestReviewRequest** | [**PostApprovalRequestReviewRequest**](PostApprovalRequestReviewRequest.md) |  | 

### Return type

[**ApprovalRequestResponse**](ApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

