# FeatureFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]FeatureFlag**](FeatureFlag.md) | An array of feature flags | 
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**TotalCount** | Pointer to **int32** | The total number of flags | [optional] 
**TotalCountWithDifferences** | Pointer to **int32** | The number of flags that have differences between environments. Only shown when query parameter &lt;code&gt;compare&lt;/code&gt; is &lt;code&gt;true&lt;/code&gt;. | [optional] 

## Methods

### NewFeatureFlags

`func NewFeatureFlags(items []FeatureFlag, links map[string]Link, ) *FeatureFlags`

NewFeatureFlags instantiates a new FeatureFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagsWithDefaults

`func NewFeatureFlagsWithDefaults() *FeatureFlags`

NewFeatureFlagsWithDefaults instantiates a new FeatureFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *FeatureFlags) GetItems() []FeatureFlag`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FeatureFlags) GetItemsOk() (*[]FeatureFlag, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FeatureFlags) SetItems(v []FeatureFlag)`

SetItems sets Items field to given value.


### GetLinks

`func (o *FeatureFlags) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FeatureFlags) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FeatureFlags) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetTotalCount

`func (o *FeatureFlags) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FeatureFlags) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FeatureFlags) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *FeatureFlags) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalCountWithDifferences

`func (o *FeatureFlags) GetTotalCountWithDifferences() int32`

GetTotalCountWithDifferences returns the TotalCountWithDifferences field if non-nil, zero value otherwise.

### GetTotalCountWithDifferencesOk

`func (o *FeatureFlags) GetTotalCountWithDifferencesOk() (*int32, bool)`

GetTotalCountWithDifferencesOk returns a tuple with the TotalCountWithDifferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCountWithDifferences

`func (o *FeatureFlags) SetTotalCountWithDifferences(v int32)`

SetTotalCountWithDifferences sets TotalCountWithDifferences field to given value.

### HasTotalCountWithDifferences

`func (o *FeatureFlags) HasTotalCountWithDifferences() bool`

HasTotalCountWithDifferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


