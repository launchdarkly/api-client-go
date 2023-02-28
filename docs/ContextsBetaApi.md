# \ContextsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContextInstances**](ContextsBetaApi.md#DeleteContextInstances) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/{id} | Delete context instances
[**EvaluateContextInstance**](ContextsBetaApi.md#EvaluateContextInstance) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/flags/evaluate | Evaluate flags for context instance
[**GetContextAttributeNames**](ContextsBetaApi.md#GetContextAttributeNames) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes | Get context attribute names
[**GetContextAttributeValues**](ContextsBetaApi.md#GetContextAttributeValues) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes/{attributeName} | Get context attribute values
[**GetContextInstances**](ContextsBetaApi.md#GetContextInstances) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/{id} | Get context instances
[**GetContextKindsByProjectKey**](ContextsBetaApi.md#GetContextKindsByProjectKey) | **Get** /api/v2/projects/{projectKey}/context-kinds | Get context kinds
[**GetContexts**](ContextsBetaApi.md#GetContexts) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/{kind}/{key} | Get contexts
[**PutContextKind**](ContextsBetaApi.md#PutContextKind) | **Put** /api/v2/projects/{projectKey}/context-kinds/{key} | Create or update context kind
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


## EvaluateContextInstance

> ContextInstanceEvaluations EvaluateContextInstance(ctx, projectKey, environmentKey).RequestBody(requestBody).Limit(limit).Offset(offset).Sort(sort).Filter(filter).Execute()

Evaluate flags for context instance



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
    limit := int64(789) // int64 | The number of feature flags to return. Defaults to -1, which returns all flags (optional)
    offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
    sort := "sort_example" // string | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order (optional)
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form field:value. Supports the same filters as the List Feature Flags API. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.EvaluateContextInstance(context.Background(), projectKey, environmentKey).RequestBody(requestBody).Limit(limit).Offset(offset).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.EvaluateContextInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EvaluateContextInstance`: ContextInstanceEvaluations
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.EvaluateContextInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateContextInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **limit** | **int64** | The number of feature flags to return. Defaults to -1, which returns all flags | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **sort** | **string** | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order | 
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form field:value. Supports the same filters as the List Feature Flags API. | 

### Return type

[**ContextInstanceEvaluations**](ContextInstanceEvaluations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
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

> ContextInstances GetContextInstances(ctx, projectKey, environmentKey, id).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()

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
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching context instances. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContextInstances(context.Background(), projectKey, environmentKey, id).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
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
 **includeTotalCount** | **bool** | Specifies whether to include or omit the total count of matching context instances. Defaults to true. | 

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


## GetContextKindsByProjectKey

> ContextKindsCollectionRep GetContextKindsByProjectKey(ctx, projectKey).Execute()

Get context kinds



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContextKindsByProjectKey(context.Background(), projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.GetContextKindsByProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextKindsByProjectKey`: ContextKindsCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.GetContextKindsByProjectKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextKindsByProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContextKindsCollectionRep**](ContextKindsCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContexts

> Contexts GetContexts(ctx, projectKey, environmentKey, kind, key).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()

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
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching contexts. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.GetContexts(context.Background(), projectKey, environmentKey, kind, key).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
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
 **includeTotalCount** | **bool** | Specifies whether to include or omit the total count of matching contexts. Defaults to true. | 

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


## PutContextKind

> UpsertResponseRep PutContextKind(ctx, projectKey, key).UpsertContextKindPayload(upsertContextKindPayload).Execute()

Create or update context kind



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
    key := "key_example" // string | The context kind key
    upsertContextKindPayload := *openapiclient.NewUpsertContextKindPayload("organization") // UpsertContextKindPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.PutContextKind(context.Background(), projectKey, key).UpsertContextKindPayload(upsertContextKindPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsBetaApi.PutContextKind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutContextKind`: UpsertResponseRep
    fmt.Fprintf(os.Stdout, "Response from `ContextsBetaApi.PutContextKind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**key** | **string** | The context kind key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutContextKindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upsertContextKindPayload** | [**UpsertContextKindPayload**](UpsertContextKindPayload.md) |  | 

### Return type

[**UpsertResponseRep**](UpsertResponseRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchContextInstances

> ContextInstances SearchContextInstances(ctx, projectKey, environmentKey).ContextInstanceSearch(contextInstanceSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()

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
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching context instances. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.SearchContextInstances(context.Background(), projectKey, environmentKey).ContextInstanceSearch(contextInstanceSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
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
 **includeTotalCount** | **bool** | Specifies whether to include or omit the total count of matching context instances. Defaults to true. | 

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

> Contexts SearchContexts(ctx, projectKey, environmentKey).ContextSearch(contextSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()

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
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching contexts. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsBetaApi.SearchContexts(context.Background(), projectKey, environmentKey).ContextSearch(contextSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
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
 **includeTotalCount** | **bool** | Specifies whether to include or omit the total count of matching contexts. Defaults to true. | 

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

