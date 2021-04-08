# ApprovalRequestConfigBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A name that describes the changes you would like to apply to a feature flag configuration | [default to null]
**Instructions** | [***SemanticPatchInstruction**](SemanticPatchInstruction.md) |  | [default to null]
**NotifyMemberIds** | **[]string** |  | [default to null]
**Comment** | **string** | comment will be included in audit log item for change. | [optional] [default to null]
**ExecutionDate** | **int64** | Timestamp for when instructions will be executed | [optional] [default to null]
**OperatingOnId** | **string** | ID of scheduled change to edit or delete | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


