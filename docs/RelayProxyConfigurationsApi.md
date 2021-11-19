# \RelayProxyConfigurationsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRelayAutoConfig**](RelayProxyConfigurationsApi.md#DeleteRelayAutoConfig) | **Delete** /api/v2/account/relay-auto-configs/{id} | Delete Relay Proxy config by ID
[**GetRelayProxyConfig**](RelayProxyConfigurationsApi.md#GetRelayProxyConfig) | **Get** /api/v2/account/relay-auto-configs/{id} | Get Relay Proxy config
[**GetRelayProxyConfigs**](RelayProxyConfigurationsApi.md#GetRelayProxyConfigs) | **Get** /api/v2/account/relay-auto-configs | List Relay Proxy configs
[**PatchRelayAutoConfig**](RelayProxyConfigurationsApi.md#PatchRelayAutoConfig) | **Patch** /api/v2/account/relay-auto-configs/{id} | Update a Relay Proxy config
[**PostRelayAutoConfig**](RelayProxyConfigurationsApi.md#PostRelayAutoConfig) | **Post** /api/v2/account/relay-auto-configs | Create a new Relay Proxy config
[**ResetRelayAutoConfig**](RelayProxyConfigurationsApi.md#ResetRelayAutoConfig) | **Post** /api/v2/account/relay-auto-configs/{id}/reset | Reset Relay Proxy configuration key



## DeleteRelayAutoConfig

> DeleteRelayAutoConfig(ctx, id).Execute()

Delete Relay Proxy config by ID



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
    id := "id_example" // string | The relay auto config id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelayProxyConfigurationsApi.DeleteRelayAutoConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelayProxyConfigurationsApi.DeleteRelayAutoConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The relay auto config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRelayAutoConfigRequest struct via the builder pattern


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


## GetRelayProxyConfig

> RelayAutoConfigRep GetRelayProxyConfig(ctx, id).Execute()

Get Relay Proxy config



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
    id := "id_example" // string | The relay auto config id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelayProxyConfigurationsApi.GetRelayProxyConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelayProxyConfigurationsApi.GetRelayProxyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelayProxyConfig`: RelayAutoConfigRep
    fmt.Fprintf(os.Stdout, "Response from `RelayProxyConfigurationsApi.GetRelayProxyConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The relay auto config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelayProxyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RelayAutoConfigRep**](RelayAutoConfigRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelayProxyConfigs

> RelayAutoConfigCollectionRep GetRelayProxyConfigs(ctx).Execute()

List Relay Proxy configs



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
    resp, r, err := api_client.RelayProxyConfigurationsApi.GetRelayProxyConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelayProxyConfigurationsApi.GetRelayProxyConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelayProxyConfigs`: RelayAutoConfigCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `RelayProxyConfigurationsApi.GetRelayProxyConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelayProxyConfigsRequest struct via the builder pattern


### Return type

[**RelayAutoConfigCollectionRep**](RelayAutoConfigCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRelayAutoConfig

> RelayAutoConfigRep PatchRelayAutoConfig(ctx, id).PatchWithComment(patchWithComment).Execute()

Update a Relay Proxy config



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
    id := "id_example" // string | The relay auto config id
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/biscuits", interface{}(Chocolate Digestive))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelayProxyConfigurationsApi.PatchRelayAutoConfig(context.Background(), id).PatchWithComment(patchWithComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelayProxyConfigurationsApi.PatchRelayAutoConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRelayAutoConfig`: RelayAutoConfigRep
    fmt.Fprintf(os.Stdout, "Response from `RelayProxyConfigurationsApi.PatchRelayAutoConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The relay auto config id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRelayAutoConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchWithComment** | [**PatchWithComment**](PatchWithComment.md) |  | 

### Return type

[**RelayAutoConfigRep**](RelayAutoConfigRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRelayAutoConfig

> RelayAutoConfigRep PostRelayAutoConfig(ctx).RelayAutoConfigPost(relayAutoConfigPost).Execute()

Create a new Relay Proxy config



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
    relayAutoConfigPost := *openapiclient.NewRelayAutoConfigPost("Name_example", []openapiclient.StatementRep{*openapiclient.NewStatementRep()}) // RelayAutoConfigPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelayProxyConfigurationsApi.PostRelayAutoConfig(context.Background()).RelayAutoConfigPost(relayAutoConfigPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelayProxyConfigurationsApi.PostRelayAutoConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRelayAutoConfig`: RelayAutoConfigRep
    fmt.Fprintf(os.Stdout, "Response from `RelayProxyConfigurationsApi.PostRelayAutoConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRelayAutoConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relayAutoConfigPost** | [**RelayAutoConfigPost**](RelayAutoConfigPost.md) |  | 

### Return type

[**RelayAutoConfigRep**](RelayAutoConfigRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetRelayAutoConfig

> RelayAutoConfigRep ResetRelayAutoConfig(ctx, id).Expiry(expiry).Execute()

Reset Relay Proxy configuration key



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
    id := "id_example" // string | The Relay Proxy configuration ID
    expiry := int64(789) // int64 | An expiration time for the old Relay Proxy configuration key, expressed as a Unix epoch time in milliseconds. By default, the Relay Proxy configuration will expire immediately. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelayProxyConfigurationsApi.ResetRelayAutoConfig(context.Background(), id).Expiry(expiry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelayProxyConfigurationsApi.ResetRelayAutoConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetRelayAutoConfig`: RelayAutoConfigRep
    fmt.Fprintf(os.Stdout, "Response from `RelayProxyConfigurationsApi.ResetRelayAutoConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Relay Proxy configuration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetRelayAutoConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiry** | **int64** | An expiration time for the old Relay Proxy configuration key, expressed as a Unix epoch time in milliseconds. By default, the Relay Proxy configuration will expire immediately. | 

### Return type

[**RelayAutoConfigRep**](RelayAutoConfigRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

