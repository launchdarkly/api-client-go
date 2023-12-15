# \DefaultApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet**](DefaultApi.md#ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/discovered-metric-events | Get Discovered Metric Events
[**DeleteSdkConfigHandler**](DefaultApi.md#DeleteSdkConfigHandler) | **Delete** /private/accounts/{acctId}/projects/{projId}/environments/{envId}/flags/{flagKey}/sdk-config | Delete SDK config
[**GetSdkConfigHandler**](DefaultApi.md#GetSdkConfigHandler) | **Get** /private/accounts/{acctId}/projects/{projId}/environments/{envId}/flags/{flagKey}/sdk-config | Get SDK config
[**PutSdkConfigHandler**](DefaultApi.md#PutSdkConfigHandler) | **Put** /private/accounts/{acctId}/projects/{projId}/environments/{envId}/flags/{flagKey}/sdk-config | Put SDK config



## ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet

> []DiscoveredMetricEventWithMetrics ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet(ctx, projectKey, environmentKey).StartDate(startDate).EndDate(endDate).Execute()

Get Discovered Metric Events



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
    startDate := "startDate_example" // string | Start Date: YYYY-MM-DD (optional)
    endDate := "endDate_example" // string | End Date: YYYY-MM-DD (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet(context.Background(), projectKey, environmentKey).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet`: []DiscoveredMetricEventWithMetrics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsProjectKeyEnvironmentsEnvironmentKeyDiscoveredMetricEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | Start Date: YYYY-MM-DD | 
 **endDate** | **string** | End Date: YYYY-MM-DD | 

### Return type

[**[]DiscoveredMetricEventWithMetrics**](DiscoveredMetricEventWithMetrics.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSdkConfigHandler

> DeleteSdkConfigHandler(ctx, acctId, projId, envId, flagKey).Execute()

Delete SDK config



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
    acctId := "acctId_example" // string | Account id
    projId := "projId_example" // string | Project id
    envId := "envId_example" // string | Environment id
    flagKey := "flagKey_example" // string | Flag key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSdkConfigHandler(context.Background(), acctId, projId, envId, flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSdkConfigHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acctId** | **string** | Account id | 
**projId** | **string** | Project id | 
**envId** | **string** | Environment id | 
**flagKey** | **string** | Flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdkConfigHandlerRequest struct via the builder pattern


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


## GetSdkConfigHandler

> GetSdkConfigHandler(ctx, acctId, projId, envId, flagKey).Execute()

Get SDK config



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
    acctId := "acctId_example" // string | Account id
    projId := "projId_example" // string | Project id
    envId := "envId_example" // string | Environment id
    flagKey := "flagKey_example" // string | Flag key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSdkConfigHandler(context.Background(), acctId, projId, envId, flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSdkConfigHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acctId** | **string** | Account id | 
**projId** | **string** | Project id | 
**envId** | **string** | Environment id | 
**flagKey** | **string** | Flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdkConfigHandlerRequest struct via the builder pattern


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


## PutSdkConfigHandler

> PutSdkConfigHandler(ctx, acctId, projId, envId, flagKey).SdkConfig(sdkConfig).Execute()

Put SDK config



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
    acctId := "acctId_example" // string | Account id
    projId := "projId_example" // string | Project id
    envId := "envId_example" // string | Environment id
    flagKey := "flagKey_example" // string | Flag key
    sdkConfig := *openapiclient.NewSdkConfig() // SdkConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PutSdkConfigHandler(context.Background(), acctId, projId, envId, flagKey).SdkConfig(sdkConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutSdkConfigHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acctId** | **string** | Account id | 
**projId** | **string** | Project id | 
**envId** | **string** | Environment id | 
**flagKey** | **string** | Flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSdkConfigHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sdkConfig** | [**SdkConfig**](SdkConfig.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

