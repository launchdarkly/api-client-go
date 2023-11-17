# MaintainerInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to [**MemberInput**](MemberInput.md) |  | [optional] 
**Team** | Pointer to [**TeamInput**](TeamInput.md) |  | [optional] 

## Methods

### NewMaintainerInput

`func NewMaintainerInput() *MaintainerInput`

NewMaintainerInput instantiates a new MaintainerInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintainerInputWithDefaults

`func NewMaintainerInputWithDefaults() *MaintainerInput`

NewMaintainerInputWithDefaults instantiates a new MaintainerInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *MaintainerInput) GetMember() MemberInput`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MaintainerInput) GetMemberOk() (*MemberInput, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MaintainerInput) SetMember(v MemberInput)`

SetMember sets Member field to given value.

### HasMember

`func (o *MaintainerInput) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetTeam

`func (o *MaintainerInput) GetTeam() TeamInput`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MaintainerInput) GetTeamOk() (*TeamInput, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MaintainerInput) SetTeam(v TeamInput)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MaintainerInput) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


