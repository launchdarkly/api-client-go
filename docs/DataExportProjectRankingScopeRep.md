# DataExportProjectRankingScopeRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The included usage source. | 
**Dimension** | **string** | The ranked dimension. | 
**From** | **int64** | Inclusive UTC range start in Unix milliseconds. | 
**To** | **int64** | Exclusive UTC range end in Unix milliseconds. | 

## Methods

### NewDataExportProjectRankingScopeRep

`func NewDataExportProjectRankingScopeRep(source string, dimension string, from int64, to int64, ) *DataExportProjectRankingScopeRep`

NewDataExportProjectRankingScopeRep instantiates a new DataExportProjectRankingScopeRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExportProjectRankingScopeRepWithDefaults

`func NewDataExportProjectRankingScopeRepWithDefaults() *DataExportProjectRankingScopeRep`

NewDataExportProjectRankingScopeRepWithDefaults instantiates a new DataExportProjectRankingScopeRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *DataExportProjectRankingScopeRep) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataExportProjectRankingScopeRep) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataExportProjectRankingScopeRep) SetSource(v string)`

SetSource sets Source field to given value.


### GetDimension

`func (o *DataExportProjectRankingScopeRep) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *DataExportProjectRankingScopeRep) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *DataExportProjectRankingScopeRep) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetFrom

`func (o *DataExportProjectRankingScopeRep) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DataExportProjectRankingScopeRep) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DataExportProjectRankingScopeRep) SetFrom(v int64)`

SetFrom sets From field to given value.


### GetTo

`func (o *DataExportProjectRankingScopeRep) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DataExportProjectRankingScopeRep) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DataExportProjectRankingScopeRep) SetTo(v int64)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


