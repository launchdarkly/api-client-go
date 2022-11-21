# BulkEditMembersRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]string** | A list of members IDs of the members who were successfully updated. | [optional] 
**Errors** | Pointer to **[]map[string]string** | A list of member IDs and errors for the members whose updates failed. | [optional] 

## Methods

### NewBulkEditMembersRep

`func NewBulkEditMembersRep() *BulkEditMembersRep`

NewBulkEditMembersRep instantiates a new BulkEditMembersRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditMembersRepWithDefaults

`func NewBulkEditMembersRepWithDefaults() *BulkEditMembersRep`

NewBulkEditMembersRepWithDefaults instantiates a new BulkEditMembersRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *BulkEditMembersRep) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *BulkEditMembersRep) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *BulkEditMembersRep) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *BulkEditMembersRep) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetErrors

`func (o *BulkEditMembersRep) GetErrors() []map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkEditMembersRep) GetErrorsOk() (*[]map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkEditMembersRep) SetErrors(v []map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkEditMembersRep) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


