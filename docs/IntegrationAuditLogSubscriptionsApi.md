# \IntegrationAuditLogSubscriptionsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](IntegrationAuditLogSubscriptionsApi.md#CreateSubscription) | **Post** /api/v2/integrations/{integrationKey} | Create audit log subscription
[**DeleteSubscription**](IntegrationAuditLogSubscriptionsApi.md#DeleteSubscription) | **Delete** /api/v2/integrations/{integrationKey}/{id} | Delete audit log subscription
[**GetSubscriptionByID**](IntegrationAuditLogSubscriptionsApi.md#GetSubscriptionByID) | **Get** /api/v2/integrations/{integrationKey}/{id} | Get audit log subscription by ID
[**GetSubscriptions**](IntegrationAuditLogSubscriptionsApi.md#GetSubscriptions) | **Get** /api/v2/integrations/{integrationKey} | Get audit log subscriptions by integration
[**UpdateSubscription**](IntegrationAuditLogSubscriptionsApi.md#UpdateSubscription) | **Patch** /api/v2/integrations/{integrationKey}/{id} | Update audit log subscription



## CreateSubscription

> Integration CreateSubscription(ctx, integrationKey).SubscriptionPost(subscriptionPost).Execute()

Create audit log subscription



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
	integrationKey := "integrationKey_example" // string | The integration key
	subscriptionPost := *openapiclient.NewSubscriptionPost("Example audit log subscription.", map[string]interface{}{"key": interface{}(123)}) // SubscriptionPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAuditLogSubscriptionsApi.CreateSubscription(context.Background(), integrationKey).SubscriptionPost(subscriptionPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAuditLogSubscriptionsApi.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAuditLogSubscriptionsApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionPost** | [**SubscriptionPost**](SubscriptionPost.md) |  | 

### Return type

[**Integration**](Integration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, integrationKey, id).Execute()

Delete audit log subscription



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
	integrationKey := "integrationKey_example" // string | The integration key
	id := "id_example" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAuditLogSubscriptionsApi.DeleteSubscription(context.Background(), integrationKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAuditLogSubscriptionsApi.DeleteSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


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


## GetSubscriptionByID

> Integration GetSubscriptionByID(ctx, integrationKey, id).Execute()

Get audit log subscription by ID



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
	integrationKey := "integrationKey_example" // string | The integration key
	id := "id_example" // string | The subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAuditLogSubscriptionsApi.GetSubscriptionByID(context.Background(), integrationKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAuditLogSubscriptionsApi.GetSubscriptionByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionByID`: Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAuditLogSubscriptionsApi.GetSubscriptionByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 
**id** | **string** | The subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Integration**](Integration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> Integrations GetSubscriptions(ctx, integrationKey).Execute()

Get audit log subscriptions by integration



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
	integrationKey := "integrationKey_example" // string | The integration key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAuditLogSubscriptionsApi.GetSubscriptions(context.Background(), integrationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAuditLogSubscriptionsApi.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: Integrations
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAuditLogSubscriptionsApi.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integrations**](Integrations.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> Integration UpdateSubscription(ctx, integrationKey, id).PatchOperation(patchOperation).Execute()

Update audit log subscription



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
	integrationKey := "integrationKey_example" // string | The integration key
	id := "id_example" // string | The ID of the audit log subscription
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAuditLogSubscriptionsApi.UpdateSubscription(context.Background(), integrationKey, id).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAuditLogSubscriptionsApi.UpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscription`: Integration
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAuditLogSubscriptionsApi.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationKey** | **string** | The integration key | 
**id** | **string** | The ID of the audit log subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**Integration**](Integration.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

