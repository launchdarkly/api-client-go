# DataExportProjectRankingPageRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The bounded page size requested. | 
**Count** | **int32** | The number of items on this page. | 
**TotalCount** | **int32** | The number of matching projects with usage. | 
**NextCursor** | Pointer to **string** | An opaque cursor for the next page. | [optional] 

## Methods

### NewDataExportProjectRankingPageRep

`func NewDataExportProjectRankingPageRep(limit int32, count int32, totalCount int32, ) *DataExportProjectRankingPageRep`

NewDataExportProjectRankingPageRep instantiates a new DataExportProjectRankingPageRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExportProjectRankingPageRepWithDefaults

`func NewDataExportProjectRankingPageRepWithDefaults() *DataExportProjectRankingPageRep`

NewDataExportProjectRankingPageRepWithDefaults instantiates a new DataExportProjectRankingPageRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *DataExportProjectRankingPageRep) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DataExportProjectRankingPageRep) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DataExportProjectRankingPageRep) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetCount

`func (o *DataExportProjectRankingPageRep) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DataExportProjectRankingPageRep) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DataExportProjectRankingPageRep) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotalCount

`func (o *DataExportProjectRankingPageRep) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DataExportProjectRankingPageRep) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DataExportProjectRankingPageRep) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetNextCursor

`func (o *DataExportProjectRankingPageRep) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *DataExportProjectRankingPageRep) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *DataExportProjectRankingPageRep) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *DataExportProjectRankingPageRep) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


