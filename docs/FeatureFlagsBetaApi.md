# \FeatureFlagsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDependentFlags**](FeatureFlagsBetaApi.md#GetDependentFlags) | **Get** /api/v2/flags/{projectKey}/{featureFlagKey}/dependent-flags | List dependent feature flags
[**GetDependentFlagsByEnv**](FeatureFlagsBetaApi.md#GetDependentFlagsByEnv) | **Get** /api/v2/flags/{projectKey}/{environmentKey}/{featureFlagKey}/dependent-flags | List dependent feature flags by environment
[**PostMigrationSafetyIssues**](FeatureFlagsBetaApi.md#PostMigrationSafetyIssues) | **Post** /api/v2/projects/{projectKey}/flags/{flagKey}/environments/{environmentKey}/migration-safety-issues | Get migration safety issues



## GetDependentFlags

> MultiEnvironmentDependentFlags GetDependentFlags(ctx, projectKey, featureFlagKey).Execute()

List dependent feature flags



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagsBetaApi.GetDependentFlags(context.Background(), projectKey, featureFlagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsBetaApi.GetDependentFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDependentFlags`: MultiEnvironmentDependentFlags
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsBetaApi.GetDependentFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDependentFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MultiEnvironmentDependentFlags**](MultiEnvironmentDependentFlags.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDependentFlagsByEnv

> DependentFlagsByEnvironment GetDependentFlagsByEnv(ctx, projectKey, environmentKey, featureFlagKey).Execute()

List dependent feature flags by environment



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
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagsBetaApi.GetDependentFlagsByEnv(context.Background(), projectKey, environmentKey, featureFlagKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsBetaApi.GetDependentFlagsByEnv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDependentFlagsByEnv`: DependentFlagsByEnvironment
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsBetaApi.GetDependentFlagsByEnv`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDependentFlagsByEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DependentFlagsByEnvironment**](DependentFlagsByEnvironment.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | The project key
    flagKey := "flagKey_example" // string | The migration flag key
    environmentKey := "environmentKey_example" // string | The environment key
    flagSempatch := *openapiclient.NewFlagSempatch([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // FlagSempatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagsBetaApi.PostMigrationSafetyIssues(context.Background(), projectKey, flagKey, environmentKey).FlagSempatch(flagSempatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsBetaApi.PostMigrationSafetyIssues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMigrationSafetyIssues`: []MigrationSafetyIssueRep
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsBetaApi.PostMigrationSafetyIssues`: %v\n", resp)
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

