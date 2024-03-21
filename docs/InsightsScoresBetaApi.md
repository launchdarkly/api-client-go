# \InsightsScoresBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInsightGroup**](InsightsScoresBetaApi.md#CreateInsightGroup) | **Post** /api/v2/engineering-insights/insights/group | Create insight group
[**DeleteInsightGroup**](InsightsScoresBetaApi.md#DeleteInsightGroup) | **Delete** /api/v2/engineering-insights/insights/groups/{insightGroupKey} | Delete insight group
[**GetInsightGroup**](InsightsScoresBetaApi.md#GetInsightGroup) | **Get** /api/v2/engineering-insights/insights/groups/{insightGroupKey} | Get insight group
[**GetInsightGroups**](InsightsScoresBetaApi.md#GetInsightGroups) | **Get** /api/v2/engineering-insights/insights/groups | List insight groups
[**GetInsightsScores**](InsightsScoresBetaApi.md#GetInsightsScores) | **Get** /api/v2/engineering-insights/insights/scores | Get insight scores
[**PatchInsightGroup**](InsightsScoresBetaApi.md#PatchInsightGroup) | **Patch** /api/v2/engineering-insights/insights/groups/{insightGroupKey} | Patch insight group



## CreateInsightGroup

> InsightGroup CreateInsightGroup(ctx).PostInsightGroupParams(postInsightGroupParams).Execute()

Create insight group



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
    postInsightGroupParams := *openapiclient.NewPostInsightGroupParams("Production - All Apps", "default-production-all-apps", "default", "production") // PostInsightGroupParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsScoresBetaApi.CreateInsightGroup(context.Background()).PostInsightGroupParams(postInsightGroupParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsScoresBetaApi.CreateInsightGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInsightGroup`: InsightGroup
    fmt.Fprintf(os.Stdout, "Response from `InsightsScoresBetaApi.CreateInsightGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInsightGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postInsightGroupParams** | [**PostInsightGroupParams**](PostInsightGroupParams.md) |  | 

### Return type

[**InsightGroup**](InsightGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInsightGroup

> DeleteInsightGroup(ctx, insightGroupKey).Execute()

Delete insight group



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
    insightGroupKey := "insightGroupKey_example" // string | The insight group key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsScoresBetaApi.DeleteInsightGroup(context.Background(), insightGroupKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsScoresBetaApi.DeleteInsightGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightGroupKey** | **string** | The insight group key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInsightGroupRequest struct via the builder pattern


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


## GetInsightGroup

> InsightGroup GetInsightGroup(ctx, insightGroupKey).Expand(expand).Execute()

Get insight group



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
    insightGroupKey := "insightGroupKey_example" // string | The insight group key
    expand := "expand_example" // string | Options: `scores`, `environment` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsScoresBetaApi.GetInsightGroup(context.Background(), insightGroupKey).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsScoresBetaApi.GetInsightGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightGroup`: InsightGroup
    fmt.Fprintf(os.Stdout, "Response from `InsightsScoresBetaApi.GetInsightGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightGroupKey** | **string** | The insight group key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Options: &#x60;scores&#x60;, &#x60;environment&#x60; | 

### Return type

[**InsightGroup**](InsightGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightGroups

> InsightGroupCollection GetInsightGroups(ctx).Limit(limit).Offset(offset).Sort(sort).Query(query).Expand(expand).Execute()

List insight groups



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
    limit := int64(789) // int64 | The number of insight groups to return. Default is 20. Must be between 1 and 20 inclusive. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
    sort := "sort_example" // string | Sort flag list by field. Prefix field with <code>-</code> to sort in descending order. Allowed fields: name (optional)
    query := "query_example" // string | Filter list of insights groups by name. (optional)
    expand := "expand_example" // string | Options: `scores`, `environment`, `metadata` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsScoresBetaApi.GetInsightGroups(context.Background()).Limit(limit).Offset(offset).Sort(sort).Query(query).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsScoresBetaApi.GetInsightGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightGroups`: InsightGroupCollection
    fmt.Fprintf(os.Stdout, "Response from `InsightsScoresBetaApi.GetInsightGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The number of insight groups to return. Default is 20. Must be between 1 and 20 inclusive. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **sort** | **string** | Sort flag list by field. Prefix field with &lt;code&gt;-&lt;/code&gt; to sort in descending order. Allowed fields: name | 
 **query** | **string** | Filter list of insights groups by name. | 
 **expand** | **string** | Options: &#x60;scores&#x60;, &#x60;environment&#x60;, &#x60;metadata&#x60; | 

### Return type

[**InsightGroupCollection**](InsightGroupCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsightsScores

> InsightScores GetInsightsScores(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Execute()

Get insight scores



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
    applicationKey := "applicationKey_example" // string | Comma separated list of application keys (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsScoresBetaApi.GetInsightsScores(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsScoresBetaApi.GetInsightsScores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightsScores`: InsightScores
    fmt.Fprintf(os.Stdout, "Response from `InsightsScoresBetaApi.GetInsightsScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightsScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 

### Return type

[**InsightScores**](InsightScores.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchInsightGroup

> InsightGroup PatchInsightGroup(ctx, insightGroupKey).PatchOperation(patchOperation).Execute()

Patch insight group



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
    insightGroupKey := "insightGroupKey_example" // string | The insight group key
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsScoresBetaApi.PatchInsightGroup(context.Background(), insightGroupKey).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsScoresBetaApi.PatchInsightGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchInsightGroup`: InsightGroup
    fmt.Fprintf(os.Stdout, "Response from `InsightsScoresBetaApi.PatchInsightGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightGroupKey** | **string** | The insight group key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchInsightGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**InsightGroup**](InsightGroup.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

