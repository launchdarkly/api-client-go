# \AccountMembersApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMember**](AccountMembersApi.md#DeleteMember) | **Delete** /api/v2/members/{id} | Delete account member
[**GetMember**](AccountMembersApi.md#GetMember) | **Get** /api/v2/members/{id} | Get account member
[**GetMembers**](AccountMembersApi.md#GetMembers) | **Get** /api/v2/members | List account members
[**PatchMember**](AccountMembersApi.md#PatchMember) | **Patch** /api/v2/members/{id} | Modify an account member
[**PostMemberTeams**](AccountMembersApi.md#PostMemberTeams) | **Post** /api/v2/members/{id}/teams | Add a member to teams
[**PostMembers**](AccountMembersApi.md#PostMembers) | **Post** /api/v2/members | Invite new members



## DeleteMember

> DeleteMember(ctx, id).Execute()

Delete account member



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
    id := "id_example" // string | The member ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountMembersApi.DeleteMember(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersApi.DeleteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


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


## GetMember

> Member GetMember(ctx, id).Execute()

Get account member



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
    id := "id_example" // string | The member ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountMembersApi.GetMember(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersApi.GetMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMember`: Member
    fmt.Fprintf(os.Stdout, "Response from `AccountMembersApi.GetMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Member**](Member.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembers

> Members GetMembers(ctx).Limit(limit).Offset(offset).Filter(filter).Sort(sort).Execute()

List account members



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
    limit := int64(789) // int64 | The number of members to return in the response. Defaults to 20. (optional)
    offset := int64(789) // int64 | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next `limit` items. (optional)
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form `field:value`. Supported fields are explained above. (optional)
    sort := "sort_example" // string | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountMembersApi.GetMembers(context.Background()).Limit(limit).Offset(offset).Filter(filter).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersApi.GetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMembers`: Members
    fmt.Fprintf(os.Stdout, "Response from `AccountMembersApi.GetMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The number of members to return in the response. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next &#x60;limit&#x60; items. | 
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form &#x60;field:value&#x60;. Supported fields are explained above. | 
 **sort** | **string** | A comma-separated list of fields to sort by. Fields prefixed by a dash ( - ) sort in descending order. | 

### Return type

[**Members**](Members.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMember

> Member PatchMember(ctx, id).PatchOperation(patchOperation).Execute()

Modify an account member



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
    id := "id_example" // string | The member ID
    patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation("replace", "/exampleField", interface{}(new example value))} // []PatchOperation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountMembersApi.PatchMember(context.Background(), id).PatchOperation(patchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersApi.PatchMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMember`: Member
    fmt.Fprintf(os.Stdout, "Response from `AccountMembersApi.PatchMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

[**Member**](Member.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMemberTeams

> Member PostMemberTeams(ctx, id).MemberTeamsPostInput(memberTeamsPostInput).Execute()

Add a member to teams



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
    id := "id_example" // string | The member ID
    memberTeamsPostInput := *openapiclient.NewMemberTeamsPostInput([]string{"TeamKeys_example"}) // MemberTeamsPostInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountMembersApi.PostMemberTeams(context.Background(), id).MemberTeamsPostInput(memberTeamsPostInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersApi.PostMemberTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMemberTeams`: Member
    fmt.Fprintf(os.Stdout, "Response from `AccountMembersApi.PostMemberTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The member ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMemberTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberTeamsPostInput** | [**MemberTeamsPostInput**](MemberTeamsPostInput.md) |  | 

### Return type

[**Member**](Member.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMembers

> Members PostMembers(ctx).NewMemberForm(newMemberForm).Execute()

Invite new members



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
    newMemberForm := []openapiclient.NewMemberForm{*openapiclient.NewNewMemberForm("Email_example")} // []NewMemberForm | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountMembersApi.PostMembers(context.Background()).NewMemberForm(newMemberForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersApi.PostMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMembers`: Members
    fmt.Fprintf(os.Stdout, "Response from `AccountMembersApi.PostMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newMemberForm** | [**[]NewMemberForm**](NewMemberForm.md) |  | 

### Return type

[**Members**](Members.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

