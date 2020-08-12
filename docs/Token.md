# Token

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***Links**](Links.md) |  | [optional] [default to null]
**Id** | **string** | The unique resource id. | [optional] [default to null]
**OwnerId** | **string** | The unique resource id. | [optional] [default to null]
**MemberId** | **string** | The unique resource id. | [optional] [default to null]
**Member** | [***Member**](Member.md) |  | [optional] [default to null]
**CreationDate** | **int64** | A unix epoch time in milliseconds specifying the creation time of this access token. | [optional] [default to null]
**LastModified** | **int64** | A unix epoch time in milliseconds specifying the last time this access token was modified. | [optional] [default to null]
**LastUsed** | **int64** | A unix epoch time in milliseconds specifying the last time this access token was used to authorize access to the LaunchDarkly REST API. | [optional] [default to null]
**Token** | **string** | The last 4 digits of the unique secret key for this access token. If creating or resetting the token, this will be the full token secret. | [optional] [default to null]
**Name** | **string** | A human-friendly name for the access token | [optional] [default to null]
**Role** | **string** | The name of a built-in role for the token | [optional] [default to null]
**CustomRoleIds** | **[]string** | A list of custom role IDs to use as access limits for the access token | [optional] [default to null]
**InlineRole** | [**[]Statement**](Statement.md) |  | [optional] [default to null]
**ServiceToken** | **bool** | Whether the token will be a service token https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens | [optional] [default to null]
**DefaultApiVersion** | **int32** | The default API version for this token | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


