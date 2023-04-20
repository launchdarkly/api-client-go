# \ContextSettingsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutContextFlagSetting**](ContextSettingsApi.md#PutContextFlagSetting) | **Put** /api/v2/projects/{projectKey}/environments/{environmentKey}/contexts/{contextKind}/{contextKey}/flags/{featureFlagKey} | Update flag settings for context



## PutContextFlagSetting

> PutContextFlagSetting(ctx, projectKey, environmentKey, contextKind, contextKey, featureFlagKey).ValuePut(valuePut).Execute()

Update flag settings for context



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
    contextKind := "contextKind_example" // string | The context kind
    contextKey := "contextKey_example" // string | The context key
    featureFlagKey := "featureFlagKey_example" // string | The feature flag key
    valuePut := *openapiclient.NewValuePut() // ValuePut | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextSettingsApi.PutContextFlagSetting(context.Background(), projectKey, environmentKey, contextKind, contextKey, featureFlagKey).ValuePut(valuePut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextSettingsApi.PutContextFlagSetting``: %v\n", err)
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
**contextKind** | **string** | The context kind | 
**contextKey** | **string** | The context key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutContextFlagSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **valuePut** | [**ValuePut**](ValuePut.md) |  | 

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

