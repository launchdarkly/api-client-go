# \ApprovalsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApprovalRequest**](ApprovalsApi.md#DeleteApprovalRequest) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id} | Delete approval request
[**GetApproval**](ApprovalsApi.md#GetApproval) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id} | Get approval request
[**GetApprovals**](ApprovalsApi.md#GetApprovals) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Get all approval requests
[**PostApprovalRequest**](ApprovalsApi.md#PostApprovalRequest) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Create approval request
[**PostApprovalRequestApplyRequest**](ApprovalsApi.md#PostApprovalRequestApplyRequest) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id}/apply | Apply approval request
[**PostApprovalRequestReview**](ApprovalsApi.md#PostApprovalRequestReview) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id}/reviews | Review approval request
[**PostFlagCopyConfigApprovalRequest**](ApprovalsApi.md#PostFlagCopyConfigApprovalRequest) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/copy/environments/{environmentKey}/approval-requests-flag-copy | Create approval request to copy flag configurations across environments



## DeleteApprovalRequest

> DeleteApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalsApi.DeleteApprovalRequest(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
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
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApproval

> WebFlagConfigApprovalRequestResponse GetApproval(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalsApi.GetApproval(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.GetApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApproval`: WebFlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.GetApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**WebFlagConfigApprovalRequestResponse**](WebFlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovals

> WebFlagConfigApprovalRequestResponse GetApprovals(ctx, projectKey, featureFlagKey, environmentKey).Execute()

Get all approval requests



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalsApi.GetApprovals(context.Background(), projectKey, featureFlagKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.GetApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovals`: WebFlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.GetApprovals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WebFlagConfigApprovalRequestResponse**](WebFlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApprovalRequest

> WebFlagConfigApprovalRequestResponse PostApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey).ApprovalsEndpointsCreateFlagConfigApprovalRequestRequest(approvalsEndpointsCreateFlagConfigApprovalRequestRequest).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    approvalsEndpointsCreateFlagConfigApprovalRequestRequest := *openapiclient.NewApprovalsEndpointsCreateFlagConfigApprovalRequestRequest() // ApprovalsEndpointsCreateFlagConfigApprovalRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalsApi.PostApprovalRequest(context.Background(), projectKey, featureFlagKey, environmentKey).ApprovalsEndpointsCreateFlagConfigApprovalRequestRequest(approvalsEndpointsCreateFlagConfigApprovalRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequest`: WebFlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **approvalsEndpointsCreateFlagConfigApprovalRequestRequest** | [**ApprovalsEndpointsCreateFlagConfigApprovalRequestRequest**](ApprovalsEndpointsCreateFlagConfigApprovalRequestRequest.md) |  | 

### Return type

[**WebFlagConfigApprovalRequestResponse**](WebFlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApprovalRequestApplyRequest

> WebFlagConfigApprovalRequestResponse PostApprovalRequestApplyRequest(ctx, projectKey, featureFlagKey, environmentKey, id).ApprovalsEndpointsPostApprovalRequestApplyRequest(approvalsEndpointsPostApprovalRequestApplyRequest).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID
    approvalsEndpointsPostApprovalRequestApplyRequest := *openapiclient.NewApprovalsEndpointsPostApprovalRequestApplyRequest() // ApprovalsEndpointsPostApprovalRequestApplyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalsApi.PostApprovalRequestApplyRequest(context.Background(), projectKey, featureFlagKey, environmentKey, id).ApprovalsEndpointsPostApprovalRequestApplyRequest(approvalsEndpointsPostApprovalRequestApplyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequestApplyRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestApplyRequest`: WebFlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequestApplyRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestApplyRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **approvalsEndpointsPostApprovalRequestApplyRequest** | [**ApprovalsEndpointsPostApprovalRequestApplyRequest**](ApprovalsEndpointsPostApprovalRequestApplyRequest.md) |  | 

### Return type

[**WebFlagConfigApprovalRequestResponse**](WebFlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApprovalRequestReview

> WebFlagConfigApprovalRequestResponse PostApprovalRequestReview(ctx, projectKey, featureFlagKey, environmentKey, id).ApprovalsEndpointsPostApprovalRequestReviewRequest(approvalsEndpointsPostApprovalRequestReviewRequest).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The feature flag approval request ID
    approvalsEndpointsPostApprovalRequestReviewRequest := *openapiclient.NewApprovalsEndpointsPostApprovalRequestReviewRequest() // ApprovalsEndpointsPostApprovalRequestReviewRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalsApi.PostApprovalRequestReview(context.Background(), projectKey, featureFlagKey, environmentKey, id).ApprovalsEndpointsPostApprovalRequestReviewRequest(approvalsEndpointsPostApprovalRequestReviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostApprovalRequestReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostApprovalRequestReview`: WebFlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostApprovalRequestReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The feature flag approval request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApprovalRequestReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **approvalsEndpointsPostApprovalRequestReviewRequest** | [**ApprovalsEndpointsPostApprovalRequestReviewRequest**](ApprovalsEndpointsPostApprovalRequestReviewRequest.md) |  | 

### Return type

[**WebFlagConfigApprovalRequestResponse**](WebFlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlagCopyConfigApprovalRequest

> WebFlagConfigApprovalRequestResponse PostFlagCopyConfigApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey).ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest(approvalsEndpointsCreateCopyFlagConfigApprovalRequestRequest).Execute()

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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    approvalsEndpointsCreateCopyFlagConfigApprovalRequestRequest := *openapiclient.NewApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest() // ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalsApi.PostFlagCopyConfigApprovalRequest(context.Background(), projectKey, featureFlagKey, environmentKey).ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest(approvalsEndpointsCreateCopyFlagConfigApprovalRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.PostFlagCopyConfigApprovalRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFlagCopyConfigApprovalRequest`: WebFlagConfigApprovalRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.PostFlagCopyConfigApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagCopyConfigApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **approvalsEndpointsCreateCopyFlagConfigApprovalRequestRequest** | [**ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest**](ApprovalsEndpointsCreateCopyFlagConfigApprovalRequestRequest.md) |  | 

### Return type

[**WebFlagConfigApprovalRequestResponse**](WebFlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

