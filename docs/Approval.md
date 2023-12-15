# Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationRequest** | Pointer to [**CreationRequest**](CreationRequest.md) |  | [optional] 
**DeletionRequest** | Pointer to [**DeletionRequest**](DeletionRequest.md) |  | [optional] 
**EnvironmentFormVariables** | Pointer to [**[]FormVariable**](FormVariable.md) |  | [optional] 
**FlagFormVariables** | Pointer to [**[]FormVariable**](FormVariable.md) |  | [optional] 
**MemberListRequest** | Pointer to [**MemberListRequest**](MemberListRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PostApplyRequest** | Pointer to [**PostApplyRequest**](PostApplyRequest.md) |  | [optional] 
**StatusRequest** | Pointer to [**StatusRequest**](StatusRequest.md) |  | [optional] 

## Methods

### NewApproval

`func NewApproval() *Approval`

NewApproval instantiates a new Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWithDefaults

`func NewApprovalWithDefaults() *Approval`

NewApprovalWithDefaults instantiates a new Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationRequest

`func (o *Approval) GetCreationRequest() CreationRequest`

GetCreationRequest returns the CreationRequest field if non-nil, zero value otherwise.

### GetCreationRequestOk

`func (o *Approval) GetCreationRequestOk() (*CreationRequest, bool)`

GetCreationRequestOk returns a tuple with the CreationRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationRequest

`func (o *Approval) SetCreationRequest(v CreationRequest)`

SetCreationRequest sets CreationRequest field to given value.

### HasCreationRequest

`func (o *Approval) HasCreationRequest() bool`

HasCreationRequest returns a boolean if a field has been set.

### GetDeletionRequest

`func (o *Approval) GetDeletionRequest() DeletionRequest`

GetDeletionRequest returns the DeletionRequest field if non-nil, zero value otherwise.

### GetDeletionRequestOk

`func (o *Approval) GetDeletionRequestOk() (*DeletionRequest, bool)`

GetDeletionRequestOk returns a tuple with the DeletionRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionRequest

`func (o *Approval) SetDeletionRequest(v DeletionRequest)`

SetDeletionRequest sets DeletionRequest field to given value.

### HasDeletionRequest

`func (o *Approval) HasDeletionRequest() bool`

HasDeletionRequest returns a boolean if a field has been set.

### GetEnvironmentFormVariables

`func (o *Approval) GetEnvironmentFormVariables() []FormVariable`

GetEnvironmentFormVariables returns the EnvironmentFormVariables field if non-nil, zero value otherwise.

### GetEnvironmentFormVariablesOk

`func (o *Approval) GetEnvironmentFormVariablesOk() (*[]FormVariable, bool)`

GetEnvironmentFormVariablesOk returns a tuple with the EnvironmentFormVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentFormVariables

`func (o *Approval) SetEnvironmentFormVariables(v []FormVariable)`

SetEnvironmentFormVariables sets EnvironmentFormVariables field to given value.

### HasEnvironmentFormVariables

`func (o *Approval) HasEnvironmentFormVariables() bool`

HasEnvironmentFormVariables returns a boolean if a field has been set.

### GetFlagFormVariables

`func (o *Approval) GetFlagFormVariables() []FormVariable`

GetFlagFormVariables returns the FlagFormVariables field if non-nil, zero value otherwise.

### GetFlagFormVariablesOk

`func (o *Approval) GetFlagFormVariablesOk() (*[]FormVariable, bool)`

GetFlagFormVariablesOk returns a tuple with the FlagFormVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagFormVariables

`func (o *Approval) SetFlagFormVariables(v []FormVariable)`

SetFlagFormVariables sets FlagFormVariables field to given value.

### HasFlagFormVariables

`func (o *Approval) HasFlagFormVariables() bool`

HasFlagFormVariables returns a boolean if a field has been set.

### GetMemberListRequest

`func (o *Approval) GetMemberListRequest() MemberListRequest`

GetMemberListRequest returns the MemberListRequest field if non-nil, zero value otherwise.

### GetMemberListRequestOk

`func (o *Approval) GetMemberListRequestOk() (*MemberListRequest, bool)`

GetMemberListRequestOk returns a tuple with the MemberListRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberListRequest

`func (o *Approval) SetMemberListRequest(v MemberListRequest)`

SetMemberListRequest sets MemberListRequest field to given value.

### HasMemberListRequest

`func (o *Approval) HasMemberListRequest() bool`

HasMemberListRequest returns a boolean if a field has been set.

### GetName

`func (o *Approval) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Approval) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Approval) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Approval) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPostApplyRequest

`func (o *Approval) GetPostApplyRequest() PostApplyRequest`

GetPostApplyRequest returns the PostApplyRequest field if non-nil, zero value otherwise.

### GetPostApplyRequestOk

`func (o *Approval) GetPostApplyRequestOk() (*PostApplyRequest, bool)`

GetPostApplyRequestOk returns a tuple with the PostApplyRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostApplyRequest

`func (o *Approval) SetPostApplyRequest(v PostApplyRequest)`

SetPostApplyRequest sets PostApplyRequest field to given value.

### HasPostApplyRequest

`func (o *Approval) HasPostApplyRequest() bool`

HasPostApplyRequest returns a boolean if a field has been set.

### GetStatusRequest

`func (o *Approval) GetStatusRequest() StatusRequest`

GetStatusRequest returns the StatusRequest field if non-nil, zero value otherwise.

### GetStatusRequestOk

`func (o *Approval) GetStatusRequestOk() (*StatusRequest, bool)`

GetStatusRequestOk returns a tuple with the StatusRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusRequest

`func (o *Approval) SetStatusRequest(v StatusRequest)`

SetStatusRequest sets StatusRequest field to given value.

### HasStatusRequest

`func (o *Approval) HasStatusRequest() bool`

HasStatusRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


