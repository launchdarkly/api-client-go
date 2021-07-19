# AccessRepDenied

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to [**AccessDeniedReasonRep**](AccessDeniedReasonRep.md) |  | [optional] 

## Methods

### NewAccessRepDenied

`func NewAccessRepDenied() *AccessRepDenied`

NewAccessRepDenied instantiates a new AccessRepDenied object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRepDeniedWithDefaults

`func NewAccessRepDeniedWithDefaults() *AccessRepDenied`

NewAccessRepDeniedWithDefaults instantiates a new AccessRepDenied object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AccessRepDenied) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccessRepDenied) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccessRepDenied) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AccessRepDenied) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetReason

`func (o *AccessRepDenied) GetReason() AccessDeniedReasonRep`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccessRepDenied) GetReasonOk() (*AccessDeniedReasonRep, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccessRepDenied) SetReason(v AccessDeniedReasonRep)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AccessRepDenied) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


