# \AccountUsageBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContextsClientsideUsage**](AccountUsageBetaApi.md#GetContextsClientsideUsage) | **Get** /api/v2/usage/clientside-contexts | Get contexts clientside usage
[**GetContextsServersideUsage**](AccountUsageBetaApi.md#GetContextsServersideUsage) | **Get** /api/v2/usage/serverside-contexts | Get contexts serverside usage
[**GetContextsTotalUsage**](AccountUsageBetaApi.md#GetContextsTotalUsage) | **Get** /api/v2/usage/total-contexts | Get contexts total usage
[**GetDataExportEventsUsage**](AccountUsageBetaApi.md#GetDataExportEventsUsage) | **Get** /api/v2/usage/data-export-events | Get data export events usage
[**GetEvaluationsUsage**](AccountUsageBetaApi.md#GetEvaluationsUsage) | **Get** /api/v2/usage/evaluations/{projectKey}/{environmentKey}/{featureFlagKey} | Get evaluations usage
[**GetEventsUsage**](AccountUsageBetaApi.md#GetEventsUsage) | **Get** /api/v2/usage/events/{type} | Get events usage
[**GetExperimentationEventsUsage**](AccountUsageBetaApi.md#GetExperimentationEventsUsage) | **Get** /api/v2/usage/experimentation-events | Get experimentation events usage
[**GetExperimentationKeysUsage**](AccountUsageBetaApi.md#GetExperimentationKeysUsage) | **Get** /api/v2/usage/experimentation-keys | Get experimentation keys usage
[**GetMAUClientsideUsage**](AccountUsageBetaApi.md#GetMAUClientsideUsage) | **Get** /api/v2/usage/clientside-mau | Get MAU clientside usage
[**GetMAUTotalUsage**](AccountUsageBetaApi.md#GetMAUTotalUsage) | **Get** /api/v2/usage/total-mau | Get MAU total usage
[**GetMauSdksByType**](AccountUsageBetaApi.md#GetMauSdksByType) | **Get** /api/v2/usage/mau/sdks | Get MAU SDKs by type
[**GetMauUsage**](AccountUsageBetaApi.md#GetMauUsage) | **Get** /api/v2/usage/mau | Get MAU usage
[**GetMauUsageByCategory**](AccountUsageBetaApi.md#GetMauUsageByCategory) | **Get** /api/v2/usage/mau/bycategory | Get MAU usage by category
[**GetObservabilityErrorsUsage**](AccountUsageBetaApi.md#GetObservabilityErrorsUsage) | **Get** /api/v2/usage/observability/errors | Get observability errors usage
[**GetObservabilityLogsUsage**](AccountUsageBetaApi.md#GetObservabilityLogsUsage) | **Get** /api/v2/usage/observability/logs | Get observability logs usage
[**GetObservabilitySessionsUsage**](AccountUsageBetaApi.md#GetObservabilitySessionsUsage) | **Get** /api/v2/usage/observability/sessions | Get observability sessions usage
[**GetObservabilityTracesUsage**](AccountUsageBetaApi.md#GetObservabilityTracesUsage) | **Get** /api/v2/usage/observability/traces | Get observability traces usage
[**GetServiceConnectionsUsage**](AccountUsageBetaApi.md#GetServiceConnectionsUsage) | **Get** /api/v2/usage/service-connections | Get service connections usage
[**GetStreamUsage**](AccountUsageBetaApi.md#GetStreamUsage) | **Get** /api/v2/usage/streams/{source} | Get stream usage
[**GetStreamUsageBySdkVersion**](AccountUsageBetaApi.md#GetStreamUsageBySdkVersion) | **Get** /api/v2/usage/streams/{source}/bysdkversion | Get stream usage by SDK version
[**GetStreamUsageSdkversion**](AccountUsageBetaApi.md#GetStreamUsageSdkversion) | **Get** /api/v2/usage/streams/{source}/sdkversions | Get stream usage SDK versions



## GetContextsClientsideUsage

> SeriesListRep GetContextsClientsideUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ContextKind(contextKind).SdkName(sdkName).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get contexts clientside usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	contextKind := "contextKind_example" // string | A context kind to filter results by. Can be specified multiple times, one query parameter per context kind. (optional)
	sdkName := "sdkName_example" // string | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. (optional)
	anonymous := "anonymous_example" // string | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.<br/>Valid values: `true`, `false`. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. `contextKind` is always included as a grouping dimension. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `projectId`, `environmentId`, `sdkName`, `sdkAppId`, `anonymousV2`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetContextsClientsideUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ContextKind(contextKind).SdkName(sdkName).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetContextsClientsideUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextsClientsideUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetContextsClientsideUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContextsClientsideUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **contextKind** | **string** | A context kind to filter results by. Can be specified multiple times, one query parameter per context kind. | 
 **sdkName** | **string** | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. | 
 **anonymous** | **string** | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.&lt;br/&gt;Valid values: &#x60;true&#x60;, &#x60;false&#x60;. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. &#x60;contextKind&#x60; is always included as a grouping dimension. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;projectId&#x60;, &#x60;environmentId&#x60;, &#x60;sdkName&#x60;, &#x60;sdkAppId&#x60;, &#x60;anonymousV2&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextsServersideUsage

> SeriesListRep GetContextsServersideUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ContextKind(contextKind).SdkName(sdkName).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get contexts serverside usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	contextKind := "contextKind_example" // string | A context kind to filter results by. Can be specified multiple times, one query parameter per context kind. (optional)
	sdkName := "sdkName_example" // string | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. (optional)
	anonymous := "anonymous_example" // string | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.<br/>Valid values: `true`, `false`. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. `contextKind` is always included as a grouping dimension. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `projectId`, `environmentId`, `sdkName`, `sdkAppId`, `anonymousV2`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetContextsServersideUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ContextKind(contextKind).SdkName(sdkName).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetContextsServersideUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextsServersideUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetContextsServersideUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContextsServersideUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **contextKind** | **string** | A context kind to filter results by. Can be specified multiple times, one query parameter per context kind. | 
 **sdkName** | **string** | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. | 
 **anonymous** | **string** | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.&lt;br/&gt;Valid values: &#x60;true&#x60;, &#x60;false&#x60;. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. &#x60;contextKind&#x60; is always included as a grouping dimension. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;projectId&#x60;, &#x60;environmentId&#x60;, &#x60;sdkName&#x60;, &#x60;sdkAppId&#x60;, &#x60;anonymousV2&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextsTotalUsage

> SeriesListRep GetContextsTotalUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ContextKind(contextKind).SdkName(sdkName).SdkType(sdkType).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get contexts total usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	contextKind := "contextKind_example" // string | A context kind to filter results by. Can be specified multiple times, one query parameter per context kind. (optional)
	sdkName := "sdkName_example" // string | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. (optional)
	sdkType := "sdkType_example" // string | An SDK type to filter results by. Can be specified multiple times, one query parameter per SDK type. (optional)
	anonymous := "anonymous_example" // string | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.<br/>Valid values: `true`, `false`. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. `contextKind` is always included as a grouping dimension. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `projectId`, `environmentId`, `sdkName`, `sdkType`, `sdkAppId`, `anonymousV2`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetContextsTotalUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ContextKind(contextKind).SdkName(sdkName).SdkType(sdkType).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetContextsTotalUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextsTotalUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetContextsTotalUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContextsTotalUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **contextKind** | **string** | A context kind to filter results by. Can be specified multiple times, one query parameter per context kind. | 
 **sdkName** | **string** | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. | 
 **sdkType** | **string** | An SDK type to filter results by. Can be specified multiple times, one query parameter per SDK type. | 
 **anonymous** | **string** | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.&lt;br/&gt;Valid values: &#x60;true&#x60;, &#x60;false&#x60;. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. &#x60;contextKind&#x60; is always included as a grouping dimension. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;projectId&#x60;, &#x60;environmentId&#x60;, &#x60;sdkName&#x60;, &#x60;sdkType&#x60;, &#x60;sdkAppId&#x60;, &#x60;anonymousV2&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataExportEventsUsage

> SeriesListRep GetDataExportEventsUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).EventKind(eventKind).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get data export events usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	eventKind := "eventKind_example" // string | An event kind to filter results by. Can be specified multiple times, one query parameter per event kind. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `environmentId`, `eventKind`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. `monthly` granularity is only supported with the **month_to_date** aggregation type.<br/>Valid values: `daily`, `hourly`, `monthly`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetDataExportEventsUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).EventKind(eventKind).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetDataExportEventsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataExportEventsUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetDataExportEventsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDataExportEventsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **eventKind** | **string** | An event kind to filter results by. Can be specified multiple times, one query parameter per event kind. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;environmentId&#x60;, &#x60;eventKind&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. &#x60;monthly&#x60; granularity is only supported with the **month_to_date** aggregation type.&lt;br/&gt;Valid values: &#x60;daily&#x60;, &#x60;hourly&#x60;, &#x60;monthly&#x60;. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvaluationsUsage

> SeriesListRep GetEvaluationsUsage(ctx, projectKey, environmentKey, featureFlagKey).From(from).To(to).Tz(tz).Execute()

Get evaluations usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	environmentKey := "environmentKey_example" // string | The environment key
	featureFlagKey := "featureFlagKey_example" // string | The feature flag key
	from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
	tz := "tz_example" // string | The timezone to use for breaks between days when returning daily data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetEvaluationsUsage(context.Background(), projectKey, environmentKey, featureFlagKey).From(from).To(to).Tz(tz).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetEvaluationsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvaluationsUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetEvaluationsUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**environmentKey** | **string** | The environment key | 
**featureFlagKey** | **string** | The feature flag key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvaluationsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **tz** | **string** | The timezone to use for breaks between days when returning daily data. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsUsage

> SeriesListRep GetEventsUsage(ctx, type_).From(from).To(to).Execute()

Get events usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	type_ := "type__example" // string | The type of event to retrieve. Must be either `received` or `published`.
	from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 24 hours ago. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetEventsUsage(context.Background(), type_).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetEventsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetEventsUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of event to retrieve. Must be either &#x60;received&#x60; or &#x60;published&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 24 hours ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentationEventsUsage

> SeriesListRep GetExperimentationEventsUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).EventKey(eventKey).EventKind(eventKind).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get experimentation events usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	eventKey := "eventKey_example" // string | An event key to filter results by. Can be specified multiple times, one query parameter per event key. (optional)
	eventKind := "eventKind_example" // string | An event kind to filter results by. Can be specified multiple times, one query parameter per event kind. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `environmentId`, `eventKey`, `eventKind`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. `monthly` granularity is only supported with the **month_to_date** aggregation type.<br/>Valid values: `daily`, `hourly`, `monthly`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetExperimentationEventsUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).EventKey(eventKey).EventKind(eventKind).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetExperimentationEventsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentationEventsUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetExperimentationEventsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentationEventsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **eventKey** | **string** | An event key to filter results by. Can be specified multiple times, one query parameter per event key. | 
 **eventKind** | **string** | An event kind to filter results by. Can be specified multiple times, one query parameter per event kind. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;environmentId&#x60;, &#x60;eventKey&#x60;, &#x60;eventKind&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. &#x60;monthly&#x60; granularity is only supported with the **month_to_date** aggregation type.&lt;br/&gt;Valid values: &#x60;daily&#x60;, &#x60;hourly&#x60;, &#x60;monthly&#x60;. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExperimentationKeysUsage

> SeriesListRep GetExperimentationKeysUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ExperimentId(experimentId).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get experimentation keys usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	experimentId := "experimentId_example" // string | An experiment ID to filter results by. Can be specified multiple times, one query parameter per experiment ID. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `projectId`, `environmentId`, `experimentId`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. `monthly` granularity is only supported with the **month_to_date** aggregation type.<br/>Valid values: `daily`, `hourly`, `monthly`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetExperimentationKeysUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ExperimentId(experimentId).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetExperimentationKeysUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExperimentationKeysUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetExperimentationKeysUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExperimentationKeysUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **experimentId** | **string** | An experiment ID to filter results by. Can be specified multiple times, one query parameter per experiment ID. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;projectId&#x60;, &#x60;environmentId&#x60;, &#x60;experimentId&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. &#x60;monthly&#x60; granularity is only supported with the **month_to_date** aggregation type.&lt;br/&gt;Valid values: &#x60;daily&#x60;, &#x60;hourly&#x60;, &#x60;monthly&#x60;. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMAUClientsideUsage

> SeriesListRep GetMAUClientsideUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).SdkName(sdkName).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get MAU clientside usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	sdkName := "sdkName_example" // string | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. (optional)
	anonymous := "anonymous_example" // string | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.<br/>Valid values: `true`, `false`. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `projectId`, `environmentId`, `sdkName`, `sdkAppId`, `anonymousV2`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetMAUClientsideUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).SdkName(sdkName).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMAUClientsideUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMAUClientsideUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMAUClientsideUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMAUClientsideUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **sdkName** | **string** | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. | 
 **anonymous** | **string** | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.&lt;br/&gt;Valid values: &#x60;true&#x60;, &#x60;false&#x60;. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;projectId&#x60;, &#x60;environmentId&#x60;, &#x60;sdkName&#x60;, &#x60;sdkAppId&#x60;, &#x60;anonymousV2&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMAUTotalUsage

> SeriesListRep GetMAUTotalUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).SdkName(sdkName).SdkType(sdkType).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get MAU total usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	sdkName := "sdkName_example" // string | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. (optional)
	sdkType := "sdkType_example" // string | An SDK type to filter results by. Can be specified multiple times, one query parameter per SDK type. (optional)
	anonymous := "anonymous_example" // string | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.<br/>Valid values: `true`, `false`. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `projectId`, `environmentId`, `sdkName`, `sdkType`, `sdkAppId`, `anonymousV2`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetMAUTotalUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).SdkName(sdkName).SdkType(sdkType).Anonymous(anonymous).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMAUTotalUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMAUTotalUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMAUTotalUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMAUTotalUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **sdkName** | **string** | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. | 
 **sdkType** | **string** | An SDK type to filter results by. Can be specified multiple times, one query parameter per SDK type. | 
 **anonymous** | **string** | An anonymous value to filter results by. Can be specified multiple times, one query parameter per anonymous value.&lt;br/&gt;Valid values: &#x60;true&#x60;, &#x60;false&#x60;. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;projectId&#x60;, &#x60;environmentId&#x60;, &#x60;sdkName&#x60;, &#x60;sdkType&#x60;, &#x60;sdkAppId&#x60;, &#x60;anonymousV2&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMauSdksByType

> SdkListRep GetMauSdksByType(ctx).From(from).To(to).Sdktype(sdktype).Execute()

Get MAU SDKs by type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The data returned starts from this timestamp. Defaults to seven days ago. The timestamp is in Unix milliseconds, for example, 1656694800000. (optional)
	to := "to_example" // string | The data returned ends at this timestamp. Defaults to the current time. The timestamp is in Unix milliseconds, for example, 1657904400000. (optional)
	sdktype := "sdktype_example" // string | The type of SDK with monthly active users (MAU) to list. Must be either `client` or `server`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetMauSdksByType(context.Background()).From(from).To(to).Sdktype(sdktype).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMauSdksByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMauSdksByType`: SdkListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMauSdksByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMauSdksByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The data returned starts from this timestamp. Defaults to seven days ago. The timestamp is in Unix milliseconds, for example, 1656694800000. | 
 **to** | **string** | The data returned ends at this timestamp. Defaults to the current time. The timestamp is in Unix milliseconds, for example, 1657904400000. | 
 **sdktype** | **string** | The type of SDK with monthly active users (MAU) to list. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Return type

[**SdkListRep**](SdkListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMauUsage

> SeriesListRep GetMauUsage(ctx).From(from).To(to).Project(project).Environment(environment).Sdktype(sdktype).Sdk(sdk).Anonymous(anonymous).Groupby(groupby).AggregationType(aggregationType).ContextKind(contextKind).Execute()

Get MAU usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
	project := "project_example" // string | A project key to filter results to. Can be specified multiple times, one query parameter per project key, to view data for multiple projects. (optional)
	environment := "environment_example" // string | An environment key to filter results to. When using this parameter, exactly one project key must also be set. Can be specified multiple times as separate query parameters to view data for multiple environments within a single project. (optional)
	sdktype := "sdktype_example" // string | An SDK type to filter results to. Can be specified multiple times, one query parameter per SDK type. Valid values: client, server (optional)
	sdk := "sdk_example" // string | An SDK name to filter results to. Can be specified multiple times, one query parameter per SDK. (optional)
	anonymous := "anonymous_example" // string | If specified, filters results to either anonymous or nonanonymous users. (optional)
	groupby := "groupby_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions (for example, to group by both project and SDK). Valid values: project, environment, sdktype, sdk, anonymous, contextKind, sdkAppId (optional)
	aggregationType := "aggregationType_example" // string | If specified, queries for rolling 30-day, month-to-date, or daily incremental counts. Default is rolling 30-day. Valid values: rolling_30d, month_to_date, daily_incremental (optional)
	contextKind := "contextKind_example" // string | Filters results to the specified context kinds. Can be specified multiple times, one query parameter per context kind. If not set, queries for the user context kind. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetMauUsage(context.Background()).From(from).To(to).Project(project).Environment(environment).Sdktype(sdktype).Sdk(sdk).Anonymous(anonymous).Groupby(groupby).AggregationType(aggregationType).ContextKind(contextKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMauUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMauUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMauUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMauUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **project** | **string** | A project key to filter results to. Can be specified multiple times, one query parameter per project key, to view data for multiple projects. | 
 **environment** | **string** | An environment key to filter results to. When using this parameter, exactly one project key must also be set. Can be specified multiple times as separate query parameters to view data for multiple environments within a single project. | 
 **sdktype** | **string** | An SDK type to filter results to. Can be specified multiple times, one query parameter per SDK type. Valid values: client, server | 
 **sdk** | **string** | An SDK name to filter results to. Can be specified multiple times, one query parameter per SDK. | 
 **anonymous** | **string** | If specified, filters results to either anonymous or nonanonymous users. | 
 **groupby** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions (for example, to group by both project and SDK). Valid values: project, environment, sdktype, sdk, anonymous, contextKind, sdkAppId | 
 **aggregationType** | **string** | If specified, queries for rolling 30-day, month-to-date, or daily incremental counts. Default is rolling 30-day. Valid values: rolling_30d, month_to_date, daily_incremental | 
 **contextKind** | **string** | Filters results to the specified context kinds. Can be specified multiple times, one query parameter per context kind. If not set, queries for the user context kind. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMauUsageByCategory

> SeriesListRep GetMauUsageByCategory(ctx).From(from).To(to).Execute()

Get MAU usage by category



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetMauUsageByCategory(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetMauUsageByCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMauUsageByCategory`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetMauUsageByCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMauUsageByCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObservabilityErrorsUsage

> SeriesListRep GetObservabilityErrorsUsage(ctx).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()

Get observability errors usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetObservabilityErrorsUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetObservabilityErrorsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObservabilityErrorsUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetObservabilityErrorsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObservabilityErrorsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObservabilityLogsUsage

> SeriesListRep GetObservabilityLogsUsage(ctx).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()

Get observability logs usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetObservabilityLogsUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetObservabilityLogsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObservabilityLogsUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetObservabilityLogsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObservabilityLogsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObservabilitySessionsUsage

> SeriesListRep GetObservabilitySessionsUsage(ctx).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()

Get observability sessions usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetObservabilitySessionsUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetObservabilitySessionsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObservabilitySessionsUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetObservabilitySessionsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObservabilitySessionsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObservabilityTracesUsage

> SeriesListRep GetObservabilityTracesUsage(ctx).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()

Get observability traces usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. Valid values depend on `aggregationType`: **month_to_date** supports `daily` and `monthly`; **incremental** and **rolling_30d** support `daily` only. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`, `rolling_30d`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetObservabilityTracesUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).Granularity(granularity).AggregationType(aggregationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetObservabilityTracesUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObservabilityTracesUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetObservabilityTracesUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObservabilityTracesUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix seconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix seconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. Valid values depend on &#x60;aggregationType&#x60;: **month_to_date** supports &#x60;daily&#x60; and &#x60;monthly&#x60;; **incremental** and **rolling_30d** support &#x60;daily&#x60; only. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;, &#x60;rolling_30d&#x60;. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConnectionsUsage

> SeriesListRepFloat GetServiceConnectionsUsage(ctx).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ConnectionType(connectionType).RelayVersion(relayVersion).SdkName(sdkName).SdkVersion(sdkVersion).SdkType(sdkType).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()

Get service connections usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	from := "from_example" // string | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. (optional)
	projectKey := "projectKey_example" // string | A project key to filter results by. Can be specified multiple times, one query parameter per project key. (optional)
	environmentKey := "environmentKey_example" // string | An environment key to filter results by. If specified, exactly one `projectKey` must be provided. Can be specified multiple times, one query parameter per environment key. (optional)
	connectionType := "connectionType_example" // string | A connection type to filter results by. Can be specified multiple times, one query parameter per connection type. (optional)
	relayVersion := "relayVersion_example" // string | A relay version to filter results by. Can be specified multiple times, one query parameter per relay version. (optional)
	sdkName := "sdkName_example" // string | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. (optional)
	sdkVersion := "sdkVersion_example" // string | An SDK version to filter results by. Can be specified multiple times, one query parameter per SDK version. (optional)
	sdkType := "sdkType_example" // string | An SDK type to filter results by. Can be specified multiple times, one query parameter per SDK type. (optional)
	groupBy := "groupBy_example" // string | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.<br/>Valid values: `projectId`, `environmentId`, `connectionType`, `relayVersion`, `sdkName`, `sdkVersion`, `sdkType`. (optional)
	aggregationType := "aggregationType_example" // string | Specifies the aggregation method. Defaults to `month_to_date`.<br/>Valid values: `month_to_date`, `incremental`. (optional)
	granularity := "granularity_example" // string | Specifies the data granularity. Defaults to `daily`. `monthly` granularity is only supported with the **month_to_date** aggregation type.<br/>Valid values: `daily`, `hourly`, `monthly`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetServiceConnectionsUsage(context.Background()).From(from).To(to).ProjectKey(projectKey).EnvironmentKey(environmentKey).ConnectionType(connectionType).RelayVersion(relayVersion).SdkName(sdkName).SdkVersion(sdkVersion).SdkType(sdkType).GroupBy(groupBy).AggregationType(aggregationType).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetServiceConnectionsUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceConnectionsUsage`: SeriesListRepFloat
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetServiceConnectionsUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceConnectionsUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The series of data returned starts from this timestamp (Unix milliseconds). Defaults to the beginning of the current month. | 
 **to** | **string** | The series of data returned ends at this timestamp (Unix milliseconds). Defaults to the current time. | 
 **projectKey** | **string** | A project key to filter results by. Can be specified multiple times, one query parameter per project key. | 
 **environmentKey** | **string** | An environment key to filter results by. If specified, exactly one &#x60;projectKey&#x60; must be provided. Can be specified multiple times, one query parameter per environment key. | 
 **connectionType** | **string** | A connection type to filter results by. Can be specified multiple times, one query parameter per connection type. | 
 **relayVersion** | **string** | A relay version to filter results by. Can be specified multiple times, one query parameter per relay version. | 
 **sdkName** | **string** | An SDK name to filter results by. Can be specified multiple times, one query parameter per SDK name. | 
 **sdkVersion** | **string** | An SDK version to filter results by. Can be specified multiple times, one query parameter per SDK version. | 
 **sdkType** | **string** | An SDK type to filter results by. Can be specified multiple times, one query parameter per SDK type. | 
 **groupBy** | **string** | If specified, returns data for each distinct value of the given field. Can be specified multiple times to group data by multiple dimensions, one query parameter per dimension.&lt;br/&gt;Valid values: &#x60;projectId&#x60;, &#x60;environmentId&#x60;, &#x60;connectionType&#x60;, &#x60;relayVersion&#x60;, &#x60;sdkName&#x60;, &#x60;sdkVersion&#x60;, &#x60;sdkType&#x60;. | 
 **aggregationType** | **string** | Specifies the aggregation method. Defaults to &#x60;month_to_date&#x60;.&lt;br/&gt;Valid values: &#x60;month_to_date&#x60;, &#x60;incremental&#x60;. | 
 **granularity** | **string** | Specifies the data granularity. Defaults to &#x60;daily&#x60;. &#x60;monthly&#x60; granularity is only supported with the **month_to_date** aggregation type.&lt;br/&gt;Valid values: &#x60;daily&#x60;, &#x60;hourly&#x60;, &#x60;monthly&#x60;. | 

### Return type

[**SeriesListRepFloat**](SeriesListRepFloat.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamUsage

> SeriesListRep GetStreamUsage(ctx, source).From(from).To(to).Tz(tz).Execute()

Get stream usage



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	source := "source_example" // string | The source of streaming connections to describe. Must be either `client` or `server`.
	from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 30 days ago. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
	tz := "tz_example" // string | The timezone to use for breaks between days when returning daily data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetStreamUsage(context.Background(), source).From(from).To(to).Tz(tz).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetStreamUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamUsage`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetStreamUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | The source of streaming connections to describe. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 30 days ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **tz** | **string** | The timezone to use for breaks between days when returning daily data. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamUsageBySdkVersion

> SeriesListRep GetStreamUsageBySdkVersion(ctx, source).From(from).To(to).Tz(tz).Sdk(sdk).Version(version).Execute()

Get stream usage by SDK version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	source := "source_example" // string | The source of streaming connections to describe. Must be either `client` or `server`.
	from := "from_example" // string | The series of data returned starts from this timestamp. Defaults to 24 hours ago. (optional)
	to := "to_example" // string | The series of data returned ends at this timestamp. Defaults to the current time. (optional)
	tz := "tz_example" // string | The timezone to use for breaks between days when returning daily data. (optional)
	sdk := "sdk_example" // string | If included, this filters the returned series to only those that match this SDK name. (optional)
	version := "version_example" // string | If included, this filters the returned series to only those that match this SDK version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetStreamUsageBySdkVersion(context.Background(), source).From(from).To(to).Tz(tz).Sdk(sdk).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetStreamUsageBySdkVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamUsageBySdkVersion`: SeriesListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetStreamUsageBySdkVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | The source of streaming connections to describe. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamUsageBySdkVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The series of data returned starts from this timestamp. Defaults to 24 hours ago. | 
 **to** | **string** | The series of data returned ends at this timestamp. Defaults to the current time. | 
 **tz** | **string** | The timezone to use for breaks between days when returning daily data. | 
 **sdk** | **string** | If included, this filters the returned series to only those that match this SDK name. | 
 **version** | **string** | If included, this filters the returned series to only those that match this SDK version. | 

### Return type

[**SeriesListRep**](SeriesListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStreamUsageSdkversion

> SdkVersionListRep GetStreamUsageSdkversion(ctx, source).Execute()

Get stream usage SDK versions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	source := "source_example" // string | The source of streaming connections to describe. Must be either `client` or `server`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountUsageBetaApi.GetStreamUsageSdkversion(context.Background(), source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountUsageBetaApi.GetStreamUsageSdkversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStreamUsageSdkversion`: SdkVersionListRep
	fmt.Fprintf(os.Stdout, "Response from `AccountUsageBetaApi.GetStreamUsageSdkversion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | The source of streaming connections to describe. Must be either &#x60;client&#x60; or &#x60;server&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamUsageSdkversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SdkVersionListRep**](SdkVersionListRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

