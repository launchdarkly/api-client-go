# ApprovalRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique resource id. | [optional] [default to null]
**Version** | **int32** |  | [optional] [default to null]
**CreationDate** | **int32** | A unix epoch time in milliseconds specifying the date the approval request was requested | [optional] [default to null]
**RequestorId** | **string** | The id of the member that requested the change | [optional] [default to null]
**ReviewStatus** | [***ApprovalRequestReviewStatus**](ApprovalRequestReviewStatus.md) |  | [optional] [default to null]
**Status** | **string** | | Name      | Description | | ---------:| ----------- | | pending   | the approval request has not been applied yet | | completed | the approval request has been applied successfully | | scheduled | the approval request for a scheduled change has been applied successfully | | failed    | the approval request has been applied but the changes were not applied successfully |  | [optional] [default to null]
**AppliedByMemberID** | **string** | The id of the member that applied the approval request | [optional] [default to null]
**AppliedDate** | **int32** | A unix epoch time in milliseconds specifying the date the approval request was applied | [optional] [default to null]
**AllReviews** | [**[]ApprovalRequestReview**](ApprovalRequestReview.md) |  | [optional] [default to null]
**NotifyMemberIds** | **[]string** |  | [optional] [default to null]
**Instructions** | [***SemanticPatchInstruction**](SemanticPatchInstruction.md) |  | [optional] [default to null]
**ExecutionDate** | **int64** | Timestamp for when instructions will be executed | [optional] [default to null]
**OperatingOnId** | **string** | ID of scheduled change to edit or delete | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


