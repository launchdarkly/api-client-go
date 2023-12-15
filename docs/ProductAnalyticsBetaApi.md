# \ProductAnalyticsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMetricsAnalyticsEventsQuery**](ProductAnalyticsBetaApi.md#PostMetricsAnalyticsEventsQuery) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/metrics/query | Query metrics analytics events



## PostMetricsAnalyticsEventsQuery

> ProductAnalyticsQueryResultsRep PostMetricsAnalyticsEventsQuery(ctx, projectKey, environmentKey).PostProductAnalyticsQueryRep(postProductAnalyticsQueryRep).Execute()

Query metrics analytics events



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
    postProductAnalyticsQueryRep := *openapiclient.NewPostProductAnalyticsQueryRep(int64(123), int64(123), []openapiclient.MetricQueryRep{*openapiclient.NewMetricQueryRep("MetricKey_example")}) // PostProductAnalyticsQueryRep | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAnalyticsBetaApi.PostMetricsAnalyticsEventsQuery(context.Background(), projectKey, environmentKey).PostProductAnalyticsQueryRep(postProductAnalyticsQueryRep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAnalyticsBetaApi.PostMetricsAnalyticsEventsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMetricsAnalyticsEventsQuery`: ProductAnalyticsQueryResultsRep
    fmt.Fprintf(os.Stdout, "Response from `ProductAnalyticsBetaApi.PostMetricsAnalyticsEventsQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMetricsAnalyticsEventsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postProductAnalyticsQueryRep** | [**PostProductAnalyticsQueryRep**](PostProductAnalyticsQueryRep.md) |  | 

### Return type

[**ProductAnalyticsQueryResultsRep**](ProductAnalyticsQueryResultsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

