# \ReleasesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReleaseForFlag**](ReleasesBetaApi.md#CreateReleaseForFlag) | **Put** /api/v2/projects/{projectKey}/flags/{flagKey}/release | Create a new release for flag
[**DeleteReleaseByFlagKey**](ReleasesBetaApi.md#DeleteReleaseByFlagKey) | **Delete** /api/v2/flags/{projectKey}/{flagKey}/release | Delete a release for flag
[**GetReleaseByFlagKey**](ReleasesBetaApi.md#GetReleaseByFlagKey) | **Get** /api/v2/flags/{projectKey}/{flagKey}/release | Get release for flag
[**PatchReleaseByFlagKey**](ReleasesBetaApi.md#PatchReleaseByFlagKey) | **Patch** /api/v2/flags/{projectKey}/{flagKey}/release | Patch release for flag
[**UpdatePhaseStatus**](ReleasesBetaApi.md#UpdatePhaseStatus) | **Put** /api/v2/projects/{projectKey}/flags/{flagKey}/release/phases/{phaseId} | Update phase status for release



## CreateReleaseForFlag

> Release CreateReleaseForFlag(ctx, projectKey, flagKey).CreateReleaseInput(createReleaseInput).Execute()

Create a new release for flag



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
    createReleaseInput := *openapiclient.NewCreateReleaseInput("ReleasePipelineKey_example") // CreateReleaseInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesBetaApi.CreateReleaseForFlag(context.Background(), projectKey, flagKey).CreateReleaseInput(createReleaseInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesBetaApi.CreateReleaseForFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReleaseForFlag`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleasesBetaApi.CreateReleaseForFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**flagKey** | **string** | The flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseForFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createReleaseInput** | [**CreateReleaseInput**](CreateReleaseInput.md) |  | 

### Return type

[**Release**](Release.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReleaseByFlagKey

> DeleteReleaseByFlagKey(ctx, projectKey, flagKey).Execute()

Delete a release for flag



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
    resp, r, err := apiClient.ReleasesBetaApi.DeleteReleaseByFlagKey(context.Background(), projectKey, flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesBetaApi.DeleteReleaseByFlagKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteReleaseByFlagKeyRequest struct via the builder pattern


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

> Release PatchReleaseByFlagKey(ctx, projectKey, flagKey).PatchOperation(patchOperation).Execute()

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
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesBetaApi.PatchReleaseByFlagKey(context.Background(), projectKey, flagKey).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesBetaApi.PatchReleaseByFlagKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchReleaseByFlagKey`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleasesBetaApi.PatchReleaseByFlagKey`: %v\n", resp)
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

[**Release**](Release.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhaseStatus

> Release UpdatePhaseStatus(ctx, projectKey, flagKey, phaseId).UpdatePhaseStatusInput(updatePhaseStatusInput).Execute()

Update phase status for release



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
    phaseId := "phaseId_example" // string | The phase ID
    updatePhaseStatusInput := *openapiclient.NewUpdatePhaseStatusInput() // UpdatePhaseStatusInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasesBetaApi.UpdatePhaseStatus(context.Background(), projectKey, flagKey, phaseId).UpdatePhaseStatusInput(updatePhaseStatusInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasesBetaApi.UpdatePhaseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePhaseStatus`: Release
    fmt.Fprintf(os.Stdout, "Response from `ReleasesBetaApi.UpdatePhaseStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**flagKey** | **string** | The flag key | 
**phaseId** | **string** | The phase ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhaseStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updatePhaseStatusInput** | [**UpdatePhaseStatusInput**](UpdatePhaseStatusInput.md) |  | 

### Return type

[**Release**](Release.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

