# \ContextsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContextInstances**](ContextsBetaApi.md#DeleteContextInstances) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/{id} | Delete context instances
[**GetContextAttributeNames**](ContextsBetaApi.md#GetContextAttributeNames) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes | Get context attribute names
[**GetContextAttributeValues**](ContextsBetaApi.md#GetContextAttributeValues) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes/{attributeName} | Get context attribute values
[**GetContextInstances**](ContextsBetaApi.md#GetContextInstances) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/{id} | Get context instances
[**GetContexts**](ContextsBetaApi.md#GetContexts) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/{kind}/{key} | Get contexts
[**PutFlagSettingForContext**](ContextsBetaApi.md#PutFlagSettingForContext) | **Put** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/{contextKind}/{contextKey}/flags/{featureFlagKey} | Update flag settings for context
[**SearchContextInstances**](ContextsBetaApi.md#SearchContextInstances) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/search | Search for context instances
[**SearchContexts**](ContextsBetaApi.md#SearchContexts) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/search | Search for contexts



## DeleteContextInstances

> DeleteContextInstances(ctx, projectKey, environmentKey, id).Execute()

Delete context instances



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
    id := "id_example" // string | The context instance ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.DeleteContextInstances(context.Background(), projectKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.DeleteContextInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The context instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextInstancesRequest struct via the builder pattern


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


## GetContextAttributeNames

> ContextAttributeNamesCollection GetContextAttributeNames(ctx, projectKey, environmentKey).Filter(filter).Execute()

Get context attribute names



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
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts `kind` filters, with the `equals` operator, and `name` filters, with the `startsWith` operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContextAttributeNames(context.Background(), projectKey, environmentKey).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.GetContextAttributeNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextAttributeNames`: ContextAttributeNamesCollection
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.GetContextAttributeNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextAttributeNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts &#x60;kind&#x60; filters, with the &#x60;equals&#x60; operator, and &#x60;name&#x60; filters, with the &#x60;startsWith&#x60; operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). | 

### Return type

[**ContextAttributeNamesCollection**](ContextAttributeNamesCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextAttributeValues

> ContextAttributeValuesCollection GetContextAttributeValues(ctx, projectKey, environmentKey, attributeName).Filter(filter).Execute()

Get context attribute values



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
    attributeName := "attributeName_example" // string | The attribute name
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts `kind` filters, with the `equals` operator, and `value` filters, with the `startsWith` operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContextAttributeValues(context.Background(), projectKey, environmentKey, attributeName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.GetContextAttributeValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextAttributeValues`: ContextAttributeValuesCollection
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.GetContextAttributeValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**attributeName** | **string** | The attribute name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextAttributeValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts &#x60;kind&#x60; filters, with the &#x60;equals&#x60; operator, and &#x60;value&#x60; filters, with the &#x60;startsWith&#x60; operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). | 

### Return type

[**ContextAttributeValuesCollection**](ContextAttributeValuesCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextInstances

> ContextInstances GetContextInstances(ctx, projectKey, environmentKey, id).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()

Get context instances



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
    id := "id_example" // string | The context instance ID
    limit := int64(789) // int64 | Specifies the maximum number of context instances to return (max: 50, default: 20) (optional)
    continuationToken := "continuationToken_example" // string | Limits results to context instances with sort values after the value specified. You can use this for pagination, however, we recommend using the `next` link we provide instead. (optional)
    sort := "sort_example" // string | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying `ts` for this value, or descending order by specifying `-ts`. (optional)
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts an `applicationId` filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContextInstances(context.Background(), projectKey, environmentKey, id).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.GetContextInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextInstances`: ContextInstances
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.GetContextInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The context instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int64** | Specifies the maximum number of context instances to return (max: 50, default: 20) | 
 **continuationToken** | **string** | Limits results to context instances with sort values after the value specified. You can use this for pagination, however, we recommend using the &#x60;next&#x60; link we provide instead. | 
 **sort** | **string** | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying &#x60;ts&#x60; for this value, or descending order by specifying &#x60;-ts&#x60;. | 
 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts an &#x60;applicationId&#x60; filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). | 

### Return type

[**ContextInstances**](ContextInstances.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContexts

> Contexts GetContexts(ctx, projectKey, environmentKey, kind, key).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()

Get contexts



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
    kind := "kind_example" // string | The context kind
    key := "key_example" // string | The context key
    limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 50, default: 20) (optional)
    continuationToken := "continuationToken_example" // string | Limits results to contexts with sort values after the value specified. You can use this for pagination, however, we recommend using the `next` link we provide instead. (optional)
    sort := "sort_example" // string | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying `ts` for this value, or descending order by specifying `-ts`. (optional)
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts an `applicationId` filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContexts(context.Background(), projectKey, environmentKey, kind, key).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.GetContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContexts`: Contexts
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.GetContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**kind** | **string** | The context kind | 
**key** | **string** | The context key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **int64** | Specifies the maximum number of items in the collection to return (max: 50, default: 20) | 
 **continuationToken** | **string** | Limits results to contexts with sort values after the value specified. You can use this for pagination, however, we recommend using the &#x60;next&#x60; link we provide instead. | 
 **sort** | **string** | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying &#x60;ts&#x60; for this value, or descending order by specifying &#x60;-ts&#x60;. | 
 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts an &#x60;applicationId&#x60; filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). | 

### Return type

[**Contexts**](Contexts.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFlagSettingForContext

> PutFlagSettingForContext(ctx, projectKey, environmentKey, contextKind, contextKey, featureFlagKey).ValuePut(valuePut).Execute()

Update flag settings for context



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
    contextKind := "contextKind_example" // string | The context kind
    contextKey := "contextKey_example" // string | The context key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    valuePut := *openapiclient.NewValuePut() // ValuePut | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.PutFlagSettingForContext(context.Background(), projectKey, environmentKey, contextKind, contextKey, featureFlagKey).ValuePut(valuePut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.PutFlagSettingForContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**contextKind** | **string** | The context kind | 
**contextKey** | **string** | The context key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFlagSettingForContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **valuePut** | [**ValuePut**](ValuePut.md) |  | 

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


## SearchContextInstances

> ContextInstances SearchContextInstances(ctx, projectKey, environmentKey).ContextInstanceSearch(contextInstanceSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()

Search for context instances



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
    contextInstanceSearch := *openapiclient.NewContextInstanceSearch() // ContextInstanceSearch | 
    limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 50, default: 20) (optional)
    continuationToken := "continuationToken_example" // string | Limits results to context instances with sort values after the value specified. You can use this for pagination, however, we recommend using the `next` link we provide instead. (optional)
    sort := "sort_example" // string | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying `ts` for this value, or descending order by specifying `-ts`. (optional)
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts an `applicationId` filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.SearchContextInstances(context.Background(), projectKey, environmentKey).ContextInstanceSearch(contextInstanceSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.SearchContextInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContextInstances`: ContextInstances
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.SearchContextInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchContextInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contextInstanceSearch** | [**ContextInstanceSearch**](ContextInstanceSearch.md) |  | 
 **limit** | **int64** | Specifies the maximum number of items in the collection to return (max: 50, default: 20) | 
 **continuationToken** | **string** | Limits results to context instances with sort values after the value specified. You can use this for pagination, however, we recommend using the &#x60;next&#x60; link we provide instead. | 
 **sort** | **string** | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying &#x60;ts&#x60; for this value, or descending order by specifying &#x60;-ts&#x60;. | 
 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts an &#x60;applicationId&#x60; filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). | 

### Return type

[**ContextInstances**](ContextInstances.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchContexts

> Contexts SearchContexts(ctx, projectKey, environmentKey).ContextSearch(contextSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()

Search for contexts



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
    contextSearch := *openapiclient.NewContextSearch() // ContextSearch | 
    limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 50, default: 20) (optional)
    continuationToken := "continuationToken_example" // string | Limits results to contexts with sort values after the value specified. You can use this for pagination, however, we recommend using the `next` link we provide instead. (optional)
    sort := "sort_example" // string | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying `ts` for this value, or descending order by specifying `-ts`. (optional)
    filter := "filter_example" // string | A comma-separated list of context filters. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.SearchContexts(context.Background(), projectKey, environmentKey).ContextSearch(contextSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.SearchContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContexts`: Contexts
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.SearchContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contextSearch** | [**ContextSearch**](ContextSearch.md) |  | 
 **limit** | **int64** | Specifies the maximum number of items in the collection to return (max: 50, default: 20) | 
 **continuationToken** | **string** | Limits results to contexts with sort values after the value specified. You can use this for pagination, however, we recommend using the &#x60;next&#x60; link we provide instead. | 
 **sort** | **string** | Specifies a field by which to sort. LaunchDarkly supports sorting by timestamp in ascending order by specifying &#x60;ts&#x60; for this value, or descending order by specifying &#x60;-ts&#x60;. | 
 **filter** | **string** | A comma-separated list of context filters. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts-(beta)#filtering-contexts-and-context-instances). | 

### Return type

[**Contexts**](Contexts.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

