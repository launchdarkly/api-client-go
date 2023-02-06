# \ApprovalsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApprovalRequest**](ApprovalsBetaApi.md#GetApprovalRequest) | **Get** /api/v2/approval-requests/{id} | Get approval request
[**GetApprovalRequests**](ApprovalsBetaApi.md#GetApprovalRequests) | **Get** /api/v2/approval-requests | List approval requests



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
    expand := "expand_example" // string | A comma-separated list of fields the user wants to expand in the response. Supported fields are explained above. (optional)
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
 **expand** | **string** | A comma-separated list of fields the user wants to expand in the response. Supported fields are explained above. | 
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

