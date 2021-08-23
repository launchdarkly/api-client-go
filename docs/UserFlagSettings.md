# UserFlagSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**map[string]UserFlagSetting**](UserFlagSetting.md) |  | 
**Links** | [**map[string]Link**](Link.md) |  | 

## Methods

### NewUserFlagSettings

`func NewUserFlagSettings(items map[string]UserFlagSetting, links map[string]Link, ) *UserFlagSettings`

NewUserFlagSettings instantiates a new UserFlagSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFlagSettingsWithDefaults

`func NewUserFlagSettingsWithDefaults() *UserFlagSettings`

NewUserFlagSettingsWithDefaults instantiates a new UserFlagSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserFlagSettings) GetItems() map[string]UserFlagSetting`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserFlagSettings) GetItemsOk() (*map[string]UserFlagSetting, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserFlagSettings) SetItems(v map[string]UserFlagSetting)`

SetItems sets Items field to given value.


### GetLinks

`func (o *UserFlagSettings) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFlagSettings) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFlagSettings) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


