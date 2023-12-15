# \ApplicationsPrivateApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostApplicationAndVersionWithFlagsPrivate**](ApplicationsPrivateApi.md#PostApplicationAndVersionWithFlagsPrivate) | **Post** /api/private/applications | Create applications and versions with flag associations



## PostApplicationAndVersionWithFlagsPrivate

> PostApplicationAndVersionWithFlagsPrivate(ctx).CreateApplicationsAndVersionsInput(createApplicationsAndVersionsInput).Execute()

Create applications and versions with flag associations



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
    createApplicationsAndVersionsInput := *openapiclient.NewCreateApplicationsAndVersionsInput() // CreateApplicationsAndVersionsInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsPrivateApi.PostApplicationAndVersionWithFlagsPrivate(context.Background()).CreateApplicationsAndVersionsInput(createApplicationsAndVersionsInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsPrivateApi.PostApplicationAndVersionWithFlagsPrivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApplicationAndVersionWithFlagsPrivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApplicationsAndVersionsInput** | [**CreateApplicationsAndVersionsInput**](CreateApplicationsAndVersionsInput.md) |  | 

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

