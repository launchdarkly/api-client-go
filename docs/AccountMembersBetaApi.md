# \AccountMembersBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchMembers**](AccountMembersBetaApi.md#PatchMembers) | **Patch** /api/v2/members | Modify account members



## PatchMembers

> BulkEditMembersRep PatchMembers(ctx).MembersPatchInput(membersPatchInput).Execute()

Modify account members



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
    membersPatchInput := *openapiclient.NewMembersPatchInput([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // MembersPatchInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountMembersBetaApi.PatchMembers(context.Background()).MembersPatchInput(membersPatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountMembersBetaApi.PatchMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMembers`: BulkEditMembersRep
    fmt.Fprintf(os.Stdout, "Response from `AccountMembersBetaApi.PatchMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **membersPatchInput** | [**MembersPatchInput**](MembersPatchInput.md) |  | 

### Return type

[**BulkEditMembersRep**](BulkEditMembersRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

