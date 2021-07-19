# AccessDeniedRep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to [**AccessDeniedReasonRep**](AccessDeniedReasonRep.md) |  | [optional] 

## Methods

### NewAccessDeniedRep

`func NewAccessDeniedRep() *AccessDeniedRep`

NewAccessDeniedRep instantiates a new AccessDeniedRep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDeniedRepWithDefaults

`func NewAccessDeniedRepWithDefaults() *AccessDeniedRep`

NewAccessDeniedRepWithDefaults instantiates a new AccessDeniedRep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AccessDeniedRep) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccessDeniedRep) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccessDeniedRep) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AccessDeniedRep) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetReason

`func (o *AccessDeniedRep) GetReason() AccessDeniedReasonRep`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccessDeniedRep) GetReasonOk() (*AccessDeniedReasonRep, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccessDeniedRep) SetReason(v AccessDeniedReasonRep)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AccessDeniedRep) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


