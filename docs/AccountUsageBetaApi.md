# \AccountUsageBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvaluationsUsage**](AccountUsageBetaApi.md#GetEvaluationsUsage) | **Get** /api/v2/usage/evaluations/{projKey}/{envKey}/{flagKey} | Get evaluations usage
[**GetEventsUsage**](AccountUsageBetaApi.md#GetEventsUsage) | **Get** /api/v2/usage/events/{type} | Get events usage
[**GetMauSdksByType**](AccountUsageBetaApi.md#GetMauSdksByType) | **Get** /api/v2/usage/mau/sdks | Get MAU SDKs by type
[**GetMauUsage**](AccountUsageBetaApi.md#GetMauUsage) | **Get** /api/v2/usage/mau | Get MAU usage
[**GetMauUsageByCategory**](AccountUsageBetaApi.md#GetMauUsageByCategory) | **Get** /api/v2/usage/mau/bycategory | Get MAU usage by category
[**GetStreamUsage**](AccountUsageBetaApi.md#GetStreamUsage) | **Get** /api/v2/usage/streams/{source} | Get stream usage
[**GetStreamUsageBySdkVersion**](AccountUsageBetaApi.md#GetStreamUsageBySdkVersion) | **Get** /api/v2/usage/streams/{source}/bysdkversion | Get stream usage by SDK version
[**GetStreamUsageSdkversion**](AccountUsageBetaApi.md#GetStreamUsageSdkversion) | **Get** /api/v2/usage/streams/{source}/sdkversions | Get stream usage SDK versions



## GetEvaluationsUsage

> SeriesListRep GetEvaluationsUsage(ctx, projKey, envKey, flagKey).From(from).To(to).Tz(tz).Execute()

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
    projKey := "projKey_example" // string | The project key.
    envKey := "envKey_example" // string | The environment key.
    flagKey := "flagKey_example" // string | The feature flag's key.
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
    tz := "tz_example" // string | The timezone to use for breaks between days when returning daily data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetEvaluationsUsage(context.Background(), projKey, envKey, flagKey).From(from).To(to).Tz(tz).Execute()
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
**projKey** | **string** | The project key. | 
**envKey** | **string** | The environment key. | 
**flagKey** | **string** | The feature flag&#39;s key. | 

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetEventsUsage(context.Background(), type_).From(from).To(to).Execute()
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
    from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to seven days ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
    sdktype := "sdktype_example" // string | The type of SDK with monthly active users (MAU) to list. Must be either `client` or `server` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetMauSdksByType(context.Background()).From(from).To(to).Sdktype(sdktype).Execute()
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
 **from** | **string** | The series of data returned starts from this timestamp. Defaults to seven days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **sdktype** | **string** | The type of SDK with monthly active users (MAU) to list. Must be either &#x60;client&#x60; or &#x60;server&#x60; | 

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

> SeriesListRep GetMauUsage(ctx).From(from).To(to).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetMauUsage(context.Background()).From(from).To(to).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetMauUsageByCategory(context.Background()).From(from).To(to).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetStreamUsage(context.Background(), source).From(from).To(to).Tz(tz).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetStreamUsageBySdkVersion(context.Background(), source).From(from).To(to).Tz(tz).Sdk(sdk).Version(version).Execute()
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUsageBetaApi.GetStreamUsageSdkversion(context.Background(), source).Execute()
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

