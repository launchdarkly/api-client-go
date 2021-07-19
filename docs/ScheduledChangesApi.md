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

> ScheduledChangesRep GetFeatureFlagScheduledChange(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

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
    // response from `GetFeatureFlagScheduledChange`: ScheduledChangesRep
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

[**ScheduledChangesRep**](ScheduledChangesRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagConfigScheduledChanges

> ScheduledChangesCollectionRep GetFlagConfigScheduledChanges(ctx, projectKey, featureFlagKey, environmentKey).Execute()

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
    // response from `GetFlagConfigScheduledChanges`: ScheduledChangesCollectionRep
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

[**ScheduledChangesCollectionRep**](ScheduledChangesCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagConfigScheduledChangesConflicts

> WebConflictResponse GetFlagConfigScheduledChangesConflicts(ctx, projectKey, featureFlagKey, environmentKey).WebReportFlagScheduledChangesInput(webReportFlagScheduledChangesInput).Execute()

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
    webReportFlagScheduledChangesInput := *openapiclient.NewWebReportFlagScheduledChangesInput() // WebReportFlagScheduledChangesInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.GetFlagConfigScheduledChangesConflicts(context.Background(), projectKey, featureFlagKey, environmentKey).WebReportFlagScheduledChangesInput(webReportFlagScheduledChangesInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.GetFlagConfigScheduledChangesConflicts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagConfigScheduledChangesConflicts`: WebConflictResponse
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



 **webReportFlagScheduledChangesInput** | [**WebReportFlagScheduledChangesInput**](WebReportFlagScheduledChangesInput.md) |  | 

### Return type

[**WebConflictResponse**](WebConflictResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFlagConfigScheduledChange

> ScheduledChangesRep PatchFlagConfigScheduledChange(ctx, projectKey, featureFlagKey, environmentKey, id).WebFlagScheduledChangesInput(webFlagScheduledChangesInput).Execute()

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
    id := "id_example" // string | The scheduled change id
    webFlagScheduledChangesInput := *openapiclient.NewWebFlagScheduledChangesInput() // WebFlagScheduledChangesInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.PatchFlagConfigScheduledChange(context.Background(), projectKey, featureFlagKey, environmentKey, id).WebFlagScheduledChangesInput(webFlagScheduledChangesInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.PatchFlagConfigScheduledChange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFlagConfigScheduledChange`: ScheduledChangesRep
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
**id** | **string** | The scheduled change id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlagConfigScheduledChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **webFlagScheduledChangesInput** | [**WebFlagScheduledChangesInput**](WebFlagScheduledChangesInput.md) |  | 

### Return type

[**ScheduledChangesRep**](ScheduledChangesRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlagConfigScheduledChanges

> ScheduledChangesRep PostFlagConfigScheduledChanges(ctx, projectKey, featureFlagKey, environmentKey).WebPostFlagScheduledChangesInput(webPostFlagScheduledChangesInput).IgnoreConflicts(ignoreConflicts).Execute()

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
    webPostFlagScheduledChangesInput := *openapiclient.NewWebPostFlagScheduledChangesInput() // WebPostFlagScheduledChangesInput | 
    ignoreConflicts := true // bool | Pass in true to error on conflicts with current flag state or existing scheduled changes (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledChangesApi.PostFlagConfigScheduledChanges(context.Background(), projectKey, featureFlagKey, environmentKey).WebPostFlagScheduledChangesInput(webPostFlagScheduledChangesInput).IgnoreConflicts(ignoreConflicts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledChangesApi.PostFlagConfigScheduledChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostFlagConfigScheduledChanges`: ScheduledChangesRep
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



 **webPostFlagScheduledChangesInput** | [**WebPostFlagScheduledChangesInput**](WebPostFlagScheduledChangesInput.md) |  | 
 **ignoreConflicts** | **bool** | Pass in true to error on conflicts with current flag state or existing scheduled changes | 

### Return type

[**ScheduledChangesRep**](ScheduledChangesRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

