# \UserSettingsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpiringFlagsForUser**](UserSettingsApi.md#GetExpiringFlagsForUser) | **Get** /api/v2/users/{projKey}/{userKey}/expiring-user-targets/{envKey} | Get expiring dates on flags for user
[**GetUserFlagSetting**](UserSettingsApi.md#GetUserFlagSetting) | **Get** /api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey} | Get flag setting for user
[**GetUserFlagSettings**](UserSettingsApi.md#GetUserFlagSettings) | **Get** /api/v2/users/{projKey}/{envKey}/{key}/flags | List flag settings for user
[**PatchExpiringFlagsForUser**](UserSettingsApi.md#PatchExpiringFlagsForUser) | **Patch** /api/v2/users/{projKey}/{userKey}/expiring-user-targets/{envKey} | Update expiring user target for flags
[**PutFlagSetting**](UserSettingsApi.md#PutFlagSetting) | **Put** /api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey} | Update flag settings for user



## GetExpiringFlagsForUser

> ExpiringUserTargetGetResponse GetExpiringFlagsForUser(ctx, projKey, userKey, envKey).Execute()

Get expiring dates on flags for user



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
    userKey := "userKey_example" // string | The user key.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users within one environment so you can manage them together.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.GetExpiringFlagsForUser(context.Background(), projKey, userKey, envKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsApi.GetExpiringFlagsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpiringFlagsForUser`: ExpiringUserTargetGetResponse
    fmt.Fprintf(os.Stdout, "Response from `UserSettingsApi.GetExpiringFlagsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**userKey** | **string** | The user key. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpiringFlagsForUserRequest struct via the builder pattern


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


## GetUserFlagSetting

> UserFlagSetting GetUserFlagSetting(ctx, projKey, envKey, key, featureKey).Execute()

Get flag setting for user



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
    key := "key_example" // string | The user key
    featureKey := "featureKey_example" // string | The feature flag key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.GetUserFlagSetting(context.Background(), projKey, envKey, key, featureKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsApi.GetUserFlagSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserFlagSetting`: UserFlagSetting
    fmt.Fprintf(os.Stdout, "Response from `UserSettingsApi.GetUserFlagSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**key** | **string** | The user key | 
**featureKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFlagSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UserFlagSetting**](UserFlagSetting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFlagSettings

> UserFlagSettings GetUserFlagSettings(ctx, projKey, envKey, key).Execute()

List flag settings for user



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
    key := "key_example" // string | The user key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.GetUserFlagSettings(context.Background(), projKey, envKey, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsApi.GetUserFlagSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserFlagSettings`: UserFlagSettings
    fmt.Fprintf(os.Stdout, "Response from `UserSettingsApi.GetUserFlagSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**key** | **string** | The user key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFlagSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserFlagSettings**](UserFlagSettings.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExpiringFlagsForUser

> ExpiringUserTargetPatchResponse PatchExpiringFlagsForUser(ctx, projKey, userKey, envKey).PatchWithComment(patchWithComment).Execute()

Update expiring user target for flags



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
    userKey := "userKey_example" // string | The user key.
    envKey := "envKey_example" // string | The environment key. This connects flag configurations and users within one environment so you can manage them together.
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/biscuits", interface{}(Chocolate Digestive))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.PatchExpiringFlagsForUser(context.Background(), projKey, userKey, envKey).PatchWithComment(patchWithComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsApi.PatchExpiringFlagsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchExpiringFlagsForUser`: ExpiringUserTargetPatchResponse
    fmt.Fprintf(os.Stdout, "Response from `UserSettingsApi.PatchExpiringFlagsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key. This connects flags within one project so you can manage them together. | 
**userKey** | **string** | The user key. | 
**envKey** | **string** | The environment key. This connects flag configurations and users within one environment so you can manage them together. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringFlagsForUserRequest struct via the builder pattern


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


## PutFlagSetting

> PutFlagSetting(ctx, projKey, envKey, key, featureKey).ValuePut(valuePut).Execute()

Update flag settings for user



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
    key := "key_example" // string | The user key
    featureKey := "featureKey_example" // string | The feature flag key
    valuePut := *openapiclient.NewValuePut() // ValuePut | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.PutFlagSetting(context.Background(), projKey, envKey, key, featureKey).ValuePut(valuePut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsApi.PutFlagSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**key** | **string** | The user key | 
**featureKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFlagSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **valuePut** | [**ValuePut**](ValuePut.md) |  | 

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

