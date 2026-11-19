# ToolSkillReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillKey** | **string** | The key of the agent skill. | 
**SkillVersion** | **int32** | The version of the agent skill that pins the tool. | 
**ToolVersion** | **int32** | The version of the tool the skill version pins. | 

## Methods

### NewToolSkillReference

`func NewToolSkillReference(skillKey string, skillVersion int32, toolVersion int32, ) *ToolSkillReference`

NewToolSkillReference instantiates a new ToolSkillReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolSkillReferenceWithDefaults

`func NewToolSkillReferenceWithDefaults() *ToolSkillReference`

NewToolSkillReferenceWithDefaults instantiates a new ToolSkillReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillKey

`func (o *ToolSkillReference) GetSkillKey() string`

GetSkillKey returns the SkillKey field if non-nil, zero value otherwise.

### GetSkillKeyOk

`func (o *ToolSkillReference) GetSkillKeyOk() (*string, bool)`

GetSkillKeyOk returns a tuple with the SkillKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillKey

`func (o *ToolSkillReference) SetSkillKey(v string)`

SetSkillKey sets SkillKey field to given value.


### GetSkillVersion

`func (o *ToolSkillReference) GetSkillVersion() int32`

GetSkillVersion returns the SkillVersion field if non-nil, zero value otherwise.

### GetSkillVersionOk

`func (o *ToolSkillReference) GetSkillVersionOk() (*int32, bool)`

GetSkillVersionOk returns a tuple with the SkillVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillVersion

`func (o *ToolSkillReference) SetSkillVersion(v int32)`

SetSkillVersion sets SkillVersion field to given value.


### GetToolVersion

`func (o *ToolSkillReference) GetToolVersion() int32`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *ToolSkillReference) GetToolVersionOk() (*int32, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *ToolSkillReference) SetToolVersion(v int32)`

SetToolVersion sets ToolVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


