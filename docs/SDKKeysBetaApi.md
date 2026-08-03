# \SDKKeysBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSdkKeyByKey**](SDKKeysBetaApi.md#DeleteSdkKeyByKey) | **Delete** /api/v2/projects/{projectKey}/environments/{environmentKey}/sdk-keys/{sdkKeyKey} | Delete SDK key
[**GetSdkKeyByKey**](SDKKeysBetaApi.md#GetSdkKeyByKey) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/sdk-keys/{sdkKeyKey} | Get SDK key
[**GetSdkKeys**](SDKKeysBetaApi.md#GetSdkKeys) | **Get** /api/v2/projects/{projectKey}/environments/{environmentKey}/sdk-keys | Get all environment SDK keys
[**PatchSdkKeyByKey**](SDKKeysBetaApi.md#PatchSdkKeyByKey) | **Patch** /api/v2/projects/{projectKey}/environments/{environmentKey}/sdk-keys/{sdkKeyKey} | Update SDK key
[**PostSdkKey**](SDKKeysBetaApi.md#PostSdkKey) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/sdk-keys | Create SDK key



## DeleteSdkKeyByKey

> DeleteSdkKeyByKey(ctx, projectKey, environmentKey, sdkKeyKey).LDAPIVersion(lDAPIVersion).Execute()

Delete SDK key



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
	environmentKey := "production" // string | 
	sdkKeyKey := "my-sdk-key" // string | The user-defined identifying key of the SDK key. This is used solely to identify an SDK key and is distinct from the value field, which is the actual SDK key value.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SDKKeysBetaApi.DeleteSdkKeyByKey(context.Background(), projectKey, environmentKey, sdkKeyKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDKKeysBetaApi.DeleteSdkKeyByKey``: %v\n", err)
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
**sdkKeyKey** | **string** | The user-defined identifying key of the SDK key. This is used solely to identify an SDK key and is distinct from the value field, which is the actual SDK key value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdkKeyByKeyRequest struct via the builder pattern


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


## GetSdkKeyByKey

> SdkKey GetSdkKeyByKey(ctx, projectKey, environmentKey, sdkKeyKey).LDAPIVersion(lDAPIVersion).Execute()

Get SDK key



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
	environmentKey := "production" // string | 
	sdkKeyKey := "my-sdk-key" // string | The user-defined identifying key of the SDK key. This is used solely to identify an SDK key and is distinct from the value field, which is the actual SDK key value.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDKKeysBetaApi.GetSdkKeyByKey(context.Background(), projectKey, environmentKey, sdkKeyKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDKKeysBetaApi.GetSdkKeyByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSdkKeyByKey`: SdkKey
	fmt.Fprintf(os.Stdout, "Response from `SDKKeysBetaApi.GetSdkKeyByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 
**sdkKeyKey** | **string** | The user-defined identifying key of the SDK key. This is used solely to identify an SDK key and is distinct from the value field, which is the actual SDK key value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdkKeyByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 




### Return type

[**SdkKey**](SdkKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdkKeys

> SdkKeysForGetSdkKeys GetSdkKeys(ctx, projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Filter(filter).Limit(limit).Offset(offset).Sort(sort).Execute()

Get all environment SDK keys



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
	environmentKey := "production" // string | 
	filter := "filter_example" // string | A filter to apply to the list of SDK keys. Supports the following fields and operators: `kind` (equals \"sdk\" or \"mobile\"), `active` (equals true or false). Format: `field:value`. Example: `kind:sdk`. (optional)
	limit := int32(56) // int32 | The number of SDK keys to return. Defaults to 20. Maximum is 100. (optional)
	offset := int32(56) // int32 | The number of SDK keys to skip. Used for pagination. (optional)
	sort := "sort_example" // string | A sort to apply to the list of SDK keys. Supported field: `createdAt`. Prefix the field with `-` to sort in descending order (for example, `-createdAt`). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDKKeysBetaApi.GetSdkKeys(context.Background(), projectKey, environmentKey).LDAPIVersion(lDAPIVersion).Filter(filter).Limit(limit).Offset(offset).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDKKeysBetaApi.GetSdkKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSdkKeys`: SdkKeysForGetSdkKeys
	fmt.Fprintf(os.Stdout, "Response from `SDKKeysBetaApi.GetSdkKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdkKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **filter** | **string** | A filter to apply to the list of SDK keys. Supports the following fields and operators: &#x60;kind&#x60; (equals \&quot;sdk\&quot; or \&quot;mobile\&quot;), &#x60;active&#x60; (equals true or false). Format: &#x60;field:value&#x60;. Example: &#x60;kind:sdk&#x60;. | 
 **limit** | **int32** | The number of SDK keys to return. Defaults to 20. Maximum is 100. | 
 **offset** | **int32** | The number of SDK keys to skip. Used for pagination. | 
 **sort** | **string** | A sort to apply to the list of SDK keys. Supported field: &#x60;createdAt&#x60;. Prefix the field with &#x60;-&#x60; to sort in descending order (for example, &#x60;-createdAt&#x60;). | 

### Return type

[**SdkKeysForGetSdkKeys**](SdkKeysForGetSdkKeys.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSdkKeyByKey

> SdkKey PatchSdkKeyByKey(ctx, projectKey, environmentKey, sdkKeyKey).LDAPIVersion(lDAPIVersion).SdkKeyPatch(sdkKeyPatch).Version(version).Execute()

Update SDK key



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
	environmentKey := "production" // string | 
	sdkKeyKey := "my-sdk-key" // string | The user-defined identifying key of the SDK key. This is used solely to identify an SDK key and is distinct from the value field, which is the actual SDK key value.
	sdkKeyPatch := *openapiclient.NewSdkKeyPatch() // SdkKeyPatch | An array of patches for updating an existing SDK key. The following fields are supported: `name`, `description`, `expiry`.
	version := int32(56) // int32 | The version of the SDK key for optimistic locking. If provided, the update will only succeed if the current version matches. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDKKeysBetaApi.PatchSdkKeyByKey(context.Background(), projectKey, environmentKey, sdkKeyKey).LDAPIVersion(lDAPIVersion).SdkKeyPatch(sdkKeyPatch).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDKKeysBetaApi.PatchSdkKeyByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSdkKeyByKey`: SdkKey
	fmt.Fprintf(os.Stdout, "Response from `SDKKeysBetaApi.PatchSdkKeyByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 
**sdkKeyKey** | **string** | The user-defined identifying key of the SDK key. This is used solely to identify an SDK key and is distinct from the value field, which is the actual SDK key value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSdkKeyByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



 **sdkKeyPatch** | [**SdkKeyPatch**](SdkKeyPatch.md) | An array of patches for updating an existing SDK key. The following fields are supported: &#x60;name&#x60;, &#x60;description&#x60;, &#x60;expiry&#x60;. | 
 **version** | **int32** | The version of the SDK key for optimistic locking. If provided, the update will only succeed if the current version matches. | 

### Return type

[**SdkKey**](SdkKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSdkKey

> SdkKey PostSdkKey(ctx, projectKey, environmentKey).LDAPIVersion(lDAPIVersion).SdkKeyPost(sdkKeyPost).Execute()

Create SDK key



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
	environmentKey := "production" // string | 
	sdkKeyPost := *openapiclient.NewSdkKeyPost("Key_example", "Name_example") // SdkKeyPost | Parameters for creating a new SDK key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SDKKeysBetaApi.PostSdkKey(context.Background(), projectKey, environmentKey).LDAPIVersion(lDAPIVersion).SdkKeyPost(sdkKeyPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDKKeysBetaApi.PostSdkKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSdkKey`: SdkKey
	fmt.Fprintf(os.Stdout, "Response from `SDKKeysBetaApi.PostSdkKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**environmentKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSdkKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **sdkKeyPost** | [**SdkKeyPost**](SdkKeyPost.md) | Parameters for creating a new SDK key | 

### Return type

[**SdkKey**](SdkKey.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

