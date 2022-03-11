# FlagLinkMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]Link**](Link.md) |  | 
**Id** | **string** |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 

## Methods

### NewFlagLinkMember

`func NewFlagLinkMember(links map[string]Link, id string, ) *FlagLinkMember`

NewFlagLinkMember instantiates a new FlagLinkMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagLinkMemberWithDefaults

`func NewFlagLinkMemberWithDefaults() *FlagLinkMember`

NewFlagLinkMemberWithDefaults instantiates a new FlagLinkMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlagLinkMember) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlagLinkMember) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlagLinkMember) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetId

`func (o *FlagLinkMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagLinkMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagLinkMember) SetId(v string)`

SetId sets Id field to given value.


### GetFirstName

`func (o *FlagLinkMember) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *FlagLinkMember) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *FlagLinkMember) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *FlagLinkMember) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *FlagLinkMember) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *FlagLinkMember) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *FlagLinkMember) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *FlagLinkMember) HasLastName() bool`

HasLastName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


