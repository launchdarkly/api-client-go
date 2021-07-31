# MemberSummaryRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]CoreLink**](CoreLink.md) |  | 
**Id** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 
**Email** | **string** |  | 

## Methods

### NewMemberSummaryRep

`func NewMemberSummaryRep(links map[string]CoreLink, id string, role string, email string, ) *MemberSummaryRep`

NewMemberSummaryRep instantiates a new MemberSummaryRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSummaryRepWithDefaults

`func NewMemberSummaryRepWithDefaults() *MemberSummaryRep`

NewMemberSummaryRepWithDefaults instantiates a new MemberSummaryRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *MemberSummaryRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MemberSummaryRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MemberSummaryRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *MemberSummaryRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberSummaryRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberSummaryRep) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *MemberSummaryRep) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MemberSummaryRep) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MemberSummaryRep) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MemberSummaryRep) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *MemberSummaryRep) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MemberSummaryRep) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MemberSummaryRep) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MemberSummaryRep) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *MemberSummaryRep) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberSummaryRep) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberSummaryRep) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *MemberSummaryRep) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberSummaryRep) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberSummaryRep) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


