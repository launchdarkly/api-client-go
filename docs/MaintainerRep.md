# MaintainerRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The ID of the maintainer member, or the key of the maintainer team | 
**Kind** | **string** | The type of the maintainer | 
**Member** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**Team** | Pointer to [**MemberTeamSummaryRep**](MemberTeamSummaryRep.md) |  | [optional] 

## Methods

### NewMaintainerRep

`func NewMaintainerRep(key string, kind string, ) *MaintainerRep`

NewMaintainerRep instantiates a new MaintainerRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintainerRepWithDefaults

`func NewMaintainerRepWithDefaults() *MaintainerRep`

NewMaintainerRepWithDefaults instantiates a new MaintainerRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MaintainerRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MaintainerRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MaintainerRep) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *MaintainerRep) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MaintainerRep) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MaintainerRep) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMember

`func (o *MaintainerRep) GetMember() MemberSummary`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MaintainerRep) GetMemberOk() (*MemberSummary, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MaintainerRep) SetMember(v MemberSummary)`

SetMember sets Member field to given value.

### HasMember

`func (o *MaintainerRep) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetTeam

`func (o *MaintainerRep) GetTeam() MemberTeamSummaryRep`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MaintainerRep) GetTeamOk() (*MemberTeamSummaryRep, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MaintainerRep) SetTeam(v MemberTeamSummaryRep)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MaintainerRep) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


