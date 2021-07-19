# CustomRoleRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]InlineResponse200**](InlineResponse200.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**[]StatementRep**](StatementRep.md) |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 

## Methods

### NewCustomRoleRep

`func NewCustomRoleRep() *CustomRoleRep`

NewCustomRoleRep instantiates a new CustomRoleRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRoleRepWithDefaults

`func NewCustomRoleRepWithDefaults() *CustomRoleRep`

NewCustomRoleRepWithDefaults instantiates a new CustomRoleRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CustomRoleRep) GetLinks() map[string]InlineResponse200`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomRoleRep) GetLinksOk() (*map[string]InlineResponse200, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomRoleRep) SetLinks(v map[string]InlineResponse200)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomRoleRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *CustomRoleRep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomRoleRep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomRoleRep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomRoleRep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *CustomRoleRep) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CustomRoleRep) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CustomRoleRep) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CustomRoleRep) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *CustomRoleRep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomRoleRep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomRoleRep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomRoleRep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CustomRoleRep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomRoleRep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomRoleRep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomRoleRep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicy

`func (o *CustomRoleRep) GetPolicy() []StatementRep`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CustomRoleRep) GetPolicyOk() (*[]StatementRep, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CustomRoleRep) SetPolicy(v []StatementRep)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CustomRoleRep) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetAccess

`func (o *CustomRoleRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CustomRoleRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CustomRoleRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CustomRoleRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


