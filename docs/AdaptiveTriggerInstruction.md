# AdaptiveTriggerInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**VariationId** | **string** |  | 
**RuleId** | Pointer to **string** | The targeting rule this adaptive trigger acts on. Required for updateRuleVariationOrRollout instructions. | [optional] 

## Methods

### NewAdaptiveTriggerInstruction

`func NewAdaptiveTriggerInstruction(kind string, variationId string, ) *AdaptiveTriggerInstruction`

NewAdaptiveTriggerInstruction instantiates a new AdaptiveTriggerInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdaptiveTriggerInstructionWithDefaults

`func NewAdaptiveTriggerInstructionWithDefaults() *AdaptiveTriggerInstruction`

NewAdaptiveTriggerInstructionWithDefaults instantiates a new AdaptiveTriggerInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AdaptiveTriggerInstruction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AdaptiveTriggerInstruction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AdaptiveTriggerInstruction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetVariationId

`func (o *AdaptiveTriggerInstruction) GetVariationId() string`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *AdaptiveTriggerInstruction) GetVariationIdOk() (*string, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *AdaptiveTriggerInstruction) SetVariationId(v string)`

SetVariationId sets VariationId field to given value.


### GetRuleId

`func (o *AdaptiveTriggerInstruction) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *AdaptiveTriggerInstruction) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *AdaptiveTriggerInstruction) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *AdaptiveTriggerInstruction) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


