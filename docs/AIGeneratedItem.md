# AIGeneratedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the item | 
**Description** | **string** | Description of the item | 
**Message** | **string** | Message from AI | 
**SuggestedItems** | Pointer to [**[]SuggestedItem**](SuggestedItem.md) | Items that are similar to the generated item | [optional] 

## Methods

### NewAIGeneratedItem

`func NewAIGeneratedItem(name string, description string, message string, ) *AIGeneratedItem`

NewAIGeneratedItem instantiates a new AIGeneratedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIGeneratedItemWithDefaults

`func NewAIGeneratedItemWithDefaults() *AIGeneratedItem`

NewAIGeneratedItemWithDefaults instantiates a new AIGeneratedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AIGeneratedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AIGeneratedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AIGeneratedItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AIGeneratedItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AIGeneratedItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AIGeneratedItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMessage

`func (o *AIGeneratedItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AIGeneratedItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AIGeneratedItem) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSuggestedItems

`func (o *AIGeneratedItem) GetSuggestedItems() []SuggestedItem`

GetSuggestedItems returns the SuggestedItems field if non-nil, zero value otherwise.

### GetSuggestedItemsOk

`func (o *AIGeneratedItem) GetSuggestedItemsOk() (*[]SuggestedItem, bool)`

GetSuggestedItemsOk returns a tuple with the SuggestedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedItems

`func (o *AIGeneratedItem) SetSuggestedItems(v []SuggestedItem)`

SetSuggestedItems sets SuggestedItems field to given value.

### HasSuggestedItems

`func (o *AIGeneratedItem) HasSuggestedItems() bool`

HasSuggestedItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


