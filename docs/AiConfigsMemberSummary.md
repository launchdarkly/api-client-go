# AiConfigsMemberSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]AiConfigsLink**](AiConfigsLink.md) | The location and content type of related resources | 
**Id** | **string** | The member&#39;s ID | 
**FirstName** | Pointer to **string** | The member&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The member&#39;s last name | [optional] 
**Role** | **string** | The member&#39;s base role. If the member has no additional roles, this role will be in effect. | 
**Email** | **string** | The member&#39;s email address | 

## Methods

### NewAiConfigsMemberSummary

`func NewAiConfigsMemberSummary(links map[string]AiConfigsLink, id string, role string, email string, ) *AiConfigsMemberSummary`

NewAiConfigsMemberSummary instantiates a new AiConfigsMemberSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigsMemberSummaryWithDefaults

`func NewAiConfigsMemberSummaryWithDefaults() *AiConfigsMemberSummary`

NewAiConfigsMemberSummaryWithDefaults instantiates a new AiConfigsMemberSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AiConfigsMemberSummary) GetLinks() map[string]AiConfigsLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AiConfigsMemberSummary) GetLinksOk() (*map[string]AiConfigsLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AiConfigsMemberSummary) SetLinks(v map[string]AiConfigsLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *AiConfigsMemberSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiConfigsMemberSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiConfigsMemberSummary) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *AiConfigsMemberSummary) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AiConfigsMemberSummary) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AiConfigsMemberSummary) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AiConfigsMemberSummary) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AiConfigsMemberSummary) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AiConfigsMemberSummary) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AiConfigsMemberSummary) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AiConfigsMemberSummary) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRole

`func (o *AiConfigsMemberSummary) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AiConfigsMemberSummary) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AiConfigsMemberSummary) SetRole(v string)`

SetRole sets Role field to given value.


### GetEmail

`func (o *AiConfigsMemberSummary) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AiConfigsMemberSummary) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AiConfigsMemberSummary) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


