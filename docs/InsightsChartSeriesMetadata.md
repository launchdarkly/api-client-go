# InsightsChartSeriesMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the series | 
**Count** | Pointer to **int64** | Aggregate count of the series values | [optional] 
**Bounds** | Pointer to [**[]InsightsChartBounds**](InsightsChartBounds.md) | Bounds for the series data | [optional] 

## Methods

### NewInsightsChartSeriesMetadata

`func NewInsightsChartSeriesMetadata(name string, ) *InsightsChartSeriesMetadata`

NewInsightsChartSeriesMetadata instantiates a new InsightsChartSeriesMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsChartSeriesMetadataWithDefaults

`func NewInsightsChartSeriesMetadataWithDefaults() *InsightsChartSeriesMetadata`

NewInsightsChartSeriesMetadataWithDefaults instantiates a new InsightsChartSeriesMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InsightsChartSeriesMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightsChartSeriesMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightsChartSeriesMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCount

`func (o *InsightsChartSeriesMetadata) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InsightsChartSeriesMetadata) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InsightsChartSeriesMetadata) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *InsightsChartSeriesMetadata) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetBounds

`func (o *InsightsChartSeriesMetadata) GetBounds() []InsightsChartBounds`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *InsightsChartSeriesMetadata) GetBoundsOk() (*[]InsightsChartBounds, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *InsightsChartSeriesMetadata) SetBounds(v []InsightsChartBounds)`

SetBounds sets Bounds field to given value.

### HasBounds

`func (o *InsightsChartSeriesMetadata) HasBounds() bool`

HasBounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


