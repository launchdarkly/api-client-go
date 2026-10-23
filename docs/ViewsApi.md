# \ViewsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateView**](ViewsApi.md#CreateView) | **Post** /api/v2/projects/{projectKey}/views | Create view
[**DeleteView**](ViewsApi.md#DeleteView) | **Delete** /api/v2/projects/{projectKey}/views/{viewKey} | Delete view
[**GetLinkedResources**](ViewsApi.md#GetLinkedResources) | **Get** /api/v2/projects/{projectKey}/views/{viewKey}/linked/{resourceType} | Get linked resources
[**GetLinkedViews**](ViewsApi.md#GetLinkedViews) | **Get** /api/v2/projects/{projectKey}/view-associations/{resourceType}/{resourceKey} | Get linked views for a given resource
[**GetView**](ViewsApi.md#GetView) | **Get** /api/v2/projects/{projectKey}/views/{viewKey} | Get view
[**GetViews**](ViewsApi.md#GetViews) | **Get** /api/v2/projects/{projectKey}/views | List views
[**LinkResource**](ViewsApi.md#LinkResource) | **Post** /api/v2/projects/{projectKey}/views/{viewKey}/link/{resourceType} | Link resource
[**UnlinkResource**](ViewsApi.md#UnlinkResource) | **Delete** /api/v2/projects/{projectKey}/views/{viewKey}/link/{resourceType} | Unlink resource
[**UpdateView**](ViewsApi.md#UpdateView) | **Patch** /api/v2/projects/{projectKey}/views/{viewKey} | Update view



## CreateView

> View CreateView(ctx, projectKey).ViewPost(viewPost).Execute()

Create view



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
	projectKey := "default" // string | 
	viewPost := *openapiclient.NewViewPost("Key_example", "Name_example") // ViewPost | View object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.CreateView(context.Background(), projectKey).ViewPost(viewPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.CreateView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.CreateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewPost** | [**ViewPost**](ViewPost.md) | View object to create | 

### Return type

[**View**](View.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteView

> DeleteView(ctx, projectKey, viewKey).Execute()

Delete view



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
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewsApi.DeleteView(context.Background(), projectKey, viewKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.DeleteView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**viewKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteViewRequest struct via the builder pattern


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


## GetLinkedResources

> ViewLinkedResources GetLinkedResources(ctx, projectKey, viewKey, resourceType).Limit(limit).Offset(offset).Sort(sort).Query(query).Filter(filter).Expand(expand).Execute()

Get linked resources



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
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	resourceType := "flags" // string | 
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	sort := "sort_example" // string | Field to sort by. Default field is `linkedAt`, default order is ascending. (optional) (default to "linkedAt")
	query := "query_example" // string | Case-insensitive search query for linked resources. Matches resource key and, when expanded, resource name. (optional)
	filter := "filter_example" // string | Optional resource filter expression for linked resources. - Supported for `flags` and `segments` resource types. - Uses the same syntax as link/unlink and list endpoints. - For `segments`, `environmentId` is required when `filter` is provided.  (optional)
	expand := []string{"Expand_example"} // []string | A comma-separated list of fields to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.GetLinkedResources(context.Background(), projectKey, viewKey, resourceType).Limit(limit).Offset(offset).Sort(sort).Query(query).Filter(filter).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetLinkedResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedResources`: ViewLinkedResources
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetLinkedResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**viewKey** | **string** |  | 
**resourceType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The number of views to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **sort** | **string** | Field to sort by. Default field is &#x60;linkedAt&#x60;, default order is ascending. | [default to &quot;linkedAt&quot;]
 **query** | **string** | Case-insensitive search query for linked resources. Matches resource key and, when expanded, resource name. | 
 **filter** | **string** | Optional resource filter expression for linked resources. - Supported for &#x60;flags&#x60; and &#x60;segments&#x60; resource types. - Uses the same syntax as link/unlink and list endpoints. - For &#x60;segments&#x60;, &#x60;environmentId&#x60; is required when &#x60;filter&#x60; is provided.  | 
 **expand** | **[]string** | A comma-separated list of fields to expand. | 

### Return type

[**ViewLinkedResources**](ViewLinkedResources.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedViews

> Views GetLinkedViews(ctx, projectKey, resourceType, resourceKey).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()

Get linked views for a given resource



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
	projectKey := "default" // string | 
	resourceType := "flags" // string | 
	resourceKey := "my-flag" // string | 
	environmentId := "6890ff25c3e3830ba1a352e4" // string | Environment ID. Required when resourceType is 'segments' (optional)
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.GetLinkedViews(context.Background(), projectKey, resourceType, resourceKey).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetLinkedViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedViews`: Views
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetLinkedViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**resourceType** | **string** |  | 
**resourceKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **environmentId** | **string** | Environment ID. Required when resourceType is &#39;segments&#39; | 
 **limit** | **int32** | The number of views to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**Views**](Views.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetView

> View GetView(ctx, projectKey, viewKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()

Get view



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
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	sort := "sort_example" // string | A sort to apply to the list of views. (optional)
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of views. Supports the following fields and operators: `name` (equals, notEquals, startsWith, contains, anyOf), `key` (equals, notEquals, startsWith, contains, anyOf), `tag` (equals, anyOf), `maintainerId` (equals, anyOf), `isPayloadView` (equals). (optional)
	expand := []string{"Expand_example"} // []string | A comma-separated list of fields to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.GetView(context.Background(), projectKey, viewKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**viewKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **string** | A sort to apply to the list of views. | 
 **limit** | **int32** | The number of views to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list of views. Supports the following fields and operators: &#x60;name&#x60; (equals, notEquals, startsWith, contains, anyOf), &#x60;key&#x60; (equals, notEquals, startsWith, contains, anyOf), &#x60;tag&#x60; (equals, anyOf), &#x60;maintainerId&#x60; (equals, anyOf), &#x60;isPayloadView&#x60; (equals). | 
 **expand** | **[]string** | A comma-separated list of fields to expand. | 

### Return type

[**View**](View.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViews

> Views GetViews(ctx, projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()

List views



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
	projectKey := "default" // string | 
	sort := "sort_example" // string | A sort to apply to the list of views. (optional)
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of views. Supports the following fields and operators: `name` (equals, notEquals, startsWith, contains, anyOf), `key` (equals, notEquals, startsWith, contains, anyOf), `tag` (equals, anyOf), `maintainerId` (equals, anyOf), `isPayloadView` (equals). (optional)
	expand := []string{"Expand_example"} // []string | A comma-separated list of fields to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.GetViews(context.Background(), projectKey).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.GetViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViews`: Views
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.GetViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | A sort to apply to the list of views. | 
 **limit** | **int32** | The number of views to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list of views. Supports the following fields and operators: &#x60;name&#x60; (equals, notEquals, startsWith, contains, anyOf), &#x60;key&#x60; (equals, notEquals, startsWith, contains, anyOf), &#x60;tag&#x60; (equals, anyOf), &#x60;maintainerId&#x60; (equals, anyOf), &#x60;isPayloadView&#x60; (equals). | 
 **expand** | **[]string** | A comma-separated list of fields to expand. | 

### Return type

[**Views**](Views.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkResource

> LinkResourceSuccessResponse LinkResource(ctx, projectKey, viewKey, resourceType).ViewLinkRequest(viewLinkRequest).Execute()

Link resource



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
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	resourceType := "flags" // string | 
	viewLinkRequest := openapiclient.ViewLinkRequest{ViewLinkRequestFilter: openapiclient.NewViewLinkRequestFilter("maintainerId:507f1f77bcf86cd799439011,tags:backend+beta")} // ViewLinkRequest | Resources to link to the view. You can provide explicit keys/IDs, filters, or both. - Flags: identified by key or filtered by maintainerId, maintainerTeamKey, tags, state, query - Segments: identified by segment ID or filtered by tags, query, unbounded 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.LinkResource(context.Background(), projectKey, viewKey, resourceType).ViewLinkRequest(viewLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.LinkResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkResource`: LinkResourceSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.LinkResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**viewKey** | **string** |  | 
**resourceType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **viewLinkRequest** | [**ViewLinkRequest**](ViewLinkRequest.md) | Resources to link to the view. You can provide explicit keys/IDs, filters, or both. - Flags: identified by key or filtered by maintainerId, maintainerTeamKey, tags, state, query - Segments: identified by segment ID or filtered by tags, query, unbounded  | 

### Return type

[**LinkResourceSuccessResponse**](LinkResourceSuccessResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkResource

> UnlinkResourceSuccessResponse UnlinkResource(ctx, projectKey, viewKey, resourceType).ViewLinkRequest(viewLinkRequest).Execute()

Unlink resource



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
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	resourceType := "flags" // string | 
	viewLinkRequest := openapiclient.ViewLinkRequest{ViewLinkRequestFilter: openapiclient.NewViewLinkRequestFilter("maintainerId:507f1f77bcf86cd799439011,tags:backend+beta")} // ViewLinkRequest | The resource to link to the view. Flags are identified by key. Segments are identified by segment ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.UnlinkResource(context.Background(), projectKey, viewKey, resourceType).ViewLinkRequest(viewLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UnlinkResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlinkResource`: UnlinkResourceSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UnlinkResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**viewKey** | **string** |  | 
**resourceType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **viewLinkRequest** | [**ViewLinkRequest**](ViewLinkRequest.md) | The resource to link to the view. Flags are identified by key. Segments are identified by segment ID. | 

### Return type

[**UnlinkResourceSuccessResponse**](UnlinkResourceSuccessResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateView

> View UpdateView(ctx, projectKey, viewKey).ViewPatch(viewPatch).Execute()

Update view



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
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	viewPatch := *openapiclient.NewViewPatch() // ViewPatch | A JSON representation of the view including only the fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsApi.UpdateView(context.Background(), projectKey, viewKey).ViewPatch(viewPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.UpdateView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewsApi.UpdateView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**viewKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewPatch** | [**ViewPatch**](ViewPatch.md) | A JSON representation of the view including only the fields to update. | 

### Return type

[**View**](View.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

