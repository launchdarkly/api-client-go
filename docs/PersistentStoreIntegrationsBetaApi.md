# \PersistentStoreIntegrationsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBigSegmentStoreIntegration**](PersistentStoreIntegrationsBetaApi.md#CreateBigSegmentStoreIntegration) | **Post** /api/v2/integration-capabilities/big-segment-store/{projectKey}/{environmentKey}/{integrationKey} | Create big segment store integration
[**DeleteBigSegmentStoreIntegration**](PersistentStoreIntegrationsBetaApi.md#DeleteBigSegmentStoreIntegration) | **Delete** /api/v2/integration-capabilities/big-segment-store/{projectKey}/{environmentKey}/{integrationKey}/{integrationId} | Delete big segment store integration
[**GetBigSegmentStoreIntegration**](PersistentStoreIntegrationsBetaApi.md#GetBigSegmentStoreIntegration) | **Get** /api/v2/integration-capabilities/big-segment-store/{projectKey}/{environmentKey}/{integrationKey}/{integrationId} | Get big segment store integration by ID
[**GetBigSegmentStoreIntegrations**](PersistentStoreIntegrationsBetaApi.md#GetBigSegmentStoreIntegrations) | **Get** /api/v2/integration-capabilities/big-segment-store | List all big segment store integrations
[**PatchBigSegmentStoreIntegration**](PersistentStoreIntegrationsBetaApi.md#PatchBigSegmentStoreIntegration) | **Patch** /api/v2/integration-capabilities/big-segment-store/{projectKey}/{environmentKey}/{integrationKey}/{integrationId} | Update big segment store integration



## CreateBigSegmentStoreIntegration

> BigSegmentStoreIntegration CreateBigSegmentStoreIntegration(ctx, projectKey, environmentKey, integrationKey).IntegrationDeliveryConfigurationPost(integrationDeliveryConfigurationPost).Execute()

Create big segment store integration



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
	environmentKey := "environmentKey_example" // string | The environment key
	integrationKey := "integrationKey_example" // string | The integration key, either `redis` or `dynamodb`
	integrationDeliveryConfigurationPost := *openapiclient.NewIntegrationDeliveryConfigurationPost(map[string]interface{}{"key": interface{}(123)}) // IntegrationDeliveryConfigurationPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersistentStoreIntegrationsBetaApi.CreateBigSegmentStoreIntegration(context.Background(), projectKey, environmentKey, integrationKey).IntegrationDeliveryConfigurationPost(integrationDeliveryConfigurationPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersistentStoreIntegrationsBetaApi.CreateBigSegmentStoreIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBigSegmentStoreIntegration`: BigSegmentStoreIntegration
	fmt.Fprintf(os.Stdout, "Response from `PersistentStoreIntegrationsBetaApi.CreateBigSegmentStoreIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**integrationKey** | **string** | The integration key, either &#x60;redis&#x60; or &#x60;dynamodb&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBigSegmentStoreIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **integrationDeliveryConfigurationPost** | [**IntegrationDeliveryConfigurationPost**](IntegrationDeliveryConfigurationPost.md) |  | 

### Return type

[**BigSegmentStoreIntegration**](BigSegmentStoreIntegration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBigSegmentStoreIntegration

> DeleteBigSegmentStoreIntegration(ctx, projectKey, environmentKey, integrationKey, integrationId).Execute()

Delete big segment store integration



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
	environmentKey := "environmentKey_example" // string | The environment key
	integrationKey := "integrationKey_example" // string | The integration key, either `redis` or `dynamodb`
	integrationId := "integrationId_example" // string | The integration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersistentStoreIntegrationsBetaApi.DeleteBigSegmentStoreIntegration(context.Background(), projectKey, environmentKey, integrationKey, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersistentStoreIntegrationsBetaApi.DeleteBigSegmentStoreIntegration``: %v\n", err)
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
**integrationKey** | **string** | The integration key, either &#x60;redis&#x60; or &#x60;dynamodb&#x60; | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBigSegmentStoreIntegrationRequest struct via the builder pattern


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


## GetBigSegmentStoreIntegration

> BigSegmentStoreIntegration GetBigSegmentStoreIntegration(ctx, projectKey, environmentKey, integrationKey, integrationId).Execute()

Get big segment store integration by ID



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
	environmentKey := "environmentKey_example" // string | The environment key
	integrationKey := "integrationKey_example" // string | The integration key, either `redis` or `dynamodb`
	integrationId := "integrationId_example" // string | The integration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersistentStoreIntegrationsBetaApi.GetBigSegmentStoreIntegration(context.Background(), projectKey, environmentKey, integrationKey, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersistentStoreIntegrationsBetaApi.GetBigSegmentStoreIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBigSegmentStoreIntegration`: BigSegmentStoreIntegration
	fmt.Fprintf(os.Stdout, "Response from `PersistentStoreIntegrationsBetaApi.GetBigSegmentStoreIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**integrationKey** | **string** | The integration key, either &#x60;redis&#x60; or &#x60;dynamodb&#x60; | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBigSegmentStoreIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BigSegmentStoreIntegration**](BigSegmentStoreIntegration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBigSegmentStoreIntegrations

> BigSegmentStoreIntegrationCollection GetBigSegmentStoreIntegrations(ctx).Execute()

List all big segment store integrations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersistentStoreIntegrationsBetaApi.GetBigSegmentStoreIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersistentStoreIntegrationsBetaApi.GetBigSegmentStoreIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBigSegmentStoreIntegrations`: BigSegmentStoreIntegrationCollection
	fmt.Fprintf(os.Stdout, "Response from `PersistentStoreIntegrationsBetaApi.GetBigSegmentStoreIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBigSegmentStoreIntegrationsRequest struct via the builder pattern


### Return type

[**BigSegmentStoreIntegrationCollection**](BigSegmentStoreIntegrationCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchBigSegmentStoreIntegration

> BigSegmentStoreIntegration PatchBigSegmentStoreIntegration(ctx, projectKey, environmentKey, integrationKey, integrationId).PatchOperation(patchOperation).Execute()

Update big segment store integration



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
	environmentKey := "environmentKey_example" // string | The environment key
	integrationKey := "integrationKey_example" // string | The integration key, either `redis` or `dynamodb`
	integrationId := "integrationId_example" // string | The integration ID
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersistentStoreIntegrationsBetaApi.PatchBigSegmentStoreIntegration(context.Background(), projectKey, environmentKey, integrationKey, integrationId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersistentStoreIntegrationsBetaApi.PatchBigSegmentStoreIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchBigSegmentStoreIntegration`: BigSegmentStoreIntegration
	fmt.Fprintf(os.Stdout, "Response from `PersistentStoreIntegrationsBetaApi.PatchBigSegmentStoreIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**integrationKey** | **string** | The integration key, either &#x60;redis&#x60; or &#x60;dynamodb&#x60; | 
**integrationId** | **string** | The integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBigSegmentStoreIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**BigSegmentStoreIntegration**](BigSegmentStoreIntegration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

