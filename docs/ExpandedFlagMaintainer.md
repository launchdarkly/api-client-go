# ExpandedFlagMaintainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The ID of the maintainer member, or the key of the maintainer team | 
**Kind** | **string** | The type of the maintainer | 
**Member** | Pointer to [**ViewsMemberSummary**](ViewsMemberSummary.md) |  | [optional] 
**Team** | Pointer to [**ViewsMemberTeamSummaryRep**](ViewsMemberTeamSummaryRep.md) |  | [optional] 

## Methods

### NewExpandedFlagMaintainer

`func NewExpandedFlagMaintainer(key string, kind string, ) *ExpandedFlagMaintainer`

NewExpandedFlagMaintainer instantiates a new ExpandedFlagMaintainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedFlagMaintainerWithDefaults

`func NewExpandedFlagMaintainerWithDefaults() *ExpandedFlagMaintainer`

NewExpandedFlagMaintainerWithDefaults instantiates a new ExpandedFlagMaintainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ExpandedFlagMaintainer) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ExpandedFlagMaintainer) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ExpandedFlagMaintainer) SetKey(v string)`

SetKey sets Key field to given value.


### GetKind

`func (o *ExpandedFlagMaintainer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExpandedFlagMaintainer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExpandedFlagMaintainer) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMember

`func (o *ExpandedFlagMaintainer) GetMember() ViewsMemberSummary`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ExpandedFlagMaintainer) GetMemberOk() (*ViewsMemberSummary, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ExpandedFlagMaintainer) SetMember(v ViewsMemberSummary)`

SetMember sets Member field to given value.

### HasMember

`func (o *ExpandedFlagMaintainer) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetTeam

`func (o *ExpandedFlagMaintainer) GetTeam() ViewsMemberTeamSummaryRep`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ExpandedFlagMaintainer) GetTeamOk() (*ViewsMemberTeamSummaryRep, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ExpandedFlagMaintainer) SetTeam(v ViewsMemberTeamSummaryRep)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ExpandedFlagMaintainer) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


