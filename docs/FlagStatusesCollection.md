# FlagStatusesCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Items** | [**[]FeatureFlagStatusAcrossEnvironments**](FeatureFlagStatusAcrossEnvironments.md) |  | 

## Methods

### NewFlagStatusesCollection

`func NewFlagStatusesCollection(links map[string]Link, items []FeatureFlagStatusAcrossEnvironments, ) *FlagStatusesCollection`

NewFlagStatusesCollection instantiates a new FlagStatusesCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagStatusesCollectionWithDefaults

`func NewFlagStatusesCollectionWithDefaults() *FlagStatusesCollection`

NewFlagStatusesCollectionWithDefaults instantiates a new FlagStatusesCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagStatusesCollection) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagStatusesCollection) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagStatusesCollection) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetItems

`func (o *FlagStatusesCollection) GetItems() []FeatureFlagStatusAcrossEnvironments`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FlagStatusesCollection) GetItemsOk() (*[]FeatureFlagStatusAcrossEnvironments, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FlagStatusesCollection) SetItems(v []FeatureFlagStatusAcrossEnvironments)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


