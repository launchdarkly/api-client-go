# \FeatureFlagsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFeatureFlag**](FeatureFlagsApi.md#CopyFeatureFlag) | **Post** /api/v2/flags/{projKey}/{featureFlagKey}/copy | Copy feature flag
[**DeleteFeatureFlag**](FeatureFlagsApi.md#DeleteFeatureFlag) | **Delete** /api/v2/flags/{projKey}/{key} | Delete flag by key
[**GetDependentFlags**](FeatureFlagsApi.md#GetDependentFlags) | **Get** /api/v2/flags/{projKey}/{flagKey}/dependent-flags | List dependent feature flags
[**GetDependentFlagsByEnv**](FeatureFlagsApi.md#GetDependentFlagsByEnv) | **Get** /api/v2/flags/{projKey}/{envKey}/{flagKey}/dependent-flags | List dependent feature flags by environment
[**GetFeatureFlag**](FeatureFlagsApi.md#GetFeatureFlag) | **Get** /api/v2/flags/{projKey}/{key} | Get feature flag
[**GetFeatureFlagStatus**](FeatureFlagsApi.md#GetFeatureFlagStatus) | **Get** /api/v2/flag-statuses/{projKey}/{envKey}/{key} | List feature flag statuses
[**GetFeatureFlagStatusAcrossEnvironments**](FeatureFlagsApi.md#GetFeatureFlagStatusAcrossEnvironments) | **Get** /api/v2/flag-status/{projKey}/{key} | Get feature flag status
[**GetFeatureFlagStatuses**](FeatureFlagsApi.md#GetFeatureFlagStatuses) | **Get** /api/v2/flag-statuses/{projKey}/{envKey} | List feature flag statuses
[**GetFeatureFlags**](FeatureFlagsApi.md#GetFeatureFlags) | **Get** /api/v2/flags/{projKey} | List feature flags
[**PatchFeatureFlag**](FeatureFlagsApi.md#PatchFeatureFlag) | **Patch** /api/v2/flags/{projKey}/{key} | Update feature flag
[**PostFeatureFlag**](FeatureFlagsApi.md#PostFeatureFlag) | **Post** /api/v2/flags/{projKey} | Create a feature flag



## CopyFeatureFlag

> GlobalFlagRep CopyFeatureFlag(ctx, projKey, featureFlagKey).FlagsFlagCopyConfigPost(flagsFlagCopyConfigPost).Execute()

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
    flagsFlagCopyConfigPost := *openapiclient.NewFlagsFlagCopyConfigPost(*openapiclient.NewFlagsFlagCopyConfigEnvironment("Key_example"), *openapiclient.NewFlagsFlagCopyConfigEnvironment("Key_example")) // FlagsFlagCopyConfigPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.CopyFeatureFlag(context.Background(), projKey, featureFlagKey).FlagsFlagCopyConfigPost(flagsFlagCopyConfigPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CopyFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyFeatureFlag`: GlobalFlagRep
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


 **flagsFlagCopyConfigPost** | [**FlagsFlagCopyConfigPost**](FlagsFlagCopyConfigPost.md) |  | 

### Return type

[**GlobalFlagRep**](GlobalFlagRep.md)

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

Delete flag by key



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


## GetDependentFlags

> MultiEnvDependentFlagsCollectionRep GetDependentFlags(ctx, projKey, flagKey).Execute()

List dependent feature flags



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
    flagKey := "flagKey_example" // string | The flag key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetDependentFlags(context.Background(), projKey, flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetDependentFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDependentFlags`: MultiEnvDependentFlagsCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetDependentFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**flagKey** | **string** | The flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDependentFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MultiEnvDependentFlagsCollectionRep**](MultiEnvDependentFlagsCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDependentFlagsByEnv

> DependentFlagsCollectionRep GetDependentFlagsByEnv(ctx, projKey, envKey, flagKey).Execute()

List dependent feature flags by environment



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
    flagKey := "flagKey_example" // string | The flag key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetDependentFlagsByEnv(context.Background(), projKey, envKey, flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetDependentFlagsByEnv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDependentFlagsByEnv`: DependentFlagsCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetDependentFlagsByEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**flagKey** | **string** | The flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDependentFlagsByEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DependentFlagsCollectionRep**](DependentFlagsCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlag

> GlobalFlagRep GetFeatureFlag(ctx, projKey, key).Env(env).Execute()

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
    // response from `GetFeatureFlag`: GlobalFlagRep
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

[**GlobalFlagRep**](GlobalFlagRep.md)

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

> FlagStatusRepFromEnvSummaries GetFeatureFlagStatusAcrossEnvironments(ctx, projKey, key).Env(env).Execute()

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
    key := "key_example" // string | The feature flag key
    env := "env_example" // string | Filter configurations by environment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments(context.Background(), projKey, key).Env(env).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlagStatusAcrossEnvironments`: FlagStatusRepFromEnvSummaries
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


 **env** | **string** | Filter configurations by environment | 

### Return type

[**FlagStatusRepFromEnvSummaries**](FlagStatusRepFromEnvSummaries.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagStatuses

> FlagStatusCollectionRep GetFeatureFlagStatuses(ctx, projKey, envKey).Execute()

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
    // response from `GetFeatureFlagStatuses`: FlagStatusCollectionRep
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

[**FlagStatusCollectionRep**](FlagStatusCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlags

> GlobalFlagCollectionRep GetFeatureFlags(ctx, projKey).Env(env).Tag(tag).Limit(limit).Offset(offset).Query(query).Archived(archived).Summary(summary).Type_(type_).HasExperiment(hasExperiment).HasDataExport(hasDataExport).Evaluated(evaluated).FilterEnv(filterEnv).Execute()

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
    limit := int64(789) // int64 | The number of feature flags to return. Defaults to -1, which returns all flags. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first 10 items and then returns the next limit items. (optional)
    query := "query_example" // string | A string that matches against the flags' keys and names. It is not case sensitive. (optional)
    archived := true // bool | A boolean to filter the list to archived flags. When this is absent, only unarchived flags will be returned. (optional)
    summary := true // bool | By default in api version >= 1, flags will _not_ include their list of prerequisites, targets or rules.  Set summary=0 to include these fields for each flag returned. (optional)
    type_ := "type__example" // string | A string that lets you filter by temporary or permanent flags. (optional)
    hasExperiment := true // bool | A boolean with values of `true` or `false`. It returns any flags that have an attached metric. (optional)
    hasDataExport := true // bool | A boolean with values of `true` or `false`. That returns any flags that are currently exporting data in the specified environment. This includes flags that are exporting data from experiments. This filter also requires a `filterEnv` field, set to a valid environment key. For example, `filter=hasExperiment:true,filterEnv:production`. (optional)
    evaluated := TODO // map[string]interface{} | An object that contains a key of after and a value in Unix time in milliseconds. This returns all flags that have been evaluated since the time you specify in the environment provided. This filter also requires a `filterEnv` field to be set to a valid environment. For example, `filter=evaluated:{ (optional)
    filterEnv := "filterEnv_example" // string | A string with the key of a valid environment. The `filterEnv` field is for filters that are environment specific. If there are multiple environment specific filters, only declare this parameter once. For example, `filter=evaluated:{ (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.GetFeatureFlags(context.Background(), projKey).Env(env).Tag(tag).Limit(limit).Offset(offset).Query(query).Archived(archived).Summary(summary).Type_(type_).HasExperiment(hasExperiment).HasDataExport(hasDataExport).Evaluated(evaluated).FilterEnv(filterEnv).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlags`: GlobalFlagCollectionRep
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
 **limit** | **int64** | The number of feature flags to return. Defaults to -1, which returns all flags. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first 10 items and then returns the next limit items. | 
 **query** | **string** | A string that matches against the flags&#39; keys and names. It is not case sensitive. | 
 **archived** | **bool** | A boolean to filter the list to archived flags. When this is absent, only unarchived flags will be returned. | 
 **summary** | **bool** | By default in api version &gt;&#x3D; 1, flags will _not_ include their list of prerequisites, targets or rules.  Set summary&#x3D;0 to include these fields for each flag returned. | 
 **type_** | **string** | A string that lets you filter by temporary or permanent flags. | 
 **hasExperiment** | **bool** | A boolean with values of &#x60;true&#x60; or &#x60;false&#x60;. It returns any flags that have an attached metric. | 
 **hasDataExport** | **bool** | A boolean with values of &#x60;true&#x60; or &#x60;false&#x60;. That returns any flags that are currently exporting data in the specified environment. This includes flags that are exporting data from experiments. This filter also requires a &#x60;filterEnv&#x60; field, set to a valid environment key. For example, &#x60;filter&#x3D;hasExperiment:true,filterEnv:production&#x60;. | 
 **evaluated** | [**map[string]interface{}**](map[string]interface{}.md) | An object that contains a key of after and a value in Unix time in milliseconds. This returns all flags that have been evaluated since the time you specify in the environment provided. This filter also requires a &#x60;filterEnv&#x60; field to be set to a valid environment. For example, &#x60;filter&#x3D;evaluated:{ | 
 **filterEnv** | **string** | A string with the key of a valid environment. The &#x60;filterEnv&#x60; field is for filters that are environment specific. If there are multiple environment specific filters, only declare this parameter once. For example, &#x60;filter&#x3D;evaluated:{ | 

### Return type

[**GlobalFlagCollectionRep**](GlobalFlagCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFeatureFlag

> GlobalFlagRep PatchFeatureFlag(ctx, projKey, key).PatchWithComment(patchWithComment).Execute()

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
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.JSONPatchElt{*openapiclient.NewJSONPatchElt("replace", "/biscuits", interface{}(Chocolate Digestive))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.PatchFeatureFlag(context.Background(), projKey, key).PatchWithComment(patchWithComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PatchFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFeatureFlag`: GlobalFlagRep
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

[**GlobalFlagRep**](GlobalFlagRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeatureFlag

> GlobalFlagRep PostFeatureFlag(ctx, projKey).FlagPost(flagPost).Clone(clone).Execute()

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
    flagPost := *openapiclient.NewFlagPost("Name_example", "Key_example") // FlagPost | 
    clone := "clone_example" // string | The key of the feature flag to be cloned. The key identifies the flag in your code. For example, setting `clone=flagKey` copies the full targeting configuration for all environments, including `on/off` state, from the original flag to the new flag. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureFlagsApi.PostFeatureFlag(context.Background(), projKey).FlagPost(flagPost).Clone(clone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PostFeatureFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFeatureFlag`: GlobalFlagRep
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

 **flagPost** | [**FlagPost**](FlagPost.md) |  | 
 **clone** | **string** | The key of the feature flag to be cloned. The key identifies the flag in your code. For example, setting &#x60;clone&#x3D;flagKey&#x60; copies the full targeting configuration for all environments, including &#x60;on/off&#x60; state, from the original flag to the new flag. | 

### Return type

[**GlobalFlagRep**](GlobalFlagRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

