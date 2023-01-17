# ValuePut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Setting** | Pointer to **interface{}** | The variation value to set for the user. Must match the flag&#39;s variation type. | [optional] 
**Comment** | Pointer to **string** | Optional comment describing the change | [optional] 

## Methods

### NewValuePut

`func NewValuePut() *ValuePut`

NewValuePut instantiates a new ValuePut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuePutWithDefaults

`func NewValuePutWithDefaults() *ValuePut`

NewValuePutWithDefaults instantiates a new ValuePut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetting

`func (o *ValuePut) GetSetting() interface{}`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *ValuePut) GetSettingOk() (*interface{}, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *ValuePut) SetSetting(v interface{})`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *ValuePut) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSettingNil

`func (o *ValuePut) SetSettingNil(b bool)`

 SetSettingNil sets the value for Setting to be an explicit nil

### UnsetSetting
`func (o *ValuePut) UnsetSetting()`

UnsetSetting ensures that no value is present for Setting, not even an explicit nil
### GetComment

`func (o *ValuePut) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ValuePut) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ValuePut) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ValuePut) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


