# \TagsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTags**](TagsApi.md#GetTags) | **Get** /api/v2/tags | List tags



## GetTags

> TagsCollection GetTags(ctx).Kind(kind).Pre(pre).Archived(archived).Limit(limit).Offset(offset).AsOf(asOf).Execute()

List tags



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
    kind := []string{"Inner_example"} // []string | Fetch tags associated with the specified resource type. Options are `flag`, `project`, `environment`, `segment`. Returns all types by default. (optional)
    pre := "pre_example" // string | Return tags with the specified prefix (optional)
    archived := true // bool | Whether or not to return archived flags (optional)
    limit := int32(56) // int32 | The number of tags to return. Maximum is 1000. (optional)
    offset := int32(56) // int32 | The index of the first tag to return. Default is 0. (optional)
    asOf := "asOf_example" // string | The time to retrieve tags as of. Default is the current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsApi.GetTags(context.Background()).Kind(kind).Pre(pre).Archived(archived).Limit(limit).Offset(offset).AsOf(asOf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: TagsCollection
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **[]string** | Fetch tags associated with the specified resource type. Options are &#x60;flag&#x60;, &#x60;project&#x60;, &#x60;environment&#x60;, &#x60;segment&#x60;. Returns all types by default. | 
 **pre** | **string** | Return tags with the specified prefix | 
 **archived** | **bool** | Whether or not to return archived flags | 
 **limit** | **int32** | The number of tags to return. Maximum is 1000. | 
 **offset** | **int32** | The index of the first tag to return. Default is 0. | 
 **asOf** | **string** | The time to retrieve tags as of. Default is the current time. | 

### Return type

[**TagsCollection**](TagsCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

