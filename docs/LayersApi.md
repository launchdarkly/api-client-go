# \LayersApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLayer**](LayersApi.md#CreateLayer) | **Post** /api/v2/projects/{projectKey}/layers | Create layer
[**GetLayers**](LayersApi.md#GetLayers) | **Get** /api/v2/projects/{projectKey}/layers | Get layers
[**UpdateLayer**](LayersApi.md#UpdateLayer) | **Patch** /api/v2/projects/{projectKey}/layers/{layerKey} | Update layer



## CreateLayer

> LayerRep CreateLayer(ctx, projectKey).LayerPost(layerPost).Execute()

Create layer



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
	layerPost := *openapiclient.NewLayerPost("checkout-flow", "Checkout Flow", "Description_example") // LayerPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LayersApi.CreateLayer(context.Background(), projectKey).LayerPost(layerPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersApi.CreateLayer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLayer`: LayerRep
	fmt.Fprintf(os.Stdout, "Response from `LayersApi.CreateLayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **layerPost** | [**LayerPost**](LayerPost.md) |  | 

### Return type

[**LayerRep**](LayerRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLayers

> LayerCollectionRep GetLayers(ctx, projectKey).Filter(filter).Execute()

Get layers



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
	filter := "filter_example" // string | A comma-separated list of filters. This endpoint only accepts filtering by `experimentKey`. The filter returns layers which include that experiment for the selected environment(s). For example: `filter=reservations.experimentKey contains expKey`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LayersApi.GetLayers(context.Background(), projectKey).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersApi.GetLayers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLayers`: LayerCollectionRep
	fmt.Fprintf(os.Stdout, "Response from `LayersApi.GetLayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A comma-separated list of filters. This endpoint only accepts filtering by &#x60;experimentKey&#x60;. The filter returns layers which include that experiment for the selected environment(s). For example: &#x60;filter&#x3D;reservations.experimentKey contains expKey&#x60;. | 

### Return type

[**LayerCollectionRep**](LayerCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLayer

> LayerRep UpdateLayer(ctx, projectKey, layerKey).LayerPatchInput(layerPatchInput).Execute()

Update layer



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
	layerKey := "layerKey_example" // string | The layer key
	layerPatchInput := *openapiclient.NewLayerPatchInput([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // LayerPatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LayersApi.UpdateLayer(context.Background(), projectKey, layerKey).LayerPatchInput(layerPatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LayersApi.UpdateLayer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLayer`: LayerRep
	fmt.Fprintf(os.Stdout, "Response from `LayersApi.UpdateLayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**layerKey** | **string** | The layer key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **layerPatchInput** | [**LayerPatchInput**](LayerPatchInput.md) |  | 

### Return type

[**LayerRep**](LayerRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

