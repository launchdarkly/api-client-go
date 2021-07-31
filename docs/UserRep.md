# UserRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPing** | Pointer to **time.Time** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**ModelsUser**](ModelsUser.md) |  | [optional] 
**SortValue** | Pointer to **interface{}** |  | [optional] 
**Links** | Pointer to [**map[string]CoreLink**](CoreLink.md) |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 

## Methods

### NewUserRep

`func NewUserRep() *UserRep`

NewUserRep instantiates a new UserRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRepWithDefaults

`func NewUserRepWithDefaults() *UserRep`

NewUserRepWithDefaults instantiates a new UserRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPing

`func (o *UserRep) GetLastPing() time.Time`

GetLastPing returns the LastPing field if non-nil, zero value otherwise.

### GetLastPingOk

`func (o *UserRep) GetLastPingOk() (*time.Time, bool)`

GetLastPingOk returns a tuple with the LastPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPing

`func (o *UserRep) SetLastPing(v time.Time)`

SetLastPing sets LastPing field to given value.

### HasLastPing

`func (o *UserRep) HasLastPing() bool`

HasLastPing returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *UserRep) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *UserRep) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *UserRep) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *UserRep) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetOwnerId

`func (o *UserRep) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UserRep) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UserRep) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UserRep) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetUser

`func (o *UserRep) GetUser() ModelsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserRep) GetUserOk() (*ModelsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserRep) SetUser(v ModelsUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UserRep) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSortValue

`func (o *UserRep) GetSortValue() interface{}`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *UserRep) GetSortValueOk() (*interface{}, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *UserRep) SetSortValue(v interface{})`

SetSortValue sets SortValue field to given value.

### HasSortValue

`func (o *UserRep) HasSortValue() bool`

HasSortValue returns a boolean if a field has been set.

### SetSortValueNil

`func (o *UserRep) SetSortValueNil(b bool)`

 SetSortValueNil sets the value for SortValue to be an explicit nil

### UnsetSortValue
`func (o *UserRep) UnsetSortValue()`

UnsetSortValue ensures that no value is present for SortValue, not even an explicit nil
### GetLinks

`func (o *UserRep) GetLinks() map[string]CoreLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserRep) GetLinksOk() (*map[string]CoreLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserRep) SetLinks(v map[string]CoreLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserRep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccess

`func (o *UserRep) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UserRep) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UserRep) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UserRep) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


