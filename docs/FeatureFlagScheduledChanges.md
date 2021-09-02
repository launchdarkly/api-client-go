# FeatureFlagScheduledChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]FeatureFlagScheduledChange**](FeatureFlagScheduledChange.md) |  | 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewFeatureFlagScheduledChanges

`func NewFeatureFlagScheduledChanges(items []FeatureFlagScheduledChange, ) *FeatureFlagScheduledChanges`

NewFeatureFlagScheduledChanges instantiates a new FeatureFlagScheduledChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagScheduledChangesWithDefaults

`func NewFeatureFlagScheduledChangesWithDefaults() *FeatureFlagScheduledChanges`

NewFeatureFlagScheduledChangesWithDefaults instantiates a new FeatureFlagScheduledChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *FeatureFlagScheduledChanges) GetItems() []FeatureFlagScheduledChange`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FeatureFlagScheduledChanges) GetItemsOk() (*[]FeatureFlagScheduledChange, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FeatureFlagScheduledChanges) SetItems(v []FeatureFlagScheduledChange)`

SetItems sets Items field to given value.


### GetLinks

`func (o *FeatureFlagScheduledChanges) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlagScheduledChanges) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlagScheduledChanges) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FeatureFlagScheduledChanges) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


