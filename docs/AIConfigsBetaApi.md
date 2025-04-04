# \AIConfigsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAIConfig**](AIConfigsBetaApi.md#DeleteAIConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Delete AI Config
[**DeleteAIConfigVariation**](AIConfigsBetaApi.md#DeleteAIConfigVariation) | **Delete** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Delete AI Config variation
[**DeleteModelConfig**](AIConfigsBetaApi.md#DeleteModelConfig) | **Delete** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Delete an AI model config
[**GetAIConfig**](AIConfigsBetaApi.md#GetAIConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Get AI Config
[**GetAIConfigMetrics**](AIConfigsBetaApi.md#GetAIConfigMetrics) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics | Get AI Config metrics
[**GetAIConfigMetricsByVariation**](AIConfigsBetaApi.md#GetAIConfigMetricsByVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/metrics-by-variation | Get AI Config metrics by variation
[**GetAIConfigVariation**](AIConfigsBetaApi.md#GetAIConfigVariation) | **Get** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Get AI Config variation
[**GetAIConfigs**](AIConfigsBetaApi.md#GetAIConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs | List AI Configs
[**GetModelConfig**](AIConfigsBetaApi.md#GetModelConfig) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs/{modelConfigKey} | Get AI model config
[**ListModelConfigs**](AIConfigsBetaApi.md#ListModelConfigs) | **Get** /api/v2/projects/{projectKey}/ai-configs/model-configs | List AI model configs
[**PatchAIConfig**](AIConfigsBetaApi.md#PatchAIConfig) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey} | Update AI Config
[**PatchAIConfigVariation**](AIConfigsBetaApi.md#PatchAIConfigVariation) | **Patch** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations/{variationKey} | Update AI Config variation
[**PostAIConfig**](AIConfigsBetaApi.md#PostAIConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs | Create new AI Config
[**PostAIConfigVariation**](AIConfigsBetaApi.md#PostAIConfigVariation) | **Post** /api/v2/projects/{projectKey}/ai-configs/{configKey}/variations | Create AI Config variation
[**PostModelConfig**](AIConfigsBetaApi.md#PostModelConfig) | **Post** /api/v2/projects/{projectKey}/ai-configs/model-configs | Create an AI model config



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
    openapiclient "./openapi"
)

func main() {
    lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
    projectKey := "default" // string | 
    configKey := "configKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIConfigsBetaApi.DeleteAIConfig(context.Background(), projectKey, configKey).LDAPIVersion(lDAPIVersion).Execute()
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
    openapiclient "./openapi"
)

func main() {
    lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
    projectKey := "projectKey_example" // string | 
    configKey := "configKey_example" // string | 
    variationKey := "variationKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIConfigsBetaApi.DeleteAIConfigVariation(context.Background(), projectKey, configKey, variationKey).LDAPIVersion(lDAPIVersion).Execute()
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
    openapiclient "./openapi"
)

func main() {
    lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
    projectKey := "default" // string | 
    modelConfigKey := "modelConfigKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIConfigsBetaApi.DeleteModelConfig(context.Background(), projectKey, modelConfigKey).LDAPIVersion(lDAPIVersion).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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


## ListModelConfigs

> []ModelConfig ListModelConfigs(ctx, projectKey).LDAPIVersion(lDAPIVersion).Execute()

List AI model configs



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
    lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
    projectKey := "default" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIConfigsBetaApi.ListModelConfigs(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Execute()
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
)

func main() {
    lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
    projectKey := "projectKey_example" // string | 
    configKey := "configKey_example" // string | 
    aIConfigVariationPost := *openapiclient.NewAIConfigVariationPost("Key_example", []openapiclient.Message{*openapiclient.NewMessage("Content_example", "Role_example")}, map[string]interface{}(123), "Name_example") // AIConfigVariationPost | AI Config variation object to create

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
    openapiclient "./openapi"
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

