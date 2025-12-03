# TrustPolicyStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** | The effect of trust policy statement | [optional] 
**Action** | Pointer to **[]string** | The action of trust policy statement | [optional] 
**Principal** | Pointer to **interface{}** | The principal of trust policy statement | [optional] 
**Condition** | Pointer to **interface{}** | The condition of trust policy statement | [optional] 

## Methods

### NewTrustPolicyStatement

`func NewTrustPolicyStatement() *TrustPolicyStatement`

NewTrustPolicyStatement instantiates a new TrustPolicyStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustPolicyStatementWithDefaults

`func NewTrustPolicyStatementWithDefaults() *TrustPolicyStatement`

NewTrustPolicyStatementWithDefaults instantiates a new TrustPolicyStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *TrustPolicyStatement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *TrustPolicyStatement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *TrustPolicyStatement) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *TrustPolicyStatement) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetAction

`func (o *TrustPolicyStatement) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TrustPolicyStatement) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TrustPolicyStatement) SetAction(v []string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TrustPolicyStatement) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPrincipal

`func (o *TrustPolicyStatement) GetPrincipal() interface{}`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *TrustPolicyStatement) GetPrincipalOk() (*interface{}, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *TrustPolicyStatement) SetPrincipal(v interface{})`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *TrustPolicyStatement) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *TrustPolicyStatement) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *TrustPolicyStatement) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetCondition

`func (o *TrustPolicyStatement) GetCondition() interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *TrustPolicyStatement) GetConditionOk() (*interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *TrustPolicyStatement) SetCondition(v interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *TrustPolicyStatement) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *TrustPolicyStatement) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *TrustPolicyStatement) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


