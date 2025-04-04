# AiConfigsAccessDenied

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Reason** | [**AiConfigsAccessDeniedReason**](AiConfigsAccessDeniedReason.md) |  | 

## Methods

### NewAiConfigsAccessDenied

`func NewAiConfigsAccessDenied(action string, reason AiConfigsAccessDeniedReason, ) *AiConfigsAccessDenied`

NewAiConfigsAccessDenied instantiates a new AiConfigsAccessDenied object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigsAccessDeniedWithDefaults

`func NewAiConfigsAccessDeniedWithDefaults() *AiConfigsAccessDenied`

NewAiConfigsAccessDeniedWithDefaults instantiates a new AiConfigsAccessDenied object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AiConfigsAccessDenied) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AiConfigsAccessDenied) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AiConfigsAccessDenied) SetAction(v string)`

SetAction sets Action field to given value.


### GetReason

`func (o *AiConfigsAccessDenied) GetReason() AiConfigsAccessDeniedReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AiConfigsAccessDenied) GetReasonOk() (*AiConfigsAccessDeniedReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AiConfigsAccessDenied) SetReason(v AiConfigsAccessDeniedReason)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


