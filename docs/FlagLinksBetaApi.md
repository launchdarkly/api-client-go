# \FlagLinksBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlagLink**](FlagLinksBetaApi.md#CreateFlagLink) | **Post** /api/v2/flag-links/projects/{projectKey}/flags/{featureFlagKey} | Create flag link
[**DeleteFlagLink**](FlagLinksBetaApi.md#DeleteFlagLink) | **Delete** /api/v2/flag-links/projects/{projectKey}/flags/{featureFlagKey}/{id} | Delete flag link
[**GetFlagLinks**](FlagLinksBetaApi.md#GetFlagLinks) | **Get** /api/v2/flag-links/projects/{projectKey}/flags/{featureFlagKey} | List flag links
[**UpdateFlagLink**](FlagLinksBetaApi.md#UpdateFlagLink) | **Patch** /api/v2/flag-links/projects/{projectKey}/flags/{featureFlagKey}/{id} | Update flag link



## CreateFlagLink

> FlagLinkRep CreateFlagLink(ctx, projectKey, featureFlagKey).FlagLinkPost(flagLinkPost).Execute()

Create flag link



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    flagLinkPost := *openapiclient.NewFlagLinkPost() // FlagLinkPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagLinksBetaApi.CreateFlagLink(context.Background(), projectKey, featureFlagKey).FlagLinkPost(flagLinkPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagLinksBetaApi.CreateFlagLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlagLink`: FlagLinkRep
    fmt.Fprintf(os.Stdout, "Response from `FlagLinksBetaApi.CreateFlagLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlagLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flagLinkPost** | [**FlagLinkPost**](FlagLinkPost.md) |  | 

### Return type

[**FlagLinkRep**](FlagLinkRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlagLink

> DeleteFlagLink(ctx, projectKey, featureFlagKey, id).Execute()

Delete flag link



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    id := "id_example" // string | The flag link ID or Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagLinksBetaApi.DeleteFlagLink(context.Background(), projectKey, featureFlagKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagLinksBetaApi.DeleteFlagLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**id** | **string** | The flag link ID or Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagLinkRequest struct via the builder pattern


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


## GetFlagLinks

> FlagLinkCollectionRep GetFlagLinks(ctx, projectKey, featureFlagKey).Execute()

List flag links



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagLinksBetaApi.GetFlagLinks(context.Background(), projectKey, featureFlagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagLinksBetaApi.GetFlagLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagLinks`: FlagLinkCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `FlagLinksBetaApi.GetFlagLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlagLinkCollectionRep**](FlagLinkCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlagLink

> FlagLinkRep UpdateFlagLink(ctx, projectKey, featureFlagKey, id).PatchOperation(patchOperation).Execute()

Update flag link



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    id := "id_example" // string | The flag link ID
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlagLinksBetaApi.UpdateFlagLink(context.Background(), projectKey, featureFlagKey, id).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlagLinksBetaApi.UpdateFlagLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlagLink`: FlagLinkRep
    fmt.Fprintf(os.Stdout, "Response from `FlagLinksBetaApi.UpdateFlagLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**id** | **string** | The flag link ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlagLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**FlagLinkRep**](FlagLinkRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

