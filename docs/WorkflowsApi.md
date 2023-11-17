# \WorkflowsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkflow**](WorkflowsApi.md#DeleteWorkflow) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/workflows/{workflowId} | Delete workflow
[**GetCustomWorkflow**](WorkflowsApi.md#GetCustomWorkflow) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/workflows/{workflowId} | Get custom workflow
[**GetWorkflows**](WorkflowsApi.md#GetWorkflows) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/workflows | Get workflows
[**PostWorkflow**](WorkflowsApi.md#PostWorkflow) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/workflows | Create workflow



## DeleteWorkflow

> DeleteWorkflow(ctx, projectKey, featureFlagKey, environmentKey, workflowId).Execute()

Delete workflow



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
    workflowId := "workflowId_example" // string | The workflow id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.DeleteWorkflow(context.Background(), projectKey, featureFlagKey, environmentKey, workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.DeleteWorkflow``: %v\n", err)
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
**workflowId** | **string** | The workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowRequest struct via the builder pattern


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


## GetCustomWorkflow

> CustomWorkflowOutput GetCustomWorkflow(ctx, projectKey, featureFlagKey, environmentKey, workflowId).Execute()

Get custom workflow



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
    workflowId := "workflowId_example" // string | The workflow ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetCustomWorkflow(context.Background(), projectKey, featureFlagKey, environmentKey, workflowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetCustomWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomWorkflow`: CustomWorkflowOutput
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetCustomWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**workflowId** | **string** | The workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**CustomWorkflowOutput**](CustomWorkflowOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflows

> CustomWorkflowsListingOutput GetWorkflows(ctx, projectKey, featureFlagKey, environmentKey).Status(status).Sort(sort).Execute()

Get workflows



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
    status := "status_example" // string | Filter results by workflow status. Valid status filters are `active`, `completed`, and `failed`. (optional)
    sort := "sort_example" // string | A field to sort the items by. Prefix field by a dash ( - ) to sort in descending order. This endpoint supports sorting by `creationDate` or `stopDate`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflows(context.Background(), projectKey, featureFlagKey, environmentKey).Status(status).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflows`: CustomWorkflowsListingOutput
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflows`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **string** | Filter results by workflow status. Valid status filters are &#x60;active&#x60;, &#x60;completed&#x60;, and &#x60;failed&#x60;. | 
 **sort** | **string** | A field to sort the items by. Prefix field by a dash ( - ) to sort in descending order. This endpoint supports sorting by &#x60;creationDate&#x60; or &#x60;stopDate&#x60;. | 

### Return type

[**CustomWorkflowsListingOutput**](CustomWorkflowsListingOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWorkflow

> CustomWorkflowOutput PostWorkflow(ctx, projectKey, featureFlagKey, environmentKey).CustomWorkflowInput(customWorkflowInput).TemplateKey(templateKey).DryRun(dryRun).Execute()

Create workflow



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
    customWorkflowInput := *openapiclient.NewCustomWorkflowInput("Progressive rollout starting in two days") // CustomWorkflowInput | 
    templateKey := "templateKey_example" // string | The template key to apply as a starting point for the new workflow (optional)
    dryRun := true // bool | Whether to call the endpoint in dry-run mode (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.PostWorkflow(context.Background(), projectKey, featureFlagKey, environmentKey).CustomWorkflowInput(customWorkflowInput).TemplateKey(templateKey).DryRun(dryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.PostWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWorkflow`: CustomWorkflowOutput
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.PostWorkflow`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPostWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **customWorkflowInput** | [**CustomWorkflowInput**](CustomWorkflowInput.md) |  | 
 **templateKey** | **string** | The template key to apply as a starting point for the new workflow | 
 **dryRun** | **bool** | Whether to call the endpoint in dry-run mode | 

### Return type

[**CustomWorkflowOutput**](CustomWorkflowOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

