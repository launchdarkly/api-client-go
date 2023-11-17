# \CustomRolesApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomRole**](CustomRolesApi.md#DeleteCustomRole) | **Delete** /api/v2/roles/{customRoleKey} | Delete custom role
[**GetCustomRole**](CustomRolesApi.md#GetCustomRole) | **Get** /api/v2/roles/{customRoleKey} | Get custom role
[**GetCustomRoles**](CustomRolesApi.md#GetCustomRoles) | **Get** /api/v2/roles | List custom roles
[**PatchCustomRole**](CustomRolesApi.md#PatchCustomRole) | **Patch** /api/v2/roles/{customRoleKey} | Update custom role
[**PostCustomRole**](CustomRolesApi.md#PostCustomRole) | **Post** /api/v2/roles | Create custom role



## DeleteCustomRole

> DeleteCustomRole(ctx, customRoleKey).Execute()

Delete custom role



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
    customRoleKey := "customRoleKey_example" // string | The custom role key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomRolesApi.DeleteCustomRole(context.Background(), customRoleKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.DeleteCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customRoleKey** | **string** | The custom role key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomRoleRequest struct via the builder pattern


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


## GetCustomRole

> CustomRole GetCustomRole(ctx, customRoleKey).Execute()

Get custom role



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
    customRoleKey := "customRoleKey_example" // string | The custom role key or ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomRolesApi.GetCustomRole(context.Background(), customRoleKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.GetCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomRole`: CustomRole
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.GetCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customRoleKey** | **string** | The custom role key or ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomRole**](CustomRole.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomRoles

> CustomRoles GetCustomRoles(ctx).Execute()

List custom roles



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
    resp, r, err := apiClient.CustomRolesApi.GetCustomRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.GetCustomRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomRoles`: CustomRoles
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.GetCustomRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomRolesRequest struct via the builder pattern


### Return type

[**CustomRoles**](CustomRoles.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCustomRole

> CustomRole PatchCustomRole(ctx, customRoleKey).PatchWithComment(patchWithComment).Execute()

Update custom role



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
    customRoleKey := "customRoleKey_example" // string | The custom role key
    patchWithComment := *openapiclient.NewPatchWithComment([]openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))}) // PatchWithComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomRolesApi.PatchCustomRole(context.Background(), customRoleKey).PatchWithComment(patchWithComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.PatchCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCustomRole`: CustomRole
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.PatchCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customRoleKey** | **string** | The custom role key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchWithComment** | [**PatchWithComment**](PatchWithComment.md) |  | 

### Return type

[**CustomRole**](CustomRole.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCustomRole

> CustomRole PostCustomRole(ctx).CustomRolePost(customRolePost).Execute()

Create custom role



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
    customRolePost := *openapiclient.NewCustomRolePost("Ops team", "role-key-123abc", []openapiclient.StatementPost{*openapiclient.NewStatementPost("allow")}) // CustomRolePost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomRolesApi.PostCustomRole(context.Background()).CustomRolePost(customRolePost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.PostCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCustomRole`: CustomRole
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.PostCustomRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customRolePost** | [**CustomRolePost**](CustomRolePost.md) |  | 

### Return type

[**CustomRole**](CustomRole.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

