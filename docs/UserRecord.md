# UserRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastPing** | Pointer to **time.Time** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**SortValue** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewUserRecord

`func NewUserRecord() *UserRecord`

NewUserRecord instantiates a new UserRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRecordWithDefaults

`func NewUserRecordWithDefaults() *UserRecord`

NewUserRecordWithDefaults instantiates a new UserRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPing

`func (o *UserRecord) GetLastPing() time.Time`

GetLastPing returns the LastPing field if non-nil, zero value otherwise.

### GetLastPingOk

`func (o *UserRecord) GetLastPingOk() (*time.Time, bool)`

GetLastPingOk returns a tuple with the LastPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPing

`func (o *UserRecord) SetLastPing(v time.Time)`

SetLastPing sets LastPing field to given value.

### HasLastPing

`func (o *UserRecord) HasLastPing() bool`

HasLastPing returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *UserRecord) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *UserRecord) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *UserRecord) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *UserRecord) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetOwnerId

`func (o *UserRecord) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UserRecord) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UserRecord) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UserRecord) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetUser

`func (o *UserRecord) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserRecord) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserRecord) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UserRecord) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSortValue

`func (o *UserRecord) GetSortValue() interface{}`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *UserRecord) GetSortValueOk() (*interface{}, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *UserRecord) SetSortValue(v interface{})`

SetSortValue sets SortValue field to given value.

### HasSortValue

`func (o *UserRecord) HasSortValue() bool`

HasSortValue returns a boolean if a field has been set.

### SetSortValueNil

`func (o *UserRecord) SetSortValueNil(b bool)`

 SetSortValueNil sets the value for SortValue to be an explicit nil

### UnsetSortValue
`func (o *UserRecord) UnsetSortValue()`

UnsetSortValue ensures that no value is present for SortValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


