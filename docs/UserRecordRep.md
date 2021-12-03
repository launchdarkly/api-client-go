# UserRecordRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPing** | Pointer to **time.Time** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**SortValue** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewUserRecordRep

`func NewUserRecordRep() *UserRecordRep`

NewUserRecordRep instantiates a new UserRecordRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRecordRepWithDefaults

`func NewUserRecordRepWithDefaults() *UserRecordRep`

NewUserRecordRepWithDefaults instantiates a new UserRecordRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPing

`func (o *UserRecordRep) GetLastPing() time.Time`

GetLastPing returns the LastPing field if non-nil, zero value otherwise.

### GetLastPingOk

`func (o *UserRecordRep) GetLastPingOk() (*time.Time, bool)`

GetLastPingOk returns a tuple with the LastPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPing

`func (o *UserRecordRep) SetLastPing(v time.Time)`

SetLastPing sets LastPing field to given value.

### HasLastPing

`func (o *UserRecordRep) HasLastPing() bool`

HasLastPing returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *UserRecordRep) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *UserRecordRep) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *UserRecordRep) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *UserRecordRep) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetOwnerId

`func (o *UserRecordRep) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UserRecordRep) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UserRecordRep) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UserRecordRep) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetUser

`func (o *UserRecordRep) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserRecordRep) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserRecordRep) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UserRecordRep) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSortValue

`func (o *UserRecordRep) GetSortValue() interface{}`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *UserRecordRep) GetSortValueOk() (*interface{}, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *UserRecordRep) SetSortValue(v interface{})`

SetSortValue sets SortValue field to given value.

### HasSortValue

`func (o *UserRecordRep) HasSortValue() bool`

HasSortValue returns a boolean if a field has been set.

### SetSortValueNil

`func (o *UserRecordRep) SetSortValueNil(b bool)`

 SetSortValueNil sets the value for SortValue to be an explicit nil

### UnsetSortValue
`func (o *UserRecordRep) UnsetSortValue()`

UnsetSortValue ensures that no value is present for SortValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


