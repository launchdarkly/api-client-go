# \ExperimentsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExperiment**](ExperimentsBetaApi.md#CreateExperiment) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/experiments | Create experiment
[**CreateIteration**](ExperimentsBetaApi.md#CreateIteration) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/experiments/{experimentKey}/iterations | Create iteration
[**GetExperiment**](ExperimentsBetaApi.md#GetExperiment) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/experiments/{experimentKey} | Get experiment
[**GetExperimentResults**](ExperimentsBetaApi.md#GetExperimentResults) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/experiments/{experimentKey}/metrics/{metricKey}/results | Get experiment results
[**GetExperimentResultsForMetricGroup**](ExperimentsBetaApi.md#GetExperimentResultsForMetricGroup) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/experiments/{experimentKey}/metric-groups/{metricGroupKey}/results | Get experiment results for metric group
[**GetExperimentationSettings**](ExperimentsBetaApi.md#GetExperimentationSettings) | **Get** /api/v2/projects/{projectKey}/experimentation-settings | Get experimentation settings
[**GetExperiments**](ExperimentsBetaApi.md#GetExperiments) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/experiments | Get experiments
[**GetLegacyExperimentResults**](ExperimentsBetaApi.md#GetLegacyExperimentResults) | **Get** /api/v2/flags/{projectKey}/{featureFlagKey}/experiments/{environmentKey}/{metricKey} | Get legacy experiment results (deprecated)
[**PatchExperiment**](ExperimentsBetaApi.md#PatchExperiment) | **Patch** /api/v2/projects/{projectKey}/environments/{environmentKey}/experiments/{experimentKey} | Patch experiment
[**PutExperimentationSettings**](ExperimentsBetaApi.md#PutExperimentationSettings) | **Put** /api/v2/projects/{projectKey}/experimentation-settings | Update experimentation settings
[**ResetExperiment**](ExperimentsBetaApi.md#ResetExperiment) | **Delete** /api/v2/flags/{projectKey}/{featureFlagKey}/experiments/{environmentKey}/{metricKey}/results | Reset experiment results



## CreateExperiment

> Experiment CreateExperiment(ctx, projectKey, environmentKey).ExperimentPost(experimentPost).Execute()

Create experiment



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
    experimentPost := *openapiclient.NewExperimentPost("Example experiment", "experiment-key-123abc", *openapiclient.NewIterationInput("Example hypothesis, the new button placement will increase conversion", []openapiclient.MetricInput{*openapiclient.NewMetricInput("metric-key-123abc", true)}, []openapiclient.TreatmentInput{*openapiclient.NewTreatmentInput("Treatment 1", true, "10", []openapiclient.TreatmentParameterInput{*openapiclient.NewTreatmentParameterInput("example-flag-for-experiment", "e432f62b-55f6-49dd-a02f-eb24acf39d05")})}, map[string]FlagInput{"key": *openapiclient.NewFlagInput("e432f62b-55f6-49dd-a02f-eb24acf39d05", int32(12))})) // ExperimentPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.CreateExperiment(context.Background(), projectKey, environmentKey).ExperimentPost(experimentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.CreateExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExperiment`: Experiment
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.CreateExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **experimentPost** | [**ExperimentPost**](ExperimentPost.md) |  | 

### Return type

[**Experiment**](Experiment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIteration

> IterationRep CreateIteration(ctx, projectKey, environmentKey, experimentKey).IterationInput(iterationInput).Execute()

Create iteration



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
    experimentKey := "experimentKey_example" // string | The experiment key
    iterationInput := *openapiclient.NewIterationInput("Example hypothesis, the new button placement will increase conversion", []openapiclient.MetricInput{*openapiclient.NewMetricInput("metric-key-123abc", true)}, []openapiclient.TreatmentInput{*openapiclient.NewTreatmentInput("Treatment 1", true, "10", []openapiclient.TreatmentParameterInput{*openapiclient.NewTreatmentParameterInput("example-flag-for-experiment", "e432f62b-55f6-49dd-a02f-eb24acf39d05")})}, map[string]FlagInput{"key": *openapiclient.NewFlagInput("e432f62b-55f6-49dd-a02f-eb24acf39d05", int32(12))}) // IterationInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.CreateIteration(context.Background(), projectKey, environmentKey, experimentKey).IterationInput(iterationInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.CreateIteration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIteration`: IterationRep
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.CreateIteration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**experimentKey** | **string** | The experiment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **iterationInput** | [**IterationInput**](IterationInput.md) |  | 

### Return type

[**IterationRep**](IterationRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperiment

> Experiment GetExperiment(ctx, projectKey, environmentKey, experimentKey).Expand(expand).Execute()

Get experiment



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
    experimentKey := "experimentKey_example" // string | The experiment key
    expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. Supported fields are explained above. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.GetExperiment(context.Background(), projectKey, environmentKey, experimentKey).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.GetExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperiment`: Experiment
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.GetExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**experimentKey** | **string** | The experiment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. Supported fields are explained above. | 

### Return type

[**Experiment**](Experiment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentResults

> ExperimentBayesianResultsRep GetExperimentResults(ctx, projectKey, environmentKey, experimentKey, metricKey).IterationId(iterationId).Execute()

Get experiment results



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
    experimentKey := "experimentKey_example" // string | The experiment key
    metricKey := "metricKey_example" // string | The metric key
    iterationId := "iterationId_example" // string | The iteration ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.GetExperimentResults(context.Background(), projectKey, environmentKey, experimentKey, metricKey).IterationId(iterationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.GetExperimentResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperimentResults`: ExperimentBayesianResultsRep
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.GetExperimentResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**experimentKey** | **string** | The experiment key | 
**metricKey** | **string** | The metric key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **iterationId** | **string** | The iteration ID | 

### Return type

[**ExperimentBayesianResultsRep**](ExperimentBayesianResultsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentResultsForMetricGroup

> MetricGroupResultsRep GetExperimentResultsForMetricGroup(ctx, projectKey, environmentKey, experimentKey, metricGroupKey).IterationId(iterationId).Execute()

Get experiment results for metric group



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
    experimentKey := "experimentKey_example" // string | The experiment key
    metricGroupKey := "metricGroupKey_example" // string | The metric group key
    iterationId := "iterationId_example" // string | The iteration ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.GetExperimentResultsForMetricGroup(context.Background(), projectKey, environmentKey, experimentKey, metricGroupKey).IterationId(iterationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.GetExperimentResultsForMetricGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperimentResultsForMetricGroup`: MetricGroupResultsRep
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.GetExperimentResultsForMetricGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**experimentKey** | **string** | The experiment key | 
**metricGroupKey** | **string** | The metric group key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentResultsForMetricGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **iterationId** | **string** | The iteration ID | 

### Return type

[**MetricGroupResultsRep**](MetricGroupResultsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentationSettings

> ExperimentationSettingsRep GetExperimentationSettings(ctx, projectKey).Execute()

Get experimentation settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.GetExperimentationSettings(context.Background(), projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.GetExperimentationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperimentationSettings`: ExperimentationSettingsRep
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.GetExperimentationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExperimentationSettingsRep**](ExperimentationSettingsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperiments

> ExperimentCollectionRep GetExperiments(ctx, projectKey, environmentKey).Limit(limit).Offset(offset).Filter(filter).Expand(expand).LifecycleState(lifecycleState).Execute()

Get experiments



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
    limit := int64(789) // int64 | The maximum number of experiments to return. Defaults to 20. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form `field:value`. Supported fields are explained above. (optional)
    expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. Supported fields are explained above. (optional)
    lifecycleState := "lifecycleState_example" // string | A comma-separated list of experiment archived states. Supports `archived`, `active`, or both. Defaults to `active` experiments. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.GetExperiments(context.Background(), projectKey, environmentKey).Limit(limit).Offset(offset).Filter(filter).Expand(expand).LifecycleState(lifecycleState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.GetExperiments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperiments`: ExperimentCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.GetExperiments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int64** | The maximum number of experiments to return. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form &#x60;field:value&#x60;. Supported fields are explained above. | 
 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. Supported fields are explained above. | 
 **lifecycleState** | **string** | A comma-separated list of experiment archived states. Supports &#x60;archived&#x60;, &#x60;active&#x60;, or both. Defaults to &#x60;active&#x60; experiments. | 

### Return type

[**ExperimentCollectionRep**](ExperimentCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLegacyExperimentResults

> ExperimentResults GetLegacyExperimentResults(ctx, projectKey, featureFlagKey, environmentKey, metricKey).From(from).To(to).Execute()

Get legacy experiment results (deprecated)



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    metricKey := "metricKey_example" // string | The metric key
    from := int64(789) // int64 | A timestamp denoting the start of the data collection period, expressed as a Unix epoch time in milliseconds. (optional)
    to := int64(789) // int64 | A timestamp denoting the end of the data collection period, expressed as a Unix epoch time in milliseconds. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.GetLegacyExperimentResults(context.Background(), projectKey, featureFlagKey, environmentKey, metricKey).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.GetLegacyExperimentResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLegacyExperimentResults`: ExperimentResults
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.GetLegacyExperimentResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**metricKey** | **string** | The metric key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLegacyExperimentResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **from** | **int64** | A timestamp denoting the start of the data collection period, expressed as a Unix epoch time in milliseconds. | 
 **to** | **int64** | A timestamp denoting the end of the data collection period, expressed as a Unix epoch time in milliseconds. | 

### Return type

[**ExperimentResults**](ExperimentResults.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExperiment

> Experiment PatchExperiment(ctx, projectKey, environmentKey, experimentKey).ExperimentPatchInput(experimentPatchInput).Execute()

Patch experiment



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
    experimentKey := "experimentKey_example" // string | The experiment key
    experimentPatchInput := *openapiclient.NewExperimentPatchInput([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // ExperimentPatchInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.PatchExperiment(context.Background(), projectKey, environmentKey, experimentKey).ExperimentPatchInput(experimentPatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.PatchExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchExperiment`: Experiment
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.PatchExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**experimentKey** | **string** | The experiment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **experimentPatchInput** | [**ExperimentPatchInput**](ExperimentPatchInput.md) |  | 

### Return type

[**Experiment**](Experiment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutExperimentationSettings

> ExperimentationSettingsRep PutExperimentationSettings(ctx, projectKey).ExperimentationSettingsPut(experimentationSettingsPut).Execute()

Update experimentation settings



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
    experimentationSettingsPut := *openapiclient.NewExperimentationSettingsPut([]openapiclient.RandomizationUnitInput{*openapiclient.NewRandomizationUnitInput("user", true, "StandardRandomizationUnit_example")}) // ExperimentationSettingsPut | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.PutExperimentationSettings(context.Background(), projectKey).ExperimentationSettingsPut(experimentationSettingsPut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.PutExperimentationSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutExperimentationSettings`: ExperimentationSettingsRep
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.PutExperimentationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExperimentationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **experimentationSettingsPut** | [**ExperimentationSettingsPut**](ExperimentationSettingsPut.md) |  | 

### Return type

[**ExperimentationSettingsRep**](ExperimentationSettingsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetExperiment

> ResetExperiment(ctx, projectKey, featureFlagKey, environmentKey, metricKey).Execute()

Reset experiment results



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    environmentKey := "environmentKey_example" // string | The environment key
    metricKey := "metricKey_example" // string | The metric's key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExperimentsBetaApi.ResetExperiment(context.Background(), projectKey, featureFlagKey, environmentKey, metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.ResetExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**metricKey** | **string** | The metric&#39;s key | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetExperimentRequest struct via the builder pattern


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

