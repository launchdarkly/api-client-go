# ApplicationMaintainerRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to [**MemberSummary**](MemberSummary.md) |  | [optional] 
**Team** | Pointer to [**MemberTeamSummaryRep**](MemberTeamSummaryRep.md) |  | [optional] 

## Methods

### NewApplicationMaintainerRep

`func NewApplicationMaintainerRep() *ApplicationMaintainerRep`

NewApplicationMaintainerRep instantiates a new ApplicationMaintainerRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationMaintainerRepWithDefaults

`func NewApplicationMaintainerRepWithDefaults() *ApplicationMaintainerRep`

NewApplicationMaintainerRepWithDefaults instantiates a new ApplicationMaintainerRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *ApplicationMaintainerRep) GetMember() MemberSummary`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ApplicationMaintainerRep) GetMemberOk() (*MemberSummary, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ApplicationMaintainerRep) SetMember(v MemberSummary)`

SetMember sets Member field to given value.

### HasMember

`func (o *ApplicationMaintainerRep) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetTeam

`func (o *ApplicationMaintainerRep) GetTeam() MemberTeamSummaryRep`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ApplicationMaintainerRep) GetTeamOk() (*MemberTeamSummaryRep, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ApplicationMaintainerRep) SetTeam(v MemberTeamSummaryRep)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ApplicationMaintainerRep) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


