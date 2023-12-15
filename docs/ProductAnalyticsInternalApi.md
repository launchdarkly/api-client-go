# \ProductAnalyticsInternalApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductAnalyticsContextsDimensionValues**](ProductAnalyticsInternalApi.md#GetProductAnalyticsContextsDimensionValues) | **Get** /internal/projects/{projectKey}/environments/{environmentKey}/analytics/metrics/{metricKey}/dimensions/{dimensionKey}/contexts/{contextKindKey}/values | List of dimension values for a context dimension for product analytics
[**GetProductAnalyticsEventsDimensionValues**](ProductAnalyticsInternalApi.md#GetProductAnalyticsEventsDimensionValues) | **Get** /internal/projects/{projectKey}/environments/{environmentKey}/analytics/metrics/{metricKey}/dimensions/{dimensionKey}/events | List of dimension values for an event dimension for product analytics
[**GetProductAnalyticsMetricDimensions**](ProductAnalyticsInternalApi.md#GetProductAnalyticsMetricDimensions) | **Get** /internal/projects/{projectKey}/environments/{environmentKey}/analytics/metrics/{metricKey}/dimensions | List metric dimensions for product analytics
[**GetProductAnalyticsMetrics**](ProductAnalyticsInternalApi.md#GetProductAnalyticsMetrics) | **Get** /internal/projects/{projectKey}/environments/{environmentKey}/analytics/metrics | List metrics for product analytics



## GetProductAnalyticsContextsDimensionValues

> MetricDimensionValuesRep GetProductAnalyticsContextsDimensionValues(ctx, projectKey, environmentKey, metricKey, dimensionKey, contextKindKey).StartsWith(startsWith).Limit(limit).Execute()

List of dimension values for a context dimension for product analytics



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
    metricKey := "metricKey_example" // string | The metric key
    dimensionKey := "dimensionKey_example" // string | The dimension key
    contextKindKey := "contextKindKey_example" // string | The context kind key
    startsWith := "startsWith_example" // string | A prefix on dimension value (optional)
    limit := int64(789) // int64 | The number of dimension values to return. Defaults to 50. Maximum limit is 100. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAnalyticsInternalApi.GetProductAnalyticsContextsDimensionValues(context.Background(), projectKey, environmentKey, metricKey, dimensionKey, contextKindKey).StartsWith(startsWith).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsInternalApi.GetProductAnalyticsContextsDimensionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductAnalyticsContextsDimensionValues`: MetricDimensionValuesRep
    fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsInternalApi.GetProductAnalyticsContextsDimensionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**metricKey** | **string** | The metric key | 
**dimensionKey** | **string** | The dimension key | 
**contextKindKey** | **string** | The context kind key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductAnalyticsContextsDimensionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **startsWith** | **string** | A prefix on dimension value | 
 **limit** | **int64** | The number of dimension values to return. Defaults to 50. Maximum limit is 100. | 

### Return type

[**MetricDimensionValuesRep**](MetricDimensionValuesRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductAnalyticsEventsDimensionValues

> MetricDimensionValuesRep GetProductAnalyticsEventsDimensionValues(ctx, projectKey, environmentKey, metricKey, dimensionKey).StartsWith(startsWith).Limit(limit).Execute()

List of dimension values for an event dimension for product analytics



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
    metricKey := "metricKey_example" // string | The metric key
    dimensionKey := "dimensionKey_example" // string | The dimension key
    startsWith := "startsWith_example" // string | A prefix on dimension value (optional)
    limit := int64(789) // int64 | The number of dimension values to return. Defaults to 50. Maximum limit is 100. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAnalyticsInternalApi.GetProductAnalyticsEventsDimensionValues(context.Background(), projectKey, environmentKey, metricKey, dimensionKey).StartsWith(startsWith).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsInternalApi.GetProductAnalyticsEventsDimensionValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductAnalyticsEventsDimensionValues`: MetricDimensionValuesRep
    fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsInternalApi.GetProductAnalyticsEventsDimensionValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**metricKey** | **string** | The metric key | 
**dimensionKey** | **string** | The dimension key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductAnalyticsEventsDimensionValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **startsWith** | **string** | A prefix on dimension value | 
 **limit** | **int64** | The number of dimension values to return. Defaults to 50. Maximum limit is 100. | 

### Return type

[**MetricDimensionValuesRep**](MetricDimensionValuesRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductAnalyticsMetricDimensions

> MetricDimensionsRep GetProductAnalyticsMetricDimensions(ctx, projectKey, environmentKey, metricKey).StartsWith(startsWith).Limit(limit).Offset(offset).Execute()

List metric dimensions for product analytics



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
    metricKey := "metricKey_example" // string | The metric key
    startsWith := "startsWith_example" // string | A prefix on metric dimension name (optional)
    limit := int64(789) // int64 | The number of dimensions for a metric to return. Defaults to 50. Maximum limit is 100. (optional)
    offset := int64(789) // int64 | Where to start in the list. For example, an offset of 10 skips the first ten items and then returns the next items in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAnalyticsInternalApi.GetProductAnalyticsMetricDimensions(context.Background(), projectKey, environmentKey, metricKey).StartsWith(startsWith).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsInternalApi.GetProductAnalyticsMetricDimensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductAnalyticsMetricDimensions`: MetricDimensionsRep
    fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsInternalApi.GetProductAnalyticsMetricDimensions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**metricKey** | **string** | The metric key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductAnalyticsMetricDimensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startsWith** | **string** | A prefix on metric dimension name | 
 **limit** | **int64** | The number of dimensions for a metric to return. Defaults to 50. Maximum limit is 100. | 
 **offset** | **int64** | Where to start in the list. For example, an offset of 10 skips the first ten items and then returns the next items in the list. | 

### Return type

[**MetricDimensionsRep**](MetricDimensionsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductAnalyticsMetrics

> ProductAnalyticsMetricsRep GetProductAnalyticsMetrics(ctx, projectKey, environmentKey).StartsWith(startsWith).Limit(limit).Offset(offset).Execute()

List metrics for product analytics



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
    startsWith := "startsWith_example" // string | A prefix on metric name (optional)
    limit := int64(789) // int64 | The number of metrics to return. Defaults to 50. Maximum limit is 100. (optional)
    offset := int64(789) // int64 | Where to start in the list. For example, an offset of 10 skips the first ten items and then returns the next items in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAnalyticsInternalApi.GetProductAnalyticsMetrics(context.Background(), projectKey, environmentKey).StartsWith(startsWith).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsInternalApi.GetProductAnalyticsMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductAnalyticsMetrics`: ProductAnalyticsMetricsRep
    fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsInternalApi.GetProductAnalyticsMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductAnalyticsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startsWith** | **string** | A prefix on metric name | 
 **limit** | **int64** | The number of metrics to return. Defaults to 50. Maximum limit is 100. | 
 **offset** | **int64** | Where to start in the list. For example, an offset of 10 skips the first ten items and then returns the next items in the list. | 

### Return type

[**ProductAnalyticsMetricsRep**](ProductAnalyticsMetricsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

