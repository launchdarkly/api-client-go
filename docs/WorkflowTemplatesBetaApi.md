# \WorkflowTemplatesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowTemplate**](WorkflowTemplatesBetaApi.md#CreateWorkflowTemplate) | **Post** /api/v2/templates | Create workflow template
[**DeleteWorkflowTemplate**](WorkflowTemplatesBetaApi.md#DeleteWorkflowTemplate) | **Delete** /api/v2/templates/{templateKey} | Delete workflow template
[**GetWorkflowTemplates**](WorkflowTemplatesBetaApi.md#GetWorkflowTemplates) | **Get** /api/v2/templates | Get workflow templates



## CreateWorkflowTemplate

> WorkflowTemplateOutput CreateWorkflowTemplate(ctx).CreateWorkflowTemplateInput(createWorkflowTemplateInput).Execute()

Create workflow template



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
    createWorkflowTemplateInput := *openapiclient.NewCreateWorkflowTemplateInput("Key_example") // CreateWorkflowTemplateInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTemplatesBetaApi.CreateWorkflowTemplate(context.Background()).CreateWorkflowTemplateInput(createWorkflowTemplateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplatesBetaApi.CreateWorkflowTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowTemplate`: WorkflowTemplateOutput
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplatesBetaApi.CreateWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowTemplateInput** | [**CreateWorkflowTemplateInput**](CreateWorkflowTemplateInput.md) |  | 

### Return type

[**WorkflowTemplateOutput**](WorkflowTemplateOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowTemplate

> DeleteWorkflowTemplate(ctx, templateKey).Execute()

Delete workflow template



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
    templateKey := "templateKey_example" // string | The template key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTemplatesBetaApi.DeleteWorkflowTemplate(context.Background(), templateKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplatesBetaApi.DeleteWorkflowTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateKey** | **string** | The template key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowTemplateRequest struct via the builder pattern


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


## GetWorkflowTemplates

> WorkflowTemplatesListingOutputRep GetWorkflowTemplates(ctx).Search(search).Execute()

Get workflow templates



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
    search := "search_example" // string | The substring in either the name or description of a template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowTemplatesBetaApi.GetWorkflowTemplates(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplatesBetaApi.GetWorkflowTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowTemplates`: WorkflowTemplatesListingOutputRep
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplatesBetaApi.GetWorkflowTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | The substring in either the name or description of a template | 

### Return type

[**WorkflowTemplatesListingOutputRep**](WorkflowTemplatesListingOutputRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

