# FlagLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmptyState** | Pointer to [**EmptyState**](EmptyState.md) |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**map[string]MetadataItem**](MetadataItem.md) |  | [optional] 
**UiBlocks** | Pointer to [**UIBlocks**](UIBlocks.md) |  | [optional] 

## Methods

### NewFlagLink

`func NewFlagLink() *FlagLink`

NewFlagLink instantiates a new FlagLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagLinkWithDefaults

`func NewFlagLinkWithDefaults() *FlagLink`

NewFlagLinkWithDefaults instantiates a new FlagLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmptyState

`func (o *FlagLink) GetEmptyState() EmptyState`

GetEmptyState returns the EmptyState field if non-nil, zero value otherwise.

### GetEmptyStateOk

`func (o *FlagLink) GetEmptyStateOk() (*EmptyState, bool)`

GetEmptyStateOk returns a tuple with the EmptyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyState

`func (o *FlagLink) SetEmptyState(v EmptyState)`

SetEmptyState sets EmptyState field to given value.

### HasEmptyState

`func (o *FlagLink) HasEmptyState() bool`

HasEmptyState returns a boolean if a field has been set.

### GetHeader

`func (o *FlagLink) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FlagLink) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FlagLink) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *FlagLink) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetMetadata

`func (o *FlagLink) GetMetadata() map[string]MetadataItem`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FlagLink) GetMetadataOk() (*map[string]MetadataItem, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FlagLink) SetMetadata(v map[string]MetadataItem)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FlagLink) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUiBlocks

`func (o *FlagLink) GetUiBlocks() UIBlocks`

GetUiBlocks returns the UiBlocks field if non-nil, zero value otherwise.

### GetUiBlocksOk

`func (o *FlagLink) GetUiBlocksOk() (*UIBlocks, bool)`

GetUiBlocksOk returns a tuple with the UiBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiBlocks

`func (o *FlagLink) SetUiBlocks(v UIBlocks)`

SetUiBlocks sets UiBlocks field to given value.

### HasUiBlocks

`func (o *FlagLink) HasUiBlocks() bool`

HasUiBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


