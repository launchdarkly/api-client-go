# \InsightsChartsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeploymentFrequencyChart**](InsightsChartsBetaApi.md#GetDeploymentFrequencyChart) | **Get** /api/v2/engineering-insights/charts/deployments/frequency | Get deployment frequency chart data
[**GetFlagStatusChart**](InsightsChartsBetaApi.md#GetFlagStatusChart) | **Get** /api/v2/engineering-insights/charts/flags/status | Get flag status chart data
[**GetLeadTimeChart**](InsightsChartsBetaApi.md#GetLeadTimeChart) | **Get** /api/v2/engineering-insights/charts/lead-time | Get lead time chart data
[**GetReleaseFrequencyChart**](InsightsChartsBetaApi.md#GetReleaseFrequencyChart) | **Get** /api/v2/engineering-insights/charts/releases/frequency | Get replease frequency chart data
[**GetStaleFlagsChart**](InsightsChartsBetaApi.md#GetStaleFlagsChart) | **Get** /api/v2/engineering-insights/charts/flags/stale | Get stale flags chart data



## GetDeploymentFrequencyChart

> InsightsChart GetDeploymentFrequencyChart(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).From(from).To(to).BucketType(bucketType).BucketMs(bucketMs).GroupBy(groupBy).Expand(expand).Execute()

Get deployment frequency chart data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | The project key (optional)
    environmentKey := "environmentKey_example" // string | The environment key (optional)
    applicationKey := "applicationKey_example" // string | Comma separated list of application keys (optional)
    from := time.Now() // time.Time | Unix timestamp in milliseconds. Default value is 7 days ago. (optional)
    to := time.Now() // time.Time | Unix timestamp in milliseconds. Default value is now. (optional)
    bucketType := "bucketType_example" // string | Specify type of bucket. Options: `rolling`, `hour`, `day`. Default: `rolling`. (optional)
    bucketMs := int64(789) // int64 | Duration of intervals for x-axis in milliseconds. Default value is one day (`86400000` milliseconds). (optional)
    groupBy := "groupBy_example" // string | Options: `application`, `kind` (optional)
    expand := "expand_example" // string | Options: `metrics` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsChartsBetaApi.GetDeploymentFrequencyChart(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).From(from).To(to).BucketType(bucketType).BucketMs(bucketMs).GroupBy(groupBy).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsChartsBetaApi.GetDeploymentFrequencyChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeploymentFrequencyChart`: InsightsChart
    fmt.Fprintf(os.Stdout, "Response from `InsightsChartsBetaApi.GetDeploymentFrequencyChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentFrequencyChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 
 **from** | **time.Time** | Unix timestamp in milliseconds. Default value is 7 days ago. | 
 **to** | **time.Time** | Unix timestamp in milliseconds. Default value is now. | 
 **bucketType** | **string** | Specify type of bucket. Options: &#x60;rolling&#x60;, &#x60;hour&#x60;, &#x60;day&#x60;. Default: &#x60;rolling&#x60;. | 
 **bucketMs** | **int64** | Duration of intervals for x-axis in milliseconds. Default value is one day (&#x60;86400000&#x60; milliseconds). | 
 **groupBy** | **string** | Options: &#x60;application&#x60;, &#x60;kind&#x60; | 
 **expand** | **string** | Options: &#x60;metrics&#x60; | 

### Return type

[**InsightsChart**](InsightsChart.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagStatusChart

> InsightsChart GetFlagStatusChart(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Execute()

Get flag status chart data



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
    resp, r, err := apiClient.InsightsChartsBetaApi.GetFlagStatusChart(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsChartsBetaApi.GetFlagStatusChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagStatusChart`: InsightsChart
    fmt.Fprintf(os.Stdout, "Response from `InsightsChartsBetaApi.GetFlagStatusChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagStatusChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 

### Return type

[**InsightsChart**](InsightsChart.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeadTimeChart

> InsightsChart GetLeadTimeChart(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).From(from).To(to).BucketType(bucketType).BucketMs(bucketMs).GroupBy(groupBy).Expand(expand).Execute()

Get lead time chart data



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
    environmentKey := "environmentKey_example" // string | The environment key (optional)
    applicationKey := "applicationKey_example" // string | Comma separated list of application keys (optional)
    from := int64(789) // int64 | Unix timestamp in milliseconds. Default value is 7 days ago. (optional)
    to := int64(789) // int64 | Unix timestamp in milliseconds. Default value is now. (optional)
    bucketType := "bucketType_example" // string | Specify type of bucket. Options: `rolling`, `hour`, `day`. Default: `rolling`. (optional)
    bucketMs := int64(789) // int64 | Duration of intervals for x-axis in milliseconds. Default value is one day (`86400000` milliseconds). (optional)
    groupBy := "groupBy_example" // string | Options: `application`, `stage`. Default: `stage`. (optional)
    expand := "expand_example" // string | Options: `metrics`, `percentiles`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsChartsBetaApi.GetLeadTimeChart(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).From(from).To(to).BucketType(bucketType).BucketMs(bucketMs).GroupBy(groupBy).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsChartsBetaApi.GetLeadTimeChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLeadTimeChart`: InsightsChart
    fmt.Fprintf(os.Stdout, "Response from `InsightsChartsBetaApi.GetLeadTimeChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeadTimeChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 
 **from** | **int64** | Unix timestamp in milliseconds. Default value is 7 days ago. | 
 **to** | **int64** | Unix timestamp in milliseconds. Default value is now. | 
 **bucketType** | **string** | Specify type of bucket. Options: &#x60;rolling&#x60;, &#x60;hour&#x60;, &#x60;day&#x60;. Default: &#x60;rolling&#x60;. | 
 **bucketMs** | **int64** | Duration of intervals for x-axis in milliseconds. Default value is one day (&#x60;86400000&#x60; milliseconds). | 
 **groupBy** | **string** | Options: &#x60;application&#x60;, &#x60;stage&#x60;. Default: &#x60;stage&#x60;. | 
 **expand** | **string** | Options: &#x60;metrics&#x60;, &#x60;percentiles&#x60;. | 

### Return type

[**InsightsChart**](InsightsChart.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseFrequencyChart

> InsightsChart GetReleaseFrequencyChart(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).HasExperiments(hasExperiments).Global(global).GroupBy(groupBy).From(from).To(to).BucketType(bucketType).BucketMs(bucketMs).Expand(expand).Execute()

Get replease frequency chart data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | The project key
    environmentKey := "environmentKey_example" // string | The environment key
    applicationKey := "applicationKey_example" // string | Comma separated list of application keys (optional)
    hasExperiments := true // bool | Filter events to those associated with an experiment (optional)
    global := "global_example" // string | Filter to include or exclude global events. Default value is `include`. Options: `include`, `exclude` (optional)
    groupBy := "groupBy_example" // string | Property to group results by. Options: `impact` (optional)
    from := time.Now() // time.Time | Unix timestamp in milliseconds. Default value is 7 days ago. (optional)
    to := time.Now() // time.Time | Unix timestamp in milliseconds. Default value is now. (optional)
    bucketType := "bucketType_example" // string | Specify type of bucket. Options: `rolling`, `hour`, `day`. Default: `rolling`. (optional)
    bucketMs := int64(789) // int64 | Duration of intervals for x-axis in milliseconds. Default value is one day (`86400000` milliseconds). (optional)
    expand := "expand_example" // string | Options: `metrics` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsChartsBetaApi.GetReleaseFrequencyChart(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).HasExperiments(hasExperiments).Global(global).GroupBy(groupBy).From(from).To(to).BucketType(bucketType).BucketMs(bucketMs).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsChartsBetaApi.GetReleaseFrequencyChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleaseFrequencyChart`: InsightsChart
    fmt.Fprintf(os.Stdout, "Response from `InsightsChartsBetaApi.GetReleaseFrequencyChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseFrequencyChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 
 **hasExperiments** | **bool** | Filter events to those associated with an experiment | 
 **global** | **string** | Filter to include or exclude global events. Default value is &#x60;include&#x60;. Options: &#x60;include&#x60;, &#x60;exclude&#x60; | 
 **groupBy** | **string** | Property to group results by. Options: &#x60;impact&#x60; | 
 **from** | **time.Time** | Unix timestamp in milliseconds. Default value is 7 days ago. | 
 **to** | **time.Time** | Unix timestamp in milliseconds. Default value is now. | 
 **bucketType** | **string** | Specify type of bucket. Options: &#x60;rolling&#x60;, &#x60;hour&#x60;, &#x60;day&#x60;. Default: &#x60;rolling&#x60;. | 
 **bucketMs** | **int64** | Duration of intervals for x-axis in milliseconds. Default value is one day (&#x60;86400000&#x60; milliseconds). | 
 **expand** | **string** | Options: &#x60;metrics&#x60; | 

### Return type

[**InsightsChart**](InsightsChart.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaleFlagsChart

> InsightsChart GetStaleFlagsChart(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).GroupBy(groupBy).MaintainerId(maintainerId).MaintainerTeamKey(maintainerTeamKey).Expand(expand).Execute()

Get stale flags chart data



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
    groupBy := "groupBy_example" // string | Property to group results by. Options: `maintainer` (optional)
    maintainerId := "maintainerId_example" // string | Comma-separated list of individual maintainers to filter results. (optional)
    maintainerTeamKey := "maintainerTeamKey_example" // string | Comma-separated list of team maintainer keys to filter results. (optional)
    expand := "expand_example" // string | Options: `metrics` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsChartsBetaApi.GetStaleFlagsChart(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).GroupBy(groupBy).MaintainerId(maintainerId).MaintainerTeamKey(maintainerTeamKey).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsChartsBetaApi.GetStaleFlagsChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStaleFlagsChart`: InsightsChart
    fmt.Fprintf(os.Stdout, "Response from `InsightsChartsBetaApi.GetStaleFlagsChart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStaleFlagsChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 
 **groupBy** | **string** | Property to group results by. Options: &#x60;maintainer&#x60; | 
 **maintainerId** | **string** | Comma-separated list of individual maintainers to filter results. | 
 **maintainerTeamKey** | **string** | Comma-separated list of team maintainer keys to filter results. | 
 **expand** | **string** | Options: &#x60;metrics&#x60; | 

### Return type

[**InsightsChart**](InsightsChart.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

