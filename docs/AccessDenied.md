# AccessDenied

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Reason** | [**AccessDeniedReason**](AccessDeniedReason.md) |  | 

## Methods

### NewAccessDenied

`func NewAccessDenied(action string, reason AccessDeniedReason, ) *AccessDenied`

NewAccessDenied instantiates a new AccessDenied object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessDeniedWithDefaults

`func NewAccessDeniedWithDefaults() *AccessDenied`

NewAccessDeniedWithDefaults instantiates a new AccessDenied object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AccessDenied) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccessDenied) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccessDenied) SetAction(v string)`

SetAction sets Action field to given value.


### GetReason

`func (o *AccessDenied) GetReason() AccessDeniedReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AccessDenied) GetReasonOk() (*AccessDeniedReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AccessDenied) SetReason(v AccessDeniedReason)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


