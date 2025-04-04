# \AnnouncementsApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnnouncementPublic**](AnnouncementsApi.md#CreateAnnouncementPublic) | **Post** /api/v2/announcements | Create an announcement
[**DeleteAnnouncementPublic**](AnnouncementsApi.md#DeleteAnnouncementPublic) | **Delete** /api/v2/announcements/{announcementId} | Delete an announcement
[**GetAnnouncementsPublic**](AnnouncementsApi.md#GetAnnouncementsPublic) | **Get** /api/v2/announcements | Get announcements
[**UpdateAnnouncementPublic**](AnnouncementsApi.md#UpdateAnnouncementPublic) | **Patch** /api/v2/announcements/{announcementId} | Update an announcement



## CreateAnnouncementPublic

> AnnouncementResponse CreateAnnouncementPublic(ctx).CreateAnnouncementBody(createAnnouncementBody).Execute()

Create an announcement



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
    createAnnouncementBody := *openapiclient.NewCreateAnnouncementBody(true, "System Maintenance Notice", "**Important Update:**

Please be aware of the upcoming maintenance scheduled for *October 31st, 2024*. The system will be unavailable from **12:00 AM** to **4:00 AM**.", int64(1731439812), "warning") // CreateAnnouncementBody | Announcement request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnouncementsApi.CreateAnnouncementPublic(context.Background()).CreateAnnouncementBody(createAnnouncementBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsApi.CreateAnnouncementPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnnouncementPublic`: AnnouncementResponse
    fmt.Fprintf(os.Stdout, "Response from `AnnouncementsApi.CreateAnnouncementPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnnouncementPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAnnouncementBody** | [**CreateAnnouncementBody**](CreateAnnouncementBody.md) | Announcement request body | 

### Return type

[**AnnouncementResponse**](AnnouncementResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnnouncementPublic

> DeleteAnnouncementPublic(ctx, announcementId).Execute()

Delete an announcement



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
    announcementId := "1234567890" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnouncementsApi.DeleteAnnouncementPublic(context.Background(), announcementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsApi.DeleteAnnouncementPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**announcementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnouncementPublicRequest struct via the builder pattern


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


## GetAnnouncementsPublic

> GetAnnouncementsPublic200Response GetAnnouncementsPublic(ctx).Status(status).Limit(limit).Offset(offset).Execute()

Get announcements



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
    status := "active" // string | Filter announcements by status. (optional)
    limit := int32(56) // int32 | The number of announcements to return. (optional)
    offset := int32(56) // int32 | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query `limit`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnouncementsApi.GetAnnouncementsPublic(context.Background()).Status(status).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsApi.GetAnnouncementsPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnnouncementsPublic`: GetAnnouncementsPublic200Response
    fmt.Fprintf(os.Stdout, "Response from `AnnouncementsApi.GetAnnouncementsPublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnouncementsPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter announcements by status. | 
 **limit** | **int32** | The number of announcements to return. | 
 **offset** | **int32** | Where to start in the list. Use this with pagination. For example, an offset of 10 skips the first ten items and then returns the next items in the list, up to the query &#x60;limit&#x60;. | 

### Return type

[**GetAnnouncementsPublic200Response**](GetAnnouncementsPublic200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnnouncementPublic

> AnnouncementResponse UpdateAnnouncementPublic(ctx, announcementId).AnnouncementPatchOperation(announcementPatchOperation).Execute()

Update an announcement



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
    announcementId := "1234567890" // string | 
    announcementPatchOperation := []openapiclient.AnnouncementPatchOperation{*openapiclient.NewAnnouncementPatchOperation("replace", "/exampleField")} // []AnnouncementPatchOperation | Update announcement request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnouncementsApi.UpdateAnnouncementPublic(context.Background(), announcementId).AnnouncementPatchOperation(announcementPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsApi.UpdateAnnouncementPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAnnouncementPublic`: AnnouncementResponse
    fmt.Fprintf(os.Stdout, "Response from `AnnouncementsApi.UpdateAnnouncementPublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**announcementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnnouncementPublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **announcementPatchOperation** | [**[]AnnouncementPatchOperation**](AnnouncementPatchOperation.md) | Update announcement request body | 

### Return type

[**AnnouncementResponse**](AnnouncementResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

