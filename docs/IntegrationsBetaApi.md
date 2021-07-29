# \IntegrationsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateConnection**](IntegrationsBetaApi.md#ValidateConnection) | **Post** /integrations/{integrationKey}/{id}/validate | Validate connection



## ValidateConnection

> IntegrationSubscriptionTestEventRep ValidateConnection(ctx, integrationKey, id).Execute()

Validate connection



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
    integrationKey := "integrationKey_example" // string | The integration key
    id := "id_example" // string | The integration configuration ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IntegrationsBetaApi.ValidateConnection(context.Background(), integrationKey, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsBetaApi.ValidateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateConnection`: IntegrationSubscriptionTestEventRep
    fmt.Fprintf(os.Stdout, "Response from `IntegrationsBetaApi.ValidateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 
**id** | **string** | The integration configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IntegrationSubscriptionTestEventRep**](IntegrationSubscriptionTestEventRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

