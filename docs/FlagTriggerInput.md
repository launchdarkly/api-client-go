# FlagTriggerInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **[]map[string]interface{}** | The action to perform when triggering. It should pass an array with a single {\&quot;kind\&quot;: &lt;flag_action&gt;} object. Currently supported flag actions are \&quot;turnFlagOn\&quot; and \&quot;turnFlagOff\&quot;. | [optional] 

## Methods

### NewFlagTriggerInput

`func NewFlagTriggerInput() *FlagTriggerInput`

NewFlagTriggerInput instantiates a new FlagTriggerInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagTriggerInputWithDefaults

`func NewFlagTriggerInputWithDefaults() *FlagTriggerInput`

NewFlagTriggerInputWithDefaults instantiates a new FlagTriggerInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *FlagTriggerInput) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FlagTriggerInput) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FlagTriggerInput) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FlagTriggerInput) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInstructions

`func (o *FlagTriggerInput) GetInstructions() []map[string]interface{}`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *FlagTriggerInput) GetInstructionsOk() (*[]map[string]interface{}, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *FlagTriggerInput) SetInstructions(v []map[string]interface{})`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *FlagTriggerInput) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


