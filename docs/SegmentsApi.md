# \SegmentsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBigSegmentExport**](SegmentsApi.md#CreateBigSegmentExport) | **Post** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/exports | Create big segment export
[**CreateBigSegmentImport**](SegmentsApi.md#CreateBigSegmentImport) | **Post** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/imports | Create big segment import
[**DeleteSegment**](SegmentsApi.md#DeleteSegment) | **Delete** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey} | Delete segment
[**GetBigSegmentExport**](SegmentsApi.md#GetBigSegmentExport) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/exports/{exportID} | Get big segment export
[**GetBigSegmentImport**](SegmentsApi.md#GetBigSegmentImport) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/imports/{importID} | Get big segment import
[**GetContextInstanceSegmentsMembershipByEnv**](SegmentsApi.md#GetContextInstanceSegmentsMembershipByEnv) | **Post** /api/v2/projects/{projectKey}/environments/{environmentKey}/segments/evaluate | List segment memberships for context instance
[**GetExpiringTargetsForSegment**](SegmentsApi.md#GetExpiringTargetsForSegment) | **Get** /api/v2/segments/{projectKey}/{segmentKey}/expiring-targets/{environmentKey} | Get expiring targets for segment
[**GetExpiringUserTargetsForSegment**](SegmentsApi.md#GetExpiringUserTargetsForSegment) | **Get** /api/v2/segments/{projectKey}/{segmentKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for segment
[**GetSegment**](SegmentsApi.md#GetSegment) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey} | Get segment
[**GetSegmentMembershipForContext**](SegmentsApi.md#GetSegmentMembershipForContext) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/contexts/{contextKey} | Get big segment membership for context
[**GetSegmentMembershipForUser**](SegmentsApi.md#GetSegmentMembershipForUser) | **Get** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/users/{userKey} | Get big segment membership for user
[**GetSegments**](SegmentsApi.md#GetSegments) | **Get** /api/v2/segments/{projectKey}/{environmentKey} | List segments
[**PatchExpiringTargetsForSegment**](SegmentsApi.md#PatchExpiringTargetsForSegment) | **Patch** /api/v2/segments/{projectKey}/{segmentKey}/expiring-targets/{environmentKey} | Update expiring targets for segment
[**PatchExpiringUserTargetsForSegment**](SegmentsApi.md#PatchExpiringUserTargetsForSegment) | **Patch** /api/v2/segments/{projectKey}/{segmentKey}/expiring-user-targets/{environmentKey} | Update expiring user targets for segment
[**PatchSegment**](SegmentsApi.md#PatchSegment) | **Patch** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey} | Patch segment
[**PostSegment**](SegmentsApi.md#PostSegment) | **Post** /api/v2/segments/{projectKey}/{environmentKey} | Create segment
[**UpdateBigSegmentContextTargets**](SegmentsApi.md#UpdateBigSegmentContextTargets) | **Post** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/contexts | Update context targets on a big segment
[**UpdateBigSegmentTargets**](SegmentsApi.md#UpdateBigSegmentTargets) | **Post** /api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}/users | Update user context targets on a big segment



## CreateBigSegmentExport

> CreateBigSegmentExport(ctx, projectKey, environmentKey, segmentKey).Execute()

Create big segment export



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
	segmentKey := "segmentKey_example" // string | The segment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SegmentsApi.CreateBigSegmentExport(context.Background(), projectKey, environmentKey, segmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.CreateBigSegmentExport``: %v\n", err)
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
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBigSegmentExportRequest struct via the builder pattern


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


## CreateBigSegmentImport

> CreateBigSegmentImport(ctx, projectKey, environmentKey, segmentKey).File(file).Mode(mode).WaitOnApprovals(waitOnApprovals).Execute()

Create big segment import



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
	segmentKey := "segmentKey_example" // string | The segment key
	file := os.NewFile(1234, "some_file") // *os.File | CSV file containing keys (optional)
	mode := "mode_example" // string | Import mode. Use either `merge` or `replace` (optional)
	waitOnApprovals := true // bool | Whether to wait for approvals before processing the import (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SegmentsApi.CreateBigSegmentImport(context.Background(), projectKey, environmentKey, segmentKey).File(file).Mode(mode).WaitOnApprovals(waitOnApprovals).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.CreateBigSegmentImport``: %v\n", err)
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
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBigSegmentImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | ***os.File** | CSV file containing keys | 
 **mode** | **string** | Import mode. Use either &#x60;merge&#x60; or &#x60;replace&#x60; | 
 **waitOnApprovals** | **bool** | Whether to wait for approvals before processing the import | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSegment

> DeleteSegment(ctx, projectKey, environmentKey, segmentKey).Execute()

Delete segment



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
	segmentKey := "segmentKey_example" // string | The segment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SegmentsApi.DeleteSegment(context.Background(), projectKey, environmentKey, segmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.DeleteSegment``: %v\n", err)
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
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSegmentRequest struct via the builder pattern


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


## GetBigSegmentExport

> Export GetBigSegmentExport(ctx, projectKey, environmentKey, segmentKey, exportID).Execute()

Get big segment export



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
	segmentKey := "segmentKey_example" // string | The segment key
	exportID := "exportID_example" // string | The export ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetBigSegmentExport(context.Background(), projectKey, environmentKey, segmentKey, exportID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetBigSegmentExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBigSegmentExport`: Export
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetBigSegmentExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 
**exportID** | **string** | The export ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBigSegmentExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Export**](Export.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBigSegmentImport

> Import GetBigSegmentImport(ctx, projectKey, environmentKey, segmentKey, importID).Execute()

Get big segment import



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
	segmentKey := "segmentKey_example" // string | The segment key
	importID := "importID_example" // string | The import ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetBigSegmentImport(context.Background(), projectKey, environmentKey, segmentKey, importID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetBigSegmentImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBigSegmentImport`: Import
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetBigSegmentImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 
**importID** | **string** | The import ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBigSegmentImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**Import**](Import.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextInstanceSegmentsMembershipByEnv

> ContextInstanceSegmentMemberships GetContextInstanceSegmentsMembershipByEnv(ctx, projectKey, environmentKey).RequestBody(requestBody).Execute()

List segment memberships for context instance



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetContextInstanceSegmentsMembershipByEnv(context.Background(), projectKey, environmentKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetContextInstanceSegmentsMembershipByEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextInstanceSegmentsMembershipByEnv`: ContextInstanceSegmentMemberships
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetContextInstanceSegmentsMembershipByEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextInstanceSegmentsMembershipByEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**ContextInstanceSegmentMemberships**](ContextInstanceSegmentMemberships.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpiringTargetsForSegment

> ExpiringTargetGetResponse GetExpiringTargetsForSegment(ctx, projectKey, environmentKey, segmentKey).Execute()

Get expiring targets for segment



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
	segmentKey := "segmentKey_example" // string | The segment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetExpiringTargetsForSegment(context.Background(), projectKey, environmentKey, segmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetExpiringTargetsForSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpiringTargetsForSegment`: ExpiringTargetGetResponse
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetExpiringTargetsForSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpiringTargetsForSegmentRequest struct via the builder pattern


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


## GetExpiringUserTargetsForSegment

> ExpiringUserTargetGetResponse GetExpiringUserTargetsForSegment(ctx, projectKey, environmentKey, segmentKey).Execute()

Get expiring user targets for segment



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
	segmentKey := "segmentKey_example" // string | The segment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetExpiringUserTargetsForSegment(context.Background(), projectKey, environmentKey, segmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetExpiringUserTargetsForSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpiringUserTargetsForSegment`: ExpiringUserTargetGetResponse
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetExpiringUserTargetsForSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpiringUserTargetsForSegmentRequest struct via the builder pattern


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


## GetSegment

> UserSegment GetSegment(ctx, projectKey, environmentKey, segmentKey).Execute()

Get segment



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
	segmentKey := "segmentKey_example" // string | The segment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetSegment(context.Background(), projectKey, environmentKey, segmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSegment`: UserSegment
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentMembershipForContext

> BigSegmentTarget GetSegmentMembershipForContext(ctx, projectKey, environmentKey, segmentKey, contextKey).Execute()

Get big segment membership for context



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
	segmentKey := "segmentKey_example" // string | The segment key
	contextKey := "contextKey_example" // string | The context key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetSegmentMembershipForContext(context.Background(), projectKey, environmentKey, segmentKey, contextKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegmentMembershipForContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSegmentMembershipForContext`: BigSegmentTarget
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegmentMembershipForContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 
**contextKey** | **string** | The context key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentMembershipForContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BigSegmentTarget**](BigSegmentTarget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentMembershipForUser

> BigSegmentTarget GetSegmentMembershipForUser(ctx, projectKey, environmentKey, segmentKey, userKey).Execute()

Get big segment membership for user



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
	segmentKey := "segmentKey_example" // string | The segment key
	userKey := "userKey_example" // string | The user key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetSegmentMembershipForUser(context.Background(), projectKey, environmentKey, segmentKey, userKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegmentMembershipForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSegmentMembershipForUser`: BigSegmentTarget
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegmentMembershipForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 
**userKey** | **string** | The user key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentMembershipForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BigSegmentTarget**](BigSegmentTarget.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegments

> UserSegments GetSegments(ctx, projectKey, environmentKey).Limit(limit).Offset(offset).Sort(sort).Filter(filter).Execute()

List segments



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
	limit := int64(789) // int64 | The number of segments to return. Defaults to 20. (optional)
	offset := int64(789) // int64 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)
	sort := "sort_example" // string | Accepts sorting order and fields. Fields can be comma separated. Possible fields are 'creationDate', 'name', 'lastModified'. Example: `sort=name` sort by names ascending or `sort=-name,creationDate` sort by names descending and creationDate ascending. (optional)
	filter := "filter_example" // string | Accepts filter by `excludedKeys`, `external`, `includedKeys`, `query`, `tags`, `unbounded`, `view`. To learn more about the filter syntax, read the  'Filtering segments' section above. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.GetSegments(context.Background(), projectKey, environmentKey).Limit(limit).Offset(offset).Sort(sort).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSegments`: UserSegments
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int64** | The number of segments to return. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 
 **sort** | **string** | Accepts sorting order and fields. Fields can be comma separated. Possible fields are &#39;creationDate&#39;, &#39;name&#39;, &#39;lastModified&#39;. Example: &#x60;sort&#x3D;name&#x60; sort by names ascending or &#x60;sort&#x3D;-name,creationDate&#x60; sort by names descending and creationDate ascending. | 
 **filter** | **string** | Accepts filter by &#x60;excludedKeys&#x60;, &#x60;external&#x60;, &#x60;includedKeys&#x60;, &#x60;query&#x60;, &#x60;tags&#x60;, &#x60;unbounded&#x60;, &#x60;view&#x60;. To learn more about the filter syntax, read the  &#39;Filtering segments&#39; section above. | 

### Return type

[**UserSegments**](UserSegments.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExpiringTargetsForSegment

> ExpiringTargetPatchResponse PatchExpiringTargetsForSegment(ctx, projectKey, environmentKey, segmentKey).PatchSegmentExpiringTargetInputRep(patchSegmentExpiringTargetInputRep).Execute()

Update expiring targets for segment



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
	segmentKey := "segmentKey_example" // string | The segment key
	patchSegmentExpiringTargetInputRep := *openapiclient.NewPatchSegmentExpiringTargetInputRep([]openapiclient.PatchSegmentExpiringTargetInstruction{*openapiclient.NewPatchSegmentExpiringTargetInstruction("addExpiringTarget", "ContextKey_example", "user", "TargetType_example")}) // PatchSegmentExpiringTargetInputRep | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.PatchExpiringTargetsForSegment(context.Background(), projectKey, environmentKey, segmentKey).PatchSegmentExpiringTargetInputRep(patchSegmentExpiringTargetInputRep).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PatchExpiringTargetsForSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExpiringTargetsForSegment`: ExpiringTargetPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PatchExpiringTargetsForSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringTargetsForSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchSegmentExpiringTargetInputRep** | [**PatchSegmentExpiringTargetInputRep**](PatchSegmentExpiringTargetInputRep.md) |  | 

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


## PatchExpiringUserTargetsForSegment

> ExpiringUserTargetPatchResponse PatchExpiringUserTargetsForSegment(ctx, projectKey, environmentKey, segmentKey).PatchSegmentRequest(patchSegmentRequest).Execute()

Update expiring user targets for segment



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
	segmentKey := "segmentKey_example" // string | The segment key
	patchSegmentRequest := *openapiclient.NewPatchSegmentRequest([]openapiclient.PatchSegmentInstruction{*openapiclient.NewPatchSegmentInstruction("addExpireUserTargetDate", "UserKey_example", "TargetType_example")}) // PatchSegmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.PatchExpiringUserTargetsForSegment(context.Background(), projectKey, environmentKey, segmentKey).PatchSegmentRequest(patchSegmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PatchExpiringUserTargetsForSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExpiringUserTargetsForSegment`: ExpiringUserTargetPatchResponse
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PatchExpiringUserTargetsForSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpiringUserTargetsForSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchSegmentRequest** | [**PatchSegmentRequest**](PatchSegmentRequest.md) |  | 

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


## PatchSegment

> UserSegment PatchSegment(ctx, projectKey, environmentKey, segmentKey).PatchWithComment(patchWithComment).DryRun(dryRun).Execute()

Patch segment



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
	segmentKey := "segmentKey_example" // string | The segment key
	patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")}) // PatchWithComment | 
	dryRun := true // bool | If true, the patch will be validated but not persisted. Returns a preview of the segment after the patch is applied. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.PatchSegment(context.Background(), projectKey, environmentKey, segmentKey).PatchWithComment(patchWithComment).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PatchSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSegment`: UserSegment
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PatchSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchWithComment** | [**PatchWithComment**](PatchWithComment.md) |  | 
 **dryRun** | **bool** | If true, the patch will be validated but not persisted. Returns a preview of the segment after the patch is applied. | 

### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSegment

> UserSegment PostSegment(ctx, projectKey, environmentKey).SegmentBody(segmentBody).Execute()

Create segment



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
	segmentBody := *openapiclient.NewSegmentBody("Example segment", "segment-key-123abc") // SegmentBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SegmentsApi.PostSegment(context.Background(), projectKey, environmentKey).SegmentBody(segmentBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.PostSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSegment`: UserSegment
	fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.PostSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **segmentBody** | [**SegmentBody**](SegmentBody.md) |  | 

### Return type

[**UserSegment**](UserSegment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBigSegmentContextTargets

> UpdateBigSegmentContextTargets(ctx, projectKey, environmentKey, segmentKey).SegmentUserState(segmentUserState).Execute()

Update context targets on a big segment



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
	segmentKey := "segmentKey_example" // string | The segment key
	segmentUserState := *openapiclient.NewSegmentUserState() // SegmentUserState | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SegmentsApi.UpdateBigSegmentContextTargets(context.Background(), projectKey, environmentKey, segmentKey).SegmentUserState(segmentUserState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.UpdateBigSegmentContextTargets``: %v\n", err)
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
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBigSegmentContextTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **segmentUserState** | [**SegmentUserState**](SegmentUserState.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBigSegmentTargets

> UpdateBigSegmentTargets(ctx, projectKey, environmentKey, segmentKey).SegmentUserState(segmentUserState).Execute()

Update user context targets on a big segment



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
	segmentKey := "segmentKey_example" // string | The segment key
	segmentUserState := *openapiclient.NewSegmentUserState() // SegmentUserState | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SegmentsApi.UpdateBigSegmentTargets(context.Background(), projectKey, environmentKey, segmentKey).SegmentUserState(segmentUserState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.UpdateBigSegmentTargets``: %v\n", err)
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
**segmentKey** | **string** | The segment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBigSegmentTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **segmentUserState** | [**SegmentUserState**](SegmentUserState.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

