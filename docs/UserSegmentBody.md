# UserSegmentBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A human-friendly name for the user segment. | [default to null]
**Key** | **string** | A unique key that will be used to reference the user segment in feature flags. | [default to null]
**Description** | **string** | A description for the user segment. | [optional] [default to null]
**Unbounded** | **bool** | Controls whether this is considered a \&quot;big segment\&quot; which can support an unlimited numbers of users. Include/exclude lists sent with this payload are not used in big segments. Contact your account manager for early access to this feature. | [optional] [default to null]
**Tags** | **[]string** | Tags for the user segment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


