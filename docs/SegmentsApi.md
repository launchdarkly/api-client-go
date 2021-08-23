# \SegmentsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSegment**](SegmentsApi.md#DeleteSegment) | **Delete** /api/v2/segments/{projKey}/{envKey}/{key} | Delete segment
[**GetBigSegmentTarget**](SegmentsApi.md#GetBigSegmentTarget) | **Get** /api/v2/segments/{projKey}/{envKey}/{key}/users/{userKey} | Get user in Big Segment
[**GetExpiringUserTargetsOnSegment**](SegmentsApi.md#GetExpiringUserTargetsOnSegment) | **Get** /api/v2/segments/{projKey}/{segmentKey}/expiring-user-targets/{envKey} | Get expiring user targets for segment
[**GetSegment**](SegmentsApi.md#GetSegment) | **Get** /api/v2/segments/{projKey}/{envKey}/{key} | Get segment
[**GetSegments**](SegmentsApi.md#GetSegments) | **Get** /api/v2/segments/{projKey}/{envKey} | List segments
[**PatchExpiringUserTargetsOnSegment**](SegmentsApi.md#PatchExpiringUserTargetsOnSegment) | **Patch** /api/v2/segments/{projKey}/{segmentKey}/expiring-user-targets/{envKey} | Update expiring user targets on segment
[**PatchSegment**](SegmentsApi.md#PatchSegment) | **Patch** /api/v2/segments/{projKey}/{envKey}/{key} | Patch segment
[**PostSegment**](SegmentsApi.md#PostSegment) | **Post** /api/v2/segments/{projKey}/{envKey} | Create segment
[**UpdateBigSegmentTargets**](SegmentsApi.md#UpdateBigSegmentTargets) | **Post** /api/v2/segments/{projKey}/{envKey}/{key}/users | Update targets on a Big Segment



## DeleteSegment

> DeleteSegment(ctx, projKey, envKey, key).Execute()

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
    projKey := "projKey_example" // string | The project key. This connects flags within one project so you can manage them together.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users within one environment so you can manage them together.
    key := "key_example" // string | The user segment key. The key identifies the user segment in your code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.DeleteSegment(context.Background(), projKey, envKey, key).Execute()
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
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 
**key** | **string** | The user segment key. The key identifies the user segment in your code. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBigSegmentTarget

> BigSegmentTargetChanges GetBigSegmentTarget(ctx, projKey, envKey, key, userKey).Execute()

Get user in Big Segment



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
    projKey := "projKey_example" // string | The project key.
    envKey := "envKey_example" // string | The environment key.
    key := "key_example" // string | The segment key.
    userKey := "userKey_example" // string | The user key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.GetBigSegmentTarget(context.Background(), projKey, envKey, key, userKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetBigSegmentTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBigSegmentTarget`: BigSegmentTargetChanges
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetBigSegmentTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. | 
**envKey** | **string** | The environment key. | 
**key** | **string** | The segment key. | 
**userKey** | **string** | The user key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBigSegmentTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BigSegmentTargetChanges**](BigSegmentTargetChanges.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpiringUserTargetsOnSegment

> ExpiringUserTargetGetResponse GetExpiringUserTargetsOnSegment(ctx, projKey, envKey, segmentKey).Execute()

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
    projKey := "projKey_example" // string | The project key. This connects flags within one project so you can manage them together.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users within one environment so you can manage them together.
    segmentKey := "segmentKey_example" // string | The user segment key. The key identifies the user segment in your code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.GetExpiringUserTargetsOnSegment(context.Background(), projKey, envKey, segmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetExpiringUserTargetsOnSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpiringUserTargetsOnSegment`: ExpiringUserTargetGetResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetExpiringUserTargetsOnSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 
**segmentKey** | **string** | The user segment key. The key identifies the user segment in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpiringUserTargetsOnSegmentRequest struct via the builder pattern


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

> UserSegment GetSegment(ctx, projKey, envKey, key).Execute()

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
    projKey := "projKey_example" // string | The project key. This connects flags within one project so you can manage them together.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users under one environment so you can manage them together.
    key := "key_example" // string | The segment key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.GetSegment(context.Background(), projKey, envKey, key).Execute()
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
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users under one environment so you can manage them together. | 
**key** | **string** | The segment key | 

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


## GetSegments

> UserSegments GetSegments(ctx, projKey, envKey).Execute()

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
    projKey := "projKey_example" // string | The project key. This connects flags within one project so you can manage them together.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users within one environment so you can manage them together.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.GetSegments(context.Background(), projKey, envKey).Execute()
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
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 

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


## PatchExpiringUserTargetsOnSegment

> ExpiringUserTargetPatchResponse PatchExpiringUserTargetsOnSegment(ctx, projKey, envKey, segmentKey).PatchSegmentRequest(patchSegmentRequest).Execute()

Update expiring user targets on segment



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
    projKey := "projKey_example" // string | The project key. This connects flags within one project so you can manage them together.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users within one environment so you can manage them together.
    segmentKey := "segmentKey_example" // string | The user segment key. The key identifies the user segment in your code.
    patchSegmentRequest := *openapiclient.NewPatchSegmentRequest([]openapiclient.PatchSegmentInstruction{*openapiclient.NewPatchSegmentInstruction("Kind_example", "UserKey_example", "TargetType_example")}) // PatchSegmentRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.PatchExpiringUserTargetsOnSegment(context.Background(), projKey, envKey, segmentKey).PatchSegmentRequest(patchSegmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PatchExpiringUserTargetsOnSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchExpiringUserTargetsOnSegment`: ExpiringUserTargetPatchResponse
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PatchExpiringUserTargetsOnSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 
**segmentKey** | **string** | The user segment key. The key identifies the user segment in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringUserTargetsOnSegmentRequest struct via the builder pattern


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

> UserSegment PatchSegment(ctx, projKey, envKey, key).PatchWithComment(patchWithComment).Execute()

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
    projKey := "projKey_example" // string | The project key. This connects flags within one project so you can manage them together.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users under one environment so you can manage them together.
    key := "key_example" // string | The user segment key. The key identifies the user segment in your code.
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.JSONPatchElt{*openapiclient.NewJSONPatchElt("replace", "/biscuits", interface{}(Chocolate Digestive))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.PatchSegment(context.Background(), projKey, envKey, key).PatchWithComment(patchWithComment).Execute()
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
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users under one environment so you can manage them together. | 
**key** | **string** | The user segment key. The key identifies the user segment in your code. | 

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

> UserSegment PostSegment(ctx, projKey, envKey).SegmentPost(segmentPost).Execute()

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
    projKey := "projKey_example" // string | The project key. This connects flags within one project so you can manage them together.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users within one environment so you can manage them together.
    segmentPost := *openapiclient.NewSegmentPost("Name_example", "Key_example") // SegmentPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.PostSegment(context.Background(), projKey, envKey).SegmentPost(segmentPost).Execute()
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
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentPost** | [**SegmentPost**](SegmentPost.md) |  | 

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

> UpdateBigSegmentTargets(ctx, projKey, envKey, key).SegmentUserState(segmentUserState).Execute()

Update targets on a Big Segment



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
    projKey := "projKey_example" // string | The project key.
    envKey := "envKey_example" // string | The environment key.
    key := "key_example" // string | The segment key.
    segmentUserState := *openapiclient.NewSegmentUserState() // SegmentUserState | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SegmentsApi.UpdateBigSegmentTargets(context.Background(), projKey, envKey, key).SegmentUserState(segmentUserState).Execute()
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
**projKey** | **string** | The project key. | 
**envKey** | **string** | The environment key. | 
**key** | **string** | The segment key. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

