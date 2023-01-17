# \SegmentsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSegment**](SegmentsApi.md#DeleteSegment) | **Delete** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey} | Delete segment
[**GetExpiringUserTargetsForSegment**](SegmentsApi.md#GetExpiringUserTargetsForSegment) | **Get** /api/v2/segments/{projectKey}/{segmentKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for segment
[**GetSegment**](SegmentsApi.md#GetSegment) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey} | Get segment
[**GetSegmentMembershipForUser**](SegmentsApi.md#GetSegmentMembershipForUser) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/users/{userKey} | Get Big Segment membership for user
[**GetSegments**](SegmentsApi.md#GetSegments) | **Get** /api/v2/segments/{projectKey}/{environmentKey} | List segments
[**PatchExpiringUserTargetsForSegment**](SegmentsApi.md#PatchExpiringUserTargetsForSegment) | **Patch** /api/v2/segments/{projectKey}/{segmentKey}/expiring-user-targets/{environmentKey} | Update expiring user targets for segment
[**PatchSegment**](SegmentsApi.md#PatchSegment) | **Patch** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey} | Patch segment
[**PostSegment**](SegmentsApi.md#PostSegment) | **Post** /api/v2/segments/{projectKey}/{environmentKey} | Create segment
[**UpdateBigSegmentTargets**](SegmentsApi.md#UpdateBigSegmentTargets) | **Post** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/users | Update user context targets on a Big Segment



## DeleteSegment

> DeleteSegment(ctx, projectKey, environmentKey, segmentKey).Execute()

Delete segment



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
    resp, r, err := apiClient.SegmentsApi.DeleteSegment(context.Background(), projectKey, environmentKey, segmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.DeleteSegment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSegmentRequest struct via the builder pattern


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


## GetExpiringUserTargetsForSegment

> ExpiringUserTargetGetResponse GetExpiringUserTargetsForSegment(ctx, projectKey, environmentKey, segmentKey).Execute()

Get expiring user targets for segment



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
    resp, r, err := apiClient.SegmentsApi.GetExpiringUserTargetsForSegment(context.Background(), projectKey, environmentKey, segmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetExpiringUserTargetsForSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpiringUserTargetsForSegment`: ExpiringUserTargetGetResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetExpiringUserTargetsForSegment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetExpiringUserTargetsForSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ExpiringUserTargetGetResponse**](ExpiringUserTargetGetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegment

> UserSegment GetSegment(ctx, projectKey, environmentKey, segmentKey).Execute()

Get segment



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
    resp, r, err := apiClient.SegmentsApi.GetSegment(context.Background(), projectKey, environmentKey, segmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegment`: UserSegment
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentMembershipForUser

> BigSegmentTarget GetSegmentMembershipForUser(ctx, projectKey, environmentKey, segmentKey, userKey).Execute()

Get Big Segment membership for user



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
    userKey := "userKey_example" // string | The user key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.GetSegmentMembershipForUser(context.Background(), projectKey, environmentKey, segmentKey, userKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegmentMembershipForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentMembershipForUser`: BigSegmentTarget
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegmentMembershipForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 
**userKey** | **string** | The user key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentMembershipForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BigSegmentTarget**](BigSegmentTarget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegments

> UserSegments GetSegments(ctx, projectKey, environmentKey).Execute()

List segments



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.GetSegments(context.Background(), projectKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegments`: UserSegments
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserSegments**](UserSegments.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExpiringUserTargetsForSegment

> ExpiringUserTargetPatchResponse PatchExpiringUserTargetsForSegment(ctx, projectKey, environmentKey, segmentKey).PatchSegmentRequest(patchSegmentRequest).Execute()

Update expiring user targets for segment



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
    patchSegmentRequest := *openapiclient.NewPatchSegmentRequest([]openapiclient.PatchSegmentInstruction{*openapiclient.NewPatchSegmentInstruction("addExpireUserTargetDate", "UserKey_example", "TargetType_example")}) // PatchSegmentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.PatchExpiringUserTargetsForSegment(context.Background(), projectKey, environmentKey, segmentKey).PatchSegmentRequest(patchSegmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PatchExpiringUserTargetsForSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchExpiringUserTargetsForSegment`: ExpiringUserTargetPatchResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PatchExpiringUserTargetsForSegment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPatchExpiringUserTargetsForSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchSegmentRequest** | [**PatchSegmentRequest**](PatchSegmentRequest.md) |  | 

### Return type

[**ExpiringUserTargetPatchResponse**](ExpiringUserTargetPatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSegment

> UserSegment PatchSegment(ctx, projectKey, environmentKey, segmentKey).PatchWithComment(patchWithComment).Execute()

Patch segment



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
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.PatchSegment(context.Background(), projectKey, environmentKey, segmentKey).PatchWithComment(patchWithComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PatchSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSegment`: UserSegment
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PatchSegment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPatchSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchWithComment** | [**PatchWithComment**](PatchWithComment.md) |  | 

### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSegment

> UserSegment PostSegment(ctx, projectKey, environmentKey).SegmentBody(segmentBody).Execute()

Create segment



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
    segmentBody := *openapiclient.NewSegmentBody("Example segment", "example-segment") // SegmentBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.PostSegment(context.Background(), projectKey, environmentKey).SegmentBody(segmentBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PostSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSegment`: UserSegment
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PostSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentBody** | [**SegmentBody**](SegmentBody.md) |  | 

### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBigSegmentTargets

> UpdateBigSegmentTargets(ctx, projectKey, environmentKey, segmentKey).SegmentUserState(segmentUserState).Execute()

Update user context targets on a Big Segment



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
    segmentUserState := *openapiclient.NewSegmentUserState() // SegmentUserState | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.UpdateBigSegmentTargets(context.Background(), projectKey, environmentKey, segmentKey).SegmentUserState(segmentUserState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.UpdateBigSegmentTargets``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateBigSegmentTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **segmentUserState** | [**SegmentUserState**](SegmentUserState.md) |  | 

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

