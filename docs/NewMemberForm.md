# NewMemberForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The member&#39;s email | 
**Password** | Pointer to **string** | The member&#39;s password | [optional] 
**FirstName** | Pointer to **string** | The member&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The member&#39;s last name | [optional] 
**Role** | Pointer to **string** | The member&#39;s built-in role | [optional] 
**CustomRoles** | Pointer to **[]string** | An array of the member&#39;s custom roles | [optional] 
**TeamKeys** | Pointer to **[]string** | An array of the member&#39;s teams | [optional] 
**RoleAttributes** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewNewMemberForm

`func NewNewMemberForm(email string, ) *NewMemberForm`

NewNewMemberForm instantiates a new NewMemberForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMemberFormWithDefaults

`func NewNewMemberFormWithDefaults() *NewMemberForm`

NewNewMemberFormWithDefaults instantiates a new NewMemberForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *NewMemberForm) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewMemberForm) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewMemberForm) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *NewMemberForm) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NewMemberForm) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NewMemberForm) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NewMemberForm) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFirstName

`func (o *NewMemberForm) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NewMemberForm) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NewMemberForm) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *NewMemberForm) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *NewMemberForm) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *NewMemberForm) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *NewMemberForm) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *NewMemberForm) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *NewMemberForm) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NewMemberForm) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NewMemberForm) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NewMemberForm) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCustomRoles

`func (o *NewMemberForm) GetCustomRoles() []string`

GetCustomRoles returns the CustomRoles field if non-nil, zero value otherwise.

### GetCustomRolesOk

`func (o *NewMemberForm) GetCustomRolesOk() (*[]string, bool)`

GetCustomRolesOk returns a tuple with the CustomRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoles

`func (o *NewMemberForm) SetCustomRoles(v []string)`

SetCustomRoles sets CustomRoles field to given value.

### HasCustomRoles

`func (o *NewMemberForm) HasCustomRoles() bool`

HasCustomRoles returns a boolean if a field has been set.

### GetTeamKeys

`func (o *NewMemberForm) GetTeamKeys() []string`

GetTeamKeys returns the TeamKeys field if non-nil, zero value otherwise.

### GetTeamKeysOk

`func (o *NewMemberForm) GetTeamKeysOk() (*[]string, bool)`

GetTeamKeysOk returns a tuple with the TeamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamKeys

`func (o *NewMemberForm) SetTeamKeys(v []string)`

SetTeamKeys sets TeamKeys field to given value.

### HasTeamKeys

`func (o *NewMemberForm) HasTeamKeys() bool`

HasTeamKeys returns a boolean if a field has been set.

### GetRoleAttributes

`func (o *NewMemberForm) GetRoleAttributes() map[string][]string`

GetRoleAttributes returns the RoleAttributes field if non-nil, zero value otherwise.

### GetRoleAttributesOk

`func (o *NewMemberForm) GetRoleAttributesOk() (*map[string][]string, bool)`

GetRoleAttributesOk returns a tuple with the RoleAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributes

`func (o *NewMemberForm) SetRoleAttributes(v map[string][]string)`

SetRoleAttributes sets RoleAttributes field to given value.

### HasRoleAttributes

`func (o *NewMemberForm) HasRoleAttributes() bool`

HasRoleAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


