# StatisticCollectionRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | [**map[string][]StatisticRep**](array.md) | A map of flag keys to a list of code reference statistics for each code repository in which the flag key appears | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 

## Methods

### NewStatisticCollectionRep

`func NewStatisticCollectionRep(flags map[string][]StatisticRep, links map[string]Link, ) *StatisticCollectionRep`

NewStatisticCollectionRep instantiates a new StatisticCollectionRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticCollectionRepWithDefaults

`func NewStatisticCollectionRepWithDefaults() *StatisticCollectionRep`

NewStatisticCollectionRepWithDefaults instantiates a new StatisticCollectionRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *StatisticCollectionRep) GetFlags() map[string][]StatisticRep`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *StatisticCollectionRep) GetFlagsOk() (*map[string][]StatisticRep, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *StatisticCollectionRep) SetFlags(v map[string][]StatisticRep)`

SetFlags sets Flags field to given value.


### GetLinks

`func (o *StatisticCollectionRep) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StatisticCollectionRep) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StatisticCollectionRep) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


