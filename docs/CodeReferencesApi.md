# \CodeReferencesApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBranches**](CodeReferencesApi.md#DeleteBranches) | **Post** /api/v2/code-refs/repositories/{repo}/branch-delete-tasks | Delete branches
[**DeleteRepository**](CodeReferencesApi.md#DeleteRepository) | **Delete** /api/v2/code-refs/repositories/{repo} | Delete repository
[**GetBranch**](CodeReferencesApi.md#GetBranch) | **Get** /api/v2/code-refs/repositories/{repo}/branches/{branch} | Get branch
[**GetBranches**](CodeReferencesApi.md#GetBranches) | **Get** /api/v2/code-refs/repositories/{repo}/branches | List branches
[**GetExtinctions**](CodeReferencesApi.md#GetExtinctions) | **Get** /api/v2/code-refs/extinctions | List extinctions
[**GetRepositories**](CodeReferencesApi.md#GetRepositories) | **Get** /api/v2/code-refs/repositories | List repositories
[**GetRepository**](CodeReferencesApi.md#GetRepository) | **Get** /api/v2/code-refs/repositories/{repo} | Get repository
[**GetRootStatistic**](CodeReferencesApi.md#GetRootStatistic) | **Get** /api/v2/code-refs/statistics | Get number of code references for flags
[**GetStatistics**](CodeReferencesApi.md#GetStatistics) | **Get** /api/v2/code-refs/statistics/{projKey} | Get number of code references for flags
[**PatchRepository**](CodeReferencesApi.md#PatchRepository) | **Patch** /api/v2/code-refs/repositories/{repo} | Update repository
[**PostExtinction**](CodeReferencesApi.md#PostExtinction) | **Post** /api/v2/code-refs/repositories/{repo}/branches/{branch} | Post extinction
[**PostRepository**](CodeReferencesApi.md#PostRepository) | **Post** /api/v2/code-refs/repositories | Create repository
[**PutBranch**](CodeReferencesApi.md#PutBranch) | **Put** /api/v2/code-refs/repositories/{repo}/branches/{branch} | Upsert branch



## DeleteBranches

> DeleteBranches(ctx, repo).RequestBody(requestBody).Execute()

Delete branches



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
    repo := "repo_example" // string | The repo name to delete branches for.
    requestBody := []string{"Property_example"} // []string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.DeleteBranches(context.Background(), repo).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.DeleteBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repo name to delete branches for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepository

> DeleteRepository(ctx, repo).Execute()

Delete repository



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
    repo := "repo_example" // string | The repository name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.DeleteRepository(context.Background(), repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.DeleteRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repository name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranch

> ApiBranchRep GetBranch(ctx, repo, branch).ProjKey(projKey).FlagKey(flagKey).Execute()

Get branch



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
    repo := "repo_example" // string | The repository name
    branch := "branch_example" // string | The url-encoded branch name
    projKey := "projKey_example" // string | Filter results to a specific project (optional)
    flagKey := "flagKey_example" // string | Filter results to a specific flag key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.GetBranch(context.Background(), repo, branch).ProjKey(projKey).FlagKey(flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.GetBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBranch`: ApiBranchRep
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.GetBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repository name | 
**branch** | **string** | The url-encoded branch name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projKey** | **string** | Filter results to a specific project | 
 **flagKey** | **string** | Filter results to a specific flag key | 

### Return type

[**ApiBranchRep**](ApiBranchRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranches

> ApiBranchCollectionRep GetBranches(ctx, repo).Execute()

List branches



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
    repo := "repo_example" // string | The repository name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.GetBranches(context.Background(), repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.GetBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBranches`: ApiBranchCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.GetBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repository name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiBranchCollectionRep**](ApiBranchCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtinctions

> ApiExtinctionCollectionRep GetExtinctions(ctx).RepoName(repoName).BranchName(branchName).ProjKey(projKey).FlagKey(flagKey).Execute()

List extinctions



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
    repoName := "repoName_example" // string | Filter results to a specific repository (optional)
    branchName := "branchName_example" // string | Filter results to a specific branch (optional)
    projKey := "projKey_example" // string | Filter results to a specific project (optional)
    flagKey := "flagKey_example" // string | Filter results to a specific flag key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.GetExtinctions(context.Background()).RepoName(repoName).BranchName(branchName).ProjKey(projKey).FlagKey(flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.GetExtinctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtinctions`: ApiExtinctionCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.GetExtinctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtinctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoName** | **string** | Filter results to a specific repository | 
 **branchName** | **string** | Filter results to a specific branch | 
 **projKey** | **string** | Filter results to a specific project | 
 **flagKey** | **string** | Filter results to a specific flag key | 

### Return type

[**ApiExtinctionCollectionRep**](ApiExtinctionCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories

> ApiRepositoryCollectionRep GetRepositories(ctx).WithBranches(withBranches).WithReferencesForDefaultBranch(withReferencesForDefaultBranch).ProjKey(projKey).FlagKey(flagKey).Execute()

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
    withBranches := "withBranches_example" // string | If set to any value, the endpoint returns repositories with associated branch data (optional)
    withReferencesForDefaultBranch := "withReferencesForDefaultBranch_example" // string | If set to any value, the endpoint returns repositories with associated branch data, as well as code references for the default git branch (optional)
    projKey := "projKey_example" // string | A LaunchDarkly project key. If provided, this filters code reference results to the specified project. (optional)
    flagKey := "flagKey_example" // string | If set to any value, the endpoint returns repositories with associated branch data, as well as code references for the default git branch (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.GetRepositories(context.Background()).WithBranches(withBranches).WithReferencesForDefaultBranch(withReferencesForDefaultBranch).ProjKey(projKey).FlagKey(flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.GetRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositories`: ApiRepositoryCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.GetRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withBranches** | **string** | If set to any value, the endpoint returns repositories with associated branch data | 
 **withReferencesForDefaultBranch** | **string** | If set to any value, the endpoint returns repositories with associated branch data, as well as code references for the default git branch | 
 **projKey** | **string** | A LaunchDarkly project key. If provided, this filters code reference results to the specified project. | 
 **flagKey** | **string** | If set to any value, the endpoint returns repositories with associated branch data, as well as code references for the default git branch | 

### Return type

[**ApiRepositoryCollectionRep**](ApiRepositoryCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> ApiRepositoryRep GetRepository(ctx, repo).Execute()

Get repository



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
    repo := "repo_example" // string | The repository name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.GetRepository(context.Background(), repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.GetRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepository`: ApiRepositoryRep
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repository name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiRepositoryRep**](ApiRepositoryRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootStatistic

> ApiStatisticsRoot GetRootStatistic(ctx).Execute()

Get number of code references for flags



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.GetRootStatistic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.GetRootStatistic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRootStatistic`: ApiStatisticsRoot
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.GetRootStatistic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootStatisticRequest struct via the builder pattern


### Return type

[**ApiStatisticsRoot**](ApiStatisticsRoot.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatistics

> ApiStatisticCollectionRep GetStatistics(ctx, projKey).FlagKey(flagKey).Execute()

Get number of code references for flags



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
    projKey := "projKey_example" // string | The project key
    flagKey := "flagKey_example" // string | The feature flag key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.GetStatistics(context.Background(), projKey).FlagKey(flagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.GetStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatistics`: ApiStatisticCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.GetStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flagKey** | **string** | The feature flag key | 

### Return type

[**ApiStatisticCollectionRep**](ApiStatisticCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRepository

> ApiRepositoryRep PatchRepository(ctx, repo).JSONPatchElt(jSONPatchElt).Execute()

Update repository



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
    repo := "repo_example" // string | The repository name
    jSONPatchElt := []openapiclient.JSONPatchElt{*openapiclient.NewJSONPatchElt("replace", "/biscuits", interface{}(Chocolate Digestive))} // []JSONPatchElt | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.PatchRepository(context.Background(), repo).JSONPatchElt(jSONPatchElt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.PatchRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRepository`: ApiRepositoryRep
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.PatchRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repository name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSONPatchElt** | [**[]JSONPatchElt**](JSONPatchElt.md) |  | 

### Return type

[**ApiRepositoryRep**](ApiRepositoryRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExtinction

> PostExtinction(ctx, repo, branch).InlineObject(inlineObject).Execute()

Post extinction



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
    repo := "repo_example" // string | The repository name
    branch := "branch_example" // string | The url-encoded branch name
    inlineObject := []openapiclient.InlineObject{*openapiclient.NewInlineObject()} // []InlineObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.PostExtinction(context.Background(), repo, branch).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.PostExtinction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repository name | 
**branch** | **string** | The url-encoded branch name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExtinctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject** | [**[]InlineObject**](InlineObject.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRepository

> PostRepository(ctx).ApiRepositoryPost(apiRepositoryPost).Execute()

Create repository



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
    apiRepositoryPost := *openapiclient.NewApiRepositoryPost() // ApiRepositoryPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.PostRepository(context.Background()).ApiRepositoryPost(apiRepositoryPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.PostRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRepositoryPost** | [**ApiRepositoryPost**](ApiRepositoryPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBranch

> PutBranch(ctx, repo, branch).CoderefsBranch(coderefsBranch).Execute()

Upsert branch



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
    repo := "repo_example" // string | The repository name
    branch := "branch_example" // string | The url-encoded branch name
    coderefsBranch := *openapiclient.NewCoderefsBranch() // CoderefsBranch | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CodeReferencesApi.PutBranch(context.Background(), repo, branch).CoderefsBranch(coderefsBranch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.PutBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | The repository name | 
**branch** | **string** | The url-encoded branch name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **coderefsBranch** | [**CoderefsBranch**](CoderefsBranch.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

