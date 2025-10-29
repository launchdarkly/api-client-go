# \MetricsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricGroup**](MetricsBetaApi.md#CreateMetricGroup) | **Post** /api/v2/projects/{projectKey}/metric-groups | Create metric group
[**DeleteMetricGroup**](MetricsBetaApi.md#DeleteMetricGroup) | **Delete** /api/v2/projects/{projectKey}/metric-groups/{metricGroupKey} | Delete metric group
[**GetMetricGroup**](MetricsBetaApi.md#GetMetricGroup) | **Get** /api/v2/projects/{projectKey}/metric-groups/{metricGroupKey} | Get metric group
[**GetMetricGroups**](MetricsBetaApi.md#GetMetricGroups) | **Get** /api/v2/projects/{projectKey}/metric-groups | List metric groups
[**PatchMetricGroup**](MetricsBetaApi.md#PatchMetricGroup) | **Patch** /api/v2/projects/{projectKey}/metric-groups/{metricGroupKey} | Patch metric group



## CreateMetricGroup

> MetricGroupRep CreateMetricGroup(ctx, projectKey).MetricGroupPost(metricGroupPost).Execute()

Create metric group



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
	metricGroupPost := *openapiclient.NewMetricGroupPost("My metric group", "funnel", "569fdeadbeef1644facecafe", []string{"Tags_example"}, []openapiclient.MetricInMetricGroupInput{*openapiclient.NewMetricInMetricGroupInput("metric-key-123abc", "Step 1")}) // MetricGroupPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsBetaApi.CreateMetricGroup(context.Background(), projectKey).MetricGroupPost(metricGroupPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsBetaApi.CreateMetricGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMetricGroup`: MetricGroupRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsBetaApi.CreateMetricGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricGroupPost** | [**MetricGroupPost**](MetricGroupPost.md) |  | 

### Return type

[**MetricGroupRep**](MetricGroupRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetricGroup

> DeleteMetricGroup(ctx, projectKey, metricGroupKey).Execute()

Delete metric group



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
	metricGroupKey := "metricGroupKey_example" // string | The metric group key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetricsBetaApi.DeleteMetricGroup(context.Background(), projectKey, metricGroupKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsBetaApi.DeleteMetricGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**metricGroupKey** | **string** | The metric group key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricGroupRequest struct via the builder pattern


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


## GetMetricGroup

> MetricGroupRep GetMetricGroup(ctx, projectKey, metricGroupKey).Expand(expand).Execute()

Get metric group



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
	metricGroupKey := "metricGroupKey_example" // string | The metric group key
	expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsBetaApi.GetMetricGroup(context.Background(), projectKey, metricGroupKey).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsBetaApi.GetMetricGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricGroup`: MetricGroupRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsBetaApi.GetMetricGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**metricGroupKey** | **string** | The metric group key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. | 

### Return type

[**MetricGroupRep**](MetricGroupRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricGroups

> MetricGroupCollectionRep GetMetricGroups(ctx, projectKey).Filter(filter).Sort(sort).Expand(expand).Limit(limit).Offset(offset).Execute()

List metric groups



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
	filter := "filter_example" // string | Accepts filter by `experimentStatus`, `query`, `kind`, `hasConnections`, `maintainerIds`, and `maintainerTeamKey`. Example: `filter=experimentStatus equals 'running' and query equals 'test'`. (optional)
	sort := "sort_example" // string | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order. Read the endpoint description for a full list of available sort fields. (optional)
	expand := "expand_example" // string | This parameter is reserved for future use and is not currently supported on this endpoint. (optional)
	limit := int64(789) // int64 | The number of metric groups to return in the response. Defaults to 20. Maximum limit is 50. (optional)
	offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and returns the next `limit` items. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsBetaApi.GetMetricGroups(context.Background(), projectKey).Filter(filter).Sort(sort).Expand(expand).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsBetaApi.GetMetricGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricGroups`: MetricGroupCollectionRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsBetaApi.GetMetricGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Accepts filter by &#x60;experimentStatus&#x60;, &#x60;query&#x60;, &#x60;kind&#x60;, &#x60;hasConnections&#x60;, &#x60;maintainerIds&#x60;, and &#x60;maintainerTeamKey&#x60;. Example: &#x60;filter&#x3D;experimentStatus equals &#39;running&#39; and query equals &#39;test&#39;&#x60;. | 
 **sort** | **string** | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order. Read the endpoint description for a full list of available sort fields. | 
 **expand** | **string** | This parameter is reserved for future use and is not currently supported on this endpoint. | 
 **limit** | **int64** | The number of metric groups to return in the response. Defaults to 20. Maximum limit is 50. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and returns the next &#x60;limit&#x60; items. | 

### Return type

[**MetricGroupCollectionRep**](MetricGroupCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMetricGroup

> MetricGroupRep PatchMetricGroup(ctx, projectKey, metricGroupKey).PatchOperation(patchOperation).Execute()

Patch metric group



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
	metricGroupKey := "metricGroupKey_example" // string | The metric group key
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsBetaApi.PatchMetricGroup(context.Background(), projectKey, metricGroupKey).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsBetaApi.PatchMetricGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMetricGroup`: MetricGroupRep
	fmt.Fprintf(os.Stdout, "Response from `MetricsBetaApi.PatchMetricGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**metricGroupKey** | **string** | The metric group key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMetricGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**MetricGroupRep**](MetricGroupRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

