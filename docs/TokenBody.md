# TokenBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the access token | [optional] [default to null]
**Role** | **string** | The name of a built-in role for the token | [optional] [default to null]
**CustomRoleIds** | **[]string** | A list of custom role IDs to use as access limits for the access token | [optional] [default to null]
**InlineRole** | [**[]Statement**](Statement.md) |  | [optional] [default to null]
**ServiceToken** | **bool** | Whether the token will be a service token https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens | [optional] [default to null]
**DefaultApiVersion** | **int32** | The default API version for this token | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


