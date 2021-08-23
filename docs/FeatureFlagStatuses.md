# FeatureFlagStatuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Items** | Pointer to [**[]FlagStatusRep**](FlagStatusRep.md) |  | [optional] 

## Methods

### NewFeatureFlagStatuses

`func NewFeatureFlagStatuses(links map[string]Link, ) *FeatureFlagStatuses`

NewFeatureFlagStatuses instantiates a new FeatureFlagStatuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagStatusesWithDefaults

`func NewFeatureFlagStatusesWithDefaults() *FeatureFlagStatuses`

NewFeatureFlagStatusesWithDefaults instantiates a new FeatureFlagStatuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FeatureFlagStatuses) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlagStatuses) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlagStatuses) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetItems

`func (o *FeatureFlagStatuses) GetItems() []FlagStatusRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FeatureFlagStatuses) GetItemsOk() (*[]FlagStatusRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FeatureFlagStatuses) SetItems(v []FlagStatusRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *FeatureFlagStatuses) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


