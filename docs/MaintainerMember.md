# MaintainerMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 

## Methods

### NewMaintainerMember

`func NewMaintainerMember(id string, email string, role string, ) *MaintainerMember`

NewMaintainerMember instantiates a new MaintainerMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintainerMemberWithDefaults

`func NewMaintainerMemberWithDefaults() *MaintainerMember`

NewMaintainerMemberWithDefaults instantiates a new MaintainerMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintainerMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintainerMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintainerMember) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *MaintainerMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MaintainerMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MaintainerMember) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *MaintainerMember) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MaintainerMember) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MaintainerMember) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MaintainerMember) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *MaintainerMember) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MaintainerMember) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MaintainerMember) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MaintainerMember) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *MaintainerMember) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MaintainerMember) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MaintainerMember) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


