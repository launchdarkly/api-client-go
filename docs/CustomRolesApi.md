# \CustomRolesApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomRole**](CustomRolesApi.md#DeleteCustomRole) | **Delete** /api/v2/roles/{key} | Delete custom role
[**GetCustomRole**](CustomRolesApi.md#GetCustomRole) | **Get** /api/v2/roles/{key} | Get custom role
[**GetCustomRoles**](CustomRolesApi.md#GetCustomRoles) | **Get** /api/v2/roles | List custom roles
[**PatchCustomRole**](CustomRolesApi.md#PatchCustomRole) | **Patch** /api/v2/roles/{key} | Update custom role
[**PostCustomRole**](CustomRolesApi.md#PostCustomRole) | **Post** /api/v2/roles | Create custom role



## DeleteCustomRole

> DeleteCustomRole(ctx, key).Execute()

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
    key := "key_example" // string | The key of the custom role to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomRolesApi.DeleteCustomRole(context.Background(), key).Execute()
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
**key** | **string** | The key of the custom role to delete | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomRole

> CustomRoleRep GetCustomRole(ctx, key).Execute()

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
    key := "key_example" // string | The custom role's key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomRolesApi.GetCustomRole(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.GetCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomRole`: CustomRoleRep
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.GetCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The custom role&#39;s key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomRoleRep**](CustomRoleRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomRoles

> CustomRoleCollectionRep GetCustomRoles(ctx).Execute()

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomRolesApi.GetCustomRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.GetCustomRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomRoles`: CustomRoleCollectionRep
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.GetCustomRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomRolesRequest struct via the builder pattern


### Return type

[**CustomRoleCollectionRep**](CustomRoleCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCustomRole

> CustomRoleRep PatchCustomRole(ctx, key).JSONPatchElt(jSONPatchElt).Execute()

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
    key := "key_example" // string | The key of the custom role to update
    jSONPatchElt := []openapiclient.JSONPatchElt{*openapiclient.NewJSONPatchElt()} // []JSONPatchElt | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomRolesApi.PatchCustomRole(context.Background(), key).JSONPatchElt(jSONPatchElt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.PatchCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCustomRole`: CustomRoleRep
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.PatchCustomRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the custom role to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSONPatchElt** | [**[]JSONPatchElt**](JSONPatchElt.md) |  | 

### Return type

[**CustomRoleRep**](CustomRoleRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCustomRole

> CustomRoleRep PostCustomRole(ctx).RolesStatementPost(rolesStatementPost).Execute()

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
    rolesStatementPost := []openapiclient.RolesStatementPost{*openapiclient.NewRolesStatementPost()} // []RolesStatementPost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomRolesApi.PostCustomRole(context.Background()).RolesStatementPost(rolesStatementPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomRolesApi.PostCustomRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCustomRole`: CustomRoleRep
    fmt.Fprintf(os.Stdout, "Response from `CustomRolesApi.PostCustomRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCustomRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rolesStatementPost** | [**[]RolesStatementPost**](RolesStatementPost.md) |  | 

### Return type

[**CustomRoleRep**](CustomRoleRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

