# \AIConfigsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAIConfig**](AIConfigsBetaApi.md#DeleteAIConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Delete AI Config
[**DeleteAIConfigVariation**](AIConfigsBetaApi.md#DeleteAIConfigVariation) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Delete AI Config variation
[**DeleteAITool**](AIConfigsBetaApi.md#DeleteAITool) | **Delete** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Delete AI tool
[**DeleteModelConfig**](AIConfigsBetaApi.md#DeleteModelConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Delete an AI model config
[**DeleteRestrictedModels**](AIConfigsBetaApi.md#DeleteRestrictedModels) | **Delete** /api/v2/projects/{projectKey}/ai-configs/model-configs/restricted | Remove AI models from the restricted list
[**GetAIConfig**](AIConfigsBetaApi.md#GetAIConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Get AI Config
[**GetAIConfigMetrics**](AIConfigsBetaApi.md#GetAIConfigMetrics) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics | Get AI Config metrics
[**GetAIConfigMetricsByVariation**](AIConfigsBetaApi.md#GetAIConfigMetricsByVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics-by-variation | Get AI Config metrics by variation
[**GetAIConfigTargeting**](AIConfigsBetaApi.md#GetAIConfigTargeting) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting | Show an AI Config&#39;s targeting
[**GetAIConfigVariation**](AIConfigsBetaApi.md#GetAIConfigVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Get AI Config variation
[**GetAIConfigs**](AIConfigsBetaApi.md#GetAIConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs | List AI Configs
[**GetAITool**](AIConfigsBetaApi.md#GetAITool) | **Get** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Get AI tool
[**GetModelConfig**](AIConfigsBetaApi.md#GetModelConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Get AI model config
[**ListAIToolVersions**](AIConfigsBetaApi.md#ListAIToolVersions) | **Get** /api/v2/projects/{projectKey}/ai-tools/{toolKey}/versions | List AI tool versions
[**ListAITools**](AIConfigsBetaApi.md#ListAITools) | **Get** /api/v2/projects/{projectKey}/ai-tools | List AI tools
[**ListModelConfigs**](AIConfigsBetaApi.md#ListModelConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs | List AI model configs
[**PatchAIConfig**](AIConfigsBetaApi.md#PatchAIConfig) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Update AI Config
[**PatchAIConfigTargeting**](AIConfigsBetaApi.md#PatchAIConfigTargeting) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey}/targeting | Update AI Config targeting
[**PatchAIConfigVariation**](AIConfigsBetaApi.md#PatchAIConfigVariation) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Update AI Config variation
[**PatchAITool**](AIConfigsBetaApi.md#PatchAITool) | **Patch** /api/v2/projects/{projectKey}/ai-tools/{toolKey} | Update AI tool
[**PostAIConfig**](AIConfigsBetaApi.md#PostAIConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs | Create new AI Config
[**PostAIConfigVariation**](AIConfigsBetaApi.md#PostAIConfigVariation) | **Post** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations | Create AI Config variation
[**PostAITool**](AIConfigsBetaApi.md#PostAITool) | **Post** /api/v2/projects/{projectKey}/ai-tools | Create an AI tool
[**PostModelConfig**](AIConfigsBetaApi.md#PostModelConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs/model-configs | Create an AI model config
[**PostRestrictedModels**](AIConfigsBetaApi.md#PostRestrictedModels) | **Post** /api/v2/projects/{projectKey}/ai-configs/model-configs/restricted | Add AI models to the restricted list



## DeleteAIConfig

> DeleteAIConfig(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsBetaApi.DeleteAIConfig(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.DeleteAIConfig``: %v\n", err)
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


## DeleteAIConfigVariation

> DeleteAIConfigVariation(ctx, projectKey, configKey, variationKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 
	variationKey := "variationKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsBetaApi.DeleteAIConfigVariation(context.Background(), projectKey, configKey, variationKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.DeleteAIConfigVariation``: %v\n", err)
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


## DeleteAITool

> DeleteAITool(ctx, projectKey, toolKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	toolKey := "toolKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsBetaApi.DeleteAITool(context.Background(), projectKey, toolKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.DeleteAITool``: %v\n", err)
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

> DeleteModelConfig(ctx, projectKey, modelConfigKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	modelConfigKey := "modelConfigKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsBetaApi.DeleteModelConfig(context.Background(), projectKey, modelConfigKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.DeleteModelConfig``: %v\n", err)
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


## DeleteRestrictedModels

> DeleteRestrictedModels(ctx, projectKey).LDAPIVersion(lDAPIVersion).RestrictedModelsRequest(restrictedModelsRequest).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	restrictedModelsRequest := *openapiclient.NewRestrictedModelsRequest([]string{"Keys_example"}) // RestrictedModelsRequest | List of AI model keys to remove from the restricted list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIConfigsBetaApi.DeleteRestrictedModels(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).RestrictedModelsRequest(restrictedModelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.DeleteRestrictedModels``: %v\n", err)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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

> AIConfig GetAIConfig(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetAIConfig(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetAIConfig`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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

> Metrics GetAIConfigMetrics(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).From(from).To(to).Env(env).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 
	from := int32(56) // int32 | The starting time, as milliseconds since epoch (inclusive).
	to := int32(56) // int32 | The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after `from`.
	env := "env_example" // string | An environment key. Only metrics from this environment will be included.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetAIConfigMetrics(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).From(from).To(to).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetAIConfigMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigMetrics`: Metrics
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetAIConfigMetrics`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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

> []MetricByVariation GetAIConfigMetricsByVariation(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).From(from).To(to).Env(env).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 
	from := int32(56) // int32 | The starting time, as milliseconds since epoch (inclusive).
	to := int32(56) // int32 | The ending time, as milliseconds since epoch (exclusive). May not be more than 100 days after `from`.
	env := "env_example" // string | An environment key. Only metrics from this environment will be included.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetAIConfigMetricsByVariation(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).From(from).To(to).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetAIConfigMetricsByVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigMetricsByVariation`: []MetricByVariation
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetAIConfigMetricsByVariation`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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

> AIConfigTargeting GetAIConfigTargeting(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetAIConfigTargeting(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetAIConfigTargeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigTargeting`: AIConfigTargeting
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetAIConfigTargeting`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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

> AIConfigVariationsResponse GetAIConfigVariation(ctx, projectKey, configKey, variationKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	configKey := "default" // string | 
	variationKey := "default" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetAIConfigVariation(context.Background(), projectKey, configKey, variationKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigVariation`: AIConfigVariationsResponse
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetAIConfigVariation`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 




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

> AIConfigs GetAIConfigs(ctx, projectKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	sort := "sort_example" // string | A sort to apply to the list of AI Configs. (optional)
	limit := int32(56) // int32 | The number of AI Configs to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of AI Configs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetAIConfigs(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetAIConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAIConfigs`: AIConfigs
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetAIConfigs`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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

> AITool GetAITool(ctx, projectKey, toolKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	toolKey := "toolKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetAITool(context.Background(), projectKey, toolKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetAITool`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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


## GetModelConfig

> ModelConfig GetModelConfig(ctx, projectKey, modelConfigKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	modelConfigKey := "default" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.GetModelConfig(context.Background(), projectKey, modelConfigKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.GetModelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelConfig`: ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.GetModelConfig`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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

> AITools ListAIToolVersions(ctx, projectKey, toolKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	toolKey := "toolKey_example" // string | 
	sort := "sort_example" // string | A sort to apply to the list of AI Configs. (optional)
	limit := int32(56) // int32 | The number of AI Configs to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.ListAIToolVersions(context.Background(), projectKey, toolKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.ListAIToolVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAIToolVersions`: AITools
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.ListAIToolVersions`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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

> AITools ListAITools(ctx, projectKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	sort := "sort_example" // string | A sort to apply to the list of AI Configs. (optional)
	limit := int32(56) // int32 | The number of AI Configs to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of AI Configs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.ListAITools(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.ListAITools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAITools`: AITools
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.ListAITools`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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


## ListModelConfigs

> []ModelConfig ListModelConfigs(ctx, projectKey).LDAPIVersion(lDAPIVersion).Restricted(restricted).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	restricted := true // bool | Whether to return only restricted models (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.ListModelConfigs(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Restricted(restricted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.ListModelConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListModelConfigs`: []ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.ListModelConfigs`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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

> AIConfig PatchAIConfig(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).AIConfigPatch(aIConfigPatch).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 
	aIConfigPatch := *openapiclient.NewAIConfigPatch() // AIConfigPatch | AI Config object to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PatchAIConfig(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).AIConfigPatch(aIConfigPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PatchAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PatchAIConfig`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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

> AIConfigTargeting PatchAIConfigTargeting(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).AIConfigTargetingPatch(aIConfigTargetingPatch).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 
	aIConfigTargetingPatch := *openapiclient.NewAIConfigTargetingPatch("EnvironmentKey_example", []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // AIConfigTargetingPatch | AI Config targeting semantic patch instructions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PatchAIConfigTargeting(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).AIConfigTargetingPatch(aIConfigTargetingPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PatchAIConfigTargeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfigTargeting`: AIConfigTargeting
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PatchAIConfigTargeting`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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

> AIConfigVariation PatchAIConfigVariation(ctx, projectKey, configKey, variationKey).LDAPIVersion(lDAPIVersion).AIConfigVariationPatch(aIConfigVariationPatch).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 
	variationKey := "variationKey_example" // string | 
	aIConfigVariationPatch := *openapiclient.NewAIConfigVariationPatch() // AIConfigVariationPatch | AI Config variation object to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PatchAIConfigVariation(context.Background(), projectKey, configKey, variationKey).LDAPIVersion(lDAPIVersion).AIConfigVariationPatch(aIConfigVariationPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PatchAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAIConfigVariation`: AIConfigVariation
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PatchAIConfigVariation`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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

> AITool PatchAITool(ctx, projectKey, toolKey).LDAPIVersion(lDAPIVersion).AIToolPatch(aIToolPatch).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	toolKey := "toolKey_example" // string | 
	aIToolPatch := *openapiclient.NewAIToolPatch() // AIToolPatch | AI tool object to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PatchAITool(context.Background(), projectKey, toolKey).LDAPIVersion(lDAPIVersion).AIToolPatch(aIToolPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PatchAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PatchAITool`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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


## PostAIConfig

> AIConfig PostAIConfig(ctx, projectKey).LDAPIVersion(lDAPIVersion).AIConfigPost(aIConfigPost).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	aIConfigPost := *openapiclient.NewAIConfigPost("Key_example", "Name_example") // AIConfigPost | AI Config object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PostAIConfig(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).AIConfigPost(aIConfigPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PostAIConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAIConfig`: AIConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PostAIConfig`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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

> AIConfigVariation PostAIConfigVariation(ctx, projectKey, configKey).LDAPIVersion(lDAPIVersion).AIConfigVariationPost(aIConfigVariationPost).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	configKey := "configKey_example" // string | 
	aIConfigVariationPost := *openapiclient.NewAIConfigVariationPost("Key_example", "Name_example") // AIConfigVariationPost | AI Config variation object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PostAIConfigVariation(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).AIConfigVariationPost(aIConfigVariationPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PostAIConfigVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAIConfigVariation`: AIConfigVariation
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PostAIConfigVariation`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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

> AITool PostAITool(ctx, projectKey).LDAPIVersion(lDAPIVersion).AIToolPost(aIToolPost).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "projectKey_example" // string | 
	aIToolPost := *openapiclient.NewAIToolPost("Key_example", map[string]interface{}(123)) // AIToolPost | AI tool object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PostAITool(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).AIToolPost(aIToolPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PostAITool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAITool`: AITool
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PostAITool`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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


## PostModelConfig

> ModelConfig PostModelConfig(ctx, projectKey).LDAPIVersion(lDAPIVersion).ModelConfigPost(modelConfigPost).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	modelConfigPost := *openapiclient.NewModelConfigPost("Name_example", "Key_example", "Id_example") // ModelConfigPost | AI model config object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PostModelConfig(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).ModelConfigPost(modelConfigPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PostModelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostModelConfig`: ModelConfig
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PostModelConfig`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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

> RestrictedModelsResponse PostRestrictedModels(ctx, projectKey).LDAPIVersion(lDAPIVersion).RestrictedModelsRequest(restrictedModelsRequest).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	restrictedModelsRequest := *openapiclient.NewRestrictedModelsRequest([]string{"Keys_example"}) // RestrictedModelsRequest | List of AI model keys to add to the restricted list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIConfigsBetaApi.PostRestrictedModels(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).RestrictedModelsRequest(restrictedModelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIConfigsBetaApi.PostRestrictedModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRestrictedModels`: RestrictedModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `AIConfigsBetaApi.PostRestrictedModels`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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

