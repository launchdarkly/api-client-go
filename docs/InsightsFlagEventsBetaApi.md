# \InsightsFlagEventsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlagEvents**](InsightsFlagEventsBetaApi.md#GetFlagEvents) | **Get** /api/v2/engineering-insights/flag-events | List flag events



## GetFlagEvents

> FlagEventCollectionRep GetFlagEvents(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Query(query).ImpactSize(impactSize).HasExperiments(hasExperiments).Global(global).Expand(expand).Limit(limit).From(from).To(to).After(after).Before(before).Execute()

List flag events



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
	applicationKey := "applicationKey_example" // string | Comma separated list of application keys (optional)
	query := "query_example" // string | Filter events by flag key (optional)
	impactSize := "impactSize_example" // string | Filter events by impact size. A small impact created a less than 20% change in the proportion of end users receiving one or more flag variations. A medium impact created between a 20%-80% change. A large impact created a more than 80% change. Options: `none`, `small`, `medium`, `large` (optional)
	hasExperiments := true // bool | Filter events to those associated with an experiment (`true`) or without an experiment (`false`) (optional)
	global := "global_example" // string | Filter to include or exclude global events. Default value is `include`. Options: `include`, `exclude` (optional)
	expand := "expand_example" // string | Expand properties in response. Options: `experiments` (optional)
	limit := int64(789) // int64 | The number of deployments to return. Default is 20. Maximum allowed is 100. (optional)
	from := int64(789) // int64 | Unix timestamp in milliseconds. Default value is 7 days ago. (optional)
	to := int64(789) // int64 | Unix timestamp in milliseconds. Default value is now. (optional)
	after := "after_example" // string | Identifier used for pagination (optional)
	before := "before_example" // string | Identifier used for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsFlagEventsBetaApi.GetFlagEvents(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Query(query).ImpactSize(impactSize).HasExperiments(hasExperiments).Global(global).Expand(expand).Limit(limit).From(from).To(to).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsFlagEventsBetaApi.GetFlagEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagEvents`: FlagEventCollectionRep
	fmt.Fprintf(os.Stdout, "Response from `InsightsFlagEventsBetaApi.GetFlagEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | The environment key | 
 **applicationKey** | **string** | Comma separated list of application keys | 
 **query** | **string** | Filter events by flag key | 
 **impactSize** | **string** | Filter events by impact size. A small impact created a less than 20% change in the proportion of end users receiving one or more flag variations. A medium impact created between a 20%-80% change. A large impact created a more than 80% change. Options: &#x60;none&#x60;, &#x60;small&#x60;, &#x60;medium&#x60;, &#x60;large&#x60; | 
 **hasExperiments** | **bool** | Filter events to those associated with an experiment (&#x60;true&#x60;) or without an experiment (&#x60;false&#x60;) | 
 **global** | **string** | Filter to include or exclude global events. Default value is &#x60;include&#x60;. Options: &#x60;include&#x60;, &#x60;exclude&#x60; | 
 **expand** | **string** | Expand properties in response. Options: &#x60;experiments&#x60; | 
 **limit** | **int64** | The number of deployments to return. Default is 20. Maximum allowed is 100. | 
 **from** | **int64** | Unix timestamp in milliseconds. Default value is 7 days ago. | 
 **to** | **int64** | Unix timestamp in milliseconds. Default value is now. | 
 **after** | **string** | Identifier used for pagination | 
 **before** | **string** | Identifier used for pagination | 

### Return type

[**FlagEventCollectionRep**](FlagEventCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

