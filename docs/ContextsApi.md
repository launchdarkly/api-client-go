# \ContextsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteContextInstances**](ContextsApi.md#DeleteContextInstances) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/{id} | Delete context instances
[**EvaluateContextInstance**](ContextsApi.md#EvaluateContextInstance) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/flags/evaluate | Evaluate flags for context instance
[**GetContextAttributeNames**](ContextsApi.md#GetContextAttributeNames) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes | Get context attribute names
[**GetContextAttributeValues**](ContextsApi.md#GetContextAttributeValues) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes/{attributeName} | Get context attribute values
[**GetContextInstances**](ContextsApi.md#GetContextInstances) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/{id} | Get context instances
[**GetContextKindsByProjectKey**](ContextsApi.md#GetContextKindsByProjectKey) | **Get** /api/v2/projects/{projectKey}/context-kinds | Get context kinds
[**GetContexts**](ContextsApi.md#GetContexts) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/{kind}/{key} | Get contexts
[**PutContextKind**](ContextsApi.md#PutContextKind) | **Put** /api/v2/projects/{projectKey}/context-kinds/{key} | Create or update context kind
[**SearchContextInstances**](ContextsApi.md#SearchContextInstances) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-instances/search | Search for context instances
[**SearchContexts**](ContextsApi.md#SearchContexts) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/search | Search for contexts



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
    resp, r, err := apiClient.ContextsApi.DeleteContextInstances(context.Background(), projectKey, environmentKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.DeleteContextInstances``: %v\n", err)
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
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form `field operator value`. Supported fields are explained above. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.EvaluateContextInstance(context.Background(), projectKey, environmentKey).RequestBody(requestBody).Limit(limit).Offset(offset).Sort(sort).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.EvaluateContextInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EvaluateContextInstance`: ContextInstanceEvaluations
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.EvaluateContextInstance`: %v\n", resp)
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
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form &#x60;field operator value&#x60;. Supported fields are explained above. | 

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

> ContextAttributeNamesCollection GetContextAttributeNames(ctx, projectKey, environmentKey).Filter(filter).Limit(limit).Execute()

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
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts `kind` filters, with the `equals` operator, and `name` filters, with the `startsWith` operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). (optional)
    limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 100, default: 100) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.GetContextAttributeNames(context.Background(), projectKey, environmentKey).Filter(filter).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetContextAttributeNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextAttributeNames`: ContextAttributeNamesCollection
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetContextAttributeNames`: %v\n", resp)
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


 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts &#x60;kind&#x60; filters, with the &#x60;equals&#x60; operator, and &#x60;name&#x60; filters, with the &#x60;startsWith&#x60; operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). | 
 **limit** | **int64** | Specifies the maximum number of items in the collection to return (max: 100, default: 100) | 

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

> ContextAttributeValuesCollection GetContextAttributeValues(ctx, projectKey, environmentKey, attributeName).Filter(filter).Limit(limit).Execute()

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
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts `kind` filters, with the `equals` operator, and `value` filters, with the `startsWith` operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). (optional)
    limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 100, default: 50) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.GetContextAttributeValues(context.Background(), projectKey, environmentKey, attributeName).Filter(filter).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetContextAttributeValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextAttributeValues`: ContextAttributeValuesCollection
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetContextAttributeValues`: %v\n", resp)
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



 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts &#x60;kind&#x60; filters, with the &#x60;equals&#x60; operator, and &#x60;value&#x60; filters, with the &#x60;startsWith&#x60; operator. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). | 
 **limit** | **int64** | Specifies the maximum number of items in the collection to return (max: 100, default: 50) | 

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
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts an `applicationId` filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). (optional)
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching context instances. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.GetContextInstances(context.Background(), projectKey, environmentKey, id).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetContextInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextInstances`: ContextInstances
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetContextInstances`: %v\n", resp)
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
 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts an &#x60;applicationId&#x60; filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). | 
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
    resp, r, err := apiClient.ContextsApi.GetContextKindsByProjectKey(context.Background(), projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetContextKindsByProjectKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextKindsByProjectKey`: ContextKindsCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetContextKindsByProjectKey`: %v\n", resp)
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
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts an `applicationId` filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). (optional)
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching contexts. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.GetContexts(context.Background(), projectKey, environmentKey, kind, key).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.GetContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContexts`: Contexts
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.GetContexts`: %v\n", resp)
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
 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts an &#x60;applicationId&#x60; filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). | 
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
    resp, r, err := apiClient.ContextsApi.PutContextKind(context.Background(), projectKey, key).UpsertContextKindPayload(upsertContextKindPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.PutContextKind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutContextKind`: UpsertResponseRep
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.PutContextKind`: %v\n", resp)
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
    filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts an `applicationId` filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). (optional)
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching context instances. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.SearchContextInstances(context.Background(), projectKey, environmentKey).ContextInstanceSearch(contextInstanceSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.SearchContextInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContextInstances`: ContextInstances
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.SearchContextInstances`: %v\n", resp)
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
 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts an &#x60;applicationId&#x60; filter. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). | 
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
    filter := "filter_example" // string | A comma-separated list of context filters. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). (optional)
    includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching contexts. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextsApi.SearchContexts(context.Background(), projectKey, environmentKey).ContextSearch(contextSearch).Limit(limit).ContinuationToken(continuationToken).Sort(sort).Filter(filter).IncludeTotalCount(includeTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextsApi.SearchContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchContexts`: Contexts
    fmt.Fprintf(os.Stdout, "Response from `ContextsApi.SearchContexts`: %v\n", resp)
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
 **filter** | **string** | A comma-separated list of context filters. To learn more about the filter syntax, read [Filtering contexts and context instances](/tag/Contexts#filtering-contexts-and-context-instances). | 
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

