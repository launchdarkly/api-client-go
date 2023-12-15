# Title

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]UIBlockElement**](UIBlockElement.md) |  | [optional] 
**LinkToReference** | Pointer to **bool** |  | [optional] 

## Methods

### NewTitle

`func NewTitle() *Title`

NewTitle instantiates a new Title object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTitleWithDefaults

`func NewTitleWithDefaults() *Title`

NewTitleWithDefaults instantiates a new Title object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *Title) GetElements() []UIBlockElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *Title) GetElementsOk() (*[]UIBlockElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *Title) SetElements(v []UIBlockElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *Title) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetLinkToReference

`func (o *Title) GetLinkToReference() bool`

GetLinkToReference returns the LinkToReference field if non-nil, zero value otherwise.

### GetLinkToReferenceOk

`func (o *Title) GetLinkToReferenceOk() (*bool, bool)`

GetLinkToReferenceOk returns a tuple with the LinkToReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToReference

`func (o *Title) SetLinkToReference(v bool)`

SetLinkToReference sets LinkToReference field to given value.

### HasLinkToReference

`func (o *Title) HasLinkToReference() bool`

HasLinkToReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


