# \AgentControlApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAIConfig**](AgentControlApi.md#DeleteAIConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Delete AI Config
[**DeleteAIConfigVariation**](AgentControlApi.md#DeleteAIConfigVariation) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Delete AI Config variation
[**DeleteAITool**](AgentControlApi.md#DeleteAITool) | **Delete** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Delete AI tool
[**DeleteAgentGraph**](AgentControlApi.md#DeleteAgentGraph) | **Delete** /api/v2/projects/{projectKey}/agent-graphs/{graphKey} | Delete agent graph
[**DeleteAgentOptimization**](AgentControlApi.md#DeleteAgentOptimization) | **Delete** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey} | Delete an agent optimization
[**DeleteAgentOptimizationRun**](AgentControlApi.md#DeleteAgentOptimizationRun) | **Delete** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/runs/{runId} | Delete an agent optimization run
[**DeleteModelConfig**](AgentControlApi.md#DeleteModelConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Delete an AI model config
[**DeletePromptSnippet**](AgentControlApi.md#DeletePromptSnippet) | **Delete** /api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey} | Delete a prompt snippet
[**DeleteRestrictedModels**](AgentControlApi.md#DeleteRestrictedModels) | **Delete** /api/v2/projects/{projectKey}/ai-configs/model-configs/restricted | Remove AI models from the restricted list
[**GetAIConfig**](AgentControlApi.md#GetAIConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Get AI Config
[**GetAIConfigMetrics**](AgentControlApi.md#GetAIConfigMetrics) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics | Get AI Config metrics
[**GetAIConfigMetricsByVariation**](AgentControlApi.md#GetAIConfigMetricsByVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics-by-variation | Get AI Config metrics by variation
[**GetAIConfigQuickStats**](AgentControlApi.md#GetAIConfigQuickStats) | **Get** /api/v2/projects/{projectKey}/ai-configs/quick-stats | Get AI Config quick stats
[**GetAIConfigTargeting**](AgentControlApi.md#GetAIConfigTargeting) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting | Show an AI Config&#39;s targeting
[**GetAIConfigVariation**](AgentControlApi.md#GetAIConfigVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Get AI Config variation
[**GetAIConfigs**](AgentControlApi.md#GetAIConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs | List AI Configs
[**GetAITool**](AgentControlApi.md#GetAITool) | **Get** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Get AI tool
[**GetAgentGraph**](AgentControlApi.md#GetAgentGraph) | **Get** /api/v2/projects/{projectKey}/agent-graphs/{graphKey} | Get agent graph
[**GetAgentOptimization**](AgentControlApi.md#GetAgentOptimization) | **Get** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey} | Get an agent optimization
[**GetModelConfig**](AgentControlApi.md#GetModelConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Get AI model config
[**GetPromptSnippet**](AgentControlApi.md#GetPromptSnippet) | **Get** /api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey} | Get a prompt snippet
[**ListAIToolVersions**](AgentControlApi.md#ListAIToolVersions) | **Get** /api/v2/projects/{projectKey}/ai-tools/{toolKey}/versions | List AI tool versions
[**ListAITools**](AgentControlApi.md#ListAITools) | **Get** /api/v2/projects/{projectKey}/ai-tools | List AI tools
[**ListAgentGraphs**](AgentControlApi.md#ListAgentGraphs) | **Get** /api/v2/projects/{projectKey}/agent-graphs | List agent graphs
[**ListAgentOptimizationResults**](AgentControlApi.md#ListAgentOptimizationResults) | **Get** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/results | List agent optimization runs
[**ListAgentOptimizationResultsByRunId**](AgentControlApi.md#ListAgentOptimizationResultsByRunId) | **Get** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/runs/{runId}/results | List agent optimization results for a run
[**ListAgentOptimizations**](AgentControlApi.md#ListAgentOptimizations) | **Get** /api/v2/projects/{projectKey}/agent-optimizations | List agent optimizations
[**ListAllAgentOptimizationResults**](AgentControlApi.md#ListAllAgentOptimizationResults) | **Get** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/all-results | List all agent optimization results across versions
[**ListModelConfigs**](AgentControlApi.md#ListModelConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs | List AI model configs
[**ListPromptSnippetReferences**](AgentControlApi.md#ListPromptSnippetReferences) | **Get** /api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey}/references | List prompt snippet references
[**ListPromptSnippetVersions**](AgentControlApi.md#ListPromptSnippetVersions) | **Get** /api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey}/versions | List prompt snippet versions
[**ListPromptSnippets**](AgentControlApi.md#ListPromptSnippets) | **Get** /api/v2/projects/{projectKey}/ai-configs/prompt-snippets | List prompt snippets
[**PatchAIConfig**](AgentControlApi.md#PatchAIConfig) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Update AI Config
[**PatchAIConfigTargeting**](AgentControlApi.md#PatchAIConfigTargeting) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting | Update AI Config targeting
[**PatchAIConfigVariation**](AgentControlApi.md#PatchAIConfigVariation) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Update AI Config variation
[**PatchAITool**](AgentControlApi.md#PatchAITool) | **Patch** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Update AI tool
[**PatchAgentGraph**](AgentControlApi.md#PatchAgentGraph) | **Patch** /api/v2/projects/{projectKey}/agent-graphs/{graphKey} | Update agent graph
[**PatchAgentOptimization**](AgentControlApi.md#PatchAgentOptimization) | **Patch** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey} | Update an agent optimization
[**PatchAgentOptimizationResult**](AgentControlApi.md#PatchAgentOptimizationResult) | **Patch** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/results/{resultId} | Update an agent optimization result
[**PatchPromptSnippet**](AgentControlApi.md#PatchPromptSnippet) | **Patch** /api/v2/projects/{projectKey}/ai-configs/prompt-snippets/{snippetKey} | Update a prompt snippet
[**PostAIConfig**](AgentControlApi.md#PostAIConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs | Create new AI Config
[**PostAIConfigVariation**](AgentControlApi.md#PostAIConfigVariation) | **Post** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations | Create AI Config variation
[**PostAITool**](AgentControlApi.md#PostAITool) | **Post** /api/v2/projects/{projectKey}/ai-tools | Create an AI tool
[**PostAgentGraph**](AgentControlApi.md#PostAgentGraph) | **Post** /api/v2/projects/{projectKey}/agent-graphs | Create new agent graph
[**PostAgentOptimization**](AgentControlApi.md#PostAgentOptimization) | **Post** /api/v2/projects/{projectKey}/agent-optimizations | Create agent optimization
[**PostAgentOptimizationResult**](AgentControlApi.md#PostAgentOptimizationResult) | **Post** /api/v2/projects/{projectKey}/agent-optimizations/{optimizationKey}/results | Create agent optimization result
[**PostModelConfig**](AgentControlApi.md#PostModelConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs/model-configs | Create an AI model config
[**PostPromptSnippet**](AgentControlApi.md#PostPromptSnippet) | **Post** /api/v2/projects/{projectKey}/ai-configs/prompt-snippets | Create a prompt snippet
[**PostRestrictedModels**](AgentControlApi.md#PostRestrictedModels) | **Post** /api/v2/projects/{projectKey}/ai-configs/model-configs/restricted | Add AI models to the restricted list



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
	r, err := apiClient.AgentControlApi.DeleteAIConfig(context.Background(), projectKey, configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteAIConfig``: %v\n", err)
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
	r, err := apiClient.AgentControlApi.DeleteAIConfigVariation(context.Background(), projectKey, configKey, variationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteAIConfigVariation``: %v\n", err)
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
	r, err := apiClient.AgentControlApi.DeleteAITool(context.Background(), projectKey, toolKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteAITool``: %v\n", err)
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
	r, err := apiClient.AgentControlApi.DeleteAgentGraph(context.Background(), projectKey, graphKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteAgentGraph``: %v\n", err)
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


## DeleteAgentOptimization

> DeleteAgentOptimization(ctx, projectKey, optimizationKey).Execute()

Delete an agent optimization



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentControlApi.DeleteAgentOptimization(context.Background(), projectKey, optimizationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteAgentOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentOptimizationRequest struct via the builder pattern


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


## DeleteAgentOptimizationRun

> DeleteAgentOptimizationRun(ctx, projectKey, optimizationKey, runId).LDAPIVersion(lDAPIVersion).Execute()

Delete an agent optimization run



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
	optimizationKey := "optimizationKey_example" // string | 
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentControlApi.DeleteAgentOptimizationRun(context.Background(), projectKey, optimizationKey, runId).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteAgentOptimizationRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentOptimizationRunRequest struct via the builder pattern


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
	r, err := apiClient.AgentControlApi.DeleteModelConfig(context.Background(), projectKey, modelConfigKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteModelConfig``: %v\n", err)
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


## DeletePromptSnippet

> DeletePromptSnippet(ctx, projectKey, snippetKey).Execute()

Delete a prompt snippet



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
	snippetKey := "snippetKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AgentControlApi.DeletePromptSnippet(context.Background(), projectKey, snippetKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeletePromptSnippet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**snippetKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePromptSnippetRequest struct via the builder pattern


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
	r, err := apiClient.AgentControlApi.DeleteRestrictedModels(context.Background(), projectKey).RestrictedModelsRequest(restrictedModelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.DeleteRestrictedModels``: %v\n", err)
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
	resp, r, err := apiClient.AgentControlApi.GetAIConfig(context.Background(), projectKey, configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAIConfig`: %v\n", resp)
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

> Metrics GetAIConfigMetrics(ctx, projectKey, configKey).From(from).To(to).Env(env).ContextKind(contextKind).ContextKey(contextKey).Execute()

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
	contextKind := "contextKind_example" // string | A context kind. Only metrics from events that include a context of this kind are included. Required if `contextKey` is provided. (optional)
	contextKey := "contextKey_example" // string | A context key. Only metrics from events whose context of the `contextKind` kind has this key are included. Requires `contextKind`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.GetAIConfigMetrics(context.Background(), projectKey, configKey).From(from).To(to).Env(env).ContextKind(contextKind).ContextKey(contextKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAIConfigMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAIConfigMetrics`: %v\n", resp)
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
 **contextKind** | **string** | A context kind. Only metrics from events that include a context of this kind are included. Required if &#x60;contextKey&#x60; is provided. | 
 **contextKey** | **string** | A context key. Only metrics from events whose context of the &#x60;contextKind&#x60; kind has this key are included. Requires &#x60;contextKind&#x60;. | 

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

> []MetricByVariation GetAIConfigMetricsByVariation(ctx, projectKey, configKey).From(from).To(to).Env(env).ContextKind(contextKind).ContextKey(contextKey).Execute()

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
	contextKind := "contextKind_example" // string | A context kind. Only metrics from events that include a context of this kind are included. Required if `contextKey` is provided. (optional)
	contextKey := "contextKey_example" // string | A context key. Only metrics from events whose context of the `contextKind` kind has this key are included. Requires `contextKind`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.GetAIConfigMetricsByVariation(context.Background(), projectKey, configKey).From(from).To(to).Env(env).ContextKind(contextKind).ContextKey(contextKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAIConfigMetricsByVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigMetricsByVariation`: []MetricByVariation
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAIConfigMetricsByVariation`: %v\n", resp)
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
 **contextKind** | **string** | A context kind. Only metrics from events that include a context of this kind are included. Required if &#x60;contextKey&#x60; is provided. | 
 **contextKey** | **string** | A context key. Only metrics from events whose context of the &#x60;contextKind&#x60; kind has this key are included. Requires &#x60;contextKind&#x60;. | 

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


## GetAIConfigQuickStats

> QuickStats GetAIConfigQuickStats(ctx, projectKey).Env(env).Execute()

Get AI Config quick stats



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
	env := "env_example" // string | An environment key. Only metrics from this environment will be included.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.GetAIConfigQuickStats(context.Background(), projectKey).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAIConfigQuickStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigQuickStats`: QuickStats
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAIConfigQuickStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAIConfigQuickStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | An environment key. Only metrics from this environment will be included. | 

### Return type

[**QuickStats**](QuickStats.md)

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
	resp, r, err := apiClient.AgentControlApi.GetAIConfigTargeting(context.Background(), projectKey, configKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAIConfigTargeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigTargeting`: AIConfigTargeting
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAIConfigTargeting`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.GetAIConfigVariation(context.Background(), projectKey, configKey, variationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigVariation`: AIConfigVariationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAIConfigVariation`: %v\n", resp)
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
	sort := "sort_example" // string | A sort to apply to the list of AgentControl configs. (optional)
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.GetAIConfigs(context.Background(), projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAIConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigs`: AIConfigs
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAIConfigs`: %v\n", resp)
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

 **sort** | **string** | A sort to apply to the list of AgentControl configs. | 
 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list. | 

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
	resp, r, err := apiClient.AgentControlApi.GetAITool(context.Background(), projectKey, toolKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAITool`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.GetAgentGraph(context.Background(), projectKey, graphKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAgentGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentGraph`: AgentGraph
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAgentGraph`: %v\n", resp)
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


## GetAgentOptimization

> AgentOptimization GetAgentOptimization(ctx, projectKey, optimizationKey).Execute()

Get an agent optimization



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.GetAgentOptimization(context.Background(), projectKey, optimizationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetAgentOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentOptimization`: AgentOptimization
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetAgentOptimization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentOptimization**](AgentOptimization.md)

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
	resp, r, err := apiClient.AgentControlApi.GetModelConfig(context.Background(), projectKey, modelConfigKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetModelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelConfig`: ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetModelConfig`: %v\n", resp)
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


## GetPromptSnippet

> PromptSnippet GetPromptSnippet(ctx, projectKey, snippetKey).Execute()

Get a prompt snippet



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
	snippetKey := "snippetKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.GetPromptSnippet(context.Background(), projectKey, snippetKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.GetPromptSnippet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPromptSnippet`: PromptSnippet
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.GetPromptSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**snippetKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPromptSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PromptSnippet**](PromptSnippet.md)

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
	sort := "sort_example" // string | A sort to apply to the list of AgentControl configs. (optional)
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListAIToolVersions(context.Background(), projectKey, toolKey).Sort(sort).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListAIToolVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIToolVersions`: AITools
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListAIToolVersions`: %v\n", resp)
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


 **sort** | **string** | A sort to apply to the list of AgentControl configs. | 
 **limit** | **int32** | The number of resources to return. | 
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
	sort := "sort_example" // string | A sort to apply to the list of AgentControl configs. (optional)
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListAITools(context.Background(), projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListAITools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAITools`: AITools
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListAITools`: %v\n", resp)
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

 **sort** | **string** | A sort to apply to the list of AgentControl configs. | 
 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list. | 

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

> AgentGraphs ListAgentGraphs(ctx, projectKey).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Filter(filter).Execute()

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
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListAgentGraphs(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListAgentGraphs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentGraphs`: AgentGraphs
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListAgentGraphs`: %v\n", resp)
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

 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list. | 

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


## ListAgentOptimizationResults

> AgentOptimizationResults ListAgentOptimizationResults(ctx, projectKey, optimizationKey).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.AgentControlApi.ListAgentOptimizationResults(context.Background(), projectKey, optimizationKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListAgentOptimizationResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentOptimizationResults`: AgentOptimizationResults
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListAgentOptimizationResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentOptimizationResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**AgentOptimizationResults**](AgentOptimizationResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentOptimizationResultsByRunId

> AgentOptimizationResults ListAgentOptimizationResultsByRunId(ctx, projectKey, optimizationKey, runId).Execute()

List agent optimization results for a run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListAgentOptimizationResultsByRunId(context.Background(), projectKey, optimizationKey, runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListAgentOptimizationResultsByRunId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentOptimizationResultsByRunId`: AgentOptimizationResults
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListAgentOptimizationResultsByRunId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentOptimizationResultsByRunIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AgentOptimizationResults**](AgentOptimizationResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentOptimizations

> AgentOptimizations ListAgentOptimizations(ctx, projectKey).Limit(limit).Offset(offset).Filter(filter).Execute()

List agent optimizations



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
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListAgentOptimizations(context.Background(), projectKey).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListAgentOptimizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAgentOptimizations`: AgentOptimizations
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListAgentOptimizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentOptimizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list. | 

### Return type

[**AgentOptimizations**](AgentOptimizations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllAgentOptimizationResults

> AgentOptimizationResults ListAllAgentOptimizationResults(ctx, projectKey, optimizationKey).Limit(limit).Offset(offset).Execute()

List all agent optimization results across versions



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
	resp, r, err := apiClient.AgentControlApi.ListAllAgentOptimizationResults(context.Background(), projectKey, optimizationKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListAllAgentOptimizationResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllAgentOptimizationResults`: AgentOptimizationResults
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListAllAgentOptimizationResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllAgentOptimizationResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**AgentOptimizationResults**](AgentOptimizationResults.md)

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
	resp, r, err := apiClient.AgentControlApi.ListModelConfigs(context.Background(), projectKey).Restricted(restricted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListModelConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListModelConfigs`: []ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListModelConfigs`: %v\n", resp)
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


## ListPromptSnippetReferences

> SnippetReferences ListPromptSnippetReferences(ctx, projectKey, snippetKey).Limit(limit).Offset(offset).Execute()

List prompt snippet references



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
	snippetKey := "snippetKey_example" // string | 
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListPromptSnippetReferences(context.Background(), projectKey, snippetKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListPromptSnippetReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPromptSnippetReferences`: SnippetReferences
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListPromptSnippetReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**snippetKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPromptSnippetReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**SnippetReferences**](SnippetReferences.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPromptSnippetVersions

> PromptSnippets ListPromptSnippetVersions(ctx, projectKey, snippetKey).Limit(limit).Offset(offset).Execute()

List prompt snippet versions



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
	snippetKey := "snippetKey_example" // string | 
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListPromptSnippetVersions(context.Background(), projectKey, snippetKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListPromptSnippetVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPromptSnippetVersions`: PromptSnippets
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListPromptSnippetVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**snippetKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPromptSnippetVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**PromptSnippets**](PromptSnippets.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPromptSnippets

> PromptSnippets ListPromptSnippets(ctx, projectKey).Limit(limit).Offset(offset).Filter(filter).Execute()

List prompt snippets



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
	limit := int32(56) // int32 | The number of resources to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.ListPromptSnippets(context.Background(), projectKey).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.ListPromptSnippets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPromptSnippets`: PromptSnippets
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.ListPromptSnippets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPromptSnippetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of resources to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list. | 

### Return type

[**PromptSnippets**](PromptSnippets.md)

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
	resp, r, err := apiClient.AgentControlApi.PatchAIConfig(context.Background(), projectKey, configKey).AIConfigPatch(aIConfigPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchAIConfig`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.PatchAIConfigTargeting(context.Background(), projectKey, configKey).AIConfigTargetingPatch(aIConfigTargetingPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchAIConfigTargeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfigTargeting`: AIConfigTargeting
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchAIConfigTargeting`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.PatchAIConfigVariation(context.Background(), projectKey, configKey, variationKey).AIConfigVariationPatch(aIConfigVariationPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfigVariation`: AIConfigVariation
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchAIConfigVariation`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.PatchAITool(context.Background(), projectKey, toolKey).AIToolPatch(aIToolPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchAITool`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.PatchAgentGraph(context.Background(), projectKey, graphKey).LDAPIVersion(lDAPIVersion).AgentGraphPatch(agentGraphPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchAgentGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentGraph`: AgentGraph
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchAgentGraph`: %v\n", resp)
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


## PatchAgentOptimization

> AgentOptimization PatchAgentOptimization(ctx, projectKey, optimizationKey).AgentOptimizationPatch(agentOptimizationPatch).Execute()

Update an agent optimization



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
	agentOptimizationPatch := *openapiclient.NewAgentOptimizationPatch() // AgentOptimizationPatch | Agent optimization fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.PatchAgentOptimization(context.Background(), projectKey, optimizationKey).AgentOptimizationPatch(agentOptimizationPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchAgentOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentOptimization`: AgentOptimization
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchAgentOptimization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentOptimizationPatch** | [**AgentOptimizationPatch**](AgentOptimizationPatch.md) | Agent optimization fields to update | 

### Return type

[**AgentOptimization**](AgentOptimization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentOptimizationResult

> AgentOptimizationResult PatchAgentOptimizationResult(ctx, projectKey, optimizationKey, resultId).AgentOptimizationResultPatch(agentOptimizationResultPatch).Execute()

Update an agent optimization result



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
	resultId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	agentOptimizationResultPatch := *openapiclient.NewAgentOptimizationResultPatch() // AgentOptimizationResultPatch | Agent optimization result fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.PatchAgentOptimizationResult(context.Background(), projectKey, optimizationKey, resultId).AgentOptimizationResultPatch(agentOptimizationResultPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchAgentOptimizationResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentOptimizationResult`: AgentOptimizationResult
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchAgentOptimizationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 
**resultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentOptimizationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **agentOptimizationResultPatch** | [**AgentOptimizationResultPatch**](AgentOptimizationResultPatch.md) | Agent optimization result fields to update | 

### Return type

[**AgentOptimizationResult**](AgentOptimizationResult.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPromptSnippet

> PromptSnippet PatchPromptSnippet(ctx, projectKey, snippetKey).PromptSnippetPatch(promptSnippetPatch).Execute()

Update a prompt snippet



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
	snippetKey := "snippetKey_example" // string | 
	promptSnippetPatch := *openapiclient.NewPromptSnippetPatch() // PromptSnippetPatch | Prompt snippet fields to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.PatchPromptSnippet(context.Background(), projectKey, snippetKey).PromptSnippetPatch(promptSnippetPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PatchPromptSnippet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPromptSnippet`: PromptSnippet
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PatchPromptSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**snippetKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPromptSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **promptSnippetPatch** | [**PromptSnippetPatch**](PromptSnippetPatch.md) | Prompt snippet fields to update | 

### Return type

[**PromptSnippet**](PromptSnippet.md)

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
	resp, r, err := apiClient.AgentControlApi.PostAIConfig(context.Background(), projectKey).AIConfigPost(aIConfigPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostAIConfig`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.PostAIConfigVariation(context.Background(), projectKey, configKey).AIConfigVariationPost(aIConfigVariationPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAIConfigVariation`: AIConfigVariation
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostAIConfigVariation`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.PostAITool(context.Background(), projectKey).AIToolPost(aIToolPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostAITool`: %v\n", resp)
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
	resp, r, err := apiClient.AgentControlApi.PostAgentGraph(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).AgentGraphPost(agentGraphPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostAgentGraph``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentGraph`: AgentGraph
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostAgentGraph`: %v\n", resp)
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


## PostAgentOptimization

> AgentOptimization PostAgentOptimization(ctx, projectKey).AgentOptimizationPost(agentOptimizationPost).Execute()

Create agent optimization



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
	agentOptimizationPost := *openapiclient.NewAgentOptimizationPost("Key_example", "AiConfigKey_example", int32(123), "JudgeModel_example") // AgentOptimizationPost | Agent optimization object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.PostAgentOptimization(context.Background(), projectKey).AgentOptimizationPost(agentOptimizationPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostAgentOptimization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentOptimization`: AgentOptimization
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostAgentOptimization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentOptimizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentOptimizationPost** | [**AgentOptimizationPost**](AgentOptimizationPost.md) | Agent optimization object to create | 

### Return type

[**AgentOptimization**](AgentOptimization.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAgentOptimizationResult

> AgentOptimizationResult PostAgentOptimizationResult(ctx, projectKey, optimizationKey).AgentOptimizationResultPost(agentOptimizationResultPost).Execute()

Create agent optimization result



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
	agentOptimizationResultPost := *openapiclient.NewAgentOptimizationResultPost("RunId_example", int32(123), int32(123), "Instructions_example", "UserInput_example") // AgentOptimizationResultPost | Agent optimization result to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.PostAgentOptimizationResult(context.Background(), projectKey, optimizationKey).AgentOptimizationResultPost(agentOptimizationResultPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostAgentOptimizationResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAgentOptimizationResult`: AgentOptimizationResult
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostAgentOptimizationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**optimizationKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAgentOptimizationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agentOptimizationResultPost** | [**AgentOptimizationResultPost**](AgentOptimizationResultPost.md) | Agent optimization result to create | 

### Return type

[**AgentOptimizationResult**](AgentOptimizationResult.md)

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
	resp, r, err := apiClient.AgentControlApi.PostModelConfig(context.Background(), projectKey).ModelConfigPost(modelConfigPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostModelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostModelConfig`: ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostModelConfig`: %v\n", resp)
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


## PostPromptSnippet

> PromptSnippet PostPromptSnippet(ctx, projectKey).PromptSnippetPost(promptSnippetPost).Execute()

Create a prompt snippet



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
	promptSnippetPost := *openapiclient.NewPromptSnippetPost("Key_example", "Name_example", "Text_example") // PromptSnippetPost | Prompt snippet object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentControlApi.PostPromptSnippet(context.Background(), projectKey).PromptSnippetPost(promptSnippetPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostPromptSnippet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPromptSnippet`: PromptSnippet
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostPromptSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPromptSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **promptSnippetPost** | [**PromptSnippetPost**](PromptSnippetPost.md) | Prompt snippet object to create | 

### Return type

[**PromptSnippet**](PromptSnippet.md)

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
	resp, r, err := apiClient.AgentControlApi.PostRestrictedModels(context.Background(), projectKey).RestrictedModelsRequest(restrictedModelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlApi.PostRestrictedModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRestrictedModels`: RestrictedModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentControlApi.PostRestrictedModels`: %v\n", resp)
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

