# \ReleasePipelinesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteReleasePipeline**](ReleasePipelinesBetaApi.md#DeleteReleasePipeline) | **Delete** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey} | Delete release pipeline
[**GetAllReleasePipelines**](ReleasePipelinesBetaApi.md#GetAllReleasePipelines) | **Get** /api/v2/projects/{projectKey}/release-pipelines | Get all release pipelines
[**GetReleasePipelineByKey**](ReleasePipelinesBetaApi.md#GetReleasePipelineByKey) | **Get** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey} | Get release pipeline by key
[**PatchReleasePipeline**](ReleasePipelinesBetaApi.md#PatchReleasePipeline) | **Patch** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey} | Update a release pipeline
[**PostReleasePipeline**](ReleasePipelinesBetaApi.md#PostReleasePipeline) | **Post** /api/v2/projects/{projectKey}/release-pipelines | Create a release pipeline



## DeleteReleasePipeline

> DeleteReleasePipeline(ctx, projectKey, pipelineKey).Execute()

Delete release pipeline



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
    pipelineKey := "pipelineKey_example" // string | The release pipeline key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasePipelinesBetaApi.DeleteReleasePipeline(context.Background(), projectKey, pipelineKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesBetaApi.DeleteReleasePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**pipelineKey** | **string** | The release pipeline key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleasePipelineRequest struct via the builder pattern


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


## GetAllReleasePipelines

> ReleasePipelineCollection GetAllReleasePipelines(ctx, projectKey).Execute()

Get all release pipelines



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasePipelinesBetaApi.GetAllReleasePipelines(context.Background(), projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesBetaApi.GetAllReleasePipelines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllReleasePipelines`: ReleasePipelineCollection
    fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesBetaApi.GetAllReleasePipelines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReleasePipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReleasePipelineCollection**](ReleasePipelineCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasePipelineByKey

> ReleasePipeline GetReleasePipelineByKey(ctx, projectKey, pipelineKey).Execute()

Get release pipeline by key



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
    pipelineKey := "pipelineKey_example" // string | The release pipeline key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasePipelinesBetaApi.GetReleasePipelineByKey(context.Background(), projectKey, pipelineKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesBetaApi.GetReleasePipelineByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleasePipelineByKey`: ReleasePipeline
    fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesBetaApi.GetReleasePipelineByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**pipelineKey** | **string** | The release pipeline key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasePipelineByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReleasePipeline**](ReleasePipeline.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchReleasePipeline

> ReleasePipeline PatchReleasePipeline(ctx, projectKey, pipelineKey).Execute()

Update a release pipeline



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
    pipelineKey := "pipelineKey_example" // string | The release pipeline key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasePipelinesBetaApi.PatchReleasePipeline(context.Background(), projectKey, pipelineKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesBetaApi.PatchReleasePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchReleasePipeline`: ReleasePipeline
    fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesBetaApi.PatchReleasePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**pipelineKey** | **string** | The release pipeline key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchReleasePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReleasePipeline**](ReleasePipeline.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReleasePipeline

> ReleasePipeline PostReleasePipeline(ctx, projectKey).CreateReleasePipelineInput(createReleasePipelineInput).Execute()

Create a release pipeline



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
    createReleasePipelineInput := *openapiclient.NewCreateReleasePipelineInput("standard-pipeline", "Standard Pipeline", []openapiclient.CreatePhaseInput{*openapiclient.NewCreatePhaseInput([]openapiclient.AudiencePost{*openapiclient.NewAudiencePost("EnvironmentKey_example", "Name_example")}, "Phase 1 - Testing")}) // CreateReleasePipelineInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasePipelinesBetaApi.PostReleasePipeline(context.Background(), projectKey).CreateReleasePipelineInput(createReleasePipelineInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesBetaApi.PostReleasePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostReleasePipeline`: ReleasePipeline
    fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesBetaApi.PostReleasePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReleasePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createReleasePipelineInput** | [**CreateReleasePipelineInput**](CreateReleasePipelineInput.md) |  | 

### Return type

[**ReleasePipeline**](ReleasePipeline.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

