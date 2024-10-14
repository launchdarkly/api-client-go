# \OAuth2ClientsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOAuth2Client**](OAuth2ClientsApi.md#CreateOAuth2Client) | **Post** /api/v2/oauth/clients | Create a LaunchDarkly OAuth 2.0 client
[**DeleteOAuthClient**](OAuth2ClientsApi.md#DeleteOAuthClient) | **Delete** /api/v2/oauth/clients/{clientId} | Delete OAuth 2.0 client
[**GetOAuthClientById**](OAuth2ClientsApi.md#GetOAuthClientById) | **Get** /api/v2/oauth/clients/{clientId} | Get client by ID
[**GetOAuthClients**](OAuth2ClientsApi.md#GetOAuthClients) | **Get** /api/v2/oauth/clients | Get clients
[**PatchOAuthClient**](OAuth2ClientsApi.md#PatchOAuthClient) | **Patch** /api/v2/oauth/clients/{clientId} | Patch client by ID



## CreateOAuth2Client

> Client CreateOAuth2Client(ctx).OauthClientPost(oauthClientPost).Execute()

Create a LaunchDarkly OAuth 2.0 client



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
    oauthClientPost := *openapiclient.NewOauthClientPost() // OauthClientPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuth2ClientsApi.CreateOAuth2Client(context.Background()).OauthClientPost(oauthClientPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientsApi.CreateOAuth2Client``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOAuth2Client`: Client
    fmt.Fprintf(os.Stdout, "Response from `OAuth2ClientsApi.CreateOAuth2Client`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOAuth2ClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauthClientPost** | [**OauthClientPost**](OauthClientPost.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOAuthClient

> DeleteOAuthClient(ctx, clientId).Execute()

Delete OAuth 2.0 client



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
    clientId := "clientId_example" // string | The client ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuth2ClientsApi.DeleteOAuthClient(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientsApi.DeleteOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOAuthClientRequest struct via the builder pattern


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


## GetOAuthClientById

> Client GetOAuthClientById(ctx, clientId).Execute()

Get client by ID



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
    clientId := "clientId_example" // string | The client ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuth2ClientsApi.GetOAuthClientById(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientsApi.GetOAuthClientById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthClientById`: Client
    fmt.Fprintf(os.Stdout, "Response from `OAuth2ClientsApi.GetOAuthClientById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthClientByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Client**](Client.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthClients

> ClientCollection GetOAuthClients(ctx).Execute()

Get clients



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
    resp, r, err := apiClient.OAuth2ClientsApi.GetOAuthClients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientsApi.GetOAuthClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOAuthClients`: ClientCollection
    fmt.Fprintf(os.Stdout, "Response from `OAuth2ClientsApi.GetOAuthClients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthClientsRequest struct via the builder pattern


### Return type

[**ClientCollection**](ClientCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOAuthClient

> Client PatchOAuthClient(ctx, clientId).PatchOperation(patchOperation).Execute()

Patch client by ID



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
    clientId := "clientId_example" // string | The client ID
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField")} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuth2ClientsApi.PatchOAuthClient(context.Background(), clientId).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuth2ClientsApi.PatchOAuthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchOAuthClient`: Client
    fmt.Fprintf(os.Stdout, "Response from `OAuth2ClientsApi.PatchOAuthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOAuthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

