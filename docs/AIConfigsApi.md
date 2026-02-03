# \AIConfigsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAIConfig**](AIConfigsApi.md#DeleteAIConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Delete AI Config
[**DeleteAIConfigVariation**](AIConfigsApi.md#DeleteAIConfigVariation) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Delete AI Config variation
[**DeleteAITool**](AIConfigsApi.md#DeleteAITool) | **Delete** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Delete AI tool
[**DeleteAgentGraph**](AIConfigsApi.md#DeleteAgentGraph) | **Delete** /api/v2/projects/{projectKey}/agent-graphs/{graphKey} | Delete agent graph
[**DeleteModelConfig**](AIConfigsApi.md#DeleteModelConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Delete an AI model config
[**DeleteRestrictedModels**](AIConfigsApi.md#DeleteRestrictedModels) | **Delete** /api/v2/projects/{projectKey}/ai-configs/model-configs/restricted | Remove AI models from the restricted list
[**GetAIConfig**](AIConfigsApi.md#GetAIConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Get AI Config
[**GetAIConfigMetrics**](AIConfigsApi.md#GetAIConfigMetrics) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics | Get AI Config metrics
[**GetAIConfigMetricsByVariation**](AIConfigsApi.md#GetAIConfigMetricsByVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics-by-variation | Get AI Config metrics by variation
[**GetAIConfigTargeting**](AIConfigsApi.md#GetAIConfigTargeting) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting | Show an AI Config&#39;s targeting
[**GetAIConfigVariation**](AIConfigsApi.md#GetAIConfigVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Get AI Config variation
[**GetAIConfigs**](AIConfigsApi.md#GetAIConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs | List AI Configs
[**GetAITool**](AIConfigsApi.md#GetAITool) | **Get** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Get AI tool
[**GetAgentGraph**](AIConfigsApi.md#GetAgentGraph) | **Get** /api/v2/projects/{projectKey}/agent-graphs/{graphKey} | Get agent graph
[**GetModelConfig**](AIConfigsApi.md#GetModelConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Get AI model config
[**ListAIToolVersions**](AIConfigsApi.md#ListAIToolVersions) | **Get** /api/v2/projects/{projectKey}/ai-tools/{toolKey}/versions | List AI tool versions
[**ListAITools**](AIConfigsApi.md#ListAITools) | **Get** /api/v2/projects/{projectKey}/ai-tools | List AI tools
[**ListAgentGraphs**](AIConfigsApi.md#ListAgentGraphs) | **Get** /api/v2/projects/{projectKey}/agent-graphs | List agent graphs
[**ListModelConfigs**](AIConfigsApi.md#ListModelConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs | List AI model configs
[**PatchAIConfig**](AIConfigsApi.md#PatchAIConfig) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Update AI Config
[**PatchAIConfigTargeting**](AIConfigsApi.md#PatchAIConfigTargeting) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting | Update AI Config targeting
[**PatchAIConfigVariation**](AIConfigsApi.md#PatchAIConfigVariation) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Update AI Config variation
[**PatchAITool**](AIConfigsApi.md#PatchAITool) | **Patch** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Update AI tool
[**PatchAgentGraph**](AIConfigsApi.md#PatchAgentGraph) | **Patch** /api/v2/projects/{projectKey}/agent-graphs/{graphKey} | Update agent graph
[**PostAIConfig**](AIConfigsApi.md#PostAIConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs | Create new AI Config
[**PostAIConfigVariation**](AIConfigsApi.md#PostAIConfigVariation) | **Post** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations | Create AI Config variation
[**PostAITool**](AIConfigsApi.md#PostAITool) | **Post** /api/v2/projects/{projectKey}/ai-tools | Create an AI tool
[**PostAgentGraph**](AIConfigsApi.md#PostAgentGraph) | **Post** /api/v2/projects/{projectKey}/agent-graphs | Create new agent graph
[**PostModelConfig**](AIConfigsApi.md#PostModelConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs/model-configs | Create an AI model config
[**PostRestrictedModels**](AIConfigsApi.md#PostRestrictedModels) | **Post** /api/v2/projects/{projectKey}/ai-configs/model-configs/restricted | Add AI models to the restricted list



## DeleteAIConfig

> DeleteAIConfig(ctx, projectKey, configKey).Execute()

Delete AI Config



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
	projectKey := "default" // string | 
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsApi.DeleteAIConfig(context.Background(), projectKey, configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.DeleteAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAIConfigRequest struct via the builder pattern


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


## DeleteAIConfigVariation

> DeleteAIConfigVariation(ctx, projectKey, configKey, variationKey).Execute()

Delete AI Config variation



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
	configKey := "configKey_example" // string | 
	variationKey := "variationKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsApi.DeleteAIConfigVariation(context.Background(), projectKey, configKey, variationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.DeleteAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 
**variationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAIConfigVariationRequest struct via the builder pattern


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


## DeleteAITool

> DeleteAITool(ctx, projectKey, toolKey).Execute()

Delete AI tool



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsApi.DeleteAITool(context.Background(), projectKey, toolKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.DeleteAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**toolKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAIToolRequest struct via the builder pattern


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


## DeleteAgentGraph

> DeleteAgentGraph(ctx, projectKey, graphKey).LDAPIVersion(lDAPIVersion).Execute()

Delete agent graph



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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	graphKey := "graphKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsApi.DeleteAgentGraph(context.Background(), projectKey, graphKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.DeleteAgentGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**graphKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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


## DeleteModelConfig

> DeleteModelConfig(ctx, projectKey, modelConfigKey).Execute()

Delete an AI model config



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
	projectKey := "default" // string | 
	modelConfigKey := "modelConfigKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsApi.DeleteModelConfig(context.Background(), projectKey, modelConfigKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.DeleteModelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**modelConfigKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModelConfigRequest struct via the builder pattern


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


## DeleteRestrictedModels

> DeleteRestrictedModels(ctx, projectKey).RestrictedModelsRequest(restrictedModelsRequest).Execute()

Remove AI models from the restricted list



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
	projectKey := "default" // string | 
	restrictedModelsRequest := *openapiclient.NewRestrictedModelsRequest([]string{"Keys_example"}) // RestrictedModelsRequest | List of AI model keys to remove from the restricted list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsApi.DeleteRestrictedModels(context.Background(), projectKey).RestrictedModelsRequest(restrictedModelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.DeleteRestrictedModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRestrictedModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restrictedModelsRequest** | [**RestrictedModelsRequest**](RestrictedModelsRequest.md) | List of AI model keys to remove from the restricted list | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIConfig

> AIConfig GetAIConfig(ctx, projectKey, configKey).Execute()

Get AI Config



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
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAIConfig(context.Background(), projectKey, configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAIConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AIConfig**](AIConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIConfigMetrics

> Metrics GetAIConfigMetrics(ctx, projectKey, configKey).From(from).To(to).Env(env).Execute()

Get AI Config metrics



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
	configKey := "configKey_example" // string | 
	from := int32(56) // int32 | The starting time, as milliseconds since epoch (inclusive).
	to := int32(56) // int32 | The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after `from`.
	env := "env_example" // string | An environment key. Only metrics from this environment will be included.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAIConfigMetrics(context.Background(), projectKey, configKey).From(from).To(to).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAIConfigMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAIConfigMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **int32** | The starting time, as milliseconds since epoch (inclusive). | 
 **to** | **int32** | The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after &#x60;from&#x60;. | 
 **env** | **string** | An environment key. Only metrics from this environment will be included. | 

### Return type

[**Metrics**](Metrics.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIConfigMetricsByVariation

> []MetricByVariation GetAIConfigMetricsByVariation(ctx, projectKey, configKey).From(from).To(to).Env(env).Execute()

Get AI Config metrics by variation



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
	configKey := "configKey_example" // string | 
	from := int32(56) // int32 | The starting time, as milliseconds since epoch (inclusive).
	to := int32(56) // int32 | The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after `from`.
	env := "env_example" // string | An environment key. Only metrics from this environment will be included.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAIConfigMetricsByVariation(context.Background(), projectKey, configKey).From(from).To(to).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAIConfigMetricsByVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigMetricsByVariation`: []MetricByVariation
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAIConfigMetricsByVariation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigMetricsByVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **int32** | The starting time, as milliseconds since epoch (inclusive). | 
 **to** | **int32** | The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after &#x60;from&#x60;. | 
 **env** | **string** | An environment key. Only metrics from this environment will be included. | 

### Return type

[**[]MetricByVariation**](MetricByVariation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIConfigTargeting

> AIConfigTargeting GetAIConfigTargeting(ctx, projectKey, configKey).Execute()

Show an AI Config's targeting



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
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAIConfigTargeting(context.Background(), projectKey, configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAIConfigTargeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigTargeting`: AIConfigTargeting
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAIConfigTargeting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigTargetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AIConfigTargeting**](AIConfigTargeting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIConfigVariation

> AIConfigVariationsResponse GetAIConfigVariation(ctx, projectKey, configKey, variationKey).Execute()

Get AI Config variation



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
	projectKey := "default" // string | 
	configKey := "default" // string | 
	variationKey := "default" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAIConfigVariation(context.Background(), projectKey, configKey, variationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigVariation`: AIConfigVariationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAIConfigVariation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 
**variationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AIConfigVariationsResponse**](AIConfigVariationsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAIConfigs

> AIConfigs GetAIConfigs(ctx, projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()

List AI Configs



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
	projectKey := "default" // string | 
	sort := "sort_example" // string | A sort to apply to the list of AI Configs. (optional)
	limit := int32(56) // int32 | The number of AI Configs to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of AI Configs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAIConfigs(context.Background(), projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAIConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigs`: AIConfigs
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAIConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | A sort to apply to the list of AI Configs. | 
 **limit** | **int32** | The number of AI Configs to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list of AI Configs. | 

### Return type

[**AIConfigs**](AIConfigs.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAITool

> AITool GetAITool(ctx, projectKey, toolKey).Execute()

Get AI tool



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAITool(context.Background(), projectKey, toolKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAITool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**toolKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AITool**](AITool.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentGraph

> AgentGraph GetAgentGraph(ctx, projectKey, graphKey).LDAPIVersion(lDAPIVersion).Execute()

Get agent graph



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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	graphKey := "graphKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetAgentGraph(context.Background(), projectKey, graphKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetAgentGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentGraph`: AgentGraph
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetAgentGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**graphKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



### Return type

[**AgentGraph**](AgentGraph.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelConfig

> ModelConfig GetModelConfig(ctx, projectKey, modelConfigKey).Execute()

Get AI model config



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
	projectKey := "default" // string | 
	modelConfigKey := "default" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.GetModelConfig(context.Background(), projectKey, modelConfigKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.GetModelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelConfig`: ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.GetModelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**modelConfigKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelConfig**](ModelConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAIToolVersions

> AITools ListAIToolVersions(ctx, projectKey, toolKey).Sort(sort).Limit(limit).Offset(offset).Execute()

List AI tool versions



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
	sort := "sort_example" // string | A sort to apply to the list of AI Configs. (optional)
	limit := int32(56) // int32 | The number of AI Configs to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.ListAIToolVersions(context.Background(), projectKey, toolKey).Sort(sort).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.ListAIToolVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIToolVersions`: AITools
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.ListAIToolVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**toolKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **string** | A sort to apply to the list of AI Configs. | 
 **limit** | **int32** | The number of AI Configs to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**AITools**](AITools.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAITools

> AITools ListAITools(ctx, projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()

List AI tools



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
	sort := "sort_example" // string | A sort to apply to the list of AI Configs. (optional)
	limit := int32(56) // int32 | The number of AI Configs to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of AI Configs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.ListAITools(context.Background(), projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.ListAITools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAITools`: AITools
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.ListAITools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAIToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | A sort to apply to the list of AI Configs. | 
 **limit** | **int32** | The number of AI Configs to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list of AI Configs. | 

### Return type

[**AITools**](AITools.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentGraphs

> AgentGraphs ListAgentGraphs(ctx, projectKey).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Execute()

List agent graphs



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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	limit := int32(56) // int32 | The number of AI Configs to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.ListAgentGraphs(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.ListAgentGraphs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentGraphs`: AgentGraphs
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.ListAgentGraphs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentGraphsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **limit** | **int32** | The number of AI Configs to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**AgentGraphs**](AgentGraphs.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelConfigs

> []ModelConfig ListModelConfigs(ctx, projectKey).Restricted(restricted).Execute()

List AI model configs



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
	projectKey := "default" // string | 
	restricted := true // bool | Whether to return only restricted models (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.ListModelConfigs(context.Background(), projectKey).Restricted(restricted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.ListModelConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListModelConfigs`: []ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.ListModelConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListModelConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restricted** | **bool** | Whether to return only restricted models | 

### Return type

[**[]ModelConfig**](ModelConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAIConfig

> AIConfig PatchAIConfig(ctx, projectKey, configKey).AIConfigPatch(aIConfigPatch).Execute()

Update AI Config



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
	configKey := "configKey_example" // string | 
	aIConfigPatch := *openapiclient.NewAIConfigPatch() // AIConfigPatch | AI Config object to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PatchAIConfig(context.Background(), projectKey, configKey).AIConfigPatch(aIConfigPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PatchAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PatchAIConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAIConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aIConfigPatch** | [**AIConfigPatch**](AIConfigPatch.md) | AI Config object to update | 

### Return type

[**AIConfig**](AIConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAIConfigTargeting

> AIConfigTargeting PatchAIConfigTargeting(ctx, projectKey, configKey).AIConfigTargetingPatch(aIConfigTargetingPatch).Execute()

Update AI Config targeting



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
	configKey := "configKey_example" // string | 
	aIConfigTargetingPatch := *openapiclient.NewAIConfigTargetingPatch("EnvironmentKey_example", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // AIConfigTargetingPatch | AI Config targeting semantic patch instructions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PatchAIConfigTargeting(context.Background(), projectKey, configKey).AIConfigTargetingPatch(aIConfigTargetingPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PatchAIConfigTargeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfigTargeting`: AIConfigTargeting
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PatchAIConfigTargeting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAIConfigTargetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aIConfigTargetingPatch** | [**AIConfigTargetingPatch**](AIConfigTargetingPatch.md) | AI Config targeting semantic patch instructions | 

### Return type

[**AIConfigTargeting**](AIConfigTargeting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAIConfigVariation

> AIConfigVariation PatchAIConfigVariation(ctx, projectKey, configKey, variationKey).AIConfigVariationPatch(aIConfigVariationPatch).Execute()

Update AI Config variation



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
	configKey := "configKey_example" // string | 
	variationKey := "variationKey_example" // string | 
	aIConfigVariationPatch := *openapiclient.NewAIConfigVariationPatch() // AIConfigVariationPatch | AI Config variation object to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PatchAIConfigVariation(context.Background(), projectKey, configKey, variationKey).AIConfigVariationPatch(aIConfigVariationPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PatchAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfigVariation`: AIConfigVariation
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PatchAIConfigVariation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 
**variationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAIConfigVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **aIConfigVariationPatch** | [**AIConfigVariationPatch**](AIConfigVariationPatch.md) | AI Config variation object to update | 

### Return type

[**AIConfigVariation**](AIConfigVariation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAITool

> AITool PatchAITool(ctx, projectKey, toolKey).AIToolPatch(aIToolPatch).Execute()

Update AI tool



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
	aIToolPatch := *openapiclient.NewAIToolPatch() // AIToolPatch | AI tool object to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PatchAITool(context.Background(), projectKey, toolKey).AIToolPatch(aIToolPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PatchAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PatchAITool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**toolKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAIToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aIToolPatch** | [**AIToolPatch**](AIToolPatch.md) | AI tool object to update | 

### Return type

[**AITool**](AITool.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentGraph

> AgentGraph PatchAgentGraph(ctx, projectKey, graphKey).LDAPIVersion(lDAPIVersion).AgentGraphPatch(agentGraphPatch).Execute()

Update agent graph



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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	graphKey := "graphKey_example" // string | 
	agentGraphPatch := *openapiclient.NewAgentGraphPatch() // AgentGraphPatch | Agent graph object to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PatchAgentGraph(context.Background(), projectKey, graphKey).LDAPIVersion(lDAPIVersion).AgentGraphPatch(agentGraphPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PatchAgentGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentGraph`: AgentGraph
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PatchAgentGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**graphKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **agentGraphPatch** | [**AgentGraphPatch**](AgentGraphPatch.md) | Agent graph object to update | 

### Return type

[**AgentGraph**](AgentGraph.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAIConfig

> AIConfig PostAIConfig(ctx, projectKey).AIConfigPost(aIConfigPost).Execute()

Create new AI Config



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
	aIConfigPost := *openapiclient.NewAIConfigPost("Key_example", "Name_example") // AIConfigPost | AI Config object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PostAIConfig(context.Background(), projectKey).AIConfigPost(aIConfigPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PostAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PostAIConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAIConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aIConfigPost** | [**AIConfigPost**](AIConfigPost.md) | AI Config object to create | 

### Return type

[**AIConfig**](AIConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAIConfigVariation

> AIConfigVariation PostAIConfigVariation(ctx, projectKey, configKey).AIConfigVariationPost(aIConfigVariationPost).Execute()

Create AI Config variation



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
	configKey := "configKey_example" // string | 
	aIConfigVariationPost := *openapiclient.NewAIConfigVariationPost("Key_example", "Name_example") // AIConfigVariationPost | AI Config variation object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PostAIConfigVariation(context.Background(), projectKey, configKey).AIConfigVariationPost(aIConfigVariationPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PostAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAIConfigVariation`: AIConfigVariation
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PostAIConfigVariation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**configKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAIConfigVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aIConfigVariationPost** | [**AIConfigVariationPost**](AIConfigVariationPost.md) | AI Config variation object to create | 

### Return type

[**AIConfigVariation**](AIConfigVariation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAITool

> AITool PostAITool(ctx, projectKey).AIToolPost(aIToolPost).Execute()

Create an AI tool



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
	aIToolPost := *openapiclient.NewAIToolPost("Key_example", map[string]interface{}(123)) // AIToolPost | AI tool object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PostAITool(context.Background(), projectKey).AIToolPost(aIToolPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PostAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PostAITool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAIToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aIToolPost** | [**AIToolPost**](AIToolPost.md) | AI tool object to create | 

### Return type

[**AITool**](AITool.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentGraph

> AgentGraph PostAgentGraph(ctx, projectKey).LDAPIVersion(lDAPIVersion).AgentGraphPost(agentGraphPost).Execute()

Create new agent graph



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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	agentGraphPost := *openapiclient.NewAgentGraphPost("Key_example", "Name_example") // AgentGraphPost | Agent graph object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PostAgentGraph(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).AgentGraphPost(agentGraphPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PostAgentGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentGraph`: AgentGraph
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PostAgentGraph`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentGraphRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **agentGraphPost** | [**AgentGraphPost**](AgentGraphPost.md) | Agent graph object to create | 

### Return type

[**AgentGraph**](AgentGraph.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostModelConfig

> ModelConfig PostModelConfig(ctx, projectKey).ModelConfigPost(modelConfigPost).Execute()

Create an AI model config



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
	projectKey := "default" // string | 
	modelConfigPost := *openapiclient.NewModelConfigPost("Name_example", "Key_example", "Id_example") // ModelConfigPost | AI model config object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PostModelConfig(context.Background(), projectKey).ModelConfigPost(modelConfigPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PostModelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostModelConfig`: ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PostModelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostModelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelConfigPost** | [**ModelConfigPost**](ModelConfigPost.md) | AI model config object to create | 

### Return type

[**ModelConfig**](ModelConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRestrictedModels

> RestrictedModelsResponse PostRestrictedModels(ctx, projectKey).RestrictedModelsRequest(restrictedModelsRequest).Execute()

Add AI models to the restricted list



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
	projectKey := "default" // string | 
	restrictedModelsRequest := *openapiclient.NewRestrictedModelsRequest([]string{"Keys_example"}) // RestrictedModelsRequest | List of AI model keys to add to the restricted list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsApi.PostRestrictedModels(context.Background(), projectKey).RestrictedModelsRequest(restrictedModelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsApi.PostRestrictedModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRestrictedModels`: RestrictedModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsApi.PostRestrictedModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRestrictedModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restrictedModelsRequest** | [**RestrictedModelsRequest**](RestrictedModelsRequest.md) | List of AI model keys to add to the restricted list. | 

### Return type

[**RestrictedModelsResponse**](RestrictedModelsResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

