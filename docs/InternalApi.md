# \InternalApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](InternalApi.md#GetAccount) | **Get** /internal/account | Get account
[**GetActions**](InternalApi.md#GetActions) | **Get** /internal/actions | Get actions
[**GetEnvironments**](InternalApi.md#GetEnvironments) | **Get** /internal/projects/{projectKey}/environments | Get environments
[**GetMigrationFlagMetrics**](InternalApi.md#GetMigrationFlagMetrics) | **Get** /internal/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/migration-metrics | List migration flag metrics
[**GetSelectedEnvironments**](InternalApi.md#GetSelectedEnvironments) | **Get** /internal/profile/selected-environments/{projectKey} | Get selected environments
[**InternalGetWorkflows**](InternalApi.md#InternalGetWorkflows) | **Get** /internal/projects/{projectKey}/flags/{featureFlagKey}/workflows | Get workflows across environments
[**InternalSearchFlags**](InternalApi.md#InternalSearchFlags) | **Post** /internal/projects/{projectKey}/flags/search | Search for multiple flags with environment configurations
[**PutSelectedEnvironments**](InternalApi.md#PutSelectedEnvironments) | **Put** /internal/profile/selected-environments/{projectKey} | Put selected environments



## GetAccount

> Account GetAccount(ctx).Execute()

Get account



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActions

> map[string][]ActionAndDescription GetActions(ctx).Execute()

Get actions



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetActions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActions`: map[string][]ActionAndDescription
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetActions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionsRequest struct via the builder pattern


### Return type

[**map[string][]ActionAndDescription**](array.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironments

> InternalEnvironmentCollection GetEnvironments(ctx, projectKey).Env(env).Q(q).Execute()

Get environments



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
    env := "env_example" // string | Filter by environment key (optional)
    q := "q_example" // string | Filter environments by name or key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetEnvironments(context.Background(), projectKey).Env(env).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironments`: InternalEnvironmentCollection
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Filter by environment key | 
 **q** | **string** | Filter environments by name or key | 

### Return type

[**InternalEnvironmentCollection**](InternalEnvironmentCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMigrationFlagMetrics

> MigrationMetricsRep GetMigrationFlagMetrics(ctx, projectKey, flagKey, environmentKey).From(from).To(to).Rules(rules).Execute()

List migration flag metrics



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
    flagKey := "flagKey_example" // string | The migration flag key
    environmentKey := "environmentKey_example" // string | The environment key
    from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to 7 days ago. (optional)
    to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to now. (optional)
    rules := "rules_example" // string | A comma-separated list of up to 10 rules whose metrics should be included in the response. Use the rule IDs or `fallthrough` to identify the rules. Defaults to only returning flag overview metrics. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetMigrationFlagMetrics(context.Background(), projectKey, flagKey, environmentKey).From(from).To(to).Rules(rules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetMigrationFlagMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMigrationFlagMetrics`: MigrationMetricsRep
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetMigrationFlagMetrics`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMigrationFlagMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to 7 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to now. | 
 **rules** | **string** | A comma-separated list of up to 10 rules whose metrics should be included in the response. Use the rule IDs or &#x60;fallthrough&#x60; to identify the rules. Defaults to only returning flag overview metrics. | 

### Return type

[**MigrationMetricsRep**](MigrationMetricsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectedEnvironments

> SelectedEnvironments GetSelectedEnvironments(ctx, projectKey).Execute()

Get selected environments



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.GetSelectedEnvironments(context.Background(), projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.GetSelectedEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelectedEnvironments`: SelectedEnvironments
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.GetSelectedEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectedEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SelectedEnvironments**](SelectedEnvironments.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalGetWorkflows

> CustomWorkflowsListingOutput InternalGetWorkflows(ctx, projectKey, featureFlagKey).Status(status).Sort(sort).Execute()

Get workflows across environments



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    status := "status_example" // string | Filter results by workflow status. Valid status filters are `active`, `completed`, and `failed`. (optional)
    sort := "sort_example" // string | A field to sort the items by. Prefix field by a dash ( - ) to sort in descending order. This endpoint supports sorting by `creationDate` or `stopDate`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.InternalGetWorkflows(context.Background(), projectKey, featureFlagKey).Status(status).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.InternalGetWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InternalGetWorkflows`: CustomWorkflowsListingOutput
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.InternalGetWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInternalGetWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter results by workflow status. Valid status filters are &#x60;active&#x60;, &#x60;completed&#x60;, and &#x60;failed&#x60;. | 
 **sort** | **string** | A field to sort the items by. Prefix field by a dash ( - ) to sort in descending order. This endpoint supports sorting by &#x60;creationDate&#x60; or &#x60;stopDate&#x60;. | 

### Return type

[**CustomWorkflowsListingOutput**](CustomWorkflowsListingOutput.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InternalSearchFlags

> FeatureFlags InternalSearchFlags(ctx, projectKey).FlagStatusesQuery(flagStatusesQuery).Expand(expand).Execute()

Search for multiple flags with environment configurations



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
    flagStatusesQuery := *openapiclient.NewFlagStatusesQuery([]string{"EnvironmentKeys_example"}, []string{"FlagKeys_example"}) // FlagStatusesQuery | 
    expand := "expand_example" // string | Properties to expand. Options: `codeReferences`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.InternalSearchFlags(context.Background(), projectKey).FlagStatusesQuery(flagStatusesQuery).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.InternalSearchFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InternalSearchFlags`: FeatureFlags
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.InternalSearchFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiInternalSearchFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flagStatusesQuery** | [**FlagStatusesQuery**](FlagStatusesQuery.md) |  | 
 **expand** | **string** | Properties to expand. Options: &#x60;codeReferences&#x60;. | 

### Return type

[**FeatureFlags**](FeatureFlags.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSelectedEnvironments

> SelectedEnvironments PutSelectedEnvironments(ctx, projectKey).SelectedEnvironmentsPayload(selectedEnvironmentsPayload).Execute()

Put selected environments



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
    selectedEnvironmentsPayload := *openapiclient.NewSelectedEnvironmentsPayload("SelectedEnvironmentKey_example", []string{"EnvironmentKeys_example"}) // SelectedEnvironmentsPayload | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InternalApi.PutSelectedEnvironments(context.Background(), projectKey).SelectedEnvironmentsPayload(selectedEnvironmentsPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalApi.PutSelectedEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSelectedEnvironments`: SelectedEnvironments
    fmt.Fprintf(os.Stdout, "Response from `InternalApi.PutSelectedEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSelectedEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedEnvironmentsPayload** | [**SelectedEnvironmentsPayload**](SelectedEnvironmentsPayload.md) |  | 

### Return type

[**SelectedEnvironments**](SelectedEnvironments.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

