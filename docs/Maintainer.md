# Maintainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**MaintainerMember** | Pointer to [**ViewsMaintainerMember**](ViewsMaintainerMember.md) |  | [optional] 
**MaintainerTeam** | Pointer to [**ViewsMaintainerTeam**](ViewsMaintainerTeam.md) |  | [optional] 

## Methods

### NewMaintainer

`func NewMaintainer(id string, kind string, ) *Maintainer`

NewMaintainer instantiates a new Maintainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintainerWithDefaults

`func NewMaintainerWithDefaults() *Maintainer`

NewMaintainerWithDefaults instantiates a new Maintainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Maintainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Maintainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Maintainer) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *Maintainer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Maintainer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Maintainer) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMaintainerMember

`func (o *Maintainer) GetMaintainerMember() ViewsMaintainerMember`

GetMaintainerMember returns the MaintainerMember field if non-nil, zero value otherwise.

### GetMaintainerMemberOk

`func (o *Maintainer) GetMaintainerMemberOk() (*ViewsMaintainerMember, bool)`

GetMaintainerMemberOk returns a tuple with the MaintainerMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerMember

`func (o *Maintainer) SetMaintainerMember(v ViewsMaintainerMember)`

SetMaintainerMember sets MaintainerMember field to given value.

### HasMaintainerMember

`func (o *Maintainer) HasMaintainerMember() bool`

HasMaintainerMember returns a boolean if a field has been set.

### GetMaintainerTeam

`func (o *Maintainer) GetMaintainerTeam() ViewsMaintainerTeam`

GetMaintainerTeam returns the MaintainerTeam field if non-nil, zero value otherwise.

### GetMaintainerTeamOk

`func (o *Maintainer) GetMaintainerTeamOk() (*ViewsMaintainerTeam, bool)`

GetMaintainerTeamOk returns a tuple with the MaintainerTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerTeam

`func (o *Maintainer) SetMaintainerTeam(v ViewsMaintainerTeam)`

SetMaintainerTeam sets MaintainerTeam field to given value.

### HasMaintainerTeam

`func (o *Maintainer) HasMaintainerTeam() bool`

HasMaintainerTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


