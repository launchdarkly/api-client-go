# \FeatureFlagsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFeatureFlag**](FeatureFlagsApi.md#CopyFeatureFlag) | **Post** /api/v2/flags/{projKey}/{featureFlagKey}/copy | Copy feature flag
[**DeleteFeatureFlag**](FeatureFlagsApi.md#DeleteFeatureFlag) | **Delete** /api/v2/flags/{projKey}/{key} | Delete feature flag
[**GetExpiringUserTargets**](FeatureFlagsApi.md#GetExpiringUserTargets) | **Get** /api/v2/flags/{projKey}/{flagKey}/expiring-user-targets/{envKey} | Get expiring user targets for feature flag
[**GetFeatureFlag**](FeatureFlagsApi.md#GetFeatureFlag) | **Get** /api/v2/flags/{projKey}/{key} | Get feature flag
[**GetFeatureFlagStatus**](FeatureFlagsApi.md#GetFeatureFlagStatus) | **Get** /api/v2/flag-statuses/{projKey}/{envKey}/{key} | Get feature flag status
[**GetFeatureFlagStatusAcrossEnvironments**](FeatureFlagsApi.md#GetFeatureFlagStatusAcrossEnvironments) | **Get** /api/v2/flag-status/{projKey}/{key} | Get flag status across environments
[**GetFeatureFlagStatuses**](FeatureFlagsApi.md#GetFeatureFlagStatuses) | **Get** /api/v2/flag-statuses/{projKey}/{envKey} | List feature flag statuses
[**GetFeatureFlags**](FeatureFlagsApi.md#GetFeatureFlags) | **Get** /api/v2/flags/{projKey} | List feature flags
[**PatchExpiringUserTargets**](FeatureFlagsApi.md#PatchExpiringUserTargets) | **Patch** /api/v2/flags/{projKey}/{flagKey}/expiring-user-targets/{envKey} | Update expiring user targets on feature flag
[**PatchFeatureFlag**](FeatureFlagsApi.md#PatchFeatureFlag) | **Patch** /api/v2/flags/{projKey}/{key} | Update feature flag
[**PostFeatureFlag**](FeatureFlagsApi.md#PostFeatureFlag) | **Post** /api/v2/flags/{projKey} | Create a feature flag



## CopyFeatureFlag

> FeatureFlag CopyFeatureFlag(ctx, projKey, featureFlagKey).FlagCopyConfigPost(flagCopyConfigPost).Execute()

Copy feature flag



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key. The key identifies the flag in your code.
    flagCopyConfigPost := *openapiclient.NewFlagCopyConfigPost(*openapiclient.NewFlagCopyConfigEnvironment("Key_example"), *openapiclient.NewFlagCopyConfigEnvironment("Key_example")) // FlagCopyConfigPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.CopyFeatureFlag(context.Background(), projKey, featureFlagKey).FlagCopyConfigPost(flagCopyConfigPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CopyFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyFeatureFlag`: FeatureFlag
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.CopyFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**featureFlagKey** | **string** | The feature flag&#39;s key. The key identifies the flag in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flagCopyConfigPost** | [**FlagCopyConfigPost**](FlagCopyConfigPost.md) |  | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeatureFlag

> DeleteFeatureFlag(ctx, projKey, key).Execute()

Delete feature flag



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
    key := "key_example" // string | The feature flag's key. The key identifies the flag in your code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.DeleteFeatureFlag(context.Background(), projKey, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.DeleteFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**key** | **string** | The feature flag&#39;s key. The key identifies the flag in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureFlagRequest struct via the builder pattern


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


## GetExpiringUserTargets

> ExpiringUserTargetGetResponse GetExpiringUserTargets(ctx, projKey, envKey, flagKey).Execute()

Get expiring user targets for feature flag



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
    flagKey := "flagKey_example" // string | The feature flag key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetExpiringUserTargets(context.Background(), projKey, envKey, flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetExpiringUserTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpiringUserTargets`: ExpiringUserTargetGetResponse
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetExpiringUserTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 
**flagKey** | **string** | The feature flag key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpiringUserTargetsRequest struct via the builder pattern


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


## GetFeatureFlag

> FeatureFlag GetFeatureFlag(ctx, projKey, key).Env(env).Execute()

Get feature flag



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
    key := "key_example" // string | The feature flag key
    env := "env_example" // string | Filter configurations by environment (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetFeatureFlag(context.Background(), projKey, key).Env(env).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlag`: FeatureFlag
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**key** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **string** | Filter configurations by environment | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagStatus

> FlagStatusRep GetFeatureFlagStatus(ctx, projKey, envKey, key).Execute()

Get feature flag status



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
    key := "key_example" // string | The feature flag key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetFeatureFlagStatus(context.Background(), projKey, envKey, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlagStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlagStatus`: FlagStatusRep
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlagStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**key** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlagStatusRep**](FlagStatusRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagStatusAcrossEnvironments

> FeatureFlagStatusAcrossEnvironments GetFeatureFlagStatusAcrossEnvironments(ctx, projKey, key).Env(env).Execute()

Get flag status across environments



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
    key := "key_example" // string | The feature flag key
    env := "env_example" // string | Optional environment filter (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments(context.Background(), projKey, key).Env(env).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlagStatusAcrossEnvironments`: FeatureFlagStatusAcrossEnvironments
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**key** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagStatusAcrossEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **string** | Optional environment filter | 

### Return type

[**FeatureFlagStatusAcrossEnvironments**](FeatureFlagStatusAcrossEnvironments.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagStatuses

> FeatureFlagStatuses GetFeatureFlagStatuses(ctx, projKey, envKey).Execute()

List feature flag statuses



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
    envKey := "envKey_example" // string | Filter configurations by environment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetFeatureFlagStatuses(context.Background(), projKey, envKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlagStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlagStatuses`: FeatureFlagStatuses
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlagStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | Filter configurations by environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FeatureFlagStatuses**](FeatureFlagStatuses.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlags

> FeatureFlags GetFeatureFlags(ctx, projKey).Env(env).Tag(tag).Limit(limit).Offset(offset).Query(query).Archived(archived).Summary(summary).Filter(filter).Sort(sort).Execute()

List feature flags



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
    env := "env_example" // string | Filter configurations by environment (optional)
    tag := "tag_example" // string | Filter feature flags by tag (optional)
    limit := int64(789) // int64 | The number of feature flags to return. Defaults to -1, which returns all flags (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next limit items (optional)
    query := "query_example" // string | A string that matches against the flags' keys and names. It is not case sensitive (optional)
    archived := true // bool | A boolean to filter the list to archived flags. When this is absent, only unarchived flags will be returned (optional)
    summary := true // bool | By default in API version >= 1, flags will _not_ include their list of prerequisites, targets or rules.  Set summary=0 to include these fields for each flag returned (optional)
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form field:value (optional)
    sort := "sort_example" // string | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetFeatureFlags(context.Background(), projKey).Env(env).Tag(tag).Limit(limit).Offset(offset).Query(query).Archived(archived).Summary(summary).Filter(filter).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlags`: FeatureFlags
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Filter configurations by environment | 
 **tag** | **string** | Filter feature flags by tag | 
 **limit** | **int64** | The number of feature flags to return. Defaults to -1, which returns all flags | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next limit items | 
 **query** | **string** | A string that matches against the flags&#39; keys and names. It is not case sensitive | 
 **archived** | **bool** | A boolean to filter the list to archived flags. When this is absent, only unarchived flags will be returned | 
 **summary** | **bool** | By default in API version &gt;&#x3D; 1, flags will _not_ include their list of prerequisites, targets or rules.  Set summary&#x3D;0 to include these fields for each flag returned | 
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form field:value | 
 **sort** | **string** | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order | 

### Return type

[**FeatureFlags**](FeatureFlags.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExpiringUserTargets

> ExpiringUserTargetPatchResponse PatchExpiringUserTargets(ctx, projKey, envKey, flagKey).PatchWithComment(patchWithComment).Execute()

Update expiring user targets on feature flag



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
    flagKey := "flagKey_example" // string | The feature flag key.
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/biscuits", interface{}(Chocolate Digestive))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.PatchExpiringUserTargets(context.Background(), projKey, envKey, flagKey).PatchWithComment(patchWithComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PatchExpiringUserTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchExpiringUserTargets`: ExpiringUserTargetPatchResponse
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PatchExpiringUserTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 
**flagKey** | **string** | The feature flag key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringUserTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchWithComment** | [**PatchWithComment**](PatchWithComment.md) |  | 

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


## PatchFeatureFlag

> FeatureFlag PatchFeatureFlag(ctx, projKey, key).PatchWithComment(patchWithComment).Execute()

Update feature flag



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
    key := "key_example" // string | The feature flag's key. The key identifies the flag in your code.
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/biscuits", interface{}(Chocolate Digestive))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.PatchFeatureFlag(context.Background(), projKey, key).PatchWithComment(patchWithComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PatchFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFeatureFlag`: FeatureFlag
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PatchFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**key** | **string** | The feature flag&#39;s key. The key identifies the flag in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchWithComment** | [**PatchWithComment**](PatchWithComment.md) |  | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeatureFlag

> FeatureFlag PostFeatureFlag(ctx, projKey).FeatureFlagBody(featureFlagBody).Clone(clone).Execute()

Create a feature flag



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
    featureFlagBody := *openapiclient.NewFeatureFlagBody("Name_example", "Key_example") // FeatureFlagBody | 
    clone := "clone_example" // string | The key of the feature flag to be cloned. The key identifies the flag in your code. For example, setting `clone=flagKey` copies the full targeting configuration for all environments, including `on/off` state, from the original flag to the new flag. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.PostFeatureFlag(context.Background(), projKey).FeatureFlagBody(featureFlagBody).Clone(clone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PostFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFeatureFlag`: FeatureFlag
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PostFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureFlagBody** | [**FeatureFlagBody**](FeatureFlagBody.md) |  | 
 **clone** | **string** | The key of the feature flag to be cloned. The key identifies the flag in your code. For example, setting &#x60;clone&#x3D;flagKey&#x60; copies the full targeting configuration for all environments, including &#x60;on/off&#x60; state, from the original flag to the new flag. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

