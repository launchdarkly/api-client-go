# \FollowFlagsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFlagFollowers**](FollowFlagsBetaApi.md#DeleteFlagFollowers) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/followers/{memberId} | Remove a member as a follower to a flag in a project and environment
[**GetFlagFollowers**](FollowFlagsBetaApi.md#GetFlagFollowers) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/followers | Get followers of a flag in a project and environment
[**PutFlagFollowers**](FollowFlagsBetaApi.md#PutFlagFollowers) | **Put** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/followers/{memberId} | Add a member as a follower to a flag in a project and environment



## DeleteFlagFollowers

> DeleteFlagFollowers(ctx, projectKey, featureFlagKey, environmentKey, memberId).Execute()

Remove a member as a follower to a flag in a project and environment



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
    memberId := "memberId_example" // string | memberId of the member following the flag resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FollowFlagsBetaApi.DeleteFlagFollowers(context.Background(), projectKey, featureFlagKey, environmentKey, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowFlagsBetaApi.DeleteFlagFollowers``: %v\n", err)
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
**memberId** | **string** | memberId of the member following the flag resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagFollowersRequest struct via the builder pattern


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


## GetFlagFollowers

> FlagFollowersGetRep GetFlagFollowers(ctx, projectKey, featureFlagKey, environmentKey).Execute()

Get followers of a flag in a project and environment



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FollowFlagsBetaApi.GetFlagFollowers(context.Background(), projectKey, featureFlagKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowFlagsBetaApi.GetFlagFollowers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagFollowers`: FlagFollowersGetRep
    fmt.Fprintf(os.Stdout, "Response from `FollowFlagsBetaApi.GetFlagFollowers`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFlagFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlagFollowersGetRep**](FlagFollowersGetRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFlagFollowers

> PutFlagFollowers(ctx, projectKey, featureFlagKey, environmentKey, memberId).Execute()

Add a member as a follower to a flag in a project and environment



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
    memberId := "memberId_example" // string | memberId of the member following the flag resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FollowFlagsBetaApi.PutFlagFollowers(context.Background(), projectKey, featureFlagKey, environmentKey, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FollowFlagsBetaApi.PutFlagFollowers``: %v\n", err)
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
**memberId** | **string** | memberId of the member following the flag resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFlagFollowersRequest struct via the builder pattern


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

