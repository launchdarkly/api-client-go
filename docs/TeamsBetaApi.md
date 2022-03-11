# \TeamsBetaApi

All URIs are relative to *https://app.launchdarkly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTeam**](TeamsBetaApi.md#DeleteTeam) | **Delete** /api/v2/teams/{teamKey} | Delete team
[**GetTeam**](TeamsBetaApi.md#GetTeam) | **Get** /api/v2/teams/{teamKey} | Get team
[**GetTeamMaintainers**](TeamsBetaApi.md#GetTeamMaintainers) | **Get** /api/v2/teams/{teamKey}/maintainers | Get team maintainers
[**GetTeamRoles**](TeamsBetaApi.md#GetTeamRoles) | **Get** /api/v2/teams/{teamKey}/roles | Get team custom roles
[**GetTeams**](TeamsBetaApi.md#GetTeams) | **Get** /api/v2/teams | List teams
[**PatchTeam**](TeamsBetaApi.md#PatchTeam) | **Patch** /api/v2/teams/{teamKey} | Update team
[**PostTeam**](TeamsBetaApi.md#PostTeam) | **Post** /api/v2/teams | Create team
[**PostTeamMembers**](TeamsBetaApi.md#PostTeamMembers) | **Post** /api/v2/teams/{teamKey}/members | Add members to team



## DeleteTeam

> DeleteTeam(ctx, teamKey).Execute()

Delete team



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
    teamKey := "teamKey_example" // string | The team key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.DeleteTeam(context.Background(), teamKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.DeleteTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string** | The team key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


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


## GetTeam

> Team GetTeam(ctx, teamKey).Execute()

Get team



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
    teamKey := "teamKey_example" // string | The team key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.GetTeam(context.Background(), teamKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.GetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeam`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string** | The team key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Team**](Team.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMaintainers

> TeamMaintainers GetTeamMaintainers(ctx, teamKey).Limit(limit).Offset(offset).Execute()

Get team maintainers



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
    teamKey := "teamKey_example" // string | The team key
    limit := int64(789) // int64 | The number of maintainers to return in the response. Defaults to 20. (optional)
    offset := int64(789) // int64 | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next `limit` items. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.GetTeamMaintainers(context.Background(), teamKey).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.GetTeamMaintainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamMaintainers`: TeamMaintainers
    fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.GetTeamMaintainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string** | The team key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMaintainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The number of maintainers to return in the response. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next &#x60;limit&#x60; items. | 

### Return type

[**TeamMaintainers**](TeamMaintainers.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamRoles

> TeamCustomRoles GetTeamRoles(ctx, teamKey).Limit(limit).Offset(offset).Execute()

Get team custom roles



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
    teamKey := "teamKey_example" // string | The team key
    limit := int64(789) // int64 | The number of roles to return in the response. Defaults to 20. (optional)
    offset := int64(789) // int64 | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next `limit` items. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.GetTeamRoles(context.Background(), teamKey).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.GetTeamRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamRoles`: TeamCustomRoles
    fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.GetTeamRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string** | The team key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | The number of roles to return in the response. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next &#x60;limit&#x60; items. | 

### Return type

[**TeamCustomRoles**](TeamCustomRoles.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeams

> Teams GetTeams(ctx).Limit(limit).Offset(offset).Filter(filter).Execute()

List teams



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
    limit := int64(789) // int64 | The number of teams to return in the response. Defaults to 20. (optional)
    offset := int64(789) // int64 | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next `limit` items. (optional)
    filter := "filter_example" // string | A comma-separated list of filters. Each filter is of the form `field:value`. Supported fields are explained above. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.GetTeams(context.Background()).Limit(limit).Offset(offset).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.GetTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeams`: Teams
    fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.GetTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | The number of teams to return in the response. Defaults to 20. | 
 **offset** | **int64** | Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first ten items and then return the next &#x60;limit&#x60; items. | 
 **filter** | **string** | A comma-separated list of filters. Each filter is of the form &#x60;field:value&#x60;. Supported fields are explained above. | 

### Return type

[**Teams**](Teams.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTeam

> Team PatchTeam(ctx, teamKey).TeamPatchInput(teamPatchInput).Execute()

Update team



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
    teamKey := "teamKey_example" // string | The team key
    teamPatchInput := *openapiclient.NewTeamPatchInput() // TeamPatchInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.PatchTeam(context.Background(), teamKey).TeamPatchInput(teamPatchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.PatchTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchTeam`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.PatchTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string** | The team key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamPatchInput** | [**TeamPatchInput**](TeamPatchInput.md) |  | 

### Return type

[**Team**](Team.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeam

> Team PostTeam(ctx).TeamPostInput(teamPostInput).Execute()

Create team



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
    teamPostInput := *openapiclient.NewTeamPostInput("Key_example", "Name_example") // TeamPostInput | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.PostTeam(context.Background()).TeamPostInput(teamPostInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.PostTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTeam`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.PostTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamPostInput** | [**TeamPostInput**](TeamPostInput.md) |  | 

### Return type

[**Team**](Team.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTeamMembers

> TeamImportsRep PostTeamMembers(ctx, teamKey).File(file).Execute()

Add members to team



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
    teamKey := "teamKey_example" // string | The team key
    file := os.NewFile(1234, "some_file") // *os.File | CSV file containing email addresses (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsBetaApi.PostTeamMembers(context.Background(), teamKey).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsBetaApi.PostTeamMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTeamMembers`: TeamImportsRep
    fmt.Fprintf(os.Stdout, "Response from `TeamsBetaApi.PostTeamMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamKey** | **string** | The team key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | CSV file containing email addresses | 

### Return type

[**TeamImportsRep**](TeamImportsRep.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

