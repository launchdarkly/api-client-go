# EnvironmentApprovalSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceKind** | **string** | The approvals system used. | [optional] [default to null]
**Required** | **bool** | Whether any changes to flags in this environment will require approval. | [optional] [default to null]
**CanReviewOwnRequest** | **bool** | Whether requesters can approve or decline their own request. They may always comment. | [optional] [default to null]
**MinNumApprovals** | **int64** | The number of approvals required before an approval request can be applied. | [optional] [default to null]
**CanApplyDeclinedChanges** | **bool** | Whether changes can be applied as long as minNumApprovals is met, regardless of if any reviewers have declined a request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


