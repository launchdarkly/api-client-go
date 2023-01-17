# BulkEditTeamsRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberIDs** | Pointer to **[]string** | A list of member IDs of the members who were added to the teams. | [optional] 
**TeamKeys** | Pointer to **[]string** | A list of team keys of the teams that were successfully updated. | [optional] 
**Errors** | Pointer to **[]map[string]string** | A list of team keys and errors for the teams whose updates failed. | [optional] 

## Methods

### NewBulkEditTeamsRep

`func NewBulkEditTeamsRep() *BulkEditTeamsRep`

NewBulkEditTeamsRep instantiates a new BulkEditTeamsRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditTeamsRepWithDefaults

`func NewBulkEditTeamsRepWithDefaults() *BulkEditTeamsRep`

NewBulkEditTeamsRepWithDefaults instantiates a new BulkEditTeamsRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberIDs

`func (o *BulkEditTeamsRep) GetMemberIDs() []string`

GetMemberIDs returns the MemberIDs field if non-nil, zero value otherwise.

### GetMemberIDsOk

`func (o *BulkEditTeamsRep) GetMemberIDsOk() (*[]string, bool)`

GetMemberIDsOk returns a tuple with the MemberIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIDs

`func (o *BulkEditTeamsRep) SetMemberIDs(v []string)`

SetMemberIDs sets MemberIDs field to given value.

### HasMemberIDs

`func (o *BulkEditTeamsRep) HasMemberIDs() bool`

HasMemberIDs returns a boolean if a field has been set.

### GetTeamKeys

`func (o *BulkEditTeamsRep) GetTeamKeys() []string`

GetTeamKeys returns the TeamKeys field if non-nil, zero value otherwise.

### GetTeamKeysOk

`func (o *BulkEditTeamsRep) GetTeamKeysOk() (*[]string, bool)`

GetTeamKeysOk returns a tuple with the TeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamKeys

`func (o *BulkEditTeamsRep) SetTeamKeys(v []string)`

SetTeamKeys sets TeamKeys field to given value.

### HasTeamKeys

`func (o *BulkEditTeamsRep) HasTeamKeys() bool`

HasTeamKeys returns a boolean if a field has been set.

### GetErrors

`func (o *BulkEditTeamsRep) GetErrors() []map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkEditTeamsRep) GetErrorsOk() (*[]map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkEditTeamsRep) SetErrors(v []map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkEditTeamsRep) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


