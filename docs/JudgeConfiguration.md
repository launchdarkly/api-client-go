# JudgeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Judges** | Pointer to [**[]JudgeAttachment**](JudgeAttachment.md) | List of judges for this variation. When updating, this replaces all existing judge attachments, and if empty, removes all judge attachments.  | [optional] 

## Methods

### NewJudgeConfiguration

`func NewJudgeConfiguration() *JudgeConfiguration`

NewJudgeConfiguration instantiates a new JudgeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeConfigurationWithDefaults

`func NewJudgeConfigurationWithDefaults() *JudgeConfiguration`

NewJudgeConfigurationWithDefaults instantiates a new JudgeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJudges

`func (o *JudgeConfiguration) GetJudges() []JudgeAttachment`

GetJudges returns the Judges field if non-nil, zero value otherwise.

### GetJudgesOk

`func (o *JudgeConfiguration) GetJudgesOk() (*[]JudgeAttachment, bool)`

GetJudgesOk returns a tuple with the Judges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudges

`func (o *JudgeConfiguration) SetJudges(v []JudgeAttachment)`

SetJudges sets Judges field to given value.

### HasJudges

`func (o *JudgeConfiguration) HasJudges() bool`

HasJudges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


