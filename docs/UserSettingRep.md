# UserSettingRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**Setting** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewUserSettingRep

`func NewUserSettingRep() *UserSettingRep`

NewUserSettingRep instantiates a new UserSettingRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingRepWithDefaults

`func NewUserSettingRepWithDefaults() *UserSettingRep`

NewUserSettingRepWithDefaults instantiates a new UserSettingRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UserSettingRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSettingRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSettingRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserSettingRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetValue

`func (o *UserSettingRep) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserSettingRep) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserSettingRep) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UserSettingRep) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UserSettingRep) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UserSettingRep) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSetting

`func (o *UserSettingRep) GetSetting() interface{}`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *UserSettingRep) GetSettingOk() (*interface{}, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *UserSettingRep) SetSetting(v interface{})`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *UserSettingRep) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *UserSettingRep) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *UserSettingRep) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


