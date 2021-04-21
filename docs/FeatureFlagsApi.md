# \FeatureFlagsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFeatureFlag**](FeatureFlagsApi.md#CopyFeatureFlag) | **Post** /flags/{projectKey}/{featureFlagKey}/copy | Copies the feature flag configuration from one environment to the same feature flag in another environment.
[**DeleteApprovalRequest**](FeatureFlagsApi.md#DeleteApprovalRequest) | **Delete** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId} | Delete an approval request for a feature flag config
[**DeleteFeatureFlag**](FeatureFlagsApi.md#DeleteFeatureFlag) | **Delete** /flags/{projectKey}/{featureFlagKey} | Delete a feature flag in all environments. Be careful-- only delete feature flags that are no longer being used by your application.
[**DeleteFlagConfigScheduledChanges**](FeatureFlagsApi.md#DeleteFlagConfigScheduledChanges) | **Delete** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{scheduledChangeId} | Delete a scheduled change on a feature flag in an environment.
[**FlagsProjectKeyEnvironmentKeyFeatureFlagKeyDependentFlagsGet**](FeatureFlagsApi.md#FlagsProjectKeyEnvironmentKeyFeatureFlagKeyDependentFlagsGet) | **Get** /flags/{projectKey}/{environmentKey}/{featureFlagKey}/dependent-flags | Get dependent flags for the flag in the environment specified in path parameters
[**FlagsProjectKeyFeatureFlagKeyDependentFlagsGet**](FeatureFlagsApi.md#FlagsProjectKeyFeatureFlagKeyDependentFlagsGet) | **Get** /flags/{projectKey}/{featureFlagKey}/dependent-flags | Get dependent flags across all environments for the flag specified in the path parameters
[**GetApprovalRequest**](FeatureFlagsApi.md#GetApprovalRequest) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId} | Get a single approval request for a feature flag config
[**GetApprovalRequests**](FeatureFlagsApi.md#GetApprovalRequests) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Get all approval requests for a feature flag config
[**GetExpiringUserTargets**](FeatureFlagsApi.md#GetExpiringUserTargets) | **Get** /flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for feature flag
[**GetFeatureFlag**](FeatureFlagsApi.md#GetFeatureFlag) | **Get** /flags/{projectKey}/{featureFlagKey} | Get a single feature flag by key.
[**GetFeatureFlagStatus**](FeatureFlagsApi.md#GetFeatureFlagStatus) | **Get** /flag-statuses/{projectKey}/{environmentKey}/{featureFlagKey} | Get the status for a particular feature flag.
[**GetFeatureFlagStatusAcrossEnvironments**](FeatureFlagsApi.md#GetFeatureFlagStatusAcrossEnvironments) | **Get** /flag-status/{projectKey}/{featureFlagKey} | Get the status for a particular feature flag across environments
[**GetFeatureFlagStatuses**](FeatureFlagsApi.md#GetFeatureFlagStatuses) | **Get** /flag-statuses/{projectKey}/{environmentKey} | Get a list of statuses for all feature flags. The status includes the last time the feature flag was requested, as well as the state of the flag.
[**GetFeatureFlags**](FeatureFlagsApi.md#GetFeatureFlags) | **Get** /flags/{projectKey} | Get a list of all features in the given project.
[**GetFlagConfigScheduledChange**](FeatureFlagsApi.md#GetFlagConfigScheduledChange) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{scheduledChangeId} | Get a scheduled change on a feature flag by id.
[**GetFlagConfigScheduledChanges**](FeatureFlagsApi.md#GetFlagConfigScheduledChanges) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | Get all scheduled workflows for a feature flag by key.
[**GetFlagConfigScheduledChangesConflicts**](FeatureFlagsApi.md#GetFlagConfigScheduledChangesConflicts) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes-conflicts | Lists conflicts between the given instructions and any existing scheduled changes for the feature flag. The actual HTTP verb should be REPORT, not POST.
[**PatchExpiringUserTargets**](FeatureFlagsApi.md#PatchExpiringUserTargets) | **Patch** /flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Update, add, or delete expiring user targets on feature flag
[**PatchFeatureFlag**](FeatureFlagsApi.md#PatchFeatureFlag) | **Patch** /flags/{projectKey}/{featureFlagKey} | Perform a partial update to a feature.
[**PatchFlagConfigScheduledChange**](FeatureFlagsApi.md#PatchFlagConfigScheduledChange) | **Patch** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes/{scheduledChangeId} | Updates an existing scheduled-change on a feature flag in an environment.
[**PostApplyApprovalRequest**](FeatureFlagsApi.md#PostApplyApprovalRequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId}/apply | Apply approval request for a feature flag config
[**PostApprovalRequest**](FeatureFlagsApi.md#PostApprovalRequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId} | Create an approval request for a feature flag config
[**PostFeatureFlag**](FeatureFlagsApi.md#PostFeatureFlag) | **Post** /flags/{projectKey} | Creates a new feature flag.
[**PostFlagConfigScheduledChanges**](FeatureFlagsApi.md#PostFlagConfigScheduledChanges) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/scheduled-changes | Creates a new scheduled change for a feature flag.
[**PostReviewApprovalRequest**](FeatureFlagsApi.md#PostReviewApprovalRequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{approvalRequestId}/review | Review approval request for a feature flag config


# **CopyFeatureFlag**
> FeatureFlag CopyFeatureFlag(ctx, projectKey, featureFlagKey, featureFlagCopyBody)
Copies the feature flag configuration from one environment to the same feature flag in another environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **featureFlagCopyBody** | [**FeatureFlagCopyBody**](FeatureFlagCopyBody.md)| Copy feature flag configurations between environments. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApprovalRequest**
> DeleteApprovalRequest(ctx, projectKey, environmentKey, featureFlagKey, approvalRequestId, optional)
Delete an approval request for a feature flag config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **approvalRequestId** | **string**| The approval request ID | 
 **optional** | ***FeatureFlagsApiDeleteApprovalRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiDeleteApprovalRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **approvalRequestConfigBody** | [**optional.Interface of ApprovalRequestConfigBody**](ApprovalRequestConfigBody.md)| Create a new approval request | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFeatureFlag**
> DeleteFeatureFlag(ctx, projectKey, featureFlagKey)
Delete a feature flag in all environments. Be careful-- only delete feature flags that are no longer being used by your application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFlagConfigScheduledChanges**
> DeleteFlagConfigScheduledChanges(ctx, projectKey, featureFlagKey, environmentKey, scheduledChangeId)
Delete a scheduled change on a feature flag in an environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **scheduledChangeId** | **string**| The id of the scheduled change | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FlagsProjectKeyEnvironmentKeyFeatureFlagKeyDependentFlagsGet**
> DependentFlagsByEnvironment FlagsProjectKeyEnvironmentKeyFeatureFlagKeyDependentFlagsGet(ctx, projectKey, environmentKey, featureFlagKey)
Get dependent flags for the flag in the environment specified in path parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**DependentFlagsByEnvironment**](DependentFlagsByEnvironment.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FlagsProjectKeyFeatureFlagKeyDependentFlagsGet**
> MultiEnvironmentDependentFlags FlagsProjectKeyFeatureFlagKeyDependentFlagsGet(ctx, projectKey, featureFlagKey)
Get dependent flags across all environments for the flag specified in the path parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**MultiEnvironmentDependentFlags**](MultiEnvironmentDependentFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApprovalRequest**
> ApprovalRequests GetApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey, approvalRequestId)
Get a single approval request for a feature flag config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **approvalRequestId** | **string**| The approval request ID | 

### Return type

[**ApprovalRequests**](ApprovalRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApprovalRequests**
> ApprovalRequests GetApprovalRequests(ctx, projectKey, featureFlagKey, environmentKey)
Get all approval requests for a feature flag config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

[**ApprovalRequests**](ApprovalRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpiringUserTargets**
> UserTargetingExpirationForFlags GetExpiringUserTargets(ctx, projectKey, environmentKey, featureFlagKey)
Get expiring user targets for feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**UserTargetingExpirationForFlags**](UserTargetingExpirationForFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlag**
> FeatureFlag GetFeatureFlag(ctx, projectKey, featureFlagKey, optional)
Get a single feature flag by key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
 **optional** | ***FeatureFlagsApiGetFeatureFlagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiGetFeatureFlagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | [**optional.Interface of []string**](string.md)| By default, each feature will include configurations for each environment. You can filter environments with the env query parameter. For example, setting env&#x3D;[\&quot;production\&quot;] will restrict the returned configurations to just your production environment. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatus**
> FeatureFlagStatus GetFeatureFlagStatus(ctx, projectKey, environmentKey, featureFlagKey)
Get the status for a particular feature flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**FeatureFlagStatus**](FeatureFlagStatus.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatusAcrossEnvironments**
> FeatureFlagStatusAcrossEnvironments GetFeatureFlagStatusAcrossEnvironments(ctx, projectKey, featureFlagKey)
Get the status for a particular feature flag across environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**FeatureFlagStatusAcrossEnvironments**](FeatureFlagStatusAcrossEnvironments.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatuses**
> FeatureFlagStatuses GetFeatureFlagStatuses(ctx, projectKey, environmentKey)
Get a list of statuses for all feature flags. The status includes the last time the feature flag was requested, as well as the state of the flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

[**FeatureFlagStatuses**](FeatureFlagStatuses.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlags**
> FeatureFlags GetFeatureFlags(ctx, projectKey, optional)
Get a list of all features in the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
 **optional** | ***FeatureFlagsApiGetFeatureFlagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiGetFeatureFlagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | [**optional.Interface of []string**](string.md)| By default, each feature will include configurations for each environment. You can filter environments with the env query parameter. For example, setting env&#x3D;[\&quot;production\&quot;] will restrict the returned configurations to just your production environment. | 
 **summary** | **optional.Bool**| By default in api version &gt;&#x3D; 1, flags will _not_ include their list of prerequisites, targets or rules.  Set summary&#x3D;0 to include these fields for each flag returned. | 
 **archived** | **optional.Bool**| When set to 1, only archived flags will be included in the list of flags returned.  By default, archived flags are not included in the list of flags. | 
 **limit** | **optional.Float32**| The number of objects to return. Defaults to -1, which returns everything. | 
 **offset** | **optional.Float32**| Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first 10 items and then return the next limit items. | 
 **filter** | **optional.String**| A comma-separated list of filters. Each filter is of the form field:value. | 
 **sort** | **optional.String**| A comma-separated list of fields to sort by. A field prefixed by a - will be sorted in descending order. | 
 **tag** | **optional.String**| Filter by tag. A tag can be used to group flags across projects. | 

### Return type

[**FeatureFlags**](FeatureFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlagConfigScheduledChange**
> FeatureFlagScheduledChange GetFlagConfigScheduledChange(ctx, projectKey, featureFlagKey, environmentKey, scheduledChangeId)
Get a scheduled change on a feature flag by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **scheduledChangeId** | **string**| The id of the scheduled change | 

### Return type

[**FeatureFlagScheduledChange**](FeatureFlagScheduledChange.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlagConfigScheduledChanges**
> FeatureFlagScheduledChanges GetFlagConfigScheduledChanges(ctx, projectKey, featureFlagKey, environmentKey)
Get all scheduled workflows for a feature flag by key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

[**FeatureFlagScheduledChanges**](FeatureFlagScheduledChanges.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlagConfigScheduledChangesConflicts**
> FeatureFlagScheduledChangesConflicts GetFlagConfigScheduledChangesConflicts(ctx, projectKey, featureFlagKey, environmentKey, flagConfigScheduledChangesConflictsBody)
Lists conflicts between the given instructions and any existing scheduled changes for the feature flag. The actual HTTP verb should be REPORT, not POST.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **flagConfigScheduledChangesConflictsBody** | [**FlagConfigScheduledChangesConflictsBody**](FlagConfigScheduledChangesConflictsBody.md)| Used to determine if a semantic patch will result in conflicts with scheduled changes on a feature flag. | 

### Return type

[**FeatureFlagScheduledChangesConflicts**](FeatureFlagScheduledChangesConflicts.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchExpiringUserTargets**
> UserTargetingExpirationForFlags PatchExpiringUserTargets(ctx, projectKey, environmentKey, featureFlagKey, semanticPatchWithComment)
Update, add, or delete expiring user targets on feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **semanticPatchWithComment** | [**interface{}**](interface{}.md)| Requires a Semantic Patch representation of the desired changes to the resource. &#39;https://apidocs.launchdarkly.com/reference#updates-via-semantic-patches&#39;. The addition of comments is also supported. | 

### Return type

[**UserTargetingExpirationForFlags**](UserTargetingExpirationForFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFeatureFlag**
> FeatureFlag PatchFeatureFlag(ctx, projectKey, featureFlagKey, patchComment)
Perform a partial update to a feature.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **patchComment** | [**PatchComment**](PatchComment.md)| Requires a JSON Patch representation of the desired changes to the project, and an optional comment. &#39;http://jsonpatch.com/&#39; Feature flag patches also support JSON Merge Patch format. &#39;https://tools.ietf.org/html/rfc7386&#39; The addition of comments is also supported. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFlagConfigScheduledChange**
> FeatureFlagScheduledChange PatchFlagConfigScheduledChange(ctx, projectKey, featureFlagKey, flagConfigScheduledChangesPatchBody, environmentKey, scheduledChangeId)
Updates an existing scheduled-change on a feature flag in an environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **flagConfigScheduledChangesPatchBody** | [**FlagConfigScheduledChangesPatchBody**](FlagConfigScheduledChangesPatchBody.md)| Update scheduled changes on a feature flag. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **scheduledChangeId** | **string**| The id of the scheduled change | 

### Return type

[**FeatureFlagScheduledChange**](FeatureFlagScheduledChange.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApplyApprovalRequest**
> ApprovalRequests PostApplyApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey, approvalRequestId, approvalRequestApplyConfigBody)
Apply approval request for a feature flag config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **approvalRequestId** | **string**| The approval request ID | 
  **approvalRequestApplyConfigBody** | [**ApprovalRequestApplyConfigBody**](ApprovalRequestApplyConfigBody.md)| Apply an approval request | 

### Return type

[**ApprovalRequests**](ApprovalRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApprovalRequest**
> ApprovalRequest PostApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey, approvalRequestId, optional)
Create an approval request for a feature flag config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **approvalRequestId** | **string**| The approval request ID | 
 **optional** | ***FeatureFlagsApiPostApprovalRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiPostApprovalRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **approvalRequestConfigBody** | [**optional.Interface of ApprovalRequestConfigBody**](ApprovalRequestConfigBody.md)| Create a new approval request | 

### Return type

[**ApprovalRequest**](ApprovalRequest.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFeatureFlag**
> FeatureFlag PostFeatureFlag(ctx, projectKey, featureFlagBody, optional)
Creates a new feature flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagBody** | [**FeatureFlagBody**](FeatureFlagBody.md)| Create a new feature flag. | 
 **optional** | ***FeatureFlagsApiPostFeatureFlagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagsApiPostFeatureFlagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clone** | **optional.String**| The key of the feature flag to be cloned. The key identifies the flag in your code.  For example, setting clone&#x3D;flagKey will copy the full targeting configuration for all environments (including on/off state) from the original flag to the new flag. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFlagConfigScheduledChanges**
> FeatureFlagScheduledChange PostFlagConfigScheduledChanges(ctx, projectKey, flagConfigScheduledChangesPostBody, featureFlagKey, environmentKey)
Creates a new scheduled change for a feature flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **flagConfigScheduledChangesPostBody** | [**FlagConfigScheduledChangesPostBody**](FlagConfigScheduledChangesPostBody.md)| Create scheduled changes on a feature flag. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

[**FeatureFlagScheduledChange**](FeatureFlagScheduledChange.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReviewApprovalRequest**
> ApprovalRequests PostReviewApprovalRequest(ctx, projectKey, featureFlagKey, environmentKey, approvalRequestId, approvalRequestReviewConfigBody)
Review approval request for a feature flag config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **approvalRequestId** | **string**| The approval request ID | 
  **approvalRequestReviewConfigBody** | [**ApprovalRequestReviewConfigBody**](ApprovalRequestReviewConfigBody.md)| Review an approval request | 

### Return type

[**ApprovalRequests**](ApprovalRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

