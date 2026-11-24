# \EnvironmentReleasePlansBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDefaultEnvironmentReleasePlan**](EnvironmentReleasePlansBetaApi.md#DeleteDefaultEnvironmentReleasePlan) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey}/default-release-plan | Clear the default release plan
[**DeleteEnvironmentReleasePlan**](EnvironmentReleasePlansBetaApi.md#DeleteEnvironmentReleasePlan) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey}/release-plans/{releasePlanKey} | Delete a release plan
[**GetDefaultEnvironmentReleasePlan**](EnvironmentReleasePlansBetaApi.md#GetDefaultEnvironmentReleasePlan) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/default-release-plan | Get the default release plan
[**GetEnvironmentReleasePlan**](EnvironmentReleasePlansBetaApi.md#GetEnvironmentReleasePlan) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/release-plans/{releasePlanKey} | Get a release plan
[**GetEnvironmentReleasePlans**](EnvironmentReleasePlansBetaApi.md#GetEnvironmentReleasePlans) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/release-plans | List release plans
[**PostEnvironmentReleasePlan**](EnvironmentReleasePlansBetaApi.md#PostEnvironmentReleasePlan) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/release-plans | Create a release plan
[**PutDefaultEnvironmentReleasePlan**](EnvironmentReleasePlansBetaApi.md#PutDefaultEnvironmentReleasePlan) | **Put** /api/v2/projects/{projectKey}/environments/{environmentKey}/default-release-plan | Set the default release plan
[**PutEnvironmentReleasePlan**](EnvironmentReleasePlansBetaApi.md#PutEnvironmentReleasePlan) | **Put** /api/v2/projects/{projectKey}/environments/{environmentKey}/release-plans/{releasePlanKey} | Update a release plan



## DeleteDefaultEnvironmentReleasePlan

> DeleteDefaultEnvironmentReleasePlan(ctx, projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Execute()

Clear the default release plan



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentReleasePlansBetaApi.DeleteDefaultEnvironmentReleasePlan(context.Background(), projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.DeleteDefaultEnvironmentReleasePlan``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultEnvironmentReleasePlanRequest struct via the builder pattern


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


## DeleteEnvironmentReleasePlan

> DeleteEnvironmentReleasePlan(ctx, projectKey, environmentKey, releasePlanKey).LDAPIVersion(lDAPIVersion).Execute()

Delete a release plan



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key
	releasePlanKey := "checkout-flow-rollout" // string | The release plan key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentReleasePlansBetaApi.DeleteEnvironmentReleasePlan(context.Background(), projectKey, environmentKey, releasePlanKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.DeleteEnvironmentReleasePlan``: %v\n", err)
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
**releasePlanKey** | **string** | The release plan key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentReleasePlanRequest struct via the builder pattern


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


## GetDefaultEnvironmentReleasePlan

> EnvironmentReleasePlan GetDefaultEnvironmentReleasePlan(ctx, projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Execute()

Get the default release plan



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasePlansBetaApi.GetDefaultEnvironmentReleasePlan(context.Background(), projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.GetDefaultEnvironmentReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultEnvironmentReleasePlan`: EnvironmentReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasePlansBetaApi.GetDefaultEnvironmentReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultEnvironmentReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



### Return type

[**EnvironmentReleasePlan**](EnvironmentReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentReleasePlan

> EnvironmentReleasePlan GetEnvironmentReleasePlan(ctx, projectKey, environmentKey, releasePlanKey).LDAPIVersion(lDAPIVersion).Execute()

Get a release plan



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key
	releasePlanKey := "checkout-flow-rollout" // string | The release plan key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasePlansBetaApi.GetEnvironmentReleasePlan(context.Background(), projectKey, environmentKey, releasePlanKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.GetEnvironmentReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentReleasePlan`: EnvironmentReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasePlansBetaApi.GetEnvironmentReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**releasePlanKey** | **string** | The release plan key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 




### Return type

[**EnvironmentReleasePlan**](EnvironmentReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentReleasePlans

> EnvironmentReleasePlans GetEnvironmentReleasePlans(ctx, projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Filter(filter).Limit(limit).Offset(offset).Execute()

List release plans



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key
	filter := "query contains "checkout"" // string | A queryfilter expression that narrows the list by text match. Supported field and operators: `query` (`equals`, `contains`). When present, only release plans whose name or key contains the text, case-insensitively, are returned, and totalCount counts the matches. Surrounding whitespace is ignored, and the text is at most 256 bytes. Omitting the filter lists every release plan in the environment. (optional)
	limit := int32(56) // int32 | The number of release plans to return. Defaults to 20. Maximum is 100. (optional)
	offset := int32(56) // int32 | The number of release plans to skip. Used for pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasePlansBetaApi.GetEnvironmentReleasePlans(context.Background(), projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Filter(filter).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.GetEnvironmentReleasePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentReleasePlans`: EnvironmentReleasePlans
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasePlansBetaApi.GetEnvironmentReleasePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentReleasePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **filter** | **string** | A queryfilter expression that narrows the list by text match. Supported field and operators: &#x60;query&#x60; (&#x60;equals&#x60;, &#x60;contains&#x60;). When present, only release plans whose name or key contains the text, case-insensitively, are returned, and totalCount counts the matches. Surrounding whitespace is ignored, and the text is at most 256 bytes. Omitting the filter lists every release plan in the environment. | 
 **limit** | **int32** | The number of release plans to return. Defaults to 20. Maximum is 100. | 
 **offset** | **int32** | The number of release plans to skip. Used for pagination. | 

### Return type

[**EnvironmentReleasePlans**](EnvironmentReleasePlans.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnvironmentReleasePlan

> EnvironmentReleasePlan PostEnvironmentReleasePlan(ctx, projectKey, environmentKey).LDAPIVersion(lDAPIVersion).PostEnvironmentReleasePlanRequest(postEnvironmentReleasePlanRequest).Execute()

Create a release plan



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key
	postEnvironmentReleasePlanRequest := *openapiclient.NewPostEnvironmentReleasePlanRequest("Checkout flow rollout", "checkout-flow-rollout", *openapiclient.NewEnvironmentReleasePlanProgressiveRelease("user", []openapiclient.EnvironmentReleasePlanPhase{*openapiclient.NewEnvironmentReleasePlanPhase([]openapiclient.EnvironmentReleasePlanStep{*openapiclient.NewEnvironmentReleasePlanStep(int64(50000), int64(3600000))})})) // PostEnvironmentReleasePlanRequest | Release plan to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasePlansBetaApi.PostEnvironmentReleasePlan(context.Background(), projectKey, environmentKey).LDAPIVersion(lDAPIVersion).PostEnvironmentReleasePlanRequest(postEnvironmentReleasePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.PostEnvironmentReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEnvironmentReleasePlan`: EnvironmentReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasePlansBetaApi.PostEnvironmentReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEnvironmentReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **postEnvironmentReleasePlanRequest** | [**PostEnvironmentReleasePlanRequest**](PostEnvironmentReleasePlanRequest.md) | Release plan to create | 

### Return type

[**EnvironmentReleasePlan**](EnvironmentReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDefaultEnvironmentReleasePlan

> EnvironmentReleasePlan PutDefaultEnvironmentReleasePlan(ctx, projectKey, environmentKey).LDAPIVersion(lDAPIVersion).PutDefaultEnvironmentReleasePlanRequest(putDefaultEnvironmentReleasePlanRequest).Execute()

Set the default release plan



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key
	putDefaultEnvironmentReleasePlanRequest := *openapiclient.NewPutDefaultEnvironmentReleasePlanRequest("checkout-flow-rollout") // PutDefaultEnvironmentReleasePlanRequest | The plan to set as the environment's default

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasePlansBetaApi.PutDefaultEnvironmentReleasePlan(context.Background(), projectKey, environmentKey).LDAPIVersion(lDAPIVersion).PutDefaultEnvironmentReleasePlanRequest(putDefaultEnvironmentReleasePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.PutDefaultEnvironmentReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDefaultEnvironmentReleasePlan`: EnvironmentReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasePlansBetaApi.PutDefaultEnvironmentReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDefaultEnvironmentReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **putDefaultEnvironmentReleasePlanRequest** | [**PutDefaultEnvironmentReleasePlanRequest**](PutDefaultEnvironmentReleasePlanRequest.md) | The plan to set as the environment&#39;s default | 

### Return type

[**EnvironmentReleasePlan**](EnvironmentReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEnvironmentReleasePlan

> EnvironmentReleasePlan PutEnvironmentReleasePlan(ctx, projectKey, environmentKey, releasePlanKey).LDAPIVersion(lDAPIVersion).PutEnvironmentReleasePlanRequest(putEnvironmentReleasePlanRequest).Execute()

Update a release plan



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
	projectKey := "default" // string | The project key
	environmentKey := "production" // string | The environment key
	releasePlanKey := "checkout-flow-rollout" // string | The release plan key
	putEnvironmentReleasePlanRequest := *openapiclient.NewPutEnvironmentReleasePlanRequest("Checkout flow rollout", *openapiclient.NewEnvironmentReleasePlanProgressiveRelease("user", []openapiclient.EnvironmentReleasePlanPhase{*openapiclient.NewEnvironmentReleasePlanPhase([]openapiclient.EnvironmentReleasePlanStep{*openapiclient.NewEnvironmentReleasePlanStep(int64(50000), int64(3600000))})})) // PutEnvironmentReleasePlanRequest | The release plan's replacement configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasePlansBetaApi.PutEnvironmentReleasePlan(context.Background(), projectKey, environmentKey, releasePlanKey).LDAPIVersion(lDAPIVersion).PutEnvironmentReleasePlanRequest(putEnvironmentReleasePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasePlansBetaApi.PutEnvironmentReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEnvironmentReleasePlan`: EnvironmentReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasePlansBetaApi.PutEnvironmentReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**releasePlanKey** | **string** | The release plan key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutEnvironmentReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



 **putEnvironmentReleasePlanRequest** | [**PutEnvironmentReleasePlanRequest**](PutEnvironmentReleasePlanRequest.md) | The release plan&#39;s replacement configuration | 

### Return type

[**EnvironmentReleasePlan**](EnvironmentReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

