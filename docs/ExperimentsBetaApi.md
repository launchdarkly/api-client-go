# \ExperimentsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExperiment**](ExperimentsBetaApi.md#GetExperiment) | **Get** /api/v2/flags/{projKey}/{flagKey}/experiments/{envKey}/{metricKey} | Get experiment results
[**ResetExperiment**](ExperimentsBetaApi.md#ResetExperiment) | **Delete** /api/v2/flags/{projKey}/{flagKey}/experiments/{envKey}/{metricKey}/results | Reset experiment results



## GetExperiment

> ExperimentResultsRep GetExperiment(ctx, projKey, flagKey, envKey, metricKey).From(from).To(to).Execute()

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
    projKey := "projKey_example" // string | The project key
    flagKey := "flagKey_example" // string | The flag key
    envKey := "envKey_example" // string | The environment key
    metricKey := "metricKey_example" // string | The metric key
    from := int64(789) // int64 | A timestamp denoting the start of the data collection period, expressed as a Unix epoch time in milliseconds. (optional)
    to := int64(789) // int64 | A timestamp denoting the end of the data collection period, expressed as a Unix epoch time in milliseconds. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentsBetaApi.GetExperiment(context.Background(), projKey, flagKey, envKey, metricKey).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentsBetaApi.GetExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperiment`: ExperimentResultsRep
    fmt.Fprintf(os.Stdout, "Response from `ExperimentsBetaApi.GetExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**flagKey** | **string** | The flag key | 
**envKey** | **string** | The environment key | 
**metricKey** | **string** | The metric key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **from** | **int64** | A timestamp denoting the start of the data collection period, expressed as a Unix epoch time in milliseconds. | 
 **to** | **int64** | A timestamp denoting the end of the data collection period, expressed as a Unix epoch time in milliseconds. | 

### Return type

[**ExperimentResultsRep**](ExperimentResultsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetExperiment

> ResetExperiment(ctx, projKey, flagKey, envKey, metricKey).Execute()

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
    projKey := "projKey_example" // string | The project key
    flagKey := "flagKey_example" // string | The feature flag's key
    envKey := "envKey_example" // string | The environment key
    metricKey := "metricKey_example" // string | The metric's key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentsBetaApi.ResetExperiment(context.Background(), projKey, flagKey, envKey, metricKey).Execute()
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
**projKey** | **string** | The project key | 
**flagKey** | **string** | The feature flag&#39;s key | 
**envKey** | **string** | The environment key | 
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

