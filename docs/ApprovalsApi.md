# \ApprovalsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApprovalRequest**](ApprovalsApi.md#DeleteApprovalRequest) | **Delete** /api/v2/approval-requests/{id} | Delete approval request
[**DeleteApprovalRequestForFlag**](ApprovalsApi.md#DeleteApprovalRequestForFlag) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id} | Delete approval request for a flag
[**GetApprovalForFlag**](ApprovalsApi.md#GetApprovalForFlag) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id} | Get approval request for a flag
[**GetApprovalRequest**](ApprovalsApi.md#GetApprovalRequest) | **Get** /api/v2/approval-requests/{id} | Get approval request
[**GetApprovalRequests**](ApprovalsApi.md#GetApprovalRequests) | **Get** /api/v2/approval-requests | List approval requests
[**GetApprovalsForFlag**](ApprovalsApi.md#GetApprovalsForFlag) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | List approval requests for a flag
[**PostApprovalRequest**](ApprovalsApi.md#PostApprovalRequest) | **Post** /api/v2/approval-requests | Create approval request
[**PostApprovalRequestApply**](ApprovalsApi.md#PostApprovalRequestApply) | **Post** /api/v2/approval-requests/{id}/apply | Apply approval request
[**PostApprovalRequestApplyForFlag**](ApprovalsApi.md#PostApprovalRequestApplyForFlag) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id}/apply | Apply approval request for a flag
[**PostApprovalRequestForFlag**](ApprovalsApi.md#PostApprovalRequestForFlag) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Create approval request for a flag
[**PostApprovalRequestReview**](ApprovalsApi.md#PostApprovalRequestReview) | **Post** /api/v2/approval-requests/{id}/reviews | Review approval request
[**PostApprovalRequestReviewForFlag**](ApprovalsApi.md#PostApprovalRequestReviewForFlag) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id}/reviews | Review approval request for a flag
[**PostFlagCopyConfigApprovalRequest**](ApprovalsApi.md#PostFlagCopyConfigApprovalRequest) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests-flag-copy | Create approval request to copy flag configurations across environments



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
    resp, r, err := apiClient.ApprovalsApi.DeleteApprovalRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.DeleteApprovalRequest``: %v\n", err)
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


## DeleteApprovalRequestForFlag

> DeleteApprovalRequestForFlag(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

Delete approval request for a flag



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.DeleteApprovalRequestForFlag(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.DeleteApprovalRequestForFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApprovalRequestForFlagRequest struct via the builder pattern


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


## GetApprovalForFlag

> FlagConfigApprovalRequestResponse GetApprovalForFlag(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

Get approval request for a flag



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.GetApprovalForFlag(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.GetApprovalForFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalForFlag`: FlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.GetApprovalForFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalForFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FlagConfigApprovalRequestResponse**](FlagConfigApprovalRequestResponse.md)

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
    resp, r, err := apiClient.ApprovalsApi.GetApprovalRequest(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.GetApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequest`: ExpandableApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.GetApprovalRequest`: %v\n", resp)
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
    limit := int64(789) // int64 | The number of approvals to return. Defaults to 20. Maximum limit is 200. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.GetApprovalRequests(context.Background()).Filter(filter).Expand(expand).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.GetApprovalRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalRequests`: ExpandableApprovalRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.GetApprovalRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form &#x60;field operator value&#x60;. Supported fields are explained above. | 
 **expand** | **string** | A comma-separated list of fields to expand in the response. Supported fields are explained above. | 
 **limit** | **int64** | The number of approvals to return. Defaults to 20. Maximum limit is 200. | 
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


## GetApprovalsForFlag

> FlagConfigApprovalRequestsResponse GetApprovalsForFlag(ctx, projectKey, featureFlagKey, environmentKey).Execute()

List approval requests for a flag



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.GetApprovalsForFlag(context.Background(), projectKey, featureFlagKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.GetApprovalsForFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalsForFlag`: FlagConfigApprovalRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.GetApprovalsForFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalsForFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlagConfigApprovalRequestsResponse**](FlagConfigApprovalRequestsResponse.md)

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
    resp, r, err := apiClient.ApprovalsApi.PostApprovalRequest(context.Background()).CreateApprovalRequestRequest(createApprovalRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequest`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequest`: %v\n", resp)
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
    resp, r, err := apiClient.ApprovalsApi.PostApprovalRequestApply(context.Background(), id).PostApprovalRequestApplyRequest(postApprovalRequestApplyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequestApply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestApply`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequestApply`: %v\n", resp)
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


## PostApprovalRequestApplyForFlag

> FlagConfigApprovalRequestResponse PostApprovalRequestApplyForFlag(ctx, projectKey, featureFlagKey, environmentKey, id).PostApprovalRequestApplyRequest(postApprovalRequestApplyRequest).Execute()

Apply approval request for a flag



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID
    postApprovalRequestApplyRequest := *openapiclient.NewPostApprovalRequestApplyRequest() // PostApprovalRequestApplyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.PostApprovalRequestApplyForFlag(context.Background(), projectKey, featureFlagKey, environmentKey, id).PostApprovalRequestApplyRequest(postApprovalRequestApplyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequestApplyForFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestApplyForFlag`: FlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequestApplyForFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestApplyForFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **postApprovalRequestApplyRequest** | [**PostApprovalRequestApplyRequest**](PostApprovalRequestApplyRequest.md) |  | 

### Return type

[**FlagConfigApprovalRequestResponse**](FlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApprovalRequestForFlag

> FlagConfigApprovalRequestResponse PostApprovalRequestForFlag(ctx, projectKey, featureFlagKey, environmentKey).CreateFlagConfigApprovalRequestRequest(createFlagConfigApprovalRequestRequest).Execute()

Create approval request for a flag



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    createFlagConfigApprovalRequestRequest := *openapiclient.NewCreateFlagConfigApprovalRequestRequest("Requesting to update targeting", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // CreateFlagConfigApprovalRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.PostApprovalRequestForFlag(context.Background(), projectKey, featureFlagKey, environmentKey).CreateFlagConfigApprovalRequestRequest(createFlagConfigApprovalRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequestForFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestForFlag`: FlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequestForFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestForFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createFlagConfigApprovalRequestRequest** | [**CreateFlagConfigApprovalRequestRequest**](CreateFlagConfigApprovalRequestRequest.md) |  | 

### Return type

[**FlagConfigApprovalRequestResponse**](FlagConfigApprovalRequestResponse.md)

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
    resp, r, err := apiClient.ApprovalsApi.PostApprovalRequestReview(context.Background(), id).PostApprovalRequestReviewRequest(postApprovalRequestReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequestReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestReview`: ApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequestReview`: %v\n", resp)
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


## PostApprovalRequestReviewForFlag

> FlagConfigApprovalRequestResponse PostApprovalRequestReviewForFlag(ctx, projectKey, featureFlagKey, environmentKey, id).PostApprovalRequestReviewRequest(postApprovalRequestReviewRequest).Execute()

Review approval request for a flag



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID
    postApprovalRequestReviewRequest := *openapiclient.NewPostApprovalRequestReviewRequest() // PostApprovalRequestReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.PostApprovalRequestReviewForFlag(context.Background(), projectKey, featureFlagKey, environmentKey, id).PostApprovalRequestReviewRequest(postApprovalRequestReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequestReviewForFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestReviewForFlag`: FlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequestReviewForFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestReviewForFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **postApprovalRequestReviewRequest** | [**PostApprovalRequestReviewRequest**](PostApprovalRequestReviewRequest.md) |  | 

### Return type

[**FlagConfigApprovalRequestResponse**](FlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlagCopyConfigApprovalRequest

> FlagConfigApprovalRequestResponse PostFlagCopyConfigApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey).CreateCopyFlagConfigApprovalRequestRequest(createCopyFlagConfigApprovalRequestRequest).Execute()

Create approval request to copy flag configurations across environments



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key for the target environment
    createCopyFlagConfigApprovalRequestRequest := *openapiclient.NewCreateCopyFlagConfigApprovalRequestRequest("copy flag settings to another environment", *openapiclient.NewSourceFlag("environment-key-123abc")) // CreateCopyFlagConfigApprovalRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.PostFlagCopyConfigApprovalRequest(context.Background(), projectKey, featureFlagKey, environmentKey).CreateCopyFlagConfigApprovalRequestRequest(createCopyFlagConfigApprovalRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostFlagCopyConfigApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFlagCopyConfigApprovalRequest`: FlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostFlagCopyConfigApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key for the target environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagCopyConfigApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createCopyFlagConfigApprovalRequestRequest** | [**CreateCopyFlagConfigApprovalRequestRequest**](CreateCopyFlagConfigApprovalRequestRequest.md) |  | 

### Return type

[**FlagConfigApprovalRequestResponse**](FlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

