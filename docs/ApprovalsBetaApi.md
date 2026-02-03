# \ApprovalsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApprovalRequestSettings**](ApprovalsBetaApi.md#GetApprovalRequestSettings) | **Get** /api/v2/approval-requests/projects/{projectKey}/settings | Get approval request settings
[**PatchApprovalRequest**](ApprovalsBetaApi.md#PatchApprovalRequest) | **Patch** /api/v2/approval-requests/{id} | Update approval request
[**PatchApprovalRequestSettings**](ApprovalsBetaApi.md#PatchApprovalRequestSettings) | **Patch** /api/v2/approval-requests/projects/{projectKey}/settings | Update approval request settings
[**PatchFlagConfigApprovalRequest**](ApprovalsBetaApi.md#PatchFlagConfigApprovalRequest) | **Patch** /api/v2/projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{id} | Update flag approval request



## GetApprovalRequestSettings

> map[string]ApprovalRequestSettingWithEnvs GetApprovalRequestSettings(ctx, projectKey).LDAPIVersion(lDAPIVersion).EnvironmentKey(environmentKey).ResourceKind(resourceKind).Expand(expand).Execute()

Get approval request settings



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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	environmentKey := "environmentKey_example" // string | An environment key filter to apply to the approval request settings. (optional)
	resourceKind := "resourceKind_example" // string | A resource kind filter to apply to the approval request settings. (optional)
	expand := "default,strict" // string | A comma-separated list of fields to expand in the response. Options include 'default' and 'strict'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsBetaApi.GetApprovalRequestSettings(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).EnvironmentKey(environmentKey).ResourceKind(resourceKind).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.GetApprovalRequestSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApprovalRequestSettings`: map[string]ApprovalRequestSettingWithEnvs
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.GetApprovalRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **environmentKey** | **string** | An environment key filter to apply to the approval request settings. | 
 **resourceKind** | **string** | A resource kind filter to apply to the approval request settings. | 
 **expand** | **string** | A comma-separated list of fields to expand in the response. Options include &#39;default&#39; and &#39;strict&#39;. | 

### Return type

[**map[string]ApprovalRequestSettingWithEnvs**](ApprovalRequestSettingWithEnvs.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApprovalRequest

> FlagConfigApprovalRequestResponse PatchApprovalRequest(ctx, id).ApprovalRequestPatchInput(approvalRequestPatchInput).Execute()

Update approval request



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
	id := "id_example" // string | The approval ID
	approvalRequestPatchInput := *openapiclient.NewApprovalRequestPatchInput([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // ApprovalRequestPatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsBetaApi.PatchApprovalRequest(context.Background(), id).ApprovalRequestPatchInput(approvalRequestPatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.PatchApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApprovalRequest`: FlagConfigApprovalRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.PatchApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The approval ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approvalRequestPatchInput** | [**ApprovalRequestPatchInput**](ApprovalRequestPatchInput.md) |  | 

### Return type

[**FlagConfigApprovalRequestResponse**](FlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApprovalRequestSettings

> map[string]ApprovalRequestSettingWithEnvs PatchApprovalRequestSettings(ctx, projectKey).LDAPIVersion(lDAPIVersion).ApprovalRequestSettingsPatch(approvalRequestSettingsPatch).Execute()

Update approval request settings



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
	lDAPIVersion := "lDAPIVersion_example" // string | Version of the endpoint.
	projectKey := "default" // string | 
	approvalRequestSettingsPatch := *openapiclient.NewApprovalRequestSettingsPatch("EnvironmentKey_example", "ResourceKind_example") // ApprovalRequestSettingsPatch | Approval request settings to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsBetaApi.PatchApprovalRequestSettings(context.Background(), projectKey).LDAPIVersion(lDAPIVersion).ApprovalRequestSettingsPatch(approvalRequestSettingsPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.PatchApprovalRequestSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApprovalRequestSettings`: map[string]ApprovalRequestSettingWithEnvs
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.PatchApprovalRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApprovalRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lDAPIVersion** | **string** | Version of the endpoint. | 

 **approvalRequestSettingsPatch** | [**ApprovalRequestSettingsPatch**](ApprovalRequestSettingsPatch.md) | Approval request settings to update | 

### Return type

[**map[string]ApprovalRequestSettingWithEnvs**](ApprovalRequestSettingWithEnvs.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFlagConfigApprovalRequest

> FlagConfigApprovalRequestResponse PatchFlagConfigApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey, id).Execute()

Update flag approval request



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
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key
	environmentKey := "environmentKey_example" // string | The environment key
	id := "id_example" // string | The approval ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsBetaApi.PatchFlagConfigApprovalRequest(context.Background(), projectKey, featureFlagKey, environmentKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsBetaApi.PatchFlagConfigApprovalRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFlagConfigApprovalRequest`: FlagConfigApprovalRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `ApprovalsBetaApi.PatchFlagConfigApprovalRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**featureFlagKey** | **string** | The feature flag key | 
**environmentKey** | **string** | The environment key | 
**id** | **string** | The approval ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFlagConfigApprovalRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FlagConfigApprovalRequestResponse**](FlagConfigApprovalRequestResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

