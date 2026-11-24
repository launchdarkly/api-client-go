# SegmentScheduledChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]SegmentScheduledChangesLink**](SegmentScheduledChangesLink.md) | The location and content type of related resources | 
**Status** | [**SegmentScheduledChangeStatus**](SegmentScheduledChangeStatus.md) |  | 
**Id** | **string** | The ID of this scheduled change | 
**CreationDate** | **int64** |  | 
**MaintainerId** | Pointer to **string** | The ID of the member who scheduled the change | [optional] 
**MaintainerServiceTokenId** | Pointer to **string** | The ID of the service token that scheduled the change | [optional] 
**ExecutionDate** | **int64** |  | 
**Comment** | Pointer to **string** | The comment describing the scheduled change | [optional] 
**Instructions** | **[]map[string]interface{}** | The segment semantic patch instructions to apply on the execution date | 

## Methods

### NewSegmentScheduledChange

`func NewSegmentScheduledChange(links map[string]SegmentScheduledChangesLink, status SegmentScheduledChangeStatus, id string, creationDate int64, executionDate int64, instructions []map[string]interface{}, ) *SegmentScheduledChange`

NewSegmentScheduledChange instantiates a new SegmentScheduledChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentScheduledChangeWithDefaults

`func NewSegmentScheduledChangeWithDefaults() *SegmentScheduledChange`

NewSegmentScheduledChangeWithDefaults instantiates a new SegmentScheduledChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SegmentScheduledChange) GetLinks() map[string]SegmentScheduledChangesLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SegmentScheduledChange) GetLinksOk() (*map[string]SegmentScheduledChangesLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SegmentScheduledChange) SetLinks(v map[string]SegmentScheduledChangesLink)`

SetLinks sets Links field to given value.


### GetStatus

`func (o *SegmentScheduledChange) GetStatus() SegmentScheduledChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SegmentScheduledChange) GetStatusOk() (*SegmentScheduledChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SegmentScheduledChange) SetStatus(v SegmentScheduledChangeStatus)`

SetStatus sets Status field to given value.


### GetId

`func (o *SegmentScheduledChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SegmentScheduledChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SegmentScheduledChange) SetId(v string)`

SetId sets Id field to given value.


### GetCreationDate

`func (o *SegmentScheduledChange) GetCreationDate() int64`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SegmentScheduledChange) GetCreationDateOk() (*int64, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SegmentScheduledChange) SetCreationDate(v int64)`

SetCreationDate sets CreationDate field to given value.


### GetMaintainerId

`func (o *SegmentScheduledChange) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *SegmentScheduledChange) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *SegmentScheduledChange) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *SegmentScheduledChange) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### GetMaintainerServiceTokenId

`func (o *SegmentScheduledChange) GetMaintainerServiceTokenId() string`

GetMaintainerServiceTokenId returns the MaintainerServiceTokenId field if non-nil, zero value otherwise.

### GetMaintainerServiceTokenIdOk

`func (o *SegmentScheduledChange) GetMaintainerServiceTokenIdOk() (*string, bool)`

GetMaintainerServiceTokenIdOk returns a tuple with the MaintainerServiceTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerServiceTokenId

`func (o *SegmentScheduledChange) SetMaintainerServiceTokenId(v string)`

SetMaintainerServiceTokenId sets MaintainerServiceTokenId field to given value.

### HasMaintainerServiceTokenId

`func (o *SegmentScheduledChange) HasMaintainerServiceTokenId() bool`

HasMaintainerServiceTokenId returns a boolean if a field has been set.

### GetExecutionDate

`func (o *SegmentScheduledChange) GetExecutionDate() int64`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *SegmentScheduledChange) GetExecutionDateOk() (*int64, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *SegmentScheduledChange) SetExecutionDate(v int64)`

SetExecutionDate sets ExecutionDate field to given value.


### GetComment

`func (o *SegmentScheduledChange) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SegmentScheduledChange) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SegmentScheduledChange) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SegmentScheduledChange) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *SegmentScheduledChange) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *SegmentScheduledChange) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *SegmentScheduledChange) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


