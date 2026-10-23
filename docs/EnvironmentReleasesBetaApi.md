# \EnvironmentReleasesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentRelease**](EnvironmentReleasesBetaApi.md#GetEnvironmentRelease) | **Get** /api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases/{releaseId} | Get an environment release
[**ListEnvironmentReleases**](EnvironmentReleasesBetaApi.md#ListEnvironmentReleases) | **Get** /api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases | List environment releases
[**StartEnvironmentRelease**](EnvironmentReleasesBetaApi.md#StartEnvironmentRelease) | **Post** /api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases | Start an environment release
[**StopEnvironmentRelease**](EnvironmentReleasesBetaApi.md#StopEnvironmentRelease) | **Put** /api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/releases/{releaseId}/stop | Stop an environment release



## GetEnvironmentRelease

> EnvironmentRelease GetEnvironmentRelease(ctx, projectKey, flagKey, environmentKey, releaseId).LDAPIVersion(lDAPIVersion).Expand(expand).Execute()

Get an environment release



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
	projectKey := "projectKey_example" // string | 
	flagKey := "flagKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	releaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	expand := []string{"Expand_example"} // []string | Optional fields to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasesBetaApi.GetEnvironmentRelease(context.Background(), projectKey, flagKey, environmentKey, releaseId).LDAPIVersion(lDAPIVersion).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasesBetaApi.GetEnvironmentRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentRelease`: EnvironmentRelease
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasesBetaApi.GetEnvironmentRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**flagKey** | **string** |  | 
**environmentKey** | **string** |  | 
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 




 **expand** | **[]string** | Optional fields to expand. | 

### Return type

[**EnvironmentRelease**](EnvironmentRelease.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEnvironmentReleases

> EnvironmentReleaseCollection ListEnvironmentReleases(ctx, projectKey, flagKey, environmentKey).LDAPIVersion(lDAPIVersion).Filter(filter).Limit(limit).Offset(offset).Snapshot(snapshot).Execute()

List environment releases



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
	projectKey := "projectKey_example" // string | 
	flagKey := "flagKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	filter := "kind equals guarded" // string | A queryfilter expression that narrows the list. Supported fields and operators: `kind` (`equals`) with values `progressive` or `guarded`.  (optional)
	limit := int32(56) // int32 | The number of releases to return. Defaults to 25. Maximum is 100. (optional) (default to 25)
	offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten releases and then returns the next releases in the list, up to the query `limit`. (optional) (default to 0)
	snapshot := "snapshot_example" // string | Pins the read to a moment in time (nanoseconds since the Unix epoch) so pages walked in sequence reflect one consistent snapshot. Follow the pagination links rather than constructing this value. Out-of-range values are ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasesBetaApi.ListEnvironmentReleases(context.Background(), projectKey, flagKey, environmentKey).LDAPIVersion(lDAPIVersion).Filter(filter).Limit(limit).Offset(offset).Snapshot(snapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasesBetaApi.ListEnvironmentReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnvironmentReleases`: EnvironmentReleaseCollection
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasesBetaApi.ListEnvironmentReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**flagKey** | **string** |  | 
**environmentKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnvironmentReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



 **filter** | **string** | A queryfilter expression that narrows the list. Supported fields and operators: &#x60;kind&#x60; (&#x60;equals&#x60;) with values &#x60;progressive&#x60; or &#x60;guarded&#x60;.  | 
 **limit** | **int32** | The number of releases to return. Defaults to 25. Maximum is 100. | [default to 25]
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten releases and then returns the next releases in the list, up to the query &#x60;limit&#x60;. | [default to 0]
 **snapshot** | **string** | Pins the read to a moment in time (nanoseconds since the Unix epoch) so pages walked in sequence reflect one consistent snapshot. Follow the pagination links rather than constructing this value. Out-of-range values are ignored. | 

### Return type

[**EnvironmentReleaseCollection**](EnvironmentReleaseCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEnvironmentRelease

> EnvironmentRelease StartEnvironmentRelease(ctx, projectKey, flagKey, environmentKey).LDAPIVersion(lDAPIVersion).StartEnvironmentReleaseRequest(startEnvironmentReleaseRequest).Execute()

Start an environment release



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
	projectKey := "projectKey_example" // string | 
	flagKey := "flagKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	startEnvironmentReleaseRequest := *openapiclient.NewStartEnvironmentReleaseRequest("OriginalVariationId_example", "TargetVariationId_example") // StartEnvironmentReleaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasesBetaApi.StartEnvironmentRelease(context.Background(), projectKey, flagKey, environmentKey).LDAPIVersion(lDAPIVersion).StartEnvironmentReleaseRequest(startEnvironmentReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasesBetaApi.StartEnvironmentRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartEnvironmentRelease`: EnvironmentRelease
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasesBetaApi.StartEnvironmentRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**flagKey** | **string** |  | 
**environmentKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartEnvironmentReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



 **startEnvironmentReleaseRequest** | [**StartEnvironmentReleaseRequest**](StartEnvironmentReleaseRequest.md) |  | 

### Return type

[**EnvironmentRelease**](EnvironmentRelease.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopEnvironmentRelease

> EnvironmentRelease StopEnvironmentRelease(ctx, projectKey, flagKey, environmentKey, releaseId).LDAPIVersion(lDAPIVersion).StopEnvironmentReleaseRequest(stopEnvironmentReleaseRequest).Execute()

Stop an environment release



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
	projectKey := "projectKey_example" // string | 
	flagKey := "flagKey_example" // string | 
	environmentKey := "environmentKey_example" // string | 
	releaseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	stopEnvironmentReleaseRequest := *openapiclient.NewStopEnvironmentReleaseRequest(openapiclient.EnvironmentReleaseResolution("roll_forward_all_phases")) // StopEnvironmentReleaseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentReleasesBetaApi.StopEnvironmentRelease(context.Background(), projectKey, flagKey, environmentKey, releaseId).LDAPIVersion(lDAPIVersion).StopEnvironmentReleaseRequest(stopEnvironmentReleaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentReleasesBetaApi.StopEnvironmentRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopEnvironmentRelease`: EnvironmentRelease
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentReleasesBetaApi.StopEnvironmentRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**flagKey** | **string** |  | 
**environmentKey** | **string** |  | 
**releaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopEnvironmentReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 




 **stopEnvironmentReleaseRequest** | [**StopEnvironmentReleaseRequest**](StopEnvironmentReleaseRequest.md) |  | 

### Return type

[**EnvironmentRelease**](EnvironmentRelease.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

