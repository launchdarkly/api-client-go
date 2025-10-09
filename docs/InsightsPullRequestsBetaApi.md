# \InsightsPullRequestsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPullRequests**](InsightsPullRequestsBetaApi.md#GetPullRequests) | **Get** /api/v2/engineering-insights/pull-requests | List pull requests



## GetPullRequests

> PullRequestCollectionRep GetPullRequests(ctx).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Status(status).Query(query).Limit(limit).Expand(expand).Sort(sort).From(from).To(to).After(after).Before(before).Execute()

List pull requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/launchdarkly/api-client-go"
)

func main() {
	projectKey := "projectKey_example" // string | The project key
	environmentKey := "environmentKey_example" // string | Required if you are using the <code>sort</code> parameter's <code>leadTime</code> option to sort pull requests. (optional)
	applicationKey := "applicationKey_example" // string | Filter the results to pull requests deployed to a comma separated list of applications (optional)
	status := "status_example" // string | Filter results to pull requests with the given status. Options: `open`, `merged`, `closed`, `deployed`. (optional)
	query := "query_example" // string | Filter list of pull requests by title or author (optional)
	limit := int64(789) // int64 | The number of pull requests to return. Default is 20. Maximum allowed is 100. (optional)
	expand := "expand_example" // string | Expand properties in response. Options: `deployments`, `flagReferences`, `leadTime`. (optional)
	sort := "sort_example" // string | Sort results. Requires the `environmentKey` to be set. Options: `leadTime` (asc) and `-leadTime` (desc). When query option is excluded, default sort is by created or merged date. (optional)
	from := time.Now() // time.Time | Unix timestamp in milliseconds. Default value is 7 days ago. (optional)
	to := time.Now() // time.Time | Unix timestamp in milliseconds. Default value is now. (optional)
	after := "after_example" // string | Identifier used for pagination (optional)
	before := "before_example" // string | Identifier used for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InsightsPullRequestsBetaApi.GetPullRequests(context.Background()).ProjectKey(projectKey).EnvironmentKey(environmentKey).ApplicationKey(applicationKey).Status(status).Query(query).Limit(limit).Expand(expand).Sort(sort).From(from).To(to).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InsightsPullRequestsBetaApi.GetPullRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequests`: PullRequestCollectionRep
	fmt.Fprintf(os.Stdout, "Response from `InsightsPullRequestsBetaApi.GetPullRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The project key | 
 **environmentKey** | **string** | Required if you are using the &lt;code&gt;sort&lt;/code&gt; parameter&#39;s &lt;code&gt;leadTime&lt;/code&gt; option to sort pull requests. | 
 **applicationKey** | **string** | Filter the results to pull requests deployed to a comma separated list of applications | 
 **status** | **string** | Filter results to pull requests with the given status. Options: &#x60;open&#x60;, &#x60;merged&#x60;, &#x60;closed&#x60;, &#x60;deployed&#x60;. | 
 **query** | **string** | Filter list of pull requests by title or author | 
 **limit** | **int64** | The number of pull requests to return. Default is 20. Maximum allowed is 100. | 
 **expand** | **string** | Expand properties in response. Options: &#x60;deployments&#x60;, &#x60;flagReferences&#x60;, &#x60;leadTime&#x60;. | 
 **sort** | **string** | Sort results. Requires the &#x60;environmentKey&#x60; to be set. Options: &#x60;leadTime&#x60; (asc) and &#x60;-leadTime&#x60; (desc). When query option is excluded, default sort is by created or merged date. | 
 **from** | **time.Time** | Unix timestamp in milliseconds. Default value is 7 days ago. | 
 **to** | **time.Time** | Unix timestamp in milliseconds. Default value is now. | 
 **after** | **string** | Identifier used for pagination | 
 **before** | **string** | Identifier used for pagination | 

### Return type

[**PullRequestCollectionRep**](PullRequestCollectionRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

