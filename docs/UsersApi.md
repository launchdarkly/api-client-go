# \UsersApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /api/v2/users/{projKey}/{envKey}/{key} | Delete user
[**GetSearchUsers**](UsersApi.md#GetSearchUsers) | **Get** /api/v2/user-search/{projKey}/{envKey} | Find users
[**GetUser**](UsersApi.md#GetUser) | **Get** /api/v2/users/{projKey}/{envKey}/{key} | Get user
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /api/v2/users/{projKey}/{envKey} | List users



## DeleteUser

> DeleteUser(ctx, projKey, envKey, key).Execute()

Delete user



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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    key := "key_example" // string | The user key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DeleteUser(context.Background(), projKey, envKey, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**key** | **string** | The user key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetSearchUsers

> UserListRep GetSearchUsers(ctx, projKey, envKey).Q(q).Limit(limit).Offset(offset).After(after).Execute()

Find users



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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    q := "q_example" // string | Full-text search for users based on name, first name, last name, e-mail address, or key (optional)
    limit := int64(789) // int64 | Specifies the maximum number of items in the collection to return (max: 50, default: 20) (optional)
    offset := int64(789) // int64 | Specifies the first item to return in the collection (optional)
    after := int64(789) // int64 | A unix epoch time in milliseconds specifying the maximum last time a user requested a feature flag from LaunchDarkly (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetSearchUsers(context.Background(), projKey, envKey).Q(q).Limit(limit).Offset(offset).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetSearchUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchUsers`: UserListRep
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetSearchUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Full-text search for users based on name, first name, last name, e-mail address, or key | 
 **limit** | **int64** | Specifies the maximum number of items in the collection to return (max: 50, default: 20) | 
 **offset** | **int64** | Specifies the first item to return in the collection | 
 **after** | **int64** | A unix epoch time in milliseconds specifying the maximum last time a user requested a feature flag from LaunchDarkly | 

### Return type

[**UserListRep**](UserListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserRep GetUser(ctx, projKey, envKey, key).Limit(limit).Execute()

Get user



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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    key := "key_example" // string | The user key
    limit := "limit_example" // string | The number of elements to return per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUser(context.Background(), projKey, envKey, key).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserRep
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 
**key** | **string** | The user key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **string** | The number of elements to return per page | 

### Return type

[**UserRep**](UserRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> UserListRep GetUsers(ctx, projKey, envKey).Limit(limit).Execute()

List users



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
    projKey := "projKey_example" // string | The project key
    envKey := "envKey_example" // string | The environment key
    limit := int64(789) // int64 | The number of elements to return per page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUsers(context.Background(), projKey, envKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: UserListRep
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projKey** | **string** | The project key | 
**envKey** | **string** | The environment key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int64** | The number of elements to return per page | 

### Return type

[**UserListRep**](UserListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

