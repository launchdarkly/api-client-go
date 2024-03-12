# InsightScores

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**InsightPeriod**](InsightPeriod.md) |  | 
**LastPeriod** | [**InsightPeriod**](InsightPeriod.md) |  | 
**Scores** | [**InsightGroupScores**](InsightGroupScores.md) |  | 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewInsightScores

`func NewInsightScores(period InsightPeriod, lastPeriod InsightPeriod, scores InsightGroupScores, ) *InsightScores`

NewInsightScores instantiates a new InsightScores object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightScoresWithDefaults

`func NewInsightScoresWithDefaults() *InsightScores`

NewInsightScoresWithDefaults instantiates a new InsightScores object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *InsightScores) GetPeriod() InsightPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *InsightScores) GetPeriodOk() (*InsightPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *InsightScores) SetPeriod(v InsightPeriod)`

SetPeriod sets Period field to given value.


### GetLastPeriod

`func (o *InsightScores) GetLastPeriod() InsightPeriod`

GetLastPeriod returns the LastPeriod field if non-nil, zero value otherwise.

### GetLastPeriodOk

`func (o *InsightScores) GetLastPeriodOk() (*InsightPeriod, bool)`

GetLastPeriodOk returns a tuple with the LastPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPeriod

`func (o *InsightScores) SetLastPeriod(v InsightPeriod)`

SetLastPeriod sets LastPeriod field to given value.


### GetScores

`func (o *InsightScores) GetScores() InsightGroupScores`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *InsightScores) GetScoresOk() (*InsightGroupScores, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *InsightScores) SetScores(v InsightGroupScores)`

SetScores sets Scores field to given value.


### GetLinks

`func (o *InsightScores) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InsightScores) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InsightScores) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InsightScores) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


