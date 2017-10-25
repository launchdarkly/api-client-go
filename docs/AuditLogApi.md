# \AuditLogApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogEntries**](AuditLogApi.md#GetAuditLogEntries) | **Get** /auditlog | Fetch a list of all audit log entries
[**GetAuditLogEntry**](AuditLogApi.md#GetAuditLogEntry) | **Get** /auditlog/{resourceId} | Get an audit log entry by ID


# **GetAuditLogEntries**
> AuditLogEntries GetAuditLogEntries()

Fetch a list of all audit log entries


### Parameters
This endpoint does not need any parameter.

### Return type

[**AuditLogEntries**](AuditLogEntries.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuditLogEntry**
> AuditLogEntry GetAuditLogEntry($resourceId)

Get an audit log entry by ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **string**| The resource ID | 

### Return type

[**AuditLogEntry**](AuditLogEntry.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

