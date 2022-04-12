# UserFlagSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Value** | **interface{}** |  | 
**Setting** | **interface{}** |  | 
**Reason** | Pointer to [**EvaluationReason**](EvaluationReason.md) |  | [optional] 

## Methods

### NewUserFlagSetting

`func NewUserFlagSetting(links map[string]Link, value interface{}, setting interface{}, ) *UserFlagSetting`

NewUserFlagSetting instantiates a new UserFlagSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFlagSettingWithDefaults

`func NewUserFlagSettingWithDefaults() *UserFlagSetting`

NewUserFlagSettingWithDefaults instantiates a new UserFlagSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UserFlagSetting) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFlagSetting) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFlagSetting) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetValue

`func (o *UserFlagSetting) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserFlagSetting) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserFlagSetting) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *UserFlagSetting) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UserFlagSetting) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSetting

`func (o *UserFlagSetting) GetSetting() interface{}`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *UserFlagSetting) GetSettingOk() (*interface{}, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *UserFlagSetting) SetSetting(v interface{})`

SetSetting sets Setting field to given value.


### SetSettingNil

`func (o *UserFlagSetting) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *UserFlagSetting) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetReason

`func (o *UserFlagSetting) GetReason() EvaluationReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UserFlagSetting) GetReasonOk() (*EvaluationReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UserFlagSetting) SetReason(v EvaluationReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UserFlagSetting) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


