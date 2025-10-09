# \HoldoutsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllHoldouts**](HoldoutsBetaApi.md#GetAllHoldouts) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/holdouts | Get all holdouts
[**GetHoldout**](HoldoutsBetaApi.md#GetHoldout) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/holdouts/{holdoutKey} | Get holdout
[**GetHoldoutById**](HoldoutsBetaApi.md#GetHoldoutById) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/holdouts/id/{holdoutId} | Get Holdout by Id
[**PatchHoldout**](HoldoutsBetaApi.md#PatchHoldout) | **Patch** /api/v2/projects/{projectKey}/environments/{environmentKey}/holdouts/{holdoutKey} | Patch holdout
[**PostHoldout**](HoldoutsBetaApi.md#PostHoldout) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/holdouts | Create holdout



## GetAllHoldouts

> HoldoutsCollectionRep GetAllHoldouts(ctx, projectKey, environmentKey).Limit(limit).Offset(offset).Execute()

Get all holdouts

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
	limit := int64(789) // int64 | The number of holdouts to return in the response. Defaults to 20 (optional)
	offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an `offset` of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsBetaApi.GetAllHoldouts(context.Background(), projectKey, environmentKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsBetaApi.GetAllHoldouts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllHoldouts`: HoldoutsCollectionRep
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsBetaApi.GetAllHoldouts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllHoldoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int64** | The number of holdouts to return in the response. Defaults to 20 | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an &#x60;offset&#x60; of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**HoldoutsCollectionRep**](HoldoutsCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHoldout

> HoldoutDetailRep GetHoldout(ctx, projectKey, environmentKey, holdoutKey).Expand(expand).Execute()

Get holdout



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
	holdoutKey := "holdoutKey_example" // string | The holdout experiment key
	expand := "expand_example" // string | A comma-separated list of properties that can reveal additional information in the response. Supported fields are explained above. Holdout experiment expansion fields have no prefix. Related experiment expansion fields have `rel-` as a prefix. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsBetaApi.GetHoldout(context.Background(), projectKey, environmentKey, holdoutKey).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsBetaApi.GetHoldout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHoldout`: HoldoutDetailRep
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsBetaApi.GetHoldout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**holdoutKey** | **string** | The holdout experiment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** | A comma-separated list of properties that can reveal additional information in the response. Supported fields are explained above. Holdout experiment expansion fields have no prefix. Related experiment expansion fields have &#x60;rel-&#x60; as a prefix. | 

### Return type

[**HoldoutDetailRep**](HoldoutDetailRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHoldoutById

> HoldoutRep GetHoldoutById(ctx, projectKey, environmentKey, holdoutId).Execute()

Get Holdout by Id

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
	holdoutId := "holdoutId_example" // string | The holdout experiment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsBetaApi.GetHoldoutById(context.Background(), projectKey, environmentKey, holdoutId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsBetaApi.GetHoldoutById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHoldoutById`: HoldoutRep
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsBetaApi.GetHoldoutById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**holdoutId** | **string** | The holdout experiment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHoldoutByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**HoldoutRep**](HoldoutRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchHoldout

> HoldoutRep PatchHoldout(ctx, projectKey, environmentKey, holdoutKey).HoldoutPatchInput(holdoutPatchInput).Execute()

Patch holdout



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
	holdoutKey := "holdoutKey_example" // string | The holdout key
	holdoutPatchInput := *openapiclient.NewHoldoutPatchInput([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // HoldoutPatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsBetaApi.PatchHoldout(context.Background(), projectKey, environmentKey, holdoutKey).HoldoutPatchInput(holdoutPatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsBetaApi.PatchHoldout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchHoldout`: HoldoutRep
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsBetaApi.PatchHoldout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**holdoutKey** | **string** | The holdout key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchHoldoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **holdoutPatchInput** | [**HoldoutPatchInput**](HoldoutPatchInput.md) |  | 

### Return type

[**HoldoutRep**](HoldoutRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostHoldout

> HoldoutRep PostHoldout(ctx, projectKey, environmentKey).HoldoutPostRequest(holdoutPostRequest).Execute()

Create holdout



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
	holdoutPostRequest := *openapiclient.NewHoldoutPostRequest() // HoldoutPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HoldoutsBetaApi.PostHoldout(context.Background(), projectKey, environmentKey).HoldoutPostRequest(holdoutPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HoldoutsBetaApi.PostHoldout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostHoldout`: HoldoutRep
	fmt.Fprintf(os.Stdout, "Response from `HoldoutsBetaApi.PostHoldout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostHoldoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **holdoutPostRequest** | [**HoldoutPostRequest**](HoldoutPostRequest.md) |  | 

### Return type

[**HoldoutRep**](HoldoutRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

