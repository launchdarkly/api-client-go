# \EnvironmentsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEnvironment**](EnvironmentsApi.md#DeleteEnvironment) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey} | Delete environment
[**GetEnvironment**](EnvironmentsApi.md#GetEnvironment) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey} | Get environment
[**PatchEnvironment**](EnvironmentsApi.md#PatchEnvironment) | **Patch** /api/v2/projects/{projectKey}/environments/{environmentKey} | Update environment
[**PostEnvironment**](EnvironmentsApi.md#PostEnvironment) | **Post** /api/v2/projects/{projectKey}/environments | Create environment
[**PostEnvironmentKey**](EnvironmentsApi.md#PostEnvironmentKey) | **Post** /api/v2/projects/{projectKey}/environments/{envKey}/apiKey | Create environment SDK key



## DeleteEnvironment

> DeleteEnvironment(ctx, projectKey, environmentKey).Execute()

Delete environment



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.DeleteEnvironment(context.Background(), projectKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironment``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## GetEnvironment

> Environment GetEnvironment(ctx, projectKey, environmentKey).Execute()

Get environment



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.GetEnvironment(context.Background(), projectKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEnvironment

> Environment PatchEnvironment(ctx, projectKey, environmentKey).PatchOperation(patchOperation).Execute()

Update environment



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
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/biscuits", interface{}(Chocolate Digestive))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.PatchEnvironment(context.Background(), projectKey, environmentKey).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.PatchEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.PatchEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnvironment

> Environment PostEnvironment(ctx, projectKey).EnvironmentPost(environmentPost).Execute()

Create environment



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
    environmentPost := *openapiclient.NewEnvironmentPost("My Environment", "my-environment", "F5A623") // EnvironmentPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.PostEnvironment(context.Background(), projectKey).EnvironmentPost(environmentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.PostEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.PostEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentPost** | [**EnvironmentPost**](EnvironmentPost.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnvironmentKey

> Environment PostEnvironmentKey(ctx, projectKey, envKey).Expiry(expiry).Execute()

Create environment SDK key



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
    envKey := "envKey_example" // string | The environment key
    expiry := int64(789) // int64 | The time at which you want the old SDK key to expire, in UNIX milliseconds. By default, the key expires immediately. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EnvironmentsApi.PostEnvironmentKey(context.Background(), projectKey, envKey).Expiry(expiry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.PostEnvironmentKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostEnvironmentKey`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.PostEnvironmentKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEnvironmentKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expiry** | **int64** | The time at which you want the old SDK key to expire, in UNIX milliseconds. By default, the key expires immediately. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

