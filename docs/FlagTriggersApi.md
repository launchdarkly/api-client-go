# \FlagTriggersApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTriggerWorkflow**](FlagTriggersApi.md#CreateTriggerWorkflow) | **Post** /api/v2/flags/{projectKey}/{featureFlagKey}/triggers/{environmentKey} | Create flag trigger
[**DeleteTriggerWorkflow**](FlagTriggersApi.md#DeleteTriggerWorkflow) | **Delete** /api/v2/flags/{projectKey}/{featureFlagKey}/triggers/{environmentKey}/{id} | Delete flag trigger
[**GetTriggerWorkflowById**](FlagTriggersApi.md#GetTriggerWorkflowById) | **Get** /api/v2/flags/{projectKey}/{featureFlagKey}/triggers/{environmentKey}/{id} | Get flag trigger by ID
[**GetTriggerWorkflows**](FlagTriggersApi.md#GetTriggerWorkflows) | **Get** /api/v2/flags/{projectKey}/{featureFlagKey}/triggers/{environmentKey} | List flag triggers
[**PatchTriggerWorkflow**](FlagTriggersApi.md#PatchTriggerWorkflow) | **Patch** /api/v2/flags/{projectKey}/{featureFlagKey}/triggers/{environmentKey}/{id} | Update flag trigger



## CreateTriggerWorkflow

> TriggerWorkflowRep CreateTriggerWorkflow(ctx, projectKey, environmentKey, featureFlagKey).TriggerPost(triggerPost).Execute()

Create flag trigger



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    triggerPost := *openapiclient.NewTriggerPost("generic-trigger") // TriggerPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlagTriggersApi.CreateTriggerWorkflow(context.Background(), projectKey, environmentKey, featureFlagKey).TriggerPost(triggerPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagTriggersApi.CreateTriggerWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTriggerWorkflow`: TriggerWorkflowRep
    fmt.Fprintf(os.Stdout, "Response from `FlagTriggersApi.CreateTriggerWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTriggerWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **triggerPost** | [**TriggerPost**](TriggerPost.md) |  | 

### Return type

[**TriggerWorkflowRep**](TriggerWorkflowRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTriggerWorkflow

> DeleteTriggerWorkflow(ctx, projectKey, environmentKey, featureFlagKey, id).Execute()

Delete flag trigger



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    id := "id_example" // string | The flag trigger ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlagTriggersApi.DeleteTriggerWorkflow(context.Background(), projectKey, environmentKey, featureFlagKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagTriggersApi.DeleteTriggerWorkflow``: %v\n", err)
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
**featureFlagKey** | **string** | The feature flag key | 
**id** | **string** | The flag trigger ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTriggerWorkflowRequest struct via the builder pattern


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


## GetTriggerWorkflowById

> TriggerWorkflowRep GetTriggerWorkflowById(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

Get flag trigger by ID



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
    id := "id_example" // string | The flag trigger ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlagTriggersApi.GetTriggerWorkflowById(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagTriggersApi.GetTriggerWorkflowById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTriggerWorkflowById`: TriggerWorkflowRep
    fmt.Fprintf(os.Stdout, "Response from `FlagTriggersApi.GetTriggerWorkflowById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The flag trigger ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerWorkflowByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**TriggerWorkflowRep**](TriggerWorkflowRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTriggerWorkflows

> TriggerWorkflowCollectionRep GetTriggerWorkflows(ctx, projectKey, environmentKey, featureFlagKey).Execute()

List flag triggers



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlagTriggersApi.GetTriggerWorkflows(context.Background(), projectKey, environmentKey, featureFlagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagTriggersApi.GetTriggerWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTriggerWorkflows`: TriggerWorkflowCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `FlagTriggersApi.GetTriggerWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTriggerWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TriggerWorkflowCollectionRep**](TriggerWorkflowCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTriggerWorkflow

> TriggerWorkflowRep PatchTriggerWorkflow(ctx, projectKey, environmentKey, featureFlagKey, id).FlagTriggerInput(flagTriggerInput).Execute()

Update flag trigger



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    id := "id_example" // string | The flag trigger ID
    flagTriggerInput := *openapiclient.NewFlagTriggerInput() // FlagTriggerInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FlagTriggersApi.PatchTriggerWorkflow(context.Background(), projectKey, environmentKey, featureFlagKey, id).FlagTriggerInput(flagTriggerInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagTriggersApi.PatchTriggerWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchTriggerWorkflow`: TriggerWorkflowRep
    fmt.Fprintf(os.Stdout, "Response from `FlagTriggersApi.PatchTriggerWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 
**id** | **string** | The flag trigger ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTriggerWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **flagTriggerInput** | [**FlagTriggerInput**](FlagTriggerInput.md) |  | 

### Return type

[**TriggerWorkflowRep**](TriggerWorkflowRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

