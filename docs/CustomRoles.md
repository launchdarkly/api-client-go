# CustomRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Items** | Pointer to [**[]CustomRole**](CustomRole.md) |  | [optional] 

## Methods

### NewCustomRoles

`func NewCustomRoles() *CustomRoles`

NewCustomRoles instantiates a new CustomRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRolesWithDefaults

`func NewCustomRolesWithDefaults() *CustomRoles`

NewCustomRolesWithDefaults instantiates a new CustomRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CustomRoles) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomRoles) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomRoles) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomRoles) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *CustomRoles) GetItems() []CustomRole`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CustomRoles) GetItemsOk() (*[]CustomRole, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CustomRoles) SetItems(v []CustomRole)`

SetItems sets Items field to given value.

### HasItems

`func (o *CustomRoles) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


