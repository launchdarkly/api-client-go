# \UserSettingsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserFlagSetting**](UserSettingsApi.md#GetUserFlagSetting) | **Get** /api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey} | Get flag setting for user
[**GetUserFlagSettings**](UserSettingsApi.md#GetUserFlagSettings) | **Get** /api/v2/users/{projKey}/{envKey}/{key}/flags | List flag settings for user
[**PutFlagSetting**](UserSettingsApi.md#PutFlagSetting) | **Put** /api/v2/users/{projKey}/{envKey}/{key}/flags/{featureKey} | Update flag settings for user



## GetUserFlagSetting

> UserSettingRep GetUserFlagSetting(ctx, projKey, envKey, key, featureKey).Execute()

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
    // response from `GetUserFlagSetting`: UserSettingRep
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

[**UserSettingRep**](UserSettingRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFlagSettings

> UserSettingsCollection GetUserFlagSettings(ctx, projKey, envKey, key).Execute()

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
    // response from `GetUserFlagSettings`: UserSettingsCollection
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

[**UserSettingsCollection**](UserSettingsCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFlagSetting

> PutFlagSetting(ctx, projKey, envKey, key, featureKey).Api2ValuePut(api2ValuePut).Execute()

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
    api2ValuePut := *openapiclient.NewApi2ValuePut() // Api2ValuePut | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.PutFlagSetting(context.Background(), projKey, envKey, key, featureKey).Api2ValuePut(api2ValuePut).Execute()
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




 **api2ValuePut** | [**Api2ValuePut**](Api2ValuePut.md) |  | 

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

