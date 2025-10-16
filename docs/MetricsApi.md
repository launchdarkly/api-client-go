# \MetricsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMetric**](MetricsApi.md#DeleteMetric) | **Delete** /api/v2/metrics/{projectKey}/{metricKey} | Delete metric
[**GetMetric**](MetricsApi.md#GetMetric) | **Get** /api/v2/metrics/{projectKey}/{metricKey} | Get metric
[**GetMetrics**](MetricsApi.md#GetMetrics) | **Get** /api/v2/metrics/{projectKey} | List metrics
[**PatchMetric**](MetricsApi.md#PatchMetric) | **Patch** /api/v2/metrics/{projectKey}/{metricKey} | Update metric
[**PostMetric**](MetricsApi.md#PostMetric) | **Post** /api/v2/metrics/{projectKey} | Create metric



## DeleteMetric

> DeleteMetric(ctx, projectKey, metricKey).Execute()

Delete metric



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
	projectKey := "projectKey_example" // string | The project key
	metricKey := "metricKey_example" // string | The metric key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetricsApi.DeleteMetric(context.Background(), projectKey, metricKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**metricKey** | **string** | The metric key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricRequest struct via the builder pattern


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


## GetMetric

> MetricRep GetMetric(ctx, projectKey, metricKey).Expand(expand).VersionId(versionId).Execute()

Get metric



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
	projectKey := "projectKey_example" // string | The project key
	metricKey := "metricKey_example" // string | The metric key
	expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. Supported fields are `experiments`, `experimentCount`, `metricGroups`, `metricGroupCount`, `eventSources`, `guardedRollouts`, `guardedRolloutCount`, `lastUsedInExperiment`, and `lastUsedInGuardedRollout`. (optional)
	versionId := "versionId_example" // string | The specific version ID of the metric (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.GetMetric(context.Background(), projectKey, metricKey).Expand(expand).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetric`: MetricRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**metricKey** | **string** | The metric key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. Supported fields are &#x60;experiments&#x60;, &#x60;experimentCount&#x60;, &#x60;metricGroups&#x60;, &#x60;metricGroupCount&#x60;, &#x60;eventSources&#x60;, &#x60;guardedRollouts&#x60;, &#x60;guardedRolloutCount&#x60;, &#x60;lastUsedInExperiment&#x60;, and &#x60;lastUsedInGuardedRollout&#x60;. | 
 **versionId** | **string** | The specific version ID of the metric | 

### Return type

[**MetricRep**](MetricRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetrics

> MetricCollectionRep GetMetrics(ctx, projectKey).Expand(expand).Limit(limit).Offset(offset).Sort(sort).Filter(filter).Execute()

List metrics



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
	projectKey := "projectKey_example" // string | The project key
	expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. (optional)
	limit := int64(789) // int64 | The number of metrics to return in the response. Defaults to 20. Maximum limit is 50. (optional)
	offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and returns the next `limit` items. (optional)
	sort := "sort_example" // string | A field to sort the items by. Prefix field by a dash ( - ) to sort in descending order. This endpoint supports sorting by `createdAt` or `name`. (optional)
	filter := "filter_example" // string | A comma-separated list of filters. This endpoint accepts filtering by `query`, `tags`, 'eventKind', 'isNumeric', 'unitAggregationType`, `hasConnections`, `maintainerIds`, `maintainerTeamKey` and `view`. To learn more about the filter syntax, read the 'Filtering metrics' section above. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.GetMetrics(context.Background(), projectKey).Expand(expand).Limit(limit).Offset(offset).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetrics`: MetricCollectionRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. | 
 **limit** | **int64** | The number of metrics to return in the response. Defaults to 20. Maximum limit is 50. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and returns the next &#x60;limit&#x60; items. | 
 **sort** | **string** | A field to sort the items by. Prefix field by a dash ( - ) to sort in descending order. This endpoint supports sorting by &#x60;createdAt&#x60; or &#x60;name&#x60;. | 
 **filter** | **string** | A comma-separated list of filters. This endpoint accepts filtering by &#x60;query&#x60;, &#x60;tags&#x60;, &#39;eventKind&#39;, &#39;isNumeric&#39;, &#39;unitAggregationType&#x60;, &#x60;hasConnections&#x60;, &#x60;maintainerIds&#x60;, &#x60;maintainerTeamKey&#x60; and &#x60;view&#x60;. To learn more about the filter syntax, read the &#39;Filtering metrics&#39; section above. | 

### Return type

[**MetricCollectionRep**](MetricCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMetric

> MetricRep PatchMetric(ctx, projectKey, metricKey).PatchOperation(patchOperation).Execute()

Update metric



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
	projectKey := "projectKey_example" // string | The project key
	metricKey := "metricKey_example" // string | The metric key
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.PatchMetric(context.Background(), projectKey, metricKey).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.PatchMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMetric`: MetricRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.PatchMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**metricKey** | **string** | The metric key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**MetricRep**](MetricRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMetric

> MetricRep PostMetric(ctx, projectKey).MetricPost(metricPost).Execute()

Create metric



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
	projectKey := "projectKey_example" // string | The project key
	metricPost := *openapiclient.NewMetricPost("metric-key-123abc", "custom") // MetricPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsApi.PostMetric(context.Background(), projectKey).MetricPost(metricPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.PostMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMetric`: MetricRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsApi.PostMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricPost** | [**MetricPost**](MetricPost.md) |  | 

### Return type

[**MetricRep**](MetricRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

