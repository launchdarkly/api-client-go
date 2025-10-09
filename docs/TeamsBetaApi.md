# \TeamsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchTeams**](TeamsBetaApi.md#PatchTeams) | **Patch** /api/v2/teams | Update teams



## PatchTeams

> BulkEditTeamsRep PatchTeams(ctx).TeamsPatchInput(teamsPatchInput).Execute()

Update teams



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
	teamsPatchInput := *openapiclient.NewTeamsPatchInput([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // TeamsPatchInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsBetaApi.PatchTeams(context.Background()).TeamsPatchInput(teamsPatchInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.PatchTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTeams`: BulkEditTeamsRep
	fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.PatchTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamsPatchInput** | [**TeamsPatchInput**](TeamsPatchInput.md) |  | 

### Return type

[**BulkEditTeamsRep**](BulkEditTeamsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

