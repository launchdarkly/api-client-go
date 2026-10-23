# \AdaptiveTriggersApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdaptiveTrigger**](AdaptiveTriggersApi.md#DeleteAdaptiveTrigger) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey}/adaptive-triggers/{triggerID} | Delete an adaptive trigger
[**GetAdaptiveTrigger**](AdaptiveTriggersApi.md#GetAdaptiveTrigger) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/adaptive-triggers/{triggerID} | Get an adaptive trigger
[**ListAdaptiveTriggers**](AdaptiveTriggersApi.md#ListAdaptiveTriggers) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/adaptive-triggers | List adaptive triggers
[**PatchAdaptiveTrigger**](AdaptiveTriggersApi.md#PatchAdaptiveTrigger) | **Patch** /api/v2/projects/{projectKey}/environments/{environmentKey}/adaptive-triggers/{triggerID} | Update an adaptive trigger
[**PostAdaptiveTrigger**](AdaptiveTriggersApi.md#PostAdaptiveTrigger) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/adaptive-triggers | Create an adaptive trigger
[**PostDisableAdaptiveTrigger**](AdaptiveTriggersApi.md#PostDisableAdaptiveTrigger) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/adaptive-triggers/{triggerID}/disable | Disable an adaptive trigger
[**PostEnableAdaptiveTrigger**](AdaptiveTriggersApi.md#PostEnableAdaptiveTrigger) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/adaptive-triggers/{triggerID}/enable | Enable an adaptive trigger



## DeleteAdaptiveTrigger

> DeleteAdaptiveTrigger(ctx, projectKey, environmentKey, triggerID).Execute()

Delete an adaptive trigger



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
	projectKey := "projectKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	triggerID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdaptiveTriggersApi.DeleteAdaptiveTrigger(context.Background(), projectKey, environmentKey, triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveTriggersApi.DeleteAdaptiveTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 
**triggerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdaptiveTriggerRequest struct via the builder pattern


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


## GetAdaptiveTrigger

> AdaptiveTrigger GetAdaptiveTrigger(ctx, projectKey, environmentKey, triggerID).Execute()

Get an adaptive trigger



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
	projectKey := "projectKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	triggerID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveTriggersApi.GetAdaptiveTrigger(context.Background(), projectKey, environmentKey, triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveTriggersApi.GetAdaptiveTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdaptiveTrigger`: AdaptiveTrigger
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveTriggersApi.GetAdaptiveTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 
**triggerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdaptiveTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AdaptiveTrigger**](AdaptiveTrigger.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAdaptiveTriggers

> AdaptiveTriggers ListAdaptiveTriggers(ctx, projectKey, environmentKey).ResourceKind(resourceKind).ResourceKey(resourceKey).Limit(limit).Offset(offset).Execute()

List adaptive triggers



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
	projectKey := "projectKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	resourceKind := openapiclient.AdaptiveTriggerResourceKind("flag") // AdaptiveTriggerResourceKind | 
	resourceKey := "resourceKey_example" // string | 
	limit := int32(56) // int32 | The number of items to return. (optional) (default to 25)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveTriggersApi.ListAdaptiveTriggers(context.Background(), projectKey, environmentKey).ResourceKind(resourceKind).ResourceKey(resourceKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveTriggersApi.ListAdaptiveTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdaptiveTriggers`: AdaptiveTriggers
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveTriggersApi.ListAdaptiveTriggers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAdaptiveTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceKind** | [**AdaptiveTriggerResourceKind**](AdaptiveTriggerResourceKind.md) |  | 
 **resourceKey** | **string** |  | 
 **limit** | **int32** | The number of items to return. | [default to 25]
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | [default to 0]

### Return type

[**AdaptiveTriggers**](AdaptiveTriggers.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAdaptiveTrigger

> AdaptiveTrigger PatchAdaptiveTrigger(ctx, projectKey, environmentKey, triggerID).UpdateAdaptiveTriggerRequest(updateAdaptiveTriggerRequest).Execute()

Update an adaptive trigger



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
	projectKey := "projectKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	triggerID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateAdaptiveTriggerRequest := *openapiclient.NewUpdateAdaptiveTriggerRequest() // UpdateAdaptiveTriggerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveTriggersApi.PatchAdaptiveTrigger(context.Background(), projectKey, environmentKey, triggerID).UpdateAdaptiveTriggerRequest(updateAdaptiveTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveTriggersApi.PatchAdaptiveTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAdaptiveTrigger`: AdaptiveTrigger
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveTriggersApi.PatchAdaptiveTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 
**triggerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAdaptiveTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateAdaptiveTriggerRequest** | [**UpdateAdaptiveTriggerRequest**](UpdateAdaptiveTriggerRequest.md) |  | 

### Return type

[**AdaptiveTrigger**](AdaptiveTrigger.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAdaptiveTrigger

> AdaptiveTrigger PostAdaptiveTrigger(ctx, projectKey, environmentKey).CreateAdaptiveTriggerRequest(createAdaptiveTriggerRequest).Execute()

Create an adaptive trigger



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
	projectKey := "projectKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	createAdaptiveTriggerRequest := *openapiclient.NewCreateAdaptiveTriggerRequest("Name_example", openapiclient.AdaptiveTriggerResourceKind("flag"), "ResourceKey_example", *openapiclient.NewObservabilityAlertSource("Type_example", int32(123)), *openapiclient.NewAdaptiveTriggerTask([]openapiclient.AdaptiveTriggerInstruction{*openapiclient.NewAdaptiveTriggerInstruction("Kind_example", "VariationId_example")})) // CreateAdaptiveTriggerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdaptiveTriggersApi.PostAdaptiveTrigger(context.Background(), projectKey, environmentKey).CreateAdaptiveTriggerRequest(createAdaptiveTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveTriggersApi.PostAdaptiveTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAdaptiveTrigger`: AdaptiveTrigger
	fmt.Fprintf(os.Stdout, "Response from `AdaptiveTriggersApi.PostAdaptiveTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAdaptiveTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createAdaptiveTriggerRequest** | [**CreateAdaptiveTriggerRequest**](CreateAdaptiveTriggerRequest.md) |  | 

### Return type

[**AdaptiveTrigger**](AdaptiveTrigger.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDisableAdaptiveTrigger

> PostDisableAdaptiveTrigger(ctx, projectKey, environmentKey, triggerID).Execute()

Disable an adaptive trigger



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
	projectKey := "projectKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	triggerID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdaptiveTriggersApi.PostDisableAdaptiveTrigger(context.Background(), projectKey, environmentKey, triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveTriggersApi.PostDisableAdaptiveTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 
**triggerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDisableAdaptiveTriggerRequest struct via the builder pattern


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


## PostEnableAdaptiveTrigger

> PostEnableAdaptiveTrigger(ctx, projectKey, environmentKey, triggerID).Execute()

Enable an adaptive trigger



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
	projectKey := "projectKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	triggerID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdaptiveTriggersApi.PostEnableAdaptiveTrigger(context.Background(), projectKey, environmentKey, triggerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdaptiveTriggersApi.PostEnableAdaptiveTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 
**triggerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEnableAdaptiveTriggerRequest struct via the builder pattern


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

