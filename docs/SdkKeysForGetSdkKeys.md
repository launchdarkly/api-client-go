# SdkKeysForGetSdkKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**SdkKeysSelfLink**](SdkKeysSelfLink.md) |  | [optional] 
**Items** | [**[]SdkKey**](SdkKey.md) |  | 
**TotalCount** | **int32** | The total number of SDK keys matching the query, before pagination. | 

## Methods

### NewSdkKeysForGetSdkKeys

`func NewSdkKeysForGetSdkKeys(items []SdkKey, totalCount int32, ) *SdkKeysForGetSdkKeys`

NewSdkKeysForGetSdkKeys instantiates a new SdkKeysForGetSdkKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkKeysForGetSdkKeysWithDefaults

`func NewSdkKeysForGetSdkKeysWithDefaults() *SdkKeysForGetSdkKeys`

NewSdkKeysForGetSdkKeysWithDefaults instantiates a new SdkKeysForGetSdkKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SdkKeysForGetSdkKeys) GetLinks() SdkKeysSelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SdkKeysForGetSdkKeys) GetLinksOk() (*SdkKeysSelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SdkKeysForGetSdkKeys) SetLinks(v SdkKeysSelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SdkKeysForGetSdkKeys) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *SdkKeysForGetSdkKeys) GetItems() []SdkKey`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SdkKeysForGetSdkKeys) GetItemsOk() (*[]SdkKey, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SdkKeysForGetSdkKeys) SetItems(v []SdkKey)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *SdkKeysForGetSdkKeys) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SdkKeysForGetSdkKeys) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SdkKeysForGetSdkKeys) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


