# \FlagImportConfigurationsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlagImportConfiguration**](FlagImportConfigurationsBetaApi.md#CreateFlagImportConfiguration) | **Post** /api/v2/integration-capabilities/flag-import/{projectKey}/{integrationKey} | Create a flag import configuration
[**DeleteFlagImportConfiguration**](FlagImportConfigurationsBetaApi.md#DeleteFlagImportConfiguration) | **Delete** /api/v2/integration-capabilities/flag-import/{projectKey}/{integrationKey}/{integrationId} | Delete a flag import configuration
[**GetFlagImportConfiguration**](FlagImportConfigurationsBetaApi.md#GetFlagImportConfiguration) | **Get** /api/v2/integration-capabilities/flag-import/{projectKey}/{integrationKey}/{integrationId} | Get a single flag import configuration
[**GetFlagImportConfigurations**](FlagImportConfigurationsBetaApi.md#GetFlagImportConfigurations) | **Get** /api/v2/integration-capabilities/flag-import | List all flag import configurations
[**PatchFlagImportConfiguration**](FlagImportConfigurationsBetaApi.md#PatchFlagImportConfiguration) | **Patch** /api/v2/integration-capabilities/flag-import/{projectKey}/{integrationKey}/{integrationId} | Update a flag import configuration
[**TriggerFlagImportJob**](FlagImportConfigurationsBetaApi.md#TriggerFlagImportJob) | **Post** /api/v2/integration-capabilities/flag-import/{projectKey}/{integrationKey}/{integrationId}/trigger | Trigger a single flag import run



## CreateFlagImportConfiguration

> FlagImportIntegration CreateFlagImportConfiguration(ctx, projectKey, integrationKey).FlagImportConfigurationPost(flagImportConfigurationPost).Execute()

Create a flag import configuration



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
    integrationKey := "integrationKey_example" // string | The integration key
    flagImportConfigurationPost := *openapiclient.NewFlagImportConfigurationPost(map[string]interface{}{"key": interface{}(123)}) // FlagImportConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagImportConfigurationsBetaApi.CreateFlagImportConfiguration(context.Background(), projectKey, integrationKey).FlagImportConfigurationPost(flagImportConfigurationPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagImportConfigurationsBetaApi.CreateFlagImportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlagImportConfiguration`: FlagImportIntegration
    fmt.Fprintf(os.Stdout, "Response from `FlagImportConfigurationsBetaApi.CreateFlagImportConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**integrationKey** | **string** | The integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlagImportConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flagImportConfigurationPost** | [**FlagImportConfigurationPost**](FlagImportConfigurationPost.md) |  | 

### Return type

[**FlagImportIntegration**](FlagImportIntegration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlagImportConfiguration

> DeleteFlagImportConfiguration(ctx, projectKey, integrationKey, integrationId).Execute()

Delete a flag import configuration



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
    integrationKey := "integrationKey_example" // string | The integration key
    integrationId := "integrationId_example" // string | The integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagImportConfigurationsBetaApi.DeleteFlagImportConfiguration(context.Background(), projectKey, integrationKey, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagImportConfigurationsBetaApi.DeleteFlagImportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**integrationKey** | **string** | The integration key | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagImportConfigurationRequest struct via the builder pattern


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


## GetFlagImportConfiguration

> FlagImportIntegration GetFlagImportConfiguration(ctx, projectKey, integrationKey, integrationId).Execute()

Get a single flag import configuration



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
    integrationKey := "integrationKey_example" // string | The integration key, for example, `split`
    integrationId := "integrationId_example" // string | The integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagImportConfigurationsBetaApi.GetFlagImportConfiguration(context.Background(), projectKey, integrationKey, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagImportConfigurationsBetaApi.GetFlagImportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagImportConfiguration`: FlagImportIntegration
    fmt.Fprintf(os.Stdout, "Response from `FlagImportConfigurationsBetaApi.GetFlagImportConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**integrationKey** | **string** | The integration key, for example, &#x60;split&#x60; | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagImportConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlagImportIntegration**](FlagImportIntegration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagImportConfigurations

> FlagImportIntegrationCollection GetFlagImportConfigurations(ctx).Execute()

List all flag import configurations



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagImportConfigurationsBetaApi.GetFlagImportConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagImportConfigurationsBetaApi.GetFlagImportConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagImportConfigurations`: FlagImportIntegrationCollection
    fmt.Fprintf(os.Stdout, "Response from `FlagImportConfigurationsBetaApi.GetFlagImportConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagImportConfigurationsRequest struct via the builder pattern


### Return type

[**FlagImportIntegrationCollection**](FlagImportIntegrationCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFlagImportConfiguration

> FlagImportIntegration PatchFlagImportConfiguration(ctx, projectKey, integrationKey, integrationId).PatchOperation(patchOperation).Execute()

Update a flag import configuration



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
    integrationKey := "integrationKey_example" // string | The integration key
    integrationId := "integrationId_example" // string | The integration ID
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagImportConfigurationsBetaApi.PatchFlagImportConfiguration(context.Background(), projectKey, integrationKey, integrationId).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagImportConfigurationsBetaApi.PatchFlagImportConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFlagImportConfiguration`: FlagImportIntegration
    fmt.Fprintf(os.Stdout, "Response from `FlagImportConfigurationsBetaApi.PatchFlagImportConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**integrationKey** | **string** | The integration key | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlagImportConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**FlagImportIntegration**](FlagImportIntegration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerFlagImportJob

> Object TriggerFlagImportJob(ctx, projectKey, integrationKey, integrationId).Execute()

Trigger a single flag import run



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
    integrationKey := "integrationKey_example" // string | The integration key
    integrationId := "integrationId_example" // string | The integration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagImportConfigurationsBetaApi.TriggerFlagImportJob(context.Background(), projectKey, integrationKey, integrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagImportConfigurationsBetaApi.TriggerFlagImportJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TriggerFlagImportJob`: Object
    fmt.Fprintf(os.Stdout, "Response from `FlagImportConfigurationsBetaApi.TriggerFlagImportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**integrationKey** | **string** | The integration key | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerFlagImportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Object**](Object.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

