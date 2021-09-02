# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPing** | Pointer to **time.Time** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**SortValue** | Pointer to **interface{}** |  | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Access** | Pointer to [**AccessRep**](AccessRep.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPing

`func (o *User) GetLastPing() time.Time`

GetLastPing returns the LastPing field if non-nil, zero value otherwise.

### GetLastPingOk

`func (o *User) GetLastPingOk() (*time.Time, bool)`

GetLastPingOk returns a tuple with the LastPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPing

`func (o *User) SetLastPing(v time.Time)`

SetLastPing sets LastPing field to given value.

### HasLastPing

`func (o *User) HasLastPing() bool`

HasLastPing returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *User) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *User) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *User) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *User) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetOwnerId

`func (o *User) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *User) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *User) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *User) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetUser

`func (o *User) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *User) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *User) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *User) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSortValue

`func (o *User) GetSortValue() interface{}`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *User) GetSortValueOk() (*interface{}, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *User) SetSortValue(v interface{})`

SetSortValue sets SortValue field to given value.

### HasSortValue

`func (o *User) HasSortValue() bool`

HasSortValue returns a boolean if a field has been set.

### SetSortValueNil

`func (o *User) SetSortValueNil(b bool)`

 SetSortValueNil sets the value for SortValue to be an explicit nil

### UnsetSortValue
`func (o *User) UnsetSortValue()`

UnsetSortValue ensures that no value is present for SortValue, not even an explicit nil
### GetLinks

`func (o *User) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *User) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *User) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *User) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccess

`func (o *User) GetAccess() AccessRep`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *User) GetAccessOk() (*AccessRep, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *User) SetAccess(v AccessRep)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *User) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


