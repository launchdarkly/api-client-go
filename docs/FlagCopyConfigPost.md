# FlagCopyConfigPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**FlagCopyConfigEnvironment**](FlagCopyConfigEnvironment.md) |  | 
**Target** | [**FlagCopyConfigEnvironment**](FlagCopyConfigEnvironment.md) |  | 
**Comment** | Pointer to **string** |  | [optional] 
**IncludedActions** | Pointer to **[]string** |  | [optional] 
**ExcludedActions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFlagCopyConfigPost

`func NewFlagCopyConfigPost(source FlagCopyConfigEnvironment, target FlagCopyConfigEnvironment, ) *FlagCopyConfigPost`

NewFlagCopyConfigPost instantiates a new FlagCopyConfigPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagCopyConfigPostWithDefaults

`func NewFlagCopyConfigPostWithDefaults() *FlagCopyConfigPost`

NewFlagCopyConfigPostWithDefaults instantiates a new FlagCopyConfigPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *FlagCopyConfigPost) GetSource() FlagCopyConfigEnvironment`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FlagCopyConfigPost) GetSourceOk() (*FlagCopyConfigEnvironment, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FlagCopyConfigPost) SetSource(v FlagCopyConfigEnvironment)`

SetSource sets Source field to given value.


### GetTarget

`func (o *FlagCopyConfigPost) GetTarget() FlagCopyConfigEnvironment`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FlagCopyConfigPost) GetTargetOk() (*FlagCopyConfigEnvironment, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FlagCopyConfigPost) SetTarget(v FlagCopyConfigEnvironment)`

SetTarget sets Target field to given value.


### GetComment

`func (o *FlagCopyConfigPost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FlagCopyConfigPost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FlagCopyConfigPost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FlagCopyConfigPost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIncludedActions

`func (o *FlagCopyConfigPost) GetIncludedActions() []string`

GetIncludedActions returns the IncludedActions field if non-nil, zero value otherwise.

### GetIncludedActionsOk

`func (o *FlagCopyConfigPost) GetIncludedActionsOk() (*[]string, bool)`

GetIncludedActionsOk returns a tuple with the IncludedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedActions

`func (o *FlagCopyConfigPost) SetIncludedActions(v []string)`

SetIncludedActions sets IncludedActions field to given value.

### HasIncludedActions

`func (o *FlagCopyConfigPost) HasIncludedActions() bool`

HasIncludedActions returns a boolean if a field has been set.

### GetExcludedActions

`func (o *FlagCopyConfigPost) GetExcludedActions() []string`

GetExcludedActions returns the ExcludedActions field if non-nil, zero value otherwise.

### GetExcludedActionsOk

`func (o *FlagCopyConfigPost) GetExcludedActionsOk() (*[]string, bool)`

GetExcludedActionsOk returns a tuple with the ExcludedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedActions

`func (o *FlagCopyConfigPost) SetExcludedActions(v []string)`

SetExcludedActions sets ExcludedActions field to given value.

### HasExcludedActions

`func (o *FlagCopyConfigPost) HasExcludedActions() bool`

HasExcludedActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


