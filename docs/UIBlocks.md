# UIBlocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**FlagLinkContext**](FlagLinkContext.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 

## Methods

### NewUIBlocks

`func NewUIBlocks() *UIBlocks`

NewUIBlocks instantiates a new UIBlocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIBlocksWithDefaults

`func NewUIBlocksWithDefaults() *UIBlocks`

NewUIBlocksWithDefaults instantiates a new UIBlocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *UIBlocks) GetContext() FlagLinkContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UIBlocks) GetContextOk() (*FlagLinkContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UIBlocks) SetContext(v FlagLinkContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *UIBlocks) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDescription

`func (o *UIBlocks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UIBlocks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UIBlocks) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UIBlocks) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImage

`func (o *UIBlocks) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UIBlocks) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UIBlocks) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *UIBlocks) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *UIBlocks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UIBlocks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UIBlocks) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UIBlocks) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *UIBlocks) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UIBlocks) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UIBlocks) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UIBlocks) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


