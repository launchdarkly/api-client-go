# \InsightsRepositoriesBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateRepositoriesAndProjects**](InsightsRepositoriesBetaApi.md#AssociateRepositoriesAndProjects) | **Put** /api/v2/engineering-insights/repositories/projects | Associate repositories with projects
[**DeleteRepositoryProject**](InsightsRepositoriesBetaApi.md#DeleteRepositoryProject) | **Delete** /api/v2/engineering-insights/repositories/{repositoryKey}/projects/{projectKey} | Remove repository project association
[**GetInsightsRepositories**](InsightsRepositoriesBetaApi.md#GetInsightsRepositories) | **Get** /api/v2/engineering-insights/repositories | List repositories



## AssociateRepositoriesAndProjects

> InsightsRepositoryProjectCollection AssociateRepositoriesAndProjects(ctx).InsightsRepositoryProjectMappings(insightsRepositoryProjectMappings).Execute()

Associate repositories with projects



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
    insightsRepositoryProjectMappings := *openapiclient.NewInsightsRepositoryProjectMappings([]openapiclient.InsightsRepositoryProject{*openapiclient.NewInsightsRepositoryProject("launchdarkly/LaunchDarkly-Docs", "default")}) // InsightsRepositoryProjectMappings | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsRepositoriesBetaApi.AssociateRepositoriesAndProjects(context.Background()).InsightsRepositoryProjectMappings(insightsRepositoryProjectMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsRepositoriesBetaApi.AssociateRepositoriesAndProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssociateRepositoriesAndProjects`: InsightsRepositoryProjectCollection
    fmt.Fprintf(os.Stdout, "Response from `InsightsRepositoriesBetaApi.AssociateRepositoriesAndProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssociateRepositoriesAndProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insightsRepositoryProjectMappings** | [**InsightsRepositoryProjectMappings**](InsightsRepositoryProjectMappings.md) |  | 

### Return type

[**InsightsRepositoryProjectCollection**](InsightsRepositoryProjectCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepositoryProject

> DeleteRepositoryProject(ctx, repositoryKey, projectKey).Execute()

Remove repository project association



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
    repositoryKey := "repositoryKey_example" // string | The repository key
    projectKey := "projectKey_example" // string | The project key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsRepositoriesBetaApi.DeleteRepositoryProject(context.Background(), repositoryKey, projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsRepositoriesBetaApi.DeleteRepositoryProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryKey** | **string** | The repository key | 
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryProjectRequest struct via the builder pattern


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


## GetInsightsRepositories

> InsightsRepositoryCollection GetInsightsRepositories(ctx).Expand(expand).Execute()

List repositories



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
    expand := "expand_example" // string | Expand properties in response. Options: `projects` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightsRepositoriesBetaApi.GetInsightsRepositories(context.Background()).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightsRepositoriesBetaApi.GetInsightsRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsightsRepositories`: InsightsRepositoryCollection
    fmt.Fprintf(os.Stdout, "Response from `InsightsRepositoriesBetaApi.GetInsightsRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInsightsRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expand properties in response. Options: &#x60;projects&#x60; | 

### Return type

[**InsightsRepositoryCollection**](InsightsRepositoryCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

