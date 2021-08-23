# \ScheduledChangesApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFlagConfigScheduledChanges**](ScheduledChangesApi.md#DeleteFlagConfigScheduledChanges) | **Delete** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{id} | Delete scheduled changes workflow
[**GetFeatureFlagScheduledChange**](ScheduledChangesApi.md#GetFeatureFlagScheduledChange) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{id} | Get a scheduled change
[**GetFlagConfigScheduledChanges**](ScheduledChangesApi.md#GetFlagConfigScheduledChanges) | **Get** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | List scheduled changes
[**GetFlagConfigScheduledChangesConflicts**](ScheduledChangesApi.md#GetFlagConfigScheduledChangesConflicts) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes-conflicts | List conflicts for proposed instructions
[**PatchFlagConfigScheduledChange**](ScheduledChangesApi.md#PatchFlagConfigScheduledChange) | **Patch** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{id} | Update scheduled changes workflow
[**PostFlagConfigScheduledChanges**](ScheduledChangesApi.md#PostFlagConfigScheduledChanges) | **Post** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | Create scheduled changes workflow



## DeleteFlagConfigScheduledChanges

> DeleteFlagConfigScheduledChanges(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

Delete scheduled changes workflow



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The scheduled change id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.DeleteFlagConfigScheduledChanges(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.DeleteFlagConfigScheduledChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The scheduled change id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagConfigScheduledChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagScheduledChange

> ScheduledChange GetFeatureFlagScheduledChange(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

Get a scheduled change



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The scheduled change id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.GetFeatureFlagScheduledChange(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.GetFeatureFlagScheduledChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureFlagScheduledChange`: ScheduledChange
    fmt.Fprintf(os.Stdout, "Response from `ScheduledChangesApi.GetFeatureFlagScheduledChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The scheduled change id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagScheduledChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ScheduledChange**](ScheduledChange.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagConfigScheduledChanges

> ScheduledChanges GetFlagConfigScheduledChanges(ctx, projectKey, featureFlagKey, environmentKey).Execute()

List scheduled changes



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.GetFlagConfigScheduledChanges(context.Background(), projectKey, featureFlagKey, environmentKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.GetFlagConfigScheduledChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagConfigScheduledChanges`: ScheduledChanges
    fmt.Fprintf(os.Stdout, "Response from `ScheduledChangesApi.GetFlagConfigScheduledChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagConfigScheduledChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ScheduledChanges**](ScheduledChanges.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagConfigScheduledChangesConflicts

> ConflictResponse GetFlagConfigScheduledChangesConflicts(ctx, projectKey, featureFlagKey, environmentKey).ReportFlagScheduledChangesInput(reportFlagScheduledChangesInput).Execute()

List conflicts for proposed instructions



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    reportFlagScheduledChangesInput := *openapiclient.NewReportFlagScheduledChangesInput(int64(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // ReportFlagScheduledChangesInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.GetFlagConfigScheduledChangesConflicts(context.Background(), projectKey, featureFlagKey, environmentKey).ReportFlagScheduledChangesInput(reportFlagScheduledChangesInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.GetFlagConfigScheduledChangesConflicts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagConfigScheduledChangesConflicts`: ConflictResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduledChangesApi.GetFlagConfigScheduledChangesConflicts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagConfigScheduledChangesConflictsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reportFlagScheduledChangesInput** | [**ReportFlagScheduledChangesInput**](ReportFlagScheduledChangesInput.md) |  | 

### Return type

[**ConflictResponse**](ConflictResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFlagConfigScheduledChange

> ScheduledChange PatchFlagConfigScheduledChange(ctx, projectKey, featureFlagKey, environmentKey, id).FlagScheduledChangesInput(flagScheduledChangesInput).IgnoreConflicts(ignoreConflicts).Execute()

Update scheduled changes workflow



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    id := "id_example" // string | The scheduled change ID
    flagScheduledChangesInput := *openapiclient.NewFlagScheduledChangesInput([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // FlagScheduledChangesInput | 
    ignoreConflicts := true // bool | Whether or not to succeed or fail when the new instructions conflict with existing scheduled changes (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.PatchFlagConfigScheduledChange(context.Background(), projectKey, featureFlagKey, environmentKey, id).FlagScheduledChangesInput(flagScheduledChangesInput).IgnoreConflicts(ignoreConflicts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.PatchFlagConfigScheduledChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFlagConfigScheduledChange`: ScheduledChange
    fmt.Fprintf(os.Stdout, "Response from `ScheduledChangesApi.PatchFlagConfigScheduledChange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The scheduled change ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlagConfigScheduledChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **flagScheduledChangesInput** | [**FlagScheduledChangesInput**](FlagScheduledChangesInput.md) |  | 
 **ignoreConflicts** | **bool** | Whether or not to succeed or fail when the new instructions conflict with existing scheduled changes | 

### Return type

[**ScheduledChange**](ScheduledChange.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlagConfigScheduledChanges

> ScheduledChange PostFlagConfigScheduledChanges(ctx, projectKey, featureFlagKey, environmentKey).PostFlagScheduledChangesInput(postFlagScheduledChangesInput).IgnoreConflicts(ignoreConflicts).Execute()

Create scheduled changes workflow



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag's key
    environmentKey := "environmentKey_example" // string | The environment key
    postFlagScheduledChangesInput := *openapiclient.NewPostFlagScheduledChangesInput(int64(123), []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // PostFlagScheduledChangesInput | 
    ignoreConflicts := true // bool | Whether or not to succeed or fail when the new instructions conflict with existing scheduled changes (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.PostFlagConfigScheduledChanges(context.Background(), projectKey, featureFlagKey, environmentKey).PostFlagScheduledChangesInput(postFlagScheduledChangesInput).IgnoreConflicts(ignoreConflicts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.PostFlagConfigScheduledChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFlagConfigScheduledChanges`: ScheduledChange
    fmt.Fprintf(os.Stdout, "Response from `ScheduledChangesApi.PostFlagConfigScheduledChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag&#39;s key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagConfigScheduledChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postFlagScheduledChangesInput** | [**PostFlagScheduledChangesInput**](PostFlagScheduledChangesInput.md) |  | 
 **ignoreConflicts** | **bool** | Whether or not to succeed or fail when the new instructions conflict with existing scheduled changes | 

### Return type

[**ScheduledChange**](ScheduledChange.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

