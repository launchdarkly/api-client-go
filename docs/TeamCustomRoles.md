# TeamCustomRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | The number of custom roles assigned to this team | [optional] 
**Items** | Pointer to [**[]TeamCustomRole**](TeamCustomRole.md) | An array of the custom roles that have been assigned to this team | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) | The location and content type of related resources | [optional] 

## Methods

### NewTeamCustomRoles

`func NewTeamCustomRoles() *TeamCustomRoles`

NewTeamCustomRoles instantiates a new TeamCustomRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamCustomRolesWithDefaults

`func NewTeamCustomRolesWithDefaults() *TeamCustomRoles`

NewTeamCustomRolesWithDefaults instantiates a new TeamCustomRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *TeamCustomRoles) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TeamCustomRoles) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TeamCustomRoles) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TeamCustomRoles) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *TeamCustomRoles) GetItems() []TeamCustomRole`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TeamCustomRoles) GetItemsOk() (*[]TeamCustomRole, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TeamCustomRoles) SetItems(v []TeamCustomRole)`

SetItems sets Items field to given value.

### HasItems

`func (o *TeamCustomRoles) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLinks

`func (o *TeamCustomRoles) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamCustomRoles) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamCustomRoles) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamCustomRoles) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


