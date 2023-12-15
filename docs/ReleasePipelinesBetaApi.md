# \ReleasePipelinesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteReleasePipeline**](ReleasePipelinesBetaApi.md#DeleteReleasePipeline) | **Delete** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey} | Delete release pipeline
[**GetAllReleasePipelines**](ReleasePipelinesBetaApi.md#GetAllReleasePipelines) | **Get** /api/v2/projects/{projectKey}/release-pipelines | Get all release pipelines
[**GetAllReleaseProgressionsForReleasePipeline**](ReleasePipelinesBetaApi.md#GetAllReleaseProgressionsForReleasePipeline) | **Get** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey}/releases | Get release progressions for release pipeline
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


## GetAllReleaseProgressionsForReleasePipeline

> ReleaseProgressionCollection GetAllReleaseProgressionsForReleasePipeline(ctx, projectKey, pipelineKey).Filter(filter).Limit(limit).Offset(offset).Execute()

Get release progressions for release pipeline



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
    pipelineKey := "pipelineKey_example" // string | The pipeline key
    filter := "filter_example" // string | Accepts filter by `status` and `activePhaseId`. `status` can take a value of `completed` or `active`. `activePhaseId` takes a UUID and will filter results down to releases active on the specified phase. Providing `status equals completed` along with an `activePhaseId` filter will return an error as they are disjoint sets of data. The combination of `status equals active` and `activePhaseId` will return the same results as `activePhaseId` alone. (optional)
    limit := int64(789) // int64 | The maximum number of items to return. Defaults to 20. (optional)
    offset := int64(789) // int64 | Where to start in the list. Defaults to 0. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReleasePipelinesBetaApi.GetAllReleaseProgressionsForReleasePipeline(context.Background(), projectKey, pipelineKey).Filter(filter).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesBetaApi.GetAllReleaseProgressionsForReleasePipeline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllReleaseProgressionsForReleasePipeline`: ReleaseProgressionCollection
    fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesBetaApi.GetAllReleaseProgressionsForReleasePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**pipelineKey** | **string** | The pipeline key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReleaseProgressionsForReleasePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | Accepts filter by &#x60;status&#x60; and &#x60;activePhaseId&#x60;. &#x60;status&#x60; can take a value of &#x60;completed&#x60; or &#x60;active&#x60;. &#x60;activePhaseId&#x60; takes a UUID and will filter results down to releases active on the specified phase. Providing &#x60;status equals completed&#x60; along with an &#x60;activePhaseId&#x60; filter will return an error as they are disjoint sets of data. The combination of &#x60;status equals active&#x60; and &#x60;activePhaseId&#x60; will return the same results as &#x60;activePhaseId&#x60; alone. | 
 **limit** | **int64** | The maximum number of items to return. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. Defaults to 0. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**ReleaseProgressionCollection**](ReleaseProgressionCollection.md)

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

