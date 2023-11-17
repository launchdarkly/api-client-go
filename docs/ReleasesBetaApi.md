# \ReleasesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReleaseByFlagKey**](ReleasesBetaApi.md#GetReleaseByFlagKey) | **Get** /api/v2/flags/{projectKey}/{flagKey}/release | Get release for flag
[**PatchReleaseByFlagKey**](ReleasesBetaApi.md#PatchReleaseByFlagKey) | **Patch** /api/v2/flags/{projectKey}/{flagKey}/release | Patch release for flag



## GetReleaseByFlagKey

> Release GetReleaseByFlagKey(ctx, projectKey, flagKey).Execute()

Get release for flag



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
    flagKey := "flagKey_example" // string | The flag key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesBetaApi.GetReleaseByFlagKey(context.Background(), projectKey, flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesBetaApi.GetReleaseByFlagKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleaseByFlagKey`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleasesBetaApi.GetReleaseByFlagKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**flagKey** | **string** | The flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseByFlagKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Release**](Release.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchReleaseByFlagKey

> PatchReleaseByFlagKey(ctx, projectKey, flagKey).PatchOperation(patchOperation).Execute()

Patch release for flag



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
    flagKey := "flagKey_example" // string | The flag key
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesBetaApi.PatchReleaseByFlagKey(context.Background(), projectKey, flagKey).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesBetaApi.PatchReleaseByFlagKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**flagKey** | **string** | The flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchReleaseByFlagKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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

