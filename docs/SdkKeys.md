# SdkKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**SdkKeysSelfLink**](SdkKeysSelfLink.md) |  | [optional] 
**Items** | [**[]SdkKeyListItem**](SdkKeyListItem.md) |  | 
**TotalCount** | **int32** | The total number of SDK keys matching the query, before pagination. | 

## Methods

### NewSdkKeys

`func NewSdkKeys(items []SdkKeyListItem, totalCount int32, ) *SdkKeys`

NewSdkKeys instantiates a new SdkKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkKeysWithDefaults

`func NewSdkKeysWithDefaults() *SdkKeys`

NewSdkKeysWithDefaults instantiates a new SdkKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SdkKeys) GetLinks() SdkKeysSelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SdkKeys) GetLinksOk() (*SdkKeysSelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SdkKeys) SetLinks(v SdkKeysSelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SdkKeys) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *SdkKeys) GetItems() []SdkKeyListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SdkKeys) GetItemsOk() (*[]SdkKeyListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SdkKeys) SetItems(v []SdkKeyListItem)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *SdkKeys) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SdkKeys) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SdkKeys) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


