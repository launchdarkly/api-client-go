# \UsersBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserAttributeNames**](UsersBetaApi.md#GetUserAttributeNames) | **Get** /api/v2/user-attributes/{projectKey}/{environmentKey} | Get user attribute names



## GetUserAttributeNames

> UserAttributeNamesRep GetUserAttributeNames(ctx, projectKey, environmentKey).Execute()

Get user attribute names



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
	projectKey := "projectKey_example" // string | The project key
	environmentKey := "environmentKey_example" // string | The environment key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersBetaApi.GetUserAttributeNames(context.Background(), projectKey, environmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersBetaApi.GetUserAttributeNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAttributeNames`: UserAttributeNamesRep
	fmt.Fprintf(os.Stdout, "Response from `UsersBetaApi.GetUserAttributeNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAttributeNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserAttributeNamesRep**](UserAttributeNamesRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

