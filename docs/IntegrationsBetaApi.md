# \IntegrationsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIntegrationConfiguration**](IntegrationsBetaApi.md#CreateIntegrationConfiguration) | **Post** /api/v2/integration-configurations/keys/{integrationKey} | Create integration configuration
[**DeleteIntegrationConfiguration**](IntegrationsBetaApi.md#DeleteIntegrationConfiguration) | **Delete** /api/v2/integration-configurations/{integrationConfigurationId} | Delete integration configuration
[**GetAllIntegrationConfigurations**](IntegrationsBetaApi.md#GetAllIntegrationConfigurations) | **Get** /api/v2/integration-configurations/keys/{integrationKey} | Get all configurations for the integration
[**GetIntegrationConfiguration**](IntegrationsBetaApi.md#GetIntegrationConfiguration) | **Get** /api/v2/integration-configurations/{integrationConfigurationId} | Get an integration configuration
[**UpdateIntegrationConfiguration**](IntegrationsBetaApi.md#UpdateIntegrationConfiguration) | **Patch** /api/v2/integration-configurations/{integrationConfigurationId} | Update integration configuration



## CreateIntegrationConfiguration

> IntegrationConfigurationsRep CreateIntegrationConfiguration(ctx, integrationKey).IntegrationConfigurationPost(integrationConfigurationPost).Execute()

Create integration configuration



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
	integrationKey := "integrationKey_example" // string | The integration key
	integrationConfigurationPost := *openapiclient.NewIntegrationConfigurationPost("Example integration configuration", map[string]interface{}{"key": interface{}(123)}) // IntegrationConfigurationPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsBetaApi.CreateIntegrationConfiguration(context.Background(), integrationKey).IntegrationConfigurationPost(integrationConfigurationPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsBetaApi.CreateIntegrationConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIntegrationConfiguration`: IntegrationConfigurationsRep
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsBetaApi.CreateIntegrationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIntegrationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationConfigurationPost** | [**IntegrationConfigurationPost**](IntegrationConfigurationPost.md) |  | 

### Return type

[**IntegrationConfigurationsRep**](IntegrationConfigurationsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationConfiguration

> DeleteIntegrationConfiguration(ctx, integrationConfigurationId).Execute()

Delete integration configuration



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
	integrationConfigurationId := "integrationConfigurationId_example" // string | The ID of the integration configuration to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsBetaApi.DeleteIntegrationConfiguration(context.Background(), integrationConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsBetaApi.DeleteIntegrationConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationConfigurationId** | **string** | The ID of the integration configuration to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationConfigurationRequest struct via the builder pattern


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


## GetAllIntegrationConfigurations

> IntegrationConfigurationCollectionRep GetAllIntegrationConfigurations(ctx, integrationKey).Execute()

Get all configurations for the integration



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
	integrationKey := "integrationKey_example" // string | Integration key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsBetaApi.GetAllIntegrationConfigurations(context.Background(), integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsBetaApi.GetAllIntegrationConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllIntegrationConfigurations`: IntegrationConfigurationCollectionRep
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsBetaApi.GetAllIntegrationConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | Integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIntegrationConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationConfigurationCollectionRep**](IntegrationConfigurationCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationConfiguration

> IntegrationConfigurationsRep GetIntegrationConfiguration(ctx, integrationConfigurationId).Execute()

Get an integration configuration



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
	integrationConfigurationId := "integrationConfigurationId_example" // string | Integration configuration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsBetaApi.GetIntegrationConfiguration(context.Background(), integrationConfigurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsBetaApi.GetIntegrationConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationConfiguration`: IntegrationConfigurationsRep
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsBetaApi.GetIntegrationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationConfigurationId** | **string** | Integration configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationConfigurationsRep**](IntegrationConfigurationsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIntegrationConfiguration

> IntegrationConfigurationsRep UpdateIntegrationConfiguration(ctx, integrationConfigurationId).PatchOperation(patchOperation).Execute()

Update integration configuration



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
	integrationConfigurationId := "integrationConfigurationId_example" // string | The ID of the integration configuration
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsBetaApi.UpdateIntegrationConfiguration(context.Background(), integrationConfigurationId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsBetaApi.UpdateIntegrationConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIntegrationConfiguration`: IntegrationConfigurationsRep
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsBetaApi.UpdateIntegrationConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationConfigurationId** | **string** | The ID of the integration configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIntegrationConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**IntegrationConfigurationsRep**](IntegrationConfigurationsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

