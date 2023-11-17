# FlagSempatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | **[]map[string]interface{}** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewFlagSempatch

`func NewFlagSempatch(instructions []map[string]interface{}, ) *FlagSempatch`

NewFlagSempatch instantiates a new FlagSempatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagSempatchWithDefaults

`func NewFlagSempatchWithDefaults() *FlagSempatch`

NewFlagSempatchWithDefaults instantiates a new FlagSempatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *FlagSempatch) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *FlagSempatch) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *FlagSempatch) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.


### GetComment

`func (o *FlagSempatch) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FlagSempatch) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FlagSempatch) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FlagSempatch) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


