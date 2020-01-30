# EnvironmentPost

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the new environment. | [default to null]
**Key** | **string** | A project-unique key for the new environment. | [default to null]
**Color** | **string** | A color swatch (as an RGB hex value with no leading &#39;#&#39;, e.g. C8C8C8). | [default to null]
**DefaultTtl** | **float32** | The default TTL for the new environment. | [optional] [default to null]
**SecureMode** | **bool** | Determines whether the environment is in secure mode. | [optional] [default to null]
**DefaultTrackEvents** | **bool** | Set to true to send detailed event information for newly created flags. | [optional] [default to null]
**Tags** | **[]string** | An array of tags for this environment. | [optional] [default to null]
**RequireComments** | **bool** | Determines if this environment requires comments for flag and segment changes. | [optional] [default to null]
**ConfirmChanges** | **bool** | Determines if this environment requires confirmation for flag and segment changes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


