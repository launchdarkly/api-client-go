# \IPAllowlistBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpAllowlistEntry**](IPAllowlistBetaApi.md#CreateIpAllowlistEntry) | **Post** /api/v2/account/ip-allowlist | Create IP Allowlist Entry
[**DeleteIpAllowlistEntry**](IPAllowlistBetaApi.md#DeleteIpAllowlistEntry) | **Delete** /api/v2/account/ip-allowlist/{id} | Delete IP Allowlist Entry
[**GetIpAllowlist**](IPAllowlistBetaApi.md#GetIpAllowlist) | **Get** /api/v2/account/ip-allowlist | Get IP Allowlist
[**PatchIpAllowlistConfig**](IPAllowlistBetaApi.md#PatchIpAllowlistConfig) | **Patch** /api/v2/account/ip-allowlist | Update IP Allowlist Configuration
[**PatchIpAllowlistEntry**](IPAllowlistBetaApi.md#PatchIpAllowlistEntry) | **Patch** /api/v2/account/ip-allowlist/{id} | Update IP Allowlist Entry Description



## CreateIpAllowlistEntry

> IpAllowlistEntryResponse CreateIpAllowlistEntry(ctx).CreateIpAllowlistEntryRequest(createIpAllowlistEntryRequest).Execute()

Create IP Allowlist Entry



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
	createIpAllowlistEntryRequest := *openapiclient.NewCreateIpAllowlistEntryRequest("203.0.113.0/24") // CreateIpAllowlistEntryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAllowlistBetaApi.CreateIpAllowlistEntry(context.Background()).CreateIpAllowlistEntryRequest(createIpAllowlistEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistBetaApi.CreateIpAllowlistEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIpAllowlistEntry`: IpAllowlistEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `IPAllowlistBetaApi.CreateIpAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIpAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIpAllowlistEntryRequest** | [**CreateIpAllowlistEntryRequest**](CreateIpAllowlistEntryRequest.md) |  | 

### Return type

[**IpAllowlistEntryResponse**](IpAllowlistEntryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIpAllowlistEntry

> DeleteIpAllowlistEntry(ctx, id).Execute()

Delete IP Allowlist Entry



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
	id := "id_example" // string | Unique identifier for the allowlist entry

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IPAllowlistBetaApi.DeleteIpAllowlistEntry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistBetaApi.DeleteIpAllowlistEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the allowlist entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpAllowlistEntryRequest struct via the builder pattern


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


## GetIpAllowlist

> IpAllowlistResponse GetIpAllowlist(ctx).Execute()

Get IP Allowlist



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAllowlistBetaApi.GetIpAllowlist(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistBetaApi.GetIpAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpAllowlist`: IpAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `IPAllowlistBetaApi.GetIpAllowlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpAllowlistRequest struct via the builder pattern


### Return type

[**IpAllowlistResponse**](IpAllowlistResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIpAllowlistConfig

> IpAllowlistResponse PatchIpAllowlistConfig(ctx).PatchIpAllowlistConfigRequest(patchIpAllowlistConfigRequest).Execute()

Update IP Allowlist Configuration



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
	patchIpAllowlistConfigRequest := *openapiclient.NewPatchIpAllowlistConfigRequest() // PatchIpAllowlistConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAllowlistBetaApi.PatchIpAllowlistConfig(context.Background()).PatchIpAllowlistConfigRequest(patchIpAllowlistConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistBetaApi.PatchIpAllowlistConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchIpAllowlistConfig`: IpAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `IPAllowlistBetaApi.PatchIpAllowlistConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchIpAllowlistConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchIpAllowlistConfigRequest** | [**PatchIpAllowlistConfigRequest**](PatchIpAllowlistConfigRequest.md) |  | 

### Return type

[**IpAllowlistResponse**](IpAllowlistResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIpAllowlistEntry

> IpAllowlistEntryResponse PatchIpAllowlistEntry(ctx, id).PatchIpAllowlistEntryRequest(patchIpAllowlistEntryRequest).Execute()

Update IP Allowlist Entry Description



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
	id := "id_example" // string | Unique identifier for the allowlist entry
	patchIpAllowlistEntryRequest := *openapiclient.NewPatchIpAllowlistEntryRequest("Updated description") // PatchIpAllowlistEntryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAllowlistBetaApi.PatchIpAllowlistEntry(context.Background(), id).PatchIpAllowlistEntryRequest(patchIpAllowlistEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistBetaApi.PatchIpAllowlistEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchIpAllowlistEntry`: IpAllowlistEntryResponse
	fmt.Fprintf(os.Stdout, "Response from `IPAllowlistBetaApi.PatchIpAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier for the allowlist entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIpAllowlistEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchIpAllowlistEntryRequest** | [**PatchIpAllowlistEntryRequest**](PatchIpAllowlistEntryRequest.md) |  | 

### Return type

[**IpAllowlistEntryResponse**](IpAllowlistEntryResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

