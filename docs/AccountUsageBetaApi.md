# \AccountUsageBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCCMUsage**](AccountUsageBetaApi.md#GetCCMUsage) | **Get** /api/v2/usage/ccm | Get client connection minutes usage
[**GetCMCIUsage**](AccountUsageBetaApi.md#GetCMCIUsage) | **Get** /api/v2/usage/cmci | Get client-side monthly context instances usage
[**GetDataExportEventsUsage**](AccountUsageBetaApi.md#GetDataExportEventsUsage) | **Get** /api/v2/usage/data-export-events | Get data export events usage
[**GetEvaluationsUsage**](AccountUsageBetaApi.md#GetEvaluationsUsage) | **Get** /api/v2/usage/evaluations/{projectKey}/{environmentKey}/{featureFlagKey} | Get evaluations usage
[**GetEventsUsage**](AccountUsageBetaApi.md#GetEventsUsage) | **Get** /api/v2/usage/events/{type} | Get events usage
[**GetExperimentationKeysUsage**](AccountUsageBetaApi.md#GetExperimentationKeysUsage) | **Get** /api/v2/usage/experimentation-keys | Get experimentation keys usage
[**GetExperimentationUnitsUsage**](AccountUsageBetaApi.md#GetExperimentationUnitsUsage) | **Get** /api/v2/usage/experimentation-units | Get experimentation units usage
[**GetMauSdksByType**](AccountUsageBetaApi.md#GetMauSdksByType) | **Get** /api/v2/usage/mau/sdks | Get MAU SDKs by type
[**GetMauUsage**](AccountUsageBetaApi.md#GetMauUsage) | **Get** /api/v2/usage/mau | Get MAU usage
[**GetMauUsageByCategory**](AccountUsageBetaApi.md#GetMauUsageByCategory) | **Get** /api/v2/usage/mau/bycategory | Get MAU usage by category
[**GetSCMUsage**](AccountUsageBetaApi.md#GetSCMUsage) | **Get** /api/v2/usage/scm | Get server connection minutes usage
[**GetServiceConnectionUsage**](AccountUsageBetaApi.md#GetServiceConnectionUsage) | **Get** /api/v2/usage/service-connections | Get service connection usage
[**GetStreamUsage**](AccountUsageBetaApi.md#GetStreamUsage) | **Get** /api/v2/usage/streams/{source} | Get stream usage
[**GetStreamUsageBySdkVersion**](AccountUsageBetaApi.md#GetStreamUsageBySdkVersion) | **Get** /api/v2/usage/streams/{source}/bysdkversion | Get stream usage by SDK version
[**GetStreamUsageSdkversion**](AccountUsageBetaApi.md#GetStreamUsageSdkversion) | **Get** /api/v2/usage/streams/{source}/sdkversions | Get stream usage SDK versions



## GetCCMUsage

> SeriesIntervalsRep GetCCMUsage(ctx).From(from).To(to).Execute()

Get client connection minutes usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetCCMUsage(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetCCMUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCCMUsage`: SeriesIntervalsRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetCCMUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCCMUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 

### Return type

[**SeriesIntervalsRep**](SeriesIntervalsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCMCIUsage

> SeriesIntervalsRep GetCMCIUsage(ctx).From(from).To(to).Execute()

Get client-side monthly context instances usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetCMCIUsage(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetCMCIUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCMCIUsage`: SeriesIntervalsRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetCMCIUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCMCIUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 

### Return type

[**SeriesIntervalsRep**](SeriesIntervalsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataExportEventsUsage

> SeriesIntervalsRep GetDataExportEventsUsage(ctx).From(from).To(to).Execute()

Get data export events usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetDataExportEventsUsage(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetDataExportEventsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataExportEventsUsage`: SeriesIntervalsRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetDataExportEventsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataExportEventsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 

### Return type

[**SeriesIntervalsRep**](SeriesIntervalsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvaluationsUsage

> SeriesListRep GetEvaluationsUsage(ctx, projectKey, environmentKey, featureFlagKey).From(from).To(to).Tz(tz).Execute()

Get evaluations usage



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
    tz := "tz_example" // string | The timezone to use for breaks between days when returning daily data. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetEvaluationsUsage(context.Background(), projectKey, environmentKey, featureFlagKey).From(from).To(to).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetEvaluationsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvaluationsUsage`: SeriesListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetEvaluationsUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvaluationsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **tz** | **string** | The timezone to use for breaks between days when returning daily data. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsUsage

> SeriesListRep GetEventsUsage(ctx, type_).From(from).To(to).Execute()

Get events usage



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
    type_ := "type__example" // string | The type of event to retrieve. Must be either `received` or `published`.
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 24 hours ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetEventsUsage(context.Background(), type_).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetEventsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsUsage`: SeriesListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetEventsUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of event to retrieve. Must be either &#x60;received&#x60; or &#x60;published&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 24 hours ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentationKeysUsage

> SeriesIntervalsRep GetExperimentationKeysUsage(ctx).From(from).To(to).Execute()

Get experimentation keys usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetExperimentationKeysUsage(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetExperimentationKeysUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperimentationKeysUsage`: SeriesIntervalsRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetExperimentationKeysUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentationKeysUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 

### Return type

[**SeriesIntervalsRep**](SeriesIntervalsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentationUnitsUsage

> SeriesIntervalsRep GetExperimentationUnitsUsage(ctx).From(from).To(to).Execute()

Get experimentation units usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetExperimentationUnitsUsage(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetExperimentationUnitsUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExperimentationUnitsUsage`: SeriesIntervalsRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetExperimentationUnitsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentationUnitsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 

### Return type

[**SeriesIntervalsRep**](SeriesIntervalsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMauSdksByType

> SdkListRep GetMauSdksByType(ctx).From(from).To(to).Sdktype(sdktype).Execute()

Get MAU SDKs by type



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
    from := "from_example" // string | The data returned starts from this timestamp. Defaults to seven days ago. The timestamp is in Unix milliseconds, for example, 1656694800000. (optional)
    to := "to_example" // string | The data returned ends at this timestamp. Defaults to the current time. The timestamp is in Unix milliseconds, for example, 1657904400000. (optional)
    sdktype := "sdktype_example" // string | The type of SDK with monthly active users (MAU) to list. Must be either `client` or `server`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetMauSdksByType(context.Background()).From(from).To(to).Sdktype(sdktype).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMauSdksByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMauSdksByType`: SdkListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMauSdksByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMauSdksByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The data returned starts from this timestamp. Defaults to seven days ago. The timestamp is in Unix milliseconds, for example, 1656694800000. | 
 **to** | **string** | The data returned ends at this timestamp. Defaults to the current time. The timestamp is in Unix milliseconds, for example, 1657904400000. | 
 **sdktype** | **string** | The type of SDK with monthly active users (MAU) to list. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Return type

[**SdkListRep**](SdkListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMauUsage

> SeriesListRep GetMauUsage(ctx).From(from).To(to).Project(project).Environment(environment).Sdktype(sdktype).Sdk(sdk).Anonymous(anonymous).Groupby(groupby).Execute()

Get MAU usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
    project := "project_example" // string | A project key to filter results to. Can be specified multiple times, one query parameter per project key, to view data for multiple projects. (optional)
    environment := "environment_example" // string | An environment key to filter results to. When using this parameter, exactly one project key must also be set. Can be specified multiple times as separate query parameters to view data for multiple environments within a single project. (optional)
    sdktype := "sdktype_example" // string | An SDK type to filter results to. Can be specified multiple times, one query parameter per SDK type. Valid values: client, server (optional)
    sdk := "sdk_example" // string | An SDK name to filter results to. Can be specified multiple times, one query parameter per SDK. (optional)
    anonymous := "anonymous_example" // string | If specified, filters results to either anonymous or nonanonymous users. (optional)
    groupby := "groupby_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions (for example, to group by both project and SDK). Valid values: project, environment, sdktype, sdk, anonymous (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetMauUsage(context.Background()).From(from).To(to).Project(project).Environment(environment).Sdktype(sdktype).Sdk(sdk).Anonymous(anonymous).Groupby(groupby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMauUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMauUsage`: SeriesListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMauUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMauUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **project** | **string** | A project key to filter results to. Can be specified multiple times, one query parameter per project key, to view data for multiple projects. | 
 **environment** | **string** | An environment key to filter results to. When using this parameter, exactly one project key must also be set. Can be specified multiple times as separate query parameters to view data for multiple environments within a single project. | 
 **sdktype** | **string** | An SDK type to filter results to. Can be specified multiple times, one query parameter per SDK type. Valid values: client, server | 
 **sdk** | **string** | An SDK name to filter results to. Can be specified multiple times, one query parameter per SDK. | 
 **anonymous** | **string** | If specified, filters results to either anonymous or nonanonymous users. | 
 **groupby** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions (for example, to group by both project and SDK). Valid values: project, environment, sdktype, sdk, anonymous | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMauUsageByCategory

> SeriesListRep GetMauUsageByCategory(ctx).From(from).To(to).Execute()

Get MAU usage by category



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
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetMauUsageByCategory(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMauUsageByCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMauUsageByCategory`: SeriesListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMauUsageByCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMauUsageByCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSCMUsage

> SeriesIntervalsRep GetSCMUsage(ctx).From(from).To(to).Execute()

Get server connection minutes usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetSCMUsage(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetSCMUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSCMUsage`: SeriesIntervalsRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetSCMUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSCMUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 

### Return type

[**SeriesIntervalsRep**](SeriesIntervalsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConnectionUsage

> SeriesIntervalsRep GetServiceConnectionUsage(ctx).From(from).To(to).Execute()

Get service connection usage



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
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetServiceConnectionUsage(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetServiceConnectionUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceConnectionUsage`: SeriesIntervalsRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetServiceConnectionUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceConnectionUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 

### Return type

[**SeriesIntervalsRep**](SeriesIntervalsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamUsage

> SeriesListRep GetStreamUsage(ctx, source).From(from).To(to).Tz(tz).Execute()

Get stream usage



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
    source := "source_example" // string | The source of streaming connections to describe. Must be either `client` or `server`.
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
    tz := "tz_example" // string | The timezone to use for breaks between days when returning daily data. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetStreamUsage(context.Background(), source).From(from).To(to).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetStreamUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamUsage`: SeriesListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetStreamUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | The source of streaming connections to describe. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **tz** | **string** | The timezone to use for breaks between days when returning daily data. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamUsageBySdkVersion

> SeriesListRep GetStreamUsageBySdkVersion(ctx, source).From(from).To(to).Tz(tz).Sdk(sdk).Version(version).Execute()

Get stream usage by SDK version



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
    source := "source_example" // string | The source of streaming connections to describe. Must be either `client` or `server`.
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 24 hours ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
    tz := "tz_example" // string | The timezone to use for breaks between days when returning daily data. (optional)
    sdk := "sdk_example" // string | If included, this filters the returned series to only those that match this SDK name. (optional)
    version := "version_example" // string | If included, this filters the returned series to only those that match this SDK version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetStreamUsageBySdkVersion(context.Background(), source).From(from).To(to).Tz(tz).Sdk(sdk).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetStreamUsageBySdkVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamUsageBySdkVersion`: SeriesListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetStreamUsageBySdkVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | The source of streaming connections to describe. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamUsageBySdkVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 24 hours ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **tz** | **string** | The timezone to use for breaks between days when returning daily data. | 
 **sdk** | **string** | If included, this filters the returned series to only those that match this SDK name. | 
 **version** | **string** | If included, this filters the returned series to only those that match this SDK version. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamUsageSdkversion

> SdkVersionListRep GetStreamUsageSdkversion(ctx, source).Execute()

Get stream usage SDK versions



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
    source := "source_example" // string | The source of streaming connections to describe. Must be either `client` or `server`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountUsageBetaApi.GetStreamUsageSdkversion(context.Background(), source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetStreamUsageSdkversion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamUsageSdkversion`: SdkVersionListRep
    fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetStreamUsageSdkversion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | The source of streaming connections to describe. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamUsageSdkversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SdkVersionListRep**](SdkVersionListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

