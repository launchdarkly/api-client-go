# AiConfigsAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Denied** | [**[]AiConfigsAccessDenied**](AiConfigsAccessDenied.md) |  | 
**Allowed** | [**[]AiConfigsAccessAllowedRep**](AiConfigsAccessAllowedRep.md) |  | 

## Methods

### NewAiConfigsAccess

`func NewAiConfigsAccess(denied []AiConfigsAccessDenied, allowed []AiConfigsAccessAllowedRep, ) *AiConfigsAccess`

NewAiConfigsAccess instantiates a new AiConfigsAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiConfigsAccessWithDefaults

`func NewAiConfigsAccessWithDefaults() *AiConfigsAccess`

NewAiConfigsAccessWithDefaults instantiates a new AiConfigsAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenied

`func (o *AiConfigsAccess) GetDenied() []AiConfigsAccessDenied`

GetDenied returns the Denied field if non-nil, zero value otherwise.

### GetDeniedOk

`func (o *AiConfigsAccess) GetDeniedOk() (*[]AiConfigsAccessDenied, bool)`

GetDeniedOk returns a tuple with the Denied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenied

`func (o *AiConfigsAccess) SetDenied(v []AiConfigsAccessDenied)`

SetDenied sets Denied field to given value.


### GetAllowed

`func (o *AiConfigsAccess) GetAllowed() []AiConfigsAccessAllowedRep`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *AiConfigsAccess) GetAllowedOk() (*[]AiConfigsAccessAllowedRep, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *AiConfigsAccess) SetAllowed(v []AiConfigsAccessAllowedRep)`

SetAllowed sets Allowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


