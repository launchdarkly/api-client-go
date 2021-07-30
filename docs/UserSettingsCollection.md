# UserSettingsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**map[string]UserSettingRep**](UserSettingRep.md) |  | [optional] 
**Links** | Pointer to [**[]CoreLink**](CoreLink.md) |  | [optional] 

## Methods

### NewUserSettingsCollection

`func NewUserSettingsCollection() *UserSettingsCollection`

NewUserSettingsCollection instantiates a new UserSettingsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsCollectionWithDefaults

`func NewUserSettingsCollectionWithDefaults() *UserSettingsCollection`

NewUserSettingsCollectionWithDefaults instantiates a new UserSettingsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserSettingsCollection) GetItems() map[string]UserSettingRep`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserSettingsCollection) GetItemsOk() (*map[string]UserSettingRep, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserSettingsCollection) SetItems(v map[string]UserSettingRep)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserSettingsCollection) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *UserSettingsCollection) GetLinks() []CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSettingsCollection) GetLinksOk() (*[]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSettingsCollection) SetLinks(v []CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserSettingsCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


