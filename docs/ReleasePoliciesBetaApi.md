# \ReleasePoliciesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteReleasePolicy**](ReleasePoliciesBetaApi.md#DeleteReleasePolicy) | **Delete** /api/v2/projects/{projectKey}/release-policies/{policyKey} | Delete a release policy
[**GetReleasePolicies**](ReleasePoliciesBetaApi.md#GetReleasePolicies) | **Get** /api/v2/projects/{projectKey}/release-policies | List release policies
[**GetReleasePolicy**](ReleasePoliciesBetaApi.md#GetReleasePolicy) | **Get** /api/v2/projects/{projectKey}/release-policies/{policyKey} | Get a release policy by key
[**PostReleasePoliciesOrder**](ReleasePoliciesBetaApi.md#PostReleasePoliciesOrder) | **Post** /api/v2/projects/{projectKey}/release-policies/order | Update the order of existing release policies
[**PostReleasePolicy**](ReleasePoliciesBetaApi.md#PostReleasePolicy) | **Post** /api/v2/projects/{projectKey}/release-policies | Create a release policy
[**PutReleasePolicy**](ReleasePoliciesBetaApi.md#PutReleasePolicy) | **Put** /api/v2/projects/{projectKey}/release-policies/{policyKey} | Update a release policy



## DeleteReleasePolicy

> DeleteReleasePolicy(ctx, projectKey, policyKey).LDAPIVersion(lDAPIVersion).Execute()

Delete a release policy



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
	policyKey := "production-release" // string | The human-readable key of the release policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReleasePoliciesBetaApi.DeleteReleasePolicy(context.Background(), projectKey, policyKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePoliciesBetaApi.DeleteReleasePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**policyKey** | **string** | The human-readable key of the release policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleasePolicyRequest struct via the builder pattern


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


## GetReleasePolicies

> ReleasePoliciesResponse GetReleasePolicies(ctx, projectKey).LDAPIVersion(lDAPIVersion).ExcludeDefault(excludeDefault).Execute()

List release policies



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
	excludeDefault := false // bool | When true, exclude the default release policy from the response. When false or omitted, include the default policy if an environment filter is present. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePoliciesBetaApi.GetReleasePolicies(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).ExcludeDefault(excludeDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePoliciesBetaApi.GetReleasePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleasePolicies`: ReleasePoliciesResponse
	fmt.Fprintf(os.Stdout, "Response from `ReleasePoliciesBetaApi.GetReleasePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **excludeDefault** | **bool** | When true, exclude the default release policy from the response. When false or omitted, include the default policy if an environment filter is present. | [default to false]

### Return type

[**ReleasePoliciesResponse**](ReleasePoliciesResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasePolicy

> ReleasePolicy GetReleasePolicy(ctx, projectKey, policyKey).LDAPIVersion(lDAPIVersion).Execute()

Get a release policy by key



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
	policyKey := "production-release" // string | The release policy key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePoliciesBetaApi.GetReleasePolicy(context.Background(), projectKey, policyKey).LDAPIVersion(lDAPIVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePoliciesBetaApi.GetReleasePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleasePolicy`: ReleasePolicy
	fmt.Fprintf(os.Stdout, "Response from `ReleasePoliciesBetaApi.GetReleasePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**policyKey** | **string** | The release policy key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 



### Return type

[**ReleasePolicy**](ReleasePolicy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReleasePoliciesOrder

> []ReleasePolicy PostReleasePoliciesOrder(ctx, projectKey).LDAPIVersion(lDAPIVersion).RequestBody(requestBody).Execute()

Update the order of existing release policies



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
	requestBody := []string{"Property_example"} // []string | Array of release policy keys in the desired rank order (scoped policies only). These keys must include _all_ of the scoped release policies for the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePoliciesBetaApi.PostReleasePoliciesOrder(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePoliciesBetaApi.PostReleasePoliciesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReleasePoliciesOrder`: []ReleasePolicy
	fmt.Fprintf(os.Stdout, "Response from `ReleasePoliciesBetaApi.PostReleasePoliciesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReleasePoliciesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **requestBody** | **[]string** | Array of release policy keys in the desired rank order (scoped policies only). These keys must include _all_ of the scoped release policies for the project. | 

### Return type

[**[]ReleasePolicy**](ReleasePolicy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReleasePolicy

> ReleasePolicy PostReleasePolicy(ctx, projectKey).LDAPIVersion(lDAPIVersion).PostReleasePolicyRequest(postReleasePolicyRequest).Execute()

Create a release policy



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
	postReleasePolicyRequest := *openapiclient.NewPostReleasePolicyRequest(openapiclient.ReleaseMethod("guarded-release"), "Production Release", "production-release") // PostReleasePolicyRequest | Release policy to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePoliciesBetaApi.PostReleasePolicy(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).PostReleasePolicyRequest(postReleasePolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePoliciesBetaApi.PostReleasePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReleasePolicy`: ReleasePolicy
	fmt.Fprintf(os.Stdout, "Response from `ReleasePoliciesBetaApi.PostReleasePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostReleasePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **postReleasePolicyRequest** | [**PostReleasePolicyRequest**](PostReleasePolicyRequest.md) | Release policy to create | 

### Return type

[**ReleasePolicy**](ReleasePolicy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReleasePolicy

> ReleasePolicy PutReleasePolicy(ctx, projectKey, policyKey).LDAPIVersion(lDAPIVersion).PutReleasePolicyRequest(putReleasePolicyRequest).Execute()

Update a release policy



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
	policyKey := "production-release" // string | The human-readable key of the release policy
	putReleasePolicyRequest := *openapiclient.NewPutReleasePolicyRequest(openapiclient.ReleaseMethod("guarded-release"), "Production Release") // PutReleasePolicyRequest | Release policy data to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleasePoliciesBetaApi.PutReleasePolicy(context.Background(), projectKey, policyKey).LDAPIVersion(lDAPIVersion).PutReleasePolicyRequest(putReleasePolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleasePoliciesBetaApi.PutReleasePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutReleasePolicy`: ReleasePolicy
	fmt.Fprintf(os.Stdout, "Response from `ReleasePoliciesBetaApi.PutReleasePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**policyKey** | **string** | The human-readable key of the release policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutReleasePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 


 **putReleasePolicyRequest** | [**PutReleasePolicyRequest**](PutReleasePolicyRequest.md) | Release policy data to update | 

### Return type

[**ReleasePolicy**](ReleasePolicy.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

