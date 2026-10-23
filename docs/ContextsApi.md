# \ContextsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluateContextInstance**](ContextsApi.md#EvaluateContextInstance) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/flags/evaluate | Evaluate flags for context instance
[**GetContextAttributeNames**](ContextsApi.md#GetContextAttributeNames) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes | Get context attribute names
[**GetContextAttributeValues**](ContextsApi.md#GetContextAttributeValues) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/context-attributes/{attributeName} | Get context attribute values
[**GetContextKindsByProjectKey**](ContextsApi.md#GetContextKindsByProjectKey) | **Get** /api/v2/projects/{projectKey}/context-kinds | Get context kinds
[**GetContexts**](ContextsApi.md#GetContexts) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/{kind}/{key} | Get contexts
[**PutContextKind**](ContextsApi.md#PutContextKind) | **Put** /api/v2/projects/{projectKey}/context-kinds/{key} | Create or update context kind
[**SearchContexts**](ContextsApi.md#SearchContexts) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/search | Search for contexts



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
	openapiclient "github.com/launchdarkly/api-client-go"
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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	environmentKey := "environmentKey_example" // string | The environment key
	filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts `kind` filters, with the `equals` operator, and `name` filters, with the `startsWith` operator. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). (optional)
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


 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts &#x60;kind&#x60; filters, with the &#x60;equals&#x60; operator, and &#x60;name&#x60; filters, with the &#x60;startsWith&#x60; operator. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). | 
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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	environmentKey := "environmentKey_example" // string | The environment key
	attributeName := "attributeName_example" // string | The attribute name
	filter := "filter_example" // string | A comma-separated list of context filters. This endpoint requires a `kind` filter, with the `equals` operator, and also accepts a `value` filter, with the `startsWith` operator. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). (optional)
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



 **filter** | **string** | A comma-separated list of context filters. This endpoint requires a &#x60;kind&#x60; filter, with the &#x60;equals&#x60; operator, and also accepts a &#x60;value&#x60; filter, with the &#x60;startsWith&#x60; operator. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). | 
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


## GetContextKindsByProjectKey

> ContextKindsCollectionRep GetContextKindsByProjectKey(ctx, projectKey).Expand(expand).Execute()

Get context kinds



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. Options: `environmentObservations`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContextsApi.GetContextKindsByProjectKey(context.Background(), projectKey).Expand(expand).Execute()
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

 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. Options: &#x60;environmentObservations&#x60;. | 

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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	environmentKey := "environmentKey_example" // string | The environment key
	kind := "kind_example" // string | The context kind
	key := "key_example" // string | The context key
	limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 50, default: 20) (optional)
	continuationToken := "continuationToken_example" // string | An opaque cursor that identifies the next page of results. You can use this for pagination, however, we recommend using the `next` link we provide instead. (optional)
	sort := "sort_example" // string | Deprecated. This endpoint does not support sorting. Requests that include this parameter are rejected with a `400` response, so omit it. Results are always sorted by context kind and key. (optional)
	filter := "filter_example" // string | A comma-separated list of context filters. This endpoint only accepts an `applicationId` filter. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). (optional)
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
 **continuationToken** | **string** | An opaque cursor that identifies the next page of results. You can use this for pagination, however, we recommend using the &#x60;next&#x60; link we provide instead. | 
 **sort** | **string** | Deprecated. This endpoint does not support sorting. Requests that include this parameter are rejected with a &#x60;400&#x60; response, so omit it. Results are always sorted by context kind and key. | 
 **filter** | **string** | A comma-separated list of context filters. This endpoint only accepts an &#x60;applicationId&#x60; filter. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). | 
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
	openapiclient "github.com/launchdarkly/api-client-go"
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
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	environmentKey := "environmentKey_example" // string | The environment key
	contextSearch := *openapiclient.NewContextSearch() // ContextSearch | 
	limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 50, default: 20) (optional)
	continuationToken := "continuationToken_example" // string | An opaque cursor that identifies the next page of results. You can use this for pagination, however, we recommend using the `next` link we provide instead. (optional)
	sort := "sort_example" // string | Deprecated. This endpoint does not support sorting. Requests that include this parameter are rejected with a `400` response, so omit it. Results are always sorted by context kind and key. (optional)
	filter := "filter_example" // string | A comma-separated list of context filters. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). (optional)
	includeTotalCount := true // bool | Specifies whether to include or omit the total count of matching contexts. Defaults to true. The total count is approximate and may be omitted if it is unavailable. (optional)

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
 **continuationToken** | **string** | An opaque cursor that identifies the next page of results. You can use this for pagination, however, we recommend using the &#x60;next&#x60; link we provide instead. | 
 **sort** | **string** | Deprecated. This endpoint does not support sorting. Requests that include this parameter are rejected with a &#x60;400&#x60; response, so omit it. Results are always sorted by context kind and key. | 
 **filter** | **string** | A comma-separated list of context filters. To learn more about the filter syntax, read [Filtering contexts](https://launchdarkly.com/docs/ld-docs/api/contexts#filtering-contexts). | 
 **includeTotalCount** | **bool** | Specifies whether to include or omit the total count of matching contexts. Defaults to true. The total count is approximate and may be omitted if it is unavailable. | 

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

