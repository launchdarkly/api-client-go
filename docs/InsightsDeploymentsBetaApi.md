# \InsightsDeploymentsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeploymentEvent**](InsightsDeploymentsBetaApi.md#CreateDeploymentEvent) | **Post** /api/v2/engineering-insights/deployment-events | Create deployment event
[**GetDeployment**](InsightsDeploymentsBetaApi.md#GetDeployment) | **Get** /api/v2/engineering-insights/deployments/{deploymentID} | Get deployment
[**GetDeployments**](InsightsDeploymentsBetaApi.md#GetDeployments) | **Get** /api/v2/engineering-insights/deployments | List deployments
[**UpdateDeployment**](InsightsDeploymentsBetaApi.md#UpdateDeployment) | **Patch** /api/v2/engineering-insights/deployments/{deploymentID} | Update deployment



## CreateDeploymentEvent

> CreateDeploymentEvent(ctx).PostDeploymentEventInput(postDeploymentEventInput).Execute()

Create deployment event



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    postDeploymentEventInput := *openapiclient.NewPostDeploymentEventInput("default", "production", "billing-service", "a90a8a2", "started") // PostDeploymentEventInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsDeploymentsBetaApi.CreateDeploymentEvent(context.Background()).PostDeploymentEventInput(postDeploymentEventInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsDeploymentsBetaApi.CreateDeploymentEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeploymentEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postDeploymentEventInput** | [**PostDeploymentEventInput**](PostDeploymentEventInput.md) |  | 

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


## GetDeployment

> DeploymentRep GetDeployment(ctx, deploymentID).Expand(expand).Execute()

Get deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deploymentID := "deploymentID_example" // string | The deployment ID
    expand := "expand_example" // string | Expand properties in response. Options: `pullRequests`, `flagReferences` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsDeploymentsBetaApi.GetDeployment(context.Background(), deploymentID).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsDeploymentsBetaApi.GetDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployment`: DeploymentRep
    fmt.Fprintf(os.Stdout, "Response from `InsightsDeploymentsBetaApi.GetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentID** | **string** | The deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Expand properties in response. Options: &#x60;pullRequests&#x60;, &#x60;flagReferences&#x60; | 

### Return type

[**DeploymentRep**](DeploymentRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployments

> DeploymentCollectionRep GetDeployments(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Limit(limit).Expand(expand).From(from).To(to).After(after).Before(before).Kind(kind).Status(status).Execute()

List deployments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | The project key
    environmentKey := "environmentKey_example" // string | The environment key
    applicationKey := "applicationKey_example" // string | Comma separated list of application keys (optional)
    limit := int64(789) // int64 | The number of deployments to return. Default is 20. Maximum allowed is 100. (optional)
    expand := "expand_example" // string | Expand properties in response. Options: `pullRequests`, `flagReferences` (optional)
    from := int64(789) // int64 | Unix timestamp in milliseconds. Default value is 7 days ago. (optional)
    to := int64(789) // int64 | Unix timestamp in milliseconds. Default value is now. (optional)
    after := "after_example" // string | Identifier used for pagination (optional)
    before := "before_example" // string | Identifier used for pagination (optional)
    kind := "kind_example" // string | The deployment kind (optional)
    status := "status_example" // string | The deployment status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsDeploymentsBetaApi.GetDeployments(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Limit(limit).Expand(expand).From(from).To(to).After(after).Before(before).Kind(kind).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsDeploymentsBetaApi.GetDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeployments`: DeploymentCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `InsightsDeploymentsBetaApi.GetDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 
 **limit** | **int64** | The number of deployments to return. Default is 20. Maximum allowed is 100. | 
 **expand** | **string** | Expand properties in response. Options: &#x60;pullRequests&#x60;, &#x60;flagReferences&#x60; | 
 **from** | **int64** | Unix timestamp in milliseconds. Default value is 7 days ago. | 
 **to** | **int64** | Unix timestamp in milliseconds. Default value is now. | 
 **after** | **string** | Identifier used for pagination | 
 **before** | **string** | Identifier used for pagination | 
 **kind** | **string** | The deployment kind | 
 **status** | **string** | The deployment status | 

### Return type

[**DeploymentCollectionRep**](DeploymentCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeployment

> DeploymentRep UpdateDeployment(ctx, deploymentID).PatchOperation(patchOperation).Execute()

Update deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deploymentID := "deploymentID_example" // string | The deployment ID
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsDeploymentsBetaApi.UpdateDeployment(context.Background(), deploymentID).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsDeploymentsBetaApi.UpdateDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeployment`: DeploymentRep
    fmt.Fprintf(os.Stdout, "Response from `InsightsDeploymentsBetaApi.UpdateDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentID** | **string** | The deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**DeploymentRep**](DeploymentRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

