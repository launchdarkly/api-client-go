# \AIBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGenerateExperiment**](AIBetaApi.md#PostGenerateExperiment) | **Post** /api/v2/ai/projects/{projectKey}/environments/{environmentKey}/generate-experiment | Generate an experiment config using AI.
[**PostGenerateVariations**](AIBetaApi.md#PostGenerateVariations) | **Post** /api/v2/ai/generate-variations | Generate new variations for a feature flag



## PostGenerateExperiment

> AIGeneratedExperimentConfig PostGenerateExperiment(ctx, projectKey, environmentKey).GenerateExperimentPost(generateExperimentPost).Execute()

Generate an experiment config using AI.



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
    projectKey := "projectKey_example" // string | Project key
    environmentKey := "environmentKey_example" // string | Environment key
    generateExperimentPost := *openapiclient.NewGenerateExperimentPost() // GenerateExperimentPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIBetaApi.PostGenerateExperiment(context.Background(), projectKey, environmentKey).GenerateExperimentPost(generateExperimentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIBetaApi.PostGenerateExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGenerateExperiment`: AIGeneratedExperimentConfig
    fmt.Fprintf(os.Stdout, "Response from `AIBetaApi.PostGenerateExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | Project key | 
**environmentKey** | **string** | Environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGenerateExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **generateExperimentPost** | [**GenerateExperimentPost**](GenerateExperimentPost.md) |  | 

### Return type

[**AIGeneratedExperimentConfig**](AIGeneratedExperimentConfig.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGenerateVariations

> AIGeneratedVariations PostGenerateVariations(ctx).GenerateVariationsPost(generateVariationsPost).Execute()

Generate new variations for a feature flag



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
    generateVariationsPost := *openapiclient.NewGenerateVariationsPost() // GenerateVariationsPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AIBetaApi.PostGenerateVariations(context.Background()).GenerateVariationsPost(generateVariationsPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIBetaApi.PostGenerateVariations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGenerateVariations`: AIGeneratedVariations
    fmt.Fprintf(os.Stdout, "Response from `AIBetaApi.PostGenerateVariations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGenerateVariationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateVariationsPost** | [**GenerateVariationsPost**](GenerateVariationsPost.md) |  | 

### Return type

[**AIGeneratedVariations**](AIGeneratedVariations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

