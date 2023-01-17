# \UserSettingsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpiringFlagsForUser**](UserSettingsApi.md#GetExpiringFlagsForUser) | **Get** /api/v2/users/{projectKey}/{userKey}/expiring-user-targets/{environmentKey} | Get expiring dates on flags for user
[**GetUserFlagSetting**](UserSettingsApi.md#GetUserFlagSetting) | **Get** /api/v2/users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Get flag setting for user
[**GetUserFlagSettings**](UserSettingsApi.md#GetUserFlagSettings) | **Get** /api/v2/users/{projectKey}/{environmentKey}/{userKey}/flags | List flag settings for user
[**PatchExpiringFlagsForUser**](UserSettingsApi.md#PatchExpiringFlagsForUser) | **Patch** /api/v2/users/{projectKey}/{userKey}/expiring-user-targets/{environmentKey} | Update expiring user target for flags
[**PutFlagSetting**](UserSettingsApi.md#PutFlagSetting) | **Put** /api/v2/users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Update flag settings for user



## GetExpiringFlagsForUser

> ExpiringUserTargetGetResponse GetExpiringFlagsForUser(ctx, projectKey, userKey, environmentKey).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    userKey := "userKey_example" // string | The user key
    environmentKey := "environmentKey_example" // string | The environment key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSettingsApi.GetExpiringFlagsForUser(context.Background(), projectKey, userKey, environmentKey).Execute()
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
**projectKey** | **string** | The project key | 
**userKey** | **string** | The user key | 
**environmentKey** | **string** | The environment key | 

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

> UserFlagSetting GetUserFlagSetting(ctx, projectKey, environmentKey, userKey, featureFlagKey).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    environmentKey := "environmentKey_example" // string | The environment key
    userKey := "userKey_example" // string | The user key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSettingsApi.GetUserFlagSetting(context.Background(), projectKey, environmentKey, userKey, featureFlagKey).Execute()
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
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**userKey** | **string** | The user key | 
**featureFlagKey** | **string** | The feature flag key | 

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

> UserFlagSettings GetUserFlagSettings(ctx, projectKey, environmentKey, userKey).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    environmentKey := "environmentKey_example" // string | The environment key
    userKey := "userKey_example" // string | The user key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSettingsApi.GetUserFlagSettings(context.Background(), projectKey, environmentKey, userKey).Execute()
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
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**userKey** | **string** | The user key | 

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

> ExpiringUserTargetPatchResponse PatchExpiringFlagsForUser(ctx, projectKey, userKey, environmentKey).PatchUsersRequest(patchUsersRequest).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    userKey := "userKey_example" // string | The user key
    environmentKey := "environmentKey_example" // string | The environment key
    patchUsersRequest := *openapiclient.NewPatchUsersRequest([]openapiclient.InstructionUserRequest{*openapiclient.NewInstructionUserRequest("addExpireUserTargetDate", "sample-flag-key", "ce12d345-a1b2-4fb5-a123-ab123d4d5f5d")}) // PatchUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSettingsApi.PatchExpiringFlagsForUser(context.Background(), projectKey, userKey, environmentKey).PatchUsersRequest(patchUsersRequest).Execute()
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
**projectKey** | **string** | The project key | 
**userKey** | **string** | The user key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringFlagsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchUsersRequest** | [**PatchUsersRequest**](PatchUsersRequest.md) |  | 

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

> PutFlagSetting(ctx, projectKey, environmentKey, userKey, featureFlagKey).ValuePut(valuePut).Execute()

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
    projectKey := "projectKey_example" // string | The project key
    environmentKey := "environmentKey_example" // string | The environment key
    userKey := "userKey_example" // string | The user key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    valuePut := *openapiclient.NewValuePut() // ValuePut | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSettingsApi.PutFlagSetting(context.Background(), projectKey, environmentKey, userKey, featureFlagKey).ValuePut(valuePut).Execute()
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
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**userKey** | **string** | The user key | 
**featureFlagKey** | **string** | The feature flag key | 

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

