# \AIConfigsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAIToolReferences**](AIConfigsApi.md#ListAIToolReferences) | **Get** /api/v2/projects/{projectKey}/ai-tools/{toolKey}/references | List AI tool references
[**ListAgentOptimizationRuns**](AIConfigsApi.md#ListAgentOptimizationRuns) | **Get** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/runs | List agent optimization runs



## ListAIToolReferences

> ToolReferences ListAIToolReferences(ctx, projectKey, toolKey).Limit(limit).Offset(offset).Execute()

List AI tool references



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
	projectKey := "projectKey_example" // string | 
	toolKey := "toolKey_example" // string | 
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.ListAIToolReferences(context.Background(), projectKey, toolKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.ListAIToolReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIToolReferences`: ToolReferences
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.ListAIToolReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**toolKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**ToolReferences**](ToolReferences.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentOptimizationRuns

> AgentOptimizationRuns ListAgentOptimizationRuns(ctx, projectKey, optimizationKey).Limit(limit).Offset(offset).Execute()

List agent optimization runs



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
	projectKey := "projectKey_example" // string | 
	optimizationKey := "optimizationKey_example" // string | 
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.ListAgentOptimizationRuns(context.Background(), projectKey, optimizationKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.ListAgentOptimizationRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentOptimizationRuns`: AgentOptimizationRuns
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.ListAgentOptimizationRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentOptimizationRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**AgentOptimizationRuns**](AgentOptimizationRuns.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

