# InsightsChartSeriesDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **int64** | X-axis value | 
**Y** | **int64** | Y-axis value | 
**Values** | Pointer to **map[string]interface{}** | Additional values for the data point | [optional] 

## Methods

### NewInsightsChartSeriesDataPoint

`func NewInsightsChartSeriesDataPoint(x int64, y int64, ) *InsightsChartSeriesDataPoint`

NewInsightsChartSeriesDataPoint instantiates a new InsightsChartSeriesDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsChartSeriesDataPointWithDefaults

`func NewInsightsChartSeriesDataPointWithDefaults() *InsightsChartSeriesDataPoint`

NewInsightsChartSeriesDataPointWithDefaults instantiates a new InsightsChartSeriesDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *InsightsChartSeriesDataPoint) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *InsightsChartSeriesDataPoint) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *InsightsChartSeriesDataPoint) SetX(v int64)`

SetX sets X field to given value.


### GetY

`func (o *InsightsChartSeriesDataPoint) GetY() int64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *InsightsChartSeriesDataPoint) GetYOk() (*int64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *InsightsChartSeriesDataPoint) SetY(v int64)`

SetY sets Y field to given value.


### GetValues

`func (o *InsightsChartSeriesDataPoint) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InsightsChartSeriesDataPoint) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InsightsChartSeriesDataPoint) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *InsightsChartSeriesDataPoint) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


