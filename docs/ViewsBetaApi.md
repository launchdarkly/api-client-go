# \ViewsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateView**](ViewsBetaApi.md#CreateView) | **Post** /api/v2/projects/{projectKey}/views | Create view
[**DeleteView**](ViewsBetaApi.md#DeleteView) | **Delete** /api/v2/projects/{projectKey}/views/{viewKey} | Delete view
[**GetLinkedResources**](ViewsBetaApi.md#GetLinkedResources) | **Get** /api/v2/projects/{projectKey}/views/{viewKey}/linked/{resourceType} | Get linked resources
[**GetLinkedViews**](ViewsBetaApi.md#GetLinkedViews) | **Get** /api/v2/projects/{projectKey}/view-associations/{resourceType}/{resourceKey} | Get linked views for a given resource
[**GetView**](ViewsBetaApi.md#GetView) | **Get** /api/v2/projects/{projectKey}/views/{viewKey} | Get view
[**GetViews**](ViewsBetaApi.md#GetViews) | **Get** /api/v2/projects/{projectKey}/views | List views
[**LinkResource**](ViewsBetaApi.md#LinkResource) | **Post** /api/v2/projects/{projectKey}/views/{viewKey}/link/{resourceType} | Link resource
[**UnlinkResource**](ViewsBetaApi.md#UnlinkResource) | **Delete** /api/v2/projects/{projectKey}/views/{viewKey}/link/{resourceType} | Unlink resource
[**UpdateView**](ViewsBetaApi.md#UpdateView) | **Patch** /api/v2/projects/{projectKey}/views/{viewKey} | Update view



## CreateView

> View CreateView(ctx, projectKey).LDAPIVersion(lDAPIVersion).ViewPost(viewPost).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	viewPost := *openapiclient.NewViewPost("Key_example", "Name_example") // ViewPost | View object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.CreateView(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).ViewPost(viewPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.CreateView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.CreateView`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

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

> DeleteView(ctx, projectKey, viewKey).LDAPIVersion(lDAPIVersion).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ViewsBetaApi.DeleteView(context.Background(), projectKey, viewKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.DeleteView``: %v\n", err)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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

> ViewLinkedResources GetLinkedResources(ctx, projectKey, viewKey, resourceType).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Sort(sort).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	resourceType := "flags" // string | 
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	sort := "sort_example" // string | Field to sort by. Default field is `linkedAt`, default order is ascending. (optional) (default to "linkedAt")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.GetLinkedResources(context.Background(), projectKey, viewKey, resourceType).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.GetLinkedResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedResources`: ViewLinkedResources
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.GetLinkedResources`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



 **limit** | **int32** | The number of views to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **sort** | **string** | Field to sort by. Default field is &#x60;linkedAt&#x60;, default order is ascending. | [default to &quot;linkedAt&quot;]

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

> Views GetLinkedViews(ctx, projectKey, resourceType, resourceKey).LDAPIVersion(lDAPIVersion).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	resourceType := "flags" // string | 
	resourceKey := "my-flag" // string | 
	environmentId := "6890ff25c3e3830ba1a352e4" // string | Environment ID. Required when resourceType is 'segments' (optional)
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.GetLinkedViews(context.Background(), projectKey, resourceType, resourceKey).LDAPIVersion(lDAPIVersion).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.GetLinkedViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedViews`: Views
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.GetLinkedViews`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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

> View GetView(ctx, projectKey, viewKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	sort := "sort_example" // string | A sort to apply to the list of views. (optional)
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of views. (optional)
	expand := []string{"Expand_example"} // []string | A comma-separated list of fields to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.GetView(context.Background(), projectKey, viewKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.GetView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.GetView`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **sort** | **string** | A sort to apply to the list of views. | 
 **limit** | **int32** | The number of views to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list of views. | 
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

> Views GetViews(ctx, projectKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	sort := "sort_example" // string | A sort to apply to the list of views. (optional)
	limit := int32(56) // int32 | The number of views to return. (optional)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	filter := "filter_example" // string | A filter to apply to the list of views. (optional)
	expand := []string{"Expand_example"} // []string | A comma-separated list of fields to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.GetViews(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Sort(sort).Limit(limit).Offset(offset).Filter(filter).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.GetViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViews`: Views
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.GetViews`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **sort** | **string** | A sort to apply to the list of views. | 
 **limit** | **int32** | The number of views to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **filter** | **string** | A filter to apply to the list of views. | 
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

> LinkResourceSuccessResponse LinkResource(ctx, projectKey, viewKey, resourceType).LDAPIVersion(lDAPIVersion).ViewLinkRequest(viewLinkRequest).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	resourceType := "flags" // string | 
	viewLinkRequest := openapiclient.ViewLinkRequest{ViewLinkRequestKeys: openapiclient.NewViewLinkRequestKeys([]string{"Keys_example"})} // ViewLinkRequest | The resource to link to the view. Flags are identified by key. Segments are identified by segment ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.LinkResource(context.Background(), projectKey, viewKey, resourceType).LDAPIVersion(lDAPIVersion).ViewLinkRequest(viewLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.LinkResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkResource`: LinkResourceSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.LinkResource`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



 **viewLinkRequest** | [**ViewLinkRequest**](ViewLinkRequest.md) | The resource to link to the view. Flags are identified by key. Segments are identified by segment ID. | 

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

> UnlinkResourceSuccessResponse UnlinkResource(ctx, projectKey, viewKey, resourceType).LDAPIVersion(lDAPIVersion).ViewLinkRequest(viewLinkRequest).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	resourceType := "flags" // string | 
	viewLinkRequest := openapiclient.ViewLinkRequest{ViewLinkRequestKeys: openapiclient.NewViewLinkRequestKeys([]string{"Keys_example"})} // ViewLinkRequest | The resource to link to the view. Flags are identified by key. Segments are identified by segment ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.UnlinkResource(context.Background(), projectKey, viewKey, resourceType).LDAPIVersion(lDAPIVersion).ViewLinkRequest(viewLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.UnlinkResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlinkResource`: UnlinkResourceSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.UnlinkResource`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 



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

> View UpdateView(ctx, projectKey, viewKey).LDAPIVersion(lDAPIVersion).ViewPatch(viewPatch).Execute()

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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	viewKey := "my-view" // string | 
	viewPatch := *openapiclient.NewViewPatch() // ViewPatch | A JSON representation of the view including only the fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ViewsBetaApi.UpdateView(context.Background(), projectKey, viewKey).LDAPIVersion(lDAPIVersion).ViewPatch(viewPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ViewsBetaApi.UpdateView``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateView`: View
	fmt.Fprintf(os.Stdout, "Response from `ViewsBetaApi.UpdateView`: %v\n", resp)
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
 **lDAPIVersion** | **string** | Version of the endpoint. | 


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

