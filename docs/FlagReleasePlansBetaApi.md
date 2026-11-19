# \FlagReleasePlansBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlagReleasePlan**](FlagReleasePlansBetaApi.md#CreateFlagReleasePlan) | **Post** /api/v2/projects/{projectKey}/flag-release-plans | Create a flag release plan for a project
[**DeleteFlagReleasePlan**](FlagReleasePlansBetaApi.md#DeleteFlagReleasePlan) | **Delete** /api/v2/projects/{projectKey}/flag-release-plans/{planKey} | Delete a flag release plan in a project
[**GetFlagReleasePlan**](FlagReleasePlansBetaApi.md#GetFlagReleasePlan) | **Get** /api/v2/projects/{projectKey}/flag-release-plans/{planKey} | Get a flag release plan in a project by plan key
[**ListFlagReleasePlans**](FlagReleasePlansBetaApi.md#ListFlagReleasePlans) | **Get** /api/v2/projects/{projectKey}/flag-release-plans | List a project&#39;s flag release plans
[**UpdateFlagReleasePlan**](FlagReleasePlansBetaApi.md#UpdateFlagReleasePlan) | **Put** /api/v2/projects/{projectKey}/flag-release-plans/{planKey} | Replace a flag release plan configuration in a project



## CreateFlagReleasePlan

> FlagReleasePlan CreateFlagReleasePlan(ctx, projectKey).LDAPIVersion(lDAPIVersion).CreateFlagReleasePlanRequest(createFlagReleasePlanRequest).Execute()

Create a flag release plan for a project



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
	createFlagReleasePlanRequest := *openapiclient.NewCreateFlagReleasePlanRequest("Name_example", []openapiclient.FlagReleasePlanEnvironment{*openapiclient.NewFlagReleasePlanEnvironment("EnvironmentKey_example", openapiclient.TriggerType("flag_eval"))}, "Key_example") // CreateFlagReleasePlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagReleasePlansBetaApi.CreateFlagReleasePlan(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).CreateFlagReleasePlanRequest(createFlagReleasePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagReleasePlansBetaApi.CreateFlagReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlagReleasePlan`: FlagReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `FlagReleasePlansBetaApi.CreateFlagReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlagReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **createFlagReleasePlanRequest** | [**CreateFlagReleasePlanRequest**](CreateFlagReleasePlanRequest.md) |  | 

### Return type

[**FlagReleasePlan**](FlagReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlagReleasePlan

> DeleteFlagReleasePlan(ctx, projectKey, planKey).LDAPIVersion(lDAPIVersion).Execute()

Delete a flag release plan in a project



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
	planKey := "planKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlagReleasePlansBetaApi.DeleteFlagReleasePlan(context.Background(), projectKey, planKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagReleasePlansBetaApi.DeleteFlagReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**planKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagReleasePlanRequest struct via the builder pattern


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


## GetFlagReleasePlan

> FlagReleasePlan GetFlagReleasePlan(ctx, projectKey, planKey).LDAPIVersion(lDAPIVersion).Execute()

Get a flag release plan in a project by plan key



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
	planKey := "planKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagReleasePlansBetaApi.GetFlagReleasePlan(context.Background(), projectKey, planKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagReleasePlansBetaApi.GetFlagReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagReleasePlan`: FlagReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `FlagReleasePlansBetaApi.GetFlagReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**planKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



### Return type

[**FlagReleasePlan**](FlagReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlagReleasePlans

> ListFlagReleasePlansResponse ListFlagReleasePlans(ctx, projectKey).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Execute()

List a project's flag release plans



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
	limit := int32(56) // int32 | The number of flag release plans to return. (optional) (default to 25)
	offset := int32(56) // int32 | The number of flag release plans to skip. Used for pagination. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagReleasePlansBetaApi.ListFlagReleasePlans(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagReleasePlansBetaApi.ListFlagReleasePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFlagReleasePlans`: ListFlagReleasePlansResponse
	fmt.Fprintf(os.Stdout, "Response from `FlagReleasePlansBetaApi.ListFlagReleasePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFlagReleasePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **limit** | **int32** | The number of flag release plans to return. | [default to 25]
 **offset** | **int32** | The number of flag release plans to skip. Used for pagination. | [default to 0]

### Return type

[**ListFlagReleasePlansResponse**](ListFlagReleasePlansResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlagReleasePlan

> FlagReleasePlan UpdateFlagReleasePlan(ctx, projectKey, planKey).LDAPIVersion(lDAPIVersion).PutFlagReleasePlanRequest(putFlagReleasePlanRequest).Execute()

Replace a flag release plan configuration in a project



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
	planKey := "planKey_example" // string | 
	putFlagReleasePlanRequest := *openapiclient.NewPutFlagReleasePlanRequest("Name_example", []openapiclient.FlagReleasePlanEnvironment{*openapiclient.NewFlagReleasePlanEnvironment("EnvironmentKey_example", openapiclient.TriggerType("flag_eval"))}) // PutFlagReleasePlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagReleasePlansBetaApi.UpdateFlagReleasePlan(context.Background(), projectKey, planKey).LDAPIVersion(lDAPIVersion).PutFlagReleasePlanRequest(putFlagReleasePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagReleasePlansBetaApi.UpdateFlagReleasePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFlagReleasePlan`: FlagReleasePlan
	fmt.Fprintf(os.Stdout, "Response from `FlagReleasePlansBetaApi.UpdateFlagReleasePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**planKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlagReleasePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **putFlagReleasePlanRequest** | [**PutFlagReleasePlanRequest**](PutFlagReleasePlanRequest.md) |  | 

### Return type

[**FlagReleasePlan**](FlagReleasePlan.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

