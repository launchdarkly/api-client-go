# \AuditLogApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogEntries**](AuditLogApi.md#GetAuditLogEntries) | **Get** /api/v2/auditlog | List audit log feature flag entries
[**GetAuditLogEntry**](AuditLogApi.md#GetAuditLogEntry) | **Get** /api/v2/auditlog/{id} | Get audit log entry



## GetAuditLogEntries

> AuditLogEntryListingRepCollection GetAuditLogEntries(ctx).Before(before).After(after).Q(q).Limit(limit).Spec(spec).Execute()

List audit log feature flag entries



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
    before := int64(789) // int64 | A timestamp filter, expressed as a Unix epoch time in milliseconds.  All entries this returns occurred before the timestamp. (optional)
    after := int64(789) // int64 | A timestamp filter, expressed as a Unix epoch time in milliseconds. All entries this returns occurred after the timestamp. (optional)
    q := "q_example" // string | Text to search for. You can search for the full or partial name of the resource, or full or partial email address of the member who made a change. (optional)
    limit := int64(789) // int64 | A limit on the number of audit log entries that return. Set between 1 and 20. (optional)
    spec := "spec_example" // string | A resource specifier that lets you filter audit log listings by resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogApi.GetAuditLogEntries(context.Background()).Before(before).After(after).Q(q).Limit(limit).Spec(spec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.GetAuditLogEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogEntries`: AuditLogEntryListingRepCollection
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.GetAuditLogEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **int64** | A timestamp filter, expressed as a Unix epoch time in milliseconds.  All entries this returns occurred before the timestamp. | 
 **after** | **int64** | A timestamp filter, expressed as a Unix epoch time in milliseconds. All entries this returns occurred after the timestamp. | 
 **q** | **string** | Text to search for. You can search for the full or partial name of the resource, or full or partial email address of the member who made a change. | 
 **limit** | **int64** | A limit on the number of audit log entries that return. Set between 1 and 20. | 
 **spec** | **string** | A resource specifier that lets you filter audit log listings by resource | 

### Return type

[**AuditLogEntryListingRepCollection**](AuditLogEntryListingRepCollection.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogEntry

> AuditLogEntryRep GetAuditLogEntry(ctx, id).Execute()

Get audit log entry



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
    id := "id_example" // string | The ID of the audit log entry

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogApi.GetAuditLogEntry(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.GetAuditLogEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLogEntry`: AuditLogEntryRep
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.GetAuditLogEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the audit log entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditLogEntryRep**](AuditLogEntryRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

