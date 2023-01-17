# StatisticsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]StatisticRep**](StatisticRep.md) | A list of code reference statistics for each code repository | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewStatisticsRep

`func NewStatisticsRep() *StatisticsRep`

NewStatisticsRep instantiates a new StatisticsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsRepWithDefaults

`func NewStatisticsRepWithDefaults() *StatisticsRep`

NewStatisticsRepWithDefaults instantiates a new StatisticsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *StatisticsRep) GetItems() []StatisticRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *StatisticsRep) GetItemsOk() (*[]StatisticRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *StatisticsRep) SetItems(v []StatisticRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *StatisticsRep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *StatisticsRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StatisticsRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StatisticsRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StatisticsRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


