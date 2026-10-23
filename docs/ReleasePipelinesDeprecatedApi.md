# \ReleasePipelinesDeprecatedApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReleaseForFlag**](ReleasePipelinesDeprecatedApi.md#CreateReleaseForFlag) | **Put** /api/v2/projects/{projectKey}/flags/{flagKey}/release | Create a new release for flag
[**DeleteReleaseByFlagKey**](ReleasePipelinesDeprecatedApi.md#DeleteReleaseByFlagKey) | **Delete** /api/v2/flags/{projectKey}/{flagKey}/release | Delete a release for flag
[**DeleteReleasePipeline**](ReleasePipelinesDeprecatedApi.md#DeleteReleasePipeline) | **Delete** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey} | Delete release pipeline
[**GetAllReleasePipelines**](ReleasePipelinesDeprecatedApi.md#GetAllReleasePipelines) | **Get** /api/v2/projects/{projectKey}/release-pipelines | Get all release pipelines
[**GetAllReleaseProgressionsForReleasePipeline**](ReleasePipelinesDeprecatedApi.md#GetAllReleaseProgressionsForReleasePipeline) | **Get** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey}/releases | Get release progressions for release pipeline
[**GetReleaseByFlagKey**](ReleasePipelinesDeprecatedApi.md#GetReleaseByFlagKey) | **Get** /api/v2/flags/{projectKey}/{flagKey}/release | Get release for flag
[**GetReleasePipelineByKey**](ReleasePipelinesDeprecatedApi.md#GetReleasePipelineByKey) | **Get** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey} | Get release pipeline by key
[**PatchReleaseByFlagKey**](ReleasePipelinesDeprecatedApi.md#PatchReleaseByFlagKey) | **Patch** /api/v2/flags/{projectKey}/{flagKey}/release | Patch release for flag
[**PostReleasePipeline**](ReleasePipelinesDeprecatedApi.md#PostReleasePipeline) | **Post** /api/v2/projects/{projectKey}/release-pipelines | Create a release pipeline
[**PutReleasePipeline**](ReleasePipelinesDeprecatedApi.md#PutReleasePipeline) | **Put** /api/v2/projects/{projectKey}/release-pipelines/{pipelineKey} | Update a release pipeline
[**UpdatePhaseStatus**](ReleasePipelinesDeprecatedApi.md#UpdatePhaseStatus) | **Put** /api/v2/projects/{projectKey}/flags/{flagKey}/release/phases/{phaseId} | Update phase status for release



## CreateReleaseForFlag

> PipelineRelease CreateReleaseForFlag(ctx, projectKey, flagKey).CreatePipelineReleaseInput(createPipelineReleaseInput).Execute()

Create a new release for flag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	flagKey := "flagKey_example" // string | The flag key
	createPipelineReleaseInput := *openapiclient.NewCreatePipelineReleaseInput("ReleasePipelineKey_example") // CreatePipelineReleaseInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.CreateReleaseForFlag(context.Background(), projectKey, flagKey).CreatePipelineReleaseInput(createPipelineReleaseInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.CreateReleaseForFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReleaseForFlag`: PipelineRelease
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.CreateReleaseForFlag`: %v\n", resp)
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


 **createPipelineReleaseInput** | [**CreatePipelineReleaseInput**](CreatePipelineReleaseInput.md) |  | 

### Return type

[**PipelineRelease**](PipelineRelease.md)

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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	flagKey := "flagKey_example" // string | The flag key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReleasePipelinesDeprecatedApi.DeleteReleaseByFlagKey(context.Background(), projectKey, flagKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.DeleteReleaseByFlagKey``: %v\n", err)
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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	pipelineKey := "pipelineKey_example" // string | The release pipeline key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReleasePipelinesDeprecatedApi.DeleteReleasePipeline(context.Background(), projectKey, pipelineKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.DeleteReleasePipeline``: %v\n", err)
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

> ReleasePipelineCollection GetAllReleasePipelines(ctx, projectKey).Filter(filter).Limit(limit).Offset(offset).Execute()

Get all release pipelines



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form field:value. Read the endpoint description for a full list of available filter fields. (optional)
	limit := int64(789) // int64 | The maximum number of items to return. Defaults to 20. (optional)
	offset := int64(789) // int64 | Where to start in the list. Defaults to 0. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.GetAllReleasePipelines(context.Background(), projectKey).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.GetAllReleasePipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllReleasePipelines`: ReleasePipelineCollection
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.GetAllReleasePipelines`: %v\n", resp)
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

 **filter** | **string** | A comma-separated list of filters. Each filter is of the form field:value. Read the endpoint description for a full list of available filter fields. | 
 **limit** | **int64** | The maximum number of items to return. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. Defaults to 0. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

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

> PipelineReleaseProgressionCollection GetAllReleaseProgressionsForReleasePipeline(ctx, projectKey, pipelineKey).Filter(filter).Limit(limit).Offset(offset).Execute()

Get release progressions for release pipeline



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	pipelineKey := "pipelineKey_example" // string | The pipeline key
	filter := "filter_example" // string | Accepts filter by `status` and `activePhaseId`. `status` can take a value of `completed` or `active`. `activePhaseId` takes a UUID and will filter results down to releases active on the specified phase. Providing `status equals completed` along with an `activePhaseId` filter will return an error as they are disjoint sets of data. The combination of `status equals active` and `activePhaseId` will return the same results as `activePhaseId` alone. (optional)
	limit := int64(789) // int64 | The maximum number of items to return. Defaults to 20. (optional)
	offset := int64(789) // int64 | Where to start in the list. Defaults to 0. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.GetAllReleaseProgressionsForReleasePipeline(context.Background(), projectKey, pipelineKey).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.GetAllReleaseProgressionsForReleasePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllReleaseProgressionsForReleasePipeline`: PipelineReleaseProgressionCollection
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.GetAllReleaseProgressionsForReleasePipeline`: %v\n", resp)
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

[**PipelineReleaseProgressionCollection**](PipelineReleaseProgressionCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseByFlagKey

> PipelineRelease GetReleaseByFlagKey(ctx, projectKey, flagKey).Execute()

Get release for flag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	flagKey := "flagKey_example" // string | The flag key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.GetReleaseByFlagKey(context.Background(), projectKey, flagKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.GetReleaseByFlagKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseByFlagKey`: PipelineRelease
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.GetReleaseByFlagKey`: %v\n", resp)
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

[**PipelineRelease**](PipelineRelease.md)

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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	pipelineKey := "pipelineKey_example" // string | The release pipeline key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.GetReleasePipelineByKey(context.Background(), projectKey, pipelineKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.GetReleasePipelineByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleasePipelineByKey`: ReleasePipeline
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.GetReleasePipelineByKey`: %v\n", resp)
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


## PatchReleaseByFlagKey

> PipelineRelease PatchReleaseByFlagKey(ctx, projectKey, flagKey).PatchOperation(patchOperation).Execute()

Patch release for flag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	flagKey := "flagKey_example" // string | The flag key
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.PatchReleaseByFlagKey(context.Background(), projectKey, flagKey).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.PatchReleaseByFlagKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchReleaseByFlagKey`: PipelineRelease
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.PatchReleaseByFlagKey`: %v\n", resp)
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

[**PipelineRelease**](PipelineRelease.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	createReleasePipelineInput := *openapiclient.NewCreateReleasePipelineInput("standard-pipeline", "Standard Pipeline", []openapiclient.CreatePhaseInput{*openapiclient.NewCreatePhaseInput([]openapiclient.AudiencePost{*openapiclient.NewAudiencePost("EnvironmentKey_example", "Name_example")}, "Phase 1 - Testing")}) // CreateReleasePipelineInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.PostReleasePipeline(context.Background(), projectKey).CreateReleasePipelineInput(createReleasePipelineInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.PostReleasePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReleasePipeline`: ReleasePipeline
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.PostReleasePipeline`: %v\n", resp)
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


## PutReleasePipeline

> ReleasePipeline PutReleasePipeline(ctx, projectKey, pipelineKey).UpdateReleasePipelineInput(updateReleasePipelineInput).Execute()

Update a release pipeline



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	pipelineKey := "pipelineKey_example" // string | The release pipeline key
	updateReleasePipelineInput := *openapiclient.NewUpdateReleasePipelineInput("Standard Pipeline", []openapiclient.CreatePhaseInput{*openapiclient.NewCreatePhaseInput([]openapiclient.AudiencePost{*openapiclient.NewAudiencePost("EnvironmentKey_example", "Name_example")}, "Phase 1 - Testing")}) // UpdateReleasePipelineInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.PutReleasePipeline(context.Background(), projectKey, pipelineKey).UpdateReleasePipelineInput(updateReleasePipelineInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.PutReleasePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutReleasePipeline`: ReleasePipeline
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.PutReleasePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**pipelineKey** | **string** | The release pipeline key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutReleasePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateReleasePipelineInput** | [**UpdateReleasePipelineInput**](UpdateReleasePipelineInput.md) |  | 

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


## UpdatePhaseStatus

> PipelineRelease UpdatePhaseStatus(ctx, projectKey, flagKey, phaseId).UpdatePhaseStatusInput(updatePhaseStatusInput).Execute()

Update phase status for release



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	flagKey := "flagKey_example" // string | The flag key
	phaseId := "phaseId_example" // string | The phase ID
	updatePhaseStatusInput := *openapiclient.NewUpdatePhaseStatusInput() // UpdatePhaseStatusInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePipelinesDeprecatedApi.UpdatePhaseStatus(context.Background(), projectKey, flagKey, phaseId).UpdatePhaseStatusInput(updatePhaseStatusInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePipelinesDeprecatedApi.UpdatePhaseStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePhaseStatus`: PipelineRelease
	fmt.Fprintf(os.Stdout, "Response from `ReleasePipelinesDeprecatedApi.UpdatePhaseStatus`: %v\n", resp)
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

[**PipelineRelease**](PipelineRelease.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

