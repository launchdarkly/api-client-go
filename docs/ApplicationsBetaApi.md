# \ApplicationsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplication**](ApplicationsBetaApi.md#DeleteApplication) | **Delete** /api/v2/applications/{applicationKey} | Delete application
[**DeleteApplicationVersion**](ApplicationsBetaApi.md#DeleteApplicationVersion) | **Delete** /api/v2/applications/{applicationKey}/versions/{versionKey} | Delete application version
[**GetApplication**](ApplicationsBetaApi.md#GetApplication) | **Get** /api/v2/applications/{applicationKey} | Get application by key
[**GetApplicationVersions**](ApplicationsBetaApi.md#GetApplicationVersions) | **Get** /api/v2/applications/{applicationKey}/versions | Get application versions by application key
[**GetApplications**](ApplicationsBetaApi.md#GetApplications) | **Get** /api/v2/applications | Get applications
[**PatchApplication**](ApplicationsBetaApi.md#PatchApplication) | **Patch** /api/v2/applications/{applicationKey} | Update application
[**PatchApplicationVersion**](ApplicationsBetaApi.md#PatchApplicationVersion) | **Patch** /api/v2/applications/{applicationKey}/versions/{versionKey} | Update application version



## DeleteApplication

> DeleteApplication(ctx, applicationKey).Execute()

Delete application



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
    applicationKey := "applicationKey_example" // string | The application key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsBetaApi.DeleteApplication(context.Background(), applicationKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## DeleteApplicationVersion

> DeleteApplicationVersion(ctx, applicationKey, versionKey).Execute()

Delete application version



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
    applicationKey := "applicationKey_example" // string | The application key
    versionKey := "versionKey_example" // string | The application version key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsBetaApi.DeleteApplicationVersion(context.Background(), applicationKey, versionKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaApi.DeleteApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 
**versionKey** | **string** | The application version key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationVersionRequest struct via the builder pattern


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


## GetApplication

> ApplicationRep GetApplication(ctx, applicationKey).Expand(expand).Execute()

Get application by key



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
    applicationKey := "applicationKey_example" // string | The application key
    expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. Options: `flags`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsBetaApi.GetApplication(context.Background(), applicationKey).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: ApplicationRep
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. Options: &#x60;flags&#x60;. | 

### Return type

[**ApplicationRep**](ApplicationRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationVersions

> ApplicationVersionsCollectionRep GetApplicationVersions(ctx, applicationKey).Filter(filter).Limit(limit).Offset(offset).Sort(sort).Execute()

Get application versions by application key



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
    applicationKey := "applicationKey_example" // string | The application key
    filter := "filter_example" // string | Accepts filter by `key`, `name`, `supported`, and `autoAdded`. Example: `filter=key equals 'test-key'`. To learn more about the filter syntax, read [Filtering applications and application versions](/tag/Applications-(beta)#filtering-contexts-and-context-instances). (optional)
    limit := int64(789) // int64 | The number of versions to return. Defaults to 50. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
    sort := "sort_example" // string | Accepts sorting order and fields. Fields can be comma separated. Possible fields are `creationDate`, `name`. Examples: `sort=name` sort by names ascending, `sort=-name,creationDate` sort by names descending and creationDate ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsBetaApi.GetApplicationVersions(context.Background(), applicationKey).Filter(filter).Limit(limit).Offset(offset).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaApi.GetApplicationVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationVersions`: ApplicationVersionsCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaApi.GetApplicationVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Accepts filter by &#x60;key&#x60;, &#x60;name&#x60;, &#x60;supported&#x60;, and &#x60;autoAdded&#x60;. Example: &#x60;filter&#x3D;key equals &#39;test-key&#39;&#x60;. To learn more about the filter syntax, read [Filtering applications and application versions](/tag/Applications-(beta)#filtering-contexts-and-context-instances). | 
 **limit** | **int64** | The number of versions to return. Defaults to 50. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **sort** | **string** | Accepts sorting order and fields. Fields can be comma separated. Possible fields are &#x60;creationDate&#x60;, &#x60;name&#x60;. Examples: &#x60;sort&#x3D;name&#x60; sort by names ascending, &#x60;sort&#x3D;-name,creationDate&#x60; sort by names descending and creationDate ascending. | 

### Return type

[**ApplicationVersionsCollectionRep**](ApplicationVersionsCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationCollectionRep GetApplications(ctx).Filter(filter).Limit(limit).Offset(offset).Sort(sort).Expand(expand).Execute()

Get applications



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
    filter := "filter_example" // string | Accepts filter by `key`, `name`, `kind`, and `autoAdded`. Example: `filter=kind anyOf ['mobile', 'server'],key equals 'test-key'`. To learn more about the filter syntax, read [Filtering applications and application versions](/tag/Applications-(beta)#filtering-contexts-and-context-instances). (optional)
    limit := int64(789) // int64 | The number of applications to return. Defaults to 10. (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
    sort := "sort_example" // string | Accepts sorting order and fields. Fields can be comma separated. Possible fields are `creationDate`, `name`. Examples: `sort=name` sort by names ascending, `sort=-name,creationDate` sort by names descending and creationDate ascending. (optional)
    expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. Options: `flags`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsBetaApi.GetApplications(context.Background()).Filter(filter).Limit(limit).Offset(offset).Sort(sort).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: ApplicationCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Accepts filter by &#x60;key&#x60;, &#x60;name&#x60;, &#x60;kind&#x60;, and &#x60;autoAdded&#x60;. Example: &#x60;filter&#x3D;kind anyOf [&#39;mobile&#39;, &#39;server&#39;],key equals &#39;test-key&#39;&#x60;. To learn more about the filter syntax, read [Filtering applications and application versions](/tag/Applications-(beta)#filtering-contexts-and-context-instances). | 
 **limit** | **int64** | The number of applications to return. Defaults to 10. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **sort** | **string** | Accepts sorting order and fields. Fields can be comma separated. Possible fields are &#x60;creationDate&#x60;, &#x60;name&#x60;. Examples: &#x60;sort&#x3D;name&#x60; sort by names ascending, &#x60;sort&#x3D;-name,creationDate&#x60; sort by names descending and creationDate ascending. | 
 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. Options: &#x60;flags&#x60;. | 

### Return type

[**ApplicationCollectionRep**](ApplicationCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApplication

> ApplicationRep PatchApplication(ctx, applicationKey).PatchOperation(patchOperation).Execute()

Update application



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
    applicationKey := "applicationKey_example" // string | The application key
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsBetaApi.PatchApplication(context.Background(), applicationKey).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaApi.PatchApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApplication`: ApplicationRep
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaApi.PatchApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**ApplicationRep**](ApplicationRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApplicationVersion

> ApplicationVersionRep PatchApplicationVersion(ctx, applicationKey, versionKey).PatchOperation(patchOperation).Execute()

Update application version



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
    applicationKey := "applicationKey_example" // string | The application key
    versionKey := "versionKey_example" // string | The application version key
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsBetaApi.PatchApplicationVersion(context.Background(), applicationKey, versionKey).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsBetaApi.PatchApplicationVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApplicationVersion`: ApplicationVersionRep
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsBetaApi.PatchApplicationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The application key | 
**versionKey** | **string** | The application version key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**ApplicationVersionRep**](ApplicationVersionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

