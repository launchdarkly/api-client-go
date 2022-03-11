# PermissionGrantInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionSet** | Pointer to **string** |  | [optional] 
**Actions** | Pointer to **[]string** |  | [optional] 
**MemberIDs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPermissionGrantInput

`func NewPermissionGrantInput() *PermissionGrantInput`

NewPermissionGrantInput instantiates a new PermissionGrantInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantInputWithDefaults

`func NewPermissionGrantInputWithDefaults() *PermissionGrantInput`

NewPermissionGrantInputWithDefaults instantiates a new PermissionGrantInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionSet

`func (o *PermissionGrantInput) GetActionSet() string`

GetActionSet returns the ActionSet field if non-nil, zero value otherwise.

### GetActionSetOk

`func (o *PermissionGrantInput) GetActionSetOk() (*string, bool)`

GetActionSetOk returns a tuple with the ActionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSet

`func (o *PermissionGrantInput) SetActionSet(v string)`

SetActionSet sets ActionSet field to given value.

### HasActionSet

`func (o *PermissionGrantInput) HasActionSet() bool`

HasActionSet returns a boolean if a field has been set.

### GetActions

`func (o *PermissionGrantInput) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PermissionGrantInput) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PermissionGrantInput) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PermissionGrantInput) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetMemberIDs

`func (o *PermissionGrantInput) GetMemberIDs() []string`

GetMemberIDs returns the MemberIDs field if non-nil, zero value otherwise.

### GetMemberIDsOk

`func (o *PermissionGrantInput) GetMemberIDsOk() (*[]string, bool)`

GetMemberIDsOk returns a tuple with the MemberIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIDs

`func (o *PermissionGrantInput) SetMemberIDs(v []string)`

SetMemberIDs sets MemberIDs field to given value.

### HasMemberIDs

`func (o *PermissionGrantInput) HasMemberIDs() bool`

HasMemberIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


