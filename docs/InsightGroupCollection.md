# InsightGroupCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** | The total number of insight groups | 
**Items** | [**[]InsightGroup**](InsightGroup.md) | A list of insight groups | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 
**Metadata** | Pointer to [**InsightGroupCollectionMetadata**](InsightGroupCollectionMetadata.md) |  | [optional] 
**ScoreMetadata** | Pointer to [**InsightGroupCollectionScoreMetadata**](InsightGroupCollectionScoreMetadata.md) |  | [optional] 

## Methods

### NewInsightGroupCollection

`func NewInsightGroupCollection(totalCount int32, items []InsightGroup, ) *InsightGroupCollection`

NewInsightGroupCollection instantiates a new InsightGroupCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightGroupCollectionWithDefaults

`func NewInsightGroupCollectionWithDefaults() *InsightGroupCollection`

NewInsightGroupCollectionWithDefaults instantiates a new InsightGroupCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *InsightGroupCollection) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *InsightGroupCollection) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *InsightGroupCollection) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetItems

`func (o *InsightGroupCollection) GetItems() []InsightGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InsightGroupCollection) GetItemsOk() (*[]InsightGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InsightGroupCollection) SetItems(v []InsightGroup)`

SetItems sets Items field to given value.


### GetLinks

`func (o *InsightGroupCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InsightGroupCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InsightGroupCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InsightGroupCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *InsightGroupCollection) GetMetadata() InsightGroupCollectionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InsightGroupCollection) GetMetadataOk() (*InsightGroupCollectionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InsightGroupCollection) SetMetadata(v InsightGroupCollectionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InsightGroupCollection) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScoreMetadata

`func (o *InsightGroupCollection) GetScoreMetadata() InsightGroupCollectionScoreMetadata`

GetScoreMetadata returns the ScoreMetadata field if non-nil, zero value otherwise.

### GetScoreMetadataOk

`func (o *InsightGroupCollection) GetScoreMetadataOk() (*InsightGroupCollectionScoreMetadata, bool)`

GetScoreMetadataOk returns a tuple with the ScoreMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreMetadata

`func (o *InsightGroupCollection) SetScoreMetadata(v InsightGroupCollectionScoreMetadata)`

SetScoreMetadata sets ScoreMetadata field to given value.

### HasScoreMetadata

`func (o *InsightGroupCollection) HasScoreMetadata() bool`

HasScoreMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


