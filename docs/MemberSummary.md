# MemberSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) | The location and content type of related resources | 
**Id** | **string** | The member&#39;s ID | 
**FirstName** | Pointer to **string** | The member&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The member&#39;s last name | [optional] 
**Role** | **string** | The member&#39;s base role. If the member has no additional roles, this role will be in effect. | 
**Email** | **string** | The member&#39;s email address | 

## Methods

### NewMemberSummary

`func NewMemberSummary(links map[string]Link, id string, role string, email string, ) *MemberSummary`

NewMemberSummary instantiates a new MemberSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSummaryWithDefaults

`func NewMemberSummaryWithDefaults() *MemberSummary`

NewMemberSummaryWithDefaults instantiates a new MemberSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *MemberSummary) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MemberSummary) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MemberSummary) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *MemberSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberSummary) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *MemberSummary) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MemberSummary) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MemberSummary) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MemberSummary) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *MemberSummary) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MemberSummary) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MemberSummary) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MemberSummary) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *MemberSummary) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberSummary) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberSummary) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *MemberSummary) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberSummary) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberSummary) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


