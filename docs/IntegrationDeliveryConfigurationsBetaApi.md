# \IntegrationDeliveryConfigurationsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIntegrationDeliveryConfiguration**](IntegrationDeliveryConfigurationsBetaApi.md#CreateIntegrationDeliveryConfiguration) | **Post** /api/v2/integration-capabilities/featureStore/{projectKey}/{environmentKey}/{integrationKey} | Create delivery configuration
[**DeleteIntegrationDeliveryConfiguration**](IntegrationDeliveryConfigurationsBetaApi.md#DeleteIntegrationDeliveryConfiguration) | **Delete** /api/v2/integration-capabilities/featureStore/{projectKey}/{environmentKey}/{integrationKey}/{id} | Delete delivery configuration
[**GetIntegrationDeliveryConfigurationByEnvironment**](IntegrationDeliveryConfigurationsBetaApi.md#GetIntegrationDeliveryConfigurationByEnvironment) | **Get** /api/v2/integration-capabilities/featureStore/{projectKey}/{environmentKey} | Get delivery configurations by environment
[**GetIntegrationDeliveryConfigurationById**](IntegrationDeliveryConfigurationsBetaApi.md#GetIntegrationDeliveryConfigurationById) | **Get** /api/v2/integration-capabilities/featureStore/{projectKey}/{environmentKey}/{integrationKey}/{id} | Get delivery configuration by ID
[**GetIntegrationDeliveryConfigurations**](IntegrationDeliveryConfigurationsBetaApi.md#GetIntegrationDeliveryConfigurations) | **Get** /api/v2/integration-capabilities/featureStore | List all delivery configurations
[**PatchIntegrationDeliveryConfiguration**](IntegrationDeliveryConfigurationsBetaApi.md#PatchIntegrationDeliveryConfiguration) | **Patch** /api/v2/integration-capabilities/featureStore/{projectKey}/{environmentKey}/{integrationKey}/{id} | Update delivery configuration
[**ValidateIntegrationDeliveryConfiguration**](IntegrationDeliveryConfigurationsBetaApi.md#ValidateIntegrationDeliveryConfiguration) | **Post** /api/v2/integration-capabilities/featureStore/{projectKey}/{environmentKey}/{integrationKey}/{id}/validate | Validate delivery configuration



## CreateIntegrationDeliveryConfiguration

> IntegrationDeliveryConfiguration CreateIntegrationDeliveryConfiguration(ctx, projectKey, environmentKey, integrationKey).IntegrationDeliveryConfigurationPost(integrationDeliveryConfigurationPost).Execute()

Create delivery configuration



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
    integrationKey := "integrationKey_example" // string | The integration key
    integrationDeliveryConfigurationPost := *openapiclient.NewIntegrationDeliveryConfigurationPost(map[string]interface{}{"key": interface{}(123)}) // IntegrationDeliveryConfigurationPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationDeliveryConfigurationsBetaApi.CreateIntegrationDeliveryConfiguration(context.Background(), projectKey, environmentKey, integrationKey).IntegrationDeliveryConfigurationPost(integrationDeliveryConfigurationPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationDeliveryConfigurationsBetaApi.CreateIntegrationDeliveryConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIntegrationDeliveryConfiguration`: IntegrationDeliveryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationDeliveryConfigurationsBetaApi.CreateIntegrationDeliveryConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**integrationKey** | **string** | The integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIntegrationDeliveryConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **integrationDeliveryConfigurationPost** | [**IntegrationDeliveryConfigurationPost**](IntegrationDeliveryConfigurationPost.md) |  | 

### Return type

[**IntegrationDeliveryConfiguration**](IntegrationDeliveryConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationDeliveryConfiguration

> DeleteIntegrationDeliveryConfiguration(ctx, projectKey, environmentKey, integrationKey, id).Execute()

Delete delivery configuration



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
    integrationKey := "integrationKey_example" // string | The integration key
    id := "id_example" // string | The configuration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationDeliveryConfigurationsBetaApi.DeleteIntegrationDeliveryConfiguration(context.Background(), projectKey, environmentKey, integrationKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationDeliveryConfigurationsBetaApi.DeleteIntegrationDeliveryConfiguration``: %v\n", err)
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
**integrationKey** | **string** | The integration key | 
**id** | **string** | The configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationDeliveryConfigurationRequest struct via the builder pattern


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


## GetIntegrationDeliveryConfigurationByEnvironment

> IntegrationDeliveryConfigurationCollection GetIntegrationDeliveryConfigurationByEnvironment(ctx, projectKey, environmentKey).Execute()

Get delivery configurations by environment



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurationByEnvironment(context.Background(), projectKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurationByEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationDeliveryConfigurationByEnvironment`: IntegrationDeliveryConfigurationCollection
    fmt.Fprintf(os.Stdout, "Response from `IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurationByEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDeliveryConfigurationByEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IntegrationDeliveryConfigurationCollection**](IntegrationDeliveryConfigurationCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationDeliveryConfigurationById

> IntegrationDeliveryConfiguration GetIntegrationDeliveryConfigurationById(ctx, projectKey, environmentKey, integrationKey, id).Execute()

Get delivery configuration by ID



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
    integrationKey := "integrationKey_example" // string | The integration key
    id := "id_example" // string | The configuration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurationById(context.Background(), projectKey, environmentKey, integrationKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationDeliveryConfigurationById`: IntegrationDeliveryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**integrationKey** | **string** | The integration key | 
**id** | **string** | The configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDeliveryConfigurationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**IntegrationDeliveryConfiguration**](IntegrationDeliveryConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationDeliveryConfigurations

> IntegrationDeliveryConfigurationCollection GetIntegrationDeliveryConfigurations(ctx).Execute()

List all delivery configurations



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
    resp, r, err := apiClient.IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntegrationDeliveryConfigurations`: IntegrationDeliveryConfigurationCollection
    fmt.Fprintf(os.Stdout, "Response from `IntegrationDeliveryConfigurationsBetaApi.GetIntegrationDeliveryConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationDeliveryConfigurationsRequest struct via the builder pattern


### Return type

[**IntegrationDeliveryConfigurationCollection**](IntegrationDeliveryConfigurationCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIntegrationDeliveryConfiguration

> IntegrationDeliveryConfiguration PatchIntegrationDeliveryConfiguration(ctx, projectKey, environmentKey, integrationKey, id).PatchOperation(patchOperation).Execute()

Update delivery configuration



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
    integrationKey := "integrationKey_example" // string | The integration key
    id := "id_example" // string | The configuration ID
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationDeliveryConfigurationsBetaApi.PatchIntegrationDeliveryConfiguration(context.Background(), projectKey, environmentKey, integrationKey, id).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationDeliveryConfigurationsBetaApi.PatchIntegrationDeliveryConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchIntegrationDeliveryConfiguration`: IntegrationDeliveryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationDeliveryConfigurationsBetaApi.PatchIntegrationDeliveryConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**integrationKey** | **string** | The integration key | 
**id** | **string** | The configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIntegrationDeliveryConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**IntegrationDeliveryConfiguration**](IntegrationDeliveryConfiguration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateIntegrationDeliveryConfiguration

> IntegrationDeliveryConfigurationResponse ValidateIntegrationDeliveryConfiguration(ctx, projectKey, environmentKey, integrationKey, id).Execute()

Validate delivery configuration



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
    integrationKey := "integrationKey_example" // string | The integration key
    id := "id_example" // string | The configuration ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationDeliveryConfigurationsBetaApi.ValidateIntegrationDeliveryConfiguration(context.Background(), projectKey, environmentKey, integrationKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationDeliveryConfigurationsBetaApi.ValidateIntegrationDeliveryConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateIntegrationDeliveryConfiguration`: IntegrationDeliveryConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `IntegrationDeliveryConfigurationsBetaApi.ValidateIntegrationDeliveryConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**integrationKey** | **string** | The integration key | 
**id** | **string** | The configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateIntegrationDeliveryConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**IntegrationDeliveryConfigurationResponse**](IntegrationDeliveryConfigurationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

