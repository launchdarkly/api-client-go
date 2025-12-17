# \FeatureFlagsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFeatureFlag**](FeatureFlagsApi.md#CopyFeatureFlag) | **Post** /api/v2/flags/{projectKey}/{featureFlagKey}/copy | Copy feature flag
[**DeleteFeatureFlag**](FeatureFlagsApi.md#DeleteFeatureFlag) | **Delete** /api/v2/flags/{projectKey}/{featureFlagKey} | Delete feature flag
[**GetExpiringContextTargets**](FeatureFlagsApi.md#GetExpiringContextTargets) | **Get** /api/v2/flags/{projectKey}/{featureFlagKey}/expiring-targets/{environmentKey} | Get expiring context targets for feature flag
[**GetExpiringUserTargets**](FeatureFlagsApi.md#GetExpiringUserTargets) | **Get** /api/v2/flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for feature flag
[**GetFeatureFlag**](FeatureFlagsApi.md#GetFeatureFlag) | **Get** /api/v2/flags/{projectKey}/{featureFlagKey} | Get feature flag
[**GetFeatureFlagStatus**](FeatureFlagsApi.md#GetFeatureFlagStatus) | **Get** /api/v2/flag-statuses/{projectKey}/{environmentKey}/{featureFlagKey} | Get feature flag status
[**GetFeatureFlagStatusAcrossEnvironments**](FeatureFlagsApi.md#GetFeatureFlagStatusAcrossEnvironments) | **Get** /api/v2/flag-status/{projectKey}/{featureFlagKey} | Get flag status across environments
[**GetFeatureFlagStatuses**](FeatureFlagsApi.md#GetFeatureFlagStatuses) | **Get** /api/v2/flag-statuses/{projectKey}/{environmentKey} | List feature flag statuses
[**GetFeatureFlags**](FeatureFlagsApi.md#GetFeatureFlags) | **Get** /api/v2/flags/{projectKey} | List feature flags
[**PatchExpiringTargets**](FeatureFlagsApi.md#PatchExpiringTargets) | **Patch** /api/v2/flags/{projectKey}/{featureFlagKey}/expiring-targets/{environmentKey} | Update expiring context targets on feature flag
[**PatchExpiringUserTargets**](FeatureFlagsApi.md#PatchExpiringUserTargets) | **Patch** /api/v2/flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Update expiring user targets on feature flag
[**PatchFeatureFlag**](FeatureFlagsApi.md#PatchFeatureFlag) | **Patch** /api/v2/flags/{projectKey}/{featureFlagKey} | Update feature flag
[**PostFeatureFlag**](FeatureFlagsApi.md#PostFeatureFlag) | **Post** /api/v2/flags/{projectKey} | Create a feature flag
[**PostMigrationSafetyIssues**](FeatureFlagsApi.md#PostMigrationSafetyIssues) | **Post** /api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/migration-safety-issues | Get migration safety issues



## CopyFeatureFlag

> FeatureFlag CopyFeatureFlag(ctx, projectKey, featureFlagKey).FlagCopyConfigPost(flagCopyConfigPost).Execute()

Copy feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key. The key identifies the flag in your code.
	flagCopyConfigPost := *openapiclient.NewFlagCopyConfigPost(*openapiclient.NewFlagCopyConfigEnvironment("Key_example"), *openapiclient.NewFlagCopyConfigEnvironment("Key_example")) // FlagCopyConfigPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.CopyFeatureFlag(context.Background(), projectKey, featureFlagKey).FlagCopyConfigPost(flagCopyConfigPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.CopyFeatureFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyFeatureFlag`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.CopyFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key. The key identifies the flag in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flagCopyConfigPost** | [**FlagCopyConfigPost**](FlagCopyConfigPost.md) |  | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeatureFlag

> DeleteFeatureFlag(ctx, projectKey, featureFlagKey).Execute()

Delete feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key. The key identifies the flag in your code.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeatureFlagsApi.DeleteFeatureFlag(context.Background(), projectKey, featureFlagKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.DeleteFeatureFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key. The key identifies the flag in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureFlagRequest struct via the builder pattern


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


## GetExpiringContextTargets

> ExpiringTargetGetResponse GetExpiringContextTargets(ctx, projectKey, environmentKey, featureFlagKey).Execute()

Get expiring context targets for feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.GetExpiringContextTargets(context.Background(), projectKey, environmentKey, featureFlagKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetExpiringContextTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpiringContextTargets`: ExpiringTargetGetResponse
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetExpiringContextTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpiringContextTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ExpiringTargetGetResponse**](ExpiringTargetGetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpiringUserTargets

> ExpiringUserTargetGetResponse GetExpiringUserTargets(ctx, projectKey, environmentKey, featureFlagKey).Execute()

Get expiring user targets for feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.GetExpiringUserTargets(context.Background(), projectKey, environmentKey, featureFlagKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetExpiringUserTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpiringUserTargets`: ExpiringUserTargetGetResponse
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetExpiringUserTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpiringUserTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ExpiringUserTargetGetResponse**](ExpiringUserTargetGetResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlag

> FeatureFlag GetFeatureFlag(ctx, projectKey, featureFlagKey).Env(env).Expand(expand).Execute()

Get feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key
	env := "env_example" // string | Filter configurations by environment (optional)
	expand := "expand_example" // string | A comma-separated list of fields to expand in the response. Supported fields are explained above. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.GetFeatureFlag(context.Background(), projectKey, featureFlagKey).Env(env).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlag`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **string** | Filter configurations by environment | 
 **expand** | **string** | A comma-separated list of fields to expand in the response. Supported fields are explained above. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagStatus

> FlagStatusRep GetFeatureFlagStatus(ctx, projectKey, environmentKey, featureFlagKey).Execute()

Get feature flag status



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.GetFeatureFlagStatus(context.Background(), projectKey, environmentKey, featureFlagKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlagStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagStatus`: FlagStatusRep
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlagStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlagStatusRep**](FlagStatusRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagStatusAcrossEnvironments

> FeatureFlagStatusAcrossEnvironments GetFeatureFlagStatusAcrossEnvironments(ctx, projectKey, featureFlagKey).Env(env).Execute()

Get flag status across environments



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key
	env := "env_example" // string | Optional environment filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments(context.Background(), projectKey, featureFlagKey).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagStatusAcrossEnvironments`: FeatureFlagStatusAcrossEnvironments
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlagStatusAcrossEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagStatusAcrossEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **string** | Optional environment filter | 

### Return type

[**FeatureFlagStatusAcrossEnvironments**](FeatureFlagStatusAcrossEnvironments.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlagStatuses

> FeatureFlagStatuses GetFeatureFlagStatuses(ctx, projectKey, environmentKey).Execute()

List feature flag statuses



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.GetFeatureFlagStatuses(context.Background(), projectKey, environmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlagStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagStatuses`: FeatureFlagStatuses
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlagStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FeatureFlagStatuses**](FeatureFlagStatuses.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureFlags

> FeatureFlags GetFeatureFlags(ctx, projectKey).Env(env).Tag(tag).Limit(limit).Offset(offset).Archived(archived).Summary(summary).Filter(filter).Sort(sort).Compare(compare).Expand(expand).Execute()

List feature flags



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
	env := "env_example" // string | Filter configurations by environment (optional)
	tag := "tag_example" // string | Filter feature flags by tag (optional)
	limit := int64(789) // int64 | The number of feature flags to return. Defaults to 20. (optional)
	offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	archived := true // bool | Deprecated, use `filter=archived:true` instead. A boolean to filter the list to archived flags. When this is absent, only unarchived flags will be returned (optional)
	summary := true // bool | By default, flags do _not_ include their lists of prerequisites, targets, or rules for each environment. Set `summary=0` and include the `env` query parameter to include these fields for each flag returned. (optional)
	filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form field:value. Read the endpoint description for a full list of available filter fields. (optional)
	sort := "sort_example" // string | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order. Read the endpoint description for a full list of available sort fields. (optional)
	compare := true // bool | Deprecated, unavailable in API version `20240415`. A boolean to filter results by only flags that have differences between environments. (optional)
	expand := "expand_example" // string | A comma-separated list of fields to expand in the response. Supported fields are explained above. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.GetFeatureFlags(context.Background(), projectKey).Env(env).Tag(tag).Limit(limit).Offset(offset).Archived(archived).Summary(summary).Filter(filter).Sort(sort).Compare(compare).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.GetFeatureFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlags`: FeatureFlags
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.GetFeatureFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Filter configurations by environment | 
 **tag** | **string** | Filter feature flags by tag | 
 **limit** | **int64** | The number of feature flags to return. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **archived** | **bool** | Deprecated, use &#x60;filter&#x3D;archived:true&#x60; instead. A boolean to filter the list to archived flags. When this is absent, only unarchived flags will be returned | 
 **summary** | **bool** | By default, flags do _not_ include their lists of prerequisites, targets, or rules for each environment. Set &#x60;summary&#x3D;0&#x60; and include the &#x60;env&#x60; query parameter to include these fields for each flag returned. | 
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form field:value. Read the endpoint description for a full list of available filter fields. | 
 **sort** | **string** | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order. Read the endpoint description for a full list of available sort fields. | 
 **compare** | **bool** | Deprecated, unavailable in API version &#x60;20240415&#x60;. A boolean to filter results by only flags that have differences between environments. | 
 **expand** | **string** | A comma-separated list of fields to expand in the response. Supported fields are explained above. | 

### Return type

[**FeatureFlags**](FeatureFlags.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExpiringTargets

> ExpiringTargetPatchResponse PatchExpiringTargets(ctx, projectKey, environmentKey, featureFlagKey).PatchFlagsRequest(patchFlagsRequest).Execute()

Update expiring context targets on feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key
	patchFlagsRequest := *openapiclient.NewPatchFlagsRequest([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // PatchFlagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.PatchExpiringTargets(context.Background(), projectKey, environmentKey, featureFlagKey).PatchFlagsRequest(patchFlagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PatchExpiringTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExpiringTargets`: ExpiringTargetPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PatchExpiringTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchFlagsRequest** | [**PatchFlagsRequest**](PatchFlagsRequest.md) |  | 

### Return type

[**ExpiringTargetPatchResponse**](ExpiringTargetPatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExpiringUserTargets

> ExpiringUserTargetPatchResponse PatchExpiringUserTargets(ctx, projectKey, environmentKey, featureFlagKey).PatchFlagsRequest(patchFlagsRequest).Execute()

Update expiring user targets on feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key
	patchFlagsRequest := *openapiclient.NewPatchFlagsRequest([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // PatchFlagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.PatchExpiringUserTargets(context.Background(), projectKey, environmentKey, featureFlagKey).PatchFlagsRequest(patchFlagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PatchExpiringUserTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExpiringUserTargets`: ExpiringUserTargetPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PatchExpiringUserTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringUserTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchFlagsRequest** | [**PatchFlagsRequest**](PatchFlagsRequest.md) |  | 

### Return type

[**ExpiringUserTargetPatchResponse**](ExpiringUserTargetPatchResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFeatureFlag

> FeatureFlag PatchFeatureFlag(ctx, projectKey, featureFlagKey).PatchWithComment(patchWithComment).IgnoreConflicts(ignoreConflicts).DryRun(dryRun).Execute()

Update feature flag



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key. The key identifies the flag in your code.
	patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")}) // PatchWithComment | 
	ignoreConflicts := true // bool | If true, the patch will be applied even if it causes a pending scheduled change or approval request to fail. (optional)
	dryRun := true // bool | If true, the patch will be validated but not persisted. Returns a preview of the flag after the patch is applied. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.PatchFeatureFlag(context.Background(), projectKey, featureFlagKey).PatchWithComment(patchWithComment).IgnoreConflicts(ignoreConflicts).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PatchFeatureFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFeatureFlag`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PatchFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key. The key identifies the flag in your code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchWithComment** | [**PatchWithComment**](PatchWithComment.md) |  | 
 **ignoreConflicts** | **bool** | If true, the patch will be applied even if it causes a pending scheduled change or approval request to fail. | 
 **dryRun** | **bool** | If true, the patch will be validated but not persisted. Returns a preview of the flag after the patch is applied. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFeatureFlag

> FeatureFlag PostFeatureFlag(ctx, projectKey).FeatureFlagBody(featureFlagBody).Clone(clone).Execute()

Create a feature flag



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
	featureFlagBody := *openapiclient.NewFeatureFlagBody("My flag", "flag-key-123abc") // FeatureFlagBody | 
	clone := "clone_example" // string | The key of the feature flag to be cloned. The key identifies the flag in your code. For example, setting `clone=flagKey` copies the full targeting configuration for all environments, including `on/off` state, from the original flag to the new flag. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.PostFeatureFlag(context.Background(), projectKey).FeatureFlagBody(featureFlagBody).Clone(clone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PostFeatureFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFeatureFlag`: FeatureFlag
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PostFeatureFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFeatureFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureFlagBody** | [**FeatureFlagBody**](FeatureFlagBody.md) |  | 
 **clone** | **string** | The key of the feature flag to be cloned. The key identifies the flag in your code. For example, setting &#x60;clone&#x3D;flagKey&#x60; copies the full targeting configuration for all environments, including &#x60;on/off&#x60; state, from the original flag to the new flag. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMigrationSafetyIssues

> []MigrationSafetyIssueRep PostMigrationSafetyIssues(ctx, projectKey, flagKey, environmentKey).FlagSempatch(flagSempatch).Execute()

Get migration safety issues



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
	flagKey := "flagKey_example" // string | The migration flag key
	environmentKey := "environmentKey_example" // string | The environment key
	flagSempatch := *openapiclient.NewFlagSempatch([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // FlagSempatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsApi.PostMigrationSafetyIssues(context.Background(), projectKey, flagKey, environmentKey).FlagSempatch(flagSempatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsApi.PostMigrationSafetyIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMigrationSafetyIssues`: []MigrationSafetyIssueRep
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsApi.PostMigrationSafetyIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**flagKey** | **string** | The migration flag key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMigrationSafetyIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **flagSempatch** | [**FlagSempatch**](FlagSempatch.md) |  | 

### Return type

[**[]MigrationSafetyIssueRep**](MigrationSafetyIssueRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

